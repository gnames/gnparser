package gnparser

var (
	// Version is the version of the gnparser package. When Makefile is
	// used, the version is calculated out of Git tags.
	Version = "v1.6.4+"
	// Build is a timestamp of when Makefile was used to compile
	// the gnparser code. If go build was used, Build stays empty.
	Build string
)
