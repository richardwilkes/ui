// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

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
	if size.Width <= 0 || size.Height <= 0 {
		button.SetSizer(button)
	} else {
		button.SetSizer(&imageButtonSizer{button: button, size: size})
	}
	button.SetPaintHandler(button)
	button.SetMouseDownHandler(button)
	button.SetMouseDraggedHandler(button)
	button.SetMouseUpHandler(button)
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

// OnPaint implements PaintHandler
func (button *ImageButton) OnPaint(g Graphics, dirty Rect) {
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
		g.DrawImageInRect(img, bounds)
	}
}

// OnMouseDown implements MouseDownHandler
func (button *ImageButton) OnMouseDown(where Point, keyModifiers int, which int, clickCount int) bool {
	button.pressed = true
	button.Repaint()
	return false
}

// OnMouseDragged implements MouseDraggedHandler
func (button *ImageButton) OnMouseDragged(where Point, keyModifiers int) {
	bounds := button.LocalInsetBounds()
	pressed := bounds.Contains(where)
	if button.pressed != pressed {
		button.pressed = pressed
		button.Repaint()
	}
}

// OnMouseUp implements MouseUpHandler
func (button *ImageButton) OnMouseUp(where Point, keyModifiers int) {
	button.pressed = false
	button.Repaint()
	if button.OnClick != nil {
		bounds := button.LocalInsetBounds()
		if bounds.Contains(where) {
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
