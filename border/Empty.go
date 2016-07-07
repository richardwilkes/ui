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
	"github.com/richardwilkes/go-ui/geom"
	"github.com/richardwilkes/go-ui/graphics"
)

// Empty is a Border that contains empty space, effectively providing an empty margin.
type Empty struct {
	insets geom.Insets
}

// NewEmpty creates a new Empty Border with the specified insets.
func NewEmpty(insets geom.Insets) Border {
	return &Empty{insets: insets}
}

// Insets -- implements the Border interface.
func (e *Empty) Insets() geom.Insets {
	return e.insets
}

// Paint -- implements the Border interface.
func (e *Empty) Paint(gc graphics.Context, bounds geom.Rect) {
	// Does nothing
}
