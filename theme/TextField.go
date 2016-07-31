// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package theme

import (
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/font"
)

var (
	// StdTextField is the theme all new TextFields get by default.
	StdTextField = NewTextField()
)

// TextField contains the theme elements for TextFields.
type TextField struct {
	BorderColor             color.Color // The color to use for the border.
	DisabledBackgroundColor color.Color // The color to use for the background when disabled.
	Font                    *font.Font  // The font to use.
	HorizontalMargin        float32     // The margin on the left and right side of the text.
	VerticalMargin          float32     // The margin on the top and bottom of the text.
	MinimumTextWidth        float32     // The minimum space to permit for text.
}

// NewTextField creates a new TextField theme.
func NewTextField() *TextField {
	theme := &TextField{}
	theme.Init()
	return theme
}

// Init initializes the theme with its default values.
func (theme *TextField) Init() {
	theme.BorderColor = color.Background.AdjustBrightness(-0.25)
	theme.DisabledBackgroundColor = color.Background
	theme.Font = font.Acquire(font.UserDesc)
	theme.HorizontalMargin = 4
	theme.VerticalMargin = 1
	theme.MinimumTextWidth = 10
}
