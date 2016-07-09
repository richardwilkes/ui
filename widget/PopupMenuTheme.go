// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package widget

var (
	// StdPopupMenuTheme is the theme all new PopupMenus get by default.
	StdPopupMenuTheme = NewPopupMenuTheme()
)

// PopupMenuTheme contains the theme elements for PopupMenus.
type PopupMenuTheme struct {
	ButtonTheme
}

// NewPopupMenuTheme creates a new button theme.
func NewPopupMenuTheme() *PopupMenuTheme {
	theme := &PopupMenuTheme{}
	theme.Init()
	return theme
}

// Init initializes the theme with its default values.
func (theme *PopupMenuTheme) Init() {
	theme.ButtonTheme.Init()
}
