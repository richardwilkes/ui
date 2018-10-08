package scrollbar

// Pager objects can provide line and page information for scrolling.
type Pager interface {
	// LineScrollAmount is called to determine how far to scroll in the given direction to reveal
	// a full 'line' of content. A positive value should be returned regardless of the direction,
	// although negative values will behave as if they were positive.
	LineScrollAmount(horizontal, towardsStart bool) float64
	// PageScrollAmount is called to determine how far to scroll in the given direction to reveal
	// a full 'page' of content. A positive value should be returned regardless of the direction,
	// although negative values will behave as if they were positive.
	PageScrollAmount(horizontal, towardsStart bool) float64
}
