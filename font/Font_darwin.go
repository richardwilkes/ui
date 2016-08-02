// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package font

// #cgo darwin LDFLAGS: -framework Cocoa
// #include <stdlib.h>
// #include <CoreText/CoreText.h>
import "C"

import (
	"unsafe"
)

func platformAvailableFontFamilies() []string {
	var names []string
	unique := make(map[string]bool)
	allFontsCollection := C.CTFontCollectionCreateFromAvailableFonts(nil)
	allFonts := C.CTFontCollectionCreateMatchingFontDescriptors(allFontsCollection)
	for i := C.CFArrayGetCount(allFonts) - 1; i >= 0; i-- {
		family := stringFromDescAttr(C.CTFontDescriptorRef(C.CFArrayGetValueAtIndex(allFonts, i)), C.kCTFontFamilyNameAttribute)
		if _, exists := unique[family]; !exists {
			unique[family] = true
			names = append(names, family)
		}
	}
	C.CFRelease(allFonts)
	C.CFRelease(allFontsCollection)
	return names
}

func platformNewFont(desc Desc) *Font {
	desc.Style &= UserSettable
	if desc.Condensed() && desc.Expanded() {
		desc.Style &= ^Expanded
	}
	font := C.CTFontCreateWithName(cfStringFromString(desc.Family), C.CGFloat(desc.Size), nil)
	if (desc.Style & UserSettable) != 0 {
		var traits C.CTFontSymbolicTraits
	tryAgain:
		if desc.Bold() {
			traits |= C.kCTFontBoldTrait
		}
		if desc.Italic() {
			traits |= C.kCTFontItalicTrait
		}
		if desc.Condensed() {
			traits |= C.kCTFontCondensedTrait
		} else if desc.Expanded() {
			traits |= C.kCTFontExpandedTrait
		}
		adjustedFont := C.CTFontCreateCopyWithSymbolicTraits(font, C.CGFloat(desc.Size), nil, traits, C.kCTFontItalicTrait|C.kCTFontBoldTrait|C.kCTFontExpandedTrait|C.kCTFontCondensedTrait)
		if adjustedFont == nil {
			traits = 0
			if desc.Expanded() {
				desc.Style &= ^Expanded
			} else if desc.Condensed() {
				desc.Style &= ^Condensed
			} else if desc.Italic() {
				desc.Style &= ^Italic
			} else if desc.Bold() {
				desc.Style &= ^Bold
			}
			if (desc.Style & UserSettable) != 0 {
				goto tryAgain
			}
			adjustedFont = font
		}
		font = adjustedFont
	}
	C.CFRetain(font)
	traits := C.CTFontGetSymbolicTraits(font)
	desc.Style &= ^(Condensed | Expanded)
	if traits&C.kCTFontBoldTrait == C.kCTFontBoldTrait {
		desc.Style |= ActualBold
	}
	if traits&C.kCTFontItalicTrait == C.kCTFontItalicTrait {
		desc.Style |= ActualItalic
	}
	if traits&C.kCTFontCondensedTrait == C.kCTFontCondensedTrait {
		desc.Style |= Condensed
	}
	if traits&C.kCTFontExpandedTrait == C.kCTFontExpandedTrait {
		desc.Style |= Expanded
	}
	if traits&C.kCTFontMonoSpaceTrait == C.kCTFontMonoSpaceTrait {
		desc.Style |= Monospaced
	}
	return &Font{font: unsafe.Pointer(font), desc: desc}
}

func (f *Font) platformDispose() {
	C.CFRelease(f.font)
}

func (f *Font) platformAscent() float32 {
	return float32(C.CTFontGetAscent(f.font))
}

func (f *Font) platformDescent() float32 {
	return float32(C.CTFontGetDescent(f.font))
}

func (f *Font) platformLeading() float32 {
	return float32(C.CTFontGetLeading(f.font))
}

func (f *Font) platformWidth(str string) float32 {
	as := f.createAttributedString(str)
	line := C.CTLineCreateWithAttributedString(as)
	width := float32(C.CTLineGetTypographicBounds(line, nil, nil, nil))
	C.CFRelease(line)
	C.CFRelease(as)
	return width
}

func (f *Font) createAttributedString(str string) C.CFMutableAttributedStringRef {
	as := C.CFAttributedStringCreateMutable(C.kCFAllocatorDefault, 0)
	C.CFAttributedStringBeginEditing(as)
	s := cfStringFromString(str)
	C.CFAttributedStringReplaceString(as, C.CFRangeMake(0, 0), s)
	C.CFAttributedStringSetAttribute(as, C.CFRangeMake(0, C.CFStringGetLength(s)), C.kCTFontAttributeName, C.CFTypeRef(f.PlatformPtr()))
	C.CFAttributedStringEndEditing(as)
	return as
}

