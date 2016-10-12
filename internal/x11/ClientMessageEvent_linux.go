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
	"unsafe"
	// #cgo linux LDFLAGS: -lX11
	// #include <X11/Xlib.h>
	"C"
)

const (
	ClientMessageType = EventType(C.ClientMessage)
)

type ClientMessageSubType C.Atom

var (
	ProtocolsSubType ClientMessageSubType
	TaskSubType      ClientMessageSubType
)

type Protocol C.Atom

var (
	DeleteWindowProtocol Protocol
)

type ClientMessageEvent C.XClientMessageEvent

func (evt *ClientMessageEvent) Window() Window {
	return Window(evt.window)
}

func (evt *ClientMessageEvent) Format() int {
	return int(evt.format)
}

func (evt *ClientMessageEvent) SubType() ClientMessageSubType {
	return ClientMessageSubType(evt.message_type)
}

func (evt *ClientMessageEvent) Protocol() Protocol {
	return *(*Protocol)(unsafe.Pointer(&evt.data))
}

func (evt *ClientMessageEvent) TaskID() uint64 {
	return *(*uint64)(unsafe.Pointer(&evt.data))
}
