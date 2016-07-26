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
	"bytes"
	"fmt"
	"reflect"
	"time"
)

// Possible KeyMask values.
const (
	CapsLockKeyMask KeyMask = 1 << iota
	ShiftKeyMask
	ControlKeyMask
	OptionKeyMask
	CommandKeyMask   // On platforms that don't have a distinct command key, this will also be set if the Control key is pressed.
	NonStickyKeyMask = ShiftKeyMask | ControlKeyMask | OptionKeyMask | CommandKeyMask
	AllKeyMask       = CapsLockKeyMask | NonStickyKeyMask
)

// KeyMask contains flags indicating which modifier keys were down when an event occurred.
type KeyMask int

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
	UserEvent = 100
)

// EventHandler is called to handle a single event.
type EventHandler func(event *Event)

// Event holds the data associated with an event.
type Event struct {
	Type         int       // Valid for all events.
	When         time.Time // Valid for all events.
	Target       Widget    // Valid for all events
	GC           Graphics  // Valid only for PaintEvent.
	DirtyRect    Rect      // Valid only for PaintEvent.
	Where        Point     // In window coordinates. Valid for MouseDownEvent, MouseDraggedEvent, MouseUpEvent, MouseEnteredEvent, MouseMovedEvent, and MouseWheelEvent.
	Delta        Point     // Valid only for MouseWheelEvent. The amount scrolled in each direction.
	KeyModifiers KeyMask   // Valid for MouseDownEvent, MouseDraggedEvent, MouseUpEvent, MouseEnteredEvent, MouseMovedEvent, MouseExitedEvent and MouseWheelEvent.
	Button       int       // Valid only for MouseDownEvent. The button that is down.
	Clicks       int       // Valid only for MouseDownEvent. The number of consecutive clicks in the widget.
	ToolTip      string    // Valid only for ToolTipEvent. Set this to the text to display for the tooltip.
	KeyCode      int       // Valid for KeyDownEvent and KeyUpEvent.
	KeyTyped     rune      // Valid only for KeyTypedEvent.
	Repeat       bool      // Valid for KeyDownEvent and KeyTypedEvent. Set to true if the key was auto-generated.
	CascadeUp    bool      // Valid for all events. true if this event should be cascaded up to parents.
	Discard      bool      // Valid for MouseDownEvent. Set to true to if the mouse down should be ignored.
	Done         bool      // Valid for all events. Set to true to stop processing this event.
}

func (event *Event) Dispatch() {
	target := event.Target
	for target != nil {
		if handlers, ok := target.EventHandlers()[event.Type]; ok {
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
		target = target.Parent()
	}
}

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
	case ToolTipEvent:
		buffer.WriteString("ToolTipEvent")
	case ResizeEvent:
		buffer.WriteString("ResizeEvent")
	case KeyDownEvent:
		buffer.WriteString("KeyDownEvent")
	case KeyTypedEvent:
		buffer.WriteString("KeyTypedEvent")
	case KeyUpEvent:
		buffer.WriteString("KeyUpEvent")
	default:
		buffer.WriteString(fmt.Sprintf("Custom%dEvent", event.Type))
	}
	buffer.WriteString(fmt.Sprintf("[When: %v, Target: %v", event.When, reflect.ValueOf(event.Target).Pointer()))
	switch event.Type {
	case PaintEvent:
		buffer.WriteString(fmt.Sprintf(", DirtyRect: %v", event.DirtyRect))
	case MouseDownEvent:
		buffer.WriteString(fmt.Sprintf(", Where: %v, KeyModifiers: %s, Button: %d, Clicks: %d", event.Where, event.KeyModifiersAsString(), event.Button, event.Clicks))
	case MouseDraggedEvent:
		buffer.WriteString(fmt.Sprintf(", Where: %v, KeyModifiers: %s", event.Where, event.KeyModifiersAsString()))
	case MouseUpEvent:
		buffer.WriteString(fmt.Sprintf(", Where: %v, KeyModifiers: %s", event.Where, event.KeyModifiersAsString()))
	case MouseEnteredEvent:
		buffer.WriteString(fmt.Sprintf(", Where: %v, KeyModifiers: %s", event.Where, event.KeyModifiersAsString()))
	case MouseMovedEvent:
		buffer.WriteString(fmt.Sprintf(", Where: %v, KeyModifiers: %s", event.Where, event.KeyModifiersAsString()))
	case MouseExitedEvent:
		buffer.WriteString(fmt.Sprintf(", KeyModifiers: %s", event.KeyModifiersAsString()))
	case MouseWheelEvent:
		buffer.WriteString(fmt.Sprintf(", Where: %v, Delta: %v, KeyModifiers: %s", event.Where, event.Delta, event.KeyModifiersAsString()))
	case KeyDownEvent:
		buffer.WriteString(fmt.Sprintf(", KeyCode: %d", event.KeyCode))
		if event.Repeat {
			buffer.WriteString(", Repeat")
		}
	case KeyTypedEvent:
		buffer.WriteString(fmt.Sprintf(", KeyCode: %d, KeyTyped: %v", event.KeyCode, event.KeyTyped))
		if event.Repeat {
			buffer.WriteString(", Repeat")
		}
	case KeyUpEvent:
		buffer.WriteString(fmt.Sprintf(", KeyCode: %d", event.KeyCode))
		if event.Repeat {
			buffer.WriteString(", Repeat")
		}
	}
	if event.CascadeUp {
		buffer.WriteString(", CascadeUp")
	}
	if event.Type == MouseDownEvent && event.Discard {
		buffer.WriteString(", Discard")
	}
	if event.Done {
		buffer.WriteString(", Done")
	}
	buffer.WriteString("]")
	return buffer.String()
}

func (event *Event) KeyModifiersAsString() string {
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
