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
	"github.com/richardwilkes/ui/draw"
	"unsafe"
)

func platformSystemCursor(id int) unsafe.Pointer {
	// RAW: Implement for Windows
	switch id {
	case arrow:
		return nil
	case text:
		return nil
	case verticalText:
		return nil
	case crossHair:
		return nil
	case closedHand:
		return nil
	case openHand:
		return nil
	case pointingHand:
		return nil
	case resizeLeft:
		return nil
	case resizeRight:
		return nil
	case resizeLeftRight:
		return nil
	case resizeUp:
		return nil
	case resizeDown:
		return nil
	case resizeUpDown:
		return nil
	case disappearingItem:
		return nil
	case notAllowed:
		return nil
	case dragLink:
		return nil
	case dragCopy:
		return nil
	case contextMenu:
		return nil
	default:
		panic(fmt.Sprintf("Invalid system cursor ID (%d)", id))
	}
}

func platformNewCursor(imgData *draw.ImageData, hotSpot geom.Point) unsafe.Pointer {
	// RAW: Implement for Windows
	return nil
}

func platformDisposeCursor(cursor *Cursor) {
	// RAW: Implement for Windows
}
