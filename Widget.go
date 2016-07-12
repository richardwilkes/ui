// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

// A Widget is the basic user interface block that interacts with the user.
type Widget interface {
	// Sizer returns the Sizer for this widget, if any.
	Sizer() Sizer
	// SetSizer sets the Sizer for this widget. May be nil.
	SetSizer(sizer Sizer)

	// Layout returns the Layout for this widget, if any.
	Layout() Layout
	// SetLayout sets the Layout for this widget. May be nil.
	SetLayout(layout Layout)
	// NeedLayout returns true if this widget needs to have its children laid out.
	NeedLayout() bool
	// SetNeedLayout sets the whether this widget needs to have its children lait out.
	SetNeedLayout(needLayout bool)
	// LayoutData returns the layout data, if any, associated with this widget.
	LayoutData() interface{}
	// SetLayoutData sets layout data on this widget. May be nil.
	SetLayoutData(data interface{})
	// ValidateLayout triggers any layout that needs to be run by this widget or its children.
	ValidateLayout()

	// Border returns the Border for this widget, if any.
	Border() Border
	// SetBorder sets the Border for this widget. May be nil.
	SetBorder(border Border)

	// PaintHandler returns the PaintHandler for this widget, if any.
	PaintHandler() PaintHandler
	// SetPaintHandler sets the PaintHandler for this widget. May be nil.
	SetPaintHandler(handler PaintHandler)
	// Repaint marks this widget for painting at the next update.
	Repaint()
	// RepaintBounds marks the area 'bounds' in local coordinates within the widget for painting at
	// the next update.
	RepaintBounds(bounds Rect)
	// Paint is called by its owning window when a widget needs to be drawn. 'g' is the graphics
	// context to use. It has already had its clip set to the 'dirty' rectangle. 'dirty' is the
	// area that needs to be drawn.
	Paint(g Graphics, dirty Rect)

	// MouseDownHandler returns the MouseDownHandler for this widget, if any.
	MouseDownHandler() MouseDownHandler
	// MouseDownHandler sets the MouseDownHandler for this widget. May be nil.
	SetMouseDownHandler(handler MouseDownHandler)

	// MouseDraggedHandler returns the MouseDraggedHandler for this widget, if any.
	MouseDraggedHandler() MouseDraggedHandler
	// SetMouseDraggedHandler sets the MouseDraggedHandler for this widget. May be nil.
	SetMouseDraggedHandler(handler MouseDraggedHandler)

	// MouseUpHandler returns the MouseUpHandler for this widget, if any.
	MouseUpHandler() MouseUpHandler
	// SetMouseUpHandler sets the MouseUpHandler for this widget. May be nil.
	SetMouseUpHandler(handler MouseUpHandler)

	// MouseEnteredHandler returns the MouseEnteredHandler for this widget, if any.
	MouseEnteredHandler() MouseEnteredHandler
	// SetMouseEnteredHandler sets the MouseEnteredHandler for this widget. May be nil.
	SetMouseEnteredHandler(handler MouseEnteredHandler)

	// MouseMovedHandler returns the MouseMovedHandler for this widget, if any.
	MouseMovedHandler() MouseMovedHandler
	// SetMouseMovedHandler sets the MouseMovedHandler for this widget. May be nil.
	SetMouseMovedHandler(handler MouseMovedHandler)

	// MouseExitedHandler returns the MouseExitedHandler for this widget, if any.
	MouseExitedHandler() MouseExitedHandler
	// SetMouseExitedHandler sets the MouseExitedHandler for this widget. May be nil.
	SetMouseExitedHandler(handler MouseExitedHandler)

	// ToolTipHandler returns the ToolTipHandler for this widget, if any.
	ToolTipHandler() ToolTipHandler
	// SetToolTipHandler sets the ToolTipHandler for this widget. May be nil.
	SetToolTipHandler(handler ToolTipHandler)

	// Enabled returns true if this widget is currently enabled and can receive events.
	Enabled() bool
	// SetEnabled sets this widget's enabled state.
	SetEnabled(enabled bool)
	// Focused returns true if this widget's has the keyboard focus.
	Focused() bool
	// SetFocused sets this widget's focus state.
	SetFocused(focused bool)

	// Children returns the direct descendents of this widget.
	Children() []Widget
	// IndexOfChild returns the index of the specified child, or -1 if the passed in widget is not
	// a child of this widget.
	IndexOfChild(child Widget) int
	// AddChild adds 'child' as a child of this widget, removing it from any previous parent it may
	// have had.
	AddChild(child Widget)
	// AddChildAtIndex adds 'child' as a child of this widget at the 'index', removing it from any
	// previous parent it may have had.
	AddChildAtIndex(child Widget, index int)
	// RemoveChild removes 'child' from this widget. If 'child' is not a direct descendent of this
	// widget, nothing happens.
	RemoveChild(child Widget)
	// RemoveChildAtIndex removes the child widget at 'index' from this widget. If 'index' is out
	// of range, nothing happens.
	RemoveChildAtIndex(index int)
	// RemoveFromParent removes this widget from its parent, if any.
	RemoveFromParent()
	// Parent returns the parent widget, if any.
	Parent() Widget
	// SetParent sets `parent` to be the parent of this widget. It does not add this widget to the
	// parent as a child. Call AddChild or AddChildAtIndex for that.
	SetParent(parent Widget)
	// Window returns the containing window, if any.
	Window() *Window
	// RootOfWindow returns true if this widget is the root widget of a window.
	RootOfWindow() bool

	// Bounds returns the location and size of the widget in its parent's coordinate system.
	Bounds() Rect
	// LocalBounds returns the location and size of the widget in local coordinates.
	LocalBounds() Rect
	// LocalInsetBounds returns the location and size of the widget in local coordinates after
	// adjusting for any Border it may have.
	LocalInsetBounds() Rect
	// SetBounds sets the location and size of the widget in its parent's coordinate system.
	SetBounds(bounds Rect)
	// Location returns the location of this widget in its parent's coordinate system.
	Location() Point
	// SetLocation sets the location of this widget in its parent's coordinate system.
	SetLocation(pt Point)
	// Size returns the size of this widget.
	Size() Size
	// SetSize sets the size of this widget.
	SetSize(size Size)

	// WidgetAt returns the leaf-most child widget containing 'pt', or this widget if no child
	// is found.
	WidgetAt(pt Point) Widget

	// ToWindow converts widget-local coordinates into window coordinates.
	ToWindow(pt Point) Point
	// FromWindow converts window coordinates into widget-local coordinates.
	FromWindow(pt Point) Point

	// Background returns the background color of this widget.
	Background() Color
	// SetBackground sets the background color of this widget.
	SetBackground(color Color)
}
