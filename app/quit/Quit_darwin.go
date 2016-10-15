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
	// #cgo CFLAGS: -x objective-c
	// #cgo LDFLAGS: -framework Cocoa
	// #include <Cocoa/Cocoa.h>
	// void attemptQuit() { [NSApp terminate:nil]; }
	// void appMayQuitNow(int quit) { [NSApp replyToApplicationShouldTerminate:quit]; }
	"C"
)

func platformAttemptQuit() {
	C.attemptQuit()
}

func platformAppMayQuitNow(quit bool) {
	var mayQuit C.int
	if quit {
		mayQuit = 1
	} else {
		mayQuit = 0
	}
	C.appMayQuitNow(mayQuit)
}
