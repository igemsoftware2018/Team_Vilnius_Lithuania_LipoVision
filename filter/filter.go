package filter

import "gocv.io/x/gocv"

//Filter Defines a computation on a frame,
//that may do either noise removal or something else.
type Filter interface {

	//Produce Responsible for allocating a new frame with manipulations on it
	Produce(gocv.Mat) (gocv.Mat, error)

	//Manipulate Responsible for changing an existing frame
	Apply(*gocv.Mat) error
}

//ApplyFilters Applies a chain of manipulations on a frame.
//The order of manipulations is the same as the slice.
//Will stop midway if errors encountered.
func ApplyFilters(frame *gocv.Mat, filters []Filter) error {
	for _, filter := range filters {
		err := filter.Apply(frame)
		if err != nil {
			return err
		}
	}
	return nil
}

//CreateFrames Applies a slice of Filters, to produce a slice of frames.
//The order of frames produces has to be the same as the order of input filters.
//It will not complete if errors in filters will be encountered.
func CreateFrames(frame gocv.Mat, filters []Filter) ([]gocv.Mat, error) {
	resultFrames := make([]gocv.Mat, len(filters))
	for index, filter := range filters {
		var err error
		resultFrames[index], err = filter.Produce(frame)
		if err != nil {
			return resultFrames, err
		}
	}
	return resultFrames, nil
}
