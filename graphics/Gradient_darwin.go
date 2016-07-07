// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package graphics

// #cgo darwin LDFLAGS: -framework Cocoa
// #include <CoreGraphics/CoreGraphics.h>
import "C"

func (g *Gradient) toPlatform() C.CGGradientRef {
	colorSpace := C.CGColorSpaceCreateDeviceRGB()
	count := len(g.Stops)
	components := make([]C.CGFloat, count*4)
	locs := make([]C.CGFloat, count)
	for i, one := range g.Stops {
		j := i * 4
		components[j] = C.CGFloat(one.Color.RedIntensity())
		components[j+1] = C.CGFloat(one.Color.GreenIntensity())
		components[j+2] = C.CGFloat(one.Color.BlueIntensity())
		components[j+3] = C.CGFloat(one.Color.AlphaIntensity())
		locs[i] = C.CGFloat(one.Location)
	}
	gradient := C.CGGradientCreateWithColorComponents(colorSpace, &components[0], &locs[0], C.size_t(count))
	C.CGGradientRetain(gradient)
	C.CGColorSpaceRelease(colorSpace)
	return gradient
}
