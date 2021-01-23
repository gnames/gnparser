package web

import (
	"github.com/gnames/gnparser"
)

// GNparserService is an interface that provides functionality for
// GNparser RESTful service.
type GNparserService interface {
	gnparser.GNparser
	// Ping is a method to check if the service is running. Returns "pong".
	Ping() string
	// Port returns the port of the service.
	Port() int
}
