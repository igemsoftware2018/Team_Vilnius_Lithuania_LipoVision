package device

import (
	"fmt"
	"net/http"
	"strconv"
)

// DropletGenomicsDevice implements default Device interface
type DropletGenomicsDevice struct {
	IPAddress         string
	HTTPPort          int
	PumpDataPort      int
	RecondingDataPort int
	PumpExperiment    int
	Pumps             []pump
}

func (device *DropletGenomicsDevice) DefinePumpExperiment(numberOfPumps int) {
	//TODO : get endpoints from device in GMC
}

func (device *DropletGenomicsDevice) EstablishPumpsWithId() {
	device.Pumps = make([]pump, device.PumpExperiment, device.PumpExperiment)
	for i := 0; i < device.PumpExperiment; i++ {
		device.Pumps[i].PumpID = float64(i)
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

func (device *DropletGenomicsDevice) GetPumpValues(selectedPump int) (answer string) {
	if selectedPump == -1 {
		for _, pump := range device.Pumps {
			if !pump.updatePumpValues(device.setupGetValuesURL()) {
				panic("ERROR")
			}
			answer += formatPumpValues(&pump)
		}
		return
	} else {
		answer += formatPumpValues(&device.Pumps[selectedPump])
		return
	}
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

/* Helper functions */
func setupDeviceURL(device *DropletGenomicsDevice) string {
	return fmt.Sprintf("http://%v:%v", device.IPAddress, strconv.Itoa(device.HTTPPort))
}
func (device *DropletGenomicsDevice) setupGetValuesURL() string {
	return fmt.Sprintf("http://%v:%v/refresh", device.IPAddress, strconv.Itoa(device.PumpDataPort))
}
func formatPumpValues(p *pump) string {
	return fmt.Sprintf("Pump: %.0f\n\tVolume Target: %.0f\n\tPurge Rate: %.0f\n\tRateW: %.0f\n\tVolume: %.0f\n\tStatus: %v\n\tName: %s\n\tDirection: %v\n\tSyringe: %.2f\n\tUsed: %v\n\tVolumeTargetW: %.0f\n\tVolumeW: %.0f\n\tRate: %.0f\n\tStalled: %v\n\tForce: %.0f\n", p.PumpID, p.VolumeTarget, p.PurgeRate, p.RateW, p.Volume, p.Status, p.Name, p.Direction, p.Syringe, p.Used, p.VolumeTargetW, p.VolumeW, p.Rate, p.Stalled, p.Force)
}
