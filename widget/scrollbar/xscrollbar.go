package scrollbar

import (
	"fmt"
	"math"
	"time"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/layout"
	"github.com/richardwilkes/ui/widget"
)

const (
	none part = iota
	thumb
	lineUp
	lineDown
	pageUp
	pageDown
)

type part int

// ScrollBar represents a widget for controlling scrolling.
type ScrollBar struct {
	widget.Block
	Target     Scrollable // The target of the scrollbar.
	Theme      *Theme     // The theme the scrollbar will use to draw itself.
	pressed    part
	sequence   int
	thumbDown  float64
	horizontal bool
}

// New creates a new scrollbar.
func New(horizontal bool, target Scrollable) *ScrollBar {
	sb := &ScrollBar{Theme: StdTheme, Target: target, horizontal: horizontal}
	sb.InitTypeAndID(sb)
	sb.Describer = func() string { return fmt.Sprintf("ScrollBar #%d", sb.ID()) }
	sb.SetSizer(sb)
	handlers := sb.EventHandlers()
	handlers.Add(event.PaintType, sb.paint)
	handlers.Add(event.MouseDownType, sb.mouseDown)
	handlers.Add(event.MouseDraggedType, sb.mouseDragged)
	handlers.Add(event.MouseUpType, sb.mouseUp)
	return sb
}

// Sizes implements the Sizer interface.
func (sb *ScrollBar) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	if sb.horizontal {
		min.Width = sb.Theme.Size * 2
		min.Height = sb.Theme.Size
		pref.Width = sb.Theme.Size * 2
		pref.Height = sb.Theme.Size
		max.Width = layout.DefaultMax
		max.Height = sb.Theme.Size
	} else {
		min.Width = sb.Theme.Size
		min.Height = sb.Theme.Size * 2
		pref.Width = sb.Theme.Size
		pref.Height = sb.Theme.Size * 2
		max.Width = sb.Theme.Size
		max.Height = layout.DefaultMax
	}
	if border := sb.Border(); border != nil {
		insets := border.Insets()
		min.AddInsets(insets)
		pref.AddInsets(insets)
		max.AddInsets(insets)
	}
	return min, pref, max
}

func (sb *ScrollBar) paint(evt event.Event) {
	bounds := sb.LocalInsetBounds()
	if sb.horizontal {
		bounds.Height = sb.Theme.Size
	} else {
		bounds.Width = sb.Theme.Size
	}
	bgColor := sb.baseBackground(none)
	gc := evt.(*event.Paint).GC()
	gc.SetColor(bgColor)
	gc.FillRect(bounds)
	gc.SetColor(bgColor.AdjustBrightness(sb.Theme.OutlineAdjustment))
	gc.StrokeRect(bounds)
	sb.drawLineButton(gc, lineDown)
	if sb.pressed == pageUp || sb.pressed == pageDown {
		bounds = sb.partRect(sb.pressed)
		if !bounds.IsEmpty() {
			if sb.horizontal {
				bounds.Y++
				bounds.Height -= 2
			} else {
				bounds.X++
				bounds.Width -= 2
			}
			gc.SetColor(sb.baseBackground(sb.pressed))
			gc.FillRect(bounds)
		}
	}
	sb.drawThumb(gc)
	sb.drawLineButton(gc, lineUp)
}

func (sb *ScrollBar) mouseDown(evt event.Event) {
	sb.sequence++
	where := sb.FromWindow(evt.(*event.MouseDown).Where())
	what := sb.over(where)
	if sb.partEnabled(what) {
		sb.pressed = what
		switch what {
		case thumb:
			if sb.horizontal {
				sb.thumbDown = where.X - sb.partRect(what).X
			} else {
				sb.thumbDown = where.Y - sb.partRect(what).Y
			}
		case lineUp, lineDown, pageUp, pageDown:
			sb.scheduleRepeat(what, sb.Theme.InitialRepeatDelay)
		}
		sb.Repaint()
	}
}

