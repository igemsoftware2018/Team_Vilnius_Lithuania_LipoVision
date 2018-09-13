package filter

import (
	"image"

	"gocv.io/x/gocv"
)

//CreateVerticalFilter Creates a Vertical filter
func CreateVerticalFilter() Vertical {
	return Vertical{}
}

//VerticalFilter Finds vertical element of a a frame
type Vertical struct {
}

//Apply Applies vertical element filter to frame
func (vf Vertical) Apply(frame *gocv.Mat) error {
	var verticalsize = frame.Rows() / 5
	verticalStructure := gocv.GetStructuringElement(0, image.Pt(1, verticalsize))
	gocv.Erode(*frame, frame, verticalStructure)
	gocv.Dilate(*frame, frame, verticalStructure)

	return nil
}

//Produce Produces a frame with vertical line filter applied
func (vf Vertical) Produce(frame gocv.Mat) (gocv.Mat, error) {
	resultingFrame := frame.Clone()
	err := vf.Apply(&resultingFrame)
	return resultingFrame, err
}
