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
	// #cgo darwin LDFLAGS: -framework Cocoa
	// #include "App_darwin.h"
	"C"
)

func platformStartUserInterface() {
	C.platformStartUserInterface()
}

func platformAppName() string {
	return C.GoString(C.platformAppName())
}

func platformHideApp() {
	C.platformHideApp()
}

func platformHideOtherApps() {
	C.platformHideOtherApps()
}

func platformShowAllApps() {
	C.platformShowAllApps()
}

//export callbackAppShouldQuit
func callbackAppShouldQuit() int {
	return int(quit.AppShouldQuit())
}

//export callbackAppShouldQuitAfterLastWindowClosed
func callbackAppShouldQuitAfterLastWindowClosed() bool {
	return quit.AppShouldQuitAfterLastWindowClosed()
}

//export callbackAppWillQuit
func callbackAppWillQuit() {
	quit.AppWillQuit()
}