func (sb *ScrollBar) mouseDragged(evt event.Event) {
	if sb.pressed == thumb {
		var pos float64
		rect := sb.partRect(lineUp)
		where := sb.FromWindow(evt.(*event.MouseDragged).Where())
		if sb.horizontal {
			pos = where.X - (sb.thumbDown + rect.X + rect.Width - 1)
		} else {
			pos = where.Y - (sb.thumbDown + rect.Y + rect.Height - 1)
		}
		sb.SetScrolledPosition(pos / sb.thumbScale())
	}
}

func (sb *ScrollBar) mouseUp(evt event.Event) {
	sb.pressed = none
	sb.Repaint()
}

func (sb *ScrollBar) scheduleRepeat(which part, delay time.Duration) {
	window := sb.Window()
	if window.Valid() {
		current := sb.sequence
		switch which {
		case lineUp:
			sb.SetScrolledPosition(sb.Target.ScrolledPosition(sb.horizontal) - math.Abs(sb.Target.LineScrollAmount(sb.horizontal, true)))
		case lineDown:
			sb.SetScrolledPosition(sb.Target.ScrolledPosition(sb.horizontal) + math.Abs(sb.Target.LineScrollAmount(sb.horizontal, false)))
		case pageUp:
			sb.SetScrolledPosition(sb.Target.ScrolledPosition(sb.horizontal) - math.Abs(sb.Target.PageScrollAmount(sb.horizontal, true)))
		case pageDown:
			sb.SetScrolledPosition(sb.Target.ScrolledPosition(sb.horizontal) + math.Abs(sb.Target.PageScrollAmount(sb.horizontal, false)))
		default:
			return
		}
		window.InvokeAfter(func() {
			if current == sb.sequence && sb.pressed == which {
				sb.scheduleRepeat(which, sb.Theme.RepeatDelay)
			}
		}, delay)
	}
}

func (sb *ScrollBar) over(where geom.Point) part {
	for i := thumb; i <= pageDown; i++ {
		rect := sb.partRect(i)
		if rect.ContainsPoint(where) {
			return i
		}
	}
	return none
}

func (sb *ScrollBar) thumbScale() float64 {
	var scale float64 = 1
	content := sb.Target.ContentSize(sb.horizontal)
	visible := sb.Target.VisibleSize(sb.horizontal)
	if content-visible > 0 {
		var size float64
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

func (sb *ScrollBar) partRect(which part) geom.Rect {
	var result geom.Rect
	switch which {
	case thumb:
		if sb.Target != nil {
			content := sb.Target.ContentSize(sb.horizontal)
			visible := sb.Target.VisibleSize(sb.horizontal)
			if content-visible > 0 {
				pos := sb.Target.ScrolledPosition(sb.horizontal)
				full := sb.LocalInsetBounds()
				if sb.horizontal {
					full.X += sb.Theme.Size - 1
					full.Width -= sb.Theme.Size*2 - 2
					full.Height = sb.Theme.Size
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
						result = full
					}
				} else {
					full.Y += sb.Theme.Size - 1
					full.Height -= sb.Theme.Size*2 - 2
					full.Width = sb.Theme.Size
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
						result = full
					}
				}
			}
		}
	case lineUp:
		result = sb.LocalInsetBounds()
		result.Width = sb.Theme.Size
		result.Height = sb.Theme.Size
	case lineDown:
		result = sb.LocalInsetBounds()
		if sb.horizontal {
			result.X += result.Width - sb.Theme.Size
		} else {
			result.Y += result.Height - sb.Theme.Size
		}
		result.Width = sb.Theme.Size
		result.Height = sb.Theme.Size
	case pageUp:
		result = sb.partRect(lineUp)
		thumb := sb.partRect(thumb)
		if sb.horizontal {
			result.X += result.Width
			result.Width = thumb.X - result.X
		} else {
			result.Y += result.Height
			result.Height = thumb.Y - result.Y
		}
	case pageDown:
		result = sb.partRect(lineDown)
		thumb := sb.partRect(thumb)
		if sb.horizontal {
			x := thumb.X + thumb.Width
			result.Width = result.X - x
			result.X = x
		} else {
			y := thumb.Y + thumb.Height
			result.Height = result.Y - y
			result.Y = y
		}
	}
	return result
}

