// Package video describes video fetching as a device
package video

import (
	"context"
	"time"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
	log "github.com/sirupsen/logrus"
	"gocv.io/x/gocv"
)

// Create returns a video.Device configured with given setttings
func Create(videoPath string, framerate int) Device {
	log.Info("video file set as: ", videoPath)
	log.Info("framerate set as: ", framerate)
	return Device{videoPath: videoPath,
		camera: Camera{FrameRate: framerate}}
}

// Device defines a mock device for gocv video retrieval
type Device struct {
	videoPath string
	camera    Camera
}

// Stream fetches frames on certain times, to mimic stream
func (dev Device) Stream(ctx context.Context) <-chan device.Frame {
	stream := make(chan device.Frame, dev.camera.FrameRate)

	capture, err := gocv.VideoCaptureFile(dev.videoPath)
	if err != nil {
		return nil
	}
	go func() {
	FrameFetch:
		for {
			select {
			case <-ctx.Done():
				break FrameFetch
			default:
				frame := gocv.NewMat()
				if !capture.Read(&frame) {
					log.Warn("could not read frame from video")
					continue
				}

				img, err := frame.ToImage()
				if err != nil {
					log.Warn("could not convert gocv.Mat to image.Image")
					continue
				}
				stream <- Frame{frame: img}
				time.Sleep((time.Duration)((int)(time.Second) / dev.camera.FrameRate))
			}
		}
	}()
	return stream
}

// Available checks if device is reachable,
// a video file is always reachable
func (Device) Available() bool {
	return true
}
