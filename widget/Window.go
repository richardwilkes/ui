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
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/app"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/geom"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/layout"
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
	eventHandlers   *event.Handlers
	root            ui.Widget
	focus           ui.Widget
	lastMouseWidget ui.Widget
	lastToolTip     string
	style           WindowStyleMask
	inMouseDown     bool
	inLiveResize    bool
}

var (
	windowMap = make(map[C.uiWindow]*Window)
)

// AllWindowsToFront attempts to bring all of the application's windows to the foreground.
func AllWindowsToFront() {
	C.uiBringAllWindowsToFront()
}

// KeyWindow returns the window that currently has the keyboard focus, or nil if none of your
// application's windows has the keyboard focus.
func KeyWindow() ui.Window {
	if window, ok := windowMap[C.uiGetKeyWindow()]; ok {
		return window
	}
	return nil
}

// NewWindow creates a new window at the specified location with the specified style.
func NewWindow(where geom.Point, styleMask WindowStyleMask) *Window {
	return NewWindowWithContentSize(where, geom.Size{Width: 100, Height: 100}, styleMask)
}

// NewWindowWithContentSize creates a new window at the specified location with the specified style and content size.
func NewWindowWithContentSize(where geom.Point, contentSize geom.Size, styleMask WindowStyleMask) *Window {
	bounds := geom.Rect{Point: where, Size: contentSize}
	window := &Window{window: C.uiNewWindow(toCRect(bounds), C.int(styleMask)), style: styleMask}
	windowMap[window.window] = window
	root := NewBlock()
	root.SetBackground(color.Background)
	root.window = window
	root.bounds = window.ContentLocalBounds()
	window.root = root
	return window
}

// EventHandlers implements the event.Target interface.
func (window *Window) EventHandlers() *event.Handlers {
	if window.eventHandlers == nil {
		window.eventHandlers = &event.Handlers{}
	}
	return window.eventHandlers
}

// ParentTarget implements the event.Target interface.
func (window *Window) ParentTarget() event.Target {
	return &app.App
}

// Title implements the ui.Window interface.
func (window *Window) Title() string {
	return C.GoString(C.uiGetWindowTitle(window.window))
}

// SetTitle implements the ui.Window interface.
func (window *Window) SetTitle(title string) {
	cTitle := C.CString(title)
	C.uiSetWindowTitle(window.window, cTitle)
	C.free(unsafe.Pointer(cTitle))
}

// Frame implements the ui.Window interface.
func (window *Window) Frame() geom.Rect {
	return toRect(C.uiGetWindowFrame(window.window))
}

// Location implements the ui.Window interface.
func (window *Window) Location() geom.Point {
	return toPoint(C.uiGetWindowPosition(window.window))
}

// SetLocation implements the ui.Window interface.
func (window *Window) SetLocation(pt geom.Point) {
	C.uiSetWindowPosition(window.window, C.float(pt.X), C.float(pt.Y))
}

// Size implements the ui.Window interface.
func (window *Window) Size() geom.Size {
	return toSize(C.uiGetWindowSize(window.window))
}

// SetSize implements the ui.Window interface.
func (window *Window) SetSize(size geom.Size) {
	C.uiSetWindowSize(window.window, C.float(size.Width), C.float(size.Height))
}

// ContentFrame implements the ui.Window interface.
func (window *Window) ContentFrame() geom.Rect {
	return toRect(C.uiGetWindowContentFrame(window.window))
}

// ContentLocalBounds implements the ui.Window interface.
func (window *Window) ContentLocalBounds() geom.Rect {
	size := C.uiGetWindowContentSize(window.window)
	return geom.Rect{Size: geom.Size{Width: float32(size.width), Height: float32(size.height)}}
}

// ContentLocation implements the ui.Window interface.
func (window *Window) ContentLocation() geom.Point {
	return toPoint(C.uiGetWindowContentPosition(window.window))
}

// SetContentLocation implements the ui.Window interface.
func (window *Window) SetContentLocation(pt geom.Point) {
	C.uiSetWindowContentPosition(window.window, C.float(pt.X), C.float(pt.Y))
}

