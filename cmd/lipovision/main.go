package main

import (
	"flag"
	"strings"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/gui"
	"github.com/gotk3/gotk3/gtk"
	log "github.com/sirupsen/logrus"
)

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

	content, err := gui.NewMainWidget()
	if err != nil {
		panic(err)
	}
	win.Add(content.Root())

	win.ShowAll()
	gtk.Main()
}
