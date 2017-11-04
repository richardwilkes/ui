// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package label

import (
	"fmt"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/font"
	"github.com/richardwilkes/ui/widget"
)

// Label represents a non-interactive piece of text.
type Label struct {
	widget.Block
	text       string
	font       *font.Font
	foreground color.Color
}

// New creates a label with the specified text.
func New(text string) *Label {
	return NewWithFont(text, font.Label)
}

// NewWithFont creates a label with the specified text and font.
func NewWithFont(text string, font *font.Font) *Label {
	label := &Label{text: text, foreground: color.Black, font: font}
	label.InitTypeAndID(label)
	label.Describer = func() string { return fmt.Sprintf("Label #%d (%s)", label.ID(), label.text) }
	label.SetSizer(label)
	label.EventHandlers().Add(event.PaintType, label.paint)
	return label
}

// Sizes implements Sizer
func (label *Label) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	size := label.font.Measure(label.text)
	size.GrowToInteger()
	size.ConstrainForHint(hint)
	if border := label.Border(); border != nil {
		size.AddInsets(border.Insets())
	}
	return size, size, size
}

func (label *Label) paint(evt event.Event) {
	bounds := label.LocalInsetBounds()
	gc := evt.(*event.Paint).GC()
	gc.SetColor(label.foreground)
	size := label.font.Measure(label.text)
	gc.DrawString(bounds.X, bounds.Y+(bounds.Height-size.Height)/2, label.text, label.font)
}

// SetForeground sets the color used when drawing the text.
func (label *Label) SetForeground(color color.Color) {
	if label.foreground != color {
		label.foreground = color
		label.Repaint()
	}
}
