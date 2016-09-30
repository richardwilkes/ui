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
	"github.com/richardwilkes/i18n"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/menu/factory"
)

// AddAppMenu adds a standard 'application' menu to the menu bar.
func AddAppMenu() (appMenu menu.Menu, aboutItem menu.Item, prefsItem menu.Item) {
	name := AppName()
	appMenu = factory.NewMenu(name)

	aboutItem = factory.NewItem(i18n.Text("About ")+name, nil)
	appMenu.AddItem(aboutItem)
	appMenu.AddItem(factory.NewSeparator())

	prefsItem = factory.NewItemWithKey(i18n.Text("Preferences…"), keys.VK_Comma, nil)
	appMenu.AddItem(prefsItem)
	appMenu.AddItem(factory.NewSeparator())

	if factory.AddServicesMenu(appMenu) {
		appMenu.AddItem(factory.NewSeparator())
	}

	appMenu.AddItem(factory.NewItemWithKey(i18n.Text("Hide ")+name, keys.VK_H, func(evt event.Event) { HideApp() }))
	appMenu.AddItem(factory.NewItemWithKeyAndModifiers(i18n.Text("Hide Others"), keys.VK_H, keys.OptionModifier|keys.PlatformMenuModifier(), func(evt event.Event) { HideOtherApps() }))
	appMenu.AddItem(factory.NewItem(i18n.Text("Show All"), func(evt event.Event) { ShowAllApps() }))
	appMenu.AddItem(factory.NewSeparator())

	appMenu.AddItem(factory.NewItemWithKey(i18n.Text("Quit ")+name, keys.VK_Q, func(evt event.Event) { AttemptQuit() }))

	factory.AppBar().AddMenu(appMenu)

	return appMenu, aboutItem, prefsItem
}

// AddWindowMenu adds a standard 'Window' menu to the menu bar.
func AddWindowMenu() menu.Menu {
	windowMenu := factory.NewMenu(i18n.Text("Window"))

	item := factory.NewItemWithKey(i18n.Text("Minimize"), keys.VK_M, func(evt event.Event) {
		window := KeyWindow()
		if window != nil {
			window.Minimize()
		}
	})
	item.EventHandlers().Add(event.ValidateType, func(evt event.Event) {
		w := KeyWindow()
		if w == nil || !w.Minimizable() {
			evt.(*event.Validate).MarkInvalid()
		}
	})
	windowMenu.AddItem(item)

	item = factory.NewItemWithKey(i18n.Text("Zoom"), keys.VK_BackSlash, func(evt event.Event) {
		window := KeyWindow()
		if window != nil {
			window.Zoom()
		}
	})
	item.EventHandlers().Add(event.ValidateType, func(evt event.Event) {
		w := KeyWindow()
		if w == nil || !w.Resizable() {
			evt.(*event.Validate).MarkInvalid()
		}
	})
	windowMenu.AddItem(item)
	windowMenu.AddItem(factory.NewSeparator())

	windowMenu.AddItem(factory.NewItem(i18n.Text("Bring All to Front"), func(evt event.Event) { AllWindowsToFront() }))

	factory.SetWindowMenu(windowMenu)
	factory.AppBar().AddMenu(windowMenu)
	return windowMenu
}

// AddHelpMenu adds a standard 'Help' menu to the menu bar.
func AddHelpMenu() menu.Menu {
	helpMenu := factory.NewMenu(i18n.Text("Help"))
	factory.SetHelpMenu(helpMenu)
	factory.AppBar().AddMenu(helpMenu)
	return helpMenu
}
