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

	"github.com/richardwilkes/toolbox/xmath/geom"
)

// UpdateCursor is generated when a cursor is being requested for the widget.
type UpdateCursor struct {
	target   Target
	where    geom.Point
	finished bool
}

// SendUpdateCursor sends a new UpdateCursor event. 'target' is the widget the mouse is over.
// 'where' is the location in the window where the mouse is. Returns true if the event was handled.
// If the event was not handled, the caller may want to explicitly set the cursor for the window.
func SendUpdateCursor(target Target, where geom.Point) bool {
	evt := &UpdateCursor{target: target, where: where}
	Dispatch(evt)
	return evt.Finished()
}

// Type returns the event type ID.
func (e *UpdateCursor) Type() Type {
	return UpdateCursorType
}

// Target the original target of the event.
func (e *UpdateCursor) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *UpdateCursor) Cascade() bool {
	return false
}

// Where returns the location in the window the mouse is.
func (e *UpdateCursor) Where() geom.Point {
	return e.where
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *UpdateCursor) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *UpdateCursor) Finish() {
	e.finished = true
}

// String implements the fmt.Stringer interface.
func (e *UpdateCursor) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("UpdateCursor[Where: [%v], Target: %v", e.where, e.target))
	if e.finished {
		buffer.WriteString(", Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
