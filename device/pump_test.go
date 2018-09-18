package device

import (
	"fmt"
	"testing"
)

func TestUpdatePumpValues(t *testing.T) {

	var dummyPump pump
	dummyPump.updatePumpValues("http://localhost:5000/refresh")
	fmt.Printf("Pump purge rate is: %v ", dummyPump.PurgeRate)
}

func TestTogglePump(t *testing.T) {
	var dummyPump pump
	if !dummyPump.togglePump("http://localhost:5000/refresh", true) {
		t.Errorf("%v", "Start not worked")
	}
}
