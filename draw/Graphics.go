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
)

// Graphics provides a graphics context for drawing into.
type Graphics interface {
	// Save pushes a copy of the current graphics state onto the stack for the context.
	// The current path is not saved as part of this call.
	Save()
	// Restore sets the current graphics state to the state most recently saved by call to Save().
	Restore()
	// Opacity returns the opacity level used when drawing.
	Opacity() float32
	// SetOpacity sets the opacity level used when drawing. Values can range from 0 (transparent)
	// to 1 (opaque).
	SetOpacity(opacity float32)
	// FillColor returns the current fill color.
	FillColor() color.Color
	// SetFillColor sets the current fill color.
	SetFillColor(color color.Color)
	// StrokeColor returns the current stroke color.
	StrokeColor() color.Color
	// SetStrokeColor sets the current stroke color.
	SetStrokeColor(color color.Color)
	// StrokeWidth returns the current stroke width.
	StrokeWidth() float32
	// SetStrokeWidth sets the current stroke width.
	SetStrokeWidth(width float32)
	// Font returns the current font.
	Font() *font.Font
	// SetFont sets the current font.
	SetFont(font *font.Font)
	// StrokeLine draws a line between two points.
	StrokeLine(x1, y1, x2, y2 float32)
	// FillRect fills a rectangle in the specified bounds.
	FillRect(bounds geom.Rect)
	// StrokeRect draws a rectangle in the specified bounds.
	StrokeRect(bounds geom.Rect)
	// FillEllipse fills an ellipse in the specified bounds.
	FillEllipse(bounds geom.Rect)
	// StrokeEllipse draws an ellipse in the specified bounds.
	StrokeEllipse(bounds geom.Rect)
	// FillPath fills the current path using the Non-Zero winding rule
	// (https://en.wikipedia.org/wiki/Nonzero-rule), then clears the current path state.
	FillPath()
	// FillPathEvenOdd fills the current path using the Even-Odd winding rule
	// (https://en.wikipedia.org/wiki/Even–odd_rule), then clears the current path state.
	FillPathEvenOdd()
	// StrokePath strokes the current path, then clears the current path state.
	StrokePath()
	// FillAndStrokePath fills and strokes the current path, then clears the current path state.
	FillAndStrokePath()
	// BeginPath creates a new empty path in the context.
	BeginPath()
	// ClosePath closes and terminates the current path’s subpath.
	ClosePath()
	// MoveTo begins a new subpath at the specified point.
	MoveTo(x, y float32)
	// LineTo appends a straight line segment from the current point to the specified point.
	LineTo(x, y float32)
	// Arc adds an arc of a circle to the current path, possibly preceded by a straight line
	// segment.
	Arc(cx, cy, radius, startAngleRadians, endAngleRadians float32, clockwise bool)
	// ArcTo adds an arc of a circle to the current path, using a radius and tangent points.
	ArcTo(x1, y1, x2, y2, radius float32)
	// CurveTo appends a cubic Bezier curve from the current point, using the provided controls
	// points and end point.
	CurveTo(cp1x, cp1y, cp2x, cp2y, x, y float32)
	// QuadCurveTo appends a quadratic Bezier curve from the current point, using a control point
	// and an end point.
	QuadCurveTo(cpx, cpy, x, y float32)
	// AddPath adds a previously created Path object to the current path.
	AddPath(path *geom.Path)
	// Clip sets the clipping path to the intersection of the current clipping path with the area
	// defined by the current path using the Non-Zero winding rule
	// (https://en.wikipedia.org/wiki/Nonzero-rule), then clears the current path state.
	Clip()
	// ClipEvenOdd sets the clipping path to the intersection of the current clipping path with
	// the area defined by the current path using the Even-Odd winding rule
	// (https://en.wikipedia.org/wiki/Even–odd_rule), then clears the current path state.
	ClipEvenOdd()
	// ClipRect sets the clipping path to the intersection of the current clipping path with the
	// area defined by the specified rectangle, then clears the current path state.
	ClipRect(bounds geom.Rect)
	// DrawLinearGradient from sx, sy to ex, ey.
	DrawLinearGradient(gradient *Gradient, sx, sy, ex, ey float32)
	// DrawRadialGradient from sx, sy to ex, ey.
	DrawRadialGradient(gradient *Gradient, scx, scy, startRadius, ecx, ecy, endRadius float32)
	// DragImage at the specified location.
	DrawImage(img *Image, where geom.Point)
	// DragImage in the bounds, scaling if necessary.
	DrawImageInRect(img *Image, bounds geom.Rect)
	// DrawText at the specified location. Returns the space taken up by the text.
	DrawText(x, y float32, str string, mode TextMode) geom.Size
	// DrawTextConstrained at the specified location. The text will be wrapped to fit within the
	// specified bounds. Returns the space taken up by the text and the number of bytes of the text
	// that were drawn.
	DrawTextConstrained(bounds geom.Rect, str string, mode TextMode) (actual geom.Size, fit int)
	// DrawAttributedText at the specified location. Returns the space taken up by the text.
	DrawAttributedText(x, y float32, str *Text, mode TextMode) geom.Size
	// DrawAttributedTextConstrained at the specified location. The text will be wrapped to fit
	// within the specified bounds. Returns the space taken up by the text and the number of bytes
	// of the text that were drawn.
	DrawAttributedTextConstrained(bounds geom.Rect, str *Text, mode TextMode) (actual geom.Size, fit int)
	// Translate the coordinate system.
	Translate(x, y float32)
	// Scale the coordinate system.
	Scale(x, y float32)
	// Rotate the coordinate system.
	Rotate(angleInRadians float32)
}
