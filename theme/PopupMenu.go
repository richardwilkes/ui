// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package theme

var (
	// StdPopupMenu is the theme all new PopupMenus get by default.
	StdPopupMenu = NewPopupMenu()
)

// PopupMenu contains the theme elements for PopupMenus.
type PopupMenu struct {
	Button
}

// NewPopupMenu creates a new button theme.
func NewPopupMenu() *PopupMenu {
	theme := &PopupMenu{}
	theme.Init()
	return theme
}

// Init initializes the theme with its default values.
func (theme *PopupMenu) Init() {
	theme.Button.Init()
}
