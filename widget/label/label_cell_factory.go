package label

import (
	"fmt"
	"reflect"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/border"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/font"
)

// CellFactory provides simple text cells.
type CellFactory struct {
	// Height is returned when CellHeight() is called.
	Height float64
}

// CellHeight implements the widget.CellFactory interface.
func (f *CellFactory) CellHeight() float64 {
	return f.Height
}

// CreateCell implements the widget.CellFactory interface.
func (f *CellFactory) CreateCell(owner ui.Widget, element interface{}, index int, selected, focused bool) ui.Widget {
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
