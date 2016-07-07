// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package graphics

// #cgo darwin LDFLAGS: -framework Cocoa
// #include <stdio.h>
// #include <CoreGraphics/CoreGraphics.h>
// #include <CoreText/CoreText.h>
import "C"

import (
	"unsafe"

	"github.com/richardwilkes/go-ui/color"
	"github.com/richardwilkes/go-ui/font"
	"github.com/richardwilkes/go-ui/geom"
	"github.com/richardwilkes/go-ui/image"
)

const (
	drawingModeFill = iota
	drawingModeEOFill
	drawingModeStroke
	drawingModeFillStroke
	drawingModeEOFillStroke
)

const (
	gradientDrawsBeforeStartLocation = 1 << 0
	gradientDrawsAfterEndLocation    = 1 << 1
	gradientOverflow                 = gradientDrawsBeforeStartLocation | gradientDrawsAfterEndLocation
)

const (
	unlimited = 1000000
)

type context struct {
	gc    unsafe.Pointer
	stack []*state
}

type state struct {
	opacity     float32
	fillColor   color.Color
	strokeColor color.Color
	strokeWidth float32
	font        *font.Font
}

// NewContext creates a new graphics context to draw with.
func NewContext(gc unsafe.Pointer) Context {
	c := &context{gc: gc}
	c.stack = append(c.stack, &state{opacity: 1, fillColor: color.White, strokeColor: color.Black, strokeWidth: 1})
	return c
}

// Save implements Context.
func (gc *context) Save() {
	gs := *gc.stack[len(gc.stack)-1]
	gc.stack = append(gc.stack, &gs)
	C.CGContextSaveGState(gc.gc)
}

// Restore implements Context.
func (gc *context) Restore() {
	gc.stack[len(gc.stack)-1] = nil
	gc.stack = gc.stack[:len(gc.stack)-1]
	C.CGContextRestoreGState(gc.gc)
}

// Opacity implements Context.
func (gc *context) Opacity() float32 {
	return gc.stack[len(gc.stack)-1].opacity
}

// SetOpacity implements Context.
func (gc *context) SetOpacity(opacity float32) {
	if opacity < 0 {
		opacity = 0
	} else if opacity > 1 {
		opacity = 1
	}
	gc.stack[len(gc.stack)-1].opacity = opacity
	C.CGContextSetAlpha(gc.gc, C.CGFloat(opacity))
}

// FillColor implements Context.
func (gc *context) FillColor() color.Color {
	return gc.stack[len(gc.stack)-1].fillColor
}

// SetFillColor implements Context.
func (gc *context) SetFillColor(color color.Color) {
	gc.stack[len(gc.stack)-1].fillColor = color
	C.CGContextSetRGBFillColor(gc.gc, C.CGFloat(color.RedIntensity()), C.CGFloat(color.GreenIntensity()), C.CGFloat(color.BlueIntensity()), C.CGFloat(color.AlphaIntensity()))
}

// StrokeColor implements Context.
func (gc *context) StrokeColor() color.Color {
	return gc.stack[len(gc.stack)-1].strokeColor
}

// SetStrokeColor implements Context.
func (gc *context) SetStrokeColor(color color.Color) {
	gc.stack[len(gc.stack)-1].strokeColor = color
	C.CGContextSetRGBStrokeColor(gc.gc, C.CGFloat(color.RedIntensity()), C.CGFloat(color.GreenIntensity()), C.CGFloat(color.BlueIntensity()), C.CGFloat(color.AlphaIntensity()))
}

// StrokeWidth implements Context.
func (gc *context) StrokeWidth() float32 {
	return gc.stack[len(gc.stack)-1].strokeWidth
}

// SetStrokeWidth implements Context.
func (gc *context) SetStrokeWidth(width float32) {
	if width > 0 {
		gc.stack[len(gc.stack)-1].strokeWidth = width
		C.CGContextSetLineWidth(gc.gc, C.CGFloat(width))
	}
}

