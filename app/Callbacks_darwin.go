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
	"github.com/richardwilkes/ui/menu/macmenus"

	"C"
)

//export callbackAppShouldQuit
// nolint: deadcode
func callbackAppShouldQuit() int {
	return int(quit.AppShouldQuit())
}

//export callbackAppShouldQuitAfterLastWindowClosed
// nolint: deadcode
func callbackAppShouldQuitAfterLastWindowClosed() bool {
	return quit.AppShouldQuitAfterLastWindowClosed()
}

//export callbackAppWillQuit
// nolint: deadcode
func callbackAppWillQuit() {
	event.SendAppWillQuit()
}

//export callbackAppWillFinishStartup
// nolint: deadcode
func callbackAppWillFinishStartup() {
	macmenus.Install()
	event.SendAppWillFinishStartup()
}

//export callbackAppDidFinishStartup
// nolint: deadcode
func callbackAppDidFinishStartup() {
	event.SendAppDidFinishStartup()
}

//export callbackAppWillBecomeActive
// nolint: deadcode
func callbackAppWillBecomeActive() {
	event.SendAppWillActivate()
}

//export callbackAppDidBecomeActive
// nolint: deadcode
func callbackAppDidBecomeActive() {
	event.SendAppDidActivate()
}

//export callbackAppWillResignActive
// nolint: deadcode
func callbackAppWillResignActive() {
	event.SendAppWillDeactivate()
}

//export callbackAppDidResignActive
// nolint: deadcode
func callbackAppDidResignActive() {
	event.SendAppDidDeactivate()
}
