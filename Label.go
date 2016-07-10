// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

// Label represents a non-interactive piece of text.
type Label struct {
	Block
	text       string
	font       *Font
	foreground Color
	alignment  Alignment
}

// NewLabel creates a label with the specified text.
func NewLabel(text string) *Label {
	return NewLabelWithFont(text, AcquireFont(LabelFontDesc))
}

// NewLabelWithFont creates a label with the specified text and font.
func NewLabelWithFont(text string, font *Font) *Label {
	label := &Label{}
	label.Init(text, BlackColor, font, AlignStart)
	return label
}

// Init initializes the label.
func (label *Label) Init(text string, foreground Color, font *Font, alignment Alignment) {
	label.Block.Init()
	label.text = text
	label.foreground = foreground
	label.font = font
	label.alignment = alignment
	label.OnPaint = func(g Graphics, dirty Rect, inLiveResize bool) {
		g.DrawAttributedTextConstrained(label.LocalInsetBounds(), label.attributedString(), TextModeFill)
	}
	label.Sizes = func(hint Size) (min, pref, max Size) {
		insets := label.Insets()
		size, _ := label.attributedString().MeasureConstrained(hint)
		size.AddInsets(insets)
		return size, size, size
	}
}

// SetAlignment sets the alignment used when drawing the text.
func (label *Label) SetAlignment(align Alignment) {
	if label.alignment != align {
		label.alignment = align
		label.Repaint()
	}
}

func (label *Label) attributedString() *AttributedString {
	str := NewAttributedString(label.text, label.foreground, label.font)
	str.SetAlignment(0, 0, label.alignment)
	return str
}