// Font implements Context.
func (gc *context) Font() *font.Font {
	return gc.stack[len(gc.stack)-1].font
}

// SetFont implements Context.
func (gc *context) SetFont(font *font.Font) {
	gc.stack[len(gc.stack)-1].font = font
}

// StrokeLine implements Context.
// To ensure the line is aligned to pixel boundaries, 0.5 is added to each coordinate.
func (gc *context) StrokeLine(x1, y1, x2, y2 float32) {
	gc.BeginPath()
	gc.MoveTo(x1+0.5, y1+0.5)
	gc.LineTo(x2+0.5, y2+0.5)
	gc.StrokePath()
}

// FillRect implements Context.
func (gc *context) FillRect(bounds geom.Rect) {
	C.CGContextFillRect(gc.gc, toCGRect(bounds))
}

// StrokeRect implements Context.
// To ensure the rectangle is aligned to pixel boundaries, 0.5 is added to the origin coordinates
// and 1 is subtracted from the size values.
func (gc *context) StrokeRect(bounds geom.Rect) {
	if bounds.Width <= 1 || bounds.Height <= 1 {
		savedColor := gc.FillColor()
		gc.SetFillColor(gc.StrokeColor())
		gc.FillRect(bounds)
		gc.SetFillColor(savedColor)
	} else {
		bounds.InsetUniform(0.5)
		C.CGContextStrokeRect(gc.gc, toCGRect(bounds))
	}
}

// FillEllipse implements Context.
func (gc *context) FillEllipse(bounds geom.Rect) {
	C.CGContextFillEllipseInRect(gc.gc, toCGRect(bounds))
}

// StrokeEllipse implements Context.
// To ensure the ellipse is aligned to pixel boundaries, 0.5 is added to the origin coordinates
// and 1 is subtracted from the size values.
func (gc *context) StrokeEllipse(bounds geom.Rect) {
	if bounds.Width <= 1 || bounds.Height <= 1 {
		savedColor := gc.FillColor()
		gc.SetFillColor(gc.StrokeColor())
		gc.FillEllipse(bounds)
		gc.SetFillColor(savedColor)
	} else {
		bounds.InsetUniform(0.5)
		C.CGContextStrokeEllipseInRect(gc.gc, toCGRect(bounds))
	}
}

// FillPath implements Context.
func (gc *context) FillPath() {
	C.CGContextFillPath(gc.gc)
}

// FillPathEvenOdd implements Context.
func (gc *context) FillPathEvenOdd() {
	C.CGContextEOFillPath(gc.gc)
}

// StrokePath implements Context.
func (gc *context) StrokePath() {
	C.CGContextStrokePath(gc.gc)
}

// FillAndStrokePath implements Context.
func (gc *context) FillAndStrokePath() {
	C.CGContextDrawPath(gc.gc, drawingModeFillStroke)
}

// BeginPath implements Context.
func (gc *context) BeginPath() {
	C.CGContextBeginPath(gc.gc)
}

// ClosePath implements Context.
func (gc *context) ClosePath() {
	C.CGContextClosePath(gc.gc)
}

// MoveTo implements Context.
func (gc *context) MoveTo(x, y float32) {
	C.CGContextMoveToPoint(gc.gc, C.CGFloat(x), C.CGFloat(y))
}

// LineTo implements Context.
func (gc *context) LineTo(x, y float32) {
	C.CGContextAddLineToPoint(gc.gc, C.CGFloat(x), C.CGFloat(y))
}

// Arc implements Context.
func (gc *context) Arc(cx, cy, radius, startAngleRadians, endAngleRadians float32, clockwise bool) {
	var cw int
	// Invert clockwise to accommodate flipped view and convert to an int
	if clockwise {
		cw = 0
	} else {
		cw = 1
	}
	C.CGContextAddArc(gc.gc, C.CGFloat(cx), C.CGFloat(cy), C.CGFloat(radius), C.CGFloat(startAngleRadians), C.CGFloat(endAngleRadians), C.int(cw))
}

