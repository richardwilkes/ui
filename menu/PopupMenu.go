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
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/geom"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/layout"
	"github.com/richardwilkes/ui/theme"
	"github.com/richardwilkes/ui/widget"
)

// PopupMenu represents a clickable button that displays a menu of choices.
type PopupMenu struct {
	widget.Block
	Theme         *theme.PopupMenu // The theme the popup menu will use to draw itself.
	OnSelection   func()           // Called when a new item is chosen by the user.
	items         []interface{}
	selectedIndex int
}

type separationMarker struct {
}

// NewPopupMenu creates a new PopupMenu.
func NewPopupMenu() *PopupMenu {
	pm := &PopupMenu{}
	pm.selectedIndex = -1
	pm.Theme = theme.StdPopupMenu
	pm.SetFocusable(true)
	pm.SetSizer(pm)
	handlers := pm.EventHandlers()
	handlers.Add(event.Paint, pm.paint)
	handlers.Add(event.MouseDown, pm.mouseDown)
	handlers.Add(event.FocusGained, pm.focusChanged)
	handlers.Add(event.FocusLost, pm.focusChanged)
	handlers.Add(event.KeyDown, pm.keyDown)
	return pm
}

// Sizes implements Sizer
func (pm *PopupMenu) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	var hSpace = pm.Theme.HorizontalMargin*3 + 2
	var vSpace = pm.Theme.VerticalMargin*2 + 2
	if hint.Width != layout.NoHint {
		hint.Width -= hSpace
		if hint.Width < pm.Theme.MinimumTextWidth {
			hint.Width = pm.Theme.MinimumTextWidth
		}
	}
	if hint.Height != layout.NoHint {
		hint.Height -= vSpace
		if hint.Height < 1 {
			hint.Height = 1
		}
	}
	size := pm.measureConstrained(hint)
	size.Width += hSpace + size.Height*0.75
	size.Height += vSpace
	size.GrowToInteger()
	if border := pm.Border(); border != nil {
		size.AddInsets(border.Insets())
	}
	return size, size, layout.DefaultMaxSize(size)
}

func (pm *PopupMenu) paint(event *event.Event) {
	var hSpace = pm.Theme.HorizontalMargin*2 + 2
	var vSpace = pm.Theme.VerticalMargin*2 + 2
	bounds := pm.LocalInsetBounds()
	path := geom.NewPath()
	path.MoveTo(bounds.X, bounds.Y+pm.Theme.CornerRadius)
	path.QuadCurveTo(bounds.X, bounds.Y, bounds.X+pm.Theme.CornerRadius, bounds.Y)
	path.LineTo(bounds.X+bounds.Width-pm.Theme.CornerRadius, bounds.Y)
	path.QuadCurveTo(bounds.X+bounds.Width, bounds.Y, bounds.X+bounds.Width, bounds.Y+pm.Theme.CornerRadius)
	path.LineTo(bounds.X+bounds.Width, bounds.Y+bounds.Height-pm.Theme.CornerRadius)
	path.QuadCurveTo(bounds.X+bounds.Width, bounds.Y+bounds.Height, bounds.X+bounds.Width-pm.Theme.CornerRadius, bounds.Y+bounds.Height)
	path.LineTo(bounds.X+pm.Theme.CornerRadius, bounds.Y+bounds.Height)
	path.QuadCurveTo(bounds.X, bounds.Y+bounds.Height, bounds.X, bounds.Y+bounds.Height-pm.Theme.CornerRadius)
	path.ClosePath()
	gc := event.GC
	gc.AddPath(path)
	gc.Clip()
	base := pm.BaseBackground()
	gc.DrawLinearGradient(pm.Theme.Gradient(base), bounds.X+bounds.Width/2, bounds.Y+1, bounds.X+bounds.Width/2, bounds.Y+bounds.Height-1)
	gc.AddPath(path)
	gc.SetStrokeColor(base.AdjustBrightness(pm.Theme.OutlineAdjustment))
	gc.StrokePath()
	triWidth := (bounds.Height*0.75 - vSpace)
	triHeight := triWidth / 2
	gc.BeginPath()
	gc.MoveTo(bounds.X+bounds.Width-pm.Theme.HorizontalMargin, bounds.Y+(bounds.Height-triHeight)/2)
	gc.LineTo(bounds.X+bounds.Width-pm.Theme.HorizontalMargin-triWidth, bounds.Y+(bounds.Height-triHeight)/2)
	gc.LineTo(bounds.X+bounds.Width-pm.Theme.HorizontalMargin-triWidth/2, bounds.Y+(bounds.Height-triHeight)/2+triHeight)
	gc.ClosePath()
	gc.SetFillColor(pm.TextColor())
	gc.FillPath()
	bounds.X += pm.Theme.HorizontalMargin + 1
	bounds.Y += pm.Theme.VerticalMargin + 1
	bounds.Height -= vSpace
	bounds.Width -= hSpace + bounds.Height
	gc.DrawAttributedTextConstrained(bounds, pm.title(), draw.TextModeFill)
}

