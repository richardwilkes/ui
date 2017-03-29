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
	"unsafe"

	"github.com/richardwilkes/geom"

	// #cgo pkg-config: cairo
	// #include <cairo/cairo-xlib.h>
	"C"
)

func NewSurface(surface unsafe.Pointer, size geom.Size) *Surface {
	return &Surface{surface: (*C.cairo_surface_t)(surface), size: size}
}

func (surface *Surface) SetSize(size geom.Size) {
	if surface.size != size {
		surface.size = size
		C.cairo_xlib_surface_set_size(surface.surface, C.int(size.Width), C.int(size.Height))
	}
}
