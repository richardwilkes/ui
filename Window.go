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
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/keys"
	"time"
	"unsafe"
)

// #cgo darwin LDFLAGS: -framework Cocoa -framework Quartz
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
	window          C.uiWindow
	ShouldClose     func() bool // Called to ask if closing the window is permitted. Return true if it is.
	DidClose        func()      // Called when the window has been closed.
	rootBlock       *Block
	focus           Widget
	lastMouseWidget Widget
	lastToolTip     string
	inMouseDown     bool
	inLiveResize    bool
}

var (
	windowMap = make(map[C.uiWindow]*Window)
)

// NewWindow creates a new window at the specified location with the specified style.
func NewWindow(where draw.Point, styleMask WindowStyleMask) *Window {
	return NewWindowWithContentSize(where, draw.Size{Width: 100, Height: 100}, styleMask)
}

// NewWindowWithContentSize creates a new window at the specified location with the specified style and content size.
func NewWindowWithContentSize(where draw.Point, contentSize draw.Size, styleMask WindowStyleMask) *Window {
	bounds := draw.Rect{Point: where, Size: contentSize}
	window := &Window{window: C.uiNewWindow(toCRect(bounds), C.int(styleMask))}
	windowMap[window.window] = window
	window.rootBlock = NewBlock()
	window.rootBlock.SetBackground(color.Background)
	window.rootBlock.window = window
	window.rootBlock.bounds = window.ContentLocalBounds()
	return window
}

//export drawWindow
func drawWindow(cWindow C.uiWindow, g unsafe.Pointer, bounds C.uiRect, inLiveResize bool) {
	if window, ok := windowMap[cWindow]; ok {
		window.rootBlock.ValidateLayout()
		window.inLiveResize = inLiveResize
		window.rootBlock.Paint(draw.NewGraphics(g), toRect(bounds))
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
		keyMask := KeyMask(keyModifiers)
		discardMouseDown := false
		where := draw.Point{X: x, Y: y}
		var widget Widget
		if window.inMouseDown {
			widget = window.lastMouseWidget
		} else {
			widget = window.rootBlock.WidgetAt(where)
			if widget == nil {
				panic("widget is nil")
			}
			if eventType == C.uiMouseMoved && widget != window.lastMouseWidget {
				if window.lastMouseWidget != nil {
					event := &Event{Type: MouseExitedEvent, When: time.Now(), Target: window.lastMouseWidget, KeyModifiers: keyMask}
					event.Dispatch()
				}
				eventType = C.uiMouseEntered
			}
		}
		switch eventType {
		case C.uiMouseDown:
			if widget.Enabled() {
				event := &Event{Type: MouseDownEvent, When: time.Now(), Target: widget, Where: where, KeyModifiers: keyMask, Button: button, Clicks: clickCount}
				event.Dispatch()
				discardMouseDown = event.Discard
			}
		case C.uiMouseDragged:
			if widget.Enabled() {
				event := &Event{Type: MouseDraggedEvent, When: time.Now(), Target: widget, Where: where, KeyModifiers: keyMask}
				event.Dispatch()
			}
		case C.uiMouseUp:
			if widget.Enabled() {
				event := &Event{Type: MouseUpEvent, When: time.Now(), Target: widget, Where: where, KeyModifiers: keyMask}
				event.Dispatch()
			}
		case C.uiMouseEntered:
			event := &Event{Type: MouseEnteredEvent, When: time.Now(), Target: widget, Where: where, KeyModifiers: keyMask}
			event.Dispatch()
			window.updateToolTip(widget, where)
		case C.uiMouseMoved:
			event := &Event{Type: MouseMovedEvent, When: time.Now(), Target: widget, Where: where, KeyModifiers: keyMask}
			event.Dispatch()
			window.updateToolTip(widget, where)
		case C.uiMouseExited:
			event := &Event{Type: MouseExitedEvent, When: time.Now(), Target: widget, KeyModifiers: keyMask}
			event.Dispatch()
		default:
			panic(fmt.Sprintf("Unknown event type: %d", eventType))
		}
		window.lastMouseWidget = widget
		if eventType == C.uiMouseDown {
			if !discardMouseDown {
				window.inMouseDown = true
			}
		} else if window.inMouseDown && eventType == C.uiMouseUp {
			window.inMouseDown = false
		}
	}
}

