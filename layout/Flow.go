// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package layout

import (
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/xmath"
	"math"
)

// Flow lays out the children of its widget left-to-right, then top-to-bottom at their preferred
// sizes, if possible.
type Flow struct {
	widget   ui.Widget
	hSpacing float32
	vSpacing float32
}

// NewFlow creates a new Flow layout and sets it on the widget.
func NewFlow(widget ui.Widget) *Flow {
	layout := &Flow{widget: widget}
	widget.SetLayout(layout)
	return layout
}

// HorizontalSpacing returns the horizontal spacing between widgets.
func (flow *Flow) HorizontalSpacing() float32 {
	return flow.hSpacing
}

// SetHorizontalSpacing sets the horizontal spacing between widgets.
func (flow *Flow) SetHorizontalSpacing(spacing float32) *Flow {
	flow.hSpacing = xmath.MaxFloat32(spacing, 0)
	return flow
}

// VerticalSpacing returns the vertical spacing between rows.
func (flow *Flow) VerticalSpacing() float32 {
	return flow.vSpacing
}

// SetVerticalSpacing sets the vertical spacing between rows.
func (flow *Flow) SetVerticalSpacing(spacing float32) *Flow {
	flow.vSpacing = xmath.MaxFloat32(spacing, 0)
	return flow
}

// Sizes implements the Layout interface.
func (flow *Flow) Sizes(hint draw.Size) (min, pref, max draw.Size) {
	var insets draw.Insets
	if border := flow.widget.Border(); border != nil {
		insets = border.Insets()
	}
	if hint.Width < 0 {
		hint.Width = math.MaxFloat32
	}
	if hint.Height < 0 {
		hint.Height = math.MaxFloat32
	}
	width := hint.Width - (insets.Left + insets.Right)
	pt := draw.Point{X: insets.Left, Y: insets.Top}
	result := draw.Size{Width: pt.Y, Height: pt.Y}
	availWidth := width
	availHeight := hint.Height - (insets.Top + insets.Bottom)
	var maxHeight float32
	var largestChildMin draw.Size
	noHint := draw.Size{Width: NoHint, Height: NoHint}
	for _, child := range flow.widget.Children() {
		min, pref, _ := Sizes(child, noHint)
		if largestChildMin.Width < min.Width {
			largestChildMin.Width = min.Width
		}
		if largestChildMin.Height < min.Height {
			largestChildMin.Height = min.Height
		}
		if pref.Width > availWidth {
			if min.Width <= availWidth {
				pref.Width = availWidth
			} else if pt.X == insets.Left {
				pref.Width = min.Width
			} else {
				pt.X = insets.Left
				pt.Y += maxHeight + flow.vSpacing
				availWidth = width
				availHeight -= maxHeight + flow.vSpacing
				maxHeight = 0
				if pref.Width > availWidth {
					if min.Width <= availWidth {
						pref.Width = availWidth
					} else {
						pref.Width = min.Width
					}
				}
			}
			savedWidth := pref.Width
			min, pref, _ = Sizes(child, draw.Size{Width: pref.Width, Height: NoHint})
			pref.Width = savedWidth
			if pref.Height > availHeight {
				if min.Height <= availHeight {
					pref.Height = availHeight
				} else {
					pref.Height = min.Height
				}
			}
		}
		extent := pt.X + pref.Width
		if result.Width < extent {
			result.Width = extent
		}
		extent = pt.Y + pref.Height
		if result.Height < extent {
			result.Height = extent
		}
		if maxHeight < pref.Height {
			maxHeight = pref.Height
		}
		availWidth -= pref.Width + flow.hSpacing
		if availWidth <= 0 {
			pt.X = insets.Left
			pt.Y += maxHeight + flow.vSpacing
			availWidth = width
			availHeight -= maxHeight + flow.vSpacing
			maxHeight = 0
		} else {
			pt.X += pref.Width + flow.hSpacing
		}
	}
	result.Width += insets.Right
	result.Height += insets.Bottom
	largestChildMin.Width += insets.Left + insets.Right
	largestChildMin.Height += insets.Top + insets.Bottom
	return largestChildMin, result, DefaultMaxSize(result)
}

// Layout implements the Layout interface.
func (flow *Flow) Layout() {
	var insets draw.Insets
	if border := flow.widget.Border(); border != nil {
		insets = border.Insets()
	}
	size := flow.widget.Bounds().Size
	width := size.Width - (insets.Left + insets.Right)
	pt := draw.Point{X: insets.Left, Y: insets.Top}
	availWidth := width
	availHeight := size.Height - (insets.Top + insets.Bottom)
	var maxHeight float32
	noHint := draw.Size{Width: NoHint, Height: NoHint}
	for _, child := range flow.widget.Children() {
		min, pref, _ := Sizes(child, noHint)
		if pref.Width > availWidth {
			if min.Width <= availWidth {
				pref.Width = availWidth
			} else if pt.X == insets.Left {
				pref.Width = min.Width
			} else {
				pt.X = insets.Left
				pt.Y += maxHeight + flow.vSpacing
				availWidth = width
				availHeight -= maxHeight + flow.vSpacing
				maxHeight = 0
				if pref.Width > availWidth {
					if min.Width <= availWidth {
						pref.Width = availWidth
					} else {
						pref.Width = min.Width
					}
				}
			}
			savedWidth := pref.Width
			min, pref, _ = Sizes(child, draw.Size{Width: pref.Width, Height: NoHint})
			pref.Width = savedWidth
			if pref.Height > availHeight {
				if min.Height <= availHeight {
					pref.Height = availHeight
				} else {
					pref.Height = min.Height
				}
			}
		}
		child.SetBounds(draw.Rect{Point: pt, Size: pref})
		if maxHeight < pref.Height {
			maxHeight = pref.Height
		}
		availWidth -= pref.Width + flow.hSpacing
		if availWidth <= 0 {
			pt.X = insets.Left
			pt.Y += maxHeight + flow.vSpacing
			availWidth = width
			availHeight -= maxHeight + flow.vSpacing
			maxHeight = 0
		} else {
			pt.X += pref.Width + flow.hSpacing
		}
	}
}
