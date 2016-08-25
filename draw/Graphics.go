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
	"github.com/richardwilkes/ui/draw/compositing"
	"github.com/richardwilkes/ui/font"
	"math"
	// #cgo darwin LDFLAGS: -framework Cocoa -framework Quartz
	// #cgo linux LDFLAGS: -lX11
	// #cgo pkg-config: pangocairo
	// #include <pango/pangocairo.h>
	"C"
)

const (
	unlimitedSize = 1000000
	oneThird      = 1.0 / 3.0
	twoThirds     = 2.0 / 3.0
)

const (
	// AntiAliasKindDefault indicates that the default anti-aliasing for the graphics subsystem and
	// target device should be used.
	AntiAliasKindDefault AntiAliasKind = C.CAIRO_ANTIALIAS_DEFAULT
	// AntiAliasKindNone indicates no anti-aliasing should be used.
	AntiAliasKindNone AntiAliasKind = C.CAIRO_ANTIALIAS_NONE
	// AntiAliasKindGray indicates a single-color anti-aliasing should be used.
	AntiAliasKindGray AntiAliasKind = C.CAIRO_ANTIALIAS_GRAY
	// AntiAliasKindSubPixel indicates anti-aliasing should be perfomed by taking advantage of the
	// order of sub-pixel elements on devices such as LCD panels.
	AntiAliasKindSubPixel AntiAliasKind = C.CAIRO_ANTIALIAS_SUBPIXEL
	// AntiAliasKindFast indicates that some anti-aliasing should be performed, but speed should be
	// preferred over quality.
	AntiAliasKindFast AntiAliasKind = C.CAIRO_ANTIALIAS_FAST
	// AntiAliasKindGood indicates that the backend should balance quality against performance.
	AntiAliasKindGood AntiAliasKind = C.CAIRO_ANTIALIAS_GOOD
	// AntiAliasKindBest indicates that the backend should render at the highest quality,
	// sacrificing speed if necessary.
	AntiAliasKindBest AntiAliasKind = C.CAIRO_ANTIALIAS_BEST
)

// AntiAliasKind holds the type of anti-aliasing that is being performed.
type AntiAliasKind int

const (
	// FillRuleWinding uses this algorithm: if the path crosses the ray from left-to-right, counts
	// +1. If the path crosses the ray from right to left, counts -1. Left and right are determined
	// from the perspective of looking along the ray from the starting point. If the total count
	// is non-zero, the point will be filled.
	FillRuleWinding FillRule = C.CAIRO_FILL_RULE_WINDING
	// FillRuleEvenOdd uses this algorithm: counts the total number of intersections, without regard
	// to the orientation of the contour. If the total number of intersections is odd, the point
	// will be filled.
	FillRuleEvenOdd FillRule = C.CAIRO_FILL_RULE_EVEN_ODD
)

// FillRule holds the type of fill rule to use when filling paths.
type FillRule int

// Possible line caps.
const (
	LineCapButt   LineCap = C.CAIRO_LINE_CAP_BUTT
	LineCapRound  LineCap = C.CAIRO_LINE_CAP_ROUND
	LineCapSquare LineCap = C.CAIRO_LINE_CAP_SQUARE
)

// LineCap indicates how to render the endpoints of a path when stroking it.
type LineCap int

// Possible line joins.
const (
	LineJoinMiter LineJoin = C.CAIRO_LINE_JOIN_MITER
	LineJoinRound LineJoin = C.CAIRO_LINE_JOIN_ROUND
	LineJoinBevel LineJoin = C.CAIRO_LINE_JOIN_BEVEL
)

// LineJoin indicates how to join the endpoints of a path when stroking it.
type LineJoin int

type CairoContext *C.cairo_t

// Graphics provides a graphics context for drawing into.
type Graphics struct {
	gc CairoContext
}

// NewGraphics creates a new Graphics object with a Cairo graphics context.
func NewGraphics(cc CairoContext) *Graphics {
	gc := &Graphics{gc: cc}
	gc.SetStrokeWidth(1)
	gc.SetColor(color.Black)
	return gc
}

