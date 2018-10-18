package x11

import (
	// #cgo pkg-config: x11
	// #include <X11/Xlib.h>
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui/keys"
)

type CrossingEvent C.XCrossingEvent

func (evt *CrossingEvent) Window() Window {
	return Window(evt.window)
}

func (evt *CrossingEvent) Where() geom.Point {
	return geom.Point{X: float64(evt.x), Y: float64(evt.y)}
}

func (evt *CrossingEvent) Modifiers() keys.Modifiers {
	return Modifiers(evt.state)
}

func (evt *CrossingEvent) ToEvent() *Event {
	return (*Event)(unsafe.Pointer(evt))
}
