package main

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"gocv.io/x/gocv"
)

func main() {

	// Fetching required input data
	capture := captureVideo(os.Args[1])
	var template = templateFetch()

	// Create basic original window
	originalWindow := gocv.NewWindow("Original")

	// Matching region detection variables
	regionSet := false
	var regionRect image.Rectangle
	var rectangleColor = color.RGBA{255, 255, 255, 0}

	// Object for Background Subtraction
	var fgbg = gocv.NewBackgroundSubtractorKNN()

	cancelled := false
	previousFrame := gocv.NewMat()
	frameCount := 0

	for !cancelled {
		frame := gocv.NewMat()
		frameOk := capture.Read(&frame)
		if !frameOk {
			break
		}

		frameOriginal := frame.Clone()
		gocv.CvtColor(frame, &frame, gocv.ColorBGRToGray)
		gocv.Threshold(frame, &frame, 125, 255, gocv.ThresholdBinaryInv)

		if !regionSet {
			regionResult := gocv.NewMat()
			templateDims := template.Size()
			gocv.MatchTemplate(frame, template, &regionResult, gocv.TmCcoeff, gocv.NewMat())
			_, _, _, maxLoc := gocv.MinMaxLoc(regionResult)
			fmt.Printf("Region found: [%d, %d]\n", maxLoc.X, maxLoc.Y)
			regionRect.Min = maxLoc
			regionRect.Max = image.Point{X: maxLoc.X + templateDims[0], Y: maxLoc.Y + templateDims[1]}
		} else {
			cropped := frame.Region(regionRect)
			croppedWindow := gocv.NewWindow("Cropped")

			frameForSubtraction := frameOriginal.Region(regionRect)
			fgbg.Apply(frameForSubtraction, &frameForSubtraction)
			gocv.Threshold(frameForSubtraction, &frameForSubtraction, 125, 255, gocv.ThresholdBinary)

			removeMovingNoisePixels(frameCount, &previousFrame, &frame)
			removeBorderPixelNoise(&cropped)
			findVerticalElements(&cropped)

			gocv.CvtColor(cropped, &cropped, gocv.ColorBGRToGray)

			applyWhiteLinesFullHeight(&cropped)

			biggestX := findBiggestXOfWhitePixel(&cropped)
			// Merge vertical line frame + Subtracted moving bubble frame
			gocv.AddWeighted(cropped, 1, frameForSubtraction, 1, 0, &cropped)

			if isDanger(cropped, biggestX) {
				gocv.PutText(&cropped, "DANGER!", image.Pt(cropped.Cols()/8, cropped.Rows()/4*3), 0, 0.3, color.RGBA{255, 255, 255, 9}, 1)
			}
			croppedWindow.IMShow(cropped)

		}
		gocv.Rectangle(&frame, regionRect, rectangleColor, 2)

		originalWindow.IMShow(frame)
		if originalWindow.WaitKey(200)&0xFF == 'r' {
			regionSet = true
		}

		previousFrame = frame.Clone()
		frameCount++
	}

}
