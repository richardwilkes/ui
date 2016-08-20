// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

import (
	"C"
	"github.com/richardwilkes/ui/event"
)

//export validateMenuItem
func validateMenuItem(cMenuItem platformMenuItem) bool {
	if item, ok := itemMap[cMenuItem]; ok {
		evt := event.NewValidate(item)
		event.Dispatch(evt)
		return evt.Valid()
	}
	return true
}

//export handleMenuItem
func handleMenuItem(cMenuItem platformMenuItem) {
	if item, ok := itemMap[cMenuItem]; ok {
		event.Dispatch(event.NewSelection(item))
	}
}
