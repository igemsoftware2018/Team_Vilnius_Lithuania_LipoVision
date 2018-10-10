package dropletgenomics_test

import (
	"testing"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device/dropletgenomics"
)

func TestCreation(t *testing.T) {
	device := dropletgenomics.Create(4)

	if device.Camera() == nil {
		t.Error("camera is nil")
	}

	if device.NumPumps() != 4 {
		t.Error("incorrect number of devices")
	}
}
