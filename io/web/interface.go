package web

import (
	"github.com/gnames/gnparser"
)

// GNParserService is an interface that provides functionality for
// GNParser RESTful service.
type GNParserService interface {
	gnparser.GNParser
	// Ping is a method to check if the service is running. Returns "pong".
	Ping() string
	// Port returns the port of the service.
	Port() int
}
