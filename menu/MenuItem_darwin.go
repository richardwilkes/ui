// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package menu

import (
	"github.com/richardwilkes/ui/keys"
)

// #cgo darwin LDFLAGS: -framework Cocoa
// #include "Menu_darwin.h"
import "C"

func (item *Item) platformSubMenu() PlatformMenu {
	return PlatformMenu(C.platformGetSubMenu(item.item))
}

func (item *Item) platformSetSubMenu(subMenu *Menu) {
	C.platformSetSubMenu(item.item, subMenu.menu)
}

func (item *Item) platformSetKeyModifierMask(modifiers keys.Modifiers) {
	C.platformSetKeyModifierMask(item.item, C.int(modifiers))
}
