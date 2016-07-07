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
	"github.com/richardwilkes/go-ui/maths"
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

// Alignment constants.
const (
	Beginning Alignment = iota
	Middle
	End
	Fill
)

// Alignment specifies how to align an object within its available space.
type Alignment int

// The Layoutable interface should be implemented by objects that want to participate in layout.
type Layoutable interface {
	// ComputeSizes returns the minimum, preferred, and maximum sizes of the Layoutable. The hint's
	// values will be either NoHint or a specific value if that particular dimension has already
	// been determined.
	ComputeSizes(hint geom.Size) (min, pref, max geom.Size)
	// LayoutChildren returns the Layoutables contained by the target.
	LayoutChildren() []Layoutable
	// LayoutData returns any layout data that is associated with the Layoutable.
	LayoutData() interface{}
	// SetLayoutData sets the layout data for the Layoutable.
	SetLayoutData(data interface{})
	// Bounds returns the location and size of the Layoutable.
	Bounds() geom.Rect
	// SetBounds sets the location and size of the Layoutable.
	SetBounds(bounds geom.Rect)
	// Insets returns the margins for the Layoutable.
	Insets() geom.Insets
}

// The Layout interface should be implemented by objects that provide layout services.
type Layout interface {
	// ComputeSizes returns the minimum, preferred, and maximum sizes of the target. The hint's
	// values will be either NoHint or a specific value if that particular dimension has already
	// been determined.
	ComputeSizes(target Layoutable, hint geom.Size) (min, pref, max geom.Size)
	// Layout is called to layout the target and its children.
	Layout(target Layoutable)
}

var (
	// NoHintSize is a convenience for passing to layouts when you don't have any particular
	// size constraints in mind. Should be treated as read-only.
	NoHintSize = geom.Size{Width: NoHint, Height: NoHint}
)

// DefaultMaxSize returns the size that is at least as large as DefaultMax in both dimensions, but
// larger if the preferred size that is passed in is larger.
func DefaultMaxSize(pref geom.Size) geom.Size {
	return geom.Size{Width: maths.MaxFloat32(DefaultMax, pref.Width), Height: maths.MaxFloat32(DefaultMax, pref.Height)}
}
