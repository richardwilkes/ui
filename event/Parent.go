package event

var (
	globalTarget Target
)

// GlobalTarget returns the global, top-level target.
func GlobalTarget() Target {
	return globalTarget
}

// SetGlobalTarget sets the global, top-level target.
func SetGlobalTarget(target Target) {
	globalTarget = target
}
