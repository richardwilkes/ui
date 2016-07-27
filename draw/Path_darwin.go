// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package draw

// #cgo darwin LDFLAGS: -framework Cocoa
// #include <CoreGraphics/CoreGraphics.h>
import "C"

func (p *Path) toPlatform() C.CGPathRef {
	path := C.CGPathCreateMutable()
	for _, node := range p.data {
		switch t := node.(type) {
		case *moveToPathNode:
			C.CGPathMoveToPoint(path, nil, C.CGFloat(t.x), C.CGFloat(t.y))
		case *lineToPathNode:
			C.CGPathAddLineToPoint(path, nil, C.CGFloat(t.x), C.CGFloat(t.y))
		case *arcPathNode:
			C.CGPathAddArc(path, nil, C.CGFloat(t.cx), C.CGFloat(t.cy), C.CGFloat(t.radius), C.CGFloat(t.startAngleRadians), C.CGFloat(t.endAngleRadians), C._Bool(t.clockwise))
		case *arcToPathNode:
			C.CGPathAddArcToPoint(path, nil, C.CGFloat(t.x1), C.CGFloat(t.y1), C.CGFloat(t.x2), C.CGFloat(t.y2), C.CGFloat(t.radius))
		case *curveToPathNode:
			C.CGPathAddCurveToPoint(path, nil, C.CGFloat(t.cp1x), C.CGFloat(t.cp1y), C.CGFloat(t.cp2x), C.CGFloat(t.cp2y), C.CGFloat(t.x), C.CGFloat(t.y))
		case *quadCurveToPathNode:
			C.CGPathAddQuadCurveToPoint(path, nil, C.CGFloat(t.cpx), C.CGFloat(t.cpy), C.CGFloat(t.x), C.CGFloat(t.y))
		case *rectPathNode:
			C.CGPathAddRect(path, nil, toCGRect(t.bounds))
		case *ellipsePathNode:
			C.CGPathAddEllipseInRect(path, nil, toCGRect(t.bounds))
		case *closePathNode:
			C.CGPathCloseSubpath(path)
		default:
			panic("Unknown path node type")
		}
	}
	return C.CGPathRef(path)
}

func toCGRect(bounds Rect) C.CGRect {
	return C.CGRectMake(C.CGFloat(bounds.X), C.CGFloat(bounds.Y), C.CGFloat(bounds.Width), C.CGFloat(bounds.Height))
}
