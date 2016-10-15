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
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/internal/task"
	"github.com/richardwilkes/ui/internal/x11"
	"math"
	"syscall"
	"time"
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
			processKeyEvent(event.ToKeyEvent(), platformKeyDown)
		case x11.KeyReleaseType:
			processKeyEvent(event.ToKeyEvent(), platformKeyUp)
		case x11.ButtonPressType:
			processButtonPressEvent(event.ToButtonEvent())
		case x11.ButtonReleaseType:
			processButtonReleaseEvent(event.ToButtonEvent())
		case x11.MotionNotifyType:
			processMotionEvent(event.ToMotionEvent())
		case x11.EnterNotifyType:
			processCrossingEvent(event.ToCrossingEvent(), platformMouseEntered)
		case x11.LeaveNotifyType:
			processCrossingEvent(event.ToCrossingEvent(), platformMouseExited)
		case x11.FocusInType:
			processFocusInEvent(event.ToFocusChangeEvent())
		case x11.FocusOutType:
			processFocusOutEvent(event.ToFocusChangeEvent())
		case x11.ExposeType:
			processExposeEvent(event.ToExposeEvent())
		case x11.DestroyWindowType:
			processDestroyWindowEvent(event.ToDestroyWindowEvent())
		case x11.ConfigureType:
			processConfigureEvent(event.ToConfigureEvent())
		case x11.ClientMessageType:
			processClientEvent(event.ToClientMessageEvent())
		}
	}
}

func processKeyEvent(evt *x11.KeyEvent, eventType platformEventType) {
	if window, ok := windowMap[platformWindow(uintptr(evt.Window()))]; ok {
		code, ch := evt.CodeAndChar()
		window.keyEvent(eventType, evt.Modifiers(), code, ch, false)
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
			window.mouseWheelEvent(evt.Modifiers(), where.X, where.Y, dir.X, dir.Y)
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
			window.mouseEvent(platformMouseDown, evt.Modifiers(), lastMouseDownButton, clickCount, where.X, where.Y)
		}
	}
}

func processButtonReleaseEvent(evt *x11.ButtonEvent) {
	if window, ok := windowMap[platformWindow(uintptr(evt.Window()))]; ok {
		if !evt.IsScrollWheel() {
			where := evt.Where()
			lastMouseDownButton = -1
			window.mouseEvent(platformMouseUp, evt.Modifiers(), evt.Button(), clickCount, where.X, where.Y)
		}
	}
}

func processMotionEvent(evt *x11.MotionEvent) {
	var eventType platformEventType
	var button int
	var translate bool
	target := platformWindow(uintptr(evt.Window()))
	if lastMouseDownButton != -1 {
		translate = target != lastMouseDownWindow
		eventType = platformMouseDragged
		target = lastMouseDownWindow
		button = lastMouseDownButton
	} else {
		eventType = platformMouseMoved
	}
	if window, ok := windowMap[target]; ok {
		where := evt.Where()
		if translate {
			// RAW: Translate coordinates appropriately
			fmt.Println("need translation for mouse drag")
		}
		window.mouseEvent(eventType, evt.Modifiers(), button, 0, where.X, where.Y)
	}
}

func processCrossingEvent(evt *x11.CrossingEvent, eventType platformEventType) {
	if window, ok := windowMap[platformWindow(uintptr(evt.Window()))]; ok {
		where := evt.Where()
		window.mouseEvent(eventType, evt.Modifiers(), 0, 0, where.X, where.Y)
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
			if other := x11.NextEventOfTypeForWindow(x11.ExposeType, wnd); other != nil {
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
			if other := x11.NextEventOfTypeForWindow(x11.ConfigureType, wnd); other != nil {
				evt = other.ToConfigureEvent()
			} else {
				break
			}
		}
		win.ignoreRepaint = true
		size := win.ContentFrame().Size
		win.root.SetSize(size)
		win.root.ValidateLayout()
		win.ignoreRepaint = false
		(*x11.Surface)(win.surface).SetSize(size)
		win.Repaint()
	}
}

func processClientEvent(evt *x11.ClientMessageEvent) {
	switch evt.SubType() {
	case x11.ProtocolsSubType:
		if evt.Format() == 32 && evt.Protocol() == x11.DeleteWindowProtocol {
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
	if WindowCount() > 0 {
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
