// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package scrollarea

import (
	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui/border"
	"github.com/richardwilkes/ui/color"
)

var (
	// StdTheme is the theme all new ScrollAreas get by default.
	StdTheme = NewTheme()
)

// Theme contains the theme elements for ScrollAreas.
type Theme struct {
	Border      border.Border // The border to use when not focused.
	FocusBorder border.Border // The border to use when focused.
}

// NewTheme creates a new ScrollArea theme.
func NewTheme() *Theme {
	theme := &Theme{}
	theme.Init()
	return theme
}

// Init initializes the theme with its default values.
func (theme *Theme) Init() {
	theme.Border = border.NewLine(color.Background.AdjustBrightness(-0.25), geom.NewUniformInsets(1))
	lineBorder := border.NewLine(color.KeyboardFocus, geom.NewUniformInsets(2))
	lineBorder.NoInset = true
	theme.FocusBorder = lineBorder
}
