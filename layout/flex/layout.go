package flex

import (
	"math"

	"github.com/richardwilkes/toolbox/xmath"
	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/draw/align"
	"github.com/richardwilkes/ui/layout"
)

// Layout lays out the children of its widget based on the Data assigned to
// each child.
type Layout struct {
	Columns      int             // Number of columns
	HSpacing     float64         // Horizontal spacing between columns
	VSpacing     float64         // Vertical spacing between rows
	HAlign       align.Alignment // Horizontal alignment of the widget within its space
	VAlign       align.Alignment // Vertical alignment of the widget within its space
	EqualColumns bool            // Each column uses the same amount of horizontal space if true
	widget       ui.Widget
	rows         int
}

// NewLayout creates a new Flex layout and sets it on the widget.
func NewLayout(widget ui.Widget) *Layout {
	layout := &Layout{
		Columns:  1,
		HSpacing: 4,
		VSpacing: 2,
		HAlign:   align.Start,
		VAlign:   align.Start,
		widget:   widget,
	}
	widget.SetLayout(layout)
	return layout
}

// Sizes implements the Layout interface.
func (l *Layout) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	min = l.layout(geom.Point{}, layout.NoHintSize, false, true)
	pref = l.layout(geom.Point{}, layout.NoHintSize, false, false)
	if border := l.widget.Border(); border != nil {
		insets := border.Insets()
		min.AddInsets(insets)
		pref.AddInsets(insets)
	}
	return min, pref, layout.DefaultMaxSize(pref)
}

// Layout implements the Layout interface.
func (l *Layout) Layout() {
	var insets geom.Insets
	if border := l.widget.Border(); border != nil {
		insets = border.Insets()
	}
	hint := l.widget.Bounds().Size
	hint.SubtractInsets(insets)
	l.layout(geom.Point{X: insets.Left, Y: insets.Top}, hint, true, false)
}

func (l *Layout) layout(location geom.Point, hint geom.Size, move, useMinimumSize bool) geom.Size {
	var totalSize geom.Size
	if l.Columns > 0 {
		children := l.prepChildren(useMinimumSize)
		if len(children) > 0 {
			if l.HSpacing < 0 {
				l.HSpacing = 0
			}
			if l.VSpacing < 0 {
				l.VSpacing = 0
			}
			grid := l.buildGrid(children)
			widths := l.adjustColumnWidths(hint.Width, grid)
			l.wrap(hint.Width, grid, widths, useMinimumSize)
			heights := l.adjustRowHeights(hint.Height, grid)
			totalSize.Width += l.HSpacing * float64(l.Columns-1)
			totalSize.Height += l.VSpacing * float64(l.rows-1)
			for i := 0; i < l.Columns; i++ {
				totalSize.Width += widths[i]
			}
			for i := 0; i < l.rows; i++ {
				totalSize.Height += heights[i]
			}
			if move {
				if totalSize.Width < hint.Width {
					if l.HAlign == align.Middle {
						location.X += xmath.Round((hint.Width - totalSize.Width) / 2)
					} else if l.HAlign == align.End {
						location.X += hint.Width - totalSize.Width
					}
				}
				if totalSize.Height < hint.Height {
					if l.VAlign == align.Middle {
						location.Y += xmath.Round((hint.Height - totalSize.Height) / 2)
					} else if l.VAlign == align.End {
						location.Y += hint.Height - totalSize.Height
					}
				}
				l.positionChildren(location, grid, widths, heights)
			}
		}
	}
	return totalSize
}

func (l *Layout) prepChildren(useMinimumSize bool) []ui.Widget {
	children := l.widget.Children()
	for _, child := range children {
		getDataFromWidget(child).computeCacheSize(child, layout.NoHintSize, useMinimumSize)
	}
	return children
}

func getDataFromWidget(w ui.Widget) *Data {
	if data, ok := w.LayoutData().(*Data); ok {
		return data
	}
	data := NewData()
	w.SetLayoutData(data)
	return data
}

