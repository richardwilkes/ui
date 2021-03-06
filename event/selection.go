package event

import (
	"bytes"
	"fmt"
)

// Selection is generated when an item is selected.
type Selection struct {
	target   Target
	finished bool
}

// NewSelection creates a new Selection event. 'target' is the widget that is being acted on.
func NewSelection(target Target) *Selection {
	return &Selection{target: target}
}

// Type returns the event type ID.
func (e *Selection) Type() Type {
	return SelectionType
}

// Target the original target of the event.
func (e *Selection) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *Selection) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *Selection) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *Selection) Finish() {
	e.finished = true
}

// String implements the fmt.Stringer interface.
func (e *Selection) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Selection[Target: %v", e.target))
	if e.finished {
		buffer.WriteString(", Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
