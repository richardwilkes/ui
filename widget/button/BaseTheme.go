// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package button

import (
	"time"

	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/draw"
)

// BaseTheme contains the common theme elements used in all buttons.
type BaseTheme struct {
	ClickAnimationTime    time.Duration // The amount of time to spend animating the click action.
	CornerRadius          float64       // The amount of rounding to use on the corners.
	Background            color.Color   // The background color when enabled but not pressed or focused.
	BackgroundWhenPressed color.Color   // The background color when enabled and pressed.
	GradientAdjustment    float64       // The amount to vary the color when creating the background gradient.
	DisabledAdjustment    float64       // The amount to adjust the background brightness when disabled.
	OutlineAdjustment     float64       // The amount to adjust the background brightness when using it to draw the button outline.
}

// Init initializes the theme with its default values.
func (theme *BaseTheme) Init() {
	theme.ClickAnimationTime = time.Millisecond * 100
	theme.CornerRadius = 6
	theme.Background = color.White
	theme.BackgroundWhenPressed = color.KeyboardFocus
	theme.GradientAdjustment = 0.15
	theme.DisabledAdjustment = -0.05
	theme.OutlineAdjustment = -0.5
}

// Gradient returns a gradient for the specified color.
func (theme *BaseTheme) Gradient(base color.Color) *draw.Gradient {
	return draw.NewEvenlySpacedGradient(base.AdjustBrightness(theme.GradientAdjustment), base.AdjustBrightness(-theme.GradientAdjustment))
}
