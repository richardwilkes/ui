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
	InsertMapping(vkUpArrow, &Mapping{KeyCode: VK_Up, Name: UpName})
	InsertMapping(vkLeftArrow, &Mapping{KeyCode: VK_Left, Name: LeftName})
	InsertMapping(vkDownArrow, &Mapping{KeyCode: VK_Down, Name: DownName})
	InsertMapping(vkRightArrow, &Mapping{KeyCode: VK_Right, Name: RightName})
	insertKeyCodeMapping(&Mapping{KeyCode: VK_Insert, Name: InsertName}) // Not on a Mac keyboard
	InsertMapping(vkHome, &Mapping{KeyCode: VK_Home, Name: HomeName})
	InsertMapping(vkEnd, &Mapping{KeyCode: VK_End, Name: EndName})
	InsertMapping(vkDelete, &Mapping{KeyCode: VK_Backspace, Name: BackspaceName})
	InsertMapping(vkTab, &Mapping{KeyCode: VK_Tab, KeyChar: '\t', Name: TabName})
	InsertMapping(vkReturn, &Mapping{KeyCode: VK_Return, KeyChar: '\n', Name: ReturnName})
	InsertMapping(vkANSI_KeypadEnter, &Mapping{KeyCode: VK_NumPadEnter, KeyChar: '\n', Name: NumPadEnterName})
	insertKeyCodeMapping(&Mapping{KeyCode: VK_Eject, Name: EjectName}) // Swallowed by the system
	InsertMapping(vkShift, &Mapping{KeyCode: VK_ShiftLeft, Name: LeftShiftName})
	InsertMapping(vkRightShift, &Mapping{KeyCode: VK_ShiftRight, Name: RightShiftName})
	InsertMapping(vkControl, &Mapping{KeyCode: VK_ControlLeft, Name: LeftControlName})
	InsertMapping(vkRightControl, &Mapping{KeyCode: VK_ControlRight, Name: RightControlName})
	InsertMapping(vkOption, &Mapping{KeyCode: VK_OptionLeft, Name: LeftOptionName})
	InsertMapping(vkRightOption, &Mapping{KeyCode: VK_OptionRight, Name: RightOptionName})
	InsertMapping(vkCommand, &Mapping{KeyCode: VK_CommandLeft, Name: LeftCommandName})
	InsertMapping(vkRightCommand, &Mapping{KeyCode: VK_CommandRight, Name: RightCommandName})
	InsertMapping(vkCapsLock, &Mapping{KeyCode: VK_CapsLock, Name: CapsLockName})
	insertKeyCodeMapping(&Mapping{KeyCode: VK_Menu, Name: MenuName}) // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VK_Fn, Name: FnName})     // Swallowed by the system
	InsertMapping(vkEscape, &Mapping{KeyCode: VK_Escape, KeyChar: '\x1b', Name: EscapeName})
	InsertMapping(vkPageUp, &Mapping{KeyCode: VK_PageUp, Name: PageUpName})
	InsertMapping(vkPageDown, &Mapping{KeyCode: VK_PageDown, Name: PageDownName})
	InsertMapping(vkSpace, &Mapping{KeyCode: VK_Space, KeyChar: ' ', Name: SpaceName})
	insertASCIIKeyCodeMapping(vkANSI_Quote, VK_Quote)
	insertASCIIKeyCodeMapping(vkANSI_Comma, VK_Comma)
	insertASCIIKeyCodeMapping(vkANSI_Minus, VK_Minus)
	insertASCIIKeyCodeMapping(vkANSI_Period, VK_Period)
	insertASCIIKeyCodeMapping(vkANSI_Slash, VK_Slash)
	insertASCIIKeyCodeMapping(vkANSI_0, VK_0)
	insertASCIIKeyCodeMapping(vkANSI_1, VK_1)
	insertASCIIKeyCodeMapping(vkANSI_2, VK_2)
	insertASCIIKeyCodeMapping(vkANSI_3, VK_3)
	insertASCIIKeyCodeMapping(vkANSI_4, VK_4)
	insertASCIIKeyCodeMapping(vkANSI_5, VK_5)
	insertASCIIKeyCodeMapping(vkANSI_6, VK_6)
	insertASCIIKeyCodeMapping(vkANSI_7, VK_7)
	insertASCIIKeyCodeMapping(vkANSI_8, VK_8)
	insertASCIIKeyCodeMapping(vkANSI_9, VK_9)
	insertASCIIKeyCodeMapping(vkANSI_Semicolon, VK_SemiColon)
	insertASCIIKeyCodeMapping(vkANSI_Equal, VK_Equal)
	insertASCIIKeyCodeMapping(vkANSI_A, VK_A)
	insertASCIIKeyCodeMapping(vkANSI_B, VK_B)
	insertASCIIKeyCodeMapping(vkANSI_C, VK_C)
	insertASCIIKeyCodeMapping(vkANSI_D, VK_D)
	insertASCIIKeyCodeMapping(vkANSI_E, VK_E)
	insertASCIIKeyCodeMapping(vkANSI_F, VK_F)
	insertASCIIKeyCodeMapping(vkANSI_G, VK_G)
	insertASCIIKeyCodeMapping(vkANSI_H, VK_H)
	insertASCIIKeyCodeMapping(vkANSI_I, VK_I)
	insertASCIIKeyCodeMapping(vkANSI_J, VK_J)
	insertASCIIKeyCodeMapping(vkANSI_K, VK_K)
	insertASCIIKeyCodeMapping(vkANSI_L, VK_L)
	insertASCIIKeyCodeMapping(vkANSI_M, VK_M)
	insertASCIIKeyCodeMapping(vkANSI_N, VK_N)
	insertASCIIKeyCodeMapping(vkANSI_O, VK_O)
	insertASCIIKeyCodeMapping(vkANSI_P, VK_P)
	insertASCIIKeyCodeMapping(vkANSI_Q, VK_Q)
	insertASCIIKeyCodeMapping(vkANSI_R, VK_R)
	insertASCIIKeyCodeMapping(vkANSI_S, VK_S)
	insertASCIIKeyCodeMapping(vkANSI_T, VK_T)
	insertASCIIKeyCodeMapping(vkANSI_U, VK_U)
	insertASCIIKeyCodeMapping(vkANSI_V, VK_V)
	insertASCIIKeyCodeMapping(vkANSI_W, VK_W)
	insertASCIIKeyCodeMapping(vkANSI_X, VK_X)
	insertASCIIKeyCodeMapping(vkANSI_Y, VK_Y)
	insertASCIIKeyCodeMapping(vkANSI_Z, VK_Z)
	insertASCIIKeyCodeMapping(vkANSI_LeftBracket, VK_LeftBracket)
	insertASCIIKeyCodeMapping(vkANSI_Backslash, VK_BackSlash)
	insertASCIIKeyCodeMapping(vkANSI_RightBracket, VK_RightBracket)
	insertASCIIKeyCodeMapping(vkANSI_Grave, VK_Backtick)
	InsertMapping(vkForwardDelete, &Mapping{KeyCode: VK_Delete, Name: DeleteName})
	InsertMapping(vkANSI_Keypad0, &Mapping{KeyCode: VK_NumPad0, KeyChar: '0', Name: NumPad0Name})
	InsertMapping(vkANSI_Keypad1, &Mapping{KeyCode: VK_NumPad1, KeyChar: '1', Name: NumPad1Name})
	InsertMapping(vkANSI_Keypad2, &Mapping{KeyCode: VK_NumPad2, KeyChar: '2', Name: NumPad2Name})
	InsertMapping(vkANSI_Keypad3, &Mapping{KeyCode: VK_NumPad3, KeyChar: '3', Name: NumPad3Name})
	InsertMapping(vkANSI_Keypad4, &Mapping{KeyCode: VK_NumPad4, KeyChar: '4', Name: NumPad4Name})
	InsertMapping(vkANSI_Keypad5, &Mapping{KeyCode: VK_NumPad5, KeyChar: '5', Name: NumPad5Name})
	InsertMapping(vkANSI_Keypad6, &Mapping{KeyCode: VK_NumPad6, KeyChar: '6', Name: NumPad6Name})
	InsertMapping(vkANSI_Keypad7, &Mapping{KeyCode: VK_NumPad7, KeyChar: '7', Name: NumPad7Name})
	InsertMapping(vkANSI_Keypad8, &Mapping{KeyCode: VK_NumPad8, KeyChar: '8', Name: NumPad8Name})
	InsertMapping(vkANSI_Keypad9, &Mapping{KeyCode: VK_NumPad9, KeyChar: '9', Name: NumPad9Name})
	insertKeyCodeMapping(&Mapping{KeyCode: VK_NumLock, Name: NumLockName})           // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VK_NumPadUp, Name: NumPadUpName})         // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VK_NumPadLeft, Name: NumPadLeftName})     // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VK_NumPadDown, Name: NumPadDownName})     // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VK_NumPadRight, Name: NumPadRightName})   // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VK_NumPadCenter, Name: NumPadCenterName}) // Not on a Mac keyboard
	InsertMapping(vkANSI_KeypadClear, &Mapping{KeyCode: VK_NumPadClear, Name: NumPadClearName})
	InsertMapping(vkANSI_KeypadDivide, &Mapping{KeyCode: VK_NumPadDivide, KeyChar: '/', Name: NumPadDivideName})
	InsertMapping(vkANSI_KeypadMultiply, &Mapping{KeyCode: VK_NumPadMultiply, KeyChar: '*', Name: NumPadMultiplyName})
	InsertMapping(vkANSI_KeypadMinus, &Mapping{KeyCode: VK_NumPadMinus, KeyChar: '-', Name: NumPadMinusName})
	InsertMapping(vkANSI_KeypadPlus, &Mapping{KeyCode: VK_NumPadAdd, KeyChar: '+', Name: NumPadAddName})
	InsertMapping(vkANSI_KeypadDecimal, &Mapping{KeyCode: VK_NumPadDecimal, KeyChar: '.', Name: NumPadDecimalName})
	insertKeyCodeMapping(&Mapping{KeyCode: VK_NumPadDelete, Name: NumPadDeleteName})     // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VK_NumPadHome, Name: NumPadHomeName})         // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VK_NumPadEnd, Name: NumPadEndName})           // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VK_NumPadPageUp, Name: NumPadPageUpName})     // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VK_NumPadPageDown, Name: NumPadPageDownName}) // Not on a Mac keyboard
	InsertMapping(vkF1, &Mapping{KeyCode: VK_F1, Name: F1Name})
	InsertMapping(vkF2, &Mapping{KeyCode: VK_F2, Name: F2Name})
	InsertMapping(vkF3, &Mapping{KeyCode: VK_F3, Name: F3Name})
	InsertMapping(vkF4, &Mapping{KeyCode: VK_F4, Name: F4Name})
	InsertMapping(vkF5, &Mapping{KeyCode: VK_F5, Name: F5Name})
	InsertMapping(vkF6, &Mapping{KeyCode: VK_F6, Name: F6Name})
	InsertMapping(vkF7, &Mapping{KeyCode: VK_F7, Name: F7Name})
	InsertMapping(vkF8, &Mapping{KeyCode: VK_F8, Name: F8Name})
	InsertMapping(vkF9, &Mapping{KeyCode: VK_F9, Name: F9Name})
	InsertMapping(vkF10, &Mapping{KeyCode: VK_F10, Name: F10Name})
	InsertMapping(vkF11, &Mapping{KeyCode: VK_F11, Name: F11Name})
	InsertMapping(vkF12, &Mapping{KeyCode: VK_F12, Name: F12Name})
	InsertMapping(vkF13, &Mapping{KeyCode: VK_F13, Name: F13Name})
	InsertMapping(vkF14, &Mapping{KeyCode: VK_F14, Name: F14Name})
	InsertMapping(vkF15, &Mapping{KeyCode: VK_F15, Name: F15Name})
	InsertMapping(vkF16, &Mapping{KeyCode: VK_F16, Name: F16Name})
	InsertMapping(vkF17, &Mapping{KeyCode: VK_F17, Name: F17Name})
	InsertMapping(vkF18, &Mapping{KeyCode: VK_F18, Name: F18Name})
	InsertMapping(vkF19, &Mapping{KeyCode: VK_F19, Name: F19Name})
}

func insertASCIIKeyCodeMapping(scanCode int, keyCode int) {
	InsertMapping(scanCode, &Mapping{KeyCode: keyCode, KeyChar: rune(keyCode), Dynamic: true, Name: string(rune(keyCode))})
}
