package dropletgenomics

import (
	"context"
	"image"
)

// CreateFrame creates a frame with given frame and context
func CreateFrame(ctx context.Context, frame image.Image) Frame {
	fCtx, cancel := context.WithCancel(ctx)
	return Frame{ctx: fCtx, cancel: cancel, frame: frame}
}

// Frame defines how dropletgenomics frame looks like. It's a struct to be passed as message.
// Getting the internal frame ends this struct's lifetime.
type Frame struct {

	// frame is an underlying frame
	frame image.Image

	// ctx defines Frame structure lifetime
	ctx context.Context

	// cancel ends Frame lifetime
	cancel context.CancelFunc
}

// Skip checks if struct is passed it's lifetime
// Since it's a frame, it means it should be skipped
func (f Frame) Skip() <-chan struct{} {
	return f.ctx.Done()
}

// Frame returns the underlying frame
// Upon frame retrieval, cancel is called, ending frame's lifetime
func (f Frame) Frame() image.Image {
	defer f.cancel()
	return f.frame
}
