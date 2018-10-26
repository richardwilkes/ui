package window

import (
	"time"
	"unsafe"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/cursor"
)

// Window represents a window on the display.
type Window commonWindow

type platformWindow unsafe.Pointer

func platformGetKeyWindow() platformWindow {
	// RAW: Implement for Windows
	return nil
}

func platformBringAllWindowsToFront() {
	// RAW: Implement for Windows
}

func platformHideCursorUntilMouseMoves() {
	// RAW: Implement for Windows
}

func platformNewWindow(bounds geom.Rect, styleMask StyleMask) *Window {
	// RAW: Implement for Windows
	return nil
}

func platformNewPopupWindow(parent ui.Window, bounds geom.Rect) *Window {
	// RAW: Implement for Windows
	return nil
}

func (window *Window) platformClose() {
	// RAW: Implement for Windows
	window.Dispose()
}

func (window *Window) platformTitle() string {
	// RAW: Implement for Windows
	return ""
}

func (window *Window) platformSetTitle(title string) {
	// RAW: Implement for Windows
}

func (window *Window) platformFrame() geom.Rect {
	// RAW: Implement for Windows
	return geom.Rect{}
}

func (window *Window) platformSetFrame(bounds geom.Rect) {
	// RAW: Implement for Windows
}

func (window *Window) platformContentFrame() geom.Rect {
	// RAW: Implement for Windows
	return geom.Rect{}
}

func (window *Window) platformToFront() {
	// RAW: Implement for Windows
}

func (window *Window) platformRepaint(bounds geom.Rect) {
	// RAW: Implement for Windows
}

func (window *Window) platformFlushPainting() {
	// RAW: Implement for Windows
}

func (window *Window) platformMinimize() {
	// RAW: Implement for Windows
}

func (window *Window) platformZoom() {
	// RAW: Implement for Windows
}

func (window *Window) platformSetCursor(c *cursor.Cursor) {
	// RAW: Implement for Windows
}

func (window *Window) platformInvoke(id uint64) {
	// RAW: Implement for Windows
}

func (window *Window) platformInvokeAfter(id uint64, after time.Duration) {
	// RAW: Implement for Windows
}
