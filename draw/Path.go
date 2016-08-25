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
)

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

// Append the other path to the end of this path.
func (p *Path) Append(other *Path) {
	p.data = append(p.data, other.data...)
}

// BeginSubPath creates a new empty sub-path. Any existing path is not affected.
func (p *Path) BeginSubPath() {
	p.data = append(p.data, &subPathNode{})
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

type subPathNode struct {
}

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
