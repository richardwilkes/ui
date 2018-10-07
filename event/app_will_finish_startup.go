package event

import (
	"bytes"
)

// AppWillFinishStartup is generated immediately prior to the application finishing its startup
// sequence, before it has been asked to open any files.
type AppWillFinishStartup struct {
	target   Target
	finished bool
}

// SendAppWillFinishStartup sends a new AppWillFinishStartup event.
func SendAppWillFinishStartup() {
	Dispatch(&AppWillFinishStartup{target: GlobalTarget()})
}

// Type returns the event type ID.
func (e *AppWillFinishStartup) Type() Type {
	return AppWillFinishStartupType
}

// Target the original target of the event.
func (e *AppWillFinishStartup) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *AppWillFinishStartup) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *AppWillFinishStartup) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *AppWillFinishStartup) Finish() {
	e.finished = true
}

// String implements the fmt.Stringer interface.
func (e *AppWillFinishStartup) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("AppWillFinishStartup[")
	if e.finished {
		buffer.WriteString("Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
