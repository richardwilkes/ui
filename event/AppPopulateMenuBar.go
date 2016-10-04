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
	"fmt"
)

// AppPopulateMenuBar is generated when the menu bar needs to be populated. On platforms with a
// global menu bar, this will occur once, while on other platforms it will occur once for each
// window with a title bar.
type AppPopulateMenuBar struct {
	target   Target
	id       int64
	finished bool
}

// SendAppLastWindowClosed sends a new AppPopulateMenuBar event.
func SendAppPopulateMenuBar(id int64) {
	Dispatch(&AppPopulateMenuBar{target: GlobalTarget(), id: id})
}

// Type returns the event type ID.
func (e *AppPopulateMenuBar) Type() Type {
	return AppPopulateMenuBarType
}

// Target the original target of the event.
func (e *AppPopulateMenuBar) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *AppPopulateMenuBar) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *AppPopulateMenuBar) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *AppPopulateMenuBar) Finish() {
	e.finished = true
}

// ID returns the id of the window that requires its menu bar populated, or 0 if this is for the
// global menu bar.
func (e *AppPopulateMenuBar) ID() int64 {
	return e.id
}

// String implements the fmt.Stringer interface.
func (e *AppPopulateMenuBar) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("AppPopulateMenuBar[")
	buffer.WriteString("window id: ")
	buffer.WriteString(fmt.Sprint(e.id))
	if e.finished {
		buffer.WriteString(", Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
