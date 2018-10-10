package filter

import "gocv.io/x/gocv"

// CreateCvtColorFilter creates a filter with color spec
func CreateCvtColorFilter(convCode gocv.ColorConversionCode) CvtColor {
	return CvtColor{convCode: convCode}
}

// CvtColor implements OpenCV's Color conversion feature
type CvtColor struct {
	convCode gocv.ColorConversionCode
}

// Produce returns a frame with conversion applied
func (cvt CvtColor) Produce(frame gocv.Mat) (gocv.Mat, error) {
	result := frame.Clone()
	err := cvt.Apply(&result)
	return result, err
}

// Apply performs convert on given image
func (cvt CvtColor) Apply(frame *gocv.Mat) error {
	gocv.CvtColor(*frame, frame, cvt.convCode)
	return nil
}
