// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package draw

import (
	"github.com/richardwilkes/geom"
	// #cgo pkg-config: pangocairo
	// #include <pango/pangocairo.h>
	"C"
)

func toCairoMatrix(matrix *geom.Matrix) *C.cairo_matrix_t {
	return &C.cairo_matrix_t{xx: C.double(matrix.XX), yx: C.double(matrix.YX), xy: C.double(matrix.XY), yy: C.double(matrix.YY), x0: C.double(matrix.X0), y0: C.double(matrix.Y0)}
}

func fromCairoMatrix(matrix *C.cairo_matrix_t) *geom.Matrix {
	return &geom.Matrix{XX: float64(matrix.xx), YX: float64(matrix.yx), XY: float64(matrix.xy), YY: float64(matrix.yy), X0: float64(matrix.x0), Y0: float64(matrix.y0)}
}
