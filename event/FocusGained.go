package event

import (
	"bytes"
	"fmt"
)

// FocusGained is generated when a widget gains the keyboard focus.
type FocusGained struct {
	target   Target
	finished bool
}

// NewFocusGained creates a new FocusGained event. 'target' is the widget that is gaining the
// keyboard focus.
func NewFocusGained(target Target) *FocusGained {
	return &FocusGained{target: target}
}

// Type returns the event type ID.
func (e *FocusGained) Type() Type {
	return FocusGainedType
}

// Target the original target of the event.
func (e *FocusGained) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *FocusGained) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *FocusGained) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *FocusGained) Finish() {
	e.finished = true
}

// String implements the fmt.Stringer interface.
func (e *FocusGained) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("FocusGained[Target: %v", e.target))
	if e.finished {
		buffer.WriteString(", Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
