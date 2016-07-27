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
	"github.com/richardwilkes/ui/draw"
	"time"
)

var (
	// StdScrollBar is the theme all new ScrollBars get by default.
	StdScrollBar = NewScrollBar()
)

// ScrollBar contains the theme elements for ScrollBars.
type ScrollBar struct {
	InitialRepeatDelay    time.Duration // The amount of time to wait before triggering the first repeating event.
	RepeatDelay           time.Duration // The amount of time to wait before triggering a repeating event.
	Background            color.Color   // The background color when enabled but not pressed or focused.
	BackgroundWhenPressed color.Color   // The background color when enabled and pressed.
	MarkWhenLight         color.Color   // The color to use for control marks when the background is considered to be 'light'.
	MarkWhenDark          color.Color   // The color to use for control marks when the background is considered to be 'dark'.
	MarkWhenDisabled      color.Color   // The color to use for control marks when disabled.
	GradientAdjustment    float32       // The amount to vary the color when creating the background gradient.
	DisabledAdjustment    float32       // The amount to adjust the background brightness when disabled.
	OutlineAdjustment     float32       // The amount to adjust the background brightness when using it to draw the button outline.
	Size                  float32       // The height of a horizontal scrollbar or the width of a vertical scrollbar.
}

// NewScrollBar creates a new image button theme.
func NewScrollBar() *ScrollBar {
	theme := &ScrollBar{}
	theme.Init()
	return theme
}

// Init initializes the theme with its default values.
func (theme *ScrollBar) Init() {
	theme.InitialRepeatDelay = time.Millisecond * 250
	theme.RepeatDelay = time.Millisecond * 75
	theme.Background = color.White
	theme.BackgroundWhenPressed = color.KeyboardFocus
	theme.MarkWhenLight = color.Black
	theme.MarkWhenDark = color.White
	theme.MarkWhenDisabled = color.Gray
	theme.GradientAdjustment = 0.15
	theme.DisabledAdjustment = -0.05
	theme.OutlineAdjustment = -0.5
	theme.Size = 16
}

// Gradient returns a gradient for the specified color.
func (theme *ScrollBar) Gradient(base color.Color) *draw.Gradient {
	return draw.NewEvenlySpacedGradient(base.AdjustBrightness(theme.GradientAdjustment), base.AdjustBrightness(-theme.GradientAdjustment))
}
