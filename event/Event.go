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
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/geom"
	"reflect"
	"time"
)

const (
	// PaintEvent is generated when a widget needs to be drawn.
	PaintEvent = iota
	// MouseDownEvent is generated when a mouse button is pressed on a widget.
	MouseDownEvent
	// MouseDraggedEvent is generated when the mouse is moved within a widget while a mouse button
	// is down.
	MouseDraggedEvent
	// MouseUpEvent is generated when the mouse button is released after a mouse button press
	// occurred within a widget.
	MouseUpEvent
	// MouseEnteredEvent is generated when the mouse enters a widget.
	MouseEnteredEvent
	// MouseMovedEvent is generated when the mouse moves within a widget, except when a mouse
	// button is also down (a MouseDraggedEvent is generated for that).
	MouseMovedEvent
	// MouseExitedEvent is generated when the mouse exits a widget.
	MouseExitedEvent
	// MouseWheelEvent is generated when the mouse wheel is used over a widget.
	MouseWheelEvent
	// FocusGainedEvent is generated when a widget gains the keyboard focus.
	FocusGainedEvent
	// FocusLostEvent is generated when a widget loses the keyboard focus.
	FocusLostEvent
	// KeyDownEvent is generated when a key is pressed.
	KeyDownEvent
	// KeyTypedEvent is generated when a rune has been generated on the keyboard.
	KeyTypedEvent
	// KeyUpEvent is generated when a key is released.
	KeyUpEvent
	// ToolTipEvent is generated when a tooltip is being requested for the widget.
	ToolTipEvent
	// ResizeEvent is generated when a widget is resized.
	ResizeEvent
	// ClosingEvent is generated when a window is asked to close. Set Discard to true to cancel the
	// closing.
	ClosingEvent
	// ClosedEvent is generated when a window is closed.
	ClosedEvent
	// UserEvent should be used as the base value for application custom events.
	UserEvent = 10000
)

// Event holds the data associated with an event.
type Event struct {
	Type         int           // Valid for all events.
	When         time.Time     // Valid for all events.
	Target       Target        // Valid for all events
	GC           draw.Graphics // Valid only for PaintEvent.
	DirtyRect    geom.Rect     // Valid only for PaintEvent.
	Where        geom.Point    // In window coordinates. Valid for MouseDownEvent, MouseDraggedEvent, MouseUpEvent, MouseEnteredEvent, MouseMovedEvent, and MouseWheelEvent.
	Delta        geom.Point    // Valid only for MouseWheelEvent. The amount scrolled in each direction.
	KeyModifiers KeyMask       // Valid for MouseDownEvent, MouseDraggedEvent, MouseUpEvent, MouseEnteredEvent, MouseMovedEvent, MouseExitedEvent, MouseWheelEvent, KeyDownEvent, KeyTypedEvent, and KeyUpEvent.
	Button       int           // Valid only for MouseDownEvent. The button that is down.
	Clicks       int           // Valid only for MouseDownEvent. The number of consecutive clicks in the widget.
	ToolTip      string        // Valid only for ToolTipEvent. Set this to the text to display for the tooltip.
	KeyCode      int           // Valid for KeyDownEvent and KeyUpEvent.
	KeyTyped     rune          // Valid only for KeyTypedEvent.
	Repeat       bool          // Valid for KeyDownEvent and KeyTypedEvent. Set to true if the key was auto-generated.
	CascadeUp    bool          // Valid for all events. true if this event should be cascaded up to parents.
	Discard      bool          // Valid for MouseDownEvent, KeyDownEvent, and ClosingEvent. Set to true to if the event should be ignored (i.e. don't do processing that would have side-effects).
	Done         bool          // Valid for all events. Set to true to stop processing this event.
}

// Dispatch the event. If there is more than one handler for the event type registered with the
// target, they will each be given a chance to handle the event in order. Should one of them set
// the Done flag on the event, processing will halt immediately. Once the target has been given an
// opportunity to process the event, if the event's CascadeUp flag is set, its parent will then be
// given the chance. This will continue until there are no more parents, the event's Done flag is
// set, or the event's CascadeUp flag is unset.
func (event *Event) Dispatch() {
	target := event.Target
	for target != nil {
		if handlers, ok := target.EventHandlers().Lookup(event.Type); ok {
			for _, handler := range handlers {
				handler(event)
				if event.Done {
					return
				}
			}
		}
		if !event.CascadeUp {
			return
		}
		target = target.ParentTarget()
	}
}

// ShiftDown returns true if the shift key is being pressed.
func (event *Event) ShiftDown() bool {
	return event.KeyModifiers&ShiftKeyMask == ShiftKeyMask
}

// OptionDown returns true if the option/alt key is being pressed.
func (event *Event) OptionDown() bool {
	return event.KeyModifiers&OptionKeyMask == OptionKeyMask
}

// ControlDown returns true if the control key is being pressed.
func (event *Event) ControlDown() bool {
	return event.KeyModifiers&ControlKeyMask == ControlKeyMask
}

// CommandDown returns true if the command/meta key is being pressed.
func (event *Event) CommandDown() bool {
	return event.KeyModifiers&CommandKeyMask == CommandKeyMask
}

