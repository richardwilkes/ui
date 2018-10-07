package button

var (
	// StdButton is the theme all new Buttons get by default.
	StdButton = NewTheme()
)

// Theme contains the theme elements for Buttons.
type Theme struct {
	BaseTextTheme
	HorizontalMargin float64 // The margin on the left and right side of the text.
	VerticalMargin   float64 // The margin on the top and bottom of the text.
	MinimumTextWidth float64 // The minimum space to permit for text.
}

// NewTheme creates a new button theme.
func NewTheme() *Theme {
	theme := &Theme{}
	theme.Init()
	return theme
}

// Init initializes the theme with its default values.
func (theme *Theme) Init() {
	theme.BaseTextTheme.Init()
	theme.HorizontalMargin = 8
	theme.VerticalMargin = 1
	theme.MinimumTextWidth = 10
}
