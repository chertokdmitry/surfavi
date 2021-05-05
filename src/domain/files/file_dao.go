package files

import (
	"gitlab.com/chertokdmitry/surfavi/src/message"
	"gitlab.com/chertokdmitry/surfavi/src/utils/logger"
	"os"
	"path/filepath"
	"time"
)

func GetFileListAvi() []string {
	var files []string
	root := message.AviDir

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		logger.Error(message.ErrFileWalk, err)
	}

	return files
}

func GetFileListCams(startTime time.Time, id string) []string {
	var files []string
	root := message.CamsDir + id

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.ModTime().After(startTime) && !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		logger.Error(message.ErrFileWalk, err)
	}

	return files
}
