package font

import (
	"github.com/richardwilkes/toolbox/i18n"

	// #cgo pkg-config: pangocairo
	// #include <pango/pangocairo.h>
	"C"
)

// Possible font slant variants.
const (
	SlantNormal  Slant = C.PANGO_STYLE_NORMAL
	SlantOblique Slant = C.PANGO_STYLE_OBLIQUE
	SlantItalic  Slant = C.PANGO_STYLE_ITALIC
)

// Slant represents the angle that a font is skewed when drawn.
type Slant C.PangoStyle

// String returns the name of the slant.
func (s Slant) String() string {
	switch s {
	case SlantNormal:
		return i18n.Text("Normal")
	case SlantOblique:
		return i18n.Text("Oblique")
	case SlantItalic:
		return i18n.Text("Italic")
	default:
		return i18n.Text("Unknown")
	}
}
