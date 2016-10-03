// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package menu

import (
	"fmt"
	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui/border"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/layout/flex"
	"github.com/richardwilkes/ui/widget"
)

type MenuBar struct {
	widget.Block
}

func NewMenuBar() *MenuBar {
	bar := &MenuBar{}
	bar.Describer = func() string { return fmt.Sprintf("MenuBar #%d", bar.ID()) }
	bar.SetBorder(border.NewLine(color.Background.AdjustBrightness(-0.25), geom.Insets{Top: 0, Left: 0, Bottom: 1, Right: 0}))
	flex.NewLayout(bar)
	return bar
}

func (bar *MenuBar) AddMenu(menu *Menu) {
	bar.AddChild(menu.item)
	menu.attachToBottom = true
	switch layout := bar.Layout().(type) {
	case *flex.Flex:
		layout.SetColumns(len(bar.Children()))
	}
}
