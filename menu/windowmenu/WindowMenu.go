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
	"github.com/richardwilkes/ui/window"
)

// Install adds a standard 'Window' menu to the end of the menu bar.
func Install(bar menu.Bar) {
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
	windowMenu.AppendItem(item)

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
	windowMenu.AppendItem(item)
	windowMenu.AppendItem(menu.NewSeparator())

	windowMenu.AppendItem(menu.NewItem(i18n.Text("Bring All to Front"), func(evt event.Event) { window.AllWindowsToFront() }))

	bar.AppendMenu(windowMenu)
	bar.SetupSpecialMenu(menu.WindowMenu, windowMenu)
}
