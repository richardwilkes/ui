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
	"fmt"
	"unsafe"

	"github.com/richardwilkes/go-ui/color"
	"github.com/richardwilkes/go-ui/geom"
	"github.com/richardwilkes/go-ui/graphics"
	"github.com/richardwilkes/go-ui/layout"
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

// Represents a window on the display.
type Window struct {
	window         C.uiWindow
	ShouldClose    func() bool // Called to ask if closing the window is permitted. Return true if it is.
	DidClose       func()      // Called when the window has been closed.
	rootBlock      *Block
	lastMouseBlock *Block
	lastToolTip    string
	inMouseDown    bool
}

var (
	windowMap = make(map[C.uiWindow]*Window)
)

// NewWindow creates a new window at the specified location with the specified style.
func NewWindow(where geom.Point, styleMask WindowStyleMask) *Window {
	return NewWindowWithContentSize(where, geom.Size{Width: 100, Height: 100}, styleMask)
}

// NewWindowWithContentSize creates a new window at the specified location with the specified style and content size.
func NewWindowWithContentSize(where geom.Point, contentSize geom.Size, styleMask WindowStyleMask) *Window {
	bounds := geom.Rect{Point: where, Size: contentSize}
	window := &Window{window: C.uiNewWindow(toCRect(bounds), C.int(styleMask))}
	windowMap[window.window] = window
	window.rootBlock = NewBlock()
	window.rootBlock.SetBackground(color.Background)
	window.rootBlock.window = window
	window.rootBlock.bounds = window.ContentLocalBounds()
	return window
}

//export drawWindow
func drawWindow(cWindow C.uiWindow, gc unsafe.Pointer, bounds C.uiRect, inLiveResize bool) {
	if window, ok := windowMap[cWindow]; ok {
		window.rootBlock.ValidateLayout()
		window.rootBlock.paint(graphics.NewContext(gc), toRect(bounds), inLiveResize)
	}
}

//export windowResized
func windowResized(cWindow C.uiWindow) {
	if window, ok := windowMap[cWindow]; ok {
		window.rootBlock.SetSize(window.ContentSize())
	}
}

//export handleMouseEvent
func handleMouseEvent(cWindow C.uiWindow, eventType, keyModifiers, button, clickCount int, x, y float32) {
	if window, ok := windowMap[cWindow]; ok {
		where := geom.Point{X: x, Y: y}
		var block *Block
		if window.inMouseDown {
			block = window.lastMouseBlock
		} else {
			block = window.rootBlock.BlockAt(where)
			if block == nil {
				panic("block is nil")
			}
			if eventType == C.uiMouseMoved && block != window.lastMouseBlock {
				if window.lastMouseBlock != nil && window.lastMouseBlock.OnMouseExited != nil {
					window.lastMouseBlock.OnMouseExited(window.lastMouseBlock.FromWindow(where), keyModifiers)
				}
				eventType = C.uiMouseEntered
			}
		}
		switch eventType {
		case C.uiMouseDown:
			if block.OnMouseDown != nil && block.Enabled() {
				block.OnMouseDown(block.FromWindow(where), keyModifiers, button, clickCount)
			}
		case C.uiMouseDragged:
			if block.OnMouseDragged != nil && block.Enabled() {
				block.OnMouseDragged(block.FromWindow(where), keyModifiers)
			}
		case C.uiMouseUp:
			if block.OnMouseUp != nil && block.Enabled() {
				block.OnMouseUp(block.FromWindow(where), keyModifiers)
			}
		case C.uiMouseEntered:
			where = block.FromWindow(where)
			if block.OnMouseEntered != nil {
				block.OnMouseEntered(where, keyModifiers)
			}
			window.setToolTip(block.ToolTip(where))
		case C.uiMouseMoved:
			where = block.FromWindow(where)
			if block.OnMouseMoved != nil {
				block.OnMouseMoved(where, keyModifiers)
			}
			window.setToolTip(block.ToolTip(where))
		case C.uiMouseExited:
			if block.OnMouseExited != nil {
				block.OnMouseExited(block.FromWindow(where), keyModifiers)
			}
		default:
			panic(fmt.Sprintf("Unknown event type: %d", eventType))
		}
		window.lastMouseBlock = block
		if eventType == C.uiMouseDown {
			window.inMouseDown = true
		} else if window.inMouseDown && eventType == C.uiMouseUp {
			window.inMouseDown = false
		}
	}
}

