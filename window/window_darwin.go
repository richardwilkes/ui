package window

import (
	// #cgo CFLAGS: -x objective-c
	// #cgo LDFLAGS: -framework Cocoa -framework Quartz
	// #cgo pkg-config: pangocairo
	// #include "window_darwin.h"
	"C"
	"time"
	"unsafe"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/cursor"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/internal/task"
	"github.com/richardwilkes/ui/keys"
)

// Window represents a window on the display.
type Window commonWindow

type platformWindow unsafe.Pointer

func platformGetKeyWindow() platformWindow {
	return platformWindow(C.getKeyWindow())
}

func platformBringAllWindowsToFront() {
	C.bringAllWindowsToFront()
}

func platformHideCursorUntilMouseMoves() {
	C.hideCursorUntilMouseMoves()
}

func platformNewWindow(bounds geom.Rect, styleMask StyleMask) *Window {
	return &Window{
		window: platformWindow(C.newWindow(C.double(bounds.X), C.double(bounds.Y), C.double(bounds.Width), C.double(bounds.Height), C.int(styleMask))),
	}
}

func platformNewPopupWindow(parent ui.Window, bounds geom.Rect) *Window {
	return platformNewWindow(bounds, BorderlessWindowMask)
}

func (window *Window) platformClose() {
	C.closeWindow(C.platformWindow(window.window))
}

func (window *Window) platformTitle() string {
	return C.GoString(C.getWindowTitle(C.platformWindow(window.window)))
}

func (window *Window) platformSetTitle(title string) {
	cTitle := C.CString(title)
	C.setWindowTitle(C.platformWindow(window.window), cTitle)
	C.free(unsafe.Pointer(cTitle))
}

func (window *Window) platformFrame() geom.Rect {
	var bounds geom.Rect
	C.getWindowFrame(C.platformWindow(window.window), (*C.double)(&bounds.X), (*C.double)(&bounds.Y), (*C.double)(&bounds.Width), (*C.double)(&bounds.Height))
	return bounds
}

func (window *Window) platformSetFrame(bounds geom.Rect) {
	C.setWindowFrame(C.platformWindow(window.window), C.double(bounds.X), C.double(bounds.Y), C.double(bounds.Width), C.double(bounds.Height))
}

func (window *Window) platformContentFrame() geom.Rect {
	var bounds geom.Rect
	C.getWindowContentFrame(C.platformWindow(window.window), (*C.double)(&bounds.X), (*C.double)(&bounds.Y), (*C.double)(&bounds.Width), (*C.double)(&bounds.Height))
	return bounds
}

func (window *Window) platformToFront() {
	C.bringWindowToFront(C.platformWindow(window.window))
}

func (window *Window) platformRepaint(bounds geom.Rect) {
	C.repaintWindow(C.platformWindow(window.window), C.double(bounds.X), C.double(bounds.Y), C.double(bounds.Width), C.double(bounds.Height))
}

func (window *Window) platformFlushPainting() {
	C.flushPainting(C.platformWindow(window.window))
}

func (window *Window) platformMinimize() {
	C.minimizeWindow(C.platformWindow(window.window))
}

func (window *Window) platformZoom() {
	C.zoomWindow(C.platformWindow(window.window))
}

func (window *Window) platformSetCursor(c *cursor.Cursor) {
	C.setCursor(C.platformWindow(window.window), c.PlatformPtr())
}

func (window *Window) platformInvoke(id uint64) {
	C.invoke(C.ulong(id))
}

func (window *Window) platformInvokeAfter(id uint64, after time.Duration) {
	C.invokeAfter(C.ulong(id), C.long(after.Nanoseconds()))
}

//export drawWindow
func drawWindow(cWindow platformWindow, gc *C.cairo_t, x, y, width, height float64) {
	if window, ok := windowMap[cWindow]; ok {
		window.paint(draw.NewGraphics(draw.CairoContext(unsafe.Pointer(gc))), geom.Rect{Point: geom.Point{X: x, Y: y}, Size: geom.Size{Width: width, Height: height}})
	}
}

