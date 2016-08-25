// Copyright (c) 2016 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package main

import (
	"fmt"
	"github.com/richardwilkes/geom"
	"github.com/richardwilkes/ui"
	"github.com/richardwilkes/ui/Demo/images"
	"github.com/richardwilkes/ui/border"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/font"
	"github.com/richardwilkes/ui/menu"
	"unicode"
)

var (
	aboutWindow *ui.Window
)

func main() {
	// event.TraceAllEvents = true
	// event.TraceEventTypes = append(event.TraceEventTypes, event.MouseDownType, event.MouseDraggedType, event.MouseUpType)
	ui.App.EventHandlers().Add(event.AppWillFinishStartupType, func(evt event.Event) {
		createMenuBar()
		w1 := createButtonsWindow()
		w2 := createButtonsWindow()
		frame1 := w1.Frame()
		frame2 := w2.Frame()
		frame2.X = frame1.X + frame1.Width + 5
		frame2.Y = frame1.Y
		w2.SetFrame(frame2)
	})
	ui.StartUserInterface()
}

func createMenuBar() {
	_, aboutItem, prefsItem := ui.AddAppMenu()
	aboutItem.EventHandlers().Add(event.SelectionType, createAboutWindow)
	prefsItem.EventHandlers().Add(event.SelectionType, createPreferencesWindow)

	fileMenu := menu.Bar().AddMenu("File")
	fileMenu.AddItem("Open", "o")
	fileMenu.AddSeparator()
	item := fileMenu.AddItem("Close", "w")
	handlers := item.EventHandlers()
	handlers.Add(event.SelectionType, func(evt event.Event) {
		window := ui.KeyWindow()
		if window != nil && window.Closable() {
			window.AttemptClose()
		}
	})
	handlers.Add(event.ValidateType, func(evt event.Event) {
		window := ui.KeyWindow()
		if window == nil || !window.Closable() {
			evt.(*event.Validate).MarkInvalid()
		}
	})

	m := menu.Bar().AddMenu("Edit")
	ui.AddCutItem(m)
	ui.AddCopyItem(m)
	ui.AddPasteItem(m)
	m.AddSeparator()
	ui.AddDeleteItem(m)
	ui.AddSelectAllItem(m)

	ui.AddWindowMenu()
	ui.AddHelpMenu()
}

func createButtonsWindow() *ui.Window {
	wnd := ui.NewWindow(geom.Point{}, ui.StdWindowMask)
	wnd.SetTitle("Demo")

	root := wnd.RootWidget()
	root.SetBorder(border.NewEmpty(geom.Insets{Top: 10, Left: 10, Bottom: 10, Right: 10}))
	ui.NewPrecision(root).SetVerticalSpacing(10)

	buttonsPanel := createButtonsPanel()
	buttonsPanel.SetLayoutData(ui.NewPrecisionData().SetHorizontalGrab(true))
	root.AddChild(buttonsPanel)

	addSeparator(root)

	checkBoxPanel := createCheckBoxPanel()
	checkBoxPanel.SetLayoutData(ui.NewPrecisionData().SetHorizontalGrab(true))
	root.AddChild(checkBoxPanel)

	addSeparator(root)

	radioButtonsPanel := createRadioButtonsPanel()
	radioButtonsPanel.SetLayoutData(ui.NewPrecisionData().SetHorizontalGrab(true))
	root.AddChild(radioButtonsPanel)

	addSeparator(root)

	popupMenusPanel := createPopupMenusPanel()
	popupMenusPanel.SetLayoutData(ui.NewPrecisionData().SetHorizontalGrab(true))
	root.AddChild(popupMenusPanel)

	addSeparator(root)

	wrapper := ui.NewBlock()
	ui.NewPrecision(wrapper).SetColumns(2).SetEqualColumns(true).SetHorizontalSpacing(10)
	wrapper.SetLayoutData(ui.NewPrecisionData().SetHorizontalGrab(true).SetHorizontalAlignment(draw.AlignFill))
	textFieldsPanel := createTextFieldsPanel()
	textFieldsPanel.SetLayoutData(ui.NewPrecisionData().SetHorizontalGrab(true).SetHorizontalAlignment(draw.AlignFill))
	wrapper.AddChild(textFieldsPanel)
	wrapper.AddChild(createListPanel())
	root.AddChild(wrapper)

	addSeparator(root)

	img, err := draw.AcquireImageFromURL("http://allwallpapersnew.com/wp-content/gallery/stock-photos-for-free/grassy_field_sunset___free_stock_by_kevron2001-d5blgkr.jpg")
	if err == nil {
		content := ui.NewImageLabel(img)
		content.SetFocusable(true)
		_, prefSize, _ := ui.Sizes(content, ui.NoHintSize)
		content.SetSize(prefSize)
		scrollArea := ui.NewScrollArea(content, ui.ScrollContentUnmodified)
		scrollArea.SetLayoutData(ui.NewPrecisionData().SetHorizontalAlignment(draw.AlignFill).SetVerticalAlignment(draw.AlignFill).SetHorizontalGrab(true).SetVerticalGrab(true))
		root.AddChild(scrollArea)
	} else {
		fmt.Println(err)
	}

	wnd.Pack()
	wnd.ToFront()
	return wnd
}

