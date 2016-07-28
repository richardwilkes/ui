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
	"github.com/richardwilkes/ui/geom"
	"reflect"
)

// MouseExited is generated when the mouse leaves a widget.
type MouseExited struct {
	target    Target
	where     geom.Point
	modifiers KeyMask
	finished  bool
	discarded bool
}

// NewMouseExited creates a new MouseExited event. 'target' is the widget that mouse is leaving.
// 'where' is the location in the window where the mouse is. 'modifiers' are the keyboard
// modifiers keys that were down.
func NewMouseExited(target Target, where geom.Point, modifiers KeyMask) *MouseExited {
	return &MouseExited{target: target, where: where, modifiers: modifiers}
}

// Type returns the event type ID.
func (e *MouseExited) Type() Type {
	return MouseExitedType
}

// Target the original target of the event.
func (e *MouseExited) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *MouseExited) Cascade() bool {
	return false
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *MouseExited) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *MouseExited) Finish() {
	e.finished = true
}

// Where returns the location in the window the mouse is.
func (e *MouseExited) Where() geom.Point {
	return e.where
}

// Modifiers returns the key modifiers that were down.
func (e *MouseExited) Modifiers() KeyMask {
	return e.modifiers
}

// String implements the fmt.Stringer interface.
func (e *MouseExited) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("MouseExited[Where: [%v], Target: %v", e.where, reflect.ValueOf(e.target).Pointer()))
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
