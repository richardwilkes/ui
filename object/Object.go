// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package object

import (
	"sync/atomic"
)

var (
	nextID uint64 = 0
)

type Object interface {
	ID() uint64
	Self() interface{}
	Is(other Object) bool
}

type Base struct {
	id   uint64
	self interface{}
}

func (obj *Base) InitTypeAndID(self interface{}) {
	obj.id = atomic.AddUint64(&nextID, 1)
	obj.self = self
}

func (obj *Base) ID() uint64 {
	if obj.id == 0 {
		panic("InitTypeAndID() must be called before use")
	}
	return obj.id
}

func (obj *Base) Self() interface{} {
	return obj.self
}

func (obj *Base) Is(other Object) bool {
	return obj.id == other.ID()
}
