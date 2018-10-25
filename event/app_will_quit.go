package event

import (
	"bytes"

	"github.com/richardwilkes/toolbox/atexit"
)

// AppWillQuit is generated immediately prior to the application quitting.
type AppWillQuit struct {
	target   Target
	finished bool
}

// SendAppWillQuit sends a new AppWillQuit event.
func SendAppWillQuit() {
	Dispatch(&AppWillQuit{target: GlobalTarget()})
	atexit.Exit(0)
}

// Type returns the event type ID.
func (e *AppWillQuit) Type() Type {
	return AppWillQuitType
}

// Target the original target of the event.
func (e *AppWillQuit) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *AppWillQuit) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *AppWillQuit) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *AppWillQuit) Finish() {
	e.finished = true
}

// String implements the fmt.Stringer interface.
func (e *AppWillQuit) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("AppWillQuit[")
	if e.finished {
		buffer.WriteString("Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
