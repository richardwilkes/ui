// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

// AddAppMenu adds a standard 'application' menu to the menu bar, attaching the aboutAction to the
// About menu item and the preferencesAction to the Preferences menu item.
func AddAppMenu(aboutAction, preferencesAction MenuAction) *Menu {
	name := AppName()
	appMenu := MenuBar().AddMenu(name)
	appMenu.AddItem("About "+name, "", aboutAction, nil)
	appMenu.AddSeparator()
	appMenu.AddItem("Preferences…", ",", preferencesAction, nil)
	appMenu.AddSeparator()
	SetServicesMenu(appMenu.AddMenu("Services"))
	appMenu.AddSeparator()
	appMenu.AddItem("Hide "+name, "h", func(item *MenuItem) { HideApp() }, nil)
	appMenu.AddItem("Hide Others", "h", func(item *MenuItem) { HideOtherApps() }, nil).SetKeyModifiers(OptionKeyMask | CommandKeyMask)
	appMenu.AddItem("Show All", "", func(item *MenuItem) { ShowAllApps() }, nil)
	appMenu.AddSeparator()
	appMenu.AddItem("Quit "+name, "q", func(item *MenuItem) { AttemptQuit() }, nil)
	return appMenu
}

// AddWindowMenu adds a standard 'Window' menu to the menu bar.
func AddWindowMenu() *Menu {
	windowMenu := MenuBar().AddMenu("Window")
	windowMenu.AddItem("Minimize", "m", func(item *MenuItem) {
		window := KeyWindow()
		if window != nil {
			window.Minimize()
		}
	}, func(item *MenuItem) bool { return KeyWindow() != nil })
	SetWindowMenu(windowMenu)
	windowMenu.AddItem("Zoom", "\\", func(item *MenuItem) {
		window := KeyWindow()
		if window != nil {
			window.Zoom()
		}
	}, func(item *MenuItem) bool { return KeyWindow() != nil })
	SetWindowMenu(windowMenu)
	windowMenu.AddSeparator()
	windowMenu.AddItem("Bring All to Front", "", func(item *MenuItem) { AllWindowsToFront() }, nil)
	return windowMenu
}

// AddHelpMenu adds a standard 'Help' menu to the menu bar.
func AddHelpMenu() *Menu {
	helpMenu := MenuBar().AddMenu("Help")
	SetHelpMenu(helpMenu)
	return helpMenu
}
