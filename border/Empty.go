package border

import (
	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui/draw"
)

// Empty is a border that contains empty space, effectively providing an empty margin.
type Empty struct {
	insets geom.Insets
}

// NewEmpty creates a new empty border with the specified insets.
func NewEmpty(insets geom.Insets) *Empty {
	return &Empty{insets: insets}
}

// Insets implements the Border interface.
func (e *Empty) Insets() geom.Insets {
	return e.insets
}

// Draw implements the Border interface.
func (e *Empty) Draw(gc *draw.Graphics, bounds geom.Rect) {
	// Does nothing
}