// Dispose of the underlying Cairo graphics context
func (gc *Graphics) Dispose() {
	C.cairo_destroy(gc.gc)
	gc.gc = nil
}

// Save pushes a copy of the current graphics state onto the stack for the context.
// The current path is not saved as part of this call.
func (gc *Graphics) Save() {
	C.cairo_save(gc.gc)
}

// Restore sets the current graphics state to the state most recently saved by call to Save().
func (gc *Graphics) Restore() {
	C.cairo_restore(gc.gc)
}

// Translate the coordinate system.
func (gc *Graphics) Translate(x, y float64) {
	C.cairo_translate(gc.gc, C.double(x), C.double(y))
}

// Scale the coordinate system.
func (gc *Graphics) Scale(x, y float64) {
	C.cairo_scale(gc.gc, C.double(x), C.double(y))
}

// Rotate the coordinate system.
func (gc *Graphics) Rotate(angleInRadians float64) {
	C.cairo_rotate(gc.gc, C.double(angleInRadians))
}

// Transform the coordinate system by the matrix.
func (gc *Graphics) Transform(matrix *geom.Matrix) {
	C.cairo_transform(gc.gc, toCairoMatrix(matrix))
}

// Matrix returns the current transformation matrix.
func (gc *Graphics) Matrix() *geom.Matrix {
	var matrix C.cairo_matrix_t
	C.cairo_get_matrix(gc.gc, &matrix)
	return fromCairoMatrix(&matrix)
}

// SetMatrix sets the current transformation matrix.
func (gc *Graphics) SetMatrix(matrix *geom.Matrix) {
	C.cairo_set_matrix(gc.gc, toCairoMatrix(matrix))
}

// SetIdentityMatrix sets the current transformation matrix to the identity matrix.
func (gc *Graphics) SetIdentityMatrix() {
	C.cairo_identity_matrix(gc.gc)
}

// UserToDevice converts the user-space coordinates to device-space.
func (gc *Graphics) UserToDevice(ux, uy float64) (x, y float64) {
	dx := C.double(ux)
	dy := C.double(uy)
	C.cairo_user_to_device(gc.gc, &dx, &dy)
	return float64(dx), float64(dy)
}

// UserToDeviceDistance converts the user-space distance to device-space.
func (gc *Graphics) UserToDeviceDistance(uw, uh float64) (w, h float64) {
	dw := C.double(uw)
	dh := C.double(uh)
	C.cairo_user_to_device_distance(gc.gc, &dw, &dh)
	return float64(dw), float64(dh)
}

// DeviceToUser converts the device-space coordinates to user-space.
func (gc *Graphics) DeviceToUser(ux, uy float64) (x, y float64) {
	dx := C.double(ux)
	dy := C.double(uy)
	C.cairo_device_to_user(gc.gc, &dx, &dy)
	return float64(dx), float64(dy)
}

// DeviceToUserDistance converts the device-space distance to user-space.
func (gc *Graphics) DeviceToUserDistance(uw, uh float64) (w, h float64) {
	dw := C.double(uw)
	dh := C.double(uh)
	C.cairo_device_to_user_distance(gc.gc, &dw, &dh)
	return float64(dw), float64(dh)
}

// BeginPath creates a new empty path in the context.
func (gc *Graphics) BeginPath() {
	C.cairo_new_path(gc.gc)
}

// AddPath adds a previously created Path object to the current path.
func (gc *Graphics) AddPath(path *Path) {
	for _, node := range path.data {
		switch t := node.(type) {
		case *moveToPathNode:
			gc.MoveTo(t.x, t.y)
		case *lineToPathNode:
			gc.LineTo(t.x, t.y)
		case *arcPathNode:
			gc.Arc(t.cx, t.cy, t.radius, t.startAngleRadians, t.endAngleRadians, t.clockwise)
		case *curveToPathNode:
			gc.CurveTo(t.cp1x, t.cp1y, t.cp2x, t.cp2y, t.x, t.y)
		case *quadCurveToPathNode:
			gc.QuadCurveTo(t.cpx, t.cpy, t.x, t.y)
		case *rectPathNode:
			gc.Rect(t.bounds)
		case *ellipsePathNode:
			gc.Ellipse(t.bounds)
		case *subPathNode:
			gc.BeginSubPath()
		case *closePathNode:
			gc.ClosePath()
		default:
			panic("Unknown path node type")
		}
	}
}