func createListPanel() ui.Widget {
	list := ui.NewList(&ui.TextCellFactory{})
	list.Append("One",
		"Two",
		"Three with some long text to make it interesting",
		"Four",
		"Five")
	list.EventHandlers().Add(event.SelectionType, func(evt event.Event) {
		fmt.Print("Selection changed in list. Now:")
		index := -1
		first := true
		for {
			index = list.Selection.NextSet(index + 1)
			if index == -1 {
				break
			}
			if first {
				first = false
			} else {
				fmt.Print(",")
			}
			fmt.Printf(" %d", index)
		}
		fmt.Println()
	})
	list.EventHandlers().Add(event.ClickType, func(evt event.Event) {
		fmt.Println("Double-clicked on list")
	})
	_, prefSize, _ := ui.Sizes(list, ui.NoHintSize)
	list.SetSize(prefSize)
	scrollArea := ui.NewScrollArea(list, ui.ScrollContentFill)
	scrollArea.SetLayoutData(ui.NewPrecisionData().SetHorizontalAlignment(draw.AlignFill).SetVerticalAlignment(draw.AlignFill).SetHorizontalGrab(true).SetVerticalGrab(true))
	return scrollArea
}

func addSeparator(root ui.Widget) {
	sep := ui.NewSeparator(true)
	sep.SetLayoutData(ui.NewPrecisionData().SetHorizontalAlignment(draw.AlignFill))
	root.AddChild(sep)
}

func createButtonsPanel() ui.Widget {
	panel := ui.NewBlock()
	ui.NewFlow(panel).SetHorizontalSpacing(5).SetVerticalSpacing(5).SetVerticallyCentered(true)

	createButton("Press Me", panel)
	createButton("Disabled", panel).SetEnabled(false)

	img, err := draw.AcquireImageFromFile(images.FS, "/home.png")
	if err == nil {
		createImageButton(img, "Home", panel)
		createImageButton(img, "Home (disabled)", panel).SetEnabled(false)
	} else {
		fmt.Println(err)
	}

	img, err = draw.AcquireImageFromFile(images.FS, "/classic-apple-logo.png")
	if err == nil {
		createImageButton(img, "Classic Apple Logo", panel)
		createImageButton(img, "Classic Apple Logo (disabled)", panel).SetEnabled(false)
	} else {
		fmt.Println(err)
	}

	return panel
}

func createButton(title string, panel ui.Widget) *ui.Button {
	button := ui.NewButton(title)
	button.EventHandlers().Add(event.ClickType, func(evt event.Event) { fmt.Printf("The button '%s' was clicked.\n", title) })
	ui.NewSimpleToolTip(button, fmt.Sprintf("This is the tooltip for the '%s' button.", title))
	panel.AddChild(button)
	return button
}

func createImageButton(img *draw.Image, name string, panel ui.Widget) *ui.ImageButton {
	size := img.Size()
	size.Width /= 2
	size.Height /= 2
	button := ui.NewImageButtonWithImageSize(img, size)
	button.EventHandlers().Add(event.ClickType, func(evt event.Event) { fmt.Printf("The button '%s' was clicked.\n", name) })
	ui.NewSimpleToolTip(button, name)
	panel.AddChild(button)
	return button
}

func createCheckBoxPanel() ui.Widget {
	panel := ui.NewBlock()
	ui.NewPrecision(panel)
	createCheckBox("Press Me", panel)
	createCheckBox("Initially Mixed", panel).SetState(ui.Mixed)
	createCheckBox("Disabled", panel).SetEnabled(false)
	checkbox := createCheckBox("Disabled w/Check", panel)
	checkbox.SetEnabled(false)
	checkbox.SetState(ui.Checked)
	return panel
}

func createCheckBox(title string, panel ui.Widget) *ui.CheckBox {
	checkbox := ui.NewCheckBox(title)
	checkbox.EventHandlers().Add(event.ClickType, func(evt event.Event) { fmt.Printf("The checkbox '%s' was clicked.\n", title) })
	ui.NewSimpleToolTip(checkbox, fmt.Sprintf("This is the tooltip for the '%s' checkbox.", title))
	panel.AddChild(checkbox)
	return checkbox
}

func createRadioButtonsPanel() ui.Widget {
	panel := ui.NewBlock()
	ui.NewPrecision(panel)

	group := ui.NewRadioButtonGroup()
	first := createRadioButton("First", panel, group)
	createRadioButton("Second", panel, group)
	createRadioButton("Third (disabled)", panel, group).SetEnabled(false)
	createRadioButton("Fourth", panel, group)
	group.Select(first)

	return panel
}

