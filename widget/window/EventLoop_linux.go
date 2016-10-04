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
	"github.com/richardwilkes/ui/keys"
	"math"
	"syscall"
	"time"
	"unsafe"
	// #cgo linux LDFLAGS: -lX11 -lcairo
	// #include <stdlib.h>
	// #include <stdio.h>
	// #include <string.h>
	// #include <X11/Xlib.h>
	// #include <X11/keysym.h>
	// #include <X11/Xutil.h>
	// #include <cairo/cairo.h>
	// #include <cairo/cairo-xlib.h>
	// #include "Types.h"
	"C"
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
	running               bool
	quitting              bool
	awaitingQuit          bool
	display               *C.Display
	wmProtocolsAtom       C.Atom
	wmDeleteAtom          C.Atom
	goTaskAtom            C.Atom
	clickCount            int
	lastClick             time.Time
	lastClickSpot         geom.Point
	lastClickButton       int = -1
	lastMouseDownWindow   platformWindow
	lastMouseDownButton   int = -1
	lastKnownWindowBounds     = make(map[platformWindow]geom.Rect)
)

func InitializeDisplay() {
	C.XInitThreads()
	if display = C.XOpenDisplay(nil); display == nil {
		panic("Failed to open the X11 display")
	}
	wmProtocolsAtom = C.XInternAtom(display, C.CString("WM_PROTOCOLS"), C.False)
	wmDeleteAtom = C.XInternAtom(display, C.CString("WM_DELETE_WINDOW"), C.False)
	goTaskAtom = C.XInternAtom(display, C.CString("GoTask"), C.False)
	running = true
}

func RunEventLoop() {
	for running {
		var event C.XEvent
		C.XNextEvent(display, &event)
		processOneEvent(&event)
	}
}

func processOneEvent(evt *C.XEvent) {
	anyEvent := (*C.XAnyEvent)(unsafe.Pointer(evt))
	window := platformWindow(uintptr(anyEvent.window))
	switch anyEvent._type {
	case C.KeyPress:
		processKeyEvent(evt, window, platformKeyDown)
	case C.KeyRelease:
		processKeyEvent(evt, window, platformKeyUp)
	case C.ButtonPress:
		buttonEvent := (*C.XButtonEvent)(unsafe.Pointer(evt))
		if isScrollWheelButton(buttonEvent.button) {
			var dx, dy float64
			switch buttonEvent.button {
			case 4: // Up
				dy = -1
			case 5: // Down
				dy = 1
			case 6: // Left
				dx = -1
			case 7: // Right
				dx = 1
			}
			handleWindowMouseWheelEvent(window, platformMouseWheel, convertKeyMask(buttonEvent.state), float64(buttonEvent.x), float64(buttonEvent.y), dx, dy)
		} else {
			lastMouseDownButton = getButton(buttonEvent.button)
			lastMouseDownWindow = window
			x := float64(buttonEvent.x)
			y := float64(buttonEvent.y)
			now := time.Now()
			if lastClickButton == lastMouseDownButton && now.Sub(lastClick) <= DoubleClickTime && math.Abs(lastClickSpot.X-x) <= DoubleClickDistance && math.Abs(lastClickSpot.Y-y) <= DoubleClickDistance {
				clickCount++
			} else {
				clickCount = 1
			}
			lastClick = now
			lastClickButton = lastMouseDownButton
			lastClickSpot.X = x
			lastClickSpot.Y = y
			handleWindowMouseEvent(window, platformMouseDown, convertKeyMask(buttonEvent.state), lastMouseDownButton, clickCount, x, y)
		}
	case C.ButtonRelease:
		buttonEvent := (*C.XButtonEvent)(unsafe.Pointer(evt))
		if !isScrollWheelButton(buttonEvent.button) {
			lastMouseDownButton = -1
			handleWindowMouseEvent(window, platformMouseUp, convertKeyMask(buttonEvent.state), getButton(buttonEvent.button), clickCount, float64(buttonEvent.x), float64(buttonEvent.y))
		}
	case C.MotionNotify:
		motionEvent := (*C.XMotionEvent)(unsafe.Pointer(evt))
		if lastMouseDownButton != -1 {
			if window != lastMouseDownWindow {
				// RAW: Translate coordinates appropriately
				fmt.Println("need translation for mouse drag")
			}
			handleWindowMouseEvent(lastMouseDownWindow, platformMouseDragged, convertKeyMask(motionEvent.state), lastMouseDownButton, 0, float64(motionEvent.x), float64(motionEvent.y))
		} else {
			handleWindowMouseEvent(window, platformMouseMoved, convertKeyMask(motionEvent.state), 0, 0, float64(motionEvent.x), float64(motionEvent.y))
		}
	case C.EnterNotify:
		crossingEvent := (*C.XCrossingEvent)(unsafe.Pointer(evt))
		handleWindowMouseEvent(window, platformMouseEntered, convertKeyMask(crossingEvent.state), 0, 0, float64(crossingEvent.x), float64(crossingEvent.y))
	case C.LeaveNotify:
		crossingEvent := (*C.XCrossingEvent)(unsafe.Pointer(evt))
		handleWindowMouseEvent(window, platformMouseExited, convertKeyMask(crossingEvent.state), 0, 0, float64(crossingEvent.x), float64(crossingEvent.y))
	case C.FocusIn:
		event.SendAppWillActivate()
		event.SendAppDidActivate()
		windowGainedKey(window)
	case C.FocusOut:
		windowLostKey(window)
		event.SendAppWillDeactivate()
		event.SendAppDidDeactivate()
	case C.Expose:
		if win, ok := windowMap[window]; ok {
			exposeEvent := (*C.XExposeEvent)(unsafe.Pointer(evt))
			gc := C.cairo_create(win.surface)
			C.cairo_set_line_width(gc, 1)
			C.cairo_rectangle(gc, C.double(exposeEvent.x), C.double(exposeEvent.y), C.double(exposeEvent.width), C.double(exposeEvent.height))
			C.cairo_clip(gc)
			drawWindow(window, gc, platformRect{x: C.double(exposeEvent.x), y: C.double(exposeEvent.y), width: C.double(exposeEvent.width), height: C.double(exposeEvent.height)}, false)
			C.cairo_destroy(gc)
		}
	case C.DestroyNotify:
		windowDidClose(window)
		if WindowCount() == 0 {
			//			finishQuit()
			//			if appShouldQuitAfterLastWindowClosed() {
			//				platformAttemptQuit()
			//			}
		}
	case C.ConfigureNotify:
		var other C.XEvent
		for C.XCheckTypedWindowEvent(display, anyEvent.window, C.ConfigureNotify, &other) != 0 {
			// Collect up the last resize event for this window that is already in the queue and use that one instead
			evt = &other
		}
		if win, ok := windowMap[window]; ok {
			win.ignoreRepaint = true
			configEvent := (*C.XConfigureEvent)(unsafe.Pointer(evt))
			lastKnownWindowBounds[window] = geom.Rect{Point: geom.Point{X: float64(configEvent.x), Y: float64(configEvent.y)}, Size: geom.Size{Width: float64(configEvent.width), Height: float64(configEvent.height)}}
			windowResized(window)
			win.root.ValidateLayout()
			win.ignoreRepaint = false
			size := win.ContentFrame().Size
			C.cairo_xlib_surface_set_size(win.surface, C.int(size.Width), C.int(size.Height))
		}
	case C.ClientMessage:
		clientEvent := (*C.XClientMessageEvent)(unsafe.Pointer(evt))
		switch clientEvent.message_type {
		case wmProtocolsAtom:
			if clientEvent.format == 32 {
				data := (*C.Atom)(unsafe.Pointer(&clientEvent.data))
				if *data == wmDeleteAtom {
					if windowShouldClose(window) {
						if win, ok := windowMap[window]; ok {
							win.Close()
							windowDidClose(window)
						}
					}
				}
			}
		case goTaskAtom:
			data := (*uint64)(unsafe.Pointer(&clientEvent.data))
			dispatchTask(*data)
		}
	}
}

