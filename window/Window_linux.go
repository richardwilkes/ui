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
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/cursor"
	"os"
	"time"
	"unsafe"
	// #cgo linux LDFLAGS: -lX11 -lcairo
	// #include <stdlib.h>
	// #include <stdio.h>
	// #include <string.h>
	// #include <X11/Xlib.h>
	// #include <X11/keysym.h>
	// #include <X11/Xutil.h>
	// #include <X11/Xatom.h>
	// #include <cairo/cairo.h>
	// #include <cairo/cairo-xlib.h>
	// #include "Types.h"
	"C"
)

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
	C.XGetInputFocus(display, &focus, &revert)
	return toPlatformWindow(focus)
}

func platformBringAllWindowsToFront() {
	// RAW: Implement for Linux
}

func platformHideCursorUntilMouseMoves() {
	// RAW: Implement for Linux
}

func platformNewWindow(bounds geom.Rect, styleMask WindowStyleMask) (window platformWindow, surface platformSurface) {
	screen := C.XDefaultScreen(display)
	attr, mask := prepareCommonWindowAttributes()
	win := C.XCreateWindow(display, C.XRootWindow(display, screen), C.int(bounds.X), C.int(bounds.Y), C.uint(bounds.Width), C.uint(bounds.Height), 0, C.CopyFromParent, C.InputOutput, nil, mask, attr)
	C.XSelectInput(display, win, C.KeyPressMask|C.KeyReleaseMask|C.ButtonPressMask|C.ButtonReleaseMask|C.EnterWindowMask|C.LeaveWindowMask|C.ExposureMask|C.PointerMotionMask|C.ExposureMask|C.VisibilityChangeMask|C.StructureNotifyMask|C.FocusChangeMask)
	C.XSetWMProtocols(display, win, &wmDeleteAtom, C.True)
	pid := os.Getpid()
	C.XChangeProperty(display, win, wmPidAtom, C.XA_CARDINAL, 32, C.PropModeReplace, (*C.uchar)(unsafe.Pointer(&pid)), 1)
	winType := wmWindowTypeNormalAtom
	C.XChangeProperty(display, win, wmWindowTypeAtom, C.XA_ATOM, 32, C.PropModeReplace, (*C.uchar)(unsafe.Pointer(&winType)), 1)
	setWindowHints(win, bounds)
	return toPlatformWindow(win), platformSurface(C.cairo_xlib_surface_create(display, C.Drawable(uintptr(win)), C.XDefaultVisual(display, screen), C.int(bounds.Width), C.int(bounds.Height)))
}

func platformNewMenuWindow(parent ui.Window, bounds geom.Rect) (window platformWindow, surface platformSurface) {
	screen := C.XDefaultScreen(display)
	attr, mask := prepareCommonWindowAttributes()
	attr.override_redirect = C.True
	win := C.XCreateWindow(display, C.XRootWindow(display, screen), C.int(bounds.X), C.int(bounds.Y), C.uint(bounds.Width), C.uint(bounds.Height), 0, C.CopyFromParent, C.InputOutput, nil, mask|C.CWOverrideRedirect, attr)
	C.XSelectInput(display, win, C.KeyPressMask|C.KeyReleaseMask|C.ButtonPressMask|C.ButtonReleaseMask|C.EnterWindowMask|C.LeaveWindowMask|C.ExposureMask|C.PointerMotionMask|C.ExposureMask|C.VisibilityChangeMask|C.StructureNotifyMask|C.FocusChangeMask)
	C.XSetWMProtocols(display, win, &wmDeleteAtom, C.True)
	pid := os.Getpid()
	C.XChangeProperty(display, win, wmPidAtom, C.XA_CARDINAL, 32, C.PropModeReplace, (*C.uchar)(unsafe.Pointer(&pid)), 1)
	winType := wmWindowTypeDropDownMenuAtom
	C.XChangeProperty(display, win, wmWindowTypeAtom, C.XA_ATOM, 32, C.PropModeReplace, (*C.uchar)(unsafe.Pointer(&winType)), 1)
	winState := wmWindowStateSkipTaskBarAtom
	C.XChangeProperty(display, win, wmWindowStateAtom, C.XA_ATOM, 32, C.PropModeReplace, (*C.uchar)(unsafe.Pointer(&winState)), 1)
	C.XSetTransientForHint(display, win, toXWindow(platformWindow(parent.PlatformPtr())))
	setWindowHints(win, bounds)
	return toPlatformWindow(win), platformSurface(C.cairo_xlib_surface_create(display, C.Drawable(uintptr(win)), C.XDefaultVisual(display, screen), C.int(bounds.Width), C.int(bounds.Height)))
}

func prepareCommonWindowAttributes() (attr *C.XSetWindowAttributes, mask C.ulong) {
	return &C.XSetWindowAttributes{background_pixmap: C.None, backing_store: C.WhenMapped}, C.CWBackPixmap | C.CWBackingStore
}

func setWindowHints(win C.Window, bounds geom.Rect) {
	var sizeHints C.XSizeHints
	sizeHints.x = C.int(bounds.X)
	sizeHints.y = C.int(bounds.Y)
	sizeHints.width = C.int(bounds.Width)
	sizeHints.height = C.int(bounds.Height)
	sizeHints.flags = C.PPosition | C.PSize
	C.XSetWMNormalHints(display, win, &sizeHints)
}

func (window *Wnd) platformClose() {
	C.cairo_surface_destroy(window.surface)
	C.XDestroyWindow(display, toXWindow(window.window))
	windowDidClose(window.window)
}

