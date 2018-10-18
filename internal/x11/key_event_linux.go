package x11

import (
	// #cgo pkg-config: x11
	// #include <X11/Xutil.h>
	"C"
	"unsafe"

	"github.com/richardwilkes/ui/keys"
)

type KeyEvent C.XKeyEvent

func (evt *KeyEvent) Window() Window {
	return Window(evt.window)
}

func (evt *KeyEvent) Modifiers() keys.Modifiers {
	return Modifiers(evt.state)
}

func (evt *KeyEvent) CodeAndChar() (code int, ch rune) {
	var buffer [5]C.char
	var keySym C.KeySym
	buffer[C.XLookupString((*C.XKeyEvent)(evt), &buffer[0], C.int(len(buffer)-1), &keySym, nil)] = 0
	code, ch = keys.Transform(int(keySym), C.GoString(&buffer[0]))
	return
}

func (evt *KeyEvent) When() C.Time {
	return evt.time
}

func (evt *KeyEvent) ToEvent() *Event {
	return (*Event)(unsafe.Pointer(evt))
}
