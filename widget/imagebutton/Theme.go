package imagebutton

import (
	"github.com/richardwilkes/ui/widget/button"
)

var (
	// StdImageButton is the theme all new ImageButtons get by default.
	StdImageButton = NewTheme()
)

// Theme contains the theme elements for ImageButtons.
type Theme struct {
	button.BaseTheme
	HorizontalMargin float64 // The margin on the left and right side of the image.
	VerticalMargin   float64 // The margin on the top and bottom of the image.
}

// NewTheme creates a new image button theme.
func NewTheme() *Theme {
	theme := &Theme{}
	theme.Init()
	return theme
}

// Init initializes the theme with its default values.
func (theme *Theme) Init() {
	theme.BaseTheme.Init()
	theme.HorizontalMargin = 4
	theme.VerticalMargin = 4
}
