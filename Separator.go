// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

import (
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/geom"
)

// Separator provides a simple vertical or horizontal separator line.
type Separator struct {
	Block
	horizontal bool
}

// NewSeparator creates a new separator.
func NewSeparator(horizontal bool) *Separator {
	sep := &Separator{}
	sep.horizontal = horizontal
	sep.SetSizer(sep)
	sep.EventHandlers().Add(event.PaintType, sep.paint)
	return sep
}

// Sizes implements Sizer
func (sep *Separator) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	if sep.horizontal {
		if hint.Width == NoHint {
			pref.Width = 1
		} else {
			pref.Width = hint.Width
		}
		min.Width = 1
		max.Width = DefaultMax
		min.Height = 1
		pref.Height = 1
		max.Height = 1
	} else {
		if hint.Height == NoHint {
			pref.Height = 1
		} else {
			pref.Height = hint.Height
		}
		min.Height = 1
		max.Height = DefaultMax
		min.Width = 1
		pref.Width = 1
		max.Width = 1
	}
	if border := sep.Border(); border != nil {
		insets := border.Insets()
		min.AddInsets(insets)
		pref.AddInsets(insets)
		max.AddInsets(insets)
	}
	return min, pref, max
}

func (sep *Separator) paint(evt event.Event) {
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
	gc := evt.(*event.Paint).GC()
	gc.SetFillColor(color.Background.AdjustBrightness(-0.25))
	gc.FillRect(bounds)
}
