// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package keys

// These are virtual key codes derived by experimentation from a US keyboard on a Mac. I'll
// probably have to have a per-platform description for these, along with some way to deal with
// non-US layouts. For now, though, this will do...
const (
	BaseMappedKeyCode = 1000
	A                 = 0
	S                 = 1
	D                 = 2
	F                 = 3
	H                 = 4
	G                 = 5
	Z                 = 6
	X                 = 7
	C                 = 8
	V                 = 9
	B                 = 11
	Q                 = 12
	W                 = 13
	E                 = 14
	R                 = 15
	Y                 = 16
	T                 = 17
	Num1              = 18
	Num2              = 19
	Num3              = 20
	Num4              = 21
	Num6              = 22
	Num5              = 23
	Equals            = 24
	Num9              = 25
	Num7              = 26
	Minus             = 27
	Num8              = 28
	Num0              = 29
	RightBracket      = 30
	O                 = 31
	U                 = 32
	LeftBracket       = 33
	I                 = 34
	P                 = 35
	Return            = 36
	L                 = 37
	J                 = 38
	Quote             = 39
	K                 = 40
	Semicolon         = 41
	BackSlash         = 42
	Comma             = 43
	Slash             = 44
	N                 = 45
	M                 = 46
	Period            = 47
	Tab               = 48
	Space             = 49
	Backtick          = 50
	Backspace         = 51
	Escape            = 53
	LeftCmd           = 54
	LeftMeta          = 54
	RightCmd          = 55
	RightMeta         = 55
	LeftShift         = 56
	CapsLock          = 57
	LeftAlt           = 58
	LeftOption        = 58
	LeftControl       = 59
	RightShift        = 60
	RightAlt          = 61
	RightOption       = 61
	RightControl      = 62
	NumPadDecimal     = 65
	NumPadMultiply    = 67
	NumPadAdd         = 69
	NumPadClear       = 71
	NumLock           = NumPadClear // Should these be unique?
	NumPadDivide      = 75
	NumPadEnter       = 76
	NumPadMinus       = 78
	NumPad0           = 82
	NumPad1           = 83
	NumPad2           = 84
	NumPad3           = 85
	NumPad4           = 86
	NumPad5           = 87
	NumPad6           = 88
	NumPad7           = 89
	NumPad8           = 91
	NumPad9           = 92
	F5                = 96
	F6                = 97
	F7                = 98
	F3                = 99
	F8                = 100
	F9                = 101
	F11               = 103
	F13               = 105
	F14               = 107
	F10               = 109
	F12               = 111
	F15               = 113
	Home              = 115
	PageUp            = 116
	Del               = 117
	F4                = 118
	End               = 119
	F2                = 120
	PageDown          = 121
	F1                = 122
	Left              = 123
	Right             = 124
	Down              = 125
	Up                = 126

	Insert         = -1 // RAW: Need code on Mac (Eject key?)
	Menu           = -1 // RAW: Need code on Mac (Fn key?)
	NumPadDelete   = -1 // RAW: Need code on Mac
	NumPadHome     = -1
	NumPadUp       = -1
	NumPadPageUp   = -1
	NumPadLeft     = -1
	NumPadRight    = -1
	NumPadEnd      = -1
	NumPadDown     = -1
	NumPadPageDown = -1
	NumPadCenter   = -1
)

// Mapping provides a mapping between key codes and the rune they represent, if any.
type Mapping struct {
	KeyCode int
	KeyChar rune
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
// keyCodes have a value equal or greater than BaseMappedKeyCode.
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
