// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

// Possible values for CheckBoxState.
const (
	Unchecked CheckBoxState = iota
	Mixed
	Checked
)

// CheckBoxState represents the current state of the checkbox.
type CheckBoxState int

// CheckBox represents a clickable checkbox with an optional label.
type CheckBox struct {
	Block
	Theme   *CheckBoxTheme // The theme the checkbox will use to draw itself.
	OnClick func()         // Called when the checkbox is clicked.
	Title   string         // An optional title for the checkbox.
	state   CheckBoxState
	pressed bool
}

// NewCheckBox creates a new checkbox with the specified title.
func NewCheckBox(title string) *CheckBox {
	checkbox := &CheckBox{}
	checkbox.Init(title)
	return checkbox
}

// Init initializes the checkbox.
func (checkbox *CheckBox) Init(title string) {
	checkbox.Block.Init()
	checkbox.Title = title
	checkbox.Theme = StdCheckBoxTheme
	checkbox.Sizes = func(hint Size) (min, pref, max Size) {
		var size Size
		box := checkbox.Theme.Font.Ascent()
		if checkbox.Title != "" {
			if hint.Width != NoLayoutHint {
				hint.Width -= checkbox.Theme.HorizontalGap + box
				if hint.Width < 1 {
					hint.Width = 1
				}
			}
			if hint.Height != NoLayoutHint {
				if hint.Height < 1 {
					hint.Height = 1
				}
			}
			size, _ = checkbox.attributedString().MeasureConstrained(hint)
			size.Width += checkbox.Theme.HorizontalGap + box
			if size.Height < box {
				size.Height = box
			}
		} else {
			size.Width = box
			size.Height = box
		}
		insets := checkbox.Insets()
		size.AddInsets(insets)
		return size, size, DefaultLayoutMaxSize(size)
	}
	checkbox.OnMouseDown = func(where Point, keyModifiers int, which int, clickCount int) bool {
		checkbox.pressed = true
		checkbox.Repaint()
		return false
	}
	checkbox.OnMouseDragged = func(where Point, keyModifiers int) {
		bounds := checkbox.LocalInsetBounds()
		pressed := bounds.Contains(where)
		if checkbox.pressed != pressed {
			checkbox.pressed = pressed
			checkbox.Repaint()
		}
	}
	checkbox.OnMouseUp = func(where Point, keyModifiers int) {
		checkbox.pressed = false
		if checkbox.state == Checked {
			checkbox.state = Unchecked
		} else {
			checkbox.state = Checked
		}
		checkbox.Repaint()
		if checkbox.OnClick != nil {
			bounds := checkbox.LocalInsetBounds()
			if bounds.Contains(where) {
				checkbox.OnClick()
			}
		}
	}
	checkbox.OnPaint = func(g Graphics, dirty Rect) {
		box := checkbox.Theme.Font.Ascent()
		bounds := checkbox.LocalInsetBounds()
		bounds.Width = box
		bounds.Y += (bounds.Height - box) / 2
		bounds.Height = box
		path := NewPath()
		path.MoveTo(bounds.X, bounds.Y+checkbox.Theme.CornerRadius)
		path.QuadCurveTo(bounds.X, bounds.Y, bounds.X+checkbox.Theme.CornerRadius, bounds.Y)
		path.LineTo(bounds.X+bounds.Width-checkbox.Theme.CornerRadius, bounds.Y)
		path.QuadCurveTo(bounds.X+bounds.Width, bounds.Y, bounds.X+bounds.Width, bounds.Y+checkbox.Theme.CornerRadius)
		path.LineTo(bounds.X+bounds.Width, bounds.Y+bounds.Height-checkbox.Theme.CornerRadius)
		path.QuadCurveTo(bounds.X+bounds.Width, bounds.Y+bounds.Height, bounds.X+bounds.Width-checkbox.Theme.CornerRadius, bounds.Y+bounds.Height)
		path.LineTo(bounds.X+checkbox.Theme.CornerRadius, bounds.Y+bounds.Height)
		path.QuadCurveTo(bounds.X, bounds.Y+bounds.Height, bounds.X, bounds.Y+bounds.Height-checkbox.Theme.CornerRadius)
		path.ClosePath()
		g.AddPath(path)
		g.Save()
		g.Clip()
		base := checkbox.BaseBackground()
		if checkbox.Enabled() {
			g.DrawLinearGradient(checkbox.Theme.Gradient(base), bounds.X+bounds.Width/2, bounds.Y+1, bounds.X+bounds.Width/2, bounds.Y+bounds.Height-1)
		} else {
			g.SetFillColor(BackgroundColor)
			g.FillRect(bounds)
		}
		g.AddPath(path)
		g.SetStrokeColor(base.AdjustBrightness(checkbox.Theme.OutlineAdjustment))
		g.StrokePath()
		g.Restore()
		switch checkbox.state {
		case Mixed:
			g.Save()
			g.SetStrokeColor(checkbox.stateColor(base))
			g.SetStrokeWidth(2)
			g.StrokeLine(bounds.X+bounds.Width*0.25, bounds.Y+bounds.Height*0.5, bounds.X+bounds.Width*0.7, bounds.Y+bounds.Height*0.5)
			g.Restore()
		case Checked:
			g.Save()
			g.SetStrokeColor(checkbox.stateColor(base))
			g.SetStrokeWidth(2)
			g.BeginPath()
			g.MoveTo(bounds.X+bounds.Width*0.25, bounds.Y+bounds.Height*0.55)
			g.LineTo(bounds.X+bounds.Width*0.45, bounds.Y+bounds.Height*0.7)
			g.LineTo(bounds.X+bounds.Width*0.75, bounds.Y+bounds.Height*0.3)
			g.StrokePath()
			g.Restore()
		}
		if checkbox.Title != "" {
			bounds = checkbox.LocalInsetBounds()
			bounds.X += box + checkbox.Theme.HorizontalGap
			bounds.Width -= box + checkbox.Theme.HorizontalGap
			if bounds.Width > 0 {
				g.DrawAttributedTextConstrained(bounds, checkbox.attributedString(), TextModeFill)
			}
		}
	}
}

