package app

// Possible termination responses
const (
	Cancel QuitResponse = iota
	Now
	Later // Must make a call to MayQuitNow() at some point in the future.
)

// QuitResponse is used to respond to requests for app termination.
type QuitResponse int
