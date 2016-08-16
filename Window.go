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
	"github.com/richardwilkes/ui/cursor"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/geom"
	"github.com/richardwilkes/ui/keys"
	"unsafe"
)

// #cgo darwin LDFLAGS: -framework Cocoa -framework Quartz
// #cgo linux LDFLAGS: -lX11
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
	window          C.platformWindow
	eventHandlers   *event.Handlers
	root            Widget
	focus           Widget
	lastMouseWidget Widget
	lastToolTip     string
	lastCursor      *cursor.Cursor
	style           WindowStyleMask
	inMouseDown     bool
	inLiveResize    bool
}

var (
	windowMap = make(map[C.platformWindow]*Window)
	diacritic int
)

// AllWindowsToFront attempts to bring all of the application's windows to the foreground.
func AllWindowsToFront() {
	C.platformBringAllWindowsToFront()
}

// KeyWindow returns the window that currently has the keyboard focus, or nil if none of your
// application's windows has the keyboard focus.
func KeyWindow() *Window {
	if window, ok := windowMap[C.platformGetKeyWindow()]; ok {
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
	window := &Window{window: C.platformNewWindow(toCRect(bounds), C.int(styleMask)), style: styleMask}
	windowMap[window.window] = window
	root := NewBlock()
	root.SetBackground(color.Background)
	root.window = window
	root.bounds = window.ContentLocalBounds()
	window.root = root
	return window
}

// AttemptClose closes the window if a Closing event permits it.
func (window *Window) AttemptClose() {
	if windowShouldClose(window.window) {
		window.Close()
	}
}

// Close the window.
func (window *Window) Close() {
	C.platformCloseWindow(window.window)
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
	return &App
}

// Title returns the title of this window.
func (window *Window) Title() string {
	return C.GoString(C.platformGetWindowTitle(window.window))
}

// SetTitle sets the title of this window.
func (window *Window) SetTitle(title string) {
	cTitle := C.CString(title)
	C.platformSetWindowTitle(window.window, cTitle)
	C.free(unsafe.Pointer(cTitle))
}

// Frame returns the boundaries in display coordinates of the frame of this window (i.e. the
// area that includes both the content and its border and window controls).
func (window *Window) Frame() geom.Rect {
	return toRect(C.platformGetWindowFrame(window.window))
}

// Location returns the upper left corner of the window in display coordinates.
func (window *Window) Location() geom.Point {
	return toPoint(C.platformGetWindowPosition(window.window))
}

// SetLocation moves the upper left corner of the window to the specified point in display
// coordinates.
func (window *Window) SetLocation(pt geom.Point) {
	C.platformSetWindowPosition(window.window, C.float(pt.X), C.float(pt.Y))
}

// Size returns the size of the window, including its frame and window controls.
func (window *Window) Size() geom.Size {
	return toSize(C.platformGetWindowSize(window.window))
}

// SetSize sets the size of the window.
func (window *Window) SetSize(size geom.Size) {
	C.platformSetWindowSize(window.window, C.float(size.Width), C.float(size.Height))
}

// ContentFrame returns the boundaries of the root content widget of this window.
func (window *Window) ContentFrame() geom.Rect {
	return toRect(C.platformGetWindowContentFrame(window.window))
}

// ContentLocalBounds returns the local boundaries of the content widget of this window.
func (window *Window) ContentLocalBounds() geom.Rect {
	size := C.platformGetWindowContentSize(window.window)
	return geom.Rect{Size: geom.Size{Width: float32(size.width), Height: float32(size.height)}}
}

// ContentLocation returns the upper left corner of the content widget in display coordinates.
func (window *Window) ContentLocation() geom.Point {
	return toPoint(C.platformGetWindowContentPosition(window.window))
}

// SetContentLocation moves the window such that the upper left corner of the content widget is
// at the specified point in display coordinates.
func (window *Window) SetContentLocation(pt geom.Point) {
	C.platformSetWindowContentPosition(window.window, C.float(pt.X), C.float(pt.Y))
}

// ContentSize returns the size of the content widget.
func (window *Window) ContentSize() geom.Size {
	return toSize(C.platformGetWindowContentSize(window.window))
}

// SetContentSize sets the size of the window to fit the specified content size.
func (window *Window) SetContentSize(size geom.Size) {
	C.platformSetWindowContentSize(window.window, C.float(size.Width), C.float(size.Height))
}

// Pack sets the window's content size to match the preferred size of the root widget.
func (window *Window) Pack() {
	_, pref, _ := Sizes(window.root, NoHintSize)
	window.SetContentSize(pref)
}

// RootWidget returns the root widget of the window.
func (window *Window) RootWidget() Widget {
	return window.root
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
			event.Dispatch(event.NewFocusLost(window.focus))
		}
		window.focus = target
		if target != nil {
			event.Dispatch(event.NewFocusGained(target))
		}
	}
}

