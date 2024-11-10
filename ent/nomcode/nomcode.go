package nomcode

import (
	"strings"
)

// Code represents a nomenclatural code.
type Code int

// Constants for different nomenclatural codes.
const (
	Unknown    Code = iota // Unknown code
	Zoological             // Zoological code
	Botanical              // Botanical code
	Cultivar               // Cultivar code
	Bacterial              // Bacterial code
)

// New creates a new Code from a string representation.
// It accepts short codes ('b', 'z', 'c') and full names
// ('botanical', 'zoological', 'cultivar') as well as
// official abbreviations ('icn', 'iczn', 'icncp').
// The input string is case-insensitive.
func New(s string) Code {
	s = strings.ToLower(s)
	switch s {
	case "bot", "botanical", "icn":
		return Botanical
	case "zoo", "zoological", "iczn":
		return Zoological
	case "cult", "cultivar", "icncp":
		return Cultivar
	case "bact", "bacterial", "icnp":
		return Bacterial
	default:
		return Unknown
	}
}

// String returns the official abbreviation of the nomenclatural code.
func (c Code) String() string {
	switch c {
	case Zoological:
		return "ICZN"
	case Botanical:
		return "ICN"
	case Cultivar:
		return "ICNCP"
	case Bacterial:
		return "ICNP"
	default:
		return ""
	}
}
