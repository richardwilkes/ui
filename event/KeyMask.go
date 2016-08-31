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
)

// Possible KeyMask values.
const (
	CapsLockKeyMask KeyMask = 1 << iota
	ShiftKeyMask
	ControlKeyMask
	OptionKeyMask
	CommandKeyMask
	NonStickyKeyMask = ShiftKeyMask | ControlKeyMask | OptionKeyMask | CommandKeyMask
	AllKeyMask       = CapsLockKeyMask | NonStickyKeyMask
)

// KeyMask contains flags indicating which modifier keys were down when an event occurred.
type KeyMask int

// CapsLockDown returns true if the caps lock key is being pressed.
func (km KeyMask) CapsLockDown() bool {
	return km&CapsLockKeyMask == CapsLockKeyMask
}

// ShiftDown returns true if the shift key is being pressed.
func (km KeyMask) ShiftDown() bool {
	return km&ShiftKeyMask == ShiftKeyMask
}

// ControlDown returns true if the control key is being pressed.
func (km KeyMask) ControlDown() bool {
	return km&ControlKeyMask == ControlKeyMask
}

// OptionDown returns true if the option/alt key is being pressed.
func (km KeyMask) OptionDown() bool {
	return km&OptionKeyMask == OptionKeyMask
}

// CommandDown returns true if the command/meta key is being pressed.
func (km KeyMask) CommandDown() bool {
	return km&CommandKeyMask == CommandKeyMask
}

// PlatformMenuKeyMaskDown returns true if the platform's menu command key is being pressed.
func (km KeyMask) PlatformMenuKeyMaskDown() bool {
	mask := PlatformMenuKeyMask()
	return km&mask == mask
}

// String implements the fmt.Stringer interface.
func (km KeyMask) String() string {
	if km == 0 {
		return ""
	}
	var buffer bytes.Buffer
	km.append(&buffer, ControlKeyMask, "\u2303")
	km.append(&buffer, OptionKeyMask, "\u2387")
	km.append(&buffer, ShiftKeyMask, "\u21e7")
	km.append(&buffer, CapsLockKeyMask, "\u21ea")
	km.append(&buffer, CommandKeyMask, "\u2318")
	return buffer.String()
}

func (km KeyMask) append(buffer *bytes.Buffer, mask KeyMask, name string) {
	if km&mask == mask {
		buffer.WriteString(name)
	}
}
