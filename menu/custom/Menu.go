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
	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/border"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/layout"
	"github.com/richardwilkes/ui/layout/flex"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/widget"
	"github.com/richardwilkes/ui/window"
)

type Menu struct {
	widget.Block
	item           *MenuItem
	wnd            ui.Window
	attachToBottom bool
}

func NewMenu(title string) *Menu {
	mnu := &Menu{item: NewItem(title, nil)}
	mnu.item.menu = mnu
	mnu.Describer = func() string {
		return fmt.Sprintf("Menu #%d (%s)", mnu.ID(), mnu.Title())
	}
	mnu.SetBorder(border.NewLine(color.Gray, geom.Insets{Top: 1, Left: 1, Bottom: 1, Right: 1}))
	mnu.item.EventHandlers().Add(event.SelectionType, mnu.open)
	flex.NewLayout(mnu).SetEqualColumns(true)
	return mnu
}

// Title returns the title of this menu.
func (mnu *Menu) Title() string {
	return mnu.item.Title()
}

// AppendItem appends an item at the end of this menu.
func (mnu *Menu) AppendItem(item menu.Item) {
	mnu.InsertItem(item, -1)
}

// InsertItem inserts an item at the specified item index within this menu. Pass in a negative
// index to append to the end.
func (mnu *Menu) InsertItem(item menu.Item, index int) {
	if actual, ok := item.(ui.Widget); ok {
		mnu.AddChildAtIndex(actual, index)
		actual.EventHandlers().Add(event.ClosingType, mnu.close)
		actual.SetLayoutData(flex.NewData().SetHorizontalGrab(true).SetHorizontalAlignment(draw.AlignFill))
	}
}

// AppendMenu appends an item with a sub-menu at the end of this menu.
func (mnu *Menu) AppendMenu(subMenu menu.Menu) {
	mnu.InsertMenu(subMenu, -1)
}

// InsertMenu inserts an item with a sub-menu at the specified item index within this menu. Pass
// in a negative index to append to the end.
func (mnu *Menu) InsertMenu(subMenu menu.Menu, index int) {
	if actual, ok := subMenu.(*Menu); ok {
		mnu.InsertItem(actual.item, index)
	}
}

// Remove the item at the specified index from this menu.
func (mnu *Menu) Remove(index int) {
	mnu.RemoveChildAtIndex(index)
}

// Count of items in this menu.
func (mnu *Menu) Count() int {
	return len(mnu.Children())
}

// Item at the specified index, or nil.
func (mnu *Menu) Item(index int) menu.Item {
	return mnu.Children()[index].(menu.Item)
}

func (mnu *Menu) adjustItems(evt event.Event) {
	var largest float64
	for _, child := range mnu.Children() {
		switch item := child.(type) {
		case *MenuItem:
			pos := item.calculateAcceleratorPosition()
			if largest < pos {
				largest = pos
			}
		}
	}
	for _, child := range mnu.Children() {
		switch item := child.(type) {
		case *MenuItem:
			item.pos = largest
			item.highlighted = false
			item.validate()
		}
	}
}

func (mnu *Menu) open(evt event.Event) {
	bounds := mnu.item.LocalBounds()
	where := mnu.item.ToWindow(bounds.Point)
	size := mnu.preparePopup(mnu.item.Window(), &where, 0)
	if mnu.attachToBottom {
		where.Y += bounds.Height
	} else {
		where.X += bounds.Width
	}
	mnu.showPopup(where, size)
}

func (mnu *Menu) close(evt event.Event) {
	if mnu.wnd != nil {
		wnd := mnu.wnd
		mnu.wnd = nil
		wnd.Close()
		mnu.item.setMenuOpen(false)
	}
}

// Dispose releases any operating system resources associated with this menu. It will also
// call Dispose() on all menu items it contains.
func (mnu *Menu) Dispose() {
	for _, child := range mnu.Children() {
		switch item := child.(type) {
		case menu.Item:
			item.Dispose()
		}
	}
}

// Popup displays the menu within the window. An attempt will be made to position the 'item'
// at 'where' within the window.
func (mnu *Menu) Popup(windowID int64, where geom.Point, width float64, item menu.Item) {
	wnd := window.ByID(windowID)
	if wnd != nil {
		size := mnu.preparePopup(wnd, &where, width)
		if item != nil {
			where.Add(geom.Point{X: 0, Y: -item.(*MenuItem).Location().Y})
		}
		mnu.showPopup(where, size)
	}
}

func (mnu *Menu) preparePopup(wnd ui.Window, where *geom.Point, width float64) geom.Size {
	mnu.adjustItems(nil)
	_, pref, _ := mnu.Layout().Sizes(layout.NoHintSize)
	if pref.Width < width {
		pref.Width = width
	}
	mnu.SetBounds(geom.Rect{Size: pref})
	mnu.Layout().Layout()
	where.Add(wnd.ContentFrame().Point)
	return pref
}

func (mnu *Menu) showPopup(where geom.Point, size geom.Size) {
	focus := window.KeyWindow()
	wnd := window.NewMenuWindow(focus, where, size)
	wnd.Content().AddChild(mnu)
	wnd.EventHandlers().Add(event.FocusLostType, mnu.close)
	wnd.ToFront()
	mnu.wnd = wnd
	mnu.item.setMenuOpen(true)
}

func (mnu *Menu) processKeyDown(evt *event.KeyDown) bool {
	for _, child := range mnu.Children() {
		switch item := child.(type) {
		case *MenuItem:
			if item.processKeyDown(evt) {
				return true
			}
		}
	}
	return false
}
