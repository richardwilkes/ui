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
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/geom"
	"unsafe"
)

// Window represents a window on the display.
type Window interface {
	event.Target

	// Title returns the title of this window.
	Title() string
	// SetTitle sets the title of this window.
	SetTitle(title string)

	// Frame returns the boundaries in display coordinates of the frame of this window (i.e. the
	// area that includes both the content and its border and window controls).
	Frame() geom.Rect
	// Location returns the upper left corner of the window in display coordinates.
	Location() geom.Point
	// SetLocation moves the upper left corner of the window to the specified point in display
	// coordinates.
	SetLocation(pt geom.Point)
	// Size returns the size of the window, including its frame and window controls.
	Size() geom.Size
	// SetSize sets the size of the window.
	SetSize(size geom.Size)

	// ContentFrame returns the boundaries of the root content widget of this window.
	ContentFrame() geom.Rect
	// ContentLocalBounds returns the local boundaries of the content widget of this window.
	ContentLocalBounds() geom.Rect
	// ContentLocation returns the upper left corner of the content widget in display coordinates.
	ContentLocation() geom.Point
	// SetContentLocation moves the window such that the upper left corner of the content widget is
	// at the specified point in display coordinates.
	SetContentLocation(pt geom.Point)
	// ContentSize returns the size of the content widget.
	ContentSize() geom.Size
	// SetContentSize sets the size of the window to fit the specified content size.
	SetContentSize(size geom.Size)
	// Pack sets the window's content size to match the preferred size of the root widget.
	Pack()

	// RootWidget returns the root widget of the window.
	RootWidget() Widget

	// Focus returns the widget with the keyboard focus in this window.
	Focus() Widget
	// SetFocus sets the keyboard focus to the specified target.
	SetFocus(target Widget)
	// FocusNext moves the keyboard focus to the next focusable widget.
	FocusNext()
	// FocusPrevious moves the keyboard focus to the previous focusable widget.
	FocusPrevious()
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
	ScalingFactor() float32

	// Minimize performs the platform's minimize function on the window.
	Minimize()
	// Zoom performs the platform's zoom funcion on the window.
	Zoom()

	// PlatformPtr returns a pointer to the underlying platform-specific data.
	PlatformPtr() unsafe.Pointer
}
