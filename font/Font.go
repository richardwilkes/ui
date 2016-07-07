// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package font

import (
	"sort"
	"unsafe"
)

// Font holds font information.
type Font struct {
	font unsafe.Pointer
	desc Desc
}

func newFont(desc Desc) *Font {
	return platformNew(desc)
}

// Desc returns a copy of the font's description information.
func (f *Font) Desc() Desc {
	return f.desc
}

func (f *Font) dispose() {
	f.platformDispose()
	f.font = nil
}

// Ascent of the Font, i.e. the distance from the baseline to the top of a typical capital letter.
func (f *Font) Ascent() float32 {
	return f.platformAscent()
}

// Descent of the Font, i.e. the distance from the baseline to the bottom of the typical letter
// that has a descender, such as a lower case 'g'.
func (f *Font) Descent() float32 {
	return f.platformDescent()
}

// Leading of the Font, i.e. the recommended distance between the bottom of the descender line
// to the top of the next line.
func (f *Font) Leading() float32 {
	return f.platformLeading()
}

// PlatformPointer returns the underlying platform data structure for the font.
// Not intended for use outside of the github.com/richardwilkes/go-ui package and its descendants.
func (f *Font) PlatformPointer() unsafe.Pointer {
	return f.font
}

// String -- implements the fmt.Stringer interface.
func (f *Font) String() string {
	return f.desc.String()
}

// AvailableFontFamilies retrieves the names of the installed font families.
func AvailableFontFamilies() []string {
	families := platformAvailableFontFamilies()
	sort.Strings(families)
	return families
}
