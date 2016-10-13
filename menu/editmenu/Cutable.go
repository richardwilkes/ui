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
	"github.com/richardwilkes/ui/window"
)

// Cutable defines the methods required of objects that can respond to the Cut menu item.
type Cutable interface {
	// CanCut returns true if Cut() can be called successfully.
	CanCut() bool
	// Cut the data from the object and copy it to the clipboard.
	Cut()
}

// AppendCutItem appends the standard Cut menu item to the specified menu.
func AppendCutItem(m menu.Menu) {
	InsertCutItem(m, -1)
}

// InsertCutItem adds the standard Cut menu item to the specified menu.
func InsertCutItem(m menu.Menu, index int) {
	item := menu.NewItemWithKey(i18n.Text("Cut"), keys.VK_X, Cut)
	item.EventHandlers().Add(event.ValidateType, CanCut)
	m.InsertItem(item, index)
}

// Cut the data from the current keyboard focus and copy it to the clipboard.
func Cut(evt event.Event) {
	wnd := window.KeyWindow()
	if wnd != nil {
		focus := wnd.Focus()
		if c, ok := focus.(Cutable); ok {
			c.Cut()
		}
	}
}

// CanCut returns true if Cut() can be called successfully.
func CanCut(evt event.Event) {
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
}
