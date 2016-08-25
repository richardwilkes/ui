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
	"unsafe"
)

// #cgo darwin LDFLAGS: -framework Cocoa
// #include <stdlib.h>
// #include "Menu_darwin.h"
import "C"

func platformMenuBar() PlatformMenu {
	return PlatformMenu(C.platformMenuBar())
}

func platformNewMenu(title string) PlatformMenu {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	return PlatformMenu(C.platformNewMenu(cTitle))
}

func (menu *Menu) platformItem(index int) PlatformItem {
	return PlatformItem(C.platformGetMenuItem(menu.menu, C.int(index)))
}

func (menu *Menu) platformAddItem(title string, key string) PlatformItem {
	cTitle := C.CString(title)
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cTitle))
	defer C.free(unsafe.Pointer(cKey))
	return PlatformItem(C.platformAddMenuItem(menu.menu, cTitle, cKey))
}

func (menu *Menu) platformAddSeparator() PlatformItem {
	return PlatformItem(C.platformAddSeparator(menu.menu))
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

func (menu *Menu) platformDispose() {
	C.platformDisposeMenu(menu.menu)
}
