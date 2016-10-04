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
	"github.com/richardwilkes/ui/app"
	"github.com/richardwilkes/ui/border"
	"github.com/richardwilkes/ui/draw"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/font"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/layout"
	"github.com/richardwilkes/ui/layout/flex"
	"github.com/richardwilkes/ui/layout/flow"
	"github.com/richardwilkes/ui/menu/appmenu"
	"github.com/richardwilkes/ui/menu/editmenu"
	"github.com/richardwilkes/ui/menu/factory"
	"github.com/richardwilkes/ui/menu/helpmenu"
	"github.com/richardwilkes/ui/menu/windowmenu"
	"github.com/richardwilkes/ui/widget"
	"github.com/richardwilkes/ui/widget/button"
	"github.com/richardwilkes/ui/widget/checkbox"
	"github.com/richardwilkes/ui/widget/imagebutton"
	"github.com/richardwilkes/ui/widget/imagelabel"
	"github.com/richardwilkes/ui/widget/label"
	"github.com/richardwilkes/ui/widget/list"
	"github.com/richardwilkes/ui/widget/popupmenu"
	"github.com/richardwilkes/ui/widget/radiobutton"
	"github.com/richardwilkes/ui/widget/scrollarea"
	"github.com/richardwilkes/ui/widget/separator"
	"github.com/richardwilkes/ui/widget/textfield"
	"github.com/richardwilkes/ui/widget/window"
	"unicode"
)

var (
	aboutWindow ui.Window
)

func main() {
	// event.TraceAllEvents = true
	// event.TraceEventTypes = append(event.TraceEventTypes, event.MouseDownType, event.MouseDraggedType, event.MouseUpType)
	app.App.EventHandlers().Add(event.AppWillFinishStartupType, func(evt event.Event) {
		createMenuBar()
		w1 := createButtonsWindow("Demo #1")
		w2 := createButtonsWindow("Demo #2")
		frame1 := w1.Frame()
		frame2 := w2.Frame()
		frame2.X = frame1.X + frame1.Width + 5
		frame2.Y = frame1.Y
		w2.SetFrame(frame2)
	})
	app.StartUserInterface()
}

func createMenuBar() {
	_, aboutItem, prefsItem := appmenu.Install()
	aboutItem.EventHandlers().Add(event.SelectionType, createAboutWindow)
	prefsItem.EventHandlers().Add(event.SelectionType, createPreferencesWindow)
	createFileMenu()
	createEditMenu()
	windowmenu.Install(-1)
	helpmenu.Install(-1)
}

func createFileMenu() {
	fileMenu := factory.NewMenu("File")

	fileMenu.InsertItem(factory.NewItemWithKey("Open", keys.VK_O, nil), -1)
	fileMenu.InsertItem(factory.NewSeparator(), -1)

	item := factory.NewItemWithKey("Close", keys.VK_W, func(evt event.Event) {
		wnd := window.KeyWindow()
		if wnd != nil && wnd.Closable() {
			wnd.AttemptClose()
		}
	})
	item.EventHandlers().Add(event.ValidateType, func(evt event.Event) {
		wnd := window.KeyWindow()
		if wnd == nil || !wnd.Closable() {
			evt.(*event.Validate).MarkInvalid()
		}
	})
	fileMenu.InsertItem(item, -1)

	factory.AppBar().InsertMenu(fileMenu, -1)
}

func createEditMenu() {
	editMenu := factory.NewMenu("Edit")

	editmenu.InsertCutItem(editMenu, -1)
	editmenu.InsertCopyItem(editMenu, -1)
	editmenu.InsertPasteItem(editMenu, -1)
	editMenu.InsertItem(factory.NewSeparator(), -1)
	editmenu.InsertDeleteItem(editMenu, -1)
	editmenu.InsertSelectAllItem(editMenu, -1)

	factory.AppBar().InsertMenu(editMenu, -1)
}

