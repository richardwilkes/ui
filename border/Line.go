package border

import (
	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/draw"
)

// Line is a border that draws a line along some or all of its sides.
type Line struct {
	insets geom.Insets
	color  color.Color
	// NoInset will cause the Insets() method to return zeroes if true.
	NoInset bool
}

// NewLine creates a new line border. The insets represent how thick the border will be drawn
// on that edge.
func NewLine(color color.Color, insets geom.Insets) *Line {
	return &Line{insets: insets, color: color}
}

// Insets implements the Border interface.
func (line *Line) Insets() geom.Insets {
	if line.NoInset {
		return geom.Insets{}
	}
	return line.insets
}

// Draw implements the Border interface.
func (line *Line) Draw(gc *draw.Graphics, bounds geom.Rect) {
	clip := bounds
	clip.Inset(line.insets)
	gc.Save()
	gc.BeginPath()
	gc.MoveTo(bounds.X, bounds.Y)
	gc.LineTo(bounds.X+bounds.Width, bounds.Y)
	gc.LineTo(bounds.X+bounds.Width, bounds.Y+bounds.Height)
	gc.LineTo(bounds.X, bounds.Y+bounds.Height)
	gc.LineTo(bounds.X, bounds.Y)
	gc.MoveTo(clip.X, clip.Y)
	gc.LineTo(clip.X+clip.Width, clip.Y)
	gc.LineTo(clip.X+clip.Width, clip.Y+clip.Height)
	gc.LineTo(clip.X, clip.Y+clip.Height)
	gc.LineTo(clip.X, clip.Y)
	gc.SetFillRule(draw.FillRuleEvenOdd)
	gc.Clip()
	gc.SetColor(line.color)
	gc.FillRect(bounds)
	gc.Restore()
}
