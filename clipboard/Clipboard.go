// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package clipboard

var (
	lastChangeCount = -1
	dataTypes       []string
)

// Clear the clipboard contents and prepare it for calls to clipboard.SetData.
func Clear() {
	platformClear()
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
	changeCount := platformChangeCount()
	if changeCount != lastChangeCount {
		lastChangeCount = changeCount
		dataTypes = platformTypes()
	}
	return dataTypes
}

// Data returns the bytes associated with the specified data type on the clipboard. An empty byte
// slice will be returned if no such data type is present.
func Data(dataType string) []byte {
	return platformGetData(dataType)
}

// SetData sets the bytes associated with a particular data type. To provide multiple flavors, first
// call clipboard.Clear() followed by calls to clipboard.SetData() with each flavor of data.
func SetData(dataType string, bytes []byte) {
	platformSetData(dataType, bytes)
}
