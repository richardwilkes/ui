package draw

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/xmath/geom"

	// #cgo pkg-config: cairo
	// #include <cairo.h>
	"C"
)

// CairoContentType holds the type of content for a surface.
type CairoContentType C.cairo_content_t

// Possible CairoContentTypes.
const (
	ColorContent         CairoContentType = C.CAIRO_CONTENT_COLOR
	AlphaContent         CairoContentType = C.CAIRO_CONTENT_ALPHA
	ColorAndAlphaContent CairoContentType = C.CAIRO_CONTENT_COLOR_ALPHA
)

// Surface holds the content of a drawing surface.
type Surface struct {
	surface *C.cairo_surface_t
	size    geom.Size
}

// Size returns the size in pixels of the surface.
func (surface *Surface) Size() geom.Size {
	return surface.size
}

// Destroy a surface.
func (surface *Surface) Destroy() {
	C.cairo_surface_destroy(surface.surface)
}

// NewCairoContext creates a new CairoContext with the specified bounds.
func (surface *Surface) NewCairoContext(bounds geom.Rect) CairoContext {
	return CairoContext(unsafe.Pointer(C.cairo_create(surface.surface)))
}

// CreateSimilar creates a new surface similar to this surface, but with the specified
// content type and size.
func (surface *Surface) CreateSimilar(contentType CairoContentType, size geom.Size) *Surface {
	return &Surface{surface: C.cairo_surface_create_similar(surface.surface, C.cairo_content_t(contentType), C.int(size.Width), C.int(size.Height)), size: size}
}
