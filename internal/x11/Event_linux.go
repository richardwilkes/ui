// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package x11

import (
	"github.com/richardwilkes/ui/keys"
	"unsafe"
	// #cgo linux LDFLAGS: -lX11
	// #include <X11/Xlib.h>
	"C"
)

type EventType int

type Event C.XEvent

func (evt *Event) Type() EventType {
	return EventType((*C.XAnyEvent)(unsafe.Pointer(evt))._type)
}

func (evt *Event) Window() Window {
	return Window((*C.XAnyEvent)(unsafe.Pointer(evt)).window)
}

func (evt *Event) ToKeyEvent() *KeyEvent {
	return (*KeyEvent)(unsafe.Pointer(evt))
}

func (evt *Event) ToButtonEvent() *ButtonEvent {
	return (*ButtonEvent)(unsafe.Pointer(evt))
}

func (evt *Event) ToMotionEvent() *MotionEvent {
	return (*MotionEvent)(unsafe.Pointer(evt))
}

func (evt *Event) ToCrossingEvent() *CrossingEvent {
	return (*CrossingEvent)(unsafe.Pointer(evt))
}

func (evt *Event) ToFocusChangeEvent() *FocusChangeEvent {
	return (*FocusChangeEvent)(unsafe.Pointer(evt))
}

func (evt *Event) ToExposeEvent() *ExposeEvent {
	return (*ExposeEvent)(unsafe.Pointer(evt))
}

func (evt *Event) ToDestroyWindowEvent() *DestroyWindowEvent {
	return (*DestroyWindowEvent)(unsafe.Pointer(evt))
}

func (evt *Event) ToConfigureEvent() *ConfigureEvent {
	return (*ConfigureEvent)(unsafe.Pointer(evt))
}

func (evt *Event) ToClientMessageEvent() *ClientMessageEvent {
	return (*ClientMessageEvent)(unsafe.Pointer(evt))
}

func (evt *Event) ToMapWindowEvent() *MapWindowEvent {
	return (*MapWindowEvent)(unsafe.Pointer(evt))
}

func Modifiers(state C.uint) keys.Modifiers {
	var modifiers keys.Modifiers
	if state&C.LockMask != 0 {
		modifiers |= keys.CapsLockModifier
	}
	if state&C.ShiftMask != 0 {
		modifiers |= keys.ShiftModifier
	}
	if state&C.ControlMask != 0 {
		modifiers |= keys.ControlModifier
	}
	if state&C.Mod1Mask != 0 {
		modifiers |= keys.OptionModifier
	}
	if state&C.Mod4Mask != 0 {
		modifiers |= keys.CommandModifier
	}
	return modifiers
}
