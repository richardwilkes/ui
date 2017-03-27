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
	InsertMapping(vkUpArrow, &Mapping{KeyCode: VirtualKeyUp, Name: UpName})
	InsertMapping(vkLeftArrow, &Mapping{KeyCode: VirtualKeyLeft, Name: LeftName})
	InsertMapping(vkDownArrow, &Mapping{KeyCode: VirtualKeyDown, Name: DownName})
	InsertMapping(vkRightArrow, &Mapping{KeyCode: VirtualKeyRight, Name: RightName})
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyInsert, Name: InsertName}) // Not on a Mac keyboard
	InsertMapping(vkHome, &Mapping{KeyCode: VirtualKeyHome, Name: HomeName})
	InsertMapping(vkEnd, &Mapping{KeyCode: VirtualKeyEnd, Name: EndName})
	InsertMapping(vkDelete, &Mapping{KeyCode: VirtualKeyBackspace, Name: BackspaceName})
	InsertMapping(vkTab, &Mapping{KeyCode: VirtualKeyTab, KeyChar: '\t', Name: TabName})
	InsertMapping(vkReturn, &Mapping{KeyCode: VirtualKeyReturn, KeyChar: '\n', Name: ReturnName})
	InsertMapping(vkANSI_KeypadEnter, &Mapping{KeyCode: VirtualKeyNumPadEnter, KeyChar: '\n', Name: NumPadEnterName})
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyEject, Name: EjectName}) // Swallowed by the system
	InsertMapping(vkShift, &Mapping{KeyCode: VirtualKeyShiftLeft, Name: LeftShiftName})
	InsertMapping(vkRightShift, &Mapping{KeyCode: VirtualKeyShiftRight, Name: RightShiftName})
	InsertMapping(vkControl, &Mapping{KeyCode: VirtualKeyControlLeft, Name: LeftControlName})
	InsertMapping(vkRightControl, &Mapping{KeyCode: VirtualKeyControlRight, Name: RightControlName})
	InsertMapping(vkOption, &Mapping{KeyCode: VirtualKeyOptionLeft, Name: LeftOptionName})
	InsertMapping(vkRightOption, &Mapping{KeyCode: VirtualKeyOptionRight, Name: RightOptionName})
	InsertMapping(vkCommand, &Mapping{KeyCode: VirtualKeyCommandLeft, Name: LeftCommandName})
	InsertMapping(vkRightCommand, &Mapping{KeyCode: VirtualKeyCommandRight, Name: RightCommandName})
	InsertMapping(vkCapsLock, &Mapping{KeyCode: VirtualKeyCapsLock, Name: CapsLockName})
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyMenu, Name: MenuName}) // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyFn, Name: FnName})     // Swallowed by the system
	InsertMapping(vkEscape, &Mapping{KeyCode: VirtualKeyEscape, KeyChar: '\x1b', Name: EscapeName})
	InsertMapping(vkPageUp, &Mapping{KeyCode: VirtualKeyPageUp, Name: PageUpName})
	InsertMapping(vkPageDown, &Mapping{KeyCode: VirtualKeyPageDown, Name: PageDownName})
	InsertMapping(vkSpace, &Mapping{KeyCode: VirtualKeySpace, KeyChar: ' ', Name: SpaceName})
	insertASCIIKeyCodeMapping(vkANSI_Quote, VirtualKeyQuote)
	insertASCIIKeyCodeMapping(vkANSI_Comma, VirtualKeyComma)
	insertASCIIKeyCodeMapping(vkANSI_Minus, VirtualKeyMinus)
	insertASCIIKeyCodeMapping(vkANSI_Period, VirtualKeyPeriod)
	insertASCIIKeyCodeMapping(vkANSI_Slash, VirtualKeySlash)
	insertASCIIKeyCodeMapping(vkANSI_0, VirtualKey0)
	insertASCIIKeyCodeMapping(vkANSI_1, VirtualKey1)
	insertASCIIKeyCodeMapping(vkANSI_2, VirtualKey2)
	insertASCIIKeyCodeMapping(vkANSI_3, VirtualKey3)
	insertASCIIKeyCodeMapping(vkANSI_4, VirtualKey4)
	insertASCIIKeyCodeMapping(vkANSI_5, VirtualKey5)
	insertASCIIKeyCodeMapping(vkANSI_6, VirtualKey6)
	insertASCIIKeyCodeMapping(vkANSI_7, VirtualKey7)
	insertASCIIKeyCodeMapping(vkANSI_8, VirtualKey8)
	insertASCIIKeyCodeMapping(vkANSI_9, VirtualKey9)
	insertASCIIKeyCodeMapping(vkANSI_Semicolon, VirtualKeySemiColon)
	insertASCIIKeyCodeMapping(vkANSI_Equal, VirtualKeyEqual)
	insertASCIIKeyCodeMapping(vkANSI_A, VirtualKeyA)
	insertASCIIKeyCodeMapping(vkANSI_B, VirtualKeyB)
	insertASCIIKeyCodeMapping(vkANSI_C, VirtualKeyC)
	insertASCIIKeyCodeMapping(vkANSI_D, VirtualKeyD)
	insertASCIIKeyCodeMapping(vkANSI_E, VirtualKeyE)
	insertASCIIKeyCodeMapping(vkANSI_F, VirtualKeyF)
	insertASCIIKeyCodeMapping(vkANSI_G, VirtualKeyG)
	insertASCIIKeyCodeMapping(vkANSI_H, VirtualKeyH)
	insertASCIIKeyCodeMapping(vkANSI_I, VirtualKeyI)
	insertASCIIKeyCodeMapping(vkANSI_J, VirtualKeyJ)
	insertASCIIKeyCodeMapping(vkANSI_K, VirtualKeyK)
	insertASCIIKeyCodeMapping(vkANSI_L, VirtualKeyL)
	insertASCIIKeyCodeMapping(vkANSI_M, VirtualKeyM)
	insertASCIIKeyCodeMapping(vkANSI_N, VirtualKeyN)
	insertASCIIKeyCodeMapping(vkANSI_O, VirtualKeyO)
	insertASCIIKeyCodeMapping(vkANSI_P, VirtualKeyP)
	insertASCIIKeyCodeMapping(vkANSI_Q, VirtualKeyQ)
	insertASCIIKeyCodeMapping(vkANSI_R, VirtualKeyR)
	insertASCIIKeyCodeMapping(vkANSI_S, VirtualKeyS)
	insertASCIIKeyCodeMapping(vkANSI_T, VirtualKeyT)
	insertASCIIKeyCodeMapping(vkANSI_U, VirtualKeyU)
	insertASCIIKeyCodeMapping(vkANSI_V, VirtualKeyV)
	insertASCIIKeyCodeMapping(vkANSI_W, VirtualKeyW)
	insertASCIIKeyCodeMapping(vkANSI_X, VirtualKeyX)
	insertASCIIKeyCodeMapping(vkANSI_Y, VirtualKeyY)
	insertASCIIKeyCodeMapping(vkANSI_Z, VirtualKeyZ)
	insertASCIIKeyCodeMapping(vkANSI_LeftBracket, VirtualKeyLeftBracket)
	insertASCIIKeyCodeMapping(vkANSI_Backslash, VirtualKeyBackSlash)
	insertASCIIKeyCodeMapping(vkANSI_RightBracket, VirtualKeyRightBracket)
	insertASCIIKeyCodeMapping(vkANSI_Grave, VirtualKeyBacktick)
	InsertMapping(vkForwardDelete, &Mapping{KeyCode: VirtualKeyDelete, Name: DeleteName})
	InsertMapping(vkANSI_Keypad0, &Mapping{KeyCode: VirtualKeyNumPad0, KeyChar: '0', Name: NumPad0Name})
	InsertMapping(vkANSI_Keypad1, &Mapping{KeyCode: VirtualKeyNumPad1, KeyChar: '1', Name: NumPad1Name})
	InsertMapping(vkANSI_Keypad2, &Mapping{KeyCode: VirtualKeyNumPad2, KeyChar: '2', Name: NumPad2Name})
	InsertMapping(vkANSI_Keypad3, &Mapping{KeyCode: VirtualKeyNumPad3, KeyChar: '3', Name: NumPad3Name})
	InsertMapping(vkANSI_Keypad4, &Mapping{KeyCode: VirtualKeyNumPad4, KeyChar: '4', Name: NumPad4Name})
	InsertMapping(vkANSI_Keypad5, &Mapping{KeyCode: VirtualKeyNumPad5, KeyChar: '5', Name: NumPad5Name})
	InsertMapping(vkANSI_Keypad6, &Mapping{KeyCode: VirtualKeyNumPad6, KeyChar: '6', Name: NumPad6Name})
	InsertMapping(vkANSI_Keypad7, &Mapping{KeyCode: VirtualKeyNumPad7, KeyChar: '7', Name: NumPad7Name})
	InsertMapping(vkANSI_Keypad8, &Mapping{KeyCode: VirtualKeyNumPad8, KeyChar: '8', Name: NumPad8Name})
	InsertMapping(vkANSI_Keypad9, &Mapping{KeyCode: VirtualKeyNumPad9, KeyChar: '9', Name: NumPad9Name})
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyNumLock, Name: NumLockName})           // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyNumPadUp, Name: NumPadUpName})         // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyNumPadLeft, Name: NumPadLeftName})     // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyNumPadDown, Name: NumPadDownName})     // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyNumPadRight, Name: NumPadRightName})   // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyNumPadCenter, Name: NumPadCenterName}) // Not on a Mac keyboard
	InsertMapping(vkANSI_KeypadClear, &Mapping{KeyCode: VirtualKeyNumPadClear, Name: NumPadClearName})
	InsertMapping(vkANSI_KeypadDivide, &Mapping{KeyCode: VirtualKeyNumPadDivide, KeyChar: '/', Name: NumPadDivideName})
	InsertMapping(vkANSI_KeypadMultiply, &Mapping{KeyCode: VirtualKeyNumPadMultiply, KeyChar: '*', Name: NumPadMultiplyName})
	InsertMapping(vkANSI_KeypadMinus, &Mapping{KeyCode: VirtualKeyNumPadMinus, KeyChar: '-', Name: NumPadMinusName})
	InsertMapping(vkANSI_KeypadPlus, &Mapping{KeyCode: VirtualKeyNumPadAdd, KeyChar: '+', Name: NumPadAddName})
	InsertMapping(vkANSI_KeypadDecimal, &Mapping{KeyCode: VirtualKeyNumPadDecimal, KeyChar: '.', Name: NumPadDecimalName})
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyNumPadDelete, Name: NumPadDeleteName})     // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyNumPadHome, Name: NumPadHomeName})         // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyNumPadEnd, Name: NumPadEndName})           // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyNumPadPageUp, Name: NumPadPageUpName})     // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyNumPadPageDown, Name: NumPadPageDownName}) // Not on a Mac keyboard
	InsertMapping(vkF1, &Mapping{KeyCode: VirtualKeyF1, Name: F1Name})
	InsertMapping(vkF2, &Mapping{KeyCode: VirtualKeyF2, Name: F2Name})
	InsertMapping(vkF3, &Mapping{KeyCode: VirtualKeyF3, Name: F3Name})
	InsertMapping(vkF4, &Mapping{KeyCode: VirtualKeyF4, Name: F4Name})
	InsertMapping(vkF5, &Mapping{KeyCode: VirtualKeyF5, Name: F5Name})
	InsertMapping(vkF6, &Mapping{KeyCode: VirtualKeyF6, Name: F6Name})
	InsertMapping(vkF7, &Mapping{KeyCode: VirtualKeyF7, Name: F7Name})
	InsertMapping(vkF8, &Mapping{KeyCode: VirtualKeyF8, Name: F8Name})
	InsertMapping(vkF9, &Mapping{KeyCode: VirtualKeyF9, Name: F9Name})
	InsertMapping(vkF10, &Mapping{KeyCode: VirtualKeyF10, Name: F10Name})
	InsertMapping(vkF11, &Mapping{KeyCode: VirtualKeyF11, Name: F11Name})
	InsertMapping(vkF12, &Mapping{KeyCode: VirtualKeyF12, Name: F12Name})
	InsertMapping(vkF13, &Mapping{KeyCode: VirtualKeyF13, Name: F13Name})
	InsertMapping(vkF14, &Mapping{KeyCode: VirtualKeyF14, Name: F14Name})
	InsertMapping(vkF15, &Mapping{KeyCode: VirtualKeyF15, Name: F15Name})
	InsertMapping(vkF16, &Mapping{KeyCode: VirtualKeyF16, Name: F16Name})
	InsertMapping(vkF17, &Mapping{KeyCode: VirtualKeyF17, Name: F17Name})
	InsertMapping(vkF18, &Mapping{KeyCode: VirtualKeyF18, Name: F18Name})
	InsertMapping(vkF19, &Mapping{KeyCode: VirtualKeyF19, Name: F19Name})
}

func insertASCIIKeyCodeMapping(scanCode int, keyCode int) {
	InsertMapping(scanCode, &Mapping{KeyCode: keyCode, KeyChar: rune(keyCode), Dynamic: true, Name: string(rune(keyCode))})
}
