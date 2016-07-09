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
	"github.com/richardwilkes/go-ui/widget"
)

// AddAppMenu adds a standard 'application' menu to the menu bar, attaching the aboutAction to the
// About menu item and the preferencesAction to the Preferences menu item.
func AddAppMenu(aboutAction, preferencesAction widget.MenuAction) *widget.Menu {
	appMenu := widget.MenuBar().AddMenu(Name)
	appMenu.AddItem("About "+Name, "", aboutAction, nil)
	appMenu.AddSeparator()
	appMenu.AddItem("Preferences…", ",", preferencesAction, nil)
	appMenu.AddSeparator()
	widget.SetServicesMenu(appMenu.AddMenu("Services"))
	appMenu.AddSeparator()
	appMenu.AddItem("Hide "+Name, "h", func(item *widget.MenuItem) { Hide() }, nil)
	appMenu.AddItem("Hide Others", "h", func(item *widget.MenuItem) { HideOthers() }, nil).SetKeyModifiers(widget.OptionKeyMask | widget.CommandKeyMask)
	appMenu.AddItem("Show All", "", func(item *widget.MenuItem) { ShowAll() }, nil)
	appMenu.AddSeparator()
	appMenu.AddItem("Quit "+Name, "q", func(item *widget.MenuItem) { AttemptQuit() }, nil)
	return appMenu
}

// AddWindowMenu adds a standard 'Window' menu to the menu bar.
func AddWindowMenu() *widget.Menu {
	windowMenu := widget.MenuBar().AddMenu("Window")
	windowMenu.AddItem("Minimize", "m", func(item *widget.MenuItem) {
		window := widget.KeyWindow()
		if window != nil {
			window.Minimize()
		}
	}, func(item *widget.MenuItem) bool { return widget.KeyWindow() != nil })
	widget.SetWindowMenu(windowMenu)
	windowMenu.AddItem("Zoom", "\\", func(item *widget.MenuItem) {
		window := widget.KeyWindow()
		if window != nil {
			window.Zoom()
		}
	}, func(item *widget.MenuItem) bool { return widget.KeyWindow() != nil })
	widget.SetWindowMenu(windowMenu)
	windowMenu.AddSeparator()
	windowMenu.AddItem("Bring All to Front", "", func(item *widget.MenuItem) { widget.AllWindowsToFront() }, nil)
	return windowMenu
}

// AddHelpMenu adds a standard 'Help' menu to the menu bar.
func AddHelpMenu() *widget.Menu {
	helpMenu := widget.MenuBar().AddMenu("Help")
	widget.SetHelpMenu(helpMenu)
	return helpMenu
}
