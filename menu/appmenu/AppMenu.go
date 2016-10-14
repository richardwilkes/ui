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
	"runtime"
)

// Install adds a standard 'application' menu to the front of the menu bar.
func Install(bar menu.Bar) (appMenu menu.Menu, aboutItem menu.Item, prefsItem menu.Item) {
	name := app.AppName()
	appMenu = menu.NewMenu(name)

	aboutItem = menu.NewItem(i18n.Text("About ")+name, nil)
	appMenu.AppendItem(aboutItem)

	appMenu.AppendItem(menu.NewSeparator())
	prefsItem = menu.NewItemWithKey(i18n.Text("Preferences…"), keys.VK_Comma, nil)
	appMenu.AppendItem(prefsItem)

	if runtime.GOOS == "darwin" {
		appMenu.AppendItem(menu.NewSeparator())
		servicesMenu := menu.NewMenu(i18n.Text("Services"))
		appMenu.AppendMenu(servicesMenu)
		bar.SetupSpecialMenu(menu.ServicesMenu, servicesMenu)
	}

	appMenu.AppendItem(menu.NewSeparator())
	appMenu.AppendItem(menu.NewItemWithKey(i18n.Text("Hide ")+name, keys.VK_H, func(evt event.Event) { app.HideApp() }))
	if runtime.GOOS == "darwin" {
		appMenu.AppendItem(menu.NewItemWithKeyAndModifiers(i18n.Text("Hide Others"), keys.VK_H, keys.OptionModifier|keys.PlatformMenuModifier(), func(evt event.Event) { app.HideOtherApps() }))
		appMenu.AppendItem(menu.NewItem(i18n.Text("Show All"), func(evt event.Event) { app.ShowAllApps() }))
	}

	appMenu.AppendItem(menu.NewSeparator())
	appMenu.AppendItem(menu.NewItemWithKey(i18n.Text("Quit ")+name, keys.VK_Q, func(evt event.Event) { quit.AttemptQuit() }))

	bar.InsertMenu(appMenu, 0)

	return appMenu, aboutItem, prefsItem
}
