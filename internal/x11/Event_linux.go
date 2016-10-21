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
	// #cgo pkg-config: x11
	// #include <X11/Xlib.h>
	"C"
)

const (
	KeyPressType = 2 + iota
	KeyReleaseType
	ButtonPressType
	ButtonReleaseType
	MotionNotifyType
	EnterNotifyType
	LeaveNotifyType
	FocusInType
	FocusOutType
	KeymapNotifyType
	ExposeType
	GraphicsExposeType
	NoExposeType
	VisibilityNotifyType
	CreateNotifyType
	DestroyNotifyType
	UnmapNotifyType
	MapNotifyType
	MapRequestType
	ReparentNotifyType
	ConfigureNotifyType
	ConfigureRequestType
	GravityNotifyType
	ResizeRequestType
	CirculateNotifyType
	CirculateRequestType
	PropertyNotifyType
	SelectionClearType
	SelectionRequestType
	SelectionNotifyType
	ColormapNotifyType
	ClientMessageType
	MappingNotifyType
	GenericEventType
	LASTEventType
)

type Event C.XEvent

type Eventable interface {
	ToEvent() *Event
}

func (evt *Event) Type() int {
	return int((*C.XAnyEvent)(unsafe.Pointer(evt))._type)
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

func (evt *Event) ToPropertyEvent() *PropertyEvent {
	return (*PropertyEvent)(unsafe.Pointer(evt))
}

func (evt *Event) ToSelectionRequestEvent() *SelectionRequestEvent {
	return (*SelectionRequestEvent)(unsafe.Pointer(evt))
}

func (evt *Event) ToSelectionEvent() *SelectionEvent {
	return (*SelectionEvent)(unsafe.Pointer(evt))
}

func (evt *Event) ToSelectionClearEvent() *SelectionClearEvent {
	return (*SelectionClearEvent)(unsafe.Pointer(evt))
}

func (evt *Event) ToEvent() *Event {
	return (*Event)(evt)
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
