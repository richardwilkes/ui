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
	"github.com/richardwilkes/ui/keys"
)

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
	checkbox.Title = title
	checkbox.Theme = StdCheckBoxTheme
	checkbox.SetFocusable(true)
	checkbox.SetSizer(checkbox)
	checkbox.AddEventHandler(PaintEvent, checkbox.paint)
	checkbox.AddEventHandler(MouseDownEvent, checkbox.mouseDown)
	checkbox.AddEventHandler(MouseDraggedEvent, checkbox.mouseDragged)
	checkbox.AddEventHandler(MouseUpEvent, checkbox.mouseUp)
	checkbox.AddEventHandler(FocusGainedEvent, checkbox.focusChanged)
	checkbox.AddEventHandler(FocusLostEvent, checkbox.focusChanged)
	checkbox.AddEventHandler(KeyDownEvent, checkbox.keyDown)
	return checkbox
}

// Sizes implements Sizer
func (checkbox *CheckBox) Sizes(hint Size) (min, pref, max Size) {
	var size Size
	box := CeilFloat32(checkbox.Theme.Font.Ascent())
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
		size.GrowToInteger()
		size.Width += checkbox.Theme.HorizontalGap + box
		if size.Height < box {
			size.Height = box
		}
	} else {
		size.Width = box
		size.Height = box
	}
	if border := checkbox.Border(); border != nil {
		size.AddInsets(border.Insets())
	}
	return size, size, DefaultLayoutMaxSize(size)
}

func (checkbox *CheckBox) paint(event *Event) {
	box := CeilFloat32(checkbox.Theme.Font.Ascent())
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
	gc := event.GC
	gc.AddPath(path)
	gc.Save()
	gc.Clip()
	base := checkbox.BaseBackground()
	if checkbox.Enabled() {
		gc.DrawLinearGradient(checkbox.Theme.Gradient(base), bounds.X+bounds.Width/2, bounds.Y+1, bounds.X+bounds.Width/2, bounds.Y+bounds.Height-1)
	} else {
		gc.SetFillColor(BackgroundColor)
		gc.FillRect(bounds)
	}
	gc.AddPath(path)
	gc.SetStrokeColor(base.AdjustBrightness(checkbox.Theme.OutlineAdjustment))
	gc.StrokePath()
	gc.Restore()
	switch checkbox.state {
	case Mixed:
		gc.Save()
		gc.SetStrokeColor(checkbox.stateColor(base))
		gc.SetStrokeWidth(2)
		gc.StrokeLine(bounds.X+bounds.Width*0.25, bounds.Y+bounds.Height*0.5, bounds.X+bounds.Width*0.7, bounds.Y+bounds.Height*0.5)
		gc.Restore()
	case Checked:
		gc.Save()
		gc.SetStrokeColor(checkbox.stateColor(base))
		gc.SetStrokeWidth(2)
		gc.BeginPath()
		gc.MoveTo(bounds.X+bounds.Width*0.25, bounds.Y+bounds.Height*0.55)
		gc.LineTo(bounds.X+bounds.Width*0.45, bounds.Y+bounds.Height*0.7)
		gc.LineTo(bounds.X+bounds.Width*0.75, bounds.Y+bounds.Height*0.3)
		gc.StrokePath()
		gc.Restore()
	}
	if checkbox.Title != "" {
		bounds = checkbox.LocalInsetBounds()
		bounds.X += box + checkbox.Theme.HorizontalGap
		bounds.Width -= box + checkbox.Theme.HorizontalGap
		if bounds.Width > 0 {
			gc.DrawAttributedTextConstrained(bounds, checkbox.attributedString(), TextModeFill)
		}
	}
}

func (checkbox *CheckBox) mouseDown(event *Event) {
	checkbox.pressed = true
	checkbox.Repaint()
}

func (checkbox *CheckBox) mouseDragged(event *Event) {
	bounds := checkbox.LocalInsetBounds()
	pressed := bounds.Contains(checkbox.FromWindow(event.Where))
	if checkbox.pressed != pressed {
		checkbox.pressed = pressed
		checkbox.Repaint()
	}
}

func (checkbox *CheckBox) mouseUp(event *Event) {
	checkbox.pressed = false
	checkbox.Repaint()
	bounds := checkbox.LocalInsetBounds()
	if bounds.Contains(checkbox.FromWindow(event.Where)) {
		checkbox.doClick()
	}
}

func (checkbox *CheckBox) doClick() {
	if checkbox.state == Checked {
		checkbox.state = Unchecked
	} else {
		checkbox.state = Checked
	}
	checkbox.Repaint()
	if checkbox.OnClick != nil {
		checkbox.OnClick()
	}
}

func (checkbox *CheckBox) focusChanged(event *Event) {
	checkbox.Repaint()
}

func (checkbox *CheckBox) keyDown(event *Event) {
	if event.KeyCode == keys.Return || event.KeyCode == keys.Enter || event.KeyCode == keys.Space {
		checkbox.doClick()
	}
}

func (checkbox *CheckBox) stateColor(base Color) Color {
	if !checkbox.Enabled() {
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
	case !checkbox.Enabled():
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
	if !checkbox.Enabled() {
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
