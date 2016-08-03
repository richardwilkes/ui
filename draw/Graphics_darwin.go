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

// #cgo darwin LDFLAGS: -framework Cocoa
// #include <stdio.h>
// #include <CoreGraphics/CoreGraphics.h>
// #include <CoreText/CoreText.h>
import "C"

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
	C.CGContextSaveGState(gc.gc)
}

// Restore implements Graphics.
func (gc *graphics) Restore() {
	gc.stack[len(gc.stack)-1] = nil
	gc.stack = gc.stack[:len(gc.stack)-1]
	C.CGContextRestoreGState(gc.gc)
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
	C.CGContextSetAlpha(gc.gc, C.CGFloat(opacity))
}

// FillColor implements Graphics.
func (gc *graphics) FillColor() color.Color {
	return gc.stack[len(gc.stack)-1].fillColor
}

// SetFillColor implements Graphics.
func (gc *graphics) SetFillColor(color color.Color) {
	gc.stack[len(gc.stack)-1].fillColor = color
	C.CGContextSetRGBFillColor(gc.gc, C.CGFloat(color.RedIntensity()), C.CGFloat(color.GreenIntensity()), C.CGFloat(color.BlueIntensity()), C.CGFloat(color.AlphaIntensity()))
}

// StrokeColor implements Graphics.
func (gc *graphics) StrokeColor() color.Color {
	return gc.stack[len(gc.stack)-1].strokeColor
}

// SetStrokeColor implements Graphics.
func (gc *graphics) SetStrokeColor(color color.Color) {
	gc.stack[len(gc.stack)-1].strokeColor = color
	C.CGContextSetRGBStrokeColor(gc.gc, C.CGFloat(color.RedIntensity()), C.CGFloat(color.GreenIntensity()), C.CGFloat(color.BlueIntensity()), C.CGFloat(color.AlphaIntensity()))
}

// StrokeWidth implements Graphics.
func (gc *graphics) StrokeWidth() float32 {
	return gc.stack[len(gc.stack)-1].strokeWidth
}

// SetStrokeWidth implements Graphics.
func (gc *graphics) SetStrokeWidth(width float32) {
	if width > 0 {
		gc.stack[len(gc.stack)-1].strokeWidth = width
		C.CGContextSetLineWidth(gc.gc, C.CGFloat(width))
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
	C.CGContextFillRect(gc.gc, toCGRect(bounds))
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
		C.CGContextStrokeRect(gc.gc, toCGRect(bounds))
	}
}

// FillEllipse implements Graphics.
func (gc *graphics) FillEllipse(bounds geom.Rect) {
	C.CGContextFillEllipseInRect(gc.gc, toCGRect(bounds))
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
		C.CGContextStrokeEllipseInRect(gc.gc, toCGRect(bounds))
	}
}

// FillPath implements Graphics.
func (gc *graphics) FillPath() {
	C.CGContextFillPath(gc.gc)
}

// FillPathEvenOdd implements Graphics.
func (gc *graphics) FillPathEvenOdd() {
	C.CGContextEOFillPath(gc.gc)
}

// StrokePath implements Graphics.
func (gc *graphics) StrokePath() {
	C.CGContextStrokePath(gc.gc)
}

// FillAndStrokePath implements Graphics.
func (gc *graphics) FillAndStrokePath() {
	C.CGContextDrawPath(gc.gc, drawingModeFillStroke)
}

// BeginPath implements Graphics.
func (gc *graphics) BeginPath() {
	C.CGContextBeginPath(gc.gc)
}

// ClosePath implements Graphics.
func (gc *graphics) ClosePath() {
	C.CGContextClosePath(gc.gc)
}

// MoveTo implements Graphics.
func (gc *graphics) MoveTo(x, y float32) {
	C.CGContextMoveToPoint(gc.gc, C.CGFloat(x), C.CGFloat(y))
}

// LineTo implements Graphics.
func (gc *graphics) LineTo(x, y float32) {
	C.CGContextAddLineToPoint(gc.gc, C.CGFloat(x), C.CGFloat(y))
}

// Arc implements Graphics.
func (gc *graphics) Arc(cx, cy, radius, startAngleRadians, endAngleRadians float32, clockwise bool) {
	var cw int
	// Invert clockwise to accommodate flipped view and convert to an int
	if clockwise {
		cw = 0
	} else {
		cw = 1
	}
	C.CGContextAddArc(gc.gc, C.CGFloat(cx), C.CGFloat(cy), C.CGFloat(radius), C.CGFloat(startAngleRadians), C.CGFloat(endAngleRadians), C.int(cw))
}

// ArcTo implements Graphics.
func (gc *graphics) ArcTo(x1, y1, x2, y2, radius float32) {
	C.CGContextAddArcToPoint(gc.gc, C.CGFloat(x1), C.CGFloat(y1), C.CGFloat(x2), C.CGFloat(y2), C.CGFloat(radius))
}

// CurveTo implements Graphics.
func (gc *graphics) CurveTo(cp1x, cp1y, cp2x, cp2y, x, y float32) {
	C.CGContextAddCurveToPoint(gc.gc, C.CGFloat(cp1x), C.CGFloat(cp1y), C.CGFloat(cp2x), C.CGFloat(cp2y), C.CGFloat(x), C.CGFloat(y))
}

