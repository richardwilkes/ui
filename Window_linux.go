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
	C.XGetInputFocus(xDisplay, &focus, &revert)
	return toPlatformWindow(focus)
}

func platformBringAllWindowsToFront() {
	// RAW: Implement for Linux
}

func platformHideCursorUntilMouseMoves() {
	// RAW: Implement for Linux
}

func platformNewWindow(bounds geom.Rect, styleMask WindowStyleMask) platformWindow {
	screen := C.XDefaultScreen(xDisplay)
	window := C.XCreateSimpleWindow(xDisplay, C.XRootWindow(xDisplay, screen), C.int(bounds.X), C.int(bounds.Y), C.uint(bounds.Width), C.uint(bounds.Height), 1, C.XBlackPixel(xDisplay, screen), C.XWhitePixel(xDisplay, screen))
	C.XSelectInput(xDisplay, window, C.KeyPressMask|C.KeyReleaseMask|C.ButtonPressMask|C.ButtonReleaseMask|C.EnterWindowMask|C.LeaveWindowMask|C.ExposureMask|C.PointerMotionMask|C.ExposureMask|C.VisibilityChangeMask|C.StructureNotifyMask|C.FocusChangeMask)
	C.XSetWMProtocols(xDisplay, window, &wmDeleteMessage, 1)
	C.XMapWindow(xDisplay, window)
	// Move it back to the original location, as the window manager might have set it somewhere else
	C.XMoveWindow(xDisplay, window, C.int(bounds.X), C.int(bounds.Y))
	xWindowCount++
	return toPlatformWindow(window)
}

func (window *Window) platformClose() {
	xWindowCount--
	C.XDestroyWindow(xDisplay, toXWindow(window.window))
}

func (window *Window) platformTitle() string {
	var result *C.char
	C.XFetchName(xDisplay, toXWindow(window.window), &result)
	if result == nil {
		return ""
	}
	defer C.XFree(result)
	return C.GoString(result)
}

func (window *Window) platformSetTitle(title string) {
	cTitle := C.CString(title)
	C.XStoreName(xDisplay, toXWindow(window.window), cTitle)
	C.free(unsafe.Pointer(cTitle))
}

func (window *Window) platformFrame() geom.Rect {
	var root C.Window
	var x, y C.int
	var width, height, border, depth C.uint
	C.XGetGeometry(xDisplay, toXDrawable(window.window), &root, &x, &y, &width, &height, &border, &depth)
	return geom.Rect{Point: geom.Point{X: float32(x), Y: float32(y)}, Size: geom.Size{Width: float32(width), Height: float32(height)}}
}

func (window *Window) platformSetFrame(bounds geom.Rect) {
	C.XMoveResizeWindow(xDisplay, toXWindow(window.window), C.int(bounds.X), C.int(bounds.Y), C.uint(bounds.Width), C.uint(bounds.Height))
}

func (window *Window) platformContentFrame() geom.Rect {
	// RAW: Implement for Linux
	return window.platformFrame()
}

func (window *Window) platformToFront() {
	C.XRaiseWindow(xDisplay, toXWindow(window.window))
}

func (window *Window) platformRepaint(bounds geom.Rect) {
	event := C.XExposeEvent{_type: C.Expose, window: toXWindow(window.window), x: C.int(bounds.X), y: C.int(bounds.Y), width: C.int(bounds.Width), height: C.int(bounds.Height)}
	C.XSendEvent(xDisplay, toXWindow(window.window), 0, C.ExposureMask, (*C.XEvent)(unsafe.Pointer(&event)))
}

func (window *Window) platformFlushPainting() {
	C.XFlush(xDisplay)
}

func (window *Window) platformScalingFactor() float32 {
	// RAW: Implement for Linux
	return 1
}

func (window *Window) platformMinimize() {
	C.XIconifyWindow(xDisplay, toXWindow(window.window), C.XDefaultScreen(xDisplay))
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
