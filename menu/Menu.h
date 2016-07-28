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

typedef void *uiMenu;
typedef void *uiMenuItem;
typedef void *uiWindow;

uiMenu getMainMenu();
void setMainMenu(uiMenu menuBar);
uiMenu uiNewMenu(const char *title);
void uiDisposeMenu(uiMenu menu);
int uiMenuItemCount(uiMenu menu);
uiMenuItem uiGetMenuItem(uiMenu menu, int index);
uiMenuItem uiAddMenuItem(uiMenu menu, const char *title, const char *key);
uiMenuItem uiAddSeparator(uiMenu menu);
void uiSetKeyModifierMask(uiMenuItem item, int mask);
uiMenu uiGetSubMenu(uiMenuItem item);
void uiSetSubMenu(uiMenuItem item, uiMenu subMenu);
void uiSetServicesMenu(uiMenu menu);
void uiSetWindowMenu(uiMenu menu);
void uiSetHelpMenu(uiMenu menu);
void uiPopupMenu(uiWindow window, uiMenu menu, float x, float y, uiMenuItem itemAtLocation);

#endif // __RW_GOUI_MENU__
