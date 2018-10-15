package gui

import "github.com/gotk3/gotk3/gtk"

// NewRegionControl creates a control widget
func NewRegionControl() (*RegionControl, error) {
	frame, frameErr := gtk.FrameNew("Reference frame")
	if frameErr != nil {
		return nil, frameErr
	}

	image, imgErr := gtk.ImageNewFromFile("template-intersection.png")
	if imgErr != nil {
		return nil, imgErr
	}

	frame.Add(image)
	return &RegionControl{rootFrame: frame}, nil
}

// RegionControl is the cut region widget for reference
type RegionControl struct {
	Control

	rootFrame *gtk.Frame
}

// Root returns the root element of this control
func (rc RegionControl) Root() gtk.IWidget {
	return rc.rootFrame
}
