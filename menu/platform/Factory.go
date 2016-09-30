// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package platform

import (
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
)

var (
	menuMap = make(map[cMenu]*platformMenu)
	itemMap = make(map[cItem]*platformItem)
)

// AppBar returns the application menu bar.
func AppBar() menu.Bar {
	if menu, ok := menuMap[platformBar()]; ok {
		return menu
	}
	menu := NewMenu("")
	platformSetBar(menu.(*platformMenu).menu)
	return menu.(*platformMenu)
}

// NewMenu creates a new menu.
func NewMenu(title string) menu.Menu {
	menu := &platformMenu{title: title, menu: platformNewMenu(title)}
	menuMap[menu.menu] = menu
	return menu
}

// NewItem creates a new item with no key accelerator.
func NewItem(title string, handler event.Handler) menu.Item {
	return NewItemWithKey(title, 0, handler)
}

// NewItemWithKey creates a new item with a key accelerator using the platform-default modifiers.
func NewItemWithKey(title string, keyCode int, handler event.Handler) menu.Item {
	return NewItemWithKeyAndModifiers(title, keyCode, keys.PlatformMenuModifier(), handler)
}

// NewItemWithKeyAndModifiers creates a new item.
func NewItemWithKeyAndModifiers(title string, keyCode int, modifiers keys.Modifiers, handler event.Handler) menu.Item {
	item := &platformItem{item: platformNewItem(title, keyCode, modifiers), title: title, keyCode: keyCode, keyModifiers: modifiers, enabled: true}
	if handler != nil {
		item.EventHandlers().Add(event.SelectionType, handler)
	}
	itemMap[item.item] = item
	return item
}

// NewSeparator creates a new separator item.
func NewSeparator() menu.Item {
	item := &platformItem{item: platformNewSeparator()}
	itemMap[item.item] = item
	return item
}
