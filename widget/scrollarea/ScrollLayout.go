// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package scrollarea

import (
	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/layout"
)

type scrollLayout struct {
	sa *ScrollArea
}

func newScrollLayout(sa *ScrollArea) *scrollLayout {
	layout := &scrollLayout{sa: sa}
	sa.SetLayout(layout)
	return layout
}

// Sizes implements the Layout interface.
func (sl *scrollLayout) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	_, hBarSize, _ := ui.Sizes(sl.sa.hBar, layout.NoHintSize)
	_, vBarSize, _ := ui.Sizes(sl.sa.vBar, layout.NoHintSize)
	min.Width = vBarSize.Width * 2
	min.Height = hBarSize.Height * 2
	if sl.sa.content != nil {
		_, pref, _ = ui.Sizes(sl.sa.content, hint)
	}
	if border := sl.sa.Border(); border != nil {
		insets := border.Insets()
		min.AddInsets(insets)
		pref.AddInsets(insets)
		max.AddInsets(insets)
	}
	return min, pref, layout.DefaultMaxSize(pref)
}

// Layout implements the Layout interface.
func (sl *scrollLayout) Layout() {
	_, hBarSize, _ := ui.Sizes(sl.sa.hBar, layout.NoHintSize)
	_, vBarSize, _ := ui.Sizes(sl.sa.vBar, layout.NoHintSize)
	needHBar := false
	needVBar := false
	var insets geom.Insets
	if border := sl.sa.Border(); border != nil {
		insets = border.Insets()
	}
	bounds := sl.sa.LocalInsetBounds()
	visibleSize := bounds.Size
	var contentSize geom.Size
	var prefContentSize geom.Size
	if sl.sa.content != nil {
		_, prefContentSize, _ = ui.Sizes(sl.sa.content, layout.NoHintSize)
		contentSize = prefContentSize
		switch sl.sa.behavior {
		case FillWidth:
			if visibleSize.Width > contentSize.Width {
				contentSize.Width = visibleSize.Width
			}
		case FillHeight:
			if visibleSize.Height > contentSize.Height {
				contentSize.Height = visibleSize.Height
			}
		case Fill:
			if visibleSize.Width > contentSize.Width {
				contentSize.Width = visibleSize.Width
			}
			if visibleSize.Height > contentSize.Height {
				contentSize.Height = visibleSize.Height
			}
		default:
		}
	}
	if visibleSize.Width < contentSize.Width {
		visibleSize.Height -= hBarSize.Height
		if insets.Bottom >= 1 {
			visibleSize.Height++
		}
		if sl.sa.behavior == FillHeight || sl.sa.behavior == Fill {
			if visibleSize.Height > prefContentSize.Height {
				contentSize.Height = visibleSize.Height
			}
		}
		needHBar = true
	}
	if visibleSize.Height < contentSize.Height {
		visibleSize.Width -= vBarSize.Width
		if insets.Right >= 1 {
			visibleSize.Width++
		}
		if sl.sa.behavior == FillWidth || sl.sa.behavior == Fill {
			if visibleSize.Width > prefContentSize.Width {
				contentSize.Width = visibleSize.Width
			}
		}
		needVBar = true
	}
	if !needHBar && visibleSize.Width < contentSize.Width {
		visibleSize.Height -= hBarSize.Height
		if insets.Bottom >= 1 {
			visibleSize.Height++
		}
		if sl.sa.behavior == FillHeight || sl.sa.behavior == Fill {
			if visibleSize.Height > prefContentSize.Height {
				contentSize.Height = visibleSize.Height
			}
		}
		needHBar = true
	}
	if needHBar {
		if sl.sa.hBar.Parent() == nil {
			sl.sa.AddChild(&sl.sa.hBar.Block)
		}
	} else {
		sl.sa.hBar.RemoveFromParent()
	}
	if needVBar {
		if sl.sa.vBar.Parent() == nil {
			sl.sa.AddChild(&sl.sa.vBar.Block)
		}
	} else {
		sl.sa.vBar.RemoveFromParent()
	}
	sl.sa.view.SetBounds(geom.Rect{Point: bounds.Point, Size: visibleSize})
	if needHBar {
		hBarSize.Width = visibleSize.Width
		barBounds := geom.Rect{Point: geom.Point{X: bounds.X, Y: bounds.Y + visibleSize.Height}, Size: hBarSize}
		if insets.Left >= 1 {
			barBounds.X--
			barBounds.Width++
		}
		if insets.Right >= 1 {
			barBounds.Width++
		}
		sl.sa.hBar.SetBounds(barBounds)
	}
	if needVBar {
		vBarSize.Height = visibleSize.Height
		barBounds := geom.Rect{Point: geom.Point{X: bounds.X + visibleSize.Width, Y: bounds.Y}, Size: vBarSize}
		if insets.Top >= 1 {
			barBounds.Y--
			barBounds.Height++
		}
		if insets.Bottom >= 1 {
			barBounds.Height++
		}
		sl.sa.vBar.SetBounds(barBounds)
	}
	if sl.sa.content != nil {
		sl.sa.content.SetSize(contentSize)
	}
}
