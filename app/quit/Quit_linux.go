// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package quit

import (
	"fmt"
	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui/internal/iapp"
	"github.com/richardwilkes/ui/internal/iwindow"
	"github.com/richardwilkes/ui/keys"
	"math"
	"syscall"
	"time"
	"unsafe"
	// #cgo linux LDFLAGS: -lX11 -lcairo
	// #include <X11/Xlib.h>
	// #include <X11/keysym.h>
	// #include <X11/Xutil.h>
	// #include <cairo/cairo.h>
	// #include <cairo/cairo-xlib.h>
	"C"
)

var (
	quitting     bool
	awaitingQuit bool
)

func platformAttemptQuit() {
	switch appShouldQuit() {
	case Cancel:
	case Later:
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
	if window.Count > 0 {
		for _, w := range Windows() {
			w.Close()
		}
	} else {
		finishQuit()
	}
}

func finishQuit() {
	if quitting {
		iapp.Running = false
		C.XCloseDisplay(xDisplay)
		xDisplay = nil
		syscall.Exit(0)
	}
}
