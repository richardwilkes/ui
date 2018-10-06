// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package window

import (
	"time"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/cursor"
)

// Window represents a window on the display.
type Window commonWindow

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

func platformNewWindow(bounds geom.Rect, styleMask StyleMask) *Window {
	// RAW: Implement for Windows
	return nil
}

func platformNewPopupWindow(parent ui.Window, bounds geom.Rect) *Window {
	// RAW: Implement for Windows
	return nil
}

func (window *Window) platformClose() {
	// RAW: Implement for Windows
	window.Dispose()
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

func (window *Window) platformMinimize() {
	// RAW: Implement for Windows
}

func (window *Window) platformZoom() {
	// RAW: Implement for Windows
}

func (window *Window) platformSetCursor(c *cursor.Cursor) {
	// RAW: Implement for Windows
}

func (window *Window) platformInvoke(id uint64) {
	// RAW: Implement for Windows
}

func (window *Window) platformInvokeAfter(id uint64, after time.Duration) {
	// RAW: Implement for Windows
}
