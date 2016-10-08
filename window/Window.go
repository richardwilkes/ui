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
	"fmt"
	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/cursor"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/id"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/layout"
	"github.com/richardwilkes/ui/menu"
	"sync"
	"time"
	"unsafe"
)

// Wnd represents a window on the display.
type Wnd struct {
	id              int64
	window          platformWindow
	surface         platformSurface // Currently only used by Linux
	eventHandlers   *event.Handlers
	owner           ui.Window
	root            *RootView
	focus           ui.Widget
	lastMouseWidget ui.Widget
	lastToolTip     string
	lastCursor      *cursor.Cursor
	style           WindowStyleMask
	inMouseDown     bool
	inLiveResize    bool
	ignoreRepaint   bool // Currently only used by Linux
	wasMapped       bool // Currently only used by Linux
}

var (
	windowMap        = make(map[platformWindow]*Wnd)
	windowIDMap      = make(map[int64]*Wnd)
	diacritic        int
	nextInvocationID uint64 = 1
	dispatchMapLock  sync.Mutex
	dispatchMap      = make(map[uint64]func())
)

// AllWindowsToFront attempts to bring all of the application's windows to the foreground.
func AllWindowsToFront() {
	platformBringAllWindowsToFront()
}

func WindowCount() int {
	return len(windowMap)
}

func ByID(id int64) ui.Window {
	return windowIDMap[id]
}

// Windows returns a slice containing the current set of open windows.
func Windows() []ui.Window {
	list := make([]ui.Window, 0, len(windowMap))
	for _, w := range windowMap {
		list = append(list, w)
	}
	return list
}

// KeyWindow returns the window that currently has the keyboard focus, or nil if none of your
// application's windows has the keyboard focus.
func KeyWindow() ui.Window {
	if window, ok := windowMap[platformGetKeyWindow()]; ok {
		if window.owner != nil {
			return window.owner
		}
		return window
	}
	return nil
}

// NewWindow creates a new window at the specified location with the specified style.
func NewWindow(where geom.Point, styleMask WindowStyleMask) *Wnd {
	return NewWindowWithContentSize(where, geom.Size{Width: 100, Height: 100}, styleMask)
}

// NewWindowWithContentSize creates a new window at the specified location with the specified style and content size.
func NewWindowWithContentSize(where geom.Point, contentSize geom.Size, styleMask WindowStyleMask) *Wnd {
	bounds := geom.Rect{Point: where, Size: contentSize}
	win, surface := platformNewWindow(bounds, styleMask)
	return newWindow(win, styleMask, surface)
}

func NewMenuWindow(parent ui.Window, where geom.Point, contentSize geom.Size) *Wnd {
	bounds := geom.Rect{Point: where, Size: contentSize}
	win, surface := platformNewMenuWindow(parent, bounds)
	wnd := newWindow(win, BorderlessWindowMask, surface)
	wnd.owner = parent
	return wnd
}

func newWindow(win platformWindow, styleMask WindowStyleMask, surface platformSurface) *Wnd {
	window := &Wnd{window: win, surface: surface, style: styleMask}
	windowMap[window.window] = window
	windowIDMap[window.ID()] = window
	window.root = newRootView(window)
	if styleMask != BorderlessWindowMask && !menu.Global() {
		bar := menu.AppBar(window.ID())
		if bar != nil {
			window.root.SetMenuBar(bar)
			event.SendAppPopulateMenuBar(window.ID())
		}
	}
	handlers := window.EventHandlers()
	handlers.Add(event.FocusGainedType, func(evt event.Event) { window.repaintFocus() })
	handlers.Add(event.FocusLostType, func(evt event.Event) { window.repaintFocus() })
	return window
}

func (window *Wnd) String() string {
	// Can't call window.Title() here, as the window may have been closed already
	return fmt.Sprintf("Window #%d", window.ID())
}

// ID returns the unique ID for this window.
func (window *Wnd) ID() int64 {
	if window.id == 0 {
		window.id = id.NextID()
	}
	return window.id
}

func (window *Wnd) Owner() ui.Window {
	return window.owner
}

func (window *Wnd) repaintFocus() {
	if focus := window.Focus(); focus != nil {
		focus.Repaint()
	}
}

