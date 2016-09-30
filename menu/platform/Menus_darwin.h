// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#ifndef __RW_GOUI_MENU__
#define __RW_GOUI_MENU__

#include "Types.h"

Menu platformBar();
void platformSetBar(Menu bar);
Menu platformNewMenu(const char *title);
Item platformNewSeparator();
Item platformNewItem(const char *title, const char *key, int modifiers);
void platformDisposeItem(Item item);
void platformDisposeMenu(Menu menu);
Menu platformSubMenu(Item item);
void platformSetSubMenu(Item item, Menu subMenu);
int platformItemCount(Menu menu);
Item platformItem(Menu menu, int index);
void platformAddItem(Menu menu, Item item);
void platformInsertItem(Menu menu, Item item, int index);
void platformRemove(Menu menu, int index);
void platformSetServicesMenu(Menu menu);
void platformSetWindowMenu(Menu menu);
void platformSetHelpMenu(Menu menu);

#endif // __RW_GOUI_MENU__
