// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package color

// #cgo darwin LDFLAGS: -framework Cocoa
// #include "SystemColors_darwin.h"
import "C"

func init() {
	KeyboardFocus = Color(C.platformKeyboardFocusColor())
	SelectedTextBackground = Color(C.platformSelectedTextBackgroundColor())
	SelectedText = Color(C.platformSelectedTextColor())
	TextBackground = Color(C.platformTextBackgroundColor())
	Text = Color(C.platformTextColor())
}
