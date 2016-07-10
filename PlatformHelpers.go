// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

// #include "Window.h"
import "C"

func toRect(bounds C.uiRect) Rect {
	return Rect{Point: Point{X: float32(bounds.x), Y: float32(bounds.y)}, Size: Size{Width: float32(bounds.width), Height: float32(bounds.height)}}
}

func toCRect(bounds Rect) C.uiRect {
	return C.uiRect{x: C.float(bounds.X), y: C.float(bounds.Y), width: C.float(bounds.Width), height: C.float(bounds.Height)}
}

func toPoint(pt C.uiPoint) Point {
	return Point{X: float32(pt.x), Y: float32(pt.y)}
}

func toSize(size C.uiSize) Size {
	return Size{Width: float32(size.width), Height: float32(size.height)}
}
