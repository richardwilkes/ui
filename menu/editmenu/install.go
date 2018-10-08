package editmenu

import (
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/ui/menu"
)

// Install adds a standard 'Edit' menu to the end of the menu bar.
func Install(bar menu.Bar) menu.Menu {
	editMenu := menu.NewMenu(i18n.Text("Edit"))

	AppendCutItem(editMenu)
	AppendCopyItem(editMenu)
	AppendPasteItem(editMenu)

	editMenu.AppendItem(menu.NewSeparator())
	AppendDeleteItem(editMenu)
	AppendSelectAllItem(editMenu)

	bar.AppendMenu(editMenu)
	return editMenu
}
