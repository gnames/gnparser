package grammar

//go:generate peg grammar.peg

import (
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
)

const endSymbol rune = 1114112

/* The rule types inferred from the grammar are below. */
type pegRule uint8

const (
	ruleUnknown pegRule = iota
	ruleSciName
	ruleTail
	ruleName
	ruleHybridFormula
	ruleHybridFormulaFull
	ruleHybridFormulaPart
	ruleNamedHybrid
	ruleNamedSpeciesHybrid
	ruleNamedGenusHybrid
	ruleSingleName
	ruleNameUninomial
	ruleNameApprox
	ruleNameComp
	ruleNameSpecies
	ruleGenusWord
	ruleInfraspGroup
	ruleInfraspEpithet
	ruleSpeciesEpithet
	ruleComparison
	ruleRank
	ruleRankOtherUncommon
	ruleRankOther
	ruleRankVar
	ruleRankForma
	ruleRankSsp
	ruleSubGenusOrSuperspecies
	ruleSubGenus
	ruleUninomialCombo
	ruleUninomialCombo1
	ruleUninomialCombo2
	ruleRankUninomial
	ruleUninomial
	ruleUninomialWord
	ruleAbbrGenus
	ruleCapWord
	ruleCapWord1
	ruleCapWord2
	ruleTwoLetterGenus
	ruleWord
	ruleWord1
	ruleWordStartsWithDigit
	ruleWord2
	ruleWordApostr
	ruleWord4
	ruleHybridChar
	ruleApproxNameIgnored
	ruleApproximation
	ruleAuthorship
	ruleAuthorshipCombo
	ruleOriginalAuthorship
	ruleOriginalAuthorshipComb
	ruleCombinationAuthorship
	ruleBasionymAuthorshipYearMisformed
	ruleBasionymAuthorship
	ruleBasionymAuthorship1
	ruleBasionymAuthorship2Parens
	ruleAuthorsGroup
	ruleAuthorsTeam
	ruleAuthorSep
	ruleAuthorSep1
	ruleAuthorSep2
	ruleAuthorEx
	ruleAuthorEmend
	ruleAuthor
	ruleAuthor1
	ruleAuthor2
	ruleUnknownAuthor
	ruleAuthorWord
	ruleAuthorEtAl
	ruleAuthorWord2
	ruleAuthorWord3
	ruleAuthorWordSoft
	ruleCapAuthorWord
	ruleAllCapsAuthorWord
	ruleFilius
	ruleAuthorPrefixGlued
	ruleAuthorPrefix
	ruleAuthorPrefix2
	ruleAuthorPrefix1
	ruleAuthorUpperChar
	ruleAuthorLowerChar
	ruleYear
	ruleYearRange
	ruleYearWithDot
	ruleYearApprox
	ruleYearWithPage
	ruleYearWithParens
	ruleYearWithChar
	ruleYearNum
	ruleNameUpperChar
	ruleUpperCharExtended
	ruleUpperChar
	ruleNameLowerChar
	ruleMiscodedChar
	ruleLowerCharExtended
	ruleLowerChar
	ruleSpaceCharEOI
	rulenums
	rulelASCII
	rulehASCII
	ruleapostr
	ruledash
	rule_
	ruleMultipleSpace
	ruleSingleSpace
	ruleOtherSpace
	ruleAction0
)

var rul3s = [...]string{
	"Unknown",
	"SciName",
	"Tail",
	"Name",
	"HybridFormula",
	"HybridFormulaFull",
	"HybridFormulaPart",
	"NamedHybrid",
	"NamedSpeciesHybrid",
	"NamedGenusHybrid",
	"SingleName",
	"NameUninomial",
	"NameApprox",
	"NameComp",
	"NameSpecies",
	"GenusWord",
	"InfraspGroup",
	"InfraspEpithet",
	"SpeciesEpithet",
	"Comparison",
	"Rank",
	"RankOtherUncommon",
	"RankOther",
	"RankVar",
	"RankForma",
	"RankSsp",
	"SubGenusOrSuperspecies",
	"SubGenus",
	"UninomialCombo",
	"UninomialCombo1",
	"UninomialCombo2",
	"RankUninomial",
	"Uninomial",
	"UninomialWord",
	"AbbrGenus",
	"CapWord",
	"CapWord1",
	"CapWord2",
	"TwoLetterGenus",
	"Word",
	"Word1",
	"WordStartsWithDigit",
	"Word2",
	"WordApostr",
	"Word4",
	"HybridChar",
	"ApproxNameIgnored",
	"Approximation",
	"Authorship",
	"AuthorshipCombo",
	"OriginalAuthorship",
	"OriginalAuthorshipComb",
	"CombinationAuthorship",
	"BasionymAuthorshipYearMisformed",
	"BasionymAuthorship",
	"BasionymAuthorship1",
	"BasionymAuthorship2Parens",
	"AuthorsGroup",
	"AuthorsTeam",
	"AuthorSep",
	"AuthorSep1",
	"AuthorSep2",
	"AuthorEx",
	"AuthorEmend",
	"Author",
	"Author1",
	"Author2",
	"UnknownAuthor",
	"AuthorWord",
	"AuthorEtAl",
	"AuthorWord2",
	"AuthorWord3",
	"AuthorWordSoft",
	"CapAuthorWord",
	"AllCapsAuthorWord",
	"Filius",
	"AuthorPrefixGlued",
	"AuthorPrefix",
	"AuthorPrefix2",
	"AuthorPrefix1",
	"AuthorUpperChar",
	"AuthorLowerChar",
	"Year",
	"YearRange",
	"YearWithDot",
	"YearApprox",
	"YearWithPage",
	"YearWithParens",
	"YearWithChar",
	"YearNum",
	"NameUpperChar",
	"UpperCharExtended",
	"UpperChar",
	"NameLowerChar",
	"MiscodedChar",
	"LowerCharExtended",
	"LowerChar",
	"SpaceCharEOI",
	"nums",
	"lASCII",
	"hASCII",
	"apostr",
	"dash",
	"_",
	"MultipleSpace",
	"SingleSpace",
	"OtherSpace",
	"Action0",
}

type token32 struct {
	pegRule
	begin, end uint32
}

func (t *token32) String() string {
	return fmt.Sprintf("\x1B[34m%v\x1B[m %v %v", rul3s[t.pegRule], t.begin, t.end)
}

type node32 struct {
	token32
	up, next *node32
}

func (node *node32) print(w io.Writer, pretty bool, buffer string) {
	var print func(node *node32, depth int)
	print = func(node *node32, depth int) {
		for node != nil {
			for c := 0; c < depth; c++ {
				fmt.Fprintf(w, " ")
			}
			rule := rul3s[node.pegRule]
			quote := strconv.Quote(string(([]rune(buffer)[node.begin:node.end])))
			if !pretty {
				fmt.Fprintf(w, "%v %v\n", rule, quote)
			} else {
				fmt.Fprintf(w, "\x1B[34m%v\x1B[m %v\n", rule, quote)
			}
			if node.up != nil {
				print(node.up, depth+1)
			}
			node = node.next
		}
	}
	print(node, 0)
}

func (node *node32) Print(w io.Writer, buffer string) {
	node.print(w, false, buffer)
}

func (node *node32) PrettyPrint(w io.Writer, buffer string) {
	node.print(w, true, buffer)
}

type tokens32 struct {
	tree []token32
}

func (t *tokens32) Trim(length uint32) {
	t.tree = t.tree[:length]
}

func (t *tokens32) Print() {
	for _, token := range t.tree {
		fmt.Println(token.String())
	}
}

func (t *tokens32) AST() *node32 {
	type element struct {
		node *node32
		down *element
	}
	tokens := t.Tokens()
	var stack *element
	for _, token := range tokens {
		if token.begin == token.end {
			continue
		}
		node := &node32{token32: token}
		for stack != nil && stack.node.begin >= token.begin && stack.node.end <= token.end {
			stack.node.next = node.up
			node.up = stack.node
			stack = stack.down
		}
		stack = &element{node: node, down: stack}
	}
	if stack != nil {
		return stack.node
	}
	return nil
}

func (t *tokens32) PrintSyntaxTree(buffer string) {
	t.AST().Print(os.Stdout, buffer)
}

func (t *tokens32) WriteSyntaxTree(w io.Writer, buffer string) {
	t.AST().Print(w, buffer)
}

func (t *tokens32) PrettyPrintSyntaxTree(buffer string) {
	t.AST().PrettyPrint(os.Stdout, buffer)
}

func (t *tokens32) Add(rule pegRule, begin, end, index uint32) {
	if tree := t.tree; int(index) >= len(tree) {
		expanded := make([]token32, 2*len(tree))
		copy(expanded, tree)
		t.tree = expanded
	}
	t.tree[index] = token32{
		pegRule: rule,
		begin:   begin,
		end:     end,
	}
}

func (t *tokens32) Tokens() []token32 {
	return t.tree
}

type Engine struct {
	BaseEngine

	Buffer string
	buffer []rune
	rules  [108]func() bool
	parse  func(rule ...int) error
	reset  func()
	Pretty bool
	tokens32
}

func (p *Engine) Parse(rule ...int) error {
	return p.parse(rule...)
}

func (p *Engine) Reset() {
	p.reset()
}

type textPosition struct {
	line, symbol int
}

type textPositionMap map[int]textPosition

func translatePositions(buffer []rune, positions []int) textPositionMap {
	length, translations, j, line, symbol := len(positions), make(textPositionMap, len(positions)), 0, 1, 0
	sort.Ints(positions)

search:
	for i, c := range buffer {
		if c == '\n' {
			line, symbol = line+1, 0
		} else {
			symbol++
		}
		if i == positions[j] {
			translations[positions[j]] = textPosition{line, symbol}
			for j++; j < length; j++ {
				if i != positions[j] {
					continue search
				}
			}
			break search
		}
	}

	return translations
}

type parseError struct {
	p   *Engine
	max token32
}

func (e *parseError) Error() string {
	tokens, error := []token32{e.max}, "\n"
	positions, p := make([]int, 2*len(tokens)), 0
	for _, token := range tokens {
		positions[p], p = int(token.begin), p+1
		positions[p], p = int(token.end), p+1
	}
	translations := translatePositions(e.p.buffer, positions)
	format := "parse error near %v (line %v symbol %v - line %v symbol %v):\n%v\n"
	if e.p.Pretty {
		format = "parse error near \x1B[34m%v\x1B[m (line %v symbol %v - line %v symbol %v):\n%v\n"
	}
	for _, token := range tokens {
		begin, end := int(token.begin), int(token.end)
		error += fmt.Sprintf(format,
			rul3s[token.pegRule],
			translations[begin].line, translations[begin].symbol,
			translations[end].line, translations[end].symbol,
			strconv.Quote(string(e.p.buffer[begin:end])))
	}

	return error
}

func (p *Engine) PrintSyntaxTree() {
	if p.Pretty {
		p.tokens32.PrettyPrintSyntaxTree(p.Buffer)
	} else {
		p.tokens32.PrintSyntaxTree(p.Buffer)
	}
}

func (p *Engine) WriteSyntaxTree(w io.Writer) {
	p.tokens32.WriteSyntaxTree(w, p.Buffer)
}

func (p *Engine) Execute() {
	buffer, _buffer, text, begin, end := p.Buffer, p.buffer, "", 0, 0
	for _, token := range p.Tokens() {
		switch token.pegRule {

		case ruleAction0:
			p.AddWarn(YearCharWarn)

		}
	}
	_, _, _, _, _ = buffer, _buffer, text, begin, end
}

