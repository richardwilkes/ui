// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

// RadioButton represents a radio button with an optional label.
type RadioButton struct {
	Block
	Theme    *RadioButtonTheme // The theme the button will use to draw itself.
	OnClick  func()            // Called when the button is clicked.
	Title    string            // An optional title for the button.
	group    *RadioButtonGroup
	selected bool
	pressed  bool
}

// NewRadioButton creates a new radio button with the specified title.
func NewRadioButton(title string) *RadioButton {
	button := &RadioButton{}
	button.Title = title
	button.Theme = StdRadioButtonTheme
	button.SetSizer(button)
	button.SetPaintHandler(button)
	button.SetMouseDownHandler(button)
	button.SetMouseDraggedHandler(button)
	button.SetMouseUpHandler(button)
	return button
}

// Sizes implements Sizer
func (button *RadioButton) Sizes(hint Size) (min, pref, max Size) {
	var size Size
	box := button.Theme.Font.Ascent()
	if button.Title != "" {
		if hint.Width != NoLayoutHint {
			hint.Width -= button.Theme.HorizontalGap + box
			if hint.Width < 1 {
				hint.Width = 1
			}
		}
		if hint.Height != NoLayoutHint {
			if hint.Height < 1 {
				hint.Height = 1
			}
		}
		size, _ = button.attributedString().MeasureConstrained(hint)
		size.Width += button.Theme.HorizontalGap + box
		if size.Height < box {
			size.Height = box
		}
	} else {
		size.Width = box
		size.Height = box
	}
	if border := button.Border(); border != nil {
		size.AddInsets(border.Insets())
	}
	return size, size, DefaultLayoutMaxSize(size)
}

// OnPaint implements PaintHandler
func (button *RadioButton) OnPaint(g Graphics, dirty Rect) {
	box := button.Theme.Font.Ascent()
	bounds := button.LocalInsetBounds()
	bounds.Width = box
	bounds.Y += (bounds.Height - box) / 2
	bounds.Height = box
	path := NewPath()
	path.Ellipse(bounds)
	g.AddPath(path)
	g.Save()
	g.Clip()
	base := button.BaseBackground()
	if button.Enabled() {
		g.DrawLinearGradient(button.Theme.Gradient(base), bounds.X+bounds.Width/2, bounds.Y+1, bounds.X+bounds.Width/2, bounds.Y+bounds.Height-1)
	} else {
		g.SetFillColor(BackgroundColor)
		g.FillRect(bounds)
	}
	g.AddPath(path)
	c := base.AdjustBrightness(button.Theme.OutlineAdjustment)
	g.SetStrokeColor(c)
	g.StrokePath()
	g.Restore()
	if button.selected {
		bounds.InsetUniform(0.2 * box)
		if button.Enabled() {
			c = KeyboardFocusColor
		}
		g.SetFillColor(c)
		g.FillEllipse(bounds)
	}
	if button.Title != "" {
		bounds = button.LocalInsetBounds()
		bounds.X += box + button.Theme.HorizontalGap
		bounds.Width -= box + button.Theme.HorizontalGap
		if bounds.Width > 0 {
			g.DrawAttributedTextConstrained(bounds, button.attributedString(), TextModeFill)
		}
	}
}

// OnMouseDown implements MouseDownHandler
func (button *RadioButton) OnMouseDown(where Point, keyModifiers int, which int, clickCount int) bool {
	button.pressed = true
	button.Repaint()
	return false
}

// OnMouseDragged implements MouseDraggedHandler
func (button *RadioButton) OnMouseDragged(where Point, keyModifiers int) {
	bounds := button.LocalInsetBounds()
	pressed := bounds.Contains(where)
	if button.pressed != pressed {
		button.pressed = pressed
		button.Repaint()
	}
}

// OnMouseUp implements MouseUpHandler
func (button *RadioButton) OnMouseUp(where Point, keyModifiers int) {
	button.pressed = false
	button.SetSelected(true)
	button.Repaint()
	if button.OnClick != nil {
		bounds := button.LocalInsetBounds()
		if bounds.Contains(where) {
			button.OnClick()
		}
	}
}

// BaseBackground returns this button's current base background color.
func (button *RadioButton) BaseBackground() Color {
	switch {
	case !button.Enabled():
		return button.Theme.Background.AdjustBrightness(button.Theme.DisabledAdjustment)
	case button.pressed:
		return button.Theme.BackgroundWhenPressed
	case button.Focused():
		return button.Theme.Background.Blend(KeyboardFocusColor, 0.5)
	default:
		return button.Theme.Background
	}
}

// TextColor returns this button's current text color.
func (button *RadioButton) TextColor() Color {
	if button.Enabled() {
		return button.Theme.TextWhenLight
	}
	return button.Theme.TextWhenDisabled
}

func (button *RadioButton) attributedString() *AttributedString {
	str := NewAttributedString(button.Title, button.TextColor(), button.Theme.Font)
	str.SetAlignment(0, 0, AlignStart)
	return str
}

// Selected returns true if the radio button is currently selected.
func (button *RadioButton) Selected() bool {
	return button.selected
}

// SetSelected sets the button's selected state.
func (button *RadioButton) SetSelected(selected bool) {
	if button.group != nil {
		button.group.Select(button)
	} else {
		button.setSelected(selected)
	}
}

func (button *RadioButton) setSelected(selected bool) {
	if button.selected != selected {
		button.selected = selected
		button.Repaint()
	}
}
