// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package font

import (
	"github.com/richardwilkes/i18n"
	// #cgo pkg-config: pangocairo
	// #include <pango/pangocairo.h>
	"C"
)

// Possible font slant variants.
const (
	SlantNormal  Slant = C.PANGO_STYLE_NORMAL
	SlantOblique Slant = C.PANGO_STYLE_OBLIQUE
	SlantItalic  Slant = C.PANGO_STYLE_ITALIC
)

// Slant represents the angle that a font is skewed when drawn.
type Slant C.PangoStyle

// String returns the name of the slant.
func (s Slant) String() string {
	switch s {
	case SlantNormal:
		return i18n.Text("Normal")
	case SlantOblique:
		return i18n.Text("Oblique")
	case SlantItalic:
		return i18n.Text("Italic")
	default:
		return i18n.Text("Unknown")
	}
}