// QuadCurveTo implements Graphics.
func (gc *graphics) QuadCurveTo(cpx, cpy, x, y float32) {
	C.CGContextAddQuadCurveToPoint(gc.gc, C.CGFloat(cpx), C.CGFloat(cpy), C.CGFloat(x), C.CGFloat(y))
}

// AddPath implements Graphics.
func (gc *graphics) AddPath(path *geom.Path) {
	platformPath := path.PlatformPtr()
	C.CGContextAddPath(gc.gc, platformPath)
	C.CGPathRelease(platformPath)
}

// Clip implements Graphics.
func (gc *graphics) Clip() {
	C.CGContextClip(gc.gc)
}

// ClipEvenOdd implements Graphics.
func (gc *graphics) ClipEvenOdd() {
	C.CGContextEOClip(gc.gc)
}

// ClipRect implements Graphics.
func (gc *graphics) ClipRect(bounds geom.Rect) {
	C.CGContextClipToRect(gc.gc, toCGRect(bounds))
}

// DrawLinearGradient implements Graphics.
func (gc *graphics) DrawLinearGradient(gradient *Gradient, sx, sy, ex, ey float32) {
	platformGradient := gradient.toPlatform()
	C.CGContextDrawLinearGradient(gc.gc, platformGradient, C.CGPointMake(C.CGFloat(sx), C.CGFloat(sy)), C.CGPointMake(C.CGFloat(ex), C.CGFloat(ey)), C.CGGradientDrawingOptions(gradientOverflow))
	C.CGGradientRelease(platformGradient)
}

// DrawRadialGradient implements Graphics.
func (gc *graphics) DrawRadialGradient(gradient *Gradient, scx, scy, startRadius, ecx, ecy, endRadius float32) {
	platformGradient := gradient.toPlatform()
	C.CGContextDrawRadialGradient(gc.gc, platformGradient, C.CGPointMake(C.CGFloat(scx), C.CGFloat(scy)), C.CGFloat(startRadius), C.CGPointMake(C.CGFloat(ecx), C.CGFloat(ecy)), C.CGFloat(endRadius), C.CGGradientDrawingOptions(gradientOverflow))
	C.CGGradientRelease(platformGradient)
}

// DrawImage implements Graphics.
func (gc *graphics) DrawImage(img *Image, where geom.Point) {
	gc.DrawImageInRect(img, geom.Rect{Point: where, Size: img.Size()})
}

// DrawImageInRect implements Graphics.
func (gc *graphics) DrawImageInRect(img *Image, bounds geom.Rect) {
	gc.Save()
	defer gc.Restore()
	C.CGContextTranslateCTM(gc.gc, 0, C.CGFloat(bounds.Y+bounds.Height))
	C.CGContextScaleCTM(gc.gc, 1, -1)
	bounds.Y = 0
	C.CGContextDrawImage(gc.gc, toCGRect(bounds), img.img)
}

func (gc *graphics) DrawString(x, y float32, str string) {
	gc.Save()
	gs := gc.stack[len(gc.stack)-1]
	as := C.CFAttributedStringCreateMutable(C.kCFAllocatorDefault, 0)
	C.CFAttributedStringBeginEditing(as)
	s := cfStringFromString(str)
	length := C.CFStringGetLength(s)
	C.CFAttributedStringReplaceString(as, C.CFRangeMake(0, 0), s)
	C.CFAttributedStringSetAttribute(as, C.CFRangeMake(0, length), C.kCTFontAttributeName, C.CFTypeRef(gs.font.PlatformPtr()))
	C.CFAttributedStringSetAttribute(as, C.CFRangeMake(0, length), C.kCTForegroundColorFromContextAttributeName, C.CFTypeRef(C.kCFBooleanTrue))
	C.CFAttributedStringEndEditing(as)
	line := C.CTLineCreateWithAttributedString(as)
	C.CGContextSetTextMatrix(gc.gc, C.CGAffineTransformIdentity)
	C.CGContextTranslateCTM(gc.gc, 0, C.CGFloat(gs.font.Ascent()))
	C.CGContextScaleCTM(gc.gc, 1, -1)
	C.CGContextTranslateCTM(gc.gc, C.CGFloat(x), C.CGFloat(-y))
	C.CTLineDraw(line, gc.gc)
	gc.Restore()
	C.CFRelease(line)
	C.CFRelease(as)
}

// Translate implements Graphics.
func (gc *graphics) Translate(x, y float32) {
	C.CGContextTranslateCTM(gc.gc, C.CGFloat(x), C.CGFloat(y))
}

// Scale implements Graphics.
func (gc *graphics) Scale(x, y float32) {
	C.CGContextScaleCTM(gc.gc, C.CGFloat(x), C.CGFloat(y))
}

// Rotate implements Graphics.
func (gc *graphics) Rotate(angleInRadians float32) {
	C.CGContextRotateCTM(gc.gc, C.CGFloat(angleInRadians))
}

func toCGRect(bounds geom.Rect) C.CGRect {
	return C.CGRectMake(C.CGFloat(bounds.X), C.CGFloat(bounds.Y), C.CGFloat(bounds.Width), C.CGFloat(bounds.Height))
}

func cfStringFromString(str string) C.CFStringRef {
	cstr := C.CString(str)
	cfstr := C.CFStringCreateWithBytes(nil, (*C.UInt8)(unsafe.Pointer(cstr)), C.CFIndex(len(str)), C.kCFStringEncodingUTF8, 0)
	C.free(unsafe.Pointer(cstr))
	return cfstr
}
