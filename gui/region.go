package gui

import (
	"image"

	"github.com/gotk3/gotk3/gtk"
)

// NewRegionControl creates a control widget
func NewRegionControl() (*RegionControl, error) {
	frame, frameErr := gtk.FrameNew("Reference stream")
	if frameErr != nil {
		return nil, frameErr
	}

	image, imgErr := gtk.ImageNewFromFile("template-intersection.png")
	if imgErr != nil {
		return nil, imgErr
	}

	frame.Add(image)
	return &RegionControl{rootFrame: frame, image: image}, nil
}

// RegionControl is the cut region widget for reference
type RegionControl struct {
	Control

	rootFrame *gtk.Frame

	image *gtk.Image
}

// Root returns the root element of this control
func (rc *RegionControl) Root() gtk.IWidget {
	return rc.rootFrame
}

// ShowFrame instructs the image container to display given image
func (rc *RegionControl) ShowFrame(frame image.Image) error {
	return showFrame(rc.image, frame)
}
