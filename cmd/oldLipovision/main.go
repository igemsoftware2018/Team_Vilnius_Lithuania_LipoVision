package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"os"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/filter"
	log "github.com/sirupsen/logrus"
	"gocv.io/x/gocv"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

//captureFrame Converts gocv provided handling to idiomatic Go
func captureFrame(capture *gocv.VideoCapture) (gocv.Mat, error) {
	frame := gocv.NewMat()
	ok := capture.Read(&frame)
	if !ok {
		return frame, errors.New("failed to read frame from source")
	}
	return frame, nil
}

func main() {

	capture := captureVideo(os.Args[1])
	var template = templateFetch()

	originalWindow := gocv.NewWindow("Original")

	croppedWindow := gocv.NewWindow("Cropped")
	croppedWindow.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)
	croppedWindow.SetWindowProperty(gocv.WindowPropertyAspectRatio, gocv.WindowKeepRatio)
	croppedWindow.SetWindowProperty(gocv.WindowPropertyVisible, 1)

	var regionRect image.Rectangle
	var rectangleColor = color.RGBA{255, 255, 255, 0}

	previousFrame := gocv.NewMat()
	frameCount := 0

	subtract := gocv.NewBackgroundSubtractorKNN()
	defer subtract.Close()

	frameFilters := []filter.Filter{
		filter.CreateNoiseFilter(&previousFrame, &frameCount),
	}

	regionFilters := []filter.Filter{
		filter.CreateVerticalFilter(),
		filter.CreateLineApplyFilter(),
	}

	regionSet := false
	cancelled := false
	for !cancelled {
		// frame is the original, has to represent the original,
		// so no heavy manipulations should be applied
		frame, err := captureFrame(capture)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed frame")
			break
		}

		originalFrame := frame.Clone()

		gocv.CvtColor(frame, &frame, gocv.ColorBGRToGray)
		gocv.Threshold(frame, &frame, 125, 255, gocv.ThresholdBinaryInv)

		if !regionSet {
			regionResult := gocv.NewMat()
			machRegionOfInterest(&regionResult, &frame, &template, &regionRect)
			gocv.PutText(&originalFrame, "Press 'r' to fix region in place", image.Pt(originalFrame.Cols()/8, originalFrame.Rows()/4*3), 0, 1, color.RGBA{255, 0, 0, 9}, 1)
		} else {
			rectangleColor = color.RGBA{255, 0, 0, 0}

			err := filter.ApplyFilters(&frame, frameFilters)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				break
			}

			cropped := frame.Region(regionRect)
			cropped = cropped.Clone()

			croppedForAdd := originalFrame.Region(regionRect)
			croppedForAdd = croppedForAdd.Clone()

			err = filter.ApplyFilters(&cropped, regionFilters)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				break
			}

			subtract.Apply(croppedForAdd, &croppedForAdd)

			biggestX := findBiggestXOfWhitePixel(&cropped)

			if isDanger(croppedForAdd, biggestX) {
				gocv.PutText(&cropped, "DANGER!", image.Pt(cropped.Cols()/8, cropped.Rows()/4*3), 0, 0.3, color.RGBA{255, 255, 255, 9}, 1)
			}

			gocv.AddWeighted(cropped, 1, croppedForAdd, 1, 0, &cropped)
			croppedWindow.IMShow(cropped)

		}

		gocv.Rectangle(&originalFrame, regionRect, rectangleColor, 2)
		originalWindow.IMShow(originalFrame)
		if originalWindow.WaitKey(200)&0xFF == 'r' {
			regionSet = true
		}

		previousFrame = frame.Clone()
		frameCount++
	}

}
