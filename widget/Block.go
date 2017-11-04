// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package widget

import (
	"fmt"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/border"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/layout"
	"github.com/richardwilkes/ui/object"
)

// Block is the basic graphical block in a window.
type Block struct {
	object.Base
	Describer     func() string
	eventHandlers *event.Handlers
	window        ui.Window
	parent        ui.Widget
	children      []ui.Widget
	sizer         layout.Sizer
	layout        layout.Layout
	border        border.Border
	bounds        geom.Rect
	layoutData    interface{}
	background    color.Color
	needLayout    bool
	disabled      bool
	focusable     bool
	grabsFocus    bool
}

// NewBlock creates a new, empty block.
func NewBlock() *Block {
	block := &Block{}
	block.InitTypeAndID(block)
	return block
}

func (b *Block) String() string {
	if b.Describer != nil {
		return b.Describer()
	}
	return fmt.Sprintf("Block #%d", b.ID())
}

// EventHandlers implements the event.Target interface.
func (b *Block) EventHandlers() *event.Handlers {
	if b.eventHandlers == nil {
		b.eventHandlers = &event.Handlers{}
	}
	return b.eventHandlers
}

// ParentTarget implements the event.Target interface.
func (b *Block) ParentTarget() event.Target {
	if b.parent != nil {
		return b.parent
	}
	return b.window
}

// Sizer implements the Widget interface.
func (b *Block) Sizer() layout.Sizer {
	return b.sizer
}

// SetSizer implements the Widget interface.
func (b *Block) SetSizer(sizer layout.Sizer) {
	b.sizer = sizer
}

// Layout implements the Widget interface.
func (b *Block) Layout() layout.Layout {
	return b.layout
}

