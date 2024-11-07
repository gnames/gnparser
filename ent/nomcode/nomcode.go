package nomcode

import (
	"log/slog"
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
)

// New creates a new Code from a string representation.
// It accepts short codes ('b', 'z', 'c') and full names
// ('botanical', 'zoological', 'cultivar') as well as
// official abbreviations ('icn', 'iczn', 'icncp').
// The input string is case-insensitive.
func New(s string) Code {
	sOrig := s
	s = strings.ToLower(s)
	switch s {
	case "b", "bot", "botanical", "icn":
		return Botanical
	case "z", "zoo", "zoological", "iczn":
		return Zoological
	case "c", "cult", "cultivar", "icncp":
		return Cultivar
	default:
		slog.Warn("Cannot determine nomenclatural code", "input", sOrig)
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
	default:
		return ""
	}
}
