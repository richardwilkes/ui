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
	"github.com/richardwilkes/go-ui/geom"
	"github.com/richardwilkes/go-ui/graphics"
	"github.com/richardwilkes/go-ui/layout"
)

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
	button.Init(title)
	return button
}

// Init initializes the radio button.
func (button *RadioButton) Init(title string) {
	button.Block.Init()
	button.Title = title
	button.Theme = StdRadioButtonTheme
	button.Sizes = func(hint geom.Size) (min, pref, max geom.Size) {
		var size geom.Size
		box := button.Theme.Font.Ascent()
		if button.Title != "" {
			if hint.Width != layout.NoHint {
				hint.Width -= button.Theme.HorizontalGap + box
				if hint.Width < 1 {
					hint.Width = 1
				}
			}
			if hint.Height != layout.NoHint {
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
		insets := button.Insets()
		size.AddInsets(insets)
		return size, size, layout.DefaultMaxSize(size)
	}
	button.OnMouseDown = func(where geom.Point, keyModifiers int, which int, clickCount int) bool {
		button.pressed = true
		button.Repaint()
		return false
	}
	button.OnMouseDragged = func(where geom.Point, keyModifiers int) {
		bounds := button.LocalInsetBounds()
		pressed := bounds.Contains(where)
		if button.pressed != pressed {
			button.pressed = pressed
			button.Repaint()
		}
	}
	button.OnMouseUp = func(where geom.Point, keyModifiers int) {
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
	button.OnPaint = func(gc graphics.Context, dirty geom.Rect, inLiveResize bool) {
		box := button.Theme.Font.Ascent()
		bounds := button.LocalInsetBounds()
		bounds.Width = box
		bounds.Y += (bounds.Height - box) / 2
		bounds.Height = box
		path := graphics.NewPath()
		path.Ellipse(bounds)
		gc.AddPath(path)
		gc.Save()
		gc.Clip()
		base := button.BaseBackground()
		if button.Enabled() {
			gc.DrawLinearGradient(button.Theme.Gradient(base), bounds.X+bounds.Width/2, bounds.Y+1, bounds.X+bounds.Width/2, bounds.Y+bounds.Height-1)
		} else {
			gc.SetFillColor(color.Background)
			gc.FillRect(bounds)
		}
		gc.AddPath(path)
		c := base.AdjustBrightness(button.Theme.OutlineAdjustment)
		gc.SetStrokeColor(c)
		gc.StrokePath()
		gc.Restore()
		if button.selected {
			bounds.InsetUniform(0.2 * box)
			if button.Enabled() {
				c = color.KeyboardFocus
			}
			gc.SetFillColor(c)
			gc.FillEllipse(bounds)
		}
		if button.Title != "" {
			bounds = button.LocalInsetBounds()
			bounds.X += box + button.Theme.HorizontalGap
			bounds.Width -= box + button.Theme.HorizontalGap
			if bounds.Width > 0 {
				gc.DrawAttributedTextConstrained(bounds, button.attributedString(), graphics.TextFill)
			}
		}
	}
}

// BaseBackground returns this button's current base background color.
func (button *RadioButton) BaseBackground() color.Color {
	switch {
	case button.Disabled():
		return button.Theme.Background.AdjustBrightness(button.Theme.DisabledAdjustment)
	case button.pressed:
		return button.Theme.BackgroundWhenPressed
	case button.Focused():
		return button.Theme.Background.Blend(color.KeyboardFocus, 0.5)
	default:
		return button.Theme.Background
	}
}

// TextColor returns this button's current text color.
func (button *RadioButton) TextColor() color.Color {
	if button.Disabled() {
		return button.Theme.TextWhenDisabled
	}
	return button.Theme.TextWhenLight
}

func (button *RadioButton) attributedString() *graphics.AttributedString {
	str := graphics.NewAttributedString(button.Title, button.TextColor(), button.Theme.Font)
	str.SetAlignment(0, 0, graphics.AlignLeft)
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
