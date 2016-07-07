// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package widget

var (
	// StdImageButtonTheme is the theme all new ImageButtons get by default.
	StdImageButtonTheme = NewImageButtonTheme()
)

// ImageButtonTheme contains the theme elements for ImageButtons.
type ImageButtonTheme struct {
	BaseButtonTheme
	HorizontalMargin float32 // The margin on the left and right side of the image.
	VerticalMargin   float32 // The margin on the top and bottom of the image.
}

// NewImageButtonTheme creates a new image button theme.
func NewImageButtonTheme() *ImageButtonTheme {
	theme := &ImageButtonTheme{}
	theme.Init()
	return theme
}

// Init initializes the theme with its default values.
func (theme *ImageButtonTheme) Init() {
	theme.BaseButtonTheme.Init()
	theme.HorizontalMargin = 4
	theme.VerticalMargin = 4
}
