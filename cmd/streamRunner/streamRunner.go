package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device/dropletgenomics"
	log "github.com/sirupsen/logrus"
	"gocv.io/x/gocv"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

func main() {
	dropletDevice := dropletgenomics.Create(4)
	originalWindow := gocv.NewWindow("Stream")
	ctx, cancel := context.WithCancel(context.Background())
	stream := dropletDevice.Stream(ctx)

	signalBuffer := make(chan os.Signal, 1)
	signal.Notify(signalBuffer, os.Interrupt)
	go func() {
		for sig := range signalBuffer {
			fmt.Printf("signal: %s\n", sig)
			cancel()
			break
		}
	}()

Processing:
	for {
		select {
		case <-ctx.Done():
			break Processing
		case frame := <-stream:
			mat, err := gocv.ImageToMatRGB(frame.Frame())
			if err != nil {
				fmt.Println(err)
			}
			originalWindow.IMShow(mat)
		default:
			if originalWindow.WaitKey(1)&0xFF == 'c' {
				cancel()
				break
			}
		}
	}

	originalWindow.Close()
}
