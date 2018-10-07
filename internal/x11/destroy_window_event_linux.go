package x11

import (
	"unsafe"

	// #cgo pkg-config: x11
	// #include <X11/Xlib.h>
	"C"
)

type DestroyWindowEvent C.XDestroyWindowEvent

func (evt *DestroyWindowEvent) Window() Window {
	return Window(evt.window)
}

func (evt *DestroyWindowEvent) ToEvent() *Event {
	return (*Event)(unsafe.Pointer(evt))
}
