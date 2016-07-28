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
	"reflect"
)

// EventHandler is called to handle a single event.
type EventHandler func(event *Event)

type Handlers struct {
	handlers map[int][]EventHandler
}

// Lookup returns an event handler list for the 'eventType'.
func (eh *Handlers) Lookup(eventType int) ([]EventHandler, bool) {
	handlers, ok := eh.handlers[eventType]
	return handlers, ok
}

// Add an event handler for an event type.
func (eh *Handlers) Add(eventType int, handler EventHandler) {
	if eh.handlers == nil {
		eh.handlers = make(map[int][]EventHandler)
	}
	eh.handlers[eventType] = append(eh.handlers[eventType], handler)
}

// Remove an event handler for an event type.
func (eh *Handlers) Remove(eventType int, handler EventHandler) {
	if eh.handlers != nil {
		hPtr := reflect.ValueOf(handler).Pointer()
		handlers := eh.handlers[eventType]
		for i, one := range handlers {
			if reflect.ValueOf(one).Pointer() == hPtr {
				if len(handlers) == 1 {
					delete(eh.handlers, eventType)
				} else {
					copy(handlers[i:], handlers[i+1:])
					length := len(handlers) - 1
					handlers[length] = nil
					eh.handlers[eventType] = handlers[:length]
				}
				break
			}
		}
	}
}
