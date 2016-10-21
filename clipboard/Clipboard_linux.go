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
	"github.com/richardwilkes/ui/clipboard/datatypes"
	"github.com/richardwilkes/ui/internal/x11"
)

func platformChangeCount() int {
	return x11.ClipboardChangeCount()
}

func platformClear() {
	x11.ClipboardClear()
}

func platformTypes() []string {
	return x11.ClipboardTypes()
}

func platformGetData(dataType string) []byte {
	return x11.GetClipboard(dataType)
}

func platformSetData(data []datatypes.Data) {
	x11.SetClipboard(data)
}
