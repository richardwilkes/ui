// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#include <Cocoa/Cocoa.h>
#include "Menu_darwin.h"
#include "_cgo_export.h"

platformMenu platformGetSubMenu(platformItem item) {
    NSMenuItem *mitem = (NSMenuItem *)item;
    if ([mitem hasSubmenu]) {
        return (platformMenu)[mitem submenu];
    }
    return nil;
}

void platformSetSubMenu(platformItem item, platformMenu subMenu) {
    [((NSMenuItem *)item) setSubmenu: subMenu];
}

void platformSetKeyModifierMask(platformItem item, int mask) {
	// macOS uses the same modifier mask bit order as we do, but it is shifted up by 16 bits
    [((NSMenuItem *)item) setKeyEquivalentModifierMask:mask << 16];
}
