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
	"C"
)

type Surface C.cairo_surface_t

func (surface *Surface) Destroy() {
	C.cairo_surface_destroy(surface)
}

func (surface *Surface) NewCairoContext(bounds geom.Rect) CairoContext {
	return CairoContext(unsafe.Pointer(C.cairo_create(surface)))
}
