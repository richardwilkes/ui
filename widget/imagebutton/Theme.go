// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package imagebutton

import (
	"github.com/richardwilkes/ui/widget/button"
)

var (
	// StdImageButton is the theme all new ImageButtons get by default.
	StdImageButton = NewTheme()
)

// ImageButton contains the theme elements for ImageButtons.
type Theme struct {
	button.BaseTheme
	HorizontalMargin float64 // The margin on the left and right side of the image.
	VerticalMargin   float64 // The margin on the top and bottom of the image.
}

// NewTheme creates a new image button theme.
func NewTheme() *Theme {
	theme := &Theme{}
	theme.Init()
	return theme
}

// Init initializes the theme with its default values.
func (theme *Theme) Init() {
	theme.BaseTheme.Init()
	theme.HorizontalMargin = 4
	theme.VerticalMargin = 4
}
