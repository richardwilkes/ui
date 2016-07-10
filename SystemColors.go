// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

// #cgo darwin LDFLAGS: -framework Cocoa
// #include "SystemColors.h"
import "C"

var (
	// BackgroundColor is the system color used for the window background.
	BackgroundColor = systemColor(C.backgroundColor)
	// KeyboardFocusColor is the system color used to highlight controls that have the keyboard
	// focus.
	KeyboardFocusColor = systemColor(C.keyboardFocusColor)
	// SelectedControlColor is the system color used to highlight controls that have a selection.
	SelectedControlColor = systemColor(C.selectedControlColor)
	// SelectedControlTextColor is the system color used for text in the selected portion of a
	// control.
	SelectedControlTextColor = systemColor(C.selectedControlTextColor)
	// SelectedTextBackgroundColor is the system color used for the background of selected text.
	SelectedTextBackgroundColor = systemColor(C.selectedTextBackgroundColor)
	// SelectedTextColor is the system color used for selected text.
	SelectedTextColor = systemColor(C.selectedTextColor)
	// TextBackgroundColor is the system color used for the background of editable text areas.
	TextBackgroundColor = systemColor(C.textBackgroundColor)
	// TextColor is the system color used for the text in editable text areas.
	TextColor = systemColor(C.textColor)
)

func systemColor(id C.SystemColorId) Color {
	return Color(uint32(C.uiGetSystemColor(id)))
}

// UpdateSystemColors updates the system color variables to reflect the current state of the OS.
func UpdateSystemColors() {
	BackgroundColor = systemColor(C.backgroundColor)
	KeyboardFocusColor = systemColor(C.keyboardFocusColor)
	SelectedControlColor = systemColor(C.selectedControlColor)
	SelectedControlTextColor = systemColor(C.selectedControlTextColor)
	SelectedTextBackgroundColor = systemColor(C.selectedTextBackgroundColor)
	SelectedTextColor = systemColor(C.selectedTextColor)
	TextBackgroundColor = systemColor(C.textBackgroundColor)
	TextColor = systemColor(C.textColor)
}
