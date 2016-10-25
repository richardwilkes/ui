// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package tooltip

import (
	"github.com/richardwilkes/geom"
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
