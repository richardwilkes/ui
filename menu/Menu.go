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
	"github.com/richardwilkes/ui/event"
)

var (
	parentTarget event.Target
	menuMap      = make(map[PlatformMenu]*Menu)
	itemMap      = make(map[PlatformItem]*Item)
)

// Menu represents a set of menu items.
type Menu struct {
	menu  PlatformMenu
	title string
}

// ParentTarget returns the value that will be returned on calls to a menu item's ParentTarget()
// method.
func ParentTarget() event.Target {
	return parentTarget
}

// SetParentTarget sets the value that should be returned on calls to a menu item's ParentTarget()
// method.
func SetParentTarget(target event.Target) {
	parentTarget = target
}

// Bar returns the application menu bar.
func Bar() *Menu {
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

// PlatformPtr returns the underlying platform data pointer.
func (menu *Menu) PlatformPtr() PlatformMenu {
	return menu.menu
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
func (menu *Menu) Item(index int) *Item {
	if item, ok := itemMap[menu.platformItem(index)]; ok {
		return item
	}
	return nil
}

// AddItem creates a new Item and appends it to the end of the Menu.
func (menu *Menu) AddItem(title string, key string) *Item {
	item := &Item{item: menu.platformAddItem(title, key), title: title}
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
	item := &Item{item: menu.platformAddSeparator()}
	itemMap[item.item] = item
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
