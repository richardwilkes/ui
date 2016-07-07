// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package border

import (
	"github.com/richardwilkes/go-ui/geom"
	"github.com/richardwilkes/go-ui/graphics"
)

// Compound is a Border that contains other Borders.
type Compound struct {
	borders []Border
}

// NewCompound creates a Border that contains other Borders. The first one will be drawn in the
// outermost position, with each successive one moving further into the interior.
func NewCompound(borders ...Border) Border {
	return &Compound{borders: borders}
}

// Insets -- implements the Border interface.
func (c *Compound) Insets() geom.Insets {
	insets := geom.Insets{}
	for _, one := range c.borders {
		insets.Add(one.Insets())
	}
	return insets
}

// Paint -- implements the Border interface.
func (c *Compound) Paint(gc graphics.Context, bounds geom.Rect) {
	for _, one := range c.borders {
		gc.Save()
		one.Paint(gc, bounds)
		gc.Restore()
		bounds.Inset(one.Insets())
	}
}
