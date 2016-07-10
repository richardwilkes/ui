// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

var (
	// StdButtonTheme is the theme all new Buttons get by default.
	StdButtonTheme = NewButtonTheme()
)

// ButtonTheme contains the theme elements for Buttons.
type ButtonTheme struct {
	BaseTextButtonTheme
	HorizontalMargin float32 // The margin on the left and right side of the text.
	VerticalMargin   float32 // The margin on the top and bottom of the text.
	MinimumTextWidth float32 // The minimum space to permit for text.
}

// NewButtonTheme creates a new button theme.
func NewButtonTheme() *ButtonTheme {
	theme := &ButtonTheme{}
	theme.Init()
	return theme
}

// Init initializes the theme with its default values.
func (theme *ButtonTheme) Init() {
	theme.BaseTextButtonTheme.Init()
	theme.HorizontalMargin = 8
	theme.VerticalMargin = 1
	theme.MinimumTextWidth = 10
}
