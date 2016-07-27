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
	"github.com/richardwilkes/ui/draw"
)

// PrecisionData is used to control how an object is laid out by the Precision layout.
type PrecisionData struct {
	hSpan        int
	vSpan        int
	hAlign       draw.Alignment
	vAlign       draw.Alignment
	sizeHint     draw.Size
	minSize      draw.Size
	cacheSize    draw.Size
	minCacheSize draw.Size
	hGrab        bool
	vGrab        bool
}

// NewPrecisionData creates a new PrecisionData.
func NewPrecisionData() *PrecisionData {
	return &PrecisionData{hSpan: 1, vSpan: 1, hAlign: draw.AlignStart, vAlign: draw.AlignMiddle, sizeHint: NoLayoutHintSize, minSize: NoLayoutHintSize}
}

// HorizontalAlignment returns the horizontal alignment of the widget within its space.
func (pd *PrecisionData) HorizontalAlignment() draw.Alignment {
	return pd.hAlign
}

// SetHorizontalAlignment sets the horizontal alignment of the widget within its space.
func (pd *PrecisionData) SetHorizontalAlignment(alignment draw.Alignment) *PrecisionData {
	pd.hAlign = alignment
	return pd
}

// VerticalAlignment returns the vertical alignment of the widget within its space.
func (pd *PrecisionData) VerticalAlignment() draw.Alignment {
	return pd.vAlign
}

// SetVerticalAlignment sets the vertical alignment of the widget within its space.
func (pd *PrecisionData) SetVerticalAlignment(alignment draw.Alignment) *PrecisionData {
	pd.vAlign = alignment
	return pd
}

// SizeHint returns a hint requesting a particular size of the widget.
func (pd *PrecisionData) SizeHint() draw.Size {
	return pd.sizeHint
}

// SetSizeHint sets a hint requesting a particular size of the widget.
func (pd *PrecisionData) SetSizeHint(size draw.Size) *PrecisionData {
	pd.sizeHint = size
	return pd
}

// SetWidthHint sets a hint requesting a particular width of the widget.
func (pd *PrecisionData) SetWidthHint(width float32) *PrecisionData {
	pd.sizeHint.Width = width
	return pd
}

// SetHeightHint sets a hint requesting a particular height of the widget.
func (pd *PrecisionData) SetHeightHint(height float32) *PrecisionData {
	pd.sizeHint.Height = height
	return pd
}

// HorizontalSpan returns the number of columns the widget should span.
func (pd *PrecisionData) HorizontalSpan() int {
	return pd.hSpan
}

// SetHorizontalSpan sets the number of columns the widget should span.
func (pd *PrecisionData) SetHorizontalSpan(span int) *PrecisionData {
	pd.hSpan = span
	return pd
}

// VerticalSpan returns the number of rows the widget should span.
func (pd *PrecisionData) VerticalSpan() int {
	return pd.vSpan
}

// SetVerticalSpan sets the number of rows the widget should span.
func (pd *PrecisionData) SetVerticalSpan(span int) *PrecisionData {
	pd.vSpan = span
	return pd
}

// MinSize returns an override for the minimum size of the widget.
func (pd *PrecisionData) MinSize() draw.Size {
	return pd.minSize
}

// SetMinSize sets an override for the minimum size of the widget.
func (pd *PrecisionData) SetMinSize(size draw.Size) *PrecisionData {
	pd.minSize = size
	return pd
}

// SetMinWidth sets an override for the minimum width of the widget.
func (pd *PrecisionData) SetMinWidth(width float32) *PrecisionData {
	pd.minSize.Width = width
	return pd
}

// SetMinHeight sets an override for the minimum height of the widget.
func (pd *PrecisionData) SetMinHeight(height float32) *PrecisionData {
	pd.minSize.Height = height
	return pd
}

// HorizontalGrab returns true if the widget should attempt to grab excess horizontal space.
func (pd *PrecisionData) HorizontalGrab() bool {
	return pd.hGrab
}

// SetHorizontalGrab marks the widget to attempt to grab excess horizontal space if true.
func (pd *PrecisionData) SetHorizontalGrab(grab bool) *PrecisionData {
	pd.hGrab = grab
	return pd
}

// VerticalGrab returns true if the widget should attempt to grab excess vertical space.
func (pd *PrecisionData) VerticalGrab() bool {
	return pd.vGrab
}

// SetVerticalGrab marks the widget to attempt to grab excess vertical space if true.
func (pd *PrecisionData) SetVerticalGrab(grab bool) *PrecisionData {
	pd.vGrab = grab
	return pd
}

func (pd *PrecisionData) computeCacheSize(target Widget, hint draw.Size, useMinimumSize bool) {
	pd.minCacheSize.Width = 0
	pd.minCacheSize.Height = 0
	pd.cacheSize.Width = 0
	pd.cacheSize.Height = 0
	min, pref, max := ComputeSizes(target, hint)
	if hint.Width != NoLayoutHint || hint.Height != NoLayoutHint {
		if pd.minSize.Width != NoLayoutHint {
			pd.minCacheSize.Width = pd.minSize.Width
		} else {
			pd.minCacheSize.Width = min.Width
		}
		if hint.Width != NoLayoutHint && hint.Width < pd.minCacheSize.Width {
			hint.Width = pd.minCacheSize.Width
		}
		if hint.Width != NoLayoutHint && hint.Width > max.Width {
			hint.Width = max.Width
		}

		if pd.minSize.Height != NoLayoutHint {
			pd.minCacheSize.Height = pd.minSize.Height
		} else {
			pd.minCacheSize.Height = min.Height
		}
		if hint.Height != NoLayoutHint && hint.Height < pd.minCacheSize.Height {
			hint.Height = pd.minCacheSize.Height
		}
		if hint.Height != NoLayoutHint && hint.Height > max.Height {
			hint.Height = max.Height
		}
	}
	if useMinimumSize {
		pd.cacheSize = min
		if pd.minSize.Width != NoLayoutHint {
			pd.minCacheSize.Width = pd.minSize.Width
		} else {
			pd.minCacheSize.Width = min.Width
		}
		if pd.minSize.Height != NoLayoutHint {
			pd.minCacheSize.Height = pd.minSize.Height
		} else {
			pd.minCacheSize.Height = min.Height
		}
	} else {
		pd.cacheSize = pref
	}
	if hint.Width != NoLayoutHint {
		pd.cacheSize.Width = hint.Width
	}
	if pd.minSize.Width != NoLayoutHint && pd.cacheSize.Width < pd.minSize.Width {
		pd.cacheSize.Width = pd.minSize.Width
	}
	if pd.sizeHint.Width != NoLayoutHint {
		pd.cacheSize.Width = pd.sizeHint.Width
	}
	if hint.Height != NoLayoutHint {
		pd.cacheSize.Height = hint.Height
	}
	if pd.minSize.Height != NoLayoutHint && pd.cacheSize.Height < pd.minSize.Height {
		pd.cacheSize.Height = pd.minSize.Height
	}
	if pd.sizeHint.Height != NoLayoutHint {
		pd.cacheSize.Height = pd.sizeHint.Height
	}
}
