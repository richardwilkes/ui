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
	"reflect"
	"time"
)

// Block is the basic graphical block in a window.
type Block struct {
	eventHandlers map[int][]EventHandler
	sizer         Sizer
	layout        Layout
	border        Border
	window        *Window
	parent        Widget
	children      []Widget
	bounds        Rect
	layoutData    interface{}
	background    Color
	needLayout    bool
	disabled      bool
	focused       bool
}

// NewBlock creates a new, empty block.
func NewBlock() *Block {
	return &Block{}
}

func (b *Block) EventHandlers() map[int][]EventHandler {
	return b.eventHandlers
}

func (b *Block) AddEventHandler(eventType int, handler EventHandler) {
	if b.eventHandlers == nil {
		b.eventHandlers = make(map[int][]EventHandler)
	}
	b.eventHandlers[eventType] = append(b.eventHandlers[eventType], handler)
}

func (b *Block) RemoveEventHandler(eventType int, handler EventHandler) {
	if b.eventHandlers != nil {
		hPtr := reflect.ValueOf(handler).Pointer()
		handlers := b.eventHandlers[eventType]
		for i, one := range handlers {
			if reflect.ValueOf(one).Pointer() == hPtr {
				if len(handlers) == 1 {
					delete(b.eventHandlers, eventType)
				} else {
					copy(handlers[i:], handlers[i+1:])
					length := len(handlers) - 1
					handlers[length] = nil
					b.eventHandlers[eventType] = handlers[:length]
				}
				break
			}
		}
	}
}

// Sizer implements the Widget interface.
func (b *Block) Sizer() Sizer {
	return b.sizer
}

// SetSizer implements the Widget interface.
func (b *Block) SetSizer(sizer Sizer) {
	b.sizer = sizer
}

// Layout implements the Widget interface.
func (b *Block) Layout() Layout {
	return b.layout
}

// SetLayout implements the Widget interface.
func (b *Block) SetLayout(layout Layout) {
	b.layout = layout
	b.SetNeedLayout(true)
}

// NeedLayout implements the Widget interface.
func (b *Block) NeedLayout() bool {
	return b.needLayout
}

// SetNeedLayout implements the Widget interface.
func (b *Block) SetNeedLayout(needLayout bool) {
	b.needLayout = needLayout
}

// LayoutData implements the Widget interface.
func (b *Block) LayoutData() interface{} {
	return b.layoutData
}

// SetLayoutData implements the Widget interface.
func (b *Block) SetLayoutData(data interface{}) {
	if b.layoutData != data {
		b.layoutData = data
		b.SetNeedLayout(true)
	}
}

// ValidateLayout implements the Widget interface.
func (b *Block) ValidateLayout() {
	if b.NeedLayout() {
		if layout := b.Layout(); layout != nil {
			layout.Layout()
			b.Repaint()
		}
		b.SetNeedLayout(false)
	}
	for _, child := range b.children {
		child.ValidateLayout()
	}
}

// Border implements the Widget interface.
func (b *Block) Border() Border {
	return b.border
}

// SetBorder implements the Widget interface.
func (b *Block) SetBorder(border Border) {
	b.border = border
}

// Repaint implements the Widget interface.
func (b *Block) Repaint() {
	b.RepaintBounds(b.LocalBounds())
}

// RepaintBounds implements the Widget interface.
func (b *Block) RepaintBounds(bounds Rect) {
	bounds.Intersect(b.LocalBounds())
	if !bounds.IsEmpty() {
		if p := b.Parent(); p != nil {
			bounds.X += b.bounds.X
			bounds.Y += b.bounds.Y
			p.RepaintBounds(bounds)
		} else if b.RootOfWindow() {
			b.Window().RepaintBounds(bounds)
		}
	}
}

// Paint implements the Widget interface.
func (b *Block) Paint(g Graphics, dirty Rect) {
	dirty.Intersect(b.LocalBounds())
	if !dirty.IsEmpty() {
		b.paintSelf(g, dirty)
		for _, child := range b.children {
			adjusted := dirty
			adjusted.Intersect(child.Bounds())
			if !adjusted.IsEmpty() {
				b.paintChild(child, g, adjusted)
			}
		}
	}
}

func (b *Block) paintSelf(g Graphics, dirty Rect) {
	g.Save()
	defer g.Restore()
	g.ClipRect(dirty)
	if b.background.Alpha() > 0 {
		g.SetFillColor(b.background)
		g.FillRect(dirty)
	}
	b.paintBorder(g)
	event := &Event{Type: PaintEvent, When: time.Now(), Target: b, GC: g, DirtyRect: dirty}
	event.Dispatch()
}

func (b *Block) paintBorder(g Graphics) {
	if border := b.Border(); border != nil {
		g.Save()
		defer g.Restore()
		border.PaintBorder(g, b.LocalBounds())
	}
}

func (b *Block) paintChild(child Widget, g Graphics, dirty Rect) {
	g.Save()
	defer g.Restore()
	bounds := child.Bounds()
	g.Translate(bounds.X, bounds.Y)
	dirty.X -= bounds.X
	dirty.Y -= bounds.Y
	child.Paint(g, dirty)
}

// Enabled implements the Widget interface.
func (b *Block) Enabled() bool {
	return !b.disabled
}

// SetEnabled implements the Widget interface.
func (b *Block) SetEnabled(enabled bool) {
	disabled := !enabled
	if b.disabled != disabled {
		b.disabled = disabled
		b.Repaint()
	}
}

// Focused implements the Widget interface.
func (b *Block) Focused() bool {
	return b.focused && !b.disabled
}

