package parser

import (
	"io"

	tb "github.com/gnames/gnlib/tribool"
	o "github.com/gnames/gnparser/entity/output"
	"github.com/gnames/gnparser/io/dict"
)

type BaseEngine struct {
	SN          *ScientificNameNode
	root        *node32
	Cardinality int
	Error       error
	Hybrid      bool
	Surrogate   bool
	Bacteria    tb.Tribool
	Warnings    map[o.Warning]struct{}
	Tail        string
}

func (p *Engine) FullReset() {
	p.Cardinality = 0
	p.Error = nil
	p.Hybrid = false
	p.Surrogate = false
	p.Bacteria = tb.NewTribool()
	var warnReset map[o.Warning]struct{}
	p.Warnings = warnReset
	p.Tail = ""
	p.Reset()
}

func (p *Engine) AddWarn(w o.Warning) {
	if p.Warnings == nil {
		p.Warnings = make(map[o.Warning]struct{})
	}
	if _, ok := p.Warnings[w]; !ok {
		p.Warnings[w] = struct{}{}
	}
}

func (p *Engine) IsBacteria(gen string) {
	if hom, ok := dict.Dict.Bacteria[gen]; ok {
		if hom {
			p.AddWarn(o.BacteriaMaybeWarn)
			p.Bacteria = tb.NewTribool(0)
		} else {
			p.Bacteria = tb.NewTribool(1)
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
	switch t.pegRule {
	case ruleHybridChar:
		p.Hybrid = true
	case ruleRankNotho, ruleRankUninomialNotho:
		p.Hybrid = true
		p.AddWarn(o.HybridNamedWarn)
	case ruleOtherSpace:
		p.AddWarn(o.SpaceNonStandardWarn)
	case ruleMultipleSpace:
		p.AddWarn(o.SpaceMultipleWarn)
	case ruleMiscodedChar:
		p.AddWarn(o.UTF8ConvBadWarn)
	case ruleBasionymAuthorship2Parens:
		p.AddWarn(o.AuthDoubleParensWarn)
	case ruleBasionymAuthorshipMissingParens:
		p.AddWarn(o.AuthMissingOneParensWarn)
	case ruleUpperAfterDash:
		p.AddWarn(o.GenusUpperCharAfterDash)
	case ruleLowerGreek:
		p.AddWarn(o.GreekLetterInRank)
	case ruleAuthorSepSpanish:
		p.AddWarn(o.SpanishAndAsSeparator)
	}
	if _, ok := nodeRules[t.pegRule]; ok {
		node := &node32{token32: t}
		return node, false
	}

	return node, true
}

func (p *Engine) nodeValue(n *node32) string {
	t := n.token32
	v := string([]rune(p.Buffer)[t.begin:t.end])
	return v
}

func (p *Engine) ParsedName() string {
	if p.Error != nil {
		return "noparse"
	}
	for i := len(p.tokens32.tree) - 1; i >= 0; i-- {
		t := p.tokens32.tree[i]
		if t.pegRule == ruleName {
			return string(p.buffer[t.begin:t.end])
		}
	}
	return "noparse"
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
	ruleBasionymAuthorshipMissingParens: struct{}{},
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
	ruleApostrOther:                     struct{}{},
	ruleAuthorSuffix:                    struct{}{},
}
