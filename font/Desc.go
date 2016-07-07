// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package font

import (
	"bytes"
	"fmt"
)

const (
	// BoldMask when set, indicates the font will be rendered as bold.
	BoldMask Flags = 1 << iota
	// ActualBoldMask when set, indicates the font is actually bold.
	ActualBoldMask
	// ItalicMask when set, indicates the font will be rendered as italic.
	ItalicMask
	// ActualItalicMask when set, indicates the font is actually italic.
	ActualItalicMask
	// CondensedMask when set, indicates the font will be rendered with less spacing than normal
	// between its characters. Cannot be set at the same time as ExpandedMask.
	CondensedMask
	// ExpandedMask when set, indicates the font will be rendered with more spacing than normal
	// between its characters. Cannot be set at the same time as CondensedMask.
	ExpandedMask
	// MonospacedMask when set, indicates the font has a uniform width for each character.
	MonospacedMask
	// UserSettableMask contains the masks that can be set by a client.
	// All other bits will be ignored when the font is created.
	UserSettableMask = BoldMask | ItalicMask | CondensedMask | ExpandedMask
)

// Flags holds the style flags for a Font.
type Flags uint8

// Desc holds information necessary to reconstruct a Font.
type Desc struct {
	Family string  // Family name of the Font, such as "Times New Roman".
	Size   float32 // Size of the font, in points.
	Flags  Flags   // Style flags.
}

var (
	// User is the font used by default for documents and other text under the user’s
	// control (that is, text whose font the user can normally change).
	User Desc
	// UserMonospaced is the font used by default for documents and other text under
	// the user’s control when that font is fixed-pitch.
	UserMonospaced Desc
	// System is the system font used for standard user-interface items such as window
	// titles, button labels, etc.
	System Desc
	// EmphasizedSystem is the system font used for emphasis in alerts.
	EmphasizedSystem Desc
	// SmallSystem is the standard small system font used for informative text in
	// alerts, column headings in lists, help tags, utility window titles, toolbar
	// item labels, tool palettes, tool tips, and small controls.
	SmallSystem Desc
	// SmallEmphasizedSystem is the small system font used for emphasis.
	SmallEmphasizedSystem Desc
	// Views is the font used as the default font of text in lists and tables.
	Views Desc
	// Label is the font used for labels.
	Label Desc
	// Menu is the font used for menus.
	Menu Desc
	// MenuCmdKey is the font used for menu item command key equivalents.
	MenuCmdKey Desc
)

// Bold when set, indicates the font will be rendered as bold.
func (desc Desc) Bold() bool {
	return (desc.Flags & BoldMask) == BoldMask
}

// ActualBold when set, indicates the font is actually bold.
func (desc Desc) ActualBold() bool {
	return (desc.Flags & ActualBoldMask) == ActualBoldMask
}

// Italic when set, indicates the font will be rendered as italic.
func (desc Desc) Italic() bool {
	return (desc.Flags & ItalicMask) == ItalicMask
}

// ActualItalic when set, indicates the font is actually italic.
func (desc Desc) ActualItalic() bool {
	return (desc.Flags & ActualItalicMask) == ActualItalicMask
}

// Condensed when set, indicates the font will be rendered with less spacing than normal between
// its characters.
func (desc Desc) Condensed() bool {
	return (desc.Flags & CondensedMask) == CondensedMask
}

// Expanded when set, indicates the font will be rendered with more spacing than normal between
// its characters.
func (desc Desc) Expanded() bool {
	return (desc.Flags & ExpandedMask) == ExpandedMask
}

// Monospaced when set, indicates the font has a uniform width for each character.
func (desc Desc) Monospaced() bool {
	return (desc.Flags & MonospacedMask) == MonospacedMask
}

// String -- implements the fmt.Stringer interface.
func (desc Desc) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(desc.Family)
	fmt.Fprintf(&buffer, " %v", desc.Size)
	if desc.Bold() {
		buffer.WriteString(" bold")
		if !desc.ActualBold() {
			buffer.WriteString(" (synthetic)")
		}
	}
	if desc.Italic() {
		buffer.WriteString(" italic")
		if !desc.ActualItalic() {
			buffer.WriteString(" (synthetic)")
		}
	}
	if desc.Condensed() {
		buffer.WriteString(" condensed")
	}
	if desc.Expanded() {
		buffer.WriteString(" expanded")
	}
	if desc.Monospaced() {
		buffer.WriteString(" monospaced")
	}
	return buffer.String()
}
