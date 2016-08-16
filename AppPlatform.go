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
	"github.com/richardwilkes/ui/event"
	"runtime"
)

// #cgo darwin LDFLAGS: -framework Cocoa
// #cgo linux LDFLAGS: -lX11
// #include "App.h"
import "C"

// StartUserInterface the ui.
func StartUserInterface() {
	runtime.LockOSThread()
	C.platformStart()
}

// AppName returns the application's name.
func AppName() string {
	return C.GoString(C.platformAppName())
}

// HideApp attempts to hide this application.
func HideApp() {
	C.platformHideApp()
}

// HideOtherApps attempts to hide other applications, leaving just this application visible.
func HideOtherApps() {
	C.platformHideOtherApps()
}

// ShowAllApps attempts to show all applications that are currently hidden.
func ShowAllApps() {
	C.platformShowAllApps()
}

//export appWillFinishStartup
func appWillFinishStartup() {
	event.Dispatch(event.NewAppWillFinishStartup(&App))
}

//export appDidFinishStartup
func appDidFinishStartup() {
	event.Dispatch(event.NewAppDidFinishStartup(&App))
}

// AttemptQuit initiates the termination sequence.
func AttemptQuit() {
	C.platformAttemptTerminate()
}

//export appShouldTerminate
func appShouldTerminate() C.int {
	e := event.NewAppTerminationRequested(&App)
	event.Dispatch(e)
	if e.Cancelled() {
		return C.platformTerminateCancel
	}
	if e.Delayed() {
		return C.platformTerminateLater
	}
	return C.platformTerminateNow
}

// MayTerminateNow resumes the termination sequence that was delayed by calling Delay() on the
// AppTerminationRequested event.
func MayTerminateNow(terminate bool) {
	var value C.int
	if terminate {
		value = 1
	} else {
		value = 0
	}
	C.platformAppMayTerminateNow(value)
}

//export appShouldTerminateAfterLastWindowClosed
func appShouldTerminateAfterLastWindowClosed() bool {
	e := event.NewAppLastWindowClosed(&App)
	event.Dispatch(e)
	return e.Terminate()
}

//export appWillTerminate
func appWillTerminate() {
	event.Dispatch(event.NewAppWillTerminate(&App))
}

//export appWillBecomeActive
func appWillBecomeActive() {
	event.Dispatch(event.NewAppWillActivate(&App))
}

//export appDidBecomeActive
func appDidBecomeActive() {
	event.Dispatch(event.NewAppDidActivate(&App))
}

//export appWillResignActive
func appWillResignActive() {
	event.Dispatch(event.NewAppWillDeactivate(&App))
}

//export appDidResignActive
func appDidResignActive() {
	event.Dispatch(event.NewAppDidDeactivate(&App))
}
