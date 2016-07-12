// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

// PrecisionLayout lays out the children of its widget based on the PrecisionData assigned to each
// child.
type PrecisionLayout struct {
	widget   Widget
	rows     int
	columns  int
	hSpacing float32
	vSpacing float32
	hAlign   Alignment
	vAlign   Alignment
	equal    bool
}

// NewPrecisionLayout creates a new PrecisionLayout and sets it on the widget.
func NewPrecisionLayout(widget Widget) *PrecisionLayout {
	layout := &PrecisionLayout{widget: widget, columns: 1, hSpacing: 4, vSpacing: 2, hAlign: AlignStart, vAlign: AlignStart}
	widget.SetLayout(layout)
	return layout
}

//Columns returns the number of columns.
func (p *PrecisionLayout) Columns() int {
	return p.columns
}

// SetColumns sets the number of columns.
func (p *PrecisionLayout) SetColumns(columns int) *PrecisionLayout {
	p.columns = columns
	return p
}

// EqualColumns returns true if each column will use the same amount of horizontal space.
func (p *PrecisionLayout) EqualColumns() bool {
	return p.equal
}

// SetEqualColumns sets each column to use the same amount of horizontal space if true.
func (p *PrecisionLayout) SetEqualColumns(equal bool) *PrecisionLayout {
	p.equal = equal
	return p
}

// HorizontalSpacing returns the horizontal spacing between columns.
func (p *PrecisionLayout) HorizontalSpacing() float32 {
	return p.hSpacing
}

// SetHorizontalSpacing sets the horizontal spacing between columns.
func (p *PrecisionLayout) SetHorizontalSpacing(spacing float32) *PrecisionLayout {
	p.hSpacing = spacing
	return p
}

// VerticalSpacing returns the vertical spacing between rows.
func (p *PrecisionLayout) VerticalSpacing() float32 {
	return p.vSpacing
}

// SetVerticalSpacing sets the vertical spacing between rows.
func (p *PrecisionLayout) SetVerticalSpacing(spacing float32) *PrecisionLayout {
	p.vSpacing = spacing
	return p
}

// HorizontalAlignment returns the horizontal alignment of the widget within its space.
func (p *PrecisionLayout) HorizontalAlignment() Alignment {
	return p.hAlign
}

// SetHorizontalAlignment sets the horizontal alignment of the widget within its space.
func (p *PrecisionLayout) SetHorizontalAlignment(alignment Alignment) *PrecisionLayout {
	p.hAlign = alignment
	return p
}

// VerticalAlignment returns the vertical alignment of the widget within its space.
func (p *PrecisionLayout) VerticalAlignment() Alignment {
	return p.vAlign
}

// SetVerticalAlignment sets the vertical alignment of the widget within its space.
func (p *PrecisionLayout) SetVerticalAlignment(alignment Alignment) *PrecisionLayout {
	p.vAlign = alignment
	return p
}

// Sizes implements the Layout interface.
func (p *PrecisionLayout) Sizes(hint Size) (min, pref, max Size) {
	min = p.layout(Point{}, NoLayoutHintSize, false, true)
	pref = p.layout(Point{}, NoLayoutHintSize, false, false)
	if border := p.widget.Border(); border != nil {
		insets := border.Insets()
		min.AddInsets(insets)
		pref.AddInsets(insets)
	}
	return min, pref, DefaultLayoutMaxSize(pref)
}

// Layout implements the Layout interface.
func (p *PrecisionLayout) Layout() {
	var insets Insets
	if border := p.widget.Border(); border != nil {
		insets = border.Insets()
	}
	hint := p.widget.Bounds().Size
	hint.SubtractInsets(insets)
	p.layout(Point{X: insets.Left, Y: insets.Top}, hint, true, false)
}

func (p *PrecisionLayout) layout(location Point, hint Size, move, useMinimumSize bool) Size {
	var totalSize Size
	if p.columns > 0 {
		children := p.prepChildren(useMinimumSize)
		if len(children) > 0 {
			grid := p.buildGrid(children)
			widths := p.adjustColumnWidths(hint.Width, grid)
			p.wrap(hint.Width, grid, widths, useMinimumSize)
			heights := p.adjustRowHeights(hint.Height, grid)
			totalSize.Width += p.hSpacing * float32(p.columns-1)
			totalSize.Height += p.vSpacing * float32(p.rows-1)
			for i := 0; i < p.columns; i++ {
				totalSize.Width += widths[i]
			}
			for i := 0; i < p.rows; i++ {
				totalSize.Height += heights[i]
			}
			if move {
				if totalSize.Width < hint.Width {
					if p.hAlign == AlignMiddle {
						location.X += RoundFloat32((hint.Width - totalSize.Width) / 2)
					} else if p.hAlign == AlignEnd {
						location.X += hint.Width - totalSize.Width
					}
				}
				if totalSize.Height < hint.Height {
					if p.vAlign == AlignMiddle {
						location.Y += RoundFloat32((hint.Height - totalSize.Height) / 2)
					} else if p.vAlign == AlignEnd {
						location.Y += hint.Height - totalSize.Height
					}
				}
				p.positionChildren(location, grid, widths, heights)
			}
		}
	}
	return totalSize
}

