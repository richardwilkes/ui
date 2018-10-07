package scrollbar

import (
	"time"

	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/draw"
)

var (
	// StdTheme is the theme all new ScrollBars get by default.
	StdTheme = NewTheme()
)

// Theme contains the theme elements for ScrollBars.
type Theme struct {
	InitialRepeatDelay    time.Duration // The amount of time to wait before triggering the first repeating event.
	RepeatDelay           time.Duration // The amount of time to wait before triggering a repeating event.
	Background            color.Color   // The background color when enabled but not pressed or focused.
	BackgroundWhenPressed color.Color   // The background color when enabled and pressed.
	MarkWhenLight         color.Color   // The color to use for control marks when the background is considered to be 'light'.
	MarkWhenDark          color.Color   // The color to use for control marks when the background is considered to be 'dark'.
	MarkWhenDisabled      color.Color   // The color to use for control marks when disabled.
	GradientAdjustment    float64       // The amount to vary the color when creating the background gradient.
	DisabledAdjustment    float64       // The amount to adjust the background brightness when disabled.
	OutlineAdjustment     float64       // The amount to adjust the background brightness when using it to draw the button outline.
	Size                  float64       // The height of a horizontal scrollbar or the width of a vertical scrollbar.
}

// NewTheme creates a new image button theme.
func NewTheme() *Theme {
	theme := &Theme{}
	theme.Init()
	return theme
}

// Init initializes the theme with its default values.
func (theme *Theme) Init() {
	theme.InitialRepeatDelay = time.Millisecond * 250
	theme.RepeatDelay = time.Millisecond * 75
	theme.Background = color.White
	theme.BackgroundWhenPressed = color.KeyboardFocus
	theme.MarkWhenLight = color.Black
	theme.MarkWhenDark = color.White
	theme.MarkWhenDisabled = color.Gray
	theme.GradientAdjustment = 0.15
	theme.DisabledAdjustment = -0.05
	theme.OutlineAdjustment = -0.5
	theme.Size = 16
}

// Gradient returns a gradient for the specified color.
func (theme *Theme) Gradient(base color.Color) *draw.Gradient {
	return draw.NewEvenlySpacedGradient(base.AdjustBrightness(theme.GradientAdjustment), base.AdjustBrightness(-theme.GradientAdjustment))
}