func (l *Layout) buildGrid(children []ui.Widget) [][]ui.Widget {
	var grid [][]ui.Widget
	var row, column int
	l.rows = 0
	for _, child := range children {
		data := getDataFromWidget(child)
		hSpan := xmath.MaxInt(1, xmath.MinInt(data.HSpan, l.Columns))
		vSpan := xmath.MaxInt(1, data.VSpan)
		for {
			lastRow := row + vSpan
			if lastRow >= len(grid) {
				grid = append(grid, make([]ui.Widget, l.Columns))
			}
			for column < l.Columns && grid[row][column] != nil {
				column++
			}
			endCount := column + hSpan
			if endCount <= l.Columns {
				index := column
				for index < endCount && grid[row][index] == nil {
					index++
				}
				if index == endCount {
					break
				}
				column = index
			}
			if column+hSpan >= l.Columns {
				column = 0
				row++
			}
		}
		for j := 0; j < vSpan; j++ {
			pos := row + j
			for k := 0; k < hSpan; k++ {
				grid[pos][column+k] = child
			}
		}
		l.rows = xmath.MaxInt(l.rows, row+vSpan)
		column += hSpan
	}
	return grid
}

func (l *Layout) adjustColumnWidths(width float64, grid [][]ui.Widget) []float64 {
	availableWidth := width - l.HSpacing*float64(l.Columns-1)
	expandCount := 0
	widths := make([]float64, l.Columns)
	minWidths := make([]float64, l.Columns)
	expandColumn := make([]bool, l.Columns)
	for j := 0; j < l.Columns; j++ {
		for i := 0; i < l.rows; i++ {
			data := l.getData(grid, i, j, true)
			if data != nil {
				hSpan := xmath.MaxInt(1, xmath.MinInt(data.HSpan, l.Columns))
				if hSpan == 1 {
					w := data.cacheSize.Width
					if widths[j] < w {
						widths[j] = w
					}
					if data.HGrab {
						if !expandColumn[j] {
							expandCount++
						}
						expandColumn[j] = true
					}
					minimumWidth := data.minCacheSize.Width
					if !data.HGrab || minimumWidth != 0 {
						if minimumWidth == layout.NoHint {
							w = data.cacheSize.Width
						} else {
							w = minimumWidth
						}
						if minWidths[j] < w {
							minWidths[j] = w
						}
					}
				}
			}
		}
		for i := 0; i < l.rows; i++ {
			data := l.getData(grid, i, j, false)
			if data != nil {
				hSpan := xmath.MaxInt(1, xmath.MinInt(data.HSpan, l.Columns))
				if hSpan > 1 {
					var spanWidth, spanMinWidth float64
					spanExpandCount := 0
					for k := 0; k < hSpan; k++ {
						spanWidth += widths[j-k]
						spanMinWidth += minWidths[j-k]
						if expandColumn[j-k] {
							spanExpandCount++
						}
					}
					if data.HGrab && spanExpandCount == 0 {
						expandCount++
						expandColumn[j] = true
					}
					w := data.cacheSize.Width - spanWidth - float64(hSpan-1)*l.HSpacing
					if w > 0 {
						if l.EqualColumns {
							equalWidth := math.Floor((w + spanWidth) / float64(hSpan))
							for k := 0; k < hSpan; k++ {
								if widths[j-k] < equalWidth {
									widths[j-k] = equalWidth
								}
							}
						} else {
							l.apportionExtra(w, j, spanExpandCount, hSpan, expandColumn, widths)
						}
					}
					minimumWidth := data.minCacheSize.Width
					if !data.HGrab || minimumWidth != 0 {
						if !data.HGrab || minimumWidth == layout.NoHint {
							w = data.cacheSize.Width
						} else {
							w = minimumWidth
						}
						w -= spanMinWidth + float64(hSpan-1)*l.HSpacing
						if w > 0 {
							l.apportionExtra(w, j, spanExpandCount, hSpan, expandColumn, minWidths)
						}
					}
				}
			}
		}
	}
	if l.EqualColumns {
		var minColumnWidth, columnWidth float64
		for i := 0; i < l.Columns; i++ {
			if minColumnWidth < minWidths[i] {
				minColumnWidth = minWidths[i]
			}
			if columnWidth < widths[i] {
				columnWidth = widths[i]
			}
		}
		if width != layout.NoHint && expandCount != 0 {
			columnWidth = math.Max(minColumnWidth, math.Floor(availableWidth/float64(l.Columns)))
		}
		for i := 0; i < l.Columns; i++ {
			expandColumn[i] = expandCount > 0
			widths[i] = columnWidth
		}
	} else {
		if width != layout.NoHint && expandCount > 0 {
			var totalWidth float64
			for i := 0; i < l.Columns; i++ {
				totalWidth += widths[i]
			}
			c := expandCount
			for math.Abs(totalWidth-availableWidth) > 0.01 {
				delta := (availableWidth - totalWidth) / float64(c)
				for j := 0; j < l.Columns; j++ {
					if expandColumn[j] {
						if widths[j]+delta > minWidths[j] {
							widths[j] += delta
						} else {
							widths[j] = minWidths[j]
							expandColumn[j] = false
							c--
						}
					}
				}
				for j := 0; j < l.Columns; j++ {
					for i := 0; i < l.rows; i++ {
						data := l.getData(grid, i, j, false)
						if data != nil {
							hSpan := xmath.MaxInt(1, xmath.MinInt(data.HSpan, l.Columns))
							if hSpan > 1 {
								minimumWidth := data.minCacheSize.Width
								if !data.HGrab || minimumWidth != 0 {
									var spanWidth float64
									spanExpandCount := 0
									for k := 0; k < hSpan; k++ {
										spanWidth += widths[j-k]
										if expandColumn[j-k] {
											spanExpandCount++
										}
									}
									var w float64
									if !data.HGrab || minimumWidth == layout.NoHint {
										w = data.cacheSize.Width
									} else {
										w = minimumWidth
									}
									w -= spanWidth + float64(hSpan-1)*l.HSpacing
									if w > 0 {
										l.apportionExtra(w, j, spanExpandCount, hSpan, expandColumn, widths)
									}
								}
							}
						}
					}
				}
				if c == 0 {
					break
				}
				totalWidth = 0
				for i := 0; i < l.Columns; i++ {
					totalWidth += widths[i]
				}
			}
		}
	}
	return widths
}

