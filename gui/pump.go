package gui

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

// NewPumpAndRegionContainer creates a box container for
// both RegionControl and PumpControl
func NewPumpAndRegionContainer() (gtk.IWidget, *PumpControl, *CameraControl, *RegionControl, error) {
	box, boxErr := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	if boxErr != nil {
		return nil, nil, nil, nil, boxErr
	}

	pumpControl, pumpControlErr := NewPumpControl()
	if pumpControlErr != nil {
		return nil, nil, nil, nil, pumpControlErr
	}
	box.PackStart(pumpControl.Root(), false, false, 0)

	cameraControl, cameraControlErr := NewCameraConrol()
	if cameraControlErr != nil {
		return nil, nil, nil, nil, cameraControlErr
	}
	box.PackStart(cameraControl.Root(), false, false, 0)

	region, regionErr := NewRegionControl()
	if regionErr != nil {
		return nil, nil, nil, nil, regionErr
	}
	box.PackEnd(region.Root(), true, true, 0)

	return box, pumpControl, cameraControl, region, nil
}

// NewPumpControl returns pump control collection
func NewPumpControl() (*PumpControl, error) {
	box, boxErr := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	if boxErr != nil {
		return nil, boxErr
	}

	frame, frameErr := gtk.FrameNew("Pump controls")
	if frameErr != nil {
		return nil, frameErr
	}

	var pumps []*gtk.SpinButton
	for i := 0; i < 4; i++ {
		pumpBox, pump, err := newPumpItem(i + 1)
		if err != nil {
			return nil, err
		}
		pumps = append(pumps, pump)
		box.PackStart(pumpBox, false, false, 5)
	}

	frame.Add(box)

	return &PumpControl{rootBox: frame, pumps: pumps}, nil
}

func newPumpItem(index int) (gtk.IWidget, *gtk.SpinButton, error) {
	box, boxErr := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if boxErr != nil {
		return nil, nil, boxErr
	}

	label, labelErr := gtk.LabelNew(fmt.Sprintf("Pump %d: ", index))
	if labelErr != nil {
		return nil, nil, labelErr
	}

	spinbox, spinboxErr := gtk.SpinButtonNewWithRange(-9000, 9000, 5)
	if spinboxErr != nil {
		return nil, nil, spinboxErr
	}
	spinbox.SetValue(0)

	box.PackStart(label, false, false, 0)
	box.PackStart(spinbox, true, true, 0)

	return box, spinbox, nil
}

// PumpControl is the pump control collection
type PumpControl struct {
	Control

	rootBox *gtk.Frame
	pumps   []*gtk.SpinButton
}

// Root returns the root stack component
func (pc *PumpControl) Root() gtk.IWidget {
	return pc.rootBox
}

// Pump returns pump by id
func (pc *PumpControl) Pump(index int) *gtk.SpinButton {
	return pc.pumps[index]
}
