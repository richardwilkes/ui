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

func (gc *graphics) platformSave() {
	C.CGContextSaveGState(gc.gc)
}

func (gc *graphics) platformRestore() {
	C.CGContextRestoreGState(gc.gc)
}

func (gc *graphics) platformSetOpacity(opacity float32) {
	C.CGContextSetAlpha(gc.gc, C.CGFloat(opacity))
}

func (gc *graphics) platformSetFillColor(color color.Color) {
	C.CGContextSetRGBFillColor(gc.gc, C.CGFloat(color.RedIntensity()), C.CGFloat(color.GreenIntensity()), C.CGFloat(color.BlueIntensity()), C.CGFloat(color.AlphaIntensity()))
}

func (gc *graphics) platformSetStrokeColor(color color.Color) {
	C.CGContextSetRGBStrokeColor(gc.gc, C.CGFloat(color.RedIntensity()), C.CGFloat(color.GreenIntensity()), C.CGFloat(color.BlueIntensity()), C.CGFloat(color.AlphaIntensity()))
}

func (gc *graphics) platformSetStrokeWidth(width float32) {
	C.CGContextSetLineWidth(gc.gc, C.CGFloat(width))
}

func (gc *graphics) platformFillRect(bounds geom.Rect) {
	C.CGContextFillRect(gc.gc, toCGRect(bounds))
}

func (gc *graphics) platformStrokeRect(bounds geom.Rect) {
	C.CGContextStrokeRect(gc.gc, toCGRect(bounds))
}

func (gc *graphics) platformFillEllipse(bounds geom.Rect) {
	C.CGContextFillEllipseInRect(gc.gc, toCGRect(bounds))
}

func (gc *graphics) platformStrokeEllipse(bounds geom.Rect) {
	C.CGContextStrokeEllipseInRect(gc.gc, toCGRect(bounds))
}

func (gc *graphics) platformFillPath() {
	C.CGContextFillPath(gc.gc)
}

func (gc *graphics) platformFillPathEvenOdd() {
	C.CGContextEOFillPath(gc.gc)
}

func (gc *graphics) platformStrokePath() {
	C.CGContextStrokePath(gc.gc)
}

func (gc *graphics) platformFillAndStrokePath() {
	C.CGContextDrawPath(gc.gc, drawingModeFillStroke)
}

func (gc *graphics) platformBeginPath() {
	C.CGContextBeginPath(gc.gc)
}

func (gc *graphics) platformClosePath() {
	C.CGContextClosePath(gc.gc)
}

func (gc *graphics) platformMoveTo(x, y float32) {
	C.CGContextMoveToPoint(gc.gc, C.CGFloat(x), C.CGFloat(y))
}

func (gc *graphics) platformLineTo(x, y float32) {
	C.CGContextAddLineToPoint(gc.gc, C.CGFloat(x), C.CGFloat(y))
}

func (gc *graphics) platformArc(cx, cy, radius, startAngleRadians, endAngleRadians float32, clockwise bool) {
	var cw int
	// Invert clockwise to accommodate flipped view and convert to an int
	if clockwise {
		cw = 0
	} else {
		cw = 1
	}
	C.CGContextAddArc(gc.gc, C.CGFloat(cx), C.CGFloat(cy), C.CGFloat(radius), C.CGFloat(startAngleRadians), C.CGFloat(endAngleRadians), C.int(cw))
}

func (gc *graphics) platformArcTo(x1, y1, x2, y2, radius float32) {
	C.CGContextAddArcToPoint(gc.gc, C.CGFloat(x1), C.CGFloat(y1), C.CGFloat(x2), C.CGFloat(y2), C.CGFloat(radius))
}

func (gc *graphics) platformCurveTo(cp1x, cp1y, cp2x, cp2y, x, y float32) {
	C.CGContextAddCurveToPoint(gc.gc, C.CGFloat(cp1x), C.CGFloat(cp1y), C.CGFloat(cp2x), C.CGFloat(cp2y), C.CGFloat(x), C.CGFloat(y))
}

func (gc *graphics) platformQuadCurveTo(cpx, cpy, x, y float32) {
	C.CGContextAddQuadCurveToPoint(gc.gc, C.CGFloat(cpx), C.CGFloat(cpy), C.CGFloat(x), C.CGFloat(y))
}

func (gc *graphics) platformAddPath(path *geom.Path) {
	platformPath := path.PlatformPtr()
	C.CGContextAddPath(gc.gc, platformPath)
	C.CGPathRelease(platformPath)
}

func (gc *graphics) platformClip() {
	C.CGContextClip(gc.gc)
}

func (gc *graphics) platformClipEvenOdd() {
	C.CGContextEOClip(gc.gc)
}

func (gc *graphics) platformClipRect(bounds geom.Rect) {
	C.CGContextClipToRect(gc.gc, toCGRect(bounds))
}

func (gc *graphics) platformDrawLinearGradient(gradient *Gradient, sx, sy, ex, ey float32) {
	platformGradient := gradient.platformData()
	C.CGContextDrawLinearGradient(gc.gc, platformGradient, C.CGPointMake(C.CGFloat(sx), C.CGFloat(sy)), C.CGPointMake(C.CGFloat(ex), C.CGFloat(ey)), C.CGGradientDrawingOptions(gradientOverflow))
	C.CGGradientRelease(platformGradient)
}

func (gc *graphics) platformDrawRadialGradient(gradient *Gradient, scx, scy, startRadius, ecx, ecy, endRadius float32) {
	platformGradient := gradient.platformData()
	C.CGContextDrawRadialGradient(gc.gc, platformGradient, C.CGPointMake(C.CGFloat(scx), C.CGFloat(scy)), C.CGFloat(startRadius), C.CGPointMake(C.CGFloat(ecx), C.CGFloat(ecy)), C.CGFloat(endRadius), C.CGGradientDrawingOptions(gradientOverflow))
	C.CGGradientRelease(platformGradient)
}

func (gc *graphics) platformDrawImageInRect(img *Image, bounds geom.Rect) {
	C.CGContextTranslateCTM(gc.gc, 0, C.CGFloat(bounds.Y+bounds.Height))
	C.CGContextScaleCTM(gc.gc, 1, -1)
	bounds.Y = 0
	C.CGContextDrawImage(gc.gc, toCGRect(bounds), img.img)
}

func (gc *graphics) platformDrawString(x, y float32, str string) {
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

func (gc *graphics) platformTranslate(x, y float32) {
	C.CGContextTranslateCTM(gc.gc, C.CGFloat(x), C.CGFloat(y))
}

func (gc *graphics) platformScale(x, y float32) {
	C.CGContextScaleCTM(gc.gc, C.CGFloat(x), C.CGFloat(y))
}

func (gc *graphics) platformRotate(angleInRadians float32) {
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
