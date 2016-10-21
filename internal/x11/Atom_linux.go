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
	"unsafe"
	// #cgo pkg-config: x11
	// #include <stdlib.h>
	// #include <X11/Xlib.h>
	"C"
)

type Atom C.Atom

var (
	atomMap                      = make(map[string]Atom)
	wmWindowTypeAtom             Atom
	wmWindowTypeNormalAtom       Atom
	wmWindowTypeDropDownMenuAtom Atom
	wmPidAtom                    Atom
	WindowStateAtom              Atom
	wmWindowStateSkipTaskBarAtom Atom
	wmWindowFrameExtentsAtom     Atom
	WindowStateMaximizedHAtom    Atom
	WindowStateMaximizedVAtom    Atom
	ProtocolsSubType             Atom
	TaskSubType                  Atom
	DeleteWindowSubType          Atom
	clipboardAtom                Atom
	utf8Atom                     Atom
	compoundStringAtom           Atom
	pairAtom                     Atom
	targetsAtom                  Atom
	saveTargetsAtom              Atom
	multipleAtom                 Atom
)

func initAtoms() {
	wmWindowTypeAtom = InternAtom("_NET_WM_WINDOW_TYPE")
	wmWindowTypeNormalAtom = InternAtom("_NET_WM_WINDOW_TYPE_NORMAL")
	wmWindowTypeDropDownMenuAtom = InternAtom("_NET_WM_WINDOW_TYPE_DROPDOWN_MENU")
	wmPidAtom = InternAtom("_NET_WM_PID")
	WindowStateAtom = InternAtom("_NET_WM_STATE")
	wmWindowStateSkipTaskBarAtom = InternAtom("_NET_WM_STATE_SKIP_TASKBAR")
	wmWindowFrameExtentsAtom = InternAtom("_NET_FRAME_EXTENTS")
	WindowStateMaximizedHAtom = InternAtom("_NET_WM_STATE_MAXIMIZED_HORZ")
	WindowStateMaximizedVAtom = InternAtom("_NET_WM_STATE_MAXIMIZED_VERT")
	ProtocolsSubType = InternAtom("WM_PROTOCOLS")
	TaskSubType = InternAtom("GoTask")
	DeleteWindowSubType = InternAtom("WM_DELETE_WINDOW")
	clipboardAtom = InternAtom("CLIPBOARD")
	utf8Atom = InternAtom("UTF8_STRING")
	compoundStringAtom = InternAtom("COMPOUND_STRING")
	pairAtom = InternAtom("ATOM_PAIR")
	targetsAtom = InternAtom("TARGETS")
	saveTargetsAtom = InternAtom("SAVE_TARGETS")
	multipleAtom = InternAtom("MULTIPLE")
}

func InternAtom(name string) Atom {
	if atom, ok := atomMap[name]; ok {
		return atom
	}
	atom := Atom(C.XInternAtom(display, C.CString(name), C.False))
	atomMap[name] = atom
	return atom
}

func (atom Atom) Name() string {
	str := C.XGetAtomName(display, C.Atom(atom))
	result := C.GoString(str)
	C.free(unsafe.Pointer(str))
	return result
}
