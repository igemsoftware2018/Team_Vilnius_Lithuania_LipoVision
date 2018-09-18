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

func (device *DropletGenomicsDevice) EstablishPumps() {
	device.Pumps = make([]pump, device.PumpExperiment, device.PumpExperiment)
	for i := 0; i < device.PumpExperiment; i++ {
		device.Pumps[i].PumpID = i
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

func setupDeviceURL(device *DropletGenomicsDevice) string {
	return fmt.Sprintf("http://%v:%v", device.IPAddress, strconv.Itoa(device.HTTPPort))
}
