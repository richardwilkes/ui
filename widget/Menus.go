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

// AddAppMenu adds a standard 'application' menu to the menu bar.
func AddAppMenu() (appMenu *menu.Menu, aboutItem *menu.Item, prefsItem *menu.Item) {
	name := app.Name()
	appMenu = menu.Bar().AddMenu(name)
	aboutItem = appMenu.AddItem("About "+name, "")
	appMenu.AddSeparator()
	prefsItem = appMenu.AddItem("Preferences…", ",")
	appMenu.AddSeparator()
	menu.SetServicesMenu(appMenu.AddMenu("Services"))
	appMenu.AddSeparator()
	item := appMenu.AddItem("Hide "+name, "h")
	item.EventHandlers().Add(event.SelectionType, func(evt event.Event) { app.Hide() })
	item = appMenu.AddItem("Hide Others", "h")
	item.SetKeyModifiers(event.OptionKeyMask | event.CommandKeyMask)
	item.EventHandlers().Add(event.SelectionType, func(evt event.Event) { app.HideOthers() })
	item = appMenu.AddItem("Show All", "")
	item.EventHandlers().Add(event.SelectionType, func(evt event.Event) { app.ShowAll() })
	appMenu.AddSeparator()
	item = appMenu.AddItem("Quit "+name, "q")
	item.EventHandlers().Add(event.SelectionType, func(evt event.Event) { app.AttemptQuit() })
	return appMenu, aboutItem, prefsItem
}

// AddWindowMenu adds a standard 'Window' menu to the menu bar.
func AddWindowMenu() *menu.Menu {
	windowMenu := menu.Bar().AddMenu("Window")
	item := windowMenu.AddItem("Minimize", "m")
	handlers := item.EventHandlers()
	handlers.Add(event.SelectionType, func(evt event.Event) {
		window := KeyWindow()
		if window != nil {
			window.Minimize()
		}
	})
	handlers.Add(event.ValidateType, func(evt event.Event) {
		w := KeyWindow()
		if w == nil || !w.Minimizable() {
			evt.(*event.Validate).MarkInvalid()
		}
	})
	item = windowMenu.AddItem("Zoom", "\\")
	handlers = item.EventHandlers()
	handlers.Add(event.SelectionType, func(evt event.Event) {
		window := KeyWindow()
		if window != nil {
			window.Zoom()
		}
	})
	handlers.Add(event.ValidateType, func(evt event.Event) {
		w := KeyWindow()
		if w == nil || !w.Resizable() {
			evt.(*event.Validate).MarkInvalid()
		}
	})
	menu.SetWindowMenu(windowMenu)
	windowMenu.AddSeparator()
	item = windowMenu.AddItem("Bring All to Front", "")
	item.EventHandlers().Add(event.SelectionType, func(evt event.Event) { AllWindowsToFront() })
	return windowMenu
}

// AddHelpMenu adds a standard 'Help' menu to the menu bar.
func AddHelpMenu() *menu.Menu {
	helpMenu := menu.Bar().AddMenu("Help")
	menu.SetHelpMenu(helpMenu)
	return helpMenu
}
