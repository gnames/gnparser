package gnparser

import (
	"fmt"
	"log"
)

type format int

const (
	// Compact is a JSON format without new lines and spaces.
	Compact format = iota
	// Pretty is a JSON nested easy to read format.
	Pretty
	// Simple is a flat format with only few most 'popular' fields.
	Simple
	// Debug is a format that shows complete and truncated AST for debugging.
	Debug
)

var formats = []string{"compact", "pretty", "simple", "debug"}

func (of format) String() string {
	return formats[of]
}

func newFormat(f string) format {
	gnp := NewGNparser()
	for i, v := range formats {
		if v == f {
			return format(i)
		}
	}
	err := fmt.Errorf("unknown format '%s', using default '%s' format",
		f, gnp.format.String())
	log.Println(err)
	return gnp.format
}

// OutputFormat returns string representation of the current output format
// for GNparser
func (gnp *GNparser) OutputFormat() string {
	return gnp.format.String()
}

// AvailableFormats function returns a string representation of supported
// output formats.
func AvailableFormats() []string {
	return formats
}
