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
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/geom"
)

// Line is a border that draws a line along some or all of its sides.
type Line struct {
	insets geom.Insets
	color  color.Color
}

// NewLine creates a new line border. The insets represent how thick the border will be drawn
// on that edge.
func NewLine(color color.Color, insets geom.Insets) Border {
	return &Line{insets: insets, color: color}
}

// Insets implements the Border interface.
func (line *Line) Insets() geom.Insets {
	return line.insets
}

// Draw implements the Border interface.
func (line *Line) Draw(gc draw.Graphics, bounds geom.Rect) {
	clip := bounds
	clip.Inset(line.insets)
	gc.Save()
	gc.BeginPath()
	gc.MoveTo(bounds.X, bounds.Y)
	gc.LineTo(bounds.X+bounds.Width, bounds.Y)
	gc.LineTo(bounds.X+bounds.Width, bounds.Y+bounds.Height)
	gc.LineTo(bounds.X, bounds.Y+bounds.Height)
	gc.LineTo(bounds.X, bounds.Y)
	gc.MoveTo(clip.X, clip.Y)
	gc.LineTo(clip.X+clip.Width, clip.Y)
	gc.LineTo(clip.X+clip.Width, clip.Y+clip.Height)
	gc.LineTo(clip.X, clip.Y+clip.Height)
	gc.LineTo(clip.X, clip.Y)
	gc.ClipEvenOdd()
	gc.SetFillColor(line.color)
	gc.FillRect(bounds)
	gc.Restore()
}
