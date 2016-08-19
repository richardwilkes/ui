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
	"github.com/richardwilkes/ui/event"
	"syscall"
	"unsafe"
)

// #cgo linux LDFLAGS: -lX11
// #include <X11/Xlib.h>
import "C"

var (
	running         bool
	awaitingQuit    bool
	xWindowCount    int
	xDisplay        *C.Display
	wmDeleteMessage C.Atom
)

func platformStartUserInterface() {
	if xDisplay = C.XOpenDisplay(nil); xDisplay == nil {
		panic("Failed to open the X11 display")
	}
	wmDeleteMessage = C.XInternAtom(xDisplay, C.CString("WM_DELETE_WINDOW"), C.False)
	appWillFinishStartup()
	running = true
	appDidFinishStartup()
	if xWindowCount == 0 && appShouldQuitAfterLastWindowClosed() {
		platformAttemptQuit()
	}
	var lastMouseDownWindow platformWindow
	mouseDownButton := -1
	for running {
		var event C.XEvent
		C.XNextEvent(xDisplay, &event)
		anyEvent := (*C.XAnyEvent)(unsafe.Pointer(&event))
		window := platformWindow(uintptr(anyEvent.window))
		switch anyEvent._type {
		case C.KeyPress:
			fmt.Println("KeyPress")
		case C.KeyRelease:
			fmt.Println("KeyRelease")
		case C.ButtonPress:
			buttonEvent := (*C.XButtonEvent)(unsafe.Pointer(&event))
			if isScrollWheelButton(buttonEvent.button) {
				var dx, dy float32
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
				handleWindowMouseWheelEvent(window, platformMouseWheel, convertKeyMask(buttonEvent.state), float32(buttonEvent.x), float32(buttonEvent.y), dx, dy)
			} else {
				mouseDownButton = getButton(buttonEvent.button)
				lastMouseDownWindow = window
				// RAW: Needs concept of click count
				handleWindowMouseEvent(window, platformMouseDown, convertKeyMask(buttonEvent.state), mouseDownButton, 0, float32(buttonEvent.x), float32(buttonEvent.y))
			}
		case C.ButtonRelease:
			buttonEvent := (*C.XButtonEvent)(unsafe.Pointer(&event))
			if !isScrollWheelButton(buttonEvent.button) {
				mouseDownButton = -1
				// RAW: Needs concept of click count
				handleWindowMouseEvent(window, platformMouseUp, convertKeyMask(buttonEvent.state), getButton(buttonEvent.button), 0, float32(buttonEvent.x), float32(buttonEvent.y))
			}
		case C.MotionNotify:
			motionEvent := (*C.XMotionEvent)(unsafe.Pointer(&event))
			if mouseDownButton != -1 {
				if window != lastMouseDownWindow {
					// RAW: Translate coordinates appropriately
					fmt.Println("need translation for mouse drag")
				}
				handleWindowMouseEvent(lastMouseDownWindow, platformMouseDragged, convertKeyMask(motionEvent.state), mouseDownButton, 0, float32(motionEvent.x), float32(motionEvent.y))
			} else {
				handleWindowMouseEvent(window, platformMouseMoved, convertKeyMask(motionEvent.state), 0, 0, float32(motionEvent.x), float32(motionEvent.y))
			}
		case C.EnterNotify:
			crossingEvent := (*C.XCrossingEvent)(unsafe.Pointer(&event))
			handleWindowMouseEvent(window, platformMouseEntered, convertKeyMask(crossingEvent.state), 0, 0, float32(crossingEvent.x), float32(crossingEvent.y))
		case C.LeaveNotify:
			crossingEvent := (*C.XCrossingEvent)(unsafe.Pointer(&event))
			handleWindowMouseEvent(window, platformMouseExited, convertKeyMask(crossingEvent.state), 0, 0, float32(crossingEvent.x), float32(crossingEvent.y))
		case C.FocusIn:
			appWillBecomeActive()
			appDidBecomeActive()
		case C.FocusOut:
			appWillResignActive()
			appDidResignActive()
		case C.Expose:
			exposeEvent := (*C.XExposeEvent)(unsafe.Pointer(&event))
			var values C.XGCValues
			gc := C.XCreateGC(xDisplay, C.Drawable(uintptr(window)), 0, &values)
			drawWindow(window, unsafe.Pointer(gc), platformRect{x: C.float(exposeEvent.x), y: C.float(exposeEvent.y), width: C.float(exposeEvent.width), height: C.float(exposeEvent.height)}, false)
			C.XFreeGC(xDisplay, gc)
		case C.DestroyNotify:
			windowDidClose(window)
			if xWindowCount == 0 && appShouldQuitAfterLastWindowClosed() {
				platformAttemptQuit()
			}
		case C.ConfigureNotify:
			windowResized(window)
		case C.ClientMessage:
			clientEvent := (*C.XClientMessageEvent)(unsafe.Pointer(&event))
			data := (*C.Atom)(unsafe.Pointer(&clientEvent.data))
			if *data == wmDeleteMessage {
				if windowShouldClose(window) {
					// RAW: Fix! Broke when I reorganized the Window code. For now, the next two lines do the same thing.
					xWindowCount--
					C.XDestroyWindow(xDisplay, clientEvent.window)
					//platformCloseWindow(window)
				}
			} else {
				// fmt.Println("Unhandled X11 ClientMessage")
			}
		default:
			// fmt.Printf("Unhandled event (type %d)\n", anyEvent._type)
		}
	}
}

func platformAppName() string {
	// RAW: Implement platformAppName for Linux
	return "<unknown>"
}

func platformHideApp() {
	// RAW: Implement for Linux
}

func platformHideOtherApps() {
	// RAW: Implement for Linux
}

func platformShowAllApps() {
	// RAW: Implement for Linux
}

func platformAttemptQuit() {
	switch appShouldQuit() {
	case QuitCancel:
	case QuitLater:
		awaitingQuit = true
	default:
		quitNow()
	}
}

func platformAppMayQuitNow(quit bool) {
	if awaitingQuit {
		awaitingQuit = false
		if quit {
			quitNow()
		}
	}
}

func quitNow() {
	appWillQuit()
	// RAW: Go through and tell each open window to close, ignoring any that refuse to.
	running = false
	C.XCloseDisplay(xDisplay)
	xDisplay = nil
	syscall.Exit(0)
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
	var keyMask event.KeyMask
	if state&C.LockMask == C.LockMask {
		keyMask |= event.CapsLockKeyMask
	}
	if state&C.ShiftMask == C.ShiftMask {
		keyMask |= event.ShiftKeyMask
	}
	if state&C.ControlMask == C.ControlMask {
		keyMask |= event.ControlKeyMask
	}
	if state&C.Mod1Mask == C.Mod1Mask {
		keyMask |= event.OptionKeyMask
	}
	if state&C.Mod4Mask == C.Mod4Mask {
		keyMask |= event.CommandKeyMask
	}
	return int(keyMask)
}
