package flow

import (
	"math"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/layout"
)

// Layout lays out the children of its widget left-to-right, then
// top-to-bottom at their preferred sizes, if possible.
type Layout struct {
	HSpacing float64 // Horizontal spacing between columns
	VSpacing float64 // Vertical spacing between rows
	VCenter  bool    // Center widgets vertically in their row if true
	widget   ui.Widget
}

// New creates a new Flow layout and sets it on the widget.
func New(widget ui.Widget) *Layout {
	layout := &Layout{widget: widget}
	widget.SetLayout(layout)
	return layout
}

// Sizes implements the Layout interface.
func (l *Layout) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	if l.HSpacing < 0 {
		l.HSpacing = 0
	}
	if l.VSpacing < 0 {
		l.VSpacing = 0
	}
	var insets geom.Insets
	if border := l.widget.Border(); border != nil {
		insets = border.Insets()
	}
	if hint.Width < 0 {
		hint.Width = math.MaxFloat32
	}
	if hint.Height < 0 {
		hint.Height = math.MaxFloat32
	}
	width := hint.Width - (insets.Left + insets.Right)
	pt := geom.Point{X: insets.Left, Y: insets.Top}
	result := geom.Size{Width: pt.Y, Height: pt.Y}
	availWidth := width
	availHeight := hint.Height - (insets.Top + insets.Bottom)
	var maxHeight float64
	var largestChildMin geom.Size
	noHint := geom.Size{Width: layout.NoHint, Height: layout.NoHint}
	for _, child := range l.widget.Children() {
		min, pref, _ := ui.Sizes(child, noHint)
		if largestChildMin.Width < min.Width {
			largestChildMin.Width = min.Width
		}
		if largestChildMin.Height < min.Height {
			largestChildMin.Height = min.Height
		}
		if pref.Width > availWidth {
			if min.Width <= availWidth {
				pref.Width = availWidth
			} else if pt.X == insets.Left {
				pref.Width = min.Width
			} else {
				pt.X = insets.Left
				pt.Y += maxHeight + l.VSpacing
				availWidth = width
				availHeight -= maxHeight + l.VSpacing
				maxHeight = 0
				if pref.Width > availWidth {
					if min.Width <= availWidth {
						pref.Width = availWidth
					} else {
						pref.Width = min.Width
					}
				}
			}
			savedWidth := pref.Width
			min, pref, _ = ui.Sizes(child, geom.Size{Width: pref.Width, Height: layout.NoHint})
			pref.Width = savedWidth
			if pref.Height > availHeight {
				if min.Height <= availHeight {
					pref.Height = availHeight
				} else {
					pref.Height = min.Height
				}
			}
		}
		extent := pt.X + pref.Width
		if result.Width < extent {
			result.Width = extent
		}
		extent = pt.Y + pref.Height
		if result.Height < extent {
			result.Height = extent
		}
		if maxHeight < pref.Height {
			maxHeight = pref.Height
		}
		availWidth -= pref.Width + l.HSpacing
		if availWidth <= 0 {
			pt.X = insets.Left
			pt.Y += maxHeight + l.VSpacing
			availWidth = width
			availHeight -= maxHeight + l.VSpacing
			maxHeight = 0
		} else {
			pt.X += pref.Width + l.HSpacing
		}
	}
	result.Width += insets.Right
	result.Height += insets.Bottom
	largestChildMin.Width += insets.Left + insets.Right
	largestChildMin.Height += insets.Top + insets.Bottom
	return largestChildMin, result, layout.DefaultMaxSize(result)
}

// Layout implements the Layout interface.
func (l *Layout) Layout() {
	var insets geom.Insets
	if border := l.widget.Border(); border != nil {
		insets = border.Insets()
	}
	size := l.widget.Bounds().Size
	width := size.Width - (insets.Left + insets.Right)
	pt := geom.Point{X: insets.Left, Y: insets.Top}
	availWidth := width
	availHeight := size.Height - (insets.Top + insets.Bottom)
	var maxHeight float64
	noHint := geom.Size{Width: layout.NoHint, Height: layout.NoHint}
	children := l.widget.Children()
	rects := make([]geom.Rect, len(children))
	start := 0
	for i, child := range children {
		min, pref, _ := ui.Sizes(child, noHint)
		if pref.Width > availWidth {
			if min.Width <= availWidth {
				pref.Width = availWidth
			} else if pt.X == insets.Left {
				pref.Width = min.Width
			} else {
				pt.X = insets.Left
				pt.Y += maxHeight + l.VSpacing
				availWidth = width
				availHeight -= maxHeight + l.VSpacing
				if i > start {
					l.applyRects(children[start:i], rects[start:i], maxHeight)
					start = i
				}
				maxHeight = 0
				if pref.Width > availWidth {
					if min.Width <= availWidth {
						pref.Width = availWidth
					} else {
						pref.Width = min.Width
					}
				}
			}
			savedWidth := pref.Width
			min, pref, _ = ui.Sizes(child, geom.Size{Width: pref.Width, Height: layout.NoHint})
			pref.Width = savedWidth
			if pref.Height > availHeight {
				if min.Height <= availHeight {
					pref.Height = availHeight
				} else {
					pref.Height = min.Height
				}
			}
		}
		rects[i] = geom.Rect{Point: pt, Size: pref}
		if maxHeight < pref.Height {
			maxHeight = pref.Height
		}
		availWidth -= pref.Width + l.HSpacing
		if availWidth <= 0 {
			pt.X = insets.Left
			pt.Y += maxHeight + l.VSpacing
			availWidth = width
			availHeight -= maxHeight + l.VSpacing
			l.applyRects(children[start:i+1], rects[start:i+1], maxHeight)
			start = i + 1
			maxHeight = 0
		} else {
			pt.X += pref.Width + l.HSpacing
		}
	}
	for i, child := range children {
		// RAW: Implement
		//if l.vCenter {
		//}
		child.SetBounds(rects[i])
	}
}

func (l *Layout) applyRects(children []ui.Widget, rects []geom.Rect, maxHeight float64) {
	for i, child := range children {
		if l.VCenter {
			if rects[i].Height < maxHeight {
				rects[i].Y += (maxHeight - rects[i].Height) / 2
			}
		}
		child.SetBounds(rects[i])
	}
}
