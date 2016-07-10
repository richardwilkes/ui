// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

// #cgo darwin LDFLAGS: -framework Cocoa
// #include <stdlib.h>
// #include <CoreText/CoreText.h>
import "C"

import (
	"math"
	"unsafe"
)

func (a *AttributedString) toPlatform() C.CFMutableAttributedStringRef {
	as := C.CFAttributedStringCreateMutable(C.kCFAllocatorDefault, 0)
	C.CFAttributedStringBeginEditing(as)
	C.CFAttributedStringReplaceString(as, C.CFRangeMake(0, 0), cfStringFromString(a.text))
	for _, ra := range a.attributes {
		begin := ra.begin
		length := ra.length
		if length == 0 {
			begin = 0
			length = len(a.text)
		}
		r := C.CFRangeMake(C.CFIndex(begin), C.CFIndex(length))
		switch ra.attr.key {
		case fontAttribute:
			C.CFAttributedStringSetAttribute(as, r, C.kCTFontAttributeName, C.CFTypeRef(ra.attr.value.(*Font).font))
		case foregroundAttribute:
			color := ra.attr.value.(Color)
			rgba := C.CGColorCreateGenericRGB(C.CGFloat(color.RedIntensity()), C.CGFloat(color.GreenIntensity()), C.CGFloat(color.BlueIntensity()), C.CGFloat(color.AlphaIntensity()))
			C.CFAttributedStringSetAttribute(as, r, C.kCTForegroundColorAttributeName, C.CFTypeRef(rgba))
		case alignmentAttribute:
			var align C.CTTextAlignment
			switch ra.attr.value.(Alignment) {
			case AlignStart:
				align = C.kCTTextAlignmentLeft
			case AlignMiddle:
				align = C.kCTTextAlignmentCenter
			case AlignEnd:
				align = C.kCTTextAlignmentRight
			case AlignFill:
				align = C.kCTTextAlignmentJustified
			default:
				align = C.kCTTextAlignmentLeft
			}
			alignPtr := (*C.CTTextAlignment)(C.malloc(C.size_t(unsafe.Sizeof(align))))
			*alignPtr = align
			defer C.free(unsafe.Pointer(alignPtr))
			para := C.CTParagraphStyleCreate(&C.CTParagraphStyleSetting{spec: C.kCTParagraphStyleSpecifierAlignment, valueSize: C.size_t(unsafe.Sizeof(align)), value: unsafe.Pointer(alignPtr)}, 1)
			C.CFAttributedStringSetAttribute(as, r, C.kCTParagraphStyleAttributeName, C.CFTypeRef(para))
		}
	}
	C.CFAttributedStringEndEditing(as)
	return as
}

func (a *AttributedString) platformMeasure(size Size) (actual Size, fit int) {
	attrStr := a.toPlatform()
	setter := C.CTFramesetterCreateWithAttributedString(attrStr)
	fitRange := C.CFRangeMake(0, 0)
	if size.Width < 0 {
		size.Width = math.MaxFloat32
	}
	if size.Height < 0 {
		size.Height = math.MaxFloat32
	}
	cSize := C.CTFramesetterSuggestFrameSizeWithConstraints(setter, C.CFRangeMake(0, 0), nil, C.CGSizeMake(C.CGFloat(size.Width), C.CGFloat(size.Height)), &fitRange)
	C.CFRelease(setter)
	C.CFRelease(attrStr)
	return Size{Width: float32(cSize.width), Height: float32(cSize.height)}, int(fitRange.length)
}
