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
	"github.com/richardwilkes/go-ui/event"
	"unsafe"
)

// #cgo darwin LDFLAGS: -framework Cocoa
// #include <stdlib.h>
// #include "menu.h"
import "C"

var (
	menuMap = make(map[C.uiMenu]*Menu)
	itemMap = make(map[C.uiMenuItem]*Item)
)

// Action that an Item can take.
type Action func(item *Item)

// Validator determines whether the specified Item should be enabled or not.
type Validator func(item *Item) bool

// Menu represents a set of menu items.
type Menu struct {
	menu  C.uiMenu
	title string
}

// Item represents individual actions that can be issued from a Menu.
type Item struct {
	item      C.uiMenuItem
	title     string
	action    Action
	validator Validator
}

// Bar returns the application menu bar.
func Bar() *Menu {
	if menu, ok := menuMap[C.getMainMenu()]; ok {
		return menu
	}
	menu := New("")
	C.setMainMenu(menu.menu)
	return menu
}

// New creates a new Menu.
func New(title string) *Menu {
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

// Count of Items in this Menu.
func (menu *Menu) Count() int {
	return int(C.uiMenuItemCount(menu.menu))
}

// Item at the specified index, or nil.
func (menu *Menu) Item(index int) *Item {
	if item, ok := itemMap[C.uiGetMenuItem(menu.menu, C.int(index))]; ok {
		return item
	}
	return nil
}

// AddItem creates a new Item and appends it to the end of the Menu.
func (menu *Menu) AddItem(title string, key string, action Action, validator Validator) *Item {
	cTitle := C.CString(title)
	cKey := C.CString(key)
	item := &Item{item: C.uiAddMenuItem(menu.menu, cTitle, cKey), title: title, action: action, validator: validator}
	C.free(unsafe.Pointer(cTitle))
	C.free(unsafe.Pointer(cKey))
	itemMap[item.item] = item
	return item
}

// AddMenu creates a new sub-Menu and appends it to the end of the Menu.
func (menu *Menu) AddMenu(title string) *Menu {
	item := menu.AddItem(title, "", nil, nil)
	subMenu := New(title)
	C.uiSetSubMenu(item.item, subMenu.menu)
	return subMenu
}

// AddSeparator creates a new separator and appends it to the end of the Menu.
func (menu *Menu) AddSeparator() {
	item := &Item{item: C.uiAddSeparator(menu.menu)}
	itemMap[item.item] = item
}

// Title returns this item's title.
func (item *Item) Title() string {
	return item.title
}

// SetKeyModifiers sets the Item's key equivalent modifiers. By default, a Item's modifier is set
// to event.CommandKeyMask.
func (item *Item) SetKeyModifiers(modifierMask event.KeyMask) {
	C.uiSetKeyModifierMask(item.item, C.int(modifierMask))
}

// SubMenu of this Item or nil.
func (item *Item) SubMenu() *Menu {
	if menu, ok := menuMap[C.uiGetSubMenu(item.item)]; ok {
		return menu
	}
	return nil
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
			delete(itemMap, item)
		}
		delete(menuMap, menu.menu)
		C.uiDisposeMenu(menu.menu)
		menu.menu = nil
	}
}

//export validateMenuItem
func validateMenuItem(cMenuItem C.uiMenuItem) bool {
	if item, ok := itemMap[cMenuItem]; ok {
		if item.validator != nil {
			return item.validator(item)
		}
	}
	return true
}

//export handleMenuItem
func handleMenuItem(cMenuItem C.uiMenuItem) {
	if item, ok := itemMap[cMenuItem]; ok {
		if item.action != nil {
			item.action(item)
		}
	}
}
