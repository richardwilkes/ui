package event

import (
	"bytes"
)

// AppLastWindowClosed is generated when the last window has been closed. It is sent to determine
// if the app should be remain open or not.
type AppLastWindowClosed struct {
	target     Target
	remainOpen bool
	finished   bool
}

// SendAppLastWindowClosed sends a new AppLastWindowClosed event.
func SendAppLastWindowClosed() *AppLastWindowClosed {
	evt := &AppLastWindowClosed{target: GlobalTarget()}
	Dispatch(evt)
	return evt
}

// Type returns the event type ID.
func (e *AppLastWindowClosed) Type() Type {
	return AppLastWindowClosedType
}

// Target the original target of the event.
func (e *AppLastWindowClosed) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *AppLastWindowClosed) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *AppLastWindowClosed) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *AppLastWindowClosed) Finish() {
	e.finished = true
}

// Quit returns true if the app should quit.
func (e *AppLastWindowClosed) Quit() bool {
	return !e.remainOpen
}

// RemainOpen marks the app for remaining open and not terminating due to the last window closing.
func (e *AppLastWindowClosed) RemainOpen() {
	e.remainOpen = true
	e.finished = true
}

// String implements the fmt.Stringer interface.
func (e *AppLastWindowClosed) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("AppLastWindowClosed[")
	if e.remainOpen {
		buffer.WriteString("Remain Open")
	} else {
		buffer.WriteString("Quit")
	}
	if e.finished {
		buffer.WriteString(", Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
