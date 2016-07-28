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
	"github.com/richardwilkes/ui/geom"
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
	_, hBarSize, _ := layout.Sizes(sl.sa.hBar, layout.NoHintSize)
	_, vBarSize, _ := layout.Sizes(sl.sa.vBar, layout.NoHintSize)
	min.Width = vBarSize.Width * 2
	min.Height = hBarSize.Height * 2
	if sl.sa.content != nil {
		_, pref, _ = layout.Sizes(sl.sa.content, hint)
	}
	if border := sl.sa.view.Border(); border != nil {
		insets := border.Insets()
		min.AddInsets(insets)
		pref.AddInsets(insets)
		max.AddInsets(insets)
	}
	return min, pref, layout.DefaultMaxSize(pref)
}

// Layout implements the Layout interface.
func (sl *scrollLayout) Layout() {
	_, hBarSize, _ := layout.Sizes(sl.sa.hBar, layout.NoHintSize)
	_, vBarSize, _ := layout.Sizes(sl.sa.vBar, layout.NoHintSize)
	needHBar := false
	needVBar := false
	var contentSize geom.Size
	if sl.sa.content != nil {
		contentSize = sl.sa.content.Size()
	}
	bounds := sl.sa.LocalInsetBounds()
	visibleSize := bounds.Size
	var viewInsets geom.Insets
	if border := sl.sa.view.Border(); border != nil {
		viewInsets = border.Insets()
	}
	visibleSize.SubtractInsets(viewInsets)
	if visibleSize.Width < contentSize.Width {
		visibleSize.Height -= hBarSize.Height
		needHBar = true
	}
	if visibleSize.Height < contentSize.Height {
		visibleSize.Width -= vBarSize.Width
		needVBar = true
	}
	if !needHBar && visibleSize.Width < contentSize.Width {
		visibleSize.Height -= hBarSize.Height
		needHBar = true
	}
	if needHBar {
		if sl.sa.hBar.parent == nil {
			sl.sa.AddChild(&sl.sa.hBar.Block)
		}
	} else {
		sl.sa.hBar.RemoveFromParent()
	}
	if needVBar {
		if sl.sa.vBar.parent == nil {
			sl.sa.AddChild(&sl.sa.vBar.Block)
		}
	} else {
		sl.sa.vBar.RemoveFromParent()
	}
	visibleSize.AddInsets(viewInsets)
	sl.sa.view.SetBounds(geom.Rect{Point: bounds.Point, Size: visibleSize})
	if needHBar {
		hBarSize.Width = visibleSize.Width
		sl.sa.hBar.SetBounds(geom.Rect{Point: geom.Point{X: bounds.X, Y: bounds.Y + visibleSize.Height}, Size: hBarSize})
	}
	if needVBar {
		vBarSize.Height = visibleSize.Height
		sl.sa.vBar.SetBounds(geom.Rect{Point: geom.Point{X: bounds.X + visibleSize.Width, Y: bounds.Y}, Size: vBarSize})
	}
}
