//go:generate go-bindata template-intersection.png
package main

import (
	"bytes"
	"context"
	"flag"
	"image"
	"image/png"
	"os"
	"runtime/pprof"
	"time"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device"
	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device/dropletgenomics"
	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device/video"
	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/gui"
	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/processor"
	"github.com/gotk3/gotk3/gtk"
	log "github.com/sirupsen/logrus"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

var (
	mainCtx         context.Context
	mainCancel      context.CancelFunc
	activeDevice    device.Device
	deviceSet       bool
	activeProcessor *processor.FrameProcessor
)

var (
	illuminationValue float64
	exposureValue     float64
)

func getTemplateImage() image.Image {
	imgBytes := MustAsset("template-intersection.png")
	img, imgErr := png.Decode(bytes.NewBuffer(imgBytes))
	if imgErr != nil {
		panic(imgErr)
	}

	return img
}

func chooseFileCreateDevice(win *gtk.Window) device.Device {
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
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	mainCtx, mainCancel = context.WithCancel(context.Background())
	defer mainCancel()

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window")
	}
	win.SetTitle("LipoVision")
	win.Connect("destroy", func() {
		mainCancel()
		gtk.MainQuit()
	})
	win.SetDefaultSize(850, 550)

	content, err := gui.NewMainControl()
	if err != nil {
		panic(err)
	}
	win.Add(content.Root())
	win.ShowAll()

	registerDeviceChange(content, win)

	gtk.Main()
}

func registerEventHandling(content *gui.MainControl, win *gtk.Window) {
	content.Camera.AutoAdjButton.Connect("clicked", func() {
		activeDevice.Camera().Invoke(device.CameraAutoAdjust, 0)
	})
	content.Camera.IlluminationScale.Connect("format-value", func(scale *gtk.Scale) {
		value := scale.GetValue()
		if illuminationValue != value {
			activeDevice.Camera().Invoke(device.CameraSetIllumination, value)
			illuminationValue = value
		}
	})
	content.Camera.ExposureScale.Connect("format-value", func(scale *gtk.Scale) {
		value := scale.GetValue()
		if exposureValue != value {
			activeDevice.Camera().Invoke(device.CameraSetExposure, value)
			exposureValue = value
		}
	})
	content.StreamControl.LockButton.Connect("toggled", func(btn *gtk.CheckButton) {
		if btn.GetActive() {
			activeProcessor.Set(processor.SettingRegionIsSet, int32(1))
		} else {
			activeProcessor.Set(processor.SettingRegionIsSet, int32(0))
		}
	})
	content.StreamControl.AutoButton.Connect("toggled", func(btn *gtk.CheckButton) {
		if btn.GetActive() {
			activeProcessor.Set(processor.SettingAutonomicRun, int32(1))
		} else {
			activeProcessor.Set(processor.SettingAutonomicRun, int32(0))
		}
	})
	content.Pump.Pump(0).Connect("value-changed", func(btn *gtk.SpinButton) {
		val := btn.GetValue()
		activeDevice.Pump(0).Invoke(device.PumpSetVolume, float64(val))
	})
	content.Pump.Pump(1).Connect("value-changed", func(btn *gtk.SpinButton) {
		val := btn.GetValue()
		activeDevice.Pump(1).Invoke(device.PumpSetVolume, float64(val))
	})
	content.Pump.Pump(2).Connect("value-changed", func(btn *gtk.SpinButton) {
		val := btn.GetValue()
		activeDevice.Pump(2).Invoke(device.PumpSetVolume, float64(val))
	})
	content.Pump.Pump(3).Connect("value-changed", func(btn *gtk.SpinButton) {
		val := btn.GetValue()
		activeDevice.Pump(3).Invoke(device.PumpSetVolume, float64(val))
	})
}

func registerDeviceChange(content *gui.MainControl, win *gtk.Window) {
	content.StreamControl.ComboBox.Connect("changed", func(combo *gtk.ComboBoxText) {
		mainCancel()
		mainCtx, mainCancel = context.WithCancel(context.Background())
		selection := combo.GetActiveText()
		switch selection {
		case "Video file...":
			activeDevice = chooseFileCreateDevice(win)
		case "DropletGenomics":
			activeDevice = dropletgenomics.Create(4)
		default:
			errDialog := gtk.MessageDialogNew(win, gtk.DIALOG_MODAL,
				gtk.MESSAGE_ERROR, gtk.BUTTONS_OK,
				"Chosen device %s, does not exist", selection)
			errDialog.Run()
		}

		frameHandlers := make(map[string]func(image.Image))
		frameHandlers[processor.StreamOriginal] = func(frame image.Image) {
			if err := content.StreamControl.ShowFrame(frame); err != nil {
				log.Error("Failed to show frame on main window: ", err)
			}
		}
		frameHandlers[processor.StreamRegion] = func(frame image.Image) {
			if err := content.RegionStream.ShowFrame(frame); err != nil {
				log.Error("Failed to show frame on reference: ", err)
			}
		}

		stream := activeDevice.Stream(mainCtx)
		activeProcessor = processor.NewFrameProcessor(getTemplateImage())
		activeProcessor.Launch(stream, frameHandlers)

		go func() {
			log.Info("Health monitor started")
		CheckHealth:
			for {
				time.Sleep(1 * time.Second)
				select {
				case <-mainCtx.Done():
					break CheckHealth
				default:
				}

				if activeProcessor.Get(processor.SettingAutonomicRun) != 0 {
					score := activeProcessor.Get(processor.SettingDangerScore)
					if score > 30 {
						pump2 := content.Pump.Pump(1).GetValue()
						pump3 := content.Pump.Pump(2).GetValue()

						content.Pump.Pump(1).SetValue(pump2 + 5)
						content.Pump.Pump(2).SetValue(pump3 + 5)
					}
				}
			}
			log.Info("Health monitor finished")
		}()

		if !deviceSet {
			registerEventHandling(content, win)
		}
	})
}
