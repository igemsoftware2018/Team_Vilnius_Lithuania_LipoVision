package device

import (
	"testing"
)

func TestDevicePumpValues(t *testing.T) {

	var device = DropletGenomicsDevice{"localhost", 5000, 5000, 5000, 0, nil}
	device.PumpExperiment = 2
	device.Pumps = make([]pump, device.PumpExperiment, device.PumpExperiment)
	device.EstablishPumpsWithId()
	if device.Update() {
		// fmt.Print(device.GetPumpValues(true))
	} else {
		t.Errorf("%v", "Updating pumps failed")
	}
}