// SetFocused implements the Widget interface.
func (b *Block) SetFocused(focused bool) {
	if b.focused != focused {
		b.focused = focused
		b.Repaint()
	}
}

// Children implements the Widget interface.
func (b *Block) Children() []Widget {
	return b.children
}

// IndexOfChild implements the Widget interface.
func (b *Block) IndexOfChild(child Widget) int {
	for i, one := range b.children {
		if one == child {
			return i
		}
	}
	return -1
}

// AddChild implements the Widget interface.
func (b *Block) AddChild(child Widget) {
	child.RemoveFromParent()
	b.children = append(b.children, child)
	child.SetParent(b)
	b.SetNeedLayout(true)
}

// AddChildAtIndex implements the Widget interface.
func (b *Block) AddChildAtIndex(child Widget, index int) {
	child.RemoveFromParent()
	if index < 0 {
		index = 0
	}
	if index >= len(b.children) {
		b.children = append(b.children, child)
	} else {
		b.children = append(b.children, nil)
		copy(b.children[index+1:], b.children[index:])
		b.children[index] = child
	}
	child.SetParent(b)
	b.SetNeedLayout(true)
}

// RemoveChild implements the Widget interface.
func (b *Block) RemoveChild(child Widget) {
	b.RemoveChildAtIndex(b.IndexOfChild(child))
}

// RemoveChildAtIndex implements the Widget interface.
func (b *Block) RemoveChildAtIndex(index int) {
	if index >= 0 && index < len(b.children) {
		child := b.children[index]
		copy(b.children[index:], b.children[index+1:])
		length := len(b.children) - 1
		b.children[length] = nil
		b.children = b.children[:length]
		b.SetNeedLayout(true)
		child.SetParent(nil)
	}
}

// RemoveFromParent implements the Widget interface.
func (b *Block) RemoveFromParent() {
	if p := b.Parent(); p != nil {
		p.RemoveChild(b)
	}
}

// Parent implements the Widget interface.
func (b *Block) Parent() Widget {
	return b.parent
}

// SetParent implements the Widget interface.
func (b *Block) SetParent(parent Widget) {
	b.parent = parent
}

// Window implements the Widget interface.
func (b *Block) Window() *Window {
	if b.window != nil {
		return b.window
	}
	if b.parent != nil {
		return b.parent.Window()
	}
	return nil
}

// RootOfWindow implements the Widget interface.
func (b *Block) RootOfWindow() bool {
	return b.window != nil
}

// Bounds implements the Widget interface.
func (b *Block) Bounds() Rect {
	return b.bounds
}

// LocalBounds implements the Widget interface.
func (b *Block) LocalBounds() Rect {
	return b.bounds.CopyAndZeroLocation()
}

// LocalInsetBounds implements the Widget interface.
func (b *Block) LocalInsetBounds() Rect {
	bounds := b.LocalBounds()
	if border := b.Border(); border != nil {
		bounds.Inset(border.Insets())
	}
	return bounds
}

// SetBounds implements the Widget interface.
func (b *Block) SetBounds(bounds Rect) {
	moved := b.bounds.X != bounds.X || b.bounds.Y != bounds.Y
	resized := b.bounds.Width != bounds.Width || b.bounds.Height != bounds.Height
	if moved || resized {
		b.Repaint()
		if moved {
			b.bounds.Point = bounds.Point
		}
		if resized {
			b.bounds.Size = bounds.Size
			b.SetNeedLayout(true)
			event := &Event{Type: ResizeEvent, When: time.Now(), Target: b}
			event.Dispatch()
		}
		b.Repaint()
	}
}

// Location implements the Widget interface.
func (b *Block) Location() Point {
	return b.bounds.Point
}

// SetLocation implements the Widget interface.
func (b *Block) SetLocation(pt Point) {
	if b.bounds.Point != pt {
		b.Repaint()
		b.bounds.Point = pt
		b.Repaint()
	}
}

// Size implements the Widget interface.
func (b *Block) Size() Size {
	return b.bounds.Size
}

// SetSize implements the Widget interface.
func (b *Block) SetSize(size Size) {
	if b.bounds.Size != size {
		b.Repaint()
		b.bounds.Size = size
		b.SetNeedLayout(true)
		event := &Event{Type: ResizeEvent, When: time.Now(), Target: b}
		event.Dispatch()
		b.Repaint()
	}
}

// WidgetAt implements the Widget interface.
func (b *Block) WidgetAt(pt Point) Widget {
	for _, child := range b.children {
		bounds := child.Bounds()
		if bounds.Contains(pt) {
			pt.Subtract(bounds.Point)
			return child.WidgetAt(pt)
		}
	}
	return b
}

// ToWindow implements the Widget interface.
func (b *Block) ToWindow(pt Point) Point {
	pt.Add(b.bounds.Point)
	parent := b.parent
	for parent != nil {
		pt.Add(parent.Bounds().Point)
		parent = parent.Parent()
	}
	return pt
}

// FromWindow implements the Widget interface.
func (b *Block) FromWindow(pt Point) Point {
	pt.Subtract(b.bounds.Point)
	parent := b.parent
	for parent != nil {
		pt.Subtract(parent.Bounds().Point)
		parent = parent.Parent()
	}
	return pt
}

// Background implements the Widget interface.
func (b *Block) Background() Color {
	return b.background
}

// SetBackground implements the Widget interface.
func (b *Block) SetBackground(color Color) {
	if color != b.background {
		b.background = color
		b.Repaint()
	}
}
