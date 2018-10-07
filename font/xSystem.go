package font

var (
	// User is the font used by default for documents and other text under the user’s control
	// (that is, text whose font the user can normally change).
	User *Font
	// UserMonospaced is the font used by default for documents and other text under the user’s
	// control when that font is fixed-pitch.
	UserMonospaced *Font
	// System is the system font used for standard user-interface items such as window titles,
	// button labels, etc.
	System *Font
	// EmphasizedSystem is the system font used for emphasis in alerts.
	EmphasizedSystem *Font
	// SmallSystem is the standard small system font used for informative text in alerts,
	// column headings in lists, help tags, utility window titles, toolbar item labels, tool
	// palettes, tool tips, and small controls.
	SmallSystem *Font
	// SmallEmphasizedSystem is the small system font used for emphasis.
	SmallEmphasizedSystem *Font
	// Views is the font used as the default font of text in lists and tables.
	Views *Font
	// Label is the font used for labels.
	Label *Font
	// Menu is the font used for menus.
	Menu *Font
	// MenuCmdKey is the font used for menu item command key equivalents.
	MenuCmdKey *Font
)
