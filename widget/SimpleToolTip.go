// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package widget

import (
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/event"
)

type SimpleToolTip struct {
	text string
}

func NewSimpleToolTip(target ui.Widget, text string) *SimpleToolTip {
	st := &SimpleToolTip{text: text}
	target.EventHandlers().Add(event.ToolTipType, st.tooltip)
	return st
}

func (st *SimpleToolTip) tooltip(evt event.Event) {
	evt.(*event.ToolTip).SetToolTip(st.text)
}
