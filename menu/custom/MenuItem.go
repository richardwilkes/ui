package custom

import (
	"fmt"
	"unicode/utf8"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/ui/color"
	"github.com/richardwilkes/ui/event"
	"github.com/richardwilkes/ui/keys"
	"github.com/richardwilkes/ui/menu"
	"github.com/richardwilkes/ui/widget"
)

// MenuItem represents a single menu item.
type MenuItem struct {
	widget.Block
	Theme        *Theme
	keyCode      int
	keyModifiers keys.Modifiers
	menu         *Menu
	menuParent   menuRoot
	title        string
	pos          float64
	highlighted  bool
	menuOpen     bool
}

// NewItem creates a new item with no key accelerator.
func NewItem(title string, handler event.Handler) *MenuItem {
	return NewItemWithKey(title, 0, handler)
}

// NewItemWithKey creates a new item with a key accelerator using the platform-default modifiers.
func NewItemWithKey(title string, keyCode int, handler event.Handler) *MenuItem {
	return NewItemWithKeyAndModifiers(title, keyCode, keys.PlatformMenuModifier(), handler)
}

// NewItemWithKeyAndModifiers creates a new item.
func NewItemWithKeyAndModifiers(title string, keyCode int, modifiers keys.Modifiers, handler event.Handler) *MenuItem {
	item := &MenuItem{Theme: StdTheme, title: title, keyCode: keyCode, keyModifiers: modifiers}
	item.InitTypeAndID(item)
	item.Describer = func() string { return fmt.Sprintf("MenuItem #%d (%s)", item.ID(), item.title) }
	item.SetSizer(item)
	handlers := item.EventHandlers()
	handlers.Add(event.PaintType, item.paint)
	handlers.Add(event.MouseDownType, item.mouseDown)
	handlers.Add(event.MouseDraggedType, item.mouseDragged)
	handlers.Add(event.MouseUpType, item.mouseUp)
	handlers.Add(event.MouseEnteredType, item.mouseEntered)
	handlers.Add(event.MouseMovedType, item.mouseMoved)
	handlers.Add(event.MouseExitedType, item.mouseExited)
	handlers.Add(event.KeyDownType, item.keyDown)
	if handler != nil {
		handlers.Add(event.SelectionType, handler)
	}
	return item
}

// Title returns this item's title.
func (item *MenuItem) Title() string {
	return item.title
}

// KeyCode returns the key code that can be used to trigger this item. A value of 0 indicates no
// key is attached.
func (item *MenuItem) KeyCode() int {
	return item.keyCode
}

// KeyModifiers returns the key modifiers that are required to trigger this item.
func (item *MenuItem) KeyModifiers() keys.Modifiers {
	return item.keyModifiers
}

// Sizes implements Sizer
func (item *MenuItem) Sizes(hint geom.Size) (min, pref, max geom.Size) {
	pref = item.Theme.TitleFont.Measure(item.title)
	pref.Width += item.Theme.HMargin*2 + item.Theme.KeySpacing
	pref.Height += item.Theme.VMargin * 2
	pref.GrowToInteger()
	if item.keyCode != 0 {
		mapping := keys.MappingForKeyCode(item.keyCode)
		if mapping != nil {
			keySize := item.Theme.KeyFont.Measure(mapping.Name)
			pref.Width += keySize.Width
			if pref.Height < keySize.Height {
				pref.Height = keySize.Height
			}
			mods := item.keyModifiers
			if mods == 0 {
				mods = keys.PlatformMenuModifier()
			}
			modSize := item.Theme.KeyFont.Measure(mods.String())
			modSize.GrowToInteger()
			pref.Width += modSize.Width
		}
		pref.GrowToInteger()
	}
	pref.ConstrainForHint(hint)
	if border := item.Border(); border != nil {
		pref.AddInsets(border.Insets())
	}
	return pref, pref, pref
}

func (item *MenuItem) calculateAcceleratorPosition() float64 {
	pos := item.Theme.HMargin
	if item.keyCode != 0 {
		mapping := keys.MappingForKeyCode(item.keyCode)
		if mapping != nil {
			pos += item.Theme.KeyFont.Measure(mapping.Name).Width
		}
	}
	return pos
}

