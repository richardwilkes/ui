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
	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/theme"
	"math"
	"time"
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
	Theme   *theme.CheckBox // The theme the checkbox will use to draw itself.
	Title   string          // An optional title for the checkbox.
	state   CheckBoxState
	pressed bool
}

// NewCheckBox creates a new checkbox with the specified title.
func NewCheckBox(title string) *CheckBox {
	checkbox := &CheckBox{}
	checkbox.Title = title
	checkbox.Theme = theme.StdCheckBox
	checkbox.SetFocusable(true)
	checkbox.SetSizer(checkbox)
	handlers := checkbox.EventHandlers()
	handlers.Add(event.PaintType, checkbox.paint)
	handlers.Add(event.MouseDownType, checkbox.mouseDown)
	handlers.Add(event.MouseDraggedType, checkbox.mouseDragged)
	handlers.Add(event.MouseUpType, checkbox.mouseUp)
	handlers.Add(event.FocusGainedType, checkbox.focusChanged)
	handlers.Add(event.FocusLostType, checkbox.focusChanged)
	handlers.Add(event.KeyDownType, checkbox.keyDown)
	return checkbox
}

func (checkbox *CheckBox) boxSize() float64 {
	return math.Ceil(checkbox.Theme.Font.Ascent() + checkbox.Theme.Font.Descent())
}

// Sizes implements Sizer
func (checkbox *CheckBox) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	var size geom.Size
	box := checkbox.boxSize()
	if checkbox.Title != "" {
		if hint.Width != NoHint {
			hint.Width -= checkbox.Theme.HorizontalGap + box
			if hint.Width < 1 {
				hint.Width = 1
			}
		}
		if hint.Height != NoHint {
			if hint.Height < 1 {
				hint.Height = 1
			}
		}
		size = checkbox.Theme.Font.Measure(checkbox.Title)
		size.GrowToInteger()
		size.ConstrainForHint(hint)
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
	return size, size, DefaultMaxSize(size)
}

func (checkbox *CheckBox) paint(evt event.Event) {
	box := checkbox.boxSize()
	bounds := checkbox.LocalInsetBounds()
	bounds.Width = box
	bounds.Y += (bounds.Height - box) / 2
	bounds.Height = box
	path := draw.NewPath()
	path.MoveTo(bounds.X, bounds.Y+checkbox.Theme.CornerRadius)
	path.QuadCurveTo(bounds.X, bounds.Y, bounds.X+checkbox.Theme.CornerRadius, bounds.Y)
	path.LineTo(bounds.X+bounds.Width-checkbox.Theme.CornerRadius, bounds.Y)
	path.QuadCurveTo(bounds.X+bounds.Width, bounds.Y, bounds.X+bounds.Width, bounds.Y+checkbox.Theme.CornerRadius)
	path.LineTo(bounds.X+bounds.Width, bounds.Y+bounds.Height-checkbox.Theme.CornerRadius)
	path.QuadCurveTo(bounds.X+bounds.Width, bounds.Y+bounds.Height, bounds.X+bounds.Width-checkbox.Theme.CornerRadius, bounds.Y+bounds.Height)
	path.LineTo(bounds.X+checkbox.Theme.CornerRadius, bounds.Y+bounds.Height)
	path.QuadCurveTo(bounds.X, bounds.Y+bounds.Height, bounds.X, bounds.Y+bounds.Height-checkbox.Theme.CornerRadius)
	path.ClosePath()
	gc := evt.(*event.Paint).GC()
	gc.AddPath(path)
	gc.Save()
	gc.Clip()
	base := checkbox.BaseBackground()
	gc.AddPath(path)
	if checkbox.Enabled() {
		paint := draw.NewLinearGradientPaint(checkbox.Theme.Gradient(base), bounds.X+bounds.Width/2, bounds.Y+1, bounds.X+bounds.Width/2, bounds.Y+bounds.Height-1)
		gc.SetPaint(paint)
		gc.FillPath()
		paint.Dispose()
	} else {
		gc.SetColor(color.Background)
		gc.FillPath()
	}
	gc.AddPath(path)
	gc.SetColor(base.AdjustBrightness(checkbox.Theme.OutlineAdjustment))
	gc.StrokePath()
	gc.Restore()
	switch checkbox.state {
	case Mixed:
		gc.Save()
		gc.SetColor(checkbox.stateColor(base))
		gc.SetStrokeWidth(2)
		gc.StrokeLine(bounds.X+bounds.Width*0.25, bounds.Y+bounds.Height*0.5, bounds.X+bounds.Width*0.7, bounds.Y+bounds.Height*0.5)
		gc.Restore()
	case Checked:
		gc.Save()
		gc.SetColor(checkbox.stateColor(base))
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
		if bounds.Width-(box+checkbox.Theme.HorizontalGap) > 0 {
			gc.SetColor(checkbox.TextColor())
			gc.DrawString(bounds.X+box+checkbox.Theme.HorizontalGap, bounds.Y, checkbox.Title, checkbox.Theme.Font)
		}
	}
}

