// Package dropletgenomics defines the DropletGenomics company's microfluidincs device
package dropletgenomics

import (
	"context"
	"fmt"
	"image"
	"image/png"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
)

var client http.Client

func init() {
	transport := http.Transport{}
	client = http.Client{
		Transport: &transport,
		//Timeout:   2 * time.Second,
	}
	transport.DisableKeepAlives = true
}

// Create returns a device configured with given configuration
func Create(usedPumps int) *Device {
	pumps := make([]Pump, usedPumps, usedPumps)
	for i := 0; i < usedPumps; i++ {
		pumps[i] = Pump{BaseAddr: "http://192.168.1.100:8764"}
	}
	device := Device{
		pumpExperiment: usedPumps,
		camera: Camera{
			BaseAddr: "http://192.168.1.100:8765",
		},
		pumps: pumps,
	}
	return &device
}

// Device is DropletGenomics' rendition of microfluidics devices
type Device struct {
	pumpExperiment int
	pumps          []Pump
	camera         Camera
}

// decodeFrame reads and discards the non-png header, then
// uses the png decoder to decode the frame into an image.Image
func decodeFrame(stream io.ReadCloser) (image.Image, error) {
	buffer := make([]byte, 36, 36)
	_, readErr := stream.Read(buffer)
	if readErr != nil {
		return nil, readErr
	}

	img, decodeError := png.Decode(stream)
	if decodeError != nil {
		return nil, decodeError
	}

	return img, nil
}

// removeTrail removes the two-byte trail after an image
func removeTrail(stream io.ReadCloser) {
	buffer := make([]byte, 2, 2)
	stream.Read(buffer)
}

// Stream launches async stream decoding of ctx lifetime
func (Device) Stream(ctx context.Context) <-chan Frame {
	const (
		streamEndpoint string = "http://192.168.1.100:8765/video_feed"
		frameRate      int    = 20
	)

	stream := make(chan Frame, 20)
	go func() {
		complete := false
		for !complete {
			select {
			case <-ctx.Done():
				fmt.Printf("%s\n", "Stream closed")
				complete = true
			default:
				response, err := client.Get(streamEndpoint)
				if err != nil {
					fmt.Fprintf(os.Stderr, "failed to connect to stream: %s\n", err)
					time.Sleep(time.Second)
					continue
				}
				fmt.Printf("%s\n", "Connected to stream")

				byteStream := response.Body
				for {
					img, err := decodeFrame(byteStream)
					if err != nil {
						fmt.Fprintf(os.Stderr, "decode error: %s", err)
						break
					}

					frameLifetime, cancel := context.WithTimeout(ctx, time.Second/(time.Duration)(frameRate))
					stream <- Frame{frame: img, ctx: frameLifetime, cancel: cancel}

					removeTrail(byteStream)
				}
				byteStream.Close()
			}
		}
	}()
	return stream
}

// Available determines whether connection to
// DropletGenomics device is established
func (d *Device) Available() bool {
	const url string = "http://192.168.1.100/"
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		return false
	}
	return true
}

// Camera returns the device's camera data
func (d Device) Camera() device.Client {
	return &d.camera
}

// NumPumps tells the user how many pumps are available
func (d Device) NumPumps() int {
	return len(d.pumps)
}

// Pump returns device's pump by it's id
func (d Device) Pump(index int) device.Client {
	return &d.pumps[index]
}

// RefreshAll launches refresh on all pumps
func (d Device) RefreshAll() error {
	for _, pump := range d.pumps {
		if err := pump.Invoke(device.PumpRefresh, nil); err != nil {
			return err
		}
	}
	return nil
}
