package dropletgenomics_test

import (
	"strings"
	"testing"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device/dropletgenomics"
)

func TestCreation(t *testing.T) {
	device := dropletgenomics.Create(4)

	if strings.Compare(device.Camera().BaseAddr, "") == 0 {
		t.Error("camera has no address")
	}

	if device.NumPumps() != 4 {
		t.Error("incorrect number of devices")
	}
}
