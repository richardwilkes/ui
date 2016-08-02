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
	"fmt"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/geom"
	"github.com/richardwilkes/ui/layout"
	"github.com/richardwilkes/ui/theme"
	"math"
	"strings"
	"time"
)

// TextField provides a single-line text input control.
type TextField struct {
	Block
	text            string
	watermark       string
	Theme           *theme.TextField // The theme the text field will use to draw itself.
	selectionStart  int
	selectionEnd    int
	selectionAnchor int
	forceShowUntil  time.Time
	scrollOffset    float32
	align           draw.Alignment
	showCursor      bool
	pending         bool
	extendByWord    bool
}

// NewTextField creates a new, empty, text field.
func NewTextField() *TextField {
	field := &TextField{}
	field.Theme = theme.StdTextField
	field.SetBackground(color.TextBackground)
	field.SetBorder(field.Theme.Border)
	field.SetFocusable(true)
	field.SetSizer(field)
	handlers := field.EventHandlers()
	handlers.Add(event.PaintType, field.paint)
	handlers.Add(event.FocusGainedType, field.focusGained)
	handlers.Add(event.FocusLostType, field.focusLost)
	handlers.Add(event.MouseDownType, field.mouseDown)
	return field
}

// Sizes implements Sizer
func (field *TextField) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	if hint.Width != layout.NoHint {
		if hint.Width < field.Theme.MinimumTextWidth {
			hint.Width = field.Theme.MinimumTextWidth
		}
	}
	if hint.Height != layout.NoHint {
		if hint.Height < 1 {
			hint.Height = 1
		}
	}
	var text string
	if field.text == "" {
		text = "M"
	} else {
		text = field.text
	}
	size := field.Theme.Font.Size(text)
	// Add the descent height to allow for a more balanced vertical look
	size.Height += field.Theme.Font.Descent()
	size.GrowToInteger()
	size.ConstrainForHint(hint)
	if border := field.Border(); border != nil {
		size.AddInsets(border.Insets())
	}
	return size, size, layout.DefaultMaxSize(size)
}

func (field *TextField) paint(evt event.Event) {
	bounds := field.LocalInsetBounds()
	e := evt.(*event.Paint)
	gc := e.GC()
	gc.Save()
	defer gc.Restore()
	if !field.Enabled() && field.Theme.DisabledBackgroundColor.Alpha() > 0 {
		gc.SetFillColor(field.Theme.DisabledBackgroundColor)
		gc.FillRect(e.DirtyRect())
	}
	gc.ClipRect(bounds)
	gc.SetFont(field.Theme.Font)
	descent := field.Theme.Font.Descent()
	textTop := bounds.Y + (bounds.Height-(field.Theme.Font.Height()-descent))/2
	if field.HasSelectionRange() {
		left := bounds.X + field.scrollOffset
		if field.selectionStart > 0 {
			gc.SetFillColor(color.Text)
			pre := field.text[:field.selectionStart]
			gc.DrawString(left, textTop, pre)
			left += field.Theme.Font.Width(pre)
		}
		mid := field.text[field.selectionStart:field.selectionEnd]
		right := bounds.X + field.Theme.Font.Width(field.text[:field.selectionEnd]) + field.scrollOffset
		selRect := geom.Rect{Point: geom.Point{X: left, Y: textTop - descent}, Size: geom.Size{Width: right - left, Height: field.Theme.Font.Height() + descent}}
		if field.Focused() {
			gc.SetFillColor(color.SelectedTextBackground)
			gc.FillRect(selRect)
		} else {
			gc.SetStrokeColor(color.SelectedTextBackground)
			gc.SetStrokeWidth(2)
			selRect.InsetUniform(0.5)
			gc.StrokeRect(selRect)
		}
		gc.SetFillColor(color.SelectedText)
		gc.DrawString(left, textTop, mid)
		if field.selectionStart < len(field.text) {
			gc.SetFillColor(color.Text)
			gc.DrawString(right, textTop, field.text[field.selectionEnd:])
		}
	} else if field.text == "" {
		if field.watermark != "" {
			gc.SetFillColor(color.Gray)
			gc.DrawString(bounds.X, textTop, field.watermark)
		}
	} else {
		gc.SetFillColor(color.Text)
		gc.DrawString(bounds.X, textTop, field.text)
	}
	if !field.HasSelectionRange() && field.Focused() {
		if field.showCursor {
			var cursorColor color.Color
			if field.Background().Luminance() > 0.6 {
				cursorColor = color.Black
			} else {
				cursorColor = color.White
			}
			x := bounds.X + field.Theme.Font.Width(field.text[:field.selectionEnd]) + field.scrollOffset
			gc.SetStrokeColor(cursorColor)
			gc.StrokeLine(x, textTop-descent, x, textTop+field.Theme.Font.Height()+descent-1)
		}
		field.scheduleBlink()
	}
}

