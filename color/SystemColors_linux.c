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
			break;
		case platformSelectedControlColor:
			break;
		case platformSelectedControlTextColor:
			break;
		case platformSelectedTextBackgroundColor:
			break;
		case platformSelectedTextColor:
			break;
		case platformTextBackgroundColor:
			break;
		case platformTextColor:
			break;
		default:
			break;
	}
	return color;
}
