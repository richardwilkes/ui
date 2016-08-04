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
	"github.com/richardwilkes/ui/keys"
	"reflect"
)

// KeyDown is generated when a key is pressed.
type KeyDown struct {
	target    Target
	code      int
	modifiers KeyMask
	ch        rune
	repeat    bool
	finished  bool
	discarded bool
}

// NewKeyDown creates a new KeyDown event. 'target' is the widget that has the keyboard focus.
// 'code' is the virtual key code. 'ch' is the rune (may be 0). 'autoRepeat' is true if the key is
// auto-repeating. 'modifiers' are the keyboard modifiers keys that were down.
func NewKeyDown(target Target, code int, ch rune, autoRepeat bool, modifiers KeyMask) *KeyDown {
	return &KeyDown{target: target, code: code, ch: ch, modifiers: modifiers, repeat: autoRepeat}
}

// Type returns the event type ID.
func (e *KeyDown) Type() Type {
	return KeyDownType
}

// Target the original target of the event.
func (e *KeyDown) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *KeyDown) Cascade() bool {
	return true
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *KeyDown) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *KeyDown) Finish() {
	e.finished = true
}

// IsControlActionKey returns true if the key should trigger a control, such as a button, that is
// focused.
func (e *KeyDown) IsControlActionKey() bool {
	return e.code == keys.Return || e.code == keys.Enter || e.code == keys.Space
}

// Code returns the virtual key code.
func (e *KeyDown) Code() int {
	return e.code
}

// Rune returns the rune that was typed. May be 0.
func (e *KeyDown) Rune() rune {
	return e.ch
}

// Modifiers returns the key modifiers that were down.
func (e *KeyDown) Modifiers() KeyMask {
	return e.modifiers
}

// Repeat returns true if this key was generated as part of an auto-repeating key.
func (e *KeyDown) Repeat() bool {
	return e.repeat
}

// Discarded returns true if this event should be treated as if it never happened.
func (e *KeyDown) Discarded() bool {
	return e.discarded
}

// Discard marks this event to be thrown away.
func (e *KeyDown) Discard() {
	e.discarded = true
	e.finished = true
}

// String implements the fmt.Stringer interface.
func (e *KeyDown) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("KeyDown[")
	if e.discarded {
		buffer.WriteString("Discarded, ")
	}
	buffer.WriteString(fmt.Sprintf("Code: %d, Rune '%v', Target: %v", e.code, e.ch, reflect.ValueOf(e.target).Pointer()))
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
