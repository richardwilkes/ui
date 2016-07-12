// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

const (
	scrollBarNone scrollBarPart = iota
	scrollBarThumb
	scrollBarLineUp
	scrollBarLineDown
	scrollBarPageUp
	scrollBarPageDown
)

type scrollBarPart int

// Pager objects can provide line and page information for scrolling.
type Pager interface {
	// LineScrollAmount is called to determine how far to scroll in the given direction to reveal
	// a full 'line' of content. A positive value should be returned regardless of the direction,
	// although negative values will behave as if they were positive.
	LineScrollAmount(horizontal, towardsStart bool) float32
	// PageScrollAmount is called to determine how far to scroll in the given direction to reveal
	// a full 'page' of content. A positive value should be returned regardless of the direction,
	// although negative values will behave as if they were positive.
	PageScrollAmount(horizontal, towardsStart bool) float32
}

// Scrollable objects can respond to ScrollBars.
type Scrollable interface {
	Pager
	// ScrolledPosition is called to determine the current position of the Scrollable.
	ScrolledPosition(horizontal bool) float32
	// SetScrolledPosition is called to set the current position of the Scrollable.
	SetScrolledPosition(horizontal bool, position float32)
	// VisibleSize is called to determine the size of the visible portion of the Scrollable.
	VisibleSize(horizontal bool) float32
	// ContentSize is called to determine the total size of the Scrollable.
	ContentSize(horizontal bool) float32
}

// ScrollBar represents a widget for controlling scrolling.
type ScrollBar struct {
	Block
	Theme      *ScrollBarTheme // The theme the scrollbar will use to draw itself.
	Target     Scrollable      // The target of the scrollbar.
	pressed    scrollBarPart
	thumbDown  float32
	sequence   int
	horizontal bool
}

// NewScrollBar creates a new scrollbar.
func NewScrollBar(horizontal bool, target Scrollable) *ScrollBar {
	sb := &ScrollBar{}
	sb.Theme = StdScrollBarTheme
	sb.Target = target
	sb.horizontal = horizontal
	sb.SetSizer(sb)
	sb.SetPaintHandler(sb)
	sb.SetMouseDownHandler(sb)
	sb.SetMouseDraggedHandler(sb)
	sb.SetMouseUpHandler(sb)
	return sb
}

// Sizes implements the Sizer interface.
func (sb *ScrollBar) Sizes(hint Size) (min, pref, max Size) {
	if sb.horizontal {
		min.Width = sb.Theme.Size * 2
		min.Height = sb.Theme.Size
		pref.Width = sb.Theme.Size * 2
		pref.Height = sb.Theme.Size
		max.Width = DefaultLayoutMax
		max.Height = sb.Theme.Size
	} else {
		min.Width = sb.Theme.Size
		min.Height = sb.Theme.Size * 2
		pref.Width = sb.Theme.Size
		pref.Height = sb.Theme.Size * 2
		max.Width = sb.Theme.Size
		max.Height = DefaultLayoutMax
	}
	if border := sb.Border(); border != nil {
		insets := border.Insets()
		min.AddInsets(insets)
		pref.AddInsets(insets)
		max.AddInsets(insets)
	}
	return min, pref, max
}

// OnPaint implements PaintHandler
func (sb *ScrollBar) OnPaint(g Graphics, dirty Rect) {
	bounds := sb.LocalInsetBounds()
	bgColor := sb.baseBackground(scrollBarNone)
	g.SetFillColor(bgColor)
	g.FillRect(bounds)
	g.SetStrokeColor(bgColor.AdjustBrightness(sb.Theme.OutlineAdjustment))
	g.StrokeRect(bounds)
	sb.drawLineButton(g, scrollBarLineUp)
	sb.drawLineButton(g, scrollBarLineDown)
	if sb.pressed == scrollBarPageUp || sb.pressed == scrollBarPageDown {
		bounds = sb.partRect(sb.pressed)
		if !bounds.IsEmpty() {
			if sb.horizontal {
				bounds.Y++
				bounds.Height -= 2
			} else {
				bounds.X++
				bounds.Width -= 2
			}
			g.SetFillColor(sb.baseBackground(sb.pressed))
			g.FillRect(bounds)
		}
	}
	sb.drawThumb(g)
}

