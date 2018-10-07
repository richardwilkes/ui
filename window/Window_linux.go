package window

import (
	"time"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/cursor"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/internal/x11"
)

// Window represents a window on the display.
type Window struct {
	commonWindow
	surface   *draw.Surface
	wasMapped bool
}

type platformWindow x11.Window

var (
	blankCursor *cursor.Cursor
)

func platformGetKeyWindow() platformWindow {
	return platformWindow(uintptr(x11.InputFocus()))
}

func platformBringAllWindowsToFront() {
	list := Windows()
	for i := len(list) - 1; i >= 0; i-- {
		list[i].ToFront()
	}
}

func platformHideCursorUntilMouseMoves() {
	if window := KeyWindow(); window != nil {
		if blankCursor == nil {
			blankCursor = cursor.NewCursor(&draw.ImageData{Width: 1, Height: 1, Pixels: make([]color.Color, 1)}, geom.Point{})
		}
		window.SetCursor(blankCursor)
	}
}

func platformNewWindow(bounds geom.Rect, styleMask StyleMask) *Window {
	wnd := x11.NewWindow(bounds)
	return &Window{
		commonWindow: commonWindow{window: platformWindow(uintptr(wnd))},
		surface:      wnd.NewSurface(bounds.Size),
	}
}

func platformNewPopupWindow(parent ui.Window, bounds geom.Rect) *Window {
	wnd := x11.NewPopupWindow(x11.Window(uintptr(parent.PlatformPtr())), bounds)
	return &Window{
		commonWindow: commonWindow{window: platformWindow(uintptr(wnd))},
		surface:      wnd.NewSurface(bounds.Size),
	}
}

func (window *Window) toXWindow() x11.Window {
	return x11.Window(uintptr(window.window))
}

func (window *Window) platformClose() {
	window.surface.Destroy()
	window.toXWindow().Destroy()
	window.Dispose()
}

func (window *Window) platformTitle() string {
	return window.toXWindow().Title()
}

func (window *Window) platformSetTitle(title string) {
	window.toXWindow().SetTitle(title)
}

func (window *Window) frameDecorationSpace() (top, left, bottom, right float64) {
	if window.Valid() {
		return window.toXWindow().FrameDecorationSpace()
	}
	return
}

func (window *Window) platformFrame() geom.Rect {
	bounds := window.platformContentFrame()
	top, left, bottom, right := window.frameDecorationSpace()
	bounds.X -= left
	bounds.Y -= top
	bounds.Width += left + right
	bounds.Height += top + bottom
	return bounds
}

func (window *Window) platformSetFrame(bounds geom.Rect) {
	window.toXWindow().SetFrame(bounds)
}

func (window *Window) platformContentFrame() geom.Rect {
	if window.Valid() {
		return window.toXWindow().ContentFrame()
	}
	return geom.Rect{}
}

func (window *Window) platformToFront() {
	wnd := window.toXWindow()
	if window.wasMapped {
		wnd.Raise()
	} else {
		wnd.Show()
		// Wait for window to be mapped
		for {
			if event := wnd.NextEventOfType(x11.MapNotifyType); event != nil {
				window.wasMapped = true
				wnd.Move(window.initialLocationRequest)
				if window.owner == nil {
					// Wait for window to be configured so that we have correct placement information
					for {
						if event = wnd.NextEventOfType(x11.ConfigureNotifyType); event != nil {
							processConfigureEvent(event.ToConfigureEvent())
							break
						}
						time.Sleep(time.Millisecond * 10)
					}
				}
				// This is here so that menu windows behave properly
				wnd.RequestFocus()
				break
			}
			time.Sleep(time.Millisecond * 10)
		}
	}
}

func (window *Window) platformRepaint(bounds geom.Rect) {
	window.toXWindow().Repaint(bounds)
}

func (window *Window) draw(bounds geom.Rect) {
	buffer := window.surface.CreateSimilar(draw.ColorContent, window.surface.Size())
	gc := draw.NewGraphics(buffer.NewCairoContext(bounds))
	gc.Rect(bounds)
	gc.Clip()
	window.paint(gc, bounds)
	gc.Dispose()

	gc = draw.NewGraphics(window.surface.NewCairoContext(bounds))
	gc.Rect(bounds)
	gc.Clip()
	gc.SetSurface(buffer, 0, 0)
	gc.FillClip()
	gc.Dispose()

	buffer.Destroy()
}

func (window *Window) platformFlushPainting() {
	x11.Flush()
}

func (window *Window) platformMinimize() {
	window.toXWindow().Minimize()
}

func (window *Window) platformZoom() {
	window.toXWindow().Zoom()
}

func (window *Window) platformSetCursor(c *cursor.Cursor) {
	window.toXWindow().SetCursor(x11.Cursor(uintptr(c.PlatformPtr())))
}

func (window *Window) platformInvoke(id uint64) {
	if window.Valid() {
		window.toXWindow().InvokeTask(id)
	}
}

func (window *Window) platformInvokeAfter(id uint64, after time.Duration) {
	time.AfterFunc(after, func() {
		window.platformInvoke(id)
	})
}
