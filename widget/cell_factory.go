package widget

import (
	"github.com/richardwilkes/ui"
)

// CellFactory defines methods all cell factories must implement.
type CellFactory interface {
	// CellHeight returns the height to use for the cells. A value less than 1 indicates that each
	// cell's height may be different.
	CellHeight() float64

	// CreateCell creates a new cell for 'owner' using 'element' as the content. 'index' indicates
	// which row the element came from. 'selected' indicates the cell should be created in its
	// selected state. 'focused' indicates the cell should be created in its focused state.
	CreateCell(owner ui.Widget, element interface{}, index int, selected, focused bool) ui.Widget
}
