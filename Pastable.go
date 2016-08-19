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

// Pastable defines the methods required of objects that can respond to the Paste menu item.
type Pastable interface {
	// CanPaste returns true if Paste() can be called successfully.
	CanPaste() bool
	// Paste the data from the clipboard into the object.
	Paste()
}

// AddPasteItem adds the standard Paste menu item to the specified menu.
func AddPasteItem(m *Menu) *MenuItem {
	item := m.AddItem(i18n.Text("Paste"), "v")
	handlers := item.EventHandlers()
	handlers.Add(event.SelectionType, func(evt event.Event) {
		window := KeyWindow()
		if window != nil {
			focus := window.Focus()
			if p, ok := focus.(Pastable); ok {
				p.Paste()
			}
		}
	})
	handlers.Add(event.ValidateType, func(evt event.Event) {
		valid := false
		window := KeyWindow()
		if window != nil {
			focus := window.Focus()
			if p, ok := focus.(Pastable); ok {
				valid = p.CanPaste()
			}
		}
		if !valid {
			evt.(*event.Validate).MarkInvalid()
		}
	})
	return item
}
