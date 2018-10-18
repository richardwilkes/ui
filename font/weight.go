package font

import (
	// #cgo pkg-config: pangocairo
	// #include <pango/pangocairo.h>
	"C"
	"fmt"

	"github.com/richardwilkes/toolbox/i18n"
)

// Pre-defined weights.
const (
	WeightThin       Weight = C.PANGO_WEIGHT_THIN
	WeightUltraLight Weight = C.PANGO_WEIGHT_ULTRALIGHT
	WeightLight      Weight = C.PANGO_WEIGHT_LIGHT
	WeightSemiLight  Weight = C.PANGO_WEIGHT_SEMILIGHT
	WeightBook       Weight = C.PANGO_WEIGHT_BOOK
	WeightNormal     Weight = C.PANGO_WEIGHT_NORMAL
	WeightMedium     Weight = C.PANGO_WEIGHT_MEDIUM
	WeightSemiBold   Weight = C.PANGO_WEIGHT_SEMIBOLD
	WeightBold       Weight = C.PANGO_WEIGHT_BOLD
	WeightUltraBold  Weight = C.PANGO_WEIGHT_ULTRABOLD
	WeightHeavy      Weight = C.PANGO_WEIGHT_HEAVY
	WeightUltraHeavy Weight = C.PANGO_WEIGHT_ULTRAHEAVY
)

// Weight specifies the boldness of a font. It can range from 100 to 1000.
type Weight C.PangoWeight

// String returns the name of the weight, or a value from 100 to 1000 if the value doesn't match
// one of the pre-defined weights.
func (c Weight) String() string {
	switch c {
	case WeightThin:
		return i18n.Text("Thin")
	case WeightUltraLight:
		return i18n.Text("Ultra-Light")
	case WeightLight:
		return i18n.Text("Light")
	case WeightSemiLight:
		return i18n.Text("Semi-Light")
	case WeightBook:
		return i18n.Text("Book")
	case WeightNormal:
		return i18n.Text("Normal")
	case WeightMedium:
		return i18n.Text("Medium")
	case WeightSemiBold:
		return i18n.Text("Semi-Bold")
	case WeightBold:
		return i18n.Text("Bold")
	case WeightUltraBold:
		return i18n.Text("Ultra-Bold")
	case WeightHeavy:
		return i18n.Text("Heavy")
	case WeightUltraHeavy:
		return i18n.Text("Ultra-Heavy")
	default:
		return fmt.Sprintf("%d", c)
	}
}
