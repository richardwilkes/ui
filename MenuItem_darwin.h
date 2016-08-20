// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#ifndef __RW_GOUI_MENUITEM__
#define __RW_GOUI_MENUITEM__

#include "Types.h"

platformMenu platformGetSubMenu(platformMenuItem item);
void platformSetSubMenu(platformMenuItem item, platformMenu subMenu);
void platformSetKeyModifierMask(platformMenuItem item, int mask);

#endif // __RW_GOUI_MENUITEM__
