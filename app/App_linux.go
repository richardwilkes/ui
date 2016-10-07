// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package app

import (
	"github.com/richardwilkes/ui/app/quit"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/menu/custom"
	"github.com/richardwilkes/ui/window"
)

// #cgo linux LDFLAGS: -lX11 -lcairo
// #include <X11/Xlib.h>
// #include <X11/keysym.h>
// #include <X11/Xutil.h>
// #include <cairo/cairo.h>
// #include <cairo/cairo-xlib.h>
import "C"

func platformStartUserInterface() {
	window.InitializeDisplay()
	window.LastWindowClosed = func() {
		if quit.AppShouldQuitAfterLastWindowClosed() {
			quit.AttemptQuit()
		}
	}
	custom.Install()
	event.SendAppWillFinishStartup()
	event.SendAppDidFinishStartup()
	if window.WindowCount() == 0 && quit.AppShouldQuitAfterLastWindowClosed() {
		quit.AttemptQuit()
	}
	window.RunEventLoop()
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