// SetLayout implements the Widget interface.
func (b *Block) SetLayout(layout layout.Layout) {
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
func (b *Block) Border() border.Border {
	return b.border
}

// SetBorder implements the Widget interface.
func (b *Block) SetBorder(border border.Border) {
	b.border = border
}

// Repaint implements the Widget interface.
func (b *Block) Repaint() {
	b.RepaintBounds(b.LocalBounds())
}

// RepaintBounds implements the Widget interface.
func (b *Block) RepaintBounds(bounds geom.Rect) {
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
func (b *Block) Paint(gc *draw.Graphics, dirty geom.Rect) {
	dirty.Intersect(b.LocalBounds())
	if !dirty.IsEmpty() {
		gc.Save()
		gc.Rect(dirty)
		gc.Clip()
		gc.Save()
		if b.background.Alpha() > 0 {
			gc.SetColor(b.background)
			gc.FillRect(dirty)
		}
		event.Dispatch(event.NewPaint(b, gc, dirty))
		gc.Restore()
		for _, child := range b.children {
			adjusted := dirty
			bounds := child.Bounds()
			adjusted.Intersect(bounds)
			if !adjusted.IsEmpty() {
				gc.Save()
				gc.Translate(bounds.X, bounds.Y)
				adjusted.X -= bounds.X
				adjusted.Y -= bounds.Y
				child.Paint(gc, adjusted)
				gc.Restore()
			}
		}
		if border := b.Border(); border != nil {
			gc.Save()
			border.Draw(gc, b.LocalBounds())
			gc.Restore()
		}
		gc.Restore()
	}
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

// Focusable implements the Widget interface.
func (b *Block) Focusable() bool {
	return b.focusable && !b.disabled
}

// SetFocusable implements the Widget interface.
func (b *Block) SetFocusable(focusable bool) {
	if b.focusable != focusable {
		b.focusable = focusable
	}
}

// GrabFocusWhenClickedOn implements the Widget interface.
func (b *Block) GrabFocusWhenClickedOn() bool {
	return b.grabsFocus
}

// SetGrabFocusWhenClickedOn implements the Widget interface.
func (b *Block) SetGrabFocusWhenClickedOn(grabsFocus bool) {
	if b.grabsFocus != grabsFocus {
		b.grabsFocus = grabsFocus
	}
}

// Focused implements the Widget interface.
func (b *Block) Focused() bool {
	if window := b.Window(); window != nil {
		return window.Focused() && b.Is(window.Focus())
	}
	return false
}

// Children implements the Widget interface.
func (b *Block) Children() []ui.Widget {
	return b.children
}

// IndexOfChild implements the Widget interface.
func (b *Block) IndexOfChild(child ui.Widget) int {
	for i, one := range b.children {
		if one.ID() == child.ID() {
			return i
		}
	}
	return -1
}

// AddChild implements the Widget interface.
func (b *Block) AddChild(child ui.Widget) {
	child.RemoveFromParent()
	b.children = append(b.children, child)
	child.SetParent(b)
	b.SetNeedLayout(true)
}

// AddChildAtIndex implements the Widget interface.
func (b *Block) AddChildAtIndex(child ui.Widget, index int) {
	child.RemoveFromParent()
	if index < 0 || index >= len(b.children) {
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
func (b *Block) RemoveChild(child ui.Widget) {
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
func (b *Block) Parent() ui.Widget {
	return b.parent
}

// SetParent implements the Widget interface.
func (b *Block) SetParent(parent ui.Widget) {
	b.parent = parent
}

// Window implements the Widget interface.
func (b *Block) Window() ui.Window {
	if b.window != nil {
		return b.window
	}
	if b.parent != nil {
		return b.parent.Window()
	}
	return nil
}

// SetWindow sets the window of this Block and the Block's bounds to the window's ContentLocalFrame().
func (b *Block) SetWindow(wnd ui.Window) {
	b.window = wnd
	b.bounds = wnd.ContentLocalFrame()
}

// RootOfWindow implements the Widget interface.
func (b *Block) RootOfWindow() bool {
	return b.window != nil
}

// Bounds implements the Widget interface.
func (b *Block) Bounds() geom.Rect {
	return b.bounds
}

// LocalBounds implements the Widget interface.
func (b *Block) LocalBounds() geom.Rect {
	return b.bounds.CopyAndZeroLocation()
}

// LocalInsetBounds implements the Widget interface.
func (b *Block) LocalInsetBounds() geom.Rect {
	bounds := b.LocalBounds()
	if border := b.Border(); border != nil {
		bounds.Inset(border.Insets())
	}
	return bounds
}

// SetBounds implements the Widget interface.
func (b *Block) SetBounds(bounds geom.Rect) {
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
			event.Dispatch(event.NewResized(b))
		}
		b.Repaint()
	}
}

// Location implements the Widget interface.
func (b *Block) Location() geom.Point {
	return b.bounds.Point
}

// SetLocation implements the Widget interface.
func (b *Block) SetLocation(pt geom.Point) {
	if b.bounds.Point != pt {
		b.Repaint()
		b.bounds.Point = pt
		b.Repaint()
	}
}

// Size implements the Widget interface.
func (b *Block) Size() geom.Size {
	return b.bounds.Size
}

// SetSize implements the Widget interface.
func (b *Block) SetSize(size geom.Size) {
	if b.bounds.Size != size {
		b.Repaint()
		b.bounds.Size = size
		b.SetNeedLayout(true)
		event.Dispatch(event.NewResized(b))
		b.Repaint()
	}
}

// WidgetAt implements the Widget interface.
func (b *Block) WidgetAt(pt geom.Point) ui.Widget {
	for _, child := range b.children {
		bounds := child.Bounds()
		if bounds.ContainsPoint(pt) {
			pt.Subtract(bounds.Point)
			return child.WidgetAt(pt)
		}
	}
	return b
}

// ToWindow implements the Widget interface.
func (b *Block) ToWindow(pt geom.Point) geom.Point {
	pt.Add(b.bounds.Point)
	parent := b.parent
	for parent != nil {
		pt.Add(parent.Bounds().Point)
		parent = parent.Parent()
	}
	return pt
}

// FromWindow implements the Widget interface.
func (b *Block) FromWindow(pt geom.Point) geom.Point {
	pt.Subtract(b.bounds.Point)
	parent := b.parent
	for parent != nil {
		pt.Subtract(parent.Bounds().Point)
		parent = parent.Parent()
	}
	return pt
}

// Background implements the Widget interface.
func (b *Block) Background() color.Color {
	return b.background
}

// SetBackground implements the Widget interface.
func (b *Block) SetBackground(color color.Color) {
	if color != b.background {
		b.background = color
		b.Repaint()
	}
}

// ScrollIntoView attempts to scroll the block into view.
func (b *Block) ScrollIntoView() {
	// RAW: Implement
}
