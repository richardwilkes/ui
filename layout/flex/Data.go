// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package flex

import (
	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/layout"
)

// Data is used to control how an object is laid out by the Flex layout.
type Data struct {
	hSpan        int
	vSpan        int
	sizeHint     geom.Size
	minSize      geom.Size
	cacheSize    geom.Size
	minCacheSize geom.Size
	hAlign       draw.Alignment
	vAlign       draw.Alignment
	hGrab        bool
	vGrab        bool
}

// NewData creates a new Data.
func NewData() *Data {
	return &Data{hSpan: 1, vSpan: 1, hAlign: draw.AlignStart, vAlign: draw.AlignMiddle, sizeHint: layout.NoHintSize, minSize: layout.NoHintSize}
}

// HorizontalAlignment returns the horizontal alignment of the widget within its space.
func (data *Data) HorizontalAlignment() draw.Alignment {
	return data.hAlign
}

// SetHorizontalAlignment sets the horizontal alignment of the widget within its space.
func (data *Data) SetHorizontalAlignment(alignment draw.Alignment) *Data {
	data.hAlign = alignment
	return data
}

// VerticalAlignment returns the vertical alignment of the widget within its space.
func (data *Data) VerticalAlignment() draw.Alignment {
	return data.vAlign
}

// SetVerticalAlignment sets the vertical alignment of the widget within its space.
func (data *Data) SetVerticalAlignment(alignment draw.Alignment) *Data {
	data.vAlign = alignment
	return data
}

// SizeHint returns a hint requesting a particular size of the widget.
func (data *Data) SizeHint() geom.Size {
	return data.sizeHint
}

// SetSizeHint sets a hint requesting a particular size of the widget.
func (data *Data) SetSizeHint(size geom.Size) *Data {
	data.sizeHint = size
	return data
}

// SetWidthHint sets a hint requesting a particular width of the widget.
func (data *Data) SetWidthHint(width float64) *Data {
	data.sizeHint.Width = width
	return data
}

// SetHeightHint sets a hint requesting a particular height of the widget.
func (data *Data) SetHeightHint(height float64) *Data {
	data.sizeHint.Height = height
	return data
}

// HorizontalSpan returns the number of columns the widget should span.
func (data *Data) HorizontalSpan() int {
	return data.hSpan
}

// SetHorizontalSpan sets the number of columns the widget should span.
func (data *Data) SetHorizontalSpan(span int) *Data {
	data.hSpan = span
	return data
}

// VerticalSpan returns the number of rows the widget should span.
func (data *Data) VerticalSpan() int {
	return data.vSpan
}

// SetVerticalSpan sets the number of rows the widget should span.
func (data *Data) SetVerticalSpan(span int) *Data {
	data.vSpan = span
	return data
}

// MinSize returns an override for the minimum size of the widget.
func (data *Data) MinSize() geom.Size {
	return data.minSize
}

// SetMinSize sets an override for the minimum size of the widget.
func (data *Data) SetMinSize(size geom.Size) *Data {
	data.minSize = size
	return data
}

// SetMinWidth sets an override for the minimum width of the widget.
func (data *Data) SetMinWidth(width float64) *Data {
	data.minSize.Width = width
	return data
}

// SetMinHeight sets an override for the minimum height of the widget.
func (data *Data) SetMinHeight(height float64) *Data {
	data.minSize.Height = height
	return data
}

// HorizontalGrab returns true if the widget should attempt to grab excess horizontal space.
func (data *Data) HorizontalGrab() bool {
	return data.hGrab
}

// SetHorizontalGrab marks the widget to attempt to grab excess horizontal space if true.
func (data *Data) SetHorizontalGrab(grab bool) *Data {
	data.hGrab = grab
	return data
}

// VerticalGrab returns true if the widget should attempt to grab excess vertical space.
func (data *Data) VerticalGrab() bool {
	return data.vGrab
}

// SetVerticalGrab marks the widget to attempt to grab excess vertical space if true.
func (data *Data) SetVerticalGrab(grab bool) *Data {
	data.vGrab = grab
	return data
}

func (data *Data) computeCacheSize(target ui.Widget, hint geom.Size, useMinimumSize bool) {
	data.minCacheSize.Width = 0
	data.minCacheSize.Height = 0
	data.cacheSize.Width = 0
	data.cacheSize.Height = 0
	min, pref, max := ui.Sizes(target, hint)
	if hint.Width != layout.NoHint || hint.Height != layout.NoHint {
		if data.minSize.Width != layout.NoHint {
			data.minCacheSize.Width = data.minSize.Width
		} else {
			data.minCacheSize.Width = min.Width
		}
		if hint.Width != layout.NoHint && hint.Width < data.minCacheSize.Width {
			hint.Width = data.minCacheSize.Width
		}
		if hint.Width != layout.NoHint && hint.Width > max.Width {
			hint.Width = max.Width
		}

		if data.minSize.Height != layout.NoHint {
			data.minCacheSize.Height = data.minSize.Height
		} else {
			data.minCacheSize.Height = min.Height
		}
		if hint.Height != layout.NoHint && hint.Height < data.minCacheSize.Height {
			hint.Height = data.minCacheSize.Height
		}
		if hint.Height != layout.NoHint && hint.Height > max.Height {
			hint.Height = max.Height
		}
	}
	if useMinimumSize {
		data.cacheSize = min
		if data.minSize.Width != layout.NoHint {
			data.minCacheSize.Width = data.minSize.Width
		} else {
			data.minCacheSize.Width = min.Width
		}
		if data.minSize.Height != layout.NoHint {
			data.minCacheSize.Height = data.minSize.Height
		} else {
			data.minCacheSize.Height = min.Height
		}
	} else {
		data.cacheSize = pref
	}
	if hint.Width != layout.NoHint {
		data.cacheSize.Width = hint.Width
	}
	if data.minSize.Width != layout.NoHint && data.cacheSize.Width < data.minSize.Width {
		data.cacheSize.Width = data.minSize.Width
	}
	if data.sizeHint.Width != layout.NoHint {
		data.cacheSize.Width = data.sizeHint.Width
	}
	if hint.Height != layout.NoHint {
		data.cacheSize.Height = hint.Height
	}
	if data.minSize.Height != layout.NoHint && data.cacheSize.Height < data.minSize.Height {
		data.cacheSize.Height = data.minSize.Height
	}
	if data.sizeHint.Height != layout.NoHint {
		data.cacheSize.Height = data.sizeHint.Height
	}
}
