package event

import (
	"bytes"
)

// AppDidDeactivate is generated immediately after the application has been sent to the
// background.
type AppDidDeactivate struct {
	target   Target
	finished bool
}

// SendAppDidDeactivate sends a new AppDidDeactivate event.
func SendAppDidDeactivate() {
	Dispatch(&AppDidDeactivate{target: GlobalTarget()})
}

// Type returns the event type ID.
func (e *AppDidDeactivate) Type() Type {
	return AppDidDeactivateType
}

// Target the original target of the event.
func (e *AppDidDeactivate) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *AppDidDeactivate) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *AppDidDeactivate) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *AppDidDeactivate) Finish() {
	e.finished = true
}

// String implements the fmt.Stringer interface.
func (e *AppDidDeactivate) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("AppDidDeactivate[")
	if e.finished {
		buffer.WriteString("Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