func (l *Layout) apportionExtra(extra float64, base, count, span int, expand []bool, values []float64) {
	if count == 0 {
		values[base] += extra
	} else {
		extraInt := int(math.Floor(extra))
		delta := extraInt / count
		remainder := extraInt - delta*count
		for i := 0; i < span; i++ {
			j := base - i
			if expand[j] {
				values[j] += float64(delta)
			}
		}
		for remainder > 0 {
			for i := 0; i < span; i++ {
				j := base - i
				if expand[j] {
					values[j]++
					remainder--
					if remainder == 0 {
						break
					}
				}
			}
		}
	}
}

func (l *Layout) getData(grid [][]ui.Widget, row, column int, first bool) *Data {
	block := grid[row][column]
	if block != nil {
		data := getDataFromWidget(block)
		hSpan := xmath.MaxInt(1, xmath.MinInt(data.HSpan, l.Columns))
		vSpan := xmath.MaxInt(1, data.VSpan)
		var i, j int
		if first {
			i = row + vSpan - 1
			j = column + hSpan - 1
		} else {
			i = row - vSpan + 1
			j = column - hSpan + 1
		}
		if i >= 0 && i < l.rows {
			if j >= 0 && j < l.Columns {
				if block == grid[i][j] {
					return data
				}
			}
		}
	}
	return nil
}

