package x11

import (
	// #cgo pkg-config: x11
	// #include <X11/Xlib.h>
	"C"
	"unsafe"
)

type FocusChangeEvent C.XFocusChangeEvent

func (evt *FocusChangeEvent) Window() Window {
	return Window(evt.window)
}

func (evt *FocusChangeEvent) ToEvent() *Event {
	return (*Event)(unsafe.Pointer(evt))
}
