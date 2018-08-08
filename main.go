package main

import (
	"os"

	"gocv.io/x/gocv"
)

var template gocv.Mat = gocv.IMRead("template-intersection.png", gocv.IMWritePngStrategyFiltered)
var red_line gocv.Mat = gocv.IMRead("template-intersection-lines.png", gocv.IMWritePngStrategyFiltered)

func main() {
	_, err := gocv.VideoCaptureFile(os.Args[1])
	if err != nil {
		panic(err)
	}
}
