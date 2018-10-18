package font

import (
	// #cgo darwin LDFLAGS: -framework Cocoa
	// #include <CoreText/CoreText.h>
	// #include <stdlib.h>
	"C"
	"bytes"
	"fmt"
	"unsafe"
)

func init() {
	User = systemFont(C.kCTFontUIFontUser)
	UserMonospaced = systemFont(C.kCTFontUIFontUserFixedPitch)
	System = systemFont(C.kCTFontUIFontSystem)
	EmphasizedSystem = systemFont(C.kCTFontUIFontEmphasizedSystem)
	SmallSystem = systemFont(C.kCTFontUIFontSmallSystem)
	SmallEmphasizedSystem = systemFont(C.kCTFontUIFontSmallEmphasizedSystem)
	Views = systemFont(C.kCTFontUIFontViews)
	Label = systemFont(C.kCTFontUIFontLabel)
	Menu = systemFont(C.kCTFontUIFontMenuItem)
	MenuCmdKey = systemFont(C.kCTFontUIFontMenuItemCmdKey)
}

func systemFont(fontType C.CTFontUIFontType) *Font {
	var buffer bytes.Buffer
	font := C.CTFontCreateUIFontForLanguage(fontType, 0, C.CFStringRef(C.NULL))
	buffer.WriteString(stringFromCFString(C.CTFontCopyFamilyName(font)))
	traits := C.CTFontGetSymbolicTraits(font)
	if traits&C.kCTFontBoldTrait != 0 {
		buffer.WriteString(" ")
		buffer.WriteString(WeightBold.String())
	}
	if traits&C.kCTFontItalicTrait != 0 {
		buffer.WriteString(" ")
		buffer.WriteString(SlantItalic.String())
	}
	if traits&C.kCTFontCondensedTrait != 0 {
		buffer.WriteString(" ")
		buffer.WriteString(StretchCondensed.String())
	} else if traits&C.kCTFontExpandedTrait != 0 {
		buffer.WriteString(" ")
		buffer.WriteString(StretchExpanded.String())
	}
	buffer.WriteString(" ")
	buffer.WriteString(fmt.Sprintf("%v", C.CTFontGetSize(font)))
	return NewFont(buffer.String())
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