// BeginSubPath creates a new empty sub-path in the context. Any existing path is not affected.
func (gc *Graphics) BeginSubPath() {
	C.cairo_new_sub_path(gc.gc)
}

// MoveTo begins a new subpath at the specified point.
func (gc *Graphics) MoveTo(x, y float64) {
	C.cairo_move_to(gc.gc, C.double(x), C.double(y))
}

// LineTo appends a straight line segment from the current point to the specified point.
func (gc *Graphics) LineTo(x, y float64) {
	C.cairo_line_to(gc.gc, C.double(x), C.double(y))
}

// Arc adds an arc of a circle to the current path, possibly preceded by a straight line segment.
func (gc *Graphics) Arc(cx, cy, radius, startAngleRadians, endAngleRadians float64, clockwise bool) {
	if clockwise {
		C.cairo_arc(gc.gc, C.double(cx), C.double(cy), C.double(radius), C.double(startAngleRadians), C.double(endAngleRadians))
	} else {
		C.cairo_arc_negative(gc.gc, C.double(cx), C.double(cy), C.double(radius), C.double(startAngleRadians), C.double(endAngleRadians))
	}
}

// CurveTo appends a cubic Bezier curve from the current point, using the provided controls
// points and end point.
func (gc *Graphics) CurveTo(cp1x, cp1y, cp2x, cp2y, x, y float64) {
	C.cairo_curve_to(gc.gc, C.double(cp1x), C.double(cp1y), C.double(cp2x), C.double(cp2y), C.double(x), C.double(y))
}

// QuadCurveTo appends a quadratic Bezier curve from the current point, using a control point
// and an end point.
func (gc *Graphics) QuadCurveTo(cpx, cpy, x, y float64) {
	var x0, y0 C.double
	C.cairo_get_current_point(gc.gc, &x0, &y0)
	gc.CurveTo(twoThirds*cpx+oneThird*float64(x0), twoThirds*cpy+oneThird*float64(y0), twoThirds*cpx+oneThird*x, twoThirds*cpy+oneThird*y, x, y)
}

// Rect appends a rectangle to the current path.
func (gc *Graphics) Rect(bounds geom.Rect) {
	C.cairo_rectangle(gc.gc, C.double(bounds.X), C.double(bounds.Y), C.double(bounds.Width), C.double(bounds.Height))
}

// Ellipse appends an ellipse to the current path.
func (gc *Graphics) Ellipse(bounds geom.Rect) {
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
	gc.Arc(cx, cy, math.Max(bounds.Width, bounds.Height)/2, 0, 2*math.Pi, true)
	gc.Restore()
}

// ClosePath closes and terminates the current path’s subpath.
func (gc *Graphics) ClosePath() {
	C.cairo_close_path(gc.gc)
}

// PushGroup temporarily redirects drawing to an intermediate surface known as a group. The
// redirection lasts until the group is completed by a call to PopGroup() or PopGroupToSource().
// Groups can be nested arbitrarily deep.
func (gc *Graphics) PushGroup() {
	C.cairo_push_group(gc.gc)
}

// PopGroup terminates the redirection begun by a call to PushGroup and returns a new Paint
// containing the results of all drawing operations performed to the group.
func (gc *Graphics) PopGroup() Paint {
	return Paint{pattern: C.cairo_pop_group(gc.gc)}
}

// PopGroupToPaint terminates the redirection begun by a call to PushGroup and installs the
// resulting Paint as the current source paint.
func (gc *Graphics) PopGroupToPaint() {
	C.cairo_pop_group_to_source(gc.gc)
}

// Paint returns a copy of the current paint. You are responsible for calling its Dispose() method.
func (gc *Graphics) Paint() Paint {
	pattern := C.cairo_get_source(gc.gc)
	C.cairo_pattern_reference(pattern)
	return Paint{pattern: pattern}
}

