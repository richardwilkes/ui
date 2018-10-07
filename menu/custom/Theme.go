package custom

import (
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/font"
)

var (
	// StdTheme is the theme all new MenuItems get by default.
	StdTheme = NewTheme()
)

// Theme contains the theme elements for MenuItems.
type Theme struct {
	HMargin               float64     // The amount of horizontal space on each edge.
	VMargin               float64     // The amount of horizontal space on each edge.
	KeySpacing            float64     // The amount of space between the title and the key binding.
	TitleFont             *font.Font  // The font to use for the title.
	KeyFont               *font.Font  // The font to use for the key binding.
	Background            color.Color // The color to use for the background.
	HighlightedBackground color.Color // The color to use for the background when highlighted.
	TextWhenLight         color.Color // The text color to use when the background is considered to be 'light'.
	TextWhenDark          color.Color // The text color to use when the background is considered to be 'dark'.
	TextWhenDisabled      color.Color // The text color to use when disabled.
	DisabledAdjustment    float64     // The amount to adjust the background brightness when disabled.
}

// NewTheme creates a new MenuItem theme.
func NewTheme() *Theme {
	theme := &Theme{}
	theme.Init()
	return theme
}

// Init initializes the theme with its default values.
func (theme *Theme) Init() {
	theme.HMargin = 4
	theme.VMargin = 2
	theme.KeySpacing = 10
	theme.TitleFont = font.Menu
	theme.KeyFont = font.MenuCmdKey
	theme.Background = color.Background
	theme.HighlightedBackground = color.KeyboardFocus
	theme.TextWhenLight = color.Black
	theme.TextWhenDark = color.White
	theme.TextWhenDisabled = color.Gray
	theme.DisabledAdjustment = -0.05
}
