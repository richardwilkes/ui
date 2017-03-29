// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package window

import (
	"fmt"

	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/layout"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/widget"
)

type RootView struct {
	widget.Block
	tooltip ui.Widget
	menuBar menu.Bar
	content ui.Widget
}

func newRootView(window ui.Window) *RootView {
	view := &RootView{}
	view.SetBackground(color.Background)
	view.SetWindow(window)
	view.Describer = func() string { return fmt.Sprintf("RootView #%d", view.ID()) }
	view.SetLayout(&RootLayout{view: view})
	view.content = widget.NewBlock()
	view.AddChild(view.content)
	return view
}

func (view *RootView) MenuBar() menu.Bar {
	return view.menuBar
}

func (view *RootView) SetMenuBar(bar menu.Bar) {
	if view.menuBar != nil {
		if actual, ok := view.menuBar.(ui.Widget); ok {
			view.RemoveChild(actual)
		}
	}
	view.menuBar = bar
	if actual, ok := bar.(ui.Widget); ok {
		view.AddChildAtIndex(actual, 0)
	}
}

func (view *RootView) Tooltip() ui.Widget {
	return view.tooltip
}

func (view *RootView) SetTooltip(tip ui.Widget) {
	if view.tooltip != nil {
		view.tooltip.Repaint()
		view.RemoveChild(view.tooltip)
	}
	view.tooltip = tip
	if tip != nil {
		view.AddChild(tip)
		tip.Repaint()
	}
}

func (view *RootView) Content() ui.Widget {
	return view.content
}

type RootLayout struct {
	view *RootView
}

func (lay *RootLayout) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	min, pref, max = ui.Sizes(lay.view.content, hint)
	if lay.view.menuBar != nil {
		if actual, ok := lay.view.menuBar.(ui.Widget); ok {
			_, barSize, _ := ui.Sizes(actual, layout.NoHintSize)
			lay.adjustSizeForBarSize(&min, barSize)
			lay.adjustSizeForBarSize(&pref, barSize)
			lay.adjustSizeForBarSize(&max, barSize)
		}
	}
	return
}

func (lay *RootLayout) adjustSizeForBarSize(size *geom.Size, barSize geom.Size) {
	size.Height += barSize.Height
	if size.Width < barSize.Width {
		size.Width = barSize.Width
	}
}

func (lay *RootLayout) Layout() {
	bounds := lay.view.LocalBounds()
	if lay.view.menuBar != nil {
		if actual, ok := lay.view.menuBar.(ui.Widget); ok {
			_, size, _ := ui.Sizes(actual, layout.NoHintSize)
			actual.SetBounds(geom.Rect{Size: geom.Size{Width: bounds.Width, Height: size.Height}})
			bounds.Y += size.Height
			bounds.Height -= size.Height
		}
	}
	lay.view.content.SetBounds(bounds)
}
