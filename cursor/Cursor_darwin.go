// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package cursor

import (
	"fmt"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/draw"
	"unsafe"
	// #cgo darwin LDFLAGS: -framework Cocoa
	// #include <stdlib.h>
	// #include <CoreFoundation/CoreFoundation.h>
	// #include <CoreGraphics/CoreGraphics.h>
	// #include <ImageIO/ImageIO.h>
	// #include "Cursor_darwin.h"
	"C"
)

func platformSystemCursor(id int) unsafe.Pointer {
	switch id {
	case arrow:
		return C.platformArrow()
	case text:
		return C.platformText()
	case verticalText:
		return C.platformVerticalText()
	case crossHair:
		return C.platformCrossHair()
	case closedHand:
		return C.platformClosedHand()
	case openHand:
		return C.platformOpenHand()
	case pointingHand:
		return C.platformPointingHand()
	case resizeLeft:
		return C.platformResizeLeft()
	case resizeRight:
		return C.platformResizeRight()
	case resizeLeftRight:
		return C.platformResizeLeftRight()
	case resizeUp:
		return C.platformResizeUp()
	case resizeDown:
		return C.platformResizeDown()
	case resizeUpDown:
		return C.platformResizeUpDown()
	case disappearingItem:
		return C.platformDisappearingItem()
	case notAllowed:
		return C.platformNotAllowed()
	case dragLink:
		return C.platformDragLink()
	case dragCopy:
		return C.platformDragCopy()
	case contextMenu:
		return C.platformContextMenu()
	default:
		panic(fmt.Sprintf("Invalid system cursor ID (%d)", id))
	}
}

func platformNewCursor(imgData *draw.ImageData, hotSpot geom.Point) unsafe.Pointer {
	colorspace := C.CGColorSpaceCreateWithName(C.kCGColorSpaceGenericRGB)
	defer C.CGColorSpaceRelease(colorspace)

	length := len(imgData.Pixels)
	size := C.size_t(length * 4)
	buffer := C.malloc(size)
	C.memcpy(buffer, unsafe.Pointer(&imgData.Pixels[0]), size)
	pixels := (*[1 << 30]color.Color)(buffer)

	// Perform alpha pre-multiplication, since macOS requires it
	for i := 0; i < length; i++ {
		pixels[i] = pixels[i].Premultiply()
	}

	provider := C.CGDataProviderCreateWithData(nil, buffer, size, nil)
	defer C.free(buffer)
	defer C.CGDataProviderRelease(provider)
	image := C.CGImageCreate(C.size_t(imgData.Width), C.size_t(imgData.Height), 8, 32, C.size_t(imgData.Width*4), colorspace, C.kCGBitmapByteOrder32Host|C.kCGImageAlphaPremultipliedFirst, provider, nil, false, C.kCGRenderingIntentDefault)

	return C.platformNewCursor(unsafe.Pointer(image), C.float(hotSpot.X), C.float(hotSpot.Y))
}

func platformDisposeCursor(cursor *Cursor) {
	C.platformDisposeCursor(cursor.cursor)
}
