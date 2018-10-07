package main

import (
	"context"
	"fmt"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device/dropletgenomics"
	"gocv.io/x/gocv"
)

func main() {
	dropletDevice := dropletgenomics.Create()
	originalWindow := gocv.NewWindow("Stream")
	ctx, cancel := context.WithCancel(context.Background())
	stream := dropletDevice.Stream(ctx)

	for {
		select {
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
}
