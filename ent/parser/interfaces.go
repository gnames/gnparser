package parser

import (
	"github.com/gnames/gnparser/ent/parsed"
)

// Parser is an interface that is responsible for parsing of a scientific
// name and creation of the Abstract Syntax Tree of the name-string.
type Parser interface {
	// PreprocessAndParse takes a scientific name and returns back Abstract
	// Syntax Tree of the name-string.
	PreprocessAndParse(
		name, version string,
		keepHTML, capitalize, disableCultivars bool,
	) ScientificNameNode
	Debug(name string) []byte
}

// ScientificNameNode is the Abstract Syntax Tree of a name-string.
// It contains a method to convert AST into final output.
type ScientificNameNode interface {
	// ToOutput converts AST into final output object.
	ToOutput(withDetails bool) parsed.Parsed
}

// nameData is the interface for converting AST to output elements.
type nameData interface {
	valuer
	canonizer
	worder
	authorFinder
	outputter
}

type valuer interface {
	// value function returns the complete composite value of a node.
	// for low level nodes it would be the same as Value field, for higher
	// nodes it will be a value made from all their components.
	value() string
}

type canonizer interface {
	// canonical function would return something only for nodes that do
	// contribute to canonical representation. For other nodes the return
	// value is an empty canonical structure.
	canonical() *canonical
}

type worder interface {
	// words function returns a meaning of words in a string and their positions
	words() []parsed.Word
}

type authorFinder interface {
	lastAuthorship() *authorshipNode
}

type outputter interface {
	// details creates a details structure for JSON-based outputs
	details() parsed.Details
}
