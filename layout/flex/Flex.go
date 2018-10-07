package flex

import (
	"math"

	"github.com/richardwilkes/toolbox/xmath"
	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/draw/align"
	"github.com/richardwilkes/ui/layout"
)

// Flex lays out the children of its widget based on the Data assigned to each child.
type Flex struct {
	widget   ui.Widget
	rows     int
	columns  int
	hSpacing float64
	vSpacing float64
	hAlign   align.Alignment
	vAlign   align.Alignment
	equal    bool
}

// NewLayout creates a new Flex layout and sets it on the widget.
func NewLayout(widget ui.Widget) *Flex {
	layout := &Flex{
		widget:   widget,
		columns:  1,
		hSpacing: 4,
		vSpacing: 2,
		hAlign:   align.Start,
		vAlign:   align.Start,
	}
	widget.SetLayout(layout)
	return layout
}

//Columns returns the number of columns.
func (flex *Flex) Columns() int {
	return flex.columns
}

// SetColumns sets the number of columns.
func (flex *Flex) SetColumns(columns int) *Flex {
	flex.columns = columns
	return flex
}

// EqualColumns returns true if each column will use the same amount of horizontal space.
func (flex *Flex) EqualColumns() bool {
	return flex.equal
}

// SetEqualColumns sets each column to use the same amount of horizontal space if true.
func (flex *Flex) SetEqualColumns(equal bool) *Flex {
	flex.equal = equal
	return flex
}

// HorizontalSpacing returns the horizontal spacing between columns.
func (flex *Flex) HorizontalSpacing() float64 {
	return flex.hSpacing
}

// SetHorizontalSpacing sets the horizontal spacing between columns.
func (flex *Flex) SetHorizontalSpacing(spacing float64) *Flex {
	flex.hSpacing = spacing
	return flex
}

// VerticalSpacing returns the vertical spacing between rows.
func (flex *Flex) VerticalSpacing() float64 {
	return flex.vSpacing
}

// SetVerticalSpacing sets the vertical spacing between rows.
func (flex *Flex) SetVerticalSpacing(spacing float64) *Flex {
	flex.vSpacing = spacing
	return flex
}

// HorizontalAlignment returns the horizontal alignment of the widget within its space.
func (flex *Flex) HorizontalAlignment() align.Alignment {
	return flex.hAlign
}

// SetHorizontalAlignment sets the horizontal alignment of the widget within its space.
func (flex *Flex) SetHorizontalAlignment(alignment align.Alignment) *Flex {
	flex.hAlign = alignment
	return flex
}

// VerticalAlignment returns the vertical alignment of the widget within its space.
func (flex *Flex) VerticalAlignment() align.Alignment {
	return flex.vAlign
}

// SetVerticalAlignment sets the vertical alignment of the widget within its space.
func (flex *Flex) SetVerticalAlignment(alignment align.Alignment) *Flex {
	flex.vAlign = alignment
	return flex
}

// Sizes implements the Layout interface.
func (flex *Flex) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	min = flex.layout(geom.Point{}, layout.NoHintSize, false, true)
	pref = flex.layout(geom.Point{}, layout.NoHintSize, false, false)
	if border := flex.widget.Border(); border != nil {
		insets := border.Insets()
		min.AddInsets(insets)
		pref.AddInsets(insets)
	}
	return min, pref, layout.DefaultMaxSize(pref)
}

// Layout implements the Layout interface.
func (flex *Flex) Layout() {
	var insets geom.Insets
	if border := flex.widget.Border(); border != nil {
		insets = border.Insets()
	}
	hint := flex.widget.Bounds().Size
	hint.SubtractInsets(insets)
	flex.layout(geom.Point{X: insets.Left, Y: insets.Top}, hint, true, false)
}

