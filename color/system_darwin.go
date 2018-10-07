package color

import (
	// #cgo CFLAGS: -x objective-c
	// #cgo LDFLAGS: -framework Cocoa
	// #include "system_darwin.h"
	"C"
)

func init() {
	KeyboardFocus = Color(C.keyboardFocusColor())
	SelectedTextBackground = Color(C.selectedTextBackgroundColor())
	SelectedText = Color(C.selectedTextColor())
	TextBackground = Color(C.textBackgroundColor())
	Text = Color(C.textColor())
}
