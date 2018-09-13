package filter

import "gocv.io/x/gocv"

func CreateLineApplyFilter() LineApply {
	return LineApply{}
}

//LineApply Applies white lines of full height of window when white pixel found
type LineApply struct {
}

func (la LineApply) Apply(frame *gocv.Mat) error {
	for i := 0; i < frame.Cols(); i++ {
		for j := 0; j < frame.Rows(); j++ {
			if frame.GetUCharAt(j, i) == 255 {
				for k := 0; k < frame.Rows(); k++ {
					frame.SetUCharAt(k, i, 255)
				}
			}
		}
	}
	return nil
}

func (la LineApply) Produce(frame gocv.Mat) (gocv.Mat, error) {
	resultFrame := frame.Clone()
	err := la.Apply(&resultFrame)
	return resultFrame, err
}
