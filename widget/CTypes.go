// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package widget

import (
	"github.com/richardwilkes/geom"
)

// #include "Types.h"
import "C"

// Define Go constants for the C constants. I do this solely because some of the automated tools I
// use don't work well in the presence of the 'import "C"' directive, so I'm just trying to minimize
// the files it appears in.

// Possible values for the WindowStyleMask.
const (
	BorderlessWindowMask  WindowStyleMask = C.platformBorderlessWindowMask
	TitledWindowMask                      = C.platformTitledWindowMask
	ClosableWindowMask                    = C.platformClosableWindowMask
	MinimizableWindowMask                 = C.platformMinimizableWindowMask
	ResizableWindowMask                   = C.platformResizableWindowMask
	StdWindowMask                         = TitledWindowMask | ClosableWindowMask | MinimizableWindowMask | ResizableWindowMask
)

const (
	platformMouseDown    platformEventType = C.platformMouseDown
	platformMouseDragged platformEventType = C.platformMouseDragged
	platformMouseUp      platformEventType = C.platformMouseUp
	platformMouseEntered platformEventType = C.platformMouseEntered
	platformMouseMoved   platformEventType = C.platformMouseMoved
	platformMouseExited  platformEventType = C.platformMouseExited
	platformMouseWheel   platformEventType = C.platformMouseWheel
	platformKeyDown      platformEventType = C.platformKeyDown
	platformKeyTyped     platformEventType = C.platformKeyTyped
	platformKeyUp        platformEventType = C.platformKeyUp
)

// WindowStyleMask controls the look and capabilities of a window.
type WindowStyleMask C.int

type platformEventType C.int
type platformWindow C.platformWindow
type platformSurface C.platformSurface
type platformRect C.platformRect

func toRect(r C.platformRect) geom.Rect {
	return geom.Rect{Point: geom.Point{X: float64(r.x), Y: float64(r.y)}, Size: geom.Size{Width: float64(r.width), Height: float64(r.height)}}
}

func toCRect(r geom.Rect) C.platformRect {
	return C.platformRect{x: C.double(r.X), y: C.double(r.Y), width: C.double(r.Width), height: C.double(r.Height)}
}
