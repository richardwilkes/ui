// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package widget

import (
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
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
		window.paint(draw.NewGraphics(draw.CairoContext(unsafe.Pointer(gc))), toRect(C.platformRect(bounds)), inLiveResize)
	}
}

//export windowResized
func windowResized(cWindow platformWindow) {
	if window, ok := windowMap[cWindow]; ok {
		window.root.SetSize(window.ContentFrame().Size)
	}
}

//export windowGainedKey
func windowGainedKey(cWindow platformWindow) {
	if window, ok := windowMap[cWindow]; ok {
		event.Dispatch(event.NewFocusGained(window))
	}
}

//export windowLostKey
func windowLostKey(cWindow platformWindow) {
	if window, ok := windowMap[cWindow]; ok {
		event.Dispatch(event.NewFocusLost(window))
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
		window.mouseEvent(eventType, keys.Modifiers(keyModifiers), button, clickCount, x, y)
	}
}

//export handleWindowMouseWheelEvent
func handleWindowMouseWheelEvent(cWindow platformWindow, eventType platformEventType, keyModifiers int, x, y, dx, dy float64) {
	if window, ok := windowMap[cWindow]; ok {
		window.mouseWheelEvent(eventType, keys.Modifiers(keyModifiers), x, y, dx, dy)
	}
}

//export handleCursorUpdateEvent
func handleCursorUpdateEvent(cWindow platformWindow, keyModifiers int, x, y float64) {
	if window, ok := windowMap[cWindow]; ok {
		window.cursorUpdateEvent(keys.Modifiers(keyModifiers), x, y)
	}
}

//export handleWindowKeyEvent
func handleWindowKeyEvent(cWindow platformWindow, eventType platformEventType, keyModifiers, keyCode int, chars *C.char, repeat bool) {
	if window, ok := windowMap[cWindow]; ok {
		var keyChar rune
		extractKeyChar := true
		if mapping := keys.MappingForScanCode(keyCode); mapping != nil {
			keyCode = mapping.KeyCode
			if !mapping.Dynamic {
				keyChar = mapping.KeyChar
				extractKeyChar = false
			}
		}
		if extractKeyChar && chars != nil && *chars != 0 {
			keyChar = (([]rune)(C.GoString(chars)))[0]
		}
		window.keyEvent(eventType, keys.Modifiers(keyModifiers), keyCode, keyChar, repeat)
	}
}

//export dispatchTask
func dispatchTask(id uint64) {
	if f := removeTask(id); f != nil {
		f()
	}
}