// ArcTo implements Context.
func (gc *context) ArcTo(x1, y1, x2, y2, radius float32) {
	C.CGContextAddArcToPoint(gc.gc, C.CGFloat(x1), C.CGFloat(y1), C.CGFloat(x2), C.CGFloat(y2), C.CGFloat(radius))
}

// CurveTo implements Context.
func (gc *context) CurveTo(cp1x, cp1y, cp2x, cp2y, x, y float32) {
	C.CGContextAddCurveToPoint(gc.gc, C.CGFloat(cp1x), C.CGFloat(cp1y), C.CGFloat(cp2x), C.CGFloat(cp2y), C.CGFloat(x), C.CGFloat(y))
}

// QuadCurveTo implements Context.
func (gc *context) QuadCurveTo(cpx, cpy, x, y float32) {
	C.CGContextAddQuadCurveToPoint(gc.gc, C.CGFloat(cpx), C.CGFloat(cpy), C.CGFloat(x), C.CGFloat(y))
}

// AddPath implements Context.
func (gc *context) AddPath(path *Path) {
	platformPath := path.toPlatform()
	C.CGContextAddPath(gc.gc, platformPath)
	C.CGPathRelease(platformPath)
}

// Clip implements Context.
func (gc *context) Clip() {
	C.CGContextClip(gc.gc)
}

// ClipEvenOdd implements Context.
func (gc *context) ClipEvenOdd() {
	C.CGContextEOClip(gc.gc)
}

// ClipRect implements Context.
func (gc *context) ClipRect(bounds geom.Rect) {
	C.CGContextClipToRect(gc.gc, toCGRect(bounds))
}

// DrawLinearGradient implements Context.
func (gc *context) DrawLinearGradient(gradient *Gradient, sx, sy, ex, ey float32) {
	platformGradient := gradient.toPlatform()
	C.CGContextDrawLinearGradient(gc.gc, platformGradient, C.CGPointMake(C.CGFloat(sx), C.CGFloat(sy)), C.CGPointMake(C.CGFloat(ex), C.CGFloat(ey)), C.CGGradientDrawingOptions(gradientOverflow))
	C.CGGradientRelease(platformGradient)
}

// DrawRadialGradient implements Context.
func (gc *context) DrawRadialGradient(gradient *Gradient, scx, scy, startRadius, ecx, ecy, endRadius float32) {
	platformGradient := gradient.toPlatform()
	C.CGContextDrawRadialGradient(gc.gc, platformGradient, C.CGPointMake(C.CGFloat(scx), C.CGFloat(scy)), C.CGFloat(startRadius), C.CGPointMake(C.CGFloat(ecx), C.CGFloat(ecy)), C.CGFloat(endRadius), C.CGGradientDrawingOptions(gradientOverflow))
	C.CGGradientRelease(platformGradient)
}

// DrawImage implements Context.
func (gc *context) DrawImage(img *image.Image, where geom.Point) {
	gc.DrawImageInRect(img, geom.Rect{Point: where, Size: img.Size()})
}

// DrawImageInRect implements Context.
func (gc *context) DrawImageInRect(img *image.Image, bounds geom.Rect) {
	gc.Save()
	defer gc.Restore()
	C.CGContextTranslateCTM(gc.gc, 0, C.CGFloat(bounds.Y+bounds.Height))
	C.CGContextScaleCTM(gc.gc, 1, -1)
	bounds.Y = 0
	C.CGContextDrawImage(gc.gc, toCGRect(bounds), img.PlatformPointer())
}

// DrawText implements Context.
func (gc *context) DrawText(x, y float32, str string, mode TextMode) geom.Size {
	size, _ := gc.DrawTextConstrained(geom.Rect{Point: geom.Point{X: x, Y: y}, Size: geom.Size{Width: unlimited, Height: unlimited}}, str, mode)
	return size
}

