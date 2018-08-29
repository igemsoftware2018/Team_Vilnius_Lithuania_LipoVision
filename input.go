/* Fetching required video stream and template for region
   match */

package main

import (
	"image"

	"gocv.io/x/gocv"
)

func templateFetch() (filteredTemplate gocv.Mat) {
	// 50x50 matching intersection template
	var template = gocv.IMRead("template-intersection.png", gocv.IMWritePngStrategyFiltered)
	gocv.CvtColor(template, &template, gocv.ColorBGRToGray)
	gocv.AdaptiveThreshold(template, &template, 255, gocv.AdaptiveThresholdMean, gocv.ThresholdBinaryInv, 31, 2)
	gocv.Erode(template, &template, gocv.GetStructuringElement(0, image.Pt(3, 3)))
	return template
}

func captureVideo(path string) *gocv.VideoCapture {
	capture, err := gocv.VideoCaptureFile(path)

	if err != nil {
		panic(err)
	}

	return capture
}
