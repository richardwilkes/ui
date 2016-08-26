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
	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/cursor"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"unsafe"
)

// Window represents a window on the display.
type Window struct {
	window          platformWindow
	surface         platformSurface // Currently only used by Linux
	eventHandlers   *event.Handlers
	root            Widget
	focus           Widget
	lastMouseWidget Widget
	lastToolTip     string
	lastCursor      *cursor.Cursor
	style           WindowStyleMask
	inMouseDown     bool
	inLiveResize    bool
	ignoreRepaint   bool // Currently only used by Linux
	wasMapped       bool // Currently only used by Linux
}

var (
	windowMap = make(map[platformWindow]*Window)
	diacritic int
)

// AllWindowsToFront attempts to bring all of the application's windows to the foreground.
func AllWindowsToFront() {
	platformBringAllWindowsToFront()
}

// Windows returns a slice containing the current set of open windows.
func Windows() []*Window {
	list := make([]*Window, 0, len(windowMap))
	for _, w := range windowMap {
		list = append(list, w)
	}
	return list
}

// KeyWindow returns the window that currently has the keyboard focus, or nil if none of your
// application's windows has the keyboard focus.
func KeyWindow() *Window {
	if window, ok := windowMap[platformGetKeyWindow()]; ok {
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
	win, surface := platformNewWindow(bounds, styleMask)
	window := &Window{window: win, surface: surface, style: styleMask}
	windowMap[window.window] = window
	root := NewBlock()
	root.SetBackground(color.Background)
	root.window = window
	root.bounds = window.ContentLocalFrame()
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
	window.platformClose()
}

// Valid returns true if the window is still valid (i.e. has not been closed).
func (window *Window) Valid() bool {
	_, valid := windowMap[window.window]
	return valid
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
	return window.platformTitle()
}

// SetTitle sets the title of this window.
func (window *Window) SetTitle(title string) {
	window.platformSetTitle(title)
}

// Frame returns the boundaries in display coordinates of the frame of this window (i.e. the
// area that includes both the content and its border and window controls).
func (window *Window) Frame() geom.Rect {
	return window.platformFrame()
}

// SetFrame sets the boundaries of the frame of this window.
func (window *Window) SetFrame(bounds geom.Rect) {
	window.platformSetFrame(bounds)
}

// ContentFrame returns the boundaries of the root content widget of this window.
func (window *Window) ContentFrame() geom.Rect {
	return window.platformContentFrame()
}

// SetContentFrame sets the boundaries of the root content widget of this window.
func (window *Window) SetContentFrame(bounds geom.Rect) {
	frame := window.Frame()
	cFrame := window.ContentFrame()
	bounds.X += frame.X - cFrame.X
	bounds.Y += frame.Y - cFrame.Y
	bounds.Width += frame.Width - cFrame.Width
	bounds.Height += frame.Height - cFrame.Height
	window.SetFrame(bounds)
}

// ContentLocalFrame returns the local boundaries of the content widget of this window.
func (window *Window) ContentLocalFrame() geom.Rect {
	bounds := window.ContentFrame()
	bounds.X = 0
	bounds.Y = 0
	return bounds
}

// Pack sets the window's content size to match the preferred size of the root widget.
func (window *Window) Pack() {
	_, pref, _ := Sizes(window.root, NoHintSize)
	bounds := window.ContentFrame()
	bounds.Size = pref
	window.SetContentFrame(bounds)
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
	window.platformToFront()
}

// Repaint marks this window for painting at the next update.
func (window *Window) Repaint() {
	if !window.ignoreRepaint {
		window.platformRepaint(window.ContentLocalFrame())
	}
}

// RepaintBounds marks the specified bounds within the window for painting at the next update.
func (window *Window) RepaintBounds(bounds geom.Rect) {
	if !window.ignoreRepaint {
		bounds.Intersect(window.ContentLocalFrame())
		if !bounds.IsEmpty() {
			window.platformRepaint(bounds)
		}
	}
}

// FlushPainting causes any areas marked for repainting to be painted.
func (window *Window) FlushPainting() {
	window.platformFlushPainting()
}

// InLiveResize returns true if the window is being actively resized by the user at this point
// in time. If it is, expensive painting operations should be deferred if possible to give a
// smooth resizing experience.
func (window *Window) InLiveResize() bool {
	return window.inLiveResize
}

// ScalingFactor returns the current OS scaling factor being applied to this window.
func (window *Window) ScalingFactor() float64 {
	return window.platformScalingFactor()
}

// Minimize performs the platform's minimize function on the window.
func (window *Window) Minimize() {
	window.platformMinimize()
}

// Zoom performs the platform's zoom funcion on the window.
func (window *Window) Zoom() {
	window.platformZoom()
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
		window.platformSetToolTip(tooltip)
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
		window.platformSetCursor(c)
		window.lastCursor = c
	}
}

