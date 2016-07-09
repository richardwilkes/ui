// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package widget

// #cgo darwin LDFLAGS: -framework Cocoa
// #include "Menu.h"
import "C"

// MenuAction that an MenuItem can take.
type MenuAction func(item *MenuItem)

// MenuValidator determines whether the specified MenuItem should be enabled or not.
type MenuValidator func(item *MenuItem) bool

// MenuItem represents individual actions that can be issued from a Menu.
type MenuItem struct {
	item      C.uiMenuItem
	title     string
	action    MenuAction
	validator MenuValidator
}

// Title returns this item's title.
func (item *MenuItem) Title() string {
	return item.title
}

// SetKeyModifiers sets the MenuItem's key equivalent modifiers. By default, a MenuItem's modifier is set
// to event.CommandKeyMask.
func (item *MenuItem) SetKeyModifiers(modifierMask KeyMask) {
	C.uiSetKeyModifierMask(item.item, C.int(modifierMask))
}

// SubMenu of this MenuItem or nil.
func (item *MenuItem) SubMenu() *Menu {
	if menu, ok := menuMap[C.uiGetSubMenu(item.item)]; ok {
		return menu
	}
	return nil
}

//export validateMenuItem
func validateMenuItem(cMenuItem C.uiMenuItem) bool {
	if item, ok := menuItemMap[cMenuItem]; ok {
		if item.validator != nil {
			return item.validator(item)
		}
	}
	return true
}

//export handleMenuItem
func handleMenuItem(cMenuItem C.uiMenuItem) {
	if item, ok := menuItemMap[cMenuItem]; ok {
		if item.action != nil {
			item.action(item)
		}
	}
}
