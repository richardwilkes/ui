// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package border

import (
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/geom"
)

// The Border interface should be implemented by objects that provide a border around an area.
type Border interface {
	// Insets returns the insets describing the space the border occupies on each side.
	Insets() geom.Insets
	// Draw the border into 'bounds'.
	Draw(gc *draw.Graphics, bounds geom.Rect)
}
