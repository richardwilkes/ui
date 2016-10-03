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
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/border"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/layout"
	"github.com/richardwilkes/ui/layout/flex"
	"github.com/richardwilkes/ui/widget"
	"github.com/richardwilkes/ui/widget/separator"
	"github.com/richardwilkes/ui/widget/window"
)

type Menu struct {
	widget.Block
	item           *MenuItem
	wnd            ui.Window
	attachToBottom bool
}

func NewMenu(title string) *Menu {
	menu := &Menu{item: NewMenuItem(title, 0, nil)}
	menu.item.menu = menu
	menu.Describer = func() string {
		return fmt.Sprintf("Menu #%d (%s)", menu.ID(), menu.Title())
	}
	menu.SetBorder(border.NewLine(color.Gray, geom.Insets{Top: 1, Left: 1, Bottom: 1, Right: 1}))
	menu.item.EventHandlers().Add(event.SelectionType, menu.open)
	flex.NewLayout(menu).SetEqualColumns(true)
	return menu
}

func (menu *Menu) Title() string {
	return menu.item.Title
}

func (menu *Menu) AddItem(item *MenuItem) {
	menu.AddChild(item)
	item.EventHandlers().Add(event.ClosingType, menu.close)
	item.SetLayoutData(flex.NewData().SetHorizontalGrab(true).SetHorizontalAlignment(draw.AlignFill))
}

func (menu *Menu) AddSeparator() {
	sep := separator.New(true)
	menu.AddChild(sep)
	sep.SetLayoutData(flex.NewData().SetHorizontalGrab(true).SetHorizontalAlignment(draw.AlignFill))
}

func (menu *Menu) adjustItems(evt event.Event) {
	var largest float64
	for _, child := range menu.Children() {
		switch item := child.(type) {
		case *MenuItem:
			pos := item.calculateAcceleratorPosition()
			if largest < pos {
				largest = pos
			}
		}
	}
	for _, child := range menu.Children() {
		switch item := child.(type) {
		case *MenuItem:
			item.pos = largest
		}
	}
}

func (menu *Menu) open(evt event.Event) {
	bounds := menu.item.Bounds()
	where := menu.item.ToWindow(bounds.Point)
	where.Add(menu.item.Window().ContentFrame().Point)
	if menu.attachToBottom {
		where.Y += bounds.Height
	} else {
		where.X += bounds.Width
	}
	menu.adjustItems(nil)
	_, pref, _ := menu.Layout().Sizes(layout.NoHintSize)
	menu.SetBounds(geom.Rect{Size: pref})
	menu.Layout().Layout()
	wnd := window.NewWindowWithContentSize(where, pref, window.BorderlessWindowMask)
	wnd.RootWidget().AddChild(menu)
	wnd.EventHandlers().Add(event.FocusLostType, menu.close)
	wnd.ToFront()
	menu.wnd = wnd
	menu.item.menuOpen = true
	menu.item.Repaint()
}

func (menu *Menu) close(evt event.Event) {
	if menu.wnd != nil {
		wnd := menu.wnd
		menu.wnd = nil
		wnd.Close()
		menu.item.menuOpen = false
		menu.item.Repaint()
	}
}
