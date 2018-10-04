package dropletgenomics

import (
	"context"
	"fmt"
	"image"
	"image/png"
	"net"
	"net/http"
	"time"
)

const streamEndpoint string = "http://example.com/"
const frameRate int64 = 30

type Device struct {
}

func (d Device) Stream(ctx context.Context) <-chan Frame {
	stream := make(chan Frame, 10)
	go func() {
		for {
			select {
			case <-ctx.Done():
				break
			default:
				response, err := http.Get(streamEndpoint)
				if err != nil {
					fmt.Printf("Failed to connect to stream")
					continue
				}
				byteStream := response.Body
				var decodeError error = nil
				for decodeError == nil {
					var img image.Image
					img, decodeError = png.Decode(byteStream)
					frameLifetime, cancel := context.WithTimeout(ctx, time.Second/(time.Duration)(frameRate))
					stream <- Frame{frame: img, ctx: frameLifetime, cancel: cancel}
				}
			}
		}
	}()
	return stream
}

func (d Device) Available() bool {
	timeout := time.Duration(200 * time.Millisecond)
	_, err := net.DialTimeout("tcp", streamEndpoint, timeout)
	if err != nil {
		return false
	}
	return true
}
