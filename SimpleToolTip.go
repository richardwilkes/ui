// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

import (
	"github.com/richardwilkes/ui/event"
)

// SimpleToolTip provides an easy way to add a static tooltip to a widget.
type SimpleToolTip struct {
	// Text is the text that will be used for the tooltip.
	Text string
}

// NewSimpleToolTip adds a tooltip to the target.
func NewSimpleToolTip(target Widget, text string) *SimpleToolTip {
	st := &SimpleToolTip{Text: text}
	target.EventHandlers().Add(event.ToolTipType, st.tooltip)
	return st
}

func (st *SimpleToolTip) tooltip(evt event.Event) {
	evt.(*event.ToolTip).SetToolTip(st.Text)
}
