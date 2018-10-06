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
	"sort"
	"sync"
	"unsafe"

	// #cgo pkg-config: pangocairo
	// #include <pango/pangocairo.h>
	"C"
)

var (
	familiesOnce sync.Once
	familiesAll  []*Family
	familiesMono []*Family
)

// Families returns the installed font families.
func Families() []*Family {
	familiesOnce.Do(initFamilyLists)
	return familiesAll
}

// MonospacedFamilies returns the installed font families that are monospaced.
func MonospacedFamilies() []*Family {
	familiesOnce.Do(initFamilyLists)
	return familiesMono
}

func initFamilyLists() {
	var list **C.PangoFontFamily
	var count C.int
	C.pango_font_map_list_families(C.pango_cairo_font_map_get_default(), &list, &count)
	fontFamilies := (*[1 << 30]*C.PangoFontFamily)(unsafe.Pointer(list))
	familiesAll = make([]*Family, count)
	var i C.int
	for i = 0; i < count; i++ {
		familiesAll[i] = &Family{family: fontFamilies[i]}
		if familiesAll[i].Monospaced() {
			familiesMono = append(familiesMono, familiesAll[i])
		}
	}
	C.g_free(C.gpointer(list))
	sort.Sort(familiesByName(familiesAll))
	sort.Sort(familiesByName(familiesMono))
}

// Family represents a family of related font faces. The faces in a family share a common design,
// but differ in slant, weight, width, and other aspects.
type Family struct {
	family *C.PangoFontFamily
}

// Name returns the name of the family.
func (f *Family) Name() string {
	return C.GoString(C.pango_font_family_get_name(f.family))
}

// String returns the name of the family.
func (f *Family) String() string {
	return f.Name()
}

// Monospaced returns true if the family has a fixed width.
func (f *Family) Monospaced() bool {
	return C.pango_font_family_is_monospace(f.family) != 0
}

// Faces returns the faces that make up the family. The faces share a common design, but differ in
// slant, weight, width, and other aspects.
func (f *Family) Faces() []*Face {
	var list **C.PangoFontFace
	var count C.int
	C.pango_font_family_list_faces(f.family, &list, &count)
	fontFaces := (*[1 << 30]*C.PangoFontFace)(unsafe.Pointer(list))
	faces := make([]*Face, count)
	var i C.int
	for i = 0; i < count; i++ {
		faces[i] = &Face{face: fontFaces[i]}
	}
	C.g_free(C.gpointer(list))
	sort.Sort(facesByName(faces))
	return faces
}

type familiesByName []*Family

// Len implements sort.Interface
func (f familiesByName) Len() int {
	return len(f)
}

// Swap implements sort.Interface
func (f familiesByName) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

// Less implements sort.Interface
func (f familiesByName) Less(i, j int) bool {
	return f[i].Name() < f[j].Name()
}

type facesByName []*Face

// Len implements sort.Interface
func (f facesByName) Len() int {
	return len(f)
}

// Swap implements sort.Interface
func (f facesByName) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

// Less implements sort.Interface
func (f facesByName) Less(i, j int) bool {
	return f[i].Name() < f[j].Name()
}
