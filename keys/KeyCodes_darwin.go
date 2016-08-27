// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package keys

// Taken from the old Events.h header in Carbon.
const (
	vkANSI_A              = 0x00
	vkANSI_S              = 0x01
	vkANSI_D              = 0x02
	vkANSI_F              = 0x03
	vkANSI_H              = 0x04
	vkANSI_G              = 0x05
	vkANSI_Z              = 0x06
	vkANSI_X              = 0x07
	vkANSI_C              = 0x08
	vkANSI_V              = 0x09
	vkANSI_B              = 0x0B
	vkANSI_Q              = 0x0C
	vkANSI_W              = 0x0D
	vkANSI_E              = 0x0E
	vkANSI_R              = 0x0F
	vkANSI_Y              = 0x10
	vkANSI_T              = 0x11
	vkANSI_1              = 0x12
	vkANSI_2              = 0x13
	vkANSI_3              = 0x14
	vkANSI_4              = 0x15
	vkANSI_6              = 0x16
	vkANSI_5              = 0x17
	vkANSI_Equal          = 0x18
	vkANSI_9              = 0x19
	vkANSI_7              = 0x1A
	vkANSI_Minus          = 0x1B
	vkANSI_8              = 0x1C
	vkANSI_0              = 0x1D
	vkANSI_RightBracket   = 0x1E
	vkANSI_O              = 0x1F
	vkANSI_U              = 0x20
	vkANSI_LeftBracket    = 0x21
	vkANSI_I              = 0x22
	vkANSI_P              = 0x23
	vkANSI_L              = 0x25
	vkANSI_J              = 0x26
	vkANSI_Quote          = 0x27
	vkANSI_K              = 0x28
	vkANSI_Semicolon      = 0x29
	vkANSI_Backslash      = 0x2A
	vkANSI_Comma          = 0x2B
	vkANSI_Slash          = 0x2C
	vkANSI_N              = 0x2D
	vkANSI_M              = 0x2E
	vkANSI_Period         = 0x2F
	vkANSI_Grave          = 0x32
	vkANSI_KeypadDecimal  = 0x41
	vkANSI_KeypadMultiply = 0x43
	vkANSI_KeypadPlus     = 0x45
	vkANSI_KeypadClear    = 0x47
	vkANSI_KeypadDivide   = 0x4B
	vkANSI_KeypadEnter    = 0x4C
	vkANSI_KeypadMinus    = 0x4E
	vkANSI_KeypadEquals   = 0x51
	vkANSI_Keypad0        = 0x52
	vkANSI_Keypad1        = 0x53
	vkANSI_Keypad2        = 0x54
	vkANSI_Keypad3        = 0x55
	vkANSI_Keypad4        = 0x56
	vkANSI_Keypad5        = 0x57
	vkANSI_Keypad6        = 0x58
	vkANSI_Keypad7        = 0x59
	vkANSI_Keypad8        = 0x5B
	vkANSI_Keypad9        = 0x5C
	vkReturn              = 0x24
	vkTab                 = 0x30
	vkSpace               = 0x31
	vkDelete              = 0x33
	vkEscape              = 0x35
	vkRightCommand        = 0x36
	vkCommand             = 0x37
	vkShift               = 0x38
	vkCapsLock            = 0x39
	vkOption              = 0x3A
	vkControl             = 0x3B
	vkRightShift          = 0x3C
	vkRightOption         = 0x3D
	vkRightControl        = 0x3E
	vkFunction            = 0x3F
	vkF17                 = 0x40
	vkVolumeUp            = 0x48
	vkVolumeDown          = 0x49
	vkMute                = 0x4A
	vkF18                 = 0x4F
	vkF19                 = 0x50
	vkF20                 = 0x5A
	vkF5                  = 0x60
	vkF6                  = 0x61
	vkF7                  = 0x62
	vkF3                  = 0x63
	vkF8                  = 0x64
	vkF9                  = 0x65
	vkF11                 = 0x67
	vkF13                 = 0x69
	vkF16                 = 0x6A
	vkF14                 = 0x6B
	vkF10                 = 0x6D
	vkF12                 = 0x6F
	vkF15                 = 0x71
	vkHelp                = 0x72
	vkHome                = 0x73
	vkPageUp              = 0x74
	vkForwardDelete       = 0x75
	vkF4                  = 0x76
	vkEnd                 = 0x77
	vkF2                  = 0x78
	vkPageDown            = 0x79
	vkF1                  = 0x7A
	vkLeftArrow           = 0x7B
	vkRightArrow          = 0x7C
	vkDownArrow           = 0x7D
	vkUpArrow             = 0x7E
)

