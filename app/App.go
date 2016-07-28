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
	// made to MayTerminateNow()
	TerminateLater
)

// TerminationResponse is used to determine what will occur when AttemptQuit() is called.
type TerminationResponse int

var (
	// WillFinishStartup is called prior to the application finishing its startup sequence.
	WillFinishStartup func()
	// DidFinishStartup is called after the application has finished its startup sequence.
	DidFinishStartup func()
	// ShouldTerminate is called to determine whether it is permitted to quit at this point in time.
	ShouldTerminate func() TerminationResponse
	// ShouldTerminateAfterLastWindowClosed is called when the last open window is closed to
	// determine if the application should attempt to quit.
	ShouldTerminateAfterLastWindowClosed func() bool
	// WillTerminate is called just prior to the application's termination.
	WillTerminate func()
	// WillBecomeActive is called prior to the application transitioning to the foreground.
	WillBecomeActive func()
	// DidBecomeActive is called after the application has transitioned to the foreground.
	DidBecomeActive func()
	// WillResignActive is called prior to the application transitioning to the background.
	WillResignActive func()
	// DidResignActive is called after the application has transitioned to the background.
	DidResignActive func()
)

// Start the ui.
func Start() {
	runtime.LockOSThread()
	C.uiStart()
}

// Name returns the application's name.
func Name() string {
	return C.GoString(C.uiAppName())
}

// AttemptQuit initiates the termination sequence.
func AttemptQuit() {
	C.uiAttemptTerminate()
}

// Hide attempts to hide this application.
func Hide() {
	C.uiHideApp()
}

// HideOthers attempts to hide other applications, leaving just this application visible.
func HideOthers() {
	C.uiHideOtherApps()
}

// ShowAll attempts to show all applications that are currently hidden.
func ShowAll() {
	C.uiShowAllApps()
}

//export appWillFinishStartup
func appWillFinishStartup() {
	if WillFinishStartup != nil {
		WillFinishStartup()
	}
}

//export appDidFinishStartup
func appDidFinishStartup() {
	if DidFinishStartup != nil {
		DidFinishStartup()
	}
}

//export appShouldTerminate
func appShouldTerminate() TerminationResponse {
	if ShouldTerminate != nil {
		return ShouldTerminate()
	}
	return TerminateNow
}

// MayTerminateNow resumes the termination sequence that was paused by responding with
// TerminateLater when ShouldTerminate() was called..
func MayTerminateNow(terminate bool) {
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
	if ShouldTerminateAfterLastWindowClosed != nil {
		return ShouldTerminateAfterLastWindowClosed()
	}
	return true
}

//export appWillTerminate
func appWillTerminate() {
	if WillTerminate != nil {
		WillTerminate()
	}
}

//export appWillBecomeActive
func appWillBecomeActive() {
	if WillBecomeActive != nil {
		WillBecomeActive()
	}
}

//export appDidBecomeActive
func appDidBecomeActive() {
	if DidBecomeActive != nil {
		DidBecomeActive()
	}
}

//export appWillResignActive
func appWillResignActive() {
	if WillResignActive != nil {
		WillResignActive()
	}
}

//export appDidResignActive
func appDidResignActive() {
	if DidResignActive != nil {
		DidResignActive()
	}
}
