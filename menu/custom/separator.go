package custom

import (
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/widget/separator"
)

// Separator represents a menu item separator.
type Separator struct {
	separator.Separator
}

// NewSeparator creates a new menu item separator.
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
