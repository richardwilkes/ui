// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package flow

import (
	"math"

	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/layout"
)

// Flow lays out the children of its widget left-to-right, then top-to-bottom at their preferred
// sizes, if possible.
type Flow struct {
	widget   ui.Widget
	hSpacing float64
	vSpacing float64
	vCenter  bool
}

// New creates a new Flow layout and sets it on the widget.
func New(widget ui.Widget) *Flow {
	layout := &Flow{widget: widget}
	widget.SetLayout(layout)
	return layout
}

// HorizontalSpacing returns the horizontal spacing between widgets.
func (flow *Flow) HorizontalSpacing() float64 {
	return flow.hSpacing
}

// SetHorizontalSpacing sets the horizontal spacing between widgets.
func (flow *Flow) SetHorizontalSpacing(spacing float64) *Flow {
	flow.hSpacing = math.Max(spacing, 0)
	return flow
}

// VerticalSpacing returns the vertical spacing between rows.
func (flow *Flow) VerticalSpacing() float64 {
	return flow.vSpacing
}

// SetVerticalSpacing sets the vertical spacing between rows.
func (flow *Flow) SetVerticalSpacing(spacing float64) *Flow {
	flow.vSpacing = math.Max(spacing, 0)
	return flow
}

// VerticallyCentered returns true if the widgets should be vertically centered in their row.
func (flow *Flow) VerticallyCentered() bool {
	return flow.vCenter
}

// SetVerticallyCentered sets whether the widgets should be vertically centered in their row.
func (flow *Flow) SetVerticallyCentered(center bool) *Flow {
	flow.vCenter = center
	return flow
}

// Sizes implements the Layout interface.
func (flow *Flow) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	var insets geom.Insets
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
	pt := geom.Point{X: insets.Left, Y: insets.Top}
	result := geom.Size{Width: pt.Y, Height: pt.Y}
	availWidth := width
	availHeight := hint.Height - (insets.Top + insets.Bottom)
	var maxHeight float64
	var largestChildMin geom.Size
	noHint := geom.Size{Width: layout.NoHint, Height: layout.NoHint}
	for _, child := range flow.widget.Children() {
		min, pref, _ := ui.Sizes(child, noHint)
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
			min, pref, _ = ui.Sizes(child, geom.Size{Width: pref.Width, Height: layout.NoHint})
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
	return largestChildMin, result, layout.DefaultMaxSize(result)
}

// Layout implements the Layout interface.
func (flow *Flow) Layout() {
	var insets geom.Insets
	if border := flow.widget.Border(); border != nil {
		insets = border.Insets()
	}
	size := flow.widget.Bounds().Size
	width := size.Width - (insets.Left + insets.Right)
	pt := geom.Point{X: insets.Left, Y: insets.Top}
	availWidth := width
	availHeight := size.Height - (insets.Top + insets.Bottom)
	var maxHeight float64
	noHint := geom.Size{Width: layout.NoHint, Height: layout.NoHint}
	children := flow.widget.Children()
	rects := make([]geom.Rect, len(children))
	start := 0
	for i, child := range children {
		min, pref, _ := ui.Sizes(child, noHint)
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
				if i > start {
					flow.applyRects(children[start:i], rects[start:i], maxHeight)
					start = i
				}
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
			min, pref, _ = ui.Sizes(child, geom.Size{Width: pref.Width, Height: layout.NoHint})
			pref.Width = savedWidth
			if pref.Height > availHeight {
				if min.Height <= availHeight {
					pref.Height = availHeight
				} else {
					pref.Height = min.Height
				}
			}
		}
		rects[i] = geom.Rect{Point: pt, Size: pref}
		if maxHeight < pref.Height {
			maxHeight = pref.Height
		}
		availWidth -= pref.Width + flow.hSpacing
		if availWidth <= 0 {
			pt.X = insets.Left
			pt.Y += maxHeight + flow.vSpacing
			availWidth = width
			availHeight -= maxHeight + flow.vSpacing
			flow.applyRects(children[start:i+1], rects[start:i+1], maxHeight)
			start = i + 1
			maxHeight = 0
		} else {
			pt.X += pref.Width + flow.hSpacing
		}
	}
	for i, child := range children {
		if flow.vCenter {
			// RAW: Implement
		}
		child.SetBounds(rects[i])
	}
}

func (flow *Flow) applyRects(children []ui.Widget, rects []geom.Rect, maxHeight float64) {
	for i, child := range children {
		if flow.vCenter {
			if rects[i].Height < maxHeight {
				rects[i].Y += (maxHeight - rects[i].Height) / 2
			}
		}
		child.SetBounds(rects[i])
	}
}
