package web

import (
	"github.com/gnames/gnparser"
)

type GNParserService interface {
	gnparser.GNParser
	Ping() string
	Port() int
}
