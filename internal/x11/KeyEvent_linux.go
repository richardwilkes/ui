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
	"github.com/richardwilkes/ui/keys"
	// #cgo pkg-config: x11
	// #include <X11/Xutil.h>
	"C"
)

const (
	KeyPressType   = EventType(C.KeyPress)
	KeyReleaseType = EventType(C.KeyRelease)
)

type KeyEvent C.XKeyEvent

func (evt *KeyEvent) Window() Window {
	return Window(evt.window)
}

func (evt *KeyEvent) Modifiers() keys.Modifiers {
	return Modifiers(evt.state)
}

func (evt *KeyEvent) CodeAndChar() (code int, ch rune) {
	var buffer [5]C.char
	var keySym C.KeySym
	buffer[C.XLookupString(evt, &buffer[0], C.int(len(buffer)-1), &keySym, nil)] = 0
	code, ch = keys.Transform(int(keySym), C.GoString(&buffer[0]))
	return
}