// CapsLockDown returns true if the caps lock key is being pressed.
func (event *Event) CapsLockDown() bool {
	return event.KeyModifiers&CapsLockKeyMask == CapsLockKeyMask
}

// String implements fmt.Stringer.
func (event *Event) String() string {
	var buffer bytes.Buffer
	switch event.Type {
	case PaintEvent:
		buffer.WriteString("PaintEvent")
	case MouseDownEvent:
		buffer.WriteString("MouseDownEvent")
	case MouseDraggedEvent:
		buffer.WriteString("MouseDraggedEvent")
	case MouseUpEvent:
		buffer.WriteString("MouseUpEvent")
	case MouseEnteredEvent:
		buffer.WriteString("MouseEnteredEvent")
	case MouseMovedEvent:
		buffer.WriteString("MouseMovedEvent")
	case MouseExitedEvent:
		buffer.WriteString("MouseExitedEvent")
	case MouseWheelEvent:
		buffer.WriteString("MouseWheelEvent")
	case FocusGainedEvent:
		buffer.WriteString("FocusGainedEvent")
	case FocusLostEvent:
		buffer.WriteString("FocusLostEvent")
	case KeyDownEvent:
		buffer.WriteString("KeyDownEvent")
	case KeyTypedEvent:
		buffer.WriteString("KeyTypedEvent")
	case KeyUpEvent:
		buffer.WriteString("KeyUpEvent")
	case ToolTipEvent:
		buffer.WriteString("ToolTipEvent")
	case ResizeEvent:
		buffer.WriteString("ResizeEvent")
	default:
		buffer.WriteString(fmt.Sprintf("Custom%dEvent", event.Type))
	}
	buffer.WriteString(fmt.Sprintf("[When: %v, Target: %v", event.When, reflect.ValueOf(event.Target).Pointer()))
	switch event.Type {
	case PaintEvent:
		buffer.WriteString(fmt.Sprintf(", DirtyRect: %v", event.DirtyRect))
	case MouseDownEvent:
		buffer.WriteString(fmt.Sprintf(", Where: %v, KeyModifiers: %s, Button: %d, Clicks: %d", event.Where, event.keyModifiersAsString(), event.Button, event.Clicks))
	case MouseDraggedEvent:
		buffer.WriteString(fmt.Sprintf(", Where: %v, KeyModifiers: %s", event.Where, event.keyModifiersAsString()))
	case MouseUpEvent:
		buffer.WriteString(fmt.Sprintf(", Where: %v, KeyModifiers: %s", event.Where, event.keyModifiersAsString()))
	case MouseEnteredEvent:
		buffer.WriteString(fmt.Sprintf(", Where: %v, KeyModifiers: %s", event.Where, event.keyModifiersAsString()))
	case MouseMovedEvent:
		buffer.WriteString(fmt.Sprintf(", Where: %v, KeyModifiers: %s", event.Where, event.keyModifiersAsString()))
	case MouseExitedEvent:
		buffer.WriteString(fmt.Sprintf(", KeyModifiers: %s", event.keyModifiersAsString()))
	case MouseWheelEvent:
		buffer.WriteString(fmt.Sprintf(", Where: %v, Delta: %v, KeyModifiers: %s", event.Where, event.Delta, event.keyModifiersAsString()))
	case KeyDownEvent:
		buffer.WriteString(fmt.Sprintf(", KeyModifiers: %s, KeyCode: %d", event.keyModifiersAsString(), event.KeyCode))
		if event.Repeat {
			buffer.WriteString(", Repeat")
		}
	case KeyTypedEvent:
		buffer.WriteString(fmt.Sprintf(", KeyModifiers: %s, KeyCode: %d, KeyTyped: %v", event.keyModifiersAsString(), event.KeyCode, event.KeyTyped))
		if event.Repeat {
			buffer.WriteString(", Repeat")
		}
	case KeyUpEvent:
		buffer.WriteString(fmt.Sprintf(", KeyModifiers: %s, KeyCode: %d", event.keyModifiersAsString(), event.KeyCode))
		if event.Repeat {
			buffer.WriteString(", Repeat")
		}
	}
	if event.CascadeUp {
		buffer.WriteString(", CascadeUp")
	}
	if event.Discard && (event.Type == MouseDownEvent || event.Type == KeyDownEvent) {
		buffer.WriteString(", Discard")
	}
	if event.Done {
		buffer.WriteString(", Done")
	}
	buffer.WriteString("]")
	return buffer.String()
}

func (event *Event) keyModifiersAsString() string {
	var buffer bytes.Buffer
	event.appendKeyModifier(&buffer, CapsLockKeyMask, "CapsLock")
	event.appendKeyModifier(&buffer, ShiftKeyMask, "Shift")
	event.appendKeyModifier(&buffer, ControlKeyMask, "Control")
	event.appendKeyModifier(&buffer, OptionKeyMask, "Option")
	event.appendKeyModifier(&buffer, CommandKeyMask, "Command")
	if buffer.Len() == 0 {
		buffer.WriteString("None")
	}
	return buffer.String()
}

func (event *Event) appendKeyModifier(buffer *bytes.Buffer, mask KeyMask, name string) {
	if event.KeyModifiers&mask == mask {
		if buffer.Len() > 0 {
			buffer.WriteString(" | ")
		}
		buffer.WriteString(name)
	}
}
