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
	"runtime"
)

// #cgo darwin LDFLAGS: -framework Cocoa
// #include <stdlib.h>
// #include "app.h"
import "C"

const (
	// TerminateCancel indicates the termination sequence should be aborted
	TerminateCancel TerminationResponse = iota
	// TerminateNow indicates the termination sequence should proceed immediately
	TerminateNow
	// TerminateLater indicates the termination sequence should proceed only after a call is
	// made to AppMayTerminateNow()
	TerminateLater
)

// TerminationResponse is used to determine what will occur when AttemptQuit() is called.
type TerminationResponse int

var (
	// AppWillFinishStartup is called prior to the app finishing its startup sequence.
	AppWillFinishStartup func()
	// AppDidFinishStartup is called after the app has finished its startup sequence.
	AppDidFinishStartup func()
	// AppShouldTerminate is called to determine whether it is permitted to quit at this point in
	// time.
	AppShouldTerminate func() TerminationResponse
	// AppShouldTerminateAfterLastWindowClosed is called when the last open window is closed to
	// determine if the app should attempt to quit.
	AppShouldTerminateAfterLastWindowClosed func() bool
	// AppWillTerminate is called just prior to the application's termination.
	AppWillTerminate func()
	// AppWillBecomeActive is called prior to the application transitioning to the foreground.
	AppWillBecomeActive func()
	// AppDidBecomeActive is called after the application has transitioned to the foreground.
	AppDidBecomeActive func()
	// AppWillResignActive is called prior to the applicaton transitioning to the background.
	AppWillResignActive func()
	// AppDidResignActive is called after the application has transitioned to the background.
	AppDidResignActive func()
)

// Start the ui.
func Start() {
	runtime.LockOSThread()
	C.uiStart()
}

// AppName returns the application's name.
func AppName() string {
	return C.GoString(C.uiAppName())
}

// AttemptQuit initiates the termination sequence.
func AttemptQuit() {
	C.uiAttemptTerminate()
}

// HideApp attempts to hide this application.
func HideApp() {
	C.uiHideApp()
}

// HideOtherApps attempts to hide other applications, leaving just this application visible.
func HideOtherApps() {
	C.uiHideOtherApps()
}

// ShowAllApps attempts to show all applications that are currently hidden.
func ShowAllApps() {
	C.uiShowAllApps()
}

//export appWillFinishStartup
func appWillFinishStartup() {
	if AppWillFinishStartup != nil {
		AppWillFinishStartup()
	}
}

//export appDidFinishStartup
func appDidFinishStartup() {
	if AppDidFinishStartup != nil {
		AppDidFinishStartup()
	}
}

//export appShouldTerminate
func appShouldTerminate() TerminationResponse {
	if AppShouldTerminate != nil {
		return AppShouldTerminate()
	}
	return TerminateNow
}

// AppMayTerminateNow resumes the termination sequence that was paused by responding with
// TerminateLater when ShouldTerminate() was called..
func AppMayTerminateNow(terminate bool) {
	var value C.int
	if terminate {
		value = 1
	} else {
		value = 0
	}
	C.uiAppMayTerminateNow(value)
}

//export appShouldTerminateAfterLastWindowClosed
func appShouldTerminateAfterLastWindowClosed() bool {
	if AppShouldTerminateAfterLastWindowClosed != nil {
		return AppShouldTerminateAfterLastWindowClosed()
	}
	return true
}

//export appWillTerminate
func appWillTerminate() {
	if AppWillTerminate != nil {
		AppWillTerminate()
	}
}

//export appWillBecomeActive
func appWillBecomeActive() {
	if AppWillBecomeActive != nil {
		AppWillBecomeActive()
	}
}

//export appDidBecomeActive
func appDidBecomeActive() {
	if AppDidBecomeActive != nil {
		AppDidBecomeActive()
	}
}

//export appWillResignActive
func appWillResignActive() {
	if AppWillResignActive != nil {
		AppWillResignActive()
	}
}

//export appDidResignActive
func appDidResignActive() {
	if AppDidResignActive != nil {
		AppDidResignActive()
	}
}
