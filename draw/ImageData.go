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