// SetPaint sets the current paint.
func (gc *Graphics) SetPaint(paint Paint) {
	C.cairo_set_source(gc.gc, paint.pattern)
}

// SetColor sets the current paint with a color.
func (gc *Graphics) SetColor(color color.Color) {
	C.cairo_set_source_rgba(gc.gc, C.double(color.RedIntensity()), C.double(color.GreenIntensity()), C.double(color.BlueIntensity()), C.double(color.AlphaIntensity()))
}

// SetImage sets the current paint with a tiled image. 'dx' and 'dy' specify the amount to offset
// the image pattern.
func (gc *Graphics) SetImage(img *Image, dx, dy float64) {
	C.cairo_set_source_surface(gc.gc, img.surface, C.double(dx), C.double(dy))
	C.cairo_pattern_set_extend(C.cairo_get_source(gc.gc), C.CAIRO_EXTEND_REPEAT)
}

// AntiAlias returns the current anti-aliasing mode.
func (gc *Graphics) AntiAlias() AntiAliasKind {
	return AntiAliasKind(C.cairo_get_antialias(gc.gc))
}

// SetAntiAlias sets the anti-aliasing mode of the rasterizer used for drawing shapes. This value is
// a hint and a particular backend may or may not support a particular value. Note that this option
// does not affect text rendering.
func (gc *Graphics) SetAntiAlias(kind AntiAliasKind) {
	C.cairo_set_antialias(gc.gc, C.cairo_antialias_t(kind))
}

// StrokeWidth returns the current stroke width.
func (gc *Graphics) StrokeWidth() float64 {
	return float64(C.cairo_get_line_width(gc.gc))
}

// SetStrokeWidth sets the current stroke width.
func (gc *Graphics) SetStrokeWidth(width float64) {
	C.cairo_set_line_width(gc.gc, C.double(width))
}

// Dash returns the current dash settings.
func (gc *Graphics) Dash() (segments []float64, offset float64) {
	count := C.cairo_get_dash_count(gc.gc)
	if count == 0 {
		return nil, 0
	}
	segments = make([]float64, count)
	C.cairo_get_dash(gc.gc, (*C.double)(&segments[0]), (*C.double)(&offset))
	return
}

// SetDash sets the dash pattern to be used when stroking. A dash pattern is specified by an array
// of positive values. Each value provides the length of alternate "on" and "off" portions of the
// stroke. The offset specifies an offset into the pattern at which the stroke begins. Each "on"
// segment will have caps applied as if the segment were a separate sub-path. In particular, it is
// valid to use an "on" length of 0 with LineCapRound or LineCapSquare in order to distribute dots
// or squares along a path. Note: The length values are in user-space units as evaluated at the time
// of stroking. This is not necessarily the same as the user space at the time of SetDash(). If
// 'segments' is nil, dashing is disabled. If 'segments' only contains one value, a symmetric
// pattern is assumed with alternating on and off portions of the size specified by the single
// value. If any value in 'segments' is negative, or if all values are 0, this call will be ignored.
func (gc *Graphics) SetDash(segments []float64, offset float64) {
	if segments == nil {
		C.cairo_set_dash(gc.gc, nil, 0, 0)
	} else {
		var hasNonZero bool
		for one := range segments {
			if one < 0 {
				return
			}
			if one > 0 {
				hasNonZero = true
			}
		}
		if hasNonZero {
			C.cairo_set_dash(gc.gc, (*C.double)(&segments[0]), C.int(len(segments)), C.double(offset))
		}
	}
}

// FillRule returns the current FillRule.
func (gc *Graphics) FillRule() FillRule {
	return FillRule(C.cairo_get_fill_rule(gc.gc))
}

// SetFillRule sets the current FillRule.
func (gc *Graphics) SetFillRule(rule FillRule) {
	C.cairo_set_fill_rule(gc.gc, C.cairo_fill_rule_t(rule))
}

// LineCap returns the current LineCap.
func (gc *Graphics) LineCap() LineCap {
	return LineCap(C.cairo_get_line_cap(gc.gc))
}

