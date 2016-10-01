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
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/menu/factory"
	"github.com/richardwilkes/ui/widget"
)

// AddToAppBar adds a standard 'application' menu to the menu bar.
func AddToAppBar() (appMenu menu.Menu, aboutItem menu.Item, prefsItem menu.Item) {
	name := widget.AppName()
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

	appMenu.AddItem(factory.NewItemWithKey(i18n.Text("Hide ")+name, keys.VK_H, func(evt event.Event) { widget.HideApp() }))
	appMenu.AddItem(factory.NewItemWithKeyAndModifiers(i18n.Text("Hide Others"), keys.VK_H, keys.OptionModifier|keys.PlatformMenuModifier(), func(evt event.Event) { widget.HideOtherApps() }))
	appMenu.AddItem(factory.NewItem(i18n.Text("Show All"), func(evt event.Event) { widget.ShowAllApps() }))
	appMenu.AddItem(factory.NewSeparator())

	appMenu.AddItem(factory.NewItemWithKey(i18n.Text("Quit ")+name, keys.VK_Q, func(evt event.Event) { widget.AttemptQuit() }))

	factory.AppBar().AddMenu(appMenu)

	return appMenu, aboutItem, prefsItem
}
