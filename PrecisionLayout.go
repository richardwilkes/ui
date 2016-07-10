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
	"math"
)

// PrecisionLayout lays out the children of its target based on the PrecisionData assigned to each child.
type PrecisionLayout struct {
	rows                int
	Columns             int
	HorizontalSpacing   float32
	VerticalSpacing     float32
	HorizontalAlignment Alignment
	VerticalAlignment   Alignment
	EqualColumns        bool
}

// NewPrecisionLayout creates a new PrecisionLayout.
func NewPrecisionLayout() *PrecisionLayout {
	return &PrecisionLayout{Columns: 1, HorizontalSpacing: 4, VerticalSpacing: 2, HorizontalAlignment: AlignStart, VerticalAlignment: AlignStart}
}

// SetColumns is a convenience for setting the number of columns while chaining calls together.
func (p *PrecisionLayout) SetColumns(columns int) *PrecisionLayout {
	p.Columns = columns
	return p
}

// SetEqualColumns is a convenience for setting equal columns while chaining calls together.
func (p *PrecisionLayout) SetEqualColumns(equal bool) *PrecisionLayout {
	p.EqualColumns = equal
	return p
}

// SetHorizontalSpacing is a convenience for setting the horizontal spacing while chaining calls
// together.
func (p *PrecisionLayout) SetHorizontalSpacing(spacing float32) *PrecisionLayout {
	p.HorizontalSpacing = spacing
	return p
}

// SetVerticalSpacing is a convenience for setting the vertical spacing while chaining calls
// together.
func (p *PrecisionLayout) SetVerticalSpacing(spacing float32) *PrecisionLayout {
	p.VerticalSpacing = spacing
	return p
}

// SetHorizontalAlignment is a convenience for setting the horizontal alignment while chaining
// calls together.
func (p *PrecisionLayout) SetHorizontalAlignment(alignment Alignment) *PrecisionLayout {
	p.HorizontalAlignment = alignment
	return p
}

// SetVerticalAlignment is a convenience for setting the vertical alignment while chaining calls
// together.
func (p *PrecisionLayout) SetVerticalAlignment(alignment Alignment) *PrecisionLayout {
	p.VerticalAlignment = alignment
	return p
}

// ComputeSizes implements the Layout interface
func (p *PrecisionLayout) ComputeSizes(target *Block, hint Size) (min, pref, max Size) {
	min = p.layout(target, Point{}, NoLayoutHintSize, false, true)
	pref = p.layout(target, Point{}, NoLayoutHintSize, false, false)
	insets := target.Insets()
	min.AddInsets(insets)
	pref.AddInsets(insets)
	return min, pref, DefaultLayoutMaxSize(pref)
}

// Layout implements the Layout interface
func (p *PrecisionLayout) Layout(target *Block) {
	insets := target.Insets()
	hint := target.Bounds().Size
	hint.SubtractInsets(insets)
	p.layout(target, Point{X: insets.Left, Y: insets.Top}, hint, true, false)
}

func (p *PrecisionLayout) layout(target *Block, location Point, hint Size, move, useMinimumSize bool) Size {
	var totalSize Size
	if p.Columns > 0 {
		children := p.prepChildren(target, useMinimumSize)
		if len(children) > 0 {
			grid := p.buildGrid(children)
			widths := p.adjustColumnWidths(hint.Width, grid)
			p.wrap(hint.Width, grid, widths, useMinimumSize)
			heights := p.adjustRowHeights(hint.Height, grid)
			totalSize.Width += p.HorizontalSpacing * float32(p.Columns-1)
			totalSize.Height += p.VerticalSpacing * float32(p.rows-1)
			for i := 0; i < p.Columns; i++ {
				totalSize.Width += widths[i]
			}
			for i := 0; i < p.rows; i++ {
				totalSize.Height += heights[i]
			}
			if move {
				if totalSize.Width < hint.Width {
					if p.HorizontalAlignment == AlignMiddle {
						location.X += (hint.Width - totalSize.Width) / 2
					} else if p.HorizontalAlignment == AlignEnd {
						location.X += hint.Width - totalSize.Width
					}
				}
				if totalSize.Height < hint.Height {
					if p.VerticalAlignment == AlignMiddle {
						location.Y += (hint.Height - totalSize.Height) / 2
					} else if p.VerticalAlignment == AlignEnd {
						location.Y += hint.Height - totalSize.Height
					}
				}
				p.positionChildren(location, grid, widths, heights)
			}
		}
	}
	return totalSize
}

