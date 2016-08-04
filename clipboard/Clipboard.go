// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package clipboard

import (
	"unsafe"
)

// #cgo darwin LDFLAGS: -framework Cocoa
// #include <stdlib.h>
// #include "Clipboard.h"
import "C"

var (
	lastChangeCount = -1
	types           []string
)

// Clear the clipboard contents and prepare it for calls to clipboard.SetData.
func Clear() {
	C.clearClipboard()
}

// HasType returns true if the specified data type exists on the clipboard.
func HasType(dataType string) bool {
	for _, one := range Types() {
		if one == dataType {
			return true
		}
	}
	return false
}

// Types returns the types of data currently on the clipboard.
func Types() []string {
	changeCount := int(C.clipboardChangeCount())
	if changeCount != lastChangeCount {
		lastChangeCount = changeCount
		clipTypes := (*[1 << 30]*C.char)(unsafe.Pointer(C.clipboardTypes()))
		i := 0
		for clipTypes[i] != nil {
			i++
		}
		types = make([]string, i)
		for j := 0; j < i; j++ {
			types[j] = C.GoString(clipTypes[j])
		}
		C.free(unsafe.Pointer(clipTypes))
	}
	return types
}

// Data returns the bytes associated with the specified data type on the clipboard. An empty byte
// slice will be returned if no such data type is present.
func Data(dataType string) []byte {
	cstr := C.CString(dataType)
	data := C.clipboardData(cstr)
	C.free(unsafe.Pointer(cstr))
	count := int(data.count)
	result := make([]byte, count)
	bytes := (*[1 << 30]byte)(unsafe.Pointer(data.data))
	for i := 0; i < count; i++ {
		result[i] = bytes[i]
	}
	return result
}

// SetData sets the bytes associated with a particular data type. To provide multiple flavors, first
// call clipboard.Clear() followed by calls to clipboard.SetData() with each flavor of data.
func SetData(dataType string, bytes []byte) {
	cstr := C.CString(dataType)
	C.setClipboardData(cstr, C.int(len(bytes)), unsafe.Pointer(&bytes[0]))
	C.free(unsafe.Pointer(cstr))
}
