// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package edit

import (
	"github.com/richardwilkes/i18n"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/widget/window"
)

// SelectAll defines the methods required of objects that can respond to the SelectAll menu item.
type SelectAll interface {
	// CanSelectAll returns true if SelectAll() can be called successfully.
	CanSelectAll() bool
	// SelectAll expands the selection to encompass the entire available range.
	SelectAll()
}

// InsertSelectAllItem adds the standard Select All menu item to the specified menu.
func InsertSelectAllItem(m menu.Menu, index int) menu.Item {
	item := menu.NewItemWithKey(i18n.Text("Select All"), keys.VK_A, func(evt event.Event) {
		wnd := window.KeyWindow()
		if wnd != nil {
			focus := wnd.Focus()
			if sa, ok := focus.(SelectAll); ok {
				sa.SelectAll()
			}
		}
	})
	item.EventHandlers().Add(event.ValidateType, func(evt event.Event) {
		valid := false
		wnd := window.KeyWindow()
		if wnd != nil {
			focus := wnd.Focus()
			if sa, ok := focus.(SelectAll); ok {
				valid = sa.CanSelectAll()
			}
		}
		if !valid {
			evt.(*event.Validate).MarkInvalid()
		}
	})
	m.InsertItem(item, index)
	return item
}
