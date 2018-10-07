package cursor

import (
	"fmt"
	"unsafe"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/internal/x11"
)

func platformSystemCursor(id int) unsafe.Pointer {
	var cursorID x11.SystemCursorID
	switch id {
	case arrow:
		cursorID = x11.ArrowID
	case text:
		cursorID = x11.TextID
	case verticalText:
		cursorID = x11.VerticalTextID
	case crossHair:
		cursorID = x11.CrossHairID
	case closedHand:
		cursorID = x11.ClosedHandID
	case openHand:
		cursorID = x11.OpenHandID
	case pointingHand:
		cursorID = x11.PointingHandID
	case resizeLeft:
		cursorID = x11.ResizeLeftID
	case resizeRight:
		cursorID = x11.ResizeRightID
	case resizeLeftRight:
		cursorID = x11.ResizeLeftRightID
	case resizeUp:
		cursorID = x11.ResizeUpID
	case resizeDown:
		cursorID = x11.ResizeDownID
	case resizeUpDown:
		cursorID = x11.ResizeUpDownID
	case disappearingItem:
		cursorID = x11.DisappearingItemID
	case notAllowed:
		cursorID = x11.NotAllowedID
	case dragLink:
		cursorID = x11.DragLinkID
	case dragCopy:
		cursorID = x11.DragCopyID
	case contextMenu:
		cursorID = x11.ContextMenuID
	default:
		panic(fmt.Sprintf("Invalid system cursor ID (%d)", id))
	}
	return unsafe.Pointer(uintptr(x11.SystemCursor(cursorID)))
}

func platformNewCursor(imgData *draw.ImageData, hotSpot geom.Point) unsafe.Pointer {
	return unsafe.Pointer(uintptr(x11.NewCursor(imgData, hotSpot)))
}

func platformDisposeCursor(cursor *Cursor) {
	x11.Cursor(uintptr(cursor.cursor)).Dispose()
	cursor.cursor = nil
}
