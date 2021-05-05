package main

import (
	"gitlab.com/chertokdmitry/surfavi/src/app"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	app.Run()
}
