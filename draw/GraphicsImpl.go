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
	"unsafe"
)

const (
	unlimitedSize = 1000000
)

type graphics struct {
	gc    unsafe.Pointer
	stack []*graphicsState
}

type graphicsState struct {
	opacity     float32
	fillColor   color.Color
	strokeColor color.Color
	strokeWidth float32
	font        *font.Font
}

// NewGraphics creates a new Graphics object with the platform-specific graphics context.
func NewGraphics(gc unsafe.Pointer) Graphics {
	c := &graphics{gc: gc}
	c.stack = append(c.stack, &graphicsState{opacity: 1, fillColor: color.White, strokeColor: color.Black, strokeWidth: 1})
	return c
}

// Save implements Graphics.
func (gc *graphics) Save() {
	gs := *gc.stack[len(gc.stack)-1]
	gc.stack = append(gc.stack, &gs)
	gc.platformSave()
}

// Restore implements Graphics.
func (gc *graphics) Restore() {
	gc.stack[len(gc.stack)-1] = nil
	gc.stack = gc.stack[:len(gc.stack)-1]
	gc.platformRestore()
}

// Opacity implements Graphics.
func (gc *graphics) Opacity() float32 {
	return gc.stack[len(gc.stack)-1].opacity
}

// SetOpacity implements Graphics.
func (gc *graphics) SetOpacity(opacity float32) {
	if opacity < 0 {
		opacity = 0
	} else if opacity > 1 {
		opacity = 1
	}
	gc.stack[len(gc.stack)-1].opacity = opacity
	gc.platformSetOpacity(opacity)
}

// FillColor implements Graphics.
func (gc *graphics) FillColor() color.Color {
	return gc.stack[len(gc.stack)-1].fillColor
}

// SetFillColor implements Graphics.
func (gc *graphics) SetFillColor(color color.Color) {
	gc.stack[len(gc.stack)-1].fillColor = color
	gc.platformSetFillColor(color)
}

// StrokeColor implements Graphics.
func (gc *graphics) StrokeColor() color.Color {
	return gc.stack[len(gc.stack)-1].strokeColor
}

// SetStrokeColor implements Graphics.
func (gc *graphics) SetStrokeColor(color color.Color) {
	gc.stack[len(gc.stack)-1].strokeColor = color
	gc.platformSetStrokeColor(color)
}

// StrokeWidth implements Graphics.
func (gc *graphics) StrokeWidth() float32 {
	return gc.stack[len(gc.stack)-1].strokeWidth
}

// SetStrokeWidth implements Graphics.
func (gc *graphics) SetStrokeWidth(width float32) {
	if width > 0 {
		gc.stack[len(gc.stack)-1].strokeWidth = width
		gc.platformSetStrokeWidth(width)
	}
}

// Font implements Graphics.
func (gc *graphics) Font() *font.Font {
	return gc.stack[len(gc.stack)-1].font
}

// SetFont implements Graphics.
func (gc *graphics) SetFont(font *font.Font) {
	gc.stack[len(gc.stack)-1].font = font
}

// StrokeLine implements Graphics.
// To ensure the line is aligned to pixel boundaries, 0.5 is added to each coordinate.
func (gc *graphics) StrokeLine(x1, y1, x2, y2 float32) {
	gc.BeginPath()
	gc.MoveTo(x1+0.5, y1+0.5)
	gc.LineTo(x2+0.5, y2+0.5)
	gc.StrokePath()
}

// FillRect implements Graphics.
func (gc *graphics) FillRect(bounds geom.Rect) {
	gc.platformFillRect(bounds)
}

// StrokeRect implements Graphics.
// To ensure the rectangle is aligned to pixel boundaries, 0.5 is added to the origin coordinates
// and 1 is subtracted from the size values.
func (gc *graphics) StrokeRect(bounds geom.Rect) {
	if bounds.Width <= 1 || bounds.Height <= 1 {
		savedColor := gc.FillColor()
		gc.SetFillColor(gc.StrokeColor())
		gc.FillRect(bounds)
		gc.SetFillColor(savedColor)
	} else {
		bounds.InsetUniform(0.5)
		gc.platformStrokeRect(bounds)
	}
}

// FillEllipse implements Graphics.
func (gc *graphics) FillEllipse(bounds geom.Rect) {
	gc.platformFillEllipse(bounds)
}

