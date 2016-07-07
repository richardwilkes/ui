// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package maths

// MinFloat32 returns the smaller of a or b. Note that there is no special handling for Inf, NaN,
// or +0 vs -0. If you want/need that, up-cast to float64 and use math.Min().
func MinFloat32(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

// MaxFloat32 returns the larger of a or b. Note that there is no special handling for Inf, NaN,
// or +0 vs -0. If you want/need that, up-cast to float64 and use math.Max().
func MaxFloat32(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

// MinInt returns the smaller of a or b.
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxInt returns the larger of a or b.
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
