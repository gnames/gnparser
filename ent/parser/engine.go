package parser

import (
	"io"

	"github.com/gnames/gnlib/ent/nomcode"
	"github.com/gnames/gnparser/ent/icvcn"
	"github.com/gnames/gnparser/ent/internal/preparser"
	"github.com/gnames/gnparser/ent/parsed"
	"github.com/gnames/gnparser/io/dict"
	"github.com/gnames/tribool"
)

type baseEngine struct {
	preParser          *preparser.PreParser
	icvcnParser        *icvcn.Parser
	sn                 *scientificNameNode
	root               *node32
	nodePool           []node32
	runeToByteOffset   []int
	nodePoolIdx        int
	code               nomcode.Code
	cardinality        int
	rank               string
	error              error
	hybrid             *parsed.Annotation
	graftChimera       *parsed.Annotation
	surrogate          *parsed.Annotation
	bacteria           *tribool.Tribool
	candidatus         bool
	warnings           map[parsed.Warning]struct{}
	tail               string
	cultivar           bool
	preserveDiaereses  bool
	compactAuthors     bool
}

// New creates implementation of Parser interface.
func New() Parser {
	p := Engine{}
	p.Init()
	p.preParser = preparser.New()
	ip := &icvcn.Parser{}
	ip.Init()
	p.icvcnParser = ip
	p.nodePool = make([]node32, 128)
	return &p
}

// fullReset must set all fields to empty, or results from the previous
// parse might bleed into new results.
func (p *Engine) fullReset() {
	p.cardinality = 0
	p.rank = ""
	p.error = nil
	p.hybrid = nil
	p.graftChimera = nil
	p.surrogate = nil
	p.bacteria = nil
	p.candidatus = false
	var warnReset map[parsed.Warning]struct{}
	p.warnings = warnReset
	p.tail = ""
	p.cultivar = false
	p.nodePoolIdx = 0
	p.Reset()
	p.buildRuneOffsets()
}

func (p *Engine) addWarn(w parsed.Warning) {
	if p.warnings == nil {
		p.warnings = make(map[parsed.Warning]struct{})
	}
	if _, ok := p.warnings[w]; !ok {
		p.warnings[w] = struct{}{}
	}
}

func (p *Engine) isBacteria(gen string) {
	if p.code == nomcode.Bacterial {
		bac := tribool.New(1)
		p.bacteria = &bac
	}
	if hom, ok := dict.Dict.Bacteria[gen]; ok {
		if hom {
			p.addWarn(parsed.BacteriaMaybeWarn)
			bac := tribool.New(0)
			p.bacteria = &bac
		} else {
			bac := tribool.New(1)
			p.bacteria = &bac
		}
	}
}

// outputAST assembles PEG nodes' AST structure.
func (p *Engine) outputAST() {
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
	return n.begin >= t.begin && n.end <= t.end
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
	var annot parsed.Annotation
	switch t.pegRule {
	case ruleHybridChar:
		annot = parsed.HybridAnnot
		p.hybrid = &annot
	case ruleGraftChimeraChar:
		annot = parsed.GraftChimeraAnnot
		p.hybrid = &annot
	case ruleRankNotho, ruleRankUninomialNotho:
		annot = parsed.NothoHybridAnnot
		p.hybrid = &annot
		p.addWarn(parsed.HybridNamedWarn)
	case ruleOtherSpace:
		p.addWarn(parsed.SpaceNonStandardWarn)
	case ruleMiscodedChar:
		p.addWarn(parsed.UTF8ConvBadWarn)
	case ruleAbbrSubgenus:
		p.addWarn(parsed.SubgenusAbbrWarn)
	case ruleBasionymAuthorship2Parens:
		p.addWarn(parsed.AuthDoubleParensWarn)
	case ruleBasionymAuthorshipMissingParens:
		p.addWarn(parsed.AuthMissingOneParensWarn)
	case ruleUpperAfterDash:
		p.addWarn(parsed.GenusUpperCharAfterDash)
	case ruleLowerGreek:
		p.addWarn(parsed.GreekLetterInRank)
	case ruleAuthorSepSpanish:
		p.addWarn(parsed.SpanishAndAsSeparator)
	case ruleIgnoredWord:
		p.addWarn(parsed.ContainsIgnoredAnnotation)
	}
	if nodeRuleSet[t.pegRule] {
		// Area allocation for nodes reduces GC pressure.
		// if nodePool is has space, allocate node there, if it is full
		// allocate node on the heap. Hopefully a scientific name generates
		// less than 128 nodes, so most of the time the GC will not need to
		// worry about releasing separate nodes.
		if p.nodePoolIdx < len(p.nodePool) {
			n := &p.nodePool[p.nodePoolIdx]
			p.nodePoolIdx++
			*n = node32{token32: t}
			return n, false
		}
		node = &node32{token32: t}
		return node, false
	}

	return node, true
}

