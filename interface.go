package gnparser

import (
	"context"

	"github.com/gnames/gnlib/domain/entity/gn"
	"github.com/gnames/gnlib/format"
	"github.com/gnames/gnparser/ent/nameidx"
	"github.com/gnames/gnparser/ent/parsed"
)

// GNparser is the main use-case interface. It provides methods required
// for parsing scientific names.
type GNparser interface {
	// Versioner provides a version and a build timestamp of gnparser.
	gn.Versioner
	// Parse name takes a name-string, and returns parsed results for the name.
	ParseName(string) parsed.Parsed
	// Parse names takes a slice of name-strings, and returns a slice of
	// parsed results in the same order as the input.
	ParseNames([]string) []parsed.Parsed
	// ParseNameString takes a context, an input channel that takes a
	// a name-string and its position in the input. It returns parsed results
	// that come in the same order as the input.
	ParseNameStream(context.Context, <-chan nameidx.NameIdx, chan<- parsed.Parsed)
	// Format returns currently chosen desired output format of a JSON or
	// CSV output.
	Format() format.Format
	// ChangeConfig allows to modify settings of GNparser. Changing settings
	// might modify parsing process, and the final output of results.
	ChangeConfig(opts ...Option) GNparser
}
