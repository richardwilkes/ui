package event

import (
	"github.com/richardwilkes/toolbox/log/logadapter"
)

// Type holds a unique type ID for each event type.
type Type int

// The event types
const (
	AppWillFinishStartupType Type = iota
	AppDidFinishStartupType
	AppWillActivateType
	AppDidActivateType
	AppWillDeactivateType
	AppDidDeactivateType
	AppQuitRequestedType
	AppWillQuitType
	AppLastWindowClosedType
	AppPopulateMenuBarType
	PaintType
	MouseDownType
	MouseDraggedType
	MouseUpType
	MouseEnteredType
	MouseMovedType
	MouseExitedType
	MouseWheelType
	ClickType
	SelectionType
	FocusGainedType
	FocusLostType
	KeyDownType
	KeyUpType
	ToolTipType
	UpdateCursorType
	ResizedType
	ClosingType
	ClosedType
	ValidateType
	ModifiedType
	// UserType should be used as the base value for custom application
	// events.
	UserType = 10000
)

// Event is the minimal interface that must be implemented for events.
type Event interface {
	// Type returns the event type ID.
	Type() Type
	// Target the original target of the event.
	Target() Target
	// Cascade returns true if this event should be passed to its target's
	// parent if not marked done.
	Cascade() bool
	// Finished returns true if this event has been handled and should no
	// longer be processed.
	Finished() bool
	// Finish marks this event as handled and no longer eligible for
	// processing.
	Finish()
}

var (
	// TraceAllEvents will cause all events to be logged if true. Overrides
	// TraceEventTypes.
	TraceAllEvents bool
	// TraceEventTypes will cause the types present in the slice to be logged.
	TraceEventTypes []Type
	// TraceLogger is used to log events when not nil and either of
	// TraceAllEvents is true or len(TraceEventTypes) > 0. Defaults to nil.
	TraceLogger logadapter.InfoLogger
)

// Dispatch an event. If there is more than one handler for the event type
// registered with the target, they will each be given a chance to handle the
// event in order. Should one of them set the Finished flag on the event,
// processing will halt immediately. Once the target has been given an
// opportunity to process the event, if the event's Cascade flag is set, its
// parent will then be given the chance. This will continue until there are no
// more parents or the event's Finished flag is set.
func Dispatch(e Event) {
	eventType := e.Type()
	if TraceLogger != nil {
		if TraceAllEvents {
			TraceLogger.Info(e)
		} else if len(TraceEventTypes) > 0 {
			for _, t := range TraceEventTypes {
				if t == eventType {
					TraceLogger.Info(e)
					break
				}
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
