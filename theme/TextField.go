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
	"github.com/richardwilkes/ui/border"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/font"
	"github.com/richardwilkes/ui/geom"
	"time"
)

var (
	// StdTextField is the theme all new TextFields get by default.
	StdTextField = NewTextField()
)

// TextField contains the theme elements for TextFields.
type TextField struct {
	DisabledBackgroundColor color.Color   // The color to use for the background when disabled.
	Font                    *font.Font    // The font to use.
	Border                  border.Border // The border to use when not focused.
	FocusBorder             border.Border // The border to use when focused.
	BlinkRate               time.Duration // The rate at which the cursor blinks.
	HorizontalMargin        float32       // The margin on the left and right side of the text.
	VerticalMargin          float32       // The margin on the top and bottom of the text.
	MinimumTextWidth        float32       // The minimum space to permit for text.
}

// NewTextField creates a new TextField theme.
func NewTextField() *TextField {
	theme := &TextField{}
	theme.Init()
	return theme
}

// Init initializes the theme with its default values.
func (theme *TextField) Init() {
	theme.DisabledBackgroundColor = color.Background
	theme.Font = font.Acquire(font.UserDesc)
	insets := geom.Insets{Top: 1, Left: 1, Bottom: 1, Right: 1}
	theme.Border = border.NewLine(color.Background.AdjustBrightness(-0.25), insets)
	theme.FocusBorder = border.NewLine(color.KeyboardFocus, insets)
	theme.BlinkRate = time.Millisecond * 560
	theme.HorizontalMargin = 4
	theme.VerticalMargin = 1
	theme.MinimumTextWidth = 10
}
