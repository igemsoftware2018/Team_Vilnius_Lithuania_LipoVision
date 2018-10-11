// +build !windows

package window

import (
	"bytes"
	"context"
	"image/png"

	"github.com/Vilnius-Lithuania-iGEM-2018/lipovision/device/video"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	log "github.com/sirupsen/logrus"
)

// init gtk
func init() {
	gtk.Init(nil)
}

// NewMain creates the main application window
func NewMain() (*Main, error) {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return nil, err
	}

	win.SetTitle("LipoVision")
	win.SetDefaultSize(800, 600)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	gtkImage, gtkErr := gtk.ImageNew()
	if gtkErr != nil {
		panic(gtkErr)
	}

	win.Add(gtkImage)

	return &Main{
		window: win,
		image:  gtkImage,
		events: make(map[string]chan Event),
	}, nil
}

// Main is the main application window
type Main struct {
	Window

	events map[string]chan Event
	window *gtk.Window
	image  *gtk.Image
}

// Run starts main loop, blocks
func (w *Main) Run() {
	w.window.ShowAll()
	videoDev := video.Create("video_1527101251.mp4", 20)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream := videoDev.Stream(ctx)

	go func() {
	Process:
		for {
			select {
			case <-ctx.Done():
				break Process
			case frame := <-stream:
				frameImg := frame.Frame()

				writer := new(bytes.Buffer)
				png.Encode(writer, frameImg)

				loader, loaderErr := gdk.PixbufLoaderNew()
				if loaderErr != nil {
					panic(loaderErr)
				}
				pixbuf, pixErr := loader.WriteAndReturnPixbuf(writer.Bytes())
				if pixErr != nil {
					log.Warn("failed to get pixbuf")
					continue
				}

				w.image.SetFromPixbuf(pixbuf)
			}
		}
	}()

	gtk.Main()
}

// Subscribe provides a new channel of events
func (w *Main) Subscribe(eventName string) <-chan Event {
	newEvent := make(chan Event, EventQueueLength)
	w.events[eventName] = newEvent
	return newEvent
}
