package font

import (
	"github.com/richardwilkes/toolbox/i18n"
	// #cgo pkg-config: pangocairo
	// #include <pango/pangocairo.h>
	"C"
)

// Possible capitalizations.
const (
	CapitalizationNormal Capitalization = C.PANGO_VARIANT_NORMAL
	CapitalizationSmall  Capitalization = C.PANGO_VARIANT_SMALL_CAPS
)

// Capitalization is an enumeration of capitalization possibilities.
type Capitalization C.PangoVariant

// String returns the name of the capitalization.
func (c Capitalization) String() string {
	switch c {
	case CapitalizationNormal:
		return i18n.Text("Normal")
	case CapitalizationSmall:
		return i18n.Text("Small-Caps")
	default:
		return i18n.Text("Unknown")
	}
}
