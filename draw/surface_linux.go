package draw

import (
	// #cgo pkg-config: pangocairo
	// #include <pango/pangocairo.h>
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/xmath/geom"
)

func NewSurface(surface unsafe.Pointer, size geom.Size) *Surface {
	return &Surface{surface: (*C.cairo_surface_t)(surface), size: size}
}

func (surface *Surface) SetSize(size geom.Size) {
	if surface.size != size {
		surface.size = size
		C.cairo_xlib_surface_set_size(surface.surface, C.int(size.Width), C.int(size.Height))
	}
}
