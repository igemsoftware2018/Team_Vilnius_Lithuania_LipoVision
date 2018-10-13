package gui

import (
	"fmt"

	"github.com/andlabs/ui"
)

const (
	// MainWindowWidth sets the default width
	MainWindowWidth int = 250

	// MainWindowHeight sets the default height
	MainWindowHeight int = 200
)

// NewMainWindow sets up the main window for ui lib
func NewMainWindow() {
	mainWin := MainWindow{
		window:           ui.NewWindow("LipoVision", MainWindowWidth, MainWindowHeight, false),
		grid:             ui.NewGrid(),
		optionsSeparator: ui.NewHorizontalSeparator(),
		statusSeparator:  ui.NewHorizontalSeparator(),
	}

	mainWin.window.SetChild(mainWin.grid)
	mainWin.window.SetMargined(true)

	mainWin.compose()

	mainWin.window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainWin.window.Destroy()
		return true
	})

	mainWin.window.Show()
}

type pumpForm struct {
	entry *ui.Entry
}

// MainWindow is the main widgeted window for lipovision
type MainWindow struct {
	window *ui.Window
	grid   *ui.Grid

	// Options bar
	optionsBox       *ui.Box
	label            *ui.Label
	combo            *ui.Combobox
	optionsSeparator *ui.Separator

	// Device status
	statusBox       *ui.Box
	statusSeparator *ui.Separator

	// Pump forms
	form      *ui.Form
	pumpForms []pumpForm
}

func (w *MainWindow) composeOptionsBox() {
	w.optionsBox = ui.NewHorizontalBox()

	// Combobox label
	w.label = ui.NewLabel("Device: ")

	// Device selection combobox
	w.combo = ui.NewCombobox()
	w.combo.Append("dropletgenomics")
	w.combo.Append("video")

	w.optionsBox.Append(w.label, false)
	w.optionsBox.Append(w.combo, true)

	w.grid.Append(w.optionsBox, 0, 0, MainWindowWidth,
		20, true, ui.AlignStart, false, ui.AlignStart)
}

func (w *MainWindow) compose() {
	w.composeOptionsBox()

	w.grid.InsertAt(
		w.optionsSeparator,
		w.optionsBox, ui.Bottom, MainWindowWidth,
		1, true, ui.AlignFill, false, ui.AlignFill)

	w.composeStatusForm()

	w.grid.InsertAt(
		w.statusSeparator,
		w.statusBox, ui.Bottom, MainWindowWidth,
		1, true, ui.AlignFill, false, ui.AlignFill)

	w.composePumpForm()
}

func (w *MainWindow) composeStatusForm() {
	w.statusBox = ui.NewHorizontalBox()

	w.statusBox.Append(ui.NewLabel("Device status: "), true)

	w.grid.InsertAt(w.statusBox, w.optionsSeparator, ui.Bottom, MainWindowWidth,
		30, true, ui.AlignCenter, false, ui.AlignFill)
}

func (w *MainWindow) composePumpForm() {
	const units string = "uL/min"

	w.form = ui.NewForm()
	for i := 0; i < 4; i++ {
		formBox := ui.NewHorizontalBox()
		spinbox := ui.NewSpinbox(-9000, 9000)
		formBox.Append(spinbox, true)
		formBox.Append(ui.NewLabel(units), false)

		name := fmt.Sprintf("Pump %d: ", i+1)
		w.form.Append(fmt.Sprintf(name), formBox, true)
	}
	w.grid.InsertAt(w.form, w.statusSeparator, ui.Bottom, MainWindowWidth,
		30, true, ui.AlignCenter, false, ui.AlignFill)
}
