package gui

import (
	"bytes"
	"image"
	"image/png"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func init() {
	gtk.Init(nil)
}

var encoder png.Encoder = png.Encoder{
	CompressionLevel: png.NoCompression,
}

// Control is the interface for collections of widgets
type Control interface {
	// Root returns gotk3 type of widget
	// This widget is root for all represented components of this widget
	Root() gtk.IWidget
}

func showFrame(cnt *gtk.Image, frame image.Image) error {
	buffer := new(bytes.Buffer)
	encoder.Encode(buffer, frame)

	loader, err := gdk.PixbufLoaderNew()
	if err != nil {
		return err
	}

	pixbuf, loadErr := loader.WriteAndReturnPixbuf(buffer.Bytes())
	if loadErr != nil {
		return loadErr
	}

	cnt.SetFromPixbuf(pixbuf)
	return nil
}
