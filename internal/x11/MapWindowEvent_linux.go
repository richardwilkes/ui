package x11

import (
	"unsafe"
	// #cgo pkg-config: x11
	// #include <X11/Xlib.h>
	"C"
)

type MapWindowEvent C.XMapEvent

func (evt *MapWindowEvent) Window() Window {
	return Window(evt.window)
}

func (evt *MapWindowEvent) ToEvent() *Event {
	return (*Event)(unsafe.Pointer(evt))
}
