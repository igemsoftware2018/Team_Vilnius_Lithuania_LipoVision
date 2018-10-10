package processor

import (
	"context"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
	log "github.com/sirupsen/logrus"
)

// CreateFrameProcessor Creates a frame processor with given settings
func CreateFrameProcessor() FrameProcessor {
	log.Info("FrameProcessor created")
	return FrameProcessor{}
}

// FrameProcessor Defines a processor for incoming frames of the stream
type FrameProcessor struct {
}

// Process Processes a stream of frames coming from a device
func (fp *FrameProcessor) Process(ctx context.Context, frames <-chan device.Frame) {
	go func() {
		for frame := range frames {
			select {
			case <-ctx.Done():
				log.Info("Processor stopped")
				return
			case <-frame.Skip():
				log.Info("Frame skip")
				continue
			default:
			}

			frame.Frame()
		}
	}()
}
