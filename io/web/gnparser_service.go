package web

import (
	"github.com/gnames/gnparser"
)

type gnparserService struct {
	gnparser.GNParser
	port int
}

func NewGNParserService(gnp gnparser.GNParser, port int) GNParserService {
	res := gnparserService{
		GNParser: gnp,
		port:     port,
	}
	return &res
}

func (gnps *gnparserService) Ping() string {
	return "pong"
}

func (gnps *gnparserService) Port() int {
	return gnps.port
}
