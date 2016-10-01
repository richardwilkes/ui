// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package scrollbar

// Pager objects can provide line and page information for scrolling.
type Pager interface {
	// LineScrollAmount is called to determine how far to scroll in the given direction to reveal
	// a full 'line' of content. A positive value should be returned regardless of the direction,
	// although negative values will behave as if they were positive.
	LineScrollAmount(horizontal, towardsStart bool) float64
	// PageScrollAmount is called to determine how far to scroll in the given direction to reveal
	// a full 'page' of content. A positive value should be returned regardless of the direction,
	// although negative values will behave as if they were positive.
	PageScrollAmount(horizontal, towardsStart bool) float64
}