// OnMouseDown implements MouseDownHandler
func (sb *ScrollBar) OnMouseDown(where Point, keyModifiers int, which int, clickCount int) bool {
	sb.sequence++
	part := sb.over(where)
	if sb.partEnabled(part) {
		sb.pressed = part
		switch part {
		case scrollBarThumb:
			if sb.horizontal {
				sb.thumbDown = where.X - sb.partRect(part).X
			} else {
				sb.thumbDown = where.Y - sb.partRect(part).Y
			}
		case scrollBarLineUp, scrollBarLineDown, scrollBarPageUp, scrollBarPageDown:
			sb.scheduleRepeat(part)
		}
		sb.Repaint()
	}
	return false
}

// OnMouseDragged implements MouseDraggedHandler
func (sb *ScrollBar) OnMouseDragged(where Point, keyModifiers int) {
	if sb.pressed == scrollBarThumb {
		var pos float32
		rect := sb.partRect(scrollBarLineUp)
		if sb.horizontal {
			pos = where.X - (sb.thumbDown + rect.X + rect.Width - 1)
		} else {
			pos = where.Y - (sb.thumbDown + rect.Y + rect.Height - 1)
		}
		sb.SetScrolledPosition(pos / sb.thumbScale())
	}
}

// OnMouseUp implements MouseUpHandler
func (sb *ScrollBar) OnMouseUp(where Point, keyModifiers int) {
	sb.pressed = scrollBarNone
	sb.Repaint()
}

func (sb *ScrollBar) scheduleRepeat(part scrollBarPart) {
	current := sb.sequence
	switch part {
	case scrollBarLineUp:
		sb.SetScrolledPosition(sb.Target.ScrolledPosition(sb.horizontal) - AbsFloat32(sb.Target.LineScrollAmount(sb.horizontal, true)))
	case scrollBarLineDown:
		sb.SetScrolledPosition(sb.Target.ScrolledPosition(sb.horizontal) + AbsFloat32(sb.Target.LineScrollAmount(sb.horizontal, true)))
	case scrollBarPageUp:
		sb.SetScrolledPosition(sb.Target.ScrolledPosition(sb.horizontal) - AbsFloat32(sb.Target.PageScrollAmount(sb.horizontal, true)))
	case scrollBarPageDown:
		sb.SetScrolledPosition(sb.Target.ScrolledPosition(sb.horizontal) + AbsFloat32(sb.Target.PageScrollAmount(sb.horizontal, true)))
	default:
		return
	}
	InvokeAfter(func() {
		if current == sb.sequence && sb.pressed == part {
			sb.scheduleRepeat(part)
		}
	}, sb.Theme.RepeatDelay)
}

func (sb *ScrollBar) over(where Point) scrollBarPart {
	for i := scrollBarThumb; i <= scrollBarPageDown; i++ {
		rect := sb.partRect(i)
		if rect.Contains(where) {
			return i
		}
	}
	return scrollBarNone
}

func (sb *ScrollBar) thumbScale() float32 {
	var scale float32 = 1
	content := sb.Target.ContentSize(sb.horizontal)
	visible := sb.Target.VisibleSize(sb.horizontal)
	if content-visible > 0 {
		var size float32
		min := sb.Theme.Size * 0.75
		bounds := sb.LocalInsetBounds()
		if sb.horizontal {
			size = bounds.Width
		} else {
			size = bounds.Height
		}
		size -= sb.Theme.Size*2 + 2
		if size > 0 {
			scale = size / content
			visible *= scale
			if visible < min {
				scale = (size + visible - min) / content
			}
		}
	}
	return scale
}

