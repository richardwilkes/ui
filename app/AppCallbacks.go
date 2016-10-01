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
	"C"
	"github.com/richardwilkes/ui/event"
)

// Possible termination responses
const (
	QuitCancel QuitResponse = iota
	QuitNow
	QuitLater // Must make a call to MayQuitNow() at some point in the future.
)

// QuitResponse is used to respond to requests for app termination.
type QuitResponse int

//export appWillFinishStartup
func appWillFinishStartup() {
	event.Dispatch(event.NewAppWillFinishStartup(&App))
}

//export appDidFinishStartup
func appDidFinishStartup() {
	event.Dispatch(event.NewAppDidFinishStartup(&App))
}

//export appShouldQuit
func appShouldQuit() QuitResponse {
	e := event.NewAppQuitRequested(&App)
	event.Dispatch(e)
	if e.Cancelled() {
		return QuitCancel
	}
	if e.Delayed() {
		return QuitLater
	}
	return QuitNow
}

//export appShouldQuitAfterLastWindowClosed
func appShouldQuitAfterLastWindowClosed() bool {
	e := event.NewAppLastWindowClosed(&App)
	event.Dispatch(e)
	return e.Quit()
}

//export appWillQuit
func appWillQuit() {
	event.Dispatch(event.NewAppWillQuit(&App))
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
