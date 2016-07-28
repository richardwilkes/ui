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

// AppTerminationRequested is generated to ask if it is OK to terminate the application.
type AppTerminationRequested struct {
	target    Target
	cancelled bool
	delayed   bool
	finished  bool
}

// NewAppTerminationRequested creates a new AppTerminationRequested event. 'target' is the app.
func NewAppTerminationRequested(target Target) *AppTerminationRequested {
	return &AppTerminationRequested{target: target}
}

// Type returns the event type ID.
func (e *AppTerminationRequested) Type() Type {
	return AppTerminationRequestedType
}

// Target the original target of the event.
func (e *AppTerminationRequested) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *AppTerminationRequested) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *AppTerminationRequested) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *AppTerminationRequested) Finish() {
	e.finished = true
}

// Cancelled returns true if the request has been cancelled.
func (e *AppTerminationRequested) Cancelled() bool {
	return e.cancelled
}

// Cancel the request.
func (e *AppTerminationRequested) Cancel() {
	e.cancelled = true
	e.finished = true
}

// Delayed returns true if the request should only proceed after a call to app.MayTerminateNow().
func (e *AppTerminationRequested) Delayed() bool {
	return e.delayed
}

// Delay the request.
func (e *AppTerminationRequested) Delay() {
	e.delayed = true
	e.finished = true
}

// String implements the fmt.Stringer interface.
func (e *AppTerminationRequested) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("AppTerminationRequested[")
	needComma := false
	if e.cancelled {
		buffer.WriteString("Cancelled")
		needComma = true
	}
	if e.delayed {
		if needComma {
			buffer.WriteString(", ")
		} else {
			needComma = true
		}
		buffer.WriteString("Delayed")
	}
	if e.finished {
		if needComma {
			buffer.WriteString(", ")
		} else {
			needComma = true
		}
		buffer.WriteString("Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