func (sb *ScrollBar) partRect(part scrollBarPart) Rect {
	switch part {
	case scrollBarThumb:
		if sb.Target != nil {
			content := sb.Target.ContentSize(sb.horizontal)
			visible := sb.Target.VisibleSize(sb.horizontal)
			if content-visible > 0 {
				pos := sb.Target.ScrolledPosition(sb.horizontal)
				full := sb.LocalInsetBounds()
				if sb.horizontal {
					full.X += sb.Theme.Size - 1
					full.Width -= sb.Theme.Size*2 - 2
					if full.Width > 0 {
						scale := full.Width / content
						visible *= scale
						min := sb.Theme.Size * 0.75
						if visible < min {
							scale = (full.Width + visible - min) / content
							visible = min
						}
						pos *= scale
						full.X += pos
						full.Width = visible + 1
						return full
					}
				} else {
					full.Y += sb.Theme.Size - 1
					full.Height -= sb.Theme.Size*2 - 2
					if full.Height > 0 {
						scale := full.Height / content
						visible *= scale
						min := sb.Theme.Size * 0.75
						if visible < min {
							scale = (full.Height + visible - min) / content
							visible = min
						}
						pos *= scale
						full.Y += pos
						full.Height = visible + 1
						return full
					}
				}
			}
		}
		return Rect{}
	case scrollBarLineUp:
		bounds := sb.LocalInsetBounds()
		return Rect{Point: Point{X: bounds.X, Y: bounds.Y}, Size: Size{Width: sb.Theme.Size, Height: sb.Theme.Size}}
	case scrollBarLineDown:
		bounds := sb.LocalInsetBounds()
		return Rect{Point: Point{X: bounds.X + bounds.Width - sb.Theme.Size, Y: bounds.Y + bounds.Height - sb.Theme.Size}, Size: Size{Width: sb.Theme.Size, Height: sb.Theme.Size}}
	case scrollBarPageUp:
		top := sb.partRect(scrollBarLineUp)
		thumb := sb.partRect(scrollBarThumb)
		if sb.horizontal {
			x := top.X + top.Width
			return Rect{Point: Point{X: x, Y: top.Y}, Size: Size{Width: thumb.X - x, Height: top.Height}}
		}
		y := top.Y + top.Height
		return Rect{Point: Point{X: top.X, Y: y}, Size: Size{Width: top.Width, Height: thumb.Y - y}}
	case scrollBarPageDown:
		thumb := sb.partRect(scrollBarThumb)
		bottom := sb.partRect(scrollBarLineDown)
		if sb.horizontal {
			x := thumb.X + thumb.Width
			return Rect{Point: Point{X: x, Y: bottom.Y}, Size: Size{Width: bottom.X - x, Height: bottom.Height}}
		}
		y := thumb.Y + thumb.Height
		return Rect{Point: Point{X: bottom.X, Y: y}, Size: Size{Width: bottom.Width, Height: bottom.Y - y}}
	default:
		return Rect{}
	}
}

func (sb *ScrollBar) drawThumb(g Graphics) {
	if sb.Target != nil {
		content := sb.Target.ContentSize(sb.horizontal)
		visible := sb.Target.VisibleSize(sb.horizontal)
		if content-visible > 0 {
			bgColor := sb.baseBackground(scrollBarThumb)
			pos := sb.Target.ScrolledPosition(sb.horizontal)
			full := sb.LocalInsetBounds()
			if sb.horizontal {
				full.X += sb.Theme.Size - 1
				full.Y++
				full.Width -= sb.Theme.Size*2 - 2
				full.Height -= 2
				if full.Width > 0 {
					scale := full.Width / content
					visible *= scale
					min := sb.Theme.Size * 0.75
					if visible < min {
						scale = (full.Width + visible - min) / content
						visible = min
					}
					pos *= scale
					g.Save()
					full.X += pos
					full.Width = visible
					middle := full.X + full.Width/2
					g.ClipRect(full)
					g.DrawLinearGradient(sb.Theme.Gradient(bgColor), middle, full.Y, middle, full.Y+full.Height-1)
					g.Restore()
					g.SetStrokeColor(bgColor.AdjustBrightness(sb.Theme.OutlineAdjustment))
					g.StrokeLine(full.X, full.Y-1, full.X, full.Y+full.Height)
					g.StrokeLine(full.X+full.Width, full.Y-1, full.X+full.Width, full.Y+full.Height)
					g.SetStrokeColor(sb.markColor(scrollBarThumb))
					for i := middle - 2; i < middle+3; i += 2 {
						g.StrokeLine(i, full.Y+2.5, i, full.Y+full.Height-3.5)
					}
				}
			} else {
				full.X++
				full.Y += sb.Theme.Size - 1
				full.Width -= 2
				full.Height -= sb.Theme.Size*2 - 2
				if full.Height > 0 {
					scale := full.Height / content
					visible *= scale
					min := sb.Theme.Size * 0.75
					if visible < min {
						scale = (full.Height + visible - min) / content
						visible = min
					}
					pos *= scale
					g.Save()
					full.Y += pos
					full.Height = visible
					middle := full.Y + full.Height/2
					g.ClipRect(full)
					g.DrawLinearGradient(sb.Theme.Gradient(bgColor), full.X, full.Y+full.Height/2, full.X+full.Width-1, full.Y+full.Height/2)
					g.Restore()
					g.SetStrokeColor(bgColor.AdjustBrightness(sb.Theme.OutlineAdjustment))
					g.StrokeLine(full.X-1, full.Y, full.X+full.Width, full.Y)
					g.StrokeLine(full.X-1, full.Y+full.Height, full.X+full.Width, full.Y+full.Height)
					g.SetStrokeColor(sb.markColor(scrollBarThumb))
					for i := middle - 2; i < middle+3; i += 2 {
						g.StrokeLine(full.X+2.5, i, full.X+full.Width-3.5, i)
					}
				}
			}
		}
	}
}

