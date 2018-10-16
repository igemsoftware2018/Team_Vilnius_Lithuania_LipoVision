package video_test

import (
	"testing"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device/video"
)

func TestGiveGoodValue(t *testing.T) {
	pump := video.NewPump()

	err := pump.Invoke(device.PumpReset, 0)
	if err != nil {
		t.Error("A mock pump returned an error")
	}
}

func TestGiveBadValue(t *testing.T) {
	pump := video.NewPump()

	err := pump.Invoke(device.CameraSetFrameRate, 23124)
	if err != nil {
		t.Error("A mock pump returned an error")
	}
}
