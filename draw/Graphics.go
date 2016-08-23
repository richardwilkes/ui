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
	"github.com/richardwilkes/ui/font"
	"github.com/richardwilkes/ui/geom"
	"github.com/richardwilkes/xmath"
	"math"
)

// #cgo darwin LDFLAGS: -framework Cocoa -framework Quartz
// #cgo linux LDFLAGS: -lX11
// #cgo pkg-config: pangocairo
// #include <pango/pangocairo.h>
import "C"

const (
	unlimitedSize = 1000000
	oneThird      = 1.0 / 3.0
	twoThirds     = 2.0 / 3.0
)

type CairoContext *C.cairo_t

// Graphics provides a graphics context for drawing into.
type Graphics struct {
	gc    CairoContext
	stack []*graphicsState
}

type graphicsState struct {
	color       color.Color
	strokeWidth float32
	font        *font.Font
}

// NewGraphics creates a new Graphics object with a Cairo graphics context.
func NewGraphics(cc CairoContext) *Graphics {
	gc := &Graphics{gc: cc}
	gc.stack = append(gc.stack, &graphicsState{})
	gc.SetStrokeWidth(1)
	gc.SetColor(color.Black)
	return gc
}

// Dispose of the underlying Cairo graphics context
func (gc *Graphics) Dispose() {
	C.cairo_destroy(gc.gc)
	gc.gc = nil
	gc.stack = nil
}

// Save pushes a copy of the current graphics state onto the stack for the context.
// The current path is not saved as part of this call.
func (gc *Graphics) Save() {
	gs := *gc.stack[len(gc.stack)-1]
	gc.stack = append(gc.stack, &gs)
	C.cairo_save(gc.gc)
}

// Restore sets the current graphics state to the state most recently saved by call to Save().
func (gc *Graphics) Restore() {
	gc.stack[len(gc.stack)-1] = nil
	gc.stack = gc.stack[:len(gc.stack)-1]
	C.cairo_restore(gc.gc)
}

// Color returns the current color.
func (gc *Graphics) Color() color.Color {
	return gc.stack[len(gc.stack)-1].color
}

// SetColor sets the current color.
func (gc *Graphics) SetColor(color color.Color) {
	gc.stack[len(gc.stack)-1].color = color
	C.cairo_set_source_rgba(gc.gc, C.double(color.RedIntensity()), C.double(color.GreenIntensity()), C.double(color.BlueIntensity()), C.double(color.AlphaIntensity()))
}

// StrokeWidth returns the current stroke width.
func (gc *Graphics) StrokeWidth() float32 {
	return gc.stack[len(gc.stack)-1].strokeWidth
}

// SetStrokeWidth sets the current stroke width.
func (gc *Graphics) SetStrokeWidth(width float32) {
	if width > 0 {
		gc.stack[len(gc.stack)-1].strokeWidth = width
		C.cairo_set_line_width(gc.gc, C.double(width))
	}
}

// Font returns the current font.
func (gc *Graphics) Font() *font.Font {
	return gc.stack[len(gc.stack)-1].font
}

// SetFont sets the current font.
func (gc *Graphics) SetFont(font *font.Font) {
	gc.stack[len(gc.stack)-1].font = font
}

// StrokeLine draws a line between two points. To ensure the line is aligned to pixel boundaries,
// 0.5 is added to each coordinate.
func (gc *Graphics) StrokeLine(x1, y1, x2, y2 float32) {
	gc.BeginPath()
	gc.MoveTo(x1+0.5, y1+0.5)
	gc.LineTo(x2+0.5, y2+0.5)
	gc.StrokePath()
}

// FillRect fills a rectangle in the specified bounds.
func (gc *Graphics) FillRect(bounds geom.Rect) {
	C.cairo_rectangle(gc.gc, C.double(bounds.X), C.double(bounds.Y), C.double(bounds.Width), C.double(bounds.Height))
	C.cairo_fill(gc.gc)
}

// StrokeRect draws a rectangle in the specified bounds. To ensure the rectangle is aligned to pixel
// boundaries, 0.5 is added to the origin coordinates and 1 is subtracted from the size values.
func (gc *Graphics) StrokeRect(bounds geom.Rect) {
	if bounds.Width <= 1 || bounds.Height <= 1 {
		gc.FillRect(bounds)
	} else {
		bounds.InsetUniform(0.5)
		C.cairo_rectangle(gc.gc, C.double(bounds.X), C.double(bounds.Y), C.double(bounds.Width), C.double(bounds.Height))
		C.cairo_stroke(gc.gc)
	}
}

// FillEllipse fills an ellipse in the specified bounds.
func (gc *Graphics) FillEllipse(bounds geom.Rect) {
	gc.Save()
	cx := bounds.X + bounds.Width/2
	cy := bounds.Y + bounds.Height/2
	gc.Translate(cx, cy)
	if bounds.Width > bounds.Height {
		gc.Scale(1, bounds.Height/bounds.Width)
	} else {
		gc.Scale(bounds.Width/bounds.Height, 1)
	}
	gc.Translate(-cx, -cy)
	gc.BeginPath()
	gc.Arc(cx, cy, xmath.MaxFloat32(bounds.Width, bounds.Height)/2, 0, 2*math.Pi, true)
	gc.Restore()
	gc.FillPath()
}

