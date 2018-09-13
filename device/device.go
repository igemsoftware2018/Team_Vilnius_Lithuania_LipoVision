package device

import (
	"context"

	"gocv.io/x/gocv"
)

//Frame Defines a frame structure that is a part of the stream
type Frame struct {
	frame gocv.Mat
	ctx   context.Context
}

//Frame Gets the underlying frame
func (f Frame) Frame() gocv.Mat {
	return f.frame
}

//SkippedFrame Exposes innner context Done()
//Which basically means that the frame should be skipped
func (f Frame) SkippedFrame() <-chan struct{} {
	return f.ctx.Done()
}

//Device Is a physical device that program is connecting to
type Device interface {

	//Stream Returns a device's video stream that can be cancelled
	Stream(context.Context) <-chan Frame

	//Available Checks if device is available
	Available() bool
}
