package components

import "github.com/f24-cse535/2pcbyz/pkg/types"

// Dispatcher interface is a global rule that all modules in components should follow.
// each module in components must have a dispatcher that provides start, and input method.
type Dispatcher interface {
	// Start method must be run inside a go-routine, since it blocks the process to
	// receive events from its input channel inside a loop.
	Start()
	// C method must return a channel that the dispatcher reads from inside its loop.
	C() chan *types.Packet
}
