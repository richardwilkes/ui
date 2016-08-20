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

platformMenu platformMenuBar();
void platformSetMenuBar(platformMenu menuBar);
platformMenu platformNewMenu(const char *title);
void platformDisposeMenu(platformMenu menu);
int platformMenuItemCount(platformMenu menu);
platformMenuItem platformGetMenuItem(platformMenu menu, int index);
platformMenuItem platformAddMenuItem(platformMenu menu, const char *title, const char *key);
platformMenuItem platformAddSeparator(platformMenu menu);
void platformSetServicesMenu(platformMenu menu);
void platformSetWindowMenu(platformMenu menu);
void platformSetHelpMenu(platformMenu menu);
void platformPopupMenu(platformWindow window, platformMenu menu, float x, float y, platformMenuItem itemAtLocation);

#endif // __RW_GOUI_MENU__
