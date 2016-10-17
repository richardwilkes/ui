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
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/cursor"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/internal/task"
	"github.com/richardwilkes/ui/keys"
	"time"
	"unsafe"
	// #cgo LDFLAGS: -framework Cocoa -framework Quartz
	// #cgo pkg-config: pangocairo
	// #include <stdlib.h>
	// #include "Window_darwin.h"
	"C"
)

func platformGetKeyWindow() platformWindow {
	return platformWindow(C.platformGetKeyWindow())
}

func platformBringAllWindowsToFront() {
	C.platformBringAllWindowsToFront()
}

func platformHideCursorUntilMouseMoves() {
	C.platformHideCursorUntilMouseMoves()
}

func platformNewWindow(bounds geom.Rect, styleMask WindowStyleMask) (window platformWindow, surface *draw.Surface) {
	return platformWindow(C.platformNewWindow(C.double(bounds.X), C.double(bounds.Y), C.double(bounds.Width), C.double(bounds.Height), C.int(styleMask))), nil
}

func platformNewMenuWindow(parent ui.Window, bounds geom.Rect) (window platformWindow, surface *draw.Surface) {
	return platformNewWindow(bounds, BorderlessWindowMask)
}

func (window *Wnd) platformClose() {
	C.platformCloseWindow(window.window)
}

func (window *Wnd) platformTitle() string {
	return C.GoString(C.platformGetWindowTitle(window.window))
}

func (window *Wnd) platformSetTitle(title string) {
	cTitle := C.CString(title)
	C.platformSetWindowTitle(window.window, cTitle)
	C.free(unsafe.Pointer(cTitle))
}

func (window *Wnd) platformFrame() geom.Rect {
	var bounds geom.Rect
	C.platformGetWindowFrame(window.window, (*C.double)(&bounds.X), (*C.double)(&bounds.Y), (*C.double)(&bounds.Width), (*C.double)(&bounds.Height))
	return bounds
}

func (window *Wnd) platformSetFrame(bounds geom.Rect) {
	C.platformSetWindowFrame(window.window, C.double(bounds.X), C.double(bounds.Y), C.double(bounds.Width), C.double(bounds.Height))
}

func (window *Wnd) platformContentFrame() geom.Rect {
	var bounds geom.Rect
	C.platformGetWindowContentFrame(window.window, (*C.double)(&bounds.X), (*C.double)(&bounds.Y), (*C.double)(&bounds.Width), (*C.double)(&bounds.Height))
	return bounds
}

func (window *Wnd) platformToFront() {
	C.platformBringWindowToFront(window.window)
}

func (window *Wnd) platformRepaint(bounds geom.Rect) {
	C.platformRepaintWindow(window.window, C.double(bounds.X), C.double(bounds.Y), C.double(bounds.Width), C.double(bounds.Height))
}

func (window *Wnd) platformFlushPainting() {
	C.platformFlushPainting(window.window)
}

func (window *Wnd) platformScalingFactor() float64 {
	return float64(C.platformGetWindowScalingFactor(window.window))
}

func (window *Wnd) platformMinimize() {
	C.platformMinimizeWindow(window.window)
}

func (window *Wnd) platformZoom() {
	C.platformZoomWindow(window.window)
}

func (window *Wnd) platformSetToolTip(tip string) {
	if tip != "" {
		cstr := C.CString(tip)
		C.platformSetToolTip(window.window, cstr)
		C.free(unsafe.Pointer(cstr))
	} else {
		C.platformSetToolTip(window.window, nil)
	}
}

func (window *Wnd) platformSetCursor(c *cursor.Cursor) {
	C.platformSetCursor(window.window, c.PlatformPtr())
}

func (window *Wnd) platformInvoke(id uint64) {
	C.platformInvoke(C.ulong(id))
}

func (window *Wnd) platformInvokeAfter(id uint64, after time.Duration) {
	C.platformInvokeAfter(C.ulong(id), C.long(after.Nanoseconds()))
}

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
func handleWindowMouseWheelEvent(cWindow platformWindow, keyModifiers int, x, y, dx, dy float64) {
	if window, ok := windowMap[cWindow]; ok {
		window.processMouseWheel(x, y, dx, dy, keys.Modifiers(keyModifiers))
	}
}

//export handleCursorUpdateEvent
func handleCursorUpdateEvent(cWindow platformWindow, keyModifiers int, x, y float64) {
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
func handleWindowKeyEvent(cWindow platformWindow, keyModifiers, keyCode int, chars *C.char, down, repeat bool) {
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
