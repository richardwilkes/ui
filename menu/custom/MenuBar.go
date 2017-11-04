// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package custom

import (
	"fmt"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui/border"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/layout/flex"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/widget"
	"github.com/richardwilkes/ui/window"
)

// MenuBar represents a set of menus.
type MenuBar struct {
	widget.Block
	special map[menu.SpecialMenuType]menu.Menu
}

var (
	lookingUpBar bool
)

// AppBar returns the menu bar for the given window.
func AppBar(id uint64) menu.Bar {
	if lookingUpBar {
		return nil
	}
	wnd := window.ByID(id)
	lookingUpBar = true
	bar := wnd.MenuBar()
	lookingUpBar = false
	if bar == nil {
		mb := &MenuBar{special: make(map[menu.SpecialMenuType]menu.Menu)}
		mb.InitTypeAndID(mb)
		mb.Describer = func() string { return fmt.Sprintf("MenuBar #%d", mb.ID()) }
		mb.SetBorder(border.NewLine(color.Background.AdjustBrightness(-0.25), geom.Insets{Top: 0, Left: 0, Bottom: 1, Right: 0}))
		flex.NewLayout(mb)
		bar = mb
	}
	return bar
}

// AppendMenu appends a menu at the end of this bar.
func (bar *MenuBar) AppendMenu(subMenu menu.Menu) {
	bar.InsertMenu(subMenu, -1)
}

// InsertMenu inserts a menu at the specified item index within this bar. Pass in a negative
// index to append to the end.
func (bar *MenuBar) InsertMenu(subMenu menu.Menu, index int) {
	if actual, ok := subMenu.(*Menu); ok {
		bar.AddChildAtIndex(actual.item, index)
		actual.menuParent = bar
		actual.item.menuParent = bar
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

// SpecialMenu returns the specified special menu, or nil if it has not been setup.
func (bar *MenuBar) SpecialMenu(which menu.SpecialMenuType) menu.Menu {
	return bar.special[which]
}

// SetupSpecialMenu sets up the specified special menu, which must have already been installed
// into the menu bar.
func (bar *MenuBar) SetupSpecialMenu(which menu.SpecialMenuType, mnu menu.Menu) {
	bar.special[which] = mnu
}

// ProcessKeyDown is called to process KeyDown events prior to anything else receiving them.
func (bar *MenuBar) ProcessKeyDown(evt *event.KeyDown) {
	for _, child := range bar.Children() {
		switch item := child.(type) {
		case *MenuItem:
			if item.processKeyDown(evt) {
				break
			}
		}
	}
}

func (bar *MenuBar) findRoot() menuRoot {
	return bar
}

func (bar *MenuBar) findLeafOpenMenu() *Menu {
	for _, child := range bar.Children() {
		if item, ok := child.(*MenuItem); ok && item.menu != nil && item.menuOpen {
			return item.menu.findLeafOpenMenu()
		}
	}
	return nil
}
