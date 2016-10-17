// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package draw

import (
	"github.com/richardwilkes/geom"
	"unsafe"
	// #cgo pkg-config: cairo
	// #include <cairo.h>
	// #include <cairo/cairo-xlib.h>
	"C"
)

func NewSurface(wnd Window, size geom.Size) *Surface {
	return (*Surface)(C.cairo_xlib_surface_create(display, C.Drawable(uintptr(wnd)), C.XDefaultVisual(display, C.XDefaultScreen(display)), C.int(size.Width), C.int(size.Height)))
}

func (surface *Surface) SetSize(size geom.Size) {
	C.cairo_xlib_surface_set_size(surface, C.int(size.Width), C.int(size.Height))
}
