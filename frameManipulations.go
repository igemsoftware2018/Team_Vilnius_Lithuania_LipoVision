package main

import (
	"image"

	"gocv.io/x/gocv"
)

// Comparison with previous frame for moving pixels detection and removal
func removeMovingNoisePixels(frameCount int, previousFrame *gocv.Mat, frame *gocv.Mat) {
	for i := 0; i < frame.Rows(); i++ {
		for j := 0; j < frame.Cols(); j++ {
			if frameCount > 0 && previousFrame.GetUCharAt(i, j) != frame.GetUCharAt(i, j) {
				frame.SetUCharAt(i, j, 0)
			}
		}
	}
}

// Finds only vertical elements in the frame
func findVerticalElements(cropped *gocv.Mat) {
	gocv.CvtColor(*cropped, cropped, gocv.ColorGrayToBGR)
	var verticalsize = cropped.Rows() / 5
	verticalStructure := gocv.GetStructuringElement(0, image.Pt(1, verticalsize))
	gocv.Erode(*cropped, cropped, verticalStructure)
	gocv.Dilate(*cropped, cropped, verticalStructure)
}

// Apply white lines of full height of window when white pixel found
func applyWhiteLinesFullHeight(cropped *gocv.Mat) {
	for i := 0; i < cropped.Cols(); i++ {
		for j := 0; j < cropped.Rows(); j++ {
			if cropped.GetUCharAt(j, i) == 255 {
				for k := 0; k < cropped.Rows(); k++ {
					cropped.SetUCharAt(k, i, 255)
				}

			}
		}
	}
}

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

func subtractBackground(frameOriginal *gocv.Mat, regionRect *image.Rectangle, fgbg *gocv.BackgroundSubtractorKNN) gocv.Mat {
	frameForSubtraction := frameOriginal.Region(*regionRect)
	fgbg.Apply(frameForSubtraction, &frameForSubtraction)
	gocv.Threshold(frameForSubtraction, &frameForSubtraction, 125, 255, gocv.ThresholdBinary)
	return frameForSubtraction
}
