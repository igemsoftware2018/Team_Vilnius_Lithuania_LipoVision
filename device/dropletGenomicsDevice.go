package device

import (
	"net/http"
	"strconv"
)

// DropletGenomicsDevice implements default Device interface
type DropletGenomicsDevice struct {
	IPAddress         string
	HTTPPort          int
	PumpDataPort      int
	RecordingDataPort int
	PumpExperiment    int
	Pumps             []pump
	camera            camera
}

func (device *DropletGenomicsDevice) GetVolumeTarget(pump int) float64 {
	return device.Pumps[pump].VolumeTarget
}
func (device *DropletGenomicsDevice) GetPurgeRate(pump int) float64 {
	return device.Pumps[pump].PurgeRate
}
func (device *DropletGenomicsDevice) GetPumpID(pump int) float64 {
	return device.Pumps[pump].PumpID
}
func (device *DropletGenomicsDevice) GetRateW(pump int) float64 {
	return device.Pumps[pump].PumpID
}
func (device *DropletGenomicsDevice) GetVolume(pump int) float64 {
	return device.Pumps[pump].Volume
}
func (device *DropletGenomicsDevice) GetStatus(pump int) bool {
	return device.Pumps[pump].Status
}
func (device *DropletGenomicsDevice) GetName(pump int) string {
	return device.Pumps[pump].Name
}
func (device *DropletGenomicsDevice) GetDirection(pump int) bool {
	return device.Pumps[pump].Direction
}
func (device *DropletGenomicsDevice) GetSyringe(pump int) float64 {
	return device.Pumps[pump].Syringe
}
func (device *DropletGenomicsDevice) GetUsed(pump int) bool {
	return device.Pumps[pump].Used
}
func (device *DropletGenomicsDevice) GetVolumeTargetW(pump int) float64 {
	return device.Pumps[pump].VolumeTargetW
}
func (device *DropletGenomicsDevice) GetVolumeW(pump int) float64 {
	return device.Pumps[pump].VolumeW
}
func (device *DropletGenomicsDevice) GetRate(pump int) float64 {
	return device.Pumps[pump].Rate
}
func (device *DropletGenomicsDevice) GetStalled(pump int) bool {
	return device.Pumps[pump].Stalled
}
func (device *DropletGenomicsDevice) GetForce(pump int) float64 {
	return device.Pumps[pump].Force
}

func (device *DropletGenomicsDevice) DefinePumpExperiment(numberOfPumps int) {
	//TODO : get endpoints from device in GMC
}

func (device *DropletGenomicsDevice) EstablishPumpsWithID() {
	device.Pumps = make([]pump, device.PumpExperiment, device.PumpExperiment)
	for i := 0; i < device.PumpExperiment; i++ {
		device.Pumps[i] = NewPump(i)
	}
}

func (device *DropletGenomicsDevice) Update() bool {
	for i := 0; i < len(device.Pumps); i++ {
		if !device.Pumps[i].updatePumpValues("http://" + device.IPAddress + ":" + strconv.Itoa(device.PumpDataPort) + "/refresh") {
			return false
		}
	}
	return true
}

func (device *DropletGenomicsDevice) Reset(pump int) bool {
	if !device.Pumps[pump].resetPump("http://" + device.IPAddress + ":" + strconv.Itoa(device.PumpDataPort) + "/update_pars") {
		return false
	}
	return true
}

func (device *DropletGenomicsDevice) ToggleWithdrawInfuse(pump int, widthdraw bool) bool {
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

func (device *DropletGenomicsDevice) DisplayPumpValues(selectedPump int) (answer string) {
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

func (device *DropletGenomicsDevice) TogglePump(selectedPump int, start bool) bool {
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

func (device *DropletGenomicsDevice) SetPumpVolume(selectedPump int, volume int) bool {
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

func (device *DropletGenomicsDevice) SetPumpTargetVolume(selectedPump int, volume int) bool {
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

func (device *DropletGenomicsDevice) SetCameraIllumination(ill float64) bool {
	if !device.camera.setIllumination("http://"+device.IPAddress+":"+strconv.Itoa(device.RecordingDataPort)+"/update", ill) {
		return false
	}
	return true
}

func (device *DropletGenomicsDevice) SetCameraExposure(ex float64) bool {
	if !device.camera.setExposure("http://"+device.IPAddress+":"+strconv.Itoa(device.RecordingDataPort)+"/update", ex) {
		return false
	}
	return true
}

func (device *DropletGenomicsDevice) SetCameraFrameRate(fr float64) bool {
	if !device.camera.setFrameRate("http://"+device.IPAddress+":"+strconv.Itoa(device.RecordingDataPort)+"/update", fr) {
		return false
	}
	return true
}

func (device *DropletGenomicsDevice) AutoAdjustCamera() bool {
	if !device.camera.autoAdjust("http://" + device.IPAddress + ":" + strconv.Itoa(device.RecordingDataPort) + "/auto_adjust") {
		return false
	}
	return true
}

// Available determines whether connection to
// DropletGenomics device is established
func (device *DropletGenomicsDevice) Available() bool {
	url := setupDeviceURL(device)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return false
	}
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return false
	}

	defer res.Body.Close()

	if res.StatusCode > 299 || res.StatusCode < 200 {
		return false
	}

	return true
}