func createButtonsWindow(title string) ui.Window {
	wnd := window.NewWindow(geom.Point{}, window.StdWindowMask)
	wnd.SetTitle(title)

	content := wnd.Content()
	content.SetBorder(border.NewEmpty(geom.Insets{Top: 10, Left: 10, Bottom: 10, Right: 10}))
	flex.NewLayout(content).SetVerticalSpacing(10)

	buttonsPanel := createButtonsPanel()
	buttonsPanel.SetLayoutData(flex.NewData().SetHorizontalGrab(true))
	content.AddChild(buttonsPanel)

	addSeparator(content)

	checkBoxPanel := createCheckBoxPanel()
	checkBoxPanel.SetLayoutData(flex.NewData().SetHorizontalGrab(true))
	content.AddChild(checkBoxPanel)

	addSeparator(content)

	radioButtonsPanel := createRadioButtonsPanel()
	radioButtonsPanel.SetLayoutData(flex.NewData().SetHorizontalGrab(true))
	content.AddChild(radioButtonsPanel)

	addSeparator(content)

	popupMenusPanel := createPopupMenusPanel()
	popupMenusPanel.SetLayoutData(flex.NewData().SetHorizontalGrab(true))
	content.AddChild(popupMenusPanel)

	addSeparator(content)

	wrapper := widget.NewBlock()
	flex.NewLayout(wrapper).SetColumns(2).SetEqualColumns(true).SetHorizontalSpacing(10)
	wrapper.SetLayoutData(flex.NewData().SetHorizontalGrab(true).SetHorizontalAlignment(draw.AlignFill))
	textFieldsPanel := createTextFieldsPanel()
	textFieldsPanel.SetLayoutData(flex.NewData().SetHorizontalGrab(true).SetHorizontalAlignment(draw.AlignFill))
	wrapper.AddChild(textFieldsPanel)
	wrapper.AddChild(createListPanel())
	content.AddChild(wrapper)

	addSeparator(content)

	img, err := draw.AcquireImageFromURL("http://legends.trollworks.com/mountains.jpg")
	if err == nil {
		imgPanel := imagelabel.New(img)
		imgPanel.SetFocusable(true)
		_, prefSize, _ := ui.Sizes(imgPanel, layout.NoHintSize)
		imgPanel.SetSize(prefSize)
		scrollArea := scrollarea.New(imgPanel, scrollarea.Unmodified)
		scrollArea.SetLayoutData(flex.NewData().SetHorizontalAlignment(draw.AlignFill).SetVerticalAlignment(draw.AlignFill).SetHorizontalGrab(true).SetVerticalGrab(true))
		content.AddChild(scrollArea)
	} else {
		fmt.Println(err)
	}

	wnd.SetFocus(textFieldsPanel.Children()[0])
	wnd.Pack()
	wnd.ToFront()
	return wnd
}

func createListPanel() ui.Widget {
	list := list.New(&label.LabelCellFactory{})
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
	_, prefSize, _ := ui.Sizes(list, layout.NoHintSize)
	list.SetSize(prefSize)
	scrollArea := scrollarea.New(list, scrollarea.Fill)
	scrollArea.SetLayoutData(flex.NewData().SetHorizontalAlignment(draw.AlignFill).SetVerticalAlignment(draw.AlignFill).SetHorizontalGrab(true).SetVerticalGrab(true))
	return scrollArea
}

func addSeparator(parent ui.Widget) {
	sep := separator.New(true)
	sep.SetLayoutData(flex.NewData().SetHorizontalAlignment(draw.AlignFill))
	parent.AddChild(sep)
}

