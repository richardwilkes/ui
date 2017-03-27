// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package scrollarea

import (
	"fmt"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/widget"
	"github.com/richardwilkes/ui/widget/scrollbar"
	"math"
)

// Possible ways to handle auto-sizing of the scroll content's preferred size.
const (
	Unmodified Behavior = iota
	FillWidth
	FillHeight
	Fill
)

// Behavior controls how auto-sizing of the scroll content's preferred size is handled.
type Behavior int

// ScrollArea provides a widget that can hold another widget and show it through a scrollable
// viewport.
type ScrollArea struct {
	widget.Block
	Theme    *Theme // The theme the ScrollArea will use to draw itself.
	hBar     *scrollbar.ScrollBar
	vBar     *scrollbar.ScrollBar
	view     *widget.Block
	content  ui.Widget
	behavior Behavior
}

// New creates a new ScrollArea with the specified block as its content. The content may be nil.
func New(content ui.Widget, behavior Behavior) *ScrollArea {
	sa := &ScrollArea{Theme: StdTheme}
	sa.InitTypeAndID(sa)
	sa.Describer = func() string { return fmt.Sprintf("ScrollArea #%d", sa.ID()) }
	sa.SetBorder(sa.Theme.Border)
	handlers := sa.EventHandlers()
	handlers.Add(event.MouseWheelType, sa.mouseWheel)
	handlers.Add(event.KeyDownType, sa.keyDown)
	handlers.Add(event.FocusGainedType, sa.focusGained)
	handlers.Add(event.FocusLostType, sa.focusLost)
	sa.view = widget.NewBlock()
	sa.view.SetBackground(color.TextBackground)
	handlers = sa.view.EventHandlers()
	handlers.Add(event.ResizedType, sa.viewResized)
	sa.SetFocusable(true)
	sa.AddChild(sa.view)
	sa.hBar = scrollbar.New(true, sa)
	sa.vBar = scrollbar.New(false, sa)
	newScrollLayout(sa)
	if content != nil {
		sa.SetContent(content, behavior)
	}
	return sa
}

// Content returns the content block.
func (sa *ScrollArea) Content() ui.Widget {
	return sa.content
}

// SetContent sets the content block, replacing any existing one.
func (sa *ScrollArea) SetContent(content ui.Widget, behavior Behavior) {
	if sa.content != nil {
		handlers := sa.content.EventHandlers()
		handlers.Remove(event.ResizedType, sa.viewResized)
		handlers.Remove(event.FocusGainedType, sa.focusGained)
		handlers.Remove(event.FocusLostType, sa.focusLost)
		sa.content.RemoveFromParent()
	}
	sa.content = content
	sa.behavior = behavior
	if sa.content != nil {
		sa.view.AddChildAtIndex(sa.content, 0)
		handlers := sa.content.EventHandlers()
		handlers.Add(event.ResizedType, sa.viewResized)
		handlers.Add(event.FocusGainedType, sa.focusGained)
		handlers.Add(event.FocusLostType, sa.focusLost)
		sa.SetFocusable(false)
	} else {
		sa.SetFocusable(true)
	}
	sa.SetNeedLayout(true)
	sa.Repaint()
}

// LineScrollAmount implements Pager and Scrollable.
func (sa *ScrollArea) LineScrollAmount(horizontal, towardsStart bool) float64 {
	if sa.content != nil {
		if s, ok := interface{}(sa.content).(scrollbar.Pager); ok {
			return s.LineScrollAmount(horizontal, towardsStart)
		}
	}
	return 16
}

