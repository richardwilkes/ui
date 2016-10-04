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
	"github.com/richardwilkes/ui/menu/custom"
)

func platformStartUserInterface() {
	custom.Install()
	// RAW: Implement for Windows
}

func platformAppName() string {
	// RAW: Implement platformAppName for Windows
	return "<unknown>"
}

func platformHideApp() {
	// RAW: Implement for Windows
}

func platformHideOtherApps() {
	// RAW: Implement for Windows
}

func platformShowAllApps() {
	// RAW: Implement for Windows
}

func platformAttemptQuit() {
	// RAW: Implement for Windows
}

func platformAppMayQuitNow(quit bool) {
	// RAW: Implement for Windows
}
