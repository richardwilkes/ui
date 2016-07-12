// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

// EmptyBorder is a Border that contains empty space, effectively providing an empty margin.
type EmptyBorder struct {
	insets Insets
}

// NewEmptyBorder creates a new Empty Border with the specified insets.
func NewEmptyBorder(insets Insets) Border {
	return &EmptyBorder{insets: insets}
}

// Insets implements the Border interface.
func (e *EmptyBorder) Insets() Insets {
	return e.insets
}

// PaintBorder implements the Border interface.
func (e *EmptyBorder) PaintBorder(g Graphics, bounds Rect) {
	// Does nothing
}
