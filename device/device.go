package device

import (
	"context"
)

//Frame Defines a frame structure that is a part of the stream
type Frame struct {
	frame interface{}
	ctx   context.Context
}

// GetFrame gets underlying Frame
func (f Frame) GetFrame() interface{} {
	return f.frame
}

//Skip Exposes innner context Done()
//Which basically means that the frame should be skipped
func (f Frame) Skip() <-chan struct{} {
	return f.ctx.Done()
}

//Device Is a physical device that program is connecting to
type Device interface {

	//Stream Returns a device's video stream that can be cancelled
	Stream(context.Context) <-chan Frame

	//Available Checks if device is available
	Available() bool
}
