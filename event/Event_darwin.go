// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package event

// #cgo darwin LDFLAGS: -framework Cocoa
// #include "Event_darwin.h"
import "C"

import (
	"sync"
	"time"
)

var (
	dispatchMapLock  sync.Mutex
	nextInvocationID C.uint64_t = 1
	dispatchMap                 = make(map[C.uint64_t]func())
)

// Invoke a task on the UI thread. The task is put into the event queue and will be run at the next
// opportunity.
func Invoke(task func()) {
	C.invoke(recordTask(task))
}

// InvokeAfter schedules a task to be run on the UI thread after waiting for the specified
// duration.
func InvokeAfter(task func(), after time.Duration) {
	C.invokeAfter(recordTask(task), C.int64_t(after.Nanoseconds()))
}

func recordTask(task func()) C.uint64_t {
	dispatchMapLock.Lock()
	defer dispatchMapLock.Unlock()
	id := nextInvocationID
	nextInvocationID++
	dispatchMap[id] = task
	return id
}

func removeTask(id C.uint64_t) func() {
	dispatchMapLock.Lock()
	defer dispatchMapLock.Unlock()
	if f, ok := dispatchMap[id]; ok {
		delete(dispatchMap, id)
		return f
	}
	return nil
}

//export dispatchInvocation
func dispatchInvocation(id C.uint64_t) {
	if f := removeTask(id); f != nil {
		f()
	}
}
