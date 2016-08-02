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
	"github.com/richardwilkes/ui/geom"
	"sort"
	"unsafe"
)

const (
	// Bold when set, indicates the font will be rendered as bold.
	Bold Style = 1 << iota
	// ActualBold when set, indicates the font is actually bold.
	ActualBold
	// Italic when set, indicates the font will be rendered as italic.
	Italic
	// ActualItalic when set, indicates the font is actually italic.
	ActualItalic
	// Condensed when set, indicates the font will be rendered with less spacing than normal
	// between its characters. Cannot be set at the same time as Expanded.
	Condensed
	// Expanded when set, indicates the font will be rendered with more spacing than normal
	// between its characters. Cannot be set at the same time as Condensed.
	Expanded
	// Monospaced when set, indicates the font has a uniform width for each character.
	Monospaced
	// UserSettable contains the masks that can be set by a client.
	// All other bits will be ignored when the font is created.
	UserSettable = Bold | Italic | Condensed | Expanded
)

const (
	userID = iota
	userMonospacedID
	systemID
	emphasizedSystemID
	smallSystemID
	smallEmphasizedSystemID
	viewsID
	labelID
	menuID
	menuCmdKeyID
)

// Style holds the style flags for a Font.
type Style uint8

// Font holds font information.
type Font struct {
	font unsafe.Pointer
	desc Desc
}

// Desc holds information necessary to reconstruct a Font.
type Desc struct {
	Family string  // Family name of the Font, such as "Times New Roman".
	Size   float32 // Size of the font, in points.
	Style  Style   // Style flags.
}

type fontRegistry struct {
	font  *Font
	count int
}

var (
	descToRegMap  = make(map[Desc]*fontRegistry, 0)
	fontToDescMap = make(map[*Font]Desc, 0)
)

var (
	// UserDesc is the font used by default for documents and other text under the user’s control
	// (that is, text whose font the user can normally change).
	UserDesc = platformDesc(userID)
	// UserMonospacedDesc is the font used by default for documents and other text under the user’s
	// control when that font is fixed-pitch.
	UserMonospacedDesc = platformDesc(userMonospacedID)
	// SystemDesc is the system font used for standard user-interface items such as window titles,
	// button labels, etc.
	SystemDesc = platformDesc(systemID)
	// EmphasizedSystemDesc is the system font used for emphasis in alerts.
	EmphasizedSystemDesc = platformDesc(emphasizedSystemID)
	// SmallSystemDesc is the standard small system font used for informative text in alerts,
	// column headings in lists, help tags, utility window titles, toolbar item labels, tool
	// palettes, tool tips, and small controls.
	SmallSystemDesc = platformDesc(smallSystemID)
	// SmallEmphasizedSystemDesc is the small system font used for emphasis.
	SmallEmphasizedSystemDesc = platformDesc(smallEmphasizedSystemID)
	// ViewsDesc is the font used as the default font of text in lists and tables.
	ViewsDesc = platformDesc(viewsID)
	// LabelDesc is the font used for labels.
	LabelDesc = platformDesc(labelID)
	// MenuDesc is the font used for menus.
	MenuDesc = platformDesc(menuID)
	// MenuCmdKeyDesc is the font used for menu item command key equivalents.
	MenuCmdKeyDesc = platformDesc(menuCmdKeyID)
)

// Acquire a font and prepares it for use.
func Acquire(desc Desc) *Font {
	var reg *fontRegistry
	var ok bool
	if reg, ok = descToRegMap[desc]; !ok {
		reg = &fontRegistry{font: platformNewFont(desc)}
		descToRegMap[desc] = reg
		fontToDescMap[reg.font] = desc
	}
	reg.count++
	return reg.font
}

// PlatformPtr returns a pointer to the underlying platform-specific data.
func (font *Font) PlatformPtr() unsafe.Pointer {
	return font.font
}

// Release a font.
func (font *Font) Release() {
	if desc, ok := fontToDescMap[font]; ok {
		if reg, ok2 := descToRegMap[desc]; ok2 {
			reg.count--
			if reg.count < 1 {
				delete(descToRegMap, desc)
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
func (font *Font) Desc() Desc {
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

// Height of the font, i.e. Ascent() + Descent().
func (font *Font) Height() float32 {
	return font.Ascent() + font.Descent()
}

// Width of the string rendered with this font.
func (font *Font) Width(str string) float32 {
	return font.platformWidth(str)
}

// Size of the string rendered with this font.
func (font *Font) Size(str string) geom.Size {
	return geom.Size{Width: font.Width(str), Height: font.Height()}
}

func (font *Font) IndexForPosition(x float32, str string) int {
	return font.platformIndexForPosition(x, str)
}

func (font *Font) PositionForIndex(index int, str string) float32 {
	return font.platformPositionForIndex(index, str)
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
func (desc Desc) Bold() bool {
	return (desc.Style & Bold) == Bold
}

// ActualBold when set, indicates the font is actually bold.
func (desc Desc) ActualBold() bool {
	return (desc.Style & ActualBold) == ActualBold
}

// Italic when set, indicates the font will be rendered as italic.
func (desc Desc) Italic() bool {
	return (desc.Style & Italic) == Italic
}

// ActualItalic when set, indicates the font is actually italic.
func (desc Desc) ActualItalic() bool {
	return (desc.Style & ActualItalic) == ActualItalic
}

// Condensed when set, indicates the font will be rendered with less spacing than normal between
// its characters.
func (desc Desc) Condensed() bool {
	return (desc.Style & Condensed) == Condensed
}

// Expanded when set, indicates the font will be rendered with more spacing than normal between
// its characters.
func (desc Desc) Expanded() bool {
	return (desc.Style & Expanded) == Expanded
}

// Monospaced when set, indicates the font has a uniform width for each character.
func (desc Desc) Monospaced() bool {
	return (desc.Style & Monospaced) == Monospaced
}

// String implements the fmt.Stringer interface.
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
