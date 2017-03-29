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

// AppQuitRequested is generated to ask if it is OK to quit the application.
type AppQuitRequested struct {
	target   Target
	canceled bool
	delayed  bool
	finished bool
}

// NewAppQuitRequested creates a new AppQuitRequested event. 'target' is the app.
func NewAppQuitRequested(target Target) *AppQuitRequested {
	return &AppQuitRequested{target: target}
}

// Type returns the event type ID.
func (e *AppQuitRequested) Type() Type {
	return AppQuitRequestedType
}

// Target the original target of the event.
func (e *AppQuitRequested) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *AppQuitRequested) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *AppQuitRequested) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *AppQuitRequested) Finish() {
	e.finished = true
}

// Canceled returns true if the request has been canceled.
func (e *AppQuitRequested) Canceled() bool {
	return e.canceled
}

// Cancel the request.
func (e *AppQuitRequested) Cancel() {
	e.canceled = true
	e.finished = true
}

// Delayed returns true if the request should only proceed after a call to app.MayTerminateNow().
func (e *AppQuitRequested) Delayed() bool {
	return e.delayed
}

// Delay the request.
func (e *AppQuitRequested) Delay() {
	e.delayed = true
	e.finished = true
}

// String implements the fmt.Stringer interface.
func (e *AppQuitRequested) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("AppQuitRequested[")
	needComma := false
	if e.canceled {
		buffer.WriteString("Canceled")
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
		}
		buffer.WriteString("Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
