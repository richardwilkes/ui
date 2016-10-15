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
	// #cgo linux LDFLAGS: -lX11 -lcairo
	// #include <X11/Xlib.h>
	"C"
)

var (
	display                      *C.Display
	wmWindowTypeAtom             C.Atom
	wmWindowTypeNormalAtom       C.Atom
	wmWindowTypeDropDownMenuAtom C.Atom
	wmPidAtom                    C.Atom
	wmWindowStateAtom            C.Atom
	wmWindowStateSkipTaskBarAtom C.Atom
	wmWindowFrameExtentsAtom     C.Atom
)

func OpenDisplay() {
	C.XInitThreads()
	if display = C.XOpenDisplay(nil); display == nil {
		panic("Failed to open the X11 display")
	}
	ProtocolsSubType = ClientMessageSubType(C.XInternAtom(display, C.CString("WM_PROTOCOLS"), C.False))
	DeleteWindowProtocol = Protocol(C.XInternAtom(display, C.CString("WM_DELETE_WINDOW"), C.False))
	wmWindowTypeAtom = C.XInternAtom(display, C.CString("_NET_WM_WINDOW_TYPE"), C.False)
	wmWindowTypeNormalAtom = C.XInternAtom(display, C.CString("_NET_WM_WINDOW_TYPE_NORMAL"), C.False)
	wmWindowTypeDropDownMenuAtom = C.XInternAtom(display, C.CString("_NET_WM_WINDOW_TYPE_DROPDOWN_MENU"), C.False)
	wmPidAtom = C.XInternAtom(display, C.CString("_NET_WM_PID"), C.False)
	wmWindowStateAtom = C.XInternAtom(display, C.CString("_NET_WM_STATE"), C.False)
	wmWindowStateSkipTaskBarAtom = C.XInternAtom(display, C.CString("_NET_WM_STATE_SKIP_TASKBAR"), C.False)
	wmWindowFrameExtentsAtom = C.XInternAtom(display, C.CString("_NET_FRAME_EXTENTS"), C.False)
	TaskSubType = ClientMessageSubType(C.XInternAtom(display, C.CString("GoTask"), C.False))
}

func CloseDisplay() {
	C.XCloseDisplay(display)
	display = nil
}

func Running() bool {
	return display != nil
}

func NextEvent() *Event {
	var event Event
	C.XNextEvent(display, (*C.XEvent)(&event))
	return &event
}

func NextEventOfTypeForWindow(eventType EventType, wnd Window) *Event {
	var event Event
	if C.XCheckTypedWindowEvent(display, C.Window(wnd), C.int(eventType), (*C.XEvent)(&event)) != 0 {
		return &event
	}
	return nil
}

func InputFocus() Window {
	var focus Window
	var revert C.int
	C.XGetInputFocus(display, (*C.Window)(&focus), &revert)
	return focus
}

func Flush() {
	C.XFlush(display)
}
