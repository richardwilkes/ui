// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

// LineBorder is a Border that draws a line along some or all of its sides.
type LineBorder struct {
	insets Insets
	color  Color
}

// NewLineBorder creates a new Line Border. The insets represent how thick the border will be drawn
// on that edge.
func NewLineBorder(color Color, insets Insets) Border {
	return &LineBorder{insets: insets, color: color}
}

// Insets implements the Border interface.
func (line *LineBorder) Insets() Insets {
	return line.insets
}

// Paint implements the Border interface.
func (line *LineBorder) Paint(g Graphics, bounds Rect) {
	clip := bounds
	clip.Inset(line.insets)
	g.Save()
	g.BeginPath()
	g.MoveTo(bounds.X, bounds.Y)
	g.LineTo(bounds.X+bounds.Width, bounds.Y)
	g.LineTo(bounds.X+bounds.Width, bounds.Y+bounds.Height)
	g.LineTo(bounds.X, bounds.Y+bounds.Height)
	g.LineTo(bounds.X, bounds.Y)
	g.MoveTo(clip.X, clip.Y)
	g.LineTo(clip.X+clip.Width, clip.Y)
	g.LineTo(clip.X+clip.Width, clip.Y+clip.Height)
	g.LineTo(clip.X, clip.Y+clip.Height)
	g.LineTo(clip.X, clip.Y)
	g.ClipEvenOdd()
	g.SetFillColor(line.color)
	g.FillRect(bounds)
	g.Restore()
}
