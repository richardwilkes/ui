// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package keys

import (
	"fmt"
)

// These are platform-independent key codes used by this framework. Keys in the printable space map
// to their ASCII equivalent. Should the framework encounter a scan code not present in this list,
// it will create a new key code for it outside the range 0-255.
const (
	VK_Up             = 1
	VK_Left           = 2
	VK_Down           = 3
	VK_Right          = 4
	VK_Insert         = 5
	VK_Home           = 6
	VK_End            = 7
	VK_Backspace      = 8
	VK_Tab            = 9
	VK_Return         = 10
	VK_NumPadEnter    = 13
	VK_Eject          = 15
	VK_ShiftLeft      = 16
	VK_ShiftRight     = 17
	VK_ControlLeft    = 18
	VK_ControlRight   = 19
	VK_OptionLeft     = 20
	VK_OptionRight    = 21
	VK_CommandLeft    = 22
	VK_CommandRight   = 23
	VK_CapsLock       = 24
	VK_Menu           = 25
	VK_Fn             = 26
	VK_Escape         = 27
	VK_PageUp         = 28
	VK_PageDown       = 29
	VK_Space          = 32
	VK_Quote          = 39
	VK_Comma          = 44
	VK_Minus          = 45
	VK_Period         = 46
	VK_Slash          = 47
	VK_0              = 48
	VK_1              = 49
	VK_2              = 50
	VK_3              = 51
	VK_4              = 52
	VK_5              = 53
	VK_6              = 54
	VK_7              = 55
	VK_8              = 56
	VK_9              = 57
	VK_SemiColon      = 59
	VK_Equal          = 61
	VK_A              = 65
	VK_B              = 66
	VK_C              = 67
	VK_D              = 68
	VK_E              = 69
	VK_F              = 70
	VK_G              = 71
	VK_H              = 72
	VK_I              = 73
	VK_J              = 74
	VK_K              = 75
	VK_L              = 76
	VK_M              = 77
	VK_N              = 78
	VK_O              = 79
	VK_P              = 80
	VK_Q              = 81
	VK_R              = 82
	VK_S              = 83
	VK_T              = 84
	VK_U              = 85
	VK_V              = 86
	VK_W              = 87
	VK_X              = 88
	VK_Y              = 89
	VK_Z              = 90
	VK_LeftBracket    = 91
	VK_BackSlash      = 92
	VK_RightBracket   = 93
	VK_Backtick       = 96
	VK_Delete         = 127
	VK_NumPad0        = 130
	VK_NumPad1        = 131
	VK_NumPad2        = 132
	VK_NumPad3        = 133
	VK_NumPad4        = 134
	VK_NumPad5        = 135
	VK_NumPad6        = 136
	VK_NumPad7        = 137
	VK_NumPad8        = 138
	VK_NumPad9        = 139
	VK_NumLock        = 140
	VK_NumPadUp       = 141
	VK_NumPadLeft     = 142
	VK_NumPadDown     = 143
	VK_NumPadRight    = 144
	VK_NumPadCenter   = 145
	VK_NumPadClear    = 146
	VK_NumPadDivide   = 147
	VK_NumPadMultiply = 148
	VK_NumPadMinus    = 149
	VK_NumPadAdd      = 150
	VK_NumPadDecimal  = 151
	VK_NumPadDelete   = 152
	VK_NumPadHome     = 153
	VK_NumPadEnd      = 154
	VK_NumPadPageUp   = 155
	VK_NumPadPageDown = 156
	VK_F1             = 201
	VK_F2             = 202
	VK_F3             = 203
	VK_F4             = 204
	VK_F5             = 205
	VK_F6             = 206
	VK_F7             = 207
	VK_F8             = 208
	VK_F9             = 209
	VK_F10            = 210
	VK_F11            = 211
	VK_F12            = 212
	VK_F13            = 213
	VK_F14            = 214
	VK_F15            = 215
	VK_F16            = 216
	VK_F17            = 217
	VK_F18            = 218
	VK_F19            = 219
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
	insertKeyCodeMapping(mapping)
}

func insertKeyCodeMapping(mapping *Mapping) {
	keyCodeToMapping[mapping.KeyCode] = mapping
}

// MappingForScanCode returns the mapping for the specified scanCode, or nil.
func MappingForScanCode(scanCode int) *Mapping {
	mapping, ok := scanCodeToMapping[scanCode]
	if !ok {
		mapping = &Mapping{KeyCode: scanCode << 8, Dynamic: true, Name: fmt.Sprintf("ScanCode %d", scanCode)}
		InsertMapping(scanCode, mapping)
	}
	return mapping
}

// MappingForKeyCode returns the mapping for the specified keyCode, or nil.
func MappingForKeyCode(keyCode int) *Mapping {
	if mapping, ok := keyCodeToMapping[keyCode]; ok {
		return mapping
	}
	return nil
}

// IsControlAction returns true if the keyCode should trigger a control, such as a button, that is
// focused.
func IsControlAction(keyCode int) bool {
	return keyCode == VK_Return || keyCode == VK_NumPadEnter || keyCode == VK_Space
}

func Transform(scanCode int, chars string) (code int, ch rune) {
	extract := true
	if mapping := MappingForScanCode(scanCode); mapping != nil {
		code = mapping.KeyCode
		if !mapping.Dynamic {
			ch = mapping.KeyChar
			extract = false
		}
	} else {
		code = scanCode
	}
	if extract && chars != "" {
		ch = (([]rune)(chars))[0]
	}
	return
}
