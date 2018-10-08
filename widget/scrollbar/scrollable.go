package scrollbar

// Scrollable objects can respond to ScrollBars.
type Scrollable interface {
	Pager
	// ScrolledPosition is called to determine the current position of the Scrollable.
	ScrolledPosition(horizontal bool) float64
	// SetScrolledPosition is called to set the current position of the Scrollable.
	SetScrolledPosition(horizontal bool, position float64)
	// VisibleSize is called to determine the size of the visible portion of the Scrollable.
	VisibleSize(horizontal bool) float64
	// ContentSize is called to determine the total size of the Scrollable.
	ContentSize(horizontal bool) float64
}
