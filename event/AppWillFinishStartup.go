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

// AppWillFinishStartup is generated immediately prior to the application finishing its startup
// sequence, before it has been asked to open any files.
type AppWillFinishStartup struct {
	target   Target
	finished bool
}

// NewAppWillFinishStartup creates a new AppWillFinishStartup event. 'target' is the app.
func NewAppWillFinishStartup(target Target) *AppWillFinishStartup {
	return &AppWillFinishStartup{target: target}
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
