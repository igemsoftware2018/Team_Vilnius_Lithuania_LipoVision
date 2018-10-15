package main

import (
	"flag"
	"strings"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device/dropletgenomics"
	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device/video"
	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/gui"
	"github.com/gotk3/gotk3/gtk"
	log "github.com/sirupsen/logrus"
)

func newDeviceWithChosenVideo(win *gtk.Window) device.Device {
	chooser, err := gtk.FileChooserDialogNewWith1Button(
		"Select video file", win, gtk.FILE_CHOOSER_ACTION_OPEN,
		"Open", gtk.RESPONSE_ACCEPT)
	if err != nil {
		log.Fatal("File chooser failed: ", err)
	}
	defer chooser.Destroy()

	filter, _ := gtk.FileFilterNew()
	filter.AddPattern("*.mp4")
	filter.SetName(".mp4")
	chooser.AddFilter(filter)

	resp := chooser.Run()
	log.Info(resp)

	videoFile := chooser.GetFilename()
	return video.Create(videoFile, 24)
}

func main() {
	deviceRequested := flag.String("device", "",
		`Specify a device to run the program with. Valid values:
	 	 > dropletgenomics
		  > video`)
	flag.Parse()

	if strings.Compare("", *deviceRequested) != 0 {
		log.Info("selected device: ", *deviceRequested)
	}

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window")
	}
	win.SetTitle("LipoVision")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.SetDefaultSize(890, 500)

	var device device.Device
	content, err := gui.NewMainControl(&device)
	if err != nil {
		panic(err)
	}
	win.Add(content.Root())
	win.ShowAll()

	content.StreamControl.ComboBox.Connect("changed", func(combo *gtk.ComboBoxText) {
		selection := combo.GetActiveText()
		switch selection {
		case "Video file...":
			device = newDeviceWithChosenVideo(win)
		case "DropletGenomics":
			device = dropletgenomics.Create(4)
		default:
			errDialog := gtk.MessageDialogNew(win, gtk.DIALOG_MODAL,
				gtk.MESSAGE_ERROR, gtk.BUTTONS_OK,
				"Chosen device %s, does not exist", selection)
			errDialog.Run()
		}
	})

	gtk.Main()
}
