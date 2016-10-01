// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package imagebutton

import (
	"fmt"
	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/widget"
	"time"
)

// ImageButton represents a clickable image button.
type ImageButton struct {
	widget.Block
	Theme         *Theme // The theme the button will use to draw itself.
	image         *draw.Image
	disabledImage *draw.Image
	pressed       bool
}

// NewImageButton creates a new button with the specified Image.
func NewImageButton(img *draw.Image) *ImageButton {
	return NewImageButtonWithImageSize(img, geom.Size{})
}

// NewImageButtonWithImageSize creates a new button with the specified Image. The image will be set
// to the specified size. The button itself will be a bit larger, based on the theme settings and
// border.
func NewImageButtonWithImageSize(img *draw.Image, size geom.Size) *ImageButton {
	button := &ImageButton{}
	button.image = img
	var err error
	if button.disabledImage, err = img.AcquireDisabled(); err != nil {
		button.disabledImage = img
	}
	button.Theme = StdImageButton
	button.Describer = func() string { return fmt.Sprintf("ImageButton #%d (%v)", button.ID(), button.Image()) }
	button.SetFocusable(true)
	if size.Width <= 0 || size.Height <= 0 {
		button.SetSizer(button)
	} else {
		button.SetSizer(&imageButtonSizer{button: button, size: size})
	}
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
func (button *ImageButton) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	size := button.image.Size()
	size.Width += button.Theme.HorizontalMargin*2 + 2
	size.Height += button.Theme.VerticalMargin*2 + 2
	if border := button.Border(); border != nil {
		size.AddInsets(border.Insets())
	}
	return size, size, size
}

func (button *ImageButton) paint(evt event.Event) {
	var hSpace = button.Theme.HorizontalMargin*2 + 2
	var vSpace = button.Theme.VerticalMargin*2 + 2
	bounds := button.LocalInsetBounds()
	path := draw.NewPath()
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
	gc.AddPath(path)
	paint := draw.NewLinearGradientPaint(button.Theme.Gradient(base), bounds.X+bounds.Width/2, bounds.Y+1, bounds.X+bounds.Width/2, bounds.Y+bounds.Height-1)
	gc.SetPaint(paint)
	gc.FillPath()
	paint.Dispose()
	gc.AddPath(path)
	gc.SetColor(base.AdjustBrightness(button.Theme.OutlineAdjustment))
	gc.StrokePath()
	bounds.X += button.Theme.HorizontalMargin + 1
	bounds.Y += button.Theme.VerticalMargin + 1
	bounds.Width -= hSpace
	bounds.Height -= vSpace
	if !bounds.IsEmpty() {
		img := button.CurrentImage()
		size := img.Size()
		if size.Width < bounds.Width {
			bounds.X += (bounds.Width - size.Width) / 2
			bounds.Width = size.Width
		}
		if size.Height < bounds.Height {
			bounds.Y += (bounds.Height - size.Height) / 2
			bounds.Height = size.Height
		}
		gc.DrawImageInRect(img, bounds)
	}
}

func (button *ImageButton) mouseDown(evt event.Event) {
	button.pressed = true
	button.Repaint()
}

func (button *ImageButton) mouseDragged(evt event.Event) {
	bounds := button.LocalInsetBounds()
	pressed := bounds.Contains(button.FromWindow(evt.(*event.MouseDragged).Where()))
	if button.pressed != pressed {
		button.pressed = pressed
		button.Repaint()
	}
}

func (button *ImageButton) mouseUp(evt event.Event) {
	button.pressed = false
	button.Repaint()
	bounds := button.LocalInsetBounds()
	if bounds.Contains(button.FromWindow(evt.(*event.MouseUp).Where())) {
		button.Click()
	}
}

func (button *ImageButton) focusChanged(evt event.Event) {
	button.Repaint()
}

// Click performs any animation associated with a click and calls the OnClick() function if it is
// set.
func (button *ImageButton) Click() {
	pressed := button.pressed
	button.pressed = true
	button.Repaint()
	button.Window().FlushPainting()
	button.pressed = pressed
	time.Sleep(button.Theme.ClickAnimationTime)
	button.Repaint()
	event.Dispatch(event.NewClick(button))
}

func (button *ImageButton) keyDown(evt event.Event) {
	if keys.IsControlAction(evt.(*event.KeyDown).Code()) {
		evt.Finish()
		button.Click()
	}
}

// Image returns this button's base image.
func (button *ImageButton) Image() *draw.Image {
	return button.image
}

// CurrentImage returns this button's current image.
func (button *ImageButton) CurrentImage() *draw.Image {
	if button.Enabled() {
		return button.image
	}
	return button.disabledImage
}

// BaseBackground returns this button's current base background color.
func (button *ImageButton) BaseBackground() color.Color {
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

type imageButtonSizer struct {
	button *ImageButton
	size   geom.Size
}

// Sizes implements Sizer
func (ibs *imageButtonSizer) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	pref = ibs.size
	pref.Width += ibs.button.Theme.HorizontalMargin*2 + 2
	pref.Height += ibs.button.Theme.VerticalMargin*2 + 2
	if border := ibs.button.Border(); border != nil {
		pref.AddInsets(border.Insets())
	}
	return pref, pref, pref
}
