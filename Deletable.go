// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

import (
	"github.com/richardwilkes/i18n"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/menu/factory"
)

// Deletable defines the methods required of objects that can respond to the Delete menu item.
type Deletable interface {
	// CanDelete returns true if Delete() can be called successfully.
	CanDelete() bool
	// Delete the data from the object and copy it to the clipboard.
	Delete()
}

// AddDeleteItem adds the standard Delete menu item to the specified menu.
func AddDeleteItem(m menu.Menu) menu.Item {
	item := factory.NewItemWithKeyAndModifiers(i18n.Text("Delete"), keys.VK_Backspace, 0, func(evt event.Event) {
		window := KeyWindow()
		if window != nil {
			focus := window.Focus()
			if c, ok := focus.(Deletable); ok {
				c.Delete()
			}
		}
	})
	item.EventHandlers().Add(event.ValidateType, func(evt event.Event) {
		valid := false
		window := KeyWindow()
		if window != nil {
			focus := window.Focus()
			if c, ok := focus.(Deletable); ok {
				valid = c.CanDelete()
			}
		}
		if !valid {
			evt.(*event.Validate).MarkInvalid()
		}
	})
	m.AddItem(item)
	return item
}