func (l *Layout) wrap(width float64, grid [][]ui.Widget, widths []float64, useMinimumSize bool) {
	if width != layout.NoHint {
		for j := 0; j < l.Columns; j++ {
			for i := 0; i < l.rows; i++ {
				data := l.getData(grid, i, j, false)
				if data != nil {
					if data.SizeHint.Height == layout.NoHint {
						hSpan := xmath.MaxInt(1, xmath.MinInt(data.HSpan, l.Columns))
						var currentWidth float64
						for k := 0; k < hSpan; k++ {
							currentWidth += widths[j-k]
						}
						currentWidth += float64(hSpan-1) * l.HSpacing
						if currentWidth != data.cacheSize.Width && data.HAlign == align.Fill || data.cacheSize.Width > currentWidth {
							data.computeCacheSize(grid[i][j], geom.Size{Width: math.Max(data.minCacheSize.Width, currentWidth), Height: layout.NoHint}, useMinimumSize)
							minimumHeight := data.MinSize.Height
							if data.VGrab && minimumHeight > 0 && data.cacheSize.Height < minimumHeight {
								data.cacheSize.Height = minimumHeight
							}
						}
					}
				}
			}
		}
	}
}

func (l *Layout) adjustRowHeights(height float64, grid [][]ui.Widget) []float64 {
	availableHeight := height - l.VSpacing*float64(l.rows-1)
	expandCount := 0
	heights := make([]float64, l.rows)
	minHeights := make([]float64, l.rows)
	expandRow := make([]bool, l.rows)
	for i := 0; i < l.rows; i++ {
		for j := 0; j < l.Columns; j++ {
			data := l.getData(grid, i, j, true)
			if data != nil {
				vSpan := xmath.MaxInt(1, xmath.MinInt(data.VSpan, l.rows))
				if vSpan == 1 {
					h := data.cacheSize.Height
					if heights[i] < h {
						heights[i] = h
					}
					if data.VGrab {
						if !expandRow[i] {
							expandCount++
						}
						expandRow[i] = true
					}
					minimumHeight := data.MinSize.Height
					if !data.VGrab || minimumHeight != 0 {
						var h float64
						if !data.VGrab || minimumHeight == layout.NoHint {
							h = data.minCacheSize.Height
						} else {
							h = minimumHeight
						}
						if minHeights[i] < h {
							minHeights[i] = h
						}
					}
				}
			}
		}
		for j := 0; j < l.Columns; j++ {
			data := l.getData(grid, i, j, false)
			if data != nil {
				vSpan := xmath.MaxInt(1, xmath.MinInt(data.VSpan, l.rows))
				if vSpan > 1 {
					var spanHeight, spanMinHeight float64
					spanExpandCount := 0
					for k := 0; k < vSpan; k++ {
						spanHeight += heights[i-k]
						spanMinHeight += minHeights[i-k]
						if expandRow[i-k] {
							spanExpandCount++
						}
					}
					if data.VGrab && spanExpandCount == 0 {
						expandCount++
						expandRow[i] = true
					}
					h := data.cacheSize.Height - spanHeight - float64(vSpan-1)*l.VSpacing
					if h > 0 {
						if spanExpandCount == 0 {
							heights[i] += h
						} else {
							delta := h / float64(spanExpandCount)
							for k := 0; k < vSpan; k++ {
								if expandRow[i-k] {
									heights[i-k] += delta
								}
							}
						}
					}
					minimumHeight := data.MinSize.Height
					if !data.VGrab || minimumHeight != 0 {
						var h float64
						if !data.VGrab || minimumHeight == layout.NoHint {
							h = data.minCacheSize.Height
						} else {
							h = minimumHeight
						}
						h -= spanMinHeight + float64(vSpan-1)*l.VSpacing
						if h > 0 {
							l.apportionExtra(h, i, spanExpandCount, vSpan, expandRow, minHeights)
						}
					}
				}
			}
		}
	}
	if height != layout.NoHint && expandCount > 0 {
		var totalHeight float64
		for i := 0; i < l.rows; i++ {
			totalHeight += heights[i]
		}
		c := expandCount
		delta := (availableHeight - totalHeight) / float64(c)
		for math.Abs(totalHeight-availableHeight) > 0.01 {
			for i := 0; i < l.rows; i++ {
				if expandRow[i] {
					if heights[i]+delta > minHeights[i] {
						heights[i] += delta
					} else {
						heights[i] = minHeights[i]
						expandRow[i] = false
						c--
					}
				}
			}
			for i := 0; i < l.rows; i++ {
				for j := 0; j < l.Columns; j++ {
					data := l.getData(grid, i, j, false)
					if data != nil {
						vSpan := xmath.MaxInt(1, xmath.MinInt(data.VSpan, l.rows))
						if vSpan > 1 {
							minimumHeight := data.MinSize.Height
							if !data.VGrab || minimumHeight != 0 {
								var spanHeight float64
								spanExpandCount := 0
								for k := 0; k < vSpan; k++ {
									spanHeight += heights[i-k]
									if expandRow[i-k] {
										spanExpandCount++
									}
								}
								var h float64
								if !data.VGrab || minimumHeight == layout.NoHint {
									h = data.minCacheSize.Height
								} else {
									h = minimumHeight
								}
								h -= spanHeight + float64(vSpan-1)*l.VSpacing
								if h > 0 {
									l.apportionExtra(h, i, spanExpandCount, vSpan, expandRow, heights)
								}
							}
						}
					}
				}
			}
			if c == 0 {
				break
			}
			totalHeight = 0
			for i := 0; i < l.rows; i++ {
				totalHeight += heights[i]
			}
			delta = (availableHeight - totalHeight) / float64(c)
		}
	}
	return heights
}

