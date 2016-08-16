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
)

// SelectAll defines the methods required of objects that can respond to the SelectAll menu item.
type SelectAll interface {
	// CanSelectAll returns true if SelectAll() can be called successfully.
	CanSelectAll() bool
	// SelectAll expands the selection to encompass the entire available range.
	SelectAll()
}

// AddSelectAllItem adds the standard Select All menu item to the specified menu.
func AddSelectAllItem(m *Menu) *Item {
	item := m.AddItem(i18n.Text("Select All"), "a")
	handlers := item.EventHandlers()
	handlers.Add(event.SelectionType, func(evt event.Event) {
		window := KeyWindow()
		if window != nil {
			focus := window.Focus()
			if sa, ok := focus.(SelectAll); ok {
				sa.SelectAll()
			}
		}
	})
	handlers.Add(event.ValidateType, func(evt event.Event) {
		valid := false
		window := KeyWindow()
		if window != nil {
			focus := window.Focus()
			if sa, ok := focus.(SelectAll); ok {
				valid = sa.CanSelectAll()
			}
		}
		if !valid {
			evt.(*event.Validate).MarkInvalid()
		}
	})
	return item
}