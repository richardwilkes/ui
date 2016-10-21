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
	// #include <X11/Xlib.h>
	"C"
)

type SelectionRequestEvent C.XSelectionRequestEvent

func (evt *SelectionRequestEvent) Owner() Window {
	return Window(evt.owner)
}

func (evt *SelectionRequestEvent) Selection() Atom {
	return Atom(evt.selection)
}

func (evt *SelectionRequestEvent) Target() Atom {
	return Atom(evt.target)
}

func (evt *SelectionRequestEvent) Property() Atom {
	return Atom(evt.property)
}

func (evt *SelectionRequestEvent) Requestor() Window {
	return Window(evt.requestor)
}

func (evt *SelectionRequestEvent) When() C.Time {
	return evt.time
}

func (evt *SelectionRequestEvent) ToEvent() *Event {
	return (*Event)(unsafe.Pointer(evt))
}

func (evt *SelectionRequestEvent) NewNotify(bad bool) *SelectionEvent {
	prop := evt.property
	if bad {
		prop = C.None
	}
	return (*SelectionEvent)(&C.XSelectionEvent{_type: C.SelectionNotify, requestor: evt.requestor, selection: evt.selection, target: evt.target, property: prop, time: evt.time})
}
