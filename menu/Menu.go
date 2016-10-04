// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package menu

import (
	"github.com/richardwilkes/geom"
)

// Menu represents a set of menu items.
type Menu interface {
	// InsertItem inserts an item at the specified item index within this menu. Pass in a negative
	// index to append to the end.
	InsertItem(item Item, index int)
	// InsertMenu inserts an item with a sub-menu at the specified item index within this menu. Pass
	// in a negative index to append to the end.
	InsertMenu(menu Menu, index int)
	// Remove the item at the specified index from this menu.
	Remove(index int)
	// Title returns the title of this menu.
	Title() string
	// Count of items in this menu.
	Count() int
	// Item at the specified index, or nil.
	Item(index int) Item
	// Popup displays the menu within the window. An attempt will be made to position the 'item'
	// at 'where' within the window.
	Popup(windowID int64, where geom.Point, item Item)
	// Dispose releases any operating system resources associated with this menu. It will also
	// call Dispose() on all menu items it contains.
	Dispose()
}

var (
	// NewMenu creates a new menu.
	NewMenu func(title string) Menu
)
