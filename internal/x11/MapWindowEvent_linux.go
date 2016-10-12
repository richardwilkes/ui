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
	// #cgo linux LDFLAGS: -lX11
	// #include <X11/Xlib.h>
	"C"
)

const (
	MapWindowType = EventType(C.MapNotify)
)

type MapWindowEvent C.XMapEvent

func (evt *MapWindowEvent) Window() Window {
	return Window(evt.window)
}
