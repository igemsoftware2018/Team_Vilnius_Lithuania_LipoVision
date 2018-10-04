package dropletgenomics

import (
	"context"
	"fmt"
	"image"
	"image/png"
	"net/http"
	"strconv"
	"time"
)

const streamEndpoint string = "http://example.com/"
const frameRate int64 = 30

type Device struct {
	IPAddress         string
	HTTPPort          int
	PumpDataPort      int
	RecordingDataPort int
	PumpExperiment    int
	Pumps             []pump
	camera            camera
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

// Available determines whether connection to
// DropletGenomics device is established
func (device *Device) Available() bool {
	url := setupDeviceURL(device)
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		return false
	}
	return true
}

func (device *Device) GetVolumeTarget(pump int) float64 {
	return device.Pumps[pump].VolumeTarget
}
func (device *Device) GetPurgeRate(pump int) float64 {
	return device.Pumps[pump].PurgeRate
}
func (device *Device) GetPumpID(pump int) float64 {
	return device.Pumps[pump].PumpID
}
func (device *Device) GetRateW(pump int) float64 {
	return device.Pumps[pump].PumpID
}
func (device *Device) GetVolume(pump int) float64 {
	return device.Pumps[pump].Volume
}
func (device *Device) GetStatus(pump int) bool {
	return device.Pumps[pump].Status
}
func (device *Device) GetName(pump int) string {
	return device.Pumps[pump].Name
}
func (device *Device) GetDirection(pump int) bool {
	return device.Pumps[pump].Direction
}
func (device *Device) GetSyringe(pump int) float64 {
	return device.Pumps[pump].Syringe
}
func (device *Device) GetUsed(pump int) bool {
	return device.Pumps[pump].Used
}
func (device *Device) GetVolumeTargetW(pump int) float64 {
	return device.Pumps[pump].VolumeTargetW
}
func (device *Device) GetVolumeW(pump int) float64 {
	return device.Pumps[pump].VolumeW
}
func (device *Device) GetRate(pump int) float64 {
	return device.Pumps[pump].Rate
}
func (device *Device) GetStalled(pump int) bool {
	return device.Pumps[pump].Stalled
}
func (device *Device) GetForce(pump int) float64 {
	return device.Pumps[pump].Force
}

func (device *Device) DefinePumpExperiment(numberOfPumps int) {
	//TODO : get endpoints from device in GMC
}

func (device *Device) EstablishPumpsWithID() {
	device.Pumps = make([]pump, device.PumpExperiment, device.PumpExperiment)
	for i := 0; i < device.PumpExperiment; i++ {
		device.Pumps[i] = NewPump(i)
	}
}

func (device *Device) Update() bool {
	for i := 0; i < len(device.Pumps); i++ {
		if !device.Pumps[i].updatePumpValues("http://" + device.IPAddress + ":" + strconv.Itoa(device.PumpDataPort) + "/refresh") {
			return false
		}
	}
	return true
}

func (device *Device) Reset(pump int) bool {
	if !device.Pumps[pump].resetPump("http://" + device.IPAddress + ":" + strconv.Itoa(device.PumpDataPort) + "/update_pars") {
		return false
	}
	return true
}

func (device *Device) ToggleWithdrawInfuse(pump int, widthdraw bool) bool {
	if pump == -1 {
		for i := 0; i < len(device.Pumps); i++ {
			if !device.Pumps[i].toggleWithdrawInfuse(device.setupToggleURL(), widthdraw) {
				return false
			}
		}
		return true
	}
	if !device.Pumps[pump].toggleWithdrawInfuse(device.setupToggleURL(), widthdraw) {
		return false
	}
	return true
}

func (device *Device) DisplayPumpValues(selectedPump int) (answer string) {
	if selectedPump == -1 {
		for _, pump := range device.Pumps {
			if !pump.updatePumpValues(device.setupGetValuesURL()) {
				panic("ERROR")
			}
			answer += formatPumpValues(&pump)
		}
		return
	}
	if !device.Pumps[selectedPump].updatePumpValues(device.setupGetValuesURL()) {
		panic("ERROR")
	}
	answer += formatPumpValues(&device.Pumps[selectedPump])
	return

}

func (device *Device) TogglePump(selectedPump int, start bool) bool {
	if device.Available() {
		if selectedPump == -1 {
			for _, pump := range device.Pumps {
				if !pump.togglePump(device.setupToggleURL(), start) {
					return false
				}
			}
			return true
		}
		if !device.Pumps[selectedPump].togglePump(device.setupToggleURL(), start) {
			return false
		}
		return true

	}
	return false
}

func (device *Device) SetPumpVolume(selectedPump int, volume int) bool {
	if device.Available() {
		if selectedPump == -1 {
			for _, pump := range device.Pumps {
				if !pump.setVolume(device.setupToggleURL(), volume) {
					return false
				}
			}
			return true
		}
		if !device.Pumps[selectedPump].setVolume(device.setupToggleURL(), volume) {
			return false
		}
		return true

	}
	return false

}

func (device *Device) SetPumpTargetVolume(selectedPump int, volume int) bool {
	if device.Available() {
		if selectedPump == -1 {
			for _, pump := range device.Pumps {
				if !pump.setTargetVolume(device.setupToggleURL(), volume) {
					return false
				}
			}
			return true
		}
		if !device.Pumps[selectedPump].setTargetVolume(device.setupToggleURL(), volume) {
			return false
		}
		return true

	}
	return false

}

func (device *Device) SetCameraIllumination(ill float64) bool {
	if !device.camera.setIllumination("http://"+device.IPAddress+":"+strconv.Itoa(device.RecordingDataPort)+"/update", ill) {
		return false
	}
	return true
}

func (device *Device) SetCameraExposure(ex float64) bool {
	if !device.camera.setExposure("http://"+device.IPAddress+":"+strconv.Itoa(device.RecordingDataPort)+"/update", ex) {
		return false
	}
	return true
}

func (device *Device) SetCameraFrameRate(fr float64) bool {
	if !device.camera.setFrameRate("http://"+device.IPAddress+":"+strconv.Itoa(device.RecordingDataPort)+"/update", fr) {
		return false
	}
	return true
}

func (device *Device) AutoAdjustCamera() bool {
	if !device.camera.autoAdjust("http://" + device.IPAddress + ":" + strconv.Itoa(device.RecordingDataPort) + "/auto_adjust") {
		return false
	}
	return true
}