func (f *Font) platformIndexForPosition(x float32, str string) int {
	as := f.createAttributedString(str)
	line := C.CTLineCreateWithAttributedString(as)
	i := C.CTLineGetStringIndexForPosition(line, C.CGPointMake(C.CGFloat(x), 0))
	C.CFRelease(line)
	C.CFRelease(as)
	return int(i)
}

func (f *Font) platformPositionForIndex(index int, str string) float32 {
	as := f.createAttributedString(str)
	line := C.CTLineCreateWithAttributedString(as)
	x := C.CTLineGetOffsetForStringIndex(line, C.CFIndex(index), nil)
	C.CFRelease(line)
	C.CFRelease(as)
	return float32(x)
}

func platformDesc(id int) Desc {
	var fontType C.int
	switch id {
	case userID:
		fontType = C.kCTFontUserFontType
	case userMonospacedID:
		fontType = C.kCTFontUserFixedPitchFontType
	case systemID:
		fontType = C.kCTFontSystemFontType
	case emphasizedSystemID:
		fontType = C.kCTFontEmphasizedSystemFontType
	case smallSystemID:
		fontType = C.kCTFontSmallSystemFontType
	case smallEmphasizedSystemID:
		fontType = C.kCTFontSmallEmphasizedSystemFontType
	case viewsID:
		fontType = C.kCTFontViewsFontType
	case labelID:
		fontType = C.kCTFontLabelFontType
	case menuID:
		fontType = C.kCTFontMenuItemFontType
	case menuCmdKeyID:
		fontType = C.kCTFontMenuItemCmdKeyFontType
	default:
		fontType = C.kCTFontUserFontType
	}
	font := C.CTFontCreateUIFontForLanguage(C.CTFontUIFontType(fontType), 0, nil)
	C.CFRetain(font)
	traits := C.CTFontGetSymbolicTraits(font)
	var desc Desc
	desc.Family = stringFromCFString(C.CTFontCopyFamilyName(font))
	desc.Size = float32(C.CTFontGetSize(font))
	if traits&C.kCTFontBoldTrait == C.kCTFontBoldTrait {
		desc.Style |= ActualBold | Bold
	}
	if traits&C.kCTFontItalicTrait == C.kCTFontItalicTrait {
		desc.Style |= ActualItalic | Italic
	}
	if traits&C.kCTFontCondensedTrait == C.kCTFontCondensedTrait {
		desc.Style |= Condensed
	}
	if traits&C.kCTFontExpandedTrait == C.kCTFontExpandedTrait {
		desc.Style |= Expanded
	}
	if traits&C.kCTFontMonoSpaceTrait == C.kCTFontMonoSpaceTrait {
		desc.Style |= Monospaced
	}
	C.CFRelease(font)
	return desc
}

func stringFromDescAttr(desc C.CTFontDescriptorRef, attrName C.CFStringRef) string {
	str := C.CFStringRef(C.CTFontDescriptorCopyAttribute(desc, attrName))
	result := stringFromCFString(str)
	C.CFRelease(str)
	return result
}

func stringFromCFString(cfStr C.CFStringRef) string {
	var freeUTF8StringPtr *C.char
	useUTF8StringPtr := C.CFStringGetCStringPtr(cfStr, C.kCFStringEncodingUTF8)
	if useUTF8StringPtr == nil {
		stringLength := C.CFStringGetLength(cfStr)
		maxBytes := 4*stringLength + 1
		freeUTF8StringPtr = (*C.char)(C.malloc(C.size_t(maxBytes)))
		C.CFStringGetCString(cfStr, freeUTF8StringPtr, maxBytes, C.kCFStringEncodingUTF8)
		useUTF8StringPtr = freeUTF8StringPtr
	}
	str := C.GoString(useUTF8StringPtr)
	if freeUTF8StringPtr != nil {
		C.free(unsafe.Pointer(freeUTF8StringPtr))
	}
	return str
}

func cfStringFromString(str string) C.CFStringRef {
	cstr := C.CString(str)
	cfstr := C.CFStringCreateWithBytes(nil, (*C.UInt8)(unsafe.Pointer(cstr)), C.CFIndex(len(str)), C.kCFStringEncodingUTF8, 0)
	C.free(unsafe.Pointer(cstr))
	return cfstr
}
