// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package factory

import (
	"github.com/richardwilkes/i18n"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/menu/custom"
	"github.com/richardwilkes/ui/menu/platform"
	"runtime"
)

var (
	// UseNative determines whether a platform's native menu implementation should be used, or the
	// custom one that is implemented by the UI framework. Currently, only macOS has a native
	// implementation. Do not mix the two.
	UseNative  = runtime.GOOS == "darwin"
	windowMenu menu.Menu
	helpMenu   menu.Menu
)

// AppBar returns the application menu bar.
func AppBar() menu.Bar {
	if UseNative {
		return platform.AppBar()
	}
	return custom.AppBar()
}

// NewMenu creates a new menu.
func NewMenu(title string) menu.Menu {
	if UseNative {
		return platform.NewMenu(title)
	}
	return custom.NewMenu(title)
}

// NewItem creates a new item with no key accelerator.
func NewItem(title string, handler event.Handler) menu.Item {
	if UseNative {
		return platform.NewItem(title, handler)
	}
	return custom.NewItem(title, handler)
}

// NewItemWithKey creates a new item with a key accelerator using the platform-default modifiers.
func NewItemWithKey(title string, keyCode int, handler event.Handler) menu.Item {
	if UseNative {
		return platform.NewItemWithKey(title, keyCode, handler)
	}
	return custom.NewItemWithKey(title, keyCode, handler)
}

// NewItemWithKeyAndModifiers creates a new item.
func NewItemWithKeyAndModifiers(title string, keyCode int, modifiers keys.Modifiers, handler event.Handler) menu.Item {
	if UseNative {
		return platform.NewItemWithKeyAndModifiers(title, keyCode, modifiers, handler)
	}
	return custom.NewItemWithKeyAndModifiers(title, keyCode, modifiers, handler)
}

// NewSeparator creates a new separator item.
func NewSeparator() menu.Item {
	if UseNative {
		return platform.NewSeparator()
	}
	return custom.NewSeparator()
}

// AddServicesMenu attempts to add the Services sub-menu to the specified menu. Currently, this is
// only applicable for macOS when using the native menus. Returns true if the menu was added.
func AddServicesMenu(menu menu.Menu) bool {
	if UseNative && runtime.GOOS == "darwin" {
		servicesMenu := NewMenu(i18n.Text("Services"))
		menu.AddMenu(servicesMenu)
		platform.SetServicesMenu(servicesMenu)
		return true
	}
	return false
}

// SetWindowMenu sets the specified menu as the menu to manipulate to show the current window list.
func SetWindowMenu(menu menu.Menu) {
	windowMenu = menu
	if UseNative && runtime.GOOS == "darwin" {
		platform.SetWindowMenu(windowMenu)
	}
}

// SetHelpMenu sets the specified menu as the menu to add help content to.
func SetHelpMenu(menu menu.Menu) {
	helpMenu = menu
	if UseNative && runtime.GOOS == "darwin" {
		platform.SetHelpMenu(helpMenu)
	}
}
