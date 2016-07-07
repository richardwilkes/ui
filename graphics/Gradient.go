// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package graphics

import (
	"bytes"
	"fmt"
	"github.com/richardwilkes/go-ui/color"
)

// ColorStop provides information about the color and position of one 'color stop' in a Gradient.
type ColorStop struct {
	Color    color.Color
	Location float32
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
		step := 1 / float32(count-1)
		var location float32
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

// CreateReversed creates a copy of the current Gradient and inverts the locations of each color
// stop in that copy.
func (g *Gradient) CreateReversed() *Gradient {
	other := &Gradient{Stops: make([]ColorStop, len(g.Stops))}
	for i, stop := range g.Stops {
		stop.Location = 1 - stop.Location
		other.Stops[i] = stop
	}
	return other
}

// String -- implements the fmt.Stringer interface.
func (g *Gradient) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Gradient:")
	for _, stop := range g.Stops {
		fmt.Fprintf(&buffer, " %v %v", stop.Color, stop.Location)
	}
	return buffer.String()
}