func processKeyEvent(evt *C.XEvent, window platformWindow, eventType platformEventType) {
	keyEvent := (*C.XKeyEvent)(unsafe.Pointer(evt))
	var buffer [5]C.char
	var keySym C.KeySym
	buffer[C.XLookupString(keyEvent, &buffer[0], C.int(len(buffer)-1), &keySym, nil)] = 0
	handleWindowKeyEvent(window, eventType, convertKeyMask(keyEvent.state), int(keySym), &buffer[0], false)
}

func paintWindow(pWindow platformWindow, gc *C.cairo_t, x, y, width, height C.double, future bool) {
	C.cairo_save(gc)
	C.cairo_rectangle(gc, x, y, width, height)
	C.cairo_clip(gc)
	drawWindow(pWindow, gc, platformRect{x: C.double(x), y: C.double(y), width: C.double(width), height: C.double(height)}, false)
	C.cairo_restore(gc)
}

func isScrollWheelButton(button C.uint) bool {
	return button > 3 && button < 8
}

func getButton(button C.uint) int {
	if button == 2 {
		return 2
	}
	if button == 3 {
		return 1
	}
	return 0
}

func convertKeyMask(state C.uint) int {
	var modifiers keys.Modifiers
	if state&C.LockMask != 0 {
		modifiers |= keys.CapsLockModifier
	}
	if state&C.ShiftMask != 0 {
		modifiers |= keys.ShiftModifier
	}
	if state&C.ControlMask != 0 {
		modifiers |= keys.ControlModifier
	}
	if state&C.Mod1Mask != 0 {
		modifiers |= keys.OptionModifier
	}
	if state&C.Mod4Mask != 0 {
		modifiers |= keys.CommandModifier
	}
	return int(modifiers)
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
	} else {
		finishQuit()
	}
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
		running = false
		C.XCloseDisplay(display)
		display = nil
		syscall.Exit(0)
	}
}