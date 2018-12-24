package grammar

type constructor func([]node) node

type node interface {
	value() string
}

var nodes map[pegRule]constructor = map[pegRule]constructor{
	ruleYear: newNodeYear,
}

type nodeYear struct {
	val string
}

func (y *nodeYear) value() string {
	return y.val
}

func newNodeYear(ns []node) node {
	var y *nodeYear
	return y
}
