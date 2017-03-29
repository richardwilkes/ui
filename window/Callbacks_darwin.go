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
	"unsafe"

	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/internal/task"
	"github.com/richardwilkes/ui/keys"
	// #cgo pkg-config: cairo
	// #include <cairo.h>
	"C"
)

//export drawWindow
func drawWindow(cWindow platformWindow, gc *C.cairo_t, x, y, width, height float64) {
	if window, ok := windowMap[cWindow]; ok {
		window.paint(draw.NewGraphics(draw.CairoContext(unsafe.Pointer(gc))), geom.Rect{Point: geom.Point{X: x, Y: y}, Size: geom.Size{Width: width, Height: height}})
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
		return window.MayClose()
	}
	return true
}

//export windowDidClose
func windowDidClose(cWindow platformWindow) {
	if window, ok := windowMap[cWindow]; ok {
		window.Dispose()
	}
}

//export handleMouseDownEvent
func handleMouseDownEvent(cWindow platformWindow, x, y float64, button, clickCount, keyModifiers int) {
	if window, ok := windowMap[cWindow]; ok {
		window.processMouseDown(x, y, button, clickCount, keys.Modifiers(keyModifiers))
	}
}

//export handleMouseDraggedEvent
func handleMouseDraggedEvent(cWindow platformWindow, x, y float64, button, keyModifiers int) {
	if window, ok := windowMap[cWindow]; ok {
		window.processMouseDragged(x, y, button, keys.Modifiers(keyModifiers))
	}
}

//export handleMouseUpEvent
func handleMouseUpEvent(cWindow platformWindow, x, y float64, button, keyModifiers int) {
	if window, ok := windowMap[cWindow]; ok {
		window.processMouseUp(x, y, button, keys.Modifiers(keyModifiers))
	}
}

//export handleMouseEnteredEvent
func handleMouseEnteredEvent(cWindow platformWindow, x, y float64, keyModifiers int) {
	if window, ok := windowMap[cWindow]; ok {
		window.processMouseEntered(x, y, keys.Modifiers(keyModifiers))
	}
}

//export handleMouseMovedEvent
func handleMouseMovedEvent(cWindow platformWindow, x, y float64, keyModifiers int) {
	if window, ok := windowMap[cWindow]; ok {
		window.processMouseMoved(x, y, keys.Modifiers(keyModifiers))
	}
}

//export handleMouseExitedEvent
func handleMouseExitedEvent(cWindow platformWindow, x, y float64, keyModifiers int) {
	if window, ok := windowMap[cWindow]; ok {
		window.processMouseExited(x, y, keys.Modifiers(keyModifiers))
	}
}

//export handleWindowMouseWheelEvent
func handleWindowMouseWheelEvent(cWindow platformWindow, x, y, dx, dy float64, keyModifiers int) {
	if window, ok := windowMap[cWindow]; ok {
		window.processMouseWheel(x, y, dx, dy, keys.Modifiers(keyModifiers))
	}
}

//export handleCursorUpdateEvent
func handleCursorUpdateEvent(cWindow platformWindow, x, y float64, keyModifiers int) {
	if window, ok := windowMap[cWindow]; ok {
		where := geom.Point{X: x, Y: y}
		var widget ui.Widget
		if window.inMouseDown {
			widget = window.lastMouseWidget
		} else {
			widget = window.root.WidgetAt(where)
			if widget == nil {
				panic("widget is nil")
			}
		}
		window.updateToolTipAndCursor(widget, where)
	}
}

//export handleWindowKeyEvent
func handleWindowKeyEvent(cWindow platformWindow, keyCode int, chars *C.char, keyModifiers int, down, repeat bool) {
	if window, ok := windowMap[cWindow]; ok {
		var str string
		if chars != nil {
			str = C.GoString(chars)
		}
		code, ch := keys.Transform(keyCode, str)
		modifiers := keys.Modifiers(keyModifiers)
		if down {
			window.processKeyDown(code, ch, modifiers, repeat)
		} else {
			window.processKeyUp(code, modifiers)
		}
	}
}

//export dispatchTask
func dispatchTask(id uint64) {
	task.Dispatch(id)
}
