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
	"bytes"
	"fmt"
	"sort"
	"unsafe"
)

const (
	// BoldStyleMask when set, indicates the font will be rendered as bold.
	BoldStyleMask FontStyleFlags = 1 << iota
	// ActualBoldStyleMask when set, indicates the font is actually bold.
	ActualBoldStyleMask
	// ItalicStyleMask when set, indicates the font will be rendered as italic.
	ItalicStyleMask
	// ActualItalicStyleMask when set, indicates the font is actually italic.
	ActualItalicStyleMask
	// CondensedStyleMask when set, indicates the font will be rendered with less spacing than
	// normal between its characters. Cannot be set at the same time as ExpandedStyleMask.
	CondensedStyleMask
	// ExpandedStyleMask when set, indicates the font will be rendered with more spacing than normal
	// between its characters. Cannot be set at the same time as CondensedStyleMask.
	ExpandedStyleMask
	// MonospacedStyleMask when set, indicates the font has a uniform width for each character.
	MonospacedStyleMask
	// UserSettableStyleMask contains the masks that can be set by a client.
	// All other bits will be ignored when the font is created.
	UserSettableStyleMask = BoldStyleMask | ItalicStyleMask | CondensedStyleMask | ExpandedStyleMask
)

const (
	userFontID = iota
	userMonospacedFontID
	systemFontID
	emphasizedSystemFontID
	smallSystemFontID
	smallEmphasizedSystemFontID
	viewsFontID
	labelFontID
	menuFontID
	menuCmdKeyFontID
)

// FontStyleFlags holds the style flags for a Font.
type FontStyleFlags uint8

// Font holds font information.
type Font struct {
	font unsafe.Pointer
	desc FontDesc
}

// FontDesc holds information necessary to reconstruct a Font.
type FontDesc struct {
	Family string         // Family name of the Font, such as "Times New Roman".
	Size   float32        // Size of the font, in points.
	Flags  FontStyleFlags // Style flags.
}

type fontRegistry struct {
	font  *Font
	count int
}

var (
	fontDescToRegMap = make(map[FontDesc]*fontRegistry, 0)
	fontToDescMap    = make(map[*Font]FontDesc, 0)
)

var (
	// UserFontDesc is the font used by default for documents and other text under the user’s
	// control (that is, text whose font the user can normally change).
	UserFontDesc = platformFontDesc(userFontID)
	// UserMonospacedFontDesc is the font used by default for documents and other text under
	// the user’s control when that font is fixed-pitch.
	UserMonospacedFontDesc = platformFontDesc(userMonospacedFontID)
	// SystemFontDesc is the system font used for standard user-interface items such as window
	// titles, button labels, etc.
	SystemFontDesc = platformFontDesc(systemFontID)
	// EmphasizedSystemFontDesc is the system font used for emphasis in alerts.
	EmphasizedSystemFontDesc = platformFontDesc(emphasizedSystemFontID)
	// SmallSystemFontDesc is the standard small system font used for informative text in
	// alerts, column headings in lists, help tags, utility window titles, toolbar
	// item labels, tool palettes, tool tips, and small controls.
	SmallSystemFontDesc = platformFontDesc(smallSystemFontID)
	// SmallEmphasizedSystemFontDesc is the small system font used for emphasis.
	SmallEmphasizedSystemFontDesc = platformFontDesc(smallEmphasizedSystemFontID)
	// ViewsFontDesc is the font used as the default font of text in lists and tables.
	ViewsFontDesc = platformFontDesc(viewsFontID)
	// LabelFontDesc is the font used for labels.
	LabelFontDesc = platformFontDesc(labelFontID)
	// MenuFontDesc is the font used for menus.
	MenuFontDesc = platformFontDesc(menuFontID)
	// MenuCmdKeyFontDesc is the font used for menu item command key equivalents.
	MenuCmdKeyFontDesc = platformFontDesc(menuCmdKeyFontID)
)

func newFont(desc FontDesc) *Font {
	return platformNewFont(desc)
}

// AcquireFont acquires a font and prepares it for use.
func AcquireFont(desc FontDesc) *Font {
	var reg *fontRegistry
	var ok bool
	if reg, ok = fontDescToRegMap[desc]; !ok {
		reg = &fontRegistry{font: newFont(desc)}
		fontDescToRegMap[desc] = reg
		fontToDescMap[reg.font] = desc
	}
	reg.count++
	return reg.font
}

// Release a font.
func (font *Font) Release() {
	if desc, ok := fontToDescMap[font]; ok {
		if reg, ok2 := fontDescToRegMap[desc]; ok2 {
			reg.count--
			if reg.count < 1 {
				delete(fontDescToRegMap, desc)
				delete(fontToDescMap, font)
				font.dispose()
			}
		} else {
			delete(fontToDescMap, font)
			font.dispose()
		}
	}
}

// Desc returns a copy of the font's description information.
func (font *Font) Desc() FontDesc {
	return font.desc
}

func (font *Font) dispose() {
	font.platformDispose()
	font.font = nil
}

// Ascent of the Font, i.e. the distance from the baseline to the top of a typical capital letter.
func (font *Font) Ascent() float32 {
	return font.platformAscent()
}

// Descent of the Font, i.e. the distance from the baseline to the bottom of the typical letter
// that has a descender, such as a lower case 'g'.
func (font *Font) Descent() float32 {
	return font.platformDescent()
}

// Leading of the Font, i.e. the recommended distance between the bottom of the descender line
// to the top of the next line.
func (font *Font) Leading() float32 {
	return font.platformLeading()
}

// String implements the fmt.Stringer interface.
func (font *Font) String() string {
	return font.desc.String()
}

// AvailableFontFamilies retrieves the names of the installed font families.
func AvailableFontFamilies() []string {
	families := platformAvailableFontFamilies()
	sort.Strings(families)
	return families
}

// Bold when set, indicates the font will be rendered as bold.
func (desc FontDesc) Bold() bool {
	return (desc.Flags & BoldStyleMask) == BoldStyleMask
}

// ActualBold when set, indicates the font is actually bold.
func (desc FontDesc) ActualBold() bool {
	return (desc.Flags & ActualBoldStyleMask) == ActualBoldStyleMask
}

// Italic when set, indicates the font will be rendered as italic.
func (desc FontDesc) Italic() bool {
	return (desc.Flags & ItalicStyleMask) == ItalicStyleMask
}

// ActualItalic when set, indicates the font is actually italic.
func (desc FontDesc) ActualItalic() bool {
	return (desc.Flags & ActualItalicStyleMask) == ActualItalicStyleMask
}

// Condensed when set, indicates the font will be rendered with less spacing than normal between
// its characters.
func (desc FontDesc) Condensed() bool {
	return (desc.Flags & CondensedStyleMask) == CondensedStyleMask
}

// Expanded when set, indicates the font will be rendered with more spacing than normal between
// its characters.
func (desc FontDesc) Expanded() bool {
	return (desc.Flags & ExpandedStyleMask) == ExpandedStyleMask
}

// Monospaced when set, indicates the font has a uniform width for each character.
func (desc FontDesc) Monospaced() bool {
	return (desc.Flags & MonospacedStyleMask) == MonospacedStyleMask
}

// String implements the fmt.Stringer interface.
func (desc FontDesc) String() string {
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
