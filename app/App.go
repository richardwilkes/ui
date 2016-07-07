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
	"unsafe"
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

type TerminationResponse int

var (
	Name                                 string
	WillFinishStartup                    func()
	DidFinishStartup                     func()
	ShouldTerminate                      func() TerminationResponse
	ShouldTerminateAfterLastWindowClosed func() bool
	WillTerminate                        func()
	WillBecomeActive                     func()
	DidBecomeActive                      func()
	WillResignActive                     func()
	DidResignActive                      func()
)

func init() {
	cAppName := C.uiAppName()
	Name = C.GoString(cAppName)
	C.free(unsafe.Pointer(cAppName))
}

// Start the ui.
func Start() {
	runtime.LockOSThread()
	C.uiStart()
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

//export willFinishStartup
func willFinishStartup() {
	if WillFinishStartup != nil {
		WillFinishStartup()
	}
}

//export didFinishStartup
func didFinishStartup() {
	if DidFinishStartup != nil {
		DidFinishStartup()
	}
}

//export shouldTerminate
func shouldTerminate() TerminationResponse {
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

//export shouldTerminateAfterLastWindowClosed
func shouldTerminateAfterLastWindowClosed() bool {
	if ShouldTerminateAfterLastWindowClosed != nil {
		return ShouldTerminateAfterLastWindowClosed()
	}
	return true
}

//export willTerminate
func willTerminate() {
	if WillTerminate != nil {
		WillTerminate()
	}
}

//export willBecomeActive
func willBecomeActive() {
	if WillBecomeActive != nil {
		WillBecomeActive()
	}
}

//export didBecomeActive
func didBecomeActive() {
	if DidBecomeActive != nil {
		DidBecomeActive()
	}
}

//export willResignActive
func willResignActive() {
	if WillResignActive != nil {
		WillResignActive()
	}
}

//export didResignActive
func didResignActive() {
	if DidResignActive != nil {
		DidResignActive()
	}
}
