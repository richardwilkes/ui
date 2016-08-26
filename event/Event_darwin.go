// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package event

import (
	"time"
	"unsafe"
	// #include <stdlib.h>
	// void invoke(uint64_t id);
	// void invokeAfter(uint64_t id, int64_t afterNanos);
	"C"
)

func platformInvoke(id uint64) {
	C.invoke(C.uint64_t(id))
}

func platformInvokeAfter(id uint64, after time.Duration) {
	C.invokeAfter(C.uint64_t(id), C.int64_t(after.Nanoseconds()))
}

//export dispatchInvocation
func dispatchInvocation(id *uint64) {
	C.free(unsafe.Pointer(id))
	DispatchInvocation(*id)
}
