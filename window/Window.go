package window

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/cursor"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/internal/task"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/layout"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/object"
	"github.com/richardwilkes/ui/widget/tooltip"
)

// StyleMask controls the look and capabilities of a window.
type StyleMask int

// Possible values for the StyleMask.
const (
	TitledWindowMask StyleMask = 1 << iota
	ClosableWindowMask
	MinimizableWindowMask
	ResizableWindowMask
	BorderlessWindowMask = 0
	StdWindowMask        = TitledWindowMask | ClosableWindowMask | MinimizableWindowMask | ResizableWindowMask
)

type commonWindow struct {
	object.Base
	window                 platformWindow
	eventHandlers          *event.Handlers
	owner                  ui.Window
	root                   *RootView
	focus                  ui.Widget
	lastMouseWidget        ui.Widget
	lastToolTip            ui.Widget
	lastTooltipShownAt     time.Time
	lastCursor             *cursor.Cursor
	style                  StyleMask
	initialLocationRequest geom.Point
	tooltipWidget          ui.Widget
	tooltipSequence        int
	inMouseDown            bool
	ignoreRepaint          bool
}

var (
	// LastWindowClosed will be called when the last window is closed, if not nil.
	LastWindowClosed func()
	windowMap        = make(map[platformWindow]*Window)
	windowIDMap      = make(map[uint64]*Window)
	windowList       = make([]*Window, 0)
)

// AllWindowsToFront attempts to bring all of the application's windows to the foreground.
func AllWindowsToFront() {
	platformBringAllWindowsToFront()
}

// Count returns the number of windows that are open.
func Count() int {
	return len(windowMap)
}

// ByID returns the window associated with the specified ID.
func ByID(id uint64) ui.Window {
	return windowIDMap[id]
}

