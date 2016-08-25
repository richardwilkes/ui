// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package menu

func platformMenuBar() platformMenu {
	// RAW: Implement for Linux
	return nil
}

func platformNewMenu(title string) platformMenu {
	// RAW: Implement for Linux
	return nil
}

func (menu *Menu) platformItem(index int) platformMenuItem {
	// RAW: Implement for Linux
	return nil
}

func (menu *Menu) platformAddItem(title string, key string) platformMenuItem {
	// RAW: Implement for Linux
	return nil
}

func (menu *Menu) platformAddSeparator() platformMenuItem {
	// RAW: Implement for Linux
	return nil
}

func (menu *Menu) platformCount() int {
	// RAW: Implement for Linux
	return 0
}

func (menu *Menu) platformSetAsMenuBar() {
	// RAW: Implement for Linux
}

func (menu *Menu) platformSetAsServicesMenu() {
	// RAW: Implement for Linux
}

func (menu *Menu) platformSetAsWindowMenu() {
	// RAW: Implement for Linux
}

func (menu *Menu) platformSetAsHelpMenu() {
	// RAW: Implement for Linux
}

func (menu *Menu) platformDispose() {
	// RAW: Implement for Linux
}