//export windowResized
func windowResized(cWindow platformWindow) {
	if window, ok := windowMap[cWindow]; ok {
		window.root.SetSize(window.ContentFrame().Size)
	}
}

//export windowGainedKey
func windowGainedKey(cWindow platformWindow) {
	if window, ok := windowMap[cWindow]; ok {
		event.Dispatch(event.NewFocusGained(window))
	}
}

//export windowLostKey
func windowLostKey(cWindow platformWindow) {
	if window, ok := windowMap[cWindow]; ok {
		event.Dispatch(event.NewFocusLost(window))
	}
}

//export windowShouldClose
func windowShouldClose(cWindow platformWindow) bool {
	if window, ok := windowMap[cWindow]; ok {
		return window.MayClose()
	}
	return true
}

//export windowDidClose
func windowDidClose(cWindow platformWindow) {
	if window, ok := windowMap[cWindow]; ok {
		window.Dispose()
	}
}

//export handleMouseDownEvent
func handleMouseDownEvent(cWindow platformWindow, x, y float64, button, clickCount, keyModifiers int) {
	if window, ok := windowMap[cWindow]; ok {
		window.processMouseDown(x, y, button, clickCount, keys.Modifiers(keyModifiers))
	}
}

//export handleMouseDraggedEvent
func handleMouseDraggedEvent(cWindow platformWindow, x, y float64, button, keyModifiers int) {
	if window, ok := windowMap[cWindow]; ok {
		window.processMouseDragged(x, y, button, keys.Modifiers(keyModifiers))
	}
}

//export handleMouseUpEvent
func handleMouseUpEvent(cWindow platformWindow, x, y float64, button, keyModifiers int) {
	if window, ok := windowMap[cWindow]; ok {
		window.processMouseUp(x, y, button, keys.Modifiers(keyModifiers))
	}
}

//export handleMouseEnteredEvent
func handleMouseEnteredEvent(cWindow platformWindow, x, y float64, keyModifiers int) {
	if window, ok := windowMap[cWindow]; ok {
		window.processMouseEntered(x, y, keys.Modifiers(keyModifiers))
	}
}

//export handleMouseMovedEvent
func handleMouseMovedEvent(cWindow platformWindow, x, y float64, keyModifiers int) {
	if window, ok := windowMap[cWindow]; ok {
		window.processMouseMoved(x, y, keys.Modifiers(keyModifiers))
	}
}

//export handleMouseExitedEvent
func handleMouseExitedEvent(cWindow platformWindow, x, y float64, keyModifiers int) {
	if window, ok := windowMap[cWindow]; ok {
		window.processMouseExited(x, y, keys.Modifiers(keyModifiers))
	}
}

//export handleWindowMouseWheelEvent
func handleWindowMouseWheelEvent(cWindow platformWindow, x, y, dx, dy float64, keyModifiers int) {
	if window, ok := windowMap[cWindow]; ok {
		window.processMouseWheel(x, y, dx, dy, keys.Modifiers(keyModifiers))
	}
}

//export handleCursorUpdateEvent
func handleCursorUpdateEvent(cWindow platformWindow, x, y float64, keyModifiers int) {
	if window, ok := windowMap[cWindow]; ok {
		where := geom.Point{X: x, Y: y}
		var widget ui.Widget
		if window.inMouseDown {
			widget = window.lastMouseWidget
		} else {
			widget = window.root.WidgetAt(where)
			if widget == nil {
				panic("widget is nil")
			}
		}
		window.updateToolTipAndCursor(widget, where)
	}
}

//export handleWindowKeyEvent
func handleWindowKeyEvent(cWindow platformWindow, keyCode int, chars *C.char, keyModifiers int, down, repeat bool) {
	if window, ok := windowMap[cWindow]; ok {
		var str string
		if chars != nil {
			str = C.GoString(chars)
		}
		code, ch := keys.Transform(keyCode, str)
		modifiers := keys.Modifiers(keyModifiers)
		if down {
			window.processKeyDown(code, ch, modifiers, repeat)
		} else {
			window.processKeyUp(code, modifiers)
		}
	}
}

//export dispatchTask
func dispatchTask(id uint64) {
	task.Dispatch(id)
}