func createRadioButton(title string, panel ui.Widget, group *ui.RadioButtonGroup) *ui.RadioButton {
	rb := ui.NewRadioButton(title)
	rb.EventHandlers().Add(event.ClickType, func(evt event.Event) { fmt.Printf("The radio button '%s' was clicked.\n", title) })
	ui.NewSimpleToolTip(rb, fmt.Sprintf("This is the tooltip for the '%s' radio button.", title))
	panel.AddChild(rb)
	group.Add(rb)
	return rb
}

func createPopupMenusPanel() ui.Widget {
	panel := ui.NewBlock()
	ui.NewPrecision(panel)

	createPopupMenu(panel, 1, "One", "Two", "Three", "", "Four", "Five", "Six")
	createPopupMenu(panel, 2, "Red", "Blue", "Green").SetEnabled(false)

	return panel
}

func createPopupMenu(panel ui.Widget, selection int, titles ...string) *ui.PopupMenu {
	p := ui.NewPopupMenu()
	ui.NewSimpleToolTip(p, fmt.Sprintf("This is the tooltip for the PopupMenu with %d items.", len(titles)))
	for _, title := range titles {
		if title == "" {
			p.AddSeparator()
		} else {
			p.AddItem(title)
		}
	}
	p.SelectIndex(selection)
	p.EventHandlers().Add(event.SelectionType, func(evt event.Event) { fmt.Printf("The '%v' item was selected from the PopupMenu.\n", p.Selected()) })
	panel.AddChild(p)
	return p
}

func createTextFieldsPanel() ui.Widget {
	panel := ui.NewBlock()
	ui.NewPrecision(panel)

	createTextField("First Text Field", panel)
	createTextField("Second Text Field (disabled)", panel).SetEnabled(false)
	createTextField("", panel).SetWatermark("Watermarked")
	field := createTextField("", panel)
	field.SetWatermark("Enter only numbers")
	field.EventHandlers().Add(event.ValidateType, func(evt event.Event) {
		e := evt.(*event.Validate)
		for _, r := range field.Text() {
			if !unicode.IsDigit(r) {
				e.MarkInvalid()
				break
			}
		}
	})

	return panel
}

func createTextField(text string, panel ui.Widget) *ui.TextField {
	field := ui.NewTextField()
	field.SetText(text)
	field.SetLayoutData(ui.NewPrecisionData().SetHorizontalGrab(true).SetHorizontalAlignment(draw.AlignFill))
	ui.NewSimpleToolTip(field, fmt.Sprintf("This is the tooltip for the '%s' text field.", text))
	panel.AddChild(field)
	return field
}

func createAboutWindow(evt event.Event) {
	if aboutWindow == nil {
		aboutWindow = ui.NewWindow(geom.Point{}, ui.TitledWindowMask|ui.ClosableWindowMask)
		aboutWindow.EventHandlers().Add(event.ClosedType, func(evt event.Event) { aboutWindow = nil })
		aboutWindow.SetTitle("About " + ui.AppName())
		root := aboutWindow.RootWidget()
		root.SetBorder(border.NewEmpty(geom.Insets{Top: 10, Left: 10, Bottom: 10, Right: 10}))
		ui.NewPrecision(root)
		title := ui.NewLabelWithFont(ui.AppName(), font.EmphasizedSystem)
		title.SetLayoutData(ui.NewPrecisionData().SetHorizontalAlignment(draw.AlignMiddle))
		root.AddChild(title)
		desc := ui.NewLabel("Simple app to demonstrate the\ncapabilities of the ui framework.")
		root.AddChild(desc)
		aboutWindow.Pack()
	}
	aboutWindow.ToFront()
}

func createPreferencesWindow(evt event.Event) {
	fmt.Println("Preferences...")
}

type scrollTarget struct {
	hpos float64
	vpos float64
}

// LineScrollAmount implements ui.Scrollable.
func (st *scrollTarget) LineScrollAmount(horizontal, towardsStart bool) float64 {
	return 1
}

// PageScrollAmount implements ui.Scrollable.
func (st *scrollTarget) PageScrollAmount(horizontal, towardsStart bool) float64 {
	return 10
}

// ScrolledPosition implements ui.Scrollable.
func (st *scrollTarget) ScrolledPosition(horizontal bool) float64 {
	if horizontal {
		return st.hpos
	}
	return st.vpos
}

// SetScrolledPosition implements ui.Scrollable.
func (st *scrollTarget) SetScrolledPosition(horizontal bool, position float64) {
	if horizontal {
		st.hpos = position
	} else {
		st.vpos = position
	}
}

// VisibleSize implements ui.Scrollable.
func (st *scrollTarget) VisibleSize(horizontal bool) float64 {
	return 10
}

// ContentSize implements ui.Scrollable.
func (st *scrollTarget) ContentSize(horizontal bool) float64 {
	return 1000
}
