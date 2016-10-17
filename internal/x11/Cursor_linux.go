// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package x11

import (
	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui/draw"
	"unsafe"
	// #cgo LDFLAGS: -lXcursor
	// #cgo pkg-config: x11
	// #include <stdlib.h>
	// #include <X11/Xlib.h>
	// #include <X11/cursorfont.h>
	// #include <X11/Xcursor/Xcursor.h>
	"C"
)

type SystemCursorID int

const (
	ArrowID            SystemCursorID = C.XC_left_ptr
	TextID             SystemCursorID = C.XC_xterm
	VerticalTextID     SystemCursorID = -1
	CrossHairID        SystemCursorID = C.XC_crosshair
	ClosedHandID       SystemCursorID = -2
	OpenHandID         SystemCursorID = -3
	PointingHandID     SystemCursorID = -4
	ResizeLeftID       SystemCursorID = C.XC_left_side
	ResizeRightID      SystemCursorID = C.XC_right_side
	ResizeLeftRightID  SystemCursorID = -5
	ResizeUpID         SystemCursorID = C.XC_top_side
	ResizeDownID       SystemCursorID = C.XC_bottom_side
	ResizeUpDownID     SystemCursorID = C.XC_double_arrow
	DisappearingItemID SystemCursorID = -6
	NotAllowedID       SystemCursorID = -7
	DragLinkID         SystemCursorID = -8
	DragCopyID         SystemCursorID = -9
	ContextMenuID      SystemCursorID = -10
)

type Cursor C.Cursor

var (
	cursors           = make(map[SystemCursorID]Cursor)
	cursorTheme       *C.char
	defaultCursorSize C.int
)

func getCursorTheme() *C.char {
	if cursorTheme == nil {
		cursorTheme = C.XcursorGetTheme(display)
	}
	return cursorTheme
}

func getDefaultCursorSize() C.int {
	if defaultCursorSize == 0 {
		defaultCursorSize = C.XcursorGetDefaultSize(display)
	}
	return defaultCursorSize
}

func SystemCursor(id SystemCursorID) Cursor {
	if cursor, ok := cursors[id]; ok {
		return cursor
	}

	var name string
	switch id {
	case ArrowID:
		name = "arrow"
	case TextID:
		name = "xterm"
	case VerticalTextID:
		// Couldn't find an equivalent
	case CrossHairID:
		name = "crosshair"
	case ClosedHandID:
		name = "grabbing"
	case OpenHandID:
		// Couldn't find an equivalent
	case PointingHandID:
		name = "hand"
	case ResizeLeftID:
		name = "left_side"
	case ResizeRightID:
		name = "right_side"
	case ResizeLeftRightID:
		name = "h_double_arrow"
	case ResizeUpID:
		name = "top_side"
	case ResizeDownID:
		name = "bottom_side"
	case ResizeUpDownID:
		name = "v_double_arrow"
	case DisappearingItemID:
		// Couldn't find an equivalent
	case NotAllowedID:
		name = "crossed_circle"
	case DragLinkID:
		name = "dnd-link"
	case DragCopyID:
		name = "dnd-copy"
	case ContextMenuID:
		// Couldn't find an equivalent
	}

	var cursor Cursor
	if name != "" {
		cstr := C.CString(name)
		ci := C.XcursorLibraryLoadImage(cstr, getCursorTheme(), getDefaultCursorSize())
		if ci != nil {
			cursor = Cursor(C.XcursorImageLoadCursor(display, ci))
			C.XcursorImageDestroy(ci)
		}
		C.free(unsafe.Pointer(cstr))
	}
	if cursor == 0 {
		if id < 0 {
			id = C.XC_left_ptr
		}
		cursor = Cursor(C.XcursorShapeLoadCursor(display, C.uint(id)))
	}

	cursors[id] = cursor
	return cursor
}

func NewCursor(imgData *draw.ImageData, hotSpot geom.Point) Cursor {
	ci := C.XcursorImageCreate(C.int(imgData.Width), C.int(imgData.Height))
	ci.xhot = C.XcursorDim(hotSpot.X)
	ci.yhot = C.XcursorDim(hotSpot.Y)
	ci.pixels = (*C.XcursorPixel)(&imgData.Pixels[0])
	cursor := C.XcursorImageLoadCursor(display, ci)
	C.XcursorImageDestroy(ci)
	return Cursor(cursor)
}

func (cursor Cursor) Dispose() {
	C.XFreeCursor(display, C.Cursor(cursor))
}