func (l *Layout) positionChildren(location geom.Point, grid [][]ui.Widget, widths []float64, heights []float64) {
	gridY := location.Y
	for i := 0; i < l.rows; i++ {
		gridX := location.X
		for j := 0; j < l.Columns; j++ {
			data := l.getData(grid, i, j, true)
			if data != nil {
				hSpan := xmath.MaxInt(1, xmath.MinInt(data.HSpan, l.Columns))
				vSpan := xmath.MaxInt(1, data.VSpan)
				var cellWidth, cellHeight float64
				for k := 0; k < hSpan; k++ {
					cellWidth += widths[j+k]
				}
				for k := 0; k < vSpan; k++ {
					cellHeight += heights[i+k]
				}
				cellWidth += l.HSpacing * float64(hSpan-1)
				childX := gridX
				childWidth := math.Min(data.cacheSize.Width, cellWidth)
				switch data.HAlign {
				case align.Middle:
					childX += math.Max(0, (cellWidth-childWidth)/2)
				case align.End:
					childX += math.Max(0, cellWidth-childWidth)
				case align.Fill:
					childWidth = cellWidth
				default:
				}
				cellHeight += l.VSpacing * float64(vSpan-1)
				childY := gridY
				childHeight := math.Min(data.cacheSize.Height, cellHeight)
				switch data.VAlign {
				case align.Middle:
					childY += math.Max(0, (cellHeight-childHeight)/2)
				case align.End:
					childY += math.Max(0, cellHeight-childHeight)
				case align.Fill:
					childHeight = cellHeight
				default:
				}
				child := grid[i][j]
				if child != nil {
					child.SetBounds(geom.Rect{Point: geom.Point{X: childX, Y: childY}, Size: geom.Size{Width: childWidth, Height: childHeight}})
				}
			}
			gridX += widths[j] + l.HSpacing
		}
		gridY += heights[i] + l.VSpacing
	}
}
