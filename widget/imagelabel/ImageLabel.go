// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package imagelabel

import (
	"fmt"
	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/widget"
)

// ImageLabel represents a non-interactive image.
type ImageLabel struct {
	widget.Block
	image *draw.Image
}

// New creates an ImageLabel with the specified image.
func New(img *draw.Image) *ImageLabel {
	return NewWithSize(img, geom.Size{})
}

// NewWithSize creates a new ImageLabel with the specified image. The image will be set to the
// specified size.
func NewWithSize(img *draw.Image, size geom.Size) *ImageLabel {
	label := &ImageLabel{image: img}
	label.InitTypeAndID(label)
	label.Describer = func() string { return fmt.Sprintf("ImageLabel #%d (%v)", label.ID(), label.image) }
	if size.Width <= 0 || size.Height <= 0 {
		label.SetSizer(label)
	} else {
		label.SetSizer(&imageLabelSizer{label: label, size: size})
	}
	label.EventHandlers().Add(event.PaintType, label.paint)
	return label
}

// Sizes implements Sizer
func (label *ImageLabel) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	size := label.image.Size()
	if border := label.Border(); border != nil {
		size.AddInsets(border.Insets())
	}
	return size, size, size
}

func (label *ImageLabel) paint(evt event.Event) {
	bounds := label.LocalInsetBounds()
	size := label.image.Size()
	if size.Width < bounds.Width {
		bounds.X += (bounds.Width - size.Width) / 2
		bounds.Width = size.Width
	}
	if size.Height < bounds.Height {
		bounds.Y += (bounds.Height - size.Height) / 2
		bounds.Height = size.Height
	}
	evt.(*event.Paint).GC().DrawImageInRect(label.image, bounds)
}

type imageLabelSizer struct {
	label *ImageLabel
	size  geom.Size
}

// Sizes implements Sizer
func (sizer *imageLabelSizer) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	pref = sizer.size
	if border := sizer.label.Border(); border != nil {
		pref.AddInsets(border.Insets())
	}
	return pref, pref, pref
}
