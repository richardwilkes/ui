// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package menu

import (
	"github.com/richardwilkes/ui/event"
)

var (
	parentTarget event.Target
)

// ParentTarget returns the value that will be returned on calls to an item's ParentTarget()
// method.
func ParentTarget() event.Target {
	return parentTarget
}

// SetParentTarget sets the value that should be returned on calls to an item's ParentTarget()
// method.
func SetParentTarget(target event.Target) {
	parentTarget = target
}
