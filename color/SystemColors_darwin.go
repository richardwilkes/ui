package color

import (
	// #cgo CFLAGS: -x objective-c
	// #cgo LDFLAGS: -framework Cocoa
	// #include <AppKit/AppKit.h>
	//
	// unsigned int convertNSColor(NSColor *color) {
	//	CGFloat red, green, blue, alpha;
	//	[[color colorUsingColorSpaceName: NSDeviceRGBColorSpace] getRed:&red green:&green blue:&blue alpha:&alpha];
	//	return ((((unsigned int)(255 * alpha + 0.5)) & 0xFF) << 24) | ((((unsigned int)(255 * red + 0.5)) & 0xFF) << 16) | ((((unsigned int)(255 * green + 0.5)) & 0xFF) << 8) | (((unsigned int)(255 * blue + 0.5)) & 0xFF);
	// }
	//
	// unsigned int platformKeyboardFocusColor() { return convertNSColor([NSColor keyboardFocusIndicatorColor]); }
	// unsigned int platformSelectedTextBackgroundColor() { return convertNSColor([NSColor selectedTextBackgroundColor]); }
	// unsigned int platformSelectedTextColor() { return convertNSColor([NSColor selectedTextColor]); }
	// unsigned int platformTextBackgroundColor() { return convertNSColor([NSColor textBackgroundColor]); }
	// unsigned int platformTextColor() { return convertNSColor([NSColor textColor]); }
	"C"
)

func init() {
	KeyboardFocus = Color(C.platformKeyboardFocusColor())
	SelectedTextBackground = Color(C.platformSelectedTextBackgroundColor())
	SelectedText = Color(C.platformSelectedTextColor())
	TextBackground = Color(C.platformTextBackgroundColor())
	Text = Color(C.platformTextColor())
}
