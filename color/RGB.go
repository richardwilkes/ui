// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package color

// RGB creates a new opaque Color from RGB values in the range 0-255.
func RGB(red, green, blue int) Color {
	return RGBA(red, green, blue, 1)
}

// RGBA creates a new Color from RGB values in the range 0-255 and an alpha value in the range 0-1.
func RGBA(red, green, blue int, alpha float64) Color {
	return Color(clamp0To1AndScale255(alpha)<<24 | clamp0To255(red)<<16 | clamp0To255(green)<<8 | clamp0To255(blue))
}

// RGBAfloat creates a new Color from RGBA values in the range 0-1.
func RGBAfloat(red, green, blue, alpha float64) Color {
	return Color(clamp0To1AndScale255(alpha)<<24 | clamp0To1AndScale255(red)<<16 | clamp0To1AndScale255(green)<<8 | clamp0To1AndScale255(blue))
}

// Red returns the red channel, in the range of 0-255.
func (c Color) Red() int {
	return int((c >> 16) & 0xFF)
}

// SetRed returns a new color based on this color, but with the red channel replaced.
func (c Color) SetRed(red int) Color {
	return RGBA(red, c.Green(), c.Blue(), c.AlphaIntensity())
}

// RedIntensity returns the red channel, in the range of 0-1.
func (c Color) RedIntensity() float64 {
	return float64(c.Red()) / 255
}

// SetRedIntensity returns a new color based on this color, but with the red channel replaced.
func (c Color) SetRedIntensity(red float64) Color {
	return RGBA(clamp0To1AndScale255(red), c.Green(), c.Blue(), c.AlphaIntensity())
}

// Green returns the green channel, in the range of 0-255.
func (c Color) Green() int {
	return int((c >> 8) & 0xFF)
}

// SetGreen returns a new color based on this color, but with the green channel replaced.
func (c Color) SetGreen(green int) Color {
	return RGBA(c.Red(), green, c.Blue(), c.AlphaIntensity())
}

// GreenIntensity returns the green channel, in the range of 0-1.
func (c Color) GreenIntensity() float64 {
	return float64(c.Green()) / 255
}

// SetGreenIntensity returns a new color based on this color, but with the green channel replaced.
func (c Color) SetGreenIntensity(green float64) Color {
	return RGBA(c.Red(), clamp0To1AndScale255(green), c.Blue(), c.AlphaIntensity())
}

// Blue returns the blue channel, in the range of 0-255.
func (c Color) Blue() int {
	return int(c & 0xFF)
}

// SetBlue returns a new color based on this color, but with the blue channel replaced.
func (c Color) SetBlue(blue int) Color {
	return RGBA(c.Red(), c.Green(), blue, c.AlphaIntensity())
}

// BlueIntensity returns the blue channel, in the range of 0-1.
func (c Color) BlueIntensity() float64 {
	return float64(c.Blue()) / 255
}

// SetBlueIntensity returns a new color based on this color, but with the blue channel replaced.
func (c Color) SetBlueIntensity(blue float64) Color {
	return RGBA(c.Red(), c.Green(), clamp0To1AndScale255(blue), c.AlphaIntensity())
}
