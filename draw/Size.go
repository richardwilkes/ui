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
	"fmt"
	"github.com/richardwilkes/xmath"
)

// Size defines a width and height.
type Size struct {
	Width, Height float32
}

// Add modifies this Size by adding the supplied Size.
func (s *Size) Add(size Size) {
	s.Width += size.Width
	s.Height += size.Height
}

// AddInsets modifies this Size by expanding it to accomodate the specified insets.
func (s *Size) AddInsets(insets Insets) {
	s.Width += insets.Left + insets.Right
	s.Height += insets.Top + insets.Bottom
}

// Subtract modifies this Size by subtracting the supplied Size.
func (s *Size) Subtract(size Size) {
	s.Width -= size.Width
	s.Height -= size.Height
}

// SubtractInsets modifies this Size by reducing it to accomodate the specified insets.
func (s *Size) SubtractInsets(insets Insets) {
	s.Width -= insets.Left + insets.Right
	s.Height -= insets.Top + insets.Bottom
}

// GrowToInteger modifies this Size such that its width and height are both the smallest integers
// greater than or equal to their original values.
func (s *Size) GrowToInteger() {
	s.Width = xmath.CeilFloat32(s.Width)
	s.Height = xmath.CeilFloat32(s.Height)
}

// String implements the fmt.Stringer interface.
func (s Size) String() string {
	return fmt.Sprintf("%v, %v", s.Width, s.Height)
}
