//dropletgenomics package defines the DropletGenomics company's microfluidincs device
package dropletgenomics

import (
	"context"
	"fmt"
	"image/png"
	"net/http"
	"time"
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

//CreateDropletGenomicsDevice returns a device configured with given configuration
func Create() *Device {
	device := Device{pumpExperiment: 4}
	device.establishPumpsWithID()
	return &device
}

//Device is DropletGenomics' rendition of microfluidics devices
type Device struct {
	pumpExperiment int
	pumps          []Pump
	camera         Camera
}

//Stream launches async stream decoding of ctx lifetime
func (device Device) Stream(ctx context.Context) <-chan Frame {
	const (
		streamEndpoint string = "http://192.168.1.100:8765/video_feed"
		frameRate      int    = 15
	)

	stream := make(chan Frame, 10)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("%s\n", "Stream closed")
				break
			default:
				response, err := client.Get(streamEndpoint)
				if err != nil {
					fmt.Printf("%s\n", "Failed to connect to stream")
					time.Sleep(time.Second)
					continue
				}
				fmt.Printf("%s\n", "Connected to stream")

				byteStream := response.Body
				for {
					buffer := make([]byte, 36, 36)
					_, readErr := byteStream.Read(buffer)
					if readErr != nil {
						fmt.Printf("Read err: %s\n", readErr)
						break
					}

					img, decodeError := png.Decode(byteStream)
					if decodeError != nil {
						fmt.Printf("%s\n", decodeError)
						break
					}
					frameLifetime, cancel := context.WithTimeout(ctx, time.Second/(time.Duration)(frameRate))
					stream <- Frame{frame: img, ctx: frameLifetime, cancel: cancel}

					buffer = make([]byte, 2, 2)
					byteStream.Read(buffer)
				}
				byteStream.Close()
			}
		}
	}()
	return stream
}

// Available determines whether connection to
// DropletGenomics device is established
func (device *Device) Available() bool {
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

//Camera returns the device's camera data
func (device Device) Camera() *Camera {
	return &device.camera
}

//NumPumps tells the user how many pumps are available
func (device Device) NumPumps() int {
	return len(device.pumps)
}

//Pump returns device's pump by it's id
func (device Device) Pump(index int) *Pump {
	return &device.pumps[index]
}

//RefreshAll launches refresh on all pumps
func (device Device) RefreshAll() error {
	for _, pump := range device.pumps {
		if err := pump.Invoke(PumpRefresh, nil); err != nil {
			return err
		}
	}
	return nil
}

//EstablishPumpsWithID creates a pump with ID
func (device *Device) establishPumpsWithID() {
	device.pumps = make([]Pump, device.pumpExperiment, device.pumpExperiment)
	for i := 0; i < device.pumpExperiment; i++ {
		device.pumps[i] = Pump{}
	}
}
