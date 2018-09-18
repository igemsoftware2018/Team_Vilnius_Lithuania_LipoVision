package main

import (
	"errors"

	"gocv.io/x/gocv"
)

//captureFrame Converts gocv provided handling to idiomatic Go
func captureFrame(capture *gocv.VideoCapture) (gocv.Mat, error) {
	frame := gocv.NewMat()
	ok := capture.Read(&frame)
	if !ok {
		return frame, errors.New("failed to read frame from source")
	}
	return frame, nil
}

// func main() {

// 	capture := captureVideo(os.Args[1])
// 	var template = templateFetch()

// 	originalWindow := gocv.NewWindow("Original")

// 	var regionRect image.Rectangle
// 	var rectangleColor = color.RGBA{255, 255, 255, 0}

// 	previousFrame := gocv.NewMat()
// 	frameCount := 0

// 	frameFilters := []filter.Filter{
// 		filter.CreateSubtractFilter(&regionRect),
// 		filter.CreateNoiseFilter(&previousFrame, &frameCount),
// 	}

// 	regionFilters := []filter.Filter{
// 		filter.CreateVerticalFilter(),
// 		filter.CreateLineApplyFilter(),
// 	}

// 	regionSet := false
// 	cancelled := false
// 	for !cancelled {
// 		// frame is the original, has to represent the original,
// 		// so no heavy manipulations should be applied
// 		frame, err := captureFrame(capture)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "Failed frame")
// 			break
// 		}

// 		gocv.CvtColor(frame, &frame, gocv.ColorBGRToGray)

// 		originalFrame := frame.Clone()

// 		gocv.Threshold(frame, &frame, 125, 255, gocv.ThresholdBinaryInv)

// 		if !regionSet {
// 			regionResult := gocv.NewMat()
// 			machRegionOfInterest(&regionResult, &frame, &template, &regionRect)
// 		} else {
// 			croppedWindow := gocv.NewWindow("Cropped")

// 			manipulatedFrame := frame.Clone()
// 			err := filter.ApplyFilters(&manipulatedFrame, frameFilters)
// 			if err != nil {
// 				fmt.Fprintln(os.Stderr, err)
// 				break
// 			}

// 			cropped := originalFrame.Region(regionRect)
// 			croppedForAdd := originalFrame.Region(regionRect)
// 			// Origginally conversion was into this type
// 			// gocv.CvtColor(*frame, frame, gocv.ColorGrayToBGR)
// 			err = filter.ApplyFilters(&cropped, regionFilters)
// 			if err != nil {
// 				fmt.Fprintln(os.Stderr, err)
// 				break
// 			}

// 			// Left in case of debugging for now
// 			// fmt.Printf("Cropped Rows: %d, cols, %d, Type: %s\n", cropped.Rows(), cropped.Cols(), cropped.Type())
// 			// fmt.Printf("CroppedForAdd Rows: %d, cols, %d, Type: %s\n", croppedForAdd.Rows(), croppedForAdd.Cols(), croppedForAdd.Type())

// 			biggestX := findBiggestXOfWhitePixel(&cropped)
// 			// Merge vertical line frame + Subtracted moving bubble frame
// 			gocv.AddWeighted(cropped, 1, croppedForAdd, 1, 0, &cropped)

// 			if isDanger(cropped, biggestX) {
// 				gocv.PutText(&cropped, "DANGER!", image.Pt(cropped.Cols()/8, cropped.Rows()/4*3), 0, 0.3, color.RGBA{255, 255, 255, 9}, 1)
// 			}
// 			croppedWindow.IMShow(cropped)

// 		}
// 		gocv.Rectangle(&frame, regionRect, rectangleColor, 2)

// 		originalWindow.IMShow(frame)
// 		if originalWindow.WaitKey(200)&0xFF == 'r' {
// 			regionSet = true
// 		}

// 		previousFrame = frame.Clone()
// 		frameCount++
// 	}

// }
