package gnparser

import (
	"github.com/gnames/gnlib/domain/entity/gn"
	"github.com/gnames/gnlib/format"
	"github.com/gnames/gnparser/config"
	"github.com/gnames/gnparser/entity/output"
)

type GNParser interface {
	gn.Versioner
	ParseName(string) output.Parsed
	ParseNames([]string) []output.Parsed
	Format() format.Format
	ChangeConfig(opts ...config.Option) GNParser
}
