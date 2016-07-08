// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package app

import (
	//"github.com/richardwilkes/go-ui/app"
	"github.com/richardwilkes/go-ui/event"
	"github.com/richardwilkes/go-ui/widget"
	"github.com/richardwilkes/go-ui/widget/menu"
)

// AddAppMenu adds a standard 'application' menu to the menu bar, attaching the aboutAction to the
// About menu item and the preferencesAction to the Preferences menu item.
func AddAppMenu(aboutAction, preferencesAction menu.Action) *menu.Menu {
	appMenu := menu.Bar().AddMenu(Name)
	appMenu.AddItem("About "+Name, "", aboutAction, nil)
	appMenu.AddSeparator()
	appMenu.AddItem("Preferences…", ",", preferencesAction, nil)
	appMenu.AddSeparator()
	menu.SetServicesMenu(appMenu.AddMenu("Services"))
	appMenu.AddSeparator()
	appMenu.AddItem("Hide "+Name, "h", func(item *menu.Item) { Hide() }, nil)
	appMenu.AddItem("Hide Others", "h", func(item *menu.Item) { HideOthers() }, nil).SetKeyModifiers(event.OptionKeyMask | event.CommandKeyMask)
	appMenu.AddItem("Show All", "", func(item *menu.Item) { ShowAll() }, nil)
	appMenu.AddSeparator()
	appMenu.AddItem("Quit "+Name, "q", func(item *menu.Item) { AttemptQuit() }, nil)
	return appMenu
}

// AddWindowMenu adds a standard 'Window' menu to the menu bar.
func AddWindowMenu() *menu.Menu {
	windowMenu := menu.Bar().AddMenu("Window")
	windowMenu.AddItem("Minimize", "m", func(item *menu.Item) {
		window := widget.KeyWindow()
		if window != nil {
			window.Minimize()
		}
	}, func(item *menu.Item) bool { return widget.KeyWindow() != nil })
	menu.SetWindowMenu(windowMenu)
	windowMenu.AddItem("Zoom", "\\", func(item *menu.Item) {
		window := widget.KeyWindow()
		if window != nil {
			window.Zoom()
		}
	}, func(item *menu.Item) bool { return widget.KeyWindow() != nil })
	menu.SetWindowMenu(windowMenu)
	windowMenu.AddSeparator()
	windowMenu.AddItem("Bring All to Front", "", func(item *menu.Item) { widget.AllWindowsToFront() }, nil)
	return windowMenu
}

// AddHelpMenu adds a standard 'Help' menu to the menu bar.
func AddHelpMenu() *menu.Menu {
	helpMenu := menu.Bar().AddMenu("Help")
	menu.SetHelpMenu(helpMenu)
	return helpMenu
}
