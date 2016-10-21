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
	"unsafe"
	// #cgo pkg-config: x11
	// #include <X11/Xlib.h>
	"C"
)

type ExposeEvent C.XExposeEvent

func NewExposeEvent(wnd Window, bounds geom.Rect) *ExposeEvent {
	return &ExposeEvent{_type: C.Expose, window: C.Window(wnd), x: C.int(bounds.X), y: C.int(bounds.Y), width: C.int(bounds.Width), height: C.int(bounds.Height)}
}

func (evt *ExposeEvent) Window() Window {
	return Window(evt.window)
}

func (evt *ExposeEvent) Bounds() geom.Rect {
	return geom.Rect{Point: geom.Point{X: float64(evt.x), Y: float64(evt.y)}, Size: geom.Size{Width: float64(evt.width), Height: float64(evt.height)}}
}

func (evt *ExposeEvent) ToEvent() *Event {
	return (*Event)(unsafe.Pointer(evt))
}
