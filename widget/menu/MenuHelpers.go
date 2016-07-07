// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package menu

import (
	"github.com/richardwilkes/go-ui/app"
	"github.com/richardwilkes/go-ui/event"
	"github.com/richardwilkes/go-ui/widget"
)

// AddAppMenu adds a standard 'application' menu to the menu bar, attaching the aboutAction to the
// About menu item and the preferencesAction to the Preferences menu item.
func AddAppMenu(aboutAction, preferencesAction Action) *Menu {
	appMenu := Bar().AddMenu(app.Name)
	appMenu.AddItem("About "+app.Name, "", aboutAction, nil)
	appMenu.AddSeparator()
	appMenu.AddItem("Preferences…", ",", preferencesAction, nil)
	appMenu.AddSeparator()
	SetServicesMenu(appMenu.AddMenu("Services"))
	appMenu.AddSeparator()
	appMenu.AddItem("Hide "+app.Name, "h", func(item *Item) { app.Hide() }, nil)
	appMenu.AddItem("Hide Others", "h", func(item *Item) { app.HideOthers() }, nil).SetKeyModifiers(event.OptionKeyMask | event.CommandKeyMask)
	appMenu.AddItem("Show All", "", func(item *Item) { app.ShowAll() }, nil)
	appMenu.AddSeparator()
	appMenu.AddItem("Quit "+app.Name, "q", func(item *Item) { app.AttemptQuit() }, nil)
	return appMenu
}

// AddWindowMenu adds a standard 'Window' menu to the menu bar.
func AddWindowMenu() *Menu {
	windowMenu := Bar().AddMenu("Window")
	windowMenu.AddItem("Minimize", "m", func(item *Item) {
		window := widget.KeyWindow()
		if window != nil {
			window.Minimize()
		}
	}, func(item *Item) bool { return widget.KeyWindow() != nil })
	SetWindowMenu(windowMenu)
	windowMenu.AddItem("Zoom", "\\", func(item *Item) {
		window := widget.KeyWindow()
		if window != nil {
			window.Zoom()
		}
	}, func(item *Item) bool { return widget.KeyWindow() != nil })
	SetWindowMenu(windowMenu)
	windowMenu.AddSeparator()
	windowMenu.AddItem("Bring All to Front", "", func(item *Item) { widget.AllWindowsToFront() }, nil)
	return windowMenu
}

// AddHelpMenu adds a standard 'Help' menu to the menu bar.
func AddHelpMenu() *Menu {
	helpMenu := Bar().AddMenu("Help")
	SetHelpMenu(helpMenu)
	return helpMenu
}
