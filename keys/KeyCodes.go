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
	A            = 0
	S            = 1
	D            = 2
	F            = 3
	H            = 4
	G            = 5
	Z            = 6
	X            = 7
	C            = 8
	V            = 9
	B            = 11
	Q            = 12
	W            = 13
	E            = 14
	R            = 15
	Y            = 16
	T            = 17
	Num1         = 18
	Num2         = 19
	Num3         = 20
	Num4         = 21
	Num6         = 22
	Num5         = 23
	Equals       = 24
	Num9         = 25
	Num7         = 26
	Minus        = 27
	Num8         = 28
	Num0         = 29
	RightBracket = 30
	O            = 31
	U            = 32
	LeftBracket  = 33
	I            = 34
	P            = 35
	Return       = 36
	L            = 37
	J            = 38
	Quote        = 39
	K            = 40
	Semicolon    = 41
	BackSlash    = 42
	Comma        = 43
	Slash        = 44
	N            = 45
	M            = 46
	Period       = 47
	Tab          = 48
	Space        = 49
	Backtick     = 50
	Backspace    = 51
	Escape       = 53
	LeftCmd      = 54
	LeftMeta     = 54
	RightCmd     = 55
	RightMeta    = 55
	LeftShift    = 56
	CapsLock     = 57
	LeftAlt      = 58
	LeftOption   = 58
	LeftControl  = 59
	RightShift   = 60
	RightAlt     = 61
	RightOption  = 61
	RightControl = 62
	NumpadPeriod = 65
	NumpadStar   = 67
	NumpadPlus   = 69
	Clear        = 71
	NumpadSlash  = 75
	Enter        = 76
	NumpadMinus  = 78
	Numpad0      = 82
	Numpad1      = 83
	Numpad2      = 84
	Numpad3      = 85
	Numpad4      = 86
	Numpad5      = 87
	Numpad6      = 88
	Numpad7      = 89
	Numpad8      = 91
	Numpad9      = 92
	F5           = 96
	F6           = 97
	F7           = 98
	F3           = 99
	F8           = 100
	F9           = 101
	F11          = 103
	F13          = 105
	F14          = 107
	F10          = 109
	F12          = 111
	F15          = 113
	Home         = 115
	PageUp       = 116
	Del          = 117
	F4           = 118
	End          = 119
	F2           = 120
	PageDown     = 121
	F1           = 122
	Left         = 123
	Right        = 124
	Down         = 125
	Up           = 126
)

// IsControlAction returns true if the key should trigger a control, such as a button, that is
// focused.
func IsControlAction(code int) bool {
	return code == Return || code == Enter || code == Space
}
