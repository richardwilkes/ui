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
	"github.com/richardwilkes/ui/cursor"
	"github.com/richardwilkes/ui/geom"
)

func platformGetKeyWindow() platformWindow {
	// RAW: Implement for Windows
	return nil
}

func platformBringAllWindowsToFront() {
	// RAW: Implement for Windows
}

func platformHideCursorUntilMouseMoves() {
	// RAW: Implement for Windows
}

func platformNewWindow(bounds geom.Rect, styleMask WindowStyleMask) platformWindow {
	// RAW: Implement for Windows
	return nil
}

func (window *Window) platformClose() {
	// RAW: Implement for Windows
}

func (window *Window) platformTitle() string {
	// RAW: Implement for Windows
	return ""
}

func (window *Window) platformSetTitle(title string) {
	// RAW: Implement for Windows
}

func (window *Window) platformFrame() geom.Rect {
	// RAW: Implement for Windows
	return geom.Rect{}
}

func (window *Window) platformSetFrame(bounds geom.Rect) {
	// RAW: Implement for Windows
}

func (window *Window) platformContentFrame() geom.Rect {
	// RAW: Implement for Windows
	return geom.Rect{}
}

func (window *Window) platformToFront() {
	// RAW: Implement for Windows
}

func (window *Window) platformRepaint(bounds geom.Rect) {
	// RAW: Implement for Windows
}

func (window *Window) platformFlushPainting() {
	// RAW: Implement for Windows
}

func (window *Window) platformScalingFactor() float32 {
	// RAW: Implement for Windows
	return 1
}

func (window *Window) platformMinimize() {
	// RAW: Implement for Windows
}

func (window *Window) platformZoom() {
	// RAW: Implement for Windows
}

func (window *Window) platformSetToolTip(tip string) {
	// RAW: Implement for Windows
}

func (window *Window) platformSetCursor(c *cursor.Cursor) {
	// RAW: Implement for Windows
}
