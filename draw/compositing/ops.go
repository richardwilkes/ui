package compositing

import (
	// #cgo pkg-config: pangocairo
	// #include <pango/pangocairo.h>
	"C"
)

const (
	// Clear clears the destination layer (bounded).
	Clear Op = C.CAIRO_OPERATOR_CLEAR
	// Source replaces the destination layer (bounded).
	Source Op = C.CAIRO_OPERATOR_SOURCE
	// Over draws the source layer on top of the destination layer (bounded).
	Over Op = C.CAIRO_OPERATOR_OVER
	// In draws the source where there was destination content (unbounded).
	In Op = C.CAIRO_OPERATOR_IN
	// Out draws the source where there was no destination content (unbounded).
	Out Op = C.CAIRO_OPERATOR_OUT
	// Atop draws the source where there was no destination content (unbounded).
	Atop Op = C.CAIRO_OPERATOR_ATOP
	// Dest ignores the source.
	Dest Op = C.CAIRO_OPERATOR_DEST
	// DestOver draws the destination on top of the source.
	DestOver Op = C.CAIRO_OPERATOR_DEST_OVER
	// DestIn leaves the destination only where there was source content (unbounded).
	DestIn Op = C.CAIRO_OPERATOR_DEST_IN
	// DestOut leaves the destination only where there was no source content.
	DestOut Op = C.CAIRO_OPERATOR_DEST_OUT
	// DestAtop leaves the destination on top of the source content and only there (unbounded).
	DestAtop Op = C.CAIRO_OPERATOR_DEST_ATOP
	// XOr source and destination are shown where there is only one of them.
	XOr Op = C.CAIRO_OPERATOR_XOR
	// Add source and destination layers are accumulated.
	Add Op = C.CAIRO_OPERATOR_ADD
	// Saturate like over, but assuming source and dest are disjoint geometries.
	Saturate Op = C.CAIRO_OPERATOR_SATURATE
	// Multiply source and destination layers are multiplied. This causes the result to be at least
	// as dark as the darker inputs.
	Multiply Op = C.CAIRO_OPERATOR_MULTIPLY
	// Screen source and destination are complemented and multiplied. This causes the result to be
	// at least as light as the lighter inputs.
	Screen Op = C.CAIRO_OPERATOR_SCREEN
	// Overlay multiplies or screens, depending on the lightness of the destination color.
	Overlay Op = C.CAIRO_OPERATOR_OVERLAY
	// Darken replaces the destination with the source if it is darker, otherwise keeps the source.
	Darken Op = C.CAIRO_OPERATOR_DARKEN
	// Lighten replaces the destination with the source if it is lighter, otherwise keeps the source.
	Lighten Op = C.CAIRO_OPERATOR_LIGHTEN
	// ColorDodge brightens the destination color to reflect the source color.
	ColorDodge Op = C.CAIRO_OPERATOR_COLOR_DODGE
	// ColorBurn darkens the destination color to reflect the source color.
	ColorBurn Op = C.CAIRO_OPERATOR_COLOR_BURN
	// HardLight multiplies or screens, dependent on source color.
	HardLight Op = C.CAIRO_OPERATOR_HARD_LIGHT
	// SoftLight darkens or lightens, dependent on source color.
	SoftLight Op = C.CAIRO_OPERATOR_SOFT_LIGHT
	// Difference takes the difference of the source and destination color.
	Difference Op = C.CAIRO_OPERATOR_DIFFERENCE
	// Exclusion produces an effect similar to difference, but with lower contrast.
	Exclusion Op = C.CAIRO_OPERATOR_EXCLUSION
	// Hue creates a color with the hue of the source and the saturation and luminosity of the
	// target.
	Hue Op = C.CAIRO_OPERATOR_HSL_HUE
	// Saturation creates a color with the saturation of the source and the hue and luminosity of
	// the target. Painting with this mode onto a gray area produces no change.
	Saturation Op = C.CAIRO_OPERATOR_HSL_SATURATION
	// Color creates a color with the hue and saturation of the source and the luminosity of the
	// target. This preserves the gray levels of the target and is useful for coloring monochrome
	// images or tinting color images.
	Color Op = C.CAIRO_OPERATOR_HSL_COLOR
	// Luminosity creates a color with the luminosity of the source and the hue and saturation of
	// the target. This produces an inverse effect to the Color Op.
	Luminosity Op = C.CAIRO_OPERATOR_HSL_LUMINOSITY
)

// Op defines how graphics operations should occur.
type Op int
