// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package button

var (
	// StdButton is the theme all new Buttons get by default.
	StdButton = NewTheme()
)

// Button contains the theme elements for Buttons.
type Theme struct {
	BaseTextTheme
	HorizontalMargin float64 // The margin on the left and right side of the text.
	VerticalMargin   float64 // The margin on the top and bottom of the text.
	MinimumTextWidth float64 // The minimum space to permit for text.
}

// NewTheme creates a new button theme.
func NewTheme() *Theme {
	theme := &Theme{}
	theme.Init()
	return theme
}

// Init initializes the theme with its default values.
func (theme *Theme) Init() {
	theme.BaseTextTheme.Init()
	theme.HorizontalMargin = 8
	theme.VerticalMargin = 1
	theme.MinimumTextWidth = 10
}
