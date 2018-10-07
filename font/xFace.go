package font

import (
	// #cgo pkg-config: pangocairo
	// #include <pango/pangocairo.h>
	"C"
)

// Face represents a group of fonts with the same family, slant, weight, and width.
type Face struct {
	face *C.PangoFontFace
}

// Name returns the name of the face.
func (f *Face) Name() string {
	return C.GoString(C.pango_font_face_get_face_name(f.face))
}

// String returns the name of the face.
func (f *Face) String() string {
	return f.Name()
}

// Synthesized returns true if the font synthesized from another variant.
func (f *Face) Synthesized() bool {
	return C.pango_font_face_is_synthesized(f.face) != 0
}

// Font returns the font representing this face. No size will have been set for the returned font.
func (f *Face) Font() *Font {
	return &Font{pfd: C.pango_font_face_describe(f.face)}
}
