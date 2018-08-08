package main

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"gocv.io/x/gocv"
)

var template = gocv.IMRead("template-intersection.png", gocv.IMWritePngStrategyFiltered)
var redLines = gocv.IMRead("template-intersection-lines.png", gocv.IMWritePngStrategyFiltered)

func main() {
	capture, err := gocv.VideoCaptureFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	originalWindow := gocv.NewWindow("Original")

	regionSet := false
	var regionRect image.Rectangle
	var rectangleColor = color.RGBA{255, 255, 255, 0}

	cancelled := false
	for !cancelled {
		frame := gocv.NewMat()
		frameOk := capture.Read(&frame)
		if !frameOk {
			break
		}

		if !regionSet {
			regionResult := gocv.NewMat()
			templateDims := template.Size()
			gocv.MatchTemplate(frame, template, &regionResult, gocv.TmCcoeff, gocv.NewMat())
			_, _, _, maxLoc := gocv.MinMaxLoc(regionResult)
			fmt.Printf("Region found: [%d, %d]\n", maxLoc.X, maxLoc.Y)
			regionRect.Min = maxLoc
			regionRect.Max = image.Point{X: maxLoc.X + templateDims[0], Y: maxLoc.Y + templateDims[1]}
		}

		gocv.Rectangle(&frame, regionRect, rectangleColor, 2)

		originalWindow.IMShow(frame)
		if originalWindow.WaitKey(200)&0xFF == 'r' {
			rectangleColor = color.RGBA{255, 0, 0, 0}
			regionSet = true
		}
	}
}