func (checkbox *CheckBox) stateColor(base Color) Color {
	if checkbox.Disabled() {
		return checkbox.Theme.TextWhenDisabled
	}
	if checkbox.BaseBackground().Luminance() > 0.65 {
		return checkbox.Theme.TextWhenLight
	}
	return checkbox.Theme.TextWhenDark
}

// BaseBackground returns this checkbox's current base background color.
func (checkbox *CheckBox) BaseBackground() Color {
	switch {
	case checkbox.Disabled():
		return checkbox.Theme.Background.AdjustBrightness(checkbox.Theme.DisabledAdjustment)
	case checkbox.pressed:
		return checkbox.Theme.BackgroundWhenPressed
	case checkbox.Focused():
		return checkbox.Theme.Background.Blend(KeyboardFocusColor, 0.5)
	default:
		return checkbox.Theme.Background
	}
}

// TextColor returns this checkbox's current text color.
func (checkbox *CheckBox) TextColor() Color {
	if checkbox.Disabled() {
		return checkbox.Theme.TextWhenDisabled
	}
	return checkbox.Theme.TextWhenLight
}

func (checkbox *CheckBox) attributedString() *AttributedString {
	str := NewAttributedString(checkbox.Title, checkbox.TextColor(), checkbox.Theme.Font)
	str.SetAlignment(0, 0, AlignStart)
	return str
}

// State returns this checkbox's current state.
func (checkbox *CheckBox) State() CheckBoxState {
	return checkbox.state
}

// SetState sets the checkbox's state.
func (checkbox *CheckBox) SetState(state CheckBoxState) {
	if checkbox.state != state {
		checkbox.state = state
		checkbox.Repaint()
	}
}
