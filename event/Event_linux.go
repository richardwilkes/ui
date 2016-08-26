// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package event

import (
	"time"
	"unsafe"
	// #cgo linux LDFLAGS: -lX11
	// #include <X11/Xlib.h>
	"C"
)

var (
	xDisplay *C.Display
	taskAtom C.Atom
)

func platformInvoke(id uint64) {
	event := C.XClientMessageEvent{_type: C.ClientMessage, message_type: taskAtom, format: 32}
	data := (*uint64)(unsafe.Pointer(&event.data))
	*data = id
	C.XSendEvent(xDisplay, 0, 0, C.NoEventMask, (*C.XEvent)(unsafe.Pointer(&event)))
	C.XFlush(xDisplay)
}

func platformInvokeAfter(id uint64, after time.Duration) {
	time.AfterFunc(after, func() {
		platformInvoke(id)
	})
}

func PlatformSetXDisplay(display unsafe.Pointer, goTaskAtom uint32) {
	xDisplay = (*C.Display)(display)
	taskAtom = C.Atom(goTaskAtom)
}
