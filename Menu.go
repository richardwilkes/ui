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
	"github.com/richardwilkes/geom"
)

var (
	menuMap = make(map[platformMenu]*Menu)
	itemMap = make(map[platformMenuItem]*MenuItem)
)

// Menu represents a set of menu items.
type Menu struct {
	menu  platformMenu
	title string
}

// MenuBar returns the application menu bar.
func MenuBar() *Menu {
	if menu, ok := menuMap[platformMenuBar()]; ok {
		return menu
	}
	menu := NewMenu("")
	menu.platformSetAsMenuBar()
	return menu
}

// NewMenu creates a new Menu.
func NewMenu(title string) *Menu {
	menu := &Menu{menu: platformNewMenu(title), title: title}
	menuMap[menu.menu] = menu
	return menu
}

// Title returns the title of this Menu.
func (menu *Menu) Title() string {
	return menu.title
}

// Count of Items in this Menu.
func (menu *Menu) Count() int {
	return menu.platformCount()
}

// Item at the specified index, or nil.
func (menu *Menu) Item(index int) *MenuItem {
	if item, ok := itemMap[menu.platformItem(index)]; ok {
		return item
	}
	return nil
}

// AddItem creates a new Item and appends it to the end of the Menu.
func (menu *Menu) AddItem(title string, key string) *MenuItem {
	item := &MenuItem{item: menu.platformAddItem(title, key), title: title}
	itemMap[item.item] = item
	return item
}

// AddMenu creates a new sub-Menu and appends it to the end of the Menu.
func (menu *Menu) AddMenu(title string) *Menu {
	item := menu.AddItem(title, "")
	subMenu := NewMenu(title)
	item.platformSetSubMenu(subMenu)
	return subMenu
}

// AddSeparator creates a new separator and appends it to the end of the Menu.
func (menu *Menu) AddSeparator() {
	item := &MenuItem{item: menu.platformAddSeparator()}
	itemMap[item.item] = item
}

// Popup shows the menu at the specified location. If itemAtLocation is specified, it also tries to
// position the menu such that the specified menu item is at that location.
func (menu *Menu) Popup(widget Widget, where geom.Point, itemAtLocation *MenuItem) {
	menu.platformPopup(widget, widget.ToWindow(where), itemAtLocation)
}

// Dispose of the Menu, releasing any operating system resources it consumed.
func (menu *Menu) Dispose() {
	if menu.menu != nil {
		count := menu.platformCount()
		for i := 0; i < count; i++ {
			if item := menu.Item(i); item != nil {
				if subMenu := menuMap[item.platformSubMenu()]; subMenu != nil {
					subMenu.Dispose()
				}
				delete(itemMap, item.item)
			}
		}
		delete(menuMap, menu.menu)
		menu.platformDispose()
		menu.menu = nil
	}
}

// SetAsServicesMenu marks the specified menu as the services menu.
func (menu *Menu) SetAsServicesMenu() {
	menu.platformSetAsServicesMenu()
}

// SetAsWindowMenu marks the specified menu as the window menu.
func (menu *Menu) SetAsWindowMenu() {
	menu.platformSetAsWindowMenu()
}

// SetAsHelpMenu marks the specified menu as the help menu.
func (menu *Menu) SetAsHelpMenu() {
	menu.platformSetAsHelpMenu()
}
