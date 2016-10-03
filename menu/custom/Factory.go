// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package custom

import (
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
	cm "github.com/richardwilkes/ui/widget/menu"
)

var (
	bar *cm.MenuBar
)

// AppBar returns the application menu bar.
func AppBar() menu.Bar {
	if bar == nil {
		bar = cm.NewMenuBar()
	}
	return bar
}

// NewMenu creates a new menu.
func NewMenu(title string) menu.Menu {
	return cm.NewMenu(title)
}

// NewItem creates a new item with no key accelerator.
func NewItem(title string, handler event.Handler) menu.Item {
	return cm.NewMenuItem(title, 0, handler)
}

// NewItemWithKey creates a new item with a key accelerator using the platform-default modifiers.
func NewItemWithKey(title string, keyCode int, handler event.Handler) menu.Item {
	return cm.NewMenuItem(title, keyCode, handler)
}

// NewItemWithKeyAndModifiers creates a new item.
func NewItemWithKeyAndModifiers(title string, keyCode int, modifiers keys.Modifiers, handler event.Handler) menu.Item {
	return cm.NewMenuItemWithModifiers(title, keyCode, modifiers, handler)
}

// NewSeparator creates a new separator item.
func NewSeparator() menu.Item {
	return cm.NewSeparator()
}
