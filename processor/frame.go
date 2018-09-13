package processor

import (
	"context"
	"fmt"

	"gocv.io/x/gocv"
)

// This will be in the device package, since the device owns the stream and sets the context
type Frame struct {
	frame gocv.Mat
	ctx   context.Context
}

func CreateFrameProcessor() FrameProcessor {
	return FrameProcessor{context: context.Background()} // Default context must be Background(), WithContext() sets an optional context
}

//FrameProcessor Defines a processor for incoming frames of the stream
type FrameProcessor struct {
	context context.Context
}

//Process Processes a stream of frames coming from a device
func (fp FrameProcessor) Process(frames <-chan Frame) error {
	for frame := range frames {
		select {
		// If Process is cancelled, exit with error
		case <-fp.context.Done():
			return fp.context.Err()

		// If Frame is cancelled, log and skip
		case <-frame.ctx.Done():
			fmt.Printf("%s", frame.ctx.Err())
			continue
		}

		// Do Frame processing here (move to unexported function preferably)
		// Needed initialization to be done in CreateFrameProcessor, held in struct itself
	}
	return nil
}

//WithContext Allows to set a Context for the Processsor
func (fp FrameProcessor) WithContext(context context.Context) FrameProcessor {
	fp.context = context
	return fp
}
