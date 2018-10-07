package event

import (
	"reflect"
)

// Handler is called to handle a single event.
type Handler func(event Event)

// Handlers maintains mapping of event types to event handlers.
type Handlers struct {
	handlers map[Type][]Handler
}

// Lookup returns an event handler list for the 'eventType'.
func (eh *Handlers) Lookup(eventType Type) ([]Handler, bool) {
	handlers, ok := eh.handlers[eventType]
	return handlers, ok
}

// Add an event handler for an event type.
func (eh *Handlers) Add(eventType Type, handler Handler) {
	if eh.handlers == nil {
		eh.handlers = make(map[Type][]Handler)
	}
	eh.handlers[eventType] = append(eh.handlers[eventType], handler)
}

// Remove an event handler for an event type.
func (eh *Handlers) Remove(eventType Type, handler Handler) {
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
