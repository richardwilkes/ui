// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package layout

import (
	"github.com/richardwilkes/go-ui/geom"
)

// PrecisionData is used to control how an object is laid out by the Precision layout.
type PrecisionData struct {
	HorizontalSpan      int
	VerticalSpan        int
	HorizontalAlignment Alignment
	VerticalAlignment   Alignment
	SizeHint            geom.Size
	MinSize             geom.Size
	cacheSize           geom.Size
	cacheMinWidth       float32
	HorizontalGrab      bool
	VerticalGrab        bool
}

// NewPrecisionData creates a new PrecisionData.
func NewPrecisionData() *PrecisionData {
	return &PrecisionData{HorizontalAlignment: Beginning, VerticalAlignment: Middle, SizeHint: NoHintSize, MinSize: NoHintSize}
}

// SetHorizontalAlignment is a convenience for setting the horizontal alignment while chaining
// calls together.
func (pd *PrecisionData) SetHorizontalAlignment(alignment Alignment) *PrecisionData {
	pd.HorizontalAlignment = alignment
	return pd
}

// SetVerticalAlignment is a convenience for setting the vertical alignment while chaining calls
// together.
func (pd *PrecisionData) SetVerticalAlignment(alignment Alignment) *PrecisionData {
	pd.VerticalAlignment = alignment
	return pd
}

// SetSizeHint is a convenience for setting the size hint while chaining calls together.
func (pd *PrecisionData) SetSizeHint(size geom.Size) *PrecisionData {
	pd.SizeHint = size
	return pd
}

// SetWidthHint is a convenience for setting the width hint while chaining calls together.
func (pd *PrecisionData) SetWidthHint(width float32) *PrecisionData {
	pd.SizeHint.Width = width
	return pd
}

// SetHeightHint is a convenience for setting the height hint while chaining calls together.
func (pd *PrecisionData) SetHeightHint(height float32) *PrecisionData {
	pd.SizeHint.Height = height
	return pd
}

// SetHorizontalSpan is a convenience for setting the horizontal span while chaining calls
// together.
func (pd *PrecisionData) SetHorizontalSpan(span int) *PrecisionData {
	pd.HorizontalSpan = span
	return pd
}

// SetVerticalSpan is a convenience for setting the vertical span while chaining calls together.
func (pd *PrecisionData) SetVerticalSpan(span int) *PrecisionData {
	pd.VerticalSpan = span
	return pd
}

// SetMinSize is a convenience for setting the minimum size while chaining calls together.
func (pd *PrecisionData) SetMinSize(size geom.Size) *PrecisionData {
	pd.MinSize = size
	return pd
}

// SetMinWidth is a convenience for setting the minimum width while chaining calls together.
func (pd *PrecisionData) SetMinWidth(width float32) *PrecisionData {
	pd.MinSize.Width = width
	return pd
}

// SetMinHeight is a convenience for setting the minimum height while chaining calls together.
func (pd *PrecisionData) SetMinHeight(height float32) *PrecisionData {
	pd.MinSize.Height = height
	return pd
}

// SetHorizontalGrab is a convenience for setting the horizontal grab while chaining calls
// together.
func (pd *PrecisionData) SetHorizontalGrab(grab bool) *PrecisionData {
	pd.HorizontalGrab = grab
	return pd
}

// SetVerticalGrab is a convenience for setting the vertical grab while chaining calls together.
func (pd *PrecisionData) SetVerticalGrab(grab bool) *PrecisionData {
	pd.VerticalGrab = grab
	return pd
}

func (pd *PrecisionData) computeCacheSize(target Layoutable, hint geom.Size, useMinimumSize bool) {
	pd.cacheMinWidth = 0
	pd.cacheSize.Width = 0
	pd.cacheSize.Height = 0
	min, pref, max := target.ComputeSizes(hint)
	if hint.Width != NoHint || hint.Height != NoHint {
		if pd.MinSize.Width != NoHint {
			pd.cacheMinWidth = pd.MinSize.Width
		} else {
			pd.cacheMinWidth = min.Width
		}
		if hint.Width != NoHint && hint.Width < pd.cacheMinWidth {
			hint.Width = pd.cacheMinWidth
		}
		if hint.Width != NoHint && hint.Width > max.Width {
			hint.Width = max.Width
		}
		if hint.Height != NoHint {
			var value float32
			if pd.MinSize.Height == NoHint {
				value = min.Height
			} else {
				value = pd.MinSize.Height
			}
			if hint.Height < value {
				hint.Height = value
			}
		}
		if hint.Height != NoHint && hint.Height > max.Height {
			hint.Height = max.Height
		}
	}
	if useMinimumSize {
		pd.cacheSize = min
		if pd.MinSize.Width != NoHint {
			pd.cacheMinWidth = pd.MinSize.Width
		} else {
			pd.cacheMinWidth = min.Width
		}
	} else {
		pd.cacheSize = pref
	}
	if hint.Width != NoHint {
		pd.cacheSize.Width = hint.Width
	} else {
		if pd.SizeHint.Width != NoHint {
			pd.cacheSize.Width = pd.SizeHint.Width
		}
		if pd.MinSize.Width != NoHint && pd.cacheSize.Width < pd.MinSize.Width {
			pd.cacheSize.Width = pd.MinSize.Width
		}
	}
	if hint.Height != NoHint {
		pd.cacheSize.Height = hint.Height
	} else {
		if pd.SizeHint.Height != NoHint {
			pd.cacheSize.Height = pd.SizeHint.Height
		}
		if pd.MinSize.Height != NoHint && pd.cacheSize.Height < pd.MinSize.Height {
			pd.cacheSize.Height = pd.MinSize.Height
		}
	}
}
