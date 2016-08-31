// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package keys

// These are platform-independent key codes used by this framework and are all negative. Key codes
// that are positive are provided by the platform-specific code.
const (
	Escape = -(iota + 2)
	ShiftLeft
	ShiftRight
	ControlLeft
	ControlRight
	OptionLeft
	OptionRight
	CommandLeft
	CommandRight
	CapsLock
	Up
	Left
	Down
	Right
	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	F9
	F10
	F11
	F12
	F13
	F14
	F15
	Return
	Tab
	Space
	Backspace
	Menu
	Insert
	Delete
	Home
	End
	PageUp
	PageDown
	NumLock
	NumPadUp
	NumPadLeft
	NumPadDown
	NumPadRight
	NumPadCenter
	NumPadClear
	NumPadDivide
	NumPadMultiply
	NumPadMinus
	NumPadAdd
	NumPadEnter
	NumPadDecimal
	NumPadDelete
	NumPadHome
	NumPadEnd
	NumPadPageUp
	NumPadPageDown
	NumPad1
	NumPad2
	NumPad3
	NumPad4
	NumPad5
	NumPad6
	NumPad7
	NumPad8
	NumPad9
	NumPad0
	Backtick
	E
	I
	N
	U
)

// Mapping provides a mapping between key codes and the rune they represent, if any.
type Mapping struct {
	KeyCode int
	KeyChar rune
	// Dynamic means the KeyChar value is only one of multiple possibilities.
	Dynamic bool
	Name    string
}

var (
	scanCodeToMapping = make(map[int]*Mapping)
	keyCodeToMapping  = make(map[int]*Mapping)
)

// InsertMapping inserts a mapping for the specified scanCode into the map used by
// MappingForScanCode and MappingForKeyCode.
func InsertMapping(scanCode int, mapping *Mapping) {
	scanCodeToMapping[scanCode] = mapping
	keyCodeToMapping[mapping.KeyCode] = mapping
}

// MappingForScanCode returns the mapping for the specified scanCode, or nil.
func MappingForScanCode(scanCode int) *Mapping {
	if mapping, ok := scanCodeToMapping[scanCode]; ok {
		return mapping
	}
	return nil
}

// MappingForKeyCode returns the mapping for the specified keyCode, or nil. Note that all mapped
// keyCodes have a negative value.
func MappingForKeyCode(keyCode int) *Mapping {
	if mapping, ok := keyCodeToMapping[keyCode]; ok {
		return mapping
	}
	return nil
}

// IsControlAction returns true if the keyCode should trigger a control, such as a button, that is
// focused.
func IsControlAction(keyCode int) bool {
	return keyCode == Return || keyCode == NumPadEnter || keyCode == Space
}
