// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package list

import (
	"fmt"
	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/border"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/layout"
	"github.com/richardwilkes/ui/widget"
	"github.com/richardwilkes/xmath"
	"math"
)

// List provides a control that allows the user to select from a list of items, represented by cells.
type List struct {
	widget.Block
	factory        widget.CellFactory
	rows           []interface{}
	Selection      xmath.BitSet
	savedSelection *xmath.BitSet
	anchor         int
	pressed        bool
}

// New creates a new List control.
func New(factory widget.CellFactory) *List {
	list := &List{factory: factory, anchor: -1}
	list.InitTypeAndID(list)
	list.Describer = func() string { return fmt.Sprintf("List #%d", list.ID()) }
	list.SetBackground(color.White)
	list.SetBorder(border.NewEmpty(geom.NewUniformInsets(2)))
	list.SetFocusable(true)
	list.SetGrabFocusWhenClickedOn(true)
	list.SetSizer(list)
	handlers := list.EventHandlers()
	handlers.Add(event.PaintType, list.paint)
	handlers.Add(event.MouseDownType, list.mouseDown)
	handlers.Add(event.MouseDraggedType, list.mouseDragged)
	handlers.Add(event.MouseUpType, list.mouseUp)
	handlers.Add(event.KeyDownType, list.keyDown)
	return list
}

// Sizes implements Sizer
func (list *List) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	max = layout.DefaultMaxSize(max)
	height := math.Ceil(list.factory.CellHeight())
	if height < 1 {
		height = layout.NoHint
	}
	size := geom.Size{Width: hint.Width, Height: height}
	for i, row := range list.rows {
		cell := list.factory.CreateCell(list, row, i, false, false)
		_, cpref, cmax := ui.Sizes(cell, size)
		cpref.GrowToInteger()
		cmax.GrowToInteger()
		if pref.Width < cpref.Width {
			pref.Width = cpref.Width
		}
		if max.Width < cmax.Width {
			max.Width = cmax.Width
		}
		if height < 1 {
			pref.Height += cpref.Height
			max.Height += cmax.Height
		}
	}
	if height >= 1 {
		count := float64(len(list.rows))
		if count < 1 {
			count = 1
		}
		pref.Height = count * height
		max.Height = count * height
		if max.Height < layout.DefaultMax {
			max.Height = layout.DefaultMax
		}
	}
	pref.GrowToInteger()
	max.GrowToInteger()
	return pref, pref, max
}

// Append values to the list of items.
func (list *List) Append(values ...interface{}) {
	list.rows = append(list.rows, values...)
	list.Repaint()
}

// Insert values at the specified index.
func (list *List) Insert(index int, values ...interface{}) {
	list.rows = append(list.rows[:index], append(values, list.rows[index:]...)...)
	list.Repaint()
}

// Remove the item at the specified index.
func (list *List) Remove(index int) {
	copy(list.rows[index:], list.rows[index+1:])
	size := len(list.rows) - 1
	list.rows[size] = nil
	list.rows = list.rows[:size]
	list.Repaint()
}

func (list *List) rowAt(y float64) (index int, top float64) {
	count := len(list.rows)
	top = list.LocalInsetBounds().Y
	cellHeight := math.Ceil(list.factory.CellHeight())
	if cellHeight < 1 {
		for index < count {
			cell := list.factory.CreateCell(list, list.rows[index], index, false, false)
			_, pref, _ := ui.Sizes(cell, layout.NoHintSize)
			pref.GrowToInteger()
			if top+pref.Height >= y {
				break
			}
			top += pref.Height
			index++
		}
	} else {
		index = int(math.Floor((y - top) / cellHeight))
		top += float64(index) * cellHeight
	}
	if index >= count {
		index = -1
		top = 0
	}
	return
}

func (list *List) mouseDown(evt event.Event) {
	list.Window().SetFocus(list)
	list.savedSelection = list.Selection.Clone()
	e := evt.(*event.MouseDown)
	if index, _ := list.rowAt(list.FromWindow(e.Where()).Y); index >= 0 {
		if e.Modifiers().CommandDown() {
			list.Selection.Flip(index)
			list.anchor = index
		} else if e.Modifiers().ShiftDown() {
			if list.anchor != -1 {
				list.Selection.SetRange(list.anchor, index)
			} else {
				list.Selection.Set(index)
				list.anchor = index
			}
		} else if list.Selection.State(index) {
			list.anchor = index
			if e.Clicks() == 2 {
				event.Dispatch(event.NewClick(list))
				e.Discard()
				return
			}
		} else {
			list.Selection.Reset()
			list.Selection.Set(index)
			list.anchor = index
		}
		if !list.Selection.Equal(list.savedSelection) {
			list.Repaint()
		}
	}
	list.pressed = true
}

func (list *List) mouseDragged(evt event.Event) {
	if list.pressed {
		e := evt.(*event.MouseDragged)
		list.Selection.Copy(list.savedSelection)
		if index, _ := list.rowAt(list.FromWindow(e.Where()).Y); index >= 0 {
			if list.anchor == -1 {
				list.anchor = index
			}
			if e.Modifiers().CommandDown() {
				list.Selection.FlipRange(list.anchor, index)
			} else if e.Modifiers().ShiftDown() {
				list.Selection.SetRange(list.anchor, index)
			} else {
				list.Selection.Reset()
				list.Selection.SetRange(list.anchor, index)
			}
			if !list.Selection.Equal(list.savedSelection) {
				list.Repaint()
			}
		}
	}
}

