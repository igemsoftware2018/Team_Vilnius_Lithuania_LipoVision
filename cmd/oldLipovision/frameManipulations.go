package main

import (
	"image"

	"gocv.io/x/gocv"
)

//find white pixel with biggest X coordinate
func findBiggestXOfWhitePixel(cropped *gocv.Mat) (biggestX int) {

	biggestX = 0
	for i := 0; i < cropped.Cols(); i++ {
		for j := 0; j < cropped.Rows(); j++ {
			if cropped.GetUCharAt(j, i) == 255 {
				if biggestX < i {
					biggestX = i
				}

			}
		}
	}
	return
}

/* Counts white pixels with bigger X value than outer right vertical line
   returns true if there are some, false otherwise
*/
func isDanger(cropped gocv.Mat, biggestX int) (isDanger bool) {
	dangerPixelsCounter := 0
	for i := 0; i < cropped.Cols(); i++ {
		for j := 0; j < cropped.Rows(); j++ {
			if cropped.GetUCharAt(j, i) == 255 && i > biggestX {
				dangerPixelsCounter++
			}
		}
	}
	// Adjust this hardcoded value for better accuracy of detection
	if dangerPixelsCounter >= 5 {
		isDanger = true
	} else {
		isDanger = false
	}
	return

}

func machRegionOfInterest(regionResult *gocv.Mat, frame *gocv.Mat, template *gocv.Mat, regionRect *image.Rectangle) {
	templateDims := template.Size()
	gocv.MatchTemplate(*frame, *template, regionResult, gocv.TmCcoeff, gocv.NewMat())
	_, _, _, maxLoc := gocv.MinMaxLoc(*regionResult)
	regionRect.Min = maxLoc
	regionRect.Max = image.Point{X: maxLoc.X + templateDims[0], Y: maxLoc.Y + templateDims[1]}
}
