// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

// Button represents a clickable text button.
type Button struct {
	Block
	Theme   *ButtonTheme // The theme the button will use to draw itself.
	OnClick func()       // Called when the button is clicked.
	Title   string       // The title of the button.
	pressed bool
}

// NewButton creates a new button with the specified title.
func NewButton(title string) *Button {
	button := &Button{}
	button.Init(title)
	return button
}

// Init initializes the button.
func (button *Button) Init(title string) {
	button.Block.Init()
	button.Title = title
	button.Theme = StdButtonTheme
	button.Sizes = func(hint Size) (min, pref, max Size) {
		var hSpace = button.Theme.HorizontalMargin*2 + 2
		var vSpace = button.Theme.VerticalMargin*2 + 2
		if hint.Width != NoLayoutHint {
			hint.Width -= hSpace
			if hint.Width < button.Theme.MinimumTextWidth {
				hint.Width = button.Theme.MinimumTextWidth
			}
		}
		if hint.Height != NoLayoutHint {
			hint.Height -= vSpace
			if hint.Height < 1 {
				hint.Height = 1
			}
		}
		size, _ := button.attributedString().MeasureConstrained(hint)
		size.Width += hSpace
		size.Height += vSpace
		insets := button.Insets()
		size.AddInsets(insets)
		return size, size, DefaultLayoutMaxSize(size)
	}
	button.OnMouseDown = func(where Point, keyModifiers int, which int, clickCount int) bool {
		button.pressed = true
		button.Repaint()
		return false
	}
	button.OnMouseDragged = func(where Point, keyModifiers int) {
		bounds := button.LocalInsetBounds()
		pressed := bounds.Contains(where)
		if button.pressed != pressed {
			button.pressed = pressed
			button.Repaint()
		}
	}
	button.OnMouseUp = func(where Point, keyModifiers int) {
		button.pressed = false
		button.Repaint()
		if button.OnClick != nil {
			bounds := button.LocalInsetBounds()
			if bounds.Contains(where) {
				button.OnClick()
			}
		}
	}
	button.OnPaint = func(g Graphics, dirty Rect, inLiveResize bool) {
		var hSpace = button.Theme.HorizontalMargin*2 + 2
		var vSpace = button.Theme.VerticalMargin*2 + 2
		bounds := button.LocalInsetBounds()
		path := NewPath()
		path.MoveTo(bounds.X, bounds.Y+button.Theme.CornerRadius)
		path.QuadCurveTo(bounds.X, bounds.Y, bounds.X+button.Theme.CornerRadius, bounds.Y)
		path.LineTo(bounds.X+bounds.Width-button.Theme.CornerRadius, bounds.Y)
		path.QuadCurveTo(bounds.X+bounds.Width, bounds.Y, bounds.X+bounds.Width, bounds.Y+button.Theme.CornerRadius)
		path.LineTo(bounds.X+bounds.Width, bounds.Y+bounds.Height-button.Theme.CornerRadius)
		path.QuadCurveTo(bounds.X+bounds.Width, bounds.Y+bounds.Height, bounds.X+bounds.Width-button.Theme.CornerRadius, bounds.Y+bounds.Height)
		path.LineTo(bounds.X+button.Theme.CornerRadius, bounds.Y+bounds.Height)
		path.QuadCurveTo(bounds.X, bounds.Y+bounds.Height, bounds.X, bounds.Y+bounds.Height-button.Theme.CornerRadius)
		path.ClosePath()
		g.AddPath(path)
		g.Clip()
		base := button.BaseBackground()
		g.DrawLinearGradient(button.Theme.Gradient(base), bounds.X+bounds.Width/2, bounds.Y+1, bounds.X+bounds.Width/2, bounds.Y+bounds.Height-1)
		g.AddPath(path)
		g.SetStrokeColor(base.AdjustBrightness(button.Theme.OutlineAdjustment))
		g.StrokePath()
		bounds.X += button.Theme.HorizontalMargin + 1
		bounds.Y += button.Theme.VerticalMargin + 1
		bounds.Width -= hSpace
		bounds.Height -= vSpace
		g.DrawAttributedTextConstrained(bounds, button.attributedString(), TextModeFill)
	}
}

// BaseBackground returns this button's current base background color.
func (button *Button) BaseBackground() Color {
	switch {
	case button.Disabled():
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
func (button *Button) TextColor() Color {
	if button.Disabled() {
		return button.Theme.TextWhenDisabled
	}
	if button.BaseBackground().Luminance() > 0.65 {
		return button.Theme.TextWhenLight
	}
	return button.Theme.TextWhenDark
}

func (button *Button) attributedString() *AttributedString {
	str := NewAttributedString(button.Title, button.TextColor(), button.Theme.Font)
	str.SetAlignment(0, 0, AlignMiddle)
	return str
}
