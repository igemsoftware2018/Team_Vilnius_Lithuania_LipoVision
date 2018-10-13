package dropletgenomics_test

import (
	"context"
	"image"
	"image/color"
	"testing"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device/dropletgenomics"
)

type testImage struct {
}

func (testImage) ColorModel() color.Model {
	return nil
}

func (testImage) Bounds() image.Rectangle {
	return image.Rectangle{}
}

func (testImage) At(x, y int) color.Color {
	return nil
}

func TestSkip(t *testing.T) {
	ctx := context.Background()
	frame := dropletgenomics.CreateFrame(ctx, testImage{})

	select {
	case <-frame.Skip():
		t.Error("frame skipped before needed")
	default:
		break
	}

	frame.Frame()
	select {
	case <-frame.Skip():
		return
	default:
		t.Error("frame lifetime did not end")
		break
	}
}
