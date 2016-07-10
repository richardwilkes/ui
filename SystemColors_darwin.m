// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#include <AppKit/AppKit.h>
#include "SystemColors.h"

unsigned int uiGetSystemColor(SystemColorId id) {
	NSColorList *list;
	NSColor *nsColor = nil;
	unsigned int color = 0;
	switch (id) {
		case backgroundColor:
			color = 0xFFECECEC;
			break;
		case keyboardFocusColor:
			nsColor = [NSColor keyboardFocusIndicatorColor];
			break;
		case selectedControlColor:
			nsColor = [NSColor alternateSelectedControlColor];
			break;
		case selectedControlTextColor:
			nsColor = [NSColor alternateSelectedControlTextColor];
			break;
		case selectedTextBackgroundColor:
			nsColor = [NSColor selectedTextBackgroundColor];
			break;
		case selectedTextColor:
			nsColor = [NSColor selectedTextColor];
			break;
		case textBackgroundColor:
			nsColor = [NSColor textBackgroundColor];
			break;
		case textColor:
			nsColor = [NSColor textColor];
			break;
		default:
			break;
	}
	if (nsColor != nil) {
		CGFloat red, green, blue, alpha;
		[[nsColor colorUsingColorSpaceName: NSDeviceRGBColorSpace] getRed:&red green:&green blue:&blue alpha:&alpha];
		color = ((((unsigned int)(255 * alpha + 0.5)) & 0xFF) << 24) | ((((unsigned int)(255 * red + 0.5)) & 0xFF) << 16) | ((((unsigned int)(255 * green + 0.5)) & 0xFF) << 8) | (((unsigned int)(255 * blue + 0.5)) & 0xFF);
	}
	return color;
}
