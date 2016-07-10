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

func toCGRect(bounds Rect) C.CGRect {
	return C.CGRectMake(C.CGFloat(bounds.X), C.CGFloat(bounds.Y), C.CGFloat(bounds.Width), C.CGFloat(bounds.Height))
}
