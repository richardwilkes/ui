package app

type driver interface {
	Start()
	Name() string
	Hide()
	HideOthers()
	ShowAll()
	AttemptQuit()
	MayQuitNow(quit bool)
}
