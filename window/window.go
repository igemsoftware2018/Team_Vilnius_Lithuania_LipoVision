package window

const EventQueueLength int = 10

// EventQueue hold the queue of events along with it's name
type EventQueue struct {
	Name  string
	Queue <-chan Event
}

// Event describes lipovision Window event
type Event interface {

	// Handle handles the window event
	Handle()
}

// Window describes a lipovision window
type Window interface {

	// Subscribe retuns a channel of window's events
	Subscribe(string) <-chan Event

	// Run starts the window event loop
	Run()
}
