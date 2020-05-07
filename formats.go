package gnparser

import (
	"fmt"
	"log"
)

type Format int

const (
	// Compact is a JSON format without new lines and spaces.
	Compact Format = iota
	// Pretty is a JSON nested easy to read format.
	Pretty
	// Simple is a flat format with only few most 'popular' fields.
	CSV
	// Debug is a format that shows complete and truncated AST for debugging.
	Debug
)

var formats = []string{"compact", "pretty", "csv", "debug"}

func (of Format) String() string {
	return formats[of]
}

func newFormat(f string) Format {
	gnp := NewGNparser()
	for i, v := range formats {
		if v == f {
			return Format(i)
		}
	}
	err := fmt.Errorf("unknown format '%s', using default '%s' format",
		f, gnp.Format.String())
	log.Println(err)
	return gnp.Format
}

// OutputFormat returns string representation of the current output format
// for GNparser
func (gnp *GNparser) OutputFormat() string {
	return gnp.Format.String()
}

// AvailableFormats function returns a string representation of supported
// output formats.
func AvailableFormats() []string {
	return formats
}
