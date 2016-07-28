// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package event

// Possible KeyMask values.
const (
	CapsLockKeyMask KeyMask = 1 << iota
	ShiftKeyMask
	ControlKeyMask
	OptionKeyMask
	CommandKeyMask   // On platforms that don't have a distinct command key, this will also be set if the Control key is pressed.
	NonStickyKeyMask = ShiftKeyMask | ControlKeyMask | OptionKeyMask | CommandKeyMask
	AllKeyMask       = CapsLockKeyMask | NonStickyKeyMask
)

// KeyMask contains flags indicating which modifier keys were down when an event occurred.
type KeyMask int
