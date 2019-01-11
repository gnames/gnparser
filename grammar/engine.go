package grammar

import (
	"io"

	"gitlab.com/gogna/gnparser/dict"
)

type BaseEngine struct {
	SN        *ScientificNameNode
	root      *node32
	Surrogate bool
	Bacteria  bool
	Warnings  map[Warning]struct{}
}

func (p *Engine) resetFields() {
	var warnReset map[Warning]struct{}
	p.Warnings = warnReset
	p.Surrogate = false
	p.Bacteria = false
}

func (p *Engine) AddWarn(w Warning) {
	if p.Warnings == nil {
		p.Warnings = make(map[Warning]struct{})
	}
	if _, ok := p.Warnings[w]; !ok {
		p.Warnings[w] = struct{}{}
	}
}

func (p *Engine) IsBacteria(gen string) {
	if hom, ok := dict.Dict.Bacteria[gen]; ok {
		if hom {
			p.AddWarn(BacteriaMaybeWarn)
		} else {
			p.Bacteria = true
		}
	}
}

func (p *Engine) OutputAST() {
	type element struct {
		node *node32
		down *element
	}
	var node *node32
	var skip bool
	var stack *element
	for _, token := range p.Tokens() {
		if node, skip = p.newNode(token); skip {
			continue
		}
		for stack != nil && stackNodeIsWithin(stack.node, token) {
			stack.node.next = node.up
			node.up = stack.node
			stack = stack.down
		}
		stack = &element{node: node, down: stack}
	}
	if stack != nil {
		p.root = stack.node
	}
}

func stackNodeIsWithin(n *node32, t token32) bool {
	return n.token32.begin >= t.begin && n.token32.end <= t.end
}

func (p *Engine) PrintOutputSyntaxTree(w io.Writer) {
	if p.root == nil || p.root.pegRule != ruleSciName {
		return
	}
	p.root.print(w, true, p.Buffer)
}

func (p *Engine) newNode(t token32) (*node32, bool) {
	var node *node32
	if _, ok := nodeRules[t.pegRule]; ok {
		node := &node32{token32: t}
		return node, false
	}
	switch t.pegRule {
	case ruleOtherSpace:
		p.AddWarn(SpaceNonStandardWarn)
	case ruleMultipleSpace:
		p.AddWarn(SpaceMultipleWarn)
	case ruleMiscodedChar:
		p.AddWarn(UTF8ConvBadWarn)
	case ruleBasionymAuthorship2Parens:
		p.AddWarn(AuthDoubleParensWarn)
	}
	return node, true
}

func (p *Engine) nodeValue(n *node32) string {
	t := n.token32
	v := string([]rune(p.Buffer)[t.begin:t.end])
	return v
}

func (p *Engine) ParsedName() string {
	if p.tokens32.tree == nil {
		return ""
	}
	for i := len(p.tokens32.tree) - 1; i >= 0; i-- {
		t := p.tokens32.tree[i]
		if t.pegRule == ruleName {
			return string(p.buffer[t.begin:t.end])
		}
	}
	return ""
}

var nodeRules = map[pegRule]struct{}{
	ruleSciName:                         struct{}{},
	ruleName:                            struct{}{},
	ruleTail:                            struct{}{},
	ruleHybridFormula:                   struct{}{},
	ruleNamedSpeciesHybrid:              struct{}{},
	ruleNamedGenusHybrid:                struct{}{},
	ruleSingleName:                      struct{}{},
	ruleNameApprox:                      struct{}{},
	ruleNameComp:                        struct{}{},
	ruleNameSpecies:                     struct{}{},
	ruleGenusWord:                       struct{}{},
	ruleInfraspGroup:                    struct{}{},
	ruleInfraspEpithet:                  struct{}{},
	ruleSpeciesEpithet:                  struct{}{},
	ruleComparison:                      struct{}{},
	ruleRank:                            struct{}{},
	ruleRankOtherUncommon:               struct{}{},
	ruleRankVar:                         struct{}{},
	ruleRankForma:                       struct{}{},
	ruleRankSsp:                         struct{}{},
	ruleSubGenusOrSuperspecies:          struct{}{},
	ruleSubGenus:                        struct{}{},
	ruleUninomialCombo:                  struct{}{},
	ruleRankUninomial:                   struct{}{},
	ruleUninomial:                       struct{}{},
	ruleUninomialWord:                   struct{}{},
	ruleAbbrGenus:                       struct{}{},
	ruleWord:                            struct{}{},
	ruleWordApostr:                      struct{}{},
	ruleWordStartsWithDigit:             struct{}{},
	ruleHybridChar:                      struct{}{},
	ruleApproxNameIgnored:               struct{}{},
	ruleApproximation:                   struct{}{},
	ruleAuthorship:                      struct{}{},
	ruleOriginalAuthorship:              struct{}{},
	ruleOriginalAuthorshipComb:          struct{}{},
	ruleCombinationAuthorship:           struct{}{},
	ruleBasionymAuthorshipYearMisformed: struct{}{},
	ruleBasionymAuthorship:              struct{}{},
	ruleAuthorsGroup:                    struct{}{},
	ruleAuthorsTeam:                     struct{}{},
	ruleAuthorSep:                       struct{}{},
	ruleAuthorEx:                        struct{}{},
	ruleAuthorEmend:                     struct{}{},
	ruleAuthor:                          struct{}{},
	ruleUnknownAuthor:                   struct{}{},
	ruleAuthorWord:                      struct{}{},
	ruleAuthorEtAl:                      struct{}{},
	ruleAllCapsAuthorWord:               struct{}{},
	ruleFilius:                          struct{}{},
	ruleAuthorPrefix:                    struct{}{},
	ruleYear:                            struct{}{},
	ruleYearRange:                       struct{}{},
	ruleYearWithDot:                     struct{}{},
	ruleYearApprox:                      struct{}{},
	ruleYearWithPage:                    struct{}{},
	ruleYearWithParens:                  struct{}{},
	ruleYearWithChar:                    struct{}{},
	ruleYearNum:                         struct{}{},
	ruleUpperCharExtended:               struct{}{},
	ruleLowerCharExtended:               struct{}{},
}