//export handleWindowMouseWheelEvent
func handleWindowMouseWheelEvent(cWindow C.uiWindow, eventType, keyModifiers int, x, y, dx, dy float32) {
	if window, ok := windowMap[cWindow]; ok {
		where := draw.Point{X: x, Y: y}
		widget := window.rootBlock.WidgetAt(where)
		if widget != nil {
			event := &Event{Type: MouseWheelEvent, When: time.Now(), Target: widget, Where: where, Delta: draw.Point{X: dx, Y: dy}, KeyModifiers: KeyMask(keyModifiers), CascadeUp: true}
			event.Dispatch()
			if window.inMouseDown {
				eventType = C.uiMouseDragged
			} else {
				eventType = C.uiMouseMoved
			}
			handleWindowMouseEvent(cWindow, eventType, keyModifiers, 0, 0, x, y)
		}
	}
}

//export handleWindowKeyEvent
func handleWindowKeyEvent(cWindow C.uiWindow, eventType, keyModifiers, keyCode int, chars *C.char, repeat bool) {
	if window, ok := windowMap[cWindow]; ok {
		keyMask := KeyMask(keyModifiers)
		switch eventType {
		case C.uiKeyDown:
			event := &Event{Type: KeyDownEvent, When: time.Now(), Target: window.Focus(), KeyModifiers: keyMask, KeyCode: keyCode, Repeat: repeat, CascadeUp: true}
			event.Dispatch()
			if !event.Discard && keyCode == keys.Tab && (keyMask&(AllKeyMask & ^ShiftKeyMask)) == 0 {
				if keyMask&ShiftKeyMask == 0 {
					window.FocusNext()
				} else {
					window.FocusPrevious()
				}
			}
		case C.uiKeyTyped:
			for _, r := range C.GoString(chars) {
				event := &Event{Type: KeyTypedEvent, When: time.Now(), Target: window.Focus(), KeyModifiers: keyMask, KeyTyped: r, Repeat: repeat, CascadeUp: true}
				event.Dispatch()
			}
		case C.uiKeyUp:
			event := &Event{Type: KeyUpEvent, When: time.Now(), Target: window.Focus(), KeyModifiers: keyMask, KeyCode: keyCode, Repeat: repeat, CascadeUp: true}
			event.Dispatch()
		default:
			panic(fmt.Sprintf("Unknown event type: %d", eventType))
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
	return C.GoString(C.uiGetWindowTitle(window.window))
}

// SetTitle sets the title of this window.
func (window *Window) SetTitle(title string) {
	cTitle := C.CString(title)
	C.uiSetWindowTitle(window.window, cTitle)
	C.free(unsafe.Pointer(cTitle))
}

func (window *Window) updateToolTip(widget Widget, where draw.Point) {
	tooltip := ""
	if widget != nil {
		event := &Event{Type: ToolTipEvent, When: time.Now(), Target: widget, Where: where}
		event.Dispatch()
		tooltip = event.ToolTip
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
func (window *Window) Frame() draw.Rect {
	return toRect(C.uiGetWindowFrame(window.window))
}

// Location returns the upper left corner of the window in display coordinates.
func (window *Window) Location() draw.Point {
	return toPoint(C.uiGetWindowPosition(window.window))
}

// SetLocation moves the upper left corner of the window to the specified point in display
// coordinates.
func (window *Window) SetLocation(pt draw.Point) {
	C.uiSetWindowPosition(window.window, C.float(pt.X), C.float(pt.Y))
}

// Size returns the size of the window, including its frame and window controls.
func (window *Window) Size() draw.Size {
	return toSize(C.uiGetWindowSize(window.window))
}

// SetSize sets the size of the window.
func (window *Window) SetSize(size draw.Size) {
	C.uiSetWindowSize(window.window, C.float(size.Width), C.float(size.Height))
}

// ContentFrame returns the boundaries of the root content widget of this window.
func (window *Window) ContentFrame() draw.Rect {
	return toRect(C.uiGetWindowContentFrame(window.window))
}

// ContentLocalBounds returns the local boundaries of the content widget of this window.
func (window *Window) ContentLocalBounds() draw.Rect {
	size := C.uiGetWindowContentSize(window.window)
	return draw.Rect{Size: draw.Size{Width: float32(size.width), Height: float32(size.height)}}
}

// ContentLocation returns the upper left corner of the content widget in display coordinates.
func (window *Window) ContentLocation() draw.Point {
	return toPoint(C.uiGetWindowContentPosition(window.window))
}

// SetContentLocation moves the window such that the upper left corner of the content widget is
// at the specified point in display coordinates.
func (window *Window) SetContentLocation(pt draw.Point) {
	C.uiSetWindowContentPosition(window.window, C.float(pt.X), C.float(pt.Y))
}

// ContentSize returns the size of the content widget.
func (window *Window) ContentSize() draw.Size {
	return toSize(C.uiGetWindowContentSize(window.window))
}

// SetContentSize sets the size of the window to fit the specified content size.
func (window *Window) SetContentSize(size draw.Size) {
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

// RootWidget returns the root widget of the window.
func (window *Window) RootWidget() Widget {
	return window.rootBlock
}

// Pack sets the window's content size to match the preferred size of the root widget.
func (window *Window) Pack() {
	_, pref, _ := ComputeSizes(window.rootBlock, NoLayoutHintSize)
	window.SetContentSize(pref)
}

// Repaint marks this window for painting at the next update.
func (window *Window) Repaint() {
	C.uiRepaintWindow(window.window, toCRect(window.ContentLocalBounds()))
}

// RepaintBounds marks the specified bounds within the window for painting at the next update.
func (window *Window) RepaintBounds(bounds draw.Rect) {
	bounds.Intersect(window.ContentLocalBounds())
	if !bounds.IsEmpty() {
		C.uiRepaintWindow(window.window, toCRect(bounds))
	}
}

// FlushPainting causes any areas marked for repainting to be painted.
func (window *Window) FlushPainting() {
	C.uiFlushPainting(window.window)
}

// Focus returns the widget with the keyboard focus in this window.
func (window *Window) Focus() Widget {
	if window.focus == nil {
		window.FocusNext()
	}
	return window.focus
}

// SetFocus sets the keyboard focus to the specified target.
func (window *Window) SetFocus(target Widget) {
	if target != nil && target.Window() == window && target != window.focus {
		if window.focus != nil {
			event := &Event{Type: FocusLostEvent, When: time.Now(), Target: window.focus}
			event.Dispatch()
		}
		window.focus = target
		if target != nil {
			event := &Event{Type: FocusGainedEvent, When: time.Now(), Target: target}
			event.Dispatch()
		}
	}
}

func collectFocusables(current Widget, target Widget, focusables []Widget) (int, []Widget) {
	match := -1
	if current.Focusable() {
		if current == target {
			match = len(focusables)
		}
		focusables = append(focusables, current)
	}
	for _, child := range current.Children() {
		var m int
		m, focusables = collectFocusables(child, target, focusables)
		if match == -1 && m != -1 {
			match = m
		}
	}
	return match, focusables
}

// FocusNext moves the keyboard focus to the next focusable widget.
func (window *Window) FocusNext() {
	current := window.focus
	if current == nil {
		current = window.rootBlock
	}
	i, focusables := collectFocusables(window.rootBlock, current, make([]Widget, 0))
	size := len(focusables)
	if size > 0 {
		i++
		if i >= size {
			i = 0
		}
		current = focusables[i]
	}
	window.SetFocus(current)
}

// FocusPrevious moves the keyboard focus to the previous focusable widget.
func (window *Window) FocusPrevious() {
	current := window.focus
	if current == nil {
		current = window.rootBlock
	}
	i, focusables := collectFocusables(window.rootBlock, current, make([]Widget, 0))
	size := len(focusables)
	if size > 0 {
		i--
		if i < 0 {
			i = size - 1
		}
		current = focusables[i]
	}
	window.SetFocus(current)
}
