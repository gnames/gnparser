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
	ruleCombinationAuthorship
	ruleBasionymAuthorshipYearMisformed
	ruleBasionymAuthorship
	ruleBasionymAuthorship1
	ruleBasionymAuthorship2
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
	ruleAuthorWord1
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
	"CombinationAuthorship",
	"BasionymAuthorshipYearMisformed",
	"BasionymAuthorship",
	"BasionymAuthorship1",
	"BasionymAuthorship2",
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
	"AuthorWord1",
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
	rules  [106]func() bool
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
		/* 1 Tail <- <.*> */
		func() bool {
			{
				position6 := position
			l7:
				{
					position8, tokenIndex8 := position, tokenIndex
					if !matchDot() {
						goto l8
					}
					goto l7
				l8:
					position, tokenIndex = position8, tokenIndex8
				}
				add(ruleTail, position6)
			}
			return true
		},
		/* 2 Name <- <(NamedHybrid / HybridFormula / SingleName)> */
		func() bool {
			position9, tokenIndex9 := position, tokenIndex
			{
				position10 := position
				{
					position11, tokenIndex11 := position, tokenIndex
					if !_rules[ruleNamedHybrid]() {
						goto l12
					}
					goto l11
				l12:
					position, tokenIndex = position11, tokenIndex11
					if !_rules[ruleHybridFormula]() {
						goto l13
					}
					goto l11
				l13:
					position, tokenIndex = position11, tokenIndex11
					if !_rules[ruleSingleName]() {
						goto l9
					}
				}
			l11:
				add(ruleName, position10)
			}
			return true
		l9:
			position, tokenIndex = position9, tokenIndex9
			return false
		},
		/* 3 HybridFormula <- <(SingleName (_ (HybridFormulaPart / HybridFormulaFull))+)> */
		func() bool {
			position14, tokenIndex14 := position, tokenIndex
			{
				position15 := position
				if !_rules[ruleSingleName]() {
					goto l14
				}
				if !_rules[rule_]() {
					goto l14
				}
				{
					position18, tokenIndex18 := position, tokenIndex
					if !_rules[ruleHybridFormulaPart]() {
						goto l19
					}
					goto l18
				l19:
					position, tokenIndex = position18, tokenIndex18
					if !_rules[ruleHybridFormulaFull]() {
						goto l14
					}
				}
			l18:
			l16:
				{
					position17, tokenIndex17 := position, tokenIndex
					if !_rules[rule_]() {
						goto l17
					}
					{
						position20, tokenIndex20 := position, tokenIndex
						if !_rules[ruleHybridFormulaPart]() {
							goto l21
						}
						goto l20
					l21:
						position, tokenIndex = position20, tokenIndex20
						if !_rules[ruleHybridFormulaFull]() {
							goto l17
						}
					}
				l20:
					goto l16
				l17:
					position, tokenIndex = position17, tokenIndex17
				}
				add(ruleHybridFormula, position15)
			}
			return true
		l14:
			position, tokenIndex = position14, tokenIndex14
			return false
		},
		/* 4 HybridFormulaFull <- <(HybridChar (_ SingleName)?)> */
		func() bool {
			position22, tokenIndex22 := position, tokenIndex
			{
				position23 := position
				if !_rules[ruleHybridChar]() {
					goto l22
				}
				{
					position24, tokenIndex24 := position, tokenIndex
					if !_rules[rule_]() {
						goto l24
					}
					if !_rules[ruleSingleName]() {
						goto l24
					}
					goto l25
				l24:
					position, tokenIndex = position24, tokenIndex24
				}
			l25:
				add(ruleHybridFormulaFull, position23)
			}
			return true
		l22:
			position, tokenIndex = position22, tokenIndex22
			return false
		},
		/* 5 HybridFormulaPart <- <(HybridChar _ SpeciesEpithet (_ InfraspGroup)?)> */
		func() bool {
			position26, tokenIndex26 := position, tokenIndex
			{
				position27 := position
				if !_rules[ruleHybridChar]() {
					goto l26
				}
				if !_rules[rule_]() {
					goto l26
				}
				if !_rules[ruleSpeciesEpithet]() {
					goto l26
				}
				{
					position28, tokenIndex28 := position, tokenIndex
					if !_rules[rule_]() {
						goto l28
					}
					if !_rules[ruleInfraspGroup]() {
						goto l28
					}
					goto l29
				l28:
					position, tokenIndex = position28, tokenIndex28
				}
			l29:
				add(ruleHybridFormulaPart, position27)
			}
			return true
		l26:
			position, tokenIndex = position26, tokenIndex26
			return false
		},
		/* 6 NamedHybrid <- <(NamedGenusHybrid / NamedSpeciesHybrid)> */
		func() bool {
			position30, tokenIndex30 := position, tokenIndex
			{
				position31 := position
				{
					position32, tokenIndex32 := position, tokenIndex
					if !_rules[ruleNamedGenusHybrid]() {
						goto l33
					}
					goto l32
				l33:
					position, tokenIndex = position32, tokenIndex32
					if !_rules[ruleNamedSpeciesHybrid]() {
						goto l30
					}
				}
			l32:
				add(ruleNamedHybrid, position31)
			}
			return true
		l30:
			position, tokenIndex = position30, tokenIndex30
			return false
		},
		/* 7 NamedSpeciesHybrid <- <(GenusWord _ HybridChar _? SpeciesEpithet)> */
		func() bool {
			position34, tokenIndex34 := position, tokenIndex
			{
				position35 := position
				if !_rules[ruleGenusWord]() {
					goto l34
				}
				if !_rules[rule_]() {
					goto l34
				}
				if !_rules[ruleHybridChar]() {
					goto l34
				}
				{
					position36, tokenIndex36 := position, tokenIndex
					if !_rules[rule_]() {
						goto l36
					}
					goto l37
				l36:
					position, tokenIndex = position36, tokenIndex36
				}
			l37:
				if !_rules[ruleSpeciesEpithet]() {
					goto l34
				}
				add(ruleNamedSpeciesHybrid, position35)
			}
			return true
		l34:
			position, tokenIndex = position34, tokenIndex34
			return false
		},
		/* 8 NamedGenusHybrid <- <(HybridChar _? SingleName)> */
		func() bool {
			position38, tokenIndex38 := position, tokenIndex
			{
				position39 := position
				if !_rules[ruleHybridChar]() {
					goto l38
				}
				{
					position40, tokenIndex40 := position, tokenIndex
					if !_rules[rule_]() {
						goto l40
					}
					goto l41
				l40:
					position, tokenIndex = position40, tokenIndex40
				}
			l41:
				if !_rules[ruleSingleName]() {
					goto l38
				}
				add(ruleNamedGenusHybrid, position39)
			}
			return true
		l38:
			position, tokenIndex = position38, tokenIndex38
			return false
		},
		/* 9 SingleName <- <(NameComp / NameApprox / NameSpecies / NameUninomial)> */
		func() bool {
			position42, tokenIndex42 := position, tokenIndex
			{
				position43 := position
				{
					position44, tokenIndex44 := position, tokenIndex
					if !_rules[ruleNameComp]() {
						goto l45
					}
					goto l44
				l45:
					position, tokenIndex = position44, tokenIndex44
					if !_rules[ruleNameApprox]() {
						goto l46
					}
					goto l44
				l46:
					position, tokenIndex = position44, tokenIndex44
					if !_rules[ruleNameSpecies]() {
						goto l47
					}
					goto l44
				l47:
					position, tokenIndex = position44, tokenIndex44
					if !_rules[ruleNameUninomial]() {
						goto l42
					}
				}
			l44:
				add(ruleSingleName, position43)
			}
			return true
		l42:
			position, tokenIndex = position42, tokenIndex42
			return false
		},
		/* 10 NameUninomial <- <(UninomialCombo / Uninomial)> */
		func() bool {
			position48, tokenIndex48 := position, tokenIndex
			{
				position49 := position
				{
					position50, tokenIndex50 := position, tokenIndex
					if !_rules[ruleUninomialCombo]() {
						goto l51
					}
					goto l50
				l51:
					position, tokenIndex = position50, tokenIndex50
					if !_rules[ruleUninomial]() {
						goto l48
					}
				}
			l50:
				add(ruleNameUninomial, position49)
			}
			return true
		l48:
			position, tokenIndex = position48, tokenIndex48
			return false
		},
		/* 11 NameApprox <- <(GenusWord (_ SpeciesEpithet)? _ Approximation ApproxNameIgnored)> */
		func() bool {
			position52, tokenIndex52 := position, tokenIndex
			{
				position53 := position
				if !_rules[ruleGenusWord]() {
					goto l52
				}
				{
					position54, tokenIndex54 := position, tokenIndex
					if !_rules[rule_]() {
						goto l54
					}
					if !_rules[ruleSpeciesEpithet]() {
						goto l54
					}
					goto l55
				l54:
					position, tokenIndex = position54, tokenIndex54
				}
			l55:
				if !_rules[rule_]() {
					goto l52
				}
				if !_rules[ruleApproximation]() {
					goto l52
				}
				if !_rules[ruleApproxNameIgnored]() {
					goto l52
				}
				add(ruleNameApprox, position53)
			}
			return true
		l52:
			position, tokenIndex = position52, tokenIndex52
			return false
		},
		/* 12 NameComp <- <(GenusWord _ Comparison (_ SpeciesEpithet)?)> */
		func() bool {
			position56, tokenIndex56 := position, tokenIndex
			{
				position57 := position
				if !_rules[ruleGenusWord]() {
					goto l56
				}
				if !_rules[rule_]() {
					goto l56
				}
				if !_rules[ruleComparison]() {
					goto l56
				}
				{
					position58, tokenIndex58 := position, tokenIndex
					if !_rules[rule_]() {
						goto l58
					}
					if !_rules[ruleSpeciesEpithet]() {
						goto l58
					}
					goto l59
				l58:
					position, tokenIndex = position58, tokenIndex58
				}
			l59:
				add(ruleNameComp, position57)
			}
			return true
		l56:
			position, tokenIndex = position56, tokenIndex56
			return false
		},
		/* 13 NameSpecies <- <(GenusWord (_? (SubGenus / SubGenusOrSuperspecies))? _ SpeciesEpithet (_ InfraspGroup)?)> */
		func() bool {
			position60, tokenIndex60 := position, tokenIndex
			{
				position61 := position
				if !_rules[ruleGenusWord]() {
					goto l60
				}
				{
					position62, tokenIndex62 := position, tokenIndex
					{
						position64, tokenIndex64 := position, tokenIndex
						if !_rules[rule_]() {
							goto l64
						}
						goto l65
					l64:
						position, tokenIndex = position64, tokenIndex64
					}
				l65:
					{
						position66, tokenIndex66 := position, tokenIndex
						if !_rules[ruleSubGenus]() {
							goto l67
						}
						goto l66
					l67:
						position, tokenIndex = position66, tokenIndex66
						if !_rules[ruleSubGenusOrSuperspecies]() {
							goto l62
						}
					}
				l66:
					goto l63
				l62:
					position, tokenIndex = position62, tokenIndex62
				}
			l63:
				if !_rules[rule_]() {
					goto l60
				}
				if !_rules[ruleSpeciesEpithet]() {
					goto l60
				}
				{
					position68, tokenIndex68 := position, tokenIndex
					if !_rules[rule_]() {
						goto l68
					}
					if !_rules[ruleInfraspGroup]() {
						goto l68
					}
					goto l69
				l68:
					position, tokenIndex = position68, tokenIndex68
				}
			l69:
				add(ruleNameSpecies, position61)
			}
			return true
		l60:
			position, tokenIndex = position60, tokenIndex60
			return false
		},
		/* 14 GenusWord <- <((AbbrGenus / UninomialWord) !(_ AuthorWord))> */
		func() bool {
			position70, tokenIndex70 := position, tokenIndex
			{
				position71 := position
				{
					position72, tokenIndex72 := position, tokenIndex
					if !_rules[ruleAbbrGenus]() {
						goto l73
					}
					goto l72
				l73:
					position, tokenIndex = position72, tokenIndex72
					if !_rules[ruleUninomialWord]() {
						goto l70
					}
				}
			l72:
				{
					position74, tokenIndex74 := position, tokenIndex
					if !_rules[rule_]() {
						goto l74
					}
					if !_rules[ruleAuthorWord]() {
						goto l74
					}
					goto l70
				l74:
					position, tokenIndex = position74, tokenIndex74
				}
				add(ruleGenusWord, position71)
			}
			return true
		l70:
			position, tokenIndex = position70, tokenIndex70
			return false
		},
		/* 15 InfraspGroup <- <(InfraspEpithet (_ InfraspEpithet)? (_ InfraspEpithet)?)> */
		func() bool {
			position75, tokenIndex75 := position, tokenIndex
			{
				position76 := position
				if !_rules[ruleInfraspEpithet]() {
					goto l75
				}
				{
					position77, tokenIndex77 := position, tokenIndex
					if !_rules[rule_]() {
						goto l77
					}
					if !_rules[ruleInfraspEpithet]() {
						goto l77
					}
					goto l78
				l77:
					position, tokenIndex = position77, tokenIndex77
				}
			l78:
				{
					position79, tokenIndex79 := position, tokenIndex
					if !_rules[rule_]() {
						goto l79
					}
					if !_rules[ruleInfraspEpithet]() {
						goto l79
					}
					goto l80
				l79:
					position, tokenIndex = position79, tokenIndex79
				}
			l80:
				add(ruleInfraspGroup, position76)
			}
			return true
		l75:
			position, tokenIndex = position75, tokenIndex75
			return false
		},
		/* 16 InfraspEpithet <- <((Rank _?)? !AuthorEx Word (_ Authorship)?)> */
		func() bool {
			position81, tokenIndex81 := position, tokenIndex
			{
				position82 := position
				{
					position83, tokenIndex83 := position, tokenIndex
					if !_rules[ruleRank]() {
						goto l83
					}
					{
						position85, tokenIndex85 := position, tokenIndex
						if !_rules[rule_]() {
							goto l85
						}
						goto l86
					l85:
						position, tokenIndex = position85, tokenIndex85
					}
				l86:
					goto l84
				l83:
					position, tokenIndex = position83, tokenIndex83
				}
			l84:
				{
					position87, tokenIndex87 := position, tokenIndex
					if !_rules[ruleAuthorEx]() {
						goto l87
					}
					goto l81
				l87:
					position, tokenIndex = position87, tokenIndex87
				}
				if !_rules[ruleWord]() {
					goto l81
				}
				{
					position88, tokenIndex88 := position, tokenIndex
					if !_rules[rule_]() {
						goto l88
					}
					if !_rules[ruleAuthorship]() {
						goto l88
					}
					goto l89
				l88:
					position, tokenIndex = position88, tokenIndex88
				}
			l89:
				add(ruleInfraspEpithet, position82)
			}
			return true
		l81:
			position, tokenIndex = position81, tokenIndex81
			return false
		},
		/* 17 SpeciesEpithet <- <(!AuthorEx Word (_? Authorship)? ','? &(SpaceCharEOI / '('))> */
		func() bool {
			position90, tokenIndex90 := position, tokenIndex
			{
				position91 := position
				{
					position92, tokenIndex92 := position, tokenIndex
					if !_rules[ruleAuthorEx]() {
						goto l92
					}
					goto l90
				l92:
					position, tokenIndex = position92, tokenIndex92
				}
				if !_rules[ruleWord]() {
					goto l90
				}
				{
					position93, tokenIndex93 := position, tokenIndex
					{
						position95, tokenIndex95 := position, tokenIndex
						if !_rules[rule_]() {
							goto l95
						}
						goto l96
					l95:
						position, tokenIndex = position95, tokenIndex95
					}
				l96:
					if !_rules[ruleAuthorship]() {
						goto l93
					}
					goto l94
				l93:
					position, tokenIndex = position93, tokenIndex93
				}
			l94:
				{
					position97, tokenIndex97 := position, tokenIndex
					if buffer[position] != rune(',') {
						goto l97
					}
					position++
					goto l98
				l97:
					position, tokenIndex = position97, tokenIndex97
				}
			l98:
				{
					position99, tokenIndex99 := position, tokenIndex
					{
						position100, tokenIndex100 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l101
						}
						goto l100
					l101:
						position, tokenIndex = position100, tokenIndex100
						if buffer[position] != rune('(') {
							goto l90
						}
						position++
					}
				l100:
					position, tokenIndex = position99, tokenIndex99
				}
				add(ruleSpeciesEpithet, position91)
			}
			return true
		l90:
			position, tokenIndex = position90, tokenIndex90
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
		/* 28 UninomialCombo1 <- <(UninomialWord _? SubGenus _? Authorship .?)> */
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
					if !_rules[rule_]() {
						goto l204
					}
					goto l205
				l204:
					position, tokenIndex = position204, tokenIndex204
				}
			l205:
				if !_rules[ruleAuthorship]() {
					goto l200
				}
				{
					position206, tokenIndex206 := position, tokenIndex
					if !matchDot() {
						goto l206
					}
					goto l207
				l206:
					position, tokenIndex = position206, tokenIndex206
				}
			l207:
				add(ruleUninomialCombo1, position201)
			}
			return true
		l200:
			position, tokenIndex = position200, tokenIndex200
			return false
		},
		/* 29 UninomialCombo2 <- <(Uninomial _? RankUninomial _? Uninomial)> */
		func() bool {
			position208, tokenIndex208 := position, tokenIndex
			{
				position209 := position
				if !_rules[ruleUninomial]() {
					goto l208
				}
				{
					position210, tokenIndex210 := position, tokenIndex
					if !_rules[rule_]() {
						goto l210
					}
					goto l211
				l210:
					position, tokenIndex = position210, tokenIndex210
				}
			l211:
				if !_rules[ruleRankUninomial]() {
					goto l208
				}
				{
					position212, tokenIndex212 := position, tokenIndex
					if !_rules[rule_]() {
						goto l212
					}
					goto l213
				l212:
					position, tokenIndex = position212, tokenIndex212
				}
			l213:
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
		/* 30 RankUninomial <- <((('s' 'e' 'c' 't') / ('s' 'u' 'b' 's' 'e' 'c' 't') / ('t' 'r' 'i' 'b') / ('s' 'u' 'b' 't' 'r' 'i' 'b') / ('s' 'u' 'b' 's' 'e' 'r') / ('s' 'e' 'r') / ('s' 'u' 'b' 'g' 'e' 'n') / ('f' 'a' 'm') / ('s' 'u' 'b' 'f' 'a' 'm') / ('s' 'u' 'p' 'e' 'r' 't' 'r' 'i' 'b')) '.'?)> */
		func() bool {
			position214, tokenIndex214 := position, tokenIndex
			{
				position215 := position
				{
					position216, tokenIndex216 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l217
					}
					position++
					if buffer[position] != rune('e') {
						goto l217
					}
					position++
					if buffer[position] != rune('c') {
						goto l217
					}
					position++
					if buffer[position] != rune('t') {
						goto l217
					}
					position++
					goto l216
				l217:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('s') {
						goto l218
					}
					position++
					if buffer[position] != rune('u') {
						goto l218
					}
					position++
					if buffer[position] != rune('b') {
						goto l218
					}
					position++
					if buffer[position] != rune('s') {
						goto l218
					}
					position++
					if buffer[position] != rune('e') {
						goto l218
					}
					position++
					if buffer[position] != rune('c') {
						goto l218
					}
					position++
					if buffer[position] != rune('t') {
						goto l218
					}
					position++
					goto l216
				l218:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('t') {
						goto l219
					}
					position++
					if buffer[position] != rune('r') {
						goto l219
					}
					position++
					if buffer[position] != rune('i') {
						goto l219
					}
					position++
					if buffer[position] != rune('b') {
						goto l219
					}
					position++
					goto l216
				l219:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('s') {
						goto l220
					}
					position++
					if buffer[position] != rune('u') {
						goto l220
					}
					position++
					if buffer[position] != rune('b') {
						goto l220
					}
					position++
					if buffer[position] != rune('t') {
						goto l220
					}
					position++
					if buffer[position] != rune('r') {
						goto l220
					}
					position++
					if buffer[position] != rune('i') {
						goto l220
					}
					position++
					if buffer[position] != rune('b') {
						goto l220
					}
					position++
					goto l216
				l220:
					position, tokenIndex = position216, tokenIndex216
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
					if buffer[position] != rune('s') {
						goto l221
					}
					position++
					if buffer[position] != rune('e') {
						goto l221
					}
					position++
					if buffer[position] != rune('r') {
						goto l221
					}
					position++
					goto l216
				l221:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('s') {
						goto l222
					}
					position++
					if buffer[position] != rune('e') {
						goto l222
					}
					position++
					if buffer[position] != rune('r') {
						goto l222
					}
					position++
					goto l216
				l222:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('s') {
						goto l223
					}
					position++
					if buffer[position] != rune('u') {
						goto l223
					}
					position++
					if buffer[position] != rune('b') {
						goto l223
					}
					position++
					if buffer[position] != rune('g') {
						goto l223
					}
					position++
					if buffer[position] != rune('e') {
						goto l223
					}
					position++
					if buffer[position] != rune('n') {
						goto l223
					}
					position++
					goto l216
				l223:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('f') {
						goto l224
					}
					position++
					if buffer[position] != rune('a') {
						goto l224
					}
					position++
					if buffer[position] != rune('m') {
						goto l224
					}
					position++
					goto l216
				l224:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('s') {
						goto l225
					}
					position++
					if buffer[position] != rune('u') {
						goto l225
					}
					position++
					if buffer[position] != rune('b') {
						goto l225
					}
					position++
					if buffer[position] != rune('f') {
						goto l225
					}
					position++
					if buffer[position] != rune('a') {
						goto l225
					}
					position++
					if buffer[position] != rune('m') {
						goto l225
					}
					position++
					goto l216
				l225:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('s') {
						goto l214
					}
					position++
					if buffer[position] != rune('u') {
						goto l214
					}
					position++
					if buffer[position] != rune('p') {
						goto l214
					}
					position++
					if buffer[position] != rune('e') {
						goto l214
					}
					position++
					if buffer[position] != rune('r') {
						goto l214
					}
					position++
					if buffer[position] != rune('t') {
						goto l214
					}
					position++
					if buffer[position] != rune('r') {
						goto l214
					}
					position++
					if buffer[position] != rune('i') {
						goto l214
					}
					position++
					if buffer[position] != rune('b') {
						goto l214
					}
					position++
				}
			l216:
				{
					position226, tokenIndex226 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l226
					}
					position++
					goto l227
				l226:
					position, tokenIndex = position226, tokenIndex226
				}
			l227:
				add(ruleRankUninomial, position215)
			}
			return true
		l214:
			position, tokenIndex = position214, tokenIndex214
			return false
		},
		/* 31 Uninomial <- <(UninomialWord (_ Authorship)?)> */
		func() bool {
			position228, tokenIndex228 := position, tokenIndex
			{
				position229 := position
				if !_rules[ruleUninomialWord]() {
					goto l228
				}
				{
					position230, tokenIndex230 := position, tokenIndex
					if !_rules[rule_]() {
						goto l230
					}
					if !_rules[ruleAuthorship]() {
						goto l230
					}
					goto l231
				l230:
					position, tokenIndex = position230, tokenIndex230
				}
			l231:
				add(ruleUninomial, position229)
			}
			return true
		l228:
			position, tokenIndex = position228, tokenIndex228
			return false
		},
		/* 32 UninomialWord <- <(CapWord / TwoLetterGenus)> */
		func() bool {
			position232, tokenIndex232 := position, tokenIndex
			{
				position233 := position
				{
					position234, tokenIndex234 := position, tokenIndex
					if !_rules[ruleCapWord]() {
						goto l235
					}
					goto l234
				l235:
					position, tokenIndex = position234, tokenIndex234
					if !_rules[ruleTwoLetterGenus]() {
						goto l232
					}
				}
			l234:
				add(ruleUninomialWord, position233)
			}
			return true
		l232:
			position, tokenIndex = position232, tokenIndex232
			return false
		},
		/* 33 AbbrGenus <- <(UpperChar LowerChar* '.')> */
		func() bool {
			position236, tokenIndex236 := position, tokenIndex
			{
				position237 := position
				if !_rules[ruleUpperChar]() {
					goto l236
				}
			l238:
				{
					position239, tokenIndex239 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l239
					}
					goto l238
				l239:
					position, tokenIndex = position239, tokenIndex239
				}
				if buffer[position] != rune('.') {
					goto l236
				}
				position++
				add(ruleAbbrGenus, position237)
			}
			return true
		l236:
			position, tokenIndex = position236, tokenIndex236
			return false
		},
		/* 34 CapWord <- <(CapWord2 / CapWord1)> */
		func() bool {
			position240, tokenIndex240 := position, tokenIndex
			{
				position241 := position
				{
					position242, tokenIndex242 := position, tokenIndex
					if !_rules[ruleCapWord2]() {
						goto l243
					}
					goto l242
				l243:
					position, tokenIndex = position242, tokenIndex242
					if !_rules[ruleCapWord1]() {
						goto l240
					}
				}
			l242:
				add(ruleCapWord, position241)
			}
			return true
		l240:
			position, tokenIndex = position240, tokenIndex240
			return false
		},
		/* 35 CapWord1 <- <(NameUpperChar NameLowerChar NameLowerChar+ '?'?)> */
		func() bool {
			position244, tokenIndex244 := position, tokenIndex
			{
				position245 := position
				if !_rules[ruleNameUpperChar]() {
					goto l244
				}
				if !_rules[ruleNameLowerChar]() {
					goto l244
				}
				if !_rules[ruleNameLowerChar]() {
					goto l244
				}
			l246:
				{
					position247, tokenIndex247 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l247
					}
					goto l246
				l247:
					position, tokenIndex = position247, tokenIndex247
				}
				{
					position248, tokenIndex248 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l248
					}
					position++
					goto l249
				l248:
					position, tokenIndex = position248, tokenIndex248
				}
			l249:
				add(ruleCapWord1, position245)
			}
			return true
		l244:
			position, tokenIndex = position244, tokenIndex244
			return false
		},
		/* 36 CapWord2 <- <(CapWord1 dash (CapWord1 / Word1))> */
		func() bool {
			position250, tokenIndex250 := position, tokenIndex
			{
				position251 := position
				if !_rules[ruleCapWord1]() {
					goto l250
				}
				if !_rules[ruledash]() {
					goto l250
				}
				{
					position252, tokenIndex252 := position, tokenIndex
					if !_rules[ruleCapWord1]() {
						goto l253
					}
					goto l252
				l253:
					position, tokenIndex = position252, tokenIndex252
					if !_rules[ruleWord1]() {
						goto l250
					}
				}
			l252:
				add(ruleCapWord2, position251)
			}
			return true
		l250:
			position, tokenIndex = position250, tokenIndex250
			return false
		},
		/* 37 TwoLetterGenus <- <(('C' 'a') / ('E' 'a') / ('G' 'e') / ('I' 'a') / ('I' 'o') / ('I' 'x') / ('L' 'o') / ('O' 'a') / ('R' 'a') / ('T' 'y') / ('U' 'a') / ('A' 'a') / ('J' 'a') / ('Z' 'u') / ('L' 'a') / ('Q' 'u') / ('A' 's') / ('B' 'a'))> */
		func() bool {
			position254, tokenIndex254 := position, tokenIndex
			{
				position255 := position
				{
					position256, tokenIndex256 := position, tokenIndex
					if buffer[position] != rune('C') {
						goto l257
					}
					position++
					if buffer[position] != rune('a') {
						goto l257
					}
					position++
					goto l256
				l257:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('E') {
						goto l258
					}
					position++
					if buffer[position] != rune('a') {
						goto l258
					}
					position++
					goto l256
				l258:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('G') {
						goto l259
					}
					position++
					if buffer[position] != rune('e') {
						goto l259
					}
					position++
					goto l256
				l259:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('I') {
						goto l260
					}
					position++
					if buffer[position] != rune('a') {
						goto l260
					}
					position++
					goto l256
				l260:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('I') {
						goto l261
					}
					position++
					if buffer[position] != rune('o') {
						goto l261
					}
					position++
					goto l256
				l261:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('I') {
						goto l262
					}
					position++
					if buffer[position] != rune('x') {
						goto l262
					}
					position++
					goto l256
				l262:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('L') {
						goto l263
					}
					position++
					if buffer[position] != rune('o') {
						goto l263
					}
					position++
					goto l256
				l263:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('O') {
						goto l264
					}
					position++
					if buffer[position] != rune('a') {
						goto l264
					}
					position++
					goto l256
				l264:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('R') {
						goto l265
					}
					position++
					if buffer[position] != rune('a') {
						goto l265
					}
					position++
					goto l256
				l265:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('T') {
						goto l266
					}
					position++
					if buffer[position] != rune('y') {
						goto l266
					}
					position++
					goto l256
				l266:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('U') {
						goto l267
					}
					position++
					if buffer[position] != rune('a') {
						goto l267
					}
					position++
					goto l256
				l267:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('A') {
						goto l268
					}
					position++
					if buffer[position] != rune('a') {
						goto l268
					}
					position++
					goto l256
				l268:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('J') {
						goto l269
					}
					position++
					if buffer[position] != rune('a') {
						goto l269
					}
					position++
					goto l256
				l269:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('Z') {
						goto l270
					}
					position++
					if buffer[position] != rune('u') {
						goto l270
					}
					position++
					goto l256
				l270:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('L') {
						goto l271
					}
					position++
					if buffer[position] != rune('a') {
						goto l271
					}
					position++
					goto l256
				l271:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('Q') {
						goto l272
					}
					position++
					if buffer[position] != rune('u') {
						goto l272
					}
					position++
					goto l256
				l272:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('A') {
						goto l273
					}
					position++
					if buffer[position] != rune('s') {
						goto l273
					}
					position++
					goto l256
				l273:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('B') {
						goto l254
					}
					position++
					if buffer[position] != rune('a') {
						goto l254
					}
					position++
				}
			l256:
				add(ruleTwoLetterGenus, position255)
			}
			return true
		l254:
			position, tokenIndex = position254, tokenIndex254
			return false
		},
		/* 38 Word <- <(!(AuthorPrefix / RankUninomial / Approximation / Word4) (WordApostr / WordStartsWithDigit / Word2 / Word1) &(SpaceCharEOI / '('))> */
		func() bool {
			position274, tokenIndex274 := position, tokenIndex
			{
				position275 := position
				{
					position276, tokenIndex276 := position, tokenIndex
					{
						position277, tokenIndex277 := position, tokenIndex
						if !_rules[ruleAuthorPrefix]() {
							goto l278
						}
						goto l277
					l278:
						position, tokenIndex = position277, tokenIndex277
						if !_rules[ruleRankUninomial]() {
							goto l279
						}
						goto l277
					l279:
						position, tokenIndex = position277, tokenIndex277
						if !_rules[ruleApproximation]() {
							goto l280
						}
						goto l277
					l280:
						position, tokenIndex = position277, tokenIndex277
						if !_rules[ruleWord4]() {
							goto l276
						}
					}
				l277:
					goto l274
				l276:
					position, tokenIndex = position276, tokenIndex276
				}
				{
					position281, tokenIndex281 := position, tokenIndex
					if !_rules[ruleWordApostr]() {
						goto l282
					}
					goto l281
				l282:
					position, tokenIndex = position281, tokenIndex281
					if !_rules[ruleWordStartsWithDigit]() {
						goto l283
					}
					goto l281
				l283:
					position, tokenIndex = position281, tokenIndex281
					if !_rules[ruleWord2]() {
						goto l284
					}
					goto l281
				l284:
					position, tokenIndex = position281, tokenIndex281
					if !_rules[ruleWord1]() {
						goto l274
					}
				}
			l281:
				{
					position285, tokenIndex285 := position, tokenIndex
					{
						position286, tokenIndex286 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l287
						}
						goto l286
					l287:
						position, tokenIndex = position286, tokenIndex286
						if buffer[position] != rune('(') {
							goto l274
						}
						position++
					}
				l286:
					position, tokenIndex = position285, tokenIndex285
				}
				add(ruleWord, position275)
			}
			return true
		l274:
			position, tokenIndex = position274, tokenIndex274
			return false
		},
		/* 39 Word1 <- <((lASCII dash)? NameLowerChar NameLowerChar+)> */
		func() bool {
			position288, tokenIndex288 := position, tokenIndex
			{
				position289 := position
				{
					position290, tokenIndex290 := position, tokenIndex
					if !_rules[rulelASCII]() {
						goto l290
					}
					if !_rules[ruledash]() {
						goto l290
					}
					goto l291
				l290:
					position, tokenIndex = position290, tokenIndex290
				}
			l291:
				if !_rules[ruleNameLowerChar]() {
					goto l288
				}
				if !_rules[ruleNameLowerChar]() {
					goto l288
				}
			l292:
				{
					position293, tokenIndex293 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l293
					}
					goto l292
				l293:
					position, tokenIndex = position293, tokenIndex293
				}
				add(ruleWord1, position289)
			}
			return true
		l288:
			position, tokenIndex = position288, tokenIndex288
			return false
		},
		/* 40 WordStartsWithDigit <- <(('1' / '2' / '3' / '4' / '5' / '6' / '7' / '8' / '9') nums? ('.' / dash)? NameLowerChar NameLowerChar NameLowerChar NameLowerChar+)> */
		func() bool {
			position294, tokenIndex294 := position, tokenIndex
			{
				position295 := position
				{
					position296, tokenIndex296 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l297
					}
					position++
					goto l296
				l297:
					position, tokenIndex = position296, tokenIndex296
					if buffer[position] != rune('2') {
						goto l298
					}
					position++
					goto l296
				l298:
					position, tokenIndex = position296, tokenIndex296
					if buffer[position] != rune('3') {
						goto l299
					}
					position++
					goto l296
				l299:
					position, tokenIndex = position296, tokenIndex296
					if buffer[position] != rune('4') {
						goto l300
					}
					position++
					goto l296
				l300:
					position, tokenIndex = position296, tokenIndex296
					if buffer[position] != rune('5') {
						goto l301
					}
					position++
					goto l296
				l301:
					position, tokenIndex = position296, tokenIndex296
					if buffer[position] != rune('6') {
						goto l302
					}
					position++
					goto l296
				l302:
					position, tokenIndex = position296, tokenIndex296
					if buffer[position] != rune('7') {
						goto l303
					}
					position++
					goto l296
				l303:
					position, tokenIndex = position296, tokenIndex296
					if buffer[position] != rune('8') {
						goto l304
					}
					position++
					goto l296
				l304:
					position, tokenIndex = position296, tokenIndex296
					if buffer[position] != rune('9') {
						goto l294
					}
					position++
				}
			l296:
				{
					position305, tokenIndex305 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l305
					}
					goto l306
				l305:
					position, tokenIndex = position305, tokenIndex305
				}
			l306:
				{
					position307, tokenIndex307 := position, tokenIndex
					{
						position309, tokenIndex309 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l310
						}
						position++
						goto l309
					l310:
						position, tokenIndex = position309, tokenIndex309
						if !_rules[ruledash]() {
							goto l307
						}
					}
				l309:
					goto l308
				l307:
					position, tokenIndex = position307, tokenIndex307
				}
			l308:
				if !_rules[ruleNameLowerChar]() {
					goto l294
				}
				if !_rules[ruleNameLowerChar]() {
					goto l294
				}
				if !_rules[ruleNameLowerChar]() {
					goto l294
				}
				if !_rules[ruleNameLowerChar]() {
					goto l294
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
				add(ruleWordStartsWithDigit, position295)
			}
			return true
		l294:
			position, tokenIndex = position294, tokenIndex294
			return false
		},
		/* 41 Word2 <- <(NameLowerChar+ dash? NameLowerChar+)> */
		func() bool {
			position313, tokenIndex313 := position, tokenIndex
			{
				position314 := position
				if !_rules[ruleNameLowerChar]() {
					goto l313
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
				{
					position317, tokenIndex317 := position, tokenIndex
					if !_rules[ruledash]() {
						goto l317
					}
					goto l318
				l317:
					position, tokenIndex = position317, tokenIndex317
				}
			l318:
				if !_rules[ruleNameLowerChar]() {
					goto l313
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
				add(ruleWord2, position314)
			}
			return true
		l313:
			position, tokenIndex = position313, tokenIndex313
			return false
		},
		/* 42 WordApostr <- <(NameLowerChar NameLowerChar* apostr Word1)> */
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
				if !_rules[ruleapostr]() {
					goto l321
				}
				if !_rules[ruleWord1]() {
					goto l321
				}
				add(ruleWordApostr, position322)
			}
			return true
		l321:
			position, tokenIndex = position321, tokenIndex321
			return false
		},
		/* 43 Word4 <- <(NameLowerChar+ '.' NameLowerChar)> */
		func() bool {
			position325, tokenIndex325 := position, tokenIndex
			{
				position326 := position
				if !_rules[ruleNameLowerChar]() {
					goto l325
				}
			l327:
				{
					position328, tokenIndex328 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l328
					}
					goto l327
				l328:
					position, tokenIndex = position328, tokenIndex328
				}
				if buffer[position] != rune('.') {
					goto l325
				}
				position++
				if !_rules[ruleNameLowerChar]() {
					goto l325
				}
				add(ruleWord4, position326)
			}
			return true
		l325:
			position, tokenIndex = position325, tokenIndex325
			return false
		},
		/* 44 HybridChar <- <'×'> */
		func() bool {
			position329, tokenIndex329 := position, tokenIndex
			{
				position330 := position
				if buffer[position] != rune('×') {
					goto l329
				}
				position++
				add(ruleHybridChar, position330)
			}
			return true
		l329:
			position, tokenIndex = position329, tokenIndex329
			return false
		},
		/* 45 ApproxNameIgnored <- <.*> */
		func() bool {
			{
				position332 := position
			l333:
				{
					position334, tokenIndex334 := position, tokenIndex
					if !matchDot() {
						goto l334
					}
					goto l333
				l334:
					position, tokenIndex = position334, tokenIndex334
				}
				add(ruleApproxNameIgnored, position332)
			}
			return true
		},
		/* 46 Approximation <- <(('s' 'p' '.' _? ('n' 'r' '.')) / ('s' 'p' '.' _? ('a' 'f' 'f' '.')) / ('m' 'o' 'n' 's' 't' '.') / '?' / ((('s' 'p' 'p') / ('n' 'r') / ('s' 'p') / ('a' 'f' 'f') / ('s' 'p' 'e' 'c' 'i' 'e' 's')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position335, tokenIndex335 := position, tokenIndex
			{
				position336 := position
				{
					position337, tokenIndex337 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l338
					}
					position++
					if buffer[position] != rune('p') {
						goto l338
					}
					position++
					if buffer[position] != rune('.') {
						goto l338
					}
					position++
					{
						position339, tokenIndex339 := position, tokenIndex
						if !_rules[rule_]() {
							goto l339
						}
						goto l340
					l339:
						position, tokenIndex = position339, tokenIndex339
					}
				l340:
					if buffer[position] != rune('n') {
						goto l338
					}
					position++
					if buffer[position] != rune('r') {
						goto l338
					}
					position++
					if buffer[position] != rune('.') {
						goto l338
					}
					position++
					goto l337
				l338:
					position, tokenIndex = position337, tokenIndex337
					if buffer[position] != rune('s') {
						goto l341
					}
					position++
					if buffer[position] != rune('p') {
						goto l341
					}
					position++
					if buffer[position] != rune('.') {
						goto l341
					}
					position++
					{
						position342, tokenIndex342 := position, tokenIndex
						if !_rules[rule_]() {
							goto l342
						}
						goto l343
					l342:
						position, tokenIndex = position342, tokenIndex342
					}
				l343:
					if buffer[position] != rune('a') {
						goto l341
					}
					position++
					if buffer[position] != rune('f') {
						goto l341
					}
					position++
					if buffer[position] != rune('f') {
						goto l341
					}
					position++
					if buffer[position] != rune('.') {
						goto l341
					}
					position++
					goto l337
				l341:
					position, tokenIndex = position337, tokenIndex337
					if buffer[position] != rune('m') {
						goto l344
					}
					position++
					if buffer[position] != rune('o') {
						goto l344
					}
					position++
					if buffer[position] != rune('n') {
						goto l344
					}
					position++
					if buffer[position] != rune('s') {
						goto l344
					}
					position++
					if buffer[position] != rune('t') {
						goto l344
					}
					position++
					if buffer[position] != rune('.') {
						goto l344
					}
					position++
					goto l337
				l344:
					position, tokenIndex = position337, tokenIndex337
					if buffer[position] != rune('?') {
						goto l345
					}
					position++
					goto l337
				l345:
					position, tokenIndex = position337, tokenIndex337
					{
						position346, tokenIndex346 := position, tokenIndex
						if buffer[position] != rune('s') {
							goto l347
						}
						position++
						if buffer[position] != rune('p') {
							goto l347
						}
						position++
						if buffer[position] != rune('p') {
							goto l347
						}
						position++
						goto l346
					l347:
						position, tokenIndex = position346, tokenIndex346
						if buffer[position] != rune('n') {
							goto l348
						}
						position++
						if buffer[position] != rune('r') {
							goto l348
						}
						position++
						goto l346
					l348:
						position, tokenIndex = position346, tokenIndex346
						if buffer[position] != rune('s') {
							goto l349
						}
						position++
						if buffer[position] != rune('p') {
							goto l349
						}
						position++
						goto l346
					l349:
						position, tokenIndex = position346, tokenIndex346
						if buffer[position] != rune('a') {
							goto l350
						}
						position++
						if buffer[position] != rune('f') {
							goto l350
						}
						position++
						if buffer[position] != rune('f') {
							goto l350
						}
						position++
						goto l346
					l350:
						position, tokenIndex = position346, tokenIndex346
						if buffer[position] != rune('s') {
							goto l335
						}
						position++
						if buffer[position] != rune('p') {
							goto l335
						}
						position++
						if buffer[position] != rune('e') {
							goto l335
						}
						position++
						if buffer[position] != rune('c') {
							goto l335
						}
						position++
						if buffer[position] != rune('i') {
							goto l335
						}
						position++
						if buffer[position] != rune('e') {
							goto l335
						}
						position++
						if buffer[position] != rune('s') {
							goto l335
						}
						position++
					}
				l346:
					{
						position351, tokenIndex351 := position, tokenIndex
						{
							position353, tokenIndex353 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l352
							}
							position, tokenIndex = position353, tokenIndex353
						}
						goto l351
					l352:
						position, tokenIndex = position351, tokenIndex351
						if buffer[position] != rune('.') {
							goto l335
						}
						position++
					}
				l351:
				}
			l337:
				add(ruleApproximation, position336)
			}
			return true
		l335:
			position, tokenIndex = position335, tokenIndex335
			return false
		},
		/* 47 Authorship <- <((AuthorshipCombo / OriginalAuthorship) &(SpaceCharEOI / ('\\' / '(' / ',' / ':')))> */
		func() bool {
			position354, tokenIndex354 := position, tokenIndex
			{
				position355 := position
				{
					position356, tokenIndex356 := position, tokenIndex
					if !_rules[ruleAuthorshipCombo]() {
						goto l357
					}
					goto l356
				l357:
					position, tokenIndex = position356, tokenIndex356
					if !_rules[ruleOriginalAuthorship]() {
						goto l354
					}
				}
			l356:
				{
					position358, tokenIndex358 := position, tokenIndex
					{
						position359, tokenIndex359 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l360
						}
						goto l359
					l360:
						position, tokenIndex = position359, tokenIndex359
						{
							position361, tokenIndex361 := position, tokenIndex
							if buffer[position] != rune('\\') {
								goto l362
							}
							position++
							goto l361
						l362:
							position, tokenIndex = position361, tokenIndex361
							if buffer[position] != rune('(') {
								goto l363
							}
							position++
							goto l361
						l363:
							position, tokenIndex = position361, tokenIndex361
							if buffer[position] != rune(',') {
								goto l364
							}
							position++
							goto l361
						l364:
							position, tokenIndex = position361, tokenIndex361
							if buffer[position] != rune(':') {
								goto l354
							}
							position++
						}
					l361:
					}
				l359:
					position, tokenIndex = position358, tokenIndex358
				}
				add(ruleAuthorship, position355)
			}
			return true
		l354:
			position, tokenIndex = position354, tokenIndex354
			return false
		},
		/* 48 AuthorshipCombo <- <(OriginalAuthorship _? CombinationAuthorship)> */
		func() bool {
			position365, tokenIndex365 := position, tokenIndex
			{
				position366 := position
				if !_rules[ruleOriginalAuthorship]() {
					goto l365
				}
				{
					position367, tokenIndex367 := position, tokenIndex
					if !_rules[rule_]() {
						goto l367
					}
					goto l368
				l367:
					position, tokenIndex = position367, tokenIndex367
				}
			l368:
				if !_rules[ruleCombinationAuthorship]() {
					goto l365
				}
				add(ruleAuthorshipCombo, position366)
			}
			return true
		l365:
			position, tokenIndex = position365, tokenIndex365
			return false
		},
		/* 49 OriginalAuthorship <- <(BasionymAuthorshipYearMisformed / AuthorsGroup / BasionymAuthorship)> */
		func() bool {
			position369, tokenIndex369 := position, tokenIndex
			{
				position370 := position
				{
					position371, tokenIndex371 := position, tokenIndex
					if !_rules[ruleBasionymAuthorshipYearMisformed]() {
						goto l372
					}
					goto l371
				l372:
					position, tokenIndex = position371, tokenIndex371
					if !_rules[ruleAuthorsGroup]() {
						goto l373
					}
					goto l371
				l373:
					position, tokenIndex = position371, tokenIndex371
					if !_rules[ruleBasionymAuthorship]() {
						goto l369
					}
				}
			l371:
				add(ruleOriginalAuthorship, position370)
			}
			return true
		l369:
			position, tokenIndex = position369, tokenIndex369
			return false
		},
		/* 50 CombinationAuthorship <- <AuthorsGroup> */
		func() bool {
			position374, tokenIndex374 := position, tokenIndex
			{
				position375 := position
				if !_rules[ruleAuthorsGroup]() {
					goto l374
				}
				add(ruleCombinationAuthorship, position375)
			}
			return true
		l374:
			position, tokenIndex = position374, tokenIndex374
			return false
		},
		/* 51 BasionymAuthorshipYearMisformed <- <('(' _? AuthorsGroup _? ')' (_? ',')? _? Year)> */
		func() bool {
			position376, tokenIndex376 := position, tokenIndex
			{
				position377 := position
				if buffer[position] != rune('(') {
					goto l376
				}
				position++
				{
					position378, tokenIndex378 := position, tokenIndex
					if !_rules[rule_]() {
						goto l378
					}
					goto l379
				l378:
					position, tokenIndex = position378, tokenIndex378
				}
			l379:
				if !_rules[ruleAuthorsGroup]() {
					goto l376
				}
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
				if buffer[position] != rune(')') {
					goto l376
				}
				position++
				{
					position382, tokenIndex382 := position, tokenIndex
					{
						position384, tokenIndex384 := position, tokenIndex
						if !_rules[rule_]() {
							goto l384
						}
						goto l385
					l384:
						position, tokenIndex = position384, tokenIndex384
					}
				l385:
					if buffer[position] != rune(',') {
						goto l382
					}
					position++
					goto l383
				l382:
					position, tokenIndex = position382, tokenIndex382
				}
			l383:
				{
					position386, tokenIndex386 := position, tokenIndex
					if !_rules[rule_]() {
						goto l386
					}
					goto l387
				l386:
					position, tokenIndex = position386, tokenIndex386
				}
			l387:
				if !_rules[ruleYear]() {
					goto l376
				}
				add(ruleBasionymAuthorshipYearMisformed, position377)
			}
			return true
		l376:
			position, tokenIndex = position376, tokenIndex376
			return false
		},
		/* 52 BasionymAuthorship <- <(BasionymAuthorship1 / BasionymAuthorship2)> */
		func() bool {
			position388, tokenIndex388 := position, tokenIndex
			{
				position389 := position
				{
					position390, tokenIndex390 := position, tokenIndex
					if !_rules[ruleBasionymAuthorship1]() {
						goto l391
					}
					goto l390
				l391:
					position, tokenIndex = position390, tokenIndex390
					if !_rules[ruleBasionymAuthorship2]() {
						goto l388
					}
				}
			l390:
				add(ruleBasionymAuthorship, position389)
			}
			return true
		l388:
			position, tokenIndex = position388, tokenIndex388
			return false
		},
		/* 53 BasionymAuthorship1 <- <('(' _? AuthorsGroup _? ')')> */
		func() bool {
			position392, tokenIndex392 := position, tokenIndex
			{
				position393 := position
				if buffer[position] != rune('(') {
					goto l392
				}
				position++
				{
					position394, tokenIndex394 := position, tokenIndex
					if !_rules[rule_]() {
						goto l394
					}
					goto l395
				l394:
					position, tokenIndex = position394, tokenIndex394
				}
			l395:
				if !_rules[ruleAuthorsGroup]() {
					goto l392
				}
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
				if buffer[position] != rune(')') {
					goto l392
				}
				position++
				add(ruleBasionymAuthorship1, position393)
			}
			return true
		l392:
			position, tokenIndex = position392, tokenIndex392
			return false
		},
		/* 54 BasionymAuthorship2 <- <('(' _? '(' _? AuthorsGroup _? ')' _? ')')> */
		func() bool {
			position398, tokenIndex398 := position, tokenIndex
			{
				position399 := position
				if buffer[position] != rune('(') {
					goto l398
				}
				position++
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
				if buffer[position] != rune('(') {
					goto l398
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
				if !_rules[ruleAuthorsGroup]() {
					goto l398
				}
				{
					position404, tokenIndex404 := position, tokenIndex
					if !_rules[rule_]() {
						goto l404
					}
					goto l405
				l404:
					position, tokenIndex = position404, tokenIndex404
				}
			l405:
				if buffer[position] != rune(')') {
					goto l398
				}
				position++
				{
					position406, tokenIndex406 := position, tokenIndex
					if !_rules[rule_]() {
						goto l406
					}
					goto l407
				l406:
					position, tokenIndex = position406, tokenIndex406
				}
			l407:
				if buffer[position] != rune(')') {
					goto l398
				}
				position++
				add(ruleBasionymAuthorship2, position399)
			}
			return true
		l398:
			position, tokenIndex = position398, tokenIndex398
			return false
		},
		/* 55 AuthorsGroup <- <(AuthorsTeam (_? AuthorEmend? AuthorEx? AuthorsTeam)?)> */
		func() bool {
			position408, tokenIndex408 := position, tokenIndex
			{
				position409 := position
				if !_rules[ruleAuthorsTeam]() {
					goto l408
				}
				{
					position410, tokenIndex410 := position, tokenIndex
					{
						position412, tokenIndex412 := position, tokenIndex
						if !_rules[rule_]() {
							goto l412
						}
						goto l413
					l412:
						position, tokenIndex = position412, tokenIndex412
					}
				l413:
					{
						position414, tokenIndex414 := position, tokenIndex
						if !_rules[ruleAuthorEmend]() {
							goto l414
						}
						goto l415
					l414:
						position, tokenIndex = position414, tokenIndex414
					}
				l415:
					{
						position416, tokenIndex416 := position, tokenIndex
						if !_rules[ruleAuthorEx]() {
							goto l416
						}
						goto l417
					l416:
						position, tokenIndex = position416, tokenIndex416
					}
				l417:
					if !_rules[ruleAuthorsTeam]() {
						goto l410
					}
					goto l411
				l410:
					position, tokenIndex = position410, tokenIndex410
				}
			l411:
				add(ruleAuthorsGroup, position409)
			}
			return true
		l408:
			position, tokenIndex = position408, tokenIndex408
			return false
		},
		/* 56 AuthorsTeam <- <(Author (AuthorSep Author)* (_? ','? _? Year)?)> */
		func() bool {
			position418, tokenIndex418 := position, tokenIndex
			{
				position419 := position
				if !_rules[ruleAuthor]() {
					goto l418
				}
			l420:
				{
					position421, tokenIndex421 := position, tokenIndex
					if !_rules[ruleAuthorSep]() {
						goto l421
					}
					if !_rules[ruleAuthor]() {
						goto l421
					}
					goto l420
				l421:
					position, tokenIndex = position421, tokenIndex421
				}
				{
					position422, tokenIndex422 := position, tokenIndex
					{
						position424, tokenIndex424 := position, tokenIndex
						if !_rules[rule_]() {
							goto l424
						}
						goto l425
					l424:
						position, tokenIndex = position424, tokenIndex424
					}
				l425:
					{
						position426, tokenIndex426 := position, tokenIndex
						if buffer[position] != rune(',') {
							goto l426
						}
						position++
						goto l427
					l426:
						position, tokenIndex = position426, tokenIndex426
					}
				l427:
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
					if !_rules[ruleYear]() {
						goto l422
					}
					goto l423
				l422:
					position, tokenIndex = position422, tokenIndex422
				}
			l423:
				add(ruleAuthorsTeam, position419)
			}
			return true
		l418:
			position, tokenIndex = position418, tokenIndex418
			return false
		},
		/* 57 AuthorSep <- <(AuthorSep1 / AuthorSep2)> */
		func() bool {
			position430, tokenIndex430 := position, tokenIndex
			{
				position431 := position
				{
					position432, tokenIndex432 := position, tokenIndex
					if !_rules[ruleAuthorSep1]() {
						goto l433
					}
					goto l432
				l433:
					position, tokenIndex = position432, tokenIndex432
					if !_rules[ruleAuthorSep2]() {
						goto l430
					}
				}
			l432:
				add(ruleAuthorSep, position431)
			}
			return true
		l430:
			position, tokenIndex = position430, tokenIndex430
			return false
		},
		/* 58 AuthorSep1 <- <(_? (',' _)? ('&' / ('e' 't') / ('a' 'n' 'd') / ('a' 'p' 'u' 'd')) _?)> */
		func() bool {
			position434, tokenIndex434 := position, tokenIndex
			{
				position435 := position
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
				{
					position438, tokenIndex438 := position, tokenIndex
					if buffer[position] != rune(',') {
						goto l438
					}
					position++
					if !_rules[rule_]() {
						goto l438
					}
					goto l439
				l438:
					position, tokenIndex = position438, tokenIndex438
				}
			l439:
				{
					position440, tokenIndex440 := position, tokenIndex
					if buffer[position] != rune('&') {
						goto l441
					}
					position++
					goto l440
				l441:
					position, tokenIndex = position440, tokenIndex440
					if buffer[position] != rune('e') {
						goto l442
					}
					position++
					if buffer[position] != rune('t') {
						goto l442
					}
					position++
					goto l440
				l442:
					position, tokenIndex = position440, tokenIndex440
					if buffer[position] != rune('a') {
						goto l443
					}
					position++
					if buffer[position] != rune('n') {
						goto l443
					}
					position++
					if buffer[position] != rune('d') {
						goto l443
					}
					position++
					goto l440
				l443:
					position, tokenIndex = position440, tokenIndex440
					if buffer[position] != rune('a') {
						goto l434
					}
					position++
					if buffer[position] != rune('p') {
						goto l434
					}
					position++
					if buffer[position] != rune('u') {
						goto l434
					}
					position++
					if buffer[position] != rune('d') {
						goto l434
					}
					position++
				}
			l440:
				{
					position444, tokenIndex444 := position, tokenIndex
					if !_rules[rule_]() {
						goto l444
					}
					goto l445
				l444:
					position, tokenIndex = position444, tokenIndex444
				}
			l445:
				add(ruleAuthorSep1, position435)
			}
			return true
		l434:
			position, tokenIndex = position434, tokenIndex434
			return false
		},
		/* 59 AuthorSep2 <- <(_? ',' _?)> */
		func() bool {
			position446, tokenIndex446 := position, tokenIndex
			{
				position447 := position
				{
					position448, tokenIndex448 := position, tokenIndex
					if !_rules[rule_]() {
						goto l448
					}
					goto l449
				l448:
					position, tokenIndex = position448, tokenIndex448
				}
			l449:
				if buffer[position] != rune(',') {
					goto l446
				}
				position++
				{
					position450, tokenIndex450 := position, tokenIndex
					if !_rules[rule_]() {
						goto l450
					}
					goto l451
				l450:
					position, tokenIndex = position450, tokenIndex450
				}
			l451:
				add(ruleAuthorSep2, position447)
			}
			return true
		l446:
			position, tokenIndex = position446, tokenIndex446
			return false
		},
		/* 60 AuthorEx <- <((('e' 'x' '.'?) / ('i' 'n')) _)> */
		func() bool {
			position452, tokenIndex452 := position, tokenIndex
			{
				position453 := position
				{
					position454, tokenIndex454 := position, tokenIndex
					if buffer[position] != rune('e') {
						goto l455
					}
					position++
					if buffer[position] != rune('x') {
						goto l455
					}
					position++
					{
						position456, tokenIndex456 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l456
						}
						position++
						goto l457
					l456:
						position, tokenIndex = position456, tokenIndex456
					}
				l457:
					goto l454
				l455:
					position, tokenIndex = position454, tokenIndex454
					if buffer[position] != rune('i') {
						goto l452
					}
					position++
					if buffer[position] != rune('n') {
						goto l452
					}
					position++
				}
			l454:
				if !_rules[rule_]() {
					goto l452
				}
				add(ruleAuthorEx, position453)
			}
			return true
		l452:
			position, tokenIndex = position452, tokenIndex452
			return false
		},
		/* 61 AuthorEmend <- <('e' 'm' 'e' 'n' 'd' '.'? _)> */
		func() bool {
			position458, tokenIndex458 := position, tokenIndex
			{
				position459 := position
				if buffer[position] != rune('e') {
					goto l458
				}
				position++
				if buffer[position] != rune('m') {
					goto l458
				}
				position++
				if buffer[position] != rune('e') {
					goto l458
				}
				position++
				if buffer[position] != rune('n') {
					goto l458
				}
				position++
				if buffer[position] != rune('d') {
					goto l458
				}
				position++
				{
					position460, tokenIndex460 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l460
					}
					position++
					goto l461
				l460:
					position, tokenIndex = position460, tokenIndex460
				}
			l461:
				if !_rules[rule_]() {
					goto l458
				}
				add(ruleAuthorEmend, position459)
			}
			return true
		l458:
			position, tokenIndex = position458, tokenIndex458
			return false
		},
		/* 62 Author <- <(Author1 / Author2 / UnknownAuthor)> */
		func() bool {
			position462, tokenIndex462 := position, tokenIndex
			{
				position463 := position
				{
					position464, tokenIndex464 := position, tokenIndex
					if !_rules[ruleAuthor1]() {
						goto l465
					}
					goto l464
				l465:
					position, tokenIndex = position464, tokenIndex464
					if !_rules[ruleAuthor2]() {
						goto l466
					}
					goto l464
				l466:
					position, tokenIndex = position464, tokenIndex464
					if !_rules[ruleUnknownAuthor]() {
						goto l462
					}
				}
			l464:
				add(ruleAuthor, position463)
			}
			return true
		l462:
			position, tokenIndex = position462, tokenIndex462
			return false
		},
		/* 63 Author1 <- <(Author2 _? Filius)> */
		func() bool {
			position467, tokenIndex467 := position, tokenIndex
			{
				position468 := position
				if !_rules[ruleAuthor2]() {
					goto l467
				}
				{
					position469, tokenIndex469 := position, tokenIndex
					if !_rules[rule_]() {
						goto l469
					}
					goto l470
				l469:
					position, tokenIndex = position469, tokenIndex469
				}
			l470:
				if !_rules[ruleFilius]() {
					goto l467
				}
				add(ruleAuthor1, position468)
			}
			return true
		l467:
			position, tokenIndex = position467, tokenIndex467
			return false
		},
		/* 64 Author2 <- <(AuthorWord (_? AuthorWord)*)> */
		func() bool {
			position471, tokenIndex471 := position, tokenIndex
			{
				position472 := position
				if !_rules[ruleAuthorWord]() {
					goto l471
				}
			l473:
				{
					position474, tokenIndex474 := position, tokenIndex
					{
						position475, tokenIndex475 := position, tokenIndex
						if !_rules[rule_]() {
							goto l475
						}
						goto l476
					l475:
						position, tokenIndex = position475, tokenIndex475
					}
				l476:
					if !_rules[ruleAuthorWord]() {
						goto l474
					}
					goto l473
				l474:
					position, tokenIndex = position474, tokenIndex474
				}
				add(ruleAuthor2, position472)
			}
			return true
		l471:
			position, tokenIndex = position471, tokenIndex471
			return false
		},
		/* 65 UnknownAuthor <- <('?' / ((('a' 'u' 'c' 't') / ('a' 'n' 'o' 'n')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position477, tokenIndex477 := position, tokenIndex
			{
				position478 := position
				{
					position479, tokenIndex479 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l480
					}
					position++
					goto l479
				l480:
					position, tokenIndex = position479, tokenIndex479
					{
						position481, tokenIndex481 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l482
						}
						position++
						if buffer[position] != rune('u') {
							goto l482
						}
						position++
						if buffer[position] != rune('c') {
							goto l482
						}
						position++
						if buffer[position] != rune('t') {
							goto l482
						}
						position++
						goto l481
					l482:
						position, tokenIndex = position481, tokenIndex481
						if buffer[position] != rune('a') {
							goto l477
						}
						position++
						if buffer[position] != rune('n') {
							goto l477
						}
						position++
						if buffer[position] != rune('o') {
							goto l477
						}
						position++
						if buffer[position] != rune('n') {
							goto l477
						}
						position++
					}
				l481:
					{
						position483, tokenIndex483 := position, tokenIndex
						{
							position485, tokenIndex485 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l484
							}
							position, tokenIndex = position485, tokenIndex485
						}
						goto l483
					l484:
						position, tokenIndex = position483, tokenIndex483
						if buffer[position] != rune('.') {
							goto l477
						}
						position++
					}
				l483:
				}
			l479:
				add(ruleUnknownAuthor, position478)
			}
			return true
		l477:
			position, tokenIndex = position477, tokenIndex477
			return false
		},
		/* 66 AuthorWord <- <(AuthorWord1 / AuthorWord2 / AuthorWord3 / AuthorPrefix)> */
		func() bool {
			position486, tokenIndex486 := position, tokenIndex
			{
				position487 := position
				{
					position488, tokenIndex488 := position, tokenIndex
					if !_rules[ruleAuthorWord1]() {
						goto l489
					}
					goto l488
				l489:
					position, tokenIndex = position488, tokenIndex488
					if !_rules[ruleAuthorWord2]() {
						goto l490
					}
					goto l488
				l490:
					position, tokenIndex = position488, tokenIndex488
					if !_rules[ruleAuthorWord3]() {
						goto l491
					}
					goto l488
				l491:
					position, tokenIndex = position488, tokenIndex488
					if !_rules[ruleAuthorPrefix]() {
						goto l486
					}
				}
			l488:
				add(ruleAuthorWord, position487)
			}
			return true
		l486:
			position, tokenIndex = position486, tokenIndex486
			return false
		},
		/* 67 AuthorWord1 <- <(('a' 'r' 'g' '.') / ('e' 't' ' ' 'a' 'l' '.' '{' '?' '}') / ((('e' 't') / '&') (' ' 'a' 'l') '.'?))> */
		func() bool {
			position492, tokenIndex492 := position, tokenIndex
			{
				position493 := position
				{
					position494, tokenIndex494 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l495
					}
					position++
					if buffer[position] != rune('r') {
						goto l495
					}
					position++
					if buffer[position] != rune('g') {
						goto l495
					}
					position++
					if buffer[position] != rune('.') {
						goto l495
					}
					position++
					goto l494
				l495:
					position, tokenIndex = position494, tokenIndex494
					if buffer[position] != rune('e') {
						goto l496
					}
					position++
					if buffer[position] != rune('t') {
						goto l496
					}
					position++
					if buffer[position] != rune(' ') {
						goto l496
					}
					position++
					if buffer[position] != rune('a') {
						goto l496
					}
					position++
					if buffer[position] != rune('l') {
						goto l496
					}
					position++
					if buffer[position] != rune('.') {
						goto l496
					}
					position++
					if buffer[position] != rune('{') {
						goto l496
					}
					position++
					if buffer[position] != rune('?') {
						goto l496
					}
					position++
					if buffer[position] != rune('}') {
						goto l496
					}
					position++
					goto l494
				l496:
					position, tokenIndex = position494, tokenIndex494
					{
						position497, tokenIndex497 := position, tokenIndex
						if buffer[position] != rune('e') {
							goto l498
						}
						position++
						if buffer[position] != rune('t') {
							goto l498
						}
						position++
						goto l497
					l498:
						position, tokenIndex = position497, tokenIndex497
						if buffer[position] != rune('&') {
							goto l492
						}
						position++
					}
				l497:
					if buffer[position] != rune(' ') {
						goto l492
					}
					position++
					if buffer[position] != rune('a') {
						goto l492
					}
					position++
					if buffer[position] != rune('l') {
						goto l492
					}
					position++
					{
						position499, tokenIndex499 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l499
						}
						position++
						goto l500
					l499:
						position, tokenIndex = position499, tokenIndex499
					}
				l500:
				}
			l494:
				add(ruleAuthorWord1, position493)
			}
			return true
		l492:
			position, tokenIndex = position492, tokenIndex492
			return false
		},
		/* 68 AuthorWord2 <- <(AuthorWord3 dash AuthorWordSoft)> */
		func() bool {
			position501, tokenIndex501 := position, tokenIndex
			{
				position502 := position
				if !_rules[ruleAuthorWord3]() {
					goto l501
				}
				if !_rules[ruledash]() {
					goto l501
				}
				if !_rules[ruleAuthorWordSoft]() {
					goto l501
				}
				add(ruleAuthorWord2, position502)
			}
			return true
		l501:
			position, tokenIndex = position501, tokenIndex501
			return false
		},
		/* 69 AuthorWord3 <- <(AuthorPrefixGlued? (AllCapsAuthorWord / CapAuthorWord) '.'?)> */
		func() bool {
			position503, tokenIndex503 := position, tokenIndex
			{
				position504 := position
				{
					position505, tokenIndex505 := position, tokenIndex
					if !_rules[ruleAuthorPrefixGlued]() {
						goto l505
					}
					goto l506
				l505:
					position, tokenIndex = position505, tokenIndex505
				}
			l506:
				{
					position507, tokenIndex507 := position, tokenIndex
					if !_rules[ruleAllCapsAuthorWord]() {
						goto l508
					}
					goto l507
				l508:
					position, tokenIndex = position507, tokenIndex507
					if !_rules[ruleCapAuthorWord]() {
						goto l503
					}
				}
			l507:
				{
					position509, tokenIndex509 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l509
					}
					position++
					goto l510
				l509:
					position, tokenIndex = position509, tokenIndex509
				}
			l510:
				add(ruleAuthorWord3, position504)
			}
			return true
		l503:
			position, tokenIndex = position503, tokenIndex503
			return false
		},
		/* 70 AuthorWordSoft <- <(((AuthorUpperChar (AuthorUpperChar+ / AuthorLowerChar+)) / AuthorLowerChar+) '.'?)> */
		func() bool {
			position511, tokenIndex511 := position, tokenIndex
			{
				position512 := position
				{
					position513, tokenIndex513 := position, tokenIndex
					if !_rules[ruleAuthorUpperChar]() {
						goto l514
					}
					{
						position515, tokenIndex515 := position, tokenIndex
						if !_rules[ruleAuthorUpperChar]() {
							goto l516
						}
					l517:
						{
							position518, tokenIndex518 := position, tokenIndex
							if !_rules[ruleAuthorUpperChar]() {
								goto l518
							}
							goto l517
						l518:
							position, tokenIndex = position518, tokenIndex518
						}
						goto l515
					l516:
						position, tokenIndex = position515, tokenIndex515
						if !_rules[ruleAuthorLowerChar]() {
							goto l514
						}
					l519:
						{
							position520, tokenIndex520 := position, tokenIndex
							if !_rules[ruleAuthorLowerChar]() {
								goto l520
							}
							goto l519
						l520:
							position, tokenIndex = position520, tokenIndex520
						}
					}
				l515:
					goto l513
				l514:
					position, tokenIndex = position513, tokenIndex513
					if !_rules[ruleAuthorLowerChar]() {
						goto l511
					}
				l521:
					{
						position522, tokenIndex522 := position, tokenIndex
						if !_rules[ruleAuthorLowerChar]() {
							goto l522
						}
						goto l521
					l522:
						position, tokenIndex = position522, tokenIndex522
					}
				}
			l513:
				{
					position523, tokenIndex523 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l523
					}
					position++
					goto l524
				l523:
					position, tokenIndex = position523, tokenIndex523
				}
			l524:
				add(ruleAuthorWordSoft, position512)
			}
			return true
		l511:
			position, tokenIndex = position511, tokenIndex511
			return false
		},
		/* 71 CapAuthorWord <- <(AuthorUpperChar AuthorLowerChar*)> */
		func() bool {
			position525, tokenIndex525 := position, tokenIndex
			{
				position526 := position
				if !_rules[ruleAuthorUpperChar]() {
					goto l525
				}
			l527:
				{
					position528, tokenIndex528 := position, tokenIndex
					if !_rules[ruleAuthorLowerChar]() {
						goto l528
					}
					goto l527
				l528:
					position, tokenIndex = position528, tokenIndex528
				}
				add(ruleCapAuthorWord, position526)
			}
			return true
		l525:
			position, tokenIndex = position525, tokenIndex525
			return false
		},
		/* 72 AllCapsAuthorWord <- <(AuthorUpperChar AuthorUpperChar+)> */
		func() bool {
			position529, tokenIndex529 := position, tokenIndex
			{
				position530 := position
				if !_rules[ruleAuthorUpperChar]() {
					goto l529
				}
				if !_rules[ruleAuthorUpperChar]() {
					goto l529
				}
			l531:
				{
					position532, tokenIndex532 := position, tokenIndex
					if !_rules[ruleAuthorUpperChar]() {
						goto l532
					}
					goto l531
				l532:
					position, tokenIndex = position532, tokenIndex532
				}
				add(ruleAllCapsAuthorWord, position530)
			}
			return true
		l529:
			position, tokenIndex = position529, tokenIndex529
			return false
		},
		/* 73 Filius <- <(('f' '.') / ('f' 'i' 'l' '.') / ('f' 'i' 'l' 'i' 'u' 's'))> */
		func() bool {
			position533, tokenIndex533 := position, tokenIndex
			{
				position534 := position
				{
					position535, tokenIndex535 := position, tokenIndex
					if buffer[position] != rune('f') {
						goto l536
					}
					position++
					if buffer[position] != rune('.') {
						goto l536
					}
					position++
					goto l535
				l536:
					position, tokenIndex = position535, tokenIndex535
					if buffer[position] != rune('f') {
						goto l537
					}
					position++
					if buffer[position] != rune('i') {
						goto l537
					}
					position++
					if buffer[position] != rune('l') {
						goto l537
					}
					position++
					if buffer[position] != rune('.') {
						goto l537
					}
					position++
					goto l535
				l537:
					position, tokenIndex = position535, tokenIndex535
					if buffer[position] != rune('f') {
						goto l533
					}
					position++
					if buffer[position] != rune('i') {
						goto l533
					}
					position++
					if buffer[position] != rune('l') {
						goto l533
					}
					position++
					if buffer[position] != rune('i') {
						goto l533
					}
					position++
					if buffer[position] != rune('u') {
						goto l533
					}
					position++
					if buffer[position] != rune('s') {
						goto l533
					}
					position++
				}
			l535:
				add(ruleFilius, position534)
			}
			return true
		l533:
			position, tokenIndex = position533, tokenIndex533
			return false
		},
		/* 74 AuthorPrefixGlued <- <(('d' '\'') / ('O' '\''))> */
		func() bool {
			position538, tokenIndex538 := position, tokenIndex
			{
				position539 := position
				{
					position540, tokenIndex540 := position, tokenIndex
					if buffer[position] != rune('d') {
						goto l541
					}
					position++
					if buffer[position] != rune('\'') {
						goto l541
					}
					position++
					goto l540
				l541:
					position, tokenIndex = position540, tokenIndex540
					if buffer[position] != rune('O') {
						goto l538
					}
					position++
					if buffer[position] != rune('\'') {
						goto l538
					}
					position++
				}
			l540:
				add(ruleAuthorPrefixGlued, position539)
			}
			return true
		l538:
			position, tokenIndex = position538, tokenIndex538
			return false
		},
		/* 75 AuthorPrefix <- <(AuthorPrefix1 / AuthorPrefix2)> */
		func() bool {
			position542, tokenIndex542 := position, tokenIndex
			{
				position543 := position
				{
					position544, tokenIndex544 := position, tokenIndex
					if !_rules[ruleAuthorPrefix1]() {
						goto l545
					}
					goto l544
				l545:
					position, tokenIndex = position544, tokenIndex544
					if !_rules[ruleAuthorPrefix2]() {
						goto l542
					}
				}
			l544:
				add(ruleAuthorPrefix, position543)
			}
			return true
		l542:
			position, tokenIndex = position542, tokenIndex542
			return false
		},
		/* 76 AuthorPrefix2 <- <(('v' '.' (_? ('d' '.'))?) / ('\'' 't'))> */
		func() bool {
			position546, tokenIndex546 := position, tokenIndex
			{
				position547 := position
				{
					position548, tokenIndex548 := position, tokenIndex
					if buffer[position] != rune('v') {
						goto l549
					}
					position++
					if buffer[position] != rune('.') {
						goto l549
					}
					position++
					{
						position550, tokenIndex550 := position, tokenIndex
						{
							position552, tokenIndex552 := position, tokenIndex
							if !_rules[rule_]() {
								goto l552
							}
							goto l553
						l552:
							position, tokenIndex = position552, tokenIndex552
						}
					l553:
						if buffer[position] != rune('d') {
							goto l550
						}
						position++
						if buffer[position] != rune('.') {
							goto l550
						}
						position++
						goto l551
					l550:
						position, tokenIndex = position550, tokenIndex550
					}
				l551:
					goto l548
				l549:
					position, tokenIndex = position548, tokenIndex548
					if buffer[position] != rune('\'') {
						goto l546
					}
					position++
					if buffer[position] != rune('t') {
						goto l546
					}
					position++
				}
			l548:
				add(ruleAuthorPrefix2, position547)
			}
			return true
		l546:
			position, tokenIndex = position546, tokenIndex546
			return false
		},
		/* 77 AuthorPrefix1 <- <((('a' 'b') / ('a' 'f') / ('b' 'i' 's') / ('d' 'a') / ('d' 'e' 'r') / ('d' 'e' 's') / ('d' 'e' 'n') / ('d' 'e' 'l') / ('d' 'e' 'l' 'l' 'a') / ('d' 'e' 'l' 'a') / ('d' 'e') / ('d' 'i') / ('d' 'u') / ('e' 'l') / ('l' 'a') / ('l' 'e') / ('t' 'e' 'r') / ('v' 'a' 'n') / ('d' '\'') / ('i' 'n' '\'' 't') / ('z' 'u' 'r') / ('v' 'o' 'n' (_ (('d' '.') / ('d' 'e' 'm')))?) / ('v' (_ 'd')?)) &_)> */
		func() bool {
			position554, tokenIndex554 := position, tokenIndex
			{
				position555 := position
				{
					position556, tokenIndex556 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l557
					}
					position++
					if buffer[position] != rune('b') {
						goto l557
					}
					position++
					goto l556
				l557:
					position, tokenIndex = position556, tokenIndex556
					if buffer[position] != rune('a') {
						goto l558
					}
					position++
					if buffer[position] != rune('f') {
						goto l558
					}
					position++
					goto l556
				l558:
					position, tokenIndex = position556, tokenIndex556
					if buffer[position] != rune('b') {
						goto l559
					}
					position++
					if buffer[position] != rune('i') {
						goto l559
					}
					position++
					if buffer[position] != rune('s') {
						goto l559
					}
					position++
					goto l556
				l559:
					position, tokenIndex = position556, tokenIndex556
					if buffer[position] != rune('d') {
						goto l560
					}
					position++
					if buffer[position] != rune('a') {
						goto l560
					}
					position++
					goto l556
				l560:
					position, tokenIndex = position556, tokenIndex556
					if buffer[position] != rune('d') {
						goto l561
					}
					position++
					if buffer[position] != rune('e') {
						goto l561
					}
					position++
					if buffer[position] != rune('r') {
						goto l561
					}
					position++
					goto l556
				l561:
					position, tokenIndex = position556, tokenIndex556
					if buffer[position] != rune('d') {
						goto l562
					}
					position++
					if buffer[position] != rune('e') {
						goto l562
					}
					position++
					if buffer[position] != rune('s') {
						goto l562
					}
					position++
					goto l556
				l562:
					position, tokenIndex = position556, tokenIndex556
					if buffer[position] != rune('d') {
						goto l563
					}
					position++
					if buffer[position] != rune('e') {
						goto l563
					}
					position++
					if buffer[position] != rune('n') {
						goto l563
					}
					position++
					goto l556
				l563:
					position, tokenIndex = position556, tokenIndex556
					if buffer[position] != rune('d') {
						goto l564
					}
					position++
					if buffer[position] != rune('e') {
						goto l564
					}
					position++
					if buffer[position] != rune('l') {
						goto l564
					}
					position++
					goto l556
				l564:
					position, tokenIndex = position556, tokenIndex556
					if buffer[position] != rune('d') {
						goto l565
					}
					position++
					if buffer[position] != rune('e') {
						goto l565
					}
					position++
					if buffer[position] != rune('l') {
						goto l565
					}
					position++
					if buffer[position] != rune('l') {
						goto l565
					}
					position++
					if buffer[position] != rune('a') {
						goto l565
					}
					position++
					goto l556
				l565:
					position, tokenIndex = position556, tokenIndex556
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
					if buffer[position] != rune('a') {
						goto l566
					}
					position++
					goto l556
				l566:
					position, tokenIndex = position556, tokenIndex556
					if buffer[position] != rune('d') {
						goto l567
					}
					position++
					if buffer[position] != rune('e') {
						goto l567
					}
					position++
					goto l556
				l567:
					position, tokenIndex = position556, tokenIndex556
					if buffer[position] != rune('d') {
						goto l568
					}
					position++
					if buffer[position] != rune('i') {
						goto l568
					}
					position++
					goto l556
				l568:
					position, tokenIndex = position556, tokenIndex556
					if buffer[position] != rune('d') {
						goto l569
					}
					position++
					if buffer[position] != rune('u') {
						goto l569
					}
					position++
					goto l556
				l569:
					position, tokenIndex = position556, tokenIndex556
					if buffer[position] != rune('e') {
						goto l570
					}
					position++
					if buffer[position] != rune('l') {
						goto l570
					}
					position++
					goto l556
				l570:
					position, tokenIndex = position556, tokenIndex556
					if buffer[position] != rune('l') {
						goto l571
					}
					position++
					if buffer[position] != rune('a') {
						goto l571
					}
					position++
					goto l556
				l571:
					position, tokenIndex = position556, tokenIndex556
					if buffer[position] != rune('l') {
						goto l572
					}
					position++
					if buffer[position] != rune('e') {
						goto l572
					}
					position++
					goto l556
				l572:
					position, tokenIndex = position556, tokenIndex556
					if buffer[position] != rune('t') {
						goto l573
					}
					position++
					if buffer[position] != rune('e') {
						goto l573
					}
					position++
					if buffer[position] != rune('r') {
						goto l573
					}
					position++
					goto l556
				l573:
					position, tokenIndex = position556, tokenIndex556
					if buffer[position] != rune('v') {
						goto l574
					}
					position++
					if buffer[position] != rune('a') {
						goto l574
					}
					position++
					if buffer[position] != rune('n') {
						goto l574
					}
					position++
					goto l556
				l574:
					position, tokenIndex = position556, tokenIndex556
					if buffer[position] != rune('d') {
						goto l575
					}
					position++
					if buffer[position] != rune('\'') {
						goto l575
					}
					position++
					goto l556
				l575:
					position, tokenIndex = position556, tokenIndex556
					if buffer[position] != rune('i') {
						goto l576
					}
					position++
					if buffer[position] != rune('n') {
						goto l576
					}
					position++
					if buffer[position] != rune('\'') {
						goto l576
					}
					position++
					if buffer[position] != rune('t') {
						goto l576
					}
					position++
					goto l556
				l576:
					position, tokenIndex = position556, tokenIndex556
					if buffer[position] != rune('z') {
						goto l577
					}
					position++
					if buffer[position] != rune('u') {
						goto l577
					}
					position++
					if buffer[position] != rune('r') {
						goto l577
					}
					position++
					goto l556
				l577:
					position, tokenIndex = position556, tokenIndex556
					if buffer[position] != rune('v') {
						goto l578
					}
					position++
					if buffer[position] != rune('o') {
						goto l578
					}
					position++
					if buffer[position] != rune('n') {
						goto l578
					}
					position++
					{
						position579, tokenIndex579 := position, tokenIndex
						if !_rules[rule_]() {
							goto l579
						}
						{
							position581, tokenIndex581 := position, tokenIndex
							if buffer[position] != rune('d') {
								goto l582
							}
							position++
							if buffer[position] != rune('.') {
								goto l582
							}
							position++
							goto l581
						l582:
							position, tokenIndex = position581, tokenIndex581
							if buffer[position] != rune('d') {
								goto l579
							}
							position++
							if buffer[position] != rune('e') {
								goto l579
							}
							position++
							if buffer[position] != rune('m') {
								goto l579
							}
							position++
						}
					l581:
						goto l580
					l579:
						position, tokenIndex = position579, tokenIndex579
					}
				l580:
					goto l556
				l578:
					position, tokenIndex = position556, tokenIndex556
					if buffer[position] != rune('v') {
						goto l554
					}
					position++
					{
						position583, tokenIndex583 := position, tokenIndex
						if !_rules[rule_]() {
							goto l583
						}
						if buffer[position] != rune('d') {
							goto l583
						}
						position++
						goto l584
					l583:
						position, tokenIndex = position583, tokenIndex583
					}
				l584:
				}
			l556:
				{
					position585, tokenIndex585 := position, tokenIndex
					if !_rules[rule_]() {
						goto l554
					}
					position, tokenIndex = position585, tokenIndex585
				}
				add(ruleAuthorPrefix1, position555)
			}
			return true
		l554:
			position, tokenIndex = position554, tokenIndex554
			return false
		},
		/* 78 AuthorUpperChar <- <(hASCII / ('À' / 'Á' / 'Â' / 'Ã' / 'Ä' / 'Å' / 'Æ' / 'Ç' / 'È' / 'É' / 'Ê' / 'Ë' / 'Ì' / 'Í' / 'Î' / 'Ï' / 'Ð' / 'Ñ' / 'Ò' / 'Ó' / 'Ô' / 'Õ' / 'Ö' / 'Ø' / 'Ù' / 'Ú' / 'Û' / 'Ü' / 'Ý' / 'Ć' / 'Č' / 'Ď' / 'İ' / 'Ķ' / 'Ĺ' / 'ĺ' / 'Ľ' / 'ľ' / 'Ł' / 'ł' / 'Ņ' / 'Ō' / 'Ő' / 'Œ' / 'Ř' / 'Ś' / 'Ŝ' / 'Ş' / 'Š' / 'Ÿ' / 'Ź' / 'Ż' / 'Ž' / 'ƒ' / 'Ǿ' / 'Ș' / 'Ț'))> */
		func() bool {
			position586, tokenIndex586 := position, tokenIndex
			{
				position587 := position
				{
					position588, tokenIndex588 := position, tokenIndex
					if !_rules[rulehASCII]() {
						goto l589
					}
					goto l588
				l589:
					position, tokenIndex = position588, tokenIndex588
					{
						position590, tokenIndex590 := position, tokenIndex
						if buffer[position] != rune('À') {
							goto l591
						}
						position++
						goto l590
					l591:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Á') {
							goto l592
						}
						position++
						goto l590
					l592:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Â') {
							goto l593
						}
						position++
						goto l590
					l593:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ã') {
							goto l594
						}
						position++
						goto l590
					l594:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ä') {
							goto l595
						}
						position++
						goto l590
					l595:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Å') {
							goto l596
						}
						position++
						goto l590
					l596:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Æ') {
							goto l597
						}
						position++
						goto l590
					l597:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ç') {
							goto l598
						}
						position++
						goto l590
					l598:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('È') {
							goto l599
						}
						position++
						goto l590
					l599:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('É') {
							goto l600
						}
						position++
						goto l590
					l600:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ê') {
							goto l601
						}
						position++
						goto l590
					l601:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ë') {
							goto l602
						}
						position++
						goto l590
					l602:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ì') {
							goto l603
						}
						position++
						goto l590
					l603:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Í') {
							goto l604
						}
						position++
						goto l590
					l604:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Î') {
							goto l605
						}
						position++
						goto l590
					l605:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ï') {
							goto l606
						}
						position++
						goto l590
					l606:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ð') {
							goto l607
						}
						position++
						goto l590
					l607:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ñ') {
							goto l608
						}
						position++
						goto l590
					l608:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ò') {
							goto l609
						}
						position++
						goto l590
					l609:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ó') {
							goto l610
						}
						position++
						goto l590
					l610:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ô') {
							goto l611
						}
						position++
						goto l590
					l611:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Õ') {
							goto l612
						}
						position++
						goto l590
					l612:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ö') {
							goto l613
						}
						position++
						goto l590
					l613:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ø') {
							goto l614
						}
						position++
						goto l590
					l614:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ù') {
							goto l615
						}
						position++
						goto l590
					l615:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ú') {
							goto l616
						}
						position++
						goto l590
					l616:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Û') {
							goto l617
						}
						position++
						goto l590
					l617:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ü') {
							goto l618
						}
						position++
						goto l590
					l618:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ý') {
							goto l619
						}
						position++
						goto l590
					l619:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ć') {
							goto l620
						}
						position++
						goto l590
					l620:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Č') {
							goto l621
						}
						position++
						goto l590
					l621:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ď') {
							goto l622
						}
						position++
						goto l590
					l622:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('İ') {
							goto l623
						}
						position++
						goto l590
					l623:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ķ') {
							goto l624
						}
						position++
						goto l590
					l624:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ĺ') {
							goto l625
						}
						position++
						goto l590
					l625:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('ĺ') {
							goto l626
						}
						position++
						goto l590
					l626:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ľ') {
							goto l627
						}
						position++
						goto l590
					l627:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('ľ') {
							goto l628
						}
						position++
						goto l590
					l628:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ł') {
							goto l629
						}
						position++
						goto l590
					l629:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('ł') {
							goto l630
						}
						position++
						goto l590
					l630:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ņ') {
							goto l631
						}
						position++
						goto l590
					l631:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ō') {
							goto l632
						}
						position++
						goto l590
					l632:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ő') {
							goto l633
						}
						position++
						goto l590
					l633:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Œ') {
							goto l634
						}
						position++
						goto l590
					l634:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ř') {
							goto l635
						}
						position++
						goto l590
					l635:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ś') {
							goto l636
						}
						position++
						goto l590
					l636:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ŝ') {
							goto l637
						}
						position++
						goto l590
					l637:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ş') {
							goto l638
						}
						position++
						goto l590
					l638:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Š') {
							goto l639
						}
						position++
						goto l590
					l639:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ÿ') {
							goto l640
						}
						position++
						goto l590
					l640:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ź') {
							goto l641
						}
						position++
						goto l590
					l641:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ż') {
							goto l642
						}
						position++
						goto l590
					l642:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ž') {
							goto l643
						}
						position++
						goto l590
					l643:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('ƒ') {
							goto l644
						}
						position++
						goto l590
					l644:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ǿ') {
							goto l645
						}
						position++
						goto l590
					l645:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ș') {
							goto l646
						}
						position++
						goto l590
					l646:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('Ț') {
							goto l586
						}
						position++
					}
				l590:
				}
			l588:
				add(ruleAuthorUpperChar, position587)
			}
			return true
		l586:
			position, tokenIndex = position586, tokenIndex586
			return false
		},
		/* 79 AuthorLowerChar <- <(lASCII / ('à' / 'á' / 'â' / 'ã' / 'ä' / 'å' / 'æ' / 'ç' / 'è' / 'é' / 'ê' / 'ë' / 'ì' / 'í' / 'î' / 'ï' / 'ð' / 'ñ' / 'ò' / 'ó' / 'ó' / 'ô' / 'õ' / 'ö' / 'ø' / 'ù' / 'ú' / 'û' / 'ü' / 'ý' / 'ÿ' / 'ā' / 'ă' / 'ą' / 'ć' / 'ĉ' / 'č' / 'ď' / 'đ' / '\'' / 'ē' / 'ĕ' / 'ė' / 'ę' / 'ě' / 'ğ' / 'ī' / 'ĭ' / 'İ' / 'ı' / 'ĺ' / 'ľ' / 'ł' / 'ń' / 'ņ' / 'ň' / 'ŏ' / 'ő' / 'œ' / 'ŕ' / 'ř' / 'ś' / 'ş' / 'š' / 'ţ' / 'ť' / 'ũ' / 'ū' / 'ŭ' / 'ů' / 'ű' / 'ź' / 'ż' / 'ž' / 'ſ' / 'ǎ' / 'ǔ' / 'ǧ' / 'ș' / 'ț' / 'ȳ' / 'ß'))> */
		func() bool {
			position647, tokenIndex647 := position, tokenIndex
			{
				position648 := position
				{
					position649, tokenIndex649 := position, tokenIndex
					if !_rules[rulelASCII]() {
						goto l650
					}
					goto l649
				l650:
					position, tokenIndex = position649, tokenIndex649
					{
						position651, tokenIndex651 := position, tokenIndex
						if buffer[position] != rune('à') {
							goto l652
						}
						position++
						goto l651
					l652:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('á') {
							goto l653
						}
						position++
						goto l651
					l653:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('â') {
							goto l654
						}
						position++
						goto l651
					l654:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ã') {
							goto l655
						}
						position++
						goto l651
					l655:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ä') {
							goto l656
						}
						position++
						goto l651
					l656:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('å') {
							goto l657
						}
						position++
						goto l651
					l657:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('æ') {
							goto l658
						}
						position++
						goto l651
					l658:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ç') {
							goto l659
						}
						position++
						goto l651
					l659:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('è') {
							goto l660
						}
						position++
						goto l651
					l660:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('é') {
							goto l661
						}
						position++
						goto l651
					l661:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ê') {
							goto l662
						}
						position++
						goto l651
					l662:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ë') {
							goto l663
						}
						position++
						goto l651
					l663:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ì') {
							goto l664
						}
						position++
						goto l651
					l664:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('í') {
							goto l665
						}
						position++
						goto l651
					l665:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('î') {
							goto l666
						}
						position++
						goto l651
					l666:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ï') {
							goto l667
						}
						position++
						goto l651
					l667:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ð') {
							goto l668
						}
						position++
						goto l651
					l668:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ñ') {
							goto l669
						}
						position++
						goto l651
					l669:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ò') {
							goto l670
						}
						position++
						goto l651
					l670:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ó') {
							goto l671
						}
						position++
						goto l651
					l671:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ó') {
							goto l672
						}
						position++
						goto l651
					l672:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ô') {
							goto l673
						}
						position++
						goto l651
					l673:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('õ') {
							goto l674
						}
						position++
						goto l651
					l674:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ö') {
							goto l675
						}
						position++
						goto l651
					l675:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ø') {
							goto l676
						}
						position++
						goto l651
					l676:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ù') {
							goto l677
						}
						position++
						goto l651
					l677:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ú') {
							goto l678
						}
						position++
						goto l651
					l678:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('û') {
							goto l679
						}
						position++
						goto l651
					l679:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ü') {
							goto l680
						}
						position++
						goto l651
					l680:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ý') {
							goto l681
						}
						position++
						goto l651
					l681:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ÿ') {
							goto l682
						}
						position++
						goto l651
					l682:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ā') {
							goto l683
						}
						position++
						goto l651
					l683:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ă') {
							goto l684
						}
						position++
						goto l651
					l684:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ą') {
							goto l685
						}
						position++
						goto l651
					l685:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ć') {
							goto l686
						}
						position++
						goto l651
					l686:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ĉ') {
							goto l687
						}
						position++
						goto l651
					l687:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('č') {
							goto l688
						}
						position++
						goto l651
					l688:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ď') {
							goto l689
						}
						position++
						goto l651
					l689:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('đ') {
							goto l690
						}
						position++
						goto l651
					l690:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('\'') {
							goto l691
						}
						position++
						goto l651
					l691:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ē') {
							goto l692
						}
						position++
						goto l651
					l692:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ĕ') {
							goto l693
						}
						position++
						goto l651
					l693:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ė') {
							goto l694
						}
						position++
						goto l651
					l694:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ę') {
							goto l695
						}
						position++
						goto l651
					l695:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ě') {
							goto l696
						}
						position++
						goto l651
					l696:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ğ') {
							goto l697
						}
						position++
						goto l651
					l697:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ī') {
							goto l698
						}
						position++
						goto l651
					l698:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ĭ') {
							goto l699
						}
						position++
						goto l651
					l699:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('İ') {
							goto l700
						}
						position++
						goto l651
					l700:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ı') {
							goto l701
						}
						position++
						goto l651
					l701:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ĺ') {
							goto l702
						}
						position++
						goto l651
					l702:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ľ') {
							goto l703
						}
						position++
						goto l651
					l703:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ł') {
							goto l704
						}
						position++
						goto l651
					l704:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ń') {
							goto l705
						}
						position++
						goto l651
					l705:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ņ') {
							goto l706
						}
						position++
						goto l651
					l706:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ň') {
							goto l707
						}
						position++
						goto l651
					l707:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ŏ') {
							goto l708
						}
						position++
						goto l651
					l708:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ő') {
							goto l709
						}
						position++
						goto l651
					l709:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('œ') {
							goto l710
						}
						position++
						goto l651
					l710:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ŕ') {
							goto l711
						}
						position++
						goto l651
					l711:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ř') {
							goto l712
						}
						position++
						goto l651
					l712:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ś') {
							goto l713
						}
						position++
						goto l651
					l713:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ş') {
							goto l714
						}
						position++
						goto l651
					l714:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('š') {
							goto l715
						}
						position++
						goto l651
					l715:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ţ') {
							goto l716
						}
						position++
						goto l651
					l716:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ť') {
							goto l717
						}
						position++
						goto l651
					l717:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ũ') {
							goto l718
						}
						position++
						goto l651
					l718:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ū') {
							goto l719
						}
						position++
						goto l651
					l719:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ŭ') {
							goto l720
						}
						position++
						goto l651
					l720:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ů') {
							goto l721
						}
						position++
						goto l651
					l721:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ű') {
							goto l722
						}
						position++
						goto l651
					l722:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ź') {
							goto l723
						}
						position++
						goto l651
					l723:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ż') {
							goto l724
						}
						position++
						goto l651
					l724:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ž') {
							goto l725
						}
						position++
						goto l651
					l725:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ſ') {
							goto l726
						}
						position++
						goto l651
					l726:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ǎ') {
							goto l727
						}
						position++
						goto l651
					l727:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ǔ') {
							goto l728
						}
						position++
						goto l651
					l728:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ǧ') {
							goto l729
						}
						position++
						goto l651
					l729:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ș') {
							goto l730
						}
						position++
						goto l651
					l730:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ț') {
							goto l731
						}
						position++
						goto l651
					l731:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ȳ') {
							goto l732
						}
						position++
						goto l651
					l732:
						position, tokenIndex = position651, tokenIndex651
						if buffer[position] != rune('ß') {
							goto l647
						}
						position++
					}
				l651:
				}
			l649:
				add(ruleAuthorLowerChar, position648)
			}
			return true
		l647:
			position, tokenIndex = position647, tokenIndex647
			return false
		},
		/* 80 Year <- <(YearRange / YearApprox / YearWithParens / YearWithPage / YearWithDot / YearWithChar / YearNum)> */
		func() bool {
			position733, tokenIndex733 := position, tokenIndex
			{
				position734 := position
				{
					position735, tokenIndex735 := position, tokenIndex
					if !_rules[ruleYearRange]() {
						goto l736
					}
					goto l735
				l736:
					position, tokenIndex = position735, tokenIndex735
					if !_rules[ruleYearApprox]() {
						goto l737
					}
					goto l735
				l737:
					position, tokenIndex = position735, tokenIndex735
					if !_rules[ruleYearWithParens]() {
						goto l738
					}
					goto l735
				l738:
					position, tokenIndex = position735, tokenIndex735
					if !_rules[ruleYearWithPage]() {
						goto l739
					}
					goto l735
				l739:
					position, tokenIndex = position735, tokenIndex735
					if !_rules[ruleYearWithDot]() {
						goto l740
					}
					goto l735
				l740:
					position, tokenIndex = position735, tokenIndex735
					if !_rules[ruleYearWithChar]() {
						goto l741
					}
					goto l735
				l741:
					position, tokenIndex = position735, tokenIndex735
					if !_rules[ruleYearNum]() {
						goto l733
					}
				}
			l735:
				add(ruleYear, position734)
			}
			return true
		l733:
			position, tokenIndex = position733, tokenIndex733
			return false
		},
		/* 81 YearRange <- <(YearNum dash (nums+ ('a' / 'b' / 'c' / 'd' / 'e' / 'f' / 'g' / 'h' / 'i' / 'j' / 'k' / 'l' / 'm' / 'n' / 'o' / 'p' / 'q' / 'r' / 's' / 't' / 'u' / 'v' / 'w' / 'x' / 'y' / 'z' / '?')*))> */
		func() bool {
			position742, tokenIndex742 := position, tokenIndex
			{
				position743 := position
				if !_rules[ruleYearNum]() {
					goto l742
				}
				if !_rules[ruledash]() {
					goto l742
				}
				if !_rules[rulenums]() {
					goto l742
				}
			l744:
				{
					position745, tokenIndex745 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l745
					}
					goto l744
				l745:
					position, tokenIndex = position745, tokenIndex745
				}
			l746:
				{
					position747, tokenIndex747 := position, tokenIndex
					{
						position748, tokenIndex748 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l749
						}
						position++
						goto l748
					l749:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('b') {
							goto l750
						}
						position++
						goto l748
					l750:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('c') {
							goto l751
						}
						position++
						goto l748
					l751:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('d') {
							goto l752
						}
						position++
						goto l748
					l752:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('e') {
							goto l753
						}
						position++
						goto l748
					l753:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('f') {
							goto l754
						}
						position++
						goto l748
					l754:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('g') {
							goto l755
						}
						position++
						goto l748
					l755:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('h') {
							goto l756
						}
						position++
						goto l748
					l756:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('i') {
							goto l757
						}
						position++
						goto l748
					l757:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('j') {
							goto l758
						}
						position++
						goto l748
					l758:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('k') {
							goto l759
						}
						position++
						goto l748
					l759:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('l') {
							goto l760
						}
						position++
						goto l748
					l760:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('m') {
							goto l761
						}
						position++
						goto l748
					l761:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('n') {
							goto l762
						}
						position++
						goto l748
					l762:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('o') {
							goto l763
						}
						position++
						goto l748
					l763:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('p') {
							goto l764
						}
						position++
						goto l748
					l764:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('q') {
							goto l765
						}
						position++
						goto l748
					l765:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('r') {
							goto l766
						}
						position++
						goto l748
					l766:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('s') {
							goto l767
						}
						position++
						goto l748
					l767:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('t') {
							goto l768
						}
						position++
						goto l748
					l768:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('u') {
							goto l769
						}
						position++
						goto l748
					l769:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('v') {
							goto l770
						}
						position++
						goto l748
					l770:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('w') {
							goto l771
						}
						position++
						goto l748
					l771:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('x') {
							goto l772
						}
						position++
						goto l748
					l772:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('y') {
							goto l773
						}
						position++
						goto l748
					l773:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('z') {
							goto l774
						}
						position++
						goto l748
					l774:
						position, tokenIndex = position748, tokenIndex748
						if buffer[position] != rune('?') {
							goto l747
						}
						position++
					}
				l748:
					goto l746
				l747:
					position, tokenIndex = position747, tokenIndex747
				}
				add(ruleYearRange, position743)
			}
			return true
		l742:
			position, tokenIndex = position742, tokenIndex742
			return false
		},
		/* 82 YearWithDot <- <(YearNum '.')> */
		func() bool {
			position775, tokenIndex775 := position, tokenIndex
			{
				position776 := position
				if !_rules[ruleYearNum]() {
					goto l775
				}
				if buffer[position] != rune('.') {
					goto l775
				}
				position++
				add(ruleYearWithDot, position776)
			}
			return true
		l775:
			position, tokenIndex = position775, tokenIndex775
			return false
		},
		/* 83 YearApprox <- <('[' _? YearNum _? ']')> */
		func() bool {
			position777, tokenIndex777 := position, tokenIndex
			{
				position778 := position
				if buffer[position] != rune('[') {
					goto l777
				}
				position++
				{
					position779, tokenIndex779 := position, tokenIndex
					if !_rules[rule_]() {
						goto l779
					}
					goto l780
				l779:
					position, tokenIndex = position779, tokenIndex779
				}
			l780:
				if !_rules[ruleYearNum]() {
					goto l777
				}
				{
					position781, tokenIndex781 := position, tokenIndex
					if !_rules[rule_]() {
						goto l781
					}
					goto l782
				l781:
					position, tokenIndex = position781, tokenIndex781
				}
			l782:
				if buffer[position] != rune(']') {
					goto l777
				}
				position++
				add(ruleYearApprox, position778)
			}
			return true
		l777:
			position, tokenIndex = position777, tokenIndex777
			return false
		},
		/* 84 YearWithPage <- <((YearWithChar / YearNum) _? ':' _? nums+)> */
		func() bool {
			position783, tokenIndex783 := position, tokenIndex
			{
				position784 := position
				{
					position785, tokenIndex785 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l786
					}
					goto l785
				l786:
					position, tokenIndex = position785, tokenIndex785
					if !_rules[ruleYearNum]() {
						goto l783
					}
				}
			l785:
				{
					position787, tokenIndex787 := position, tokenIndex
					if !_rules[rule_]() {
						goto l787
					}
					goto l788
				l787:
					position, tokenIndex = position787, tokenIndex787
				}
			l788:
				if buffer[position] != rune(':') {
					goto l783
				}
				position++
				{
					position789, tokenIndex789 := position, tokenIndex
					if !_rules[rule_]() {
						goto l789
					}
					goto l790
				l789:
					position, tokenIndex = position789, tokenIndex789
				}
			l790:
				if !_rules[rulenums]() {
					goto l783
				}
			l791:
				{
					position792, tokenIndex792 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l792
					}
					goto l791
				l792:
					position, tokenIndex = position792, tokenIndex792
				}
				add(ruleYearWithPage, position784)
			}
			return true
		l783:
			position, tokenIndex = position783, tokenIndex783
			return false
		},
		/* 85 YearWithParens <- <('(' (YearWithChar / YearNum) ')')> */
		func() bool {
			position793, tokenIndex793 := position, tokenIndex
			{
				position794 := position
				if buffer[position] != rune('(') {
					goto l793
				}
				position++
				{
					position795, tokenIndex795 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l796
					}
					goto l795
				l796:
					position, tokenIndex = position795, tokenIndex795
					if !_rules[ruleYearNum]() {
						goto l793
					}
				}
			l795:
				if buffer[position] != rune(')') {
					goto l793
				}
				position++
				add(ruleYearWithParens, position794)
			}
			return true
		l793:
			position, tokenIndex = position793, tokenIndex793
			return false
		},
		/* 86 YearWithChar <- <(YearNum lASCII Action0)> */
		func() bool {
			position797, tokenIndex797 := position, tokenIndex
			{
				position798 := position
				if !_rules[ruleYearNum]() {
					goto l797
				}
				if !_rules[rulelASCII]() {
					goto l797
				}
				if !_rules[ruleAction0]() {
					goto l797
				}
				add(ruleYearWithChar, position798)
			}
			return true
		l797:
			position, tokenIndex = position797, tokenIndex797
			return false
		},
		/* 87 YearNum <- <(('1' / '2') ('0' / '7' / '8' / '9') nums (nums / '?') '?'*)> */
		func() bool {
			position799, tokenIndex799 := position, tokenIndex
			{
				position800 := position
				{
					position801, tokenIndex801 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l802
					}
					position++
					goto l801
				l802:
					position, tokenIndex = position801, tokenIndex801
					if buffer[position] != rune('2') {
						goto l799
					}
					position++
				}
			l801:
				{
					position803, tokenIndex803 := position, tokenIndex
					if buffer[position] != rune('0') {
						goto l804
					}
					position++
					goto l803
				l804:
					position, tokenIndex = position803, tokenIndex803
					if buffer[position] != rune('7') {
						goto l805
					}
					position++
					goto l803
				l805:
					position, tokenIndex = position803, tokenIndex803
					if buffer[position] != rune('8') {
						goto l806
					}
					position++
					goto l803
				l806:
					position, tokenIndex = position803, tokenIndex803
					if buffer[position] != rune('9') {
						goto l799
					}
					position++
				}
			l803:
				if !_rules[rulenums]() {
					goto l799
				}
				{
					position807, tokenIndex807 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l808
					}
					goto l807
				l808:
					position, tokenIndex = position807, tokenIndex807
					if buffer[position] != rune('?') {
						goto l799
					}
					position++
				}
			l807:
			l809:
				{
					position810, tokenIndex810 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l810
					}
					position++
					goto l809
				l810:
					position, tokenIndex = position810, tokenIndex810
				}
				add(ruleYearNum, position800)
			}
			return true
		l799:
			position, tokenIndex = position799, tokenIndex799
			return false
		},
		/* 88 NameUpperChar <- <(UpperChar / UpperCharExtended)> */
		func() bool {
			position811, tokenIndex811 := position, tokenIndex
			{
				position812 := position
				{
					position813, tokenIndex813 := position, tokenIndex
					if !_rules[ruleUpperChar]() {
						goto l814
					}
					goto l813
				l814:
					position, tokenIndex = position813, tokenIndex813
					if !_rules[ruleUpperCharExtended]() {
						goto l811
					}
				}
			l813:
				add(ruleNameUpperChar, position812)
			}
			return true
		l811:
			position, tokenIndex = position811, tokenIndex811
			return false
		},
		/* 89 UpperCharExtended <- <('Æ' / 'Œ' / 'Ö')> */
		func() bool {
			position815, tokenIndex815 := position, tokenIndex
			{
				position816 := position
				{
					position817, tokenIndex817 := position, tokenIndex
					if buffer[position] != rune('Æ') {
						goto l818
					}
					position++
					goto l817
				l818:
					position, tokenIndex = position817, tokenIndex817
					if buffer[position] != rune('Œ') {
						goto l819
					}
					position++
					goto l817
				l819:
					position, tokenIndex = position817, tokenIndex817
					if buffer[position] != rune('Ö') {
						goto l815
					}
					position++
				}
			l817:
				add(ruleUpperCharExtended, position816)
			}
			return true
		l815:
			position, tokenIndex = position815, tokenIndex815
			return false
		},
		/* 90 UpperChar <- <hASCII> */
		func() bool {
			position820, tokenIndex820 := position, tokenIndex
			{
				position821 := position
				if !_rules[rulehASCII]() {
					goto l820
				}
				add(ruleUpperChar, position821)
			}
			return true
		l820:
			position, tokenIndex = position820, tokenIndex820
			return false
		},
		/* 91 NameLowerChar <- <(LowerChar / LowerCharExtended / MiscodedChar)> */
		func() bool {
			position822, tokenIndex822 := position, tokenIndex
			{
				position823 := position
				{
					position824, tokenIndex824 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l825
					}
					goto l824
				l825:
					position, tokenIndex = position824, tokenIndex824
					if !_rules[ruleLowerCharExtended]() {
						goto l826
					}
					goto l824
				l826:
					position, tokenIndex = position824, tokenIndex824
					if !_rules[ruleMiscodedChar]() {
						goto l822
					}
				}
			l824:
				add(ruleNameLowerChar, position823)
			}
			return true
		l822:
			position, tokenIndex = position822, tokenIndex822
			return false
		},
		/* 92 MiscodedChar <- <'�'> */
		func() bool {
			position827, tokenIndex827 := position, tokenIndex
			{
				position828 := position
				if buffer[position] != rune('�') {
					goto l827
				}
				position++
				add(ruleMiscodedChar, position828)
			}
			return true
		l827:
			position, tokenIndex = position827, tokenIndex827
			return false
		},
		/* 93 LowerCharExtended <- <('æ' / 'œ' / 'à' / 'â' / 'å' / 'ã' / 'ä' / 'á' / 'ç' / 'č' / 'é' / 'è' / 'ë' / 'í' / 'ì' / 'ï' / 'ň' / 'ñ' / 'ñ' / 'ó' / 'ò' / 'ô' / 'ø' / 'õ' / 'ö' / 'ú' / 'ù' / 'ü' / 'ŕ' / 'ř' / 'ŗ' / 'ſ' / 'š' / 'š' / 'ş' / 'ž')> */
		func() bool {
			position829, tokenIndex829 := position, tokenIndex
			{
				position830 := position
				{
					position831, tokenIndex831 := position, tokenIndex
					if buffer[position] != rune('æ') {
						goto l832
					}
					position++
					goto l831
				l832:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('œ') {
						goto l833
					}
					position++
					goto l831
				l833:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('à') {
						goto l834
					}
					position++
					goto l831
				l834:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('â') {
						goto l835
					}
					position++
					goto l831
				l835:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('å') {
						goto l836
					}
					position++
					goto l831
				l836:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('ã') {
						goto l837
					}
					position++
					goto l831
				l837:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('ä') {
						goto l838
					}
					position++
					goto l831
				l838:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('á') {
						goto l839
					}
					position++
					goto l831
				l839:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('ç') {
						goto l840
					}
					position++
					goto l831
				l840:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('č') {
						goto l841
					}
					position++
					goto l831
				l841:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('é') {
						goto l842
					}
					position++
					goto l831
				l842:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('è') {
						goto l843
					}
					position++
					goto l831
				l843:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('ë') {
						goto l844
					}
					position++
					goto l831
				l844:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('í') {
						goto l845
					}
					position++
					goto l831
				l845:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('ì') {
						goto l846
					}
					position++
					goto l831
				l846:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('ï') {
						goto l847
					}
					position++
					goto l831
				l847:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('ň') {
						goto l848
					}
					position++
					goto l831
				l848:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('ñ') {
						goto l849
					}
					position++
					goto l831
				l849:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('ñ') {
						goto l850
					}
					position++
					goto l831
				l850:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('ó') {
						goto l851
					}
					position++
					goto l831
				l851:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('ò') {
						goto l852
					}
					position++
					goto l831
				l852:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('ô') {
						goto l853
					}
					position++
					goto l831
				l853:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('ø') {
						goto l854
					}
					position++
					goto l831
				l854:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('õ') {
						goto l855
					}
					position++
					goto l831
				l855:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('ö') {
						goto l856
					}
					position++
					goto l831
				l856:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('ú') {
						goto l857
					}
					position++
					goto l831
				l857:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('ù') {
						goto l858
					}
					position++
					goto l831
				l858:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('ü') {
						goto l859
					}
					position++
					goto l831
				l859:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('ŕ') {
						goto l860
					}
					position++
					goto l831
				l860:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('ř') {
						goto l861
					}
					position++
					goto l831
				l861:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('ŗ') {
						goto l862
					}
					position++
					goto l831
				l862:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('ſ') {
						goto l863
					}
					position++
					goto l831
				l863:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('š') {
						goto l864
					}
					position++
					goto l831
				l864:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('š') {
						goto l865
					}
					position++
					goto l831
				l865:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('ş') {
						goto l866
					}
					position++
					goto l831
				l866:
					position, tokenIndex = position831, tokenIndex831
					if buffer[position] != rune('ž') {
						goto l829
					}
					position++
				}
			l831:
				add(ruleLowerCharExtended, position830)
			}
			return true
		l829:
			position, tokenIndex = position829, tokenIndex829
			return false
		},
		/* 94 LowerChar <- <lASCII> */
		func() bool {
			position867, tokenIndex867 := position, tokenIndex
			{
				position868 := position
				if !_rules[rulelASCII]() {
					goto l867
				}
				add(ruleLowerChar, position868)
			}
			return true
		l867:
			position, tokenIndex = position867, tokenIndex867
			return false
		},
		/* 95 SpaceCharEOI <- <(_ / !.)> */
		func() bool {
			position869, tokenIndex869 := position, tokenIndex
			{
				position870 := position
				{
					position871, tokenIndex871 := position, tokenIndex
					if !_rules[rule_]() {
						goto l872
					}
					goto l871
				l872:
					position, tokenIndex = position871, tokenIndex871
					{
						position873, tokenIndex873 := position, tokenIndex
						if !matchDot() {
							goto l873
						}
						goto l869
					l873:
						position, tokenIndex = position873, tokenIndex873
					}
				}
			l871:
				add(ruleSpaceCharEOI, position870)
			}
			return true
		l869:
			position, tokenIndex = position869, tokenIndex869
			return false
		},
		/* 96 nums <- <[0-9]> */
		func() bool {
			position874, tokenIndex874 := position, tokenIndex
			{
				position875 := position
				if c := buffer[position]; c < rune('0') || c > rune('9') {
					goto l874
				}
				position++
				add(rulenums, position875)
			}
			return true
		l874:
			position, tokenIndex = position874, tokenIndex874
			return false
		},
		/* 97 lASCII <- <[a-z]> */
		func() bool {
			position876, tokenIndex876 := position, tokenIndex
			{
				position877 := position
				if c := buffer[position]; c < rune('a') || c > rune('z') {
					goto l876
				}
				position++
				add(rulelASCII, position877)
			}
			return true
		l876:
			position, tokenIndex = position876, tokenIndex876
			return false
		},
		/* 98 hASCII <- <[A-Z]> */
		func() bool {
			position878, tokenIndex878 := position, tokenIndex
			{
				position879 := position
				if c := buffer[position]; c < rune('A') || c > rune('Z') {
					goto l878
				}
				position++
				add(rulehASCII, position879)
			}
			return true
		l878:
			position, tokenIndex = position878, tokenIndex878
			return false
		},
		/* 99 apostr <- <'\''> */
		func() bool {
			position880, tokenIndex880 := position, tokenIndex
			{
				position881 := position
				if buffer[position] != rune('\'') {
					goto l880
				}
				position++
				add(ruleapostr, position881)
			}
			return true
		l880:
			position, tokenIndex = position880, tokenIndex880
			return false
		},
		/* 100 dash <- <'-'> */
		func() bool {
			position882, tokenIndex882 := position, tokenIndex
			{
				position883 := position
				if buffer[position] != rune('-') {
					goto l882
				}
				position++
				add(ruledash, position883)
			}
			return true
		l882:
			position, tokenIndex = position882, tokenIndex882
			return false
		},
		/* 101 _ <- <(MultipleSpace / SingleSpace)> */
		func() bool {
			position884, tokenIndex884 := position, tokenIndex
			{
				position885 := position
				{
					position886, tokenIndex886 := position, tokenIndex
					if !_rules[ruleMultipleSpace]() {
						goto l887
					}
					goto l886
				l887:
					position, tokenIndex = position886, tokenIndex886
					if !_rules[ruleSingleSpace]() {
						goto l884
					}
				}
			l886:
				add(rule_, position885)
			}
			return true
		l884:
			position, tokenIndex = position884, tokenIndex884
			return false
		},
		/* 102 MultipleSpace <- <(SingleSpace SingleSpace+)> */
		func() bool {
			position888, tokenIndex888 := position, tokenIndex
			{
				position889 := position
				if !_rules[ruleSingleSpace]() {
					goto l888
				}
				if !_rules[ruleSingleSpace]() {
					goto l888
				}
			l890:
				{
					position891, tokenIndex891 := position, tokenIndex
					if !_rules[ruleSingleSpace]() {
						goto l891
					}
					goto l890
				l891:
					position, tokenIndex = position891, tokenIndex891
				}
				add(ruleMultipleSpace, position889)
			}
			return true
		l888:
			position, tokenIndex = position888, tokenIndex888
			return false
		},
		/* 103 SingleSpace <- <' '> */
		func() bool {
			position892, tokenIndex892 := position, tokenIndex
			{
				position893 := position
				if buffer[position] != rune(' ') {
					goto l892
				}
				position++
				add(ruleSingleSpace, position893)
			}
			return true
		l892:
			position, tokenIndex = position892, tokenIndex892
			return false
		},
		/* 105 Action0 <- <{ p.AddWarn(YearCharWarn) }> */
		func() bool {
			{
				add(ruleAction0, position)
			}
			return true
		},
	}
	p.rules = _rules
}
