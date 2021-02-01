// Package nameidx provides a structure that preserves original position
// of a name-string in an input slice.
package nameidx

// NameIdx presents an input name-string and its position in the input
// slice.
type NameIdx struct {
	// Index is the position of a string in the input slice.
	Index int

	// NameString is the input string.
	NameString string
}
