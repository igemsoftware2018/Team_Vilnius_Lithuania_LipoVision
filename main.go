package main

import (
	"image"
	"image/color"
	"os"

	"gocv.io/x/gocv"
)

func main() {

	capture := captureVideo(os.Args[1])
	var template = templateFetch()

	originalWindow := gocv.NewWindow("Original")

	regionSet := false
	var regionRect image.Rectangle
	var rectangleColor = color.RGBA{255, 255, 255, 0}

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
			machRegionOfInterest(&regionResult, &frame, &template, &regionRect)
		} else {
			cropped := frame.Region(regionRect)
			croppedWindow := gocv.NewWindow("Cropped")

			frameForSubtraction := subtractBackground(&frameOriginal, &regionRect, &fgbg)

			removeMovingNoisePixels(frameCount, &previousFrame, &frame)
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
