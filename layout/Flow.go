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
	"github.com/richardwilkes/go-ui/geom"
	"math"
)

// Flow lays out the children of its target left-to-right, then top-to-bottom at
// their preferred sizes, if possible.
type Flow struct {
	// The intra-child spacing to use.
	HGap, VGap float32
}

// ComputeSizes implements the Layout interface.
func (flow *Flow) ComputeSizes(target Layoutable, hint geom.Size) (min, pref, max geom.Size) {
	if hint.Width < 0 {
		hint.Width = math.MaxFloat32
	}
	if hint.Height < 0 {
		hint.Height = math.MaxFloat32
	}
	insets := target.Insets()
	width := hint.Width - (insets.Left + insets.Right)
	pt := geom.Point{X: insets.Left, Y: insets.Top}
	result := geom.Size{Width: pt.Y, Height: pt.Y}
	availWidth := width
	availHeight := hint.Height - (insets.Top + insets.Bottom)
	var maxHeight float32
	var largestChildMin geom.Size
	noHint := geom.Size{Width: NoHint, Height: NoHint}
	for _, child := range target.LayoutChildren() {
		min, pref, _ := child.ComputeSizes(noHint)
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
				pt.Y += maxHeight + flow.VGap
				availWidth = width
				availHeight -= maxHeight + flow.VGap
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
			min, pref, _ = child.ComputeSizes(geom.Size{Width: pref.Width, Height: NoHint})
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
		availWidth -= pref.Width + flow.HGap
		if availWidth <= 0 {
			pt.X = insets.Left
			pt.Y += maxHeight + flow.VGap
			availWidth = width
			availHeight -= maxHeight + flow.VGap
			maxHeight = 0
		} else {
			pt.X += pref.Width + flow.HGap
		}
	}
	result.Width += insets.Right
	result.Height += insets.Bottom
	largestChildMin.Width += insets.Left + insets.Right
	largestChildMin.Height += insets.Top + insets.Bottom
	return largestChildMin, result, DefaultMaxSize(result)
}

// Layout implements the Layout interface.
func (flow *Flow) Layout(target Layoutable) {
	insets := target.Insets()
	size := target.Bounds().Size
	width := size.Width - (insets.Left + insets.Right)
	pt := geom.Point{X: insets.Left, Y: insets.Top}
	availWidth := width
	availHeight := size.Height - (insets.Top + insets.Bottom)
	var maxHeight float32
	noHint := geom.Size{Width: NoHint, Height: NoHint}
	for _, child := range target.LayoutChildren() {
		min, pref, _ := child.ComputeSizes(noHint)
		if pref.Width > availWidth {
			if min.Width <= availWidth {
				pref.Width = availWidth
			} else if pt.X == insets.Left {
				pref.Width = min.Width
			} else {
				pt.X = insets.Left
				pt.Y += maxHeight + flow.VGap
				availWidth = width
				availHeight -= maxHeight + flow.VGap
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
			min, pref, _ = child.ComputeSizes(geom.Size{Width: pref.Width, Height: NoHint})
			pref.Width = savedWidth
			if pref.Height > availHeight {
				if min.Height <= availHeight {
					pref.Height = availHeight
				} else {
					pref.Height = min.Height
				}
			}
		}
		child.SetBounds(geom.Rect{Point: pt, Size: pref})
		if maxHeight < pref.Height {
			maxHeight = pref.Height
		}
		availWidth -= pref.Width + flow.HGap
		if availWidth <= 0 {
			pt.X = insets.Left
			pt.Y += maxHeight + flow.VGap
			availWidth = width
			availHeight -= maxHeight + flow.VGap
			maxHeight = 0
		} else {
			pt.X += pref.Width + flow.HGap
		}
	}
}