// StrokeEllipse draws an ellipse in the specified bounds. To ensure the ellipse is aligned to pixel
// boundaries, 0.5 is added to the origin coordinates and 1 is subtracted from the size values.
func (gc *Graphics) StrokeEllipse(bounds geom.Rect) {
	if bounds.Width <= 1 || bounds.Height <= 1 {
		gc.FillEllipse(bounds)
	} else {
		bounds.InsetUniform(0.5)
		lineWidth := gc.StrokeWidth()
		gc.Save()
		cx := bounds.X + bounds.Width/2
		cy := bounds.Y + bounds.Height/2
		gc.Translate(cx, cy)
		if bounds.Width > bounds.Height {
			gc.Scale(1, bounds.Height/bounds.Width)
		} else {
			gc.Scale(bounds.Width/bounds.Height, 1)
		}
		gc.Translate(-cx, -cy)
		gc.BeginPath()
		gc.Arc(cx, cy, xmath.MaxFloat32(bounds.Width, bounds.Height)/2, 0, 2*math.Pi, true)
		gc.Restore()
		gc.SetStrokeWidth(lineWidth)
		gc.StrokePath()
	}
}

// FillPath fills the current path using the Non-Zero winding rule
// (https://en.wikipedia.org/wiki/Nonzero-rule), then clears the current path state.
func (gc *Graphics) FillPath() {
	C.cairo_fill(gc.gc)
}

// FillPathEvenOdd fills the current path using the Even-Odd winding rule
// (https://en.wikipedia.org/wiki/Even–odd_rule), then clears the current path state.
func (gc *Graphics) FillPathEvenOdd() {
	current := C.cairo_get_fill_rule(gc.gc)
	if current != C.CAIRO_FILL_RULE_EVEN_ODD {
		C.cairo_set_fill_rule(gc.gc, C.CAIRO_FILL_RULE_EVEN_ODD)
	}
	C.cairo_fill(gc.gc)
	if current != C.CAIRO_FILL_RULE_EVEN_ODD {
		C.cairo_set_fill_rule(gc.gc, current)
	}
}

// StrokePath strokes the current path, then clears the current path state.
func (gc *Graphics) StrokePath() {
	C.cairo_stroke(gc.gc)
}

// BeginPath creates a new empty path in the context.
func (gc *Graphics) BeginPath() {
	C.cairo_new_path(gc.gc)
}

// ClosePath closes and terminates the current path’s subpath.
func (gc *Graphics) ClosePath() {
	C.cairo_close_path(gc.gc)
}

// MoveTo begins a new subpath at the specified point.
func (gc *Graphics) MoveTo(x, y float32) {
	C.cairo_move_to(gc.gc, C.double(x), C.double(y))
}

// LineTo appends a straight line segment from the current point to the specified point.
func (gc *Graphics) LineTo(x, y float32) {
	C.cairo_line_to(gc.gc, C.double(x), C.double(y))
}

// Arc adds an arc of a circle to the current path, possibly preceded by a straight line segment.
func (gc *Graphics) Arc(cx, cy, radius, startAngleRadians, endAngleRadians float32, clockwise bool) {
	if clockwise {
		C.cairo_arc(gc.gc, C.double(cx), C.double(cy), C.double(radius), C.double(startAngleRadians), C.double(endAngleRadians))
	} else {
		C.cairo_arc_negative(gc.gc, C.double(cx), C.double(cy), C.double(radius), C.double(startAngleRadians), C.double(endAngleRadians))
	}
}

// CurveTo appends a cubic Bezier curve from the current point, using the provided controls
// points and end point.
func (gc *Graphics) CurveTo(cp1x, cp1y, cp2x, cp2y, x, y float32) {
	C.cairo_curve_to(gc.gc, C.double(cp1x), C.double(cp1y), C.double(cp2x), C.double(cp2y), C.double(x), C.double(y))
}

// QuadCurveTo appends a quadratic Bezier curve from the current point, using a control point
// and an end point.
func (gc *Graphics) QuadCurveTo(cpx, cpy, x, y float32) {
	var x0, y0 C.double
	C.cairo_get_current_point(gc.gc, &x0, &y0)
	gc.CurveTo(twoThirds*cpx+oneThird*float32(x0), twoThirds*cpy+oneThird*float32(y0), twoThirds*cpx+oneThird*x, twoThirds*cpy+oneThird*y, x, y)
}

// AddPath adds a previously created Path object to the current path.
func (gc *Graphics) AddPath(path *Path) {
	path.Apply(gc.gc)
}

