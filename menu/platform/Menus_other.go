// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// +build !darwin

package platform

func platformBar() cMenu {
	return nil
}

func platformSetBar(bar cMenu) {
}

func platformNewMenu(title string) cMenu {
	return nil
}

func platformNewSeparator() cItem {
	return nil
}

func platformNewItem(title string, keyCode int, modifiers keys.Modifiers) cItem {
	return nil
}

func (menu *platformMenu) platformDispose() {
}

func (menu *platformMenu) platformItemCount() int {
	return 0
}

func (menu *platformMenu) platformItem(index int) cItem {
	return nil
}

func (menu *platformMenu) platformAddItem(item cItem) {
}

func (menu *platformMenu) platformInsertItem(index int, item cItem) {
}

func (menu *platformMenu) platformRemove(index int) {
}

func (item *platformItem) platformDispose() {
}

func (item *platformItem) platformSubMenu() cMenu {
	return nil
}

func (item *platformItem) platformSetSubMenu(subMenu cMenu) {
}

func SetServicesMenu(menu menu.Menu) {
}

func SetWindowMenu(menu menu.Menu) {
}

func SetHelpMenu(menu menu.Menu) {
}
