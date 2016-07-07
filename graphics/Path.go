// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package graphics

type moveToPathNode struct {
	x, y float32
}

type lineToPathNode struct {
	x, y float32
}

type arcPathNode struct {
	cx, cy, radius, startAngleRadians, endAngleRadians float32
	clockwise                                          bool
}

type arcToPathNode struct {
	x1, y1, x2, y2, radius float32
}

type curveToPathNode struct {
	cp1x, cp1y, cp2x, cp2y, x, y float32
}

type quadCurveToPathNode struct {
	cpx, cpy, x, y float32
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
func (p *Path) MoveTo(x, y float32) {
	p.data = append(p.data, &moveToPathNode{x: x, y: y})
}

// LineTo appends a straight line segment from the current point to the specified point.
func (p *Path) LineTo(x, y float32) {
	p.data = append(p.data, &lineToPathNode{x: x, y: y})
}

// Arc adds an arc of a circle to the current path, possibly preceded by a straight line segment.
func (p *Path) Arc(cx, cy, radius, startAngleRadians, endAngleRadians float32, clockwise bool) {
	p.data = append(p.data, &arcPathNode{cx: cx, cy: cy, radius: radius, startAngleRadians: startAngleRadians, endAngleRadians: endAngleRadians, clockwise: clockwise})
}

// ArcTo adds an arc of a circle to the current path, using a radius and tangent points.
func (p *Path) ArcTo(x1, y1, x2, y2, radius float32) {
	p.data = append(p.data, &arcToPathNode{x1: x1, y1: y1, x2: x2, y2: y2, radius: radius})
}

// CurveTo appends a cubic Bezier curve from the current point, using the provided controls points
// and end point.
func (p *Path) CurveTo(cp1x, cp1y, cp2x, cp2y, x, y float32) {
	p.data = append(p.data, &curveToPathNode{cp1x: cp1x, cp1y: cp1y, cp2x: cp2x, cp2y: cp2y, x: x, y: y})
}

// QuadCurveTo appends a quadratic Bezier curve from the current point, using a control point and
// an end point.
func (p *Path) QuadCurveTo(cpx, cpy, x, y float32) {
	p.data = append(p.data, &quadCurveToPathNode{cpx: cpx, cpy: cpy, x: x, y: y})
}

// ClosePath closes and terminates the current path’s subpath.
func (p *Path) ClosePath() {
	p.data = append(p.data, &closePathNode{})
}
