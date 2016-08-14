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
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/geom"
	"unsafe"
)

// #cgo darwin LDFLAGS: -framework Cocoa
// #include <stdlib.h>
// #include "Cursor.h"
import "C"

// Available system cursors
var (
	Arrow            = newSystemCursor(C.platformArrowID)
	Text             = newSystemCursor(C.platformTextID)
	VerticalText     = newSystemCursor(C.platformVerticalTextID)
	CrossHair        = newSystemCursor(C.platformCrossHairID)
	ClosedHand       = newSystemCursor(C.platformClosedHandID)
	OpenHand         = newSystemCursor(C.platformOpenHandID)
	PointingHand     = newSystemCursor(C.platformPointingHandID)
	ResizeLeft       = newSystemCursor(C.platformResizeLeftID)
	ResizeRight      = newSystemCursor(C.platformResizeRightID)
	ResizeLeftRight  = newSystemCursor(C.platformResizeLeftRightID)
	ResizeUp         = newSystemCursor(C.platformResizeUpID)
	ResizeDown       = newSystemCursor(C.platformResizeDownID)
	ResizeUpDown     = newSystemCursor(C.platformResizeUpDownID)
	DisappearingItem = newSystemCursor(C.platformDisappearingItemID)
	NotAllowed       = newSystemCursor(C.platformNotAllowedID)
	DragLink         = newSystemCursor(C.platformDragLinkID)
	DragCopy         = newSystemCursor(C.platformDragCopyID)
	ContextMenu      = newSystemCursor(C.platformContextMenuID)
)

// Cursor provides a graphical cursor for the mouse location.
type Cursor struct {
	id     C.int
	cursor C.uiCursor
}

func newSystemCursor(id C.int) *Cursor {
	return &Cursor{id: id}
}

// NewCursor creates a new cursor from an image.
func NewCursor(img *draw.Image, hotSpot geom.Point) *Cursor {
	return &Cursor{id: C.platformCustomID, cursor: C.platformNewCursor(img.PlatformPtr(), C.float(hotSpot.X), C.float(hotSpot.Y))}
}

// PlatformPtr returns a pointer to the underlying platform-specific data.
func (cursor *Cursor) PlatformPtr() unsafe.Pointer {
	if cursor.cursor == nil && cursor.id >= 0 && cursor.id < C.platformCustomID {
		cursor.cursor = C.platformSystemCursor(cursor.id)
	}
	return unsafe.Pointer(cursor.cursor)
}

// Dispose of the cursor. Has no effect on system-defined cursors.
func (cursor *Cursor) Dispose() {
	if cursor.id == C.platformCustomID {
		C.platformDisposeCursor(cursor.cursor)
		cursor.cursor = nil
	}
}
