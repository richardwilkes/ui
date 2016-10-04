// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package editmenu

import (
	"github.com/richardwilkes/i18n"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/menu/factory"
	"github.com/richardwilkes/ui/widget/window"
)

// Cutable defines the methods required of objects that can respond to the Cut menu item.
type Cutable interface {
	// CanCut returns true if Cut() can be called successfully.
	CanCut() bool
	// Cut the data from the object and copy it to the clipboard.
	Cut()
}

// InsertCutItem adds the standard Cut menu item to the specified menu.
func InsertCutItem(m menu.Menu, index int) menu.Item {
	item := factory.NewItemWithKey(i18n.Text("Cut"), keys.VK_X, func(evt event.Event) {
		wnd := window.KeyWindow()
		if wnd != nil {
			focus := wnd.Focus()
			if c, ok := focus.(Cutable); ok {
				c.Cut()
			}
		}
	})
	item.EventHandlers().Add(event.ValidateType, func(evt event.Event) {
		valid := false
		wnd := window.KeyWindow()
		if wnd != nil {
			focus := wnd.Focus()
			if c, ok := focus.(Cutable); ok {
				valid = c.CanCut()
			}
		}
		if !valid {
			evt.(*event.Validate).MarkInvalid()
		}
	})
	m.InsertItem(item, index)
	return item
}
