package input

// Name presents an input name-string and its position in the input
// slice.
type Name struct {
	// Index is the position of a string in the input slice.
	Index int

	// NameString is the input string.
	NameString string
}
