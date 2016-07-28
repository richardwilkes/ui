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
	"github.com/richardwilkes/ui/geom"
	"github.com/richardwilkes/xmath"
)

const (
	// NoHint is passed as a hint value when one or both dimensions have no suggested value.
	NoHint = -1
	// DefaultMax is the default value that should be used for a maximum dimension if the
	// block has no real preference and can be expanded beyond its preferred size. This is
	// intentionally not something like math.MaxFloat32 to allow basic math operations an
	// opportunity to succeed when laying out components. It is perfectly acceptable to use
	// a larger value than this, however, if that makes sense for your specific component.
	DefaultMax = 10000
)

var (
	// NoHintSize is a convenience for passing to layouts when you don't have any particular
	// size constraints in mind. Should be treated as read-only.
	NoHintSize = geom.Size{Width: NoHint, Height: NoHint}
)

// DefaultMaxSize returns the size that is at least as large as DefaultMax in both dimensions, but
// larger if the preferred size that is passed in is larger.
func DefaultMaxSize(pref geom.Size) geom.Size {
	return geom.Size{Width: xmath.MaxFloat32(DefaultMax, pref.Width), Height: xmath.MaxFloat32(DefaultMax, pref.Height)}
}

// Sizes returns the minimum, preferred, and maximum sizes the 'widget' wishes to be. It does
// this by asking the widget's Layout. If no Layout is present, then the widget's Sizer is asked.
// If no Sizer is present, then it finally uses a default set of sizes that are used for all
// components.
func Sizes(widget ui.Widget, hint geom.Size) (min, pref, max geom.Size) {
	if l := widget.Layout(); l != nil {
		return l.Sizes(hint)
	}
	if s := widget.Sizer(); s != nil {
		return s.Sizes(hint)
	}
	return geom.Size{}, geom.Size{}, DefaultMaxSize(geom.Size{})
}
