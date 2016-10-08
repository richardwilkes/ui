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
	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui/cursor"
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

func platformNewWindow(bounds geom.Rect, styleMask WindowStyleMask) (window platformWindow, surface platformSurface) {
	// RAW: Implement for Windows
	return nil, nil
}

func platformNewMenuWindow(parent ui.Window, bounds geom.Rect) (window platformWindow, surface platformSurface) {
	// RAW: Implement for Windows
	return nil, nil
}

func (window *Wnd) platformClose() {
	// RAW: Implement for Windows
	windowDidClose(window.window)
}

func (window *Wnd) platformTitle() string {
	// RAW: Implement for Windows
	return ""
}

func (window *Wnd) platformSetTitle(title string) {
	// RAW: Implement for Windows
}

func (window *Wnd) platformFrame() geom.Rect {
	// RAW: Implement for Windows
	return geom.Rect{}
}

func (window *Wnd) platformSetFrame(bounds geom.Rect) {
	// RAW: Implement for Windows
}

func (window *Wnd) platformContentFrame() geom.Rect {
	// RAW: Implement for Windows
	return geom.Rect{}
}

func (window *Wnd) platformToFront() {
	// RAW: Implement for Windows
}

func (window *Wnd) platformRepaint(bounds geom.Rect) {
	// RAW: Implement for Windows
}

func (window *Wnd) platformFlushPainting() {
	// RAW: Implement for Windows
}

func (window *Wnd) platformScalingFactor() float64 {
	// RAW: Implement for Windows
	return 1
}

func (window *Wnd) platformMinimize() {
	// RAW: Implement for Windows
}

func (window *Wnd) platformZoom() {
	// RAW: Implement for Windows
}

func (window *Wnd) platformSetToolTip(tip string) {
	// RAW: Implement for Windows
}

func (window *Wnd) platformSetCursor(c *cursor.Cursor) {
	// RAW: Implement for Windows
}

func (window *Wnd) platformInvoke(id uint64) {
	// RAW: Implement for Windows
}

func (window *Wnd) platformInvokeAfter(id uint64, after time.Duration) {
	// RAW: Implement for Windows
}
