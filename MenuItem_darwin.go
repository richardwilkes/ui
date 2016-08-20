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
	"github.com/richardwilkes/ui/event"
)

// #cgo darwin LDFLAGS: -framework Cocoa
// #include <stdlib.h>
// #include "MenuItem_darwin.h"
import "C"

func (item *MenuItem) platformSubMenu() platformMenu {
	return platformMenu(C.platformGetSubMenu(item.item))
}

func (item *MenuItem) platformSetSubMenu(subMenu *Menu) {
	C.platformSetSubMenu(item.item, subMenu.menu)
}

func (item *MenuItem) platformSetKeyModifierMask(modifierMask event.KeyMask) {
	C.platformSetKeyModifierMask(item.item, C.int(modifierMask))
}
