package x11

import (
	// #cgo pkg-config: x11
	// #include <X11/Xlib.h>
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui/keys"
)

type MotionEvent C.XMotionEvent

func (evt *MotionEvent) Window() Window {
	return Window(evt.window)
}

func (evt *MotionEvent) Where() geom.Point {
	return geom.Point{X: float64(evt.x), Y: float64(evt.y)}
}

func (evt *MotionEvent) Modifiers() keys.Modifiers {
	return Modifiers(evt.state)
}

func (evt *MotionEvent) ToEvent() *Event {
	return (*Event)(unsafe.Pointer(evt))
}