// FocusNext moves the keyboard focus to the next focusable widget.
func (window *Window) FocusNext() {
	current := window.focus
	if current == nil {
		current = window.root
	}
	i, focusables := collectFocusables(window.root, current, make([]Widget, 0))
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
		current = window.root
	}
	i, focusables := collectFocusables(window.root, current, make([]Widget, 0))
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

// ToFront attempts to bring the window to the foreground and give it the keyboard focus.
func (window *Window) ToFront() {
	C.platformBringWindowToFront(window.window)
}

// Repaint marks this window for painting at the next update.
func (window *Window) Repaint() {
	C.platformRepaintWindow(window.window, toCRect(window.ContentLocalBounds()))
}

// RepaintBounds marks the specified bounds within the window for painting at the next update.
func (window *Window) RepaintBounds(bounds geom.Rect) {
	bounds.Intersect(window.ContentLocalBounds())
	if !bounds.IsEmpty() {
		C.platformRepaintWindow(window.window, toCRect(bounds))
	}
}

// FlushPainting causes any areas marked for repainting to be painted.
func (window *Window) FlushPainting() {
	C.platformFlushPainting(window.window)
}

// InLiveResize returns true if the window is being actively resized by the user at this point
// in time. If it is, expensive painting operations should be deferred if possible to give a
// smooth resizing experience.
func (window *Window) InLiveResize() bool {
	return window.inLiveResize
}

// ScalingFactor returns the current OS scaling factor being applied to this window.
func (window *Window) ScalingFactor() float32 {
	return float32(C.platformGetWindowScalingFactor(window.window))
}

// Minimize performs the platform's minimize function on the window.
func (window *Window) Minimize() {
	C.platformMinimizeWindow(window.window)
}

// Zoom performs the platform's zoom funcion on the window.
func (window *Window) Zoom() {
	C.platformZoomWindow(window.window)
}

// PlatformPtr returns a pointer to the underlying platform-specific data.
func (window *Window) PlatformPtr() unsafe.Pointer {
	return unsafe.Pointer(window.window)
}

func (window *Window) updateToolTipAndCursor(widget Widget, where geom.Point) {
	window.updateToolTip(widget, where)
	window.updateCursor(widget, where)
}

func (window *Window) updateToolTip(widget Widget, where geom.Point) {
	tooltip := ""
	if widget != nil {
		e := event.NewToolTip(widget, where)
		event.Dispatch(e)
		tooltip = e.ToolTip()
	}
	if window.lastToolTip != tooltip {
		if tooltip != "" {
			tip := C.CString(tooltip)
			C.platformSetToolTip(window.window, tip)
			C.free(unsafe.Pointer(tip))
		} else {
			C.platformSetToolTip(window.window, nil)
		}
		window.lastToolTip = tooltip
	}
}

func (window *Window) updateCursor(widget Widget, where geom.Point) {
	c := cursor.Arrow
	if widget != nil {
		e := event.NewCursor(widget, where)
		event.Dispatch(e)
		c = e.Cursor()
	}
	if window.lastCursor != c {
		C.platformSetCursor(window.window, c.PlatformPtr())
		window.lastCursor = c
	}
}

//export drawWindow
func drawWindow(cWindow C.platformWindow, g unsafe.Pointer, bounds C.platformRect, inLiveResize bool) {
	if window, ok := windowMap[cWindow]; ok {
		window.root.ValidateLayout()
		window.inLiveResize = inLiveResize
		window.root.Paint(draw.NewGraphics(g), toRect(bounds))
		window.inLiveResize = false
	}
}

//export windowResized
func windowResized(cWindow C.platformWindow) {
	if window, ok := windowMap[cWindow]; ok {
		window.root.SetSize(window.ContentSize())
	}
}

