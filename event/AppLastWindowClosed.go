// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

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

// NewAppLastWindowClosed creates a new AppLastWindowClosed event. 'target' is the app.
func NewAppLastWindowClosed(target Target) *AppLastWindowClosed {
	return &AppLastWindowClosed{target: target}
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
