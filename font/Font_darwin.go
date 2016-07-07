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

func init() {
	User = systemFontDesc(C.kCTFontUserFontType)
	UserMonospaced = systemFontDesc(C.kCTFontUserFixedPitchFontType)
	System = systemFontDesc(C.kCTFontSystemFontType)
	EmphasizedSystem = systemFontDesc(C.kCTFontEmphasizedSystemFontType)
	SmallSystem = systemFontDesc(C.kCTFontSmallSystemFontType)
	SmallEmphasizedSystem = systemFontDesc(C.kCTFontSmallEmphasizedSystemFontType)
	Views = systemFontDesc(C.kCTFontViewsFontType)
	Label = systemFontDesc(C.kCTFontLabelFontType)
	Menu = systemFontDesc(C.kCTFontMenuItemFontType)
	MenuCmdKey = systemFontDesc(C.kCTFontMenuItemCmdKeyFontType)
}

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
	cfstr := C.CFStringCreateWithBytes(nil, (*C.UInt8)(unsafe.Pointer(cstr)), C.CFIndex(len(str)), C.kCFStringEncodingUTF8, C.Boolean(0))
	C.free(unsafe.Pointer(cstr))
	return cfstr
}

func platformNew(desc Desc) *Font {
	desc.Flags &= UserSettableMask
	if desc.Condensed() && desc.Expanded() {
		desc.Flags &= ^Flags(ExpandedMask)
	}
	font := C.CTFontCreateWithName(cfStringFromString(desc.Family), C.CGFloat(desc.Size), nil)
	if (desc.Flags & UserSettableMask) != 0 {
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
				desc.Flags &= ^Flags(ExpandedMask)
			} else if desc.Condensed() {
				desc.Flags &= ^Flags(CondensedMask)
			} else if desc.Italic() {
				desc.Flags &= ^Flags(ItalicMask)
			} else if desc.Bold() {
				desc.Flags &= ^Flags(BoldMask)
			}
			if (desc.Flags & UserSettableMask) != 0 {
				goto tryAgain
			}
			adjustedFont = font
		}
		font = adjustedFont
	}
	C.CFRetain(font)
	traits := C.CTFontGetSymbolicTraits(font)
	desc.Flags &= ^Flags(CondensedMask | ExpandedMask)
	if traits&C.kCTFontBoldTrait == C.kCTFontBoldTrait {
		desc.Flags |= ActualBoldMask
	}
	if traits&C.kCTFontItalicTrait == C.kCTFontItalicTrait {
		desc.Flags |= ActualItalicMask
	}
	if traits&C.kCTFontCondensedTrait == C.kCTFontCondensedTrait {
		desc.Flags |= CondensedMask
	}
	if traits&C.kCTFontExpandedTrait == C.kCTFontExpandedTrait {
		desc.Flags |= ExpandedMask
	}
	if traits&C.kCTFontMonoSpaceTrait == C.kCTFontMonoSpaceTrait {
		desc.Flags |= MonospacedMask
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

func systemFontDesc(fontType int) Desc {
	font := C.CTFontCreateUIFontForLanguage(C.CTFontUIFontType(fontType), 0, nil)
	C.CFRetain(font)
	traits := C.CTFontGetSymbolicTraits(font)
	var desc Desc
	desc.Family = stringFromCFString(C.CTFontCopyFamilyName(font))
	desc.Size = float32(C.CTFontGetSize(font))
	if traits&C.kCTFontBoldTrait == C.kCTFontBoldTrait {
		desc.Flags |= ActualBoldMask | BoldMask
	}
	if traits&C.kCTFontItalicTrait == C.kCTFontItalicTrait {
		desc.Flags |= ActualItalicMask | ItalicMask
	}
	if traits&C.kCTFontCondensedTrait == C.kCTFontCondensedTrait {
		desc.Flags |= CondensedMask
	}
	if traits&C.kCTFontExpandedTrait == C.kCTFontExpandedTrait {
		desc.Flags |= ExpandedMask
	}
	if traits&C.kCTFontMonoSpaceTrait == C.kCTFontMonoSpaceTrait {
		desc.Flags |= MonospacedMask
	}
	C.CFRelease(font)
	return desc
}
