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
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/layout"
	"github.com/richardwilkes/ui/theme"
	"github.com/richardwilkes/xmath"
	"time"
)

// RadioButton represents a radio button with an optional label.
type RadioButton struct {
	Block
	Theme    *theme.RadioButton // The theme the button will use to draw itself.
	OnClick  func()             // Called when the button is clicked.
	Title    string             // An optional title for the button.
	group    *RadioButtonGroup
	selected bool
	pressed  bool
}

// NewRadioButton creates a new radio button with the specified title.
func NewRadioButton(title string) *RadioButton {
	button := &RadioButton{}
	button.Title = title
	button.Theme = theme.StdRadioButton
	button.SetFocusable(true)
	button.SetSizer(button)
	handlers := button.EventHandlers()
	handlers.Add(event.PaintEvent, button.paint)
	handlers.Add(event.MouseDownEvent, button.mouseDown)
	handlers.Add(event.MouseDraggedEvent, button.mouseDragged)
	handlers.Add(event.MouseUpEvent, button.mouseUp)
	handlers.Add(event.FocusGainedEvent, button.focusChanged)
	handlers.Add(event.FocusLostEvent, button.focusChanged)
	handlers.Add(event.KeyDownEvent, button.keyDown)
	return button
}

// Sizes implements Sizer
func (button *RadioButton) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	var size geom.Size
	box := xmath.CeilFloat32(button.Theme.Font.Ascent())
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
		size, _ = button.title().MeasureConstrained(hint)
		size.GrowToInteger()
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
	return size, size, layout.DefaultMaxSize(size)
}

func (button *RadioButton) paint(event *event.Event) {
	box := xmath.CeilFloat32(button.Theme.Font.Ascent())
	bounds := button.LocalInsetBounds()
	bounds.Width = box
	bounds.Y += (bounds.Height - box) / 2
	bounds.Height = box
	path := geom.NewPath()
	path.Ellipse(bounds)
	gc := event.GC
	gc.AddPath(path)
	gc.Save()
	gc.Clip()
	base := button.BaseBackground()
	if button.Enabled() {
		gc.DrawLinearGradient(button.Theme.Gradient(base), bounds.X, bounds.Y, bounds.X, bounds.Y+bounds.Height)
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
			gc.DrawAttributedTextConstrained(bounds, button.title(), draw.TextModeFill)
		}
	}
}

func (button *RadioButton) mouseDown(event *event.Event) {
	button.pressed = true
	button.Repaint()
}

func (button *RadioButton) mouseDragged(event *event.Event) {
	bounds := button.LocalInsetBounds()
	pressed := bounds.Contains(button.FromWindow(event.Where))
	if button.pressed != pressed {
		button.pressed = pressed
		button.Repaint()
	}
}

func (button *RadioButton) mouseUp(event *event.Event) {
	button.pressed = false
	button.Repaint()
	bounds := button.LocalInsetBounds()
	if bounds.Contains(button.FromWindow(event.Where)) {
		button.Click()
	}
}

func (button *RadioButton) focusChanged(event *event.Event) {
	button.Repaint()
}

// Click performs any animation associated with a click and calls the OnClick() function if it is
// set.
func (button *RadioButton) Click() {
	button.SetSelected(true)
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

func (button *RadioButton) keyDown(event *event.Event) {
	if event.KeyCode == keys.Return || event.KeyCode == keys.Enter || event.KeyCode == keys.Space {
		event.Done = true
		button.Click()
	}
}

// BaseBackground returns this button's current base background color.
func (button *RadioButton) BaseBackground() color.Color {
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
func (button *RadioButton) TextColor() color.Color {
	if button.Enabled() {
		return button.Theme.TextWhenLight
	}
	return button.Theme.TextWhenDisabled
}

func (button *RadioButton) title() *draw.Text {
	str := draw.NewText(button.Title, button.TextColor(), button.Theme.Font)
	str.SetAlignment(0, 0, draw.AlignStart)
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
