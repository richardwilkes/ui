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
	"github.com/richardwilkes/ui/draw"
)

// Compound is a border that contains other borders.
type Compound struct {
	borders []Border
}

// NewCompound creates a border that contains other borders. The first one will be drawn in
// the outermost position, with each successive one moving further into the interior.
func NewCompound(borders ...Border) Border {
	return &Compound{borders: borders}
}

// Insets implements the Border interface.
func (c *Compound) Insets() draw.Insets {
	insets := draw.Insets{}
	for _, one := range c.borders {
		insets.Add(one.Insets())
	}
	return insets
}

// Draw implements the Border interface.
func (c *Compound) Draw(gc draw.Graphics, bounds draw.Rect) {
	for _, one := range c.borders {
		gc.Save()
		one.Draw(gc, bounds)
		gc.Restore()
		bounds.Inset(one.Insets())
	}
}
