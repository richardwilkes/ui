// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package windowmenu

import (
	"github.com/richardwilkes/i18n"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/menu/factory"
	"github.com/richardwilkes/ui/widget"
)

// AddToAppBar adds a standard 'Window' menu to the menu bar.
func AddToAppBar() menu.Menu {
	windowMenu := factory.NewMenu(i18n.Text("Window"))

	item := factory.NewItemWithKey(i18n.Text("Minimize"), keys.VK_M, func(evt event.Event) {
		window := widget.KeyWindow()
		if window != nil {
			window.Minimize()
		}
	})
	item.EventHandlers().Add(event.ValidateType, func(evt event.Event) {
		w := widget.KeyWindow()
		if w == nil || !w.Minimizable() {
			evt.(*event.Validate).MarkInvalid()
		}
	})
	windowMenu.AddItem(item)

	item = factory.NewItemWithKey(i18n.Text("Zoom"), keys.VK_BackSlash, func(evt event.Event) {
		window := widget.KeyWindow()
		if window != nil {
			window.Zoom()
		}
	})
	item.EventHandlers().Add(event.ValidateType, func(evt event.Event) {
		w := widget.KeyWindow()
		if w == nil || !w.Resizable() {
			evt.(*event.Validate).MarkInvalid()
		}
	})
	windowMenu.AddItem(item)
	windowMenu.AddItem(factory.NewSeparator())

	windowMenu.AddItem(factory.NewItem(i18n.Text("Bring All to Front"), func(evt event.Event) { widget.AllWindowsToFront() }))

	factory.SetWindowMenu(windowMenu)
	factory.AppBar().AddMenu(windowMenu)
	return windowMenu
}
