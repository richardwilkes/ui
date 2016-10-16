// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package macmenus

import (
	// typedef void *Item;
	"C"
	"github.com/richardwilkes/ui/event"
)

//export validateMenuItemCallback
func validateMenuItemCallback(menuItem C.Item) bool {
	if item, ok := itemMap[menuItem]; ok {
		evt := event.NewValidate(item)
		event.Dispatch(evt)
		item.enabled = evt.Valid()
		return item.enabled
	}
	return true
}

//export handleMenuItemCallback
func handleMenuItemCallback(menuItem C.Item) {
	if item, ok := itemMap[menuItem]; ok {
		event.Dispatch(event.NewSelection(item))
	}
}
