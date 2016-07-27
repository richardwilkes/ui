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
	"github.com/richardwilkes/ui/keys"
	"time"
)

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
	button.Title = title
	button.Theme = StdButtonTheme
	button.SetFocusable(true)
	button.SetSizer(button)
	button.AddEventHandler(PaintEvent, button.paint)
	button.AddEventHandler(MouseDownEvent, button.mouseDown)
	button.AddEventHandler(MouseDraggedEvent, button.mouseDragged)
	button.AddEventHandler(MouseUpEvent, button.mouseUp)
	button.AddEventHandler(FocusGainedEvent, button.focusChanged)
	button.AddEventHandler(FocusLostEvent, button.focusChanged)
	button.AddEventHandler(KeyDownEvent, button.keyDown)
	return button
}

// Sizes implements Sizer
func (button *Button) Sizes(hint Size) (min, pref, max Size) {
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
	size.GrowToInteger()
	size.Width += hSpace
	size.Height += vSpace
	if border := button.Border(); border != nil {
		size.AddInsets(border.Insets())
	}
	return size, size, DefaultLayoutMaxSize(size)
}

func (button *Button) paint(event *Event) {
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
	gc := event.GC
	gc.AddPath(path)
	gc.Clip()
	base := button.BaseBackground()
	gc.DrawLinearGradient(button.Theme.Gradient(base), bounds.X+bounds.Width/2, bounds.Y+1, bounds.X+bounds.Width/2, bounds.Y+bounds.Height-1)
	gc.AddPath(path)
	gc.SetStrokeColor(base.AdjustBrightness(button.Theme.OutlineAdjustment))
	gc.StrokePath()
	bounds.X += button.Theme.HorizontalMargin + 1
	bounds.Y += button.Theme.VerticalMargin + 1
	bounds.Width -= hSpace
	bounds.Height -= vSpace
	gc.DrawAttributedTextConstrained(bounds, button.attributedString(), TextModeFill)
}

func (button *Button) mouseDown(event *Event) {
	button.pressed = true
	button.Repaint()
}

func (button *Button) mouseDragged(event *Event) {
	bounds := button.LocalInsetBounds()
	pressed := bounds.Contains(button.FromWindow(event.Where))
	if button.pressed != pressed {
		button.pressed = pressed
		button.Repaint()
	}
}

func (button *Button) mouseUp(event *Event) {
	button.pressed = false
	button.Repaint()
	bounds := button.LocalInsetBounds()
	if bounds.Contains(button.FromWindow(event.Where)) {
		button.Click()
	}
}

func (button *Button) focusChanged(event *Event) {
	button.Repaint()
}

// Click performs any animation associated with a click and calls the OnClick() function if it is
// set.
func (button *Button) Click() {
	pressed := button.pressed
	button.pressed = true
	button.Repaint()
	button.Window().FlushPainting()
	button.pressed = pressed
	time.Sleep(button.Theme.ClickAnimationTime)
	button.Repaint()
	if button.OnClick != nil {
		button.OnClick()
	}
}

func (button *Button) keyDown(event *Event) {
	if event.KeyCode == keys.Return || event.KeyCode == keys.Enter || event.KeyCode == keys.Space {
		event.Done = true
		button.Click()
	}
}

// BaseBackground returns this button's current base background color.
func (button *Button) BaseBackground() color.Color {
	switch {
	case !button.Enabled():
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
func (button *Button) TextColor() color.Color {
	if !button.Enabled() {
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
