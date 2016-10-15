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
	// #cgo CFLAGS: -x objective-c
	// #cgo LDFLAGS: -framework Cocoa
	// #include <Cocoa/Cocoa.h>
	//
	// struct clipboardData {
	//	int count;
	//	const void *data;
	// };
	//
	// int clipboardChangeCount() {
	//	return [[NSPasteboard generalPasteboard] changeCount];
	// }
	//
	// void clearClipboard() {
	//	[[NSPasteboard generalPasteboard] clearContents];
	// }
	//
	// const char **clipboardTypes() {
	//	NSArray<NSString *> *types = [[NSPasteboard generalPasteboard] types];
	//	NSUInteger count = [types count];
	//	const char **result = malloc(sizeof(char *) * (count + 1));
	//	result[count] = NULL;
	//	for (int i = 0; i < count; i++) {
	//		result[i] = [[types objectAtIndex:i] UTF8String];
	//	}
	//	return result;
	// }
	//
	// struct clipboardData getClipboardData(char *type) {
	//	struct clipboardData d;
	//	NSData *nsd = [[NSPasteboard generalPasteboard] dataForType:[NSString stringWithUTF8String:type]];
	//	d.count = [nsd length];
	//	d.data = [nsd bytes];
	//	return d;
	// }
	//
	// void setClipboardData(char *type, int size, void *bytes) {
	//	[[NSPasteboard generalPasteboard] setData:[NSData dataWithBytes:bytes length:size] forType:[NSString stringWithUTF8String:type]];
	// }
	"C"
)

func platformChangeCount() int {
	return int(C.clipboardChangeCount())
}

func platformClear() {
	C.clearClipboard()
}

func platformTypes() []string {
	clipTypes := (*[1 << 30]*C.char)(unsafe.Pointer(C.clipboardTypes()))
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
	data := C.getClipboardData(cstr)
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
	C.setClipboardData(cstr, C.int(len(bytes)), unsafe.Pointer(&bytes[0]))
	C.free(unsafe.Pointer(cstr))
}
