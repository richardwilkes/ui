// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

// CellFactory defines methods all cell factories must implement.
type CellFactory interface {
	// CellHeight returns the height to use for the cells. A value less than 1 indicates that each
	// cell's height may be different.
	CellHeight() float32

	// CreateCell creates a new cell for 'owner' using 'element' as the content. 'index' indicates
	// which row the element came from. 'selected' indicates the cell should be created in its
	// selected state. 'focused' indicates the cell should be created in its focused state.
	CreateCell(owner Widget, element interface{}, index int, selected, focused bool) Widget
}