// AttemptClose closes the window if a Closing event permits it.
func (window *Wnd) AttemptClose() {
	if windowShouldClose(window.window) {
		window.Close()
	}
}

// Close the window.
func (window *Wnd) Close() {
	window.platformClose()
}

// Valid returns true if the window is still valid (i.e. has not been closed).
func (window *Wnd) Valid() bool {
	_, valid := windowMap[window.window]
	return valid
}

// EventHandlers implements the event.Target interface.
func (window *Wnd) EventHandlers() *event.Handlers {
	if window.eventHandlers == nil {
		window.eventHandlers = &event.Handlers{}
	}
	return window.eventHandlers
}

// ParentTarget implements the event.Target interface.
func (window *Wnd) ParentTarget() event.Target {
	return event.GlobalTarget()
}

// Title returns the title of this window.
func (window *Wnd) Title() string {
	return window.platformTitle()
}

// SetTitle sets the title of this window.
func (window *Wnd) SetTitle(title string) {
	window.platformSetTitle(title)
}

// Frame returns the boundaries in display coordinates of the frame of this window (i.e. the
// area that includes both the content and its border and window controls).
func (window *Wnd) Frame() geom.Rect {
	return window.platformFrame()
}

// SetFrame sets the boundaries of the frame of this window.
func (window *Wnd) SetFrame(bounds geom.Rect) {
	window.platformSetFrame(bounds)
}

// ContentFrame returns the boundaries of the root widget of this window.
func (window *Wnd) ContentFrame() geom.Rect {
	return window.platformContentFrame()
}

// SetContentFrame sets the boundaries of the root widget of this window.
func (window *Wnd) SetContentFrame(bounds geom.Rect) {
	frame := window.Frame()
	cFrame := window.ContentFrame()
	bounds.X += frame.X - cFrame.X
	bounds.Y += frame.Y - cFrame.Y
	bounds.Width += frame.Width - cFrame.Width
	bounds.Height += frame.Height - cFrame.Height
	window.SetFrame(bounds)
}

// ContentLocalFrame returns the local boundaries of the root widget of this window.
func (window *Wnd) ContentLocalFrame() geom.Rect {
	bounds := window.ContentFrame()
	bounds.X = 0
	bounds.Y = 0
	return bounds
}

// Pack sets the window's content size to match the preferred size of the root widget.
func (window *Wnd) Pack() {
	_, pref, _ := ui.Sizes(window.root, layout.NoHintSize)
	bounds := window.ContentFrame()
	bounds.Size = pref
	window.SetContentFrame(bounds)
}

// MenuBar returns the menu bar for the window. On some platforms, the menu bar is a global
// entity and the same value will be returned for all windows.
func (window *Wnd) MenuBar() menu.Bar {
	if menu.Global() {
		return menu.AppBar(0)
	}
	return window.root.MenuBar()
}

// Content returns the content widget of the window. This is not the root widget of the window,
// which contains both the content widget and the menu bar, for platforms that hold the menu bar
// within the window.
func (window *Wnd) Content() ui.Widget {
	return window.root.Content()
}

func (window *Wnd) Focused() bool {
	return window == KeyWindow()
}

// Focus returns the widget with the keyboard focus in this window.
func (window *Wnd) Focus() ui.Widget {
	if window.focus == nil {
		window.FocusNext()
	}
	return window.focus
}

