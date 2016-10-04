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
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/layout/flex"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/widget"
)

type RootView struct {
	widget.Block
}

func newRootView(window ui.Window) *RootView {
	view := &RootView{}
	view.SetBackground(color.Background)
	view.SetWindow(window)
	view.Describer = func() string { return fmt.Sprintf("RootView #%d", view.ID()) }
	flex.NewLayout(view)
	content := widget.NewBlock()
	content.SetLayoutData(flex.NewData().SetHorizontalAlignment(draw.AlignFill).SetVerticalAlignment(draw.AlignFill).SetHorizontalGrab(true).SetVerticalGrab(true))
	view.AddChild(content)
	return view
}

func (view *RootView) MenuBar() menu.Bar {
	children := view.Children()
	if len(children) > 1 {
		return children[0].(menu.Bar)
	}
	return nil
}

func (view *RootView) SetMenuBar(bar menu.Bar) {
	if actual, ok := bar.(ui.Widget); ok {
		view.AddChildAtIndex(actual, 0)
		actual.SetLayoutData(flex.NewData().SetHorizontalAlignment(draw.AlignFill).SetHorizontalGrab(true))
	}
}

func (view *RootView) Content() ui.Widget {
	children := view.Children()
	return children[len(children)-1]
}
