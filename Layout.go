// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

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

// Alignment constants.
const (
	AlignStart Alignment = iota
	AlignMiddle
	AlignEnd
	AlignFill
)

// Alignment specifies how to align an object within its available space.
type Alignment uint8

// The Layout interface should be implemented by objects that provide layout services.
type Layout interface {
	// ComputeSizes returns the minimum, preferred, and maximum sizes of the target. The hint's
	// values will be either NoLayoutHint or a specific value if that particular dimension has already
	// been determined.
	ComputeSizes(target *Block, hint Size) (min, pref, max Size)
	// Layout is called to layout the target and its children.
	Layout(target *Block)
}

var (
	// NoLayoutHintSize is a convenience for passing to layouts when you don't have any particular
	// size constraints in mind. Should be treated as read-only.
	NoLayoutHintSize = Size{Width: NoLayoutHint, Height: NoLayoutHint}
)

// DefaultLayoutMaxSize returns the size that is at least as large as DefaultLayoutMax in both dimensions, but
// larger if the preferred size that is passed in is larger.
func DefaultLayoutMaxSize(pref Size) Size {
	return Size{Width: MaxFloat32(DefaultLayoutMax, pref.Width), Height: MaxFloat32(DefaultLayoutMax, pref.Height)}
}
