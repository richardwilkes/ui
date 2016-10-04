// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package helpmenu

import (
	"github.com/richardwilkes/i18n"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/menu/factory"
)

// Install adds a standard 'Help' menu to the menu bar.
func Install(index int) menu.Menu {
	helpMenu := factory.NewMenu(i18n.Text("Help"))
	factory.SetHelpMenu(helpMenu)
	factory.AppBar().InsertMenu(helpMenu, index)
	return helpMenu
}