func (pm *PopupMenu) mouseDown(event *event.Event) {
	pm.Click()
	event.Discard = true
}

func (pm *PopupMenu) focusChanged(event *event.Event) {
	pm.Repaint()
}

func (pm *PopupMenu) keyDown(event *event.Event) {
	if event.KeyCode == keys.Return || event.KeyCode == keys.Enter || event.KeyCode == keys.Space {
		event.Done = true
		pm.Click()
	}
}

// Click performs any animation associated with a click and triggers the popup menu to appear.
func (pm *PopupMenu) Click() {
	hasItem := false
	menu := NewMenu("")
	defer menu.Dispose()
	for i := range pm.items {
		if pm.addItemToMenu(menu, i) {
			hasItem = true
		}
	}
	if hasItem {
		menu.Popup(pm, pm.LocalInsetBounds().Point, menu.Item(pm.selectedIndex))
	}
}

func (pm *PopupMenu) addItemToMenu(menu *Menu, index int) bool {
	one := pm.items[index]
	switch one.(type) {
	case *separationMarker:
		menu.AddSeparator()
		return false
	default:
		menu.AddItem(fmt.Sprintf("%v", one), "", func(item *Item) {
			pm.SelectIndex(index)
			if pm.OnSelection != nil {
				pm.OnSelection()
			}
		}, nil)
		return true
	}
}

// AddItem appends an item to the end of the PopupMenu.
func (pm *PopupMenu) AddItem(item interface{}) {
	pm.items = append(pm.items, item)
}

// AddSeparator adds a separator to the end of the PopupMenu.
func (pm *PopupMenu) AddSeparator() {
	pm.items = append(pm.items, &separationMarker{})
}

// IndexOfItem returns the index of the specified item. -1 will be returned if the item isn't
// present.
func (pm *PopupMenu) IndexOfItem(item interface{}) int {
	for i, one := range pm.items {
		if one == item {
			return i
		}
	}
	return -1
}

// RemoveItem from the PopupMenu.
func (pm *PopupMenu) RemoveItem(item interface{}) {
	pm.RemoveItemAt(pm.IndexOfItem(item))
}

// RemoveItemAt the specified index from the PopupMenu.
func (pm *PopupMenu) RemoveItemAt(index int) {
	if index >= 0 {
		length := len(pm.items)
		if index < length {
			if pm.selectedIndex == index {
				pm.selectedIndex = -1
				pm.Repaint()
			} else if pm.selectedIndex > index {
				pm.selectedIndex--
			}
			copy(pm.items[index:], pm.items[index+1:])
			length--
			pm.items[length] = nil
			pm.items = pm.items[:length]
		}
	}
}

// Selected returns the currently selected item or nil.
func (pm *PopupMenu) Selected() interface{} {
	if pm.selectedIndex >= 0 && pm.selectedIndex < len(pm.items) {
		return pm.items[pm.selectedIndex]
	}
	return nil
}

// SelectedIndex returns the currently selected item index or -1.
func (pm *PopupMenu) SelectedIndex() int {
	return pm.selectedIndex
}

// Select an item.
func (pm *PopupMenu) Select(item interface{}) {
	pm.SelectIndex(pm.IndexOfItem(item))
}

// SelectIndex selects an item by its index.
func (pm *PopupMenu) SelectIndex(index int) {
	if index != pm.selectedIndex && index >= 0 && index < len(pm.items) {
		pm.selectedIndex = index
		pm.Repaint()
	}
}

// BaseBackground returns this popup menu's current base background color.
func (pm *PopupMenu) BaseBackground() color.Color {
	switch {
	case !pm.Enabled():
		return pm.Theme.Background.AdjustBrightness(pm.Theme.DisabledAdjustment)
	case pm.Focused():
		return pm.Theme.Background.Blend(color.KeyboardFocus, 0.5)
	default:
		return pm.Theme.Background
	}
}

// TextColor returns this popup menu's current text color.
func (pm *PopupMenu) TextColor() color.Color {
	if !pm.Enabled() {
		return pm.Theme.TextWhenDisabled
	}
	if pm.BaseBackground().Luminance() > 0.65 {
		return pm.Theme.TextWhenLight
	}
	return pm.Theme.TextWhenDark
}

func (pm *PopupMenu) title() *draw.Text {
	title := ""
	if pm.selectedIndex >= 0 && pm.selectedIndex < len(pm.items) {
		title = fmt.Sprintf("%v", pm.items[pm.selectedIndex])
	}
	return draw.NewText(title, pm.TextColor(), pm.Theme.Font)
}

func (pm *PopupMenu) measureConstrained(hint geom.Size) geom.Size {
	var largest geom.Size
	for _, one := range pm.items {
		size, _ := draw.NewText(fmt.Sprintf("%v", one), color.Black, pm.Theme.Font).MeasureConstrained(hint)
		if largest.Width < size.Width {
			largest.Width = size.Width
		}
		if largest.Height < size.Height {
			largest.Height = size.Height
		}
	}
	return largest
}
