// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package font

type registry struct {
	font  *Font
	count int
}

var (
	descToRegMap  = make(map[Desc]*registry, 0)
	fontToDescMap = make(map[*Font]Desc, 0)
)

// Acquire a font.
func Acquire(desc Desc) *Font {
	var reg *registry
	var ok bool
	if reg, ok = descToRegMap[desc]; !ok {
		reg = &registry{font: newFont(desc)}
		descToRegMap[desc] = reg
		fontToDescMap[reg.font] = desc
	}
	reg.count++
	return reg.font
}

// Release a font.
func (font *Font) Release() {
	if desc, ok := fontToDescMap[font]; ok {
		if reg, ok2 := descToRegMap[desc]; ok2 {
			reg.count--
			if reg.count < 1 {
				delete(descToRegMap, desc)
				delete(fontToDescMap, font)
				font.dispose()
			}
		} else {
			delete(fontToDescMap, font)
			font.dispose()
		}
	}
}
