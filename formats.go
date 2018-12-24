package gnparser

import (
	"fmt"
	"log"
)

type format int

const (
	Compact format = iota
	Pretty
	Simple
)

var formats = []string{"compact", "pretty", "simple"}

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
	err := fmt.Errorf("Unknown format '%s', using default '%s' format.",
		f, gnp.format.String())
	log.Println(err)
	return gnp.format
}

// OutputForat returns string representation of the current output format
// for GNparser
func (gnp *GNparser) OutputFormat() string {
	return gnp.format.String()
}

// AvailableFormats function returns a string representation of supported
// output formats.
func AvailableFormats() []string {
	return formats
}

// StrToFormat function creates an internal type of a supported format
// out of string.
func StrToFormat(s string) format {
	return newFormat(s)
}
