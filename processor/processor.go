package processor

import (
	"image"
)

// Processor is a channel (in this case - frame) processing
// struct and collection of functions
type Processor interface {
	// Launch starts an asynchronious processing goroutine and returns all requested frames of processing
	Launch(<-chan struct{}, map[string]func(image.Image))
	// Set certain parameters about the Processor
	Set(string, interface{}) error
}
