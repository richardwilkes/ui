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

// ImageButton represents a clickable image button.
type ImageButton struct {
	Block
	Theme         *ImageButtonTheme // The theme the button will use to draw itself.
	OnClick       func()            // Called when the button is clicked.
	image         *Image
	disabledImage *Image
	pressed       bool
}

// NewImageButton creates a new button with the specified Image.
func NewImageButton(img *Image) *ImageButton {
	return NewImageButtonWithImageSize(img, Size{})
}

// NewImageButtonWithImageSize creates a new button with the specified Image. The image will be set
// to the specified size. The button itself will be a bit larger, based on the theme settings and
// border.
func NewImageButtonWithImageSize(img *Image, size Size) *ImageButton {
	button := &ImageButton{}
	button.image = img
	var err error
	if button.disabledImage, err = img.AcquireDisabled(); err != nil {
		button.disabledImage = img
	}
	button.Theme = StdImageButtonTheme
	button.SetFocusable(true)
	if size.Width <= 0 || size.Height <= 0 {
		button.SetSizer(button)
	} else {
		button.SetSizer(&imageButtonSizer{button: button, size: size})
	}
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
func (button *ImageButton) Sizes(hint Size) (min, pref, max Size) {
	size := button.image.Size()
	size.Width += button.Theme.HorizontalMargin*2 + 2
	size.Height += button.Theme.VerticalMargin*2 + 2
	if border := button.Border(); border != nil {
		size.AddInsets(border.Insets())
	}
	return size, size, size
}

func (button *ImageButton) paint(event *Event) {
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

func (button *ImageButton) mouseDown(event *Event) {
	button.pressed = true
	button.Repaint()
}

func (button *ImageButton) mouseDragged(event *Event) {
	bounds := button.LocalInsetBounds()
	pressed := bounds.Contains(button.FromWindow(event.Where))
	if button.pressed != pressed {
		button.pressed = pressed
		button.Repaint()
	}
}

func (button *ImageButton) mouseUp(event *Event) {
	button.pressed = false
	button.Repaint()
	if button.OnClick != nil {
		bounds := button.LocalInsetBounds()
		if bounds.Contains(button.FromWindow(event.Where)) {
			button.OnClick()
		}
	}
}

func (button *ImageButton) focusChanged(event *Event) {
	button.Repaint()
}

func (button *ImageButton) keyDown(event *Event) {
	if event.KeyCode == keys.Return || event.KeyCode == keys.Enter || event.KeyCode == keys.Space {
		if button.OnClick != nil {
			button.OnClick()
		}
	}
}

// Image returns this button's base image.
func (button *ImageButton) Image() *Image {
	return button.image
}

// CurrentImage returns this button's current image.
func (button *ImageButton) CurrentImage() *Image {
	if button.Enabled() {
		return button.image
	}
	return button.disabledImage
}

// BaseBackground returns this button's current base background color.
func (button *ImageButton) BaseBackground() Color {
	switch {
	case !button.Enabled():
		return button.Theme.Background.AdjustBrightness(button.Theme.DisabledAdjustment)
	case button.pressed:
		return button.Theme.BackgroundWhenPressed
	case button.Focused():
		return button.Theme.Background.Blend(KeyboardFocusColor, 0.5)
	default:
		return button.Theme.Background
	}
}

type imageButtonSizer struct {
	button *ImageButton
	size   Size
}

// Sizes implements Sizer
func (ibs *imageButtonSizer) Sizes(hint Size) (min, pref, max Size) {
	pref = ibs.size
	pref.Width += ibs.button.Theme.HorizontalMargin*2 + 2
	pref.Height += ibs.button.Theme.VerticalMargin*2 + 2
	if border := ibs.button.Border(); border != nil {
		pref.AddInsets(border.Insets())
	}
	return pref, pref, pref
}