// Windows returns a slice containing the current set of open windows.
func Windows() []ui.Window {
	list := make([]ui.Window, 0, len(windowList))
	for _, w := range windowList {
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
func NewWindow(where geom.Point, styleMask StyleMask) *Window {
	return NewWindowWithContentSize(where, geom.Size{Width: 100, Height: 100}, styleMask)
}

// NewWindowWithContentSize creates a new window at the specified location with the specified style and content size.
func NewWindowWithContentSize(where geom.Point, contentSize geom.Size, styleMask StyleMask) *Window {
	bounds := geom.Rect{Point: where, Size: contentSize}
	wnd := newWindow(platformNewWindow(bounds, styleMask), styleMask, where)
	windowList = append(windowList, wnd)
	return wnd
}

// NewPopupWindow creates a new popup window at the specified location and content size.
func NewPopupWindow(parent ui.Window, where geom.Point, contentSize geom.Size) *Window {
	bounds := geom.Rect{Point: where, Size: contentSize}
	wnd := newWindow(platformNewPopupWindow(parent, bounds), BorderlessWindowMask, where)
	wnd.owner = parent
	return wnd
}

func newWindow(window *Window, styleMask StyleMask, where geom.Point) *Window {
	window.style = styleMask
	window.InitTypeAndID(window)
	windowMap[window.window] = window
	windowIDMap[window.ID()] = window
	window.initialLocationRequest = where
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

func (window *Window) String() string {
	// Can't call window.Title() here, as the window may have been closed already
	return fmt.Sprintf("Window #%d", window.ID())
}

// Owner returns the owner of the window.
func (window *Window) Owner() ui.Window {
	return window.owner
}

func (window *Window) repaintFocus() {
	if focus := window.Focus(); focus != nil {
		focus.Repaint()
	}
}

// MayClose returns true if the window is permitted to close.
func (window *Window) MayClose() bool {
	evt := event.NewClosing(window)
	event.Dispatch(evt)
	return !evt.Aborted()
}

// AttemptClose closes the window if a Closing event permits it.
func (window *Window) AttemptClose() {
	if window.MayClose() {
		window.Close()
	}
}

// Close the window.
func (window *Window) Close() {
	window.platformClose()
}

// Dispose of the window.
func (window *Window) Dispose() {
	event.Dispatch(event.NewClosed(window))
	delete(windowIDMap, window.ID())
	delete(windowMap, window.window)
	if window.owner == nil {
		for i, wnd := range windowList {
			if wnd == window {
				copy(windowList[i:], windowList[i+1:])
				count := len(windowList) - 1
				windowList[count] = nil
				windowList = windowList[:count]
				break
			}
		}
	}
	if Count() == 0 && LastWindowClosed != nil {
		LastWindowClosed()
	}
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
	return event.GlobalTarget()
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

// ContentFrame returns the boundaries of the root widget of this window.
func (window *Window) ContentFrame() geom.Rect {
	return window.platformContentFrame()
}

// SetContentFrame sets the boundaries of the root widget of this window.
func (window *Window) SetContentFrame(bounds geom.Rect) {
	frame := window.Frame()
	cFrame := window.ContentFrame()
	bounds.X += frame.X - cFrame.X
	bounds.Y += frame.Y - cFrame.Y
	bounds.Width += frame.Width - cFrame.Width
	bounds.Height += frame.Height - cFrame.Height
	window.SetFrame(bounds)
}

// ContentLocalFrame returns the local boundaries of the root widget of this window.
func (window *Window) ContentLocalFrame() geom.Rect {
	bounds := window.ContentFrame()
	bounds.X = 0
	bounds.Y = 0
	return bounds
}

// Pack sets the window's content size to match the preferred size of the root widget.
func (window *Window) Pack() {
	_, pref, _ := ui.Sizes(window.root, layout.NoHintSize)
	bounds := window.ContentFrame()
	bounds.Size = pref
	window.SetContentFrame(bounds)
}

// MenuBar returns the menu bar for the window. On some platforms, the menu bar is a global
// entity and the same value will be returned for all windows.
func (window *Window) MenuBar() menu.Bar {
	if menu.Global() {
		return menu.AppBar(0)
	}
	return window.root.MenuBar()
}

// Content returns the content widget of the window. This is not the root widget of the window,
// which contains both the content widget and the menu bar, for platforms that hold the menu bar
// within the window.
func (window *Window) Content() ui.Widget {
	return window.root.Content()
}

// Focused returns true if the window has the current keyboard focus.
func (window *Window) Focused() bool {
	return window == KeyWindow()
}

// Focus returns the widget with the keyboard focus in this window.
func (window *Window) Focus() ui.Widget {
	if window.focus == nil {
		window.FocusNext()
	}
	return window.focus
}

// SetFocus sets the keyboard focus to the specified target.
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

// FocusNext moves the keyboard focus to the next focusable widget.
func (window *Window) FocusNext() {
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
func (window *Window) FocusPrevious() {
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
	return unsafe.Pointer(uintptr(window.window))
}

func (window *Window) updateToolTipAndCursor(widget ui.Widget, where geom.Point) {
	window.updateCursor(widget, where)
	window.updateToolTip(widget, where)
}

func (window *Window) updateToolTip(widget ui.Widget, where geom.Point) {
	avoid := geom.Rect{Point: widget.ToWindow(geom.Point{}), Size: widget.Size()}
	avoid.GrowToInteger()
	var tip ui.Widget
	if widget != nil {
		e := tooltip.NewEvent(widget, where, avoid)
		event.Dispatch(e)
		tip = e.ToolTip()
		avoid = e.Avoid()
	}
	if window.lastToolTip != tip || widget != window.tooltipWidget {
		wasShowing := window.root.tooltip != nil
		window.tooltipWidget = widget
		window.clearToolTip()
		window.lastToolTip = tip
		if tip != nil {
			ts := &tooltipSequencer{window: window, avoid: avoid, sequence: window.tooltipSequence}
			if wasShowing || time.Since(window.lastTooltipShownAt) < TooltipDismissal {
				ts.show()
			} else {
				window.InvokeAfter(ts.show, TooltipDelay)
			}
		}
	}
}

func (window *Window) clearToolTip() {
	window.tooltipSequence++
	window.root.SetTooltip(nil)
}

func (window *Window) updateCursor(target event.Target, where geom.Point) {
	if target != nil {
		if !event.SendUpdateCursor(target, where) {
			window.SetCursor(cursor.Arrow)
		}
	}
}

// SetCursor sets the window's current cursor.
func (window *Window) SetCursor(cur *cursor.Cursor) {
	if window.lastCursor != cur {
		window.platformSetCursor(cur)
		window.lastCursor = cur
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

func (window *Window) paint(gc *draw.Graphics, bounds geom.Rect) {
	window.root.ValidateLayout()
	window.root.Paint(gc, bounds)
}

func (window *Window) widgetForMouse(where geom.Point) ui.Widget {
	if window.inMouseDown {
		return window.lastMouseWidget
	}
	return window.root.WidgetAt(where)
}

func (window *Window) processMouseDown(x, y float64, button, clickCount int, keyModifiers keys.Modifiers) {
	window.clearToolTip()
	where := geom.Point{X: x, Y: y}
	widget := window.root.WidgetAt(where)
	if widget.Enabled() {
		if widget.Focusable() && widget.GrabFocusWhenClickedOn() {
			window.SetFocus(widget)
		}
		e := event.NewMouseDown(widget, where, keyModifiers, button, clickCount)
		event.Dispatch(e)
		if !e.Discarded() {
			window.inMouseDown = true
		}
	}
	window.lastMouseWidget = widget
}

func (window *Window) processMouseDragged(x, y float64, button int, keyModifiers keys.Modifiers) {
	where := geom.Point{X: x, Y: y}
	widget := window.widgetForMouse(where)
	if widget.Enabled() {
		event.Dispatch(event.NewMouseDragged(widget, where, button, keyModifiers))
	}
	window.lastMouseWidget = widget
}

func (window *Window) processMouseUp(x, y float64, button int, keyModifiers keys.Modifiers) {
	where := geom.Point{X: x, Y: y}
	widget := window.widgetForMouse(where)
	if widget.Enabled() {
		event.Dispatch(event.NewMouseUp(widget, where, button, keyModifiers))
	}
	window.updateToolTipAndCursor(window.root.WidgetAt(where), where)
	window.lastMouseWidget = widget
	if window.inMouseDown {
		window.inMouseDown = false
	}
}

func (window *Window) processMouseEntered(x, y float64, keyModifiers keys.Modifiers) {
	where := geom.Point{X: x, Y: y}
	widget := window.widgetForMouse(where)
	event.Dispatch(event.NewMouseEntered(widget, where, keyModifiers))
	window.updateToolTipAndCursor(widget, where)
	window.lastMouseWidget = widget
}

func (window *Window) processMouseMoved(x, y float64, keyModifiers keys.Modifiers) {
	var evt event.Event
	where := geom.Point{X: x, Y: y}
	widget := window.widgetForMouse(where)
	if !window.inMouseDown && widget != window.lastMouseWidget {
		if window.lastMouseWidget != nil {
			event.Dispatch(event.NewMouseExited(window.lastMouseWidget, where, keyModifiers))
		}
		evt = event.NewMouseEntered(widget, where, keyModifiers)
	} else {
		evt = event.NewMouseMoved(widget, where, keyModifiers)
	}
	event.Dispatch(evt)
	window.updateToolTipAndCursor(widget, where)
	window.lastMouseWidget = widget
}

func (window *Window) processMouseExited(x, y float64, keyModifiers keys.Modifiers) {
	where := geom.Point{X: x, Y: y}
	widget := window.widgetForMouse(where)
	event.Dispatch(event.NewMouseExited(widget, where, keyModifiers))
	c := cursor.Arrow
	if window.lastCursor != c {
		window.platformSetCursor(c)
		window.lastCursor = c
	}
	window.lastMouseWidget = widget
}

func (window *Window) processMouseWheel(x, y, dx, dy float64, keyModifiers keys.Modifiers) {
	where := geom.Point{X: x, Y: y}
	widget := window.root.WidgetAt(where)
	if widget != nil {
		event.Dispatch(event.NewMouseWheel(widget, geom.Point{X: dx, Y: dy}, where, keyModifiers))
		if window.inMouseDown {
			window.processMouseDragged(x, y, 0, keyModifiers)
		} else {
			window.processMouseMoved(x, y, keyModifiers)
		}
	}
}

func (window *Window) processKeyDown(keyCode int, ch rune, keyModifiers keys.Modifiers, repeat bool) {
	window.clearToolTip()
	ch = processDiacritics(keyCode, ch, keyModifiers)
	e := event.NewKeyDown(window.Focus(), keyCode, ch, keyModifiers, repeat)
	bar := window.MenuBar()
	if bar != nil {
		bar.ProcessKeyDown(e)
	}
	if !e.Discarded() && !e.Finished() {
		event.Dispatch(e)
		if !e.Discarded() && keyCode == keys.VirtualKeyTab && (keyModifiers&(keys.AllModifiers & ^keys.ShiftModifier)) == 0 {
			if keyModifiers.ShiftDown() {
				window.FocusPrevious()
			} else {
				window.FocusNext()
			}
		}
	}
}

func (window *Window) processKeyUp(keyCode int, keyModifiers keys.Modifiers) {
	event.Dispatch(event.NewKeyUp(window.Focus(), keyCode, keyModifiers))
}

// Invoke a task on the UI thread. The task is put into the system event queue and will be run at
// the next opportunity.
func (window *Window) Invoke(taskFunction func()) {
	window.platformInvoke(task.Record(taskFunction))
}

// InvokeAfter schedules a task to be run on the UI thread after waiting for the specified
// duration.
func (window *Window) InvokeAfter(taskFunction func(), after time.Duration) {
	window.platformInvokeAfter(task.Record(taskFunction), after)
}
