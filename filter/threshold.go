package filter

import "gocv.io/x/gocv"

// CreateThesholdFilter returns a configured filter
func CreateThesholdFilter(thresh float32, maxValue float32, threshType gocv.ThresholdType) Threshold {
	return Threshold{threshType: threshType, thresh: thresh, maxValue: maxValue}
}

// Threshold implements OpenCV's threshold as a filter
type Threshold struct {
	thresh     float32
	maxValue   float32
	threshType gocv.ThresholdType
}

// Produce creates a frame with threshold applied
func (thresh Threshold) Produce(frame gocv.Mat) (gocv.Mat, error) {
	result := frame.Clone()
	err := thresh.Apply(&result)
	return result, err
}

// Apply performs thresholding on given frame
func (thresh Threshold) Apply(frame *gocv.Mat) error {
	gocv.Threshold(*frame, frame, thresh.thresh, thresh.maxValue, thresh.threshType)
	return nil
}