// buildRuneOffsets builds a rune-index → byte-offset table from p.buffer,
// which is already populated by Reset(). Stored in runeToByteOffset so it
// can be reused across the pool without reallocating each parse.
func (p *Engine) buildRuneOffsets() {
	runeLen := len(p.buffer) // p.buffer = []rune(p.Buffer) + endSymbol
	if cap(p.runeToByteOffset) < runeLen {
		p.runeToByteOffset = make([]int, runeLen)
	} else {
		p.runeToByteOffset = p.runeToByteOffset[:runeLen]
	}
	ri := 0
	for bytePos := range p.Buffer {
		p.runeToByteOffset[ri] = bytePos
		ri++
	}
	// Sentinel: endSymbol position maps past the end of p.Buffer.
	p.runeToByteOffset[runeLen-1] = len(p.Buffer)
}

func (p *Engine) nodeValue(n *node32) string {
	t := n.token32
	return p.Buffer[p.runeToByteOffset[t.begin]:p.runeToByteOffset[t.end]]
}

// ParseName returns the name the nodes. In case of parsing errors
// returns string 'noparse'.
func (p *Engine) ParsedName() string {
	if p.error != nil {
		return "noparse"
	}
	for i := len(p.tree) - 1; i >= 0; i-- {
		t := p.tree[i]
		if t.pegRule == ruleName {
			return string(p.buffer[t.begin:t.end])
		}
	}
	return "noparse"
}

// nodeRuleSet is a fixed-size array for O(1) rule lookup
// (pegRule is uint8, max 255).
var nodeRuleSet [256]bool

func init() {
	for rule := range nodeRules {
		nodeRuleSet[rule] = true
	}
}

var nodeRules = map[pegRule]struct{}{
	ruleAbbrGenus:                       {},
	ruleAbbrSubgenus:                    {},
	ruleAllCapsAuthorWord:               {},
	ruleApostrOther:                     {},
	ruleApproxNameIgnored:               {},
	ruleApproximation:                   {},
	ruleAuthor:                          {},
	ruleAuthorEmend:                     {},
	ruleAuthorEtAl:                      {},
	ruleAuthorEx:                        {},
	ruleAuthorIn:                        {},
	ruleAuthorPrefix:                    {},
	ruleAuthorSep:                       {},
	ruleAuthorSuffix:                    {},
	ruleAuthorWord:                      {},
	ruleAuthorsGroup:                    {},
	ruleAuthorsTeam:                     {},
	ruleAuthorship:                      {},
	ruleBasionymAuthorship:              {},
	ruleBasionymAuthorshipMissingParens: {},
	ruleBasionymAuthorshipYearMisformed: {},
	ruleCandidatus:                      {},
	ruleCandidatusName:                  {},
	ruleCombinationAuthorship:           {},
	ruleComparison:                      {},
	ruleCultivar:                        {},
	ruleCultivarRecursive:               {},
	ruleDashOther:                       {},
	ruleDotPrefix:                       {},
	ruleFilius:                          {},
	ruleFiliusFNoSpace:                  {},
	ruleGenusWord:                       {},
	ruleGraftChimeraChar:                {},
	ruleGraftChimeraFormula:             {},
	ruleHybridChar:                      {},
	ruleHybridFormula:                   {},
	ruleInfraspEpithet:                  {},
	ruleInfraspGroup:                    {},
	ruleLowerCharExtended:               {},
	ruleName:                            {},
	ruleNameApprox:                      {},
	ruleNameComp:                        {},
	ruleNameCompSp:                      {},
	ruleNameCompIsp:                     {},
	ruleNameSpecies:                     {},
	ruleNamedGenusGraftChimera:          {},
	ruleNamedGenusHybrid:                {},
	ruleNamedSpeciesHybrid:              {},
	ruleOriginalAuthorship:              {},
	ruleOriginalAuthorshipComb:          {},
	ruleRank:                            {},
	ruleRankCultivar:                    {},
	ruleRankForma:                       {},
	ruleRankOtherUncommon:               {},
	ruleRankSsp:                         {},
	ruleRankUninomial:                   {},
	ruleRankVar:                         {},
	ruleSciName:                         {},
	ruleSingleName:                      {},
	ruleSpeciesEpithet:                  {},
	ruleSubgenus:                        {},
	ruleSubgenusOrSuperspecies:          {},
	ruleTail:                            {},
	ruleUninomial:                       {},
	ruleUninomialCombo:                  {},
	ruleUninomialWord:                   {},
	ruleUnknownAuthor:                   {},
	ruleUpperCharExtended:               {},
	ruleWord:                            {},
	ruleWordApostr:                      {},
	ruleWordStartsWithDigit:             {},
	ruleYear:                            {},
	ruleYearApprox:                      {},
	ruleYearNum:                         {},
	ruleYearRange:                       {},
	ruleYearWithChar:                    {},
	ruleYearWithDot:                     {},
	ruleYearWithPage:                    {},
	ruleYearWithParens:                  {},
}
