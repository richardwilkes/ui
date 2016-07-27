// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

import (
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/font"
)

// BaseTextButtonTheme contains the common theme elements used in all buttons that display text.
type BaseTextButtonTheme struct {
	BaseButtonTheme
	TextWhenLight    color.Color // The text color to use when the background is considered to be 'light'.
	TextWhenDark     color.Color // The text color to use when the background is considered to be 'dark'.
	TextWhenDisabled color.Color // The text color to use when disabled.
	Font             *font.Font  // The font to use.
}

// Init initializes the theme with its default values.
func (theme *BaseTextButtonTheme) Init() {
	theme.BaseButtonTheme.Init()
	theme.TextWhenLight = color.Black
	theme.TextWhenDark = color.White
	theme.TextWhenDisabled = color.Gray
	theme.Font = font.Acquire(font.SystemDesc)
}