func (flex *Flex) layout(location geom.Point, hint geom.Size, move, useMinimumSize bool) geom.Size {
	var totalSize geom.Size
	if flex.columns > 0 {
		children := flex.prepChildren(useMinimumSize)
		if len(children) > 0 {
			grid := flex.buildGrid(children)
			widths := flex.adjustColumnWidths(hint.Width, grid)
			flex.wrap(hint.Width, grid, widths, useMinimumSize)
			heights := flex.adjustRowHeights(hint.Height, grid)
			totalSize.Width += flex.hSpacing * float64(flex.columns-1)
			totalSize.Height += flex.vSpacing * float64(flex.rows-1)
			for i := 0; i < flex.columns; i++ {
				totalSize.Width += widths[i]
			}
			for i := 0; i < flex.rows; i++ {
				totalSize.Height += heights[i]
			}
			if move {
				if totalSize.Width < hint.Width {
					if flex.hAlign == align.Middle {
						location.X += xmath.Round((hint.Width - totalSize.Width) / 2)
					} else if flex.hAlign == align.End {
						location.X += hint.Width - totalSize.Width
					}
				}
				if totalSize.Height < hint.Height {
					if flex.vAlign == align.Middle {
						location.Y += xmath.Round((hint.Height - totalSize.Height) / 2)
					} else if flex.vAlign == align.End {
						location.Y += hint.Height - totalSize.Height
					}
				}
				flex.positionChildren(location, grid, widths, heights)
			}
		}
	}
	return totalSize
}

