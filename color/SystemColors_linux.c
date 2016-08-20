// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#include "SystemColors.h"

unsigned int platformSystemColor(SystemColorId id) {
	// RAW: Implement platformSystemColor for Linux
	unsigned int color = 0;
	switch (id) {
		case platformBackgroundColor:
			color = 0xFFECECEC;
			break;
		case platformKeyboardFocusColor:
			color = 0xFFCCCCFF;
			break;
		case platformSelectedControlColor:
			color = 0xFFEEEEFF;
			break;
		case platformSelectedControlTextColor:
			color = 0xFF000000;
			break;
		case platformSelectedTextBackgroundColor:
			color = 0xFF8888FF;
			break;
		case platformSelectedTextColor:
			color = 0xFFFFFFFF;
			break;
		case platformTextBackgroundColor:
			color = 0xFFFFFFFF;
			break;
		case platformTextColor:
			color = 0xFF000000;
			break;
		default:
			break;
	}
	return color;
}