//export handleWindowMouseEvent
func handleWindowMouseEvent(cWindow C.platformWindow, eventType, keyModifiers, button, clickCount int, x, y float32) {
	if window, ok := windowMap[cWindow]; ok {
		modifiers := event.KeyMask(keyModifiers)
		discardMouseDown := false
		where := geom.Point{X: x, Y: y}
		var widget Widget
		if window.inMouseDown {
			widget = window.lastMouseWidget
		} else {
			widget = window.root.WidgetAt(where)
			if widget == nil {
				panic("widget is nil")
			}
			if eventType == C.platformMouseMoved && widget != window.lastMouseWidget {
				if window.lastMouseWidget != nil {
					event.Dispatch(event.NewMouseExited(window.lastMouseWidget, where, modifiers))
				}
				eventType = C.platformMouseEntered
			}
		}
		switch eventType {
		case C.platformMouseDown:
			if widget.Enabled() {
				e := event.NewMouseDown(widget, where, modifiers, button, clickCount)
				event.Dispatch(e)
				discardMouseDown = e.Discarded()
			}
		case C.platformMouseDragged:
			if widget.Enabled() {
				event.Dispatch(event.NewMouseDragged(widget, where, modifiers, button))
			}
		case C.platformMouseUp:
			if widget.Enabled() {
				event.Dispatch(event.NewMouseUp(widget, where, modifiers, button))
			}
		case C.platformMouseEntered:
			event.Dispatch(event.NewMouseEntered(widget, where, modifiers))
			window.updateToolTipAndCursor(widget, where)
		case C.platformMouseMoved:
			event.Dispatch(event.NewMouseMoved(widget, where, modifiers))
			window.updateToolTipAndCursor(widget, where)
		case C.platformMouseExited:
			event.Dispatch(event.NewMouseExited(widget, where, modifiers))
			C.platformSetCursor(window.window, cursor.Arrow.PlatformPtr())
			window.lastCursor = cursor.Arrow
		default:
			panic(fmt.Sprintf("Unknown event type: %d", eventType))
		}
		window.lastMouseWidget = widget
		if eventType == C.platformMouseDown {
			if !discardMouseDown {
				window.inMouseDown = true
			}
		} else if window.inMouseDown && eventType == C.platformMouseUp {
			window.inMouseDown = false
		}
	}
}

//export handleCursorUpdateEvent
func handleCursorUpdateEvent(cWindow C.platformWindow) {
	if window, ok := windowMap[cWindow]; ok {
		C.platformSetCursor(window.window, window.lastCursor.PlatformPtr())
	}
}

// HideCursorUntilMouseMoves causes the cursor to disappear until it is moved.
func HideCursorUntilMouseMoves() {
	C.platformHideCursorUntilMouseMoves()
}

//export handleWindowMouseWheelEvent
func handleWindowMouseWheelEvent(cWindow C.platformWindow, eventType, keyModifiers int, x, y, dx, dy float32) {
	if window, ok := windowMap[cWindow]; ok {
		where := geom.Point{X: x, Y: y}
		widget := window.root.WidgetAt(where)
		if widget != nil {
			event.Dispatch(event.NewMouseWheel(widget, geom.Point{X: dx, Y: dy}, where, event.KeyMask(keyModifiers)))
			if window.inMouseDown {
				eventType = C.platformMouseDragged
			} else {
				eventType = C.platformMouseMoved
			}
			handleWindowMouseEvent(cWindow, eventType, keyModifiers, 0, 0, x, y)
		}
	}
}