// SetLineCap sets the current LineCap.
func (gc *Graphics) SetLineCap(cap LineCap) {
	C.cairo_set_line_cap(gc.gc, C.cairo_line_cap_t(cap))
}

// LineJoin returns the current LineJoin.
func (gc *Graphics) LineJoin() LineJoin {
	return LineJoin(C.cairo_get_line_join(gc.gc))
}

// SetLineJoin sets the current LineJoin.
func (gc *Graphics) SetLineJoin(cap LineJoin) {
	C.cairo_set_line_join(gc.gc, C.cairo_line_join_t(cap))
}

// MiterLimit returns the current miter limit.
func (gc *Graphics) MiterLimit() float64 {
	return float64(C.cairo_get_miter_limit(gc.gc))
}

// SetMiterLimit sets the current miter limit.
func (gc *Graphics) SetMiterLimit(limit float64) {
	C.cairo_set_miter_limit(gc.gc, C.double(limit))
}

// CompositingOperator returns the current compositing operator.
func (gc *Graphics) CompositingOperator() compositing.Op {
	return compositing.Op(C.cairo_get_operator(gc.gc))
}

// SetCompositingOperator sets the current compositing operator.
func (gc *Graphics) SetCompositingOperator(op compositing.Op) {
	C.cairo_set_operator(gc.gc, C.cairo_operator_t(op))
}

// Clip establishes a new clip region by intersecting the current clip region with the current path
// as it would be filled by Fill() and according to the current FillRule. After Clip(), the current
// path will be cleared from the context. The current clip region affects all drawing operations by
// effectively masking out any changes to the surface that are outside the current clip region.
// Calling Clip() can only make the clip region smaller, never larger. The current clip is part of
// the graphics state, so a temporary restriction of the clip region can be achieved by calling
// Clip() within a Save()/Restore() pair. The only other means of increasing the size of the clip
// region is ResetClip().
func (gc *Graphics) Clip() {
	C.cairo_clip(gc.gc)
}

// ClipPreserve operates as Clip(), but does not clear the path from the context.
func (gc *Graphics) ClipPreserve() {
	C.cairo_clip_preserve(gc.gc)
}

// InClip returns true if the specific location would be visible through the current clip.
func (gc *Graphics) InClip(x, y float64) bool {
	return C.cairo_in_clip(gc.gc, C.double(x), C.double(y)) != 0
}

// ResetClip resets the current clip region to its original, unrestricted state. That is, set the
// clip region to an infinitely large shape containing the target surface. Note that code meant to
// be reusable should not call ResetClip() as it will cause results unexpected by higher-level code
// which calls Clip(). Consider using Save() and Restore() around Clip() as a more robust means of
// temporarily restricting the clip region.
func (gc *Graphics) ResetClip() {
	C.cairo_reset_clip(gc.gc)
}

// FillPath fills the current path, then clears the current path state.
func (gc *Graphics) FillPath() {
	C.cairo_fill(gc.gc)
}

// FillPathPreserve fills the current path without clearing the current path state.
func (gc *Graphics) FillPathPreserve() {
	C.cairo_fill_preserve(gc.gc)
}

// InFill returns true if the specific location would be inside the area affected by FillPath() or
// FillPathPreserve(). Note that the clip area is not taken into account.
func (gc *Graphics) InFill(x, y float64) bool {
	return C.cairo_in_fill(gc.gc, C.double(x), C.double(y)) != 0
}

// MaskClipWithPaint uses the current Paint to cover the entire clip region, but uses the alpha
// channel of 'mask' as a mask (i.e. opaque areas of 'mask' are painted with the current Paint,
// transparent areas are not painted).
func (gc *Graphics) MaskClipWithPaint(mask Paint) {
	C.cairo_mask(gc.gc, mask.pattern)
}

