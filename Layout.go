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
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/xmath"
)

const (
	// NoLayoutHint is passed as a hint value when one or both dimensions have no suggested value.
	NoLayoutHint = -1
	// DefaultLayoutMax is the default value that should be used for a maximum dimension if the
	// block has no real preference and can be expanded beyond its preferred size. This is
	// intentionally not something like math.MaxFloat32 to allow basic math operations an
	// opportunity to succeed when laying out components. It is perfectly acceptable to use
	// a larger value than this, however, if that makes sense for your specific component.
	DefaultLayoutMax = 10000
)

// Sizer is called when no layout has been set for a widget. Returns the minimum, preferred, and
// maximum sizes of the widget. The hint's values will be either NoLayoutHint or a specific value
// if that particular dimension has already been determined.
type Sizer interface {
	Sizes(hint draw.Size) (min, pref, max draw.Size)
}

// The Layout interface should be implemented by objects that provide layout services.
type Layout interface {
	Sizer
	// Layout is called to layout the target and its children.
	Layout()
}

var (
	// NoLayoutHintSize is a convenience for passing to layouts when you don't have any particular
	// size constraints in mind. Should be treated as read-only.
	NoLayoutHintSize = draw.Size{Width: NoLayoutHint, Height: NoLayoutHint}
)

// DefaultLayoutMaxSize returns the size that is at least as large as DefaultLayoutMax in both dimensions, but
// larger if the preferred size that is passed in is larger.
func DefaultLayoutMaxSize(pref draw.Size) draw.Size {
	return draw.Size{Width: xmath.MaxFloat32(DefaultLayoutMax, pref.Width), Height: xmath.MaxFloat32(DefaultLayoutMax, pref.Height)}
}

// ComputeSizes returns the minimum, preferred, and maximum sizes 'widget' wishes to be. It does
// this by asking the widget's Layout. If no Layout is present, then the widget's Sizer is asked.
// If no Sizer is present, then it finally uses a default set of sizes that are used for all
// components.
func ComputeSizes(widget Widget, hint draw.Size) (min, pref, max draw.Size) {
	if l := widget.Layout(); l != nil {
		return l.Sizes(hint)
	}
	if s := widget.Sizer(); s != nil {
		return s.Sizes(hint)
	}
	return draw.Size{}, draw.Size{}, DefaultLayoutMaxSize(draw.Size{})
}
