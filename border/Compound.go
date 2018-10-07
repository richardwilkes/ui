package border

import (
	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui/draw"
)

// Compound is a border that contains other borders.
type Compound struct {
	borders []Border
}

// NewCompound creates a border that contains other borders. The first one will be drawn in
// the outermost position, with each successive one moving further into the interior.
func NewCompound(borders ...Border) *Compound {
	return &Compound{borders: borders}
}

// Insets implements the Border interface.
func (c *Compound) Insets() geom.Insets {
	insets := geom.Insets{}
	for _, one := range c.borders {
		insets.Add(one.Insets())
	}
	return insets
}

// Draw implements the Border interface.
func (c *Compound) Draw(gc *draw.Graphics, bounds geom.Rect) {
	for _, one := range c.borders {
		gc.Save()
		one.Draw(gc, bounds)
		gc.Restore()
		bounds.Inset(one.Insets())
	}
}
