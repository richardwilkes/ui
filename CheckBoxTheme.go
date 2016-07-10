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
	// StdCheckBoxTheme is the theme all new CheckBoxes get by default.
	StdCheckBoxTheme = NewCheckBoxTheme()
)

// CheckBoxTheme contains the theme elements for CheckBoxes.
type CheckBoxTheme struct {
	BaseTextButtonTheme
	HorizontalGap float32 // The gap between the checkbox graphic and its label.
}

// NewCheckBoxTheme creates a new checkbox theme.
func NewCheckBoxTheme() *CheckBoxTheme {
	theme := &CheckBoxTheme{}
	theme.Init()
	return theme
}

// Init initializes the theme with its default values.
func (theme *CheckBoxTheme) Init() {
	theme.BaseTextButtonTheme.Init()
	theme.CornerRadius = 4
	theme.HorizontalGap = 4
}
