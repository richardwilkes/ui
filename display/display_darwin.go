package display

import (
	// #cgo CFLAGS: -x objective-c
	// #cgo LDFLAGS: -framework Cocoa
	// #include "display_darwin.h"
	"C"

	"github.com/richardwilkes/toolbox/xmath/geom"
)

func platformMainDisplayBounds() geom.Rect {
	var bounds geom.Rect
	C.getMainDisplayBounds((*C.double)(&bounds.X), (*C.double)(&bounds.Y), (*C.double)(&bounds.Width), (*C.double)(&bounds.Height))
	return bounds
}
