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
	"github.com/richardwilkes/ui/menu"
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

// AddMenu appends a menu to the end of this bar.
func (bar *MenuBar) AddMenu(subMenu menu.Menu) {
	if actual, ok := subMenu.(*Menu); ok {
		bar.AddChild(actual.item)
		actual.attachToBottom = true
		switch layout := bar.Layout().(type) {
		case *flex.Flex:
			layout.SetColumns(len(bar.Children()))
		}
	}
}

// InsertMenu inserts a menu at the specified menu index within this bar.
func (bar *MenuBar) InsertMenu(index int, subMenu menu.Menu) {
	if actual, ok := subMenu.(*Menu); ok {
		bar.AddChildAtIndex(actual.item, index)
		actual.attachToBottom = true
		switch layout := bar.Layout().(type) {
		case *flex.Flex:
			layout.SetColumns(len(bar.Children()))
		}
	}
}

// Remove the menu at the specified index from this bar.
func (bar *MenuBar) Remove(index int) {
	bar.RemoveChildAtIndex(index)
}

// Count of menus in this bar.
func (bar *MenuBar) Count() int {
	return len(bar.Children())
}

// Menu at the specified index, or nil.
func (bar *MenuBar) Menu(index int) menu.Menu {
	switch item := bar.Children()[index].(type) {
	case *MenuItem:
		return item.SubMenu()
	}
	panic("Invalid child")
}
