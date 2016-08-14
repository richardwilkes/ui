// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package font

// #include <stdlib.h>
import "C"

func platformAvailableFontFamilies() []string {
	// RAW: Implement platformAvailableFontFamilies for Windows
	var names []string
	return names
}

func platformNewFont(desc Desc) *Font {
	// RAW: Implement platformNewFont for Windows
	return &Font{font: nil, desc: desc}
}

func (f *Font) platformDispose() {
	// RAW: Implement platformDispose for Windows
}

func (f *Font) platformAscent() float32 {
	// RAW: Implement platformAscent for Windows
	return 0
}

func (f *Font) platformDescent() float32 {
	// RAW: Implement platformDescent for Windows
	return 0
}

func (f *Font) platformLeading() float32 {
	// RAW: Implement platformLeading for Windows
	return 0
}

func (f *Font) platformWidth(str string) float32 {
	// RAW: Implement platformWidth for Windows
	return 0
}

func (f *Font) platformIndexForPosition(x float32, str string) int {
	// RAW: Implement platformIndexForPosition for Windows
	return 0
}

func (f *Font) platformPositionForIndex(index int, str string) float32 {
	// RAW: Implement platformPositionForIndex for Windows
	return 0
}

func platformDesc(id int) Desc {
	// RAW: Implement platformDesc for Windows
	return Desc{}
}
