// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package event

import (
	"bytes"
	"fmt"
	"github.com/richardwilkes/ui/geom"
	"reflect"
)

// ToolTip is generated when a tooltip is being requested for the widget.
type ToolTip struct {
	target   Target
	where    geom.Point
	tooltip  string
	finished bool
}

// NewToolTip creates a new ToolTip event. 'target' is the widget the mouse is over. 'where' is the
// location in the window where the mouse is.
func NewToolTip(target Target, where geom.Point) *ToolTip {
	return &ToolTip{target: target, where: where}
}

// Type returns the event type ID.
func (e *ToolTip) Type() Type {
	return ToolTipType
}

// Target the original target of the event.
func (e *ToolTip) Target() Target {
	return e.target
}

// Cascade returns true if this event should be passed to its target's parent if not marked done.
func (e *ToolTip) Cascade() bool {
	return false
}

// Where returns the location in the window the mouse is.
func (e *ToolTip) Where() geom.Point {
	return e.where
}

// ToolTip returns the text that was set for the tooltip.
func (e *ToolTip) ToolTip() string {
	return e.tooltip
}

// SetToolTip sets the text to use for the tooltip.
func (e *ToolTip) SetToolTip(text string) {
	e.tooltip = text
}

// Finished returns true if this event has been handled and should no longer be processed.
func (e *ToolTip) Finished() bool {
	return e.finished
}

// Finish marks this event as handled and no longer eligible for processing.
func (e *ToolTip) Finish() {
	e.finished = true
}

// String implements the fmt.Stringer interface.
func (e *ToolTip) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("ToolTip[Where: [%v], Target: %v", e.where, reflect.ValueOf(e.target).Pointer()))
	if e.tooltip != "" {
		buffer.WriteString(", Text: '")
		buffer.WriteString(e.tooltip)
		buffer.WriteString("'")
	}
	if e.finished {
		buffer.WriteString(", Finished")
	}
	buffer.WriteString("]")
	return buffer.String()
}
