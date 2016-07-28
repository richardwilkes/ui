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
	"reflect"
)

// KeyUp is generated when a key is released.
type KeyUp struct {
	target    Target
	code      int
	modifiers KeyMask
	finished  bool
}

// NewKeyUp creates a new KeyUp event. 'target' is the widget that has the keyboard focus.
// 'code' is the key that was typed. 'modifiers' are the keyboard modifiers keys that were down.
func NewKeyUp(target Target, code int, modifiers KeyMask) *KeyUp {
	return &KeyUp{target: target, code: code, modifiers: modifiers}
}

// Type returns the event type ID.
func (e *KeyUp) Type() Type {
	return KeyUpType
}

// Target the original target of the event.
func (e *KeyUp) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *KeyUp) Cascade() bool {
	return true
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *KeyUp) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *KeyUp) Finish() {
	e.finished = true
}

// Code returns the key that was released.
func (e *KeyUp) Code() int {
	return e.code
}

// Modifiers returns the key modifiers that were down.
func (e *KeyUp) Modifiers() KeyMask {
	return e.modifiers
}

// String implements the fmt.Stringer interface.
func (e *KeyUp) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("KeyUp[Code: %d, Target: %v", e.code, reflect.ValueOf(e.target).Pointer()))
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
