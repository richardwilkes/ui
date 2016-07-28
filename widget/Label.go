// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package widget

import (
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/font"
	"github.com/richardwilkes/ui/geom"
)

// Label represents a non-interactive piece of text.
type Label struct {
	Block
	text       string
	font       *font.Font
	foreground color.Color
	alignment  draw.Alignment
}

// NewLabel creates a label with the specified text.
func NewLabel(text string) *Label {
	return NewLabelWithFont(text, font.Acquire(font.LabelDesc))
}

// NewLabelWithFont creates a label with the specified text and font.
func NewLabelWithFont(text string, font *font.Font) *Label {
	label := &Label{}
	label.text = text
	label.foreground = color.Black
	label.font = font
	label.alignment = draw.AlignStart
	label.SetSizer(label)
	label.EventHandlers().Add(event.Paint, label.paint)
	return label
}

// Sizes implements Sizer
func (label *Label) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	size, _ := label.title().MeasureConstrained(hint)
	size.GrowToInteger()
	if border := label.Border(); border != nil {
		size.AddInsets(border.Insets())
	}
	return size, size, size
}

func (label *Label) paint(event *event.Event) {
	event.GC.DrawAttributedTextConstrained(label.LocalInsetBounds(), label.title(), draw.TextModeFill)
}

// SetForeground sets the color used when drawing the text.
func (label *Label) SetForeground(color color.Color) {
	if label.foreground != color {
		label.foreground = color
		label.Repaint()
	}
}

// SetAlignment sets the alignment used when drawing the text.
func (label *Label) SetAlignment(align draw.Alignment) {
	if label.alignment != align {
		label.alignment = align
		label.Repaint()
	}
}

func (label *Label) title() *draw.Text {
	str := draw.NewText(label.text, label.foreground, label.font)
	str.SetAlignment(0, 0, label.alignment)
	return str
}
