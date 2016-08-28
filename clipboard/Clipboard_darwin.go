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
	// #cgo darwin LDFLAGS: -framework Cocoa
	// #include <stdlib.h>
	// #include "Clipboard_darwin.h"
	"C"
)

func platformChangeCount() int {
	return int(C.platformClipboardChangeCount())
}

func platformClear() {
	C.platformClearClipboard()
}

func platformTypes() []string {
	clipTypes := (*[1 << 30]*C.char)(unsafe.Pointer(C.platformClipboardTypes()))
	i := 0
	for clipTypes[i] != nil {
		i++
	}
	types := make([]string, i)
	for j := 0; j < i; j++ {
		types[j] = C.GoString(clipTypes[j])
	}
	C.free(unsafe.Pointer(clipTypes))
	return types
}

func platformGetData(dataType string) []byte {
	cstr := C.CString(dataType)
	data := C.platformClipboardData(cstr)
	C.free(unsafe.Pointer(cstr))
	count := int(data.count)
	result := make([]byte, count)
	bytes := (*[1 << 30]byte)(unsafe.Pointer(data.data))
	for i := 0; i < count; i++ {
		result[i] = bytes[i]
	}
	return result
}

func platformSetData(dataType string, bytes []byte) {
	cstr := C.CString(dataType)
	C.platformSetClipboardData(cstr, C.int(len(bytes)), unsafe.Pointer(&bytes[0]))
	C.free(unsafe.Pointer(cstr))
}
