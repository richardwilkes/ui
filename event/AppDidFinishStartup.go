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

// AppDidFinishStartup is generated immediately after the application has finished its startup
// sequence.
type AppDidFinishStartup struct {
	target   Target
	finished bool
}

// SendAppDidFinishStartup sends a new AppDidFinishStartup event.
func SendAppDidFinishStartup() {
	Dispatch(&AppDidFinishStartup{target: GlobalTarget()})
}

// Type returns the event type ID.
func (e *AppDidFinishStartup) Type() Type {
	return AppDidFinishStartupType
}

// Target the original target of the event.
func (e *AppDidFinishStartup) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *AppDidFinishStartup) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *AppDidFinishStartup) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *AppDidFinishStartup) Finish() {
	e.finished = true
}

// String implements the fmt.Stringer interface.
func (e *AppDidFinishStartup) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("AppDidFinishStartup[")
	if e.finished {
		buffer.WriteString("Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
