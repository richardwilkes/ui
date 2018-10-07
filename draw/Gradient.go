package draw

import (
	"bytes"
	"fmt"

	"github.com/richardwilkes/ui/color"
)

// ColorStop provides information about the color and position of one 'color stop' in a Gradient.
type ColorStop struct {
	Color    color.Color
	Location float64
}

// String implements the fmt.Stringer interface.
func (cs ColorStop) String() string {
	return fmt.Sprintf("%v:%v", cs.Color, cs.Location)
}

// Gradient defines a smooth transition between colors across an area.
type Gradient struct {
	Stops []ColorStop
}

// NewGradient creates a new gradient.
func NewGradient(stops ...ColorStop) *Gradient {
	return &Gradient{Stops: stops}
}

// NewEvenlySpacedGradient creates a new gradient with the specified colors evenly spread across
// the whole range.
func NewEvenlySpacedGradient(colors ...color.Color) *Gradient {
	count := len(colors)
	gradient := &Gradient{Stops: make([]ColorStop, count)}
	switch count {
	case 0:
	case 1:
		gradient.Stops[0].Color = colors[0]
	case 2:
		gradient.Stops[0].Color = colors[0]
		gradient.Stops[1].Color = colors[1]
		gradient.Stops[1].Location = 1
	default:
		step := 1 / float64(count-1)
		var location float64
		for i, color := range colors {
			gradient.Stops[i].Color = color
			gradient.Stops[i].Location = location
			if i < count-1 {
				location += step
			} else {
				location = 1
			}
		}
	}
	return gradient
}

// Reversed creates a copy of the current Gradient and inverts the locations of each color stop
// in that copy.
func (g *Gradient) Reversed() *Gradient {
	other := &Gradient{Stops: make([]ColorStop, len(g.Stops))}
	for i, stop := range g.Stops {
		stop.Location = 1 - stop.Location
		other.Stops[i] = stop
	}
	return other
}

// String implements the fmt.Stringer interface.
func (g *Gradient) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Gradient[")
	for i, stop := range g.Stops {
		if i != 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprintf(&buffer, "[%v]", stop)
	}
	buffer.WriteString("]")
	return buffer.String()
}
