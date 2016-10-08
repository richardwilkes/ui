// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package window

import (
	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui/cursor"
	"time"
	"unsafe"
	// #cgo darwin LDFLAGS: -framework Cocoa -framework Quartz
	// #cgo pkg-config: pangocairo
	// #include <stdlib.h>
	// #include "Window_darwin.h"
	"C"
)

func platformGetKeyWindow() platformWindow {
	return platformWindow(C.platformGetKeyWindow())
}

func platformBringAllWindowsToFront() {
	C.platformBringAllWindowsToFront()
}

func platformHideCursorUntilMouseMoves() {
	C.platformHideCursorUntilMouseMoves()
}

func platformNewWindow(bounds geom.Rect, styleMask WindowStyleMask) (window platformWindow, surface platformSurface) {
	return platformWindow(C.platformNewWindow(toCRect(bounds), C.int(styleMask))), nil
}

func platformNewMenuWindow(parent ui.Window, bounds geom.Rect) (window platformWindow, surface platformSurface) {
	return platformNewWindow(bounds, BorderlessWindowMask)
}

func (window *Wnd) platformClose() {
	C.platformCloseWindow(window.window)
}

func (window *Wnd) platformTitle() string {
	return C.GoString(C.platformGetWindowTitle(window.window))
}

func (window *Wnd) platformSetTitle(title string) {
	cTitle := C.CString(title)
	C.platformSetWindowTitle(window.window, cTitle)
	C.free(unsafe.Pointer(cTitle))
}

func (window *Wnd) platformFrame() geom.Rect {
	return toRect(C.platformGetWindowFrame(window.window))
}

func (window *Wnd) platformSetFrame(bounds geom.Rect) {
	C.platformSetWindowFrame(window.window, toCRect(bounds))
}

func (window *Wnd) platformContentFrame() geom.Rect {
	return toRect(C.platformGetWindowContentFrame(window.window))
}

func (window *Wnd) platformToFront() {
	C.platformBringWindowToFront(window.window)
}

func (window *Wnd) platformRepaint(bounds geom.Rect) {
	C.platformRepaintWindow(window.window, toCRect(bounds))
}

func (window *Wnd) platformFlushPainting() {
	C.platformFlushPainting(window.window)
}

func (window *Wnd) platformScalingFactor() float64 {
	return float64(C.platformGetWindowScalingFactor(window.window))
}

func (window *Wnd) platformMinimize() {
	C.platformMinimizeWindow(window.window)
}

func (window *Wnd) platformZoom() {
	C.platformZoomWindow(window.window)
}

func (window *Wnd) platformSetToolTip(tip string) {
	if tip != "" {
		cstr := C.CString(tip)
		C.platformSetToolTip(window.window, cstr)
		C.free(unsafe.Pointer(cstr))
	} else {
		C.platformSetToolTip(window.window, nil)
	}
}

func (window *Wnd) platformSetCursor(c *cursor.Cursor) {
	C.platformSetCursor(window.window, c.PlatformPtr())
}

func (window *Wnd) platformInvoke(id uint64) {
	C.platformInvoke(C.ulong(id))
}

func (window *Wnd) platformInvokeAfter(id uint64, after time.Duration) {
	C.platformInvokeAfter(C.ulong(id), C.long(after.Nanoseconds()))
}
