// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package draw

import (
	"image"
	gocolor "image/color"

	"github.com/richardwilkes/ui/color"
)

// ImageData is the raw information that makes up an Image.
type ImageData struct {
	Width  int
	Height int
	Pixels []color.Color
}

// ColorModel returns the Image's color model. (Implementation of image.Image)
func (img *ImageData) ColorModel() gocolor.Model {
	return gocolor.NRGBAModel
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0). (Implementation of image.Image)
func (img *ImageData) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one. (Implementation of image.Image)
func (img *ImageData) At(x, y int) gocolor.Color {
	pixel := img.Pixels[y*img.Width+x]
	return gocolor.NRGBA{R: uint8(pixel.Red()), G: uint8(pixel.Green()), B: uint8(pixel.Blue()), A: uint8(pixel.Alpha())}
}
