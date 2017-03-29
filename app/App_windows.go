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
	"os"
	"path/filepath"

	"github.com/richardwilkes/ui/menu/custom"
	"github.com/richardwilkes/ui/window"
)

func platformStartUserInterface() {
	custom.Install()
	// RAW: Implement for Windows
}

func platformAppName() string {
	return filepath.Base(os.Args[0])
}

func platformHideApp() {
	for _, wnd := range window.Windows() {
		wnd.Minimize()
	}
}

func platformHideOtherApps() {
	panic("platformHideOtherApps() is not implemented")
}

func platformShowAllApps() {
	panic("platformShowAllApps() is not implemented")
}
