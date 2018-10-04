package dropletgenomics

import (
	"testing"
)

func TestUpdatePumpValues(t *testing.T) {

	var dummyPump pump
	if !dummyPump.updatePumpValues("http://localhost:5000/refresh") {
		t.Errorf("%v", "Updating Pump Values Failed")
	}
}

func TestTogglePump(t *testing.T) {

	var dummyPump pump
	dummyPump.PumpID = 0
	if !dummyPump.togglePump("http://localhost:5000/update", true) {
		t.Errorf("%v", "Toggling not worked")
	}
}
