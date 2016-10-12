// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package x11

import (
	"github.com/richardwilkes/geom"
	"os"
	"unsafe"
	// #cgo linux LDFLAGS: -lX11 -lcairo
	// #include <stdlib.h>
	// #include <X11/Xlib.h>
	// #include <X11/Xatom.h>
	// #include <X11/Xutil.h>
	"C"
)

type Window C.Window

func NewWindow(bounds geom.Rect) Window {
	attr, mask := prepareCommonWindowAttributes()
	win := C.XCreateWindow(display, C.XRootWindow(display, C.XDefaultScreen(display)), C.int(bounds.X), C.int(bounds.Y), C.uint(bounds.Width), C.uint(bounds.Height), 0, C.CopyFromParent, C.InputOutput, nil, mask, attr)
	win.applyCommonSetup()
	winType := wmWindowTypeNormalAtom
	C.XChangeProperty(display, win, wmWindowTypeAtom, C.XA_ATOM, 32, C.PropModeReplace, (*C.uchar)(unsafe.Pointer(&winType)), 1)
	win.setWindowHints(bounds)
	return Window(win)
}

func NewMenuWindow(parent Window, bounds geom.Rect) Window {
	attr, mask := prepareCommonWindowAttributes()
	attr.override_redirect = C.True
	win := C.XCreateWindow(display, C.XRootWindow(display, C.XDefaultScreen(display)), C.int(bounds.X), C.int(bounds.Y), C.uint(bounds.Width), C.uint(bounds.Height), 0, C.CopyFromParent, C.InputOutput, nil, mask|C.CWOverrideRedirect, attr)
	win.applyCommonSetup()
	winType := wmWindowTypeDropDownMenuAtom
	C.XChangeProperty(display, win, wmWindowTypeAtom, C.XA_ATOM, 32, C.PropModeReplace, (*C.uchar)(unsafe.Pointer(&winType)), 1)
	winState := wmWindowStateSkipTaskBarAtom
	C.XChangeProperty(display, win, wmWindowStateAtom, C.XA_ATOM, 32, C.PropModeReplace, (*C.uchar)(unsafe.Pointer(&winState)), 1)
	C.XSetTransientForHint(display, win, C.Window(parent))
	win.setWindowHints(bounds)
	return Window(win)
}

func prepareCommonWindowAttributes() (attr *C.XSetWindowAttributes, mask C.ulong) {
	return &C.XSetWindowAttributes{background_pixmap: C.None, backing_store: C.WhenMapped}, C.CWBackPixmap | C.CWBackingStore
}

func (wnd C.Window) applyCommonSetup() {
	C.XSelectInput(display, wnd, C.KeyPressMask|C.KeyReleaseMask|C.ButtonPressMask|C.ButtonReleaseMask|C.EnterWindowMask|C.LeaveWindowMask|C.ExposureMask|C.PointerMotionMask|C.ExposureMask|C.VisibilityChangeMask|C.StructureNotifyMask|C.FocusChangeMask)
	C.XSetWMProtocols(display, wnd, (*C.Atom)(&DeleteWindowProtocol), C.True)
	pid := os.Getpid()
	C.XChangeProperty(display, wnd, wmPidAtom, C.XA_CARDINAL, 32, C.PropModeReplace, (*C.uchar)(unsafe.Pointer(&pid)), 1)
}

func (wnd C.Window) setWindowHints(bounds geom.Rect) {
	var sizeHints C.XSizeHints
	sizeHints.x = C.int(bounds.X)
	sizeHints.y = C.int(bounds.Y)
	sizeHints.width = C.int(bounds.Width)
	sizeHints.height = C.int(bounds.Height)
	sizeHints.flags = C.PPosition | C.PSize
	C.XSetWMNormalHints(display, wnd, &sizeHints)
}

func (wnd Window) Destroy() {
	C.XDestroyWindow(display, C.Window(wnd))
}

func (wnd Window) Title() string {
	var result *C.char
	C.XFetchName(display, C.Window(wnd), &result)
	if result == nil {
		return ""
	}
	defer C.XFree(result)
	return C.GoString(result)
}

func (wnd Window) SetTitle(title string) {
	cTitle := C.CString(title)
	C.XStoreName(display, C.Window(wnd), cTitle)
	C.free(unsafe.Pointer(cTitle))
}

func (wnd Window) FrameDecorationSpace() (top, left, bottom, right float64) {
	var actualType C.Atom
	var actualFormat C.int
	var count C.ulong
	var bytes C.ulong
	var data *C.uchar
	if C.XGetWindowProperty(display, C.Window(wnd), wmWindowFrameExtentsAtom, 0, 4, C.False, C.XA_CARDINAL, &actualType, &actualFormat, &count, &bytes, &data) == C.Success {
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
	return
}

func (wnd Window) SetFrame(bounds geom.Rect) {
	C.XMoveResizeWindow(display, C.Window(wnd), C.int(bounds.X), C.int(bounds.Y), C.uint(bounds.Width), C.uint(bounds.Height))
}

func (wnd Window) ContentFrame() geom.Rect {
	var xwa C.XWindowAttributes
	if C.XGetWindowAttributes(display, C.Window(wnd), &xwa) != 0 {
		var x, y C.int
		var child C.Window
		if C.XTranslateCoordinates(display, C.Window(wnd), xwa.root, 0, 0, &x, &y, &child) != 0 {
			return geom.Rect{Point: geom.Point{X: float64(x - xwa.x), Y: float64(y - xwa.y)}, Size: geom.Size{Width: float64(xwa.width), Height: float64(xwa.height)}}
		}
	}
	return geom.Rect{}
}

func (wnd Window) Raise() {
	C.XRaiseWindow(display, C.Window(wnd))
}

func (wnd Window) Show() {
	C.XMapWindow(display, C.Window(wnd))
}

func (wnd Window) Move(where geom.Point) {
	C.XMoveWindow(display, C.Window(wnd), C.int(where.X), C.int(where.Y))
}

func (wnd Window) RequestFocus() {
	C.XSetInputFocus(display, C.Window(wnd), C.RevertToNone, C.CurrentTime)
}

func (wnd Window) Repaint(bounds geom.Rect) {
	event := C.XExposeEvent{_type: C.Expose, window: C.Window(wnd), x: C.int(bounds.X), y: C.int(bounds.Y), width: C.int(bounds.Width), height: C.int(bounds.Height)}
	C.XSendEvent(display, C.Window(wnd), 0, C.ExposureMask, (*C.XEvent)(unsafe.Pointer(&event)))
}

func (wnd Window) Minimize() {
	C.XIconifyWindow(display, C.Window(wnd), C.XDefaultScreen(display))
}

func (wnd Window) InvokeTask(id uint64) {
	event := C.XClientMessageEvent{_type: C.ClientMessage, message_type: C.Atom(TaskSubType), format: 32}
	data := (*uint64)(unsafe.Pointer(&event.data))
	*data = id
	C.XSendEvent(display, C.Window(wnd), 0, C.NoEventMask, (*C.XEvent)(unsafe.Pointer(&event)))
	Flush()
}

func (wnd Window) SetCursor(cursor Cursor) {
	C.XDefineCursor(display, C.Window(wnd), C.Cursor(cursor))
}
