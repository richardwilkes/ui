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

// #cgo linux LDFLAGS: -lX11 -lcairo
// #include <stdlib.h>
// #include <stdio.h>
// #include <string.h>
// #include <X11/Xlib.h>
// #include "globals_linux.h"
// #include "Types.h"
import "C"

func toXWindow(window platformWindow) C.Window {
	return C.Window(uintptr(window))
}

func toXDrawable(window platformWindow) C.Drawable {
	return C.Drawable(uintptr(window))
}

func toPlatformWindow(window C.Window) platformWindow {
	return platformWindow(uintptr(window))
}

func platformGetKeyWindow() platformWindow {
	var focus C.Window
	var revert C.int
	C.XGetInputFocus(C.AppGlobals.display, &focus, &revert)
	return toPlatformWindow(focus)
}

func platformBringAllWindowsToFront() {
	// RAW: Implement for Linux
}

func platformHideCursorUntilMouseMoves() {
	// RAW: Implement for Linux
}

func platformNewWindow(bounds geom.Rect, styleMask WindowStyleMask) platformWindow {
	screen := C.XDefaultScreen(C.AppGlobals.display)
	window := C.XCreateSimpleWindow(C.AppGlobals.display, C.XRootWindow(C.AppGlobals.display, screen), C.int(bounds.X), C.int(bounds.Y), C.uint(bounds.Width), C.uint(bounds.Height), 1, C.XBlackPixel(C.AppGlobals.display, screen), C.XWhitePixel(C.AppGlobals.display, screen))
	C.XSelectInput(C.AppGlobals.display, window, C.KeyPressMask|C.KeyReleaseMask|C.ButtonPressMask|C.ButtonReleaseMask|C.EnterWindowMask|C.LeaveWindowMask|C.ExposureMask|C.PointerMotionMask|C.ExposureMask|C.VisibilityChangeMask|C.StructureNotifyMask|C.FocusChangeMask)
	C.XSetWMProtocols(C.AppGlobals.display, window, &C.AppGlobals.wmDeleteMessage, 1)
	C.XMapWindow(C.AppGlobals.display, window)
	// Move it back to the original location, as the window manager might have set it somewhere else
	C.XMoveWindow(C.AppGlobals.display, window, C.int(bounds.X), C.int(bounds.Y))
	C.AppGlobals.windowCount++
	return toPlatformWindow(window)
}

func (window *Window) platformClose() {
	C.AppGlobals.windowCount--
	C.XDestroyWindow(C.AppGlobals.display, toXWindow(window.window))
}

func (window *Window) platformTitle() string {
	var result *C.char
	C.XFetchName(C.AppGlobals.display, toXWindow(window.window), &result)
	if result == nil {
		return ""
	}
	defer C.XFree(result)
	return C.GoString(result)
}

func (window *Window) platformSetTitle(title string) {
	cTitle := C.CString(title)
	C.XStoreName(C.AppGlobals.display, toXWindow(window.window), cTitle)
	C.free(unsafe.Pointer(cTitle))
}

func (window *Window) platformFrame() geom.Rect {
	var root C.Window
	var x, y C.int
	var width, height, border, depth C.uint
	C.XGetGeometry(C.AppGlobals.display, toXDrawable(window.window), &root, &x, &y, &width, &height, &border, &depth)
	return geom.Rect{Point: geom.Point{X: float32(x), Y: float32(y)}, Size: geom.Size{Width: float32(width), Height: float32(height)}}
}

func (window *Window) platformSetFrame(bounds geom.Rect) {
	C.XMoveResizeWindow(C.AppGlobals.display, toXWindow(window.window), C.int(bounds.X), C.int(bounds.Y), C.uint(bounds.Width), C.uint(bounds.Height))
}

func (window *Window) platformContentFrame() geom.Rect {
	// RAW: Implement for Linux
	return window.platformFrame()
}

func (window *Window) platformToFront() {
	C.XRaiseWindow(C.AppGlobals.display, toXWindow(window.window))
}

func (window *Window) platformRepaint(bounds geom.Rect) {
	event := C.XExposeEvent{_type: C.Expose, window: toXWindow(window.window), x: C.int(bounds.X), y: C.int(bounds.Y), width: C.int(bounds.Width), height: C.int(bounds.Height)}
	C.XSendEvent(C.AppGlobals.display, toXWindow(window.window), 0, C.ExposureMask, (*C.XEvent)(unsafe.Pointer(&event)))
}

func (window *Window) platformFlushPainting() {
	C.XFlush(C.AppGlobals.display)
}

func (window *Window) platformScalingFactor() float32 {
	// RAW: Implement for Linux
	return 1
}

func (window *Window) platformMinimize() {
	C.XIconifyWindow(C.AppGlobals.display, toXWindow(window.window), C.XDefaultScreen(C.AppGlobals.display))
}

func (window *Window) platformZoom() {
	// RAW: Implement for Linux
}

func (window *Window) platformSetToolTip(tip string) {
	// RAW: Implement for Linux
}

func (window *Window) platformSetCursor(c *cursor.Cursor) {
	// RAW: Implement for Linux
}
