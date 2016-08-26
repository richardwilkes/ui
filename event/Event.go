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
	"log"
)

// Common button ids.
const (
	LeftButton   = 0
	RightButton  = 1
	MiddleButton = 2
)

// Event is the minimal interface that must be implemented for events.
type Event interface {
	// Type returns the event type ID.
	Type() Type
	// Target the original target of the event.
	Target() Target
	// Cascade returns true if this event should be passed to its target's parent if not marked done.
	Cascade() bool
	// Finished returns true if this event has been handled and should no longer be processed.
	Finished() bool
	// Finish marks this event as handled and no longer eligible for processing.
	Finish()
}

var (
	// TraceAllEvents will cause all events to be logged if true. Overrides TraceEventTypes.
	TraceAllEvents bool
	// TraceEventTypes will cause the types present in the slice to be logged.
	TraceEventTypes []Type
)

// Dispatch an event. If there is more than one handler for the event type registered with the
// target, they will each be given a chance to handle the event in order. Should one of them set
// the Finished flag on the event, processing will halt immediately. Once the target has been given
// an opportunity to process the event, if the event's Cascade flag is set, its parent will then be
// given the chance. This will continue until there are no more parents or the event's Finished
// flag is set.
func Dispatch(e Event) {
	eventType := e.Type()
	if TraceAllEvents {
		log.Println(e)
	} else if len(TraceEventTypes) > 0 {
		for _, t := range TraceEventTypes {
			if t == eventType {
				log.Println(e)
				break
			}
		}
	}
	target := e.Target()
	for target != nil {
		if handlers, ok := target.EventHandlers().Lookup(eventType); ok {
			for _, handler := range handlers {
				handler(e)
				if e.Finished() {
					break
				}
			}
		}
		if !e.Cascade() {
			break
		}
		target = target.ParentTarget()
	}
}
