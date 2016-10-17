// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package window

// #include "Types.h"
import "C"

// Define Go constants for the C constants. I do this solely because some of the automated tools I
// use don't work well in the presence of the 'import "C"' directive, so I'm just trying to minimize
// the files it appears in.

const (
	platformMouseDown    platformEventType = C.platformMouseDown
	platformMouseDragged platformEventType = C.platformMouseDragged
	platformMouseUp      platformEventType = C.platformMouseUp
	platformMouseEntered platformEventType = C.platformMouseEntered
	platformMouseMoved   platformEventType = C.platformMouseMoved
	platformMouseExited  platformEventType = C.platformMouseExited
)

type platformEventType C.int
type platformWindow C.platformWindow
type platformSurface C.platformSurface
