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
	"github.com/richardwilkes/geom"
	"unsafe"
)

const (
	custom = iota
	arrow
	text
	verticalText
	crossHair
	closedHand
	openHand
	pointingHand
	resizeLeft
	resizeRight
	resizeLeftRight
	resizeUp
	resizeDown
	resizeUpDown
	disappearingItem
	notAllowed
	dragLink
	dragCopy
	contextMenu
)

// Available system cursors
var (
	Arrow            = &Cursor{id: arrow}
	Text             = &Cursor{id: text}
	VerticalText     = &Cursor{id: verticalText}
	CrossHair        = &Cursor{id: crossHair}
	ClosedHand       = &Cursor{id: closedHand}
	OpenHand         = &Cursor{id: openHand}
	PointingHand     = &Cursor{id: pointingHand}
	ResizeLeft       = &Cursor{id: resizeLeft}
	ResizeRight      = &Cursor{id: resizeRight}
	ResizeLeftRight  = &Cursor{id: resizeLeftRight}
	ResizeUp         = &Cursor{id: resizeUp}
	ResizeDown       = &Cursor{id: resizeDown}
	ResizeUpDown     = &Cursor{id: resizeUpDown}
	DisappearingItem = &Cursor{id: disappearingItem}
	NotAllowed       = &Cursor{id: notAllowed}
	DragLink         = &Cursor{id: dragLink}
	DragCopy         = &Cursor{id: dragCopy}
	ContextMenu      = &Cursor{id: contextMenu}
)

// Cursor provides a graphical cursor for the mouse location.
type Cursor struct {
	id     int
	cursor unsafe.Pointer
}

// NewCursor creates a new cursor from an image.
func NewCursor(imgData *draw.ImageData, hotSpot geom.Point) *Cursor {
	return &Cursor{id: custom, cursor: platformNewCursor(imgData, hotSpot.X, hotSpot.Y)}
}

// PlatformPtr returns a pointer to the underlying platform-specific data.
func (cursor *Cursor) PlatformPtr() unsafe.Pointer {
	if cursor.id != custom && cursor.cursor == nil {
		cursor.cursor = platformSystemCursor(cursor.id)
	}
	return cursor.cursor
}

// Dispose of the cursor. Has no effect on system-defined cursors.
func (cursor *Cursor) Dispose() {
	if cursor.id == custom {
		platformDisposeCursor(cursor)
		cursor.cursor = nil
	}
}
