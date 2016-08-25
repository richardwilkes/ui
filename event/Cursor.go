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
	"github.com/richardwilkes/ui/cursor"
	"github.com/richardwilkes/geom"
	"reflect"
)

// Cursor is generated when a cursor is being requested for the widget.
type Cursor struct {
	target   Target
	where    geom.Point
	cursor   *cursor.Cursor
	finished bool
}

// NewCursor creates a new Cursor event. 'target' is the widget the mouse is over. 'where' is the
// location in the window where the mouse is.
func NewCursor(target Target, where geom.Point) *Cursor {
	return &Cursor{target: target, where: where, cursor: cursor.Arrow}
}

// Type returns the event type ID.
func (e *Cursor) Type() Type {
	return CursorType
}

// Target the original target of the event.
func (e *Cursor) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *Cursor) Cascade() bool {
	return false
}

// Where returns the location in the window the mouse is.
func (e *Cursor) Where() geom.Point {
	return e.where
}

// Cursor returns the cursor that was set for the widget.
func (e *Cursor) Cursor() *cursor.Cursor {
	return e.cursor
}

// SetCursor sets the cursor to use for the widget.
func (e *Cursor) SetCursor(cursor *cursor.Cursor) {
	e.cursor = cursor
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *Cursor) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *Cursor) Finish() {
	e.finished = true
}

// String implements the fmt.Stringer interface.
func (e *Cursor) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Cursor[Where: [%v], Target: %v", e.where, reflect.ValueOf(e.target).Pointer()))
	if e.finished {
		buffer.WriteString(", Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
