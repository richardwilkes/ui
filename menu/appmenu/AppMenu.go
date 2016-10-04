// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package appmenu

import (
	"github.com/richardwilkes/i18n"
	"github.com/richardwilkes/ui/app"
	"github.com/richardwilkes/ui/app/quit"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/menu/factory"
)

// Install adds a standard 'application' menu to the menu bar.
func Install() (appMenu menu.Menu, aboutItem menu.Item, prefsItem menu.Item) {
	name := app.AppName()
	appMenu = factory.NewMenu(name)

	aboutItem = factory.NewItem(i18n.Text("About ")+name, nil)
	appMenu.InsertItem(aboutItem, -1)
	appMenu.InsertItem(factory.NewSeparator(), -1)

	prefsItem = factory.NewItemWithKey(i18n.Text("Preferences…"), keys.VK_Comma, nil)
	appMenu.InsertItem(prefsItem, -1)
	appMenu.InsertItem(factory.NewSeparator(), -1)

	if factory.AddServicesMenu(appMenu) {
		appMenu.InsertItem(factory.NewSeparator(), -1)
	}

	appMenu.InsertItem(factory.NewItemWithKey(i18n.Text("Hide ")+name, keys.VK_H, func(evt event.Event) { app.HideApp() }), -1)
	appMenu.InsertItem(factory.NewItemWithKeyAndModifiers(i18n.Text("Hide Others"), keys.VK_H, keys.OptionModifier|keys.PlatformMenuModifier(), func(evt event.Event) { app.HideOtherApps() }), -1)
	appMenu.InsertItem(factory.NewItem(i18n.Text("Show All"), func(evt event.Event) { app.ShowAllApps() }), -1)
	appMenu.InsertItem(factory.NewSeparator(), -1)

	appMenu.InsertItem(factory.NewItemWithKey(i18n.Text("Quit ")+name, keys.VK_Q, func(evt event.Event) { quit.AttemptQuit() }), -1)

	factory.AppBar().InsertMenu(appMenu, 0)

	return appMenu, aboutItem, prefsItem
}
