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

type scrollLayout struct {
	sa *ScrollArea
}

func newScrollLayout(sa *ScrollArea) *scrollLayout {
	layout := &scrollLayout{sa: sa}
	sa.SetLayout(layout)
	return layout
}

// Sizes implements the Layout interface.
func (layout *scrollLayout) Sizes(hint draw.Size) (min, pref, max draw.Size) {
	_, hBarSize, _ := ComputeSizes(layout.sa.hBar, NoLayoutHintSize)
	_, vBarSize, _ := ComputeSizes(layout.sa.vBar, NoLayoutHintSize)
	min.Width = vBarSize.Width * 2
	min.Height = hBarSize.Height * 2
	if layout.sa.content != nil {
		_, pref, _ = ComputeSizes(layout.sa.content, hint)
	}
	if border := layout.sa.view.Border(); border != nil {
		insets := border.Insets()
		min.AddInsets(insets)
		pref.AddInsets(insets)
		max.AddInsets(insets)
	}
	return min, pref, DefaultLayoutMaxSize(pref)
}

// Layout implements the Layout interface.
func (layout *scrollLayout) Layout() {
	_, hBarSize, _ := ComputeSizes(layout.sa.hBar, NoLayoutHintSize)
	_, vBarSize, _ := ComputeSizes(layout.sa.vBar, NoLayoutHintSize)
	needHBar := false
	needVBar := false
	var contentSize draw.Size
	if layout.sa.content != nil {
		contentSize = layout.sa.content.Size()
	}
	bounds := layout.sa.LocalInsetBounds()
	visibleSize := bounds.Size
	var viewInsets draw.Insets
	if border := layout.sa.view.Border(); border != nil {
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
		if layout.sa.hBar.parent == nil {
			layout.sa.AddChild(&layout.sa.hBar.Block)
		}
	} else {
		layout.sa.hBar.RemoveFromParent()
	}
	if needVBar {
		if layout.sa.vBar.parent == nil {
			layout.sa.AddChild(&layout.sa.vBar.Block)
		}
	} else {
		layout.sa.vBar.RemoveFromParent()
	}
	visibleSize.AddInsets(viewInsets)
	layout.sa.view.SetBounds(draw.Rect{Point: bounds.Point, Size: visibleSize})
	if needHBar {
		hBarSize.Width = visibleSize.Width
		layout.sa.hBar.SetBounds(draw.Rect{Point: draw.Point{X: bounds.X, Y: bounds.Y + visibleSize.Height}, Size: hBarSize})
	}
	if needVBar {
		vBarSize.Height = visibleSize.Height
		layout.sa.vBar.SetBounds(draw.Rect{Point: draw.Point{X: bounds.X + visibleSize.Width, Y: bounds.Y}, Size: vBarSize})
	}
}
