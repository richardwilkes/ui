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
	"github.com/richardwilkes/ui/geom"
	"math"
	// #cgo pkg-config: pangocairo
	// #include <pango/pangocairo.h>
	"C"
)

type moveToPathNode struct {
	x, y float64
}

type lineToPathNode struct {
	x, y float64
}

type arcPathNode struct {
	cx, cy, radius, startAngleRadians, endAngleRadians float64
	clockwise                                          bool
}

type curveToPathNode struct {
	cp1x, cp1y, cp2x, cp2y, x, y float64
}

type quadCurveToPathNode struct {
	cpx, cpy, x, y float64
}

type rectPathNode struct {
	bounds geom.Rect
}

type ellipsePathNode struct {
	bounds geom.Rect
}

type closePathNode struct {
}

// Path is a description of a series of shapes or lines.
type Path struct {
	data []interface{}
}

// NewPath creates a new, empty Path.
func NewPath() *Path {
	return &Path{}
}

// Copy copies this path's contents into a new path.
func (p *Path) Copy() *Path {
	return &Path{data: append([]interface{}(nil), p.data...)}
}

// MoveTo begins a new subpath at the specified point.
func (p *Path) MoveTo(x, y float64) {
	p.data = append(p.data, &moveToPathNode{x: x, y: y})
}

// LineTo appends a straight line segment from the current point to the specified point.
func (p *Path) LineTo(x, y float64) {
	p.data = append(p.data, &lineToPathNode{x: x, y: y})
}

// Arc adds an arc of a circle to the current path, possibly preceded by a straight line segment.
func (p *Path) Arc(cx, cy, radius, startAngleRadians, endAngleRadians float64, clockwise bool) {
	p.data = append(p.data, &arcPathNode{cx: cx, cy: cy, radius: radius, startAngleRadians: startAngleRadians, endAngleRadians: endAngleRadians, clockwise: clockwise})
}

// CurveTo appends a cubic Bezier curve from the current point, using the provided controls points
// and end point.
func (p *Path) CurveTo(cp1x, cp1y, cp2x, cp2y, x, y float64) {
	p.data = append(p.data, &curveToPathNode{cp1x: cp1x, cp1y: cp1y, cp2x: cp2x, cp2y: cp2y, x: x, y: y})
}

// QuadCurveTo appends a quadratic Bezier curve from the current point, using a control point and
// an end point.
func (p *Path) QuadCurveTo(cpx, cpy, x, y float64) {
	p.data = append(p.data, &quadCurveToPathNode{cpx: cpx, cpy: cpy, x: x, y: y})
}

// Rect adds a rectangle to the path. The rectangle is a complete subpath, i.e. it starts with a
// MoveTo and ends with a ClosePath operation.
func (p *Path) Rect(bounds geom.Rect) {
	p.data = append(p.data, &rectPathNode{bounds: bounds})
}

// Ellipse adds an ellipse to the path. The ellipse is a complete subpath, i.e. it starts with a
// MoveTo and ends with a ClosePath operation.
func (p *Path) Ellipse(bounds geom.Rect) {
	p.data = append(p.data, &ellipsePathNode{bounds: bounds})
}

// ClosePath closes and terminates the current path’s subpath.
func (p *Path) ClosePath() {
	p.data = append(p.data, &closePathNode{})
}

func (p *Path) Apply(gc CairoContext) {
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
			cx := t.bounds.X + t.bounds.Width/2
			cy := t.bounds.Y + t.bounds.Height/2
			C.cairo_save(gc)
			C.cairo_translate(gc, C.double(cx), C.double(cy))
			if t.bounds.Width > t.bounds.Height {
				C.cairo_scale(gc, 1, C.double(t.bounds.Height/t.bounds.Width))
			} else {
				C.cairo_scale(gc, C.double(t.bounds.Width/t.bounds.Height), 1)
			}
			C.cairo_translate(gc, C.double(-cx), C.double(-cy))
			C.cairo_new_path(gc)
			C.cairo_arc(gc, C.double(cx), C.double(cy), C.double(math.Max(t.bounds.Width, t.bounds.Height)/2), 0, C.double(2*math.Pi))
			C.cairo_close_path(gc)
			C.cairo_restore(gc)
		case *closePathNode:
			C.cairo_close_path(gc)
		default:
			panic("Unknown path node type")
		}
	}
}
