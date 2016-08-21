// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package draw

import (
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/geom"
)

func (gc *graphics) platformSave() {
	// RAW: Implement platformSave for Windows
}

func (gc *graphics) platformRestore() {
	// RAW: Implement platformRestore for Windows
}

func (gc *graphics) platformSetOpacity(opacity float32) {
	// RAW: Implement platformSetOpacity for Windows
}

func (gc *graphics) platformSetFillColor(color color.Color) {
	// RAW: Implement platformSetFillColor for Windows
}

func (gc *graphics) platformSetStrokeColor(color color.Color) {
	// RAW: Implement platformSetStrokeColor for Windows
}

func (gc *graphics) platformSetStrokeWidth(width float32) {
	// RAW: Implement platformSetStrokeWidth for Windows
}

func (gc *graphics) platformFillRect(bounds geom.Rect) {
	// RAW: Implement platformFillRect for Windows
}

func (gc *graphics) platformStrokeRect(bounds geom.Rect) {
	// RAW: Implement platformFillRect for Windows
}

func (gc *graphics) platformFillEllipse(bounds geom.Rect) {
	// RAW: Implement platformFillEllipse for Windows
}

func (gc *graphics) platformStrokeEllipse(bounds geom.Rect) {
	// RAW: Implement platformStrokeEllipse for Windows
}

func (gc *graphics) platformFillPath() {
	// RAW: Implement platformFillPath for Windows
}

func (gc *graphics) platformFillPathEvenOdd() {
	// RAW: Implement platformFillPathEvenOdd for Windows
}

func (gc *graphics) platformStrokePath() {
	// RAW: Implement platformStrokePath for Windows
}

func (gc *graphics) platformBeginPath() {
	// RAW: Implement platformBeginPath for Windows
}

func (gc *graphics) platformClosePath() {
	// RAW: Implement platformClosePath for Windows
}

func (gc *graphics) platformMoveTo(x, y float32) {
	// RAW: Implement platformMoveTo for Windows
}

func (gc *graphics) platformLineTo(x, y float32) {
	// RAW: Implement platformLineTo for Windows
}

func (gc *graphics) platformArc(cx, cy, radius, startAngleRadians, endAngleRadians float32, clockwise bool) {
	// RAW: Implement platformArc for Windows
}

func (gc *graphics) platformArcTo(x1, y1, x2, y2, radius float32) {
	// RAW: Implement platformArcTo for Windows
}

func (gc *graphics) platformCurveTo(cp1x, cp1y, cp2x, cp2y, x, y float32) {
	// RAW: Implement platformCurveTo for Windows
}

func (gc *graphics) platformQuadCurveTo(cpx, cpy, x, y float32) {
	// RAW: Implement platformQuadCurveTo for Windows
}

func (gc *graphics) platformAddPath(path *geom.Path) {
	// RAW: Implement platformAddPath for Windows
}

func (gc *graphics) platformClip() {
	// RAW: Implement platformClip for Windows
}

func (gc *graphics) platformClipEvenOdd() {
	// RAW: Implement platformClipEvenOdd for Windows
}

func (gc *graphics) platformClipRect(bounds geom.Rect) {
	// RAW: Implement platformClipRect for Windows
}

func (gc *graphics) platformDrawLinearGradient(gradient *Gradient, sx, sy, ex, ey float32) {
	// RAW: Implement platformDrawLinearGradient for Windows
}

func (gc *graphics) platformDrawRadialGradient(gradient *Gradient, scx, scy, startRadius, ecx, ecy, endRadius float32) {
	// RAW: Implement platformDrawRadialGradient for Windows
}

func (gc *graphics) platformDrawImageInRect(img *Image, bounds geom.Rect) {
	// RAW: Implement platformDrawImageInRect for Windows
}

func (gc *graphics) platformDrawString(x, y float32, str string) {
	// RAW: Implement platformDrawString for Windows
}

func (gc *graphics) platformTranslate(x, y float32) {
	// RAW: Implement platformTranslate for Windows
}

func (gc *graphics) platformScale(x, y float32) {
	// RAW: Implement platformScale for Windows
}

func (gc *graphics) platformRotate(angleInRadians float32) {
	// RAW: Implement platformRotate for Windows
}
