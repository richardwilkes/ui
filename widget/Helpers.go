// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package widget

import (
	"github.com/richardwilkes/ui/app"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/menu"
)

// AddAppMenu adds a standard 'application' menu to the menu bar, attaching the aboutAction to the
// About menu item and the preferencesAction to the Preferences menu item.
func AddAppMenu(aboutAction, preferencesAction menu.Action) *menu.Menu {
	name := app.Name()
	appMenu := menu.Bar().AddMenu(name)
	appMenu.AddItem("About "+name, "", aboutAction, nil)
	appMenu.AddSeparator()
	appMenu.AddItem("Preferences…", ",", preferencesAction, nil)
	appMenu.AddSeparator()
	menu.SetServicesMenu(appMenu.AddMenu("Services"))
	appMenu.AddSeparator()
	appMenu.AddItem("Hide "+name, "h", func(item *menu.Item) { app.Hide() }, nil)
	appMenu.AddItem("Hide Others", "h", func(item *menu.Item) { app.HideOthers() }, nil).SetKeyModifiers(event.OptionKeyMask | event.CommandKeyMask)
	appMenu.AddItem("Show All", "", func(item *menu.Item) { app.ShowAll() }, nil)
	appMenu.AddSeparator()
	appMenu.AddItem("Quit "+name, "q", func(item *menu.Item) { app.AttemptQuit() }, nil)
	return appMenu
}

// AddWindowMenu adds a standard 'Window' menu to the menu bar.
func AddWindowMenu() *menu.Menu {
	windowMenu := menu.Bar().AddMenu("Window")
	windowMenu.AddItem("Minimize", "m", func(item *menu.Item) {
		window := KeyWindow()
		if window != nil {
			window.Minimize()
		}
	}, func(item *menu.Item) bool { return KeyWindow() != nil })
	windowMenu.AddItem("Zoom", "\\", func(item *menu.Item) {
		window := KeyWindow()
		if window != nil {
			window.Zoom()
		}
	}, func(item *menu.Item) bool { return KeyWindow() != nil })
	menu.SetWindowMenu(windowMenu)
	windowMenu.AddSeparator()
	windowMenu.AddItem("Bring All to Front", "", func(item *menu.Item) { AllWindowsToFront() }, nil)
	return windowMenu
}

// AddHelpMenu adds a standard 'Help' menu to the menu bar.
func AddHelpMenu() *menu.Menu {
	helpMenu := menu.Bar().AddMenu("Help")
	menu.SetHelpMenu(helpMenu)
	return helpMenu
}
