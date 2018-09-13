package processor

import (
	"context"
	"fmt"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
)

func CreateFrameProcessor() FrameProcessor {
	return FrameProcessor{context: context.Background()} // Default context must be Background(), WithContext() sets an optional context
}

//FrameProcessor Defines a processor for incoming frames of the stream
type FrameProcessor struct {
	context context.Context
}

//Process Processes a stream of frames coming from a device
func (fp FrameProcessor) Process(frames <-chan device.Frame) error {
	for frame := range frames {
		select {
		// If Process is cancelled, exit with error
		case <-fp.context.Done():
			return fp.context.Err()

		// If Frame is cancelled, log and skip
		case <-frame.SkippedFrame():
			fmt.Printf("%s", "Skipped frame")
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
