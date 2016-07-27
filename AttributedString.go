// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

import (
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/font"
)

const (
	fontAttribute       = "font"
	foregroundAttribute = "foreground"
	alignmentAttribute  = "align"
)

type attribute struct {
	key   string
	value interface{}
}

type rangedAttribute struct {
	begin  int
	length int
	attr   attribute
}

// AttributedString contains a string with attributes attached to ranges within it.
type AttributedString struct {
	text       string
	attributes []*rangedAttribute
}

// NewAttributedString creates a new attributed string with the specified font and color for the
// whole string.
func NewAttributedString(text string, color color.Color, font *font.Font) *AttributedString {
	as := &AttributedString{text: text}
	as.SetFont(0, 0, font)
	as.SetForeground(0, 0, color)
	return as
}

// SetFont sets the font for the specified range.
func (a *AttributedString) SetFont(begin, length int, font *font.Font) {
	a.set(begin, length, fontAttribute, font)
}

// SetForeground sets the foreground color for the specified range.
func (a *AttributedString) SetForeground(begin, length int, color color.Color) {
	a.set(begin, length, foregroundAttribute, color)
}

// SetAlignment sets the text alignment for the specified range.
func (a *AttributedString) SetAlignment(begin, length int, align Alignment) {
	a.set(begin, length, alignmentAttribute, align)
}

func (a *AttributedString) set(begin, length int, key string, value interface{}) {
	a.attributes = append(a.attributes, &rangedAttribute{begin: begin, length: length, attr: attribute{key: key, value: value}})
}

// Measure determines the space required to display the specified text.
func (a *AttributedString) Measure() Size {
	size, _ := a.MeasureConstrained(Size{Width: -1, Height: -1})
	return size
}

// MeasureConstrained determines the space required to display the specified text. The size parameter
// allow you to constrain the space, wrapping the text within. If the size parameter contains a negative
// number, it will be treated as if no constrait were passed in for that value. Returns the space
// required as well as the amount of text that fits in that space.
func (a *AttributedString) MeasureConstrained(size Size) (actual Size, fit int) {
	return a.platformMeasure(size)
}