func createButtonsPanel() ui.Widget {
	panel := widget.NewBlock()
	flow.New(panel).SetHorizontalSpacing(5).SetVerticalSpacing(5).SetVerticallyCentered(true)

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

func createButton(title string, panel ui.Widget) *button.Button {
	button := button.New(title)
	button.EventHandlers().Add(event.ClickType, func(evt event.Event) { fmt.Printf("The button '%s' was clicked.\n", title) })
	widget.NewSimpleToolTip(button, fmt.Sprintf("This is the tooltip for the '%s' button.", title))
	panel.AddChild(button)
	return button
}

func createImageButton(img *draw.Image, name string, panel ui.Widget) *imagebutton.ImageButton {
	size := img.Size()
	size.Width /= 2
	size.Height /= 2
	button := imagebutton.NewImageButtonWithImageSize(img, size)
	button.EventHandlers().Add(event.ClickType, func(evt event.Event) { fmt.Printf("The button '%s' was clicked.\n", name) })
	widget.NewSimpleToolTip(button, name)
	panel.AddChild(button)
	return button
}

func createCheckBoxPanel() ui.Widget {
	panel := widget.NewBlock()
	flex.NewLayout(panel)
	createCheckBox("Press Me", panel)
	createCheckBox("Initially Mixed", panel).SetState(checkbox.Mixed)
	createCheckBox("Disabled", panel).SetEnabled(false)
	check := createCheckBox("Disabled w/Check", panel)
	check.SetEnabled(false)
	check.SetState(checkbox.Checked)
	return panel
}

func createCheckBox(title string, panel ui.Widget) *checkbox.CheckBox {
	check := checkbox.NewCheckBox(title)
	check.EventHandlers().Add(event.ClickType, func(evt event.Event) { fmt.Printf("The checkbox '%s' was clicked.\n", title) })
	widget.NewSimpleToolTip(check, fmt.Sprintf("This is the tooltip for the '%s' checkbox.", title))
	panel.AddChild(check)
	return check
}

func createRadioButtonsPanel() ui.Widget {
	panel := widget.NewBlock()
	flex.NewLayout(panel)

	group := radiobutton.NewGroup()
	first := createRadioButton("First", panel, group)
	createRadioButton("Second", panel, group)
	createRadioButton("Third (disabled)", panel, group).SetEnabled(false)
	createRadioButton("Fourth", panel, group)
	group.Select(first)

	return panel
}

func createRadioButton(title string, panel ui.Widget, group *radiobutton.Group) *radiobutton.RadioButton {
	rb := radiobutton.New(title)
	rb.EventHandlers().Add(event.ClickType, func(evt event.Event) { fmt.Printf("The radio button '%s' was clicked.\n", title) })
	widget.NewSimpleToolTip(rb, fmt.Sprintf("This is the tooltip for the '%s' radio button.", title))
	panel.AddChild(rb)
	group.Add(rb)
	return rb
}

func createPopupMenusPanel() ui.Widget {
	panel := widget.NewBlock()
	flex.NewLayout(panel)

	createPopupMenu(panel, 1, "One", "Two", "Three", "", "Four", "Five", "Six")
	createPopupMenu(panel, 2, "Red", "Blue", "Green").SetEnabled(false)

	return panel
}

func createPopupMenu(panel ui.Widget, selection int, titles ...string) *popupmenu.PopupMenu {
	p := popupmenu.NewPopupMenu()
	widget.NewSimpleToolTip(p, fmt.Sprintf("This is the tooltip for the PopupMenu with %d items.", len(titles)))
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
	panel := widget.NewBlock()
	flex.NewLayout(panel)

	field := createTextField("First Text Field", panel)
	createTextField("Second Text Field (disabled)", panel).SetEnabled(false)
	createTextField("", panel).SetWatermark("Watermarked")
	field = createTextField("", panel)
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

func createTextField(text string, panel ui.Widget) *textfield.TextField {
	field := textfield.New()
	field.SetText(text)
	field.SetLayoutData(flex.NewData().SetHorizontalGrab(true).SetHorizontalAlignment(draw.AlignFill))
	widget.NewSimpleToolTip(field, fmt.Sprintf("This is the tooltip for the '%s' text field.", text))
	panel.AddChild(field)
	return field
}

func createAboutWindow(evt event.Event) {
	if aboutWindow == nil {
		aboutWindow = window.NewWindow(geom.Point{}, window.TitledWindowMask|window.ClosableWindowMask)
		aboutWindow.EventHandlers().Add(event.ClosedType, func(evt event.Event) { aboutWindow = nil })
		aboutWindow.SetTitle("About " + app.AppName())
		content := aboutWindow.Content()
		content.SetBorder(border.NewEmpty(geom.Insets{Top: 10, Left: 10, Bottom: 10, Right: 10}))
		flex.NewLayout(content)
		title := label.NewWithFont(app.AppName(), font.EmphasizedSystem)
		title.SetLayoutData(flex.NewData().SetHorizontalAlignment(draw.AlignMiddle))
		content.AddChild(title)
		desc := label.New("Simple app to demonstrate the\ncapabilities of the ui framework.")
		content.AddChild(desc)
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
