package processor

import (
	"image"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
	log "github.com/sirupsen/logrus"
)

const (
	// For identification on witch stream to return
	StreamRegion     string = "region"
	StreamSubtracted string = "subtracted"
	StreamOriginal   string = "origina;"

	// Settings of this Processor
	SettingAutonomicRun string = "auto"

	frameQueueSize int = 10
)

// CreateFrameProcessor Creates a frame processor with given settings
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

// Set
func Set(id string, value interface{}) {

}
