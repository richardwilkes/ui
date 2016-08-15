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
	"github.com/richardwilkes/ui/event"
	"runtime"
)

// #cgo darwin LDFLAGS: -framework Cocoa
// #include "App.h"
import "C"

const (
	terminateCancel terminationResponse = iota
	terminateNow
	terminateLater
)

type terminationResponse int

// Application represents the overall application.
type Application struct {
	eventHandlers event.Handlers
}

// EventHandlers implements the event.Target interface.
func (app *Application) EventHandlers() *event.Handlers {
	return &app.eventHandlers
}

// ParentTarget implements the event.Target interface.
func (app *Application) ParentTarget() event.Target {
	return nil
}

var (
	// App provides the top-level event distribution point. Events that cascade will flow from the
	// widgets, to their parents, to their window, then finally to this app if not handled somewhere
	// along the line.
	App Application
)

// Start the ui.
func Start() {
	runtime.LockOSThread()
	C.platformStart()
}

// Name returns the application's name.
func Name() string {
	return C.GoString(C.platformAppName())
}

// AttemptQuit initiates the termination sequence.
func AttemptQuit() {
	C.platformAttemptTerminate()
}

// Hide attempts to hide this application.
func Hide() {
	C.platformHideApp()
}

// HideOthers attempts to hide other applications, leaving just this application visible.
func HideOthers() {
	C.platformHideOtherApps()
}

// ShowAll attempts to show all applications that are currently hidden.
func ShowAll() {
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

//export appShouldTerminate
func appShouldTerminate() terminationResponse {
	e := event.NewAppTerminationRequested(&App)
	event.Dispatch(e)
	if e.Cancelled() {
		return terminateCancel
	}
	if e.Delayed() {
		return terminateLater
	}
	return terminateNow
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
