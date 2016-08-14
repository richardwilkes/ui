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
	"github.com/richardwilkes/ui/geom"
)

func platformNewImageFromBytes(buffer []byte) *Image {
	// RAW: Implement platformNewImageFromBytes for Linux
	return nil
}

func platformNewImageFromURL(url string) *Image {
	// RAW: Implement platformNewImageFromURL for Linux
	return nil
}

func platformNewImageFromData(data *ImageData) *Image {
	// RAW: Implement platformNewImageFromData for Linux
	return nil
}

func platformNewImageFromImage(other *Image, bounds geom.Rect) *Image {
	// RAW: Implement platformNewImageFromImage for Linux
	return nil
}

func (img *Image) platformDispose() {
	// RAW: Implement platformDispose for Linux
}

func (img *Image) platformData() *ImageData {
	// RAW: Implement platformData for Linux
	return nil
}
