package x11

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/xmath/geom"
	// #cgo pkg-config: x11
	// #include <X11/Xlib.h>
	"C"
)

type ExposeEvent C.XExposeEvent

func NewExposeEvent(wnd Window, bounds geom.Rect) *ExposeEvent {
	return &ExposeEvent{_type: C.Expose, window: C.Window(wnd), x: C.int(bounds.X), y: C.int(bounds.Y), width: C.int(bounds.Width), height: C.int(bounds.Height)}
}

func (evt *ExposeEvent) Window() Window {
	return Window(evt.window)
}

func (evt *ExposeEvent) Bounds() geom.Rect {
	return geom.Rect{Point: geom.Point{X: float64(evt.x), Y: float64(evt.y)}, Size: geom.Size{Width: float64(evt.width), Height: float64(evt.height)}}
}

func (evt *ExposeEvent) ToEvent() *Event {
	return (*Event)(unsafe.Pointer(evt))
}
