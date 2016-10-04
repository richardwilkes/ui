// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

import (
	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/menu"
	"time"
	"unsafe"
)

type Window interface {
	event.Target
	// AttemptClose closes the window if a Closing event permits it.
	AttemptClose()
	// Close the window.
	Close()
	// Valid returns true if the window is still valid (i.e. has not been closed).
	Valid() bool
	// Title returns the title of this window.
	Title() string
	// SetTitle sets the title of this window.
	SetTitle(title string)
	// Frame returns the boundaries in display coordinates of the frame of this window (i.e. the
	// area that includes both the content and its border and window controls).
	Frame() geom.Rect
	// SetFrame sets the boundaries of the frame of this window.
	SetFrame(bounds geom.Rect)
	// ContentFrame returns the boundaries of the root widget of this window.
	ContentFrame() geom.Rect
	// SetContentFrame sets the boundaries of the root widget of this window.
	SetContentFrame(bounds geom.Rect)
	// ContentLocalFrame returns the local boundaries of the root widget of this window.
	ContentLocalFrame() geom.Rect
	// Pack sets the window's content size to match the preferred size of the root widget.
	Pack()
	// MenuBar returns the menu bar for the window. On some platforms, the menu bar is a global
	// entity and the same value will be returned for all windows.
	MenuBar() menu.Bar
	// Content returns the content widget of the window. This is not the root widget of the window,
	// which contains both the content widget and the menu bar, for platforms that hold the menu bar
	// within the window.
	Content() Widget
	// Focus returns the widget with the keyboard focus in this window.
	Focus() Widget
	// SetFocus sets the keyboard focus to the specified target.
	SetFocus(target Widget)
	// FocusNext moves the keyboard focus to the next focusable widget.
	FocusNext()
	// FocusPrevious moves the keyboard focus to the previous focusable widget.
	FocusPrevious()
	// Focused returns true when this window has the keyboard focus.
	Focused() bool
	// ToFront attempts to bring the window to the foreground and give it the keyboard focus.
	ToFront()
	// Repaint marks this window for painting at the next update.
	Repaint()
	// RepaintBounds marks the specified bounds within the window for painting at the next update.
	RepaintBounds(bounds geom.Rect)
	// FlushPainting causes any areas marked for repainting to be painted.
	FlushPainting()
	// InLiveResize returns true if the window is being actively resized by the user at this point
	// in time. If it is, expensive painting operations should be deferred if possible to give a
	// smooth resizing experience.
	InLiveResize() bool
	// ScalingFactor returns the current OS scaling factor being applied to this window.
	ScalingFactor() float64
	// Minimize performs the platform's minimize function on the window.
	Minimize()
	// Zoom performs the platform's zoom funcion on the window.
	Zoom()
	// PlatformPtr returns a pointer to the underlying platform-specific data.
	PlatformPtr() unsafe.Pointer
	// Closable returns true if the window was created with the ClosableWindowMask.
	Closable() bool
	// Minimizable returns true if the window was created with the MiniaturizableWindowMask.
	Minimizable() bool
	// Resizable returns true if the window was created with the ResizableWindowMask.
	Resizable() bool
	// Invoke a task on the UI thread. The task is put into the system event queue and will be run at
	// the next opportunity.
	Invoke(task func())
	// InvokeAfter schedules a task to be run on the UI thread after waiting for the specified
	// duration.
	InvokeAfter(task func(), after time.Duration)
}