func init() {
	InsertMapping(vkEscape, &Mapping{KeyCode: Escape, KeyChar: '\x1b', Name: EscapeName})
	InsertMapping(vkShift, &Mapping{KeyCode: ShiftLeft, Name: LeftShiftName})
	InsertMapping(vkRightShift, &Mapping{KeyCode: ShiftRight, Name: RightShiftName})
	InsertMapping(vkControl, &Mapping{KeyCode: ControlLeft, Name: LeftControlName})
	InsertMapping(vkRightControl, &Mapping{KeyCode: ControlRight, Name: RightControlName})
	InsertMapping(vkOption, &Mapping{KeyCode: OptionLeft, Name: LeftOptionName})
	InsertMapping(vkRightOption, &Mapping{KeyCode: OptionRight, Name: RightOptionName})
	InsertMapping(vkCommand, &Mapping{KeyCode: CommandLeft, Name: LeftCommandName})
	InsertMapping(vkRightCommand, &Mapping{KeyCode: CommandRight, Name: RightCommandName})
	InsertMapping(vkCapsLock, &Mapping{KeyCode: CapsLock, Name: CapsLockName})
	InsertMapping(vkUpArrow, &Mapping{KeyCode: Up, Name: UpName})
	InsertMapping(vkLeftArrow, &Mapping{KeyCode: Left, Name: LeftName})
	InsertMapping(vkDownArrow, &Mapping{KeyCode: Down, Name: DownName})
	InsertMapping(vkRightArrow, &Mapping{KeyCode: Right, Name: RightName})
	InsertMapping(vkF1, &Mapping{KeyCode: F1, Name: F1Name})
	InsertMapping(vkF2, &Mapping{KeyCode: F2, Name: F2Name})
	InsertMapping(vkF3, &Mapping{KeyCode: F3, Name: F3Name})
	InsertMapping(vkF4, &Mapping{KeyCode: F4, Name: F4Name})
	InsertMapping(vkF5, &Mapping{KeyCode: F5, Name: F5Name})
	InsertMapping(vkF6, &Mapping{KeyCode: F6, Name: F6Name})
	InsertMapping(vkF7, &Mapping{KeyCode: F7, Name: F7Name})
	InsertMapping(vkF8, &Mapping{KeyCode: F8, Name: F8Name})
	InsertMapping(vkF9, &Mapping{KeyCode: F9, Name: F9Name})
	InsertMapping(vkF10, &Mapping{KeyCode: F10, Name: F10Name})
	InsertMapping(vkF11, &Mapping{KeyCode: F11, Name: F11Name})
	InsertMapping(vkF12, &Mapping{KeyCode: F12, Name: F12Name})
	InsertMapping(vkF13, &Mapping{KeyCode: F13, Name: F13Name})
	InsertMapping(vkF14, &Mapping{KeyCode: F14, Name: F14Name})
	InsertMapping(vkF15, &Mapping{KeyCode: F15, Name: F15Name})
	InsertMapping(vkReturn, &Mapping{KeyCode: Return, KeyChar: '\n', Name: ReturnName})
	InsertMapping(vkTab, &Mapping{KeyCode: Tab, KeyChar: '\t', Name: TabName})
	InsertMapping(vkSpace, &Mapping{KeyCode: Space, KeyChar: ' ', Name: SpaceName})
	InsertMapping(vkDelete, &Mapping{KeyCode: Backspace, Name: BackspaceName})
	// Maps to the Fn key... swallowed by the system
	//InsertMapping(vkMenu, &Mapping{KeyCode: Menu, Name: MenuName})
	// Maps to the Help or Eject key, depending on the keyboard... swallowed by the system
	//InsertMapping(vkInsert, &Mapping{KeyCode: Insert, Name: InsertName})
	InsertMapping(vkForwardDelete, &Mapping{KeyCode: Delete, Name: DeleteName})
	InsertMapping(vkHome, &Mapping{KeyCode: Home, Name: HomeName})
	InsertMapping(vkEnd, &Mapping{KeyCode: End, Name: EndName})
	InsertMapping(vkPageUp, &Mapping{KeyCode: PageUp, Name: PageUpName})
	InsertMapping(vkPageDown, &Mapping{KeyCode: PageDown, Name: PageDownName})
	// None of the Mac keyboards I've seen have the concept of NumLock and the alternates it
	// provides
	//InsertMapping(vkNumLock, &Mapping{KeyCode: NumLock, Name: NumLockName})
	//InsertMapping(vkANSI_KeypadUp, &Mapping{KeyCode: NumPadUp, Name: NumPadUpName})
	//InsertMapping(vkANSI_KeypadLeft, &Mapping{KeyCode: NumPadLeft, Name: NumPadLeftName})
	//InsertMapping(vkANSI_KeypadDown, &Mapping{KeyCode: NumPadDown, Name: NumPadDownName})
	//InsertMapping(vkANSI_KeypadRight, &Mapping{KeyCode: NumPadRight, Name: NumPadRightName})
	//InsertMapping(vkANSI_KeypadCenter, &Mapping{KeyCode: NumPadCenter, Name: NumPadCenterName})
	InsertMapping(vkANSI_KeypadClear, &Mapping{KeyCode: NumPadClear, Name: NumPadClearName})
	InsertMapping(vkANSI_KeypadDivide, &Mapping{KeyCode: NumPadDivide, KeyChar: '/', Name: NumPadDivideName})
	InsertMapping(vkANSI_KeypadMultiply, &Mapping{KeyCode: NumPadMultiply, KeyChar: '*', Name: NumPadMultiplyName})
	InsertMapping(vkANSI_KeypadMinus, &Mapping{KeyCode: NumPadMinus, KeyChar: '-', Name: NumPadMinusName})
	InsertMapping(vkANSI_KeypadPlus, &Mapping{KeyCode: NumPadAdd, KeyChar: '+', Name: NumPadAddName})
	InsertMapping(vkANSI_KeypadEnter, &Mapping{KeyCode: NumPadEnter, KeyChar: '\n', Name: NumPadEnterName})
	InsertMapping(vkANSI_KeypadDecimal, &Mapping{KeyCode: NumPadDecimal, KeyChar: '.', Name: NumPadDecimalName})
	//InsertMapping(vkANSI_KeypadDelete, &Mapping{KeyCode: NumPadDelete, Name: NumPadDeleteName})
	//InsertMapping(vkANSI_KeypadHome, &Mapping{KeyCode: NumPadHome, Name: NumPadHomeName})
	//InsertMapping(vkANSI_KeypadEnd, &Mapping{KeyCode: NumPadEnd, Name: NumPadEndName})
	//InsertMapping(vkANSI_KeypadPage_Up, &Mapping{KeyCode: NumPadPageUp, Name: NumPadPageUpName})
	//InsertMapping(vkANSI_KeypadPage_Down, &Mapping{KeyCode: NumPadPageDown, Name: NumPadPageDownName})
	InsertMapping(vkANSI_Keypad1, &Mapping{KeyCode: NumPad1, KeyChar: '1', Name: NumPad1Name})
	InsertMapping(vkANSI_Keypad2, &Mapping{KeyCode: NumPad2, KeyChar: '2', Name: NumPad2Name})
	InsertMapping(vkANSI_Keypad3, &Mapping{KeyCode: NumPad3, KeyChar: '3', Name: NumPad3Name})
	InsertMapping(vkANSI_Keypad4, &Mapping{KeyCode: NumPad4, KeyChar: '4', Name: NumPad4Name})
	InsertMapping(vkANSI_Keypad5, &Mapping{KeyCode: NumPad5, KeyChar: '5', Name: NumPad5Name})
	InsertMapping(vkANSI_Keypad6, &Mapping{KeyCode: NumPad6, KeyChar: '6', Name: NumPad6Name})
	InsertMapping(vkANSI_Keypad7, &Mapping{KeyCode: NumPad7, KeyChar: '7', Name: NumPad7Name})
	InsertMapping(vkANSI_Keypad8, &Mapping{KeyCode: NumPad8, KeyChar: '8', Name: NumPad8Name})
	InsertMapping(vkANSI_Keypad9, &Mapping{KeyCode: NumPad9, KeyChar: '9', Name: NumPad9Name})
	InsertMapping(vkANSI_Keypad0, &Mapping{KeyCode: NumPad0, KeyChar: '0', Name: NumPad0Name})
	InsertMapping(vkANSI_Grave, &Mapping{KeyCode: Backtick, KeyChar: '`', Dynamic: true, Name: "`"})
	InsertMapping(vkANSI_E, &Mapping{KeyCode: E, KeyChar: 'e', Dynamic: true, Name: "e"})
	InsertMapping(vkANSI_I, &Mapping{KeyCode: I, KeyChar: 'i', Dynamic: true, Name: "i"})
	InsertMapping(vkANSI_N, &Mapping{KeyCode: N, KeyChar: 'n', Dynamic: true, Name: "n"})
	InsertMapping(vkANSI_U, &Mapping{KeyCode: U, KeyChar: 'u', Dynamic: true, Name: "u"})
}
