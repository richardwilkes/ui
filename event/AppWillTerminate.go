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

// AppWillTerminate is generated immediately prior to the application terminating.
type AppWillTerminate struct {
	target   Target
	finished bool
}

// NewAppWillTerminate creates a new AppWillTerminate event. 'target' is the app.
func NewAppWillTerminate(target Target) *AppWillTerminate {
	return &AppWillTerminate{target: target}
}

// Type returns the event type ID.
func (e *AppWillTerminate) Type() Type {
	return AppWillTerminateType
}

// Target the original target of the event.
func (e *AppWillTerminate) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *AppWillTerminate) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *AppWillTerminate) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *AppWillTerminate) Finish() {
	e.finished = true
}

// String implements the fmt.Stringer interface.
func (e *AppWillTerminate) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("AppWillTerminate[")
	if e.finished {
		buffer.WriteString("Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
