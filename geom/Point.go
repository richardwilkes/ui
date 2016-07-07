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
	"fmt"
)

// Point defines a location.
type Point struct {
	X, Y float32
}

// Add modifies this Point by adding the supplied coordinates.
func (p *Point) Add(pt Point) {
	p.X += pt.X
	p.Y += pt.Y
}

// Subtract modifies this Point by subtracting the supplied coordinates.
func (p *Point) Subtract(pt Point) {
	p.X -= pt.X
	p.Y -= pt.Y
}

// String -- implements the fmt.Stringer interface.
func (p *Point) String() string {
	return fmt.Sprintf("%v, %v", p.X, p.Y)
}
