package device

import (
	"context"
	"image"
)

//Frame Defines a frame structure that is a part of the stream
type Frame interface {

	//Frame Gets the underlying frame
	Frame() image.Image

	//SkippedFrame provides context.Done() like method of
	//communicating that a frame should be skipped
	Skip() <-chan struct{}
}

//Device Is a physical device that program is connecting to
type Device interface {

	//Stream Returns a device's video stream that can be cancelled
	Stream(context.Context) <-chan Frame

	//Available Checks if device is available
	Available() bool
}
