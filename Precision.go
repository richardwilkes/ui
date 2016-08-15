// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

import (
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/geom"
	"github.com/richardwilkes/xmath"
)

// Precision lays out the children of its widget based on the PrecisionData assigned to each child.
type Precision struct {
	widget   Widget
	rows     int
	columns  int
	hSpacing float32
	vSpacing float32
	hAlign   draw.Alignment
	vAlign   draw.Alignment
	equal    bool
}

// NewPrecision creates a new Precision layout and sets it on the widget.
func NewPrecision(widget Widget) *Precision {
	layout := &Precision{widget: widget, columns: 1, hSpacing: 4, vSpacing: 2, hAlign: draw.AlignStart, vAlign: draw.AlignStart}
	widget.SetLayout(layout)
	return layout
}

//Columns returns the number of columns.
func (p *Precision) Columns() int {
	return p.columns
}

// SetColumns sets the number of columns.
func (p *Precision) SetColumns(columns int) *Precision {
	p.columns = columns
	return p
}

// EqualColumns returns true if each column will use the same amount of horizontal space.
func (p *Precision) EqualColumns() bool {
	return p.equal
}

// SetEqualColumns sets each column to use the same amount of horizontal space if true.
func (p *Precision) SetEqualColumns(equal bool) *Precision {
	p.equal = equal
	return p
}

// HorizontalSpacing returns the horizontal spacing between columns.
func (p *Precision) HorizontalSpacing() float32 {
	return p.hSpacing
}

// SetHorizontalSpacing sets the horizontal spacing between columns.
func (p *Precision) SetHorizontalSpacing(spacing float32) *Precision {
	p.hSpacing = spacing
	return p
}

// VerticalSpacing returns the vertical spacing between rows.
func (p *Precision) VerticalSpacing() float32 {
	return p.vSpacing
}

// SetVerticalSpacing sets the vertical spacing between rows.
func (p *Precision) SetVerticalSpacing(spacing float32) *Precision {
	p.vSpacing = spacing
	return p
}

// HorizontalAlignment returns the horizontal alignment of the widget within its space.
func (p *Precision) HorizontalAlignment() draw.Alignment {
	return p.hAlign
}

// SetHorizontalAlignment sets the horizontal alignment of the widget within its space.
func (p *Precision) SetHorizontalAlignment(alignment draw.Alignment) *Precision {
	p.hAlign = alignment
	return p
}

// VerticalAlignment returns the vertical alignment of the widget within its space.
func (p *Precision) VerticalAlignment() draw.Alignment {
	return p.vAlign
}

// SetVerticalAlignment sets the vertical alignment of the widget within its space.
func (p *Precision) SetVerticalAlignment(alignment draw.Alignment) *Precision {
	p.vAlign = alignment
	return p
}

// Sizes implements the Layout interface.
func (p *Precision) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	min = p.layout(geom.Point{}, NoHintSize, false, true)
	pref = p.layout(geom.Point{}, NoHintSize, false, false)
	if border := p.widget.Border(); border != nil {
		insets := border.Insets()
		min.AddInsets(insets)
		pref.AddInsets(insets)
	}
	return min, pref, DefaultMaxSize(pref)
}

// Layout implements the Layout interface.
func (p *Precision) Layout() {
	var insets geom.Insets
	if border := p.widget.Border(); border != nil {
		insets = border.Insets()
	}
	hint := p.widget.Bounds().Size
	hint.SubtractInsets(insets)
	p.layout(geom.Point{X: insets.Left, Y: insets.Top}, hint, true, false)
}

func (p *Precision) layout(location geom.Point, hint geom.Size, move, useMinimumSize bool) geom.Size {
	var totalSize geom.Size
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
					if p.hAlign == draw.AlignMiddle {
						location.X += xmath.RoundFloat32((hint.Width - totalSize.Width) / 2)
					} else if p.hAlign == draw.AlignEnd {
						location.X += hint.Width - totalSize.Width
					}
				}
				if totalSize.Height < hint.Height {
					if p.vAlign == draw.AlignMiddle {
						location.Y += xmath.RoundFloat32((hint.Height - totalSize.Height) / 2)
					} else if p.vAlign == draw.AlignEnd {
						location.Y += hint.Height - totalSize.Height
					}
				}
				p.positionChildren(location, grid, widths, heights)
			}
		}
	}
	return totalSize
}

