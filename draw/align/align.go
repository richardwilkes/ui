package align

// Alignment constants.
const (
	Start Alignment = iota
	Middle
	End
	Fill
)

// Alignment specifies how to align an object within its available space.
type Alignment uint8
