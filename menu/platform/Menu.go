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
	"github.com/richardwilkes/ui/menu"
)

type platformMenu struct {
	menu  cMenu // Must be first element in struct!
	title string
}

// AddItem appends an item to the end of this menu.
func (menu *platformMenu) AddItem(item menu.Item) {
	menu.platformAddItem(item.(*platformItem).item)
}

// AddMenu appends an item with a sub-menu to the end of this menu.
func (menu *platformMenu) AddMenu(subMenu menu.Menu) {
	item := NewItem(subMenu.Title(), nil)
	item.(*platformItem).platformSetSubMenu(subMenu.(*platformMenu).menu)
	menu.AddItem(item)
}

// InsertItem inserts an item at the specified item index within this menu.
func (menu *platformMenu) InsertItem(index int, item menu.Item) {
	if index < 0 {
		index = 0
	} else {
		max := menu.Count()
		if index > max {
			index = max
		}
	}
	menu.platformInsertItem(index, item.(*platformItem).item)
}

// InsertMenu inserts an item with a sub-menu at the specified item index within this menu.
func (menu *platformMenu) InsertMenu(index int, subMenu menu.Menu) {
	item := NewItem(subMenu.Title(), nil)
	item.(*platformItem).platformSetSubMenu(subMenu.(*platformMenu).menu)
	menu.InsertItem(index, item)
}

// Remove the item at the specified index from this menu. This does not dispose of the menu item.
func (menu *platformMenu) Remove(index int) {
	if index >= 0 && index < menu.Count() {
		menu.platformRemove(index)
	}
}

// Title returns the title of this menu.
func (menu *platformMenu) Title() string {
	return menu.title
}

// Count of items in this menu.
func (menu *platformMenu) Count() int {
	return menu.platformItemCount()
}

// Item at the specified index, or nil.
func (menu *platformMenu) Item(index int) menu.Item {
	if item, ok := itemMap[menu.platformItem(index)]; ok {
		return item
	}
	return nil
}

// Dispose releases any operating system resources associated with this menu. It will also
// call Dispose() on all menu items it contains.
func (menu *platformMenu) Dispose() {
	if _, ok := menuMap[menu.menu]; ok {
		count := menu.Count()
		for i := 0; i < count; i++ {
			menu.Item(i).Dispose()
		}
		delete(menuMap, menu.menu)
		menu.platformDispose()
	}
}

// Menu at the specified index, or nil.
func (menu *platformMenu) Menu(index int) menu.Menu {
	item := menu.Item(index)
	return item.SubMenu()
}
