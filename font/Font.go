// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package font

import (
	"github.com/richardwilkes/ui/geom"
	"unsafe"
	// #cgo pkg-config: pangocairo
	// #include <pango/pangocairo.h>
	"C"
)

const (
	// PangoScale is a constant used by the Pango font engine.
	PangoScale = float64(C.PANGO_SCALE)
)

var (
	surface = C.cairo_image_surface_create(C.CAIRO_FORMAT_ARGB32, 8, 8)
	context = C.pango_font_map_create_context(C.pango_cairo_font_map_get_default())
)

// Font represents a font.
type Font struct {
	pfd           *C.PangoFontDescription
	leading       float64
	ascent        float64
	descent       float64
	monospaced    uint8
	metricsLoaded bool
}

func init() {
	// Tell Pango we want our typographic points to be 1/72 of an inch.
	C.pango_cairo_font_map_set_resolution(C.pango_cairo_font_map_get_default(), 72)
}

// NewFont creates a font from a string.
func NewFont(str string) *Font {
	cstr := C.CString(str)
	defer C.g_free(cstr)
	return &Font{pfd: C.pango_font_description_from_string(cstr)}
}

// Copy makes a copy of a font.
func (d *Font) Copy() *Font {
	desc := &Font{}
	*desc = *d
	desc.pfd = C.pango_font_description_copy(d.pfd)
	return desc
}

// Dispose of the underlying OS resources.
func (d *Font) Dispose() {
	C.pango_font_description_free(d.pfd)
	d.pfd = nil
}

// PangeFontDescription returns the pointer to the underlying Pango font description.
func (d *Font) PangoFontDescription() unsafe.Pointer {
	return unsafe.Pointer(d.pfd)
}

// Equals compares two fonts for equality. Two fonts are considered equal if the fonts they
// describe are provably identical. This means that their masks do not have to match, as long as
// other fields are all the same. Two fonts may result in identical fonts being loaded, but still
// compare false.
func (d *Font) Equals(other *Font) bool {
	return C.pango_font_description_equal(d.pfd, other.pfd) != 0
}

// BetterMatch determines if the style attributes of newDesc are a closer match for this font
// than those of oldDesc are, or if oldDesc is nil, determines if newDesc is a match at
// all. Approximate matching is done for weight and style; other style attributes must match
// exactly. Style attributes are all attributes other than family and size-related attributes.
// Approximate matching for style considers SlantOblique and SlantItalic as matches, but not as
// good a match as when the slants are equal. Note that oldDesc must match this font or be nil.
func (d *Font) BetterMatch(oldDesc *Font, newDesc *Font) bool {
	var old *C.PangoFontDescription
	if oldDesc != nil {
		old = oldDesc.pfd
	}
	return C.pango_font_description_better_match(d.pfd, old, newDesc.pfd) != 0
}

// Family returns the family name. May be empty if it has not been set.
func (d *Font) Family() string {
	return C.GoString(C.pango_font_description_get_family(d.pfd))
}

// SetFamily sets the family name.
func (d *Font) SetFamily(family string) {
	cstr := C.CString(family)
	C.pango_font_description_set_family(d.pfd, cstr)
	C.g_free(cstr)
	d.metricsLoaded = false
}

// Slant returns the font slant.
func (d *Font) Slant() Slant {
	return Slant(C.pango_font_description_get_style(d.pfd))
}

// SetSlant sets the font slant.
func (d *Font) SetSlant(slant Slant) {
	C.pango_font_description_set_style(d.pfd, C.PangoStyle(slant))
	d.metricsLoaded = false
}

// Capitalization returns the font capitalization.
func (d *Font) Capitalization() Capitalization {
	return Capitalization(C.pango_font_description_get_variant(d.pfd))
}

// SetCapitalization sets the font capitalization.
func (d *Font) SetCapitalization(capitalization Capitalization) {
	C.pango_font_description_set_variant(d.pfd, C.PangoVariant(capitalization))
	d.metricsLoaded = false
}

// Weight returns the font weight.
func (d *Font) Weight() Weight {
	return Weight(C.pango_font_description_get_weight(d.pfd))
}

// SetWeight sets the font weight.
func (d *Font) SetWeight(weight Weight) {
	C.pango_font_description_set_weight(d.pfd, C.PangoWeight(weight))
	d.metricsLoaded = false
}

// Stretch returns the font stretch.
func (d *Font) Stretch() Stretch {
	return Stretch(C.pango_font_description_get_stretch(d.pfd))
}

