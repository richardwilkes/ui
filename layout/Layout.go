package layout

import (
	"math"

	"github.com/richardwilkes/toolbox/xmath/geom"
)

const (
	// NoHint is passed as a hint value when one or both dimensions have no suggested value.
	NoHint = -1
	// DefaultMax is the default value that should be used for a maximum dimension if the
	// block has no real preference and can be expanded beyond its preferred size. This is
	// intentionally not something like math.MaxFloat32 to allow basic math operations an
	// opportunity to succeed when laying out components. It is perfectly acceptable to use
	// a larger value than this, however, if that makes sense for your specific component.
	DefaultMax = 10000
)

var (
	// NoHintSize is a convenience for passing to layouts when you don't have any particular
	// size constraints in mind. Should be treated as read-only.
	NoHintSize = geom.Size{Width: NoHint, Height: NoHint}
)

// Sizer is called when no layout has been set for a widget. Returns the minimum, preferred, and
// maximum sizes of the widget. The hint's values will be either NoHint or a specific value
// if that particular dimension has already been determined.
type Sizer interface {
	Sizes(hint geom.Size) (min, pref, max geom.Size)
}

// The Layout interface should be implemented by objects that provide layout services.
type Layout interface {
	Sizer
	// Layout is called to layout the target and its children.
	Layout()
}

// DefaultMaxSize returns the size that is at least as large as DefaultMax in both dimensions, but
// larger if the preferred size that is passed in is larger.
func DefaultMaxSize(pref geom.Size) geom.Size {
	return geom.Size{Width: math.Max(DefaultMax, pref.Width), Height: math.Max(DefaultMax, pref.Height)}
}
