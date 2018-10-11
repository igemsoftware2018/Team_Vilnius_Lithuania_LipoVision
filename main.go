package main

import (
	"flag"
	"strings"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/window"
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

	mainWindow, err := window.NewMain()
	if err != nil {
		panic(err)
	}
	mainWindow.Run()

}
