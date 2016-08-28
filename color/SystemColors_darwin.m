// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#include <AppKit/AppKit.h>
#include "SystemColors_darwin.h"

unsigned int convertNSColor(NSColor *color) {
	CGFloat red, green, blue, alpha;
	[[color colorUsingColorSpaceName: NSDeviceRGBColorSpace] getRed:&red green:&green blue:&blue alpha:&alpha];
	return ((((unsigned int)(255 * alpha + 0.5)) & 0xFF) << 24) | ((((unsigned int)(255 * red + 0.5)) & 0xFF) << 16) | ((((unsigned int)(255 * green + 0.5)) & 0xFF) << 8) | (((unsigned int)(255 * blue + 0.5)) & 0xFF);
}

unsigned int platformKeyboardFocusColor() {
	return convertNSColor([NSColor keyboardFocusIndicatorColor]);
}

unsigned int platformSelectedTextBackgroundColor() {
	return convertNSColor([NSColor selectedTextBackgroundColor]);
}

unsigned int platformSelectedTextColor() {
	return convertNSColor([NSColor selectedTextColor]);
}

unsigned int platformTextBackgroundColor() {
	return convertNSColor([NSColor textBackgroundColor]);
}

unsigned int platformTextColor() {
	return convertNSColor([NSColor textColor]);
}