func (flex *Flex) prepChildren(useMinimumSize bool) []ui.Widget {
	children := flex.widget.Children()
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

func (flex *Flex) buildGrid(children []ui.Widget) [][]ui.Widget {
	var grid [][]ui.Widget
	var row, column int
	flex.rows = 0
	for _, child := range children {
		data := getDataFromWidget(child)
		hSpan := xmath.MaxInt(1, xmath.MinInt(data.hSpan, flex.columns))
		vSpan := xmath.MaxInt(1, data.vSpan)
		for {
			lastRow := row + vSpan
			if lastRow >= len(grid) {
				grid = append(grid, make([]ui.Widget, flex.columns))
			}
			for column < flex.columns && grid[row][column] != nil {
				column++
			}
			endCount := column + hSpan
			if endCount <= flex.columns {
				index := column
				for index < endCount && grid[row][index] == nil {
					index++
				}
				if index == endCount {
					break
				}
				column = index
			}
			if column+hSpan >= flex.columns {
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
		flex.rows = xmath.MaxInt(flex.rows, row+vSpan)
		column += hSpan
	}
	return grid
}

func (flex *Flex) adjustColumnWidths(width float64, grid [][]ui.Widget) []float64 {
	availableWidth := width - flex.hSpacing*float64(flex.columns-1)
	expandCount := 0
	widths := make([]float64, flex.columns)
	minWidths := make([]float64, flex.columns)
	expandColumn := make([]bool, flex.columns)
	for j := 0; j < flex.columns; j++ {
		for i := 0; i < flex.rows; i++ {
			data := flex.getData(grid, i, j, true)
			if data != nil {
				hSpan := xmath.MaxInt(1, xmath.MinInt(data.hSpan, flex.columns))
				if hSpan == 1 {
					w := data.cacheSize.Width
					if widths[j] < w {
						widths[j] = w
					}
					if data.hGrab {
						if !expandColumn[j] {
							expandCount++
						}
						expandColumn[j] = true
					}
					minimumWidth := data.minCacheSize.Width
					if !data.hGrab || minimumWidth != 0 {
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
		for i := 0; i < flex.rows; i++ {
			data := flex.getData(grid, i, j, false)
			if data != nil {
				hSpan := xmath.MaxInt(1, xmath.MinInt(data.hSpan, flex.columns))
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
					if data.hGrab && spanExpandCount == 0 {
						expandCount++
						expandColumn[j] = true
					}
					w := data.cacheSize.Width - spanWidth - float64(hSpan-1)*flex.hSpacing
					if w > 0 {
						if flex.equal {
							equalWidth := math.Floor((w + spanWidth) / float64(hSpan))
							for k := 0; k < hSpan; k++ {
								if widths[j-k] < equalWidth {
									widths[j-k] = equalWidth
								}
							}
						} else {
							flex.apportionExtra(w, j, spanExpandCount, hSpan, expandColumn, widths)
						}
					}
					minimumWidth := data.minCacheSize.Width
					if !data.hGrab || minimumWidth != 0 {
						if !data.hGrab || minimumWidth == layout.NoHint {
							w = data.cacheSize.Width
						} else {
							w = minimumWidth
						}
						w -= spanMinWidth + float64(hSpan-1)*flex.hSpacing
						if w > 0 {
							flex.apportionExtra(w, j, spanExpandCount, hSpan, expandColumn, minWidths)
						}
					}
				}
			}
		}
	}
	if flex.equal {
		var minColumnWidth, columnWidth float64
		for i := 0; i < flex.columns; i++ {
			if minColumnWidth < minWidths[i] {
				minColumnWidth = minWidths[i]
			}
			if columnWidth < widths[i] {
				columnWidth = widths[i]
			}
		}
		if width != layout.NoHint && expandCount != 0 {
			columnWidth = math.Max(minColumnWidth, math.Floor(availableWidth/float64(flex.columns)))
		}
		for i := 0; i < flex.columns; i++ {
			expandColumn[i] = expandCount > 0
			widths[i] = columnWidth
		}
	} else {
		if width != layout.NoHint && expandCount > 0 {
			var totalWidth float64
			for i := 0; i < flex.columns; i++ {
				totalWidth += widths[i]
			}
			c := expandCount
			for math.Abs(totalWidth-availableWidth) > 0.01 {
				delta := (availableWidth - totalWidth) / float64(c)
				for j := 0; j < flex.columns; j++ {
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
				for j := 0; j < flex.columns; j++ {
					for i := 0; i < flex.rows; i++ {
						data := flex.getData(grid, i, j, false)
						if data != nil {
							hSpan := xmath.MaxInt(1, xmath.MinInt(data.hSpan, flex.columns))
							if hSpan > 1 {
								minimumWidth := data.minCacheSize.Width
								if !data.hGrab || minimumWidth != 0 {
									var spanWidth float64
									spanExpandCount := 0
									for k := 0; k < hSpan; k++ {
										spanWidth += widths[j-k]
										if expandColumn[j-k] {
											spanExpandCount++
										}
									}
									var w float64
									if !data.hGrab || minimumWidth == layout.NoHint {
										w = data.cacheSize.Width
									} else {
										w = minimumWidth
									}
									w -= spanWidth + float64(hSpan-1)*flex.hSpacing
									if w > 0 {
										flex.apportionExtra(w, j, spanExpandCount, hSpan, expandColumn, widths)
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
				for i := 0; i < flex.columns; i++ {
					totalWidth += widths[i]
				}
			}
		}
	}
	return widths
}

func (flex *Flex) apportionExtra(extra float64, base, count, span int, expand []bool, values []float64) {
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

func (flex *Flex) getData(grid [][]ui.Widget, row, column int, first bool) *Data {
	block := grid[row][column]
	if block != nil {
		data := getDataFromWidget(block)
		hSpan := xmath.MaxInt(1, xmath.MinInt(data.hSpan, flex.columns))
		vSpan := xmath.MaxInt(1, data.vSpan)
		var i, j int
		if first {
			i = row + vSpan - 1
			j = column + hSpan - 1
		} else {
			i = row - vSpan + 1
			j = column - hSpan + 1
		}
		if i >= 0 && i < flex.rows {
			if j >= 0 && j < flex.columns {
				if block == grid[i][j] {
					return data
				}
			}
		}
	}
	return nil
}

func (flex *Flex) wrap(width float64, grid [][]ui.Widget, widths []float64, useMinimumSize bool) {
	if width != layout.NoHint {
		for j := 0; j < flex.columns; j++ {
			for i := 0; i < flex.rows; i++ {
				data := flex.getData(grid, i, j, false)
				if data != nil {
					if data.sizeHint.Height == layout.NoHint {
						hSpan := xmath.MaxInt(1, xmath.MinInt(data.hSpan, flex.columns))
						var currentWidth float64
						for k := 0; k < hSpan; k++ {
							currentWidth += widths[j-k]
						}
						currentWidth += float64(hSpan-1) * flex.hSpacing
						if currentWidth != data.cacheSize.Width && data.hAlign == align.Fill || data.cacheSize.Width > currentWidth {
							data.computeCacheSize(grid[i][j], geom.Size{Width: math.Max(data.minCacheSize.Width, currentWidth), Height: layout.NoHint}, useMinimumSize)
							minimumHeight := data.minSize.Height
							if data.vGrab && minimumHeight > 0 && data.cacheSize.Height < minimumHeight {
								data.cacheSize.Height = minimumHeight
							}
						}
					}
				}
			}
		}
	}
}

func (flex *Flex) adjustRowHeights(height float64, grid [][]ui.Widget) []float64 {
	availableHeight := height - flex.vSpacing*float64(flex.rows-1)
	expandCount := 0
	heights := make([]float64, flex.rows)
	minHeights := make([]float64, flex.rows)
	expandRow := make([]bool, flex.rows)
	for i := 0; i < flex.rows; i++ {
		for j := 0; j < flex.columns; j++ {
			data := flex.getData(grid, i, j, true)
			if data != nil {
				vSpan := xmath.MaxInt(1, xmath.MinInt(data.vSpan, flex.rows))
				if vSpan == 1 {
					h := data.cacheSize.Height
					if heights[i] < h {
						heights[i] = h
					}
					if data.vGrab {
						if !expandRow[i] {
							expandCount++
						}
						expandRow[i] = true
					}
					minimumHeight := data.minSize.Height
					if !data.vGrab || minimumHeight != 0 {
						var h float64
						if !data.vGrab || minimumHeight == layout.NoHint {
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
		for j := 0; j < flex.columns; j++ {
			data := flex.getData(grid, i, j, false)
			if data != nil {
				vSpan := xmath.MaxInt(1, xmath.MinInt(data.vSpan, flex.rows))
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
					if data.vGrab && spanExpandCount == 0 {
						expandCount++
						expandRow[i] = true
					}
					h := data.cacheSize.Height - spanHeight - float64(vSpan-1)*flex.vSpacing
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
					minimumHeight := data.minSize.Height
					if !data.vGrab || minimumHeight != 0 {
						var h float64
						if !data.vGrab || minimumHeight == layout.NoHint {
							h = data.minCacheSize.Height
						} else {
							h = minimumHeight
						}
						h -= spanMinHeight + float64(vSpan-1)*flex.vSpacing
						if h > 0 {
							flex.apportionExtra(h, i, spanExpandCount, vSpan, expandRow, minHeights)
						}
					}
				}
			}
		}
	}
	if height != layout.NoHint && expandCount > 0 {
		var totalHeight float64
		for i := 0; i < flex.rows; i++ {
			totalHeight += heights[i]
		}
		c := expandCount
		delta := (availableHeight - totalHeight) / float64(c)
		for math.Abs(totalHeight-availableHeight) > 0.01 {
			for i := 0; i < flex.rows; i++ {
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
			for i := 0; i < flex.rows; i++ {
				for j := 0; j < flex.columns; j++ {
					data := flex.getData(grid, i, j, false)
					if data != nil {
						vSpan := xmath.MaxInt(1, xmath.MinInt(data.vSpan, flex.rows))
						if vSpan > 1 {
							minimumHeight := data.minSize.Height
							if !data.vGrab || minimumHeight != 0 {
								var spanHeight float64
								spanExpandCount := 0
								for k := 0; k < vSpan; k++ {
									spanHeight += heights[i-k]
									if expandRow[i-k] {
										spanExpandCount++
									}
								}
								var h float64
								if !data.vGrab || minimumHeight == layout.NoHint {
									h = data.minCacheSize.Height
								} else {
									h = minimumHeight
								}
								h -= spanHeight + float64(vSpan-1)*flex.vSpacing
								if h > 0 {
									flex.apportionExtra(h, i, spanExpandCount, vSpan, expandRow, heights)
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
			for i := 0; i < flex.rows; i++ {
				totalHeight += heights[i]
			}
			delta = (availableHeight - totalHeight) / float64(c)
		}
	}
	return heights
}

func (flex *Flex) positionChildren(location geom.Point, grid [][]ui.Widget, widths []float64, heights []float64) {
	gridY := location.Y
	for i := 0; i < flex.rows; i++ {
		gridX := location.X
		for j := 0; j < flex.columns; j++ {
			data := flex.getData(grid, i, j, true)
			if data != nil {
				hSpan := xmath.MaxInt(1, xmath.MinInt(data.hSpan, flex.columns))
				vSpan := xmath.MaxInt(1, data.vSpan)
				var cellWidth, cellHeight float64
				for k := 0; k < hSpan; k++ {
					cellWidth += widths[j+k]
				}
				for k := 0; k < vSpan; k++ {
					cellHeight += heights[i+k]
				}
				cellWidth += flex.hSpacing * float64(hSpan-1)
				childX := gridX
				childWidth := math.Min(data.cacheSize.Width, cellWidth)
				switch data.hAlign {
				case align.Middle:
					childX += math.Max(0, (cellWidth-childWidth)/2)
				case align.End:
					childX += math.Max(0, cellWidth-childWidth)
				case align.Fill:
					childWidth = cellWidth
				default:
				}
				cellHeight += flex.vSpacing * float64(vSpan-1)
				childY := gridY
				childHeight := math.Min(data.cacheSize.Height, cellHeight)
				switch data.vAlign {
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
			gridX += widths[j] + flex.hSpacing
		}
		gridY += heights[i] + flex.vSpacing
	}
}
