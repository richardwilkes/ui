// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package ui

// ToolTip provides a string that implements the ToolTipHandler interface for simple tooltips.
type ToolTip string

// OnToolTip implements the ToolTipHandler interface.
func (tip ToolTip) OnToolTip(where Point) string {
	return string(tip)
}
