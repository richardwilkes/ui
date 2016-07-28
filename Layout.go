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
	"github.com/richardwilkes/ui/geom"
)

// Sizer is called when no layout has been set for a widget. Returns the minimum, preferred, and
// maximum sizes of the widget. The hint's values will be either NoHint or a specific value
// if that particular dimension has already been determined.
type Sizer interface {
	Sizes(hint geom.Size) (min, pref, max geom.Size)
}

// The Layout interface should be implemented by objects that provide layout services.
type Layout interface {
	Sizer
	// Layout is called to layout the target and its children.
	Layout()
}
