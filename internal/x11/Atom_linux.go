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
	// #cgo pkg-config: x11
	// #include <X11/Xlib.h>
	"C"
)

type Atom C.Atom

var (
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
	multipleAtom                 Atom
)

func initAtoms() {
	wmWindowTypeAtom = NewAtom("_NET_WM_WINDOW_TYPE")
	wmWindowTypeNormalAtom = NewAtom("_NET_WM_WINDOW_TYPE_NORMAL")
	wmWindowTypeDropDownMenuAtom = NewAtom("_NET_WM_WINDOW_TYPE_DROPDOWN_MENU")
	wmPidAtom = NewAtom("_NET_WM_PID")
	WindowStateAtom = NewAtom("_NET_WM_STATE")
	wmWindowStateSkipTaskBarAtom = NewAtom("_NET_WM_STATE_SKIP_TASKBAR")
	wmWindowFrameExtentsAtom = NewAtom("_NET_FRAME_EXTENTS")
	WindowStateMaximizedHAtom = NewAtom("_NET_WM_STATE_MAXIMIZED_HORZ")
	WindowStateMaximizedVAtom = NewAtom("_NET_WM_STATE_MAXIMIZED_VERT")
	ProtocolsSubType = NewAtom("WM_PROTOCOLS")
	TaskSubType = NewAtom("GoTask")
	DeleteWindowSubType = NewAtom("WM_DELETE_WINDOW")
	clipboardAtom = NewAtom("CLIPBOARD")
	utf8Atom = NewAtom("UTF8_STRING")
	compoundStringAtom = NewAtom("COMPOUND_STRING")
	pairAtom = NewAtom("ATOM_PAIR")
	targetsAtom = NewAtom("TARGETS")
	multipleAtom = NewAtom("MULTIPLE")
}

func NewAtom(name string) Atom {
	return Atom(C.XInternAtom(display, C.CString(name), C.False))
}
