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
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"unsafe"
	// #cgo pkg-config: pangocairo
	// #include <pango/pangocairo.h>
	// #include <stdlib.h>
	// #include "Types.h"
	"C"
)

//export drawWindow
func drawWindow(cWindow platformWindow, gc *C.cairo_t, bounds platformRect, inLiveResize bool) {
	if window, ok := windowMap[cWindow]; ok {
		window.paint(draw.NewGraphics(draw.CairoContext(unsafe.Pointer(gc))), bounds.toRect(), inLiveResize)
	}
}

//export windowResized
func windowResized(cWindow platformWindow) {
	if window, ok := windowMap[cWindow]; ok {
		window.root.SetSize(window.ContentFrame().Size)
	}
}

//export windowShouldClose
func windowShouldClose(cWindow platformWindow) bool {
	if window, ok := windowMap[cWindow]; ok {
		e := event.NewClosing(window)
		event.Dispatch(e)
		return !e.Aborted()
	}
	return true
}

//export windowDidClose
func windowDidClose(cWindow platformWindow) {
	if window, ok := windowMap[cWindow]; ok {
		event.Dispatch(event.NewClosed(window))
	}
	delete(windowMap, cWindow)
}

//export handleWindowMouseEvent
func handleWindowMouseEvent(cWindow platformWindow, eventType platformEventType, keyModifiers, button, clickCount int, x, y float64) {
	if window, ok := windowMap[cWindow]; ok {
		window.mouseEvent(eventType, event.KeyMask(keyModifiers), button, clickCount, x, y)
	}
}

//export handleWindowMouseWheelEvent
func handleWindowMouseWheelEvent(cWindow platformWindow, eventType platformEventType, keyModifiers int, x, y, dx, dy float64) {
	if window, ok := windowMap[cWindow]; ok {
		window.mouseWheelEvent(eventType, event.KeyMask(keyModifiers), x, y, dx, dy)
	}
}

//export handleCursorUpdateEvent
func handleCursorUpdateEvent(cWindow platformWindow, keyModifiers int, x, y float64) {
	if window, ok := windowMap[cWindow]; ok {
		window.cursorUpdateEvent(event.KeyMask(keyModifiers), x, y)
	}
}

//export handleWindowKeyEvent
func handleWindowKeyEvent(cWindow platformWindow, eventType platformEventType, keyModifiers, keyCode int, chars *C.char, repeat bool) {
	if window, ok := windowMap[cWindow]; ok {
		var ch rune
		runes := ([]rune)(C.GoString(chars))
		if len(runes) > 0 {
			ch = runes[0]
		} else {
			ch = 0
		}
		window.keyEvent(eventType, event.KeyMask(keyModifiers), keyCode, ch, repeat)
	}
}
