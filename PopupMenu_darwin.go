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
	"github.com/richardwilkes/ui/menu"
	"reflect"
	"unsafe"
)

// #cgo darwin LDFLAGS: -framework Cocoa
// #include "PopupMenu_darwin.h"
import "C"

func platformPopupMenu(window *Window, where geom.Point, menu menu.Menu, item menu.Item) {
	mPtr := (*[1]unsafe.Pointer)(unsafe.Pointer(reflect.ValueOf(menu).Elem().UnsafeAddr()))
	iPtr := (*[1]unsafe.Pointer)(unsafe.Pointer(reflect.ValueOf(item).Elem().UnsafeAddr()))
	C.platformPopupMenu(window.PlatformPtr(), mPtr[0], C.double(where.X), C.double(where.Y), iPtr[0])
}
