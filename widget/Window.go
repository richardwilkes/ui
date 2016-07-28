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
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
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
	eventHandlers   *event.Handlers
	closeHandler    ui.CloseHandler
	root            ui.Widget
	focus           ui.Widget
	lastMouseWidget ui.Widget
	lastToolTip     string
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
func NewWindow(where draw.Point, styleMask WindowStyleMask) *Window {
	return NewWindowWithContentSize(where, draw.Size{Width: 100, Height: 100}, styleMask)
}

// NewWindowWithContentSize creates a new window at the specified location with the specified style and content size.
func NewWindowWithContentSize(where draw.Point, contentSize draw.Size, styleMask WindowStyleMask) *Window {
	bounds := draw.Rect{Point: where, Size: contentSize}
	window := &Window{window: C.uiNewWindow(toCRect(bounds), C.int(styleMask))}
	windowMap[window.window] = window
	root := NewBlock()
	root.SetBackground(color.Background)
	root.window = window
	root.bounds = window.ContentLocalBounds()
	window.root = root
	return window
}

// EventHandlers returns the current event.Handlers.
func (window *Window) EventHandlers() *event.Handlers {
	if window.eventHandlers == nil {
		window.eventHandlers = &event.Handlers{}
	}
	return window.eventHandlers
}

func (window *Window) ParentTarget() event.Target {
	return nil
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

// CloseHandler implements the ui.Window interface.
func (window *Window) CloseHandler() ui.CloseHandler {
	return window.closeHandler
}

// SetCloseHandler implements the ui.Window interface.
func (window *Window) SetCloseHandler(handler ui.CloseHandler) {
	window.closeHandler = handler
}

// Frame implements the ui.Window interface.
func (window *Window) Frame() draw.Rect {
	return toRect(C.uiGetWindowFrame(window.window))
}

// Location implements the ui.Window interface.
func (window *Window) Location() draw.Point {
	return toPoint(C.uiGetWindowPosition(window.window))
}

// SetLocation implements the ui.Window interface.
func (window *Window) SetLocation(pt draw.Point) {
	C.uiSetWindowPosition(window.window, C.float(pt.X), C.float(pt.Y))
}

// Size implements the ui.Window interface.
func (window *Window) Size() draw.Size {
	return toSize(C.uiGetWindowSize(window.window))
}

// SetSize implements the ui.Window interface.
func (window *Window) SetSize(size draw.Size) {
	C.uiSetWindowSize(window.window, C.float(size.Width), C.float(size.Height))
}

// ContentFrame implements the ui.Window interface.
func (window *Window) ContentFrame() draw.Rect {
	return toRect(C.uiGetWindowContentFrame(window.window))
}

// ContentLocalBounds implements the ui.Window interface.
func (window *Window) ContentLocalBounds() draw.Rect {
	size := C.uiGetWindowContentSize(window.window)
	return draw.Rect{Size: draw.Size{Width: float32(size.width), Height: float32(size.height)}}
}

// ContentLocation implements the ui.Window interface.
func (window *Window) ContentLocation() draw.Point {
	return toPoint(C.uiGetWindowContentPosition(window.window))
}

// SetContentLocation implements the ui.Window interface.
func (window *Window) SetContentLocation(pt draw.Point) {
	C.uiSetWindowContentPosition(window.window, C.float(pt.X), C.float(pt.Y))
}

// ContentSize implements the ui.Window interface.
func (window *Window) ContentSize() draw.Size {
	return toSize(C.uiGetWindowContentSize(window.window))
}

// SetContentSize implements the ui.Window interface.
func (window *Window) SetContentSize(size draw.Size) {
	C.uiSetWindowContentSize(window.window, C.float(size.Width), C.float(size.Height))
}

// Pack implements the ui.Window interface.
func (window *Window) Pack() {
	_, pref, _ := ui.ComputeSizes(window.root, ui.NoLayoutHintSize)
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
			event := &event.Event{Type: event.FocusLostEvent, When: time.Now(), Target: window.focus}
			event.Dispatch()
		}
		window.focus = target
		if target != nil {
			event := &event.Event{Type: event.FocusGainedEvent, When: time.Now(), Target: target}
			event.Dispatch()
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
func (window *Window) RepaintBounds(bounds draw.Rect) {
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

func (window *Window) updateToolTip(widget ui.Widget, where draw.Point) {
	tooltip := ""
	if widget != nil {
		event := &event.Event{Type: event.ToolTipEvent, When: time.Now(), Target: widget, Where: where}
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
		keyMask := event.KeyMask(keyModifiers)
		discardMouseDown := false
		where := draw.Point{X: x, Y: y}
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
					event := &event.Event{Type: event.MouseExitedEvent, When: time.Now(), Target: window.lastMouseWidget, KeyModifiers: keyMask}
					event.Dispatch()
				}
				eventType = C.uiMouseEntered
			}
		}
		switch eventType {
		case C.uiMouseDown:
			if widget.Enabled() {
				event := &event.Event{Type: event.MouseDownEvent, When: time.Now(), Target: widget, Where: where, KeyModifiers: keyMask, Button: button, Clicks: clickCount}
				event.Dispatch()
				discardMouseDown = event.Discard
			}
		case C.uiMouseDragged:
			if widget.Enabled() {
				event := &event.Event{Type: event.MouseDraggedEvent, When: time.Now(), Target: widget, Where: where, KeyModifiers: keyMask}
				event.Dispatch()
			}
		case C.uiMouseUp:
			if widget.Enabled() {
				event := &event.Event{Type: event.MouseUpEvent, When: time.Now(), Target: widget, Where: where, KeyModifiers: keyMask}
				event.Dispatch()
			}
		case C.uiMouseEntered:
			event := &event.Event{Type: event.MouseEnteredEvent, When: time.Now(), Target: widget, Where: where, KeyModifiers: keyMask}
			event.Dispatch()
			window.updateToolTip(widget, where)
		case C.uiMouseMoved:
			event := &event.Event{Type: event.MouseMovedEvent, When: time.Now(), Target: widget, Where: where, KeyModifiers: keyMask}
			event.Dispatch()
			window.updateToolTip(widget, where)
		case C.uiMouseExited:
			event := &event.Event{Type: event.MouseExitedEvent, When: time.Now(), Target: widget, KeyModifiers: keyMask}
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
		widget := window.root.WidgetAt(where)
		if widget != nil {
			event := &event.Event{Type: event.MouseWheelEvent, When: time.Now(), Target: widget, Where: where, Delta: draw.Point{X: dx, Y: dy}, KeyModifiers: event.KeyMask(keyModifiers), CascadeUp: true}
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
		keyMask := event.KeyMask(keyModifiers)
		switch eventType {
		case C.uiKeyDown:
			evt := &event.Event{Type: event.KeyDownEvent, When: time.Now(), Target: window.Focus(), KeyModifiers: keyMask, KeyCode: keyCode, Repeat: repeat, CascadeUp: true}
			evt.Dispatch()
			if !evt.Discard && keyCode == keys.Tab && (keyMask&(event.AllKeyMask & ^event.ShiftKeyMask)) == 0 {
				if keyMask&event.ShiftKeyMask == 0 {
					window.FocusNext()
				} else {
					window.FocusPrevious()
				}
			}
		case C.uiKeyTyped:
			for _, r := range C.GoString(chars) {
				event := &event.Event{Type: event.KeyTypedEvent, When: time.Now(), Target: window.Focus(), KeyModifiers: keyMask, KeyTyped: r, Repeat: repeat, CascadeUp: true}
				event.Dispatch()
			}
		case C.uiKeyUp:
			event := &event.Event{Type: event.KeyUpEvent, When: time.Now(), Target: window.Focus(), KeyModifiers: keyMask, KeyCode: keyCode, Repeat: repeat, CascadeUp: true}
			event.Dispatch()
		default:
			panic(fmt.Sprintf("Unknown event type: %d", eventType))
		}
	}
}

//export windowShouldClose
func windowShouldClose(cWindow C.uiWindow) bool {
	if window, ok := windowMap[cWindow]; ok {
		if window.closeHandler != nil {
			return window.closeHandler.WillClose()
		}
	}
	return true
}

//export windowDidClose
func windowDidClose(cWindow C.uiWindow) {
	if window, ok := windowMap[cWindow]; ok {
		if window.closeHandler != nil {
			window.closeHandler.DidClose()
		}
	}
	delete(windowMap, cWindow)
}

func toRect(bounds C.uiRect) draw.Rect {
	return draw.Rect{Point: draw.Point{X: float32(bounds.x), Y: float32(bounds.y)}, Size: draw.Size{Width: float32(bounds.width), Height: float32(bounds.height)}}
}

func toCRect(bounds draw.Rect) C.uiRect {
	return C.uiRect{x: C.float(bounds.X), y: C.float(bounds.Y), width: C.float(bounds.Width), height: C.float(bounds.Height)}
}

func toPoint(pt C.uiPoint) draw.Point {
	return draw.Point{X: float32(pt.x), Y: float32(pt.y)}
}

func toSize(size C.uiSize) draw.Size {
	return draw.Size{Width: float32(size.width), Height: float32(size.height)}
}
