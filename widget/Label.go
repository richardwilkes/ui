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
	"github.com/richardwilkes/go-ui/color"
	"github.com/richardwilkes/go-ui/font"
	"github.com/richardwilkes/go-ui/geom"
	"github.com/richardwilkes/go-ui/graphics"
)

// Label represents a non-interactive piece of text.
type Label struct {
	Block
	text       string
	foreground color.Color
	font       *font.Font
	alignment  uint8
}

// NewLabel creates a label with the specified text.
func NewLabel(text string) *Label {
	return NewLabelWithFont(text, font.Acquire(font.Label))
}

// NewLabelWithFont creates a label with the specified text and font.
func NewLabelWithFont(text string, font *font.Font) *Label {
	label := &Label{}
	label.Init(text, color.Black, font, graphics.AlignLeft)
	return label
}

// Init initializes the label.
func (label *Label) Init(text string, foreground color.Color, font *font.Font, alignment uint8) {
	label.Block.Init()
	label.text = text
	label.foreground = foreground
	label.font = font
	label.alignment = alignment
	label.OnPaint = func(gc graphics.Context, dirty geom.Rect, inLiveResize bool) {
		gc.DrawAttributedTextConstrained(label.LocalInsetBounds(), label.attributedString(), graphics.TextFill)
	}
	label.Sizes = func(hint geom.Size) (min, pref, max geom.Size) {
		insets := label.Insets()
		size, _ := label.attributedString().MeasureConstrained(hint)
		size.AddInsets(insets)
		return size, size, size
	}
}

// SetAlignment sets the alignment used when drawing the text.
func (label *Label) SetAlignment(align uint8) {
	if label.alignment != align {
		label.alignment = align
		label.Repaint()
	}
}

func (label *Label) attributedString() *graphics.AttributedString {
	str := graphics.NewAttributedString(label.text, label.foreground, label.font)
	str.SetAlignment(0, 0, label.alignment)
	return str
}
