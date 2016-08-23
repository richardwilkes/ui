// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

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
