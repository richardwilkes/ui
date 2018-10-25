package display

import (
	"github.com/richardwilkes/toolbox/xmath/geom"
)

// MainBounds returns the bounds of the main display.
func MainBounds() geom.Rect {
	return platformMainDisplayBounds()
}
