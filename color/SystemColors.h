// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#ifndef __TROLLWORKS_SYSTEM_COLORS__
#define __TROLLWORKS_SYSTEM_COLORS__

typedef enum {
	backgroundColor,
	keyboardFocusColor,
	selectedControlColor,
	selectedControlTextColor,
	selectedTextBackgroundColor,
	selectedTextColor,
	textBackgroundColor,
	textColor
} SystemColorId;

// The returned color is in the format 0xAARRGGBB
unsigned int uiGetSystemColor(SystemColorId id);

#endif // __TROLLWORKS_SYSTEM_COLORS__
