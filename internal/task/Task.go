// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package task

import (
	"sync"
)

var (
	nextInvocationID uint64 = 1
	dispatchMapLock  sync.Mutex
	dispatchMap      = make(map[uint64]func())
)

func Record(task func()) uint64 {
	dispatchMapLock.Lock()
	defer dispatchMapLock.Unlock()
	id := nextInvocationID
	nextInvocationID++
	dispatchMap[id] = task
	return id
}

func Dispatch(id uint64) {
	if f := remove(id); f != nil {
		f()
	}
}

func remove(id uint64) func() {
	dispatchMapLock.Lock()
	defer dispatchMapLock.Unlock()
	if f, ok := dispatchMap[id]; ok {
		delete(dispatchMap, id)
		return f
	}
	return nil
}
