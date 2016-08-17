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
	"github.com/richardwilkes/ui/cursor"
	"github.com/richardwilkes/ui/geom"
	"unsafe"
)

// #cgo darwin LDFLAGS: -framework Cocoa -framework Quartz
// #include <stdlib.h>
// #include "Window_darwin.h"
import "C"

func platformGetKeyWindow() platformWindow {
	return platformWindow(C.platformGetKeyWindow())
}

func platformBringAllWindowsToFront() {
	C.platformBringAllWindowsToFront()
}

func platformHideCursorUntilMouseMoves() {
	C.platformHideCursorUntilMouseMoves()
}

func platformNewWindow(bounds geom.Rect, styleMask WindowStyleMask) platformWindow {
	return platformWindow(C.platformNewWindow(toCRect(bounds), C.int(styleMask)))
}

func (window *Window) platformClose() {
	C.platformCloseWindow(window.window)
}

func (window *Window) platformTitle() string {
	return C.GoString(C.platformGetWindowTitle(window.window))
}

func (window *Window) platformSetTitle(title string) {
	cTitle := C.CString(title)
	C.platformSetWindowTitle(window.window, cTitle)
	C.free(unsafe.Pointer(cTitle))
}

func (window *Window) platformFrame() geom.Rect {
	return C.platformGetWindowFrame(window.window).toRect()
}

func (window *Window) platformSetFrame(bounds geom.Rect) {
	C.platformSetWindowFrame(window.window, toCRect(bounds))
}

func (window *Window) platformContentFrame() geom.Rect {
	return C.platformGetWindowContentFrame(window.window).toRect()
}

func (window *Window) platformToFront() {
	C.platformBringWindowToFront(window.window)
}

func (window *Window) platformRepaint(bounds geom.Rect) {
	C.platformRepaintWindow(window.window, toCRect(bounds))
}

func (window *Window) platformFlushPainting() {
	C.platformFlushPainting(window.window)
}

func (window *Window) platformScalingFactor() float32 {
	return float32(C.platformGetWindowScalingFactor(window.window))
}

func (window *Window) platformMinimize() {
	C.platformMinimizeWindow(window.window)
}

func (window *Window) platformZoom() {
	C.platformZoomWindow(window.window)
}

func (window *Window) platformSetToolTip(tip string) {
	if tip != "" {
		cstr := C.CString(tip)
		C.platformSetToolTip(window.window, cstr)
		C.free(unsafe.Pointer(cstr))
	} else {
		C.platformSetToolTip(window.window, nil)
	}
}

func (window *Window) platformSetCursor(c *cursor.Cursor) {
	C.platformSetCursor(window.window, c.PlatformPtr())
}
