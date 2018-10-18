package font

import (
	// #cgo pkg-config: pangocairo
	// #include <pango/pangocairo.h>
	"C"

	"github.com/richardwilkes/toolbox/i18n"
)

// Possible stretch values.
const (
	StretchUltraCondensed Stretch = C.PANGO_STRETCH_ULTRA_CONDENSED
	StretchExtraCondensed Stretch = C.PANGO_STRETCH_EXTRA_CONDENSED
	StretchCondensed      Stretch = C.PANGO_STRETCH_CONDENSED
	StretchSemiCondensed  Stretch = C.PANGO_STRETCH_SEMI_CONDENSED
	StretchNormal         Stretch = C.PANGO_STRETCH_NORMAL
	StretchSemiExpanded   Stretch = C.PANGO_STRETCH_SEMI_EXPANDED
	StretchExpanded       Stretch = C.PANGO_STRETCH_EXPANDED
	StretchExtraExpanded  Stretch = C.PANGO_STRETCH_EXTRA_EXPANDED
	StretchUltraExpanded  Stretch = C.PANGO_STRETCH_ULTRA_EXPANDED
)

// Stretch is an enumeration of stretch possibilities.
type Stretch C.PangoStretch

// String returns the name of the stretch.
func (s Stretch) String() string {
	switch s {
	case StretchUltraCondensed:
		return i18n.Text("Ultra-Condensed")
	case StretchExtraCondensed:
		return i18n.Text("Extra-Condensed")
	case StretchCondensed:
		return i18n.Text("Condensed")
	case StretchSemiCondensed:
		return i18n.Text("Semi-Condensed")
	case StretchNormal:
		return i18n.Text("Normal")
	case StretchSemiExpanded:
		return i18n.Text("Semi-Expanded")
	case StretchExpanded:
		return i18n.Text("Expanded")
	case StretchExtraExpanded:
		return i18n.Text("Extra-Expanded")
	case StretchUltraExpanded:
		return i18n.Text("Ultra-Expanded")
	default:
		return "Unknown"
	}
}
