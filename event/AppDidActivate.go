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

// AppDidActivate is generated immediately after the application has been brought to the
// foreground.
type AppDidActivate struct {
	target   Target
	finished bool
}

// NewAppDidActivate creates a new AppDidActivate event. 'target' is the app.
func NewAppDidActivate(target Target) *AppDidActivate {
	return &AppDidActivate{target: target}
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