// ContentSize implements the ui.Window interface.
func (window *Window) ContentSize() geom.Size {
	return toSize(C.uiGetWindowContentSize(window.window))
}

// SetContentSize implements the ui.Window interface.
func (window *Window) SetContentSize(size geom.Size) {
	C.uiSetWindowContentSize(window.window, C.float(size.Width), C.float(size.Height))
}

// Pack implements the ui.Window interface.
func (window *Window) Pack() {
	_, pref, _ := layout.Sizes(window.root, layout.NoHintSize)
	window.SetContentSize(pref)
}

// RootWidget implements the ui.Window interface.
func (window *Window) RootWidget() ui.Widget {
	return window.root
}

// Focus implements the ui.Window interface.
func (window *Window) Focus() ui.Widget {
	if window.focus == nil {
		window.FocusNext()
	}
	return window.focus
}

// SetFocus implements the ui.Window interface.
func (window *Window) SetFocus(target ui.Widget) {
	if target != nil && target.Window() == window && target != window.focus {
		if window.focus != nil {
			event.Dispatch(event.NewFocusLost(window.focus))
		}
		window.focus = target
		if target != nil {
			event.Dispatch(event.NewFocusGained(target))
		}
	}
}

// FocusNext implements the ui.Window interface.
func (window *Window) FocusNext() {
	current := window.focus
	if current == nil {
		current = window.root
	}
	i, focusables := collectFocusables(window.root, current, make([]ui.Widget, 0))
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

// FocusPrevious implements the ui.Window interface.
func (window *Window) FocusPrevious() {
	current := window.focus
	if current == nil {
		current = window.root
	}
	i, focusables := collectFocusables(window.root, current, make([]ui.Widget, 0))
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

func collectFocusables(current ui.Widget, target ui.Widget, focusables []ui.Widget) (int, []ui.Widget) {
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

// ToFront implements the ui.Window interface.
func (window *Window) ToFront() {
	C.uiBringWindowToFront(window.window)
}

// Repaint implements the ui.Window interface.
func (window *Window) Repaint() {
	C.uiRepaintWindow(window.window, toCRect(window.ContentLocalBounds()))
}

// RepaintBounds implements the ui.Window interface.
func (window *Window) RepaintBounds(bounds geom.Rect) {
	bounds.Intersect(window.ContentLocalBounds())
	if !bounds.IsEmpty() {
		C.uiRepaintWindow(window.window, toCRect(bounds))
	}
}

// FlushPainting implements the ui.Window interface.
func (window *Window) FlushPainting() {
	C.uiFlushPainting(window.window)
}

// InLiveResize implements the ui.Window interface.
func (window *Window) InLiveResize() bool {
	return window.inLiveResize
}

// ScalingFactor implements the ui.Window interface.
func (window *Window) ScalingFactor() float32 {
	return float32(C.uiGetWindowScalingFactor(window.window))
}

// Minimize implements the ui.Window interface.
func (window *Window) Minimize() {
	C.uiMinimizeWindow(window.window)
}

// Zoom implements the ui.Window interface.
func (window *Window) Zoom() {
	C.uiZoomWindow(window.window)
}

// PlatformPtr implements the ui.Window interface.
func (window *Window) PlatformPtr() unsafe.Pointer {
	return unsafe.Pointer(window.window)
}

func (window *Window) updateToolTip(widget ui.Widget, where geom.Point) {
	tooltip := ""
	if widget != nil {
		e := event.NewToolTip(widget, where)
		event.Dispatch(e)
		tooltip = e.ToolTip()
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

//export drawWindow
func drawWindow(cWindow C.uiWindow, g unsafe.Pointer, bounds C.uiRect, inLiveResize bool) {
	if window, ok := windowMap[cWindow]; ok {
		window.root.ValidateLayout()
		window.inLiveResize = inLiveResize
		window.root.Paint(draw.NewGraphics(g), toRect(bounds))
		window.inLiveResize = false
	}
}

//export windowResized
func windowResized(cWindow C.uiWindow) {
	if window, ok := windowMap[cWindow]; ok {
		window.root.SetSize(window.ContentSize())
	}
}

//export handleWindowMouseEvent
func handleWindowMouseEvent(cWindow C.uiWindow, eventType, keyModifiers, button, clickCount int, x, y float32) {
	if window, ok := windowMap[cWindow]; ok {
		modifiers := event.KeyMask(keyModifiers)
		discardMouseDown := false
		where := geom.Point{X: x, Y: y}
		var widget ui.Widget
		if window.inMouseDown {
			widget = window.lastMouseWidget
		} else {
			widget = window.root.WidgetAt(where)
			if widget == nil {
				panic("widget is nil")
			}
			if eventType == C.uiMouseMoved && widget != window.lastMouseWidget {
				if window.lastMouseWidget != nil {
					event.Dispatch(event.NewMouseExited(window.lastMouseWidget, where, modifiers))
				}
				eventType = C.uiMouseEntered
			}
		}
		switch eventType {
		case C.uiMouseDown:
			if widget.Enabled() {
				e := event.NewMouseDown(widget, where, modifiers, button, clickCount)
				event.Dispatch(e)
				discardMouseDown = e.Discarded()
			}
		case C.uiMouseDragged:
			if widget.Enabled() {
				event.Dispatch(event.NewMouseDragged(widget, where, modifiers, button))
			}
		case C.uiMouseUp:
			if widget.Enabled() {
				event.Dispatch(event.NewMouseUp(widget, where, modifiers, button))
			}
		case C.uiMouseEntered:
			event.Dispatch(event.NewMouseEntered(widget, where, modifiers))
			window.updateToolTip(widget, where)
		case C.uiMouseMoved:
			event.Dispatch(event.NewMouseMoved(widget, where, modifiers))
			window.updateToolTip(widget, where)
		case C.uiMouseExited:
			event.Dispatch(event.NewMouseExited(widget, where, modifiers))
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
		where := geom.Point{X: x, Y: y}
		widget := window.root.WidgetAt(where)
		if widget != nil {
			event.Dispatch(event.NewMouseWheel(widget, geom.Point{X: dx, Y: dy}, where, event.KeyMask(keyModifiers)))
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
		modifiers := event.KeyMask(keyModifiers)
		switch eventType {
		case C.uiKeyDown:
			e := event.NewKeyDown(window.Focus(), keyCode, repeat, modifiers)
			event.Dispatch(e)
			if !e.Discarded() && keyCode == keys.Tab && (modifiers&(event.AllKeyMask & ^event.ShiftKeyMask)) == 0 {
				if modifiers.ShiftDown() {
					window.FocusPrevious()
				} else {
					window.FocusNext()
				}
			}
		case C.uiKeyTyped:
			for _, r := range C.GoString(chars) {
				event.Dispatch(event.NewKeyTyped(window.Focus(), r, repeat, modifiers))
			}
		case C.uiKeyUp:
			event.Dispatch(event.NewKeyUp(window.Focus(), keyCode, modifiers))
		default:
			panic(fmt.Sprintf("Unknown event type: %d", eventType))
		}
	}
}

//export windowShouldClose
func windowShouldClose(cWindow C.uiWindow) bool {
	if window, ok := windowMap[cWindow]; ok {
		e := event.NewClosing(window)
		event.Dispatch(e)
		return !e.Aborted()
	}
	return true
}

//export windowDidClose
func windowDidClose(cWindow C.uiWindow) {
	if window, ok := windowMap[cWindow]; ok {
		event.Dispatch(event.NewClosed(window))
	}
	delete(windowMap, cWindow)
}

// Closable returns true if the window was created with the ClosableWindowMask.
func (window *Window) Closable() bool {
	return window.style&ClosableWindowMask == ClosableWindowMask
}

// Minimizable returns true if the window was created with the MiniaturizableWindowMask.
func (window *Window) Minimizable() bool {
	return window.style&MiniaturizableWindowMask == MiniaturizableWindowMask
}

// Resizable returns true if the window was created with the ResizableWindowMask.
func (window *Window) Resizable() bool {
	return window.style&ResizableWindowMask == ResizableWindowMask
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
