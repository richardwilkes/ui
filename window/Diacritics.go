// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package window

import (
	"github.com/richardwilkes/ui/keys"
)

var (
	diacriticState int
)

func processDiacritics(keyCode int, ch rune, keyModifiers keys.Modifiers) rune {
	if diacriticState != 0 {
		if keyModifiers&^keys.ShiftModifier == 0 {
			switch ch {
			case 'a':
				switch diacriticState {
				case keys.VK_E:
					ch = 'á'
				case keys.VK_I:
					ch = 'â'
				case keys.VK_Backtick:
					ch = 'à'
				case keys.VK_N:
					ch = 'ã'
				case keys.VK_U:
					ch = 'ä'
				}
			case 'A':
				switch diacriticState {
				case keys.VK_E:
					ch = 'Á'
				case keys.VK_I:
					ch = 'Â'
				case keys.VK_Backtick:
					ch = 'À'
				case keys.VK_N:
					ch = 'Ã'
				case keys.VK_U:
					ch = 'Ä'
				}
			case 'e':
				switch diacriticState {
				case keys.VK_E:
					ch = 'é'
				case keys.VK_I:
					ch = 'ê'
				case keys.VK_Backtick:
					ch = 'è'
				case keys.VK_U:
					ch = 'ë'
				}
			case 'E':
				switch diacriticState {
				case keys.VK_E:
					ch = 'É'
				case keys.VK_I:
					ch = 'Ê'
				case keys.VK_Backtick:
					ch = 'È'
				case keys.VK_U:
					ch = 'Ë'
				}
			case 'i':
				switch diacriticState {
				case keys.VK_E:
					ch = 'í'
				case keys.VK_I:
					ch = 'î'
				case keys.VK_Backtick:
					ch = 'ì'
				case keys.VK_U:
					ch = 'ï'
				}
			case 'I':
				switch diacriticState {
				case keys.VK_E:
					ch = 'Í'
				case keys.VK_I:
					ch = 'Î'
				case keys.VK_Backtick:
					ch = 'Ì'
				case keys.VK_U:
					ch = 'Ï'
				}
			case 'o':
				switch diacriticState {
				case keys.VK_E:
					ch = 'ó'
				case keys.VK_I:
					ch = 'ô'
				case keys.VK_Backtick:
					ch = 'ò'
				case keys.VK_N:
					ch = 'õ'
				case keys.VK_U:
					ch = 'ö'
				}
			case 'O':
				switch diacriticState {
				case keys.VK_E:
					ch = 'Ó'
				case keys.VK_I:
					ch = 'Ô'
				case keys.VK_Backtick:
					ch = 'Ò'
				case keys.VK_N:
					ch = 'Õ'
				case keys.VK_U:
					ch = 'Ö'
				}
			case 'u':
				switch diacriticState {
				case keys.VK_E:
					ch = 'ú'
				case keys.VK_I:
					ch = 'û'
				case keys.VK_Backtick:
					ch = 'ù'
				case keys.VK_U:
					ch = 'ü'
				}
			case 'U':
				switch diacriticState {
				case keys.VK_E:
					ch = 'Ú'
				case keys.VK_I:
					ch = 'Û'
				case keys.VK_Backtick:
					ch = 'Ù'
				case keys.VK_U:
					ch = 'Ü'
				}
			}
		}
		diacriticState = 0
	}
	if keyModifiers&^keys.ShiftModifier == keys.OptionModifier {
		switch keyCode {
		case keys.VK_E, keys.VK_I, keys.VK_Backtick, keys.VK_N, keys.VK_U:
			diacriticState = keyCode
		default:
			diacriticState = 0
		}
	}
	if diacriticState != 0 {
		ch = 0
	}
	return ch
}
