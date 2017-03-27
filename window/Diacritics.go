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
				case keys.VirtualKeyE:
					ch = 'á'
				case keys.VirtualKeyI:
					ch = 'â'
				case keys.VirtualKeyBacktick:
					ch = 'à'
				case keys.VirtualKeyN:
					ch = 'ã'
				case keys.VirtualKeyU:
					ch = 'ä'
				}
			case 'A':
				switch diacriticState {
				case keys.VirtualKeyE:
					ch = 'Á'
				case keys.VirtualKeyI:
					ch = 'Â'
				case keys.VirtualKeyBacktick:
					ch = 'À'
				case keys.VirtualKeyN:
					ch = 'Ã'
				case keys.VirtualKeyU:
					ch = 'Ä'
				}
			case 'e':
				switch diacriticState {
				case keys.VirtualKeyE:
					ch = 'é'
				case keys.VirtualKeyI:
					ch = 'ê'
				case keys.VirtualKeyBacktick:
					ch = 'è'
				case keys.VirtualKeyU:
					ch = 'ë'
				}
			case 'E':
				switch diacriticState {
				case keys.VirtualKeyE:
					ch = 'É'
				case keys.VirtualKeyI:
					ch = 'Ê'
				case keys.VirtualKeyBacktick:
					ch = 'È'
				case keys.VirtualKeyU:
					ch = 'Ë'
				}
			case 'i':
				switch diacriticState {
				case keys.VirtualKeyE:
					ch = 'í'
				case keys.VirtualKeyI:
					ch = 'î'
				case keys.VirtualKeyBacktick:
					ch = 'ì'
				case keys.VirtualKeyU:
					ch = 'ï'
				}
			case 'I':
				switch diacriticState {
				case keys.VirtualKeyE:
					ch = 'Í'
				case keys.VirtualKeyI:
					ch = 'Î'
				case keys.VirtualKeyBacktick:
					ch = 'Ì'
				case keys.VirtualKeyU:
					ch = 'Ï'
				}
			case 'o':
				switch diacriticState {
				case keys.VirtualKeyE:
					ch = 'ó'
				case keys.VirtualKeyI:
					ch = 'ô'
				case keys.VirtualKeyBacktick:
					ch = 'ò'
				case keys.VirtualKeyN:
					ch = 'õ'
				case keys.VirtualKeyU:
					ch = 'ö'
				}
			case 'O':
				switch diacriticState {
				case keys.VirtualKeyE:
					ch = 'Ó'
				case keys.VirtualKeyI:
					ch = 'Ô'
				case keys.VirtualKeyBacktick:
					ch = 'Ò'
				case keys.VirtualKeyN:
					ch = 'Õ'
				case keys.VirtualKeyU:
					ch = 'Ö'
				}
			case 'u':
				switch diacriticState {
				case keys.VirtualKeyE:
					ch = 'ú'
				case keys.VirtualKeyI:
					ch = 'û'
				case keys.VirtualKeyBacktick:
					ch = 'ù'
				case keys.VirtualKeyU:
					ch = 'ü'
				}
			case 'U':
				switch diacriticState {
				case keys.VirtualKeyE:
					ch = 'Ú'
				case keys.VirtualKeyI:
					ch = 'Û'
				case keys.VirtualKeyBacktick:
					ch = 'Ù'
				case keys.VirtualKeyU:
					ch = 'Ü'
				}
			}
		}
		diacriticState = 0
	}
	if keyModifiers&^keys.ShiftModifier == keys.OptionModifier {
		switch keyCode {
		case keys.VirtualKeyE, keys.VirtualKeyI, keys.VirtualKeyBacktick, keys.VirtualKeyN, keys.VirtualKeyU:
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
