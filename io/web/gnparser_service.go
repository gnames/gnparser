package web

import (
	"github.com/gnames/gnparser"
)

type gnparserService struct {
	gnparser.GNparser
	port int
}

// NewGNparserService creates a new object that implements GNparserService
// interface.
func NewGNparserService(gnp gnparser.GNparser, port int) GNparserService {
	res := gnparserService{
		GNparser: gnp,
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
