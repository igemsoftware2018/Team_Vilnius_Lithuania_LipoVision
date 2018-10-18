package video_test

import (
	"testing"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device/video"
)

func TestSetsFrameRate(t *testing.T) {
	expectedVal := 10.0
	camera := video.Camera{FrameRate: 20.0}

	if err := camera.Invoke(device.CameraSetFrameRate, expectedVal); err != nil {
		t.Error("invoke failed")
	}

	if camera.FrameRate != expectedVal {
		t.Error("returned value does not match")
	}
}

func TestUnsupported(t *testing.T) {
	camera := video.Camera{FrameRate: 20}
	if err := camera.Invoke(device.CameraSetExposure, 10); err == nil {
		t.Error("invoke passed when it shouldn't have")
	}

}
