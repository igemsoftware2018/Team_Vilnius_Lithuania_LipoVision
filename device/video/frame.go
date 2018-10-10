package video

import (
	"image"
)

// Frame defines how video frame looks like.
// It's a struct to be passed as message.
type Frame struct {
	// frame is an underlying frame
	frame image.Image
}

// Skip checks if struct is passed it's lifetime,
// since it's a video frame, there's nothing to be late for,
// so it's never skipped
func (f Frame) Skip() <-chan struct{} {
	return make(chan struct{}, 1)
}

// Frame returns the underlying frame
func (f Frame) Frame() image.Image {
	return f.frame
}
