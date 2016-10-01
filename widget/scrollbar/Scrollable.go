// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package scrollbar

// Scrollable objects can respond to ScrollBars.
type Scrollable interface {
	Pager
	// ScrolledPosition is called to determine the current position of the Scrollable.
	ScrolledPosition(horizontal bool) float64
	// SetScrolledPosition is called to set the current position of the Scrollable.
	SetScrolledPosition(horizontal bool, position float64)
	// VisibleSize is called to determine the size of the visible portion of the Scrollable.
	VisibleSize(horizontal bool) float64
	// ContentSize is called to determine the total size of the Scrollable.
	ContentSize(horizontal bool) float64
}
