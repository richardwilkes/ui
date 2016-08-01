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
	watermark       string
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
	if hint.Width != layout.NoHint {
		if hint.Width < field.Theme.MinimumTextWidth {
			hint.Width = field.Theme.MinimumTextWidth
		}
	}
	if hint.Height != layout.NoHint {
		if hint.Height < 1 {
			hint.Height = 1
		}
	}
	var text string
	if field.text == "" {
		text = "M"
	} else {
		text = field.text
	}
	size, _ := field.createAttributedText(text, color.Black).MeasureConstrained(hint)
	size.GrowToInteger()
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
	var text *draw.Text
	if field.text == "" {
		text = field.attributedWatermark()
	} else {
		text = field.attributedText()
	}
	gc.DrawAttributedTextConstrained(bounds, text, draw.TextModeFill)
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

func (field *TextField) Watermark() string {
	return field.watermark
}

func (field *TextField) SetWatermark(text string) {
	field.watermark = text
	field.Repaint()
}

func (field *TextField) attributedText() *draw.Text {
	return field.createAttributedText(field.text, color.Text)
}

func (field *TextField) attributedWatermark() *draw.Text {
	return field.createAttributedText(field.watermark, color.Gray)
}

func (field *TextField) createAttributedText(text string, textColor color.Color) *draw.Text {
	str := draw.NewText(text, textColor, field.Theme.Font)
	str.SetAlignment(0, 0, field.align)
	return str
}
