package video_test

import (
	"testing"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
)

func TestGiveGoodValue(t *testing.T) {
	pump := video.NewPump()

	err := pump.Invoke(device.PumpReset, nil)
	if err != nil {
		t.Fail("A mock pump returned an error")
	}
}

func TestGiveGoodValue(t *testing.T) {
	pump := video.NewPump()

	err := pump.Invoke(device.CameraSetFrameRate, 23124)
	if err != nil {
		t.Fail("A mock pump returned an error")
	}
}
