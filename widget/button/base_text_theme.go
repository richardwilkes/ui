package button

import (
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/font"
)

// BaseTextTheme contains the common theme elements used in all buttons that display text.
type BaseTextTheme struct {
	BaseTheme
	TextWhenLight    color.Color // The text color to use when the background is considered to be 'light'.
	TextWhenDark     color.Color // The text color to use when the background is considered to be 'dark'.
	TextWhenDisabled color.Color // The text color to use when disabled.
	Font             *font.Font  // The font to use.
}

// Init initializes the theme with its default values.
func (theme *BaseTextTheme) Init() {
	theme.BaseTheme.Init()
	theme.TextWhenLight = color.Black
	theme.TextWhenDark = color.White
	theme.TextWhenDisabled = color.Gray
	theme.Font = font.System
}
