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
	"github.com/richardwilkes/ui/keys"
)

// Item represents individual actions that can be issued from a menu.
type Item interface {
	event.Target
	// Title returns this item's title.
	Title() string
	// KeyCode returns the key code that can be used to trigger this item. A value of 0 indicates no
	// key is attached.
	KeyCode() int
	// KeyModifiers returns the key modifiers that are required to trigger this item.
	KeyModifiers() keys.Modifiers
	// SubMenu returns a sub-menu attached to this item or nil.
	SubMenu() Menu
	// Enabled returns true if this item is enabled.
	Enabled() bool
	// Dispose releases any operating system resources associated with this item.
	Dispose()
}

var (
	// NewItem creates a new item with no key accelerator.
	NewItem func(title string, handler event.Handler) Item
	// NewItemWithKey creates a new item with a key accelerator using the platform-default modifiers.
	NewItemWithKey func(title string, keyCode int, handler event.Handler) Item
	// NewItemWithKeyAndModifiers creates a new item.
	NewItemWithKeyAndModifiers func(title string, keyCode int, modifiers keys.Modifiers, handler event.Handler) Item
	// NewSeparator creates a new separator item.
	NewSeparator func() Item
)