func (item *MenuItem) paint(evt event.Event) {
	if paintEvent, ok := evt.(*event.Paint); ok {
		bounds := item.LocalInsetBounds()
		gc := paintEvent.GC()
		gc.SetColor(item.currentBackground())
		gc.FillRect(bounds)
		size := item.Theme.TitleFont.Measure(item.title)
		gc.SetColor(item.textColor())
		gc.DrawString(bounds.X+item.Theme.HMargin, bounds.Y+(bounds.Height-size.Height)/2, item.title, item.Theme.TitleFont)
		if item.keyCode != 0 && item.pos > 0 {
			mapping := keys.MappingForKeyCode(item.keyCode)
			if mapping != nil {
				size = item.Theme.KeyFont.Measure(mapping.Name)
				x := bounds.X + bounds.Width - item.pos
				y := bounds.Y + (bounds.Height-size.Height)/2
				modY := y - (item.Theme.KeyFont.Leading() + 0.5)
				needOffset := false
				for _, r := range mapping.Name {
					if r < ' ' || r > utf8.RuneSelf {
						needOffset = true
					} else {
						needOffset = false
						break
					}
				}
				if needOffset {
					y = modY
				}
				gc.DrawString(x, y, mapping.Name, item.Theme.KeyFont)
				modText := item.keyModifiers.String()
				size = item.Theme.KeyFont.Measure(modText)
				gc.DrawString(x-size.Width, modY, modText, item.Theme.KeyFont)
			}
		}
	}
}

func (item *MenuItem) currentBackground() color.Color {
	switch {
	case !item.Enabled():
		return item.Theme.Background
	case item.highlighted || item.menuOpen:
		return item.Theme.HighlightedBackground
	case item.Focused():
		return item.Theme.Background.Blend(color.KeyboardFocus, 0.5)
	default:
		return item.Theme.Background
	}
}

func (item *MenuItem) textColor() color.Color {
	if !item.Enabled() {
		return item.Theme.TextWhenDisabled
	}
	if item.currentBackground().Luminance() > 0.65 {
		return item.Theme.TextWhenLight
	}
	return item.Theme.TextWhenDark
}

func (item *MenuItem) mouseDown(evt event.Event) {
	item.highlighted = true
	item.Repaint()
}

func (item *MenuItem) setMenuOpen(menuOpen bool) {
	if item.menuOpen != menuOpen {
		item.menuOpen = menuOpen
		item.highlighted = false
		item.Repaint()
	}
}

func (item *MenuItem) mouseOver(where geom.Point) {
	if item.Enabled() {
		bounds := item.LocalInsetBounds()
		highlighted := bounds.ContainsPoint(item.FromWindow(where))
		if item.highlighted != highlighted {
			item.highlighted = highlighted
			item.Repaint()
		}
		if highlighted && item.menu != nil && !item.menuOpen {
			if leafMenu := item.findRoot().findLeafOpenMenu(); leafMenu != nil {
				leafMenu.close(nil)
				item.menu.open(nil)
			}
		}
	}
}

func (item *MenuItem) mouseDragged(evt event.Event) {
	if evt2, ok := evt.(*event.MouseDragged); ok {
		item.mouseOver(evt2.Where())
	}
}

func (item *MenuItem) mouseUp(evt event.Event) {
	item.highlighted = false
	item.Repaint()
	if mouseUp, ok := evt.(*event.MouseUp); ok {
		bounds := item.LocalInsetBounds()
		if bounds.ContainsPoint(item.FromWindow(mouseUp.Where())) {
			event.Dispatch(event.NewSelection(item))
			event.Dispatch(event.NewClosing(item))
		}
	}
}

func (item *MenuItem) mouseEntered(evt event.Event) {
	if evt2, ok := evt.(*event.MouseEntered); ok {
		item.mouseOver(evt2.Where())
	}
}

func (item *MenuItem) mouseMoved(evt event.Event) {
	if evt2, ok := evt.(*event.MouseMoved); ok {
		item.mouseOver(evt2.Where())
	}
}

func (item *MenuItem) mouseExited(evt event.Event) {
	if evt2, ok := evt.(*event.MouseExited); ok {
		item.mouseOver(evt2.Where())
	}
}

func (item *MenuItem) keyDown(evt event.Event) {
	if evt2, ok := evt.(*event.KeyDown); ok {
		if keys.IsControlAction(evt2.Code()) {
			evt.Finish()
			event.Dispatch(event.NewSelection(item))
		}
	}
}

func (item *MenuItem) validate() {
	evt := event.NewValidate(item)
	event.Dispatch(evt)
	item.SetEnabled(evt.Valid())
}

// SubMenu returns a sub-menu attached to this item or nil.
func (item *MenuItem) SubMenu() menu.Menu {
	return item.menu
}

// Dispose releases any operating system resources associated with this item.
func (item *MenuItem) Dispose() {
	// Does nothing
}

func (item *MenuItem) processKeyDown(evt *event.KeyDown) bool {
	if item.KeyCode() == evt.Code() && item.KeyModifiers() == evt.Modifiers() {
		item.validate()
		if item.Enabled() {
			event.Dispatch(event.NewSelection(item))
			evt.Finish()
		}
		return true
	} else if item.menu != nil {
		return item.menu.processKeyDown(evt)
	}
	return false
}

func (item *MenuItem) findRoot() menuRoot {
	if item.menuParent != nil {
		return item.menuParent
	}
	return item.menu
}