// HideCursorUntilMouseMoves causes the cursor to disappear until it is moved.
func HideCursorUntilMouseMoves() {
	platformHideCursorUntilMouseMoves()
}

// Closable returns true if the window was created with the ClosableWindowMask.
func (window *Window) Closable() bool {
	return window.style&ClosableWindowMask != 0
}

// Minimizable returns true if the window was created with the MiniaturizableWindowMask.
func (window *Window) Minimizable() bool {
	return window.style&MinimizableWindowMask != 0
}

// Resizable returns true if the window was created with the ResizableWindowMask.
func (window *Window) Resizable() bool {
	return window.style&ResizableWindowMask != 0
}

func (window *Window) paint(gc *draw.Graphics, bounds geom.Rect, inLiveResize bool) {
	window.root.ValidateLayout()
	window.inLiveResize = inLiveResize
	window.root.Paint(gc, bounds)
	window.inLiveResize = false
}

func (window *Window) mouseEvent(eventType platformEventType, keyModifiers event.KeyMask, button, clickCount int, x, y float64) {
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
		if eventType == platformMouseMoved && widget != window.lastMouseWidget {
			if window.lastMouseWidget != nil {
				event.Dispatch(event.NewMouseExited(window.lastMouseWidget, where, keyModifiers))
			}
			eventType = platformMouseEntered
		}
	}
	switch eventType {
	case platformMouseDown:
		if widget.Enabled() {
			e := event.NewMouseDown(widget, where, keyModifiers, button, clickCount)
			event.Dispatch(e)
			discardMouseDown = e.Discarded()
		}
	case platformMouseDragged:
		if widget.Enabled() {
			event.Dispatch(event.NewMouseDragged(widget, where, keyModifiers, button))
		}
	case platformMouseUp:
		if widget.Enabled() {
			event.Dispatch(event.NewMouseUp(widget, where, keyModifiers, button))
		}
		window.updateToolTipAndCursor(window.root.WidgetAt(where), where)
	case platformMouseEntered:
		event.Dispatch(event.NewMouseEntered(widget, where, keyModifiers))
		window.updateToolTipAndCursor(widget, where)
	case platformMouseMoved:
		event.Dispatch(event.NewMouseMoved(widget, where, keyModifiers))
		window.updateToolTipAndCursor(widget, where)
	case platformMouseExited:
		event.Dispatch(event.NewMouseExited(widget, where, keyModifiers))
		c := cursor.Arrow
		if window.lastCursor != c {
			window.platformSetCursor(c)
			window.lastCursor = c
		}
	default:
		panic(fmt.Sprintf("Unknown event type: %d", eventType))
	}
	window.lastMouseWidget = widget
	if eventType == platformMouseDown {
		if !discardMouseDown {
			window.inMouseDown = true
		}
	} else if window.inMouseDown && eventType == platformMouseUp {
		window.inMouseDown = false
	}
}

func (window *Window) mouseWheelEvent(eventType platformEventType, keyModifiers event.KeyMask, x, y, dx, dy float64) {
	where := geom.Point{X: x, Y: y}
	widget := window.root.WidgetAt(where)
	if widget != nil {
		event.Dispatch(event.NewMouseWheel(widget, geom.Point{X: dx, Y: dy}, where, keyModifiers))
		if window.inMouseDown {
			eventType = platformMouseDragged
		} else {
			eventType = platformMouseMoved
		}
		window.mouseEvent(eventType, keyModifiers, 0, 0, x, y)
	}
}

func (window *Window) cursorUpdateEvent(keyModifiers event.KeyMask, x, y float64) {
	where := geom.Point{X: x, Y: y}
	var widget Widget
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

func (window *Window) keyEvent(eventType platformEventType, keyModifiers event.KeyMask, keyCode int, ch rune, repeat bool) {
	switch eventType {
	case platformKeyDown:
		if diacritic != 0 {
			if keyModifiers&^event.ShiftKeyMask == 0 {
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
		if keyModifiers&^event.ShiftKeyMask == event.OptionKeyMask {
			switch keyCode {
			case keys.E, keys.I, keys.Backtick, keys.N, keys.U:
				diacritic = keyCode
			default:
				diacritic = 0
			}
		}
		e := event.NewKeyDown(window.Focus(), keyCode, ch, repeat, keyModifiers)
		event.Dispatch(e)
		if !e.Discarded() && keyCode == keys.Tab && (keyModifiers&(event.AllKeyMask & ^event.ShiftKeyMask)) == 0 {
			if keyModifiers.ShiftDown() {
				window.FocusPrevious()
			} else {
				window.FocusNext()
			}
		}
	case platformKeyUp:
		event.Dispatch(event.NewKeyUp(window.Focus(), keyCode, ch, keyModifiers))
	default:
		panic(fmt.Sprintf("Unknown event type: %d", eventType))
	}
}
