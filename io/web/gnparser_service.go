package web

import (
	"github.com/gnames/gnparser"
)

type gnparserService struct {
	gnparser.GNParser
	port int
}

// NewGNParserService creates a new object that implements GNParserService
// interface.
func NewGNParserService(gnp gnparser.GNParser, port int) GNParserService {
	res := gnparserService{
		GNParser: gnp,
		port:     port,
	}
	return &res
}

// Ping is a method to check a liveliness of the service, returns "pong".
func (gnps *gnparserService) Ping() string {
	return "pong"
}

// Port returns the port of the service.
func (gnps *gnparserService) Port() int {
	return gnps.port
}
