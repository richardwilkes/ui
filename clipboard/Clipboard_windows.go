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
	clipData            = make(map[string][]byte)
	clipDataChangeCount int
)

func platformChangeCount() int {
	// RAW: Implement for Windows (i.e. cross-app support)
	return clipDataChangeCount
}

func platformClear() {
	// RAW: Implement for Windows (i.e. cross-app support)
	clipData = make(map[string][]byte)
	clipDataChangeCount++
}

func platformTypes() []string {
	// RAW: Implement for Windows (i.e. cross-app support)
	types := make([]string, len(clipData))
	i := 0
	for key := range clipData {
		types[i++] = key
	}
	return types
}

func platformGetData(dataType string) []byte {
	// RAW: Implement for Windows (i.e. cross-app support)
	return clipData[dataType]
}

func platformSetData(dataType string, bytes []byte) {
	// RAW: Implement for Windows (i.e. cross-app support)
	clipData[dataType] = bytes
	clipDataChangeCount++
}
