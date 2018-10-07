package window

import (
	"math"
	"syscall"
	"time"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/internal/task"
	"github.com/richardwilkes/ui/internal/x11"
)

var (
	// DoubleClickTime holds the maximum amount of time that can elapse between two clicks for them
	// to be considered part of a multi-click event.
	DoubleClickTime time.Duration = time.Millisecond * 250
	// DoubleClickDistance holds the maximum distance subsequent clicks can be from the last click
	// when determining if a click is part of a multi-click event.
	DoubleClickDistance float64 = 5
)

var (
	quitting            bool
	awaitingQuit        bool
	clickCount          int
	lastClick           time.Time
	lastClickSpot       geom.Point
	lastClickButton     int = -1
	lastMouseDownWindow platformWindow
	lastMouseDownButton int = -1
)

func RunEventLoop() {
	for x11.Running() {
		event := x11.NextEvent()
		switch event.Type() {
		case x11.KeyPressType:
			processKeyDownEvent(event.ToKeyEvent())
		case x11.KeyReleaseType:
			processKeyUpEvent(event.ToKeyEvent())
		case x11.ButtonPressType:
			processButtonPressEvent(event.ToButtonEvent())
		case x11.ButtonReleaseType:
			processButtonReleaseEvent(event.ToButtonEvent())
		case x11.MotionNotifyType:
			processMotionEvent(event.ToMotionEvent())
		case x11.EnterNotifyType:
			processMouseEnteredEvent(event.ToCrossingEvent())
		case x11.LeaveNotifyType:
			processMouseExitedEvent(event.ToCrossingEvent())
		case x11.FocusInType:
			processFocusInEvent(event.ToFocusChangeEvent())
		case x11.FocusOutType:
			processFocusOutEvent(event.ToFocusChangeEvent())
		case x11.ExposeType:
			processExposeEvent(event.ToExposeEvent())
		case x11.DestroyNotifyType:
			processDestroyWindowEvent(event.ToDestroyWindowEvent())
		case x11.ConfigureNotifyType:
			processConfigureEvent(event.ToConfigureEvent())
		case x11.ClientMessageType:
			processClientEvent(event.ToClientMessageEvent())
		case x11.SelectionClearType:
			x11.ProcessSelectionClearEvent(event.ToSelectionClearEvent())
		case x11.SelectionRequestType:
			x11.ProcessSelectionRequestEvent(event.ToSelectionRequestEvent())
		}
	}
}

func processKeyDownEvent(evt *x11.KeyEvent) {
	if window, ok := windowMap[platformWindow(uintptr(evt.Window()))]; ok {
		code, ch := evt.CodeAndChar()
		window.processKeyDown(code, ch, evt.Modifiers(), false)
	}
}

func processKeyUpEvent(evt *x11.KeyEvent) {
	if window, ok := windowMap[platformWindow(uintptr(evt.Window()))]; ok {
		code, _ := evt.CodeAndChar()
		window.processKeyUp(code, evt.Modifiers())
	}
}

func processButtonPressEvent(evt *x11.ButtonEvent) {
	wnd := platformWindow(uintptr(evt.Window()))
	keyWindow := platformGetKeyWindow()
	if keyWindow != wnd {
		focusOut(keyWindow)
	}
	if window, ok := windowMap[platformWindow(uintptr(evt.Window()))]; ok {
		where := evt.Where()
		if evt.IsScrollWheel() {
			dir := evt.ScrollWheelDirection()
			window.processMouseWheel(where.X, where.Y, dir.X, dir.Y, evt.Modifiers())
		} else {
			lastMouseDownButton = evt.Button()
			lastMouseDownWindow = wnd
			now := time.Now()
			if lastClickButton == lastMouseDownButton && now.Sub(lastClick) <= DoubleClickTime && math.Abs(lastClickSpot.X-where.X) <= DoubleClickDistance && math.Abs(lastClickSpot.Y-where.Y) <= DoubleClickDistance {
				clickCount++
			} else {
				clickCount = 1
			}
			lastClick = now
			lastClickButton = lastMouseDownButton
			lastClickSpot = where
			window.processMouseDown(where.X, where.Y, lastMouseDownButton, clickCount, evt.Modifiers())
		}
	}
}

func processButtonReleaseEvent(evt *x11.ButtonEvent) {
	if window, ok := windowMap[platformWindow(uintptr(evt.Window()))]; ok {
		if !evt.IsScrollWheel() {
			where := evt.Where()
			lastMouseDownButton = -1
			window.processMouseUp(where.X, where.Y, evt.Button(), evt.Modifiers())
		}
	}
}

