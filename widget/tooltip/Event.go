// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package tooltip

import (
	"bytes"
	"fmt"
	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/id"
)

var (
	EventType = event.Type(id.Next())
)

// ToolTip is generated when a tooltip is being requested for the widget.
type Event struct {
	target   ui.Widget
	where    geom.Point
	avoid    geom.Rect
	tooltip  ui.Widget
	finished bool
}

// NewToolTip creates a new ToolTip event. 'target' is the widget the mouse is over. 'where' is the
// location in the window where the mouse is. 'avoid' is the area to avoid placing the tooltip
// within, using window coordinates. Typically, 'avoid' is set to the boundaries of the widget
// within the window.
func NewEvent(target ui.Widget, where geom.Point, avoid geom.Rect) *Event {
	return &Event{target: target, where: where, avoid: avoid}
}

// Type returns the event type ID.
func (e *Event) Type() event.Type {
	return event.ToolTipType
}

// Target the original target of the event.
func (e *Event) Target() event.Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *Event) Cascade() bool {
	return false
}

// Where returns the location in the window the mouse is.
func (e *Event) Where() geom.Point {
	return e.where
}

// Avoid returns the area to avoid placing the tooltip within.
func (e *Event) Avoid() geom.Rect {
	return e.avoid
}

// SetAvoid sets the area to avoid placing the tooltip within.
func (e *Event) SetAvoid(avoid geom.Rect) {
	e.avoid = avoid
}

// ToolTip returns a widget to be used as the tooltip.
func (e *Event) ToolTip() ui.Widget {
	return e.tooltip
}

// SetToolTip sets the widget to be used for the tooltip.
func (e *Event) SetToolTip(tooltip ui.Widget) {
	e.tooltip = tooltip
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *Event) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *Event) Finish() {
	e.finished = true
}

// String implements the fmt.Stringer interface.
func (e *Event) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Event[Where: [%v], Avoid: [%v], Target: %v", e.where, e.avoid, e.target))
	if e.tooltip != nil {
		buffer.WriteString(fmt.Sprintf(", Tooltip: [%v]", e.tooltip))
	}
	if e.finished {
		buffer.WriteString(", Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
