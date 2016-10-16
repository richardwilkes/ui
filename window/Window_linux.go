// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package window

import (
	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/cursor"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/internal/x11"
	"time"
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
	// RAW: Implement for Linux
}

func platformNewWindow(bounds geom.Rect, styleMask WindowStyleMask) (window platformWindow, surface platformSurface) {
	wnd := x11.NewWindow(bounds)
	return platformWindow(uintptr(wnd)), platformSurface(x11.NewSurface(wnd, bounds.Size))
}

func platformNewMenuWindow(parent ui.Window, bounds geom.Rect) (window platformWindow, surface platformSurface) {
	wnd := x11.NewMenuWindow(x11.Window(uintptr(parent.PlatformPtr())), bounds)
	return platformWindow(uintptr(wnd)), platformSurface(x11.NewSurface(wnd, bounds.Size))
}

func (window *Wnd) toXWindow() x11.Window {
	return x11.Window(uintptr(window.window))
}

func (window *Wnd) platformClose() {
	(*x11.Surface)(window.surface).Destroy()
	window.toXWindow().Destroy()
	window.Dispose()
}

func (window *Wnd) platformTitle() string {
	return window.toXWindow().Title()
}

func (window *Wnd) platformSetTitle(title string) {
	window.toXWindow().SetTitle(title)
}

func (window *Wnd) frameDecorationSpace() (top, left, bottom, right float64) {
	if window.Valid() {
		return window.toXWindow().FrameDecorationSpace()
	}
	return
}

func (window *Wnd) platformFrame() geom.Rect {
	bounds := window.platformContentFrame()
	top, left, bottom, right := window.frameDecorationSpace()
	bounds.X -= left
	bounds.Y -= top
	bounds.Width += left + right
	bounds.Height += top + bottom
	return bounds
}

func (window *Wnd) platformSetFrame(bounds geom.Rect) {
	window.toXWindow().SetFrame(bounds)
}

func (window *Wnd) platformContentFrame() geom.Rect {
	if window.Valid() {
		return window.toXWindow().ContentFrame()
	}
	return geom.Rect{}
}

func (window *Wnd) platformToFront() {
	wnd := window.toXWindow()
	if window.wasMapped {
		wnd.Raise()
	} else {
		wnd.Show()
		// Wait for window to be mapped
		for {
			if event := x11.NextEventOfTypeForWindow(x11.MapWindowType, wnd); event != nil {
				window.wasMapped = true
				wnd.Move(window.initialLocationRequest)
				if window.owner == nil {
					// Wait for window to be configured so that we have correct placement information
					for {
						if event = x11.NextEventOfTypeForWindow(x11.ConfigureType, wnd); event != nil {
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

func (window *Wnd) platformRepaint(bounds geom.Rect) {
	window.toXWindow().Repaint(bounds)
}

func (window *Wnd) draw(bounds geom.Rect) {
	gc := draw.NewGraphics((*x11.Surface)(window.surface).NewCairoContext(bounds))
	gc.Rect(bounds)
	gc.Clip()
	window.paint(gc, bounds)
	gc.Dispose()
}

func (window *Wnd) platformFlushPainting() {
	x11.Flush()
}

func (window *Wnd) platformScalingFactor() float64 {
	// RAW: Implement for Linux
	return 1
}

func (window *Wnd) platformMinimize() {
	window.toXWindow().Minimize()
}

func (window *Wnd) platformZoom() {
	// RAW: Implement for Linux
}

func (window *Wnd) platformSetToolTip(tip string) {
	// RAW: Implement for Linux
}

func (window *Wnd) platformSetCursor(c *cursor.Cursor) {
	window.toXWindow().SetCursor(x11.Cursor(uintptr(c.PlatformPtr())))
}

func (window *Wnd) platformInvoke(id uint64) {
	if window.Valid() {
		window.toXWindow().InvokeTask(id)
	}
}

func (window *Wnd) platformInvokeAfter(id uint64, after time.Duration) {
	time.AfterFunc(after, func() {
		window.platformInvoke(id)
	})
}
