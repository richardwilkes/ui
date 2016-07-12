// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

// CompoundBorder is a Border that contains other Borders.
type CompoundBorder struct {
	borders []Border
}

// NewCompoundBorder creates a Border that contains other Borders. The first one will be drawn in
// the outermost position, with each successive one moving further into the interior.
func NewCompoundBorder(borders ...Border) Border {
	return &CompoundBorder{borders: borders}
}

// Insets implements the Border interface.
func (c *CompoundBorder) Insets() Insets {
	insets := Insets{}
	for _, one := range c.borders {
		insets.Add(one.Insets())
	}
	return insets
}

// PaintBorder implements the Border interface.
func (c *CompoundBorder) PaintBorder(g Graphics, bounds Rect) {
	for _, one := range c.borders {
		g.Save()
		one.PaintBorder(g, bounds)
		g.Restore()
		bounds.Inset(one.Insets())
	}
}
