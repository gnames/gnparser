package parser

import (
	o "github.com/gnames/gnparser/entity/output"
)

type Parser interface {
	Init()
	PreprocessAndParse() *SciNameNode
}

type SciNameNode interface {
	ToOutput() o.Parsed
}

type nameData interface {
	valuer
	canonizer
	poser
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

type poser interface {
	// pos function returns a meaning of words in a string and their positions
	pos() []o.Position
}

type authorFinder interface {
	lastAuthorship() *authorshipNode
}

type outputter interface {
	// details creates a details structure for JSON-based outputs
	details() o.Details
}
