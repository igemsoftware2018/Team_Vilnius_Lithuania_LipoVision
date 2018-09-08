package filter

import (
	"image"

	"gocv.io/x/gocv"
)

func CreateSubtractFilter(region *image.Rectangle) Subtract {
	return Subtract{
		region:     region,
		subtractor: gocv.NewBackgroundSubtractorKNN(),
	}
}

//Subtract Is a background subtractor, used to find movement
type Subtract struct {
	region     *image.Rectangle
	subtractor gocv.BackgroundSubtractorKNN
}

func (s Subtract) Produce(frame gocv.Mat) (gocv.Mat, error) {
	resultingFrame := frame.Region(*s.region)
	s.subtractor.Apply(frame, &resultingFrame)
	gocv.Threshold(resultingFrame, &resultingFrame, 125, 255, gocv.ThresholdBinary)
	return resultingFrame, nil
}

func (s Subtract) Apply(frame *gocv.Mat) error {
	s.subtractor.Apply(*frame, frame)
	gocv.Threshold(*frame, frame, 125, 255, gocv.ThresholdBinary)
	return nil
}
