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
	"github.com/richardwilkes/ui/geom"
	"github.com/richardwilkes/ui/layout"
	"github.com/richardwilkes/ui/theme"
)

type TextField struct {
	Block
	Theme           *theme.TextField // The theme the text field will use to draw itself.
	text            string
	selectionStart  int
	selectionEnd    int
	selectionAnchor int
	align           draw.Alignment
}

// NewTextField creates a new, empty, text field.
func NewTextField() *TextField {
	field := &TextField{}
	field.Theme = theme.StdTextField
	field.SetBackground(color.TextBackground)
	field.SetBorder(field.Theme.Border)
	field.SetFocusable(true)
	field.SetSizer(field)
	handlers := field.EventHandlers()
	handlers.Add(event.PaintType, field.paint)
	handlers.Add(event.FocusGainedType, field.focusGained)
	handlers.Add(event.FocusLostType, field.focusLost)
	return field
}

// Sizes implements Sizer
func (field *TextField) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	var hSpace = field.Theme.HorizontalMargin*2 + 2
	var vSpace = field.Theme.VerticalMargin*2 + 2
	if hint.Width != layout.NoHint {
		hint.Width -= hSpace
		if hint.Width < field.Theme.MinimumTextWidth {
			hint.Width = field.Theme.MinimumTextWidth
		}
	}
	if hint.Height != layout.NoHint {
		hint.Height -= vSpace
		if hint.Height < 1 {
			hint.Height = 1
		}
	}
	size, _ := field.attributedText().MeasureConstrained(hint)
	size.Height += field.Theme.Font.Descent() // Give it a more balanced vertical spacing
	size.GrowToInteger()
	size.Width += hSpace
	size.Height += vSpace
	if border := field.Border(); border != nil {
		size.AddInsets(border.Insets())
	}
	return size, size, layout.DefaultMaxSize(size)
}

func (field *TextField) paint(evt event.Event) {
	e := evt.(*event.Paint)
	gc := e.GC()
	if !field.Enabled() && field.Theme.DisabledBackgroundColor.Alpha() > 0 {
		gc.SetFillColor(field.Theme.DisabledBackgroundColor)
		gc.FillRect(e.DirtyRect())
	}
	bounds := field.LocalInsetBounds()
	bounds.X += field.Theme.HorizontalMargin
	// Give it a more balanced vertical spacing by adding in the font descent, otherwise it appears
	// to sit too low, despite being perfectly centered for its overall height.
	bounds.Y += field.Theme.VerticalMargin + field.Theme.Font.Descent()
	bounds.Width -= field.Theme.HorizontalMargin * 2
	bounds.Height -= field.Theme.VerticalMargin * 2
	gc.DrawAttributedTextConstrained(bounds, field.attributedText(), draw.TextModeFill)
}

func (field *TextField) focusGained(evt event.Event) {
	field.SetBorder(field.Theme.FocusBorder)
	field.Repaint()
}

func (field *TextField) focusLost(evt event.Event) {
	field.SetBorder(field.Theme.Border)
	field.Repaint()
}

func (field *TextField) Text() string {
	return field.text
}

func (field *TextField) SetText(text string) {
	field.text = text
	field.Repaint()
}

func (field *TextField) attributedText() *draw.Text {
	str := draw.NewText(field.text, color.Text, field.Theme.Font)
	str.SetAlignment(0, 0, field.align)
	return str
}
