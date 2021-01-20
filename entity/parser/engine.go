package parser

import (
	"io"

	tb "github.com/gnames/gnlib/tribool"
	o "github.com/gnames/gnparser/entity/output"
	"github.com/gnames/gnparser/io/dict"
)

type baseEngine struct {
	sn          *scientificNameNode
	root        *node32
	cardinality int
	error       error
	hybrid      *o.Annotation
	surrogate   *o.Annotation
	bacteria    *tb.Tribool
	warnings    map[o.Warning]struct{}
	tail        string
}

// NewParser creates implementation of Parser interface.
func NewParser() Parser {
	p := Engine{}
	p.Init()
	return &p
}

func (p *Engine) fullReset() {
	p.cardinality = 0
	p.error = nil
	p.hybrid = nil
	p.surrogate = nil
	p.bacteria = nil
	var warnReset map[o.Warning]struct{}
	p.warnings = warnReset
	p.tail = ""
	p.Reset()
}

func (p *Engine) addWarn(w o.Warning) {
	if p.warnings == nil {
		p.warnings = make(map[o.Warning]struct{})
	}
	if _, ok := p.warnings[w]; !ok {
		p.warnings[w] = struct{}{}
	}
}

func (p *Engine) isBacteria(gen string) {
	if hom, ok := dict.Dict.Bacteria[gen]; ok {
		if hom {
			p.addWarn(o.BacteriaMaybeWarn)
			bac := tb.NewTribool(0)
			p.bacteria = &bac
		} else {
			bac := tb.NewTribool(1)
			p.bacteria = &bac
		}
	}
}

// OutputAST assembles PEG nodes AST structure.
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

// PrintOutputSyntaxTree outputs a simplified version of a nodes
// Abstract Syntax Tree. This method can be used for debugging purposes.
func (p *Engine) PrintOutputSyntaxTree(w io.Writer) {
	if p.root == nil || p.root.pegRule != ruleSciName {
		return
	}
	p.root.print(w, true, p.Buffer)
}

func (p *Engine) newNode(t token32) (*node32, bool) {
	var node *node32
	var annot o.Annotation
	switch t.pegRule {
	case ruleHybridChar:
		annot = o.HybridAnnot
		p.hybrid = &annot
	case ruleRankNotho, ruleRankUninomialNotho:
		annot = o.NothoHybridAnnot
		p.hybrid = &annot
		p.addWarn(o.HybridNamedWarn)
	case ruleOtherSpace:
		p.addWarn(o.SpaceNonStandardWarn)
	case ruleMultipleSpace:
		p.addWarn(o.SpaceMultipleWarn)
	case ruleMiscodedChar:
		p.addWarn(o.UTF8ConvBadWarn)
	case ruleBasionymAuthorship2Parens:
		p.addWarn(o.AuthDoubleParensWarn)
	case ruleBasionymAuthorshipMissingParens:
		p.addWarn(o.AuthMissingOneParensWarn)
	case ruleUpperAfterDash:
		p.addWarn(o.GenusUpperCharAfterDash)
	case ruleLowerGreek:
		p.addWarn(o.GreekLetterInRank)
	case ruleAuthorSepSpanish:
		p.addWarn(o.SpanishAndAsSeparator)
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

// ParseName returns the name the nodes. In case of parsing errors
// returns string 'noparse'.
func (p *Engine) ParsedName() string {
	if p.error != nil {
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
	ruleSciName:                         {},
	ruleName:                            {},
	ruleTail:                            {},
	ruleHybridFormula:                   {},
	ruleNamedSpeciesHybrid:              {},
	ruleNamedGenusHybrid:                {},
	ruleSingleName:                      {},
	ruleNameApprox:                      {},
	ruleNameComp:                        {},
	ruleNameSpecies:                     {},
	ruleGenusWord:                       {},
	ruleInfraspGroup:                    {},
	ruleInfraspEpithet:                  {},
	ruleSpeciesEpithet:                  {},
	ruleComparison:                      {},
	ruleRank:                            {},
	ruleRankOtherUncommon:               {},
	ruleRankVar:                         {},
	ruleRankForma:                       {},
	ruleRankSsp:                         {},
	ruleSubgenusOrSuperspecies:          {},
	ruleSubgenus:                        {},
	ruleUninomialCombo:                  {},
	ruleRankUninomial:                   {},
	ruleUninomial:                       {},
	ruleUninomialWord:                   {},
	ruleAbbrGenus:                       {},
	ruleWord:                            {},
	ruleWordApostr:                      {},
	ruleWordStartsWithDigit:             {},
	ruleHybridChar:                      {},
	ruleApproxNameIgnored:               {},
	ruleApproximation:                   {},
	ruleAuthorship:                      {},
	ruleOriginalAuthorship:              {},
	ruleOriginalAuthorshipComb:          {},
	ruleCombinationAuthorship:           {},
	ruleBasionymAuthorshipYearMisformed: {},
	ruleBasionymAuthorshipMissingParens: {},
	ruleBasionymAuthorship:              {},
	ruleAuthorsGroup:                    {},
	ruleAuthorsTeam:                     {},
	ruleAuthorSep:                       {},
	ruleAuthorEx:                        {},
	ruleAuthorEmend:                     {},
	ruleAuthor:                          {},
	ruleUnknownAuthor:                   {},
	ruleAuthorWord:                      {},
	ruleAuthorEtAl:                      {},
	ruleAllCapsAuthorWord:               {},
	ruleFilius:                          {},
	ruleAuthorPrefix:                    {},
	ruleYear:                            {},
	ruleYearRange:                       {},
	ruleYearWithDot:                     {},
	ruleYearApprox:                      {},
	ruleYearWithPage:                    {},
	ruleYearWithParens:                  {},
	ruleYearWithChar:                    {},
	ruleYearNum:                         {},
	ruleUpperCharExtended:               {},
	ruleLowerCharExtended:               {},
	ruleApostrOther:                     {},
	ruleAuthorSuffix:                    {},
}
