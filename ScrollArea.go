// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

// ScrollArea provides a widget that can hold another widget and show it through a scrollable
// viewport.
type ScrollArea struct {
	Block
	hBar    *ScrollBar
	vBar    *ScrollBar
	view    *Block
	content Widget
}

// NewScrollArea creates a new ScrollArea with the specified block as its content. The content may
// be nil.
func NewScrollArea(content Widget) *ScrollArea {
	sa := &ScrollArea{}
	sa.view = NewBlock()
	sa.view.SetBackground(TextBackgroundColor)
	sa.view.SetResizeHandler(sa)
	sa.AddChild(sa.view)
	sa.hBar = NewScrollBar(true, sa)
	sa.vBar = NewScrollBar(false, sa)
	sa.content = content
	if content != nil {
		sa.view.AddChild(content)
	}
	newScrollLayout(sa)
	return sa
}

// View returns the view port block.
func (sa *ScrollArea) View() *Block {
	return sa.view
}

// Content returns the content block.
func (sa *ScrollArea) Content() Widget {
	return sa.content
}

// SetContent sets the content block, replacing any existing one.
func (sa *ScrollArea) SetContent(content Widget) {
	if sa.content != nil {
		sa.content.RemoveFromParent()
	}
	sa.content = content
	if content != nil {
		sa.view.AddChildAtIndex(content, 0)
	}
	sa.SetNeedLayout(true)
	sa.Repaint()
}

// LineScrollAmount implements Pager and Scrollable.
func (sa *ScrollArea) LineScrollAmount(horizontal, towardsStart bool) float32 {
	if sa.content != nil {
		if s, ok := interface{}(sa.content).(Pager); ok {
			return s.LineScrollAmount(horizontal, towardsStart)
		}
	}
	return 16
}

// PageScrollAmount implements Pager and Scrollable.
func (sa *ScrollArea) PageScrollAmount(horizontal, towardsStart bool) float32 {
	if sa.content != nil {
		if s, ok := interface{}(sa.content).(Pager); ok {
			return s.PageScrollAmount(horizontal, towardsStart)
		}
	}
	size := sa.view.Size()
	if horizontal {
		return size.Width
	}
	return size.Height
}

// ScrolledPosition implements Scrollable.
func (sa *ScrollArea) ScrolledPosition(horizontal bool) float32 {
	if sa.content == nil {
		return 0
	}
	loc := sa.content.Location()
	if horizontal {
		return -loc.X
	}
	return -loc.Y
}

// SetScrolledPosition implements Scrollable.
func (sa *ScrollArea) SetScrolledPosition(horizontal bool, position float32) {
	if sa.content != nil {
		loc := sa.content.Location()
		if horizontal {
			loc.X = -position
		} else {
			loc.Y = -position
		}
		sa.content.SetLocation(loc)
	}
}

// VisibleSize implements Scrollable.
func (sa *ScrollArea) VisibleSize(horizontal bool) float32 {
	size := sa.view.Size()
	if horizontal {
		return size.Width
	}
	return size.Height
}

// ContentSize implements Scrollable.
func (sa *ScrollArea) ContentSize(horizontal bool) float32 {
	if sa.content == nil {
		return 0
	}
	size := sa.content.Size()
	if horizontal {
		return size.Width
	}
	return size.Height
}

// Resized implements ResizeHandler and is called for the contained view.
func (sa *ScrollArea) Resized() {
	if sa.content != nil {
		vs := sa.view.Size()
		cl := sa.content.Location()
		cs := sa.content.Size()
		nl := cl
		if cl.Y != 0 && vs.Height > cl.Y+cs.Height {
			nl.Y = MinFloat32(vs.Height-cs.Height, 0)
		}
		if cl.X != 0 && vs.Width > cl.X+cs.Width {
			nl.X = MinFloat32(vs.Width-cs.Width, 0)
		}
		if nl != cl {
			sa.content.SetLocation(nl)
		}
	}
}
