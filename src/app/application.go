package app

import (
	"gitlab.com/chertokdmitry/surfavi/src/domain/cameras"
	"gitlab.com/chertokdmitry/surfavi/src/domain/files"
	"gitlab.com/chertokdmitry/surfavi/src/domain/movies"
)

// run the app
func Run() {
	cameras := cameras.GetAll()

	for _, camera := range cameras {
		movies.MakeAvi(camera)
	}

	movies.Convert(files.GetFileListAvi())
}
