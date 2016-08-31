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

// MouseEntered is generated when the mouse moves enters a widget.
type MouseEntered struct {
	target    Target
	where     geom.Point
	modifiers keys.Modifiers
	finished  bool
	discarded bool
}

// NewMouseEntered creates a new MouseEntered event. 'target' is the widget that mouse is entering.
// 'where' is the location in the window where the mouse is. 'modifiers' are the keyboard
// modifiers keys that were down.
func NewMouseEntered(target Target, where geom.Point, modifiers keys.Modifiers) *MouseEntered {
	return &MouseEntered{target: target, where: where, modifiers: modifiers}
}

// Type returns the event type ID.
func (e *MouseEntered) Type() Type {
	return MouseEnteredType
}

// Target the original target of the event.
func (e *MouseEntered) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *MouseEntered) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *MouseEntered) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *MouseEntered) Finish() {
	e.finished = true
}

// Where returns the location in the window the mouse is.
func (e *MouseEntered) Where() geom.Point {
	return e.where
}

// Modifiers returns the key modifiers that were down.
func (e *MouseEntered) Modifiers() keys.Modifiers {
	return e.modifiers
}

// String implements the fmt.Stringer interface.
func (e *MouseEntered) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("MouseEntered[Where: [%v], Target: %v", e.where, e.target))
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
