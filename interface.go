package gnparser

import (
	"github.com/gnames/gnlib/domain/entity/gn"
	"github.com/gnames/gnparser/entity/input"
	"github.com/gnames/gnparser/entity/output"
)

type GNParser interface {
	gn.Versioner
	ParseName(string) output.Parsed
	ParseNames([]input.Name) []output.ParseResult
}
