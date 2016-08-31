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
	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui/keys"
)

// MouseMoved is generated when the mouse moves within a widget, except when a mouse button
// is also down (MouseDragged is generated for that).
type MouseMoved struct {
	target    Target
	where     geom.Point
	modifiers keys.Modifiers
	finished  bool
	discarded bool
}

// NewMouseMoved creates a new MouseMoved event. 'target' is the widget that mouse is over.
// 'where' is the location in the window where the mouse is. 'modifiers' are the keyboard
// modifiers keys that were down.
func NewMouseMoved(target Target, where geom.Point, modifiers keys.Modifiers) *MouseMoved {
	return &MouseMoved{target: target, where: where, modifiers: modifiers}
}

// Type returns the event type ID.
func (e *MouseMoved) Type() Type {
	return MouseMovedType
}

// Target the original target of the event.
func (e *MouseMoved) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *MouseMoved) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *MouseMoved) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *MouseMoved) Finish() {
	e.finished = true
}

// Where returns the location in the window the mouse is.
func (e *MouseMoved) Where() geom.Point {
	return e.where
}

// Modifiers returns the key modifiers that were down.
func (e *MouseMoved) Modifiers() keys.Modifiers {
	return e.modifiers
}

// String implements the fmt.Stringer interface.
func (e *MouseMoved) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("MouseMoved[Where: [%v], Target: %v", e.where, e.target))
	modifiers := e.modifiers.String()
	if modifiers != "" {
		buffer.WriteString(", ")
		buffer.WriteString(modifiers)
	}
	if e.finished {
		buffer.WriteString(", Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
