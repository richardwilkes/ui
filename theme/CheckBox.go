// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package theme

var (
	// StdCheckBox is the theme all new CheckBoxes get by default.
	StdCheckBox = NewCheckBox()
)

// CheckBox contains the theme elements for CheckBoxes.
type CheckBox struct {
	BaseTextButton
	HorizontalGap float32 // The gap between the checkbox graphic and its label.
}

// NewCheckBox creates a new checkbox theme.
func NewCheckBox() *CheckBox {
	theme := &CheckBox{}
	theme.Init()
	return theme
}

// Init initializes the theme with its default values.
func (theme *CheckBox) Init() {
	theme.BaseTextButton.Init()
	theme.CornerRadius = 4
	theme.HorizontalGap = 4
}
