// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package custom

import (
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
)

var (
	windowMenu menu.Menu
	helpMenu   menu.Menu
)

func Install() {
	menu.AppBar = func() menu.Bar { return AppBar() }
	menu.SpecialMenu = SpecialMenu
	menu.SetupSpecialMenu = SetupSpecialMenu
	menu.NewMenu = func(title string) menu.Menu { return NewMenu(title) }
	menu.NewItem = func(title string, handler event.Handler) menu.Item { return NewItem(title, handler) }
	menu.NewItemWithKey = func(title string, keyCode int, handler event.Handler) menu.Item {
		return NewItemWithKey(title, keyCode, handler)
	}
	menu.NewItemWithKeyAndModifiers = func(title string, keyCode int, modifiers keys.Modifiers, handler event.Handler) menu.Item {
		return NewItemWithKeyAndModifiers(title, keyCode, modifiers, handler)
	}
	menu.NewSeparator = func() menu.Item { return NewSeparator() }
}

func SpecialMenu(which menu.SpecialMenuType) menu.Menu {
	switch which {
	case menu.WindowMenu:
		return windowMenu
	case menu.HelpMenu:
		return helpMenu
	default:
		return nil
	}
}

func SetupSpecialMenu(which menu.SpecialMenuType, mnu menu.Menu) {
	switch which {
	case menu.WindowMenu:
		windowMenu = mnu
	case menu.HelpMenu:
		helpMenu = mnu
	}
}
