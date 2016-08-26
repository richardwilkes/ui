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
	"github.com/richardwilkes/ui/event"
	"syscall"
	"unsafe"
)

// #cgo linux LDFLAGS: -lX11 -lcairo
// #include <X11/Xlib.h>
// #include <cairo/cairo.h>
// #include <cairo/cairo-xlib.h>
import "C"

var (
	running             bool
	awaitingQuit        bool
	xWindowCount        int
	xDisplay            *C.Display
	wmProtocolsAtom     C.Atom
	wmDeleteAtom        C.Atom
	goTaskAtom          C.Atom
	quitting            bool
	lastMouseDownWindow platformWindow
	lastMouseDownButton = -1
)

func platformStartUserInterface() {
	if xDisplay = C.XOpenDisplay(nil); xDisplay == nil {
		panic("Failed to open the X11 display")
	}
	wmProtocolsAtom = C.XInternAtom(xDisplay, C.CString("WM_PROTOCOLS"), C.False)
	wmDeleteAtom = C.XInternAtom(xDisplay, C.CString("WM_DELETE_WINDOW"), C.False)
	goTaskAtom = C.XInternAtom(xDisplay, C.CString("GoTask"), C.False)
	event.PlatformSetXDisplay(unsafe.Pointer(xDisplay), uint32(goTaskAtom))
	running = true
	appWillFinishStartup()
	appDidFinishStartup()
	if xWindowCount == 0 && appShouldQuitAfterLastWindowClosed() {
		platformAttemptQuit()
	}
	for running {
		var event C.XEvent
		C.XNextEvent(xDisplay, &event)
		processOneEvent(event)
	}
}

func processOneEvent(evt C.XEvent) {
	anyEvent := (*C.XAnyEvent)(unsafe.Pointer(&evt))
	window := platformWindow(uintptr(anyEvent.window))
	switch anyEvent._type {
	case C.KeyPress:
		fmt.Println("KeyPress")
	case C.KeyRelease:
		fmt.Println("KeyRelease")
	case C.ButtonPress:
		buttonEvent := (*C.XButtonEvent)(unsafe.Pointer(&evt))
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
			// RAW: Needs concept of click count
			handleWindowMouseEvent(window, platformMouseDown, convertKeyMask(buttonEvent.state), lastMouseDownButton, 0, float64(buttonEvent.x), float64(buttonEvent.y))
		}
	case C.ButtonRelease:
		buttonEvent := (*C.XButtonEvent)(unsafe.Pointer(&evt))
		if !isScrollWheelButton(buttonEvent.button) {
			lastMouseDownButton = -1
			// RAW: Needs concept of click count
			handleWindowMouseEvent(window, platformMouseUp, convertKeyMask(buttonEvent.state), getButton(buttonEvent.button), 0, float64(buttonEvent.x), float64(buttonEvent.y))
		}
	case C.MotionNotify:
		motionEvent := (*C.XMotionEvent)(unsafe.Pointer(&evt))
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
		crossingEvent := (*C.XCrossingEvent)(unsafe.Pointer(&evt))
		handleWindowMouseEvent(window, platformMouseEntered, convertKeyMask(crossingEvent.state), 0, 0, float64(crossingEvent.x), float64(crossingEvent.y))
	case C.LeaveNotify:
		crossingEvent := (*C.XCrossingEvent)(unsafe.Pointer(&evt))
		handleWindowMouseEvent(window, platformMouseExited, convertKeyMask(crossingEvent.state), 0, 0, float64(crossingEvent.x), float64(crossingEvent.y))
	case C.FocusIn:
		appWillBecomeActive()
		appDidBecomeActive()
	case C.FocusOut:
		appWillResignActive()
		appDidResignActive()
	case C.Expose:
		if win, ok := windowMap[window]; ok {
			exposeEvent := (*C.XExposeEvent)(unsafe.Pointer(&evt))
			gc := C.cairo_create(win.surface)
			C.cairo_set_line_width(gc, 1)
			C.cairo_rectangle(gc, C.double(exposeEvent.x), C.double(exposeEvent.y), C.double(exposeEvent.width), C.double(exposeEvent.height))
			C.cairo_clip(gc)
			drawWindow(window, gc, platformRect{x: C.double(exposeEvent.x), y: C.double(exposeEvent.y), width: C.double(exposeEvent.width), height: C.double(exposeEvent.height)}, false)
			C.cairo_destroy(gc)
		}
	case C.DestroyNotify:
		windowDidClose(window)
		if xWindowCount == 0 {
			if quitting {
				finishQuit()
			}
			if appShouldQuitAfterLastWindowClosed() {
				platformAttemptQuit()
			}
		}
	case C.ConfigureNotify:
		var other C.XEvent
		for C.XCheckTypedWindowEvent(xDisplay, anyEvent.window, C.ConfigureNotify, &other) != 0 {
			// Collect up the last resize event for this window that is already in the queue and use that one instead
			evt = other
		}
		if win, ok := windowMap[window]; ok {
			win.ignoreRepaint = true
			configEvent := (*C.XConfigureEvent)(unsafe.Pointer(&evt))
			lastKnownWindowBounds[window] = geom.Rect{Point: geom.Point{X: float64(configEvent.x), Y: float64(configEvent.y)}, Size: geom.Size{Width: float64(configEvent.width), Height: float64(configEvent.height)}}
			windowResized(window)
			win.root.ValidateLayout()
			win.ignoreRepaint = false
			size := win.ContentFrame().Size
			C.cairo_xlib_surface_set_size(win.surface, C.int(size.Width), C.int(size.Height))
		}
	case C.ClientMessage:
		clientEvent := (*C.XClientMessageEvent)(unsafe.Pointer(&evt))
		switch clientEvent.message_type {
		case wmProtocolsAtom:
			if clientEvent.format == 32 {
				data := (*C.Atom)(unsafe.Pointer(&clientEvent.data))
				if *data == wmDeleteAtom {
					if windowShouldClose(window) {
						if win, ok := windowMap[window]; ok {
							win.Close()
						}
					}
				}
			}
		case goTaskAtom:
			if clientEvent.format == 32 {
				data := (*uint64)(unsafe.Pointer(&clientEvent.data))
				event.DispatchInvocation(*data)
			}
		}
	}
}

func paintWindow(pWindow platformWindow, gc *C.cairo_t, x, y, width, height C.double, future bool) {
	C.cairo_save(gc)
	C.cairo_rectangle(gc, x, y, width, height)
	C.cairo_clip(gc)
	drawWindow(pWindow, gc, platformRect{x: C.double(x), y: C.double(y), width: C.double(width), height: C.double(height)}, false)
	C.cairo_restore(gc)
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
		initiateQuit()
	}
}

func platformAppMayQuitNow(quit bool) {
	if awaitingQuit {
		awaitingQuit = false
		if quit {
			initiateQuit()
		}
	}
}

func initiateQuit() {
	appWillQuit()
	quitting = true
	if xWindowCount > 0 {
		for _, w := range Windows() {
			w.Close()
		}
	} else {
		finishQuit()
	}
}

func finishQuit() {
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
