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
	"github.com/richardwilkes/go-ui/color"
	"github.com/richardwilkes/go-ui/geom"
	"github.com/richardwilkes/go-ui/graphics"
	"github.com/richardwilkes/go-ui/image"
)

// ImageButton represents a clickable image button.
type ImageButton struct {
	Block
	Theme         *ImageButtonTheme // The theme the button will use to draw itself.
	OnClick       func()            // Called when the button is clicked.
	image         *image.Image
	disabledImage *image.Image
	pressed       bool
}

// NewImageButton creates a new button with the specified Image.
func NewImageButton(img *image.Image) *ImageButton {
	button := &ImageButton{}
	button.Init(img)
	return button
}

// NewImageButtonWithImageSize creates a new button with the specified Image. The image will be set
// to the specified size. The button itself will be a bit larger, based on the theme settings and
// border.
func NewImageButtonWithImageSize(img *image.Image, size geom.Size) *ImageButton {
	button := &ImageButton{}
	button.Init(img)
	button.Sizes = func(hint geom.Size) (min, pref, max geom.Size) {
		pref = size
		pref.Width += button.Theme.HorizontalMargin*2 + 2
		pref.Height += button.Theme.VerticalMargin*2 + 2
		insets := button.Insets()
		pref.AddInsets(insets)
		return pref, pref, pref
	}
	return button
}

// Init initializes the button.
func (button *ImageButton) Init(img *image.Image) {
	button.Block.Init()
	button.image = img
	var err error
	if button.disabledImage, err = img.AcquireDisabled(); err != nil {
		button.disabledImage = img
	}
	button.Theme = StdImageButtonTheme
	button.Sizes = func(hint geom.Size) (min, pref, max geom.Size) {
		size := button.image.Size()
		size.Width += button.Theme.HorizontalMargin*2 + 2
		size.Height += button.Theme.VerticalMargin*2 + 2
		insets := button.Insets()
		size.AddInsets(insets)
		return size, size, size
	}
	button.OnMouseDown = func(where geom.Point, keyModifiers int, which int, clickCount int) {
		button.pressed = true
		button.Repaint()
	}
	button.OnMouseDragged = func(where geom.Point, keyModifiers int) {
		bounds := button.LocalInsetBounds()
		pressed := bounds.Contains(where)
		if button.pressed != pressed {
			button.pressed = pressed
			button.Repaint()
		}
	}
	button.OnMouseUp = func(where geom.Point, keyModifiers int) {
		button.pressed = false
		button.Repaint()
		if button.OnClick != nil {
			bounds := button.LocalInsetBounds()
			if bounds.Contains(where) {
				button.OnClick()
			}
		}
	}
	button.OnPaint = func(gc graphics.Context, dirty geom.Rect, inLiveResize bool) {
		var hSpace = button.Theme.HorizontalMargin*2 + 2
		var vSpace = button.Theme.VerticalMargin*2 + 2
		bounds := button.LocalInsetBounds()
		path := graphics.NewPath()
		path.MoveTo(bounds.X, bounds.Y+button.Theme.CornerRadius)
		path.QuadCurveTo(bounds.X, bounds.Y, bounds.X+button.Theme.CornerRadius, bounds.Y)
		path.LineTo(bounds.X+bounds.Width-button.Theme.CornerRadius, bounds.Y)
		path.QuadCurveTo(bounds.X+bounds.Width, bounds.Y, bounds.X+bounds.Width, bounds.Y+button.Theme.CornerRadius)
		path.LineTo(bounds.X+bounds.Width, bounds.Y+bounds.Height-button.Theme.CornerRadius)
		path.QuadCurveTo(bounds.X+bounds.Width, bounds.Y+bounds.Height, bounds.X+bounds.Width-button.Theme.CornerRadius, bounds.Y+bounds.Height)
		path.LineTo(bounds.X+button.Theme.CornerRadius, bounds.Y+bounds.Height)
		path.QuadCurveTo(bounds.X, bounds.Y+bounds.Height, bounds.X, bounds.Y+bounds.Height-button.Theme.CornerRadius)
		path.ClosePath()
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
}

// Image returns this button's base image.
func (button *ImageButton) Image() *image.Image {
	return button.image
}

// CurrentImage returns this button's current image.
func (button *ImageButton) CurrentImage() *image.Image {
	if button.Enabled() {
		return button.image
	}
	return button.disabledImage
}

// BaseBackground returns this button's current base background color.
func (button *ImageButton) BaseBackground() color.Color {
	switch {
	case button.Disabled():
		return button.Theme.Background.AdjustBrightness(button.Theme.DisabledAdjustment)
	case button.pressed:
		return button.Theme.BackgroundWhenPressed
	case button.Focused():
		return button.Theme.Background.Blend(color.KeyboardFocus, 0.5)
	default:
		return button.Theme.Background
	}
}
