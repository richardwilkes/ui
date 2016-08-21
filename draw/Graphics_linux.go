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
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/geom"
)

// #cgo linux LDFLAGS: -lX11 -lcairo
// #include <stdio.h>
// #include <cairo/cairo.h>
import "C"

const (
	oneThird  = 1.0 / 3.0
	twoThirds = 2.0 / 3.0
)

func (gc *graphics) platformSave() {
	C.cairo_save(gc.gc)
}

func (gc *graphics) platformRestore() {
	C.cairo_restore(gc.gc)
}

func (gc *graphics) platformSetOpacity(opacity float32) {
	// RAW: Implement platformSetOpacity for Linux
}

func (gc *graphics) platformSetFillColor(color color.Color) {
	C.cairo_set_source_rgba(gc.gc, C.double(color.RedIntensity()), C.double(color.GreenIntensity()), C.double(color.BlueIntensity()), C.double(color.AlphaIntensity()))
}

func (gc *graphics) platformSetStrokeColor(color color.Color) {
	// RAW: Implement platformSetStrokeColor for Linux
}

func (gc *graphics) platformSetStrokeWidth(width float32) {
	C.cairo_set_line_width(gc.gc, C.double(width))
}

func (gc *graphics) platformFillRect(bounds geom.Rect) {
	C.cairo_rectangle(gc.gc, C.double(bounds.X), C.double(bounds.Y), C.double(bounds.Width), C.double(bounds.Height))
	C.cairo_fill(gc.gc)
}

func (gc *graphics) platformStrokeRect(bounds geom.Rect) {
	color := gc.StrokeColor()
	C.cairo_set_source_rgba(gc.gc, C.double(color.RedIntensity()), C.double(color.GreenIntensity()), C.double(color.BlueIntensity()), C.double(color.AlphaIntensity()))
	C.cairo_rectangle(gc.gc, C.double(bounds.X), C.double(bounds.Y), C.double(bounds.Width), C.double(bounds.Height))
	C.cairo_stroke(gc.gc)
	gc.platformSetFillColor(gc.FillColor())
}

func (gc *graphics) platformFillEllipse(bounds geom.Rect) {
	// RAW: Implement platformFillEllipse for Linux
}

func (gc *graphics) platformStrokeEllipse(bounds geom.Rect) {
	// RAW: Implement platformStrokeEllipse for Linux
}

func (gc *graphics) platformFillPath() {
	C.cairo_fill(gc.gc)
}

func (gc *graphics) platformFillPathEvenOdd() {
	current := C.cairo_get_fill_rule(gc.gc)
	if current != C.CAIRO_FILL_RULE_EVEN_ODD {
		C.cairo_set_fill_rule(gc.gc, C.CAIRO_FILL_RULE_EVEN_ODD)
	}
	C.cairo_fill(gc.gc)
	if current != C.CAIRO_FILL_RULE_EVEN_ODD {
		C.cairo_set_fill_rule(gc.gc, current)
	}
}

func (gc *graphics) platformStrokePath() {
	color := gc.StrokeColor()
	C.cairo_set_source_rgba(gc.gc, C.double(color.RedIntensity()), C.double(color.GreenIntensity()), C.double(color.BlueIntensity()), C.double(color.AlphaIntensity()))
	C.cairo_stroke(gc.gc)
	gc.platformSetFillColor(gc.FillColor())
}

func (gc *graphics) platformBeginPath() {
	C.cairo_new_path(gc.gc)
}

func (gc *graphics) platformClosePath() {
	C.cairo_close_path(gc.gc)
}

func (gc *graphics) platformMoveTo(x, y float32) {
	C.cairo_move_to(gc.gc, C.double(x), C.double(y))
}

func (gc *graphics) platformLineTo(x, y float32) {
	C.cairo_line_to(gc.gc, C.double(x), C.double(y))
}

