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
	label.text = text
	label.foreground = BlackColor
	label.font = font
	label.alignment = AlignStart
	label.SetSizer(label)
	label.SetPaintHandler(label)
	return label
}

// Sizes implements Sizer
func (label *Label) Sizes(hint Size) (min, pref, max Size) {
	size, _ := label.attributedString().MeasureConstrained(hint)
	size.GrowToInteger()
	if border := label.Border(); border != nil {
		size.AddInsets(border.Insets())
	}
	return size, size, size
}

// OnPaint implements PaintHandler
func (label *Label) OnPaint(g Graphics, dirty Rect) {
	g.DrawAttributedTextConstrained(label.LocalInsetBounds(), label.attributedString(), TextModeFill)
}

// SetForeground sets the color used when drawing the text.
func (label *Label) SetForeground(color Color) {
	if label.foreground != color {
		label.foreground = color
		label.Repaint()
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
