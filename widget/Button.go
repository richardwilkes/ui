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
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/geom"
	"github.com/richardwilkes/ui/layout"
	"github.com/richardwilkes/ui/theme"
	"time"
)

// Button represents a clickable text button.
type Button struct {
	Block
	Theme   *theme.Button // The theme the button will use to draw itself.
	Title   string        // The title of the button.
	pressed bool
}

// NewButton creates a new button with the specified title.
func NewButton(title string) *Button {
	button := &Button{}
	button.Title = title
	button.Theme = theme.StdButton
	button.SetFocusable(true)
	button.SetSizer(button)
	handlers := button.EventHandlers()
	handlers.Add(event.PaintType, button.paint)
	handlers.Add(event.MouseDownType, button.mouseDown)
	handlers.Add(event.MouseDraggedType, button.mouseDragged)
	handlers.Add(event.MouseUpType, button.mouseUp)
	handlers.Add(event.FocusGainedType, button.focusChanged)
	handlers.Add(event.FocusLostType, button.focusChanged)
	handlers.Add(event.KeyDownType, button.keyDown)
	return button
}

// Sizes implements Sizer
func (button *Button) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	var hSpace = button.Theme.HorizontalMargin*2 + 2
	var vSpace = button.Theme.VerticalMargin*2 + 2
	if hint.Width != layout.NoHint {
		hint.Width -= hSpace
		if hint.Width < button.Theme.MinimumTextWidth {
			hint.Width = button.Theme.MinimumTextWidth
		}
	}
	if hint.Height != layout.NoHint {
		hint.Height -= vSpace
		if hint.Height < 1 {
			hint.Height = 1
		}
	}
	size := button.Theme.Font.Size(button.Title)
	size.GrowToInteger()
	size.ConstrainForHint(hint)
	size.Width += hSpace
	size.Height += vSpace
	if border := button.Border(); border != nil {
		size.AddInsets(border.Insets())
	}
	return size, size, layout.DefaultMaxSize(size)
}

func (button *Button) paint(evt event.Event) {
	var hSpace = button.Theme.HorizontalMargin*2 + 2
	var vSpace = button.Theme.VerticalMargin*2 + 2
	bounds := button.LocalInsetBounds()
	path := geom.NewPath()
	path.MoveTo(bounds.X, bounds.Y+button.Theme.CornerRadius)
	path.QuadCurveTo(bounds.X, bounds.Y, bounds.X+button.Theme.CornerRadius, bounds.Y)
	path.LineTo(bounds.X+bounds.Width-button.Theme.CornerRadius, bounds.Y)
	path.QuadCurveTo(bounds.X+bounds.Width, bounds.Y, bounds.X+bounds.Width, bounds.Y+button.Theme.CornerRadius)
	path.LineTo(bounds.X+bounds.Width, bounds.Y+bounds.Height-button.Theme.CornerRadius)
	path.QuadCurveTo(bounds.X+bounds.Width, bounds.Y+bounds.Height, bounds.X+bounds.Width-button.Theme.CornerRadius, bounds.Y+bounds.Height)
	path.LineTo(bounds.X+button.Theme.CornerRadius, bounds.Y+bounds.Height)
	path.QuadCurveTo(bounds.X, bounds.Y+bounds.Height, bounds.X, bounds.Y+bounds.Height-button.Theme.CornerRadius)
	path.ClosePath()
	gc := evt.(*event.Paint).GC()
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
	gc.SetFillColor(button.TextColor())
	gc.SetFont(button.Theme.Font)
	gc.DrawString(bounds.X+(bounds.Width-button.Theme.Font.Width(button.Title))/2, bounds.Y+(bounds.Height-button.Theme.Font.Height())/2, button.Title)
}

func (button *Button) mouseDown(evt event.Event) {
	button.pressed = true
	button.Repaint()
}

func (button *Button) mouseDragged(evt event.Event) {
	bounds := button.LocalInsetBounds()
	pressed := bounds.Contains(button.FromWindow(evt.(*event.MouseDragged).Where()))
	if button.pressed != pressed {
		button.pressed = pressed
		button.Repaint()
	}
}

func (button *Button) mouseUp(evt event.Event) {
	button.pressed = false
	button.Repaint()
	bounds := button.LocalInsetBounds()
	mouseUp := evt.(*event.MouseUp)
	if bounds.Contains(button.FromWindow(mouseUp.Where())) {
		button.Click()
	}
}

func (button *Button) focusChanged(evt event.Event) {
	button.Repaint()
}

// Click makes the button behave as if a user clicked on it.
func (button *Button) Click() {
	pressed := button.pressed
	button.pressed = true
	button.Repaint()
	button.Window().FlushPainting()
	button.pressed = pressed
	time.Sleep(button.Theme.ClickAnimationTime)
	button.Repaint()
	event.Dispatch(event.NewClick(button))
}

func (button *Button) keyDown(evt event.Event) {
	if evt.(*event.KeyDown).IsControlActionKey() {
		evt.Finish()
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