func (p *PrecisionLayout) prepChildren(target *Block, useMinimumSize bool) []*Block {
	children := target.Children()
	for _, child := range children {
		var layoutData *PrecisionData
		data := child.LayoutData()
		switch v := data.(type) {
		case *PrecisionData:
			layoutData = v
		default:
			layoutData = NewPrecisionData()
			child.SetLayoutData(layoutData)
		}
		layoutData.computeCacheSize(child, NoLayoutHintSize, useMinimumSize)
	}
	return children
}

func (p *PrecisionLayout) buildGrid(children []*Block) [][]*Block {
	var grid [][]*Block
	var row, column int
	p.rows = 0
	for _, child := range children {
		data := child.LayoutData().(*PrecisionData)
		hSpan := MaxInt(1, MinInt(data.HorizontalSpan, p.Columns))
		vSpan := MaxInt(1, data.VerticalSpan)
		for {
			lastRow := row + vSpan
			if lastRow >= len(grid) {
				grid = append(grid, make([]*Block, p.Columns))
			}
			for column < p.Columns && grid[row][column] != nil {
				column++
			}
			endCount := column + hSpan
			if endCount <= p.Columns {
				index := column
				for index < endCount && grid[row][index] == nil {
					index++
				}
				if index == endCount {
					break
				}
				column = index
			}
			if column+hSpan >= p.Columns {
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

func (p *PrecisionLayout) adjustColumnWidths(width float32, grid [][]*Block) []float32 {
	availableWidth := width - p.HorizontalSpacing*float32(p.Columns-1)
	expandCount := 0
	widths := make([]float32, p.Columns)
	minWidths := make([]float32, p.Columns)
	expandColumn := make([]bool, p.Columns)
	for j := 0; j < p.Columns; j++ {
		for i := 0; i < p.rows; i++ {
			data := p.getData(grid, i, j, true)
			if data != nil {
				hSpan := MaxInt(1, MinInt(data.HorizontalSpan, p.Columns))
				if hSpan == 1 {
					w := data.cacheSize.Width
					if widths[j] < w {
						widths[j] = w
					}
					if data.HorizontalGrab {
						if !expandColumn[j] {
							expandCount++
						}
						expandColumn[j] = true
					}
					minimumWidth := data.cacheMinWidth
					if !data.HorizontalGrab || minimumWidth != 0 {
						if !data.HorizontalGrab || minimumWidth == NoLayoutHint {
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
				hSpan := MaxInt(1, MinInt(data.HorizontalSpan, p.Columns))
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
					if data.HorizontalGrab && spanExpandCount == 0 {
						expandCount++
						expandColumn[j] = true
					}
					w := data.cacheSize.Width - spanWidth - float32(hSpan-1)*p.HorizontalSpacing
					if w > 0 {
						if p.EqualColumns {
							equalWidth := (w + spanWidth) / float32(hSpan)
							for k := 0; k < hSpan; k++ {
								if widths[j-k] < equalWidth {
									widths[j-k] = equalWidth
								}
							}
						} else {
							if spanExpandCount == 0 {
								widths[j] += w
							} else {
								delta := w / float32(spanExpandCount)
								for k := 0; k < hSpan; k++ {
									if expandColumn[j-k] {
										widths[j-k] += delta
									}
								}
							}
						}
					}
					minimumWidth := data.cacheMinWidth
					if !data.HorizontalGrab || minimumWidth != 0 {
						if !data.HorizontalGrab || minimumWidth == NoLayoutHint {
							w = data.cacheSize.Width
						} else {
							w = minimumWidth
						}
						w -= spanMinWidth + float32(hSpan-1)*p.HorizontalSpacing
						if w > 0 {
							if spanExpandCount == 0 {
								minWidths[j] += w
							} else {
								delta := w / float32(spanExpandCount)
								for k := 0; k < hSpan; k++ {
									if expandColumn[j-k] {
										minWidths[j-k] += delta
									}
								}
							}
						}
					}
				}
			}
		}
	}
	if p.EqualColumns {
		var minColumnWidth, columnWidth float32
		for i := 0; i < p.Columns; i++ {
			if minColumnWidth < minWidths[i] {
				minColumnWidth = minWidths[i]
			}
			if columnWidth < widths[i] {
				columnWidth = widths[i]
			}
		}
		if width != NoLayoutHint && expandCount != 0 {
			columnWidth = MaxFloat32(minColumnWidth, availableWidth/float32(p.Columns))
		}
		for i := 0; i < p.Columns; i++ {
			expandColumn[i] = expandCount > 0
			widths[i] = columnWidth
		}
	} else {
		if width != NoLayoutHint && expandCount > 0 {
			var totalWidth float32
			for i := 0; i < p.Columns; i++ {
				totalWidth += widths[i]
			}
			c := expandCount
			delta := (availableWidth - totalWidth) / float32(c)
			for math.Abs(float64(totalWidth-availableWidth)) > 0.01 {
				for j := 0; j < p.Columns; j++ {
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
				for j := 0; j < p.Columns; j++ {
					for i := 0; i < p.rows; i++ {
						data := p.getData(grid, i, j, false)
						if data != nil {
							hSpan := MaxInt(1, MinInt(data.HorizontalSpan, p.Columns))
							if hSpan > 1 {
								minimumWidth := data.cacheMinWidth
								if !data.HorizontalGrab || minimumWidth != 0 {
									var spanWidth float32
									spanExpandCount := 0
									for k := 0; k < hSpan; k++ {
										spanWidth += widths[j-k]
										if expandColumn[j-k] {
											spanExpandCount++
										}
									}
									var w float32
									if !data.HorizontalGrab || minimumWidth == NoLayoutHint {
										w = data.cacheSize.Width
									} else {
										w = minimumWidth
									}
									w -= spanWidth + float32(hSpan-1)*p.HorizontalSpacing
									if w > 0 {
										if spanExpandCount == 0 {
											widths[j] += w
										} else {
											delta2 := w / float32(spanExpandCount)
											for k := 0; k < hSpan; k++ {
												if expandColumn[j-k] {
													widths[j-k] += delta2
												}
											}
										}
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
				for i := 0; i < p.Columns; i++ {
					totalWidth += widths[i]
				}
				delta = (availableWidth - totalWidth) / float32(c)
			}
		}
	}
	return widths
}

func (p *PrecisionLayout) getData(grid [][]*Block, row, column int, first bool) *PrecisionData {
	block := grid[row][column]
	if block != nil {
		data := block.LayoutData().(*PrecisionData)
		hSpan := MaxInt(1, MinInt(data.HorizontalSpan, p.Columns))
		vSpan := MaxInt(1, data.VerticalSpan)
		var i, j int
		if first {
			i = row + vSpan - 1
			j = column + hSpan - 1
		} else {
			i = row - vSpan + 1
			j = column - hSpan + 1
		}
		if i >= 0 && i < p.rows {
			if j >= 0 && j < p.Columns {
				if block == grid[i][j] {
					return data
				}
			}
		}
	}
	return nil
}

func (p *PrecisionLayout) wrap(width float32, grid [][]*Block, widths []float32, useMinimumSize bool) {
	if width != NoLayoutHint {
		for j := 0; j < p.Columns; j++ {
			for i := 0; i < p.rows; i++ {
				data := p.getData(grid, i, j, false)
				if data != nil {
					if data.SizeHint.Height == NoLayoutHint {
						hSpan := MaxInt(1, MinInt(data.HorizontalSpan, p.Columns))
						var currentWidth float32
						for k := 0; k < hSpan; k++ {
							currentWidth += widths[j-k]
						}
						currentWidth += float32(hSpan-1) * p.HorizontalSpacing
						if currentWidth != data.cacheSize.Width && data.HorizontalAlignment == AlignFill || data.cacheSize.Width > currentWidth {
							data.computeCacheSize(grid[i][j], Size{Width: MaxFloat32(data.cacheMinWidth, currentWidth), Height: NoLayoutHint}, useMinimumSize)
							minimumHeight := data.MinSize.Height
							if data.VerticalGrab && minimumHeight > 0 && data.cacheSize.Height < minimumHeight {
								data.cacheSize.Height = minimumHeight
							}
						}
					}
				}
			}
		}
	}
}

func (p *PrecisionLayout) adjustRowHeights(height float32, grid [][]*Block) []float32 {
	availableHeight := height - p.VerticalSpacing*float32(p.rows-1)
	expandCount := 0
	heights := make([]float32, p.rows)
	minHeights := make([]float32, p.rows)
	expandRow := make([]bool, p.rows)
	for i := 0; i < p.rows; i++ {
		for j := 0; j < p.Columns; j++ {
			data := p.getData(grid, i, j, true)
			if data != nil {
				vSpan := MaxInt(1, MinInt(data.VerticalSpan, p.rows))
				if vSpan == 1 {
					h := data.cacheSize.Height
					if heights[i] < h {
						heights[i] = h
					}
					if data.VerticalGrab {
						if !expandRow[i] {
							expandCount++
						}
						expandRow[i] = true
					}
					minimumHeight := data.MinSize.Height
					if !data.VerticalGrab || minimumHeight != 0 {
						var h float32
						if !data.VerticalGrab || minimumHeight == NoLayoutHint {
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
		for j := 0; j < p.Columns; j++ {
			data := p.getData(grid, i, j, false)
			if data != nil {
				vSpan := MaxInt(1, MinInt(data.VerticalSpan, p.rows))
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
					if data.VerticalGrab && spanExpandCount == 0 {
						expandCount++
						expandRow[i] = true
					}
					h := data.cacheSize.Height - spanHeight - float32(vSpan-1)*p.VerticalSpacing
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
					minimumHeight := data.MinSize.Height
					if !data.VerticalGrab || minimumHeight != 0 {
						var h float32
						if !data.VerticalGrab || minimumHeight == NoLayoutHint {
							h = data.cacheSize.Height
						} else {
							h = minimumHeight
						}
						h -= spanMinHeight + float32(vSpan-1)*p.VerticalSpacing
						if h > 0 {
							if spanExpandCount == 0 {
								minHeights[i] += h
							} else {
								delta := h / float32(spanExpandCount)
								for k := 0; k < vSpan; k++ {
									if expandRow[i-k] {
										minHeights[i-k] += delta
									}
								}
							}
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
		for math.Abs(float64(totalHeight-availableHeight)) > 0.01 {
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
				for j := 0; j < p.Columns; j++ {
					data := p.getData(grid, i, j, false)
					if data != nil {
						vSpan := MaxInt(1, MinInt(data.VerticalSpan, p.rows))
						if vSpan > 1 {
							minimumHeight := data.MinSize.Height
							if !data.VerticalGrab || minimumHeight != 0 {
								var spanHeight float32
								spanExpandCount := 0
								for k := 0; k < vSpan; k++ {
									spanHeight += heights[i-k]
									if expandRow[i-k] {
										spanExpandCount++
									}
								}
								var h float32
								if !data.VerticalGrab || minimumHeight == NoLayoutHint {
									h = data.cacheSize.Height
								} else {
									h = minimumHeight
								}
								h -= spanHeight + float32(vSpan-1)*p.VerticalSpacing
								if h > 0 {
									if spanExpandCount == 0 {
										heights[i] += h
									} else {
										delta2 := h / float32(spanExpandCount)
										for k := 0; k < vSpan; k++ {
											if expandRow[i-k] {
												heights[i-k] += delta2
											}
										}
									}
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

func (p *PrecisionLayout) positionChildren(location Point, grid [][]*Block, widths []float32, heights []float32) {
	gridY := location.Y
	for i := 0; i < p.rows; i++ {
		gridX := location.X
		for j := 0; j < p.Columns; j++ {
			data := p.getData(grid, i, j, true)
			if data != nil {
				hSpan := MaxInt(1, MinInt(data.HorizontalSpan, p.Columns))
				vSpan := MaxInt(1, data.VerticalSpan)
				var cellWidth, cellHeight float32
				for k := 0; k < hSpan; k++ {
					cellWidth += widths[j+k]
				}
				for k := 0; k < vSpan; k++ {
					cellHeight += heights[i+k]
				}
				cellWidth += p.HorizontalSpacing * float32(hSpan-1)
				childX := gridX
				childWidth := MinFloat32(data.cacheSize.Width, cellWidth)
				switch data.HorizontalAlignment {
				case AlignMiddle:
					childX += MaxFloat32(0, (cellWidth-childWidth)/2)
				case AlignEnd:
					childX += MaxFloat32(0, cellWidth-childWidth)
				case AlignFill:
					childWidth = cellWidth
				default:
				}
				cellHeight += p.VerticalSpacing * float32(vSpan-1)
				childY := gridY
				childHeight := MinFloat32(data.cacheSize.Height, cellHeight)
				switch data.VerticalAlignment {
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
			gridX += widths[j] + p.HorizontalSpacing
		}
		gridY += heights[i] + p.VerticalSpacing
	}
}
