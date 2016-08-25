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
	"github.com/richardwilkes/geom"
)

var (
	// StdScrollArea is the theme all new ScrollAreas get by default.
	StdScrollArea = NewScrollArea()
)

// ScrollArea contains the theme elements for ScrollAreas.
type ScrollArea struct {
	Border      border.Border // The border to use when not focused.
	FocusBorder border.Border // The border to use when focused.
}

// NewScrollArea creates a new ScrollArea theme.
func NewScrollArea() *ScrollArea {
	theme := &ScrollArea{}
	theme.Init()
	return theme
}

// Init initializes the theme with its default values.
func (theme *ScrollArea) Init() {
	theme.Border = border.NewLine(color.Background.AdjustBrightness(-0.25), geom.Insets{Top: 1, Left: 1, Bottom: 1, Right: 1})
	lineBorder := border.NewLine(color.KeyboardFocus, geom.Insets{Top: 2, Left: 2, Bottom: 2, Right: 2})
	lineBorder.NoInset = true
	theme.FocusBorder = lineBorder
}
