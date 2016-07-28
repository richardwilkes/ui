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

// AppWillDeactivate is generated immediately prior to the application being sent to the
// background.
type AppWillDeactivate struct {
	target   Target
	finished bool
}

// NewAppWillDeactivate creates a new AppWillDeactivate event. 'target' is the app.
func NewAppWillDeactivate(target Target) *AppWillDeactivate {
	return &AppWillDeactivate{target: target}
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
