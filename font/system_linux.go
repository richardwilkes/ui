package font

func init() {
	// RAW: Try to determine what the user set in system preferences...
	User = NewFont("Sans 12")
	UserMonospaced = NewFont("Monospace 10")
	System = NewFont("Sans 13")
	EmphasizedSystem = NewFont("Sans Bold 13")
	SmallSystem = NewFont("Sans 11")
	SmallEmphasizedSystem = NewFont("Sans Bold 11")
	Views = NewFont("Sans 12")
	Label = NewFont("Sans 10")
	Menu = NewFont("Sans 14")
	MenuCmdKey = NewFont("Sans 14")
}