func (sb *ScrollBar) drawLineButton(g Graphics, linePart scrollBarPart) {
	bounds := sb.partRect(linePart)
	g.Save()
	g.ClipRect(bounds)
	bgColor := sb.baseBackground(linePart)
	g.DrawLinearGradient(sb.Theme.Gradient(bgColor), bounds.X, bounds.Y, bounds.X, bounds.Y+bounds.Height)
	g.Restore()
	g.SetStrokeColor(bgColor.AdjustBrightness(sb.Theme.OutlineAdjustment))
	g.StrokeRect(bounds)
	bounds.InsetUniform(1)
	if sb.horizontal {
		triHeight := (bounds.Width * 0.75)
		triWidth := triHeight / 2
		g.BeginPath()
		left := bounds.X + (bounds.Width-triWidth)/2
		right := left + triWidth
		top := bounds.Y + (bounds.Height-triHeight)/2
		bottom := top + triHeight
		if linePart == scrollBarLineUp {
			left, right = right, left
		}
		g.MoveTo(left, top)
		g.LineTo(left, bottom)
		g.LineTo(right, top+(bottom-top)/2)
	} else {
		triWidth := (bounds.Height * 0.75)
		triHeight := triWidth / 2
		g.BeginPath()
		left := bounds.X + (bounds.Width-triWidth)/2
		right := left + triWidth
		top := bounds.Y + (bounds.Height-triHeight)/2
		bottom := top + triHeight
		if linePart == scrollBarLineUp {
			top, bottom = bottom, top
		}
		g.MoveTo(left, top)
		g.LineTo(right, top)
		g.LineTo(left+(right-left)/2, bottom)
	}
	g.ClosePath()
	g.SetFillColor(sb.markColor(linePart))
	g.FillPath()
}

func (sb *ScrollBar) baseBackground(part scrollBarPart) Color {
	switch {
	case !sb.Enabled():
		return sb.Theme.Background.AdjustBrightness(sb.Theme.DisabledAdjustment)
	case part != scrollBarNone && sb.pressed == part:
		return sb.Theme.BackgroundWhenPressed
	case sb.Focused():
		return sb.Theme.Background.Blend(KeyboardFocusColor, 0.5)
	default:
		return sb.Theme.Background
	}
}

func (sb *ScrollBar) markColor(part scrollBarPart) Color {
	if sb.partEnabled(part) {
		if sb.baseBackground(part).Luminance() > 0.65 {
			return sb.Theme.MarkWhenLight
		}
		return sb.Theme.MarkWhenDark
	}
	return sb.Theme.MarkWhenDisabled
}

func (sb *ScrollBar) partEnabled(part scrollBarPart) bool {
	if sb.Enabled() && sb.Target != nil {
		switch part {
		case scrollBarLineUp, scrollBarPageUp:
			return sb.Target.ScrolledPosition(sb.horizontal) > 0
		case scrollBarLineDown, scrollBarPageDown:
			return sb.Target.ScrolledPosition(sb.horizontal) < sb.Target.ContentSize(sb.horizontal)-sb.Target.VisibleSize(sb.horizontal)
		case scrollBarThumb:
			pos := sb.Target.ScrolledPosition(sb.horizontal)
			return pos > 0 || pos < sb.Target.ContentSize(sb.horizontal)-sb.Target.VisibleSize(sb.horizontal)
		default:
		}
	}
	return false
}

// SetScrolledPosition attempts to set the current scrolled position of this ScrollBar to the
// specified value. The value will be clipped to the available range. If no target has been set,
// then nothing will happen.
func (sb *ScrollBar) SetScrolledPosition(position float32) {
	if sb.Target != nil {
		position = MaxFloat32(MinFloat32(position, sb.Target.ContentSize(sb.horizontal)-sb.Target.VisibleSize(sb.horizontal)), 0)
		if sb.Target.ScrolledPosition(sb.horizontal) != position {
			sb.Target.SetScrolledPosition(sb.horizontal, position)
			sb.Repaint()
		}
	}
}