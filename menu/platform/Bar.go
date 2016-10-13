// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package platform

import (
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/menu"
)

type Bar struct {
	menu    *platformMenu
	special map[menu.SpecialMenuType]menu.Menu
}

var (
	appBar *Bar
)

// AppBar returns the application menu bar.
func AppBar() menu.Bar {
	if appBar == nil {
		appBar = &Bar{menu: NewMenu(""), special: make(map[menu.SpecialMenuType]menu.Menu)}
		platformSetBar(appBar.menu.menu)
	}
	return appBar
}

// AppendMenu appends a menu at the end of this bar.
func (bar *Bar) AppendMenu(subMenu menu.Menu) {
	bar.InsertMenu(subMenu, -1)
}

// InsertMenu inserts an item with a sub-menu at the specified item index within this menu. Pass
// in a negative index to append to the end.
func (bar *Bar) InsertMenu(subMenu menu.Menu, index int) {
	bar.menu.InsertMenu(subMenu, index)
}

// Remove the item at the specified index from this menu. This does not dispose of the menu item.
func (bar *Bar) Remove(index int) {
	bar.menu.Remove(index)
}

// Count of items in this menu.
func (bar *Bar) Count() int {
	return bar.menu.Count()
}

// Menu at the specified index, or nil.
func (bar *Bar) Menu(index int) menu.Menu {
	item := bar.menu.Item(index)
	return item.SubMenu()
}

// SpecialMenu returns the specified special menu, or nil if it has not been setup.
func (bar *Bar) SpecialMenu(which menu.SpecialMenuType) menu.Menu {
	return bar.special[which]
}

// SetupSpecialMenu sets up the specified special menu, which must have already been installed
// into the menu bar.
func (bar *Bar) SetupSpecialMenu(which menu.SpecialMenuType, mnu menu.Menu) {
	bar.special[which] = mnu
	switch which {
	case menu.ServicesMenu:
		SetServicesMenu(mnu)
	case menu.WindowMenu:
		SetWindowMenu(mnu)
	case menu.HelpMenu:
		SetHelpMenu(mnu)
	}
}

// ProcessKeyDown is called to process KeyDown events prior to anything else receiving them.
func (bar *Bar) ProcessKeyDown(evt *event.KeyDown) {
	// Unused. The native implementation already gets the keys before we do.
}
