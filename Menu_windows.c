// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#include <stdlib.h>
#include "_cgo_export.h"
#include "Menu.h"


platformMenu platformGetMainMenu() {
	// RAW: Implement platformGetMainMenu for Windows
	return NULL;
}

void platformSetMainMenu(platformMenu menuBar) {
	// RAW: Implement platformSetMainMenu for Windows
}

platformMenu platformNewMenu(const char *title) {
	// RAW: Implement platformNewMenu for Windows
	return NULL;
}

void platformDisposeMenu(platformMenu menu) {
	// RAW: Implement platformDisposeMenu for Windows
}

int platformMenuItemCount(platformMenu menu) {
	// RAW: Implement platformMenuItemCount for Windows
	return 0;
}

platformMenuItem platformGetMenuItem(platformMenu menu, int index) {
	// RAW: Implement platformGetMenuItem for Windows
	return NULL;
}

platformMenuItem platformAddMenuItem(platformMenu menu, const char *title, const char *key) {
	// RAW: Implement platformAddMenuItem for Windows
	return NULL;
}

platformMenuItem platformAddSeparator(platformMenu menu) {
	// RAW: Implement platformAddSeparator for Windows
	return NULL;
}

void platformSetKeyModifierMask(platformMenuItem item, int mask) {
	// RAW: Implement platformSetKeyModifierMask for Windows
}

platformMenu platformGetSubMenu(platformMenuItem item) {
	// RAW: Implement platformGetSubMenu for Windows
	return NULL;
}

void platformSetSubMenu(platformMenuItem item, platformMenu subMenu) {
	// RAW: Implement platformSetSubMenu for Windows
}

void platformSetServicesMenu(platformMenu menu) {
	// RAW: Implement platformSetServicesMenu for Windows
}

void platformSetWindowMenu(platformMenu menu) {
	// RAW: Implement platformSetWindowMenu for Windows
}

void platformSetHelpMenu(platformMenu menu) {
	// RAW: Implement platformSetHelpMenu for Windows
}

void platformPopupMenu(platformWindow window, platformMenu menu, float x, float y, platformMenuItem itemAtLocation) {
	// RAW: Implement platformPopupMenu for Windows
}