// StrokeEllipse implements Graphics.
// To ensure the ellipse is aligned to pixel boundaries, 0.5 is added to the origin coordinates
// and 1 is subtracted from the size values.
func (gc *graphics) StrokeEllipse(bounds geom.Rect) {
	if bounds.Width <= 1 || bounds.Height <= 1 {
		savedColor := gc.FillColor()
		gc.SetFillColor(gc.StrokeColor())
		gc.FillEllipse(bounds)
		gc.SetFillColor(savedColor)
	} else {
		bounds.InsetUniform(0.5)
		gc.platformStrokeEllipse(bounds)
	}
}

// FillPath implements Graphics.
func (gc *graphics) FillPath() {
	gc.platformFillPath()
}

// FillPathEvenOdd implements Graphics.
func (gc *graphics) FillPathEvenOdd() {
	gc.platformFillPathEvenOdd()
}

// StrokePath implements Graphics.
func (gc *graphics) StrokePath() {
	gc.platformStrokePath()
}

// FillAndStrokePath implements Graphics.
func (gc *graphics) FillAndStrokePath() {
	gc.platformFillAndStrokePath()
}

// BeginPath implements Graphics.
func (gc *graphics) BeginPath() {
	gc.platformBeginPath()
}

// ClosePath implements Graphics.
func (gc *graphics) ClosePath() {
	gc.platformClosePath()
}

// MoveTo implements Graphics.
func (gc *graphics) MoveTo(x, y float32) {
	gc.platformMoveTo(x, y)
}

// LineTo implements Graphics.
func (gc *graphics) LineTo(x, y float32) {
	gc.platformLineTo(x, y)
}

// Arc implements Graphics.
func (gc *graphics) Arc(cx, cy, radius, startAngleRadians, endAngleRadians float32, clockwise bool) {
	gc.platformArc(cx, cy, radius, startAngleRadians, endAngleRadians, clockwise)
}

// ArcTo implements Graphics.
func (gc *graphics) ArcTo(x1, y1, x2, y2, radius float32) {
	gc.platformArcTo(x1, y1, x2, y2, radius)
}

// CurveTo implements Graphics.
func (gc *graphics) CurveTo(cp1x, cp1y, cp2x, cp2y, x, y float32) {
	gc.platformCurveTo(cp1x, cp1y, cp2x, cp2y, x, y)
}

// QuadCurveTo implements Graphics.
func (gc *graphics) QuadCurveTo(cpx, cpy, x, y float32) {
	gc.platformQuadCurveTo(cpx, cpy, x, y)
}

// AddPath implements Graphics.
func (gc *graphics) AddPath(path *geom.Path) {
	gc.platformAddPath(path)
}

// Clip implements Graphics.
func (gc *graphics) Clip() {
	gc.platformClip()
}

// ClipEvenOdd implements Graphics.
func (gc *graphics) ClipEvenOdd() {
	gc.platformClipEvenOdd()
}

// ClipRect implements Graphics.
func (gc *graphics) ClipRect(bounds geom.Rect) {
	gc.platformClipRect(bounds)
}

// DrawLinearGradient implements Graphics.
func (gc *graphics) DrawLinearGradient(gradient *Gradient, sx, sy, ex, ey float32) {
	gc.platformDrawLinearGradient(gradient, sx, sy, ex, ey)
}

// DrawRadialGradient implements Graphics.
func (gc *graphics) DrawRadialGradient(gradient *Gradient, scx, scy, startRadius, ecx, ecy, endRadius float32) {
	gc.platformDrawRadialGradient(gradient, scx, scy, startRadius, ecx, ecy, endRadius)
}

// DrawImage implements Graphics.
func (gc *graphics) DrawImage(img *Image, where geom.Point) {
	gc.DrawImageInRect(img, geom.Rect{Point: where, Size: img.Size()})
}

// DrawImageInRect implements Graphics.
func (gc *graphics) DrawImageInRect(img *Image, bounds geom.Rect) {
	gc.Save()
	defer gc.Restore()
	gc.platformDrawImageInRect(img, bounds)
}

func (gc *graphics) DrawString(x, y float32, str string) {
	gc.platformDrawString(x, y, str)
}

// Translate implements Graphics.
func (gc *graphics) Translate(x, y float32) {
	gc.platformTranslate(x, y)
}

// Scale implements Graphics.
func (gc *graphics) Scale(x, y float32) {
	gc.platformScale(x, y)
}

// Rotate implements Graphics.
func (gc *graphics) Rotate(angleInRadians float32) {
	gc.platformRotate(angleInRadians)
}