// MaskClipWithImage uses the current Paint to cover the entire clip region, but uses the alpha
// channel of 'mask' as a mask (i.e. opaque areas of 'mask' are painted with the current Paint,
// transparent areas are not painted). The image's mask will be offset by 'dx' and 'dy'.
func (gc *Graphics) MaskClipWithImage(mask *Image, dx, dy float64) {
	C.cairo_mask_surface(gc.gc, mask.surface, C.double(dx), C.double(dy))
}

// FillClip uses the current Paint to fill the entire clip region.
func (gc *Graphics) FillClip() {
	C.cairo_paint(gc.gc)
}

// FillClipWithAlpha uses the current Paint to cover the entire clip region, but with the specified
// 'alpha' value applied.
func (gc *Graphics) FillClipWithAlpha(alpha float64) {
	C.cairo_paint_with_alpha(gc.gc, C.double(alpha))
}

// StrokePath strokes the current path, then clears the current path state.
func (gc *Graphics) StrokePath() {
	C.cairo_stroke(gc.gc)
}

// StrokePathPreserve strokes the current path without clearing the current path state.
func (gc *Graphics) StrokePathPreserve() {
	C.cairo_stroke_preserve(gc.gc)
}

// InStroke returns true if the specific location would be inside the area affected by StrokePath()
// or StrokePathPreserve(). Note that the clip area is not taken into account.
func (gc *Graphics) InStroke(x, y float64) bool {
	return C.cairo_in_stroke(gc.gc, C.double(x), C.double(y)) != 0
}

// StrokeLine draws a line between two points. To ensure the line is aligned to pixel boundaries,
// 0.5 is added to each coordinate.
func (gc *Graphics) StrokeLine(x1, y1, x2, y2 float64) {
	gc.BeginPath()
	gc.MoveTo(x1+0.5, y1+0.5)
	gc.LineTo(x2+0.5, y2+0.5)
	gc.StrokePath()
}

// StrokeRect draws a rectangle in the specified bounds. To ensure the rectangle is aligned to pixel
// boundaries, 0.5 is added to the origin coordinates and 1 is subtracted from the size values.
func (gc *Graphics) StrokeRect(bounds geom.Rect) {
	if bounds.Width <= 1 || bounds.Height <= 1 {
		gc.FillRect(bounds)
	} else {
		bounds.InsetUniform(0.5)
		gc.Rect(bounds)
		gc.StrokePath()
	}
}

// StrokeEllipse draws an ellipse in the specified bounds. To ensure the ellipse is aligned to pixel
// boundaries, 0.5 is added to the origin coordinates and 1 is subtracted from the size values.
func (gc *Graphics) StrokeEllipse(bounds geom.Rect) {
	if bounds.Width <= 1 || bounds.Height <= 1 {
		gc.FillEllipse(bounds)
	} else {
		bounds.InsetUniform(0.5)
		lineWidth := gc.StrokeWidth()
		gc.Ellipse(bounds)
		gc.SetStrokeWidth(lineWidth)
		gc.StrokePath()
	}
}

// FillRect fills a rectangle in the specified bounds.
func (gc *Graphics) FillRect(bounds geom.Rect) {
	gc.Rect(bounds)
	gc.FillPath()
}

// FillEllipse fills an ellipse in the specified bounds.
func (gc *Graphics) FillEllipse(bounds geom.Rect) {
	gc.Ellipse(bounds)
	gc.FillPath()
}

// DragImage at the specified location.
func (gc *Graphics) DrawImage(img *Image, where geom.Point) {
	gc.DrawImageInRect(img, geom.Rect{Point: where, Size: img.Size()})
}

// DrawImageInRect draws the image in the bounds, scaling if necessary.
func (gc *Graphics) DrawImageInRect(img *Image, bounds geom.Rect) {
	gc.Save()
	gc.Rect(bounds)
	gc.Clip()
	gc.Translate(bounds.X, bounds.Y)
	size := img.Size()
	gc.Scale(bounds.Width/size.Width, bounds.Height/size.Height)
	C.cairo_set_source_surface(gc.gc, img.surface, 0, 0)
	gc.FillClip()
	gc.Restore()
}

// DrawString at the specified location using the current font and fill color.
func (gc *Graphics) DrawString(x, y float64, str string, f *font.Font) {
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