func (gc *graphics) platformArc(cx, cy, radius, startAngleRadians, endAngleRadians float32, clockwise bool) {
	if clockwise {
		C.cairo_arc(gc.gc, C.double(cx), C.double(cy), C.double(radius), C.double(startAngleRadians), C.double(endAngleRadians))
	} else {
		C.cairo_arc_negative(gc.gc, C.double(cx), C.double(cy), C.double(radius), C.double(startAngleRadians), C.double(endAngleRadians))
	}
}

func (gc *graphics) platformArcTo(x1, y1, x2, y2, radius float32) {
	// RAW: Implement platformArcTo for Linux
}

func (gc *graphics) platformCurveTo(cp1x, cp1y, cp2x, cp2y, x, y float32) {
	C.cairo_curve_to(gc.gc, C.double(cp1x), C.double(cp1y), C.double(cp2x), C.double(cp2y), C.double(x), C.double(y))
}

func (gc *graphics) platformQuadCurveTo(cpx, cpy, x, y float32) {
	var x0, y0 C.double
	C.cairo_get_current_point(gc.gc, &x0, &y0)
	gc.platformCurveTo(twoThirds*cpx+oneThird*float32(x0), twoThirds*cpy+oneThird*float32(y0), twoThirds*cpx+oneThird*x, twoThirds*cpy+oneThird*y, x, y)
}

func (gc *graphics) platformAddPath(path *geom.Path) {
	path.Apply(gc.gc)
}

func (gc *graphics) platformClip() {
	C.cairo_clip(gc.gc)
}

func (gc *graphics) platformClipEvenOdd() {
	current := C.cairo_get_fill_rule(gc.gc)
	if current != C.CAIRO_FILL_RULE_EVEN_ODD {
		C.cairo_set_fill_rule(gc.gc, C.CAIRO_FILL_RULE_EVEN_ODD)
	}
	C.cairo_clip(gc.gc)
	if current != C.CAIRO_FILL_RULE_EVEN_ODD {
		C.cairo_set_fill_rule(gc.gc, current)
	}
}

func (gc *graphics) platformClipRect(bounds geom.Rect) {
	C.cairo_rectangle(gc.gc, C.double(bounds.X), C.double(bounds.Y), C.double(bounds.Width), C.double(bounds.Height))
	C.cairo_clip(gc.gc)
}

func (gc *graphics) platformDrawLinearGradient(gradient *Gradient, sx, sy, ex, ey float32) {
	pattern := C.cairo_pattern_create_linear(C.double(sx), C.double(sy), C.double(ex), C.double(ey))
	for _, one := range gradient.Stops {
		C.cairo_pattern_add_color_stop_rgba(pattern, C.double(one.Location), C.double(one.Color.RedIntensity()), C.double(one.Color.GreenIntensity()), C.double(one.Color.BlueIntensity()), C.double(one.Color.AlphaIntensity()))
	}
	C.cairo_set_source(gc.gc, pattern)
	C.cairo_paint(gc.gc)
	C.cairo_pattern_destroy(pattern)
}

func (gc *graphics) platformDrawRadialGradient(gradient *Gradient, scx, scy, startRadius, ecx, ecy, endRadius float32) {
	// RAW: Implement platformDrawRadialGradient for Linux
}

func (gc *graphics) platformDrawImageInRect(img *Image, bounds geom.Rect) {
	// RAW: Implement platformDrawImageInRect for Linux
}

func (gc *graphics) platformDrawString(x, y float32, str string) {
	C.cairo_move_to(gc.gc, C.double(x), C.double(y+14)) // RAW: Hack until actual font code is written
	C.cairo_show_text(gc.gc, C.CString(str))
}

func (gc *graphics) platformTranslate(x, y float32) {
	C.cairo_translate(gc.gc, C.double(x), C.double(y))
}

func (gc *graphics) platformScale(x, y float32) {
	C.cairo_scale(gc.gc, C.double(x), C.double(y))
}

func (gc *graphics) platformRotate(angleInRadians float32) {
	C.cairo_rotate(gc.gc, C.double(angleInRadians))
}
