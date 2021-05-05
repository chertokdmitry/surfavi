package movies

import (
	"gitlab.com/chertokdmitry/surfavi/src/domain/files"
	"gitlab.com/chertokdmitry/surfavi/src/domain/mjpeg"
	"gitlab.com/chertokdmitry/surfavi/src/message"
	"gitlab.com/chertokdmitry/surfavi/src/utils/logger"
	"image"
	_ "image/jpeg"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Job struct {
	filename string
	results  chan<- Result
}

type Result struct {
	filename string
}

func MakeAvi(id int64) {
	idStr := strconv.FormatInt(id, 10)
	t := time.Now()
	startTime := t.Add(-time.Hour * 24)
	filesList := files.GetFileListCams(startTime, idStr)
	width, height := getDimensions(filesList[0])

	aw, err := mjpeg.New(message.AviDir+idStr+".avi", width, height, 12)

	if err != nil {
		logger.Error(message.ErrMjpegNew, err)
	}

	for _, f := range filesList {
		data, err := ioutil.ReadFile(f)
		if err != nil {
			logger.Error(message.ErrReadFile, err)
		}
		err = aw.AddFrame(data)

		if err != nil {
			logger.Error(message.ErrMjpegAddFrame, err)
		}
	}

	err = aw.Close()
	if err != nil {
		logger.Error(message.ErrMjpegNew, err)
	}
}

func getDimensions(imagePath string) (int32, int32) {
	file, err := os.Open(imagePath)
	if err != nil {
		logger.Error(message.ErrOpenFile, err)
	}
	defer file.Close()

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		logger.Error(message.ErrDimFile, err)
	}
	return int32(image.Width), int32(image.Height)
}

func Convert(filenames []string) {
	var workers = runtime.NumCPU()
	jobs := make(chan Job, workers)
	results := make(chan Result, len(filenames))
	done := make(chan struct{}, workers)

	go addJobs(jobs, filenames, results)

	for i := 0; i < workers; i++ {
		go doJobs(done, jobs)
	}

	go awaitCompletion(done, workers, results)
	processResults(results)
}

func addJobs(jobs chan<- Job, filenames []string, results chan<- Result) {
	for _, filename := range filenames {
		jobs <- Job{filename, results}
	}
	close(jobs)
}

func doJobs(done chan<- struct{}, jobs <-chan Job) {
	for job := range jobs {
		job.Do()
	}
	done <- struct{}{}
}

func awaitCompletion(done <-chan struct{}, workers int, results chan Result) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(results)
}

func processResults(results <-chan Result) {
	for range results {
		<-results
	}
}

func (job Job) Do() {
	file, err := os.Open(job.filename)
	if err != nil {
		logger.Error(message.ErrOpenFile, err)
		return
	}

	imagePath := strings.TrimSuffix(job.filename, filepath.Ext(job.filename))

	_, err = exec.Command("ffmpeg", "-y", "-i", imagePath+".avi", imagePath+".mov").Output()

	if err != nil {
		logger.Error(message.ErrMjpegConvert, err)
	}

	defer file.Close()
	job.results <- Result{job.filename}
}
