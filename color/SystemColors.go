// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package color

var (
	// Background is the system color used for the window background.
	Background = RGB(236, 236, 236)
	// KeyboardFocus is the system color used to highlight controls that have the keyboard
	// focus.
	KeyboardFocus Color
	// SelectedControl is the system color used to highlight controls that have a selection.
	SelectedControl Color
	// SelectedControlText is the system color used for text in the selected portion of a control.
	SelectedControlText = Black
	// SelectedTextBackground is the system color used for the background of selected text.
	SelectedTextBackground Color
	// SelectedText is the system color used for selected text.
	SelectedText = White
	// TextBackground is the system color used for the background of editable text areas.
	TextBackground = White
	// Text is the system color used for the text in editable text areas.
	Text = Black
)
