package checkbox

import (
	"github.com/richardwilkes/ui/widget/button"
)

var (
	// StdCheckBox is the theme all new CheckBoxes get by default.
	StdCheckBox = NewTheme()
)

// Theme contains the theme elements for CheckBoxes.
type Theme struct {
	button.BaseTextTheme
	HorizontalGap float64 // The gap between the checkbox graphic and its label.
}

// NewTheme creates a new checkbox theme.
func NewTheme() *Theme {
	theme := &Theme{}
	theme.Init()
	return theme
}

// Init initializes the theme with its default values.
func (theme *Theme) Init() {
	theme.BaseTextTheme.Init()
	theme.CornerRadius = 4
	theme.HorizontalGap = 4
}
