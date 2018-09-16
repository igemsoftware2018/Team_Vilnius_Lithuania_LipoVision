package device

import (
	"fmt"
	"testing"
)

func TestUpdatePumpValues(t *testing.T) {

	var dummyPump pump
	dummyPump.updatePumpValues("http://localhost:5000/refresh", 0)
	fmt.Printf("Pump purge rate is: %v ", dummyPump.PurgeRate)
}
