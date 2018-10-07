package radiobutton

import (
	"github.com/richardwilkes/ui/widget/button"
)

var (
	// StdTheme is the theme all new RadioButtons get by default.
	StdTheme = NewTheme()
)

// Theme contains the theme elements for RadioButtons.
type Theme struct {
	button.BaseTextTheme
	HorizontalGap float64 // The gap between the radio button graphic and its label.
}

// NewTheme creates a new radio button theme.
func NewTheme() *Theme {
	theme := &Theme{}
	theme.Init()
	return theme
}

// Init initializes the theme with its default values.
func (theme *Theme) Init() {
	theme.BaseTextTheme.Init()
	theme.HorizontalGap = 4
}