func (window *Wnd) platformTitle() string {
	var result *C.char
	C.XFetchName(display, toXWindow(window.window), &result)
	if result == nil {
		return ""
	}
	defer C.XFree(result)
	return C.GoString(result)
}

func (window *Wnd) platformSetTitle(title string) {
	cTitle := C.CString(title)
	C.XStoreName(display, toXWindow(window.window), cTitle)
	C.free(unsafe.Pointer(cTitle))
}

func (window *Wnd) frameDecorationSpace() (top, left, bottom, right float64) {
	if window.Valid() {
		var actualType C.Atom
		var actualFormat C.int
		var count C.ulong
		var bytes C.ulong
		var data *C.uchar
		if C.XGetWindowProperty(display, toXWindow(window.window), wmWindowFrameExtentsAtom, 0, 4, C.False, C.XA_CARDINAL, &actualType, &actualFormat, &count, &bytes, &data) == C.Success {
			if actualType == C.XA_CARDINAL && actualFormat == 32 && count == 4 {
				fields := (*[4]C.long)(unsafe.Pointer(data))
				left = float64(fields[0])
				right = float64(fields[1])
				top = float64(fields[2])
				bottom = float64(fields[3])
			}
			if data != nil {
				C.XFree(unsafe.Pointer(data))
			}
		}
	}
	return
}

func (window *Wnd) platformFrame() geom.Rect {
	bounds := window.platformContentFrame()
	top, left, bottom, right := window.frameDecorationSpace()
	bounds.X -= left
	bounds.Y -= top
	bounds.Width += left + right
	bounds.Height += top + bottom
	return bounds
}

func (window *Wnd) platformSetFrame(bounds geom.Rect) {
	win := toXWindow(window.window)
	C.XMoveResizeWindow(display, win, C.int(bounds.X), C.int(bounds.Y), C.uint(bounds.Width), C.uint(bounds.Height))
}

func (window *Wnd) platformContentFrame() geom.Rect {
	if window.Valid() {
		win := toXWindow(window.window)
		var xwa C.XWindowAttributes
		if C.XGetWindowAttributes(display, win, &xwa) != 0 {
			var x, y C.int
			var child C.Window
			if C.XTranslateCoordinates(display, win, xwa.root, 0, 0, &x, &y, &child) != 0 {
				return geom.Rect{Point: geom.Point{X: float64(x - xwa.x), Y: float64(y - xwa.y)}, Size: geom.Size{Width: float64(xwa.width), Height: float64(xwa.height)}}
			}
		}
	}
	return geom.Rect{}
}

func (window *Wnd) platformToFront() {
	win := toXWindow(window.window)
	if window.wasMapped {
		C.XRaiseWindow(display, win)
	} else {
		C.XMapWindow(display, win)
		var other C.XEvent
		for C.XCheckTypedWindowEvent(display, win, C.MapNotify, &other) == 0 {
			// Wait for window to be mapped
		}
		window.wasMapped = true
		C.XMoveWindow(display, win, C.int(window.initialLocationRequest.X), C.int(window.initialLocationRequest.Y))
		if window.owner == nil {
			for C.XCheckTypedWindowEvent(display, win, C.ConfigureNotify, &other) == 0 {
				// Wait for window to be configured so that we have correct placement information
			}
			processConfigureEvent(&other, window.window)
		}

		// This is here so that menu windows behave properly
		C.XSetInputFocus(display, win, C.RevertToNone, C.CurrentTime)
	}
}

func (window *Wnd) platformRepaint(bounds geom.Rect) {
	event := C.XExposeEvent{_type: C.Expose, window: toXWindow(window.window), x: C.int(bounds.X), y: C.int(bounds.Y), width: C.int(bounds.Width), height: C.int(bounds.Height)}
	C.XSendEvent(display, toXWindow(window.window), 0, C.ExposureMask, (*C.XEvent)(unsafe.Pointer(&event)))
}

func (window *Wnd) draw(bounds geom.Rect) {
	gc := C.cairo_create(window.surface)
	C.cairo_set_line_width(gc, 1)
	C.cairo_rectangle(gc, C.double(bounds.X), C.double(bounds.Y), C.double(bounds.Width), C.double(bounds.Height))
	C.cairo_clip(gc)
	drawWindow(window.window, gc, platformRect{x: C.double(bounds.X), y: C.double(bounds.Y), width: C.double(bounds.Width), height: C.double(bounds.Height)}, false)
	C.cairo_destroy(gc)
}

func (window *Wnd) platformFlushPainting() {
	C.XFlush(display)
}

func (window *Wnd) platformScalingFactor() float64 {
	// RAW: Implement for Linux
	return 1
}

func (window *Wnd) platformMinimize() {
	C.XIconifyWindow(display, toXWindow(window.window), C.XDefaultScreen(display))
}

func (window *Wnd) platformZoom() {
	// RAW: Implement for Linux
}

func (window *Wnd) platformSetToolTip(tip string) {
	// RAW: Implement for Linux
}

func (window *Wnd) platformSetCursor(c *cursor.Cursor) {
	// RAW: Implement for Linux
}

func (window *Wnd) platformInvoke(id uint64) {
	if window.Valid() {
		event := C.XClientMessageEvent{_type: C.ClientMessage, message_type: goTaskAtom, format: 32}
		data := (*uint64)(unsafe.Pointer(&event.data))
		*data = id
		C.XSendEvent(display, toXWindow(window.window), 0, C.NoEventMask, (*C.XEvent)(unsafe.Pointer(&event)))
		C.XFlush(display)
	}
}

func (window *Wnd) platformInvokeAfter(id uint64, after time.Duration) {
	time.AfterFunc(after, func() {
		window.platformInvoke(id)
	})
}
