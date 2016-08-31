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

// MouseUp is generated when a mouse button is released after a MouseDown was generated in the
// widget.
type MouseUp struct {
	target    Target
	where     geom.Point
	modifiers keys.Modifiers
	button    int
	finished  bool
	discarded bool
}

// NewMouseUp creates a new MouseUp event. 'target' is the widget that was being clicked on.
// 'where' is the location in the window where the mouse is. 'modifiers' are the keyboard
// modifiers keys that were down. 'button' is the button number.
func NewMouseUp(target Target, where geom.Point, modifiers keys.Modifiers, button int) *MouseUp {
	return &MouseUp{target: target, where: where, modifiers: modifiers, button: button}
}

// Type returns the event type ID.
func (e *MouseUp) Type() Type {
	return MouseUpType
}

// Target the original target of the event.
func (e *MouseUp) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *MouseUp) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *MouseUp) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *MouseUp) Finish() {
	e.finished = true
}

// Where returns the location in the window the mouse is being released.
func (e *MouseUp) Where() geom.Point {
	return e.where
}

// Modifiers returns the key modifiers that were down.
func (e *MouseUp) Modifiers() keys.Modifiers {
	return e.modifiers
}

// Button returns the button that triggered this event.
func (e *MouseUp) Button() int {
	return e.button
}

// String implements the fmt.Stringer interface.
func (e *MouseUp) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("MouseUp[Where: [%v], Target: %v, Button: %d", e.where, e.target, e.button))
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
