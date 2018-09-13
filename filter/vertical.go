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

func (vf Vertical) Apply(frame *gocv.Mat) error {
	var verticalsize = frame.Rows() / 5
	verticalStructure := gocv.GetStructuringElement(0, image.Pt(1, verticalsize))
	gocv.Erode(*frame, frame, verticalStructure)
	gocv.Dilate(*frame, frame, verticalStructure)

	return nil
}

//Produce Finds only vertical elements in the frame
func (vf Vertical) Produce(frame gocv.Mat) (gocv.Mat, error) {
	resultingFrame := frame.Clone()
	gocv.CvtColor(frame, &resultingFrame, gocv.ColorGrayToBGR)

	var verticalsize = resultingFrame.Rows() / 5
	verticalStructure := gocv.GetStructuringElement(0, image.Pt(1, verticalsize))
	gocv.Erode(resultingFrame, &resultingFrame, verticalStructure)
	gocv.Dilate(resultingFrame, &resultingFrame, verticalStructure)

	return resultingFrame, nil
}
