package x11

import (
	"unsafe"

	// #cgo pkg-config: x11
	// #include <X11/Xlib.h>
	"C"
)

type SelectionClearEvent C.XSelectionClearEvent

func (evt *SelectionClearEvent) Window() Window {
	return Window(evt.window)
}

func (evt *SelectionClearEvent) Selection() C.Atom {
	return evt.selection
}

func (evt *SelectionClearEvent) When() C.Time {
	return evt.time
}

func (evt *SelectionClearEvent) ToEvent() *Event {
	return (*Event)(unsafe.Pointer(evt))
}