// DrawTextConstrained implements Context.
func (gc *context) DrawTextConstrained(bounds geom.Rect, str string, mode TextMode) (actual geom.Size, fit int) {
	gs := gc.stack[len(gc.stack)-1]
	return gc.DrawAttributedTextConstrained(bounds, NewAttributedString(str, gs.fillColor, gs.font), mode)
}

// DrawAttributedText implements Context.
func (gc *context) DrawAttributedText(x, y float32, str *AttributedString, mode TextMode) geom.Size {
	size, _ := gc.DrawAttributedTextConstrained(geom.Rect{Point: geom.Point{X: x, Y: y}, Size: geom.Size{Width: unlimited, Height: unlimited}}, str, mode)
	return size
}

// DrawAttributedTextConstrained implements Context.
func (gc *context) DrawAttributedTextConstrained(bounds geom.Rect, str *AttributedString, mode TextMode) (actual geom.Size, fit int) {
	gc.Save()
	defer gc.Restore()
	attrStr := str.toPlatform()
	setter := C.CTFramesetterCreateWithAttributedString(attrStr)
	fitRange := C.CFRangeMake(0, 0)
	size := C.CTFramesetterSuggestFrameSizeWithConstraints(setter, C.CFRangeMake(0, 0), nil, C.CGSizeMake(C.CGFloat(bounds.Width), C.CGFloat(bounds.Height)), &fitRange)
	C.CGContextSetTextMatrix(gc.gc, C.CGAffineTransformIdentity)
	C.CGContextSetTextDrawingMode(gc.gc, C.CGTextDrawingMode(mode))
	C.CGContextTranslateCTM(gc.gc, 0, size.height)
	C.CGContextScaleCTM(gc.gc, 1, -1)
	path := C.CGPathCreateMutable()
	bounds.Y = -bounds.Y
	bounds.Height = float32(size.height)
	C.CGPathAddRect(path, nil, toCGRect(bounds))
	frame := C.CTFramesetterCreateFrame(setter, C.CFRangeMake(0, 0), path, nil)
	C.CGPathRelease(path)
	C.CTFrameDraw(frame, gc.gc)
	C.CFRelease(frame)
	C.CFRelease(setter)
	C.CFRelease(attrStr)
	return geom.Size{Width: float32(size.width), Height: float32(size.height)}, int(fitRange.length)
}

// Translate implements Context.
func (gc *context) Translate(x, y float32) {
	C.CGContextTranslateCTM(gc.gc, C.CGFloat(x), C.CGFloat(y))
}

// Scale implements Context.
func (gc *context) Scale(x, y float32) {
	C.CGContextScaleCTM(gc.gc, C.CGFloat(x), C.CGFloat(y))
}

// Rotate implements Context.
func (gc *context) Rotate(angleInRadians float32) {
	C.CGContextRotateCTM(gc.gc, C.CGFloat(angleInRadians))
}

func (gc *context) createAttributedString(str string) C.CFMutableAttributedStringRef {
	gs := gc.stack[len(gc.stack)-1]
	return NewAttributedString(str, gs.fillColor, gs.font).toPlatform()
}

func toCGRect(bounds geom.Rect) C.CGRect {
	return C.CGRectMake(C.CGFloat(bounds.X), C.CGFloat(bounds.Y), C.CGFloat(bounds.Width), C.CGFloat(bounds.Height))
}

func cfStringFromString(str string) C.CFStringRef {
	cstr := C.CString(str)
	cfstr := C.CFStringCreateWithBytes(nil, (*C.UInt8)(unsafe.Pointer(cstr)), C.CFIndex(len(str)), C.kCFStringEncodingUTF8, C.Boolean(0))
	C.free(unsafe.Pointer(cstr))
	return cfstr
}
