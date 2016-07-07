// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package color

// #cgo darwin LDFLAGS: -framework Cocoa
// #include "SystemColors.h"
import "C"

var (
	Background             Color
	KeyboardFocus          Color
	SelectedControl        Color
	SelectedControlText    Color
	SelectedTextBackground Color
	SelectedText           Color
	TextBackground         Color
	Text                   Color
)

func init() {
	UpdateSystemColors()
}

// UpdateSystemColors updates the system color variables to reflect the current state of the OS.
func UpdateSystemColors() {
	Background = systemColor(C.backgroundColor)
	KeyboardFocus = systemColor(C.keyboardFocusColor)
	SelectedControl = systemColor(C.selectedControlColor)
	SelectedControlText = systemColor(C.selectedControlTextColor)
	SelectedTextBackground = systemColor(C.selectedTextBackgroundColor)
	SelectedText = systemColor(C.selectedTextColor)
	TextBackground = systemColor(C.textBackgroundColor)
	Text = systemColor(C.textColor)
}

func systemColor(id C.SystemColorId) Color {
	return Color(uint32(C.uiGetSystemColor(id)))
}
