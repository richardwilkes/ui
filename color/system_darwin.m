#include "system_darwin.h"

unsigned int convertNSColor(NSColor *color) {
	CGFloat red, green, blue, alpha;
	[[color colorUsingColorSpace:[NSColorSpace deviceRGBColorSpace]] getRed:&red green:&green blue:&blue alpha:&alpha];
	return ((((unsigned int)(255 * alpha + 0.5)) & 0xFF) << 24) |
			((((unsigned int)(255 * red + 0.5)) & 0xFF) << 16) |
			((((unsigned int)(255 * green + 0.5)) & 0xFF) << 8) |
			(((unsigned int)(255 * blue + 0.5)) & 0xFF);
}

unsigned int keyboardFocusColor() {
	return convertNSColor([NSColor keyboardFocusIndicatorColor]);
}

unsigned int selectedTextBackgroundColor() {
	return convertNSColor([NSColor selectedTextBackgroundColor]);
}

unsigned int selectedTextColor() {
	return convertNSColor([NSColor selectedTextColor]);
}

unsigned int textBackgroundColor() {
	return convertNSColor([NSColor textBackgroundColor]);
}

unsigned int textColor() {
	return convertNSColor([NSColor textColor]);
}
