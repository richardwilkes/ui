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
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/xmath"
)

// ScrollArea provides a widget that can hold another widget and show it through a scrollable
// viewport.
type ScrollArea struct {
	Block
	hBar    *ScrollBar
	vBar    *ScrollBar
	view    *Block
	content ui.Widget
}

// NewScrollArea creates a new ScrollArea with the specified block as its content. The content may
// be nil.
func NewScrollArea(content ui.Widget) *ScrollArea {
	sa := &ScrollArea{}
	handlers := sa.EventHandlers()
	handlers.Add(event.MouseWheelEvent, sa.mouseWheel)
	sa.view = NewBlock()
	sa.view.SetBackground(color.TextBackground)
	sa.view.EventHandlers().Add(event.ResizeEvent, sa.viewResized)
	sa.SetFocusable(true) // RAW: Revist... don't want to be focusable, but do want to participate
	handlers.Add(event.KeyDownEvent, sa.keyDown)
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
func (sa *ScrollArea) Content() ui.Widget {
	return sa.content
}

// SetContent sets the content block, replacing any existing one.
func (sa *ScrollArea) SetContent(content ui.Widget) {
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

func (sa *ScrollArea) viewResized(event *event.Event) {
	if sa.content != nil {
		vs := sa.view.Size()
		cl := sa.content.Location()
		cs := sa.content.Size()
		nl := cl
		if cl.Y != 0 && vs.Height > cl.Y+cs.Height {
			nl.Y = xmath.MinFloat32(vs.Height-cs.Height, 0)
		}
		if cl.X != 0 && vs.Width > cl.X+cs.Width {
			nl.X = xmath.MinFloat32(vs.Width-cs.Width, 0)
		}
		if nl != cl {
			sa.content.SetLocation(nl)
		}
	}
}

func (sa *ScrollArea) mouseWheel(event *event.Event) {
	if event.Delta.Y != 0 {
		sa.vBar.SetScrolledPosition(sa.ScrolledPosition(false) - event.Delta.Y*sa.LineScrollAmount(false, event.Delta.Y > 0))
	}
	if event.Delta.X != 0 {
		sa.hBar.SetScrolledPosition(sa.ScrolledPosition(true) - event.Delta.X*sa.LineScrollAmount(true, event.Delta.X > 0))
	}
	event.CascadeUp = false
}

func (sa *ScrollArea) keyDown(event *event.Event) {
	switch event.KeyCode {
	case keys.Up:
		event.Done = true
		sa.vBar.SetScrolledPosition(sa.ScrolledPosition(false) - sa.LineScrollAmount(false, true))
	case keys.Down:
		event.Done = true
		sa.vBar.SetScrolledPosition(sa.ScrolledPosition(false) + sa.LineScrollAmount(false, false))
	case keys.Left:
		event.Done = true
		sa.hBar.SetScrolledPosition(sa.ScrolledPosition(true) - sa.LineScrollAmount(true, true))
	case keys.Right:
		event.Done = true
		sa.hBar.SetScrolledPosition(sa.ScrolledPosition(true) + sa.LineScrollAmount(true, false))
	case keys.Home:
		event.Done = true
		var bar *ScrollBar
		if event.ShiftDown() {
			bar = sa.hBar
		} else {
			bar = sa.vBar
		}
		bar.SetScrolledPosition(0)
	case keys.End:
		event.Done = true
		var bar *ScrollBar
		horizontal := event.ShiftDown()
		if horizontal {
			bar = sa.hBar
		} else {
			bar = sa.vBar
		}
		bar.SetScrolledPosition(sa.ContentSize(horizontal))
	case keys.PageUp:
		event.Done = true
		var bar *ScrollBar
		horizontal := event.ShiftDown()
		if horizontal {
			bar = sa.hBar
		} else {
			bar = sa.vBar
		}
		bar.SetScrolledPosition(sa.ScrolledPosition(horizontal) - sa.PageScrollAmount(horizontal, true))
	case keys.PageDown:
		event.Done = true
		var bar *ScrollBar
		horizontal := event.ShiftDown()
		if horizontal {
			bar = sa.hBar
		} else {
			bar = sa.vBar
		}
		bar.SetScrolledPosition(sa.ScrolledPosition(horizontal) + sa.PageScrollAmount(horizontal, false))
	}
}
