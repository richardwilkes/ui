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
	"github.com/richardwilkes/ui/geom"
	"unsafe"
)

// #cgo darwin LDFLAGS: -framework Cocoa
// #include <stdlib.h>
// #include "Menu_darwin.h"
import "C"

func platformMenuBar() platformMenu {
	return platformMenu(C.platformMenuBar())
}

func platformNewMenu(title string) platformMenu {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	return platformMenu(C.platformNewMenu(cTitle))
}

func (menu *Menu) platformItem(index int) platformMenuItem {
	return platformMenuItem(C.platformGetMenuItem(menu.menu, C.int(index)))
}

func (menu *Menu) platformAddItem(title string, key string) platformMenuItem {
	cTitle := C.CString(title)
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cTitle))
	defer C.free(unsafe.Pointer(cKey))
	return platformMenuItem(C.platformAddMenuItem(menu.menu, cTitle, cKey))
}

func (menu *Menu) platformAddSeparator() platformMenuItem {
	return platformMenuItem(C.platformAddSeparator(menu.menu))
}

func (menu *Menu) platformCount() int {
	return int(C.platformMenuItemCount(menu.menu))
}

func (menu *Menu) platformSetAsMenuBar() {
	C.platformSetMenuBar(menu.menu)
}

func (menu *Menu) platformSetAsServicesMenu() {
	C.platformSetServicesMenu(menu.menu)
}

func (menu *Menu) platformSetAsWindowMenu() {
	C.platformSetWindowMenu(menu.menu)
}

func (menu *Menu) platformSetAsHelpMenu() {
	C.platformSetHelpMenu(menu.menu)
}

func (menu *Menu) platformPopup(widget Widget, where geom.Point, itemAtLocation *MenuItem) {
	C.platformPopupMenu(widget.Window().PlatformPtr(), menu.menu, C.float(where.X), C.float(where.Y), itemAtLocation.item)
}

func (menu *Menu) platformDispose() {
	C.platformDisposeMenu(menu.menu)
}
