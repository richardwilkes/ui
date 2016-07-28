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

// KeyTyped is generated when a key is typed.
type KeyTyped struct {
	target    Target
	modifiers KeyMask
	ch        rune
	repeat    bool
	finished  bool
	discarded bool
}

// NewKeyTyped creates a new KeyTyped event. 'target' is the widget that has the keyboard focus.
// 'ch' is the rune that was typed. 'autoRepeat' is true if the rune is auto-repeating.
// 'modifiers' are the keyboard modifiers keys that were down.
func NewKeyTyped(target Target, ch rune, autoRepeat bool, modifiers KeyMask) *KeyTyped {
	return &KeyTyped{target: target, ch: ch, modifiers: modifiers, repeat: autoRepeat}
}

// Type returns the event type ID.
func (e *KeyTyped) Type() Type {
	return KeyTypedType
}

// Target the original target of the event.
func (e *KeyTyped) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *KeyTyped) Cascade() bool {
	return true
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *KeyTyped) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *KeyTyped) Finish() {
	e.finished = true
}

// Rune returns the rune that was typed.
func (e *KeyTyped) Rune() rune {
	return e.ch
}

// Modifiers returns the key modifiers that were down.
func (e *KeyTyped) Modifiers() KeyMask {
	return e.modifiers
}

// Repeat returns true if this key was generated as part of an auto-repeating key.
func (e *KeyTyped) Repeat() bool {
	return e.repeat
}

// String implements the fmt.Stringer interface.
func (e *KeyTyped) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("KeyTyped[Rune: '%v', Target: %v", e.ch, reflect.ValueOf(e.target).Pointer()))
	modifiers := e.modifiers.String()
	if modifiers != "" {
		buffer.WriteString(", ")
		buffer.WriteString(modifiers)
	}
	if e.repeat {
		buffer.WriteString(", Auto-Repeat")
	}
	if e.finished {
		buffer.WriteString(", Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
