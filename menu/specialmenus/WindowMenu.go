// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package specialmenus

import (
	"github.com/richardwilkes/i18n"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/widget/window"
)

// InstallWindowMenu adds a standard 'Window' menu to the menu bar.
func InstallWindowMenu(index int) {
	windowMenu := menu.NewMenu(i18n.Text("Window"))

	item := menu.NewItemWithKey(i18n.Text("Minimize"), keys.VK_M, func(evt event.Event) {
		wnd := window.KeyWindow()
		if wnd != nil {
			wnd.Minimize()
		}
	})
	item.EventHandlers().Add(event.ValidateType, func(evt event.Event) {
		w := window.KeyWindow()
		if w == nil || !w.Minimizable() {
			evt.(*event.Validate).MarkInvalid()
		}
	})
	windowMenu.InsertItem(item, -1)

	item = menu.NewItemWithKey(i18n.Text("Zoom"), keys.VK_BackSlash, func(evt event.Event) {
		wnd := window.KeyWindow()
		if wnd != nil {
			wnd.Zoom()
		}
	})
	item.EventHandlers().Add(event.ValidateType, func(evt event.Event) {
		w := window.KeyWindow()
		if w == nil || !w.Resizable() {
			evt.(*event.Validate).MarkInvalid()
		}
	})
	windowMenu.InsertItem(item, -1)
	windowMenu.InsertItem(menu.NewSeparator(), -1)

	windowMenu.InsertItem(menu.NewItem(i18n.Text("Bring All to Front"), func(evt event.Event) { window.AllWindowsToFront() }), -1)

	menu.AppBar().InsertMenu(windowMenu, index)
	menu.SetupSpecialMenu(menu.WindowMenu, windowMenu)
}
