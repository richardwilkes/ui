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
	"fmt"
	"unsafe"
)

// #cgo darwin LDFLAGS: -framework Cocoa
// #include <stdlib.h>
// #include "Window.h"
import "C"

// Possible values for the WindowStyleMask.
const (
	BorderlessWindowMask             WindowStyleMask = 0
	TitledWindowMask                                 = 1 << 0
	ClosableWindowMask                               = 1 << 1
	MiniaturizableWindowMask                         = 1 << 2
	ResizableWindowMask                              = 1 << 3
	TexturedBackgroundWindowMask                     = 1 << 8
	UnifiedTitleAndToolbarWindowMask                 = 1 << 12
	FullScreenWindowMask                             = 1 << 14
	FullSizeContentViewWindowMask                    = 1 << 15
	StdWindowMask                                    = TitledWindowMask | ClosableWindowMask | MiniaturizableWindowMask | ResizableWindowMask
)

// WindowStyleMask controls the look and capabilities of a window.
type WindowStyleMask int

// Window represents a window on the display.
type Window struct {
	window         C.uiWindow
	ShouldClose    func() bool // Called to ask if closing the window is permitted. Return true if it is.
	DidClose       func()      // Called when the window has been closed.
	rootBlock      *Block
	lastMouseBlock Widget
	lastToolTip    string
	inMouseDown    bool
	inLiveResize   bool
}

var (
	windowMap = make(map[C.uiWindow]*Window)
)

// NewWindow creates a new window at the specified location with the specified style.
func NewWindow(where Point, styleMask WindowStyleMask) *Window {
	return NewWindowWithContentSize(where, Size{Width: 100, Height: 100}, styleMask)
}

// NewWindowWithContentSize creates a new window at the specified location with the specified style and content size.
func NewWindowWithContentSize(where Point, contentSize Size, styleMask WindowStyleMask) *Window {
	bounds := Rect{Point: where, Size: contentSize}
	window := &Window{window: C.uiNewWindow(toCRect(bounds), C.int(styleMask))}
	windowMap[window.window] = window
	window.rootBlock = NewBlock()
	window.rootBlock.SetBackground(BackgroundColor)
	window.rootBlock.window = window
	window.rootBlock.bounds = window.ContentLocalBounds()
	return window
}

//export drawWindow
func drawWindow(cWindow C.uiWindow, g unsafe.Pointer, bounds C.uiRect, inLiveResize bool) {
	if window, ok := windowMap[cWindow]; ok {
		window.rootBlock.ValidateLayout()
		window.inLiveResize = inLiveResize
		window.rootBlock.Paint(newGraphics(g), toRect(bounds))
		window.inLiveResize = false
	}
}

// InLiveResize returns true if the window is being actively resized by the user at this point in
// time. If it is, expensive painting operations should be deferred if possible to give a smooth
// resizing experience.
func (window *Window) InLiveResize() bool {
	return window.inLiveResize
}

//export windowResized
func windowResized(cWindow C.uiWindow) {
	if window, ok := windowMap[cWindow]; ok {
		window.rootBlock.SetSize(window.ContentSize())
	}
}

//export handleWindowMouseEvent
func handleWindowMouseEvent(cWindow C.uiWindow, eventType, keyModifiers, button, clickCount int, x, y float32) {
	if window, ok := windowMap[cWindow]; ok {
		discardMouseDown := false
		where := Point{X: x, Y: y}
		var block Widget
		if window.inMouseDown {
			block = window.lastMouseBlock
		} else {
			block = window.rootBlock.WidgetAt(where)
			if block == nil {
				panic("block is nil")
			}
			if eventType == C.uiMouseMoved && block != window.lastMouseBlock {
				if window.lastMouseBlock != nil {
					if h := window.lastMouseBlock.MouseExitedHandler(); h != nil {
						h.OnMouseExited(keyModifiers)
					}
				}
				eventType = C.uiMouseEntered
			}
		}
		switch eventType {
		case C.uiMouseDown:
			if block.Enabled() {
				if h := block.MouseDownHandler(); h != nil {
					discardMouseDown = h.OnMouseDown(block.FromWindow(where), keyModifiers, button, clickCount)
				}
			}
		case C.uiMouseDragged:
			if block.Enabled() {
				if h := block.MouseDraggedHandler(); h != nil {
					h.OnMouseDragged(block.FromWindow(where), keyModifiers)
				}
			}
		case C.uiMouseUp:
			if block.Enabled() {
				if h := block.MouseUpHandler(); h != nil {
					h.OnMouseUp(block.FromWindow(where), keyModifiers)
				}
			}
		case C.uiMouseEntered:
			where = block.FromWindow(where)
			if h := block.MouseEnteredHandler(); h != nil {
				h.OnMouseEntered(where, keyModifiers)
			}
			window.updateToolTip(block, where)
		case C.uiMouseMoved:
			where = block.FromWindow(where)
			if h := block.MouseMovedHandler(); h != nil {
				h.OnMouseMoved(where, keyModifiers)
			}
			window.updateToolTip(block, where)
		case C.uiMouseExited:
			if h := block.MouseExitedHandler(); h != nil {
				h.OnMouseExited(keyModifiers)
			}
		default:
			panic(fmt.Sprintf("Unknown event type: %d", eventType))
		}
		window.lastMouseBlock = block
		if eventType == C.uiMouseDown {
			if !discardMouseDown {
				window.inMouseDown = true
			}
		} else if window.inMouseDown && eventType == C.uiMouseUp {
			window.inMouseDown = false
		}
	}
}

