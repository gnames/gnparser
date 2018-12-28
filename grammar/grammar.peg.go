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
	ruleSciName1
	ruleSciName2
	ruleHybridFormula
	ruleHybridFormula1
	ruleHybridFormula2
	ruleNamedHybrid
	ruleName
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
	ruleWord2StartDigit
	ruleWord2
	ruleWord3
	ruleWord4
	ruleHybridChar
	ruleApproxName
	ruleApproxName1
	ruleApproxName2
	ruleApproxNameIgnored
	ruleApproximation
	ruleAuthorship
	ruleAuthorship1
	ruleAuthorship2
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
	ruleWordBorderChar
	rulenums
	rulelASCII
	rulehASCII
	ruleapostr
	ruledash
	rule_
	ruleAction0
)

var rul3s = [...]string{
	"Unknown",
	"SciName",
	"Tail",
	"SciName1",
	"SciName2",
	"HybridFormula",
	"HybridFormula1",
	"HybridFormula2",
	"NamedHybrid",
	"Name",
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
	"Word2StartDigit",
	"Word2",
	"Word3",
	"Word4",
	"HybridChar",
	"ApproxName",
	"ApproxName1",
	"ApproxName2",
	"ApproxNameIgnored",
	"Approximation",
	"Authorship",
	"Authorship1",
	"Authorship2",
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
	"WordBorderChar",
	"nums",
	"lASCII",
	"hASCII",
	"apostr",
	"dash",
	"_",
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
				fmt.Printf(" ")
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
	rules  [104]func() bool
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
			p.addWarn(YearCharWarn)

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
		/* 0 SciName <- <(_? SciName1 Tail !.)> */
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
				if !_rules[ruleSciName1]() {
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
		/* 2 SciName1 <- <SciName2> */
		func() bool {
			position9, tokenIndex9 := position, tokenIndex
			{
				position10 := position
				if !_rules[ruleSciName2]() {
					goto l9
				}
				add(ruleSciName1, position10)
			}
			return true
		l9:
			position, tokenIndex = position9, tokenIndex9
			return false
		},
		/* 3 SciName2 <- <Name> */
		func() bool {
			position11, tokenIndex11 := position, tokenIndex
			{
				position12 := position
				if !_rules[ruleName]() {
					goto l11
				}
				add(ruleSciName2, position12)
			}
			return true
		l11:
			position, tokenIndex = position11, tokenIndex11
			return false
		},
		/* 4 HybridFormula <- <(Name (_ (HybridFormula1 / HybridFormula2)))> */
		nil,
		/* 5 HybridFormula1 <- <(HybridChar _? SpeciesEpithet (_ InfraspGroup)?)> */
		nil,
		/* 6 HybridFormula2 <- <(HybridChar (_ Name)?)> */
		nil,
		/* 7 NamedHybrid <- <(HybridChar _? Name)> */
		nil,
		/* 8 Name <- <(NameSpecies / NameUninomial)> */
		func() bool {
			position17, tokenIndex17 := position, tokenIndex
			{
				position18 := position
				{
					position19, tokenIndex19 := position, tokenIndex
					if !_rules[ruleNameSpecies]() {
						goto l20
					}
					goto l19
				l20:
					position, tokenIndex = position19, tokenIndex19
					if !_rules[ruleNameUninomial]() {
						goto l17
					}
				}
			l19:
				add(ruleName, position18)
			}
			return true
		l17:
			position, tokenIndex = position17, tokenIndex17
			return false
		},
		/* 9 NameUninomial <- <(UninomialCombo / Uninomial)> */
		func() bool {
			position21, tokenIndex21 := position, tokenIndex
			{
				position22 := position
				{
					position23, tokenIndex23 := position, tokenIndex
					if !_rules[ruleUninomialCombo]() {
						goto l24
					}
					goto l23
				l24:
					position, tokenIndex = position23, tokenIndex23
					if !_rules[ruleUninomial]() {
						goto l21
					}
				}
			l23:
				add(ruleNameUninomial, position22)
			}
			return true
		l21:
			position, tokenIndex = position21, tokenIndex21
			return false
		},
		/* 10 NameApprox <- <(GenusWord _ Approximation (_ SpeciesEpithet)?)> */
		nil,
		/* 11 NameComp <- <(GenusWord _ Comparison (_ SpeciesEpithet)?)> */
		nil,
		/* 12 NameSpecies <- <(GenusWord (_? (SubGenus / SubGenusOrSuperspecies))? _ SpeciesEpithet (_ InfraspGroup)?)> */
		func() bool {
			position27, tokenIndex27 := position, tokenIndex
			{
				position28 := position
				if !_rules[ruleGenusWord]() {
					goto l27
				}
				{
					position29, tokenIndex29 := position, tokenIndex
					{
						position31, tokenIndex31 := position, tokenIndex
						if !_rules[rule_]() {
							goto l31
						}
						goto l32
					l31:
						position, tokenIndex = position31, tokenIndex31
					}
				l32:
					{
						position33, tokenIndex33 := position, tokenIndex
						if !_rules[ruleSubGenus]() {
							goto l34
						}
						goto l33
					l34:
						position, tokenIndex = position33, tokenIndex33
						if !_rules[ruleSubGenusOrSuperspecies]() {
							goto l29
						}
					}
				l33:
					goto l30
				l29:
					position, tokenIndex = position29, tokenIndex29
				}
			l30:
				if !_rules[rule_]() {
					goto l27
				}
				if !_rules[ruleSpeciesEpithet]() {
					goto l27
				}
				{
					position35, tokenIndex35 := position, tokenIndex
					if !_rules[rule_]() {
						goto l35
					}
					if !_rules[ruleInfraspGroup]() {
						goto l35
					}
					goto l36
				l35:
					position, tokenIndex = position35, tokenIndex35
				}
			l36:
				add(ruleNameSpecies, position28)
			}
			return true
		l27:
			position, tokenIndex = position27, tokenIndex27
			return false
		},
		/* 13 GenusWord <- <((AbbrGenus / UninomialWord) !(_ AuthorWord))> */
		func() bool {
			position37, tokenIndex37 := position, tokenIndex
			{
				position38 := position
				{
					position39, tokenIndex39 := position, tokenIndex
					if !_rules[ruleAbbrGenus]() {
						goto l40
					}
					goto l39
				l40:
					position, tokenIndex = position39, tokenIndex39
					if !_rules[ruleUninomialWord]() {
						goto l37
					}
				}
			l39:
				{
					position41, tokenIndex41 := position, tokenIndex
					if !_rules[rule_]() {
						goto l41
					}
					if !_rules[ruleAuthorWord]() {
						goto l41
					}
					goto l37
				l41:
					position, tokenIndex = position41, tokenIndex41
				}
				add(ruleGenusWord, position38)
			}
			return true
		l37:
			position, tokenIndex = position37, tokenIndex37
			return false
		},
		/* 14 InfraspGroup <- <(InfraspEpithet (_ InfraspEpithet)? (_ InfraspEpithet)?)> */
		func() bool {
			position42, tokenIndex42 := position, tokenIndex
			{
				position43 := position
				if !_rules[ruleInfraspEpithet]() {
					goto l42
				}
				{
					position44, tokenIndex44 := position, tokenIndex
					if !_rules[rule_]() {
						goto l44
					}
					if !_rules[ruleInfraspEpithet]() {
						goto l44
					}
					goto l45
				l44:
					position, tokenIndex = position44, tokenIndex44
				}
			l45:
				{
					position46, tokenIndex46 := position, tokenIndex
					if !_rules[rule_]() {
						goto l46
					}
					if !_rules[ruleInfraspEpithet]() {
						goto l46
					}
					goto l47
				l46:
					position, tokenIndex = position46, tokenIndex46
				}
			l47:
				add(ruleInfraspGroup, position43)
			}
			return true
		l42:
			position, tokenIndex = position42, tokenIndex42
			return false
		},
		/* 15 InfraspEpithet <- <((Rank _?)? !AuthorEx Word (_ Authorship)?)> */
		func() bool {
			position48, tokenIndex48 := position, tokenIndex
			{
				position49 := position
				{
					position50, tokenIndex50 := position, tokenIndex
					if !_rules[ruleRank]() {
						goto l50
					}
					{
						position52, tokenIndex52 := position, tokenIndex
						if !_rules[rule_]() {
							goto l52
						}
						goto l53
					l52:
						position, tokenIndex = position52, tokenIndex52
					}
				l53:
					goto l51
				l50:
					position, tokenIndex = position50, tokenIndex50
				}
			l51:
				{
					position54, tokenIndex54 := position, tokenIndex
					if !_rules[ruleAuthorEx]() {
						goto l54
					}
					goto l48
				l54:
					position, tokenIndex = position54, tokenIndex54
				}
				if !_rules[ruleWord]() {
					goto l48
				}
				{
					position55, tokenIndex55 := position, tokenIndex
					if !_rules[rule_]() {
						goto l55
					}
					if !_rules[ruleAuthorship]() {
						goto l55
					}
					goto l56
				l55:
					position, tokenIndex = position55, tokenIndex55
				}
			l56:
				add(ruleInfraspEpithet, position49)
			}
			return true
		l48:
			position, tokenIndex = position48, tokenIndex48
			return false
		},
		/* 16 SpeciesEpithet <- <(!AuthorEx Word (_? Authorship)? ','? &(SpaceCharEOI / '('))> */
		func() bool {
			position57, tokenIndex57 := position, tokenIndex
			{
				position58 := position
				{
					position59, tokenIndex59 := position, tokenIndex
					if !_rules[ruleAuthorEx]() {
						goto l59
					}
					goto l57
				l59:
					position, tokenIndex = position59, tokenIndex59
				}
				if !_rules[ruleWord]() {
					goto l57
				}
				{
					position60, tokenIndex60 := position, tokenIndex
					{
						position62, tokenIndex62 := position, tokenIndex
						if !_rules[rule_]() {
							goto l62
						}
						goto l63
					l62:
						position, tokenIndex = position62, tokenIndex62
					}
				l63:
					if !_rules[ruleAuthorship]() {
						goto l60
					}
					goto l61
				l60:
					position, tokenIndex = position60, tokenIndex60
				}
			l61:
				{
					position64, tokenIndex64 := position, tokenIndex
					if buffer[position] != rune(',') {
						goto l64
					}
					position++
					goto l65
				l64:
					position, tokenIndex = position64, tokenIndex64
				}
			l65:
				{
					position66, tokenIndex66 := position, tokenIndex
					{
						position67, tokenIndex67 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l68
						}
						goto l67
					l68:
						position, tokenIndex = position67, tokenIndex67
						if buffer[position] != rune('(') {
							goto l57
						}
						position++
					}
				l67:
					position, tokenIndex = position66, tokenIndex66
				}
				add(ruleSpeciesEpithet, position58)
			}
			return true
		l57:
			position, tokenIndex = position57, tokenIndex57
			return false
		},
		/* 17 Comparison <- <('c' 'f' '.'?)> */
		nil,
		/* 18 Rank <- <(RankForma / RankVar / RankSsp / RankOther / RankOtherUncommon)> */
		func() bool {
			position70, tokenIndex70 := position, tokenIndex
			{
				position71 := position
				{
					position72, tokenIndex72 := position, tokenIndex
					if !_rules[ruleRankForma]() {
						goto l73
					}
					goto l72
				l73:
					position, tokenIndex = position72, tokenIndex72
					if !_rules[ruleRankVar]() {
						goto l74
					}
					goto l72
				l74:
					position, tokenIndex = position72, tokenIndex72
					if !_rules[ruleRankSsp]() {
						goto l75
					}
					goto l72
				l75:
					position, tokenIndex = position72, tokenIndex72
					if !_rules[ruleRankOther]() {
						goto l76
					}
					goto l72
				l76:
					position, tokenIndex = position72, tokenIndex72
					if !_rules[ruleRankOtherUncommon]() {
						goto l70
					}
				}
			l72:
				add(ruleRank, position71)
			}
			return true
		l70:
			position, tokenIndex = position70, tokenIndex70
			return false
		},
		/* 19 RankOtherUncommon <- <(('*' / ('n' 'a' 't') / ('f' '.' 's' 'p') / ('m' 'u' 't' '.')) &SpaceCharEOI)> */
		func() bool {
			position77, tokenIndex77 := position, tokenIndex
			{
				position78 := position
				{
					position79, tokenIndex79 := position, tokenIndex
					if buffer[position] != rune('*') {
						goto l80
					}
					position++
					goto l79
				l80:
					position, tokenIndex = position79, tokenIndex79
					if buffer[position] != rune('n') {
						goto l81
					}
					position++
					if buffer[position] != rune('a') {
						goto l81
					}
					position++
					if buffer[position] != rune('t') {
						goto l81
					}
					position++
					goto l79
				l81:
					position, tokenIndex = position79, tokenIndex79
					if buffer[position] != rune('f') {
						goto l82
					}
					position++
					if buffer[position] != rune('.') {
						goto l82
					}
					position++
					if buffer[position] != rune('s') {
						goto l82
					}
					position++
					if buffer[position] != rune('p') {
						goto l82
					}
					position++
					goto l79
				l82:
					position, tokenIndex = position79, tokenIndex79
					if buffer[position] != rune('m') {
						goto l77
					}
					position++
					if buffer[position] != rune('u') {
						goto l77
					}
					position++
					if buffer[position] != rune('t') {
						goto l77
					}
					position++
					if buffer[position] != rune('.') {
						goto l77
					}
					position++
				}
			l79:
				{
					position83, tokenIndex83 := position, tokenIndex
					if !_rules[ruleSpaceCharEOI]() {
						goto l77
					}
					position, tokenIndex = position83, tokenIndex83
				}
				add(ruleRankOtherUncommon, position78)
			}
			return true
		l77:
			position, tokenIndex = position77, tokenIndex77
			return false
		},
		/* 20 RankOther <- <((('m' 'o' 'r' 'p' 'h' '.') / ('n' 'o' 't' 'h' 'o' 's' 'u' 'b' 's' 'p' '.') / ('c' 'o' 'n' 'v' 'a' 'r' '.') / ('p' 's' 'e' 'u' 'd' 'o' 'v' 'a' 'r' '.') / ('s' 'e' 'c' 't' '.') / ('s' 'e' 'r' '.') / ('s' 'u' 'b' 'v' 'a' 'r' '.') / ('s' 'u' 'b' 'f' '.') / ('r' 'a' 'c' 'e') / 'α' / ('β' 'β') / 'β' / 'γ' / 'δ' / 'ε' / 'φ' / 'θ' / 'μ' / ('a' '.') / ('b' '.') / ('c' '.') / ('d' '.') / ('e' '.') / ('g' '.') / ('k' '.') / ('p' 'v' '.') / ('p' 'a' 't' 'h' 'o' 'v' 'a' 'r' '.') / ('a' 'b' '.' (_? ('n' '.'))?) / ('s' 't' '.')) &SpaceCharEOI)> */
		func() bool {
			position84, tokenIndex84 := position, tokenIndex
			{
				position85 := position
				{
					position86, tokenIndex86 := position, tokenIndex
					if buffer[position] != rune('m') {
						goto l87
					}
					position++
					if buffer[position] != rune('o') {
						goto l87
					}
					position++
					if buffer[position] != rune('r') {
						goto l87
					}
					position++
					if buffer[position] != rune('p') {
						goto l87
					}
					position++
					if buffer[position] != rune('h') {
						goto l87
					}
					position++
					if buffer[position] != rune('.') {
						goto l87
					}
					position++
					goto l86
				l87:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('n') {
						goto l88
					}
					position++
					if buffer[position] != rune('o') {
						goto l88
					}
					position++
					if buffer[position] != rune('t') {
						goto l88
					}
					position++
					if buffer[position] != rune('h') {
						goto l88
					}
					position++
					if buffer[position] != rune('o') {
						goto l88
					}
					position++
					if buffer[position] != rune('s') {
						goto l88
					}
					position++
					if buffer[position] != rune('u') {
						goto l88
					}
					position++
					if buffer[position] != rune('b') {
						goto l88
					}
					position++
					if buffer[position] != rune('s') {
						goto l88
					}
					position++
					if buffer[position] != rune('p') {
						goto l88
					}
					position++
					if buffer[position] != rune('.') {
						goto l88
					}
					position++
					goto l86
				l88:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('c') {
						goto l89
					}
					position++
					if buffer[position] != rune('o') {
						goto l89
					}
					position++
					if buffer[position] != rune('n') {
						goto l89
					}
					position++
					if buffer[position] != rune('v') {
						goto l89
					}
					position++
					if buffer[position] != rune('a') {
						goto l89
					}
					position++
					if buffer[position] != rune('r') {
						goto l89
					}
					position++
					if buffer[position] != rune('.') {
						goto l89
					}
					position++
					goto l86
				l89:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('p') {
						goto l90
					}
					position++
					if buffer[position] != rune('s') {
						goto l90
					}
					position++
					if buffer[position] != rune('e') {
						goto l90
					}
					position++
					if buffer[position] != rune('u') {
						goto l90
					}
					position++
					if buffer[position] != rune('d') {
						goto l90
					}
					position++
					if buffer[position] != rune('o') {
						goto l90
					}
					position++
					if buffer[position] != rune('v') {
						goto l90
					}
					position++
					if buffer[position] != rune('a') {
						goto l90
					}
					position++
					if buffer[position] != rune('r') {
						goto l90
					}
					position++
					if buffer[position] != rune('.') {
						goto l90
					}
					position++
					goto l86
				l90:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('s') {
						goto l91
					}
					position++
					if buffer[position] != rune('e') {
						goto l91
					}
					position++
					if buffer[position] != rune('c') {
						goto l91
					}
					position++
					if buffer[position] != rune('t') {
						goto l91
					}
					position++
					if buffer[position] != rune('.') {
						goto l91
					}
					position++
					goto l86
				l91:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('s') {
						goto l92
					}
					position++
					if buffer[position] != rune('e') {
						goto l92
					}
					position++
					if buffer[position] != rune('r') {
						goto l92
					}
					position++
					if buffer[position] != rune('.') {
						goto l92
					}
					position++
					goto l86
				l92:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('s') {
						goto l93
					}
					position++
					if buffer[position] != rune('u') {
						goto l93
					}
					position++
					if buffer[position] != rune('b') {
						goto l93
					}
					position++
					if buffer[position] != rune('v') {
						goto l93
					}
					position++
					if buffer[position] != rune('a') {
						goto l93
					}
					position++
					if buffer[position] != rune('r') {
						goto l93
					}
					position++
					if buffer[position] != rune('.') {
						goto l93
					}
					position++
					goto l86
				l93:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('s') {
						goto l94
					}
					position++
					if buffer[position] != rune('u') {
						goto l94
					}
					position++
					if buffer[position] != rune('b') {
						goto l94
					}
					position++
					if buffer[position] != rune('f') {
						goto l94
					}
					position++
					if buffer[position] != rune('.') {
						goto l94
					}
					position++
					goto l86
				l94:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('r') {
						goto l95
					}
					position++
					if buffer[position] != rune('a') {
						goto l95
					}
					position++
					if buffer[position] != rune('c') {
						goto l95
					}
					position++
					if buffer[position] != rune('e') {
						goto l95
					}
					position++
					goto l86
				l95:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('α') {
						goto l96
					}
					position++
					goto l86
				l96:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('β') {
						goto l97
					}
					position++
					if buffer[position] != rune('β') {
						goto l97
					}
					position++
					goto l86
				l97:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('β') {
						goto l98
					}
					position++
					goto l86
				l98:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('γ') {
						goto l99
					}
					position++
					goto l86
				l99:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('δ') {
						goto l100
					}
					position++
					goto l86
				l100:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('ε') {
						goto l101
					}
					position++
					goto l86
				l101:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('φ') {
						goto l102
					}
					position++
					goto l86
				l102:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('θ') {
						goto l103
					}
					position++
					goto l86
				l103:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('μ') {
						goto l104
					}
					position++
					goto l86
				l104:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('a') {
						goto l105
					}
					position++
					if buffer[position] != rune('.') {
						goto l105
					}
					position++
					goto l86
				l105:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('b') {
						goto l106
					}
					position++
					if buffer[position] != rune('.') {
						goto l106
					}
					position++
					goto l86
				l106:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('c') {
						goto l107
					}
					position++
					if buffer[position] != rune('.') {
						goto l107
					}
					position++
					goto l86
				l107:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('d') {
						goto l108
					}
					position++
					if buffer[position] != rune('.') {
						goto l108
					}
					position++
					goto l86
				l108:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('e') {
						goto l109
					}
					position++
					if buffer[position] != rune('.') {
						goto l109
					}
					position++
					goto l86
				l109:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('g') {
						goto l110
					}
					position++
					if buffer[position] != rune('.') {
						goto l110
					}
					position++
					goto l86
				l110:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('k') {
						goto l111
					}
					position++
					if buffer[position] != rune('.') {
						goto l111
					}
					position++
					goto l86
				l111:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('p') {
						goto l112
					}
					position++
					if buffer[position] != rune('v') {
						goto l112
					}
					position++
					if buffer[position] != rune('.') {
						goto l112
					}
					position++
					goto l86
				l112:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('p') {
						goto l113
					}
					position++
					if buffer[position] != rune('a') {
						goto l113
					}
					position++
					if buffer[position] != rune('t') {
						goto l113
					}
					position++
					if buffer[position] != rune('h') {
						goto l113
					}
					position++
					if buffer[position] != rune('o') {
						goto l113
					}
					position++
					if buffer[position] != rune('v') {
						goto l113
					}
					position++
					if buffer[position] != rune('a') {
						goto l113
					}
					position++
					if buffer[position] != rune('r') {
						goto l113
					}
					position++
					if buffer[position] != rune('.') {
						goto l113
					}
					position++
					goto l86
				l113:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('a') {
						goto l114
					}
					position++
					if buffer[position] != rune('b') {
						goto l114
					}
					position++
					if buffer[position] != rune('.') {
						goto l114
					}
					position++
					{
						position115, tokenIndex115 := position, tokenIndex
						{
							position117, tokenIndex117 := position, tokenIndex
							if !_rules[rule_]() {
								goto l117
							}
							goto l118
						l117:
							position, tokenIndex = position117, tokenIndex117
						}
					l118:
						if buffer[position] != rune('n') {
							goto l115
						}
						position++
						if buffer[position] != rune('.') {
							goto l115
						}
						position++
						goto l116
					l115:
						position, tokenIndex = position115, tokenIndex115
					}
				l116:
					goto l86
				l114:
					position, tokenIndex = position86, tokenIndex86
					if buffer[position] != rune('s') {
						goto l84
					}
					position++
					if buffer[position] != rune('t') {
						goto l84
					}
					position++
					if buffer[position] != rune('.') {
						goto l84
					}
					position++
				}
			l86:
				{
					position119, tokenIndex119 := position, tokenIndex
					if !_rules[ruleSpaceCharEOI]() {
						goto l84
					}
					position, tokenIndex = position119, tokenIndex119
				}
				add(ruleRankOther, position85)
			}
			return true
		l84:
			position, tokenIndex = position84, tokenIndex84
			return false
		},
		/* 21 RankVar <- <(('v' 'a' 'r' 'i' 'e' 't' 'y') / ('[' 'v' 'a' 'r' '.' ']') / ('n' 'v' 'a' 'r' '.') / ('v' 'a' 'r' (&SpaceCharEOI / '.')))> */
		func() bool {
			position120, tokenIndex120 := position, tokenIndex
			{
				position121 := position
				{
					position122, tokenIndex122 := position, tokenIndex
					if buffer[position] != rune('v') {
						goto l123
					}
					position++
					if buffer[position] != rune('a') {
						goto l123
					}
					position++
					if buffer[position] != rune('r') {
						goto l123
					}
					position++
					if buffer[position] != rune('i') {
						goto l123
					}
					position++
					if buffer[position] != rune('e') {
						goto l123
					}
					position++
					if buffer[position] != rune('t') {
						goto l123
					}
					position++
					if buffer[position] != rune('y') {
						goto l123
					}
					position++
					goto l122
				l123:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('[') {
						goto l124
					}
					position++
					if buffer[position] != rune('v') {
						goto l124
					}
					position++
					if buffer[position] != rune('a') {
						goto l124
					}
					position++
					if buffer[position] != rune('r') {
						goto l124
					}
					position++
					if buffer[position] != rune('.') {
						goto l124
					}
					position++
					if buffer[position] != rune(']') {
						goto l124
					}
					position++
					goto l122
				l124:
					position, tokenIndex = position122, tokenIndex122
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
					if buffer[position] != rune('v') {
						goto l120
					}
					position++
					if buffer[position] != rune('a') {
						goto l120
					}
					position++
					if buffer[position] != rune('r') {
						goto l120
					}
					position++
					{
						position126, tokenIndex126 := position, tokenIndex
						{
							position128, tokenIndex128 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l127
							}
							position, tokenIndex = position128, tokenIndex128
						}
						goto l126
					l127:
						position, tokenIndex = position126, tokenIndex126
						if buffer[position] != rune('.') {
							goto l120
						}
						position++
					}
				l126:
				}
			l122:
				add(ruleRankVar, position121)
			}
			return true
		l120:
			position, tokenIndex = position120, tokenIndex120
			return false
		},
		/* 22 RankForma <- <((('f' 'o' 'r' 'm' 'a') / ('f' 'm' 'a') / ('f' 'o' 'r' 'm') / ('f' 'o') / 'f') (&SpaceCharEOI / '.'))> */
		func() bool {
			position129, tokenIndex129 := position, tokenIndex
			{
				position130 := position
				{
					position131, tokenIndex131 := position, tokenIndex
					if buffer[position] != rune('f') {
						goto l132
					}
					position++
					if buffer[position] != rune('o') {
						goto l132
					}
					position++
					if buffer[position] != rune('r') {
						goto l132
					}
					position++
					if buffer[position] != rune('m') {
						goto l132
					}
					position++
					if buffer[position] != rune('a') {
						goto l132
					}
					position++
					goto l131
				l132:
					position, tokenIndex = position131, tokenIndex131
					if buffer[position] != rune('f') {
						goto l133
					}
					position++
					if buffer[position] != rune('m') {
						goto l133
					}
					position++
					if buffer[position] != rune('a') {
						goto l133
					}
					position++
					goto l131
				l133:
					position, tokenIndex = position131, tokenIndex131
					if buffer[position] != rune('f') {
						goto l134
					}
					position++
					if buffer[position] != rune('o') {
						goto l134
					}
					position++
					if buffer[position] != rune('r') {
						goto l134
					}
					position++
					if buffer[position] != rune('m') {
						goto l134
					}
					position++
					goto l131
				l134:
					position, tokenIndex = position131, tokenIndex131
					if buffer[position] != rune('f') {
						goto l135
					}
					position++
					if buffer[position] != rune('o') {
						goto l135
					}
					position++
					goto l131
				l135:
					position, tokenIndex = position131, tokenIndex131
					if buffer[position] != rune('f') {
						goto l129
					}
					position++
				}
			l131:
				{
					position136, tokenIndex136 := position, tokenIndex
					{
						position138, tokenIndex138 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l137
						}
						position, tokenIndex = position138, tokenIndex138
					}
					goto l136
				l137:
					position, tokenIndex = position136, tokenIndex136
					if buffer[position] != rune('.') {
						goto l129
					}
					position++
				}
			l136:
				add(ruleRankForma, position130)
			}
			return true
		l129:
			position, tokenIndex = position129, tokenIndex129
			return false
		},
		/* 23 RankSsp <- <((('s' 's' 'p') / ('s' 'u' 'b' 's' 'p')) (&SpaceCharEOI / '.'))> */
		func() bool {
			position139, tokenIndex139 := position, tokenIndex
			{
				position140 := position
				{
					position141, tokenIndex141 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l142
					}
					position++
					if buffer[position] != rune('s') {
						goto l142
					}
					position++
					if buffer[position] != rune('p') {
						goto l142
					}
					position++
					goto l141
				l142:
					position, tokenIndex = position141, tokenIndex141
					if buffer[position] != rune('s') {
						goto l139
					}
					position++
					if buffer[position] != rune('u') {
						goto l139
					}
					position++
					if buffer[position] != rune('b') {
						goto l139
					}
					position++
					if buffer[position] != rune('s') {
						goto l139
					}
					position++
					if buffer[position] != rune('p') {
						goto l139
					}
					position++
				}
			l141:
				{
					position143, tokenIndex143 := position, tokenIndex
					{
						position145, tokenIndex145 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l144
						}
						position, tokenIndex = position145, tokenIndex145
					}
					goto l143
				l144:
					position, tokenIndex = position143, tokenIndex143
					if buffer[position] != rune('.') {
						goto l139
					}
					position++
				}
			l143:
				add(ruleRankSsp, position140)
			}
			return true
		l139:
			position, tokenIndex = position139, tokenIndex139
			return false
		},
		/* 24 SubGenusOrSuperspecies <- <('(' _? Word _? ')')> */
		func() bool {
			position146, tokenIndex146 := position, tokenIndex
			{
				position147 := position
				if buffer[position] != rune('(') {
					goto l146
				}
				position++
				{
					position148, tokenIndex148 := position, tokenIndex
					if !_rules[rule_]() {
						goto l148
					}
					goto l149
				l148:
					position, tokenIndex = position148, tokenIndex148
				}
			l149:
				if !_rules[ruleWord]() {
					goto l146
				}
				{
					position150, tokenIndex150 := position, tokenIndex
					if !_rules[rule_]() {
						goto l150
					}
					goto l151
				l150:
					position, tokenIndex = position150, tokenIndex150
				}
			l151:
				if buffer[position] != rune(')') {
					goto l146
				}
				position++
				add(ruleSubGenusOrSuperspecies, position147)
			}
			return true
		l146:
			position, tokenIndex = position146, tokenIndex146
			return false
		},
		/* 25 SubGenus <- <('(' _? UninomialWord _? ')')> */
		func() bool {
			position152, tokenIndex152 := position, tokenIndex
			{
				position153 := position
				if buffer[position] != rune('(') {
					goto l152
				}
				position++
				{
					position154, tokenIndex154 := position, tokenIndex
					if !_rules[rule_]() {
						goto l154
					}
					goto l155
				l154:
					position, tokenIndex = position154, tokenIndex154
				}
			l155:
				if !_rules[ruleUninomialWord]() {
					goto l152
				}
				{
					position156, tokenIndex156 := position, tokenIndex
					if !_rules[rule_]() {
						goto l156
					}
					goto l157
				l156:
					position, tokenIndex = position156, tokenIndex156
				}
			l157:
				if buffer[position] != rune(')') {
					goto l152
				}
				position++
				add(ruleSubGenus, position153)
			}
			return true
		l152:
			position, tokenIndex = position152, tokenIndex152
			return false
		},
		/* 26 UninomialCombo <- <(UninomialCombo1 / UninomialCombo2)> */
		func() bool {
			position158, tokenIndex158 := position, tokenIndex
			{
				position159 := position
				{
					position160, tokenIndex160 := position, tokenIndex
					if !_rules[ruleUninomialCombo1]() {
						goto l161
					}
					goto l160
				l161:
					position, tokenIndex = position160, tokenIndex160
					if !_rules[ruleUninomialCombo2]() {
						goto l158
					}
				}
			l160:
				add(ruleUninomialCombo, position159)
			}
			return true
		l158:
			position, tokenIndex = position158, tokenIndex158
			return false
		},
		/* 27 UninomialCombo1 <- <(UninomialWord _? SubGenus _? Authorship .?)> */
		func() bool {
			position162, tokenIndex162 := position, tokenIndex
			{
				position163 := position
				if !_rules[ruleUninomialWord]() {
					goto l162
				}
				{
					position164, tokenIndex164 := position, tokenIndex
					if !_rules[rule_]() {
						goto l164
					}
					goto l165
				l164:
					position, tokenIndex = position164, tokenIndex164
				}
			l165:
				if !_rules[ruleSubGenus]() {
					goto l162
				}
				{
					position166, tokenIndex166 := position, tokenIndex
					if !_rules[rule_]() {
						goto l166
					}
					goto l167
				l166:
					position, tokenIndex = position166, tokenIndex166
				}
			l167:
				if !_rules[ruleAuthorship]() {
					goto l162
				}
				{
					position168, tokenIndex168 := position, tokenIndex
					if !matchDot() {
						goto l168
					}
					goto l169
				l168:
					position, tokenIndex = position168, tokenIndex168
				}
			l169:
				add(ruleUninomialCombo1, position163)
			}
			return true
		l162:
			position, tokenIndex = position162, tokenIndex162
			return false
		},
		/* 28 UninomialCombo2 <- <(Uninomial _? RankUninomial _? Uninomial)> */
		func() bool {
			position170, tokenIndex170 := position, tokenIndex
			{
				position171 := position
				if !_rules[ruleUninomial]() {
					goto l170
				}
				{
					position172, tokenIndex172 := position, tokenIndex
					if !_rules[rule_]() {
						goto l172
					}
					goto l173
				l172:
					position, tokenIndex = position172, tokenIndex172
				}
			l173:
				if !_rules[ruleRankUninomial]() {
					goto l170
				}
				{
					position174, tokenIndex174 := position, tokenIndex
					if !_rules[rule_]() {
						goto l174
					}
					goto l175
				l174:
					position, tokenIndex = position174, tokenIndex174
				}
			l175:
				if !_rules[ruleUninomial]() {
					goto l170
				}
				add(ruleUninomialCombo2, position171)
			}
			return true
		l170:
			position, tokenIndex = position170, tokenIndex170
			return false
		},
		/* 29 RankUninomial <- <((('s' 'e' 'c' 't') / ('s' 'u' 'b' 's' 'e' 'c' 't') / ('t' 'r' 'i' 'b') / ('s' 'u' 'b' 't' 'r' 'i' 'b') / ('s' 'u' 'b' 's' 'e' 'r') / ('s' 'e' 'r') / ('s' 'u' 'b' 'g' 'e' 'n') / ('f' 'a' 'm') / ('s' 'u' 'b' 'f' 'a' 'm') / ('s' 'u' 'p' 'e' 'r' 't' 'r' 'i' 'b')) '.'?)> */
		func() bool {
			position176, tokenIndex176 := position, tokenIndex
			{
				position177 := position
				{
					position178, tokenIndex178 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l179
					}
					position++
					if buffer[position] != rune('e') {
						goto l179
					}
					position++
					if buffer[position] != rune('c') {
						goto l179
					}
					position++
					if buffer[position] != rune('t') {
						goto l179
					}
					position++
					goto l178
				l179:
					position, tokenIndex = position178, tokenIndex178
					if buffer[position] != rune('s') {
						goto l180
					}
					position++
					if buffer[position] != rune('u') {
						goto l180
					}
					position++
					if buffer[position] != rune('b') {
						goto l180
					}
					position++
					if buffer[position] != rune('s') {
						goto l180
					}
					position++
					if buffer[position] != rune('e') {
						goto l180
					}
					position++
					if buffer[position] != rune('c') {
						goto l180
					}
					position++
					if buffer[position] != rune('t') {
						goto l180
					}
					position++
					goto l178
				l180:
					position, tokenIndex = position178, tokenIndex178
					if buffer[position] != rune('t') {
						goto l181
					}
					position++
					if buffer[position] != rune('r') {
						goto l181
					}
					position++
					if buffer[position] != rune('i') {
						goto l181
					}
					position++
					if buffer[position] != rune('b') {
						goto l181
					}
					position++
					goto l178
				l181:
					position, tokenIndex = position178, tokenIndex178
					if buffer[position] != rune('s') {
						goto l182
					}
					position++
					if buffer[position] != rune('u') {
						goto l182
					}
					position++
					if buffer[position] != rune('b') {
						goto l182
					}
					position++
					if buffer[position] != rune('t') {
						goto l182
					}
					position++
					if buffer[position] != rune('r') {
						goto l182
					}
					position++
					if buffer[position] != rune('i') {
						goto l182
					}
					position++
					if buffer[position] != rune('b') {
						goto l182
					}
					position++
					goto l178
				l182:
					position, tokenIndex = position178, tokenIndex178
					if buffer[position] != rune('s') {
						goto l183
					}
					position++
					if buffer[position] != rune('u') {
						goto l183
					}
					position++
					if buffer[position] != rune('b') {
						goto l183
					}
					position++
					if buffer[position] != rune('s') {
						goto l183
					}
					position++
					if buffer[position] != rune('e') {
						goto l183
					}
					position++
					if buffer[position] != rune('r') {
						goto l183
					}
					position++
					goto l178
				l183:
					position, tokenIndex = position178, tokenIndex178
					if buffer[position] != rune('s') {
						goto l184
					}
					position++
					if buffer[position] != rune('e') {
						goto l184
					}
					position++
					if buffer[position] != rune('r') {
						goto l184
					}
					position++
					goto l178
				l184:
					position, tokenIndex = position178, tokenIndex178
					if buffer[position] != rune('s') {
						goto l185
					}
					position++
					if buffer[position] != rune('u') {
						goto l185
					}
					position++
					if buffer[position] != rune('b') {
						goto l185
					}
					position++
					if buffer[position] != rune('g') {
						goto l185
					}
					position++
					if buffer[position] != rune('e') {
						goto l185
					}
					position++
					if buffer[position] != rune('n') {
						goto l185
					}
					position++
					goto l178
				l185:
					position, tokenIndex = position178, tokenIndex178
					if buffer[position] != rune('f') {
						goto l186
					}
					position++
					if buffer[position] != rune('a') {
						goto l186
					}
					position++
					if buffer[position] != rune('m') {
						goto l186
					}
					position++
					goto l178
				l186:
					position, tokenIndex = position178, tokenIndex178
					if buffer[position] != rune('s') {
						goto l187
					}
					position++
					if buffer[position] != rune('u') {
						goto l187
					}
					position++
					if buffer[position] != rune('b') {
						goto l187
					}
					position++
					if buffer[position] != rune('f') {
						goto l187
					}
					position++
					if buffer[position] != rune('a') {
						goto l187
					}
					position++
					if buffer[position] != rune('m') {
						goto l187
					}
					position++
					goto l178
				l187:
					position, tokenIndex = position178, tokenIndex178
					if buffer[position] != rune('s') {
						goto l176
					}
					position++
					if buffer[position] != rune('u') {
						goto l176
					}
					position++
					if buffer[position] != rune('p') {
						goto l176
					}
					position++
					if buffer[position] != rune('e') {
						goto l176
					}
					position++
					if buffer[position] != rune('r') {
						goto l176
					}
					position++
					if buffer[position] != rune('t') {
						goto l176
					}
					position++
					if buffer[position] != rune('r') {
						goto l176
					}
					position++
					if buffer[position] != rune('i') {
						goto l176
					}
					position++
					if buffer[position] != rune('b') {
						goto l176
					}
					position++
				}
			l178:
				{
					position188, tokenIndex188 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l188
					}
					position++
					goto l189
				l188:
					position, tokenIndex = position188, tokenIndex188
				}
			l189:
				add(ruleRankUninomial, position177)
			}
			return true
		l176:
			position, tokenIndex = position176, tokenIndex176
			return false
		},
		/* 30 Uninomial <- <(UninomialWord (_ Authorship)?)> */
		func() bool {
			position190, tokenIndex190 := position, tokenIndex
			{
				position191 := position
				if !_rules[ruleUninomialWord]() {
					goto l190
				}
				{
					position192, tokenIndex192 := position, tokenIndex
					if !_rules[rule_]() {
						goto l192
					}
					if !_rules[ruleAuthorship]() {
						goto l192
					}
					goto l193
				l192:
					position, tokenIndex = position192, tokenIndex192
				}
			l193:
				add(ruleUninomial, position191)
			}
			return true
		l190:
			position, tokenIndex = position190, tokenIndex190
			return false
		},
		/* 31 UninomialWord <- <(CapWord / TwoLetterGenus)> */
		func() bool {
			position194, tokenIndex194 := position, tokenIndex
			{
				position195 := position
				{
					position196, tokenIndex196 := position, tokenIndex
					if !_rules[ruleCapWord]() {
						goto l197
					}
					goto l196
				l197:
					position, tokenIndex = position196, tokenIndex196
					if !_rules[ruleTwoLetterGenus]() {
						goto l194
					}
				}
			l196:
				add(ruleUninomialWord, position195)
			}
			return true
		l194:
			position, tokenIndex = position194, tokenIndex194
			return false
		},
		/* 32 AbbrGenus <- <(UpperChar LowerChar* '.')> */
		func() bool {
			position198, tokenIndex198 := position, tokenIndex
			{
				position199 := position
				if !_rules[ruleUpperChar]() {
					goto l198
				}
			l200:
				{
					position201, tokenIndex201 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l201
					}
					goto l200
				l201:
					position, tokenIndex = position201, tokenIndex201
				}
				if buffer[position] != rune('.') {
					goto l198
				}
				position++
				add(ruleAbbrGenus, position199)
			}
			return true
		l198:
			position, tokenIndex = position198, tokenIndex198
			return false
		},
		/* 33 CapWord <- <(CapWord2 / CapWord1)> */
		func() bool {
			position202, tokenIndex202 := position, tokenIndex
			{
				position203 := position
				{
					position204, tokenIndex204 := position, tokenIndex
					if !_rules[ruleCapWord2]() {
						goto l205
					}
					goto l204
				l205:
					position, tokenIndex = position204, tokenIndex204
					if !_rules[ruleCapWord1]() {
						goto l202
					}
				}
			l204:
				add(ruleCapWord, position203)
			}
			return true
		l202:
			position, tokenIndex = position202, tokenIndex202
			return false
		},
		/* 34 CapWord1 <- <(NameUpperChar NameLowerChar NameLowerChar+ '?'?)> */
		func() bool {
			position206, tokenIndex206 := position, tokenIndex
			{
				position207 := position
				if !_rules[ruleNameUpperChar]() {
					goto l206
				}
				if !_rules[ruleNameLowerChar]() {
					goto l206
				}
				if !_rules[ruleNameLowerChar]() {
					goto l206
				}
			l208:
				{
					position209, tokenIndex209 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l209
					}
					goto l208
				l209:
					position, tokenIndex = position209, tokenIndex209
				}
				{
					position210, tokenIndex210 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l210
					}
					position++
					goto l211
				l210:
					position, tokenIndex = position210, tokenIndex210
				}
			l211:
				add(ruleCapWord1, position207)
			}
			return true
		l206:
			position, tokenIndex = position206, tokenIndex206
			return false
		},
		/* 35 CapWord2 <- <(CapWord1 dash (CapWord1 / Word1))> */
		func() bool {
			position212, tokenIndex212 := position, tokenIndex
			{
				position213 := position
				if !_rules[ruleCapWord1]() {
					goto l212
				}
				if !_rules[ruledash]() {
					goto l212
				}
				{
					position214, tokenIndex214 := position, tokenIndex
					if !_rules[ruleCapWord1]() {
						goto l215
					}
					goto l214
				l215:
					position, tokenIndex = position214, tokenIndex214
					if !_rules[ruleWord1]() {
						goto l212
					}
				}
			l214:
				add(ruleCapWord2, position213)
			}
			return true
		l212:
			position, tokenIndex = position212, tokenIndex212
			return false
		},
		/* 36 TwoLetterGenus <- <(('C' 'a') / ('E' 'a') / ('G' 'e') / ('I' 'a') / ('I' 'o') / ('I' 'x') / ('L' 'o') / ('O' 'a') / ('R' 'a') / ('T' 'y') / ('U' 'a') / ('A' 'a') / ('J' 'a') / ('Z' 'u') / ('L' 'a') / ('Q' 'u') / ('A' 's') / ('B' 'a'))> */
		func() bool {
			position216, tokenIndex216 := position, tokenIndex
			{
				position217 := position
				{
					position218, tokenIndex218 := position, tokenIndex
					if buffer[position] != rune('C') {
						goto l219
					}
					position++
					if buffer[position] != rune('a') {
						goto l219
					}
					position++
					goto l218
				l219:
					position, tokenIndex = position218, tokenIndex218
					if buffer[position] != rune('E') {
						goto l220
					}
					position++
					if buffer[position] != rune('a') {
						goto l220
					}
					position++
					goto l218
				l220:
					position, tokenIndex = position218, tokenIndex218
					if buffer[position] != rune('G') {
						goto l221
					}
					position++
					if buffer[position] != rune('e') {
						goto l221
					}
					position++
					goto l218
				l221:
					position, tokenIndex = position218, tokenIndex218
					if buffer[position] != rune('I') {
						goto l222
					}
					position++
					if buffer[position] != rune('a') {
						goto l222
					}
					position++
					goto l218
				l222:
					position, tokenIndex = position218, tokenIndex218
					if buffer[position] != rune('I') {
						goto l223
					}
					position++
					if buffer[position] != rune('o') {
						goto l223
					}
					position++
					goto l218
				l223:
					position, tokenIndex = position218, tokenIndex218
					if buffer[position] != rune('I') {
						goto l224
					}
					position++
					if buffer[position] != rune('x') {
						goto l224
					}
					position++
					goto l218
				l224:
					position, tokenIndex = position218, tokenIndex218
					if buffer[position] != rune('L') {
						goto l225
					}
					position++
					if buffer[position] != rune('o') {
						goto l225
					}
					position++
					goto l218
				l225:
					position, tokenIndex = position218, tokenIndex218
					if buffer[position] != rune('O') {
						goto l226
					}
					position++
					if buffer[position] != rune('a') {
						goto l226
					}
					position++
					goto l218
				l226:
					position, tokenIndex = position218, tokenIndex218
					if buffer[position] != rune('R') {
						goto l227
					}
					position++
					if buffer[position] != rune('a') {
						goto l227
					}
					position++
					goto l218
				l227:
					position, tokenIndex = position218, tokenIndex218
					if buffer[position] != rune('T') {
						goto l228
					}
					position++
					if buffer[position] != rune('y') {
						goto l228
					}
					position++
					goto l218
				l228:
					position, tokenIndex = position218, tokenIndex218
					if buffer[position] != rune('U') {
						goto l229
					}
					position++
					if buffer[position] != rune('a') {
						goto l229
					}
					position++
					goto l218
				l229:
					position, tokenIndex = position218, tokenIndex218
					if buffer[position] != rune('A') {
						goto l230
					}
					position++
					if buffer[position] != rune('a') {
						goto l230
					}
					position++
					goto l218
				l230:
					position, tokenIndex = position218, tokenIndex218
					if buffer[position] != rune('J') {
						goto l231
					}
					position++
					if buffer[position] != rune('a') {
						goto l231
					}
					position++
					goto l218
				l231:
					position, tokenIndex = position218, tokenIndex218
					if buffer[position] != rune('Z') {
						goto l232
					}
					position++
					if buffer[position] != rune('u') {
						goto l232
					}
					position++
					goto l218
				l232:
					position, tokenIndex = position218, tokenIndex218
					if buffer[position] != rune('L') {
						goto l233
					}
					position++
					if buffer[position] != rune('a') {
						goto l233
					}
					position++
					goto l218
				l233:
					position, tokenIndex = position218, tokenIndex218
					if buffer[position] != rune('Q') {
						goto l234
					}
					position++
					if buffer[position] != rune('u') {
						goto l234
					}
					position++
					goto l218
				l234:
					position, tokenIndex = position218, tokenIndex218
					if buffer[position] != rune('A') {
						goto l235
					}
					position++
					if buffer[position] != rune('s') {
						goto l235
					}
					position++
					goto l218
				l235:
					position, tokenIndex = position218, tokenIndex218
					if buffer[position] != rune('B') {
						goto l216
					}
					position++
					if buffer[position] != rune('a') {
						goto l216
					}
					position++
				}
			l218:
				add(ruleTwoLetterGenus, position217)
			}
			return true
		l216:
			position, tokenIndex = position216, tokenIndex216
			return false
		},
		/* 37 Word <- <(!(AuthorPrefix / RankUninomial / Approximation / Word4) (Word3 / Word2StartDigit / Word2 / Word1) &(SpaceCharEOI / '('))> */
		func() bool {
			position236, tokenIndex236 := position, tokenIndex
			{
				position237 := position
				{
					position238, tokenIndex238 := position, tokenIndex
					{
						position239, tokenIndex239 := position, tokenIndex
						if !_rules[ruleAuthorPrefix]() {
							goto l240
						}
						goto l239
					l240:
						position, tokenIndex = position239, tokenIndex239
						if !_rules[ruleRankUninomial]() {
							goto l241
						}
						goto l239
					l241:
						position, tokenIndex = position239, tokenIndex239
						if !_rules[ruleApproximation]() {
							goto l242
						}
						goto l239
					l242:
						position, tokenIndex = position239, tokenIndex239
						if !_rules[ruleWord4]() {
							goto l238
						}
					}
				l239:
					goto l236
				l238:
					position, tokenIndex = position238, tokenIndex238
				}
				{
					position243, tokenIndex243 := position, tokenIndex
					if !_rules[ruleWord3]() {
						goto l244
					}
					goto l243
				l244:
					position, tokenIndex = position243, tokenIndex243
					if !_rules[ruleWord2StartDigit]() {
						goto l245
					}
					goto l243
				l245:
					position, tokenIndex = position243, tokenIndex243
					if !_rules[ruleWord2]() {
						goto l246
					}
					goto l243
				l246:
					position, tokenIndex = position243, tokenIndex243
					if !_rules[ruleWord1]() {
						goto l236
					}
				}
			l243:
				{
					position247, tokenIndex247 := position, tokenIndex
					{
						position248, tokenIndex248 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l249
						}
						goto l248
					l249:
						position, tokenIndex = position248, tokenIndex248
						if buffer[position] != rune('(') {
							goto l236
						}
						position++
					}
				l248:
					position, tokenIndex = position247, tokenIndex247
				}
				add(ruleWord, position237)
			}
			return true
		l236:
			position, tokenIndex = position236, tokenIndex236
			return false
		},
		/* 38 Word1 <- <((lASCII dash)? NameLowerChar NameLowerChar+)> */
		func() bool {
			position250, tokenIndex250 := position, tokenIndex
			{
				position251 := position
				{
					position252, tokenIndex252 := position, tokenIndex
					if !_rules[rulelASCII]() {
						goto l252
					}
					if !_rules[ruledash]() {
						goto l252
					}
					goto l253
				l252:
					position, tokenIndex = position252, tokenIndex252
				}
			l253:
				if !_rules[ruleNameLowerChar]() {
					goto l250
				}
				if !_rules[ruleNameLowerChar]() {
					goto l250
				}
			l254:
				{
					position255, tokenIndex255 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l255
					}
					goto l254
				l255:
					position, tokenIndex = position255, tokenIndex255
				}
				add(ruleWord1, position251)
			}
			return true
		l250:
			position, tokenIndex = position250, tokenIndex250
			return false
		},
		/* 39 Word2StartDigit <- <(('1' / '2' / '3' / '4' / '5' / '6' / '7' / '8' / '9') nums? ('.' / dash)? NameLowerChar NameLowerChar NameLowerChar NameLowerChar+)> */
		func() bool {
			position256, tokenIndex256 := position, tokenIndex
			{
				position257 := position
				{
					position258, tokenIndex258 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l259
					}
					position++
					goto l258
				l259:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('2') {
						goto l260
					}
					position++
					goto l258
				l260:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('3') {
						goto l261
					}
					position++
					goto l258
				l261:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('4') {
						goto l262
					}
					position++
					goto l258
				l262:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('5') {
						goto l263
					}
					position++
					goto l258
				l263:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('6') {
						goto l264
					}
					position++
					goto l258
				l264:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('7') {
						goto l265
					}
					position++
					goto l258
				l265:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('8') {
						goto l266
					}
					position++
					goto l258
				l266:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('9') {
						goto l256
					}
					position++
				}
			l258:
				{
					position267, tokenIndex267 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l267
					}
					goto l268
				l267:
					position, tokenIndex = position267, tokenIndex267
				}
			l268:
				{
					position269, tokenIndex269 := position, tokenIndex
					{
						position271, tokenIndex271 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l272
						}
						position++
						goto l271
					l272:
						position, tokenIndex = position271, tokenIndex271
						if !_rules[ruledash]() {
							goto l269
						}
					}
				l271:
					goto l270
				l269:
					position, tokenIndex = position269, tokenIndex269
				}
			l270:
				if !_rules[ruleNameLowerChar]() {
					goto l256
				}
				if !_rules[ruleNameLowerChar]() {
					goto l256
				}
				if !_rules[ruleNameLowerChar]() {
					goto l256
				}
				if !_rules[ruleNameLowerChar]() {
					goto l256
				}
			l273:
				{
					position274, tokenIndex274 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l274
					}
					goto l273
				l274:
					position, tokenIndex = position274, tokenIndex274
				}
				add(ruleWord2StartDigit, position257)
			}
			return true
		l256:
			position, tokenIndex = position256, tokenIndex256
			return false
		},
		/* 40 Word2 <- <(NameLowerChar+ dash? NameLowerChar)> */
		func() bool {
			position275, tokenIndex275 := position, tokenIndex
			{
				position276 := position
				if !_rules[ruleNameLowerChar]() {
					goto l275
				}
			l277:
				{
					position278, tokenIndex278 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l278
					}
					goto l277
				l278:
					position, tokenIndex = position278, tokenIndex278
				}
				{
					position279, tokenIndex279 := position, tokenIndex
					if !_rules[ruledash]() {
						goto l279
					}
					goto l280
				l279:
					position, tokenIndex = position279, tokenIndex279
				}
			l280:
				if !_rules[ruleNameLowerChar]() {
					goto l275
				}
				add(ruleWord2, position276)
			}
			return true
		l275:
			position, tokenIndex = position275, tokenIndex275
			return false
		},
		/* 41 Word3 <- <(NameLowerChar NameLowerChar* apostr Word1)> */
		func() bool {
			position281, tokenIndex281 := position, tokenIndex
			{
				position282 := position
				if !_rules[ruleNameLowerChar]() {
					goto l281
				}
			l283:
				{
					position284, tokenIndex284 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l284
					}
					goto l283
				l284:
					position, tokenIndex = position284, tokenIndex284
				}
				if !_rules[ruleapostr]() {
					goto l281
				}
				if !_rules[ruleWord1]() {
					goto l281
				}
				add(ruleWord3, position282)
			}
			return true
		l281:
			position, tokenIndex = position281, tokenIndex281
			return false
		},
		/* 42 Word4 <- <(NameLowerChar+ '.' NameLowerChar)> */
		func() bool {
			position285, tokenIndex285 := position, tokenIndex
			{
				position286 := position
				if !_rules[ruleNameLowerChar]() {
					goto l285
				}
			l287:
				{
					position288, tokenIndex288 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l288
					}
					goto l287
				l288:
					position, tokenIndex = position288, tokenIndex288
				}
				if buffer[position] != rune('.') {
					goto l285
				}
				position++
				if !_rules[ruleNameLowerChar]() {
					goto l285
				}
				add(ruleWord4, position286)
			}
			return true
		l285:
			position, tokenIndex = position285, tokenIndex285
			return false
		},
		/* 43 HybridChar <- <'×'> */
		nil,
		/* 44 ApproxName <- <(Uninomial _ (ApproxName1 / ApproxName2))> */
		nil,
		/* 45 ApproxName1 <- <(Approximation ApproxNameIgnored)> */
		nil,
		/* 46 ApproxName2 <- <(Word _ Approximation ApproxNameIgnored)> */
		nil,
		/* 47 ApproxNameIgnored <- <.*> */
		nil,
		/* 48 Approximation <- <(('s' 'p' '.' _? ('n' 'r' '.')) / ('s' 'p' '.' _? ('a' 'f' 'f' '.')) / ('m' 'o' 'n' 's' 't' '.') / '?' / ((('s' 'p' 'p') / ('n' 'r') / ('s' 'p') / ('a' 'f' 'f') / ('s' 'p' 'e' 'c' 'i' 'e' 's')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position294, tokenIndex294 := position, tokenIndex
			{
				position295 := position
				{
					position296, tokenIndex296 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l297
					}
					position++
					if buffer[position] != rune('p') {
						goto l297
					}
					position++
					if buffer[position] != rune('.') {
						goto l297
					}
					position++
					{
						position298, tokenIndex298 := position, tokenIndex
						if !_rules[rule_]() {
							goto l298
						}
						goto l299
					l298:
						position, tokenIndex = position298, tokenIndex298
					}
				l299:
					if buffer[position] != rune('n') {
						goto l297
					}
					position++
					if buffer[position] != rune('r') {
						goto l297
					}
					position++
					if buffer[position] != rune('.') {
						goto l297
					}
					position++
					goto l296
				l297:
					position, tokenIndex = position296, tokenIndex296
					if buffer[position] != rune('s') {
						goto l300
					}
					position++
					if buffer[position] != rune('p') {
						goto l300
					}
					position++
					if buffer[position] != rune('.') {
						goto l300
					}
					position++
					{
						position301, tokenIndex301 := position, tokenIndex
						if !_rules[rule_]() {
							goto l301
						}
						goto l302
					l301:
						position, tokenIndex = position301, tokenIndex301
					}
				l302:
					if buffer[position] != rune('a') {
						goto l300
					}
					position++
					if buffer[position] != rune('f') {
						goto l300
					}
					position++
					if buffer[position] != rune('f') {
						goto l300
					}
					position++
					if buffer[position] != rune('.') {
						goto l300
					}
					position++
					goto l296
				l300:
					position, tokenIndex = position296, tokenIndex296
					if buffer[position] != rune('m') {
						goto l303
					}
					position++
					if buffer[position] != rune('o') {
						goto l303
					}
					position++
					if buffer[position] != rune('n') {
						goto l303
					}
					position++
					if buffer[position] != rune('s') {
						goto l303
					}
					position++
					if buffer[position] != rune('t') {
						goto l303
					}
					position++
					if buffer[position] != rune('.') {
						goto l303
					}
					position++
					goto l296
				l303:
					position, tokenIndex = position296, tokenIndex296
					if buffer[position] != rune('?') {
						goto l304
					}
					position++
					goto l296
				l304:
					position, tokenIndex = position296, tokenIndex296
					{
						position305, tokenIndex305 := position, tokenIndex
						if buffer[position] != rune('s') {
							goto l306
						}
						position++
						if buffer[position] != rune('p') {
							goto l306
						}
						position++
						if buffer[position] != rune('p') {
							goto l306
						}
						position++
						goto l305
					l306:
						position, tokenIndex = position305, tokenIndex305
						if buffer[position] != rune('n') {
							goto l307
						}
						position++
						if buffer[position] != rune('r') {
							goto l307
						}
						position++
						goto l305
					l307:
						position, tokenIndex = position305, tokenIndex305
						if buffer[position] != rune('s') {
							goto l308
						}
						position++
						if buffer[position] != rune('p') {
							goto l308
						}
						position++
						goto l305
					l308:
						position, tokenIndex = position305, tokenIndex305
						if buffer[position] != rune('a') {
							goto l309
						}
						position++
						if buffer[position] != rune('f') {
							goto l309
						}
						position++
						if buffer[position] != rune('f') {
							goto l309
						}
						position++
						goto l305
					l309:
						position, tokenIndex = position305, tokenIndex305
						if buffer[position] != rune('s') {
							goto l294
						}
						position++
						if buffer[position] != rune('p') {
							goto l294
						}
						position++
						if buffer[position] != rune('e') {
							goto l294
						}
						position++
						if buffer[position] != rune('c') {
							goto l294
						}
						position++
						if buffer[position] != rune('i') {
							goto l294
						}
						position++
						if buffer[position] != rune('e') {
							goto l294
						}
						position++
						if buffer[position] != rune('s') {
							goto l294
						}
						position++
					}
				l305:
					{
						position310, tokenIndex310 := position, tokenIndex
						{
							position312, tokenIndex312 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l311
							}
							position, tokenIndex = position312, tokenIndex312
						}
						goto l310
					l311:
						position, tokenIndex = position310, tokenIndex310
						if buffer[position] != rune('.') {
							goto l294
						}
						position++
					}
				l310:
				}
			l296:
				add(ruleApproximation, position295)
			}
			return true
		l294:
			position, tokenIndex = position294, tokenIndex294
			return false
		},
		/* 49 Authorship <- <((Authorship1 / Authorship2) &(SpaceCharEOI / ('\\' / '(' / ',' / ':')))> */
		func() bool {
			position313, tokenIndex313 := position, tokenIndex
			{
				position314 := position
				{
					position315, tokenIndex315 := position, tokenIndex
					if !_rules[ruleAuthorship1]() {
						goto l316
					}
					goto l315
				l316:
					position, tokenIndex = position315, tokenIndex315
					if !_rules[ruleAuthorship2]() {
						goto l313
					}
				}
			l315:
				{
					position317, tokenIndex317 := position, tokenIndex
					{
						position318, tokenIndex318 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l319
						}
						goto l318
					l319:
						position, tokenIndex = position318, tokenIndex318
						{
							position320, tokenIndex320 := position, tokenIndex
							if buffer[position] != rune('\\') {
								goto l321
							}
							position++
							goto l320
						l321:
							position, tokenIndex = position320, tokenIndex320
							if buffer[position] != rune('(') {
								goto l322
							}
							position++
							goto l320
						l322:
							position, tokenIndex = position320, tokenIndex320
							if buffer[position] != rune(',') {
								goto l323
							}
							position++
							goto l320
						l323:
							position, tokenIndex = position320, tokenIndex320
							if buffer[position] != rune(':') {
								goto l313
							}
							position++
						}
					l320:
					}
				l318:
					position, tokenIndex = position317, tokenIndex317
				}
				add(ruleAuthorship, position314)
			}
			return true
		l313:
			position, tokenIndex = position313, tokenIndex313
			return false
		},
		/* 50 Authorship1 <- <(Authorship2 _? AuthorsGroup)> */
		func() bool {
			position324, tokenIndex324 := position, tokenIndex
			{
				position325 := position
				if !_rules[ruleAuthorship2]() {
					goto l324
				}
				{
					position326, tokenIndex326 := position, tokenIndex
					if !_rules[rule_]() {
						goto l326
					}
					goto l327
				l326:
					position, tokenIndex = position326, tokenIndex326
				}
			l327:
				if !_rules[ruleAuthorsGroup]() {
					goto l324
				}
				add(ruleAuthorship1, position325)
			}
			return true
		l324:
			position, tokenIndex = position324, tokenIndex324
			return false
		},
		/* 51 Authorship2 <- <(AuthorsGroup / BasionymAuthorship / BasionymAuthorshipYearMisformed)> */
		func() bool {
			position328, tokenIndex328 := position, tokenIndex
			{
				position329 := position
				{
					position330, tokenIndex330 := position, tokenIndex
					if !_rules[ruleAuthorsGroup]() {
						goto l331
					}
					goto l330
				l331:
					position, tokenIndex = position330, tokenIndex330
					if !_rules[ruleBasionymAuthorship]() {
						goto l332
					}
					goto l330
				l332:
					position, tokenIndex = position330, tokenIndex330
					if !_rules[ruleBasionymAuthorshipYearMisformed]() {
						goto l328
					}
				}
			l330:
				add(ruleAuthorship2, position329)
			}
			return true
		l328:
			position, tokenIndex = position328, tokenIndex328
			return false
		},
		/* 52 BasionymAuthorshipYearMisformed <- <('(' _? AuthorsGroup _? ')' (_? ',')? _? Year)> */
		func() bool {
			position333, tokenIndex333 := position, tokenIndex
			{
				position334 := position
				if buffer[position] != rune('(') {
					goto l333
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
				if !_rules[ruleAuthorsGroup]() {
					goto l333
				}
				{
					position337, tokenIndex337 := position, tokenIndex
					if !_rules[rule_]() {
						goto l337
					}
					goto l338
				l337:
					position, tokenIndex = position337, tokenIndex337
				}
			l338:
				if buffer[position] != rune(')') {
					goto l333
				}
				position++
				{
					position339, tokenIndex339 := position, tokenIndex
					{
						position341, tokenIndex341 := position, tokenIndex
						if !_rules[rule_]() {
							goto l341
						}
						goto l342
					l341:
						position, tokenIndex = position341, tokenIndex341
					}
				l342:
					if buffer[position] != rune(',') {
						goto l339
					}
					position++
					goto l340
				l339:
					position, tokenIndex = position339, tokenIndex339
				}
			l340:
				{
					position343, tokenIndex343 := position, tokenIndex
					if !_rules[rule_]() {
						goto l343
					}
					goto l344
				l343:
					position, tokenIndex = position343, tokenIndex343
				}
			l344:
				if !_rules[ruleYear]() {
					goto l333
				}
				add(ruleBasionymAuthorshipYearMisformed, position334)
			}
			return true
		l333:
			position, tokenIndex = position333, tokenIndex333
			return false
		},
		/* 53 BasionymAuthorship <- <(BasionymAuthorship1 / BasionymAuthorship2)> */
		func() bool {
			position345, tokenIndex345 := position, tokenIndex
			{
				position346 := position
				{
					position347, tokenIndex347 := position, tokenIndex
					if !_rules[ruleBasionymAuthorship1]() {
						goto l348
					}
					goto l347
				l348:
					position, tokenIndex = position347, tokenIndex347
					if !_rules[ruleBasionymAuthorship2]() {
						goto l345
					}
				}
			l347:
				add(ruleBasionymAuthorship, position346)
			}
			return true
		l345:
			position, tokenIndex = position345, tokenIndex345
			return false
		},
		/* 54 BasionymAuthorship1 <- <('(' _? AuthorsGroup _? ')')> */
		func() bool {
			position349, tokenIndex349 := position, tokenIndex
			{
				position350 := position
				if buffer[position] != rune('(') {
					goto l349
				}
				position++
				{
					position351, tokenIndex351 := position, tokenIndex
					if !_rules[rule_]() {
						goto l351
					}
					goto l352
				l351:
					position, tokenIndex = position351, tokenIndex351
				}
			l352:
				if !_rules[ruleAuthorsGroup]() {
					goto l349
				}
				{
					position353, tokenIndex353 := position, tokenIndex
					if !_rules[rule_]() {
						goto l353
					}
					goto l354
				l353:
					position, tokenIndex = position353, tokenIndex353
				}
			l354:
				if buffer[position] != rune(')') {
					goto l349
				}
				position++
				add(ruleBasionymAuthorship1, position350)
			}
			return true
		l349:
			position, tokenIndex = position349, tokenIndex349
			return false
		},
		/* 55 BasionymAuthorship2 <- <('(' _? '(' _? AuthorsGroup _? ')' _? ')')> */
		func() bool {
			position355, tokenIndex355 := position, tokenIndex
			{
				position356 := position
				if buffer[position] != rune('(') {
					goto l355
				}
				position++
				{
					position357, tokenIndex357 := position, tokenIndex
					if !_rules[rule_]() {
						goto l357
					}
					goto l358
				l357:
					position, tokenIndex = position357, tokenIndex357
				}
			l358:
				if buffer[position] != rune('(') {
					goto l355
				}
				position++
				{
					position359, tokenIndex359 := position, tokenIndex
					if !_rules[rule_]() {
						goto l359
					}
					goto l360
				l359:
					position, tokenIndex = position359, tokenIndex359
				}
			l360:
				if !_rules[ruleAuthorsGroup]() {
					goto l355
				}
				{
					position361, tokenIndex361 := position, tokenIndex
					if !_rules[rule_]() {
						goto l361
					}
					goto l362
				l361:
					position, tokenIndex = position361, tokenIndex361
				}
			l362:
				if buffer[position] != rune(')') {
					goto l355
				}
				position++
				{
					position363, tokenIndex363 := position, tokenIndex
					if !_rules[rule_]() {
						goto l363
					}
					goto l364
				l363:
					position, tokenIndex = position363, tokenIndex363
				}
			l364:
				if buffer[position] != rune(')') {
					goto l355
				}
				position++
				add(ruleBasionymAuthorship2, position356)
			}
			return true
		l355:
			position, tokenIndex = position355, tokenIndex355
			return false
		},
		/* 56 AuthorsGroup <- <(AuthorsTeam (_? AuthorEmend? AuthorEx? AuthorsTeam)?)> */
		func() bool {
			position365, tokenIndex365 := position, tokenIndex
			{
				position366 := position
				if !_rules[ruleAuthorsTeam]() {
					goto l365
				}
				{
					position367, tokenIndex367 := position, tokenIndex
					{
						position369, tokenIndex369 := position, tokenIndex
						if !_rules[rule_]() {
							goto l369
						}
						goto l370
					l369:
						position, tokenIndex = position369, tokenIndex369
					}
				l370:
					{
						position371, tokenIndex371 := position, tokenIndex
						if !_rules[ruleAuthorEmend]() {
							goto l371
						}
						goto l372
					l371:
						position, tokenIndex = position371, tokenIndex371
					}
				l372:
					{
						position373, tokenIndex373 := position, tokenIndex
						if !_rules[ruleAuthorEx]() {
							goto l373
						}
						goto l374
					l373:
						position, tokenIndex = position373, tokenIndex373
					}
				l374:
					if !_rules[ruleAuthorsTeam]() {
						goto l367
					}
					goto l368
				l367:
					position, tokenIndex = position367, tokenIndex367
				}
			l368:
				add(ruleAuthorsGroup, position366)
			}
			return true
		l365:
			position, tokenIndex = position365, tokenIndex365
			return false
		},
		/* 57 AuthorsTeam <- <(Author (AuthorSep Author)* (_? ','? _? Year)?)> */
		func() bool {
			position375, tokenIndex375 := position, tokenIndex
			{
				position376 := position
				if !_rules[ruleAuthor]() {
					goto l375
				}
			l377:
				{
					position378, tokenIndex378 := position, tokenIndex
					if !_rules[ruleAuthorSep]() {
						goto l378
					}
					if !_rules[ruleAuthor]() {
						goto l378
					}
					goto l377
				l378:
					position, tokenIndex = position378, tokenIndex378
				}
				{
					position379, tokenIndex379 := position, tokenIndex
					{
						position381, tokenIndex381 := position, tokenIndex
						if !_rules[rule_]() {
							goto l381
						}
						goto l382
					l381:
						position, tokenIndex = position381, tokenIndex381
					}
				l382:
					{
						position383, tokenIndex383 := position, tokenIndex
						if buffer[position] != rune(',') {
							goto l383
						}
						position++
						goto l384
					l383:
						position, tokenIndex = position383, tokenIndex383
					}
				l384:
					{
						position385, tokenIndex385 := position, tokenIndex
						if !_rules[rule_]() {
							goto l385
						}
						goto l386
					l385:
						position, tokenIndex = position385, tokenIndex385
					}
				l386:
					if !_rules[ruleYear]() {
						goto l379
					}
					goto l380
				l379:
					position, tokenIndex = position379, tokenIndex379
				}
			l380:
				add(ruleAuthorsTeam, position376)
			}
			return true
		l375:
			position, tokenIndex = position375, tokenIndex375
			return false
		},
		/* 58 AuthorSep <- <(AuthorSep1 / AuthorSep2)> */
		func() bool {
			position387, tokenIndex387 := position, tokenIndex
			{
				position388 := position
				{
					position389, tokenIndex389 := position, tokenIndex
					if !_rules[ruleAuthorSep1]() {
						goto l390
					}
					goto l389
				l390:
					position, tokenIndex = position389, tokenIndex389
					if !_rules[ruleAuthorSep2]() {
						goto l387
					}
				}
			l389:
				add(ruleAuthorSep, position388)
			}
			return true
		l387:
			position, tokenIndex = position387, tokenIndex387
			return false
		},
		/* 59 AuthorSep1 <- <(_? (',' _)? ('&' / ('e' 't') / ('a' 'n' 'd') / ('a' 'p' 'u' 'd')) _?)> */
		func() bool {
			position391, tokenIndex391 := position, tokenIndex
			{
				position392 := position
				{
					position393, tokenIndex393 := position, tokenIndex
					if !_rules[rule_]() {
						goto l393
					}
					goto l394
				l393:
					position, tokenIndex = position393, tokenIndex393
				}
			l394:
				{
					position395, tokenIndex395 := position, tokenIndex
					if buffer[position] != rune(',') {
						goto l395
					}
					position++
					if !_rules[rule_]() {
						goto l395
					}
					goto l396
				l395:
					position, tokenIndex = position395, tokenIndex395
				}
			l396:
				{
					position397, tokenIndex397 := position, tokenIndex
					if buffer[position] != rune('&') {
						goto l398
					}
					position++
					goto l397
				l398:
					position, tokenIndex = position397, tokenIndex397
					if buffer[position] != rune('e') {
						goto l399
					}
					position++
					if buffer[position] != rune('t') {
						goto l399
					}
					position++
					goto l397
				l399:
					position, tokenIndex = position397, tokenIndex397
					if buffer[position] != rune('a') {
						goto l400
					}
					position++
					if buffer[position] != rune('n') {
						goto l400
					}
					position++
					if buffer[position] != rune('d') {
						goto l400
					}
					position++
					goto l397
				l400:
					position, tokenIndex = position397, tokenIndex397
					if buffer[position] != rune('a') {
						goto l391
					}
					position++
					if buffer[position] != rune('p') {
						goto l391
					}
					position++
					if buffer[position] != rune('u') {
						goto l391
					}
					position++
					if buffer[position] != rune('d') {
						goto l391
					}
					position++
				}
			l397:
				{
					position401, tokenIndex401 := position, tokenIndex
					if !_rules[rule_]() {
						goto l401
					}
					goto l402
				l401:
					position, tokenIndex = position401, tokenIndex401
				}
			l402:
				add(ruleAuthorSep1, position392)
			}
			return true
		l391:
			position, tokenIndex = position391, tokenIndex391
			return false
		},
		/* 60 AuthorSep2 <- <(_? ',' _?)> */
		func() bool {
			position403, tokenIndex403 := position, tokenIndex
			{
				position404 := position
				{
					position405, tokenIndex405 := position, tokenIndex
					if !_rules[rule_]() {
						goto l405
					}
					goto l406
				l405:
					position, tokenIndex = position405, tokenIndex405
				}
			l406:
				if buffer[position] != rune(',') {
					goto l403
				}
				position++
				{
					position407, tokenIndex407 := position, tokenIndex
					if !_rules[rule_]() {
						goto l407
					}
					goto l408
				l407:
					position, tokenIndex = position407, tokenIndex407
				}
			l408:
				add(ruleAuthorSep2, position404)
			}
			return true
		l403:
			position, tokenIndex = position403, tokenIndex403
			return false
		},
		/* 61 AuthorEx <- <((('e' 'x' '.'?) / ('i' 'n')) _)> */
		func() bool {
			position409, tokenIndex409 := position, tokenIndex
			{
				position410 := position
				{
					position411, tokenIndex411 := position, tokenIndex
					if buffer[position] != rune('e') {
						goto l412
					}
					position++
					if buffer[position] != rune('x') {
						goto l412
					}
					position++
					{
						position413, tokenIndex413 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l413
						}
						position++
						goto l414
					l413:
						position, tokenIndex = position413, tokenIndex413
					}
				l414:
					goto l411
				l412:
					position, tokenIndex = position411, tokenIndex411
					if buffer[position] != rune('i') {
						goto l409
					}
					position++
					if buffer[position] != rune('n') {
						goto l409
					}
					position++
				}
			l411:
				if !_rules[rule_]() {
					goto l409
				}
				add(ruleAuthorEx, position410)
			}
			return true
		l409:
			position, tokenIndex = position409, tokenIndex409
			return false
		},
		/* 62 AuthorEmend <- <('e' 'm' 'e' 'n' 'd' '.'? _)> */
		func() bool {
			position415, tokenIndex415 := position, tokenIndex
			{
				position416 := position
				if buffer[position] != rune('e') {
					goto l415
				}
				position++
				if buffer[position] != rune('m') {
					goto l415
				}
				position++
				if buffer[position] != rune('e') {
					goto l415
				}
				position++
				if buffer[position] != rune('n') {
					goto l415
				}
				position++
				if buffer[position] != rune('d') {
					goto l415
				}
				position++
				{
					position417, tokenIndex417 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l417
					}
					position++
					goto l418
				l417:
					position, tokenIndex = position417, tokenIndex417
				}
			l418:
				if !_rules[rule_]() {
					goto l415
				}
				add(ruleAuthorEmend, position416)
			}
			return true
		l415:
			position, tokenIndex = position415, tokenIndex415
			return false
		},
		/* 63 Author <- <(Author1 / Author2 / UnknownAuthor)> */
		func() bool {
			position419, tokenIndex419 := position, tokenIndex
			{
				position420 := position
				{
					position421, tokenIndex421 := position, tokenIndex
					if !_rules[ruleAuthor1]() {
						goto l422
					}
					goto l421
				l422:
					position, tokenIndex = position421, tokenIndex421
					if !_rules[ruleAuthor2]() {
						goto l423
					}
					goto l421
				l423:
					position, tokenIndex = position421, tokenIndex421
					if !_rules[ruleUnknownAuthor]() {
						goto l419
					}
				}
			l421:
				add(ruleAuthor, position420)
			}
			return true
		l419:
			position, tokenIndex = position419, tokenIndex419
			return false
		},
		/* 64 Author1 <- <(Author2 _? Filius)> */
		func() bool {
			position424, tokenIndex424 := position, tokenIndex
			{
				position425 := position
				if !_rules[ruleAuthor2]() {
					goto l424
				}
				{
					position426, tokenIndex426 := position, tokenIndex
					if !_rules[rule_]() {
						goto l426
					}
					goto l427
				l426:
					position, tokenIndex = position426, tokenIndex426
				}
			l427:
				if !_rules[ruleFilius]() {
					goto l424
				}
				add(ruleAuthor1, position425)
			}
			return true
		l424:
			position, tokenIndex = position424, tokenIndex424
			return false
		},
		/* 65 Author2 <- <(AuthorWord (_? AuthorWord)*)> */
		func() bool {
			position428, tokenIndex428 := position, tokenIndex
			{
				position429 := position
				if !_rules[ruleAuthorWord]() {
					goto l428
				}
			l430:
				{
					position431, tokenIndex431 := position, tokenIndex
					{
						position432, tokenIndex432 := position, tokenIndex
						if !_rules[rule_]() {
							goto l432
						}
						goto l433
					l432:
						position, tokenIndex = position432, tokenIndex432
					}
				l433:
					if !_rules[ruleAuthorWord]() {
						goto l431
					}
					goto l430
				l431:
					position, tokenIndex = position431, tokenIndex431
				}
				add(ruleAuthor2, position429)
			}
			return true
		l428:
			position, tokenIndex = position428, tokenIndex428
			return false
		},
		/* 66 UnknownAuthor <- <('?' / ((('a' 'u' 'c' 't') / ('a' 'n' 'o' 'n')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position434, tokenIndex434 := position, tokenIndex
			{
				position435 := position
				{
					position436, tokenIndex436 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l437
					}
					position++
					goto l436
				l437:
					position, tokenIndex = position436, tokenIndex436
					{
						position438, tokenIndex438 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l439
						}
						position++
						if buffer[position] != rune('u') {
							goto l439
						}
						position++
						if buffer[position] != rune('c') {
							goto l439
						}
						position++
						if buffer[position] != rune('t') {
							goto l439
						}
						position++
						goto l438
					l439:
						position, tokenIndex = position438, tokenIndex438
						if buffer[position] != rune('a') {
							goto l434
						}
						position++
						if buffer[position] != rune('n') {
							goto l434
						}
						position++
						if buffer[position] != rune('o') {
							goto l434
						}
						position++
						if buffer[position] != rune('n') {
							goto l434
						}
						position++
					}
				l438:
					{
						position440, tokenIndex440 := position, tokenIndex
						{
							position442, tokenIndex442 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l441
							}
							position, tokenIndex = position442, tokenIndex442
						}
						goto l440
					l441:
						position, tokenIndex = position440, tokenIndex440
						if buffer[position] != rune('.') {
							goto l434
						}
						position++
					}
				l440:
				}
			l436:
				add(ruleUnknownAuthor, position435)
			}
			return true
		l434:
			position, tokenIndex = position434, tokenIndex434
			return false
		},
		/* 67 AuthorWord <- <(AuthorWord1 / AuthorWord2 / AuthorWord3 / AuthorPrefix)> */
		func() bool {
			position443, tokenIndex443 := position, tokenIndex
			{
				position444 := position
				{
					position445, tokenIndex445 := position, tokenIndex
					if !_rules[ruleAuthorWord1]() {
						goto l446
					}
					goto l445
				l446:
					position, tokenIndex = position445, tokenIndex445
					if !_rules[ruleAuthorWord2]() {
						goto l447
					}
					goto l445
				l447:
					position, tokenIndex = position445, tokenIndex445
					if !_rules[ruleAuthorWord3]() {
						goto l448
					}
					goto l445
				l448:
					position, tokenIndex = position445, tokenIndex445
					if !_rules[ruleAuthorPrefix]() {
						goto l443
					}
				}
			l445:
				add(ruleAuthorWord, position444)
			}
			return true
		l443:
			position, tokenIndex = position443, tokenIndex443
			return false
		},
		/* 68 AuthorWord1 <- <(('a' 'r' 'g' '.') / ('e' 't' ' ' 'a' 'l' '.' '{' '?' '}') / ((('e' 't') / '&') (' ' 'a' 'l') '.'?))> */
		func() bool {
			position449, tokenIndex449 := position, tokenIndex
			{
				position450 := position
				{
					position451, tokenIndex451 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l452
					}
					position++
					if buffer[position] != rune('r') {
						goto l452
					}
					position++
					if buffer[position] != rune('g') {
						goto l452
					}
					position++
					if buffer[position] != rune('.') {
						goto l452
					}
					position++
					goto l451
				l452:
					position, tokenIndex = position451, tokenIndex451
					if buffer[position] != rune('e') {
						goto l453
					}
					position++
					if buffer[position] != rune('t') {
						goto l453
					}
					position++
					if buffer[position] != rune(' ') {
						goto l453
					}
					position++
					if buffer[position] != rune('a') {
						goto l453
					}
					position++
					if buffer[position] != rune('l') {
						goto l453
					}
					position++
					if buffer[position] != rune('.') {
						goto l453
					}
					position++
					if buffer[position] != rune('{') {
						goto l453
					}
					position++
					if buffer[position] != rune('?') {
						goto l453
					}
					position++
					if buffer[position] != rune('}') {
						goto l453
					}
					position++
					goto l451
				l453:
					position, tokenIndex = position451, tokenIndex451
					{
						position454, tokenIndex454 := position, tokenIndex
						if buffer[position] != rune('e') {
							goto l455
						}
						position++
						if buffer[position] != rune('t') {
							goto l455
						}
						position++
						goto l454
					l455:
						position, tokenIndex = position454, tokenIndex454
						if buffer[position] != rune('&') {
							goto l449
						}
						position++
					}
				l454:
					if buffer[position] != rune(' ') {
						goto l449
					}
					position++
					if buffer[position] != rune('a') {
						goto l449
					}
					position++
					if buffer[position] != rune('l') {
						goto l449
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
				}
			l451:
				add(ruleAuthorWord1, position450)
			}
			return true
		l449:
			position, tokenIndex = position449, tokenIndex449
			return false
		},
		/* 69 AuthorWord2 <- <(AuthorWord3 dash AuthorWordSoft)> */
		func() bool {
			position458, tokenIndex458 := position, tokenIndex
			{
				position459 := position
				if !_rules[ruleAuthorWord3]() {
					goto l458
				}
				if !_rules[ruledash]() {
					goto l458
				}
				if !_rules[ruleAuthorWordSoft]() {
					goto l458
				}
				add(ruleAuthorWord2, position459)
			}
			return true
		l458:
			position, tokenIndex = position458, tokenIndex458
			return false
		},
		/* 70 AuthorWord3 <- <(AuthorPrefixGlued? AuthorUpperChar (AuthorUpperChar / AuthorLowerChar)* '.'?)> */
		func() bool {
			position460, tokenIndex460 := position, tokenIndex
			{
				position461 := position
				{
					position462, tokenIndex462 := position, tokenIndex
					if !_rules[ruleAuthorPrefixGlued]() {
						goto l462
					}
					goto l463
				l462:
					position, tokenIndex = position462, tokenIndex462
				}
			l463:
				if !_rules[ruleAuthorUpperChar]() {
					goto l460
				}
			l464:
				{
					position465, tokenIndex465 := position, tokenIndex
					{
						position466, tokenIndex466 := position, tokenIndex
						if !_rules[ruleAuthorUpperChar]() {
							goto l467
						}
						goto l466
					l467:
						position, tokenIndex = position466, tokenIndex466
						if !_rules[ruleAuthorLowerChar]() {
							goto l465
						}
					}
				l466:
					goto l464
				l465:
					position, tokenIndex = position465, tokenIndex465
				}
				{
					position468, tokenIndex468 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l468
					}
					position++
					goto l469
				l468:
					position, tokenIndex = position468, tokenIndex468
				}
			l469:
				add(ruleAuthorWord3, position461)
			}
			return true
		l460:
			position, tokenIndex = position460, tokenIndex460
			return false
		},
		/* 71 AuthorWordSoft <- <(((AuthorUpperChar (AuthorUpperChar+ / AuthorLowerChar+)) / AuthorLowerChar+) '.'?)> */
		func() bool {
			position470, tokenIndex470 := position, tokenIndex
			{
				position471 := position
				{
					position472, tokenIndex472 := position, tokenIndex
					if !_rules[ruleAuthorUpperChar]() {
						goto l473
					}
					{
						position474, tokenIndex474 := position, tokenIndex
						if !_rules[ruleAuthorUpperChar]() {
							goto l475
						}
					l476:
						{
							position477, tokenIndex477 := position, tokenIndex
							if !_rules[ruleAuthorUpperChar]() {
								goto l477
							}
							goto l476
						l477:
							position, tokenIndex = position477, tokenIndex477
						}
						goto l474
					l475:
						position, tokenIndex = position474, tokenIndex474
						if !_rules[ruleAuthorLowerChar]() {
							goto l473
						}
					l478:
						{
							position479, tokenIndex479 := position, tokenIndex
							if !_rules[ruleAuthorLowerChar]() {
								goto l479
							}
							goto l478
						l479:
							position, tokenIndex = position479, tokenIndex479
						}
					}
				l474:
					goto l472
				l473:
					position, tokenIndex = position472, tokenIndex472
					if !_rules[ruleAuthorLowerChar]() {
						goto l470
					}
				l480:
					{
						position481, tokenIndex481 := position, tokenIndex
						if !_rules[ruleAuthorLowerChar]() {
							goto l481
						}
						goto l480
					l481:
						position, tokenIndex = position481, tokenIndex481
					}
				}
			l472:
				{
					position482, tokenIndex482 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l482
					}
					position++
					goto l483
				l482:
					position, tokenIndex = position482, tokenIndex482
				}
			l483:
				add(ruleAuthorWordSoft, position471)
			}
			return true
		l470:
			position, tokenIndex = position470, tokenIndex470
			return false
		},
		/* 72 Filius <- <(('f' '.') / ('f' 'i' 'l' '.') / ('f' 'i' 'l' 'i' 'u' 's'))> */
		func() bool {
			position484, tokenIndex484 := position, tokenIndex
			{
				position485 := position
				{
					position486, tokenIndex486 := position, tokenIndex
					if buffer[position] != rune('f') {
						goto l487
					}
					position++
					if buffer[position] != rune('.') {
						goto l487
					}
					position++
					goto l486
				l487:
					position, tokenIndex = position486, tokenIndex486
					if buffer[position] != rune('f') {
						goto l488
					}
					position++
					if buffer[position] != rune('i') {
						goto l488
					}
					position++
					if buffer[position] != rune('l') {
						goto l488
					}
					position++
					if buffer[position] != rune('.') {
						goto l488
					}
					position++
					goto l486
				l488:
					position, tokenIndex = position486, tokenIndex486
					if buffer[position] != rune('f') {
						goto l484
					}
					position++
					if buffer[position] != rune('i') {
						goto l484
					}
					position++
					if buffer[position] != rune('l') {
						goto l484
					}
					position++
					if buffer[position] != rune('i') {
						goto l484
					}
					position++
					if buffer[position] != rune('u') {
						goto l484
					}
					position++
					if buffer[position] != rune('s') {
						goto l484
					}
					position++
				}
			l486:
				add(ruleFilius, position485)
			}
			return true
		l484:
			position, tokenIndex = position484, tokenIndex484
			return false
		},
		/* 73 AuthorPrefixGlued <- <(('d' '\'') / ('O' '\''))> */
		func() bool {
			position489, tokenIndex489 := position, tokenIndex
			{
				position490 := position
				{
					position491, tokenIndex491 := position, tokenIndex
					if buffer[position] != rune('d') {
						goto l492
					}
					position++
					if buffer[position] != rune('\'') {
						goto l492
					}
					position++
					goto l491
				l492:
					position, tokenIndex = position491, tokenIndex491
					if buffer[position] != rune('O') {
						goto l489
					}
					position++
					if buffer[position] != rune('\'') {
						goto l489
					}
					position++
				}
			l491:
				add(ruleAuthorPrefixGlued, position490)
			}
			return true
		l489:
			position, tokenIndex = position489, tokenIndex489
			return false
		},
		/* 74 AuthorPrefix <- <(AuthorPrefix1 / AuthorPrefix2)> */
		func() bool {
			position493, tokenIndex493 := position, tokenIndex
			{
				position494 := position
				{
					position495, tokenIndex495 := position, tokenIndex
					if !_rules[ruleAuthorPrefix1]() {
						goto l496
					}
					goto l495
				l496:
					position, tokenIndex = position495, tokenIndex495
					if !_rules[ruleAuthorPrefix2]() {
						goto l493
					}
				}
			l495:
				add(ruleAuthorPrefix, position494)
			}
			return true
		l493:
			position, tokenIndex = position493, tokenIndex493
			return false
		},
		/* 75 AuthorPrefix2 <- <(('v' '.' (_? ('d' '.'))?) / ('\'' 't'))> */
		func() bool {
			position497, tokenIndex497 := position, tokenIndex
			{
				position498 := position
				{
					position499, tokenIndex499 := position, tokenIndex
					if buffer[position] != rune('v') {
						goto l500
					}
					position++
					if buffer[position] != rune('.') {
						goto l500
					}
					position++
					{
						position501, tokenIndex501 := position, tokenIndex
						{
							position503, tokenIndex503 := position, tokenIndex
							if !_rules[rule_]() {
								goto l503
							}
							goto l504
						l503:
							position, tokenIndex = position503, tokenIndex503
						}
					l504:
						if buffer[position] != rune('d') {
							goto l501
						}
						position++
						if buffer[position] != rune('.') {
							goto l501
						}
						position++
						goto l502
					l501:
						position, tokenIndex = position501, tokenIndex501
					}
				l502:
					goto l499
				l500:
					position, tokenIndex = position499, tokenIndex499
					if buffer[position] != rune('\'') {
						goto l497
					}
					position++
					if buffer[position] != rune('t') {
						goto l497
					}
					position++
				}
			l499:
				add(ruleAuthorPrefix2, position498)
			}
			return true
		l497:
			position, tokenIndex = position497, tokenIndex497
			return false
		},
		/* 76 AuthorPrefix1 <- <((('a' 'b') / ('a' 'f') / ('b' 'i' 's') / ('d' 'a') / ('d' 'e' 'r') / ('d' 'e' 's') / ('d' 'e' 'n') / ('d' 'e' 'l') / ('d' 'e' 'l' 'l' 'a') / ('d' 'e' 'l' 'a') / ('d' 'e') / ('d' 'i') / ('d' 'u') / ('e' 'l') / ('l' 'a') / ('l' 'e') / ('t' 'e' 'r') / ('v' 'a' 'n') / ('d' '\'') / ('i' 'n' '\'' 't') / ('z' 'u' 'r') / ('v' 'o' 'n' (_ ('d' '.'))?) / ('v' (_ 'd')?)) &_)> */
		func() bool {
			position505, tokenIndex505 := position, tokenIndex
			{
				position506 := position
				{
					position507, tokenIndex507 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l508
					}
					position++
					if buffer[position] != rune('b') {
						goto l508
					}
					position++
					goto l507
				l508:
					position, tokenIndex = position507, tokenIndex507
					if buffer[position] != rune('a') {
						goto l509
					}
					position++
					if buffer[position] != rune('f') {
						goto l509
					}
					position++
					goto l507
				l509:
					position, tokenIndex = position507, tokenIndex507
					if buffer[position] != rune('b') {
						goto l510
					}
					position++
					if buffer[position] != rune('i') {
						goto l510
					}
					position++
					if buffer[position] != rune('s') {
						goto l510
					}
					position++
					goto l507
				l510:
					position, tokenIndex = position507, tokenIndex507
					if buffer[position] != rune('d') {
						goto l511
					}
					position++
					if buffer[position] != rune('a') {
						goto l511
					}
					position++
					goto l507
				l511:
					position, tokenIndex = position507, tokenIndex507
					if buffer[position] != rune('d') {
						goto l512
					}
					position++
					if buffer[position] != rune('e') {
						goto l512
					}
					position++
					if buffer[position] != rune('r') {
						goto l512
					}
					position++
					goto l507
				l512:
					position, tokenIndex = position507, tokenIndex507
					if buffer[position] != rune('d') {
						goto l513
					}
					position++
					if buffer[position] != rune('e') {
						goto l513
					}
					position++
					if buffer[position] != rune('s') {
						goto l513
					}
					position++
					goto l507
				l513:
					position, tokenIndex = position507, tokenIndex507
					if buffer[position] != rune('d') {
						goto l514
					}
					position++
					if buffer[position] != rune('e') {
						goto l514
					}
					position++
					if buffer[position] != rune('n') {
						goto l514
					}
					position++
					goto l507
				l514:
					position, tokenIndex = position507, tokenIndex507
					if buffer[position] != rune('d') {
						goto l515
					}
					position++
					if buffer[position] != rune('e') {
						goto l515
					}
					position++
					if buffer[position] != rune('l') {
						goto l515
					}
					position++
					goto l507
				l515:
					position, tokenIndex = position507, tokenIndex507
					if buffer[position] != rune('d') {
						goto l516
					}
					position++
					if buffer[position] != rune('e') {
						goto l516
					}
					position++
					if buffer[position] != rune('l') {
						goto l516
					}
					position++
					if buffer[position] != rune('l') {
						goto l516
					}
					position++
					if buffer[position] != rune('a') {
						goto l516
					}
					position++
					goto l507
				l516:
					position, tokenIndex = position507, tokenIndex507
					if buffer[position] != rune('d') {
						goto l517
					}
					position++
					if buffer[position] != rune('e') {
						goto l517
					}
					position++
					if buffer[position] != rune('l') {
						goto l517
					}
					position++
					if buffer[position] != rune('a') {
						goto l517
					}
					position++
					goto l507
				l517:
					position, tokenIndex = position507, tokenIndex507
					if buffer[position] != rune('d') {
						goto l518
					}
					position++
					if buffer[position] != rune('e') {
						goto l518
					}
					position++
					goto l507
				l518:
					position, tokenIndex = position507, tokenIndex507
					if buffer[position] != rune('d') {
						goto l519
					}
					position++
					if buffer[position] != rune('i') {
						goto l519
					}
					position++
					goto l507
				l519:
					position, tokenIndex = position507, tokenIndex507
					if buffer[position] != rune('d') {
						goto l520
					}
					position++
					if buffer[position] != rune('u') {
						goto l520
					}
					position++
					goto l507
				l520:
					position, tokenIndex = position507, tokenIndex507
					if buffer[position] != rune('e') {
						goto l521
					}
					position++
					if buffer[position] != rune('l') {
						goto l521
					}
					position++
					goto l507
				l521:
					position, tokenIndex = position507, tokenIndex507
					if buffer[position] != rune('l') {
						goto l522
					}
					position++
					if buffer[position] != rune('a') {
						goto l522
					}
					position++
					goto l507
				l522:
					position, tokenIndex = position507, tokenIndex507
					if buffer[position] != rune('l') {
						goto l523
					}
					position++
					if buffer[position] != rune('e') {
						goto l523
					}
					position++
					goto l507
				l523:
					position, tokenIndex = position507, tokenIndex507
					if buffer[position] != rune('t') {
						goto l524
					}
					position++
					if buffer[position] != rune('e') {
						goto l524
					}
					position++
					if buffer[position] != rune('r') {
						goto l524
					}
					position++
					goto l507
				l524:
					position, tokenIndex = position507, tokenIndex507
					if buffer[position] != rune('v') {
						goto l525
					}
					position++
					if buffer[position] != rune('a') {
						goto l525
					}
					position++
					if buffer[position] != rune('n') {
						goto l525
					}
					position++
					goto l507
				l525:
					position, tokenIndex = position507, tokenIndex507
					if buffer[position] != rune('d') {
						goto l526
					}
					position++
					if buffer[position] != rune('\'') {
						goto l526
					}
					position++
					goto l507
				l526:
					position, tokenIndex = position507, tokenIndex507
					if buffer[position] != rune('i') {
						goto l527
					}
					position++
					if buffer[position] != rune('n') {
						goto l527
					}
					position++
					if buffer[position] != rune('\'') {
						goto l527
					}
					position++
					if buffer[position] != rune('t') {
						goto l527
					}
					position++
					goto l507
				l527:
					position, tokenIndex = position507, tokenIndex507
					if buffer[position] != rune('z') {
						goto l528
					}
					position++
					if buffer[position] != rune('u') {
						goto l528
					}
					position++
					if buffer[position] != rune('r') {
						goto l528
					}
					position++
					goto l507
				l528:
					position, tokenIndex = position507, tokenIndex507
					if buffer[position] != rune('v') {
						goto l529
					}
					position++
					if buffer[position] != rune('o') {
						goto l529
					}
					position++
					if buffer[position] != rune('n') {
						goto l529
					}
					position++
					{
						position530, tokenIndex530 := position, tokenIndex
						if !_rules[rule_]() {
							goto l530
						}
						if buffer[position] != rune('d') {
							goto l530
						}
						position++
						if buffer[position] != rune('.') {
							goto l530
						}
						position++
						goto l531
					l530:
						position, tokenIndex = position530, tokenIndex530
					}
				l531:
					goto l507
				l529:
					position, tokenIndex = position507, tokenIndex507
					if buffer[position] != rune('v') {
						goto l505
					}
					position++
					{
						position532, tokenIndex532 := position, tokenIndex
						if !_rules[rule_]() {
							goto l532
						}
						if buffer[position] != rune('d') {
							goto l532
						}
						position++
						goto l533
					l532:
						position, tokenIndex = position532, tokenIndex532
					}
				l533:
				}
			l507:
				{
					position534, tokenIndex534 := position, tokenIndex
					if !_rules[rule_]() {
						goto l505
					}
					position, tokenIndex = position534, tokenIndex534
				}
				add(ruleAuthorPrefix1, position506)
			}
			return true
		l505:
			position, tokenIndex = position505, tokenIndex505
			return false
		},
		/* 77 AuthorUpperChar <- <(hASCII / ('À' / 'Á' / 'Â' / 'Ã' / 'Ä' / 'Å' / 'Æ' / 'Ç' / 'È' / 'É' / 'Ê' / 'Ë' / 'Ì' / 'Í' / 'Î' / 'Ï' / 'Ð' / 'Ñ' / 'Ò' / 'Ó' / 'Ô' / 'Õ' / 'Ö' / 'Ø' / 'Ù' / 'Ú' / 'Û' / 'Ü' / 'Ý' / 'Ć' / 'Č' / 'Ď' / 'İ' / 'Ķ' / 'Ĺ' / 'ĺ' / 'Ľ' / 'ľ' / 'Ł' / 'ł' / 'Ņ' / 'Ō' / 'Ő' / 'Œ' / 'Ř' / 'Ś' / 'Ŝ' / 'Ş' / 'Š' / 'Ÿ' / 'Ź' / 'Ż' / 'Ž' / 'ƒ' / 'Ǿ' / 'Ș' / 'Ț'))> */
		func() bool {
			position535, tokenIndex535 := position, tokenIndex
			{
				position536 := position
				{
					position537, tokenIndex537 := position, tokenIndex
					if !_rules[rulehASCII]() {
						goto l538
					}
					goto l537
				l538:
					position, tokenIndex = position537, tokenIndex537
					{
						position539, tokenIndex539 := position, tokenIndex
						if buffer[position] != rune('À') {
							goto l540
						}
						position++
						goto l539
					l540:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Á') {
							goto l541
						}
						position++
						goto l539
					l541:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Â') {
							goto l542
						}
						position++
						goto l539
					l542:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ã') {
							goto l543
						}
						position++
						goto l539
					l543:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ä') {
							goto l544
						}
						position++
						goto l539
					l544:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Å') {
							goto l545
						}
						position++
						goto l539
					l545:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Æ') {
							goto l546
						}
						position++
						goto l539
					l546:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ç') {
							goto l547
						}
						position++
						goto l539
					l547:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('È') {
							goto l548
						}
						position++
						goto l539
					l548:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('É') {
							goto l549
						}
						position++
						goto l539
					l549:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ê') {
							goto l550
						}
						position++
						goto l539
					l550:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ë') {
							goto l551
						}
						position++
						goto l539
					l551:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ì') {
							goto l552
						}
						position++
						goto l539
					l552:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Í') {
							goto l553
						}
						position++
						goto l539
					l553:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Î') {
							goto l554
						}
						position++
						goto l539
					l554:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ï') {
							goto l555
						}
						position++
						goto l539
					l555:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ð') {
							goto l556
						}
						position++
						goto l539
					l556:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ñ') {
							goto l557
						}
						position++
						goto l539
					l557:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ò') {
							goto l558
						}
						position++
						goto l539
					l558:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ó') {
							goto l559
						}
						position++
						goto l539
					l559:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ô') {
							goto l560
						}
						position++
						goto l539
					l560:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Õ') {
							goto l561
						}
						position++
						goto l539
					l561:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ö') {
							goto l562
						}
						position++
						goto l539
					l562:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ø') {
							goto l563
						}
						position++
						goto l539
					l563:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ù') {
							goto l564
						}
						position++
						goto l539
					l564:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ú') {
							goto l565
						}
						position++
						goto l539
					l565:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Û') {
							goto l566
						}
						position++
						goto l539
					l566:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ü') {
							goto l567
						}
						position++
						goto l539
					l567:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ý') {
							goto l568
						}
						position++
						goto l539
					l568:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ć') {
							goto l569
						}
						position++
						goto l539
					l569:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Č') {
							goto l570
						}
						position++
						goto l539
					l570:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ď') {
							goto l571
						}
						position++
						goto l539
					l571:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('İ') {
							goto l572
						}
						position++
						goto l539
					l572:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ķ') {
							goto l573
						}
						position++
						goto l539
					l573:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ĺ') {
							goto l574
						}
						position++
						goto l539
					l574:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('ĺ') {
							goto l575
						}
						position++
						goto l539
					l575:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ľ') {
							goto l576
						}
						position++
						goto l539
					l576:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('ľ') {
							goto l577
						}
						position++
						goto l539
					l577:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ł') {
							goto l578
						}
						position++
						goto l539
					l578:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('ł') {
							goto l579
						}
						position++
						goto l539
					l579:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ņ') {
							goto l580
						}
						position++
						goto l539
					l580:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ō') {
							goto l581
						}
						position++
						goto l539
					l581:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ő') {
							goto l582
						}
						position++
						goto l539
					l582:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Œ') {
							goto l583
						}
						position++
						goto l539
					l583:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ř') {
							goto l584
						}
						position++
						goto l539
					l584:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ś') {
							goto l585
						}
						position++
						goto l539
					l585:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ŝ') {
							goto l586
						}
						position++
						goto l539
					l586:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ş') {
							goto l587
						}
						position++
						goto l539
					l587:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Š') {
							goto l588
						}
						position++
						goto l539
					l588:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ÿ') {
							goto l589
						}
						position++
						goto l539
					l589:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ź') {
							goto l590
						}
						position++
						goto l539
					l590:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ż') {
							goto l591
						}
						position++
						goto l539
					l591:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ž') {
							goto l592
						}
						position++
						goto l539
					l592:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('ƒ') {
							goto l593
						}
						position++
						goto l539
					l593:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ǿ') {
							goto l594
						}
						position++
						goto l539
					l594:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ș') {
							goto l595
						}
						position++
						goto l539
					l595:
						position, tokenIndex = position539, tokenIndex539
						if buffer[position] != rune('Ț') {
							goto l535
						}
						position++
					}
				l539:
				}
			l537:
				add(ruleAuthorUpperChar, position536)
			}
			return true
		l535:
			position, tokenIndex = position535, tokenIndex535
			return false
		},
		/* 78 AuthorLowerChar <- <(lASCII / ('à' / 'á' / 'â' / 'ã' / 'ä' / 'å' / 'æ' / 'ç' / 'è' / 'é' / 'ê' / 'ë' / 'ì' / 'í' / 'î' / 'ï' / 'ð' / 'ñ' / 'ò' / 'ó' / 'ó' / 'ô' / 'õ' / 'ö' / 'ø' / 'ù' / 'ú' / 'û' / 'ü' / 'ý' / 'ÿ' / 'ā' / 'ă' / 'ą' / 'ć' / 'ĉ' / 'č' / 'ď' / 'đ' / '\'' / 'ē' / 'ĕ' / 'ė' / 'ę' / 'ě' / 'ğ' / 'ī' / 'ĭ' / 'İ' / 'ı' / 'ĺ' / 'ľ' / 'ł' / 'ń' / 'ņ' / 'ň' / 'ŏ' / 'ő' / 'œ' / 'ŕ' / 'ř' / 'ś' / 'ş' / 'š' / 'ţ' / 'ť' / 'ũ' / 'ū' / 'ŭ' / 'ů' / 'ű' / 'ź' / 'ż' / 'ž' / 'ſ' / 'ǎ' / 'ǔ' / 'ǧ' / 'ș' / 'ț' / 'ȳ' / 'ß'))> */
		func() bool {
			position596, tokenIndex596 := position, tokenIndex
			{
				position597 := position
				{
					position598, tokenIndex598 := position, tokenIndex
					if !_rules[rulelASCII]() {
						goto l599
					}
					goto l598
				l599:
					position, tokenIndex = position598, tokenIndex598
					{
						position600, tokenIndex600 := position, tokenIndex
						if buffer[position] != rune('à') {
							goto l601
						}
						position++
						goto l600
					l601:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('á') {
							goto l602
						}
						position++
						goto l600
					l602:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('â') {
							goto l603
						}
						position++
						goto l600
					l603:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ã') {
							goto l604
						}
						position++
						goto l600
					l604:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ä') {
							goto l605
						}
						position++
						goto l600
					l605:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('å') {
							goto l606
						}
						position++
						goto l600
					l606:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('æ') {
							goto l607
						}
						position++
						goto l600
					l607:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ç') {
							goto l608
						}
						position++
						goto l600
					l608:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('è') {
							goto l609
						}
						position++
						goto l600
					l609:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('é') {
							goto l610
						}
						position++
						goto l600
					l610:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ê') {
							goto l611
						}
						position++
						goto l600
					l611:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ë') {
							goto l612
						}
						position++
						goto l600
					l612:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ì') {
							goto l613
						}
						position++
						goto l600
					l613:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('í') {
							goto l614
						}
						position++
						goto l600
					l614:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('î') {
							goto l615
						}
						position++
						goto l600
					l615:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ï') {
							goto l616
						}
						position++
						goto l600
					l616:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ð') {
							goto l617
						}
						position++
						goto l600
					l617:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ñ') {
							goto l618
						}
						position++
						goto l600
					l618:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ò') {
							goto l619
						}
						position++
						goto l600
					l619:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ó') {
							goto l620
						}
						position++
						goto l600
					l620:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ó') {
							goto l621
						}
						position++
						goto l600
					l621:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ô') {
							goto l622
						}
						position++
						goto l600
					l622:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('õ') {
							goto l623
						}
						position++
						goto l600
					l623:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ö') {
							goto l624
						}
						position++
						goto l600
					l624:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ø') {
							goto l625
						}
						position++
						goto l600
					l625:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ù') {
							goto l626
						}
						position++
						goto l600
					l626:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ú') {
							goto l627
						}
						position++
						goto l600
					l627:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('û') {
							goto l628
						}
						position++
						goto l600
					l628:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ü') {
							goto l629
						}
						position++
						goto l600
					l629:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ý') {
							goto l630
						}
						position++
						goto l600
					l630:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ÿ') {
							goto l631
						}
						position++
						goto l600
					l631:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ā') {
							goto l632
						}
						position++
						goto l600
					l632:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ă') {
							goto l633
						}
						position++
						goto l600
					l633:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ą') {
							goto l634
						}
						position++
						goto l600
					l634:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ć') {
							goto l635
						}
						position++
						goto l600
					l635:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ĉ') {
							goto l636
						}
						position++
						goto l600
					l636:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('č') {
							goto l637
						}
						position++
						goto l600
					l637:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ď') {
							goto l638
						}
						position++
						goto l600
					l638:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('đ') {
							goto l639
						}
						position++
						goto l600
					l639:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('\'') {
							goto l640
						}
						position++
						goto l600
					l640:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ē') {
							goto l641
						}
						position++
						goto l600
					l641:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ĕ') {
							goto l642
						}
						position++
						goto l600
					l642:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ė') {
							goto l643
						}
						position++
						goto l600
					l643:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ę') {
							goto l644
						}
						position++
						goto l600
					l644:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ě') {
							goto l645
						}
						position++
						goto l600
					l645:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ğ') {
							goto l646
						}
						position++
						goto l600
					l646:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ī') {
							goto l647
						}
						position++
						goto l600
					l647:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ĭ') {
							goto l648
						}
						position++
						goto l600
					l648:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('İ') {
							goto l649
						}
						position++
						goto l600
					l649:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ı') {
							goto l650
						}
						position++
						goto l600
					l650:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ĺ') {
							goto l651
						}
						position++
						goto l600
					l651:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ľ') {
							goto l652
						}
						position++
						goto l600
					l652:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ł') {
							goto l653
						}
						position++
						goto l600
					l653:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ń') {
							goto l654
						}
						position++
						goto l600
					l654:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ņ') {
							goto l655
						}
						position++
						goto l600
					l655:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ň') {
							goto l656
						}
						position++
						goto l600
					l656:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ŏ') {
							goto l657
						}
						position++
						goto l600
					l657:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ő') {
							goto l658
						}
						position++
						goto l600
					l658:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('œ') {
							goto l659
						}
						position++
						goto l600
					l659:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ŕ') {
							goto l660
						}
						position++
						goto l600
					l660:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ř') {
							goto l661
						}
						position++
						goto l600
					l661:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ś') {
							goto l662
						}
						position++
						goto l600
					l662:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ş') {
							goto l663
						}
						position++
						goto l600
					l663:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('š') {
							goto l664
						}
						position++
						goto l600
					l664:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ţ') {
							goto l665
						}
						position++
						goto l600
					l665:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ť') {
							goto l666
						}
						position++
						goto l600
					l666:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ũ') {
							goto l667
						}
						position++
						goto l600
					l667:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ū') {
							goto l668
						}
						position++
						goto l600
					l668:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ŭ') {
							goto l669
						}
						position++
						goto l600
					l669:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ů') {
							goto l670
						}
						position++
						goto l600
					l670:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ű') {
							goto l671
						}
						position++
						goto l600
					l671:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ź') {
							goto l672
						}
						position++
						goto l600
					l672:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ż') {
							goto l673
						}
						position++
						goto l600
					l673:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ž') {
							goto l674
						}
						position++
						goto l600
					l674:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ſ') {
							goto l675
						}
						position++
						goto l600
					l675:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ǎ') {
							goto l676
						}
						position++
						goto l600
					l676:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ǔ') {
							goto l677
						}
						position++
						goto l600
					l677:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ǧ') {
							goto l678
						}
						position++
						goto l600
					l678:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ș') {
							goto l679
						}
						position++
						goto l600
					l679:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ț') {
							goto l680
						}
						position++
						goto l600
					l680:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ȳ') {
							goto l681
						}
						position++
						goto l600
					l681:
						position, tokenIndex = position600, tokenIndex600
						if buffer[position] != rune('ß') {
							goto l596
						}
						position++
					}
				l600:
				}
			l598:
				add(ruleAuthorLowerChar, position597)
			}
			return true
		l596:
			position, tokenIndex = position596, tokenIndex596
			return false
		},
		/* 79 Year <- <(YearRange / YearApprox / YearWithParens / YearWithPage / YearWithDot / YearWithChar / YearNum)> */
		func() bool {
			position682, tokenIndex682 := position, tokenIndex
			{
				position683 := position
				{
					position684, tokenIndex684 := position, tokenIndex
					if !_rules[ruleYearRange]() {
						goto l685
					}
					goto l684
				l685:
					position, tokenIndex = position684, tokenIndex684
					if !_rules[ruleYearApprox]() {
						goto l686
					}
					goto l684
				l686:
					position, tokenIndex = position684, tokenIndex684
					if !_rules[ruleYearWithParens]() {
						goto l687
					}
					goto l684
				l687:
					position, tokenIndex = position684, tokenIndex684
					if !_rules[ruleYearWithPage]() {
						goto l688
					}
					goto l684
				l688:
					position, tokenIndex = position684, tokenIndex684
					if !_rules[ruleYearWithDot]() {
						goto l689
					}
					goto l684
				l689:
					position, tokenIndex = position684, tokenIndex684
					if !_rules[ruleYearWithChar]() {
						goto l690
					}
					goto l684
				l690:
					position, tokenIndex = position684, tokenIndex684
					if !_rules[ruleYearNum]() {
						goto l682
					}
				}
			l684:
				add(ruleYear, position683)
			}
			return true
		l682:
			position, tokenIndex = position682, tokenIndex682
			return false
		},
		/* 80 YearRange <- <(YearNum dash (nums+ ('a' / 'b' / 'c' / 'd' / 'e' / 'f' / 'g' / 'h' / 'i' / 'j' / 'k' / 'l' / 'm' / 'n' / 'o' / 'p' / 'q' / 'r' / 's' / 't' / 'u' / 'v' / 'w' / 'x' / 'y' / 'z' / '?')*))> */
		func() bool {
			position691, tokenIndex691 := position, tokenIndex
			{
				position692 := position
				if !_rules[ruleYearNum]() {
					goto l691
				}
				if !_rules[ruledash]() {
					goto l691
				}
				if !_rules[rulenums]() {
					goto l691
				}
			l693:
				{
					position694, tokenIndex694 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l694
					}
					goto l693
				l694:
					position, tokenIndex = position694, tokenIndex694
				}
			l695:
				{
					position696, tokenIndex696 := position, tokenIndex
					{
						position697, tokenIndex697 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l698
						}
						position++
						goto l697
					l698:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('b') {
							goto l699
						}
						position++
						goto l697
					l699:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('c') {
							goto l700
						}
						position++
						goto l697
					l700:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('d') {
							goto l701
						}
						position++
						goto l697
					l701:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('e') {
							goto l702
						}
						position++
						goto l697
					l702:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('f') {
							goto l703
						}
						position++
						goto l697
					l703:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('g') {
							goto l704
						}
						position++
						goto l697
					l704:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('h') {
							goto l705
						}
						position++
						goto l697
					l705:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('i') {
							goto l706
						}
						position++
						goto l697
					l706:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('j') {
							goto l707
						}
						position++
						goto l697
					l707:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('k') {
							goto l708
						}
						position++
						goto l697
					l708:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('l') {
							goto l709
						}
						position++
						goto l697
					l709:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('m') {
							goto l710
						}
						position++
						goto l697
					l710:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('n') {
							goto l711
						}
						position++
						goto l697
					l711:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('o') {
							goto l712
						}
						position++
						goto l697
					l712:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('p') {
							goto l713
						}
						position++
						goto l697
					l713:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('q') {
							goto l714
						}
						position++
						goto l697
					l714:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('r') {
							goto l715
						}
						position++
						goto l697
					l715:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('s') {
							goto l716
						}
						position++
						goto l697
					l716:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('t') {
							goto l717
						}
						position++
						goto l697
					l717:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('u') {
							goto l718
						}
						position++
						goto l697
					l718:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('v') {
							goto l719
						}
						position++
						goto l697
					l719:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('w') {
							goto l720
						}
						position++
						goto l697
					l720:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('x') {
							goto l721
						}
						position++
						goto l697
					l721:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('y') {
							goto l722
						}
						position++
						goto l697
					l722:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('z') {
							goto l723
						}
						position++
						goto l697
					l723:
						position, tokenIndex = position697, tokenIndex697
						if buffer[position] != rune('?') {
							goto l696
						}
						position++
					}
				l697:
					goto l695
				l696:
					position, tokenIndex = position696, tokenIndex696
				}
				add(ruleYearRange, position692)
			}
			return true
		l691:
			position, tokenIndex = position691, tokenIndex691
			return false
		},
		/* 81 YearWithDot <- <(YearNum '.')> */
		func() bool {
			position724, tokenIndex724 := position, tokenIndex
			{
				position725 := position
				if !_rules[ruleYearNum]() {
					goto l724
				}
				if buffer[position] != rune('.') {
					goto l724
				}
				position++
				add(ruleYearWithDot, position725)
			}
			return true
		l724:
			position, tokenIndex = position724, tokenIndex724
			return false
		},
		/* 82 YearApprox <- <('[' _? YearNum _? ']')> */
		func() bool {
			position726, tokenIndex726 := position, tokenIndex
			{
				position727 := position
				if buffer[position] != rune('[') {
					goto l726
				}
				position++
				{
					position728, tokenIndex728 := position, tokenIndex
					if !_rules[rule_]() {
						goto l728
					}
					goto l729
				l728:
					position, tokenIndex = position728, tokenIndex728
				}
			l729:
				if !_rules[ruleYearNum]() {
					goto l726
				}
				{
					position730, tokenIndex730 := position, tokenIndex
					if !_rules[rule_]() {
						goto l730
					}
					goto l731
				l730:
					position, tokenIndex = position730, tokenIndex730
				}
			l731:
				if buffer[position] != rune(']') {
					goto l726
				}
				position++
				add(ruleYearApprox, position727)
			}
			return true
		l726:
			position, tokenIndex = position726, tokenIndex726
			return false
		},
		/* 83 YearWithPage <- <((YearWithChar / YearNum) _? ':' _? nums+)> */
		func() bool {
			position732, tokenIndex732 := position, tokenIndex
			{
				position733 := position
				{
					position734, tokenIndex734 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l735
					}
					goto l734
				l735:
					position, tokenIndex = position734, tokenIndex734
					if !_rules[ruleYearNum]() {
						goto l732
					}
				}
			l734:
				{
					position736, tokenIndex736 := position, tokenIndex
					if !_rules[rule_]() {
						goto l736
					}
					goto l737
				l736:
					position, tokenIndex = position736, tokenIndex736
				}
			l737:
				if buffer[position] != rune(':') {
					goto l732
				}
				position++
				{
					position738, tokenIndex738 := position, tokenIndex
					if !_rules[rule_]() {
						goto l738
					}
					goto l739
				l738:
					position, tokenIndex = position738, tokenIndex738
				}
			l739:
				if !_rules[rulenums]() {
					goto l732
				}
			l740:
				{
					position741, tokenIndex741 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l741
					}
					goto l740
				l741:
					position, tokenIndex = position741, tokenIndex741
				}
				add(ruleYearWithPage, position733)
			}
			return true
		l732:
			position, tokenIndex = position732, tokenIndex732
			return false
		},
		/* 84 YearWithParens <- <('(' (YearWithChar / YearNum) ')')> */
		func() bool {
			position742, tokenIndex742 := position, tokenIndex
			{
				position743 := position
				if buffer[position] != rune('(') {
					goto l742
				}
				position++
				{
					position744, tokenIndex744 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l745
					}
					goto l744
				l745:
					position, tokenIndex = position744, tokenIndex744
					if !_rules[ruleYearNum]() {
						goto l742
					}
				}
			l744:
				if buffer[position] != rune(')') {
					goto l742
				}
				position++
				add(ruleYearWithParens, position743)
			}
			return true
		l742:
			position, tokenIndex = position742, tokenIndex742
			return false
		},
		/* 85 YearWithChar <- <(YearNum lASCII Action0)> */
		func() bool {
			position746, tokenIndex746 := position, tokenIndex
			{
				position747 := position
				if !_rules[ruleYearNum]() {
					goto l746
				}
				if !_rules[rulelASCII]() {
					goto l746
				}
				if !_rules[ruleAction0]() {
					goto l746
				}
				add(ruleYearWithChar, position747)
			}
			return true
		l746:
			position, tokenIndex = position746, tokenIndex746
			return false
		},
		/* 86 YearNum <- <(('1' / '2') ('0' / '7' / '8' / '9') nums (nums / '?') '?'*)> */
		func() bool {
			position748, tokenIndex748 := position, tokenIndex
			{
				position749 := position
				{
					position750, tokenIndex750 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l751
					}
					position++
					goto l750
				l751:
					position, tokenIndex = position750, tokenIndex750
					if buffer[position] != rune('2') {
						goto l748
					}
					position++
				}
			l750:
				{
					position752, tokenIndex752 := position, tokenIndex
					if buffer[position] != rune('0') {
						goto l753
					}
					position++
					goto l752
				l753:
					position, tokenIndex = position752, tokenIndex752
					if buffer[position] != rune('7') {
						goto l754
					}
					position++
					goto l752
				l754:
					position, tokenIndex = position752, tokenIndex752
					if buffer[position] != rune('8') {
						goto l755
					}
					position++
					goto l752
				l755:
					position, tokenIndex = position752, tokenIndex752
					if buffer[position] != rune('9') {
						goto l748
					}
					position++
				}
			l752:
				if !_rules[rulenums]() {
					goto l748
				}
				{
					position756, tokenIndex756 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l757
					}
					goto l756
				l757:
					position, tokenIndex = position756, tokenIndex756
					if buffer[position] != rune('?') {
						goto l748
					}
					position++
				}
			l756:
			l758:
				{
					position759, tokenIndex759 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l759
					}
					position++
					goto l758
				l759:
					position, tokenIndex = position759, tokenIndex759
				}
				add(ruleYearNum, position749)
			}
			return true
		l748:
			position, tokenIndex = position748, tokenIndex748
			return false
		},
		/* 87 NameUpperChar <- <(UpperChar / UpperCharExtended)> */
		func() bool {
			position760, tokenIndex760 := position, tokenIndex
			{
				position761 := position
				{
					position762, tokenIndex762 := position, tokenIndex
					if !_rules[ruleUpperChar]() {
						goto l763
					}
					goto l762
				l763:
					position, tokenIndex = position762, tokenIndex762
					if !_rules[ruleUpperCharExtended]() {
						goto l760
					}
				}
			l762:
				add(ruleNameUpperChar, position761)
			}
			return true
		l760:
			position, tokenIndex = position760, tokenIndex760
			return false
		},
		/* 88 UpperCharExtended <- <('Æ' / 'Œ' / 'Ö')> */
		func() bool {
			position764, tokenIndex764 := position, tokenIndex
			{
				position765 := position
				{
					position766, tokenIndex766 := position, tokenIndex
					if buffer[position] != rune('Æ') {
						goto l767
					}
					position++
					goto l766
				l767:
					position, tokenIndex = position766, tokenIndex766
					if buffer[position] != rune('Œ') {
						goto l768
					}
					position++
					goto l766
				l768:
					position, tokenIndex = position766, tokenIndex766
					if buffer[position] != rune('Ö') {
						goto l764
					}
					position++
				}
			l766:
				add(ruleUpperCharExtended, position765)
			}
			return true
		l764:
			position, tokenIndex = position764, tokenIndex764
			return false
		},
		/* 89 UpperChar <- <hASCII> */
		func() bool {
			position769, tokenIndex769 := position, tokenIndex
			{
				position770 := position
				if !_rules[rulehASCII]() {
					goto l769
				}
				add(ruleUpperChar, position770)
			}
			return true
		l769:
			position, tokenIndex = position769, tokenIndex769
			return false
		},
		/* 90 NameLowerChar <- <(LowerChar / LowerCharExtended / MiscodedChar)> */
		func() bool {
			position771, tokenIndex771 := position, tokenIndex
			{
				position772 := position
				{
					position773, tokenIndex773 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l774
					}
					goto l773
				l774:
					position, tokenIndex = position773, tokenIndex773
					if !_rules[ruleLowerCharExtended]() {
						goto l775
					}
					goto l773
				l775:
					position, tokenIndex = position773, tokenIndex773
					if !_rules[ruleMiscodedChar]() {
						goto l771
					}
				}
			l773:
				add(ruleNameLowerChar, position772)
			}
			return true
		l771:
			position, tokenIndex = position771, tokenIndex771
			return false
		},
		/* 91 MiscodedChar <- <'�'> */
		func() bool {
			position776, tokenIndex776 := position, tokenIndex
			{
				position777 := position
				if buffer[position] != rune('�') {
					goto l776
				}
				position++
				add(ruleMiscodedChar, position777)
			}
			return true
		l776:
			position, tokenIndex = position776, tokenIndex776
			return false
		},
		/* 92 LowerCharExtended <- <('æ' / 'œ' / 'ſ' / 'à' / 'â' / 'å' / 'ã' / 'ä' / 'á' / 'ç' / 'č' / 'é' / 'è' / 'í' / 'ì' / 'ï' / 'ň' / 'ñ' / 'ñ' / 'ó' / 'ò' / 'ô' / 'ø' / 'õ' / 'ö' / 'ú' / 'ù' / 'ü' / 'ŕ' / 'ř' / 'ŗ' / 'š' / 'š' / 'ş' / 'ž')> */
		func() bool {
			position778, tokenIndex778 := position, tokenIndex
			{
				position779 := position
				{
					position780, tokenIndex780 := position, tokenIndex
					if buffer[position] != rune('æ') {
						goto l781
					}
					position++
					goto l780
				l781:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('œ') {
						goto l782
					}
					position++
					goto l780
				l782:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('ſ') {
						goto l783
					}
					position++
					goto l780
				l783:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('à') {
						goto l784
					}
					position++
					goto l780
				l784:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('â') {
						goto l785
					}
					position++
					goto l780
				l785:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('å') {
						goto l786
					}
					position++
					goto l780
				l786:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('ã') {
						goto l787
					}
					position++
					goto l780
				l787:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('ä') {
						goto l788
					}
					position++
					goto l780
				l788:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('á') {
						goto l789
					}
					position++
					goto l780
				l789:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('ç') {
						goto l790
					}
					position++
					goto l780
				l790:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('č') {
						goto l791
					}
					position++
					goto l780
				l791:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('é') {
						goto l792
					}
					position++
					goto l780
				l792:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('è') {
						goto l793
					}
					position++
					goto l780
				l793:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('í') {
						goto l794
					}
					position++
					goto l780
				l794:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('ì') {
						goto l795
					}
					position++
					goto l780
				l795:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('ï') {
						goto l796
					}
					position++
					goto l780
				l796:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('ň') {
						goto l797
					}
					position++
					goto l780
				l797:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('ñ') {
						goto l798
					}
					position++
					goto l780
				l798:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('ñ') {
						goto l799
					}
					position++
					goto l780
				l799:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('ó') {
						goto l800
					}
					position++
					goto l780
				l800:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('ò') {
						goto l801
					}
					position++
					goto l780
				l801:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('ô') {
						goto l802
					}
					position++
					goto l780
				l802:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('ø') {
						goto l803
					}
					position++
					goto l780
				l803:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('õ') {
						goto l804
					}
					position++
					goto l780
				l804:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('ö') {
						goto l805
					}
					position++
					goto l780
				l805:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('ú') {
						goto l806
					}
					position++
					goto l780
				l806:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('ù') {
						goto l807
					}
					position++
					goto l780
				l807:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('ü') {
						goto l808
					}
					position++
					goto l780
				l808:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('ŕ') {
						goto l809
					}
					position++
					goto l780
				l809:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('ř') {
						goto l810
					}
					position++
					goto l780
				l810:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('ŗ') {
						goto l811
					}
					position++
					goto l780
				l811:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('š') {
						goto l812
					}
					position++
					goto l780
				l812:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('š') {
						goto l813
					}
					position++
					goto l780
				l813:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('ş') {
						goto l814
					}
					position++
					goto l780
				l814:
					position, tokenIndex = position780, tokenIndex780
					if buffer[position] != rune('ž') {
						goto l778
					}
					position++
				}
			l780:
				add(ruleLowerCharExtended, position779)
			}
			return true
		l778:
			position, tokenIndex = position778, tokenIndex778
			return false
		},
		/* 93 LowerChar <- <([a-z] / 'ë')> */
		func() bool {
			position815, tokenIndex815 := position, tokenIndex
			{
				position816 := position
				{
					position817, tokenIndex817 := position, tokenIndex
					if c := buffer[position]; c < rune('a') || c > rune('z') {
						goto l818
					}
					position++
					goto l817
				l818:
					position, tokenIndex = position817, tokenIndex817
					if buffer[position] != rune('ë') {
						goto l815
					}
					position++
				}
			l817:
				add(ruleLowerChar, position816)
			}
			return true
		l815:
			position, tokenIndex = position815, tokenIndex815
			return false
		},
		/* 94 SpaceCharEOI <- <(_ / !.)> */
		func() bool {
			position819, tokenIndex819 := position, tokenIndex
			{
				position820 := position
				{
					position821, tokenIndex821 := position, tokenIndex
					if !_rules[rule_]() {
						goto l822
					}
					goto l821
				l822:
					position, tokenIndex = position821, tokenIndex821
					{
						position823, tokenIndex823 := position, tokenIndex
						if !matchDot() {
							goto l823
						}
						goto l819
					l823:
						position, tokenIndex = position823, tokenIndex823
					}
				}
			l821:
				add(ruleSpaceCharEOI, position820)
			}
			return true
		l819:
			position, tokenIndex = position819, tokenIndex819
			return false
		},
		/* 95 WordBorderChar <- <(_ / (';' / '.' / ',' / ';' / '(' / ')'))> */
		nil,
		/* 96 nums <- <[0-9]> */
		func() bool {
			position825, tokenIndex825 := position, tokenIndex
			{
				position826 := position
				if c := buffer[position]; c < rune('0') || c > rune('9') {
					goto l825
				}
				position++
				add(rulenums, position826)
			}
			return true
		l825:
			position, tokenIndex = position825, tokenIndex825
			return false
		},
		/* 97 lASCII <- <[a-z]> */
		func() bool {
			position827, tokenIndex827 := position, tokenIndex
			{
				position828 := position
				if c := buffer[position]; c < rune('a') || c > rune('z') {
					goto l827
				}
				position++
				add(rulelASCII, position828)
			}
			return true
		l827:
			position, tokenIndex = position827, tokenIndex827
			return false
		},
		/* 98 hASCII <- <[A-Z]> */
		func() bool {
			position829, tokenIndex829 := position, tokenIndex
			{
				position830 := position
				if c := buffer[position]; c < rune('A') || c > rune('Z') {
					goto l829
				}
				position++
				add(rulehASCII, position830)
			}
			return true
		l829:
			position, tokenIndex = position829, tokenIndex829
			return false
		},
		/* 99 apostr <- <'\''> */
		func() bool {
			position831, tokenIndex831 := position, tokenIndex
			{
				position832 := position
				if buffer[position] != rune('\'') {
					goto l831
				}
				position++
				add(ruleapostr, position832)
			}
			return true
		l831:
			position, tokenIndex = position831, tokenIndex831
			return false
		},
		/* 100 dash <- <'-'> */
		func() bool {
			position833, tokenIndex833 := position, tokenIndex
			{
				position834 := position
				if buffer[position] != rune('-') {
					goto l833
				}
				position++
				add(ruledash, position834)
			}
			return true
		l833:
			position, tokenIndex = position833, tokenIndex833
			return false
		},
		/* 101 _ <- <' '+> */
		func() bool {
			position835, tokenIndex835 := position, tokenIndex
			{
				position836 := position
				if buffer[position] != rune(' ') {
					goto l835
				}
				position++
			l837:
				{
					position838, tokenIndex838 := position, tokenIndex
					if buffer[position] != rune(' ') {
						goto l838
					}
					position++
					goto l837
				l838:
					position, tokenIndex = position838, tokenIndex838
				}
				add(rule_, position836)
			}
			return true
		l835:
			position, tokenIndex = position835, tokenIndex835
			return false
		},
		/* 103 Action0 <- <{ p.addWarn(YearCharWarn) }> */
		func() bool {
			{
				add(ruleAction0, position)
			}
			return true
		},
	}
	p.rules = _rules
}