func (p *Precision) prepChildren(useMinimumSize bool) []Widget {
	children := p.widget.Children()
	for _, child := range children {
		var layoutData *PrecisionData
		var ok bool
		if layoutData, ok = child.LayoutData().(*PrecisionData); !ok {
			layoutData = NewPrecisionData()
			child.SetLayoutData(layoutData)
		}
		layoutData.computeCacheSize(child, NoHintSize, useMinimumSize)
	}
	return children
}

func (p *Precision) buildGrid(children []Widget) [][]Widget {
	var grid [][]Widget
	var row, column int
	p.rows = 0
	for _, child := range children {
		data := child.LayoutData().(*PrecisionData)
		hSpan := xmath.MaxInt(1, xmath.MinInt(data.hSpan, p.columns))
		vSpan := xmath.MaxInt(1, data.vSpan)
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
		p.rows = xmath.MaxInt(p.rows, row+vSpan)
		column += hSpan
	}
	return grid
}

func (p *Precision) adjustColumnWidths(width float32, grid [][]Widget) []float32 {
	availableWidth := width - p.hSpacing*float32(p.columns-1)
	expandCount := 0
	widths := make([]float32, p.columns)
	minWidths := make([]float32, p.columns)
	expandColumn := make([]bool, p.columns)
	for j := 0; j < p.columns; j++ {
		for i := 0; i < p.rows; i++ {
			data := p.getData(grid, i, j, true)
			if data != nil {
				hSpan := xmath.MaxInt(1, xmath.MinInt(data.hSpan, p.columns))
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
						if minimumWidth == NoHint {
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
				hSpan := xmath.MaxInt(1, xmath.MinInt(data.hSpan, p.columns))
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
							equalWidth := xmath.FloorFloat32((w + spanWidth) / float32(hSpan))
							for k := 0; k < hSpan; k++ {
								if widths[j-k] < equalWidth {
									widths[j-k] = equalWidth
								}
							}
						} else {
							p.apportionExtra(w, j, spanExpandCount, hSpan, expandColumn, widths)
						}
					}
					minimumWidth := data.minCacheSize.Width
					if !data.hGrab || minimumWidth != 0 {
						if !data.hGrab || minimumWidth == NoHint {
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
		if width != NoHint && expandCount != 0 {
			columnWidth = xmath.MaxFloat32(minColumnWidth, xmath.FloorFloat32(availableWidth/float32(p.columns)))
		}
		for i := 0; i < p.columns; i++ {
			expandColumn[i] = expandCount > 0
			widths[i] = columnWidth
		}
	} else {
		if width != NoHint && expandCount > 0 {
			var totalWidth float32
			for i := 0; i < p.columns; i++ {
				totalWidth += widths[i]
			}
			c := expandCount
			for xmath.AbsFloat32(totalWidth-availableWidth) > 0.01 {
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
							hSpan := xmath.MaxInt(1, xmath.MinInt(data.hSpan, p.columns))
							if hSpan > 1 {
								minimumWidth := data.minCacheSize.Width
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
									if !data.hGrab || minimumWidth == NoHint {
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

func (p *Precision) apportionExtra(extra float32, base, count, span int, expand []bool, values []float32) {
	if count == 0 {
		values[base] += extra
	} else {
		extraInt := int(xmath.FloorFloat32(extra))
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

func (p *Precision) getData(grid [][]Widget, row, column int, first bool) *PrecisionData {
	block := grid[row][column]
	if block != nil {
		data := block.LayoutData().(*PrecisionData)
		hSpan := xmath.MaxInt(1, xmath.MinInt(data.hSpan, p.columns))
		vSpan := xmath.MaxInt(1, data.vSpan)
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

func (p *Precision) wrap(width float32, grid [][]Widget, widths []float32, useMinimumSize bool) {
	if width != NoHint {
		for j := 0; j < p.columns; j++ {
			for i := 0; i < p.rows; i++ {
				data := p.getData(grid, i, j, false)
				if data != nil {
					if data.sizeHint.Height == NoHint {
						hSpan := xmath.MaxInt(1, xmath.MinInt(data.hSpan, p.columns))
						var currentWidth float32
						for k := 0; k < hSpan; k++ {
							currentWidth += widths[j-k]
						}
						currentWidth += float32(hSpan-1) * p.hSpacing
						if currentWidth != data.cacheSize.Width && data.hAlign == draw.AlignFill || data.cacheSize.Width > currentWidth {
							data.computeCacheSize(grid[i][j], geom.Size{Width: xmath.MaxFloat32(data.minCacheSize.Width, currentWidth), Height: NoHint}, useMinimumSize)
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

func (p *Precision) adjustRowHeights(height float32, grid [][]Widget) []float32 {
	availableHeight := height - p.vSpacing*float32(p.rows-1)
	expandCount := 0
	heights := make([]float32, p.rows)
	minHeights := make([]float32, p.rows)
	expandRow := make([]bool, p.rows)
	for i := 0; i < p.rows; i++ {
		for j := 0; j < p.columns; j++ {
			data := p.getData(grid, i, j, true)
			if data != nil {
				vSpan := xmath.MaxInt(1, xmath.MinInt(data.vSpan, p.rows))
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
						if !data.vGrab || minimumHeight == NoHint {
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
		for j := 0; j < p.columns; j++ {
			data := p.getData(grid, i, j, false)
			if data != nil {
				vSpan := xmath.MaxInt(1, xmath.MinInt(data.vSpan, p.rows))
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
						if !data.vGrab || minimumHeight == NoHint {
							h = data.minCacheSize.Height
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
	if height != NoHint && expandCount > 0 {
		var totalHeight float32
		for i := 0; i < p.rows; i++ {
			totalHeight += heights[i]
		}
		c := expandCount
		delta := (availableHeight - totalHeight) / float32(c)
		for xmath.AbsFloat32(totalHeight-availableHeight) > 0.01 {
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
						vSpan := xmath.MaxInt(1, xmath.MinInt(data.vSpan, p.rows))
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
								if !data.vGrab || minimumHeight == NoHint {
									h = data.minCacheSize.Height
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

func (p *Precision) positionChildren(location geom.Point, grid [][]Widget, widths []float32, heights []float32) {
	gridY := location.Y
	for i := 0; i < p.rows; i++ {
		gridX := location.X
		for j := 0; j < p.columns; j++ {
			data := p.getData(grid, i, j, true)
			if data != nil {
				hSpan := xmath.MaxInt(1, xmath.MinInt(data.hSpan, p.columns))
				vSpan := xmath.MaxInt(1, data.vSpan)
				var cellWidth, cellHeight float32
				for k := 0; k < hSpan; k++ {
					cellWidth += widths[j+k]
				}
				for k := 0; k < vSpan; k++ {
					cellHeight += heights[i+k]
				}
				cellWidth += p.hSpacing * float32(hSpan-1)
				childX := gridX
				childWidth := xmath.MinFloat32(data.cacheSize.Width, cellWidth)
				switch data.hAlign {
				case draw.AlignMiddle:
					childX += xmath.MaxFloat32(0, (cellWidth-childWidth)/2)
				case draw.AlignEnd:
					childX += xmath.MaxFloat32(0, cellWidth-childWidth)
				case draw.AlignFill:
					childWidth = cellWidth
				default:
				}
				cellHeight += p.vSpacing * float32(vSpan-1)
				childY := gridY
				childHeight := xmath.MinFloat32(data.cacheSize.Height, cellHeight)
				switch data.vAlign {
				case draw.AlignMiddle:
					childY += xmath.MaxFloat32(0, (cellHeight-childHeight)/2)
				case draw.AlignEnd:
					childY += xmath.MaxFloat32(0, cellHeight-childHeight)
				case draw.AlignFill:
					childHeight = cellHeight
				default:
				}
				child := grid[i][j]
				if child != nil {
					child.SetBounds(geom.Rect{Point: geom.Point{X: childX, Y: childY}, Size: geom.Size{Width: childWidth, Height: childHeight}})
				}
			}
			gridX += widths[j] + p.hSpacing
		}
		gridY += heights[i] + p.vSpacing
	}
}
