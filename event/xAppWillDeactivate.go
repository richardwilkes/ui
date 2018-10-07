package event

import (
	"bytes"
)

// AppWillDeactivate is generated immediately prior to the application being sent to the
// background.
type AppWillDeactivate struct {
	target   Target
	finished bool
}

// SendAppWillDeactivate sends a new AppWillDeactivate event.
func SendAppWillDeactivate() {
	Dispatch(&AppWillDeactivate{target: GlobalTarget()})
}

// Type returns the event type ID.
func (e *AppWillDeactivate) Type() Type {
	return AppWillDeactivateType
}

// Target the original target of the event.
func (e *AppWillDeactivate) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *AppWillDeactivate) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *AppWillDeactivate) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *AppWillDeactivate) Finish() {
	e.finished = true
}

// String implements the fmt.Stringer interface.
func (e *AppWillDeactivate) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("AppWillDeactivate[")
	if e.finished {
		buffer.WriteString("Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
