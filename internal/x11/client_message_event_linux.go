package x11

import (
	// #cgo pkg-config: x11
	// #include <X11/Xlib.h>
	"C"
	"unsafe"
)

type ClientMessageEvent C.XClientMessageEvent

func NewClientMessageEvent(wnd Window, subType Atom, format int) *ClientMessageEvent {
	evt := &ClientMessageEvent{_type: C.ClientMessage, window: C.Window(wnd), message_type: C.Atom(subType), format: C.int(format)}
	return evt
}

func (evt *ClientMessageEvent) Window() Window {
	return Window(evt.window)
}

func (evt *ClientMessageEvent) Format() int {
	return int(evt.format)
}

func (evt *ClientMessageEvent) SubType() Atom {
	return Atom(evt.message_type)
}

func (evt *ClientMessageEvent) Protocol() Atom {
	return *(*Atom)(unsafe.Pointer(&evt.data))
}

func (evt *ClientMessageEvent) TaskID() uint64 {
	return *(*uint64)(unsafe.Pointer(&evt.data))
}

func (evt *ClientMessageEvent) ToEvent() *Event {
	return (*Event)(unsafe.Pointer(evt))
}
