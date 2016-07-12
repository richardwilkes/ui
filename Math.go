// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

import (
	"math"
)

const (
	// DegreesToRadians converts a value in degrees to radians when multiplied with the value.
	DegreesToRadians = math.Pi / 180
	// RadiansToDegrees converts a value in radians to degrees when multiplied with the value.
	RadiansToDegrees = 180 / math.Pi
)

// AbsFloat32 returns the absolute value of x.
func AbsFloat32(x float32) float32 {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0 // Removes negative sign from potential -0
	}
	return x
}

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

// MinOf3int returns the smaller of a, b, or c.
func MinOf3int(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

// MaxOf3int returns the larger of a, b, or c.
func MaxOf3int(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}

// RoundFloat32 returns the closest integer.
func RoundFloat32(x float32) float32 {
	return float32(int(x + 0.5))
}

// FloorFloat32 returns the greatest integer value less than or equal to x.
func FloorFloat32(x float32) float32 {
	return float32(math.Floor(float64(x)))
}

// CeilFloat32 returns the smallest integer value greater than or equal to x.
func CeilFloat32(x float32) float32 {
	return float32(math.Ceil(float64(x)))
}