// Clip sets the clipping path to the intersection of the current clipping path with the area
// defined by the current path using the Non-Zero winding rule
// (https://en.wikipedia.org/wiki/Nonzero-rule), then clears the current path state.
func (gc *Graphics) Clip() {
	C.cairo_clip(gc.gc)
}

// ClipEvenOdd sets the clipping path to the intersection of the current clipping path with
// the area defined by the current path using the Even-Odd winding rule
// (https://en.wikipedia.org/wiki/Even–odd_rule), then clears the current path state.
func (gc *Graphics) ClipEvenOdd() {
	current := C.cairo_get_fill_rule(gc.gc)
	if current != C.CAIRO_FILL_RULE_EVEN_ODD {
		C.cairo_set_fill_rule(gc.gc, C.CAIRO_FILL_RULE_EVEN_ODD)
	}
	C.cairo_clip(gc.gc)
	if current != C.CAIRO_FILL_RULE_EVEN_ODD {
		C.cairo_set_fill_rule(gc.gc, current)
	}
}

// ClipRect sets the clipping path to the intersection of the current clipping path with the
// area defined by the specified rectangle, then clears the current path state.
func (gc *Graphics) ClipRect(bounds geom.Rect) {
	C.cairo_rectangle(gc.gc, C.double(bounds.X), C.double(bounds.Y), C.double(bounds.Width), C.double(bounds.Height))
	C.cairo_clip(gc.gc)
}

// DrawLinearGradient from sx, sy to ex, ey.
func (gc *Graphics) DrawLinearGradient(gradient *Gradient, sx, sy, ex, ey float32) {
	pattern := C.cairo_pattern_create_linear(C.double(sx), C.double(sy), C.double(ex), C.double(ey))
	for _, one := range gradient.Stops {
		C.cairo_pattern_add_color_stop_rgba(pattern, C.double(one.Location), C.double(one.Color.RedIntensity()), C.double(one.Color.GreenIntensity()), C.double(one.Color.BlueIntensity()), C.double(one.Color.AlphaIntensity()))
	}
	C.cairo_set_source(gc.gc, pattern)
	C.cairo_paint(gc.gc)
	C.cairo_pattern_destroy(pattern)
}

// DrawRadialGradient from sx, sy to ex, ey.
func (gc *Graphics) DrawRadialGradient(gradient *Gradient, scx, scy, startRadius, ecx, ecy, endRadius float32) {
	pattern := C.cairo_pattern_create_radial(C.double(scx), C.double(scy), C.double(startRadius), C.double(ecx), C.double(ecy), C.double(endRadius))
	for _, one := range gradient.Stops {
		C.cairo_pattern_add_color_stop_rgba(pattern, C.double(one.Location), C.double(one.Color.RedIntensity()), C.double(one.Color.GreenIntensity()), C.double(one.Color.BlueIntensity()), C.double(one.Color.AlphaIntensity()))
	}
	C.cairo_set_source(gc.gc, pattern)
	C.cairo_paint(gc.gc)
	C.cairo_pattern_destroy(pattern)
}

// DragImage at the specified location.
func (gc *Graphics) DrawImage(img *Image, where geom.Point) {
	gc.DrawImageInRect(img, geom.Rect{Point: where, Size: img.Size()})
}

// DrawImageInRect draws the image in the bounds, scaling if necessary.
func (gc *Graphics) DrawImageInRect(img *Image, bounds geom.Rect) {
	gc.Save()
	gc.ClipRect(bounds)
	gc.Translate(bounds.X, bounds.Y)
	size := img.Size()
	gc.Scale(bounds.Width/size.Width, bounds.Height/size.Height)
	C.cairo_set_source_surface(gc.gc, img.PlatformPtr(), 0, 0)
	C.cairo_paint(gc.gc)
	gc.Restore()
}

// DrawString at the specified location using the current font and fill color.
func (gc *Graphics) DrawString(x, y float32, str string) {
	f := gc.Font()
	layout := C.pango_cairo_create_layout(gc.gc)
	C.pango_layout_set_font_description(layout, f.PangoFontDescription())
	C.pango_layout_set_spacing(layout, C.int(f.Leading()*font.PangoScale))
	cstr := C.CString(str)
	C.pango_layout_set_text(layout, cstr, -1)
	C.g_free(cstr)
	gc.MoveTo(x, y+f.Leading())
	C.pango_cairo_show_layout(gc.gc, layout)
	var inkRect, logicalRect C.PangoRectangle
	C.pango_layout_get_pixel_extents(layout, &inkRect, &logicalRect)
	C.g_object_unref(layout)
}

// Translate the coordinate system.
func (gc *Graphics) Translate(x, y float32) {
	C.cairo_translate(gc.gc, C.double(x), C.double(y))
}

// Scale the coordinate system.
func (gc *Graphics) Scale(x, y float32) {
	C.cairo_scale(gc.gc, C.double(x), C.double(y))
}

// Rotate the coordinate system.
func (gc *Graphics) Rotate(angleInRadians float32) {
	C.cairo_rotate(gc.gc, C.double(angleInRadians))
}
