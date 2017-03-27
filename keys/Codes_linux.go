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

func init() {
	InsertMapping(C.XK_Up, &Mapping{KeyCode: VirtualKeyUp, Name: UpName})
	InsertMapping(C.XK_Left, &Mapping{KeyCode: VirtualKeyLeft, Name: LeftName})
	InsertMapping(C.XK_Down, &Mapping{KeyCode: VirtualKeyDown, Name: DownName})
	InsertMapping(C.XK_Right, &Mapping{KeyCode: VirtualKeyRight, Name: RightName})
	InsertMapping(C.XK_Insert, &Mapping{KeyCode: VirtualKeyInsert, Name: InsertName})
	InsertMapping(C.XK_Home, &Mapping{KeyCode: VirtualKeyHome, Name: HomeName})
	InsertMapping(C.XK_End, &Mapping{KeyCode: VirtualKeyEnd, Name: EndName})
	InsertMapping(C.XK_BackSpace, &Mapping{KeyCode: VirtualKeyBackspace, Name: BackspaceName})
	InsertMapping(C.XK_Tab, &Mapping{KeyCode: VirtualKeyTab, KeyChar: '\t', Name: TabName})
	InsertMapping(C.XK_ISO_Left_Tab, &Mapping{KeyCode: VirtualKeyTab, KeyChar: '\t', Name: TabName})
	InsertMapping(C.XK_Return, &Mapping{KeyCode: VirtualKeyReturn, KeyChar: '\n', Name: ReturnName})
	InsertMapping(C.XK_KP_Enter, &Mapping{KeyCode: VirtualKeyNumPadEnter, KeyChar: '\n', Name: NumPadEnterName})
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyEject, Name: EjectName}) // Not a PC keyboard
	InsertMapping(C.XK_Shift_L, &Mapping{KeyCode: VirtualKeyShiftLeft, Name: LeftShiftName})
	InsertMapping(C.XK_Shift_R, &Mapping{KeyCode: VirtualKeyShiftRight, Name: RightShiftName})
	InsertMapping(C.XK_Control_L, &Mapping{KeyCode: VirtualKeyControlLeft, Name: LeftControlName})
	InsertMapping(C.XK_Control_R, &Mapping{KeyCode: VirtualKeyControlRight, Name: RightControlName})
	InsertMapping(C.XK_Alt_L, &Mapping{KeyCode: VirtualKeyOptionLeft, Name: LeftOptionName})
	InsertMapping(C.XK_Alt_R, &Mapping{KeyCode: VirtualKeyOptionRight, Name: RightOptionName})
	InsertMapping(C.XK_Super_L, &Mapping{KeyCode: VirtualKeyCommandLeft, Name: LeftWindowsName})
	InsertMapping(C.XK_Super_R, &Mapping{KeyCode: VirtualKeyCommandRight, Name: RightWindowsName})
	InsertMapping(C.XK_Caps_Lock, &Mapping{KeyCode: VirtualKeyCapsLock, Name: CapsLockName})
	InsertMapping(C.XK_Menu, &Mapping{KeyCode: VirtualKeyMenu, Name: MenuName})
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyFn, Name: FnName}) // Not on a PC keyboard
	InsertMapping(C.XK_Escape, &Mapping{KeyCode: VirtualKeyEscape, KeyChar: '\x1b', Name: EscapeName})
	InsertMapping(C.XK_Page_Up, &Mapping{KeyCode: VirtualKeyPageUp, Name: PageUpName})
	InsertMapping(C.XK_Page_Down, &Mapping{KeyCode: VirtualKeyPageDown, Name: PageDownName})
	InsertMapping(' ', &Mapping{KeyCode: VirtualKeySpace, KeyChar: ' ', Name: SpaceName})
	insertASCIIKeyCodeMapping('\'', VirtualKeyQuote)
	insertASCIIKeyCodeMapping(',', VirtualKeyComma)
	insertASCIIKeyCodeMapping('-', VirtualKeyMinus)
	insertASCIIKeyCodeMapping('.', VirtualKeyPeriod)
	insertASCIIKeyCodeMapping('/', VirtualKeySlash)
	insertASCIIKeyCodeMapping('0', VirtualKey0)
	insertASCIIKeyCodeMapping('1', VirtualKey1)
	insertASCIIKeyCodeMapping('2', VirtualKey2)
	insertASCIIKeyCodeMapping('3', VirtualKey3)
	insertASCIIKeyCodeMapping('4', VirtualKey4)
	insertASCIIKeyCodeMapping('5', VirtualKey5)
	insertASCIIKeyCodeMapping('6', VirtualKey6)
	insertASCIIKeyCodeMapping('7', VirtualKey7)
	insertASCIIKeyCodeMapping('8', VirtualKey8)
	insertASCIIKeyCodeMapping('9', VirtualKey9)
	insertASCIIKeyCodeMapping(';', VirtualKeySemiColon)
	insertASCIIKeyCodeMapping('=', VirtualKeyEqual)
	insertASCIILetterCodeMapping(VirtualKeyA)
	insertASCIILetterCodeMapping(VirtualKeyB)
	insertASCIILetterCodeMapping(VirtualKeyC)
	insertASCIILetterCodeMapping(VirtualKeyD)
	insertASCIILetterCodeMapping(VirtualKeyE)
	insertASCIILetterCodeMapping(VirtualKeyF)
	insertASCIILetterCodeMapping(VirtualKeyG)
	insertASCIILetterCodeMapping(VirtualKeyH)
	insertASCIILetterCodeMapping(VirtualKeyI)
	insertASCIILetterCodeMapping(VirtualKeyJ)
	insertASCIILetterCodeMapping(VirtualKeyK)
	insertASCIILetterCodeMapping(VirtualKeyL)
	insertASCIILetterCodeMapping(VirtualKeyM)
	insertASCIILetterCodeMapping(VirtualKeyN)
	insertASCIILetterCodeMapping(VirtualKeyO)
	insertASCIILetterCodeMapping(VirtualKeyP)
	insertASCIILetterCodeMapping(VirtualKeyQ)
	insertASCIILetterCodeMapping(VirtualKeyR)
	insertASCIILetterCodeMapping(VirtualKeyS)
	insertASCIILetterCodeMapping(VirtualKeyT)
	insertASCIILetterCodeMapping(VirtualKeyU)
	insertASCIILetterCodeMapping(VirtualKeyV)
	insertASCIILetterCodeMapping(VirtualKeyW)
	insertASCIILetterCodeMapping(VirtualKeyX)
	insertASCIILetterCodeMapping(VirtualKeyY)
	insertASCIILetterCodeMapping(VirtualKeyZ)
	insertASCIIKeyCodeMapping('[', VirtualKeyLeftBracket)
	insertASCIIKeyCodeMapping('\\', VirtualKeyBackSlash)
	insertASCIIKeyCodeMapping(']', VirtualKeyRightBracket)
	insertASCIIKeyCodeMapping('`', VirtualKeyBacktick)
	InsertMapping(C.XK_Delete, &Mapping{KeyCode: VirtualKeyDelete, Name: DeleteName})
	InsertMapping(C.XK_KP_0, &Mapping{KeyCode: VirtualKeyNumPad0, KeyChar: '0', Name: NumPad0Name})
	InsertMapping(C.XK_KP_1, &Mapping{KeyCode: VirtualKeyNumPad1, KeyChar: '1', Name: NumPad1Name})
	InsertMapping(C.XK_KP_2, &Mapping{KeyCode: VirtualKeyNumPad2, KeyChar: '2', Name: NumPad2Name})
	InsertMapping(C.XK_KP_3, &Mapping{KeyCode: VirtualKeyNumPad3, KeyChar: '3', Name: NumPad3Name})
	InsertMapping(C.XK_KP_4, &Mapping{KeyCode: VirtualKeyNumPad4, KeyChar: '4', Name: NumPad4Name})
	InsertMapping(C.XK_KP_5, &Mapping{KeyCode: VirtualKeyNumPad5, KeyChar: '5', Name: NumPad5Name})
	InsertMapping(C.XK_KP_6, &Mapping{KeyCode: VirtualKeyNumPad6, KeyChar: '6', Name: NumPad6Name})
	InsertMapping(C.XK_KP_7, &Mapping{KeyCode: VirtualKeyNumPad7, KeyChar: '7', Name: NumPad7Name})
	InsertMapping(C.XK_KP_8, &Mapping{KeyCode: VirtualKeyNumPad8, KeyChar: '8', Name: NumPad8Name})
	InsertMapping(C.XK_KP_9, &Mapping{KeyCode: VirtualKeyNumPad9, KeyChar: '9', Name: NumPad9Name})
	InsertMapping(C.XK_Num_Lock, &Mapping{KeyCode: VirtualKeyNumLock, Name: NumLockName})
	InsertMapping(C.XK_KP_Up, &Mapping{KeyCode: VirtualKeyNumPadUp, Name: NumPadUpName})
	InsertMapping(C.XK_KP_Left, &Mapping{KeyCode: VirtualKeyNumPadLeft, Name: NumPadLeftName})
	InsertMapping(C.XK_KP_Down, &Mapping{KeyCode: VirtualKeyNumPadDown, Name: NumPadDownName})
	InsertMapping(C.XK_KP_Right, &Mapping{KeyCode: VirtualKeyNumPadRight, Name: NumPadRightName})
	InsertMapping(C.XK_KP_Begin, &Mapping{KeyCode: VirtualKeyNumPadCenter, Name: NumPadCenterName})
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyNumPadClear, Name: NumPadClearName}) // Not on a PC keyboard
	InsertMapping(C.XK_KP_Divide, &Mapping{KeyCode: VirtualKeyNumPadDivide, KeyChar: '/', Name: NumPadDivideName})
	InsertMapping(C.XK_KP_Multiply, &Mapping{KeyCode: VirtualKeyNumPadMultiply, KeyChar: '*', Name: NumPadMultiplyName})
	InsertMapping(C.XK_KP_Subtract, &Mapping{KeyCode: VirtualKeyNumPadMinus, KeyChar: '-', Name: NumPadMinusName})
	InsertMapping(C.XK_KP_Add, &Mapping{KeyCode: VirtualKeyNumPadAdd, KeyChar: '+', Name: NumPadAddName})
	InsertMapping(C.XK_KP_Decimal, &Mapping{KeyCode: VirtualKeyNumPadDecimal, KeyChar: '.', Name: NumPadDecimalName})
	InsertMapping(C.XK_KP_Delete, &Mapping{KeyCode: VirtualKeyNumPadDelete, Name: NumPadDeleteName})
	InsertMapping(C.XK_KP_Home, &Mapping{KeyCode: VirtualKeyNumPadHome, Name: NumPadHomeName})
	InsertMapping(C.XK_KP_End, &Mapping{KeyCode: VirtualKeyNumPadEnd, Name: NumPadEndName})
	InsertMapping(C.XK_KP_Page_Up, &Mapping{KeyCode: VirtualKeyNumPadPageUp, Name: NumPadPageUpName})
	InsertMapping(C.XK_KP_Page_Down, &Mapping{KeyCode: VirtualKeyNumPadPageDown, Name: NumPadPageDownName})
	InsertMapping(C.XK_F1, &Mapping{KeyCode: VirtualKeyF1, Name: F1Name})
	InsertMapping(C.XK_F2, &Mapping{KeyCode: VirtualKeyF2, Name: F2Name})
	InsertMapping(C.XK_F3, &Mapping{KeyCode: VirtualKeyF3, Name: F3Name})
	InsertMapping(C.XK_F4, &Mapping{KeyCode: VirtualKeyF4, Name: F4Name})
	InsertMapping(C.XK_F5, &Mapping{KeyCode: VirtualKeyF5, Name: F5Name})
	InsertMapping(C.XK_F6, &Mapping{KeyCode: VirtualKeyF6, Name: F6Name})
	InsertMapping(C.XK_F7, &Mapping{KeyCode: VirtualKeyF7, Name: F7Name})
	InsertMapping(C.XK_F8, &Mapping{KeyCode: VirtualKeyF8, Name: F8Name})
	InsertMapping(C.XK_F9, &Mapping{KeyCode: VirtualKeyF9, Name: F9Name})
	InsertMapping(C.XK_F10, &Mapping{KeyCode: VirtualKeyF10, Name: F10Name})
	InsertMapping(C.XK_F11, &Mapping{KeyCode: VirtualKeyF11, Name: F11Name})
	InsertMapping(C.XK_F12, &Mapping{KeyCode: VirtualKeyF12, Name: F12Name})
	// F13 (aka PrtScn) seems to be taken over by the system
	InsertMapping(C.XK_F13, &Mapping{KeyCode: VirtualKeyF13, Name: F13Name})
	InsertMapping(C.XK_Scroll_Lock, &Mapping{KeyCode: VirtualKeyF14, Name: F14Name})
	InsertMapping(C.XK_Pause, &Mapping{KeyCode: VirtualKeyF15, Name: F15Name})
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyF16, Name: F16Name}) // Not on a PC keyboard?
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyF17, Name: F17Name}) // Not on a PC keyboard?
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyF18, Name: F18Name}) // Not on a PC keyboard?
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyF19, Name: F19Name}) // Not on a PC keyboard?
}

func insertASCIIKeyCodeMapping(ch int, keyCode int) {
	InsertMapping(ch, &Mapping{KeyCode: keyCode, KeyChar: rune(keyCode), Dynamic: true, Name: string(rune(keyCode))})
}

func insertASCIILetterCodeMapping(keyCode int) {
	mapping := &Mapping{KeyCode: keyCode, KeyChar: rune(keyCode), Dynamic: true, Name: string(rune(keyCode))}
	InsertMapping(keyCode|32, mapping)
	InsertMapping(keyCode, mapping)
}
