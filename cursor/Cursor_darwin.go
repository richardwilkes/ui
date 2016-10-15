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
	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/draw"
	"unsafe"
	// #cgo CFLAGS: -x objective-c
	// #cgo LDFLAGS: -framework Cocoa
	// #include <Cocoa/Cocoa.h>
	//
	// void *ArrowCursor() { return [NSCursor arrowCursor]; }
	// void *TextCursor() { return [NSCursor IBeamCursor]; }
	// void *VerticalTextCursor() { return [NSCursor IBeamCursorForVerticalLayout]; }
	// void *CrossHairCursor() { return [NSCursor crosshairCursor]; }
	// void *ClosedHandCursor() { return [NSCursor closedHandCursor]; }
	// void *OpenHandCursor() { return [NSCursor openHandCursor]; }
	// void *PointingHandCursor() { return [NSCursor pointingHandCursor]; }
	// void *ResizeLeftCursor() { return [NSCursor resizeLeftCursor]; }
	// void *ResizeRightCursor() { return [NSCursor resizeRightCursor]; }
	// void *ResizeLeftRightCursor() { return [NSCursor resizeLeftRightCursor]; }
	// void *ResizeUpCursor() { return [NSCursor resizeUpCursor]; }
	// void *ResizeDownCursor() { return [NSCursor resizeDownCursor]; }
	// void *ResizeUpDownCursor() { return [NSCursor resizeUpDownCursor]; }
	// void *DisappearingItemCursor() { return [NSCursor disappearingItemCursor]; }
	// void *NotAllowedCursor() { return [NSCursor operationNotAllowedCursor]; }
	// void *DragLinkCursor() { return [NSCursor dragLinkCursor]; }
	// void *DragCopyCursor() { return [NSCursor dragCopyCursor]; }
	// void *ContextMenuCursor() { return [NSCursor contextualMenuCursor]; }
	// void *NewCursor(void *img, float hotX, float hotY) { return [[[NSCursor alloc] initWithImage:img hotSpot:NSMakePoint(hotX,hotY)] retain]; }
	// void DisposeCursor(void *cursor) { [((NSCursor *)cursor) release]; }
	"C"
)

func platformSystemCursor(id int) unsafe.Pointer {
	switch id {
	case arrow:
		return C.ArrowCursor()
	case text:
		return C.TextCursor()
	case verticalText:
		return C.VerticalTextCursor()
	case crossHair:
		return C.CrossHairCursor()
	case closedHand:
		return C.ClosedHandCursor()
	case openHand:
		return C.OpenHandCursor()
	case pointingHand:
		return C.PointingHandCursor()
	case resizeLeft:
		return C.ResizeLeftCursor()
	case resizeRight:
		return C.ResizeRightCursor()
	case resizeLeftRight:
		return C.ResizeLeftRightCursor()
	case resizeUp:
		return C.ResizeUpCursor()
	case resizeDown:
		return C.ResizeDownCursor()
	case resizeUpDown:
		return C.ResizeUpDownCursor()
	case disappearingItem:
		return C.DisappearingItemCursor()
	case notAllowed:
		return C.NotAllowedCursor()
	case dragLink:
		return C.DragLinkCursor()
	case dragCopy:
		return C.DragCopyCursor()
	case contextMenu:
		return C.ContextMenuCursor()
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

	return C.NewCursor(unsafe.Pointer(image), C.float(hotSpot.X), C.float(hotSpot.Y))
}

func platformDisposeCursor(cursor *Cursor) {
	C.DisposeCursor(cursor.cursor)
}
