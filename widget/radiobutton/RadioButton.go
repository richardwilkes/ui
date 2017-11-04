// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package radiobutton

import (
	"fmt"
	"math"
	"time"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/layout"
	"github.com/richardwilkes/ui/widget"
)

// RadioButton represents a radio button with an optional label.
type RadioButton struct {
	widget.Block
	Theme    *Theme // The theme the button will use to draw itself.
	Title    string // An optional title for the button.
	group    *Group
	selected bool
	pressed  bool
}

// New creates a new radio button with the specified title.
func New(title string) *RadioButton {
	button := &RadioButton{Title: title, Theme: StdTheme}
	button.InitTypeAndID(button)
	button.Describer = func() string { return fmt.Sprintf("RadioButton #%d (%s)", button.ID(), button.Title) }
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

func (button *RadioButton) circleSize() float64 {
	return math.Ceil(button.Theme.Font.Ascent() + button.Theme.Font.Descent())
}

// Sizes implements Sizer
func (button *RadioButton) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	var size geom.Size
	circle := button.circleSize()
	if button.Title != "" {
		if hint.Width != layout.NoHint {
			hint.Width -= button.Theme.HorizontalGap + circle
			if hint.Width < 1 {
				hint.Width = 1
			}
		}
		if hint.Height != layout.NoHint {
			if hint.Height < 1 {
				hint.Height = 1
			}
		}
		size = button.Theme.Font.Measure(button.Title)
		size.GrowToInteger()
		size.ConstrainForHint(hint)
		size.Width += button.Theme.HorizontalGap + circle
		if size.Height < circle {
			size.Height = circle
		}
	} else {
		size.Width = circle
		size.Height = circle
	}
	if border := button.Border(); border != nil {
		size.AddInsets(border.Insets())
	}
	return size, size, layout.DefaultMaxSize(size)
}

func (button *RadioButton) paint(evt event.Event) {
	circle := button.circleSize()
	bounds := button.LocalInsetBounds()
	bounds.Width = circle
	bounds.Y += (bounds.Height - circle) / 2
	bounds.Height = circle
	path := draw.NewPath()
	path.Ellipse(bounds)
	gc := evt.(*event.Paint).GC()
	gc.AddPath(path)
	gc.Save()
	gc.Clip()
	base := button.BaseBackground()
	gc.AddPath(path)
	if button.Enabled() {
		paint := draw.NewLinearGradientPaint(button.Theme.Gradient(base), bounds.X, bounds.Y, bounds.X, bounds.Y+bounds.Height)
		gc.SetPaint(paint)
		gc.FillPath()
		paint.Dispose()
	} else {
		gc.SetColor(color.Background)
		gc.FillPath()
	}
	gc.AddPath(path)
	c := base.AdjustBrightness(button.Theme.OutlineAdjustment)
	gc.SetColor(c)
	gc.StrokePath()
	gc.Restore()
	if button.selected {
		bounds.InsetUniform(0.2 * circle)
		if button.Enabled() {
			c = color.KeyboardFocus
		}
		gradient := draw.NewEvenlySpacedGradient(c.AdjustBrightness(button.Theme.GradientAdjustment), c.AdjustBrightness(-button.Theme.GradientAdjustment*2))
		paint := draw.NewRadialGradientPaint(gradient, bounds.CenterX(), bounds.CenterY(), bounds.Width/4, bounds.CenterX(), bounds.CenterY(), bounds.Width/2)
		gc.SetPaint(paint)
		gc.Ellipse(bounds)
		gc.FillPath()
		paint.Dispose()
	}
	if button.Title != "" {
		bounds = button.LocalInsetBounds()
		bounds.X += circle + button.Theme.HorizontalGap
		bounds.Width -= circle + button.Theme.HorizontalGap
		if bounds.Width > 0 {
			gc.SetColor(button.TextColor())
			gc.DrawString(bounds.X, bounds.Y+(bounds.Height-button.Theme.Font.Height())/2, button.Title, button.Theme.Font)
		}
	}
}

func (button *RadioButton) mouseDown(evt event.Event) {
	button.pressed = true
	button.Repaint()
}

func (button *RadioButton) mouseDragged(evt event.Event) {
	bounds := button.LocalInsetBounds()
	pressed := bounds.ContainsPoint(button.FromWindow(evt.(*event.MouseDragged).Where()))
	if button.pressed != pressed {
		button.pressed = pressed
		button.Repaint()
	}
}

func (button *RadioButton) mouseUp(evt event.Event) {
	button.pressed = false
	button.Repaint()
	bounds := button.LocalInsetBounds()
	if bounds.ContainsPoint(button.FromWindow(evt.(*event.MouseUp).Where())) {
		button.Click()
	}
}

func (button *RadioButton) focusChanged(evt event.Event) {
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
	event.Dispatch(event.NewClick(button))
}

func (button *RadioButton) keyDown(evt event.Event) {
	if keys.IsControlAction(evt.(*event.KeyDown).Code()) {
		evt.Finish()
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
