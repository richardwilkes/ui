// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package macmenus

import (
	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/window"
)

// NewMenu creates a new menu.
func NewMenu(title string) menu.Menu {
	mnu := &platformMenu{title: title, menu: platformNewMenu(title)}
	menuMap[mnu.menu] = mnu
	return mnu
}

// AppendItem appends an item at the end of this menu.
func (mnu *platformMenu) AppendItem(item menu.Item) {
	mnu.InsertItem(item, -1)
}

// InsertItem inserts an item at the specified item index within this menu. Pass in a negative
// index to append to the end.
func (mnu *platformMenu) InsertItem(item menu.Item, index int) {
	max := mnu.Count()
	if index < 0 || index > max {
		index = max
	}
	mnu.platformInsertItem(item.(*platformItem).item, index)
}

// AppendMenu appends an item with a sub-menu at the end of this menu.
func (mnu *platformMenu) AppendMenu(subMenu menu.Menu) {
	mnu.InsertMenu(subMenu, -1)
}

// InsertMenu inserts an item with a sub-menu at the specified item index within this menu. Pass
// in a negative index to append to the end.
func (mnu *platformMenu) InsertMenu(subMenu menu.Menu, index int) {
	item := NewItem(subMenu.Title(), nil)
	item.(*platformItem).platformSetSubMenu(subMenu.(*platformMenu).menu)
	mnu.InsertItem(item, index)
}

// Remove the item at the specified index from this menu. This does not dispose of the menu item.
func (mnu *platformMenu) Remove(index int) {
	if index >= 0 && index < mnu.Count() {
		mnu.platformRemove(index)
	}
}

// Title returns the title of this menu.
func (mnu *platformMenu) Title() string {
	return mnu.title
}

// Count of items in this menu.
func (mnu *platformMenu) Count() int {
	return mnu.platformItemCount()
}

// Item at the specified index, or nil.
func (mnu *platformMenu) Item(index int) menu.Item {
	if item, ok := itemMap[mnu.platformItem(index)]; ok {
		return item
	}
	return nil
}

// Dispose releases any operating system resources associated with this menu. It will also
// call Dispose() on all menu items it contains.
func (mnu *platformMenu) Dispose() {
	if _, ok := menuMap[mnu.menu]; ok {
		count := mnu.Count()
		for i := 0; i < count; i++ {
			mnu.Item(i).Dispose()
		}
		delete(menuMap, mnu.menu)
		mnu.platformDispose()
	}
}

// Menu at the specified index, or nil.
func (mnu *platformMenu) Menu(index int) menu.Menu {
	item := mnu.Item(index)
	return item.SubMenu()
}

// Popup displays the menu within the window. An attempt will be made to position the 'item'
// at 'where' within the window.
func (mnu *platformMenu) Popup(windowID uint64, where geom.Point, width float64, item menu.Item) {
	wnd := window.ByID(windowID)
	if wnd != nil {
		mnu.platformPopup(wnd, where, item.(*platformItem).item)
	}
}
