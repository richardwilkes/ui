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
	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui/color"

	// #cgo pkg-config: pangocairo
	// #include <pango/pangocairo.h>
	"C"
)

// PaintKind possibilities.
const (
	PaintKindColor          PaintKind = C.CAIRO_PATTERN_TYPE_SOLID
	PaintKindImage          PaintKind = C.CAIRO_PATTERN_TYPE_SURFACE
	PaintKindLinearGradient PaintKind = C.CAIRO_PATTERN_TYPE_LINEAR
	PaintKindRadialGradient PaintKind = C.CAIRO_PATTERN_TYPE_RADIAL
)

// PaintKind defines the possible types of Paint.
type PaintKind int

const (
	// PaintModeNone pixels outside of the pattern are fully transparent.
	PaintModeNone PaintMode = C.CAIRO_EXTEND_NONE
	// PaintModeRepeat the pattern is tiled by repeating.
	PaintModeRepeat PaintMode = C.CAIRO_EXTEND_REPEAT
	// PaintModeReflect the pattern is tiled by reflecting at the edges.
	PaintModeReflect PaintMode = C.CAIRO_EXTEND_REFLECT
	// PaintModePad pixels outside of the pattern copy the closest pixel from the source.
	PaintModePad PaintMode = C.CAIRO_EXTEND_PAD
)

// PaintMode defines how pixels outside a pattern's space are filled in.
type PaintMode int

const (
	// FilterModeFast provides a high-performance filter, with quality similar to FilterModeNearest.
	FilterModeFast FilterMode = C.CAIRO_FILTER_FAST
	// FilterModeGood provides a reasonable-performance filter, with quality similar to
	// FilterModeBilinear.
	FilterModeGood FilterMode = C.CAIRO_FILTER_GOOD
	// FilterModeBest provides the highest quality filter available, but performance may not be
	// suitable for interactive use.
	FilterModeBest FilterMode = C.CAIRO_FILTER_BEST
	// FilterModeNearest provides a nearest-neighbor filter.
	FilterModeNearest FilterMode = C.CAIRO_FILTER_NEAREST
	// FilterModeBilinear provides linear interpolation in two dimensions.
	FilterModeBilinear FilterMode = C.CAIRO_FILTER_BILINEAR
)

// FilterMode defines the type of filtering to be used when reading pixel values.
type FilterMode int

// Paint holds patterns used when filling and stroking paths.
type Paint struct {
	pattern *C.cairo_pattern_t
}

// NewColorPaint creates a new Paint from a color.
func NewColorPaint(color color.Color) Paint {
	return Paint{pattern: C.cairo_pattern_create_rgba(C.double(color.RedIntensity()), C.double(color.GreenIntensity()), C.double(color.BlueIntensity()), C.double(color.AlphaIntensity()))}
}

// NewImagePaint creates a new Paint from an image.
func NewImagePaint(img *Image) Paint {
	return Paint{pattern: C.cairo_pattern_create_for_surface(img.surface)}
}

// NewLinearGradientPaint creates a new Paint from a gradient that is spread across the specified
// line.
func NewLinearGradientPaint(gradient *Gradient, sx, sy, ex, ey float64) Paint {
	pattern := C.cairo_pattern_create_linear(C.double(sx), C.double(sy), C.double(ex), C.double(ey))
	for _, one := range gradient.Stops {
		C.cairo_pattern_add_color_stop_rgba(pattern, C.double(one.Location), C.double(one.Color.RedIntensity()), C.double(one.Color.GreenIntensity()), C.double(one.Color.BlueIntensity()), C.double(one.Color.AlphaIntensity()))
	}
	return Paint{pattern: pattern}
}

// NewRadialGradientPaint creates a new Paint from a gradient that radiates from one circle to
// another.
func NewRadialGradientPaint(gradient *Gradient, scx, scy, startRadius, ecx, ecy, endRadius float64) Paint {
	pattern := C.cairo_pattern_create_radial(C.double(scx), C.double(scy), C.double(startRadius), C.double(ecx), C.double(ecy), C.double(endRadius))
	for _, one := range gradient.Stops {
		C.cairo_pattern_add_color_stop_rgba(pattern, C.double(one.Location), C.double(one.Color.RedIntensity()), C.double(one.Color.GreenIntensity()), C.double(one.Color.BlueIntensity()), C.double(one.Color.AlphaIntensity()))
	}
	return Paint{pattern: pattern}
}

// PaintMode returns the PaintMode of this Paint.
func (p Paint) PaintMode() PaintMode {
	return PaintMode(C.cairo_pattern_get_extend(p.pattern))
}

// SetPaintMode sets the PaintMode of this Paint.
func (p Paint) SetPaintMode(mode PaintMode) {
	C.cairo_pattern_set_extend(p.pattern, C.cairo_extend_t(mode))
}

// FilterMode returns the FilterMode of this Paint.
func (p Paint) FilterMode() FilterMode {
	return FilterMode(C.cairo_pattern_get_filter(p.pattern))
}

// SetFilterMode sets the FilterMode of this Paint.
func (p Paint) SetFilterMode(mode FilterMode) {
	C.cairo_pattern_set_filter(p.pattern, C.cairo_filter_t(mode))
}

// Matrix returns the Matrix of this Paint.
func (p Paint) Matrix() *geom.Matrix {
	var matrix C.cairo_matrix_t
	C.cairo_pattern_get_matrix(p.pattern, &matrix)
	return fromCairoMatrix(&matrix)
}

// SetMatrix sets the Matrix of this Paint.
func (p Paint) SetMatrix(matrix *geom.Matrix) {
	C.cairo_pattern_set_matrix(p.pattern, toCairoMatrix(matrix))
}

// Kind returns the PaintKind of this Paint.
func (p Paint) Kind() PaintKind {
	return PaintKind(C.cairo_pattern_get_type(p.pattern))
}

// Dispose releases the underlying resources used by this Paint.
func (p Paint) Dispose() {
	C.cairo_pattern_destroy(p.pattern)
	p.pattern = nil
}
