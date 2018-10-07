package tooltip

import (
	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/border"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/widget/label"
)

// SetText sets a text tooltip on the target.
func SetText(target ui.Widget, text string) {
	tip := label.New(text)
	tip.SetBackground(color.LightYellow)
	tip.SetBorder(border.NewCompound(border.NewLine(color.DarkGray, geom.NewUniformInsets(1)), border.NewEmpty(geom.Insets{Top: 2, Left: 4, Bottom: 2, Right: 4})))
	target.EventHandlers().Add(event.ToolTipType, func(evt event.Event) { evt.(*Event).SetToolTip(tip) })
}
