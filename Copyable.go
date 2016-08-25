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
	"github.com/richardwilkes/ui/menu"
)

// Copyable defines the methods required of objects that can respond to the Copy menu item.
type Copyable interface {
	// CanCopy returns true if Copy() can be called successfully.
	CanCopy() bool
	// Copy the data to the clipboard.
	Copy()
}

// AddCopyItem adds the standard Copy menu item to the specified menu.
func AddCopyItem(m *menu.Menu) *menu.Item {
	item := m.AddItem(i18n.Text("Copy"), "c")
	handlers := item.EventHandlers()
	handlers.Add(event.SelectionType, func(evt event.Event) {
		window := KeyWindow()
		if window != nil {
			focus := window.Focus()
			if c, ok := focus.(Copyable); ok {
				c.Copy()
			}
		}
	})
	handlers.Add(event.ValidateType, func(evt event.Event) {
		valid := false
		window := KeyWindow()
		if window != nil {
			focus := window.Focus()
			if c, ok := focus.(Copyable); ok {
				valid = c.CanCopy()
			}
		}
		if !valid {
			evt.(*event.Validate).MarkInvalid()
		}
	})
	return item
}