func (checkbox *CheckBox) mouseDown(evt event.Event) {
	checkbox.pressed = true
	checkbox.Repaint()
}

func (checkbox *CheckBox) mouseDragged(evt event.Event) {
	bounds := checkbox.LocalInsetBounds()
	pressed := bounds.Contains(checkbox.FromWindow(evt.(*event.MouseDragged).Where()))
	if checkbox.pressed != pressed {
		checkbox.pressed = pressed
		checkbox.Repaint()
	}
}

func (checkbox *CheckBox) mouseUp(evt event.Event) {
	checkbox.pressed = false
	checkbox.Repaint()
	bounds := checkbox.LocalInsetBounds()
	if bounds.Contains(checkbox.FromWindow(evt.(*event.MouseUp).Where())) {
		checkbox.Click()
	}
}

// Click performs any animation associated with a click and calls the OnClick() function if it is
// set.
func (checkbox *CheckBox) Click() {
	if checkbox.state == Checked {
		checkbox.state = Unchecked
	} else {
		checkbox.state = Checked
	}
	pressed := checkbox.pressed
	checkbox.pressed = true
	checkbox.Repaint()
	checkbox.Window().FlushPainting()
	checkbox.pressed = pressed
	time.Sleep(checkbox.Theme.ClickAnimationTime)
	checkbox.Repaint()
	event.Dispatch(event.NewClick(checkbox))
}

func (checkbox *CheckBox) focusChanged(evt event.Event) {
	checkbox.Repaint()
}

func (checkbox *CheckBox) keyDown(evt event.Event) {
	if keys.IsControlAction(evt.(*event.KeyDown).Code()) {
		evt.Finish()
		checkbox.Click()
	}
}

func (checkbox *CheckBox) stateColor(base color.Color) color.Color {
	if !checkbox.Enabled() {
		return checkbox.Theme.TextWhenDisabled
	}
	if checkbox.BaseBackground().Luminance() > 0.65 {
		return checkbox.Theme.TextWhenLight
	}
	return checkbox.Theme.TextWhenDark
}

// BaseBackground returns this checkbox's current base background color.
func (checkbox *CheckBox) BaseBackground() color.Color {
	switch {
	case !checkbox.Enabled():
		return checkbox.Theme.Background.AdjustBrightness(checkbox.Theme.DisabledAdjustment)
	case checkbox.pressed:
		return checkbox.Theme.BackgroundWhenPressed
	case checkbox.Focused():
		return checkbox.Theme.Background.Blend(color.KeyboardFocus, 0.5)
	default:
		return checkbox.Theme.Background
	}
}

// TextColor returns this checkbox's current text color.
func (checkbox *CheckBox) TextColor() color.Color {
	if !checkbox.Enabled() {
		return checkbox.Theme.TextWhenDisabled
	}
	return checkbox.Theme.TextWhenLight
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
