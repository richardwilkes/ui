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
	// #include "MenuTypes.h"
	"C"
)

// PlatformMenu represents the underlying platform menu object.
type PlatformMenu C.platformMenu

// PlatformItem represents the underlying platform menu item object.
type PlatformItem C.platformItem