// SetFocus sets the keyboard focus to the specified target.
func (window *Wnd) SetFocus(target ui.Widget) {
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
func (window *Wnd) FocusNext() {
	current := window.focus
	if current == nil {
		current = window.root.Content()
	}
	i, focusables := collectFocusables(window.root.Content(), current, make([]ui.Widget, 0))
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
func (window *Wnd) FocusPrevious() {
	current := window.focus
	if current == nil {
		current = window.root.Content()
	}
	i, focusables := collectFocusables(window.root.Content(), current, make([]ui.Widget, 0))
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

// ToFront attempts to bring the window to the foreground and give it the keyboard focus.
func (window *Wnd) ToFront() {
	window.platformToFront()
}

// Repaint marks this window for painting at the next update.
func (window *Wnd) Repaint() {
	if !window.ignoreRepaint {
		window.platformRepaint(window.ContentLocalFrame())
	}
}

// RepaintBounds marks the specified bounds within the window for painting at the next update.
func (window *Wnd) RepaintBounds(bounds geom.Rect) {
	if !window.ignoreRepaint {
		bounds.Intersect(window.ContentLocalFrame())
		if !bounds.IsEmpty() {
			window.platformRepaint(bounds)
		}
	}
}

// FlushPainting causes any areas marked for repainting to be painted.
func (window *Wnd) FlushPainting() {
	window.platformFlushPainting()
}

// InLiveResize returns true if the window is being actively resized by the user at this point
// in time. If it is, expensive painting operations should be deferred if possible to give a
// smooth resizing experience.
func (window *Wnd) InLiveResize() bool {
	return window.inLiveResize
}

// ScalingFactor returns the current OS scaling factor being applied to this window.
func (window *Wnd) ScalingFactor() float64 {
	return window.platformScalingFactor()
}

// Minimize performs the platform's minimize function on the window.
func (window *Wnd) Minimize() {
	window.platformMinimize()
}

// Zoom performs the platform's zoom funcion on the window.
func (window *Wnd) Zoom() {
	window.platformZoom()
}

// PlatformPtr returns a pointer to the underlying platform-specific data.
func (window *Wnd) PlatformPtr() unsafe.Pointer {
	return unsafe.Pointer(window.window)
}

func (window *Wnd) updateToolTipAndCursor(widget ui.Widget, where geom.Point) {
	window.updateToolTip(widget, where)
	window.updateCursor(widget, where)
}

func (window *Wnd) updateToolTip(widget ui.Widget, where geom.Point) {
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

func (window *Wnd) updateCursor(widget ui.Widget, where geom.Point) {
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
func (window *Wnd) Closable() bool {
	return window.style&ClosableWindowMask != 0
}

// Minimizable returns true if the window was created with the MiniaturizableWindowMask.
func (window *Wnd) Minimizable() bool {
	return window.style&MinimizableWindowMask != 0
}

// Resizable returns true if the window was created with the ResizableWindowMask.
func (window *Wnd) Resizable() bool {
	return window.style&ResizableWindowMask != 0
}

func (window *Wnd) paint(gc *draw.Graphics, bounds geom.Rect, inLiveResize bool) {
	window.root.ValidateLayout()
	window.inLiveResize = inLiveResize
	window.root.Paint(gc, bounds)
	window.inLiveResize = false
}

func (window *Wnd) mouseEvent(eventType platformEventType, keyModifiers keys.Modifiers, button, clickCount int, x, y float64) {
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
			if widget.Focusable() && widget.GrabFocusWhenClickedOn() {
				window.SetFocus(widget)
			}
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

func (window *Wnd) mouseWheelEvent(eventType platformEventType, keyModifiers keys.Modifiers, x, y, dx, dy float64) {
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

func (window *Wnd) cursorUpdateEvent(keyModifiers keys.Modifiers, x, y float64) {
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

func (window *Wnd) keyEvent(eventType platformEventType, keyModifiers keys.Modifiers, keyCode int, ch rune, repeat bool) {
	switch eventType {
	case platformKeyDown:
		if diacritic != 0 {
			if keyModifiers&^keys.ShiftModifier == 0 {
				switch ch {
				case 'a':
					switch diacritic {
					case keys.VK_E:
						ch = 'á'
					case keys.VK_I:
						ch = 'â'
					case keys.VK_Backtick:
						ch = 'à'
					case keys.VK_N:
						ch = 'ã'
					case keys.VK_U:
						ch = 'ä'
					}
				case 'A':
					switch diacritic {
					case keys.VK_E:
						ch = 'Á'
					case keys.VK_I:
						ch = 'Â'
					case keys.VK_Backtick:
						ch = 'À'
					case keys.VK_N:
						ch = 'Ã'
					case keys.VK_U:
						ch = 'Ä'
					}
				case 'e':
					switch diacritic {
					case keys.VK_E:
						ch = 'é'
					case keys.VK_I:
						ch = 'ê'
					case keys.VK_Backtick:
						ch = 'è'
					case keys.VK_U:
						ch = 'ë'
					}
				case 'E':
					switch diacritic {
					case keys.VK_E:
						ch = 'É'
					case keys.VK_I:
						ch = 'Ê'
					case keys.VK_Backtick:
						ch = 'È'
					case keys.VK_U:
						ch = 'Ë'
					}
				case 'i':
					switch diacritic {
					case keys.VK_E:
						ch = 'í'
					case keys.VK_I:
						ch = 'î'
					case keys.VK_Backtick:
						ch = 'ì'
					case keys.VK_U:
						ch = 'ï'
					}
				case 'I':
					switch diacritic {
					case keys.VK_E:
						ch = 'Í'
					case keys.VK_I:
						ch = 'Î'
					case keys.VK_Backtick:
						ch = 'Ì'
					case keys.VK_U:
						ch = 'Ï'
					}
				case 'o':
					switch diacritic {
					case keys.VK_E:
						ch = 'ó'
					case keys.VK_I:
						ch = 'ô'
					case keys.VK_Backtick:
						ch = 'ò'
					case keys.VK_N:
						ch = 'õ'
					case keys.VK_U:
						ch = 'ö'
					}
				case 'O':
					switch diacritic {
					case keys.VK_E:
						ch = 'Ó'
					case keys.VK_I:
						ch = 'Ô'
					case keys.VK_Backtick:
						ch = 'Ò'
					case keys.VK_N:
						ch = 'Õ'
					case keys.VK_U:
						ch = 'Ö'
					}
				case 'u':
					switch diacritic {
					case keys.VK_E:
						ch = 'ú'
					case keys.VK_I:
						ch = 'û'
					case keys.VK_Backtick:
						ch = 'ù'
					case keys.VK_U:
						ch = 'ü'
					}
				case 'U':
					switch diacritic {
					case keys.VK_E:
						ch = 'Ú'
					case keys.VK_I:
						ch = 'Û'
					case keys.VK_Backtick:
						ch = 'Ù'
					case keys.VK_U:
						ch = 'Ü'
					}
				}
			}
			diacritic = 0
		}
		if keyModifiers&^keys.ShiftModifier == keys.OptionModifier {
			switch keyCode {
			case keys.VK_E, keys.VK_I, keys.VK_Backtick, keys.VK_N, keys.VK_U:
				diacritic = keyCode
			default:
				diacritic = 0
			}
		}
		if diacritic != 0 {
			ch = 0
		}
		e := event.NewKeyDown(window.Focus(), keyCode, ch, repeat, keyModifiers)
		bar := window.MenuBar()
		if bar != nil {
			bar.ProcessKeyDown(e)
		}
		if !e.Discarded() && !e.Finished() {
			event.Dispatch(e)
			if !e.Discarded() && keyCode == keys.VK_Tab && (keyModifiers&(keys.AllModifiers & ^keys.ShiftModifier)) == 0 {
				if keyModifiers.ShiftDown() {
					window.FocusPrevious()
				} else {
					window.FocusNext()
				}
			}
		}
	case platformKeyUp:
		event.Dispatch(event.NewKeyUp(window.Focus(), keyCode, ch, keyModifiers))
	default:
		panic(fmt.Sprintf("Unknown event type: %d", eventType))
	}
}

// Invoke a task on the UI thread. The task is put into the system event queue and will be run at
// the next opportunity.
func (window *Wnd) Invoke(task func()) {
	window.platformInvoke(recordTask(task))
}

// InvokeAfter schedules a task to be run on the UI thread after waiting for the specified
// duration.
func (window *Wnd) InvokeAfter(task func(), after time.Duration) {
	window.platformInvokeAfter(recordTask(task), after)
}

func recordTask(task func()) uint64 {
	dispatchMapLock.Lock()
	defer dispatchMapLock.Unlock()
	id := nextInvocationID
	nextInvocationID++
	dispatchMap[id] = task
	return id
}

func removeTask(id uint64) func() {
	dispatchMapLock.Lock()
	defer dispatchMapLock.Unlock()
	if f, ok := dispatchMap[id]; ok {
		delete(dispatchMap, id)
		return f
	}
	return nil
}
