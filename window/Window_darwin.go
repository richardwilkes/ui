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
	// #cgo darwin LDFLAGS: -framework Cocoa -framework Quartz
	// #cgo pkg-config: pangocairo
	// #include <stdlib.h>
	// #include "Window_darwin.h"
	"C"
)

func toRect(r C.platformRect) geom.Rect {
	return geom.Rect{Point: geom.Point{X: float64(r.x), Y: float64(r.y)}, Size: geom.Size{Width: float64(r.width), Height: float64(r.height)}}
}

func toCRect(r geom.Rect) C.platformRect {
	return C.platformRect{x: C.double(r.X), y: C.double(r.Y), width: C.double(r.Width), height: C.double(r.Height)}
}

func platformGetKeyWindow() platformWindow {
	return platformWindow(C.platformGetKeyWindow())
}

func platformBringAllWindowsToFront() {
	C.platformBringAllWindowsToFront()
}

func platformHideCursorUntilMouseMoves() {
	C.platformHideCursorUntilMouseMoves()
}

func platformNewWindow(bounds geom.Rect, styleMask WindowStyleMask) (window platformWindow, surface platformSurface) {
	return platformWindow(C.platformNewWindow(toCRect(bounds), C.int(styleMask))), nil
}

func platformNewMenuWindow(parent ui.Window, bounds geom.Rect) (window platformWindow, surface platformSurface) {
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
	return toRect(C.platformGetWindowFrame(window.window))
}

func (window *Wnd) platformSetFrame(bounds geom.Rect) {
	C.platformSetWindowFrame(window.window, toCRect(bounds))
}

func (window *Wnd) platformContentFrame() geom.Rect {
	return toRect(C.platformGetWindowContentFrame(window.window))
}

func (window *Wnd) platformToFront() {
	C.platformBringWindowToFront(window.window)
}

func (window *Wnd) platformRepaint(bounds geom.Rect) {
	C.platformRepaintWindow(window.window, toCRect(bounds))
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
func drawWindow(cWindow platformWindow, gc *C.cairo_t, bounds platformRect) {
	if window, ok := windowMap[cWindow]; ok {
		window.paint(draw.NewGraphics(draw.CairoContext(unsafe.Pointer(gc))), toRect(C.platformRect(bounds)))
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

//export handleWindowMouseEvent
func handleWindowMouseEvent(cWindow platformWindow, eventType platformEventType, keyModifiers, button, clickCount int, x, y float64) {
	if window, ok := windowMap[cWindow]; ok {
		window.mouseEvent(eventType, keys.Modifiers(keyModifiers), button, clickCount, x, y)
	}
}

//export handleWindowMouseWheelEvent
func handleWindowMouseWheelEvent(cWindow platformWindow, keyModifiers int, x, y, dx, dy float64) {
	if window, ok := windowMap[cWindow]; ok {
		window.mouseWheelEvent(keys.Modifiers(keyModifiers), x, y, dx, dy)
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
		var str string
		if chars != nil {
			str = C.GoString(chars)
		}
		code, ch := keys.Transform(keyCode, str)
		window.keyEvent(eventType, keys.Modifiers(keyModifiers), code, ch, repeat)
	}
}

//export dispatchTask
func dispatchTask(id uint64) {
	task.Dispatch(id)
}
