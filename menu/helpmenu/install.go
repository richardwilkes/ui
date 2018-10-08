package helpmenu

import (
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/ui/menu"
)

// Install adds a standard 'Help' menu to the end of the menu bar.
func Install(bar menu.Bar) {
	helpMenu := menu.NewMenu(i18n.Text("Help"))
	bar.AppendMenu(helpMenu)
	bar.SetupSpecialMenu(menu.HelpMenu, helpMenu)
}
