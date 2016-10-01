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
	// #include <X11/keysym.h>
	"C"
)

// RAW: Look at all instances of insertASCIIKeyCodeMappings at substitute something appropriate
// for linux.

func init() {
	InsertMapping(C.XK_Up, &Mapping{KeyCode: VK_Up, Name: UpName})
	InsertMapping(C.XK_Left, &Mapping{KeyCode: VK_Left, Name: LeftName})
	InsertMapping(C.XK_Down, &Mapping{KeyCode: VK_Down, Name: DownName})
	InsertMapping(C.XK_Right, &Mapping{KeyCode: VK_Right, Name: RightName})
	InsertMapping(C.XK_Insert, &Mapping{KeyCode: VK_Insert, Name: InsertName})
	InsertMapping(C.XK_Home, &Mapping{KeyCode: VK_Home, Name: HomeName})
	InsertMapping(C.XK_End, &Mapping{KeyCode: VK_End, Name: EndName})
	InsertMapping(C.XK_BackSpace, &Mapping{KeyCode: VK_Backspace, Name: BackspaceName})
	InsertMapping(C.XK_Tab, &Mapping{KeyCode: VK_Tab, KeyChar: '\t', Name: TabName})
	InsertMapping(C.XK_ISO_Left_Tab, &Mapping{KeyCode: VK_Tab, KeyChar: '\t', Name: TabName})
	InsertMapping(C.XK_Return, &Mapping{KeyCode: VK_Return, KeyChar: '\n', Name: ReturnName})
	InsertMapping(C.XK_KP_Enter, &Mapping{KeyCode: VK_NumPadEnter, KeyChar: '\n', Name: NumPadEnterName})
	insertKeyCodeMapping(&Mapping{KeyCode: VK_Eject, Name: EjectName}) // Not a PC keyboard
	InsertMapping(C.XK_Shift_L, &Mapping{KeyCode: VK_ShiftLeft, Name: LeftShiftName})
	InsertMapping(C.XK_Shift_R, &Mapping{KeyCode: VK_ShiftRight, Name: RightShiftName})
	InsertMapping(C.XK_Control_L, &Mapping{KeyCode: VK_ControlLeft, Name: LeftControlName})
	InsertMapping(C.XK_Control_R, &Mapping{KeyCode: VK_ControlRight, Name: RightControlName})
	InsertMapping(C.XK_Alt_L, &Mapping{KeyCode: VK_OptionLeft, Name: LeftOptionName})
	InsertMapping(C.XK_Alt_R, &Mapping{KeyCode: VK_OptionRight, Name: RightOptionName})
	InsertMapping(C.XK_Super_L, &Mapping{KeyCode: VK_CommandLeft, Name: LeftWindowsName})
	InsertMapping(C.XK_Super_R, &Mapping{KeyCode: VK_CommandRight, Name: RightWindowsName})
	InsertMapping(C.XK_Caps_Lock, &Mapping{KeyCode: VK_CapsLock, Name: CapsLockName})
	InsertMapping(C.XK_Menu, &Mapping{KeyCode: VK_Menu, Name: MenuName})
	insertKeyCodeMapping(&Mapping{KeyCode: VK_Fn, Name: FnName}) // Not on a PC keyboard
	InsertMapping(C.XK_Escape, &Mapping{KeyCode: VK_Escape, KeyChar: '\x1b', Name: EscapeName})
	InsertMapping(C.XK_Page_Up, &Mapping{KeyCode: VK_PageUp, Name: PageUpName})
	InsertMapping(C.XK_Page_Down, &Mapping{KeyCode: VK_PageDown, Name: PageDownName})
	InsertMapping(' ', &Mapping{KeyCode: VK_Space, KeyChar: ' ', Name: SpaceName})
	insertASCIIKeyCodeMapping('\'', VK_Quote)
	insertASCIIKeyCodeMapping(',', VK_Comma)
	insertASCIIKeyCodeMapping('-', VK_Minus)
	insertASCIIKeyCodeMapping('.', VK_Period)
	insertASCIIKeyCodeMapping('/', VK_Slash)
	insertASCIIKeyCodeMapping('0', VK_0)
	insertASCIIKeyCodeMapping('1', VK_1)
	insertASCIIKeyCodeMapping('2', VK_2)
	insertASCIIKeyCodeMapping('3', VK_3)
	insertASCIIKeyCodeMapping('4', VK_4)
	insertASCIIKeyCodeMapping('5', VK_5)
	insertASCIIKeyCodeMapping('6', VK_6)
	insertASCIIKeyCodeMapping('7', VK_7)
	insertASCIIKeyCodeMapping('8', VK_8)
	insertASCIIKeyCodeMapping('9', VK_9)
	insertASCIIKeyCodeMapping(';', VK_SemiColon)
	insertASCIIKeyCodeMapping('=', VK_Equal)
	insertASCIILetterCodeMapping(VK_A)
	insertASCIILetterCodeMapping(VK_B)
	insertASCIILetterCodeMapping(VK_C)
	insertASCIILetterCodeMapping(VK_D)
	insertASCIILetterCodeMapping(VK_E)
	insertASCIILetterCodeMapping(VK_F)
	insertASCIILetterCodeMapping(VK_G)
	insertASCIILetterCodeMapping(VK_H)
	insertASCIILetterCodeMapping(VK_I)
	insertASCIILetterCodeMapping(VK_J)
	insertASCIILetterCodeMapping(VK_K)
	insertASCIILetterCodeMapping(VK_L)
	insertASCIILetterCodeMapping(VK_M)
	insertASCIILetterCodeMapping(VK_N)
	insertASCIILetterCodeMapping(VK_O)
	insertASCIILetterCodeMapping(VK_P)
	insertASCIILetterCodeMapping(VK_Q)
	insertASCIILetterCodeMapping(VK_R)
	insertASCIILetterCodeMapping(VK_S)
	insertASCIILetterCodeMapping(VK_T)
	insertASCIILetterCodeMapping(VK_U)
	insertASCIILetterCodeMapping(VK_V)
	insertASCIILetterCodeMapping(VK_W)
	insertASCIILetterCodeMapping(VK_X)
	insertASCIILetterCodeMapping(VK_Y)
	insertASCIILetterCodeMapping(VK_Z)
	insertASCIIKeyCodeMapping('[', VK_LeftBracket)
	insertASCIIKeyCodeMapping('\\', VK_BackSlash)
	insertASCIIKeyCodeMapping(']', VK_RightBracket)
	insertASCIIKeyCodeMapping('`', VK_Backtick)
	InsertMapping(C.XK_Delete, &Mapping{KeyCode: VK_Delete, Name: DeleteName})
	InsertMapping(C.XK_KP_0, &Mapping{KeyCode: VK_NumPad0, KeyChar: '0', Name: NumPad0Name})
	InsertMapping(C.XK_KP_1, &Mapping{KeyCode: VK_NumPad1, KeyChar: '1', Name: NumPad1Name})
	InsertMapping(C.XK_KP_2, &Mapping{KeyCode: VK_NumPad2, KeyChar: '2', Name: NumPad2Name})
	InsertMapping(C.XK_KP_3, &Mapping{KeyCode: VK_NumPad3, KeyChar: '3', Name: NumPad3Name})
	InsertMapping(C.XK_KP_4, &Mapping{KeyCode: VK_NumPad4, KeyChar: '4', Name: NumPad4Name})
	InsertMapping(C.XK_KP_5, &Mapping{KeyCode: VK_NumPad5, KeyChar: '5', Name: NumPad5Name})
	InsertMapping(C.XK_KP_6, &Mapping{KeyCode: VK_NumPad6, KeyChar: '6', Name: NumPad6Name})
	InsertMapping(C.XK_KP_7, &Mapping{KeyCode: VK_NumPad7, KeyChar: '7', Name: NumPad7Name})
	InsertMapping(C.XK_KP_8, &Mapping{KeyCode: VK_NumPad8, KeyChar: '8', Name: NumPad8Name})
	InsertMapping(C.XK_KP_9, &Mapping{KeyCode: VK_NumPad9, KeyChar: '9', Name: NumPad9Name})
	InsertMapping(C.XK_Num_Lock, &Mapping{KeyCode: VK_NumLock, Name: NumLockName})
	InsertMapping(C.XK_KP_Up, &Mapping{KeyCode: VK_NumPadUp, Name: NumPadUpName})
	InsertMapping(C.XK_KP_Left, &Mapping{KeyCode: VK_NumPadLeft, Name: NumPadLeftName})
	InsertMapping(C.XK_KP_Down, &Mapping{KeyCode: VK_NumPadDown, Name: NumPadDownName})
	InsertMapping(C.XK_KP_Right, &Mapping{KeyCode: VK_NumPadRight, Name: NumPadRightName})
	InsertMapping(C.XK_KP_Begin, &Mapping{KeyCode: VK_NumPadCenter, Name: NumPadCenterName})
	insertKeyCodeMapping(&Mapping{KeyCode: VK_NumPadClear, Name: NumPadClearName}) // Not on a PC keyboard
	InsertMapping(C.XK_KP_Divide, &Mapping{KeyCode: VK_NumPadDivide, KeyChar: '/', Name: NumPadDivideName})
	InsertMapping(C.XK_KP_Multiply, &Mapping{KeyCode: VK_NumPadMultiply, KeyChar: '*', Name: NumPadMultiplyName})
	InsertMapping(C.XK_KP_Subtract, &Mapping{KeyCode: VK_NumPadMinus, KeyChar: '-', Name: NumPadMinusName})
	InsertMapping(C.XK_KP_Add, &Mapping{KeyCode: VK_NumPadAdd, KeyChar: '+', Name: NumPadAddName})
	InsertMapping(C.XK_KP_Decimal, &Mapping{KeyCode: VK_NumPadDecimal, KeyChar: '.', Name: NumPadDecimalName})
	InsertMapping(C.XK_KP_Delete, &Mapping{KeyCode: VK_NumPadDelete, Name: NumPadDeleteName})
	InsertMapping(C.XK_KP_Home, &Mapping{KeyCode: VK_NumPadHome, Name: NumPadHomeName})
	InsertMapping(C.XK_KP_End, &Mapping{KeyCode: VK_NumPadEnd, Name: NumPadEndName})
	InsertMapping(C.XK_KP_Page_Up, &Mapping{KeyCode: VK_NumPadPageUp, Name: NumPadPageUpName})
	InsertMapping(C.XK_KP_Page_Down, &Mapping{KeyCode: VK_NumPadPageDown, Name: NumPadPageDownName})
	InsertMapping(C.XK_F1, &Mapping{KeyCode: VK_F1, Name: F1Name})
	InsertMapping(C.XK_F2, &Mapping{KeyCode: VK_F2, Name: F2Name})
	InsertMapping(C.XK_F3, &Mapping{KeyCode: VK_F3, Name: F3Name})
	InsertMapping(C.XK_F4, &Mapping{KeyCode: VK_F4, Name: F4Name})
	InsertMapping(C.XK_F5, &Mapping{KeyCode: VK_F5, Name: F5Name})
	InsertMapping(C.XK_F6, &Mapping{KeyCode: VK_F6, Name: F6Name})
	InsertMapping(C.XK_F7, &Mapping{KeyCode: VK_F7, Name: F7Name})
	InsertMapping(C.XK_F8, &Mapping{KeyCode: VK_F8, Name: F8Name})
	InsertMapping(C.XK_F9, &Mapping{KeyCode: VK_F9, Name: F9Name})
	InsertMapping(C.XK_F10, &Mapping{KeyCode: VK_F10, Name: F10Name})
	InsertMapping(C.XK_F11, &Mapping{KeyCode: VK_F11, Name: F11Name})
	InsertMapping(C.XK_F12, &Mapping{KeyCode: VK_F12, Name: F12Name})
	// F13 (aka PrtScn) seems to be taken over by the system
	InsertMapping(C.XK_F13, &Mapping{KeyCode: VK_F13, Name: F13Name})
	InsertMapping(C.XK_Scroll_Lock, &Mapping{KeyCode: VK_F14, Name: F14Name})
	InsertMapping(C.XK_Pause, &Mapping{KeyCode: VK_F15, Name: F15Name})
	insertKeyCodeMapping(&Mapping{KeyCode: VK_F16, Name: F16Name}) // Not on a PC keyboard?
	insertKeyCodeMapping(&Mapping{KeyCode: VK_F17, Name: F17Name}) // Not on a PC keyboard?
	insertKeyCodeMapping(&Mapping{KeyCode: VK_F18, Name: F18Name}) // Not on a PC keyboard?
	insertKeyCodeMapping(&Mapping{KeyCode: VK_F19, Name: F19Name}) // Not on a PC keyboard?
}

func insertASCIIKeyCodeMapping(ch int, keyCode int) {
	InsertMapping(ch, &Mapping{KeyCode: keyCode, KeyChar: rune(keyCode), Dynamic: true, Name: string(rune(keyCode))})
}

func insertASCIILetterCodeMapping(keyCode int) {
	mapping := &Mapping{KeyCode: keyCode, KeyChar: rune(keyCode), Dynamic: true, Name: string(rune(keyCode))}
	InsertMapping(keyCode|32, mapping)
	InsertMapping(keyCode, mapping)
}