// SetStretch sets the font stretch.
func (d *Font) SetStretch(stretch Stretch) {
	C.pango_font_description_set_stretch(d.pfd, C.PangoStretch(stretch))
	d.metricsLoaded = false
}

// Size the size of the font, in points.
func (d *Font) Size() float64 {
	return float64(C.pango_font_description_get_size(d.pfd)) / PangoScale
}

// SetSize sets the size of the font, in points.
func (d *Font) SetSize(size float64) {
	C.pango_font_description_set_size(d.pfd, C.gint(size*PangoScale))
	d.metricsLoaded = false
}

// String returns a string that can be used with NewFont.
func (d *Font) String() string {
	cstr := C.pango_font_description_to_string(d.pfd)
	defer C.g_free(cstr)
	return C.GoString(cstr)
}

// Monospaced returns true if this font has a fixed width.
func (d *Font) Monospaced() bool {
	d.loadMetrics()
	return d.monospaced == 1
}

// Leading returns the amount of space before the ascent.
func (d *Font) Leading() float64 {
	d.loadMetrics()
	return d.leading
}

// Ascent returns the amount of space used from the baseline to the top of the tallest character.
func (d *Font) Ascent() float64 {
	d.loadMetrics()
	return d.ascent
}

// Descent returns the amount of space used from the baseline to the bottom.
func (d *Font) Descent() float64 {
	d.loadMetrics()
	return d.descent
}

// Height returns the overall height of the font, effectively Leading() + Ascent() + Descent().
func (d *Font) Height() float64 {
	d.loadMetrics()
	return d.leading + d.ascent + d.descent
}

func (d *Font) loadMetrics() {
	if !d.metricsLoaded {
		d.metricsLoaded = true
		if f := C.pango_font_map_load_font(C.pango_cairo_font_map_get_default(), context, d.pfd); f != nil {
			metrics := C.pango_font_get_metrics(f, C.pango_language_get_default())
			d.ascent = float64(C.pango_font_metrics_get_ascent(metrics)) / PangoScale
			d.descent = float64(C.pango_font_metrics_get_descent(metrics)) / PangoScale
			d.leading = d.descent / 2
			C.pango_font_metrics_unref(metrics)
		} else {
			d.leading = 0
			d.ascent = 0
			d.descent = 0
		}
		// The monospaced attribute only needs to ever be determined once.
		if d.monospaced == 0 {
			d.monospaced = 2
			for _, family := range MonospacedFamilies() {
				if family.Name() == d.Family() {
					d.monospaced = 1
					break
				}
			}
		}
	}
}

// Measure the string rendered with this font. If you want to account for transformations in a
// graphics context, you must use the Measure method on the Graphics object instead.
func (d *Font) Measure(text string) geom.Size {
	layout, gc := d.createLayout(text)
	var width, height C.int
	C.pango_layout_get_size(layout, &width, &height)
	d.destroyLayout(layout, gc)
	return geom.Size{Width: float64(width) / PangoScale, Height: d.Leading() + float64(height)/PangoScale}
}

// IndexForPosition returns the rune index within the string for the specified x-coordinate, where
// 0 is the start of the string.
func (d *Font) IndexForPosition(x float64, text string) int {
	layout, gc := d.createLayout(text)
	var index, trailing C.int
	C.pango_layout_xy_to_index(layout, C.int(x*PangoScale), 0, &index, &trailing)
	d.destroyLayout(layout, gc)
	return int(index + trailing)
}

// PositionForIndex returns the x-coordinate where the specified rune index starts. The returned
// coordinate assumes 0 is the start of the string.
func (d *Font) PositionForIndex(index int, text string) float64 {
	layout, gc := d.createLayout(text)
	var x C.int
	C.pango_layout_index_to_line_x(layout, C.int(index), 0, nil, &x)
	d.destroyLayout(layout, gc)
	return float64(x) / PangoScale
}

func (d *Font) createLayout(text string) (layout *C.PangoLayout, gc *C.cairo_t) {
	gc = C.cairo_create(surface)
	layout = C.pango_cairo_create_layout(gc)
	C.pango_layout_set_font_description(layout, d.pfd)
	C.pango_layout_set_spacing(layout, C.int(d.Leading()*PangoScale))
	cstr := C.CString(text)
	C.pango_layout_set_text(layout, cstr, -1)
	C.g_free(cstr)
	return layout, gc
}

func (d *Font) destroyLayout(layout *C.PangoLayout, gc *C.cairo_t) {
	C.g_object_unref(layout)
	C.cairo_destroy(gc)
}
