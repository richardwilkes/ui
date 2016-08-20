// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package geom

import (
	"unsafe"
)

// #cgo linux LDFLAGS: -lX11 -lcairo
// #include <stdio.h>
// #include <cairo/cairo.h>
import "C"

// PlatformPtr returns a pointer to the underlying platform-specific data.
func (p *Path) PlatformPtr() unsafe.Pointer {
	// RAW: Implement PlatformPtr for Linux
	return nil
}

func (p *Path) Apply(gc unsafe.Pointer) {
	for _, node := range p.data {
		switch t := node.(type) {
		case *moveToPathNode:
			C.cairo_move_to(gc, C.double(t.x), C.double(t.y))
		case *lineToPathNode:
			C.cairo_line_to(gc, C.double(t.x), C.double(t.y))
		case *arcPathNode:
			if t.clockwise {
				C.cairo_arc(gc, C.double(t.cx), C.double(t.cy), C.double(t.radius), C.double(t.startAngleRadians), C.double(t.endAngleRadians))
			} else {
				C.cairo_arc_negative(gc, C.double(t.cx), C.double(t.cy), C.double(t.radius), C.double(t.startAngleRadians), C.double(t.endAngleRadians))
			}
		case *arcToPathNode:
			//			C.CGPathAddArcToPoint(path, nil, C.double(t.x1), C.double(t.y1), C.double(t.x2), C.double(t.y2), C.double(t.radius))
		case *curveToPathNode:
			C.cairo_curve_to(gc, C.double(t.cp1x), C.double(t.cp1y), C.double(t.cp2x), C.double(t.cp2y), C.double(t.x), C.double(t.y))
		case *quadCurveToPathNode:
			var x0, y0 C.double
			C.cairo_get_current_point(gc, &x0, &y0)
			x1 := C.double(t.cpx)
			y1 := C.double(t.cpy)
			xx := C.double(t.x)
			yy := C.double(t.y)
			C.cairo_curve_to(gc, 2.0/3.0*x1+1.0/3.0*x0, 2.0/3.0*y1+1.0/3.0*y0, 2.0/3.0*x1+1.0/3.0*xx, 2.0/3.0*y1+1.0/3.0*yy, xx, yy)
		case *rectPathNode:
			C.cairo_rectangle(gc, C.double(t.bounds.X), C.double(t.bounds.Y), C.double(t.bounds.Width), C.double(t.bounds.Height))
		case *ellipsePathNode:
			//C.CGPathAddEllipseInRect(path, nil, toCGRect(t.bounds))
		case *closePathNode:
			C.cairo_close_path(gc)
		default:
			panic("Unknown path node type")
		}
	}
}
