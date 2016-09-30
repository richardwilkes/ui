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
	"strings"
	"unsafe"
	// #cgo darwin LDFLAGS: -framework Cocoa
	// #include <stdlib.h>
	// #include "Menus_darwin.h"
	"C"
)

func platformBar() cMenu {
	return cMenu(C.platformBar())
}

func platformSetBar(bar cMenu) {
	C.platformSetBar(bar)
}

func platformNewMenu(title string) cMenu {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	return cMenu(C.platformNewMenu(cTitle))
}

func platformNewSeparator() cItem {
	return cItem(C.platformNewSeparator())
}

func platformNewItem(title string, keyCode int, modifiers keys.Modifiers) cItem {
	var keyCodeStr string
	if keyCode != 0 {
		mapping := keys.MappingForKeyCode(keyCode)
		if mapping.KeyChar != 0 {
			keyCodeStr = strings.ToLower(string(mapping.KeyChar))
		}
	}
	cTitle := C.CString(title)
	cKey := C.CString(keyCodeStr)
	defer C.free(unsafe.Pointer(cTitle))
	defer C.free(unsafe.Pointer(cKey))
	return cItem(C.platformNewItem(cTitle, cKey, C.int(modifiers)))
}

func (menu *platformMenu) platformDispose() {
	C.platformDisposeMenu(menu.menu)
}

func (menu *platformMenu) platformItemCount() int {
	return int(C.platformItemCount(menu.menu))
}

func (menu *platformMenu) platformItem(index int) cItem {
	return cItem(C.platformItem(menu.menu, C.int(index)))
}

func (menu *platformMenu) platformAddItem(item cItem) {
	C.platformAddItem(menu.menu, item)
}

func (menu *platformMenu) platformInsertItem(index int, item cItem) {
	C.platformInsertItem(menu.menu, item, C.int(index))
}

func (menu *platformMenu) platformRemove(index int) {
	C.platformRemove(menu.menu, C.int(index))
}

func (item *platformItem) platformDispose() {
	C.platformDisposeItem(item.item)
}

func (item *platformItem) platformSubMenu() cMenu {
	return cMenu(C.platformSubMenu(item.item))
}

func (item *platformItem) platformSetSubMenu(subMenu cMenu) {
	C.platformSetSubMenu(item.item, subMenu)
}

func SetServicesMenu(menu menu.Menu) {
	C.platformSetServicesMenu(menu.(*platformMenu).menu)
}

func SetWindowMenu(menu menu.Menu) {
	C.platformSetWindowMenu(menu.(*platformMenu).menu)
}

func SetHelpMenu(menu menu.Menu) {
	C.platformSetHelpMenu(menu.(*platformMenu).menu)
}

//export platformValidateMenuItem
func platformValidateMenuItem(menuItem cItem) bool {
	if item, ok := itemMap[menuItem]; ok {
		evt := event.NewValidate(item)
		event.Dispatch(evt)
		item.enabled = evt.Valid()
		return item.enabled
	}
	return true
}

//export platformHandleMenuItem
func platformHandleMenuItem(menuItem cItem) {
	if item, ok := itemMap[menuItem]; ok {
		event.Dispatch(event.NewSelection(item))
	}
}
