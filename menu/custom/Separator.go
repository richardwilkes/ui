// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package custom

import (
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/widget/separator"
)

type Separator struct {
	separator.Separator
}

func NewSeparator() *Separator {
	sep := &Separator{}
	sep.InitTypeAndID(sep)
	sep.Initialize(true)
	return sep
}

// Title returns this item's title.
func (sep *Separator) Title() string {
	return ""
}

// KeyCode returns the key code that can be used to trigger this item. A value of 0 indicates no
// key is attached.
func (sep *Separator) KeyCode() int {
	return 0
}

// KeyModifiers returns the key modifiers that are required to trigger this item.
func (sep *Separator) KeyModifiers() keys.Modifiers {
	return 0
}

// SubMenu returns a sub-menu attached to this item or nil.
func (sep *Separator) SubMenu() menu.Menu {
	return nil
}

// Dispose releases any operating system resources associated with this item.
func (sep *Separator) Dispose() {
	// Does nothing
}