func (p *Engine) Init() {
	var (
		max                  token32
		position, tokenIndex uint32
		buffer               []rune
	)
	p.reset = func() {
		max = token32{}
		position, tokenIndex = 0, 0

		p.buffer = []rune(p.Buffer)
		if len(p.buffer) == 0 || p.buffer[len(p.buffer)-1] != endSymbol {
			p.buffer = append(p.buffer, endSymbol)
		}
		buffer = p.buffer
	}
	p.reset()

	_rules := p.rules
	tree := tokens32{tree: make([]token32, math.MaxInt16)}
	p.parse = func(rule ...int) error {
		r := 1
		if len(rule) > 0 {
			r = rule[0]
		}
		matches := p.rules[r]()
		p.tokens32 = tree
		if matches {
			p.Trim(tokenIndex)
			return nil
		}
		return &parseError{p, max}
	}

	add := func(rule pegRule, begin uint32) {
		tree.Add(rule, begin, position, tokenIndex)
		tokenIndex++
		if begin != position && position > max.end {
			max = token32{rule, begin, position}
		}
	}

	matchDot := func() bool {
		if buffer[position] != endSymbol {
			position++
			return true
		}
		return false
	}

	/*matchChar := func(c byte) bool {
		if buffer[position] == c {
			position++
			return true
		}
		return false
	}*/

	/*matchRange := func(lower byte, upper byte) bool {
		if c := buffer[position]; c >= lower && c <= upper {
			position++
			return true
		}
		return false
	}*/

	_rules = [...]func() bool{
		nil,
		/* 0 SciName <- <(_? Name Tail !.)> */
		func() bool {
			position0, tokenIndex0 := position, tokenIndex
			{
				position1 := position
				{
					position2, tokenIndex2 := position, tokenIndex
					if !_rules[rule_]() {
						goto l2
					}
					goto l3
				l2:
					position, tokenIndex = position2, tokenIndex2
				}
			l3:
				if !_rules[ruleName]() {
					goto l0
				}
				if !_rules[ruleTail]() {
					goto l0
				}
				{
					position4, tokenIndex4 := position, tokenIndex
					if !matchDot() {
						goto l4
					}
					goto l0
				l4:
					position, tokenIndex = position4, tokenIndex4
				}
				add(ruleSciName, position1)
			}
			return true
		l0:
			position, tokenIndex = position0, tokenIndex0
			return false
		},
		/* 1 Tail <- <((_ / ';' / ',') .*)?> */
		func() bool {
			{
				position6 := position
				{
					position7, tokenIndex7 := position, tokenIndex
					{
						position9, tokenIndex9 := position, tokenIndex
						if !_rules[rule_]() {
							goto l10
						}
						goto l9
					l10:
						position, tokenIndex = position9, tokenIndex9
						if buffer[position] != rune(';') {
							goto l11
						}
						position++
						goto l9
					l11:
						position, tokenIndex = position9, tokenIndex9
						if buffer[position] != rune(',') {
							goto l7
						}
						position++
					}
				l9:
				l12:
					{
						position13, tokenIndex13 := position, tokenIndex
						if !matchDot() {
							goto l13
						}
						goto l12
					l13:
						position, tokenIndex = position13, tokenIndex13
					}
					goto l8
				l7:
					position, tokenIndex = position7, tokenIndex7
				}
			l8:
				add(ruleTail, position6)
			}
			return true
		},
		/* 2 Name <- <(NamedHybrid / HybridFormula / SingleName)> */
		func() bool {
			position14, tokenIndex14 := position, tokenIndex
			{
				position15 := position
				{
					position16, tokenIndex16 := position, tokenIndex
					if !_rules[ruleNamedHybrid]() {
						goto l17
					}
					goto l16
				l17:
					position, tokenIndex = position16, tokenIndex16
					if !_rules[ruleHybridFormula]() {
						goto l18
					}
					goto l16
				l18:
					position, tokenIndex = position16, tokenIndex16
					if !_rules[ruleSingleName]() {
						goto l14
					}
				}
			l16:
				add(ruleName, position15)
			}
			return true
		l14:
			position, tokenIndex = position14, tokenIndex14
			return false
		},
		/* 3 HybridFormula <- <(SingleName (_ (HybridFormulaPart / HybridFormulaFull))+)> */
		func() bool {
			position19, tokenIndex19 := position, tokenIndex
			{
				position20 := position
				if !_rules[ruleSingleName]() {
					goto l19
				}
				if !_rules[rule_]() {
					goto l19
				}
				{
					position23, tokenIndex23 := position, tokenIndex
					if !_rules[ruleHybridFormulaPart]() {
						goto l24
					}
					goto l23
				l24:
					position, tokenIndex = position23, tokenIndex23
					if !_rules[ruleHybridFormulaFull]() {
						goto l19
					}
				}
			l23:
			l21:
				{
					position22, tokenIndex22 := position, tokenIndex
					if !_rules[rule_]() {
						goto l22
					}
					{
						position25, tokenIndex25 := position, tokenIndex
						if !_rules[ruleHybridFormulaPart]() {
							goto l26
						}
						goto l25
					l26:
						position, tokenIndex = position25, tokenIndex25
						if !_rules[ruleHybridFormulaFull]() {
							goto l22
						}
					}
				l25:
					goto l21
				l22:
					position, tokenIndex = position22, tokenIndex22
				}
				add(ruleHybridFormula, position20)
			}
			return true
		l19:
			position, tokenIndex = position19, tokenIndex19
			return false
		},
		/* 4 HybridFormulaFull <- <(HybridChar (_ SingleName)?)> */
		func() bool {
			position27, tokenIndex27 := position, tokenIndex
			{
				position28 := position
				if !_rules[ruleHybridChar]() {
					goto l27
				}
				{
					position29, tokenIndex29 := position, tokenIndex
					if !_rules[rule_]() {
						goto l29
					}
					if !_rules[ruleSingleName]() {
						goto l29
					}
					goto l30
				l29:
					position, tokenIndex = position29, tokenIndex29
				}
			l30:
				add(ruleHybridFormulaFull, position28)
			}
			return true
		l27:
			position, tokenIndex = position27, tokenIndex27
			return false
		},
		/* 5 HybridFormulaPart <- <(HybridChar _ SpeciesEpithet (_ InfraspGroup)?)> */
		func() bool {
			position31, tokenIndex31 := position, tokenIndex
			{
				position32 := position
				if !_rules[ruleHybridChar]() {
					goto l31
				}
				if !_rules[rule_]() {
					goto l31
				}
				if !_rules[ruleSpeciesEpithet]() {
					goto l31
				}
				{
					position33, tokenIndex33 := position, tokenIndex
					if !_rules[rule_]() {
						goto l33
					}
					if !_rules[ruleInfraspGroup]() {
						goto l33
					}
					goto l34
				l33:
					position, tokenIndex = position33, tokenIndex33
				}
			l34:
				add(ruleHybridFormulaPart, position32)
			}
			return true
		l31:
			position, tokenIndex = position31, tokenIndex31
			return false
		},
		/* 6 NamedHybrid <- <(NamedGenusHybrid / NamedSpeciesHybrid)> */
		func() bool {
			position35, tokenIndex35 := position, tokenIndex
			{
				position36 := position
				{
					position37, tokenIndex37 := position, tokenIndex
					if !_rules[ruleNamedGenusHybrid]() {
						goto l38
					}
					goto l37
				l38:
					position, tokenIndex = position37, tokenIndex37
					if !_rules[ruleNamedSpeciesHybrid]() {
						goto l35
					}
				}
			l37:
				add(ruleNamedHybrid, position36)
			}
			return true
		l35:
			position, tokenIndex = position35, tokenIndex35
			return false
		},
		/* 7 NamedSpeciesHybrid <- <(GenusWord _ HybridChar _? SpeciesEpithet)> */
		func() bool {
			position39, tokenIndex39 := position, tokenIndex
			{
				position40 := position
				if !_rules[ruleGenusWord]() {
					goto l39
				}
				if !_rules[rule_]() {
					goto l39
				}
				if !_rules[ruleHybridChar]() {
					goto l39
				}
				{
					position41, tokenIndex41 := position, tokenIndex
					if !_rules[rule_]() {
						goto l41
					}
					goto l42
				l41:
					position, tokenIndex = position41, tokenIndex41
				}
			l42:
				if !_rules[ruleSpeciesEpithet]() {
					goto l39
				}
				add(ruleNamedSpeciesHybrid, position40)
			}
			return true
		l39:
			position, tokenIndex = position39, tokenIndex39
			return false
		},
		/* 8 NamedGenusHybrid <- <(HybridChar _? SingleName)> */
		func() bool {
			position43, tokenIndex43 := position, tokenIndex
			{
				position44 := position
				if !_rules[ruleHybridChar]() {
					goto l43
				}
				{
					position45, tokenIndex45 := position, tokenIndex
					if !_rules[rule_]() {
						goto l45
					}
					goto l46
				l45:
					position, tokenIndex = position45, tokenIndex45
				}
			l46:
				if !_rules[ruleSingleName]() {
					goto l43
				}
				add(ruleNamedGenusHybrid, position44)
			}
			return true
		l43:
			position, tokenIndex = position43, tokenIndex43
			return false
		},
		/* 9 SingleName <- <(NameComp / NameApprox / NameSpecies / NameUninomial)> */
		func() bool {
			position47, tokenIndex47 := position, tokenIndex
			{
				position48 := position
				{
					position49, tokenIndex49 := position, tokenIndex
					if !_rules[ruleNameComp]() {
						goto l50
					}
					goto l49
				l50:
					position, tokenIndex = position49, tokenIndex49
					if !_rules[ruleNameApprox]() {
						goto l51
					}
					goto l49
				l51:
					position, tokenIndex = position49, tokenIndex49
					if !_rules[ruleNameSpecies]() {
						goto l52
					}
					goto l49
				l52:
					position, tokenIndex = position49, tokenIndex49
					if !_rules[ruleNameUninomial]() {
						goto l47
					}
				}
			l49:
				add(ruleSingleName, position48)
			}
			return true
		l47:
			position, tokenIndex = position47, tokenIndex47
			return false
		},
		/* 10 NameUninomial <- <(UninomialCombo / Uninomial)> */
		func() bool {
			position53, tokenIndex53 := position, tokenIndex
			{
				position54 := position
				{
					position55, tokenIndex55 := position, tokenIndex
					if !_rules[ruleUninomialCombo]() {
						goto l56
					}
					goto l55
				l56:
					position, tokenIndex = position55, tokenIndex55
					if !_rules[ruleUninomial]() {
						goto l53
					}
				}
			l55:
				add(ruleNameUninomial, position54)
			}
			return true
		l53:
			position, tokenIndex = position53, tokenIndex53
			return false
		},
		/* 11 NameApprox <- <(GenusWord (_ SpeciesEpithet)? _ Approximation ApproxNameIgnored)> */
		func() bool {
			position57, tokenIndex57 := position, tokenIndex
			{
				position58 := position
				if !_rules[ruleGenusWord]() {
					goto l57
				}
				{
					position59, tokenIndex59 := position, tokenIndex
					if !_rules[rule_]() {
						goto l59
					}
					if !_rules[ruleSpeciesEpithet]() {
						goto l59
					}
					goto l60
				l59:
					position, tokenIndex = position59, tokenIndex59
				}
			l60:
				if !_rules[rule_]() {
					goto l57
				}
				if !_rules[ruleApproximation]() {
					goto l57
				}
				if !_rules[ruleApproxNameIgnored]() {
					goto l57
				}
				add(ruleNameApprox, position58)
			}
			return true
		l57:
			position, tokenIndex = position57, tokenIndex57
			return false
		},
		/* 12 NameComp <- <(GenusWord _ Comparison (_ SpeciesEpithet)?)> */
		func() bool {
			position61, tokenIndex61 := position, tokenIndex
			{
				position62 := position
				if !_rules[ruleGenusWord]() {
					goto l61
				}
				if !_rules[rule_]() {
					goto l61
				}
				if !_rules[ruleComparison]() {
					goto l61
				}
				{
					position63, tokenIndex63 := position, tokenIndex
					if !_rules[rule_]() {
						goto l63
					}
					if !_rules[ruleSpeciesEpithet]() {
						goto l63
					}
					goto l64
				l63:
					position, tokenIndex = position63, tokenIndex63
				}
			l64:
				add(ruleNameComp, position62)
			}
			return true
		l61:
			position, tokenIndex = position61, tokenIndex61
			return false
		},
		/* 13 NameSpecies <- <(GenusWord (_? (SubGenus / SubGenusOrSuperspecies))? _ SpeciesEpithet (_ InfraspGroup)?)> */
		func() bool {
			position65, tokenIndex65 := position, tokenIndex
			{
				position66 := position
				if !_rules[ruleGenusWord]() {
					goto l65
				}
				{
					position67, tokenIndex67 := position, tokenIndex
					{
						position69, tokenIndex69 := position, tokenIndex
						if !_rules[rule_]() {
							goto l69
						}
						goto l70
					l69:
						position, tokenIndex = position69, tokenIndex69
					}
				l70:
					{
						position71, tokenIndex71 := position, tokenIndex
						if !_rules[ruleSubGenus]() {
							goto l72
						}
						goto l71
					l72:
						position, tokenIndex = position71, tokenIndex71
						if !_rules[ruleSubGenusOrSuperspecies]() {
							goto l67
						}
					}
				l71:
					goto l68
				l67:
					position, tokenIndex = position67, tokenIndex67
				}
			l68:
				if !_rules[rule_]() {
					goto l65
				}
				if !_rules[ruleSpeciesEpithet]() {
					goto l65
				}
				{
					position73, tokenIndex73 := position, tokenIndex
					if !_rules[rule_]() {
						goto l73
					}
					if !_rules[ruleInfraspGroup]() {
						goto l73
					}
					goto l74
				l73:
					position, tokenIndex = position73, tokenIndex73
				}
			l74:
				add(ruleNameSpecies, position66)
			}
			return true
		l65:
			position, tokenIndex = position65, tokenIndex65
			return false
		},
		/* 14 GenusWord <- <((AbbrGenus / UninomialWord) !(_ AuthorWord))> */
		func() bool {
			position75, tokenIndex75 := position, tokenIndex
			{
				position76 := position
				{
					position77, tokenIndex77 := position, tokenIndex
					if !_rules[ruleAbbrGenus]() {
						goto l78
					}
					goto l77
				l78:
					position, tokenIndex = position77, tokenIndex77
					if !_rules[ruleUninomialWord]() {
						goto l75
					}
				}
			l77:
				{
					position79, tokenIndex79 := position, tokenIndex
					if !_rules[rule_]() {
						goto l79
					}
					if !_rules[ruleAuthorWord]() {
						goto l79
					}
					goto l75
				l79:
					position, tokenIndex = position79, tokenIndex79
				}
				add(ruleGenusWord, position76)
			}
			return true
		l75:
			position, tokenIndex = position75, tokenIndex75
			return false
		},
		/* 15 InfraspGroup <- <(InfraspEpithet (_ InfraspEpithet)? (_ InfraspEpithet)?)> */
		func() bool {
			position80, tokenIndex80 := position, tokenIndex
			{
				position81 := position
				if !_rules[ruleInfraspEpithet]() {
					goto l80
				}
				{
					position82, tokenIndex82 := position, tokenIndex
					if !_rules[rule_]() {
						goto l82
					}
					if !_rules[ruleInfraspEpithet]() {
						goto l82
					}
					goto l83
				l82:
					position, tokenIndex = position82, tokenIndex82
				}
			l83:
				{
					position84, tokenIndex84 := position, tokenIndex
					if !_rules[rule_]() {
						goto l84
					}
					if !_rules[ruleInfraspEpithet]() {
						goto l84
					}
					goto l85
				l84:
					position, tokenIndex = position84, tokenIndex84
				}
			l85:
				add(ruleInfraspGroup, position81)
			}
			return true
		l80:
			position, tokenIndex = position80, tokenIndex80
			return false
		},
		/* 16 InfraspEpithet <- <((Rank _?)? !AuthorEx Word (_ Authorship)?)> */
		func() bool {
			position86, tokenIndex86 := position, tokenIndex
			{
				position87 := position
				{
					position88, tokenIndex88 := position, tokenIndex
					if !_rules[ruleRank]() {
						goto l88
					}
					{
						position90, tokenIndex90 := position, tokenIndex
						if !_rules[rule_]() {
							goto l90
						}
						goto l91
					l90:
						position, tokenIndex = position90, tokenIndex90
					}
				l91:
					goto l89
				l88:
					position, tokenIndex = position88, tokenIndex88
				}
			l89:
				{
					position92, tokenIndex92 := position, tokenIndex
					if !_rules[ruleAuthorEx]() {
						goto l92
					}
					goto l86
				l92:
					position, tokenIndex = position92, tokenIndex92
				}
				if !_rules[ruleWord]() {
					goto l86
				}
				{
					position93, tokenIndex93 := position, tokenIndex
					if !_rules[rule_]() {
						goto l93
					}
					if !_rules[ruleAuthorship]() {
						goto l93
					}
					goto l94
				l93:
					position, tokenIndex = position93, tokenIndex93
				}
			l94:
				add(ruleInfraspEpithet, position87)
			}
			return true
		l86:
			position, tokenIndex = position86, tokenIndex86
			return false
		},
		/* 17 SpeciesEpithet <- <(!AuthorEx Word (_? Authorship)?)> */
		func() bool {
			position95, tokenIndex95 := position, tokenIndex
			{
				position96 := position
				{
					position97, tokenIndex97 := position, tokenIndex
					if !_rules[ruleAuthorEx]() {
						goto l97
					}
					goto l95
				l97:
					position, tokenIndex = position97, tokenIndex97
				}
				if !_rules[ruleWord]() {
					goto l95
				}
				{
					position98, tokenIndex98 := position, tokenIndex
					{
						position100, tokenIndex100 := position, tokenIndex
						if !_rules[rule_]() {
							goto l100
						}
						goto l101
					l100:
						position, tokenIndex = position100, tokenIndex100
					}
				l101:
					if !_rules[ruleAuthorship]() {
						goto l98
					}
					goto l99
				l98:
					position, tokenIndex = position98, tokenIndex98
				}
			l99:
				add(ruleSpeciesEpithet, position96)
			}
			return true
		l95:
			position, tokenIndex = position95, tokenIndex95
			return false
		},
		/* 18 Comparison <- <('c' 'f' '.'?)> */
		func() bool {
			position102, tokenIndex102 := position, tokenIndex
			{
				position103 := position
				if buffer[position] != rune('c') {
					goto l102
				}
				position++
				if buffer[position] != rune('f') {
					goto l102
				}
				position++
				{
					position104, tokenIndex104 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l104
					}
					position++
					goto l105
				l104:
					position, tokenIndex = position104, tokenIndex104
				}
			l105:
				add(ruleComparison, position103)
			}
			return true
		l102:
			position, tokenIndex = position102, tokenIndex102
			return false
		},
		/* 19 Rank <- <(RankForma / RankVar / RankSsp / RankOther / RankOtherUncommon)> */
		func() bool {
			position106, tokenIndex106 := position, tokenIndex
			{
				position107 := position
				{
					position108, tokenIndex108 := position, tokenIndex
					if !_rules[ruleRankForma]() {
						goto l109
					}
					goto l108
				l109:
					position, tokenIndex = position108, tokenIndex108
					if !_rules[ruleRankVar]() {
						goto l110
					}
					goto l108
				l110:
					position, tokenIndex = position108, tokenIndex108
					if !_rules[ruleRankSsp]() {
						goto l111
					}
					goto l108
				l111:
					position, tokenIndex = position108, tokenIndex108
					if !_rules[ruleRankOther]() {
						goto l112
					}
					goto l108
				l112:
					position, tokenIndex = position108, tokenIndex108
					if !_rules[ruleRankOtherUncommon]() {
						goto l106
					}
				}
			l108:
				add(ruleRank, position107)
			}
			return true
		l106:
			position, tokenIndex = position106, tokenIndex106
			return false
		},
		/* 20 RankOtherUncommon <- <(('*' / ('n' 'a' 't') / ('f' '.' 's' 'p') / ('m' 'u' 't' '.')) &SpaceCharEOI)> */
		func() bool {
			position113, tokenIndex113 := position, tokenIndex
			{
				position114 := position
				{
					position115, tokenIndex115 := position, tokenIndex
					if buffer[position] != rune('*') {
						goto l116
					}
					position++
					goto l115
				l116:
					position, tokenIndex = position115, tokenIndex115
					if buffer[position] != rune('n') {
						goto l117
					}
					position++
					if buffer[position] != rune('a') {
						goto l117
					}
					position++
					if buffer[position] != rune('t') {
						goto l117
					}
					position++
					goto l115
				l117:
					position, tokenIndex = position115, tokenIndex115
					if buffer[position] != rune('f') {
						goto l118
					}
					position++
					if buffer[position] != rune('.') {
						goto l118
					}
					position++
					if buffer[position] != rune('s') {
						goto l118
					}
					position++
					if buffer[position] != rune('p') {
						goto l118
					}
					position++
					goto l115
				l118:
					position, tokenIndex = position115, tokenIndex115
					if buffer[position] != rune('m') {
						goto l113
					}
					position++
					if buffer[position] != rune('u') {
						goto l113
					}
					position++
					if buffer[position] != rune('t') {
						goto l113
					}
					position++
					if buffer[position] != rune('.') {
						goto l113
					}
					position++
				}
			l115:
				{
					position119, tokenIndex119 := position, tokenIndex
					if !_rules[ruleSpaceCharEOI]() {
						goto l113
					}
					position, tokenIndex = position119, tokenIndex119
				}
				add(ruleRankOtherUncommon, position114)
			}
			return true
		l113:
			position, tokenIndex = position113, tokenIndex113
			return false
		},
		/* 21 RankOther <- <((('m' 'o' 'r' 'p' 'h' '.') / ('n' 'o' 't' 'h' 'o' 's' 'u' 'b' 's' 'p' '.') / ('c' 'o' 'n' 'v' 'a' 'r' '.') / ('p' 's' 'e' 'u' 'd' 'o' 'v' 'a' 'r' '.') / ('s' 'e' 'c' 't' '.') / ('s' 'e' 'r' '.') / ('s' 'u' 'b' 'v' 'a' 'r' '.') / ('s' 'u' 'b' 'f' '.') / ('r' 'a' 'c' 'e') / 'α' / ('β' 'β') / 'β' / 'γ' / 'δ' / 'ε' / 'φ' / 'θ' / 'μ' / ('a' '.') / ('b' '.') / ('c' '.') / ('d' '.') / ('e' '.') / ('g' '.') / ('k' '.') / ('p' 'v' '.') / ('p' 'a' 't' 'h' 'o' 'v' 'a' 'r' '.') / ('a' 'b' '.' (_? ('n' '.'))?) / ('s' 't' '.')) &SpaceCharEOI)> */
		func() bool {
			position120, tokenIndex120 := position, tokenIndex
			{
				position121 := position
				{
					position122, tokenIndex122 := position, tokenIndex
					if buffer[position] != rune('m') {
						goto l123
					}
					position++
					if buffer[position] != rune('o') {
						goto l123
					}
					position++
					if buffer[position] != rune('r') {
						goto l123
					}
					position++
					if buffer[position] != rune('p') {
						goto l123
					}
					position++
					if buffer[position] != rune('h') {
						goto l123
					}
					position++
					if buffer[position] != rune('.') {
						goto l123
					}
					position++
					goto l122
				l123:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('n') {
						goto l124
					}
					position++
					if buffer[position] != rune('o') {
						goto l124
					}
					position++
					if buffer[position] != rune('t') {
						goto l124
					}
					position++
					if buffer[position] != rune('h') {
						goto l124
					}
					position++
					if buffer[position] != rune('o') {
						goto l124
					}
					position++
					if buffer[position] != rune('s') {
						goto l124
					}
					position++
					if buffer[position] != rune('u') {
						goto l124
					}
					position++
					if buffer[position] != rune('b') {
						goto l124
					}
					position++
					if buffer[position] != rune('s') {
						goto l124
					}
					position++
					if buffer[position] != rune('p') {
						goto l124
					}
					position++
					if buffer[position] != rune('.') {
						goto l124
					}
					position++
					goto l122
				l124:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('c') {
						goto l125
					}
					position++
					if buffer[position] != rune('o') {
						goto l125
					}
					position++
					if buffer[position] != rune('n') {
						goto l125
					}
					position++
					if buffer[position] != rune('v') {
						goto l125
					}
					position++
					if buffer[position] != rune('a') {
						goto l125
					}
					position++
					if buffer[position] != rune('r') {
						goto l125
					}
					position++
					if buffer[position] != rune('.') {
						goto l125
					}
					position++
					goto l122
				l125:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('p') {
						goto l126
					}
					position++
					if buffer[position] != rune('s') {
						goto l126
					}
					position++
					if buffer[position] != rune('e') {
						goto l126
					}
					position++
					if buffer[position] != rune('u') {
						goto l126
					}
					position++
					if buffer[position] != rune('d') {
						goto l126
					}
					position++
					if buffer[position] != rune('o') {
						goto l126
					}
					position++
					if buffer[position] != rune('v') {
						goto l126
					}
					position++
					if buffer[position] != rune('a') {
						goto l126
					}
					position++
					if buffer[position] != rune('r') {
						goto l126
					}
					position++
					if buffer[position] != rune('.') {
						goto l126
					}
					position++
					goto l122
				l126:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('s') {
						goto l127
					}
					position++
					if buffer[position] != rune('e') {
						goto l127
					}
					position++
					if buffer[position] != rune('c') {
						goto l127
					}
					position++
					if buffer[position] != rune('t') {
						goto l127
					}
					position++
					if buffer[position] != rune('.') {
						goto l127
					}
					position++
					goto l122
				l127:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('s') {
						goto l128
					}
					position++
					if buffer[position] != rune('e') {
						goto l128
					}
					position++
					if buffer[position] != rune('r') {
						goto l128
					}
					position++
					if buffer[position] != rune('.') {
						goto l128
					}
					position++
					goto l122
				l128:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('s') {
						goto l129
					}
					position++
					if buffer[position] != rune('u') {
						goto l129
					}
					position++
					if buffer[position] != rune('b') {
						goto l129
					}
					position++
					if buffer[position] != rune('v') {
						goto l129
					}
					position++
					if buffer[position] != rune('a') {
						goto l129
					}
					position++
					if buffer[position] != rune('r') {
						goto l129
					}
					position++
					if buffer[position] != rune('.') {
						goto l129
					}
					position++
					goto l122
				l129:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('s') {
						goto l130
					}
					position++
					if buffer[position] != rune('u') {
						goto l130
					}
					position++
					if buffer[position] != rune('b') {
						goto l130
					}
					position++
					if buffer[position] != rune('f') {
						goto l130
					}
					position++
					if buffer[position] != rune('.') {
						goto l130
					}
					position++
					goto l122
				l130:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('r') {
						goto l131
					}
					position++
					if buffer[position] != rune('a') {
						goto l131
					}
					position++
					if buffer[position] != rune('c') {
						goto l131
					}
					position++
					if buffer[position] != rune('e') {
						goto l131
					}
					position++
					goto l122
				l131:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('α') {
						goto l132
					}
					position++
					goto l122
				l132:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('β') {
						goto l133
					}
					position++
					if buffer[position] != rune('β') {
						goto l133
					}
					position++
					goto l122
				l133:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('β') {
						goto l134
					}
					position++
					goto l122
				l134:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('γ') {
						goto l135
					}
					position++
					goto l122
				l135:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('δ') {
						goto l136
					}
					position++
					goto l122
				l136:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('ε') {
						goto l137
					}
					position++
					goto l122
				l137:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('φ') {
						goto l138
					}
					position++
					goto l122
				l138:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('θ') {
						goto l139
					}
					position++
					goto l122
				l139:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('μ') {
						goto l140
					}
					position++
					goto l122
				l140:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('a') {
						goto l141
					}
					position++
					if buffer[position] != rune('.') {
						goto l141
					}
					position++
					goto l122
				l141:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('b') {
						goto l142
					}
					position++
					if buffer[position] != rune('.') {
						goto l142
					}
					position++
					goto l122
				l142:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('c') {
						goto l143
					}
					position++
					if buffer[position] != rune('.') {
						goto l143
					}
					position++
					goto l122
				l143:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('d') {
						goto l144
					}
					position++
					if buffer[position] != rune('.') {
						goto l144
					}
					position++
					goto l122
				l144:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('e') {
						goto l145
					}
					position++
					if buffer[position] != rune('.') {
						goto l145
					}
					position++
					goto l122
				l145:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('g') {
						goto l146
					}
					position++
					if buffer[position] != rune('.') {
						goto l146
					}
					position++
					goto l122
				l146:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('k') {
						goto l147
					}
					position++
					if buffer[position] != rune('.') {
						goto l147
					}
					position++
					goto l122
				l147:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('p') {
						goto l148
					}
					position++
					if buffer[position] != rune('v') {
						goto l148
					}
					position++
					if buffer[position] != rune('.') {
						goto l148
					}
					position++
					goto l122
				l148:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('p') {
						goto l149
					}
					position++
					if buffer[position] != rune('a') {
						goto l149
					}
					position++
					if buffer[position] != rune('t') {
						goto l149
					}
					position++
					if buffer[position] != rune('h') {
						goto l149
					}
					position++
					if buffer[position] != rune('o') {
						goto l149
					}
					position++
					if buffer[position] != rune('v') {
						goto l149
					}
					position++
					if buffer[position] != rune('a') {
						goto l149
					}
					position++
					if buffer[position] != rune('r') {
						goto l149
					}
					position++
					if buffer[position] != rune('.') {
						goto l149
					}
					position++
					goto l122
				l149:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('a') {
						goto l150
					}
					position++
					if buffer[position] != rune('b') {
						goto l150
					}
					position++
					if buffer[position] != rune('.') {
						goto l150
					}
					position++
					{
						position151, tokenIndex151 := position, tokenIndex
						{
							position153, tokenIndex153 := position, tokenIndex
							if !_rules[rule_]() {
								goto l153
							}
							goto l154
						l153:
							position, tokenIndex = position153, tokenIndex153
						}
					l154:
						if buffer[position] != rune('n') {
							goto l151
						}
						position++
						if buffer[position] != rune('.') {
							goto l151
						}
						position++
						goto l152
					l151:
						position, tokenIndex = position151, tokenIndex151
					}
				l152:
					goto l122
				l150:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('s') {
						goto l120
					}
					position++
					if buffer[position] != rune('t') {
						goto l120
					}
					position++
					if buffer[position] != rune('.') {
						goto l120
					}
					position++
				}
			l122:
				{
					position155, tokenIndex155 := position, tokenIndex
					if !_rules[ruleSpaceCharEOI]() {
						goto l120
					}
					position, tokenIndex = position155, tokenIndex155
				}
				add(ruleRankOther, position121)
			}
			return true
		l120:
			position, tokenIndex = position120, tokenIndex120
			return false
		},
		/* 22 RankVar <- <(('v' 'a' 'r' 'i' 'e' 't' 'y') / ('[' 'v' 'a' 'r' '.' ']') / ('n' 'v' 'a' 'r' '.') / ('v' 'a' 'r' (&SpaceCharEOI / '.')))> */
		func() bool {
			position156, tokenIndex156 := position, tokenIndex
			{
				position157 := position
				{
					position158, tokenIndex158 := position, tokenIndex
					if buffer[position] != rune('v') {
						goto l159
					}
					position++
					if buffer[position] != rune('a') {
						goto l159
					}
					position++
					if buffer[position] != rune('r') {
						goto l159
					}
					position++
					if buffer[position] != rune('i') {
						goto l159
					}
					position++
					if buffer[position] != rune('e') {
						goto l159
					}
					position++
					if buffer[position] != rune('t') {
						goto l159
					}
					position++
					if buffer[position] != rune('y') {
						goto l159
					}
					position++
					goto l158
				l159:
					position, tokenIndex = position158, tokenIndex158
					if buffer[position] != rune('[') {
						goto l160
					}
					position++
					if buffer[position] != rune('v') {
						goto l160
					}
					position++
					if buffer[position] != rune('a') {
						goto l160
					}
					position++
					if buffer[position] != rune('r') {
						goto l160
					}
					position++
					if buffer[position] != rune('.') {
						goto l160
					}
					position++
					if buffer[position] != rune(']') {
						goto l160
					}
					position++
					goto l158
				l160:
					position, tokenIndex = position158, tokenIndex158
					if buffer[position] != rune('n') {
						goto l161
					}
					position++
					if buffer[position] != rune('v') {
						goto l161
					}
					position++
					if buffer[position] != rune('a') {
						goto l161
					}
					position++
					if buffer[position] != rune('r') {
						goto l161
					}
					position++
					if buffer[position] != rune('.') {
						goto l161
					}
					position++
					goto l158
				l161:
					position, tokenIndex = position158, tokenIndex158
					if buffer[position] != rune('v') {
						goto l156
					}
					position++
					if buffer[position] != rune('a') {
						goto l156
					}
					position++
					if buffer[position] != rune('r') {
						goto l156
					}
					position++
					{
						position162, tokenIndex162 := position, tokenIndex
						{
							position164, tokenIndex164 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l163
							}
							position, tokenIndex = position164, tokenIndex164
						}
						goto l162
					l163:
						position, tokenIndex = position162, tokenIndex162
						if buffer[position] != rune('.') {
							goto l156
						}
						position++
					}
				l162:
				}
			l158:
				add(ruleRankVar, position157)
			}
			return true
		l156:
			position, tokenIndex = position156, tokenIndex156
			return false
		},
		/* 23 RankForma <- <((('f' 'o' 'r' 'm' 'a') / ('f' 'm' 'a') / ('f' 'o' 'r' 'm') / ('f' 'o') / 'f') (&SpaceCharEOI / '.'))> */
		func() bool {
			position165, tokenIndex165 := position, tokenIndex
			{
				position166 := position
				{
					position167, tokenIndex167 := position, tokenIndex
					if buffer[position] != rune('f') {
						goto l168
					}
					position++
					if buffer[position] != rune('o') {
						goto l168
					}
					position++
					if buffer[position] != rune('r') {
						goto l168
					}
					position++
					if buffer[position] != rune('m') {
						goto l168
					}
					position++
					if buffer[position] != rune('a') {
						goto l168
					}
					position++
					goto l167
				l168:
					position, tokenIndex = position167, tokenIndex167
					if buffer[position] != rune('f') {
						goto l169
					}
					position++
					if buffer[position] != rune('m') {
						goto l169
					}
					position++
					if buffer[position] != rune('a') {
						goto l169
					}
					position++
					goto l167
				l169:
					position, tokenIndex = position167, tokenIndex167
					if buffer[position] != rune('f') {
						goto l170
					}
					position++
					if buffer[position] != rune('o') {
						goto l170
					}
					position++
					if buffer[position] != rune('r') {
						goto l170
					}
					position++
					if buffer[position] != rune('m') {
						goto l170
					}
					position++
					goto l167
				l170:
					position, tokenIndex = position167, tokenIndex167
					if buffer[position] != rune('f') {
						goto l171
					}
					position++
					if buffer[position] != rune('o') {
						goto l171
					}
					position++
					goto l167
				l171:
					position, tokenIndex = position167, tokenIndex167
					if buffer[position] != rune('f') {
						goto l165
					}
					position++
				}
			l167:
				{
					position172, tokenIndex172 := position, tokenIndex
					{
						position174, tokenIndex174 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l173
						}
						position, tokenIndex = position174, tokenIndex174
					}
					goto l172
				l173:
					position, tokenIndex = position172, tokenIndex172
					if buffer[position] != rune('.') {
						goto l165
					}
					position++
				}
			l172:
				add(ruleRankForma, position166)
			}
			return true
		l165:
			position, tokenIndex = position165, tokenIndex165
			return false
		},
		/* 24 RankSsp <- <((('s' 's' 'p') / ('s' 'u' 'b' 's' 'p')) (&SpaceCharEOI / '.'))> */
		func() bool {
			position175, tokenIndex175 := position, tokenIndex
			{
				position176 := position
				{
					position177, tokenIndex177 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l178
					}
					position++
					if buffer[position] != rune('s') {
						goto l178
					}
					position++
					if buffer[position] != rune('p') {
						goto l178
					}
					position++
					goto l177
				l178:
					position, tokenIndex = position177, tokenIndex177
					if buffer[position] != rune('s') {
						goto l175
					}
					position++
					if buffer[position] != rune('u') {
						goto l175
					}
					position++
					if buffer[position] != rune('b') {
						goto l175
					}
					position++
					if buffer[position] != rune('s') {
						goto l175
					}
					position++
					if buffer[position] != rune('p') {
						goto l175
					}
					position++
				}
			l177:
				{
					position179, tokenIndex179 := position, tokenIndex
					{
						position181, tokenIndex181 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l180
						}
						position, tokenIndex = position181, tokenIndex181
					}
					goto l179
				l180:
					position, tokenIndex = position179, tokenIndex179
					if buffer[position] != rune('.') {
						goto l175
					}
					position++
				}
			l179:
				add(ruleRankSsp, position176)
			}
			return true
		l175:
			position, tokenIndex = position175, tokenIndex175
			return false
		},
		/* 25 SubGenusOrSuperspecies <- <('(' _? NameLowerChar+ _? ')')> */
		func() bool {
			position182, tokenIndex182 := position, tokenIndex
			{
				position183 := position
				if buffer[position] != rune('(') {
					goto l182
				}
				position++
				{
					position184, tokenIndex184 := position, tokenIndex
					if !_rules[rule_]() {
						goto l184
					}
					goto l185
				l184:
					position, tokenIndex = position184, tokenIndex184
				}
			l185:
				if !_rules[ruleNameLowerChar]() {
					goto l182
				}
			l186:
				{
					position187, tokenIndex187 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l187
					}
					goto l186
				l187:
					position, tokenIndex = position187, tokenIndex187
				}
				{
					position188, tokenIndex188 := position, tokenIndex
					if !_rules[rule_]() {
						goto l188
					}
					goto l189
				l188:
					position, tokenIndex = position188, tokenIndex188
				}
			l189:
				if buffer[position] != rune(')') {
					goto l182
				}
				position++
				add(ruleSubGenusOrSuperspecies, position183)
			}
			return true
		l182:
			position, tokenIndex = position182, tokenIndex182
			return false
		},
		/* 26 SubGenus <- <('(' _? UninomialWord _? ')')> */
		func() bool {
			position190, tokenIndex190 := position, tokenIndex
			{
				position191 := position
				if buffer[position] != rune('(') {
					goto l190
				}
				position++
				{
					position192, tokenIndex192 := position, tokenIndex
					if !_rules[rule_]() {
						goto l192
					}
					goto l193
				l192:
					position, tokenIndex = position192, tokenIndex192
				}
			l193:
				if !_rules[ruleUninomialWord]() {
					goto l190
				}
				{
					position194, tokenIndex194 := position, tokenIndex
					if !_rules[rule_]() {
						goto l194
					}
					goto l195
				l194:
					position, tokenIndex = position194, tokenIndex194
				}
			l195:
				if buffer[position] != rune(')') {
					goto l190
				}
				position++
				add(ruleSubGenus, position191)
			}
			return true
		l190:
			position, tokenIndex = position190, tokenIndex190
			return false
		},
		/* 27 UninomialCombo <- <(UninomialCombo1 / UninomialCombo2)> */
		func() bool {
			position196, tokenIndex196 := position, tokenIndex
			{
				position197 := position
				{
					position198, tokenIndex198 := position, tokenIndex
					if !_rules[ruleUninomialCombo1]() {
						goto l199
					}
					goto l198
				l199:
					position, tokenIndex = position198, tokenIndex198
					if !_rules[ruleUninomialCombo2]() {
						goto l196
					}
				}
			l198:
				add(ruleUninomialCombo, position197)
			}
			return true
		l196:
			position, tokenIndex = position196, tokenIndex196
			return false
		},
		/* 28 UninomialCombo1 <- <(UninomialWord _? SubGenus (_? Authorship)?)> */
		func() bool {
			position200, tokenIndex200 := position, tokenIndex
			{
				position201 := position
				if !_rules[ruleUninomialWord]() {
					goto l200
				}
				{
					position202, tokenIndex202 := position, tokenIndex
					if !_rules[rule_]() {
						goto l202
					}
					goto l203
				l202:
					position, tokenIndex = position202, tokenIndex202
				}
			l203:
				if !_rules[ruleSubGenus]() {
					goto l200
				}
				{
					position204, tokenIndex204 := position, tokenIndex
					{
						position206, tokenIndex206 := position, tokenIndex
						if !_rules[rule_]() {
							goto l206
						}
						goto l207
					l206:
						position, tokenIndex = position206, tokenIndex206
					}
				l207:
					if !_rules[ruleAuthorship]() {
						goto l204
					}
					goto l205
				l204:
					position, tokenIndex = position204, tokenIndex204
				}
			l205:
				add(ruleUninomialCombo1, position201)
			}
			return true
		l200:
			position, tokenIndex = position200, tokenIndex200
			return false
		},
		/* 29 UninomialCombo2 <- <(Uninomial _ RankUninomial _ Uninomial)> */
		func() bool {
			position208, tokenIndex208 := position, tokenIndex
			{
				position209 := position
				if !_rules[ruleUninomial]() {
					goto l208
				}
				if !_rules[rule_]() {
					goto l208
				}
				if !_rules[ruleRankUninomial]() {
					goto l208
				}
				if !_rules[rule_]() {
					goto l208
				}
				if !_rules[ruleUninomial]() {
					goto l208
				}
				add(ruleUninomialCombo2, position209)
			}
			return true
		l208:
			position, tokenIndex = position208, tokenIndex208
			return false
		},
		/* 30 RankUninomial <- <((('s' 'e' 'c' 't') / ('s' 'u' 'b' 's' 'e' 'c' 't') / ('t' 'r' 'i' 'b') / ('s' 'u' 'b' 't' 'r' 'i' 'b') / ('s' 'u' 'b' 's' 'e' 'r') / ('s' 'e' 'r' '.') / ('s' 'u' 'b' 'g' 'e' 'n') / ('f' 'a' 'm') / ('s' 'u' 'b' 'f' 'a' 'm') / ('s' 'u' 'p' 'e' 'r' 't' 'r' 'i' 'b')) '.'?)> */
		func() bool {
			position210, tokenIndex210 := position, tokenIndex
			{
				position211 := position
				{
					position212, tokenIndex212 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l213
					}
					position++
					if buffer[position] != rune('e') {
						goto l213
					}
					position++
					if buffer[position] != rune('c') {
						goto l213
					}
					position++
					if buffer[position] != rune('t') {
						goto l213
					}
					position++
					goto l212
				l213:
					position, tokenIndex = position212, tokenIndex212
					if buffer[position] != rune('s') {
						goto l214
					}
					position++
					if buffer[position] != rune('u') {
						goto l214
					}
					position++
					if buffer[position] != rune('b') {
						goto l214
					}
					position++
					if buffer[position] != rune('s') {
						goto l214
					}
					position++
					if buffer[position] != rune('e') {
						goto l214
					}
					position++
					if buffer[position] != rune('c') {
						goto l214
					}
					position++
					if buffer[position] != rune('t') {
						goto l214
					}
					position++
					goto l212
				l214:
					position, tokenIndex = position212, tokenIndex212
					if buffer[position] != rune('t') {
						goto l215
					}
					position++
					if buffer[position] != rune('r') {
						goto l215
					}
					position++
					if buffer[position] != rune('i') {
						goto l215
					}
					position++
					if buffer[position] != rune('b') {
						goto l215
					}
					position++
					goto l212
				l215:
					position, tokenIndex = position212, tokenIndex212
					if buffer[position] != rune('s') {
						goto l216
					}
					position++
					if buffer[position] != rune('u') {
						goto l216
					}
					position++
					if buffer[position] != rune('b') {
						goto l216
					}
					position++
					if buffer[position] != rune('t') {
						goto l216
					}
					position++
					if buffer[position] != rune('r') {
						goto l216
					}
					position++
					if buffer[position] != rune('i') {
						goto l216
					}
					position++
					if buffer[position] != rune('b') {
						goto l216
					}
					position++
					goto l212
				l216:
					position, tokenIndex = position212, tokenIndex212
					if buffer[position] != rune('s') {
						goto l217
					}
					position++
					if buffer[position] != rune('u') {
						goto l217
					}
					position++
					if buffer[position] != rune('b') {
						goto l217
					}
					position++
					if buffer[position] != rune('s') {
						goto l217
					}
					position++
					if buffer[position] != rune('e') {
						goto l217
					}
					position++
					if buffer[position] != rune('r') {
						goto l217
					}
					position++
					goto l212
				l217:
					position, tokenIndex = position212, tokenIndex212
					if buffer[position] != rune('s') {
						goto l218
					}
					position++
					if buffer[position] != rune('e') {
						goto l218
					}
					position++
					if buffer[position] != rune('r') {
						goto l218
					}
					position++
					if buffer[position] != rune('.') {
						goto l218
					}
					position++
					goto l212
				l218:
					position, tokenIndex = position212, tokenIndex212
					if buffer[position] != rune('s') {
						goto l219
					}
					position++
					if buffer[position] != rune('u') {
						goto l219
					}
					position++
					if buffer[position] != rune('b') {
						goto l219
					}
					position++
					if buffer[position] != rune('g') {
						goto l219
					}
					position++
					if buffer[position] != rune('e') {
						goto l219
					}
					position++
					if buffer[position] != rune('n') {
						goto l219
					}
					position++
					goto l212
				l219:
					position, tokenIndex = position212, tokenIndex212
					if buffer[position] != rune('f') {
						goto l220
					}
					position++
					if buffer[position] != rune('a') {
						goto l220
					}
					position++
					if buffer[position] != rune('m') {
						goto l220
					}
					position++
					goto l212
				l220:
					position, tokenIndex = position212, tokenIndex212
					if buffer[position] != rune('s') {
						goto l221
					}
					position++
					if buffer[position] != rune('u') {
						goto l221
					}
					position++
					if buffer[position] != rune('b') {
						goto l221
					}
					position++
					if buffer[position] != rune('f') {
						goto l221
					}
					position++
					if buffer[position] != rune('a') {
						goto l221
					}
					position++
					if buffer[position] != rune('m') {
						goto l221
					}
					position++
					goto l212
				l221:
					position, tokenIndex = position212, tokenIndex212
					if buffer[position] != rune('s') {
						goto l210
					}
					position++
					if buffer[position] != rune('u') {
						goto l210
					}
					position++
					if buffer[position] != rune('p') {
						goto l210
					}
					position++
					if buffer[position] != rune('e') {
						goto l210
					}
					position++
					if buffer[position] != rune('r') {
						goto l210
					}
					position++
					if buffer[position] != rune('t') {
						goto l210
					}
					position++
					if buffer[position] != rune('r') {
						goto l210
					}
					position++
					if buffer[position] != rune('i') {
						goto l210
					}
					position++
					if buffer[position] != rune('b') {
						goto l210
					}
					position++
				}
			l212:
				{
					position222, tokenIndex222 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l222
					}
					position++
					goto l223
				l222:
					position, tokenIndex = position222, tokenIndex222
				}
			l223:
				add(ruleRankUninomial, position211)
			}
			return true
		l210:
			position, tokenIndex = position210, tokenIndex210
			return false
		},
		/* 31 Uninomial <- <(UninomialWord (_ Authorship)?)> */
		func() bool {
			position224, tokenIndex224 := position, tokenIndex
			{
				position225 := position
				if !_rules[ruleUninomialWord]() {
					goto l224
				}
				{
					position226, tokenIndex226 := position, tokenIndex
					if !_rules[rule_]() {
						goto l226
					}
					if !_rules[ruleAuthorship]() {
						goto l226
					}
					goto l227
				l226:
					position, tokenIndex = position226, tokenIndex226
				}
			l227:
				add(ruleUninomial, position225)
			}
			return true
		l224:
			position, tokenIndex = position224, tokenIndex224
			return false
		},
		/* 32 UninomialWord <- <(CapWord / TwoLetterGenus)> */
		func() bool {
			position228, tokenIndex228 := position, tokenIndex
			{
				position229 := position
				{
					position230, tokenIndex230 := position, tokenIndex
					if !_rules[ruleCapWord]() {
						goto l231
					}
					goto l230
				l231:
					position, tokenIndex = position230, tokenIndex230
					if !_rules[ruleTwoLetterGenus]() {
						goto l228
					}
				}
			l230:
				add(ruleUninomialWord, position229)
			}
			return true
		l228:
			position, tokenIndex = position228, tokenIndex228
			return false
		},
		/* 33 AbbrGenus <- <(UpperChar LowerChar? '.')> */
		func() bool {
			position232, tokenIndex232 := position, tokenIndex
			{
				position233 := position
				if !_rules[ruleUpperChar]() {
					goto l232
				}
				{
					position234, tokenIndex234 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l234
					}
					goto l235
				l234:
					position, tokenIndex = position234, tokenIndex234
				}
			l235:
				if buffer[position] != rune('.') {
					goto l232
				}
				position++
				add(ruleAbbrGenus, position233)
			}
			return true
		l232:
			position, tokenIndex = position232, tokenIndex232
			return false
		},
		/* 34 CapWord <- <(CapWord2 / CapWord1)> */
		func() bool {
			position236, tokenIndex236 := position, tokenIndex
			{
				position237 := position
				{
					position238, tokenIndex238 := position, tokenIndex
					if !_rules[ruleCapWord2]() {
						goto l239
					}
					goto l238
				l239:
					position, tokenIndex = position238, tokenIndex238
					if !_rules[ruleCapWord1]() {
						goto l236
					}
				}
			l238:
				add(ruleCapWord, position237)
			}
			return true
		l236:
			position, tokenIndex = position236, tokenIndex236
			return false
		},
		/* 35 CapWord1 <- <(NameUpperChar NameLowerChar NameLowerChar+ '?'?)> */
		func() bool {
			position240, tokenIndex240 := position, tokenIndex
			{
				position241 := position
				if !_rules[ruleNameUpperChar]() {
					goto l240
				}
				if !_rules[ruleNameLowerChar]() {
					goto l240
				}
				if !_rules[ruleNameLowerChar]() {
					goto l240
				}
			l242:
				{
					position243, tokenIndex243 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l243
					}
					goto l242
				l243:
					position, tokenIndex = position243, tokenIndex243
				}
				{
					position244, tokenIndex244 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l244
					}
					position++
					goto l245
				l244:
					position, tokenIndex = position244, tokenIndex244
				}
			l245:
				add(ruleCapWord1, position241)
			}
			return true
		l240:
			position, tokenIndex = position240, tokenIndex240
			return false
		},
		/* 36 CapWord2 <- <(CapWord1 dash (CapWord1 / Word1))> */
		func() bool {
			position246, tokenIndex246 := position, tokenIndex
			{
				position247 := position
				if !_rules[ruleCapWord1]() {
					goto l246
				}
				if !_rules[ruledash]() {
					goto l246
				}
				{
					position248, tokenIndex248 := position, tokenIndex
					if !_rules[ruleCapWord1]() {
						goto l249
					}
					goto l248
				l249:
					position, tokenIndex = position248, tokenIndex248
					if !_rules[ruleWord1]() {
						goto l246
					}
				}
			l248:
				add(ruleCapWord2, position247)
			}
			return true
		l246:
			position, tokenIndex = position246, tokenIndex246
			return false
		},
		/* 37 TwoLetterGenus <- <(('C' 'a') / ('E' 'a') / ('G' 'e') / ('I' 'a') / ('I' 'o') / ('I' 'x') / ('L' 'o') / ('O' 'a') / ('R' 'a') / ('T' 'y') / ('U' 'a') / ('A' 'a') / ('J' 'a') / ('Z' 'u') / ('L' 'a') / ('Q' 'u') / ('A' 's') / ('B' 'a'))> */
		func() bool {
			position250, tokenIndex250 := position, tokenIndex
			{
				position251 := position
				{
					position252, tokenIndex252 := position, tokenIndex
					if buffer[position] != rune('C') {
						goto l253
					}
					position++
					if buffer[position] != rune('a') {
						goto l253
					}
					position++
					goto l252
				l253:
					position, tokenIndex = position252, tokenIndex252
					if buffer[position] != rune('E') {
						goto l254
					}
					position++
					if buffer[position] != rune('a') {
						goto l254
					}
					position++
					goto l252
				l254:
					position, tokenIndex = position252, tokenIndex252
					if buffer[position] != rune('G') {
						goto l255
					}
					position++
					if buffer[position] != rune('e') {
						goto l255
					}
					position++
					goto l252
				l255:
					position, tokenIndex = position252, tokenIndex252
					if buffer[position] != rune('I') {
						goto l256
					}
					position++
					if buffer[position] != rune('a') {
						goto l256
					}
					position++
					goto l252
				l256:
					position, tokenIndex = position252, tokenIndex252
					if buffer[position] != rune('I') {
						goto l257
					}
					position++
					if buffer[position] != rune('o') {
						goto l257
					}
					position++
					goto l252
				l257:
					position, tokenIndex = position252, tokenIndex252
					if buffer[position] != rune('I') {
						goto l258
					}
					position++
					if buffer[position] != rune('x') {
						goto l258
					}
					position++
					goto l252
				l258:
					position, tokenIndex = position252, tokenIndex252
					if buffer[position] != rune('L') {
						goto l259
					}
					position++
					if buffer[position] != rune('o') {
						goto l259
					}
					position++
					goto l252
				l259:
					position, tokenIndex = position252, tokenIndex252
					if buffer[position] != rune('O') {
						goto l260
					}
					position++
					if buffer[position] != rune('a') {
						goto l260
					}
					position++
					goto l252
				l260:
					position, tokenIndex = position252, tokenIndex252
					if buffer[position] != rune('R') {
						goto l261
					}
					position++
					if buffer[position] != rune('a') {
						goto l261
					}
					position++
					goto l252
				l261:
					position, tokenIndex = position252, tokenIndex252
					if buffer[position] != rune('T') {
						goto l262
					}
					position++
					if buffer[position] != rune('y') {
						goto l262
					}
					position++
					goto l252
				l262:
					position, tokenIndex = position252, tokenIndex252
					if buffer[position] != rune('U') {
						goto l263
					}
					position++
					if buffer[position] != rune('a') {
						goto l263
					}
					position++
					goto l252
				l263:
					position, tokenIndex = position252, tokenIndex252
					if buffer[position] != rune('A') {
						goto l264
					}
					position++
					if buffer[position] != rune('a') {
						goto l264
					}
					position++
					goto l252
				l264:
					position, tokenIndex = position252, tokenIndex252
					if buffer[position] != rune('J') {
						goto l265
					}
					position++
					if buffer[position] != rune('a') {
						goto l265
					}
					position++
					goto l252
				l265:
					position, tokenIndex = position252, tokenIndex252
					if buffer[position] != rune('Z') {
						goto l266
					}
					position++
					if buffer[position] != rune('u') {
						goto l266
					}
					position++
					goto l252
				l266:
					position, tokenIndex = position252, tokenIndex252
					if buffer[position] != rune('L') {
						goto l267
					}
					position++
					if buffer[position] != rune('a') {
						goto l267
					}
					position++
					goto l252
				l267:
					position, tokenIndex = position252, tokenIndex252
					if buffer[position] != rune('Q') {
						goto l268
					}
					position++
					if buffer[position] != rune('u') {
						goto l268
					}
					position++
					goto l252
				l268:
					position, tokenIndex = position252, tokenIndex252
					if buffer[position] != rune('A') {
						goto l269
					}
					position++
					if buffer[position] != rune('s') {
						goto l269
					}
					position++
					goto l252
				l269:
					position, tokenIndex = position252, tokenIndex252
					if buffer[position] != rune('B') {
						goto l250
					}
					position++
					if buffer[position] != rune('a') {
						goto l250
					}
					position++
				}
			l252:
				add(ruleTwoLetterGenus, position251)
			}
			return true
		l250:
			position, tokenIndex = position250, tokenIndex250
			return false
		},
		/* 38 Word <- <(!((AuthorPrefix / RankUninomial / Approximation / Word4) SpaceCharEOI) (WordApostr / WordStartsWithDigit / Word2 / Word1) &(SpaceCharEOI / '('))> */
		func() bool {
			position270, tokenIndex270 := position, tokenIndex
			{
				position271 := position
				{
					position272, tokenIndex272 := position, tokenIndex
					{
						position273, tokenIndex273 := position, tokenIndex
						if !_rules[ruleAuthorPrefix]() {
							goto l274
						}
						goto l273
					l274:
						position, tokenIndex = position273, tokenIndex273
						if !_rules[ruleRankUninomial]() {
							goto l275
						}
						goto l273
					l275:
						position, tokenIndex = position273, tokenIndex273
						if !_rules[ruleApproximation]() {
							goto l276
						}
						goto l273
					l276:
						position, tokenIndex = position273, tokenIndex273
						if !_rules[ruleWord4]() {
							goto l272
						}
					}
				l273:
					if !_rules[ruleSpaceCharEOI]() {
						goto l272
					}
					goto l270
				l272:
					position, tokenIndex = position272, tokenIndex272
				}
				{
					position277, tokenIndex277 := position, tokenIndex
					if !_rules[ruleWordApostr]() {
						goto l278
					}
					goto l277
				l278:
					position, tokenIndex = position277, tokenIndex277
					if !_rules[ruleWordStartsWithDigit]() {
						goto l279
					}
					goto l277
				l279:
					position, tokenIndex = position277, tokenIndex277
					if !_rules[ruleWord2]() {
						goto l280
					}
					goto l277
				l280:
					position, tokenIndex = position277, tokenIndex277
					if !_rules[ruleWord1]() {
						goto l270
					}
				}
			l277:
				{
					position281, tokenIndex281 := position, tokenIndex
					{
						position282, tokenIndex282 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l283
						}
						goto l282
					l283:
						position, tokenIndex = position282, tokenIndex282
						if buffer[position] != rune('(') {
							goto l270
						}
						position++
					}
				l282:
					position, tokenIndex = position281, tokenIndex281
				}
				add(ruleWord, position271)
			}
			return true
		l270:
			position, tokenIndex = position270, tokenIndex270
			return false
		},
		/* 39 Word1 <- <((lASCII dash)? NameLowerChar NameLowerChar+)> */
		func() bool {
			position284, tokenIndex284 := position, tokenIndex
			{
				position285 := position
				{
					position286, tokenIndex286 := position, tokenIndex
					if !_rules[rulelASCII]() {
						goto l286
					}
					if !_rules[ruledash]() {
						goto l286
					}
					goto l287
				l286:
					position, tokenIndex = position286, tokenIndex286
				}
			l287:
				if !_rules[ruleNameLowerChar]() {
					goto l284
				}
				if !_rules[ruleNameLowerChar]() {
					goto l284
				}
			l288:
				{
					position289, tokenIndex289 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l289
					}
					goto l288
				l289:
					position, tokenIndex = position289, tokenIndex289
				}
				add(ruleWord1, position285)
			}
			return true
		l284:
			position, tokenIndex = position284, tokenIndex284
			return false
		},
		/* 40 WordStartsWithDigit <- <(('1' / '2' / '3' / '4' / '5' / '6' / '7' / '8' / '9') nums? ('.' / dash)? NameLowerChar NameLowerChar NameLowerChar NameLowerChar+)> */
		func() bool {
			position290, tokenIndex290 := position, tokenIndex
			{
				position291 := position
				{
					position292, tokenIndex292 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l293
					}
					position++
					goto l292
				l293:
					position, tokenIndex = position292, tokenIndex292
					if buffer[position] != rune('2') {
						goto l294
					}
					position++
					goto l292
				l294:
					position, tokenIndex = position292, tokenIndex292
					if buffer[position] != rune('3') {
						goto l295
					}
					position++
					goto l292
				l295:
					position, tokenIndex = position292, tokenIndex292
					if buffer[position] != rune('4') {
						goto l296
					}
					position++
					goto l292
				l296:
					position, tokenIndex = position292, tokenIndex292
					if buffer[position] != rune('5') {
						goto l297
					}
					position++
					goto l292
				l297:
					position, tokenIndex = position292, tokenIndex292
					if buffer[position] != rune('6') {
						goto l298
					}
					position++
					goto l292
				l298:
					position, tokenIndex = position292, tokenIndex292
					if buffer[position] != rune('7') {
						goto l299
					}
					position++
					goto l292
				l299:
					position, tokenIndex = position292, tokenIndex292
					if buffer[position] != rune('8') {
						goto l300
					}
					position++
					goto l292
				l300:
					position, tokenIndex = position292, tokenIndex292
					if buffer[position] != rune('9') {
						goto l290
					}
					position++
				}
			l292:
				{
					position301, tokenIndex301 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l301
					}
					goto l302
				l301:
					position, tokenIndex = position301, tokenIndex301
				}
			l302:
				{
					position303, tokenIndex303 := position, tokenIndex
					{
						position305, tokenIndex305 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l306
						}
						position++
						goto l305
					l306:
						position, tokenIndex = position305, tokenIndex305
						if !_rules[ruledash]() {
							goto l303
						}
					}
				l305:
					goto l304
				l303:
					position, tokenIndex = position303, tokenIndex303
				}
			l304:
				if !_rules[ruleNameLowerChar]() {
					goto l290
				}
				if !_rules[ruleNameLowerChar]() {
					goto l290
				}
				if !_rules[ruleNameLowerChar]() {
					goto l290
				}
				if !_rules[ruleNameLowerChar]() {
					goto l290
				}
			l307:
				{
					position308, tokenIndex308 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l308
					}
					goto l307
				l308:
					position, tokenIndex = position308, tokenIndex308
				}
				add(ruleWordStartsWithDigit, position291)
			}
			return true
		l290:
			position, tokenIndex = position290, tokenIndex290
			return false
		},
		/* 41 Word2 <- <(NameLowerChar+ dash? NameLowerChar+)> */
		func() bool {
			position309, tokenIndex309 := position, tokenIndex
			{
				position310 := position
				if !_rules[ruleNameLowerChar]() {
					goto l309
				}
			l311:
				{
					position312, tokenIndex312 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l312
					}
					goto l311
				l312:
					position, tokenIndex = position312, tokenIndex312
				}
				{
					position313, tokenIndex313 := position, tokenIndex
					if !_rules[ruledash]() {
						goto l313
					}
					goto l314
				l313:
					position, tokenIndex = position313, tokenIndex313
				}
			l314:
				if !_rules[ruleNameLowerChar]() {
					goto l309
				}
			l315:
				{
					position316, tokenIndex316 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l316
					}
					goto l315
				l316:
					position, tokenIndex = position316, tokenIndex316
				}
				add(ruleWord2, position310)
			}
			return true
		l309:
			position, tokenIndex = position309, tokenIndex309
			return false
		},
		/* 42 WordApostr <- <(NameLowerChar NameLowerChar* apostr Word1)> */
		func() bool {
			position317, tokenIndex317 := position, tokenIndex
			{
				position318 := position
				if !_rules[ruleNameLowerChar]() {
					goto l317
				}
			l319:
				{
					position320, tokenIndex320 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l320
					}
					goto l319
				l320:
					position, tokenIndex = position320, tokenIndex320
				}
				if !_rules[ruleapostr]() {
					goto l317
				}
				if !_rules[ruleWord1]() {
					goto l317
				}
				add(ruleWordApostr, position318)
			}
			return true
		l317:
			position, tokenIndex = position317, tokenIndex317
			return false
		},
		/* 43 Word4 <- <(NameLowerChar+ '.' NameLowerChar)> */
		func() bool {
			position321, tokenIndex321 := position, tokenIndex
			{
				position322 := position
				if !_rules[ruleNameLowerChar]() {
					goto l321
				}
			l323:
				{
					position324, tokenIndex324 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l324
					}
					goto l323
				l324:
					position, tokenIndex = position324, tokenIndex324
				}
				if buffer[position] != rune('.') {
					goto l321
				}
				position++
				if !_rules[ruleNameLowerChar]() {
					goto l321
				}
				add(ruleWord4, position322)
			}
			return true
		l321:
			position, tokenIndex = position321, tokenIndex321
			return false
		},
		/* 44 HybridChar <- <'×'> */
		func() bool {
			position325, tokenIndex325 := position, tokenIndex
			{
				position326 := position
				if buffer[position] != rune('×') {
					goto l325
				}
				position++
				add(ruleHybridChar, position326)
			}
			return true
		l325:
			position, tokenIndex = position325, tokenIndex325
			return false
		},
		/* 45 ApproxNameIgnored <- <.*> */
		func() bool {
			{
				position328 := position
			l329:
				{
					position330, tokenIndex330 := position, tokenIndex
					if !matchDot() {
						goto l330
					}
					goto l329
				l330:
					position, tokenIndex = position330, tokenIndex330
				}
				add(ruleApproxNameIgnored, position328)
			}
			return true
		},
		/* 46 Approximation <- <(('s' 'p' '.' _? ('n' 'r' '.')) / ('s' 'p' '.' _? ('a' 'f' 'f' '.')) / ('m' 'o' 'n' 's' 't' '.') / '?' / ((('s' 'p' 'p') / ('n' 'r') / ('s' 'p') / ('a' 'f' 'f') / ('s' 'p' 'e' 'c' 'i' 'e' 's')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position331, tokenIndex331 := position, tokenIndex
			{
				position332 := position
				{
					position333, tokenIndex333 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l334
					}
					position++
					if buffer[position] != rune('p') {
						goto l334
					}
					position++
					if buffer[position] != rune('.') {
						goto l334
					}
					position++
					{
						position335, tokenIndex335 := position, tokenIndex
						if !_rules[rule_]() {
							goto l335
						}
						goto l336
					l335:
						position, tokenIndex = position335, tokenIndex335
					}
				l336:
					if buffer[position] != rune('n') {
						goto l334
					}
					position++
					if buffer[position] != rune('r') {
						goto l334
					}
					position++
					if buffer[position] != rune('.') {
						goto l334
					}
					position++
					goto l333
				l334:
					position, tokenIndex = position333, tokenIndex333
					if buffer[position] != rune('s') {
						goto l337
					}
					position++
					if buffer[position] != rune('p') {
						goto l337
					}
					position++
					if buffer[position] != rune('.') {
						goto l337
					}
					position++
					{
						position338, tokenIndex338 := position, tokenIndex
						if !_rules[rule_]() {
							goto l338
						}
						goto l339
					l338:
						position, tokenIndex = position338, tokenIndex338
					}
				l339:
					if buffer[position] != rune('a') {
						goto l337
					}
					position++
					if buffer[position] != rune('f') {
						goto l337
					}
					position++
					if buffer[position] != rune('f') {
						goto l337
					}
					position++
					if buffer[position] != rune('.') {
						goto l337
					}
					position++
					goto l333
				l337:
					position, tokenIndex = position333, tokenIndex333
					if buffer[position] != rune('m') {
						goto l340
					}
					position++
					if buffer[position] != rune('o') {
						goto l340
					}
					position++
					if buffer[position] != rune('n') {
						goto l340
					}
					position++
					if buffer[position] != rune('s') {
						goto l340
					}
					position++
					if buffer[position] != rune('t') {
						goto l340
					}
					position++
					if buffer[position] != rune('.') {
						goto l340
					}
					position++
					goto l333
				l340:
					position, tokenIndex = position333, tokenIndex333
					if buffer[position] != rune('?') {
						goto l341
					}
					position++
					goto l333
				l341:
					position, tokenIndex = position333, tokenIndex333
					{
						position342, tokenIndex342 := position, tokenIndex
						if buffer[position] != rune('s') {
							goto l343
						}
						position++
						if buffer[position] != rune('p') {
							goto l343
						}
						position++
						if buffer[position] != rune('p') {
							goto l343
						}
						position++
						goto l342
					l343:
						position, tokenIndex = position342, tokenIndex342
						if buffer[position] != rune('n') {
							goto l344
						}
						position++
						if buffer[position] != rune('r') {
							goto l344
						}
						position++
						goto l342
					l344:
						position, tokenIndex = position342, tokenIndex342
						if buffer[position] != rune('s') {
							goto l345
						}
						position++
						if buffer[position] != rune('p') {
							goto l345
						}
						position++
						goto l342
					l345:
						position, tokenIndex = position342, tokenIndex342
						if buffer[position] != rune('a') {
							goto l346
						}
						position++
						if buffer[position] != rune('f') {
							goto l346
						}
						position++
						if buffer[position] != rune('f') {
							goto l346
						}
						position++
						goto l342
					l346:
						position, tokenIndex = position342, tokenIndex342
						if buffer[position] != rune('s') {
							goto l331
						}
						position++
						if buffer[position] != rune('p') {
							goto l331
						}
						position++
						if buffer[position] != rune('e') {
							goto l331
						}
						position++
						if buffer[position] != rune('c') {
							goto l331
						}
						position++
						if buffer[position] != rune('i') {
							goto l331
						}
						position++
						if buffer[position] != rune('e') {
							goto l331
						}
						position++
						if buffer[position] != rune('s') {
							goto l331
						}
						position++
					}
				l342:
					{
						position347, tokenIndex347 := position, tokenIndex
						{
							position349, tokenIndex349 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l348
							}
							position, tokenIndex = position349, tokenIndex349
						}
						goto l347
					l348:
						position, tokenIndex = position347, tokenIndex347
						if buffer[position] != rune('.') {
							goto l331
						}
						position++
					}
				l347:
				}
			l333:
				add(ruleApproximation, position332)
			}
			return true
		l331:
			position, tokenIndex = position331, tokenIndex331
			return false
		},
		/* 47 Authorship <- <((AuthorshipCombo / OriginalAuthorship) &(SpaceCharEOI / ';' / ','))> */
		func() bool {
			position350, tokenIndex350 := position, tokenIndex
			{
				position351 := position
				{
					position352, tokenIndex352 := position, tokenIndex
					if !_rules[ruleAuthorshipCombo]() {
						goto l353
					}
					goto l352
				l353:
					position, tokenIndex = position352, tokenIndex352
					if !_rules[ruleOriginalAuthorship]() {
						goto l350
					}
				}
			l352:
				{
					position354, tokenIndex354 := position, tokenIndex
					{
						position355, tokenIndex355 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l356
						}
						goto l355
					l356:
						position, tokenIndex = position355, tokenIndex355
						if buffer[position] != rune(';') {
							goto l357
						}
						position++
						goto l355
					l357:
						position, tokenIndex = position355, tokenIndex355
						if buffer[position] != rune(',') {
							goto l350
						}
						position++
					}
				l355:
					position, tokenIndex = position354, tokenIndex354
				}
				add(ruleAuthorship, position351)
			}
			return true
		l350:
			position, tokenIndex = position350, tokenIndex350
			return false
		},
		/* 48 AuthorshipCombo <- <(OriginalAuthorshipComb (_? CombinationAuthorship)?)> */
		func() bool {
			position358, tokenIndex358 := position, tokenIndex
			{
				position359 := position
				if !_rules[ruleOriginalAuthorshipComb]() {
					goto l358
				}
				{
					position360, tokenIndex360 := position, tokenIndex
					{
						position362, tokenIndex362 := position, tokenIndex
						if !_rules[rule_]() {
							goto l362
						}
						goto l363
					l362:
						position, tokenIndex = position362, tokenIndex362
					}
				l363:
					if !_rules[ruleCombinationAuthorship]() {
						goto l360
					}
					goto l361
				l360:
					position, tokenIndex = position360, tokenIndex360
				}
			l361:
				add(ruleAuthorshipCombo, position359)
			}
			return true
		l358:
			position, tokenIndex = position358, tokenIndex358
			return false
		},
		/* 49 OriginalAuthorship <- <AuthorsGroup> */
		func() bool {
			position364, tokenIndex364 := position, tokenIndex
			{
				position365 := position
				if !_rules[ruleAuthorsGroup]() {
					goto l364
				}
				add(ruleOriginalAuthorship, position365)
			}
			return true
		l364:
			position, tokenIndex = position364, tokenIndex364
			return false
		},
		/* 50 OriginalAuthorshipComb <- <(BasionymAuthorshipYearMisformed / BasionymAuthorship)> */
		func() bool {
			position366, tokenIndex366 := position, tokenIndex
			{
				position367 := position
				{
					position368, tokenIndex368 := position, tokenIndex
					if !_rules[ruleBasionymAuthorshipYearMisformed]() {
						goto l369
					}
					goto l368
				l369:
					position, tokenIndex = position368, tokenIndex368
					if !_rules[ruleBasionymAuthorship]() {
						goto l366
					}
				}
			l368:
				add(ruleOriginalAuthorshipComb, position367)
			}
			return true
		l366:
			position, tokenIndex = position366, tokenIndex366
			return false
		},
		/* 51 CombinationAuthorship <- <AuthorsGroup> */
		func() bool {
			position370, tokenIndex370 := position, tokenIndex
			{
				position371 := position
				if !_rules[ruleAuthorsGroup]() {
					goto l370
				}
				add(ruleCombinationAuthorship, position371)
			}
			return true
		l370:
			position, tokenIndex = position370, tokenIndex370
			return false
		},
		/* 52 BasionymAuthorshipYearMisformed <- <('(' _? AuthorsGroup _? ')' (_? ',')? _? Year)> */
		func() bool {
			position372, tokenIndex372 := position, tokenIndex
			{
				position373 := position
				if buffer[position] != rune('(') {
					goto l372
				}
				position++
				{
					position374, tokenIndex374 := position, tokenIndex
					if !_rules[rule_]() {
						goto l374
					}
					goto l375
				l374:
					position, tokenIndex = position374, tokenIndex374
				}
			l375:
				if !_rules[ruleAuthorsGroup]() {
					goto l372
				}
				{
					position376, tokenIndex376 := position, tokenIndex
					if !_rules[rule_]() {
						goto l376
					}
					goto l377
				l376:
					position, tokenIndex = position376, tokenIndex376
				}
			l377:
				if buffer[position] != rune(')') {
					goto l372
				}
				position++
				{
					position378, tokenIndex378 := position, tokenIndex
					{
						position380, tokenIndex380 := position, tokenIndex
						if !_rules[rule_]() {
							goto l380
						}
						goto l381
					l380:
						position, tokenIndex = position380, tokenIndex380
					}
				l381:
					if buffer[position] != rune(',') {
						goto l378
					}
					position++
					goto l379
				l378:
					position, tokenIndex = position378, tokenIndex378
				}
			l379:
				{
					position382, tokenIndex382 := position, tokenIndex
					if !_rules[rule_]() {
						goto l382
					}
					goto l383
				l382:
					position, tokenIndex = position382, tokenIndex382
				}
			l383:
				if !_rules[ruleYear]() {
					goto l372
				}
				add(ruleBasionymAuthorshipYearMisformed, position373)
			}
			return true
		l372:
			position, tokenIndex = position372, tokenIndex372
			return false
		},
		/* 53 BasionymAuthorship <- <(BasionymAuthorship1 / BasionymAuthorship2Parens)> */
		func() bool {
			position384, tokenIndex384 := position, tokenIndex
			{
				position385 := position
				{
					position386, tokenIndex386 := position, tokenIndex
					if !_rules[ruleBasionymAuthorship1]() {
						goto l387
					}
					goto l386
				l387:
					position, tokenIndex = position386, tokenIndex386
					if !_rules[ruleBasionymAuthorship2Parens]() {
						goto l384
					}
				}
			l386:
				add(ruleBasionymAuthorship, position385)
			}
			return true
		l384:
			position, tokenIndex = position384, tokenIndex384
			return false
		},
		/* 54 BasionymAuthorship1 <- <('(' _? AuthorsGroup _? ')')> */
		func() bool {
			position388, tokenIndex388 := position, tokenIndex
			{
				position389 := position
				if buffer[position] != rune('(') {
					goto l388
				}
				position++
				{
					position390, tokenIndex390 := position, tokenIndex
					if !_rules[rule_]() {
						goto l390
					}
					goto l391
				l390:
					position, tokenIndex = position390, tokenIndex390
				}
			l391:
				if !_rules[ruleAuthorsGroup]() {
					goto l388
				}
				{
					position392, tokenIndex392 := position, tokenIndex
					if !_rules[rule_]() {
						goto l392
					}
					goto l393
				l392:
					position, tokenIndex = position392, tokenIndex392
				}
			l393:
				if buffer[position] != rune(')') {
					goto l388
				}
				position++
				add(ruleBasionymAuthorship1, position389)
			}
			return true
		l388:
			position, tokenIndex = position388, tokenIndex388
			return false
		},
		/* 55 BasionymAuthorship2Parens <- <('(' _? '(' _? AuthorsGroup _? ')' _? ')')> */
		func() bool {
			position394, tokenIndex394 := position, tokenIndex
			{
				position395 := position
				if buffer[position] != rune('(') {
					goto l394
				}
				position++
				{
					position396, tokenIndex396 := position, tokenIndex
					if !_rules[rule_]() {
						goto l396
					}
					goto l397
				l396:
					position, tokenIndex = position396, tokenIndex396
				}
			l397:
				if buffer[position] != rune('(') {
					goto l394
				}
				position++
				{
					position398, tokenIndex398 := position, tokenIndex
					if !_rules[rule_]() {
						goto l398
					}
					goto l399
				l398:
					position, tokenIndex = position398, tokenIndex398
				}
			l399:
				if !_rules[ruleAuthorsGroup]() {
					goto l394
				}
				{
					position400, tokenIndex400 := position, tokenIndex
					if !_rules[rule_]() {
						goto l400
					}
					goto l401
				l400:
					position, tokenIndex = position400, tokenIndex400
				}
			l401:
				if buffer[position] != rune(')') {
					goto l394
				}
				position++
				{
					position402, tokenIndex402 := position, tokenIndex
					if !_rules[rule_]() {
						goto l402
					}
					goto l403
				l402:
					position, tokenIndex = position402, tokenIndex402
				}
			l403:
				if buffer[position] != rune(')') {
					goto l394
				}
				position++
				add(ruleBasionymAuthorship2Parens, position395)
			}
			return true
		l394:
			position, tokenIndex = position394, tokenIndex394
			return false
		},
		/* 56 AuthorsGroup <- <(AuthorsTeam (_ (AuthorEmend / AuthorEx) AuthorsTeam)?)> */
		func() bool {
			position404, tokenIndex404 := position, tokenIndex
			{
				position405 := position
				if !_rules[ruleAuthorsTeam]() {
					goto l404
				}
				{
					position406, tokenIndex406 := position, tokenIndex
					if !_rules[rule_]() {
						goto l406
					}
					{
						position408, tokenIndex408 := position, tokenIndex
						if !_rules[ruleAuthorEmend]() {
							goto l409
						}
						goto l408
					l409:
						position, tokenIndex = position408, tokenIndex408
						if !_rules[ruleAuthorEx]() {
							goto l406
						}
					}
				l408:
					if !_rules[ruleAuthorsTeam]() {
						goto l406
					}
					goto l407
				l406:
					position, tokenIndex = position406, tokenIndex406
				}
			l407:
				add(ruleAuthorsGroup, position405)
			}
			return true
		l404:
			position, tokenIndex = position404, tokenIndex404
			return false
		},
		/* 57 AuthorsTeam <- <(Author (AuthorSep Author)* (_? ','? _? Year)?)> */
		func() bool {
			position410, tokenIndex410 := position, tokenIndex
			{
				position411 := position
				if !_rules[ruleAuthor]() {
					goto l410
				}
			l412:
				{
					position413, tokenIndex413 := position, tokenIndex
					if !_rules[ruleAuthorSep]() {
						goto l413
					}
					if !_rules[ruleAuthor]() {
						goto l413
					}
					goto l412
				l413:
					position, tokenIndex = position413, tokenIndex413
				}
				{
					position414, tokenIndex414 := position, tokenIndex
					{
						position416, tokenIndex416 := position, tokenIndex
						if !_rules[rule_]() {
							goto l416
						}
						goto l417
					l416:
						position, tokenIndex = position416, tokenIndex416
					}
				l417:
					{
						position418, tokenIndex418 := position, tokenIndex
						if buffer[position] != rune(',') {
							goto l418
						}
						position++
						goto l419
					l418:
						position, tokenIndex = position418, tokenIndex418
					}
				l419:
					{
						position420, tokenIndex420 := position, tokenIndex
						if !_rules[rule_]() {
							goto l420
						}
						goto l421
					l420:
						position, tokenIndex = position420, tokenIndex420
					}
				l421:
					if !_rules[ruleYear]() {
						goto l414
					}
					goto l415
				l414:
					position, tokenIndex = position414, tokenIndex414
				}
			l415:
				add(ruleAuthorsTeam, position411)
			}
			return true
		l410:
			position, tokenIndex = position410, tokenIndex410
			return false
		},
		/* 58 AuthorSep <- <(AuthorSep1 / AuthorSep2)> */
		func() bool {
			position422, tokenIndex422 := position, tokenIndex
			{
				position423 := position
				{
					position424, tokenIndex424 := position, tokenIndex
					if !_rules[ruleAuthorSep1]() {
						goto l425
					}
					goto l424
				l425:
					position, tokenIndex = position424, tokenIndex424
					if !_rules[ruleAuthorSep2]() {
						goto l422
					}
				}
			l424:
				add(ruleAuthorSep, position423)
			}
			return true
		l422:
			position, tokenIndex = position422, tokenIndex422
			return false
		},
		/* 59 AuthorSep1 <- <(_? (',' _)? ('&' / ('e' 't') / ('a' 'n' 'd') / ('a' 'p' 'u' 'd')) _?)> */
		func() bool {
			position426, tokenIndex426 := position, tokenIndex
			{
				position427 := position
				{
					position428, tokenIndex428 := position, tokenIndex
					if !_rules[rule_]() {
						goto l428
					}
					goto l429
				l428:
					position, tokenIndex = position428, tokenIndex428
				}
			l429:
				{
					position430, tokenIndex430 := position, tokenIndex
					if buffer[position] != rune(',') {
						goto l430
					}
					position++
					if !_rules[rule_]() {
						goto l430
					}
					goto l431
				l430:
					position, tokenIndex = position430, tokenIndex430
				}
			l431:
				{
					position432, tokenIndex432 := position, tokenIndex
					if buffer[position] != rune('&') {
						goto l433
					}
					position++
					goto l432
				l433:
					position, tokenIndex = position432, tokenIndex432
					if buffer[position] != rune('e') {
						goto l434
					}
					position++
					if buffer[position] != rune('t') {
						goto l434
					}
					position++
					goto l432
				l434:
					position, tokenIndex = position432, tokenIndex432
					if buffer[position] != rune('a') {
						goto l435
					}
					position++
					if buffer[position] != rune('n') {
						goto l435
					}
					position++
					if buffer[position] != rune('d') {
						goto l435
					}
					position++
					goto l432
				l435:
					position, tokenIndex = position432, tokenIndex432
					if buffer[position] != rune('a') {
						goto l426
					}
					position++
					if buffer[position] != rune('p') {
						goto l426
					}
					position++
					if buffer[position] != rune('u') {
						goto l426
					}
					position++
					if buffer[position] != rune('d') {
						goto l426
					}
					position++
				}
			l432:
				{
					position436, tokenIndex436 := position, tokenIndex
					if !_rules[rule_]() {
						goto l436
					}
					goto l437
				l436:
					position, tokenIndex = position436, tokenIndex436
				}
			l437:
				add(ruleAuthorSep1, position427)
			}
			return true
		l426:
			position, tokenIndex = position426, tokenIndex426
			return false
		},
		/* 60 AuthorSep2 <- <(_? ',' _?)> */
		func() bool {
			position438, tokenIndex438 := position, tokenIndex
			{
				position439 := position
				{
					position440, tokenIndex440 := position, tokenIndex
					if !_rules[rule_]() {
						goto l440
					}
					goto l441
				l440:
					position, tokenIndex = position440, tokenIndex440
				}
			l441:
				if buffer[position] != rune(',') {
					goto l438
				}
				position++
				{
					position442, tokenIndex442 := position, tokenIndex
					if !_rules[rule_]() {
						goto l442
					}
					goto l443
				l442:
					position, tokenIndex = position442, tokenIndex442
				}
			l443:
				add(ruleAuthorSep2, position439)
			}
			return true
		l438:
			position, tokenIndex = position438, tokenIndex438
			return false
		},
		/* 61 AuthorEx <- <((('e' 'x' '.'?) / ('i' 'n')) _)> */
		func() bool {
			position444, tokenIndex444 := position, tokenIndex
			{
				position445 := position
				{
					position446, tokenIndex446 := position, tokenIndex
					if buffer[position] != rune('e') {
						goto l447
					}
					position++
					if buffer[position] != rune('x') {
						goto l447
					}
					position++
					{
						position448, tokenIndex448 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l448
						}
						position++
						goto l449
					l448:
						position, tokenIndex = position448, tokenIndex448
					}
				l449:
					goto l446
				l447:
					position, tokenIndex = position446, tokenIndex446
					if buffer[position] != rune('i') {
						goto l444
					}
					position++
					if buffer[position] != rune('n') {
						goto l444
					}
					position++
				}
			l446:
				if !_rules[rule_]() {
					goto l444
				}
				add(ruleAuthorEx, position445)
			}
			return true
		l444:
			position, tokenIndex = position444, tokenIndex444
			return false
		},
		/* 62 AuthorEmend <- <('e' 'm' 'e' 'n' 'd' '.'? _)> */
		func() bool {
			position450, tokenIndex450 := position, tokenIndex
			{
				position451 := position
				if buffer[position] != rune('e') {
					goto l450
				}
				position++
				if buffer[position] != rune('m') {
					goto l450
				}
				position++
				if buffer[position] != rune('e') {
					goto l450
				}
				position++
				if buffer[position] != rune('n') {
					goto l450
				}
				position++
				if buffer[position] != rune('d') {
					goto l450
				}
				position++
				{
					position452, tokenIndex452 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l452
					}
					position++
					goto l453
				l452:
					position, tokenIndex = position452, tokenIndex452
				}
			l453:
				if !_rules[rule_]() {
					goto l450
				}
				add(ruleAuthorEmend, position451)
			}
			return true
		l450:
			position, tokenIndex = position450, tokenIndex450
			return false
		},
		/* 63 Author <- <(Author1 / Author2 / UnknownAuthor)> */
		func() bool {
			position454, tokenIndex454 := position, tokenIndex
			{
				position455 := position
				{
					position456, tokenIndex456 := position, tokenIndex
					if !_rules[ruleAuthor1]() {
						goto l457
					}
					goto l456
				l457:
					position, tokenIndex = position456, tokenIndex456
					if !_rules[ruleAuthor2]() {
						goto l458
					}
					goto l456
				l458:
					position, tokenIndex = position456, tokenIndex456
					if !_rules[ruleUnknownAuthor]() {
						goto l454
					}
				}
			l456:
				add(ruleAuthor, position455)
			}
			return true
		l454:
			position, tokenIndex = position454, tokenIndex454
			return false
		},
		/* 64 Author1 <- <(Author2 _? Filius)> */
		func() bool {
			position459, tokenIndex459 := position, tokenIndex
			{
				position460 := position
				if !_rules[ruleAuthor2]() {
					goto l459
				}
				{
					position461, tokenIndex461 := position, tokenIndex
					if !_rules[rule_]() {
						goto l461
					}
					goto l462
				l461:
					position, tokenIndex = position461, tokenIndex461
				}
			l462:
				if !_rules[ruleFilius]() {
					goto l459
				}
				add(ruleAuthor1, position460)
			}
			return true
		l459:
			position, tokenIndex = position459, tokenIndex459
			return false
		},
		/* 65 Author2 <- <(AuthorWord (_? AuthorWord)*)> */
		func() bool {
			position463, tokenIndex463 := position, tokenIndex
			{
				position464 := position
				if !_rules[ruleAuthorWord]() {
					goto l463
				}
			l465:
				{
					position466, tokenIndex466 := position, tokenIndex
					{
						position467, tokenIndex467 := position, tokenIndex
						if !_rules[rule_]() {
							goto l467
						}
						goto l468
					l467:
						position, tokenIndex = position467, tokenIndex467
					}
				l468:
					if !_rules[ruleAuthorWord]() {
						goto l466
					}
					goto l465
				l466:
					position, tokenIndex = position466, tokenIndex466
				}
				add(ruleAuthor2, position464)
			}
			return true
		l463:
			position, tokenIndex = position463, tokenIndex463
			return false
		},
		/* 66 UnknownAuthor <- <('?' / ((('a' 'u' 'c' 't') / ('a' 'n' 'o' 'n')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position469, tokenIndex469 := position, tokenIndex
			{
				position470 := position
				{
					position471, tokenIndex471 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l472
					}
					position++
					goto l471
				l472:
					position, tokenIndex = position471, tokenIndex471
					{
						position473, tokenIndex473 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l474
						}
						position++
						if buffer[position] != rune('u') {
							goto l474
						}
						position++
						if buffer[position] != rune('c') {
							goto l474
						}
						position++
						if buffer[position] != rune('t') {
							goto l474
						}
						position++
						goto l473
					l474:
						position, tokenIndex = position473, tokenIndex473
						if buffer[position] != rune('a') {
							goto l469
						}
						position++
						if buffer[position] != rune('n') {
							goto l469
						}
						position++
						if buffer[position] != rune('o') {
							goto l469
						}
						position++
						if buffer[position] != rune('n') {
							goto l469
						}
						position++
					}
				l473:
					{
						position475, tokenIndex475 := position, tokenIndex
						{
							position477, tokenIndex477 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l476
							}
							position, tokenIndex = position477, tokenIndex477
						}
						goto l475
					l476:
						position, tokenIndex = position475, tokenIndex475
						if buffer[position] != rune('.') {
							goto l469
						}
						position++
					}
				l475:
				}
			l471:
				add(ruleUnknownAuthor, position470)
			}
			return true
		l469:
			position, tokenIndex = position469, tokenIndex469
			return false
		},
		/* 67 AuthorWord <- <(!(('b' / 'B') ('o' / 'O') ('l' / 'L') ('d' / 'D') ':') (AuthorEtAl / AuthorWord2 / AuthorWord3 / AuthorPrefix))> */
		func() bool {
			position478, tokenIndex478 := position, tokenIndex
			{
				position479 := position
				{
					position480, tokenIndex480 := position, tokenIndex
					{
						position481, tokenIndex481 := position, tokenIndex
						if buffer[position] != rune('b') {
							goto l482
						}
						position++
						goto l481
					l482:
						position, tokenIndex = position481, tokenIndex481
						if buffer[position] != rune('B') {
							goto l480
						}
						position++
					}
				l481:
					{
						position483, tokenIndex483 := position, tokenIndex
						if buffer[position] != rune('o') {
							goto l484
						}
						position++
						goto l483
					l484:
						position, tokenIndex = position483, tokenIndex483
						if buffer[position] != rune('O') {
							goto l480
						}
						position++
					}
				l483:
					{
						position485, tokenIndex485 := position, tokenIndex
						if buffer[position] != rune('l') {
							goto l486
						}
						position++
						goto l485
					l486:
						position, tokenIndex = position485, tokenIndex485
						if buffer[position] != rune('L') {
							goto l480
						}
						position++
					}
				l485:
					{
						position487, tokenIndex487 := position, tokenIndex
						if buffer[position] != rune('d') {
							goto l488
						}
						position++
						goto l487
					l488:
						position, tokenIndex = position487, tokenIndex487
						if buffer[position] != rune('D') {
							goto l480
						}
						position++
					}
				l487:
					if buffer[position] != rune(':') {
						goto l480
					}
					position++
					goto l478
				l480:
					position, tokenIndex = position480, tokenIndex480
				}
				{
					position489, tokenIndex489 := position, tokenIndex
					if !_rules[ruleAuthorEtAl]() {
						goto l490
					}
					goto l489
				l490:
					position, tokenIndex = position489, tokenIndex489
					if !_rules[ruleAuthorWord2]() {
						goto l491
					}
					goto l489
				l491:
					position, tokenIndex = position489, tokenIndex489
					if !_rules[ruleAuthorWord3]() {
						goto l492
					}
					goto l489
				l492:
					position, tokenIndex = position489, tokenIndex489
					if !_rules[ruleAuthorPrefix]() {
						goto l478
					}
				}
			l489:
				add(ruleAuthorWord, position479)
			}
			return true
		l478:
			position, tokenIndex = position478, tokenIndex478
			return false
		},
		/* 68 AuthorEtAl <- <(('a' 'r' 'g' '.') / ('e' 't' ' ' 'a' 'l' '.' '{' '?' '}') / ((('e' 't') / '&') (' ' 'a' 'l') '.'?))> */
		func() bool {
			position493, tokenIndex493 := position, tokenIndex
			{
				position494 := position
				{
					position495, tokenIndex495 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l496
					}
					position++
					if buffer[position] != rune('r') {
						goto l496
					}
					position++
					if buffer[position] != rune('g') {
						goto l496
					}
					position++
					if buffer[position] != rune('.') {
						goto l496
					}
					position++
					goto l495
				l496:
					position, tokenIndex = position495, tokenIndex495
					if buffer[position] != rune('e') {
						goto l497
					}
					position++
					if buffer[position] != rune('t') {
						goto l497
					}
					position++
					if buffer[position] != rune(' ') {
						goto l497
					}
					position++
					if buffer[position] != rune('a') {
						goto l497
					}
					position++
					if buffer[position] != rune('l') {
						goto l497
					}
					position++
					if buffer[position] != rune('.') {
						goto l497
					}
					position++
					if buffer[position] != rune('{') {
						goto l497
					}
					position++
					if buffer[position] != rune('?') {
						goto l497
					}
					position++
					if buffer[position] != rune('}') {
						goto l497
					}
					position++
					goto l495
				l497:
					position, tokenIndex = position495, tokenIndex495
					{
						position498, tokenIndex498 := position, tokenIndex
						if buffer[position] != rune('e') {
							goto l499
						}
						position++
						if buffer[position] != rune('t') {
							goto l499
						}
						position++
						goto l498
					l499:
						position, tokenIndex = position498, tokenIndex498
						if buffer[position] != rune('&') {
							goto l493
						}
						position++
					}
				l498:
					if buffer[position] != rune(' ') {
						goto l493
					}
					position++
					if buffer[position] != rune('a') {
						goto l493
					}
					position++
					if buffer[position] != rune('l') {
						goto l493
					}
					position++
					{
						position500, tokenIndex500 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l500
						}
						position++
						goto l501
					l500:
						position, tokenIndex = position500, tokenIndex500
					}
				l501:
				}
			l495:
				add(ruleAuthorEtAl, position494)
			}
			return true
		l493:
			position, tokenIndex = position493, tokenIndex493
			return false
		},
		/* 69 AuthorWord2 <- <(AuthorWord3 dash AuthorWordSoft)> */
		func() bool {
			position502, tokenIndex502 := position, tokenIndex
			{
				position503 := position
				if !_rules[ruleAuthorWord3]() {
					goto l502
				}
				if !_rules[ruledash]() {
					goto l502
				}
				if !_rules[ruleAuthorWordSoft]() {
					goto l502
				}
				add(ruleAuthorWord2, position503)
			}
			return true
		l502:
			position, tokenIndex = position502, tokenIndex502
			return false
		},
		/* 70 AuthorWord3 <- <(AuthorPrefixGlued? (AllCapsAuthorWord / CapAuthorWord) '.'?)> */
		func() bool {
			position504, tokenIndex504 := position, tokenIndex
			{
				position505 := position
				{
					position506, tokenIndex506 := position, tokenIndex
					if !_rules[ruleAuthorPrefixGlued]() {
						goto l506
					}
					goto l507
				l506:
					position, tokenIndex = position506, tokenIndex506
				}
			l507:
				{
					position508, tokenIndex508 := position, tokenIndex
					if !_rules[ruleAllCapsAuthorWord]() {
						goto l509
					}
					goto l508
				l509:
					position, tokenIndex = position508, tokenIndex508
					if !_rules[ruleCapAuthorWord]() {
						goto l504
					}
				}
			l508:
				{
					position510, tokenIndex510 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l510
					}
					position++
					goto l511
				l510:
					position, tokenIndex = position510, tokenIndex510
				}
			l511:
				add(ruleAuthorWord3, position505)
			}
			return true
		l504:
			position, tokenIndex = position504, tokenIndex504
			return false
		},
		/* 71 AuthorWordSoft <- <(((AuthorUpperChar (AuthorUpperChar+ / AuthorLowerChar+)) / AuthorLowerChar+) '.'?)> */
		func() bool {
			position512, tokenIndex512 := position, tokenIndex
			{
				position513 := position
				{
					position514, tokenIndex514 := position, tokenIndex
					if !_rules[ruleAuthorUpperChar]() {
						goto l515
					}
					{
						position516, tokenIndex516 := position, tokenIndex
						if !_rules[ruleAuthorUpperChar]() {
							goto l517
						}
					l518:
						{
							position519, tokenIndex519 := position, tokenIndex
							if !_rules[ruleAuthorUpperChar]() {
								goto l519
							}
							goto l518
						l519:
							position, tokenIndex = position519, tokenIndex519
						}
						goto l516
					l517:
						position, tokenIndex = position516, tokenIndex516
						if !_rules[ruleAuthorLowerChar]() {
							goto l515
						}
					l520:
						{
							position521, tokenIndex521 := position, tokenIndex
							if !_rules[ruleAuthorLowerChar]() {
								goto l521
							}
							goto l520
						l521:
							position, tokenIndex = position521, tokenIndex521
						}
					}
				l516:
					goto l514
				l515:
					position, tokenIndex = position514, tokenIndex514
					if !_rules[ruleAuthorLowerChar]() {
						goto l512
					}
				l522:
					{
						position523, tokenIndex523 := position, tokenIndex
						if !_rules[ruleAuthorLowerChar]() {
							goto l523
						}
						goto l522
					l523:
						position, tokenIndex = position523, tokenIndex523
					}
				}
			l514:
				{
					position524, tokenIndex524 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l524
					}
					position++
					goto l525
				l524:
					position, tokenIndex = position524, tokenIndex524
				}
			l525:
				add(ruleAuthorWordSoft, position513)
			}
			return true
		l512:
			position, tokenIndex = position512, tokenIndex512
			return false
		},
		/* 72 CapAuthorWord <- <(AuthorUpperChar AuthorLowerChar*)> */
		func() bool {
			position526, tokenIndex526 := position, tokenIndex
			{
				position527 := position
				if !_rules[ruleAuthorUpperChar]() {
					goto l526
				}
			l528:
				{
					position529, tokenIndex529 := position, tokenIndex
					if !_rules[ruleAuthorLowerChar]() {
						goto l529
					}
					goto l528
				l529:
					position, tokenIndex = position529, tokenIndex529
				}
				add(ruleCapAuthorWord, position527)
			}
			return true
		l526:
			position, tokenIndex = position526, tokenIndex526
			return false
		},
		/* 73 AllCapsAuthorWord <- <(AuthorUpperChar AuthorUpperChar+)> */
		func() bool {
			position530, tokenIndex530 := position, tokenIndex
			{
				position531 := position
				if !_rules[ruleAuthorUpperChar]() {
					goto l530
				}
				if !_rules[ruleAuthorUpperChar]() {
					goto l530
				}
			l532:
				{
					position533, tokenIndex533 := position, tokenIndex
					if !_rules[ruleAuthorUpperChar]() {
						goto l533
					}
					goto l532
				l533:
					position, tokenIndex = position533, tokenIndex533
				}
				add(ruleAllCapsAuthorWord, position531)
			}
			return true
		l530:
			position, tokenIndex = position530, tokenIndex530
			return false
		},
		/* 74 Filius <- <(('f' '.') / ('f' 'i' 'l' '.') / ('f' 'i' 'l' 'i' 'u' 's'))> */
		func() bool {
			position534, tokenIndex534 := position, tokenIndex
			{
				position535 := position
				{
					position536, tokenIndex536 := position, tokenIndex
					if buffer[position] != rune('f') {
						goto l537
					}
					position++
					if buffer[position] != rune('.') {
						goto l537
					}
					position++
					goto l536
				l537:
					position, tokenIndex = position536, tokenIndex536
					if buffer[position] != rune('f') {
						goto l538
					}
					position++
					if buffer[position] != rune('i') {
						goto l538
					}
					position++
					if buffer[position] != rune('l') {
						goto l538
					}
					position++
					if buffer[position] != rune('.') {
						goto l538
					}
					position++
					goto l536
				l538:
					position, tokenIndex = position536, tokenIndex536
					if buffer[position] != rune('f') {
						goto l534
					}
					position++
					if buffer[position] != rune('i') {
						goto l534
					}
					position++
					if buffer[position] != rune('l') {
						goto l534
					}
					position++
					if buffer[position] != rune('i') {
						goto l534
					}
					position++
					if buffer[position] != rune('u') {
						goto l534
					}
					position++
					if buffer[position] != rune('s') {
						goto l534
					}
					position++
				}
			l536:
				add(ruleFilius, position535)
			}
			return true
		l534:
			position, tokenIndex = position534, tokenIndex534
			return false
		},
		/* 75 AuthorPrefixGlued <- <(('d' '\'') / ('O' '\'') / ('L' '\''))> */
		func() bool {
			position539, tokenIndex539 := position, tokenIndex
			{
				position540 := position
				{
					position541, tokenIndex541 := position, tokenIndex
					if buffer[position] != rune('d') {
						goto l542
					}
					position++
					if buffer[position] != rune('\'') {
						goto l542
					}
					position++
					goto l541
				l542:
					position, tokenIndex = position541, tokenIndex541
					if buffer[position] != rune('O') {
						goto l543
					}
					position++
					if buffer[position] != rune('\'') {
						goto l543
					}
					position++
					goto l541
				l543:
					position, tokenIndex = position541, tokenIndex541
					if buffer[position] != rune('L') {
						goto l539
					}
					position++
					if buffer[position] != rune('\'') {
						goto l539
					}
					position++
				}
			l541:
				add(ruleAuthorPrefixGlued, position540)
			}
			return true
		l539:
			position, tokenIndex = position539, tokenIndex539
			return false
		},
		/* 76 AuthorPrefix <- <(AuthorPrefix1 / AuthorPrefix2)> */
		func() bool {
			position544, tokenIndex544 := position, tokenIndex
			{
				position545 := position
				{
					position546, tokenIndex546 := position, tokenIndex
					if !_rules[ruleAuthorPrefix1]() {
						goto l547
					}
					goto l546
				l547:
					position, tokenIndex = position546, tokenIndex546
					if !_rules[ruleAuthorPrefix2]() {
						goto l544
					}
				}
			l546:
				add(ruleAuthorPrefix, position545)
			}
			return true
		l544:
			position, tokenIndex = position544, tokenIndex544
			return false
		},
		/* 77 AuthorPrefix2 <- <(('v' '.' (_? ('d' '.'))?) / ('\'' 't'))> */
		func() bool {
			position548, tokenIndex548 := position, tokenIndex
			{
				position549 := position
				{
					position550, tokenIndex550 := position, tokenIndex
					if buffer[position] != rune('v') {
						goto l551
					}
					position++
					if buffer[position] != rune('.') {
						goto l551
					}
					position++
					{
						position552, tokenIndex552 := position, tokenIndex
						{
							position554, tokenIndex554 := position, tokenIndex
							if !_rules[rule_]() {
								goto l554
							}
							goto l555
						l554:
							position, tokenIndex = position554, tokenIndex554
						}
					l555:
						if buffer[position] != rune('d') {
							goto l552
						}
						position++
						if buffer[position] != rune('.') {
							goto l552
						}
						position++
						goto l553
					l552:
						position, tokenIndex = position552, tokenIndex552
					}
				l553:
					goto l550
				l551:
					position, tokenIndex = position550, tokenIndex550
					if buffer[position] != rune('\'') {
						goto l548
					}
					position++
					if buffer[position] != rune('t') {
						goto l548
					}
					position++
				}
			l550:
				add(ruleAuthorPrefix2, position549)
			}
			return true
		l548:
			position, tokenIndex = position548, tokenIndex548
			return false
		},
		/* 78 AuthorPrefix1 <- <((('a' 'b') / ('a' 'f') / ('b' 'i' 's') / ('d' 'a') / ('d' 'e' 'r') / ('d' 'e' 's') / ('d' 'e' 'n') / ('d' 'e' 'l') / ('d' 'e' 'l' 'l' 'a') / ('d' 'e' 'l' 'a') / ('d' 'e') / ('d' 'i') / ('d' 'u') / ('e' 'l') / ('l' 'a') / ('l' 'e') / ('t' 'e' 'r') / ('v' 'a' 'n') / ('d' '\'') / ('i' 'n' '\'' 't') / ('z' 'u' 'r') / ('v' 'o' 'n' (_ (('d' '.') / ('d' 'e' 'm')))?) / ('v' (_ 'd')?)) &_)> */
		func() bool {
			position556, tokenIndex556 := position, tokenIndex
			{
				position557 := position
				{
					position558, tokenIndex558 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l559
					}
					position++
					if buffer[position] != rune('b') {
						goto l559
					}
					position++
					goto l558
				l559:
					position, tokenIndex = position558, tokenIndex558
					if buffer[position] != rune('a') {
						goto l560
					}
					position++
					if buffer[position] != rune('f') {
						goto l560
					}
					position++
					goto l558
				l560:
					position, tokenIndex = position558, tokenIndex558
					if buffer[position] != rune('b') {
						goto l561
					}
					position++
					if buffer[position] != rune('i') {
						goto l561
					}
					position++
					if buffer[position] != rune('s') {
						goto l561
					}
					position++
					goto l558
				l561:
					position, tokenIndex = position558, tokenIndex558
					if buffer[position] != rune('d') {
						goto l562
					}
					position++
					if buffer[position] != rune('a') {
						goto l562
					}
					position++
					goto l558
				l562:
					position, tokenIndex = position558, tokenIndex558
					if buffer[position] != rune('d') {
						goto l563
					}
					position++
					if buffer[position] != rune('e') {
						goto l563
					}
					position++
					if buffer[position] != rune('r') {
						goto l563
					}
					position++
					goto l558
				l563:
					position, tokenIndex = position558, tokenIndex558
					if buffer[position] != rune('d') {
						goto l564
					}
					position++
					if buffer[position] != rune('e') {
						goto l564
					}
					position++
					if buffer[position] != rune('s') {
						goto l564
					}
					position++
					goto l558
				l564:
					position, tokenIndex = position558, tokenIndex558
					if buffer[position] != rune('d') {
						goto l565
					}
					position++
					if buffer[position] != rune('e') {
						goto l565
					}
					position++
					if buffer[position] != rune('n') {
						goto l565
					}
					position++
					goto l558
				l565:
					position, tokenIndex = position558, tokenIndex558
					if buffer[position] != rune('d') {
						goto l566
					}
					position++
					if buffer[position] != rune('e') {
						goto l566
					}
					position++
					if buffer[position] != rune('l') {
						goto l566
					}
					position++
					goto l558
				l566:
					position, tokenIndex = position558, tokenIndex558
					if buffer[position] != rune('d') {
						goto l567
					}
					position++
					if buffer[position] != rune('e') {
						goto l567
					}
					position++
					if buffer[position] != rune('l') {
						goto l567
					}
					position++
					if buffer[position] != rune('l') {
						goto l567
					}
					position++
					if buffer[position] != rune('a') {
						goto l567
					}
					position++
					goto l558
				l567:
					position, tokenIndex = position558, tokenIndex558
					if buffer[position] != rune('d') {
						goto l568
					}
					position++
					if buffer[position] != rune('e') {
						goto l568
					}
					position++
					if buffer[position] != rune('l') {
						goto l568
					}
					position++
					if buffer[position] != rune('a') {
						goto l568
					}
					position++
					goto l558
				l568:
					position, tokenIndex = position558, tokenIndex558
					if buffer[position] != rune('d') {
						goto l569
					}
					position++
					if buffer[position] != rune('e') {
						goto l569
					}
					position++
					goto l558
				l569:
					position, tokenIndex = position558, tokenIndex558
					if buffer[position] != rune('d') {
						goto l570
					}
					position++
					if buffer[position] != rune('i') {
						goto l570
					}
					position++
					goto l558
				l570:
					position, tokenIndex = position558, tokenIndex558
					if buffer[position] != rune('d') {
						goto l571
					}
					position++
					if buffer[position] != rune('u') {
						goto l571
					}
					position++
					goto l558
				l571:
					position, tokenIndex = position558, tokenIndex558
					if buffer[position] != rune('e') {
						goto l572
					}
					position++
					if buffer[position] != rune('l') {
						goto l572
					}
					position++
					goto l558
				l572:
					position, tokenIndex = position558, tokenIndex558
					if buffer[position] != rune('l') {
						goto l573
					}
					position++
					if buffer[position] != rune('a') {
						goto l573
					}
					position++
					goto l558
				l573:
					position, tokenIndex = position558, tokenIndex558
					if buffer[position] != rune('l') {
						goto l574
					}
					position++
					if buffer[position] != rune('e') {
						goto l574
					}
					position++
					goto l558
				l574:
					position, tokenIndex = position558, tokenIndex558
					if buffer[position] != rune('t') {
						goto l575
					}
					position++
					if buffer[position] != rune('e') {
						goto l575
					}
					position++
					if buffer[position] != rune('r') {
						goto l575
					}
					position++
					goto l558
				l575:
					position, tokenIndex = position558, tokenIndex558
					if buffer[position] != rune('v') {
						goto l576
					}
					position++
					if buffer[position] != rune('a') {
						goto l576
					}
					position++
					if buffer[position] != rune('n') {
						goto l576
					}
					position++
					goto l558
				l576:
					position, tokenIndex = position558, tokenIndex558
					if buffer[position] != rune('d') {
						goto l577
					}
					position++
					if buffer[position] != rune('\'') {
						goto l577
					}
					position++
					goto l558
				l577:
					position, tokenIndex = position558, tokenIndex558
					if buffer[position] != rune('i') {
						goto l578
					}
					position++
					if buffer[position] != rune('n') {
						goto l578
					}
					position++
					if buffer[position] != rune('\'') {
						goto l578
					}
					position++
					if buffer[position] != rune('t') {
						goto l578
					}
					position++
					goto l558
				l578:
					position, tokenIndex = position558, tokenIndex558
					if buffer[position] != rune('z') {
						goto l579
					}
					position++
					if buffer[position] != rune('u') {
						goto l579
					}
					position++
					if buffer[position] != rune('r') {
						goto l579
					}
					position++
					goto l558
				l579:
					position, tokenIndex = position558, tokenIndex558
					if buffer[position] != rune('v') {
						goto l580
					}
					position++
					if buffer[position] != rune('o') {
						goto l580
					}
					position++
					if buffer[position] != rune('n') {
						goto l580
					}
					position++
					{
						position581, tokenIndex581 := position, tokenIndex
						if !_rules[rule_]() {
							goto l581
						}
						{
							position583, tokenIndex583 := position, tokenIndex
							if buffer[position] != rune('d') {
								goto l584
							}
							position++
							if buffer[position] != rune('.') {
								goto l584
							}
							position++
							goto l583
						l584:
							position, tokenIndex = position583, tokenIndex583
							if buffer[position] != rune('d') {
								goto l581
							}
							position++
							if buffer[position] != rune('e') {
								goto l581
							}
							position++
							if buffer[position] != rune('m') {
								goto l581
							}
							position++
						}
					l583:
						goto l582
					l581:
						position, tokenIndex = position581, tokenIndex581
					}
				l582:
					goto l558
				l580:
					position, tokenIndex = position558, tokenIndex558
					if buffer[position] != rune('v') {
						goto l556
					}
					position++
					{
						position585, tokenIndex585 := position, tokenIndex
						if !_rules[rule_]() {
							goto l585
						}
						if buffer[position] != rune('d') {
							goto l585
						}
						position++
						goto l586
					l585:
						position, tokenIndex = position585, tokenIndex585
					}
				l586:
				}
			l558:
				{
					position587, tokenIndex587 := position, tokenIndex
					if !_rules[rule_]() {
						goto l556
					}
					position, tokenIndex = position587, tokenIndex587
				}
				add(ruleAuthorPrefix1, position557)
			}
			return true
		l556:
			position, tokenIndex = position556, tokenIndex556
			return false
		},
		/* 79 AuthorUpperChar <- <(hASCII / MiscodedChar / ('À' / 'Á' / 'Â' / 'Ã' / 'Ä' / 'Å' / 'Æ' / 'Ç' / 'È' / 'É' / 'Ê' / 'Ë' / 'Ì' / 'Í' / 'Î' / 'Ï' / 'Ð' / 'Ñ' / 'Ò' / 'Ó' / 'Ô' / 'Õ' / 'Ö' / 'Ø' / 'Ù' / 'Ú' / 'Û' / 'Ü' / 'Ý' / 'Ć' / 'Č' / 'Ď' / 'İ' / 'Ķ' / 'Ĺ' / 'ĺ' / 'Ľ' / 'ľ' / 'Ł' / 'ł' / 'Ņ' / 'Ō' / 'Ő' / 'Œ' / 'Ř' / 'Ś' / 'Ŝ' / 'Ş' / 'Š' / 'Ÿ' / 'Ź' / 'Ż' / 'Ž' / 'ƒ' / 'Ǿ' / 'Ș' / 'Ț'))> */
		func() bool {
			position588, tokenIndex588 := position, tokenIndex
			{
				position589 := position
				{
					position590, tokenIndex590 := position, tokenIndex
					if !_rules[rulehASCII]() {
						goto l591
					}
					goto l590
				l591:
					position, tokenIndex = position590, tokenIndex590
					if !_rules[ruleMiscodedChar]() {
						goto l592
					}
					goto l590
				l592:
					position, tokenIndex = position590, tokenIndex590
					{
						position593, tokenIndex593 := position, tokenIndex
						if buffer[position] != rune('À') {
							goto l594
						}
						position++
						goto l593
					l594:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Á') {
							goto l595
						}
						position++
						goto l593
					l595:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Â') {
							goto l596
						}
						position++
						goto l593
					l596:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ã') {
							goto l597
						}
						position++
						goto l593
					l597:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ä') {
							goto l598
						}
						position++
						goto l593
					l598:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Å') {
							goto l599
						}
						position++
						goto l593
					l599:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Æ') {
							goto l600
						}
						position++
						goto l593
					l600:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ç') {
							goto l601
						}
						position++
						goto l593
					l601:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('È') {
							goto l602
						}
						position++
						goto l593
					l602:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('É') {
							goto l603
						}
						position++
						goto l593
					l603:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ê') {
							goto l604
						}
						position++
						goto l593
					l604:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ë') {
							goto l605
						}
						position++
						goto l593
					l605:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ì') {
							goto l606
						}
						position++
						goto l593
					l606:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Í') {
							goto l607
						}
						position++
						goto l593
					l607:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Î') {
							goto l608
						}
						position++
						goto l593
					l608:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ï') {
							goto l609
						}
						position++
						goto l593
					l609:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ð') {
							goto l610
						}
						position++
						goto l593
					l610:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ñ') {
							goto l611
						}
						position++
						goto l593
					l611:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ò') {
							goto l612
						}
						position++
						goto l593
					l612:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ó') {
							goto l613
						}
						position++
						goto l593
					l613:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ô') {
							goto l614
						}
						position++
						goto l593
					l614:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Õ') {
							goto l615
						}
						position++
						goto l593
					l615:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ö') {
							goto l616
						}
						position++
						goto l593
					l616:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ø') {
							goto l617
						}
						position++
						goto l593
					l617:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ù') {
							goto l618
						}
						position++
						goto l593
					l618:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ú') {
							goto l619
						}
						position++
						goto l593
					l619:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Û') {
							goto l620
						}
						position++
						goto l593
					l620:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ü') {
							goto l621
						}
						position++
						goto l593
					l621:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ý') {
							goto l622
						}
						position++
						goto l593
					l622:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ć') {
							goto l623
						}
						position++
						goto l593
					l623:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Č') {
							goto l624
						}
						position++
						goto l593
					l624:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ď') {
							goto l625
						}
						position++
						goto l593
					l625:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('İ') {
							goto l626
						}
						position++
						goto l593
					l626:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ķ') {
							goto l627
						}
						position++
						goto l593
					l627:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ĺ') {
							goto l628
						}
						position++
						goto l593
					l628:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ĺ') {
							goto l629
						}
						position++
						goto l593
					l629:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ľ') {
							goto l630
						}
						position++
						goto l593
					l630:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ľ') {
							goto l631
						}
						position++
						goto l593
					l631:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ł') {
							goto l632
						}
						position++
						goto l593
					l632:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ł') {
							goto l633
						}
						position++
						goto l593
					l633:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ņ') {
							goto l634
						}
						position++
						goto l593
					l634:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ō') {
							goto l635
						}
						position++
						goto l593
					l635:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ő') {
							goto l636
						}
						position++
						goto l593
					l636:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Œ') {
							goto l637
						}
						position++
						goto l593
					l637:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ř') {
							goto l638
						}
						position++
						goto l593
					l638:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ś') {
							goto l639
						}
						position++
						goto l593
					l639:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ŝ') {
							goto l640
						}
						position++
						goto l593
					l640:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ş') {
							goto l641
						}
						position++
						goto l593
					l641:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Š') {
							goto l642
						}
						position++
						goto l593
					l642:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ÿ') {
							goto l643
						}
						position++
						goto l593
					l643:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ź') {
							goto l644
						}
						position++
						goto l593
					l644:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ż') {
							goto l645
						}
						position++
						goto l593
					l645:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ž') {
							goto l646
						}
						position++
						goto l593
					l646:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ƒ') {
							goto l647
						}
						position++
						goto l593
					l647:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ǿ') {
							goto l648
						}
						position++
						goto l593
					l648:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ș') {
							goto l649
						}
						position++
						goto l593
					l649:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('Ț') {
							goto l588
						}
						position++
					}
				l593:
				}
			l590:
				add(ruleAuthorUpperChar, position589)
			}
			return true
		l588:
			position, tokenIndex = position588, tokenIndex588
			return false
		},
		/* 80 AuthorLowerChar <- <(lASCII / MiscodedChar / ('à' / 'á' / 'â' / 'ã' / 'ä' / 'å' / 'æ' / 'ç' / 'è' / 'é' / 'ê' / 'ë' / 'ì' / 'í' / 'î' / 'ï' / 'ð' / 'ñ' / 'ò' / 'ó' / 'ó' / 'ô' / 'õ' / 'ö' / 'ø' / 'ù' / 'ú' / 'û' / 'ü' / 'ý' / 'ÿ' / 'ā' / 'ă' / 'ą' / 'ć' / 'ĉ' / 'č' / 'ď' / 'đ' / '\'' / 'ē' / 'ĕ' / 'ė' / 'ę' / 'ě' / 'ğ' / 'ī' / 'ĭ' / 'İ' / 'ı' / 'ĺ' / 'ľ' / 'ł' / 'ń' / 'ņ' / 'ň' / 'ŏ' / 'ő' / 'œ' / 'ŕ' / 'ř' / 'ś' / 'ş' / 'š' / 'ţ' / 'ť' / 'ũ' / 'ū' / 'ŭ' / 'ů' / 'ű' / 'ź' / 'ż' / 'ž' / 'ſ' / 'ǎ' / 'ǔ' / 'ǧ' / 'ș' / 'ț' / 'ȳ' / 'ß'))> */
		func() bool {
			position650, tokenIndex650 := position, tokenIndex
			{
				position651 := position
				{
					position652, tokenIndex652 := position, tokenIndex
					if !_rules[rulelASCII]() {
						goto l653
					}
					goto l652
				l653:
					position, tokenIndex = position652, tokenIndex652
					if !_rules[ruleMiscodedChar]() {
						goto l654
					}
					goto l652
				l654:
					position, tokenIndex = position652, tokenIndex652
					{
						position655, tokenIndex655 := position, tokenIndex
						if buffer[position] != rune('à') {
							goto l656
						}
						position++
						goto l655
					l656:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('á') {
							goto l657
						}
						position++
						goto l655
					l657:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('â') {
							goto l658
						}
						position++
						goto l655
					l658:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ã') {
							goto l659
						}
						position++
						goto l655
					l659:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ä') {
							goto l660
						}
						position++
						goto l655
					l660:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('å') {
							goto l661
						}
						position++
						goto l655
					l661:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('æ') {
							goto l662
						}
						position++
						goto l655
					l662:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ç') {
							goto l663
						}
						position++
						goto l655
					l663:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('è') {
							goto l664
						}
						position++
						goto l655
					l664:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('é') {
							goto l665
						}
						position++
						goto l655
					l665:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ê') {
							goto l666
						}
						position++
						goto l655
					l666:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ë') {
							goto l667
						}
						position++
						goto l655
					l667:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ì') {
							goto l668
						}
						position++
						goto l655
					l668:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('í') {
							goto l669
						}
						position++
						goto l655
					l669:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('î') {
							goto l670
						}
						position++
						goto l655
					l670:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ï') {
							goto l671
						}
						position++
						goto l655
					l671:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ð') {
							goto l672
						}
						position++
						goto l655
					l672:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ñ') {
							goto l673
						}
						position++
						goto l655
					l673:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ò') {
							goto l674
						}
						position++
						goto l655
					l674:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ó') {
							goto l675
						}
						position++
						goto l655
					l675:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ó') {
							goto l676
						}
						position++
						goto l655
					l676:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ô') {
							goto l677
						}
						position++
						goto l655
					l677:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('õ') {
							goto l678
						}
						position++
						goto l655
					l678:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ö') {
							goto l679
						}
						position++
						goto l655
					l679:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ø') {
							goto l680
						}
						position++
						goto l655
					l680:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ù') {
							goto l681
						}
						position++
						goto l655
					l681:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ú') {
							goto l682
						}
						position++
						goto l655
					l682:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('û') {
							goto l683
						}
						position++
						goto l655
					l683:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ü') {
							goto l684
						}
						position++
						goto l655
					l684:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ý') {
							goto l685
						}
						position++
						goto l655
					l685:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ÿ') {
							goto l686
						}
						position++
						goto l655
					l686:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ā') {
							goto l687
						}
						position++
						goto l655
					l687:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ă') {
							goto l688
						}
						position++
						goto l655
					l688:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ą') {
							goto l689
						}
						position++
						goto l655
					l689:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ć') {
							goto l690
						}
						position++
						goto l655
					l690:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ĉ') {
							goto l691
						}
						position++
						goto l655
					l691:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('č') {
							goto l692
						}
						position++
						goto l655
					l692:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ď') {
							goto l693
						}
						position++
						goto l655
					l693:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('đ') {
							goto l694
						}
						position++
						goto l655
					l694:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('\'') {
							goto l695
						}
						position++
						goto l655
					l695:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ē') {
							goto l696
						}
						position++
						goto l655
					l696:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ĕ') {
							goto l697
						}
						position++
						goto l655
					l697:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ė') {
							goto l698
						}
						position++
						goto l655
					l698:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ę') {
							goto l699
						}
						position++
						goto l655
					l699:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ě') {
							goto l700
						}
						position++
						goto l655
					l700:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ğ') {
							goto l701
						}
						position++
						goto l655
					l701:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ī') {
							goto l702
						}
						position++
						goto l655
					l702:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ĭ') {
							goto l703
						}
						position++
						goto l655
					l703:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('İ') {
							goto l704
						}
						position++
						goto l655
					l704:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ı') {
							goto l705
						}
						position++
						goto l655
					l705:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ĺ') {
							goto l706
						}
						position++
						goto l655
					l706:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ľ') {
							goto l707
						}
						position++
						goto l655
					l707:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ł') {
							goto l708
						}
						position++
						goto l655
					l708:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ń') {
							goto l709
						}
						position++
						goto l655
					l709:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ņ') {
							goto l710
						}
						position++
						goto l655
					l710:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ň') {
							goto l711
						}
						position++
						goto l655
					l711:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ŏ') {
							goto l712
						}
						position++
						goto l655
					l712:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ő') {
							goto l713
						}
						position++
						goto l655
					l713:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('œ') {
							goto l714
						}
						position++
						goto l655
					l714:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ŕ') {
							goto l715
						}
						position++
						goto l655
					l715:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ř') {
							goto l716
						}
						position++
						goto l655
					l716:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ś') {
							goto l717
						}
						position++
						goto l655
					l717:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ş') {
							goto l718
						}
						position++
						goto l655
					l718:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('š') {
							goto l719
						}
						position++
						goto l655
					l719:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ţ') {
							goto l720
						}
						position++
						goto l655
					l720:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ť') {
							goto l721
						}
						position++
						goto l655
					l721:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ũ') {
							goto l722
						}
						position++
						goto l655
					l722:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ū') {
							goto l723
						}
						position++
						goto l655
					l723:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ŭ') {
							goto l724
						}
						position++
						goto l655
					l724:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ů') {
							goto l725
						}
						position++
						goto l655
					l725:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ű') {
							goto l726
						}
						position++
						goto l655
					l726:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ź') {
							goto l727
						}
						position++
						goto l655
					l727:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ż') {
							goto l728
						}
						position++
						goto l655
					l728:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ž') {
							goto l729
						}
						position++
						goto l655
					l729:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ſ') {
							goto l730
						}
						position++
						goto l655
					l730:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ǎ') {
							goto l731
						}
						position++
						goto l655
					l731:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ǔ') {
							goto l732
						}
						position++
						goto l655
					l732:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ǧ') {
							goto l733
						}
						position++
						goto l655
					l733:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ș') {
							goto l734
						}
						position++
						goto l655
					l734:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ț') {
							goto l735
						}
						position++
						goto l655
					l735:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ȳ') {
							goto l736
						}
						position++
						goto l655
					l736:
						position, tokenIndex = position655, tokenIndex655
						if buffer[position] != rune('ß') {
							goto l650
						}
						position++
					}
				l655:
				}
			l652:
				add(ruleAuthorLowerChar, position651)
			}
			return true
		l650:
			position, tokenIndex = position650, tokenIndex650
			return false
		},
		/* 81 Year <- <(YearRange / YearApprox / YearWithParens / YearWithPage / YearWithDot / YearWithChar / YearNum)> */
		func() bool {
			position737, tokenIndex737 := position, tokenIndex
			{
				position738 := position
				{
					position739, tokenIndex739 := position, tokenIndex
					if !_rules[ruleYearRange]() {
						goto l740
					}
					goto l739
				l740:
					position, tokenIndex = position739, tokenIndex739
					if !_rules[ruleYearApprox]() {
						goto l741
					}
					goto l739
				l741:
					position, tokenIndex = position739, tokenIndex739
					if !_rules[ruleYearWithParens]() {
						goto l742
					}
					goto l739
				l742:
					position, tokenIndex = position739, tokenIndex739
					if !_rules[ruleYearWithPage]() {
						goto l743
					}
					goto l739
				l743:
					position, tokenIndex = position739, tokenIndex739
					if !_rules[ruleYearWithDot]() {
						goto l744
					}
					goto l739
				l744:
					position, tokenIndex = position739, tokenIndex739
					if !_rules[ruleYearWithChar]() {
						goto l745
					}
					goto l739
				l745:
					position, tokenIndex = position739, tokenIndex739
					if !_rules[ruleYearNum]() {
						goto l737
					}
				}
			l739:
				add(ruleYear, position738)
			}
			return true
		l737:
			position, tokenIndex = position737, tokenIndex737
			return false
		},
		/* 82 YearRange <- <(YearNum dash (nums+ ('a' / 'b' / 'c' / 'd' / 'e' / 'f' / 'g' / 'h' / 'i' / 'j' / 'k' / 'l' / 'm' / 'n' / 'o' / 'p' / 'q' / 'r' / 's' / 't' / 'u' / 'v' / 'w' / 'x' / 'y' / 'z' / '?')*))> */
		func() bool {
			position746, tokenIndex746 := position, tokenIndex
			{
				position747 := position
				if !_rules[ruleYearNum]() {
					goto l746
				}
				if !_rules[ruledash]() {
					goto l746
				}
				if !_rules[rulenums]() {
					goto l746
				}
			l748:
				{
					position749, tokenIndex749 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l749
					}
					goto l748
				l749:
					position, tokenIndex = position749, tokenIndex749
				}
			l750:
				{
					position751, tokenIndex751 := position, tokenIndex
					{
						position752, tokenIndex752 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l753
						}
						position++
						goto l752
					l753:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('b') {
							goto l754
						}
						position++
						goto l752
					l754:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('c') {
							goto l755
						}
						position++
						goto l752
					l755:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('d') {
							goto l756
						}
						position++
						goto l752
					l756:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('e') {
							goto l757
						}
						position++
						goto l752
					l757:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('f') {
							goto l758
						}
						position++
						goto l752
					l758:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('g') {
							goto l759
						}
						position++
						goto l752
					l759:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('h') {
							goto l760
						}
						position++
						goto l752
					l760:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('i') {
							goto l761
						}
						position++
						goto l752
					l761:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('j') {
							goto l762
						}
						position++
						goto l752
					l762:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('k') {
							goto l763
						}
						position++
						goto l752
					l763:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('l') {
							goto l764
						}
						position++
						goto l752
					l764:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('m') {
							goto l765
						}
						position++
						goto l752
					l765:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('n') {
							goto l766
						}
						position++
						goto l752
					l766:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('o') {
							goto l767
						}
						position++
						goto l752
					l767:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('p') {
							goto l768
						}
						position++
						goto l752
					l768:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('q') {
							goto l769
						}
						position++
						goto l752
					l769:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('r') {
							goto l770
						}
						position++
						goto l752
					l770:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('s') {
							goto l771
						}
						position++
						goto l752
					l771:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('t') {
							goto l772
						}
						position++
						goto l752
					l772:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('u') {
							goto l773
						}
						position++
						goto l752
					l773:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('v') {
							goto l774
						}
						position++
						goto l752
					l774:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('w') {
							goto l775
						}
						position++
						goto l752
					l775:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('x') {
							goto l776
						}
						position++
						goto l752
					l776:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('y') {
							goto l777
						}
						position++
						goto l752
					l777:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('z') {
							goto l778
						}
						position++
						goto l752
					l778:
						position, tokenIndex = position752, tokenIndex752
						if buffer[position] != rune('?') {
							goto l751
						}
						position++
					}
				l752:
					goto l750
				l751:
					position, tokenIndex = position751, tokenIndex751
				}
				add(ruleYearRange, position747)
			}
			return true
		l746:
			position, tokenIndex = position746, tokenIndex746
			return false
		},
		/* 83 YearWithDot <- <(YearNum '.')> */
		func() bool {
			position779, tokenIndex779 := position, tokenIndex
			{
				position780 := position
				if !_rules[ruleYearNum]() {
					goto l779
				}
				if buffer[position] != rune('.') {
					goto l779
				}
				position++
				add(ruleYearWithDot, position780)
			}
			return true
		l779:
			position, tokenIndex = position779, tokenIndex779
			return false
		},
		/* 84 YearApprox <- <('[' _? YearNum _? ']')> */
		func() bool {
			position781, tokenIndex781 := position, tokenIndex
			{
				position782 := position
				if buffer[position] != rune('[') {
					goto l781
				}
				position++
				{
					position783, tokenIndex783 := position, tokenIndex
					if !_rules[rule_]() {
						goto l783
					}
					goto l784
				l783:
					position, tokenIndex = position783, tokenIndex783
				}
			l784:
				if !_rules[ruleYearNum]() {
					goto l781
				}
				{
					position785, tokenIndex785 := position, tokenIndex
					if !_rules[rule_]() {
						goto l785
					}
					goto l786
				l785:
					position, tokenIndex = position785, tokenIndex785
				}
			l786:
				if buffer[position] != rune(']') {
					goto l781
				}
				position++
				add(ruleYearApprox, position782)
			}
			return true
		l781:
			position, tokenIndex = position781, tokenIndex781
			return false
		},
		/* 85 YearWithPage <- <((YearWithChar / YearNum) _? ':' _? nums+)> */
		func() bool {
			position787, tokenIndex787 := position, tokenIndex
			{
				position788 := position
				{
					position789, tokenIndex789 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l790
					}
					goto l789
				l790:
					position, tokenIndex = position789, tokenIndex789
					if !_rules[ruleYearNum]() {
						goto l787
					}
				}
			l789:
				{
					position791, tokenIndex791 := position, tokenIndex
					if !_rules[rule_]() {
						goto l791
					}
					goto l792
				l791:
					position, tokenIndex = position791, tokenIndex791
				}
			l792:
				if buffer[position] != rune(':') {
					goto l787
				}
				position++
				{
					position793, tokenIndex793 := position, tokenIndex
					if !_rules[rule_]() {
						goto l793
					}
					goto l794
				l793:
					position, tokenIndex = position793, tokenIndex793
				}
			l794:
				if !_rules[rulenums]() {
					goto l787
				}
			l795:
				{
					position796, tokenIndex796 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l796
					}
					goto l795
				l796:
					position, tokenIndex = position796, tokenIndex796
				}
				add(ruleYearWithPage, position788)
			}
			return true
		l787:
			position, tokenIndex = position787, tokenIndex787
			return false
		},
		/* 86 YearWithParens <- <('(' (YearWithChar / YearNum) ')')> */
		func() bool {
			position797, tokenIndex797 := position, tokenIndex
			{
				position798 := position
				if buffer[position] != rune('(') {
					goto l797
				}
				position++
				{
					position799, tokenIndex799 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l800
					}
					goto l799
				l800:
					position, tokenIndex = position799, tokenIndex799
					if !_rules[ruleYearNum]() {
						goto l797
					}
				}
			l799:
				if buffer[position] != rune(')') {
					goto l797
				}
				position++
				add(ruleYearWithParens, position798)
			}
			return true
		l797:
			position, tokenIndex = position797, tokenIndex797
			return false
		},
		/* 87 YearWithChar <- <(YearNum lASCII Action0)> */
		func() bool {
			position801, tokenIndex801 := position, tokenIndex
			{
				position802 := position
				if !_rules[ruleYearNum]() {
					goto l801
				}
				if !_rules[rulelASCII]() {
					goto l801
				}
				if !_rules[ruleAction0]() {
					goto l801
				}
				add(ruleYearWithChar, position802)
			}
			return true
		l801:
			position, tokenIndex = position801, tokenIndex801
			return false
		},
		/* 88 YearNum <- <(('1' / '2') ('0' / '7' / '8' / '9') nums (nums / '?') '?'*)> */
		func() bool {
			position803, tokenIndex803 := position, tokenIndex
			{
				position804 := position
				{
					position805, tokenIndex805 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l806
					}
					position++
					goto l805
				l806:
					position, tokenIndex = position805, tokenIndex805
					if buffer[position] != rune('2') {
						goto l803
					}
					position++
				}
			l805:
				{
					position807, tokenIndex807 := position, tokenIndex
					if buffer[position] != rune('0') {
						goto l808
					}
					position++
					goto l807
				l808:
					position, tokenIndex = position807, tokenIndex807
					if buffer[position] != rune('7') {
						goto l809
					}
					position++
					goto l807
				l809:
					position, tokenIndex = position807, tokenIndex807
					if buffer[position] != rune('8') {
						goto l810
					}
					position++
					goto l807
				l810:
					position, tokenIndex = position807, tokenIndex807
					if buffer[position] != rune('9') {
						goto l803
					}
					position++
				}
			l807:
				if !_rules[rulenums]() {
					goto l803
				}
				{
					position811, tokenIndex811 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l812
					}
					goto l811
				l812:
					position, tokenIndex = position811, tokenIndex811
					if buffer[position] != rune('?') {
						goto l803
					}
					position++
				}
			l811:
			l813:
				{
					position814, tokenIndex814 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l814
					}
					position++
					goto l813
				l814:
					position, tokenIndex = position814, tokenIndex814
				}
				add(ruleYearNum, position804)
			}
			return true
		l803:
			position, tokenIndex = position803, tokenIndex803
			return false
		},
		/* 89 NameUpperChar <- <(UpperChar / UpperCharExtended)> */
		func() bool {
			position815, tokenIndex815 := position, tokenIndex
			{
				position816 := position
				{
					position817, tokenIndex817 := position, tokenIndex
					if !_rules[ruleUpperChar]() {
						goto l818
					}
					goto l817
				l818:
					position, tokenIndex = position817, tokenIndex817
					if !_rules[ruleUpperCharExtended]() {
						goto l815
					}
				}
			l817:
				add(ruleNameUpperChar, position816)
			}
			return true
		l815:
			position, tokenIndex = position815, tokenIndex815
			return false
		},
		/* 90 UpperCharExtended <- <('Æ' / 'Œ' / 'Ö')> */
		func() bool {
			position819, tokenIndex819 := position, tokenIndex
			{
				position820 := position
				{
					position821, tokenIndex821 := position, tokenIndex
					if buffer[position] != rune('Æ') {
						goto l822
					}
					position++
					goto l821
				l822:
					position, tokenIndex = position821, tokenIndex821
					if buffer[position] != rune('Œ') {
						goto l823
					}
					position++
					goto l821
				l823:
					position, tokenIndex = position821, tokenIndex821
					if buffer[position] != rune('Ö') {
						goto l819
					}
					position++
				}
			l821:
				add(ruleUpperCharExtended, position820)
			}
			return true
		l819:
			position, tokenIndex = position819, tokenIndex819
			return false
		},
		/* 91 UpperChar <- <hASCII> */
		func() bool {
			position824, tokenIndex824 := position, tokenIndex
			{
				position825 := position
				if !_rules[rulehASCII]() {
					goto l824
				}
				add(ruleUpperChar, position825)
			}
			return true
		l824:
			position, tokenIndex = position824, tokenIndex824
			return false
		},
		/* 92 NameLowerChar <- <(LowerChar / LowerCharExtended / MiscodedChar)> */
		func() bool {
			position826, tokenIndex826 := position, tokenIndex
			{
				position827 := position
				{
					position828, tokenIndex828 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l829
					}
					goto l828
				l829:
					position, tokenIndex = position828, tokenIndex828
					if !_rules[ruleLowerCharExtended]() {
						goto l830
					}
					goto l828
				l830:
					position, tokenIndex = position828, tokenIndex828
					if !_rules[ruleMiscodedChar]() {
						goto l826
					}
				}
			l828:
				add(ruleNameLowerChar, position827)
			}
			return true
		l826:
			position, tokenIndex = position826, tokenIndex826
			return false
		},
		/* 93 MiscodedChar <- <'�'> */
		func() bool {
			position831, tokenIndex831 := position, tokenIndex
			{
				position832 := position
				if buffer[position] != rune('�') {
					goto l831
				}
				position++
				add(ruleMiscodedChar, position832)
			}
			return true
		l831:
			position, tokenIndex = position831, tokenIndex831
			return false
		},
		/* 94 LowerCharExtended <- <('æ' / 'œ' / 'à' / 'â' / 'å' / 'ã' / 'ä' / 'á' / 'ç' / 'č' / 'é' / 'è' / 'ë' / 'í' / 'ì' / 'ï' / 'ň' / 'ñ' / 'ñ' / 'ó' / 'ò' / 'ô' / 'ø' / 'õ' / 'ö' / 'ú' / 'ù' / 'ü' / 'ŕ' / 'ř' / 'ŗ' / 'ſ' / 'š' / 'š' / 'ş' / 'ž')> */
		func() bool {
			position833, tokenIndex833 := position, tokenIndex
			{
				position834 := position
				{
					position835, tokenIndex835 := position, tokenIndex
					if buffer[position] != rune('æ') {
						goto l836
					}
					position++
					goto l835
				l836:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('œ') {
						goto l837
					}
					position++
					goto l835
				l837:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('à') {
						goto l838
					}
					position++
					goto l835
				l838:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('â') {
						goto l839
					}
					position++
					goto l835
				l839:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('å') {
						goto l840
					}
					position++
					goto l835
				l840:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('ã') {
						goto l841
					}
					position++
					goto l835
				l841:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('ä') {
						goto l842
					}
					position++
					goto l835
				l842:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('á') {
						goto l843
					}
					position++
					goto l835
				l843:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('ç') {
						goto l844
					}
					position++
					goto l835
				l844:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('č') {
						goto l845
					}
					position++
					goto l835
				l845:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('é') {
						goto l846
					}
					position++
					goto l835
				l846:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('è') {
						goto l847
					}
					position++
					goto l835
				l847:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('ë') {
						goto l848
					}
					position++
					goto l835
				l848:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('í') {
						goto l849
					}
					position++
					goto l835
				l849:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('ì') {
						goto l850
					}
					position++
					goto l835
				l850:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('ï') {
						goto l851
					}
					position++
					goto l835
				l851:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('ň') {
						goto l852
					}
					position++
					goto l835
				l852:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('ñ') {
						goto l853
					}
					position++
					goto l835
				l853:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('ñ') {
						goto l854
					}
					position++
					goto l835
				l854:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('ó') {
						goto l855
					}
					position++
					goto l835
				l855:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('ò') {
						goto l856
					}
					position++
					goto l835
				l856:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('ô') {
						goto l857
					}
					position++
					goto l835
				l857:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('ø') {
						goto l858
					}
					position++
					goto l835
				l858:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('õ') {
						goto l859
					}
					position++
					goto l835
				l859:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('ö') {
						goto l860
					}
					position++
					goto l835
				l860:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('ú') {
						goto l861
					}
					position++
					goto l835
				l861:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('ù') {
						goto l862
					}
					position++
					goto l835
				l862:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('ü') {
						goto l863
					}
					position++
					goto l835
				l863:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('ŕ') {
						goto l864
					}
					position++
					goto l835
				l864:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('ř') {
						goto l865
					}
					position++
					goto l835
				l865:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('ŗ') {
						goto l866
					}
					position++
					goto l835
				l866:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('ſ') {
						goto l867
					}
					position++
					goto l835
				l867:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('š') {
						goto l868
					}
					position++
					goto l835
				l868:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('š') {
						goto l869
					}
					position++
					goto l835
				l869:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('ş') {
						goto l870
					}
					position++
					goto l835
				l870:
					position, tokenIndex = position835, tokenIndex835
					if buffer[position] != rune('ž') {
						goto l833
					}
					position++
				}
			l835:
				add(ruleLowerCharExtended, position834)
			}
			return true
		l833:
			position, tokenIndex = position833, tokenIndex833
			return false
		},
		/* 95 LowerChar <- <lASCII> */
		func() bool {
			position871, tokenIndex871 := position, tokenIndex
			{
				position872 := position
				if !_rules[rulelASCII]() {
					goto l871
				}
				add(ruleLowerChar, position872)
			}
			return true
		l871:
			position, tokenIndex = position871, tokenIndex871
			return false
		},
		/* 96 SpaceCharEOI <- <(_ / !.)> */
		func() bool {
			position873, tokenIndex873 := position, tokenIndex
			{
				position874 := position
				{
					position875, tokenIndex875 := position, tokenIndex
					if !_rules[rule_]() {
						goto l876
					}
					goto l875
				l876:
					position, tokenIndex = position875, tokenIndex875
					{
						position877, tokenIndex877 := position, tokenIndex
						if !matchDot() {
							goto l877
						}
						goto l873
					l877:
						position, tokenIndex = position877, tokenIndex877
					}
				}
			l875:
				add(ruleSpaceCharEOI, position874)
			}
			return true
		l873:
			position, tokenIndex = position873, tokenIndex873
			return false
		},
		/* 97 nums <- <[0-9]> */
		func() bool {
			position878, tokenIndex878 := position, tokenIndex
			{
				position879 := position
				if c := buffer[position]; c < rune('0') || c > rune('9') {
					goto l878
				}
				position++
				add(rulenums, position879)
			}
			return true
		l878:
			position, tokenIndex = position878, tokenIndex878
			return false
		},
		/* 98 lASCII <- <[a-z]> */
		func() bool {
			position880, tokenIndex880 := position, tokenIndex
			{
				position881 := position
				if c := buffer[position]; c < rune('a') || c > rune('z') {
					goto l880
				}
				position++
				add(rulelASCII, position881)
			}
			return true
		l880:
			position, tokenIndex = position880, tokenIndex880
			return false
		},
		/* 99 hASCII <- <[A-Z]> */
		func() bool {
			position882, tokenIndex882 := position, tokenIndex
			{
				position883 := position
				if c := buffer[position]; c < rune('A') || c > rune('Z') {
					goto l882
				}
				position++
				add(rulehASCII, position883)
			}
			return true
		l882:
			position, tokenIndex = position882, tokenIndex882
			return false
		},
		/* 100 apostr <- <'\''> */
		func() bool {
			position884, tokenIndex884 := position, tokenIndex
			{
				position885 := position
				if buffer[position] != rune('\'') {
					goto l884
				}
				position++
				add(ruleapostr, position885)
			}
			return true
		l884:
			position, tokenIndex = position884, tokenIndex884
			return false
		},
		/* 101 dash <- <'-'> */
		func() bool {
			position886, tokenIndex886 := position, tokenIndex
			{
				position887 := position
				if buffer[position] != rune('-') {
					goto l886
				}
				position++
				add(ruledash, position887)
			}
			return true
		l886:
			position, tokenIndex = position886, tokenIndex886
			return false
		},
		/* 102 _ <- <(MultipleSpace / SingleSpace)> */
		func() bool {
			position888, tokenIndex888 := position, tokenIndex
			{
				position889 := position
				{
					position890, tokenIndex890 := position, tokenIndex
					if !_rules[ruleMultipleSpace]() {
						goto l891
					}
					goto l890
				l891:
					position, tokenIndex = position890, tokenIndex890
					if !_rules[ruleSingleSpace]() {
						goto l888
					}
				}
			l890:
				add(rule_, position889)
			}
			return true
		l888:
			position, tokenIndex = position888, tokenIndex888
			return false
		},
		/* 103 MultipleSpace <- <(SingleSpace SingleSpace+)> */
		func() bool {
			position892, tokenIndex892 := position, tokenIndex
			{
				position893 := position
				if !_rules[ruleSingleSpace]() {
					goto l892
				}
				if !_rules[ruleSingleSpace]() {
					goto l892
				}
			l894:
				{
					position895, tokenIndex895 := position, tokenIndex
					if !_rules[ruleSingleSpace]() {
						goto l895
					}
					goto l894
				l895:
					position, tokenIndex = position895, tokenIndex895
				}
				add(ruleMultipleSpace, position893)
			}
			return true
		l892:
			position, tokenIndex = position892, tokenIndex892
			return false
		},
		/* 104 SingleSpace <- <(' ' / OtherSpace)> */
		func() bool {
			position896, tokenIndex896 := position, tokenIndex
			{
				position897 := position
				{
					position898, tokenIndex898 := position, tokenIndex
					if buffer[position] != rune(' ') {
						goto l899
					}
					position++
					goto l898
				l899:
					position, tokenIndex = position898, tokenIndex898
					if !_rules[ruleOtherSpace]() {
						goto l896
					}
				}
			l898:
				add(ruleSingleSpace, position897)
			}
			return true
		l896:
			position, tokenIndex = position896, tokenIndex896
			return false
		},
		/* 105 OtherSpace <- <('\u3000' / '\u00a0' / '\t' / '\r' / '\n' / '\f' / '\v')> */
		func() bool {
			position900, tokenIndex900 := position, tokenIndex
			{
				position901 := position
				{
					position902, tokenIndex902 := position, tokenIndex
					if buffer[position] != rune('\u3000') {
						goto l903
					}
					position++
					goto l902
				l903:
					position, tokenIndex = position902, tokenIndex902
					if buffer[position] != rune('\u00a0') {
						goto l904
					}
					position++
					goto l902
				l904:
					position, tokenIndex = position902, tokenIndex902
					if buffer[position] != rune('\t') {
						goto l905
					}
					position++
					goto l902
				l905:
					position, tokenIndex = position902, tokenIndex902
					if buffer[position] != rune('\r') {
						goto l906
					}
					position++
					goto l902
				l906:
					position, tokenIndex = position902, tokenIndex902
					if buffer[position] != rune('\n') {
						goto l907
					}
					position++
					goto l902
				l907:
					position, tokenIndex = position902, tokenIndex902
					if buffer[position] != rune('\f') {
						goto l908
					}
					position++
					goto l902
				l908:
					position, tokenIndex = position902, tokenIndex902
					if buffer[position] != rune('\v') {
						goto l900
					}
					position++
				}
			l902:
				add(ruleOtherSpace, position901)
			}
			return true
		l900:
			position, tokenIndex = position900, tokenIndex900
			return false
		},
		/* 107 Action0 <- <{ p.AddWarn(YearCharWarn) }> */
		func() bool {
			{
				add(ruleAction0, position)
			}
			return true
		},
	}
	p.rules = _rules
}
