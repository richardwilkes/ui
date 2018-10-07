package x11

import (
	"unsafe"
	// #cgo pkg-config: x11
	// #include <X11/Xlib.h>
	"C"
)

type PropertyEvent C.XPropertyEvent

func (evt *PropertyEvent) Time() C.Time {
	return evt.time
}

func (evt *PropertyEvent) ToEvent() *Event {
	return (*Event)(unsafe.Pointer(evt))
}
