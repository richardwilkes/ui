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
)

// MouseDown is generated when a mouse button is pressed while over a widget.
type MouseDown struct {
	target    Target
	where     geom.Point
	modifiers KeyMask
	button    int
	clicks    int
	finished  bool
	discarded bool
}

// NewMouseDown creates a new MouseDown event. 'target' is the widget that is being clicked on.
// 'where' is the location in the window the mouse is being pressed. 'modifiers' are the keyboard
// modifiers keys that were down. 'button' is the button number. 'clicks' is the number of
// consecutive clicks in this widget.
func NewMouseDown(target Target, where geom.Point, modifiers KeyMask, button int, clicks int) *MouseDown {
	return &MouseDown{target: target, where: where, modifiers: modifiers, button: button, clicks: clicks}
}

// Type returns the event type ID.
func (e *MouseDown) Type() Type {
	return MouseDownType
}

// Target the original target of the event.
func (e *MouseDown) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *MouseDown) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *MouseDown) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *MouseDown) Finish() {
	e.finished = true
}

// Where returns the location in the window the mouse is being pressed.
func (e *MouseDown) Where() geom.Point {
	return e.where
}

// Modifiers returns the key modifiers that were down.
func (e *MouseDown) Modifiers() KeyMask {
	return e.modifiers
}

// Button returns the button that triggered this event.
func (e *MouseDown) Button() int {
	return e.button
}

// Clicks returns the number of consecutive clicks in this widget at or near this location.
func (e *MouseDown) Clicks() int {
	return e.clicks
}

// Discarded returns true if this event should be treated as if it never happened.
func (e *MouseDown) Discarded() bool {
	return e.discarded
}

// Discard marks this event to be thrown away.
func (e *MouseDown) Discard() {
	e.discarded = true
	e.finished = true
}

// String implements the fmt.Stringer interface.
func (e *MouseDown) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("MouseDown[")
	if e.discarded {
		buffer.WriteString("Discarded, ")
	}
	buffer.WriteString(fmt.Sprintf("Where: [%v], Target: %v, Button: %d, Clicks: %d", e.where, e.target, e.button, e.clicks))
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