// PageScrollAmount implements Pager and Scrollable.
func (sa *ScrollArea) PageScrollAmount(horizontal, towardsStart bool) float64 {
	if sa.content != nil {
		if s, ok := interface{}(sa.content).(scrollbar.Pager); ok {
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
func (sa *ScrollArea) ScrolledPosition(horizontal bool) float64 {
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
func (sa *ScrollArea) SetScrolledPosition(horizontal bool, position float64) {
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
func (sa *ScrollArea) VisibleSize(horizontal bool) float64 {
	size := sa.view.Size()
	if horizontal {
		return size.Width
	}
	return size.Height
}

// ContentSize implements Scrollable.
func (sa *ScrollArea) ContentSize(horizontal bool) float64 {
	if sa.content == nil {
		return 0
	}
	size := sa.content.Size()
	if horizontal {
		return size.Width
	}
	return size.Height
}

func (sa *ScrollArea) viewResized(evt event.Event) {
	if sa.content != nil {
		vs := sa.view.LocalInsetBounds().Size
		cs := sa.content.Size()
		cl := sa.content.Location()
		nl := cl
		if cl.Y != 0 && vs.Height > cl.Y+cs.Height {
			nl.Y = math.Min(vs.Height-cs.Height, 0)
		}
		if cl.X != 0 && vs.Width > cl.X+cs.Width {
			nl.X = math.Min(vs.Width-cs.Width, 0)
		}
		if nl != cl {
			sa.content.SetLocation(nl)
		}
	}
}

func (sa *ScrollArea) mouseWheel(evt event.Event) {
	delta := evt.(*event.MouseWheel).Delta()
	if delta.Y != 0 {
		sa.vBar.SetScrolledPosition(sa.ScrolledPosition(false) - delta.Y*sa.LineScrollAmount(false, delta.Y > 0))
	}
	if delta.X != 0 {
		sa.hBar.SetScrolledPosition(sa.ScrolledPosition(true) - delta.X*sa.LineScrollAmount(true, delta.X > 0))
	}
	evt.Finish()
}

func (sa *ScrollArea) focusGained(evt event.Event) {
	if sa.content == nil || sa.content.ID() == evt.Target().ID() {
		sa.view.SetBorder(sa.Theme.FocusBorder)
		sa.SetNeedLayout(true)
		sa.Repaint()
	}
}

func (sa *ScrollArea) focusLost(evt event.Event) {
	if sa.content == nil || sa.content.ID() == evt.Target().ID() {
		sa.view.SetBorder(nil)
		sa.SetNeedLayout(true)
		sa.Repaint()
	}
}

func (sa *ScrollArea) keyDown(evt event.Event) {
	e := evt.(*event.KeyDown)
	switch e.Code() {
	case keys.VirtualKeyUp, keys.VirtualKeyNumPadUp:
		evt.Finish()
		sa.vBar.SetScrolledPosition(sa.ScrolledPosition(false) - sa.LineScrollAmount(false, true))
	case keys.VirtualKeyDown, keys.VirtualKeyNumPadDown:
		evt.Finish()
		sa.vBar.SetScrolledPosition(sa.ScrolledPosition(false) + sa.LineScrollAmount(false, false))
	case keys.VirtualKeyLeft, keys.VirtualKeyNumPadLeft:
		evt.Finish()
		sa.hBar.SetScrolledPosition(sa.ScrolledPosition(true) - sa.LineScrollAmount(true, true))
	case keys.VirtualKeyRight, keys.VirtualKeyNumPadRight:
		evt.Finish()
		sa.hBar.SetScrolledPosition(sa.ScrolledPosition(true) + sa.LineScrollAmount(true, false))
	case keys.VirtualKeyHome, keys.VirtualKeyNumPadHome:
		evt.Finish()
		var bar *scrollbar.ScrollBar
		if e.Modifiers().ShiftDown() {
			bar = sa.hBar
		} else {
			bar = sa.vBar
		}
		bar.SetScrolledPosition(0)
	case keys.VirtualKeyEnd, keys.VirtualKeyNumPadEnd:
		evt.Finish()
		var bar *scrollbar.ScrollBar
		horizontal := e.Modifiers().ShiftDown()
		if horizontal {
			bar = sa.hBar
		} else {
			bar = sa.vBar
		}
		bar.SetScrolledPosition(sa.ContentSize(horizontal))
	case keys.VirtualKeyPageUp, keys.VirtualKeyNumPadPageUp:
		evt.Finish()
		var bar *scrollbar.ScrollBar
		horizontal := e.Modifiers().ShiftDown()
		if horizontal {
			bar = sa.hBar
		} else {
			bar = sa.vBar
		}
		bar.SetScrolledPosition(sa.ScrolledPosition(horizontal) - sa.PageScrollAmount(horizontal, true))
	case keys.VirtualKeyPageDown, keys.VirtualKeyNumPadPageDown:
		evt.Finish()
		var bar *scrollbar.ScrollBar
		horizontal := e.Modifiers().ShiftDown()
		if horizontal {
			bar = sa.hBar
		} else {
			bar = sa.vBar
		}
		bar.SetScrolledPosition(sa.ScrolledPosition(horizontal) + sa.PageScrollAmount(horizontal, false))
	}
}
