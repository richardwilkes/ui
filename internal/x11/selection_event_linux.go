package x11

import (
	// #cgo pkg-config: x11
	// #include <X11/Xlib.h>
	"C"
	"unsafe"
)

type SelectionEvent C.XSelectionEvent

func (evt *SelectionEvent) Requestor() Window {
	return Window(evt.requestor)
}

func (evt *SelectionEvent) Selection() Atom {
	return Atom(evt.selection)
}

func (evt *SelectionEvent) Target() Atom {
	return Atom(evt.target)
}

func (evt *SelectionEvent) Property() Atom {
	return Atom(evt.property)
}

func (evt *SelectionEvent) When() C.Time {
	return evt.time
}

func (evt *SelectionEvent) ToEvent() *Event {
	return (*Event)(unsafe.Pointer(evt))
}
