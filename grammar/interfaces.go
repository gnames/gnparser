package grammar

type Valuer interface {
	// value function returns the complete composite value of a node.
	// for low level nodes it would be the same as Value field, for higher
	// nodes it will be a value made from all their components.
	value() string
}

type Canonizer interface {
	// canonical function would return something only for nodes that do
	// contribute to canonical representation. For other nodes the return
	// value is an empty canonical structure.
	canonical() *Canonical
}

type Poser interface {
	// pos function returns a meaning of words in a string and their positions
	pos() []Pos
}

type AuthorFinder interface {
	lastAuthorship() *authorshipNode
}

type Outputter interface {
	// details creates a details structure for JSON-based outputs
	details() interface{}
}

type Name interface {
	Valuer
	Canonizer
	Poser
	AuthorFinder
	Outputter
}
