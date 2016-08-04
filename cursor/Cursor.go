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
	Arrow            = newSystemCursor(C.arrowID)
	Text             = newSystemCursor(C.textID)
	VerticalText     = newSystemCursor(C.verticalTextID)
	CrossHair        = newSystemCursor(C.crossHairID)
	ClosedHand       = newSystemCursor(C.closedHandID)
	OpenHand         = newSystemCursor(C.openHandID)
	PointingHand     = newSystemCursor(C.pointingHandID)
	ResizeLeft       = newSystemCursor(C.resizeLeftID)
	ResizeRight      = newSystemCursor(C.resizeRightID)
	ResizeLeftRight  = newSystemCursor(C.resizeLeftRightID)
	ResizeUp         = newSystemCursor(C.resizeUpID)
	ResizeDown       = newSystemCursor(C.resizeDownID)
	ResizeUpDown     = newSystemCursor(C.resizeUpDownID)
	DisappearingItem = newSystemCursor(C.disappearingItemID)
	NotAllowed       = newSystemCursor(C.notAllowedID)
	DragLink         = newSystemCursor(C.dragLinkID)
	DragCopy         = newSystemCursor(C.dragCopyID)
	ContextMenu      = newSystemCursor(C.contextMenuID)
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
	return &Cursor{id: C.customID, cursor: C.newCursor(img.PlatformPtr(), C.float(hotSpot.X), C.float(hotSpot.Y))}
}

// PlatformPtr returns a pointer to the underlying platform-specific data.
func (cursor *Cursor) PlatformPtr() unsafe.Pointer {
	if cursor.cursor == nil && cursor.id >= 0 && cursor.id < C.customID {
		cursor.cursor = C.systemCursor(cursor.id)
	}
	return unsafe.Pointer(cursor.cursor)
}

// Dispose of the cursor. Has no effect on system-defined cursors.
func (cursor *Cursor) Dispose() {
	if cursor.id == C.customID {
		C.disposeCursor(cursor.cursor)
		cursor.cursor = nil
	}
}