func (sb *ScrollBar) drawThumb(g *draw.Graphics) {
	bounds := sb.partRect(thumb)
	if !bounds.IsEmpty() {
		bgColor := sb.baseBackground(thumb)
		g.Rect(bounds)
		var paint *draw.Paint
		if sb.horizontal {
			paint = draw.NewLinearGradientPaint(sb.Theme.Gradient(bgColor), bounds.X, bounds.Y, bounds.X, bounds.Y+bounds.Height)
		} else {
			paint = draw.NewLinearGradientPaint(sb.Theme.Gradient(bgColor), bounds.X, bounds.Y, bounds.X+bounds.Width, bounds.Y)
		}
		g.SetPaint(paint)
		g.FillPath()
		paint.Dispose()
		g.SetColor(bgColor.AdjustBrightness(sb.Theme.OutlineAdjustment))
		g.StrokeRect(bounds)
		g.SetColor(sb.markColor(thumb))
		var v0, v1, v2 float64
		if sb.horizontal {
			v0 = math.Floor(bounds.X + bounds.Width/2)
			d := math.Ceil(bounds.Height * 0.2)
			v1 = bounds.Y + d
			v2 = bounds.Y + bounds.Height - (d + 1)
		} else {
			v0 = math.Floor(bounds.Y + bounds.Height/2)
			d := math.Ceil(bounds.Width * 0.2)
			v1 = bounds.X + d
			v2 = bounds.X + bounds.Width - (d + 1)
		}
		for i := -1; i < 2; i++ {
			if sb.horizontal {
				x := v0 + float64(i*2)
				g.StrokeLine(x, v1, x, v2)
			} else {
				y := v0 + float64(i*2)
				g.StrokeLine(v1, y, v2, y)
			}
		}
	}
}

func (sb *ScrollBar) drawLineButton(g *draw.Graphics, linePart part) {
	bounds := sb.partRect(linePart)
	g.Save()
	g.Rect(bounds)
	bgColor := sb.baseBackground(linePart)
	paint := draw.NewLinearGradientPaint(sb.Theme.Gradient(bgColor), bounds.X, bounds.Y, bounds.X, bounds.Y+bounds.Height)
	g.SetPaint(paint)
	g.FillPath()
	paint.Dispose()
	g.SetColor(bgColor.AdjustBrightness(sb.Theme.OutlineAdjustment))
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
		if linePart == lineUp {
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
		if linePart == lineUp {
			top, bottom = bottom, top
		}
		g.MoveTo(left, top)
		g.LineTo(right, top)
		g.LineTo(left+(right-left)/2, bottom)
	}
	g.ClosePath()
	g.SetColor(sb.markColor(linePart))
	g.FillPath()
	g.Restore()
}

func (sb *ScrollBar) baseBackground(which part) color.Color {
	switch {
	case !sb.Enabled():
		return sb.Theme.Background.AdjustBrightness(sb.Theme.DisabledAdjustment)
	case which != none && sb.pressed == which:
		return sb.Theme.BackgroundWhenPressed
	case sb.Focused():
		return sb.Theme.Background.Blend(color.KeyboardFocus, 0.5)
	default:
		return sb.Theme.Background
	}
}

func (sb *ScrollBar) markColor(which part) color.Color {
	if sb.partEnabled(which) {
		if sb.baseBackground(which).Luminance() > 0.65 {
			return sb.Theme.MarkWhenLight
		}
		return sb.Theme.MarkWhenDark
	}
	return sb.Theme.MarkWhenDisabled
}

func (sb *ScrollBar) partEnabled(which part) bool {
	if sb.Enabled() && sb.Target != nil {
		switch which {
		case lineUp, pageUp:
			return sb.Target.ScrolledPosition(sb.horizontal) > 0
		case lineDown, pageDown:
			return sb.Target.ScrolledPosition(sb.horizontal) < sb.Target.ContentSize(sb.horizontal)-sb.Target.VisibleSize(sb.horizontal)
		case thumb:
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
func (sb *ScrollBar) SetScrolledPosition(position float64) {
	if sb.Target != nil {
		position = math.Max(math.Min(position, sb.Target.ContentSize(sb.horizontal)-sb.Target.VisibleSize(sb.horizontal)), 0)
		if sb.Target.ScrolledPosition(sb.horizontal) != position {
			sb.Target.SetScrolledPosition(sb.horizontal, position)
			sb.Repaint()
		}
	}
}
