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
	nextID uint64 = 1
	lock   sync.Mutex
	tasks  = make(map[uint64]func())
)

func Record(task func()) uint64 {
	if task == nil {
		panic("nil task not permitted")
	}
	lock.Lock()
	id := nextID
	nextID++
	tasks[id] = task
	lock.Unlock()
	return id
}

func Dispatch(id uint64) {
	lock.Lock()
	task := tasks[id]
	if task != nil {
		delete(tasks, id)
	}
	lock.Unlock()
	if task != nil {
		task()
	}
}
