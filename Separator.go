// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

// Separator provides a simple vertical or horizontal separator line.
type Separator struct {
	Block
	horizontal bool
}

// NewSeparator creates a new separator.
func NewSeparator(horizontal bool) *Separator {
	sep := &Separator{}
	sep.Init(horizontal)
	return sep
}

// Init initializes the separator.
func (sep *Separator) Init(horizontal bool) {
	sep.Block.Init()
	sep.horizontal = horizontal
	sep.Sizes = func(hint Size) (min, pref, max Size) {
		if sep.horizontal {
			if hint.Width == NoLayoutHint {
				pref.Width = 1
			} else {
				pref.Width = hint.Width
			}
			min.Width = 1
			max.Width = DefaultLayoutMax
			min.Height = 1
			pref.Height = 1
			max.Height = 1
		} else {
			if hint.Height == NoLayoutHint {
				pref.Height = 1
			} else {
				pref.Height = hint.Height
			}
			min.Height = 1
			max.Height = DefaultLayoutMax
			min.Width = 1
			pref.Width = 1
			max.Width = 1
		}
		insets := sep.Insets()
		min.AddInsets(insets)
		pref.AddInsets(insets)
		max.AddInsets(insets)
		return min, pref, max
	}
	sep.OnPaint = func(g Graphics, dirty Rect, inLiveResize bool) {
		bounds := sep.LocalInsetBounds()
		if sep.horizontal {
			if bounds.Height > 1 {
				bounds.Y += (bounds.Height - 1) / 2
				bounds.Height = 1
			}
		} else {
			if bounds.Width > 1 {
				bounds.X += (bounds.Width - 1) / 2
				bounds.Width = 1
			}
		}
		g.SetFillColor(BackgroundColor.AdjustBrightness(-0.25))
		g.FillRect(bounds)
	}
}
