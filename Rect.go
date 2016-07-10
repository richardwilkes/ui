// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

import (
	"fmt"
)

// Rect defines a rectangle.
type Rect struct {
	Point
	Size
}

// CopyAndZeroLocation creates a new copy of the Rect and sets the location of the copy to 0,0.
func (r *Rect) CopyAndZeroLocation() Rect {
	return Rect{Size: r.Size}
}

// IsEmpty returns true if either the width or height is zero or less.
func (r *Rect) IsEmpty() bool {
	return r.Width <= 0 || r.Height <= 0
}

// Intersect this Rect with another Rect, storing the result into this Rect.
func (r *Rect) Intersect(other Rect) {
	if r.IsEmpty() || other.IsEmpty() {
		r.Width = 0
		r.Height = 0
	} else {
		x := MaxFloat32(r.X, other.X)
		y := MaxFloat32(r.Y, other.Y)
		w := MinFloat32(r.X+r.Width, other.X+other.Width) - x
		h := MinFloat32(r.Y+r.Height, other.Y+other.Height) - y
		if w > 0 && h > 0 {
			r.X = x
			r.Y = y
			r.Width = w
			r.Height = h
		} else {
			r.Width = 0
			r.Height = 0
		}
	}
}

// Union this Rect with another Rect, storing the result into this Rect.
func (r *Rect) Union(other Rect) {
	e1 := r.IsEmpty()
	e2 := other.IsEmpty()
	if e1 && e2 {
		r.Width = 0
		r.Height = 0
	} else if e1 {
		*r = other
	} else if !e2 {
		x := MinFloat32(r.X, other.X)
		y := MinFloat32(r.Y, other.Y)
		r.Width = MaxFloat32(r.X+r.Width, other.X+other.Width) - x
		r.Height = MaxFloat32(r.Y+r.Height, other.Y+other.Height) - y
		r.X = x
		r.Y = y
	}
}

// InsetUniform insets this Rect by the specified amount on all sides. Positive values make the
// Rect smaller, while negative values make it larger.
func (r *Rect) InsetUniform(amount float32) {
	r.X += amount
	r.Y += amount
	r.Width -= amount * 2
	if r.Width < 0 {
		r.Width = 0
		r.Height = 0
	} else {
		r.Height -= amount * 2
		if r.Height < 0 {
			r.Width = 0
			r.Height = 0
		}
	}
}

// Inset this Rect by the specified Insets.
func (r *Rect) Inset(insets Insets) {
	r.X += insets.Left
	r.Y += insets.Top
	r.Width -= insets.Left + insets.Right
	if r.Width <= 0 {
		r.Width = 0
	}
	r.Height -= insets.Top + insets.Bottom
	if r.Height < 0 {
		r.Height = 0
	}
}

// Contains returns true if the coordinates are within the Rect.
func (r *Rect) Contains(pt Point) bool {
	if r.IsEmpty() {
		return false
	}
	return r.X <= pt.X && r.Y <= pt.Y && pt.X < r.X+r.Width && pt.Y < r.Y+r.Height
}

// String implements the fmt.Stringer interface.
func (r *Rect) String() string {
	return fmt.Sprintf("%v, %v", r.Point, r.Size)
}
