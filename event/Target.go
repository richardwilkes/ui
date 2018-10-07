package event

import (
	"github.com/richardwilkes/ui/object"
)

// Target marks objects that can be the target of an event.
type Target interface {
	object.Object
	// EventHandlers returns the handler mappings for this Target.
	EventHandlers() *Handlers
	// ParentTarget returns the parent of this Target, or nil.
	ParentTarget() Target
}