func (p *PrecisionLayout) prepChildren(useMinimumSize bool) []Widget {
	children := p.widget.Children()
	for _, child := range children {
		var layoutData *PrecisionData
		var ok bool
		if layoutData, ok = child.LayoutData().(*PrecisionData); !ok {
			layoutData = NewPrecisionData()
			child.SetLayoutData(layoutData)
		}
		layoutData.computeCacheSize(child, NoLayoutHintSize, useMinimumSize)
	}
	return children
}

func (p *PrecisionLayout) buildGrid(children []Widget) [][]Widget {
	var grid [][]Widget
	var row, column int
	p.rows = 0
	for _, child := range children {
		data := child.LayoutData().(*PrecisionData)
		hSpan := MaxInt(1, MinInt(data.hSpan, p.columns))
		vSpan := MaxInt(1, data.vSpan)
		for {
			lastRow := row + vSpan
			if lastRow >= len(grid) {
				grid = append(grid, make([]Widget, p.columns))
			}
			for column < p.columns && grid[row][column] != nil {
				column++
			}
			endCount := column + hSpan
			if endCount <= p.columns {
				index := column
				for index < endCount && grid[row][index] == nil {
					index++
				}
				if index == endCount {
					break
				}
				column = index
			}
			if column+hSpan >= p.columns {
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
		p.rows = MaxInt(p.rows, row+vSpan)
		column += hSpan
	}
	return grid
}

func (p *PrecisionLayout) adjustColumnWidths(width float32, grid [][]Widget) []float32 {
	availableWidth := width - p.hSpacing*float32(p.columns-1)
	expandCount := 0
	widths := make([]float32, p.columns)
	minWidths := make([]float32, p.columns)
	expandColumn := make([]bool, p.columns)
	for j := 0; j < p.columns; j++ {
		for i := 0; i < p.rows; i++ {
			data := p.getData(grid, i, j, true)
			if data != nil {
				hSpan := MaxInt(1, MinInt(data.hSpan, p.columns))
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
					minimumWidth := data.cacheMinWidth
					if !data.hGrab || minimumWidth != 0 {
						if minimumWidth == NoLayoutHint {
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
		for i := 0; i < p.rows; i++ {
			data := p.getData(grid, i, j, false)
			if data != nil {
				hSpan := MaxInt(1, MinInt(data.hSpan, p.columns))
				if hSpan > 1 {
					var spanWidth, spanMinWidth float32
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
					w := data.cacheSize.Width - spanWidth - float32(hSpan-1)*p.hSpacing
					if w > 0 {
						if p.equal {
							equalWidth := FloorFloat32((w + spanWidth) / float32(hSpan))
							for k := 0; k < hSpan; k++ {
								if widths[j-k] < equalWidth {
									widths[j-k] = equalWidth
								}
							}
						} else {
							p.apportionExtra(w, j, spanExpandCount, hSpan, expandColumn, widths)
						}
					}
					minimumWidth := data.cacheMinWidth
					if !data.hGrab || minimumWidth != 0 {
						if !data.hGrab || minimumWidth == NoLayoutHint {
							w = data.cacheSize.Width
						} else {
							w = minimumWidth
						}
						w -= spanMinWidth + float32(hSpan-1)*p.hSpacing
						if w > 0 {
							p.apportionExtra(w, j, spanExpandCount, hSpan, expandColumn, minWidths)
						}
					}
				}
			}
		}
	}
	if p.equal {
		var minColumnWidth, columnWidth float32
		for i := 0; i < p.columns; i++ {
			if minColumnWidth < minWidths[i] {
				minColumnWidth = minWidths[i]
			}
			if columnWidth < widths[i] {
				columnWidth = widths[i]
			}
		}
		if width != NoLayoutHint && expandCount != 0 {
			columnWidth = MaxFloat32(minColumnWidth, FloorFloat32(availableWidth/float32(p.columns)))
		}
		for i := 0; i < p.columns; i++ {
			expandColumn[i] = expandCount > 0
			widths[i] = columnWidth
		}
	} else {
		if width != NoLayoutHint && expandCount > 0 {
			var totalWidth float32
			for i := 0; i < p.columns; i++ {
				totalWidth += widths[i]
			}
			c := expandCount
			for AbsFloat32(totalWidth-availableWidth) > 0.01 {
				delta := (availableWidth - totalWidth) / float32(c)
				for j := 0; j < p.columns; j++ {
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
				for j := 0; j < p.columns; j++ {
					for i := 0; i < p.rows; i++ {
						data := p.getData(grid, i, j, false)
						if data != nil {
							hSpan := MaxInt(1, MinInt(data.hSpan, p.columns))
							if hSpan > 1 {
								minimumWidth := data.cacheMinWidth
								if !data.hGrab || minimumWidth != 0 {
									var spanWidth float32
									spanExpandCount := 0
									for k := 0; k < hSpan; k++ {
										spanWidth += widths[j-k]
										if expandColumn[j-k] {
											spanExpandCount++
										}
									}
									var w float32
									if !data.hGrab || minimumWidth == NoLayoutHint {
										w = data.cacheSize.Width
									} else {
										w = minimumWidth
									}
									w -= spanWidth + float32(hSpan-1)*p.hSpacing
									if w > 0 {
										p.apportionExtra(w, j, spanExpandCount, hSpan, expandColumn, widths)
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
				for i := 0; i < p.columns; i++ {
					totalWidth += widths[i]
				}
			}
		}
	}
	return widths
}

func (p *PrecisionLayout) apportionExtra(extra float32, base, count, span int, expand []bool, values []float32) {
	if count == 0 {
		values[base] += extra
	} else {
		extraInt := int(FloorFloat32(extra))
		delta := extraInt / count
		remainder := extraInt - delta*count
		for i := 0; i < span; i++ {
			j := base - i
			if expand[j] {
				values[j] += float32(delta)
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

func (p *PrecisionLayout) getData(grid [][]Widget, row, column int, first bool) *PrecisionData {
	block := grid[row][column]
	if block != nil {
		data := block.LayoutData().(*PrecisionData)
		hSpan := MaxInt(1, MinInt(data.hSpan, p.columns))
		vSpan := MaxInt(1, data.vSpan)
		var i, j int
		if first {
			i = row + vSpan - 1
			j = column + hSpan - 1
		} else {
			i = row - vSpan + 1
			j = column - hSpan + 1
		}
		if i >= 0 && i < p.rows {
			if j >= 0 && j < p.columns {
				if block == grid[i][j] {
					return data
				}
			}
		}
	}
	return nil
}

func (p *PrecisionLayout) wrap(width float32, grid [][]Widget, widths []float32, useMinimumSize bool) {
	if width != NoLayoutHint {
		for j := 0; j < p.columns; j++ {
			for i := 0; i < p.rows; i++ {
				data := p.getData(grid, i, j, false)
				if data != nil {
					if data.sizeHint.Height == NoLayoutHint {
						hSpan := MaxInt(1, MinInt(data.hSpan, p.columns))
						var currentWidth float32
						for k := 0; k < hSpan; k++ {
							currentWidth += widths[j-k]
						}
						currentWidth += float32(hSpan-1) * p.hSpacing
						if currentWidth != data.cacheSize.Width && data.hAlign == AlignFill || data.cacheSize.Width > currentWidth {
							data.computeCacheSize(grid[i][j], Size{Width: MaxFloat32(data.cacheMinWidth, currentWidth), Height: NoLayoutHint}, useMinimumSize)
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

func (p *PrecisionLayout) adjustRowHeights(height float32, grid [][]Widget) []float32 {
	availableHeight := height - p.vSpacing*float32(p.rows-1)
	expandCount := 0
	heights := make([]float32, p.rows)
	minHeights := make([]float32, p.rows)
	expandRow := make([]bool, p.rows)
	for i := 0; i < p.rows; i++ {
		for j := 0; j < p.columns; j++ {
			data := p.getData(grid, i, j, true)
			if data != nil {
				vSpan := MaxInt(1, MinInt(data.vSpan, p.rows))
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
						var h float32
						if !data.vGrab || minimumHeight == NoLayoutHint {
							h = data.cacheSize.Height
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
		for j := 0; j < p.columns; j++ {
			data := p.getData(grid, i, j, false)
			if data != nil {
				vSpan := MaxInt(1, MinInt(data.vSpan, p.rows))
				if vSpan > 1 {
					var spanHeight, spanMinHeight float32
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
					h := data.cacheSize.Height - spanHeight - float32(vSpan-1)*p.vSpacing
					if h > 0 {
						if spanExpandCount == 0 {
							heights[i] += h
						} else {
							delta := h / float32(spanExpandCount)
							for k := 0; k < vSpan; k++ {
								if expandRow[i-k] {
									heights[i-k] += delta
								}
							}
						}
					}
					minimumHeight := data.minSize.Height
					if !data.vGrab || minimumHeight != 0 {
						var h float32
						if !data.vGrab || minimumHeight == NoLayoutHint {
							h = data.cacheSize.Height
						} else {
							h = minimumHeight
						}
						h -= spanMinHeight + float32(vSpan-1)*p.vSpacing
						if h > 0 {
							p.apportionExtra(h, i, spanExpandCount, vSpan, expandRow, minHeights)
						}
					}
				}
			}
		}
	}
	if height != NoLayoutHint && expandCount > 0 {
		var totalHeight float32
		for i := 0; i < p.rows; i++ {
			totalHeight += heights[i]
		}
		c := expandCount
		delta := (availableHeight - totalHeight) / float32(c)
		for AbsFloat32(totalHeight-availableHeight) > 0.01 {
			for i := 0; i < p.rows; i++ {
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
			for i := 0; i < p.rows; i++ {
				for j := 0; j < p.columns; j++ {
					data := p.getData(grid, i, j, false)
					if data != nil {
						vSpan := MaxInt(1, MinInt(data.vSpan, p.rows))
						if vSpan > 1 {
							minimumHeight := data.minSize.Height
							if !data.vGrab || minimumHeight != 0 {
								var spanHeight float32
								spanExpandCount := 0
								for k := 0; k < vSpan; k++ {
									spanHeight += heights[i-k]
									if expandRow[i-k] {
										spanExpandCount++
									}
								}
								var h float32
								if !data.vGrab || minimumHeight == NoLayoutHint {
									h = data.cacheSize.Height
								} else {
									h = minimumHeight
								}
								h -= spanHeight + float32(vSpan-1)*p.vSpacing
								if h > 0 {
									p.apportionExtra(h, i, spanExpandCount, vSpan, expandRow, heights)
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
			for i := 0; i < p.rows; i++ {
				totalHeight += heights[i]
			}
			delta = (availableHeight - totalHeight) / float32(c)
		}
	}
	return heights
}

func (p *PrecisionLayout) positionChildren(location Point, grid [][]Widget, widths []float32, heights []float32) {
	gridY := location.Y
	for i := 0; i < p.rows; i++ {
		gridX := location.X
		for j := 0; j < p.columns; j++ {
			data := p.getData(grid, i, j, true)
			if data != nil {
				hSpan := MaxInt(1, MinInt(data.hSpan, p.columns))
				vSpan := MaxInt(1, data.vSpan)
				var cellWidth, cellHeight float32
				for k := 0; k < hSpan; k++ {
					cellWidth += widths[j+k]
				}
				for k := 0; k < vSpan; k++ {
					cellHeight += heights[i+k]
				}
				cellWidth += p.hSpacing * float32(hSpan-1)
				childX := gridX
				childWidth := MinFloat32(data.cacheSize.Width, cellWidth)
				switch data.hAlign {
				case AlignMiddle:
					childX += MaxFloat32(0, (cellWidth-childWidth)/2)
				case AlignEnd:
					childX += MaxFloat32(0, cellWidth-childWidth)
				case AlignFill:
					childWidth = cellWidth
				default:
				}
				cellHeight += p.vSpacing * float32(vSpan-1)
				childY := gridY
				childHeight := MinFloat32(data.cacheSize.Height, cellHeight)
				switch data.vAlign {
				case AlignMiddle:
					childY += MaxFloat32(0, (cellHeight-childHeight)/2)
				case AlignEnd:
					childY += MaxFloat32(0, cellHeight-childHeight)
				case AlignFill:
					childHeight = cellHeight
				default:
				}
				child := grid[i][j]
				if child != nil {
					child.SetBounds(Rect{Point: Point{X: childX, Y: childY}, Size: Size{Width: childWidth, Height: childHeight}})
				}
			}
			gridX += widths[j] + p.hSpacing
		}
		gridY += heights[i] + p.vSpacing
	}
}
