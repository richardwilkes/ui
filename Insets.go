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
	"fmt"
)

// Insets defines margins on each side of a rectangle.
type Insets struct {
	Top    float32
	Left   float32
	Bottom float32
	Right  float32
}

// Add modifies this Insets by adding the supplied Insets.
func (i *Insets) Add(insets Insets) {
	i.Top += insets.Top
	i.Left += insets.Left
	i.Bottom += insets.Bottom
	i.Right += insets.Right
}

// String implements the fmt.Stringer interface.
func (i Insets) String() string {
	return fmt.Sprintf("%v, %v, %v, %v", i.Top, i.Left, i.Bottom, i.Right)
}
