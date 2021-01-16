package gnparser

import (
	"context"

	"github.com/gnames/gnlib/domain/entity/gn"
	"github.com/gnames/gnlib/format"
	"github.com/gnames/gnparser/config"
	"github.com/gnames/gnparser/entity/input"
	"github.com/gnames/gnparser/entity/output"
)

type GNParser interface {
	gn.Versioner
	ParseName(string) output.Parsed
	ParseNames([]string) []output.Parsed
	ParseNameStream(context.Context, <-chan input.Name, chan<- output.Parsed)
	Format() format.Format
	ChangeConfig(opts ...config.Option) GNParser
}
