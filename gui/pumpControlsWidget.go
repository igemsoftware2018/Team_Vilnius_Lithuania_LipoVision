package gui

import "github.com/gotk3/gotk3/gtk"

// NewPumpControlsWidget returns pump control collection
func NewPumpControlsWidget() (*PumpControlsWidget, error) {
	stack, stackErr := gtk.StackNew()
	if stackErr != nil {
		return nil, stackErr
	}

	return &PumpControlsWidget{rootStack: stack}, nil
}

// PumpControlsWidget is the pump control collection
type PumpControlsWidget struct {
	Widget

	rootStack *gtk.Stack
}

// Widget returns the root stack component
func (pc *PumpControlsWidget) Root() gtk.IWidget {
	return pc.rootStack
}