func (list *List) mouseUp(evt event.Event) {
	if list.pressed {
		list.pressed = false
		if !list.Selection.Equal(list.savedSelection) {
			event.Dispatch(event.NewSelection(list))
		}
	}
	list.savedSelection = nil
}

func (list *List) paint(evt event.Event) {
	e := evt.(*event.Paint)
	dirty := e.DirtyRect()
	index, y := list.rowAt(dirty.Y)
	if index >= 0 {
		cellHeight := math.Ceil(list.factory.CellHeight())
		count := len(list.rows)
		ymax := dirty.Y + dirty.Height
		focused := list.Focused()
		selCount := list.Selection.Count()
		fullBounds := list.LocalBounds()
		bounds := list.LocalInsetBounds()
		gc := e.GC()
		for index < count && y < ymax {
			selected := list.Selection.State(index)
			cell := list.factory.CreateCell(list, list.rows[index], index, selected, focused && selected && selCount == 1)
			cellBounds := geom.Rect{Point: geom.Point{X: bounds.X, Y: y}, Size: geom.Size{Width: bounds.Width, Height: cellHeight}}
			if cellHeight < 1 {
				_, pref, _ := ui.Sizes(cell, layout.NoHintSize)
				pref.GrowToInteger()
				cellBounds.Height = pref.Height
			}
			cell.SetBounds(cellBounds)
			y += cellBounds.Height
			if selected {
				gc.SetColor(color.SelectedTextBackground)
				gc.FillRect(geom.Rect{Point: geom.Point{X: fullBounds.X, Y: cellBounds.Y}, Size: geom.Size{Width: fullBounds.Width, Height: cellBounds.Height}})
			}
			gc.Save()
			tl := cellBounds.Point
			dirty.Point.Subtract(tl)
			gc.Translate(cellBounds.X, cellBounds.Y)
			cellBounds.X = 0
			cellBounds.Y = 0
			cell.Paint(gc, dirty)
			dirty.Point.Add(tl)
			gc.Restore()
			index++
		}
	}
}

func (list *List) keyDown(evt event.Event) {
	e := evt.(*event.KeyDown)
	code := e.Code()
	if keys.IsControlAction(code) {
		if list.Selection.Count() > 0 {
			event.Dispatch(event.NewClick(list))
		}
	} else {
		switch code {
		case keys.VirtualKeyUp, keys.VirtualKeyNumPadUp:
			evt.Finish()
			var first int
			if list.Selection.Count() == 0 {
				first = len(list.rows) - 1
			} else {
				first = list.Selection.FirstSet() - 1
				if first < 0 {
					first = 0
				}
			}
			list.Select(e.Modifiers().ShiftDown(), first)
			event.Dispatch(event.NewSelection(list))
		case keys.VirtualKeyDown, keys.VirtualKeyNumPadDown:
			evt.Finish()
			last := list.Selection.LastSet() + 1
			if last >= len(list.rows) {
				last = len(list.rows) - 1
			}
			list.Select(e.Modifiers().ShiftDown(), last)
			event.Dispatch(event.NewSelection(list))
		case keys.VirtualKeyHome, keys.VirtualKeyNumPadHome:
			evt.Finish()
			list.Select(e.Modifiers().ShiftDown(), 0)
			event.Dispatch(event.NewSelection(list))
		case keys.VirtualKeyEnd, keys.VirtualKeyNumPadEnd:
			evt.Finish()
			list.Select(e.Modifiers().ShiftDown(), len(list.rows)-1)
			event.Dispatch(event.NewSelection(list))
		}
	}
}

func (list *List) CanSelectAll() bool {
	return list.Selection.Count() < len(list.rows)
}

func (list *List) SelectAll() {
	list.SelectRange(0, len(list.rows)-1, false)
}

// SelectRange selects items from 'start' to 'end', inclusive. If 'append' is true, then any
// existing selection is added to rather than replaced.
func (list *List) SelectRange(start, end int, append bool) {
	if !append {
		list.Selection.Reset()
		list.anchor = -1
	}
	max := len(list.rows) - 1
	start = xmath.MaxInt(xmath.MinInt(start, max), 0)
	end = xmath.MaxInt(xmath.MinInt(end, max), 0)
	list.Selection.SetRange(start, end)
	if list.anchor == -1 {
		list.anchor = start
	}
	list.Repaint()
}

// Select items at the specified indexes. If 'append' is true, then any existing selection is added
// to rather than replaced.
func (list *List) Select(append bool, index ...int) {
	if !append {
		list.Selection.Reset()
		list.anchor = -1
	}
	max := len(list.rows)
	for _, v := range index {
		if v >= 0 && v < max {
			list.Selection.Set(v)
			if list.anchor == -1 {
				list.anchor = v
			}
		}
	}
	list.Repaint()
}
