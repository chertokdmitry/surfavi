package message

const (
	AviDir           string = "/var/www/surfweather.ru/webcams/avi"
	CamsDir          string = "/var/www/surfweather.ru/webcams/"
	ErrReadAll       string = "error read all"
	ErrHttpGet       string = "error when make http get request"
	ErrReadFile      string = "error when read file"
	ErrMjpegNew      string = "error when create new mjpeg"
	ErrMjpegConvert  string = "error when convert mjpeg to mov"
	ErrMjpegAddFrame string = "error when add new frame"
	ErrOpenFile      string = "error when open file"
	ErrDimFile       string = "error when get file dimension"
	ErrFileWalk      string = "error when scan file directory"
	ErrConvertFile   string = "error when convert file"
)