//export handleWindowKeyEvent
func handleWindowKeyEvent(cWindow C.platformWindow, eventType, keyModifiers, keyCode int, chars *C.char, repeat bool) {
	if window, ok := windowMap[cWindow]; ok {
		modifiers := event.KeyMask(keyModifiers)
		var ch rune
		runes := ([]rune)(C.GoString(chars))
		if len(runes) > 0 {
			ch = runes[0]
		} else {
			ch = 0
		}
		switch eventType {
		case C.platformKeyDown:
			if diacritic != 0 {
				if modifiers&^event.ShiftKeyMask == 0 {
					switch ch {
					case 'a':
						switch diacritic {
						case keys.E:
							ch = 'á'
						case keys.I:
							ch = 'â'
						case keys.Backtick:
							ch = 'à'
						case keys.N:
							ch = 'ã'
						case keys.U:
							ch = 'ä'
						}
					case 'A':
						switch diacritic {
						case keys.E:
							ch = 'Á'
						case keys.I:
							ch = 'Â'
						case keys.Backtick:
							ch = 'À'
						case keys.N:
							ch = 'Ã'
						case keys.U:
							ch = 'Ä'
						}
					case 'e':
						switch diacritic {
						case keys.E:
							ch = 'é'
						case keys.I:
							ch = 'ê'
						case keys.Backtick:
							ch = 'è'
						case keys.U:
							ch = 'ë'
						}
					case 'E':
						switch diacritic {
						case keys.E:
							ch = 'É'
						case keys.I:
							ch = 'Ê'
						case keys.Backtick:
							ch = 'È'
						case keys.U:
							ch = 'Ë'
						}
					case 'i':
						switch diacritic {
						case keys.E:
							ch = 'í'
						case keys.I:
							ch = 'î'
						case keys.Backtick:
							ch = 'ì'
						case keys.U:
							ch = 'ï'
						}
					case 'I':
						switch diacritic {
						case keys.E:
							ch = 'Í'
						case keys.I:
							ch = 'Î'
						case keys.Backtick:
							ch = 'Ì'
						case keys.U:
							ch = 'Ï'
						}
					case 'o':
						switch diacritic {
						case keys.E:
							ch = 'ó'
						case keys.I:
							ch = 'ô'
						case keys.Backtick:
							ch = 'ò'
						case keys.N:
							ch = 'õ'
						case keys.U:
							ch = 'ö'
						}
					case 'O':
						switch diacritic {
						case keys.E:
							ch = 'Ó'
						case keys.I:
							ch = 'Ô'
						case keys.Backtick:
							ch = 'Ò'
						case keys.N:
							ch = 'Õ'
						case keys.U:
							ch = 'Ö'
						}
					case 'u':
						switch diacritic {
						case keys.E:
							ch = 'ú'
						case keys.I:
							ch = 'û'
						case keys.Backtick:
							ch = 'ù'
						case keys.U:
							ch = 'ü'
						}
					case 'U':
						switch diacritic {
						case keys.E:
							ch = 'Ú'
						case keys.I:
							ch = 'Û'
						case keys.Backtick:
							ch = 'Ù'
						case keys.U:
							ch = 'Ü'
						}
					}
				}
				diacritic = 0
			}
			if modifiers&^event.ShiftKeyMask == event.OptionKeyMask {
				switch keyCode {
				case keys.E, keys.I, keys.Backtick, keys.N, keys.U:
					diacritic = keyCode
				default:
					diacritic = 0
				}
			}
			e := event.NewKeyDown(window.Focus(), keyCode, ch, repeat, modifiers)
			event.Dispatch(e)
			if !e.Discarded() && keyCode == keys.Tab && (modifiers&(event.AllKeyMask & ^event.ShiftKeyMask)) == 0 {
				if modifiers.ShiftDown() {
					window.FocusPrevious()
				} else {
					window.FocusNext()
				}
			}
		case C.platformKeyUp:
			event.Dispatch(event.NewKeyUp(window.Focus(), keyCode, ch, modifiers))
		default:
			panic(fmt.Sprintf("Unknown event type: %d", eventType))
		}
	}
}

//export windowShouldClose
func windowShouldClose(cWindow C.platformWindow) bool {
	if window, ok := windowMap[cWindow]; ok {
		e := event.NewClosing(window)
		event.Dispatch(e)
		return !e.Aborted()
	}
	return true
}

//export windowDidClose
func windowDidClose(cWindow C.platformWindow) {
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

func toRect(bounds C.platformRect) geom.Rect {
	return geom.Rect{Point: geom.Point{X: float32(bounds.x), Y: float32(bounds.y)}, Size: geom.Size{Width: float32(bounds.width), Height: float32(bounds.height)}}
}

func toCRect(bounds geom.Rect) C.platformRect {
	return C.platformRect{x: C.float(bounds.X), y: C.float(bounds.Y), width: C.float(bounds.Width), height: C.float(bounds.Height)}
}

func toPoint(pt C.platformPoint) geom.Point {
	return geom.Point{X: float32(pt.x), Y: float32(pt.y)}
}

func toSize(size C.platformSize) geom.Size {
	return geom.Size{Width: float32(size.width), Height: float32(size.height)}
}
