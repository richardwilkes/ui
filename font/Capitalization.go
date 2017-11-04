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
	"github.com/richardwilkes/toolbox/i18n"
	// #cgo pkg-config: pangocairo
	// #include <pango/pangocairo.h>
	"C"
)

// Possible capitalizations.
const (
	CapitalizationNormal Capitalization = C.PANGO_VARIANT_NORMAL
	CapitalizationSmall  Capitalization = C.PANGO_VARIANT_SMALL_CAPS
)

// Capitalization is an enumeration of capitalization possibilities.
type Capitalization C.PangoVariant

// String returns the name of the capitalization.
func (c Capitalization) String() string {
	switch c {
	case CapitalizationNormal:
		return i18n.Text("Normal")
	case CapitalizationSmall:
		return i18n.Text("Small-Caps")
	default:
		return i18n.Text("Unknown")
	}
}
