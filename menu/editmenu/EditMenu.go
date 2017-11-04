// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package editmenu

import (
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/ui/menu"
)

// Install adds a standard 'Edit' menu to the end of the menu bar.
func Install(bar menu.Bar) menu.Menu {
	editMenu := menu.NewMenu(i18n.Text("Edit"))

	AppendCutItem(editMenu)
	AppendCopyItem(editMenu)
	AppendPasteItem(editMenu)

	editMenu.AppendItem(menu.NewSeparator())
	AppendDeleteItem(editMenu)
	AppendSelectAllItem(editMenu)

	bar.AppendMenu(editMenu)
	return editMenu
}
