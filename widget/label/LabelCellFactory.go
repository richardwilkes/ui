// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package label

import (
	"fmt"
	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/border"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/font"
	"reflect"
)

// LabelCellFactory provides simple text cells.
type LabelCellFactory struct {
	// Height is returned when CellHeight() is called.
	Height float64
}

// CellHeight implements the CellFactory interface.
func (f *LabelCellFactory) CellHeight() float64 {
	return f.Height
}

// CreateCell implements the CellFactory interface.
func (f *LabelCellFactory) CreateCell(owner ui.Widget, element interface{}, index int, selected, focused bool) ui.Widget {
	var text string
	switch v := element.(type) {
	case string:
		text = v
	case fmt.Stringer:
		text = v.String()
	default:
		text = reflect.TypeOf(element).String()
	}
	label := NewWithFont(text, font.Views)
	if selected {
		label.SetBackground(color.SelectedTextBackground)
		label.SetForeground(color.SelectedText)
	}
	label.SetBorder(border.NewEmpty(geom.NewHorizontalInsets(4)))
	return label
}
