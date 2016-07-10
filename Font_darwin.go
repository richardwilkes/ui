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

func stringFromDescAttr(desc C.CTFontDescriptorRef, attrName C.CFStringRef) string {
	str := C.CFStringRef(C.CTFontDescriptorCopyAttribute(desc, attrName))
	result := stringFromCFString(str)
	C.CFRelease(str)
	return result
}

func platformNewFont(desc FontDesc) *Font {
	desc.Flags &= UserSettableStyleMask
	if desc.Condensed() && desc.Expanded() {
		desc.Flags &= ^FontStyleFlags(ExpandedStyleMask)
	}
	font := C.CTFontCreateWithName(cfStringFromString(desc.Family), C.CGFloat(desc.Size), nil)
	if (desc.Flags & UserSettableStyleMask) != 0 {
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
				desc.Flags &= ^FontStyleFlags(ExpandedStyleMask)
			} else if desc.Condensed() {
				desc.Flags &= ^FontStyleFlags(CondensedStyleMask)
			} else if desc.Italic() {
				desc.Flags &= ^FontStyleFlags(ItalicStyleMask)
			} else if desc.Bold() {
				desc.Flags &= ^FontStyleFlags(BoldStyleMask)
			}
			if (desc.Flags & UserSettableStyleMask) != 0 {
				goto tryAgain
			}
			adjustedFont = font
		}
		font = adjustedFont
	}
	C.CFRetain(font)
	traits := C.CTFontGetSymbolicTraits(font)
	desc.Flags &= ^FontStyleFlags(CondensedStyleMask | ExpandedStyleMask)
	if traits&C.kCTFontBoldTrait == C.kCTFontBoldTrait {
		desc.Flags |= ActualBoldStyleMask
	}
	if traits&C.kCTFontItalicTrait == C.kCTFontItalicTrait {
		desc.Flags |= ActualItalicStyleMask
	}
	if traits&C.kCTFontCondensedTrait == C.kCTFontCondensedTrait {
		desc.Flags |= CondensedStyleMask
	}
	if traits&C.kCTFontExpandedTrait == C.kCTFontExpandedTrait {
		desc.Flags |= ExpandedStyleMask
	}
	if traits&C.kCTFontMonoSpaceTrait == C.kCTFontMonoSpaceTrait {
		desc.Flags |= MonospacedStyleMask
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

func platformFontDesc(id int) FontDesc {
	var fontType C.int
	switch id {
	case userFontID:
		fontType = C.kCTFontUserFontType
	case userMonospacedFontID:
		fontType = C.kCTFontUserFixedPitchFontType
	case systemFontID:
		fontType = C.kCTFontSystemFontType
	case emphasizedSystemFontID:
		fontType = C.kCTFontEmphasizedSystemFontType
	case smallSystemFontID:
		fontType = C.kCTFontSmallSystemFontType
	case smallEmphasizedSystemFontID:
		fontType = C.kCTFontSmallEmphasizedSystemFontType
	case viewsFontID:
		fontType = C.kCTFontViewsFontType
	case labelFontID:
		fontType = C.kCTFontLabelFontType
	case menuFontID:
		fontType = C.kCTFontMenuItemFontType
	case menuCmdKeyFontID:
		fontType = C.kCTFontMenuItemCmdKeyFontType
	default:
		fontType = C.kCTFontUserFontType
	}
	font := C.CTFontCreateUIFontForLanguage(C.CTFontUIFontType(fontType), 0, nil)
	C.CFRetain(font)
	traits := C.CTFontGetSymbolicTraits(font)
	var desc FontDesc
	desc.Family = stringFromCFString(C.CTFontCopyFamilyName(font))
	desc.Size = float32(C.CTFontGetSize(font))
	if traits&C.kCTFontBoldTrait == C.kCTFontBoldTrait {
		desc.Flags |= ActualBoldStyleMask | BoldStyleMask
	}
	if traits&C.kCTFontItalicTrait == C.kCTFontItalicTrait {
		desc.Flags |= ActualItalicStyleMask | ItalicStyleMask
	}
	if traits&C.kCTFontCondensedTrait == C.kCTFontCondensedTrait {
		desc.Flags |= CondensedStyleMask
	}
	if traits&C.kCTFontExpandedTrait == C.kCTFontExpandedTrait {
		desc.Flags |= ExpandedStyleMask
	}
	if traits&C.kCTFontMonoSpaceTrait == C.kCTFontMonoSpaceTrait {
		desc.Flags |= MonospacedStyleMask
	}
	C.CFRelease(font)
	return desc
}