func processMouseEnteredEvent(evt *x11.CrossingEvent) {
	if window, ok := windowMap[platformWindow(uintptr(evt.Window()))]; ok {
		where := evt.Where()
		window.processMouseEntered(where.X, where.Y, evt.Modifiers())
	}
}

func processMotionEvent(evt *x11.MotionEvent) {
	target := platformWindow(uintptr(evt.Window()))
	if lastMouseDownButton != -1 {
		if window, ok := windowMap[lastMouseDownWindow]; ok {
			where := evt.Where()
			if target != lastMouseDownWindow {
				if other, ok := windowMap[target]; ok {
					// Translate the coordinates to the window that had the mouse down
					bounds := other.ContentFrame()
					bounds.Point.Subtract(window.ContentFrame().Point)
					where.Add(bounds.Point)
				}
			}
			window.processMouseDragged(where.X, where.Y, lastMouseDownButton, evt.Modifiers())
		}
	} else {
		if window, ok := windowMap[target]; ok {
			where := evt.Where()
			window.processMouseMoved(where.X, where.Y, evt.Modifiers())
		}
	}
}

func processMouseExitedEvent(evt *x11.CrossingEvent) {
	if window, ok := windowMap[platformWindow(uintptr(evt.Window()))]; ok {
		where := evt.Where()
		window.processMouseExited(where.X, where.Y, evt.Modifiers())
	}
}

func processFocusInEvent(evt *x11.FocusChangeEvent) {
	event.SendAppWillActivate()
	event.SendAppDidActivate()
	if window, ok := windowMap[platformWindow(uintptr(evt.Window()))]; ok {
		event.Dispatch(event.NewFocusGained(window))
	}
}

func processFocusOutEvent(evt *x11.FocusChangeEvent) {
	focusOut(platformWindow(uintptr(evt.Window())))
}

func focusOut(wnd platformWindow) {
	if window, ok := windowMap[wnd]; ok {
		event.Dispatch(event.NewFocusLost(window))
	}
	event.SendAppWillDeactivate()
	event.SendAppDidDeactivate()
}

func processExposeEvent(evt *x11.ExposeEvent) {
	wnd := evt.Window()
	if win, ok := windowMap[platformWindow(uintptr(wnd))]; ok {
		bounds := evt.Bounds()
		// Collect up any other expose events for this window that are already in the queue and union their exposed area into the area we need to redraw
		for {
			if other := wnd.NextEventOfType(x11.ExposeType); other != nil {
				bounds.Union(other.ToExposeEvent().Bounds())
			} else {
				break
			}
		}
		win.draw(bounds)
	}
}

func processDestroyWindowEvent(evt *x11.DestroyWindowEvent) {
	if window, ok := windowMap[platformWindow(uintptr(evt.Window()))]; ok {
		window.Dispose()
	}
}

func processConfigureEvent(evt *x11.ConfigureEvent) {
	wnd := evt.Window()
	pwnd := platformWindow(uintptr(wnd))
	if win, ok := windowMap[pwnd]; ok {
		// Collect up the last resize event for this window that is already in the queue and use that one instead
		for {
			if other := wnd.NextEventOfType(x11.ConfigureNotifyType); other != nil {
				evt = other.ToConfigureEvent()
			} else {
				break
			}
		}
		win.ignoreRepaint = true
		size := win.ContentFrame().Size
		win.root.SetSize(size)
		win.ignoreRepaint = false
		win.surface.SetSize(size)
	}
}

func processClientEvent(evt *x11.ClientMessageEvent) {
	switch evt.SubType() {
	case x11.ProtocolsSubType:
		if evt.Format() == 32 && evt.Protocol() == x11.DeleteWindowSubType {
			wnd := platformWindow(uintptr(evt.Window()))
			if win, ok := windowMap[wnd]; ok {
				if win.MayClose() {
					win.Close()
				}
			}
		}
	case x11.TaskSubType:
		if evt.Format() == 32 {
			task.Dispatch(evt.TaskID())
		}
	}
}

func DeferQuit() {
	awaitingQuit = true
}

func StartQuit() {
	event.SendAppWillQuit()
	quitting = true
	if Count() > 0 {
		for _, w := range Windows() {
			w.Close()
		}
	}
	finishQuit()
}

func ResumeQuit(quit bool) {
	if awaitingQuit {
		awaitingQuit = false
		if quit {
			StartQuit()
		}
	}
}

func finishQuit() {
	if quitting {
		x11.CloseDisplay()
		syscall.Exit(0)
	}
}
