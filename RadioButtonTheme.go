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
	// StdRadioButtonTheme is the theme all new RadioButtons get by default.
	StdRadioButtonTheme = NewRadioButtonTheme()
)

// RadioButtonTheme contains the theme elements for RadioButtons.
type RadioButtonTheme struct {
	BaseTextButtonTheme
	HorizontalGap float32 // The gap between the radio button graphic and its label.
}

// NewRadioButtonTheme creates a new radio button theme.
func NewRadioButtonTheme() *RadioButtonTheme {
	theme := &RadioButtonTheme{}
	theme.Init()
	return theme
}

// Init initializes the theme with its default values.
func (theme *RadioButtonTheme) Init() {
	theme.BaseTextButtonTheme.Init()
	theme.HorizontalGap = 4
}
