// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

// PaintHandler is called when a widget needs to be drawn. 'g' is the Graphics context to use. It
// has already had its clip set to the dirty rectangle. 'dirty' is the area that needs to be drawn.
type PaintHandler interface {
	OnPaint(g Graphics, dirty Rect)
}

// MouseDownHandler is called when a mouse button is pressed on a widget. 'where' is the location
// of the mouse in local coordinates. 'keyModifiers' are the modifier keys that were down at the
// time of the event. 'which' is the button that is pressed. 'clickCount' is the number of
// consecutive clicks in this widget. Return true if the event was effectively discarded for the
// object, such as when a mouse press is passed off to a popup menu.
type MouseDownHandler interface {
	OnMouseDown(where Point, keyModifiers int, which int, clickCount int) bool
}

// MouseDraggedHandler is called when the mouse is moved within a widget while a mouse button is
// down. 'where' is the location of the mouse in local coordinates. 'keyModifiers' are the modifier
// keys that were down at the time of the event.
type MouseDraggedHandler interface {
	OnMouseDragged(where Point, keyModifiers int)
}

// MouseUpHandler is called when the mouse button is released after a mouse button press occurred
// within a widget. 'where' is the location of the mouse in local coordinates. 'keyModifiers' are
// the modifier keys that were down at the time of the event.
type MouseUpHandler interface {
	OnMouseUp(where Point, keyModifiers int)
}

// MouseEnteredHandler is called when the mouse enters a widget. 'where' is the location of the
// mouse in local coordinates. 'keyModifiers' are the modifier keys that were down at the time of
// the event.
type MouseEnteredHandler interface {
	OnMouseEntered(where Point, keyModifiers int)
}

// MouseMovedHandler is called when the mouse moves within a widget, except when a mouse button is
// also down (use a MouseDraggedHandler for that). 'where' is the location of the mouse in local
// coordinates. 'keyModifiers' are the modifier keys that were down at the time of the event.
type MouseMovedHandler interface {
	OnMouseMoved(where Point, keyModifiers int)
}

// MouseExitedHandler is called when the mouse exits a widget. 'keyModifiers' are the modifier keys
// that were down at the time of the event.
type MouseExitedHandler interface {
	OnMouseExited(keyModifiers int)
}

// ToolTipHandler is called when a tooltip is being requested for the widget. 'where' is the
// location of the mouse in local coordinates. Return the text of the tooltip, or an empty string
// if no tooltip should be shown.
type ToolTipHandler interface {
	OnToolTip(where Point) string
}
