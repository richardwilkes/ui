// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package font

func init() {
	// RAW: Try to determine what the user set in system preferences...
	User = NewFont("Sans 12")
	UserMonospaced = NewFont("Monospace 10")
	System = NewFont("Sans 13")
	EmphasizedSystem = NewFont("Sans Bold 13")
	SmallSystem = NewFont("Sans 11")
	SmallEmphasizedSystem = NewFont("Sans Bold 11")
	Views = NewFont("Sans 12")
	Label = NewFont("Sans 10")
	Menu = NewFont("Sans 14")
	MenuCmdKey = NewFont("Sans 14")
}
