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

type Alignment int

type PrecisionData struct {
	cacheMinWidth       float32
	cacheSize           geom.Size
	HorizontalAlignment Alignment
	VerticalAlignment   Alignment
	SizeHint            geom.Size
	HorizontalSpan      int
	VerticalSpan        int
	MinSize             geom.Size
	HorizontalGrab      bool
	VerticalGrab        bool
}

func NewPrecisionData() *PrecisionData {
	return &PrecisionData{HorizontalAlignment: Beginning, VerticalAlignment: Middle, SizeHint: NoHintSize, MinSize: NoHintSize}
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
