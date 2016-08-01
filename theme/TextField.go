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
	theme.Border = border.NewCompound(border.NewLine(color.Background.AdjustBrightness(-0.25), geom.Insets{Top: 1, Left: 1, Bottom: 1, Right: 1}), border.NewEmpty(geom.Insets{Top: 1, Left: 4, Bottom: 1, Right: 4}))
	theme.FocusBorder = border.NewCompound(border.NewLine(color.KeyboardFocus, geom.Insets{Top: 2, Left: 2, Bottom: 2, Right: 2}), border.NewEmpty(geom.Insets{Top: 0, Left: 3, Bottom: 0, Right: 3}))
	theme.BlinkRate = time.Millisecond * 560
	theme.MinimumTextWidth = 10
}
