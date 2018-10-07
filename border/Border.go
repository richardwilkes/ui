package border

import (
	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui/draw"
)

// The Border interface should be implemented by objects that provide a border around an area.
type Border interface {
	// Insets returns the insets describing the space the border occupies on each side.
	Insets() geom.Insets
	// Draw the border into 'bounds'.
	Draw(gc *draw.Graphics, bounds geom.Rect)
}
