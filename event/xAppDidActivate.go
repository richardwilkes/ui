package event

import (
	"bytes"
)

// AppDidActivate is generated immediately after the application has been brought to the
// foreground.
type AppDidActivate struct {
	target   Target
	finished bool
}

// SendAppDidActivate sends a new AppDidActivate event.
func SendAppDidActivate() {
	Dispatch(&AppDidActivate{target: GlobalTarget()})
}

// Type returns the event type ID.
func (e *AppDidActivate) Type() Type {
	return AppDidActivateType
}

// Target the original target of the event.
func (e *AppDidActivate) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *AppDidActivate) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *AppDidActivate) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *AppDidActivate) Finish() {
	e.finished = true
}

// String implements the fmt.Stringer interface.
func (e *AppDidActivate) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("AppDidActivate[")
	if e.finished {
		buffer.WriteString("Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
