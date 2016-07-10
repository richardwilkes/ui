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
	"unsafe"
)

// #cgo darwin LDFLAGS: -framework Cocoa
// #include <stdlib.h>
// #include "Menu.h"
import "C"

var (
	menuMap     = make(map[C.uiMenu]*Menu)
	menuItemMap = make(map[C.uiMenuItem]*MenuItem)
)

// Menu represents a set of menu items.
type Menu struct {
	menu  C.uiMenu
	title string
}

// MenuBar returns the application menu bar.
func MenuBar() *Menu {
	if menu, ok := menuMap[C.getMainMenu()]; ok {
		return menu
	}
	menu := NewMenu("")
	C.setMainMenu(menu.menu)
	return menu
}

// NewMenu creates a new Menu.
func NewMenu(title string) *Menu {
	cTitle := C.CString(title)
	menu := &Menu{menu: C.uiNewMenu(cTitle), title: title}
	C.free(unsafe.Pointer(cTitle))
	menuMap[menu.menu] = menu
	return menu
}

// SetServicesMenu marks the specified menu as the services menu.
func SetServicesMenu(menu *Menu) {
	C.uiSetServicesMenu(menu.menu)
}

// SetWindowMenu marks the specified menu as the window menu.
func SetWindowMenu(menu *Menu) {
	C.uiSetWindowMenu(menu.menu)
}

// SetHelpMenu marks the specified menu as the help menu.
func SetHelpMenu(menu *Menu) {
	C.uiSetHelpMenu(menu.menu)
}

// Title returns the title of this Menu.
func (menu *Menu) Title() string {
	return menu.title
}

// Count of MenuItems in this Menu.
func (menu *Menu) Count() int {
	return int(C.uiMenuItemCount(menu.menu))
}

// Item at the specified index, or nil.
func (menu *Menu) Item(index int) *MenuItem {
	if item, ok := menuItemMap[C.uiGetMenuItem(menu.menu, C.int(index))]; ok {
		return item
	}
	return nil
}

// AddItem creates a new MenuItem and appends it to the end of the Menu.
func (menu *Menu) AddItem(title string, key string, action MenuAction, validator MenuValidator) *MenuItem {
	cTitle := C.CString(title)
	cKey := C.CString(key)
	item := &MenuItem{item: C.uiAddMenuItem(menu.menu, cTitle, cKey), title: title, action: action, validator: validator}
	C.free(unsafe.Pointer(cTitle))
	C.free(unsafe.Pointer(cKey))
	menuItemMap[item.item] = item
	return item
}

// AddMenu creates a new sub-Menu and appends it to the end of the Menu.
func (menu *Menu) AddMenu(title string) *Menu {
	item := menu.AddItem(title, "", nil, nil)
	subMenu := NewMenu(title)
	C.uiSetSubMenu(item.item, subMenu.menu)
	return subMenu
}

// AddSeparator creates a new separator and appends it to the end of the Menu.
func (menu *Menu) AddSeparator() {
	item := &MenuItem{item: C.uiAddSeparator(menu.menu)}
	menuItemMap[item.item] = item
}

// Popup shows the menu at the specified location. If itemAtLocation is specified, it also tries to
// position the menu such that the specified menu item is at that location.
func (menu *Menu) Popup(block *Block, where Point, itemAtLocation *MenuItem) {
	where = block.ToWindow(where)
	C.uiPopupMenu(block.Window().window, menu.menu, C.float(where.X), C.float(where.Y), itemAtLocation.item)
}

// Dispose of the Menu, releasing any operating system resources it consumed.
func (menu *Menu) Dispose() {
	if menu.menu != nil {
		count := C.uiMenuItemCount(menu.menu)
		var i C.int
		for i = 0; i < count; i++ {
			item := C.uiGetMenuItem(menu.menu, i)
			subMenu := menuMap[C.uiGetSubMenu(item)]
			if subMenu != nil {
				subMenu.Dispose()
			}
			delete(menuItemMap, item)
		}
		delete(menuMap, menu.menu)
		C.uiDisposeMenu(menu.menu)
		menu.menu = nil
	}
}