//export shouldClose
func shouldClose(cWindow C.uiWindow) bool {
	if window, ok := windowMap[cWindow]; ok {
		if window.ShouldClose != nil {
			return window.ShouldClose()
		}
	}
	return true
}

//export didClose
func didClose(cWindow C.uiWindow) {
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

func (window *Window) setToolTip(tooltip string) {
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
func (window *Window) Frame() geom.Rect {
	return toRect(C.uiGetWindowFrame(window.window))
}

// Location returns the upper left corner of the window in display coordinates.
func (window *Window) Location() geom.Point {
	return toPoint(C.uiGetWindowPosition(window.window))
}

// SetLocation moves the upper left corner of the window to the specified point in display
// coordinates.
func (window *Window) SetLocation(pt geom.Point) {
	C.uiSetWindowPosition(window.window, C.float(pt.X), C.float(pt.Y))
}

// Size returns the size of the window, including its frame and window controls.
func (window *Window) Size() geom.Size {
	return toSize(C.uiGetWindowSize(window.window))
}

// SetSize sets the size of the window.
func (window *Window) SetSize(size geom.Size) {
	C.uiSetWindowSize(window.window, C.float(size.Width), C.float(size.Height))
}

// ContentFrame returns the boundaries of the root content block of this window.
func (window *Window) ContentFrame() geom.Rect {
	return toRect(C.uiGetWindowContentFrame(window.window))
}

// ContentLocalBounds returns the local boundaries of the content block of this window.
func (window *Window) ContentLocalBounds() geom.Rect {
	size := C.uiGetWindowContentSize(window.window)
	return geom.Rect{Size: geom.Size{Width: float32(size.width), Height: float32(size.height)}}
}

// ContentLocation returns the upper left corner of the content block in display coordinates.
func (window *Window) ContentLocation() geom.Point {
	return toPoint(C.uiGetWindowContentPosition(window.window))
}

// SetContentLocation moves the window such that the upper left corner of the content block is
// at the specified point in display coordinates.
func (window *Window) SetContentLocation(pt geom.Point) {
	C.uiSetWindowContentPosition(window.window, C.float(pt.X), C.float(pt.Y))
}

// ContentSize returns the size of the content block.
func (window *Window) ContentSize() geom.Size {
	return toSize(C.uiGetWindowContentSize(window.window))
}

// SetContentSize sets the size of the window to fit the specified content size.
func (window *Window) SetContentSize(size geom.Size) {
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
func (window *Window) RootBlock() *Block {
	return window.rootBlock
}

// Pack sets the window's content size to match the preferred size of the root block.
func (window *Window) Pack() {
	_, pref, _ := window.rootBlock.ComputeSizes(layout.NoHintSize)
	window.SetContentSize(pref)
}

// Repaint marks this window for painting at the next update.
func (window *Window) Repaint() {
	C.uiRepaintWindow(window.window, toCRect(window.ContentLocalBounds()))
}

// RepaintBounds marks the specified bounds within the window for painting at the next update.
func (window *Window) RepaintBounds(bounds geom.Rect) {
	bounds.Intersect(window.ContentLocalBounds())
	if !bounds.IsEmpty() {
		C.uiRepaintWindow(window.window, toCRect(bounds))
	}
}

func toRect(bounds C.uiRect) geom.Rect {
	return geom.Rect{Point: geom.Point{X: float32(bounds.x), Y: float32(bounds.y)}, Size: geom.Size{Width: float32(bounds.width), Height: float32(bounds.height)}}
}

func toCRect(bounds geom.Rect) C.uiRect {
	return C.uiRect{x: C.float(bounds.X), y: C.float(bounds.Y), width: C.float(bounds.Width), height: C.float(bounds.Height)}
}

func toPoint(pt C.uiPoint) geom.Point {
	return geom.Point{X: float32(pt.x), Y: float32(pt.y)}
}

func toSize(size C.uiSize) geom.Size {
	return geom.Size{Width: float32(size.width), Height: float32(size.height)}
}