//export windowShouldClose
func windowShouldClose(cWindow C.uiWindow) bool {
	if window, ok := windowMap[cWindow]; ok {
		if window.ShouldClose != nil {
			return window.ShouldClose()
		}
	}
	return true
}

//export windowDidClose
func windowDidClose(cWindow C.uiWindow) {
	if window, ok := windowMap[cWindow]; ok {
		if window.DidClose != nil {
			window.DidClose()
		}
	}
	delete(windowMap, cWindow)
}

// AllWindowsToFront attempts to bring all of the application's windows to the foreground.
func AllWindowsToFront() {
	C.uiBringAllWindowsToFront()
}

// KeyWindow returns the window that currently has the keyboard focus, or nil if none of your
// application's windows has the keyboard focus.
func KeyWindow() *Window {
	if window, ok := windowMap[C.uiGetKeyWindow()]; ok {
		return window
	}
	return nil
}

// Title returns the title of this window.
func (window *Window) Title() string {
	cTitle := C.uiGetWindowTitle(window.window)
	title := C.GoString(cTitle)
	C.free(unsafe.Pointer(cTitle))
	return title
}

// SetTitle sets the title of this window.
func (window *Window) SetTitle(title string) {
	cTitle := C.CString(title)
	C.uiSetWindowTitle(window.window, cTitle)
	C.free(unsafe.Pointer(cTitle))
}

func (window *Window) updateToolTip(widget Widget, whereInWidget Point) {
	tooltip := ""
	if widget != nil {
		if th := widget.ToolTipHandler(); th != nil {
			tooltip = th.OnToolTip(whereInWidget)
		}
	}
	if window.lastToolTip != tooltip {
		if tooltip != "" {
			tip := C.CString(tooltip)
			C.uiSetToolTip(window.window, tip)
			C.free(unsafe.Pointer(tip))
		} else {
			C.uiSetToolTip(window.window, nil)
		}
		window.lastToolTip = tooltip
	}
}

// Frame returns the boundaries in display coordinates of the frame of this window (i.e. the area
// that includes both the content and its border and window controls).
func (window *Window) Frame() Rect {
	return toRect(C.uiGetWindowFrame(window.window))
}

// Location returns the upper left corner of the window in display coordinates.
func (window *Window) Location() Point {
	return toPoint(C.uiGetWindowPosition(window.window))
}

// SetLocation moves the upper left corner of the window to the specified point in display
// coordinates.
func (window *Window) SetLocation(pt Point) {
	C.uiSetWindowPosition(window.window, C.float(pt.X), C.float(pt.Y))
}

// Size returns the size of the window, including its frame and window controls.
func (window *Window) Size() Size {
	return toSize(C.uiGetWindowSize(window.window))
}

// SetSize sets the size of the window.
func (window *Window) SetSize(size Size) {
	C.uiSetWindowSize(window.window, C.float(size.Width), C.float(size.Height))
}

// ContentFrame returns the boundaries of the root content block of this window.
func (window *Window) ContentFrame() Rect {
	return toRect(C.uiGetWindowContentFrame(window.window))
}

// ContentLocalBounds returns the local boundaries of the content block of this window.
func (window *Window) ContentLocalBounds() Rect {
	size := C.uiGetWindowContentSize(window.window)
	return Rect{Size: Size{Width: float32(size.width), Height: float32(size.height)}}
}

// ContentLocation returns the upper left corner of the content block in display coordinates.
func (window *Window) ContentLocation() Point {
	return toPoint(C.uiGetWindowContentPosition(window.window))
}

// SetContentLocation moves the window such that the upper left corner of the content block is
// at the specified point in display coordinates.
func (window *Window) SetContentLocation(pt Point) {
	C.uiSetWindowContentPosition(window.window, C.float(pt.X), C.float(pt.Y))
}

// ContentSize returns the size of the content block.
func (window *Window) ContentSize() Size {
	return toSize(C.uiGetWindowContentSize(window.window))
}

// SetContentSize sets the size of the window to fit the specified content size.
func (window *Window) SetContentSize(size Size) {
	C.uiSetWindowContentSize(window.window, C.float(size.Width), C.float(size.Height))
}

// ScalingFactor returns the current OS scaling factor being applied to this window.
func (window *Window) ScalingFactor() float32 {
	return float32(C.uiGetWindowScalingFactor(window.window))
}

// Minimize performs the platform's minimize function on the window.
func (window *Window) Minimize() {
	C.uiMinimizeWindow(window.window)
}

// Zoom performs the platform's zoom funcion on the window.
func (window *Window) Zoom() {
	C.uiZoomWindow(window.window)
}

// ToFront attempts to bring the window to the foreground and give it the keyboard focus.
func (window *Window) ToFront() {
	C.uiBringWindowToFront(window.window)
}

// RootBlock returns the root Block of the window.
func (window *Window) RootBlock() Widget {
	return window.rootBlock
}

// Pack sets the window's content size to match the preferred size of the root block.
func (window *Window) Pack() {
	_, pref, _ := ComputeSizes(window.rootBlock, NoLayoutHintSize)
	window.SetContentSize(pref)
}

// Repaint marks this window for painting at the next update.
func (window *Window) Repaint() {
	C.uiRepaintWindow(window.window, toCRect(window.ContentLocalBounds()))
}

// RepaintBounds marks the specified bounds within the window for painting at the next update.
func (window *Window) RepaintBounds(bounds Rect) {
	bounds.Intersect(window.ContentLocalBounds())
	if !bounds.IsEmpty() {
		C.uiRepaintWindow(window.window, toCRect(bounds))
	}
}
