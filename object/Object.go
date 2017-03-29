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
	nextID uint64
)

// Object defines what a generic object must implement.
type Object interface {
	// ID returns a unique ID for the object.
	ID() uint64
	// Self returns this object.
	Self() interface{}
	// Is returns true if this object is the passed-in object.
	Is(other Object) bool
}

// Base holds the base information for Objects.
type Base struct {
	id   uint64
	self interface{}
}

// InitTypeAndID initializes the object with the appropriate self-identification information.
func (obj *Base) InitTypeAndID(self interface{}) {
	obj.id = atomic.AddUint64(&nextID, 1)
	obj.self = self
}

// ID returns a unique ID for the object.
func (obj *Base) ID() uint64 {
	if obj.id == 0 {
		panic("InitTypeAndID() must be called before use")
	}
	return obj.id
}

// Self returns this object.
func (obj *Base) Self() interface{} {
	return obj.self
}

// Is returns true if this object is the passed-in object.
func (obj *Base) Is(other Object) bool {
	return obj.id == other.ID()
}