func (field *TextField) scheduleBlink() {
	if !field.pending && field.Focused() {
		field.pending = true
		event.InvokeAfter(field.blink, field.Theme.BlinkRate)
	}
}

func (field *TextField) blink() {
	field.pending = false
	if time.Now().After(field.forceShowUntil) {
		field.showCursor = !field.showCursor
		field.Repaint()
	}
	field.scheduleBlink()
}

func (field *TextField) focusGained(evt event.Event) {
	field.SetBorder(field.Theme.FocusBorder)
	field.showCursor = true
	field.Repaint()
}

func (field *TextField) focusLost(evt event.Event) {
	field.SetBorder(field.Theme.Border)
	field.Repaint()
}

func (field *TextField) mouseDown(evt event.Event) {
	field.Window().SetFocus(field)
	e := evt.(*event.MouseDown)
	if e.Button() == event.LeftButton {
		field.extendByWord = false
		switch e.Clicks() {
		case 2:
		case 3:
			field.SelectAll()
		default:
		}
	} else if e.Button() == event.RightButton {
		fmt.Println("right click")
	}
}

// Text returns the content of the field.
func (field *TextField) Text() string {
	return field.text
}

// SetText sets the content of the field. Returns true if a modification was made.
func (field *TextField) SetText(text string) bool {
	text = sanitize(text)
	if field.text != text {
		field.text = text
		field.SetSelectionToEnd()
		field.Repaint()
		event.Dispatch(event.NewModified(field))
		return true
	}
	return false
}

func sanitize(text string) string {
	return strings.NewReplacer("\n", "", "\r", "").Replace(text)
}

// Watermark returns the current watermark, if any.
func (field *TextField) Watermark() string {
	return field.watermark
}

// SetWatermark sets the watermark. The watermark is used to give the user a hint about what the
// field is for when it is empty.
func (field *TextField) SetWatermark(text string) {
	field.watermark = text
	field.Repaint()
}

// SelectedText returns the currently selected text.
func (field *TextField) SelectedText() string {
	return field.text[field.selectionStart:field.selectionEnd]
}

// HasSelectionRange returns true is a selection range is currently present.
func (field *TextField) HasSelectionRange() bool {
	return field.selectionStart < field.selectionEnd
}

// SelectionCount returns the number of characters currently selected.
func (field *TextField) SelectionCount() int {
	return field.selectionEnd - field.selectionStart
}

// Selection returns the current start and end selection indexes.
func (field *TextField) Selection() (start, end int) {
	return field.selectionStart, field.selectionEnd
}

// SelectAll selects all of the text in the field.
func (field *TextField) SelectAll() {
	field.SetSelection(0, len(field.text))
}

// SetSelectionToStart moves the cursor to the beginning of the text and removes any range that may
// have been present.
func (field *TextField) SetSelectionToStart() {
	field.SetSelection(0, 0)
}

// SetSelectionToEnd moves the cursor to the end of the text and removes any range that may have
// been present.
func (field *TextField) SetSelectionToEnd() {
	field.SetSelection(math.MaxInt64, math.MaxInt64)
}

// SetSelectionTo moves the cursor to the specified index and removes any range that may have been
// present.
func (field *TextField) SetSelectionTo(pos int) {
	field.SetSelection(pos, pos)
}

// SetSelection sets the start and end range of the selection. Values beyond either end will be
// constrained to the appropriate end. Likewise, an end value less than the start value will be
// treated as if the start and end values were the same.
func (field *TextField) SetSelection(start, end int) {
	field.setSelection(start, end, start)
}

func (field *TextField) setSelection(start, end, anchor int) {
	length := len(field.text)
	if start < 0 {
		start = 0
	} else if start > length {
		start = length
	}
	if end < start {
		end = start
	} else if end > length {
		end = length
	}
	if anchor < start {
		anchor = start
	} else if anchor > end {
		anchor = end
	}
	if field.selectionStart != start || field.selectionEnd != end || field.selectionAnchor != anchor {
		field.selectionStart = start
		field.selectionEnd = end
		field.selectionAnchor = anchor
		field.forceShowUntil = time.Now().Add(field.Theme.BlinkRate)
		field.showCursor = true
		field.Repaint()
		field.ScrollIntoView()
		field.autoScroll()
	}
}

// ScrollIntoView scrolls the insertion cursor into view.
func (field *TextField) ScrollIntoView() {
	// RAW: Implement
}

func (field *TextField) autoScroll() {
	// RAW: Implement
}
