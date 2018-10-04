package dropletgenomics

import (
	"context"
	"image"
)

//Frame defines how dropletgenomics frame should look like.
//It's a struct to be passed as message.
type Frame struct {

	//frame is an underlying frame
	frame image.Image

	//ctx defines Frame structure lifetime
	ctx context.Context

	//cancel ends Frame lifetime
	cancel context.CancelFunc
}

//Skip checks if struct is passed it's lifetime
//Since it's a frame, it means it should be skipped
func (f Frame) Skip() <-chan struct{} {
	return f.ctx.Done()
}

//Frame returns the underlying frame
//Upon frame retrieval, cancel is called, and frame's lifetime ends
func (f Frame) Frame() image.Image {
	defer f.cancel()
	return f.frame
}
