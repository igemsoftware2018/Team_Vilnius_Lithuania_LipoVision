package processor

import (
	"image"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
	log "github.com/sirupsen/logrus"
)

/* These constants define what streams are supported on this Processor.
StreamRegion     - matched region frame
StreamSubtracted - full frame that is algorithm-subtracted
StreamOriginal   - original frame, as it came from device*/
const (
	StreamRegion     string = "region"
	StreamSubtracted string = "subtracted"
	StreamOriginal   string = "origina;"
)

/* There are the settings that are settable for this Processor
SettingAutonomicRun - enables the auto-coating feature*/
const (
	SettingAutonomicRun string = "auto"
)

const (
	frameQueueSize int = 10
)

// NewFrameProcessor Creates a frame processor with given settings
func NewFrameProcessor() *FrameProcessor {
	log.WithFields(log.Fields{
		"processor": "Frame",
	}).Info("Created")
	return &FrameProcessor{}
}

// FrameProcessor Defines a processor for incoming frames of the stream
type FrameProcessor struct {
	Processor

	autonomicRun bool
}

// Launch starts a processing goroutine for the stream of frames coming from a device
func (fp *FrameProcessor) Launch(frames <-chan device.Frame, streamHandlers map[string]func(image.Image)) {
	log.WithFields(log.Fields{"processor": "Frame"}).Info("Launched")

	go func() {
		for frame := range frames {
			select {
			case <-frame.Skip():
				log.WithFields(log.Fields{"processor": "Frame"}).Warn("Frame skip")
				continue
			default:
			}

			frameImage := frame.Frame()
			if h, ok := streamHandlers[StreamOriginal]; ok {
				h(frameImage)
			}
		}
		log.WithFields(log.Fields{"processor": "Frame"}).Info("Finished")
	}()
}

// Set the configurables on this Processor
func Set(id string, value interface{}) {

}
