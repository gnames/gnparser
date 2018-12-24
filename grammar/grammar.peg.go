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
)

var rul3s = [...]string{
	"Unknown",
	"SciName",
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
	Buffer string
	buffer []rune
	rules  [102]func() bool
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
		/* 0 SciName <- <(_? SciName1 .* !.)> */
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
			l4:
				{
					position5, tokenIndex5 := position, tokenIndex
					if !matchDot() {
						goto l5
					}
					goto l4
				l5:
					position, tokenIndex = position5, tokenIndex5
				}
				{
					position6, tokenIndex6 := position, tokenIndex
					if !matchDot() {
						goto l6
					}
					goto l0
				l6:
					position, tokenIndex = position6, tokenIndex6
				}
				add(ruleSciName, position1)
			}
			return true
		l0:
			position, tokenIndex = position0, tokenIndex0
			return false
		},
		/* 1 SciName1 <- <SciName2> */
		func() bool {
			position7, tokenIndex7 := position, tokenIndex
			{
				position8 := position
				if !_rules[ruleSciName2]() {
					goto l7
				}
				add(ruleSciName1, position8)
			}
			return true
		l7:
			position, tokenIndex = position7, tokenIndex7
			return false
		},
		/* 2 SciName2 <- <Name> */
		func() bool {
			position9, tokenIndex9 := position, tokenIndex
			{
				position10 := position
				if !_rules[ruleName]() {
					goto l9
				}
				add(ruleSciName2, position10)
			}
			return true
		l9:
			position, tokenIndex = position9, tokenIndex9
			return false
		},
		/* 3 HybridFormula <- <(Name (_ (HybridFormula1 / HybridFormula2)))> */
		nil,
		/* 4 HybridFormula1 <- <(HybridChar _? SpeciesEpithet (_ InfraspGroup)?)> */
		nil,
		/* 5 HybridFormula2 <- <(HybridChar (_ Name)?)> */
		nil,
		/* 6 NamedHybrid <- <(HybridChar _? Name)> */
		nil,
		/* 7 Name <- <(NameSpecies / NameUninomial)> */
		func() bool {
			position15, tokenIndex15 := position, tokenIndex
			{
				position16 := position
				{
					position17, tokenIndex17 := position, tokenIndex
					if !_rules[ruleNameSpecies]() {
						goto l18
					}
					goto l17
				l18:
					position, tokenIndex = position17, tokenIndex17
					if !_rules[ruleNameUninomial]() {
						goto l15
					}
				}
			l17:
				add(ruleName, position16)
			}
			return true
		l15:
			position, tokenIndex = position15, tokenIndex15
			return false
		},
		/* 8 NameUninomial <- <(UninomialCombo / Uninomial)> */
		func() bool {
			position19, tokenIndex19 := position, tokenIndex
			{
				position20 := position
				{
					position21, tokenIndex21 := position, tokenIndex
					if !_rules[ruleUninomialCombo]() {
						goto l22
					}
					goto l21
				l22:
					position, tokenIndex = position21, tokenIndex21
					if !_rules[ruleUninomial]() {
						goto l19
					}
				}
			l21:
				add(ruleNameUninomial, position20)
			}
			return true
		l19:
			position, tokenIndex = position19, tokenIndex19
			return false
		},
		/* 9 NameApprox <- <(GenusWord _ Approximation (_ SpeciesEpithet)?)> */
		nil,
		/* 10 NameComp <- <(GenusWord _ Comparison (_ SpeciesEpithet)?)> */
		nil,
		/* 11 NameSpecies <- <(GenusWord (_? (SubGenus / SubGenusOrSuperspecies))? _ SpeciesEpithet (_ InfraspGroup)?)> */
		func() bool {
			position25, tokenIndex25 := position, tokenIndex
			{
				position26 := position
				if !_rules[ruleGenusWord]() {
					goto l25
				}
				{
					position27, tokenIndex27 := position, tokenIndex
					{
						position29, tokenIndex29 := position, tokenIndex
						if !_rules[rule_]() {
							goto l29
						}
						goto l30
					l29:
						position, tokenIndex = position29, tokenIndex29
					}
				l30:
					{
						position31, tokenIndex31 := position, tokenIndex
						if !_rules[ruleSubGenus]() {
							goto l32
						}
						goto l31
					l32:
						position, tokenIndex = position31, tokenIndex31
						if !_rules[ruleSubGenusOrSuperspecies]() {
							goto l27
						}
					}
				l31:
					goto l28
				l27:
					position, tokenIndex = position27, tokenIndex27
				}
			l28:
				if !_rules[rule_]() {
					goto l25
				}
				if !_rules[ruleSpeciesEpithet]() {
					goto l25
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
				add(ruleNameSpecies, position26)
			}
			return true
		l25:
			position, tokenIndex = position25, tokenIndex25
			return false
		},
		/* 12 GenusWord <- <((AbbrGenus / UninomialWord) !(_ AuthorWord))> */
		func() bool {
			position35, tokenIndex35 := position, tokenIndex
			{
				position36 := position
				{
					position37, tokenIndex37 := position, tokenIndex
					if !_rules[ruleAbbrGenus]() {
						goto l38
					}
					goto l37
				l38:
					position, tokenIndex = position37, tokenIndex37
					if !_rules[ruleUninomialWord]() {
						goto l35
					}
				}
			l37:
				{
					position39, tokenIndex39 := position, tokenIndex
					if !_rules[rule_]() {
						goto l39
					}
					if !_rules[ruleAuthorWord]() {
						goto l39
					}
					goto l35
				l39:
					position, tokenIndex = position39, tokenIndex39
				}
				add(ruleGenusWord, position36)
			}
			return true
		l35:
			position, tokenIndex = position35, tokenIndex35
			return false
		},
		/* 13 InfraspGroup <- <(InfraspEpithet (_ InfraspEpithet)? (_ InfraspEpithet)?)> */
		func() bool {
			position40, tokenIndex40 := position, tokenIndex
			{
				position41 := position
				if !_rules[ruleInfraspEpithet]() {
					goto l40
				}
				{
					position42, tokenIndex42 := position, tokenIndex
					if !_rules[rule_]() {
						goto l42
					}
					if !_rules[ruleInfraspEpithet]() {
						goto l42
					}
					goto l43
				l42:
					position, tokenIndex = position42, tokenIndex42
				}
			l43:
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
				add(ruleInfraspGroup, position41)
			}
			return true
		l40:
			position, tokenIndex = position40, tokenIndex40
			return false
		},
		/* 14 InfraspEpithet <- <((Rank _?)? !AuthorEx Word (_ Authorship)?)> */
		func() bool {
			position46, tokenIndex46 := position, tokenIndex
			{
				position47 := position
				{
					position48, tokenIndex48 := position, tokenIndex
					if !_rules[ruleRank]() {
						goto l48
					}
					{
						position50, tokenIndex50 := position, tokenIndex
						if !_rules[rule_]() {
							goto l50
						}
						goto l51
					l50:
						position, tokenIndex = position50, tokenIndex50
					}
				l51:
					goto l49
				l48:
					position, tokenIndex = position48, tokenIndex48
				}
			l49:
				{
					position52, tokenIndex52 := position, tokenIndex
					if !_rules[ruleAuthorEx]() {
						goto l52
					}
					goto l46
				l52:
					position, tokenIndex = position52, tokenIndex52
				}
				if !_rules[ruleWord]() {
					goto l46
				}
				{
					position53, tokenIndex53 := position, tokenIndex
					if !_rules[rule_]() {
						goto l53
					}
					if !_rules[ruleAuthorship]() {
						goto l53
					}
					goto l54
				l53:
					position, tokenIndex = position53, tokenIndex53
				}
			l54:
				add(ruleInfraspEpithet, position47)
			}
			return true
		l46:
			position, tokenIndex = position46, tokenIndex46
			return false
		},
		/* 15 SpeciesEpithet <- <(!AuthorEx Word (_? Authorship)? ','? &(SpaceCharEOI / '('))> */
		func() bool {
			position55, tokenIndex55 := position, tokenIndex
			{
				position56 := position
				{
					position57, tokenIndex57 := position, tokenIndex
					if !_rules[ruleAuthorEx]() {
						goto l57
					}
					goto l55
				l57:
					position, tokenIndex = position57, tokenIndex57
				}
				if !_rules[ruleWord]() {
					goto l55
				}
				{
					position58, tokenIndex58 := position, tokenIndex
					{
						position60, tokenIndex60 := position, tokenIndex
						if !_rules[rule_]() {
							goto l60
						}
						goto l61
					l60:
						position, tokenIndex = position60, tokenIndex60
					}
				l61:
					if !_rules[ruleAuthorship]() {
						goto l58
					}
					goto l59
				l58:
					position, tokenIndex = position58, tokenIndex58
				}
			l59:
				{
					position62, tokenIndex62 := position, tokenIndex
					if buffer[position] != rune(',') {
						goto l62
					}
					position++
					goto l63
				l62:
					position, tokenIndex = position62, tokenIndex62
				}
			l63:
				{
					position64, tokenIndex64 := position, tokenIndex
					{
						position65, tokenIndex65 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l66
						}
						goto l65
					l66:
						position, tokenIndex = position65, tokenIndex65
						if buffer[position] != rune('(') {
							goto l55
						}
						position++
					}
				l65:
					position, tokenIndex = position64, tokenIndex64
				}
				add(ruleSpeciesEpithet, position56)
			}
			return true
		l55:
			position, tokenIndex = position55, tokenIndex55
			return false
		},
		/* 16 Comparison <- <('c' 'f' '.'?)> */
		nil,
		/* 17 Rank <- <(RankForma / RankVar / RankSsp / RankOther / RankOtherUncommon)> */
		func() bool {
			position68, tokenIndex68 := position, tokenIndex
			{
				position69 := position
				{
					position70, tokenIndex70 := position, tokenIndex
					if !_rules[ruleRankForma]() {
						goto l71
					}
					goto l70
				l71:
					position, tokenIndex = position70, tokenIndex70
					if !_rules[ruleRankVar]() {
						goto l72
					}
					goto l70
				l72:
					position, tokenIndex = position70, tokenIndex70
					if !_rules[ruleRankSsp]() {
						goto l73
					}
					goto l70
				l73:
					position, tokenIndex = position70, tokenIndex70
					if !_rules[ruleRankOther]() {
						goto l74
					}
					goto l70
				l74:
					position, tokenIndex = position70, tokenIndex70
					if !_rules[ruleRankOtherUncommon]() {
						goto l68
					}
				}
			l70:
				add(ruleRank, position69)
			}
			return true
		l68:
			position, tokenIndex = position68, tokenIndex68
			return false
		},
		/* 18 RankOtherUncommon <- <(('*' / ('n' 'a' 't') / ('f' '.' 's' 'p') / ('m' 'u' 't' '.')) &SpaceCharEOI)> */
		func() bool {
			position75, tokenIndex75 := position, tokenIndex
			{
				position76 := position
				{
					position77, tokenIndex77 := position, tokenIndex
					if buffer[position] != rune('*') {
						goto l78
					}
					position++
					goto l77
				l78:
					position, tokenIndex = position77, tokenIndex77
					if buffer[position] != rune('n') {
						goto l79
					}
					position++
					if buffer[position] != rune('a') {
						goto l79
					}
					position++
					if buffer[position] != rune('t') {
						goto l79
					}
					position++
					goto l77
				l79:
					position, tokenIndex = position77, tokenIndex77
					if buffer[position] != rune('f') {
						goto l80
					}
					position++
					if buffer[position] != rune('.') {
						goto l80
					}
					position++
					if buffer[position] != rune('s') {
						goto l80
					}
					position++
					if buffer[position] != rune('p') {
						goto l80
					}
					position++
					goto l77
				l80:
					position, tokenIndex = position77, tokenIndex77
					if buffer[position] != rune('m') {
						goto l75
					}
					position++
					if buffer[position] != rune('u') {
						goto l75
					}
					position++
					if buffer[position] != rune('t') {
						goto l75
					}
					position++
					if buffer[position] != rune('.') {
						goto l75
					}
					position++
				}
			l77:
				{
					position81, tokenIndex81 := position, tokenIndex
					if !_rules[ruleSpaceCharEOI]() {
						goto l75
					}
					position, tokenIndex = position81, tokenIndex81
				}
				add(ruleRankOtherUncommon, position76)
			}
			return true
		l75:
			position, tokenIndex = position75, tokenIndex75
			return false
		},
		/* 19 RankOther <- <((('m' 'o' 'r' 'p' 'h' '.') / ('n' 'o' 't' 'h' 'o' 's' 'u' 'b' 's' 'p' '.') / ('c' 'o' 'n' 'v' 'a' 'r' '.') / ('p' 's' 'e' 'u' 'd' 'o' 'v' 'a' 'r' '.') / ('s' 'e' 'c' 't' '.') / ('s' 'e' 'r' '.') / ('s' 'u' 'b' 'v' 'a' 'r' '.') / ('s' 'u' 'b' 'f' '.') / ('r' 'a' 'c' 'e') / 'α' / ('β' 'β') / 'β' / 'γ' / 'δ' / 'ε' / 'φ' / 'θ' / 'μ' / ('a' '.') / ('b' '.') / ('c' '.') / ('d' '.') / ('e' '.') / ('g' '.') / ('k' '.') / ('p' 'v' '.') / ('p' 'a' 't' 'h' 'o' 'v' 'a' 'r' '.') / ('a' 'b' '.' (_? ('n' '.'))?) / ('s' 't' '.')) &SpaceCharEOI)> */
		func() bool {
			position82, tokenIndex82 := position, tokenIndex
			{
				position83 := position
				{
					position84, tokenIndex84 := position, tokenIndex
					if buffer[position] != rune('m') {
						goto l85
					}
					position++
					if buffer[position] != rune('o') {
						goto l85
					}
					position++
					if buffer[position] != rune('r') {
						goto l85
					}
					position++
					if buffer[position] != rune('p') {
						goto l85
					}
					position++
					if buffer[position] != rune('h') {
						goto l85
					}
					position++
					if buffer[position] != rune('.') {
						goto l85
					}
					position++
					goto l84
				l85:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('n') {
						goto l86
					}
					position++
					if buffer[position] != rune('o') {
						goto l86
					}
					position++
					if buffer[position] != rune('t') {
						goto l86
					}
					position++
					if buffer[position] != rune('h') {
						goto l86
					}
					position++
					if buffer[position] != rune('o') {
						goto l86
					}
					position++
					if buffer[position] != rune('s') {
						goto l86
					}
					position++
					if buffer[position] != rune('u') {
						goto l86
					}
					position++
					if buffer[position] != rune('b') {
						goto l86
					}
					position++
					if buffer[position] != rune('s') {
						goto l86
					}
					position++
					if buffer[position] != rune('p') {
						goto l86
					}
					position++
					if buffer[position] != rune('.') {
						goto l86
					}
					position++
					goto l84
				l86:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('c') {
						goto l87
					}
					position++
					if buffer[position] != rune('o') {
						goto l87
					}
					position++
					if buffer[position] != rune('n') {
						goto l87
					}
					position++
					if buffer[position] != rune('v') {
						goto l87
					}
					position++
					if buffer[position] != rune('a') {
						goto l87
					}
					position++
					if buffer[position] != rune('r') {
						goto l87
					}
					position++
					if buffer[position] != rune('.') {
						goto l87
					}
					position++
					goto l84
				l87:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('p') {
						goto l88
					}
					position++
					if buffer[position] != rune('s') {
						goto l88
					}
					position++
					if buffer[position] != rune('e') {
						goto l88
					}
					position++
					if buffer[position] != rune('u') {
						goto l88
					}
					position++
					if buffer[position] != rune('d') {
						goto l88
					}
					position++
					if buffer[position] != rune('o') {
						goto l88
					}
					position++
					if buffer[position] != rune('v') {
						goto l88
					}
					position++
					if buffer[position] != rune('a') {
						goto l88
					}
					position++
					if buffer[position] != rune('r') {
						goto l88
					}
					position++
					if buffer[position] != rune('.') {
						goto l88
					}
					position++
					goto l84
				l88:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('s') {
						goto l89
					}
					position++
					if buffer[position] != rune('e') {
						goto l89
					}
					position++
					if buffer[position] != rune('c') {
						goto l89
					}
					position++
					if buffer[position] != rune('t') {
						goto l89
					}
					position++
					if buffer[position] != rune('.') {
						goto l89
					}
					position++
					goto l84
				l89:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('s') {
						goto l90
					}
					position++
					if buffer[position] != rune('e') {
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
					goto l84
				l90:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('s') {
						goto l91
					}
					position++
					if buffer[position] != rune('u') {
						goto l91
					}
					position++
					if buffer[position] != rune('b') {
						goto l91
					}
					position++
					if buffer[position] != rune('v') {
						goto l91
					}
					position++
					if buffer[position] != rune('a') {
						goto l91
					}
					position++
					if buffer[position] != rune('r') {
						goto l91
					}
					position++
					if buffer[position] != rune('.') {
						goto l91
					}
					position++
					goto l84
				l91:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('s') {
						goto l92
					}
					position++
					if buffer[position] != rune('u') {
						goto l92
					}
					position++
					if buffer[position] != rune('b') {
						goto l92
					}
					position++
					if buffer[position] != rune('f') {
						goto l92
					}
					position++
					if buffer[position] != rune('.') {
						goto l92
					}
					position++
					goto l84
				l92:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('r') {
						goto l93
					}
					position++
					if buffer[position] != rune('a') {
						goto l93
					}
					position++
					if buffer[position] != rune('c') {
						goto l93
					}
					position++
					if buffer[position] != rune('e') {
						goto l93
					}
					position++
					goto l84
				l93:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('α') {
						goto l94
					}
					position++
					goto l84
				l94:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('β') {
						goto l95
					}
					position++
					if buffer[position] != rune('β') {
						goto l95
					}
					position++
					goto l84
				l95:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('β') {
						goto l96
					}
					position++
					goto l84
				l96:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('γ') {
						goto l97
					}
					position++
					goto l84
				l97:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('δ') {
						goto l98
					}
					position++
					goto l84
				l98:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('ε') {
						goto l99
					}
					position++
					goto l84
				l99:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('φ') {
						goto l100
					}
					position++
					goto l84
				l100:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('θ') {
						goto l101
					}
					position++
					goto l84
				l101:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('μ') {
						goto l102
					}
					position++
					goto l84
				l102:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('a') {
						goto l103
					}
					position++
					if buffer[position] != rune('.') {
						goto l103
					}
					position++
					goto l84
				l103:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('b') {
						goto l104
					}
					position++
					if buffer[position] != rune('.') {
						goto l104
					}
					position++
					goto l84
				l104:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('c') {
						goto l105
					}
					position++
					if buffer[position] != rune('.') {
						goto l105
					}
					position++
					goto l84
				l105:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('d') {
						goto l106
					}
					position++
					if buffer[position] != rune('.') {
						goto l106
					}
					position++
					goto l84
				l106:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('e') {
						goto l107
					}
					position++
					if buffer[position] != rune('.') {
						goto l107
					}
					position++
					goto l84
				l107:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('g') {
						goto l108
					}
					position++
					if buffer[position] != rune('.') {
						goto l108
					}
					position++
					goto l84
				l108:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('k') {
						goto l109
					}
					position++
					if buffer[position] != rune('.') {
						goto l109
					}
					position++
					goto l84
				l109:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('p') {
						goto l110
					}
					position++
					if buffer[position] != rune('v') {
						goto l110
					}
					position++
					if buffer[position] != rune('.') {
						goto l110
					}
					position++
					goto l84
				l110:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('p') {
						goto l111
					}
					position++
					if buffer[position] != rune('a') {
						goto l111
					}
					position++
					if buffer[position] != rune('t') {
						goto l111
					}
					position++
					if buffer[position] != rune('h') {
						goto l111
					}
					position++
					if buffer[position] != rune('o') {
						goto l111
					}
					position++
					if buffer[position] != rune('v') {
						goto l111
					}
					position++
					if buffer[position] != rune('a') {
						goto l111
					}
					position++
					if buffer[position] != rune('r') {
						goto l111
					}
					position++
					if buffer[position] != rune('.') {
						goto l111
					}
					position++
					goto l84
				l111:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('a') {
						goto l112
					}
					position++
					if buffer[position] != rune('b') {
						goto l112
					}
					position++
					if buffer[position] != rune('.') {
						goto l112
					}
					position++
					{
						position113, tokenIndex113 := position, tokenIndex
						{
							position115, tokenIndex115 := position, tokenIndex
							if !_rules[rule_]() {
								goto l115
							}
							goto l116
						l115:
							position, tokenIndex = position115, tokenIndex115
						}
					l116:
						if buffer[position] != rune('n') {
							goto l113
						}
						position++
						if buffer[position] != rune('.') {
							goto l113
						}
						position++
						goto l114
					l113:
						position, tokenIndex = position113, tokenIndex113
					}
				l114:
					goto l84
				l112:
					position, tokenIndex = position84, tokenIndex84
					if buffer[position] != rune('s') {
						goto l82
					}
					position++
					if buffer[position] != rune('t') {
						goto l82
					}
					position++
					if buffer[position] != rune('.') {
						goto l82
					}
					position++
				}
			l84:
				{
					position117, tokenIndex117 := position, tokenIndex
					if !_rules[ruleSpaceCharEOI]() {
						goto l82
					}
					position, tokenIndex = position117, tokenIndex117
				}
				add(ruleRankOther, position83)
			}
			return true
		l82:
			position, tokenIndex = position82, tokenIndex82
			return false
		},
		/* 20 RankVar <- <(('v' 'a' 'r' 'i' 'e' 't' 'y') / ('[' 'v' 'a' 'r' '.' ']') / ('n' 'v' 'a' 'r' '.') / ('v' 'a' 'r' (&SpaceCharEOI / '.')))> */
		func() bool {
			position118, tokenIndex118 := position, tokenIndex
			{
				position119 := position
				{
					position120, tokenIndex120 := position, tokenIndex
					if buffer[position] != rune('v') {
						goto l121
					}
					position++
					if buffer[position] != rune('a') {
						goto l121
					}
					position++
					if buffer[position] != rune('r') {
						goto l121
					}
					position++
					if buffer[position] != rune('i') {
						goto l121
					}
					position++
					if buffer[position] != rune('e') {
						goto l121
					}
					position++
					if buffer[position] != rune('t') {
						goto l121
					}
					position++
					if buffer[position] != rune('y') {
						goto l121
					}
					position++
					goto l120
				l121:
					position, tokenIndex = position120, tokenIndex120
					if buffer[position] != rune('[') {
						goto l122
					}
					position++
					if buffer[position] != rune('v') {
						goto l122
					}
					position++
					if buffer[position] != rune('a') {
						goto l122
					}
					position++
					if buffer[position] != rune('r') {
						goto l122
					}
					position++
					if buffer[position] != rune('.') {
						goto l122
					}
					position++
					if buffer[position] != rune(']') {
						goto l122
					}
					position++
					goto l120
				l122:
					position, tokenIndex = position120, tokenIndex120
					if buffer[position] != rune('n') {
						goto l123
					}
					position++
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
					if buffer[position] != rune('.') {
						goto l123
					}
					position++
					goto l120
				l123:
					position, tokenIndex = position120, tokenIndex120
					if buffer[position] != rune('v') {
						goto l118
					}
					position++
					if buffer[position] != rune('a') {
						goto l118
					}
					position++
					if buffer[position] != rune('r') {
						goto l118
					}
					position++
					{
						position124, tokenIndex124 := position, tokenIndex
						{
							position126, tokenIndex126 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l125
							}
							position, tokenIndex = position126, tokenIndex126
						}
						goto l124
					l125:
						position, tokenIndex = position124, tokenIndex124
						if buffer[position] != rune('.') {
							goto l118
						}
						position++
					}
				l124:
				}
			l120:
				add(ruleRankVar, position119)
			}
			return true
		l118:
			position, tokenIndex = position118, tokenIndex118
			return false
		},
		/* 21 RankForma <- <((('f' 'o' 'r' 'm' 'a') / ('f' 'm' 'a') / ('f' 'o' 'r' 'm') / ('f' 'o') / 'f') (&SpaceCharEOI / '.'))> */
		func() bool {
			position127, tokenIndex127 := position, tokenIndex
			{
				position128 := position
				{
					position129, tokenIndex129 := position, tokenIndex
					if buffer[position] != rune('f') {
						goto l130
					}
					position++
					if buffer[position] != rune('o') {
						goto l130
					}
					position++
					if buffer[position] != rune('r') {
						goto l130
					}
					position++
					if buffer[position] != rune('m') {
						goto l130
					}
					position++
					if buffer[position] != rune('a') {
						goto l130
					}
					position++
					goto l129
				l130:
					position, tokenIndex = position129, tokenIndex129
					if buffer[position] != rune('f') {
						goto l131
					}
					position++
					if buffer[position] != rune('m') {
						goto l131
					}
					position++
					if buffer[position] != rune('a') {
						goto l131
					}
					position++
					goto l129
				l131:
					position, tokenIndex = position129, tokenIndex129
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
					goto l129
				l132:
					position, tokenIndex = position129, tokenIndex129
					if buffer[position] != rune('f') {
						goto l133
					}
					position++
					if buffer[position] != rune('o') {
						goto l133
					}
					position++
					goto l129
				l133:
					position, tokenIndex = position129, tokenIndex129
					if buffer[position] != rune('f') {
						goto l127
					}
					position++
				}
			l129:
				{
					position134, tokenIndex134 := position, tokenIndex
					{
						position136, tokenIndex136 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l135
						}
						position, tokenIndex = position136, tokenIndex136
					}
					goto l134
				l135:
					position, tokenIndex = position134, tokenIndex134
					if buffer[position] != rune('.') {
						goto l127
					}
					position++
				}
			l134:
				add(ruleRankForma, position128)
			}
			return true
		l127:
			position, tokenIndex = position127, tokenIndex127
			return false
		},
		/* 22 RankSsp <- <((('s' 's' 'p') / ('s' 'u' 'b' 's' 'p')) (&SpaceCharEOI / '.'))> */
		func() bool {
			position137, tokenIndex137 := position, tokenIndex
			{
				position138 := position
				{
					position139, tokenIndex139 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l140
					}
					position++
					if buffer[position] != rune('s') {
						goto l140
					}
					position++
					if buffer[position] != rune('p') {
						goto l140
					}
					position++
					goto l139
				l140:
					position, tokenIndex = position139, tokenIndex139
					if buffer[position] != rune('s') {
						goto l137
					}
					position++
					if buffer[position] != rune('u') {
						goto l137
					}
					position++
					if buffer[position] != rune('b') {
						goto l137
					}
					position++
					if buffer[position] != rune('s') {
						goto l137
					}
					position++
					if buffer[position] != rune('p') {
						goto l137
					}
					position++
				}
			l139:
				{
					position141, tokenIndex141 := position, tokenIndex
					{
						position143, tokenIndex143 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l142
						}
						position, tokenIndex = position143, tokenIndex143
					}
					goto l141
				l142:
					position, tokenIndex = position141, tokenIndex141
					if buffer[position] != rune('.') {
						goto l137
					}
					position++
				}
			l141:
				add(ruleRankSsp, position138)
			}
			return true
		l137:
			position, tokenIndex = position137, tokenIndex137
			return false
		},
		/* 23 SubGenusOrSuperspecies <- <('(' _? Word _? ')')> */
		func() bool {
			position144, tokenIndex144 := position, tokenIndex
			{
				position145 := position
				if buffer[position] != rune('(') {
					goto l144
				}
				position++
				{
					position146, tokenIndex146 := position, tokenIndex
					if !_rules[rule_]() {
						goto l146
					}
					goto l147
				l146:
					position, tokenIndex = position146, tokenIndex146
				}
			l147:
				if !_rules[ruleWord]() {
					goto l144
				}
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
				if buffer[position] != rune(')') {
					goto l144
				}
				position++
				add(ruleSubGenusOrSuperspecies, position145)
			}
			return true
		l144:
			position, tokenIndex = position144, tokenIndex144
			return false
		},
		/* 24 SubGenus <- <('(' _? UninomialWord _? ')')> */
		func() bool {
			position150, tokenIndex150 := position, tokenIndex
			{
				position151 := position
				if buffer[position] != rune('(') {
					goto l150
				}
				position++
				{
					position152, tokenIndex152 := position, tokenIndex
					if !_rules[rule_]() {
						goto l152
					}
					goto l153
				l152:
					position, tokenIndex = position152, tokenIndex152
				}
			l153:
				if !_rules[ruleUninomialWord]() {
					goto l150
				}
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
				if buffer[position] != rune(')') {
					goto l150
				}
				position++
				add(ruleSubGenus, position151)
			}
			return true
		l150:
			position, tokenIndex = position150, tokenIndex150
			return false
		},
		/* 25 UninomialCombo <- <(UninomialCombo1 / UninomialCombo2)> */
		func() bool {
			position156, tokenIndex156 := position, tokenIndex
			{
				position157 := position
				{
					position158, tokenIndex158 := position, tokenIndex
					if !_rules[ruleUninomialCombo1]() {
						goto l159
					}
					goto l158
				l159:
					position, tokenIndex = position158, tokenIndex158
					if !_rules[ruleUninomialCombo2]() {
						goto l156
					}
				}
			l158:
				add(ruleUninomialCombo, position157)
			}
			return true
		l156:
			position, tokenIndex = position156, tokenIndex156
			return false
		},
		/* 26 UninomialCombo1 <- <(UninomialWord _? SubGenus _? Authorship .?)> */
		func() bool {
			position160, tokenIndex160 := position, tokenIndex
			{
				position161 := position
				if !_rules[ruleUninomialWord]() {
					goto l160
				}
				{
					position162, tokenIndex162 := position, tokenIndex
					if !_rules[rule_]() {
						goto l162
					}
					goto l163
				l162:
					position, tokenIndex = position162, tokenIndex162
				}
			l163:
				if !_rules[ruleSubGenus]() {
					goto l160
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
				if !_rules[ruleAuthorship]() {
					goto l160
				}
				{
					position166, tokenIndex166 := position, tokenIndex
					if !matchDot() {
						goto l166
					}
					goto l167
				l166:
					position, tokenIndex = position166, tokenIndex166
				}
			l167:
				add(ruleUninomialCombo1, position161)
			}
			return true
		l160:
			position, tokenIndex = position160, tokenIndex160
			return false
		},
		/* 27 UninomialCombo2 <- <(Uninomial _? RankUninomial _? Uninomial)> */
		func() bool {
			position168, tokenIndex168 := position, tokenIndex
			{
				position169 := position
				if !_rules[ruleUninomial]() {
					goto l168
				}
				{
					position170, tokenIndex170 := position, tokenIndex
					if !_rules[rule_]() {
						goto l170
					}
					goto l171
				l170:
					position, tokenIndex = position170, tokenIndex170
				}
			l171:
				if !_rules[ruleRankUninomial]() {
					goto l168
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
				if !_rules[ruleUninomial]() {
					goto l168
				}
				add(ruleUninomialCombo2, position169)
			}
			return true
		l168:
			position, tokenIndex = position168, tokenIndex168
			return false
		},
		/* 28 RankUninomial <- <((('s' 'e' 'c' 't') / ('s' 'u' 'b' 's' 'e' 'c' 't') / ('t' 'r' 'i' 'b') / ('s' 'u' 'b' 't' 'r' 'i' 'b') / ('s' 'u' 'b' 's' 'e' 'r') / ('s' 'e' 'r') / ('s' 'u' 'b' 'g' 'e' 'n') / ('f' 'a' 'm') / ('s' 'u' 'b' 'f' 'a' 'm') / ('s' 'u' 'p' 'e' 'r' 't' 'r' 'i' 'b')) '.'?)> */
		func() bool {
			position174, tokenIndex174 := position, tokenIndex
			{
				position175 := position
				{
					position176, tokenIndex176 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l177
					}
					position++
					if buffer[position] != rune('e') {
						goto l177
					}
					position++
					if buffer[position] != rune('c') {
						goto l177
					}
					position++
					if buffer[position] != rune('t') {
						goto l177
					}
					position++
					goto l176
				l177:
					position, tokenIndex = position176, tokenIndex176
					if buffer[position] != rune('s') {
						goto l178
					}
					position++
					if buffer[position] != rune('u') {
						goto l178
					}
					position++
					if buffer[position] != rune('b') {
						goto l178
					}
					position++
					if buffer[position] != rune('s') {
						goto l178
					}
					position++
					if buffer[position] != rune('e') {
						goto l178
					}
					position++
					if buffer[position] != rune('c') {
						goto l178
					}
					position++
					if buffer[position] != rune('t') {
						goto l178
					}
					position++
					goto l176
				l178:
					position, tokenIndex = position176, tokenIndex176
					if buffer[position] != rune('t') {
						goto l179
					}
					position++
					if buffer[position] != rune('r') {
						goto l179
					}
					position++
					if buffer[position] != rune('i') {
						goto l179
					}
					position++
					if buffer[position] != rune('b') {
						goto l179
					}
					position++
					goto l176
				l179:
					position, tokenIndex = position176, tokenIndex176
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
					if buffer[position] != rune('t') {
						goto l180
					}
					position++
					if buffer[position] != rune('r') {
						goto l180
					}
					position++
					if buffer[position] != rune('i') {
						goto l180
					}
					position++
					if buffer[position] != rune('b') {
						goto l180
					}
					position++
					goto l176
				l180:
					position, tokenIndex = position176, tokenIndex176
					if buffer[position] != rune('s') {
						goto l181
					}
					position++
					if buffer[position] != rune('u') {
						goto l181
					}
					position++
					if buffer[position] != rune('b') {
						goto l181
					}
					position++
					if buffer[position] != rune('s') {
						goto l181
					}
					position++
					if buffer[position] != rune('e') {
						goto l181
					}
					position++
					if buffer[position] != rune('r') {
						goto l181
					}
					position++
					goto l176
				l181:
					position, tokenIndex = position176, tokenIndex176
					if buffer[position] != rune('s') {
						goto l182
					}
					position++
					if buffer[position] != rune('e') {
						goto l182
					}
					position++
					if buffer[position] != rune('r') {
						goto l182
					}
					position++
					goto l176
				l182:
					position, tokenIndex = position176, tokenIndex176
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
					if buffer[position] != rune('g') {
						goto l183
					}
					position++
					if buffer[position] != rune('e') {
						goto l183
					}
					position++
					if buffer[position] != rune('n') {
						goto l183
					}
					position++
					goto l176
				l183:
					position, tokenIndex = position176, tokenIndex176
					if buffer[position] != rune('f') {
						goto l184
					}
					position++
					if buffer[position] != rune('a') {
						goto l184
					}
					position++
					if buffer[position] != rune('m') {
						goto l184
					}
					position++
					goto l176
				l184:
					position, tokenIndex = position176, tokenIndex176
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
					if buffer[position] != rune('f') {
						goto l185
					}
					position++
					if buffer[position] != rune('a') {
						goto l185
					}
					position++
					if buffer[position] != rune('m') {
						goto l185
					}
					position++
					goto l176
				l185:
					position, tokenIndex = position176, tokenIndex176
					if buffer[position] != rune('s') {
						goto l174
					}
					position++
					if buffer[position] != rune('u') {
						goto l174
					}
					position++
					if buffer[position] != rune('p') {
						goto l174
					}
					position++
					if buffer[position] != rune('e') {
						goto l174
					}
					position++
					if buffer[position] != rune('r') {
						goto l174
					}
					position++
					if buffer[position] != rune('t') {
						goto l174
					}
					position++
					if buffer[position] != rune('r') {
						goto l174
					}
					position++
					if buffer[position] != rune('i') {
						goto l174
					}
					position++
					if buffer[position] != rune('b') {
						goto l174
					}
					position++
				}
			l176:
				{
					position186, tokenIndex186 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l186
					}
					position++
					goto l187
				l186:
					position, tokenIndex = position186, tokenIndex186
				}
			l187:
				add(ruleRankUninomial, position175)
			}
			return true
		l174:
			position, tokenIndex = position174, tokenIndex174
			return false
		},
		/* 29 Uninomial <- <(UninomialWord (_ Authorship)?)> */
		func() bool {
			position188, tokenIndex188 := position, tokenIndex
			{
				position189 := position
				if !_rules[ruleUninomialWord]() {
					goto l188
				}
				{
					position190, tokenIndex190 := position, tokenIndex
					if !_rules[rule_]() {
						goto l190
					}
					if !_rules[ruleAuthorship]() {
						goto l190
					}
					goto l191
				l190:
					position, tokenIndex = position190, tokenIndex190
				}
			l191:
				add(ruleUninomial, position189)
			}
			return true
		l188:
			position, tokenIndex = position188, tokenIndex188
			return false
		},
		/* 30 UninomialWord <- <(CapWord / TwoLetterGenus)> */
		func() bool {
			position192, tokenIndex192 := position, tokenIndex
			{
				position193 := position
				{
					position194, tokenIndex194 := position, tokenIndex
					if !_rules[ruleCapWord]() {
						goto l195
					}
					goto l194
				l195:
					position, tokenIndex = position194, tokenIndex194
					if !_rules[ruleTwoLetterGenus]() {
						goto l192
					}
				}
			l194:
				add(ruleUninomialWord, position193)
			}
			return true
		l192:
			position, tokenIndex = position192, tokenIndex192
			return false
		},
		/* 31 AbbrGenus <- <(UpperChar LowerChar* '.')> */
		func() bool {
			position196, tokenIndex196 := position, tokenIndex
			{
				position197 := position
				if !_rules[ruleUpperChar]() {
					goto l196
				}
			l198:
				{
					position199, tokenIndex199 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l199
					}
					goto l198
				l199:
					position, tokenIndex = position199, tokenIndex199
				}
				if buffer[position] != rune('.') {
					goto l196
				}
				position++
				add(ruleAbbrGenus, position197)
			}
			return true
		l196:
			position, tokenIndex = position196, tokenIndex196
			return false
		},
		/* 32 CapWord <- <(CapWord2 / CapWord1)> */
		func() bool {
			position200, tokenIndex200 := position, tokenIndex
			{
				position201 := position
				{
					position202, tokenIndex202 := position, tokenIndex
					if !_rules[ruleCapWord2]() {
						goto l203
					}
					goto l202
				l203:
					position, tokenIndex = position202, tokenIndex202
					if !_rules[ruleCapWord1]() {
						goto l200
					}
				}
			l202:
				add(ruleCapWord, position201)
			}
			return true
		l200:
			position, tokenIndex = position200, tokenIndex200
			return false
		},
		/* 33 CapWord1 <- <(NameUpperChar NameLowerChar NameLowerChar+ '?'?)> */
		func() bool {
			position204, tokenIndex204 := position, tokenIndex
			{
				position205 := position
				if !_rules[ruleNameUpperChar]() {
					goto l204
				}
				if !_rules[ruleNameLowerChar]() {
					goto l204
				}
				if !_rules[ruleNameLowerChar]() {
					goto l204
				}
			l206:
				{
					position207, tokenIndex207 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l207
					}
					goto l206
				l207:
					position, tokenIndex = position207, tokenIndex207
				}
				{
					position208, tokenIndex208 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l208
					}
					position++
					goto l209
				l208:
					position, tokenIndex = position208, tokenIndex208
				}
			l209:
				add(ruleCapWord1, position205)
			}
			return true
		l204:
			position, tokenIndex = position204, tokenIndex204
			return false
		},
		/* 34 CapWord2 <- <(CapWord1 dash (CapWord1 / Word1))> */
		func() bool {
			position210, tokenIndex210 := position, tokenIndex
			{
				position211 := position
				if !_rules[ruleCapWord1]() {
					goto l210
				}
				if !_rules[ruledash]() {
					goto l210
				}
				{
					position212, tokenIndex212 := position, tokenIndex
					if !_rules[ruleCapWord1]() {
						goto l213
					}
					goto l212
				l213:
					position, tokenIndex = position212, tokenIndex212
					if !_rules[ruleWord1]() {
						goto l210
					}
				}
			l212:
				add(ruleCapWord2, position211)
			}
			return true
		l210:
			position, tokenIndex = position210, tokenIndex210
			return false
		},
		/* 35 TwoLetterGenus <- <(('C' 'a') / ('E' 'a') / ('G' 'e') / ('I' 'a') / ('I' 'o') / ('I' 'x') / ('L' 'o') / ('O' 'a') / ('R' 'a') / ('T' 'y') / ('U' 'a') / ('A' 'a') / ('J' 'a') / ('Z' 'u') / ('L' 'a') / ('Q' 'u') / ('A' 's') / ('B' 'a'))> */
		func() bool {
			position214, tokenIndex214 := position, tokenIndex
			{
				position215 := position
				{
					position216, tokenIndex216 := position, tokenIndex
					if buffer[position] != rune('C') {
						goto l217
					}
					position++
					if buffer[position] != rune('a') {
						goto l217
					}
					position++
					goto l216
				l217:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('E') {
						goto l218
					}
					position++
					if buffer[position] != rune('a') {
						goto l218
					}
					position++
					goto l216
				l218:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('G') {
						goto l219
					}
					position++
					if buffer[position] != rune('e') {
						goto l219
					}
					position++
					goto l216
				l219:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('I') {
						goto l220
					}
					position++
					if buffer[position] != rune('a') {
						goto l220
					}
					position++
					goto l216
				l220:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('I') {
						goto l221
					}
					position++
					if buffer[position] != rune('o') {
						goto l221
					}
					position++
					goto l216
				l221:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('I') {
						goto l222
					}
					position++
					if buffer[position] != rune('x') {
						goto l222
					}
					position++
					goto l216
				l222:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('L') {
						goto l223
					}
					position++
					if buffer[position] != rune('o') {
						goto l223
					}
					position++
					goto l216
				l223:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('O') {
						goto l224
					}
					position++
					if buffer[position] != rune('a') {
						goto l224
					}
					position++
					goto l216
				l224:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('R') {
						goto l225
					}
					position++
					if buffer[position] != rune('a') {
						goto l225
					}
					position++
					goto l216
				l225:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('T') {
						goto l226
					}
					position++
					if buffer[position] != rune('y') {
						goto l226
					}
					position++
					goto l216
				l226:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('U') {
						goto l227
					}
					position++
					if buffer[position] != rune('a') {
						goto l227
					}
					position++
					goto l216
				l227:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('A') {
						goto l228
					}
					position++
					if buffer[position] != rune('a') {
						goto l228
					}
					position++
					goto l216
				l228:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('J') {
						goto l229
					}
					position++
					if buffer[position] != rune('a') {
						goto l229
					}
					position++
					goto l216
				l229:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('Z') {
						goto l230
					}
					position++
					if buffer[position] != rune('u') {
						goto l230
					}
					position++
					goto l216
				l230:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('L') {
						goto l231
					}
					position++
					if buffer[position] != rune('a') {
						goto l231
					}
					position++
					goto l216
				l231:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('Q') {
						goto l232
					}
					position++
					if buffer[position] != rune('u') {
						goto l232
					}
					position++
					goto l216
				l232:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('A') {
						goto l233
					}
					position++
					if buffer[position] != rune('s') {
						goto l233
					}
					position++
					goto l216
				l233:
					position, tokenIndex = position216, tokenIndex216
					if buffer[position] != rune('B') {
						goto l214
					}
					position++
					if buffer[position] != rune('a') {
						goto l214
					}
					position++
				}
			l216:
				add(ruleTwoLetterGenus, position215)
			}
			return true
		l214:
			position, tokenIndex = position214, tokenIndex214
			return false
		},
		/* 36 Word <- <(!(AuthorPrefix / RankUninomial / Approximation / Word4) (Word3 / Word2StartDigit / Word2 / Word1) &(SpaceCharEOI / ('(' ')')))> */
		func() bool {
			position234, tokenIndex234 := position, tokenIndex
			{
				position235 := position
				{
					position236, tokenIndex236 := position, tokenIndex
					{
						position237, tokenIndex237 := position, tokenIndex
						if !_rules[ruleAuthorPrefix]() {
							goto l238
						}
						goto l237
					l238:
						position, tokenIndex = position237, tokenIndex237
						if !_rules[ruleRankUninomial]() {
							goto l239
						}
						goto l237
					l239:
						position, tokenIndex = position237, tokenIndex237
						if !_rules[ruleApproximation]() {
							goto l240
						}
						goto l237
					l240:
						position, tokenIndex = position237, tokenIndex237
						if !_rules[ruleWord4]() {
							goto l236
						}
					}
				l237:
					goto l234
				l236:
					position, tokenIndex = position236, tokenIndex236
				}
				{
					position241, tokenIndex241 := position, tokenIndex
					if !_rules[ruleWord3]() {
						goto l242
					}
					goto l241
				l242:
					position, tokenIndex = position241, tokenIndex241
					if !_rules[ruleWord2StartDigit]() {
						goto l243
					}
					goto l241
				l243:
					position, tokenIndex = position241, tokenIndex241
					if !_rules[ruleWord2]() {
						goto l244
					}
					goto l241
				l244:
					position, tokenIndex = position241, tokenIndex241
					if !_rules[ruleWord1]() {
						goto l234
					}
				}
			l241:
				{
					position245, tokenIndex245 := position, tokenIndex
					{
						position246, tokenIndex246 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l247
						}
						goto l246
					l247:
						position, tokenIndex = position246, tokenIndex246
						if buffer[position] != rune('(') {
							goto l234
						}
						position++
						if buffer[position] != rune(')') {
							goto l234
						}
						position++
					}
				l246:
					position, tokenIndex = position245, tokenIndex245
				}
				add(ruleWord, position235)
			}
			return true
		l234:
			position, tokenIndex = position234, tokenIndex234
			return false
		},
		/* 37 Word1 <- <((lASCII dash)? NameLowerChar NameLowerChar+)> */
		func() bool {
			position248, tokenIndex248 := position, tokenIndex
			{
				position249 := position
				{
					position250, tokenIndex250 := position, tokenIndex
					if !_rules[rulelASCII]() {
						goto l250
					}
					if !_rules[ruledash]() {
						goto l250
					}
					goto l251
				l250:
					position, tokenIndex = position250, tokenIndex250
				}
			l251:
				if !_rules[ruleNameLowerChar]() {
					goto l248
				}
				if !_rules[ruleNameLowerChar]() {
					goto l248
				}
			l252:
				{
					position253, tokenIndex253 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l253
					}
					goto l252
				l253:
					position, tokenIndex = position253, tokenIndex253
				}
				add(ruleWord1, position249)
			}
			return true
		l248:
			position, tokenIndex = position248, tokenIndex248
			return false
		},
		/* 38 Word2StartDigit <- <(('1' / '2' / '3' / '4' / '5' / '6' / '7' / '8' / '9') nums? ('.' / dash)? NameLowerChar NameLowerChar NameLowerChar NameLowerChar+)> */
		func() bool {
			position254, tokenIndex254 := position, tokenIndex
			{
				position255 := position
				{
					position256, tokenIndex256 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l257
					}
					position++
					goto l256
				l257:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('2') {
						goto l258
					}
					position++
					goto l256
				l258:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('3') {
						goto l259
					}
					position++
					goto l256
				l259:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('4') {
						goto l260
					}
					position++
					goto l256
				l260:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('5') {
						goto l261
					}
					position++
					goto l256
				l261:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('6') {
						goto l262
					}
					position++
					goto l256
				l262:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('7') {
						goto l263
					}
					position++
					goto l256
				l263:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('8') {
						goto l264
					}
					position++
					goto l256
				l264:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('9') {
						goto l254
					}
					position++
				}
			l256:
				{
					position265, tokenIndex265 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l265
					}
					goto l266
				l265:
					position, tokenIndex = position265, tokenIndex265
				}
			l266:
				{
					position267, tokenIndex267 := position, tokenIndex
					{
						position269, tokenIndex269 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l270
						}
						position++
						goto l269
					l270:
						position, tokenIndex = position269, tokenIndex269
						if !_rules[ruledash]() {
							goto l267
						}
					}
				l269:
					goto l268
				l267:
					position, tokenIndex = position267, tokenIndex267
				}
			l268:
				if !_rules[ruleNameLowerChar]() {
					goto l254
				}
				if !_rules[ruleNameLowerChar]() {
					goto l254
				}
				if !_rules[ruleNameLowerChar]() {
					goto l254
				}
				if !_rules[ruleNameLowerChar]() {
					goto l254
				}
			l271:
				{
					position272, tokenIndex272 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l272
					}
					goto l271
				l272:
					position, tokenIndex = position272, tokenIndex272
				}
				add(ruleWord2StartDigit, position255)
			}
			return true
		l254:
			position, tokenIndex = position254, tokenIndex254
			return false
		},
		/* 39 Word2 <- <(NameLowerChar+ dash? NameLowerChar)> */
		func() bool {
			position273, tokenIndex273 := position, tokenIndex
			{
				position274 := position
				if !_rules[ruleNameLowerChar]() {
					goto l273
				}
			l275:
				{
					position276, tokenIndex276 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l276
					}
					goto l275
				l276:
					position, tokenIndex = position276, tokenIndex276
				}
				{
					position277, tokenIndex277 := position, tokenIndex
					if !_rules[ruledash]() {
						goto l277
					}
					goto l278
				l277:
					position, tokenIndex = position277, tokenIndex277
				}
			l278:
				if !_rules[ruleNameLowerChar]() {
					goto l273
				}
				add(ruleWord2, position274)
			}
			return true
		l273:
			position, tokenIndex = position273, tokenIndex273
			return false
		},
		/* 40 Word3 <- <(NameLowerChar NameLowerChar* apostr Word1)> */
		func() bool {
			position279, tokenIndex279 := position, tokenIndex
			{
				position280 := position
				if !_rules[ruleNameLowerChar]() {
					goto l279
				}
			l281:
				{
					position282, tokenIndex282 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l282
					}
					goto l281
				l282:
					position, tokenIndex = position282, tokenIndex282
				}
				if !_rules[ruleapostr]() {
					goto l279
				}
				if !_rules[ruleWord1]() {
					goto l279
				}
				add(ruleWord3, position280)
			}
			return true
		l279:
			position, tokenIndex = position279, tokenIndex279
			return false
		},
		/* 41 Word4 <- <(NameLowerChar+ '.' NameLowerChar)> */
		func() bool {
			position283, tokenIndex283 := position, tokenIndex
			{
				position284 := position
				if !_rules[ruleNameLowerChar]() {
					goto l283
				}
			l285:
				{
					position286, tokenIndex286 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l286
					}
					goto l285
				l286:
					position, tokenIndex = position286, tokenIndex286
				}
				if buffer[position] != rune('.') {
					goto l283
				}
				position++
				if !_rules[ruleNameLowerChar]() {
					goto l283
				}
				add(ruleWord4, position284)
			}
			return true
		l283:
			position, tokenIndex = position283, tokenIndex283
			return false
		},
		/* 42 HybridChar <- <'×'> */
		nil,
		/* 43 ApproxName <- <(Uninomial _ (ApproxName1 / ApproxName2))> */
		nil,
		/* 44 ApproxName1 <- <(Approximation ApproxNameIgnored)> */
		nil,
		/* 45 ApproxName2 <- <(Word _ Approximation ApproxNameIgnored)> */
		nil,
		/* 46 ApproxNameIgnored <- <.*> */
		nil,
		/* 47 Approximation <- <(('s' 'p' '.' _? ('n' 'r' '.')) / ('s' 'p' '.' _? ('a' 'f' 'f' '.')) / ('m' 'o' 'n' 's' 't' '.') / '?' / ((('s' 'p' 'p') / ('n' 'r') / ('s' 'p') / ('a' 'f' 'f') / ('s' 'p' 'e' 'c' 'i' 'e' 's')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position292, tokenIndex292 := position, tokenIndex
			{
				position293 := position
				{
					position294, tokenIndex294 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l295
					}
					position++
					if buffer[position] != rune('p') {
						goto l295
					}
					position++
					if buffer[position] != rune('.') {
						goto l295
					}
					position++
					{
						position296, tokenIndex296 := position, tokenIndex
						if !_rules[rule_]() {
							goto l296
						}
						goto l297
					l296:
						position, tokenIndex = position296, tokenIndex296
					}
				l297:
					if buffer[position] != rune('n') {
						goto l295
					}
					position++
					if buffer[position] != rune('r') {
						goto l295
					}
					position++
					if buffer[position] != rune('.') {
						goto l295
					}
					position++
					goto l294
				l295:
					position, tokenIndex = position294, tokenIndex294
					if buffer[position] != rune('s') {
						goto l298
					}
					position++
					if buffer[position] != rune('p') {
						goto l298
					}
					position++
					if buffer[position] != rune('.') {
						goto l298
					}
					position++
					{
						position299, tokenIndex299 := position, tokenIndex
						if !_rules[rule_]() {
							goto l299
						}
						goto l300
					l299:
						position, tokenIndex = position299, tokenIndex299
					}
				l300:
					if buffer[position] != rune('a') {
						goto l298
					}
					position++
					if buffer[position] != rune('f') {
						goto l298
					}
					position++
					if buffer[position] != rune('f') {
						goto l298
					}
					position++
					if buffer[position] != rune('.') {
						goto l298
					}
					position++
					goto l294
				l298:
					position, tokenIndex = position294, tokenIndex294
					if buffer[position] != rune('m') {
						goto l301
					}
					position++
					if buffer[position] != rune('o') {
						goto l301
					}
					position++
					if buffer[position] != rune('n') {
						goto l301
					}
					position++
					if buffer[position] != rune('s') {
						goto l301
					}
					position++
					if buffer[position] != rune('t') {
						goto l301
					}
					position++
					if buffer[position] != rune('.') {
						goto l301
					}
					position++
					goto l294
				l301:
					position, tokenIndex = position294, tokenIndex294
					if buffer[position] != rune('?') {
						goto l302
					}
					position++
					goto l294
				l302:
					position, tokenIndex = position294, tokenIndex294
					{
						position303, tokenIndex303 := position, tokenIndex
						if buffer[position] != rune('s') {
							goto l304
						}
						position++
						if buffer[position] != rune('p') {
							goto l304
						}
						position++
						if buffer[position] != rune('p') {
							goto l304
						}
						position++
						goto l303
					l304:
						position, tokenIndex = position303, tokenIndex303
						if buffer[position] != rune('n') {
							goto l305
						}
						position++
						if buffer[position] != rune('r') {
							goto l305
						}
						position++
						goto l303
					l305:
						position, tokenIndex = position303, tokenIndex303
						if buffer[position] != rune('s') {
							goto l306
						}
						position++
						if buffer[position] != rune('p') {
							goto l306
						}
						position++
						goto l303
					l306:
						position, tokenIndex = position303, tokenIndex303
						if buffer[position] != rune('a') {
							goto l307
						}
						position++
						if buffer[position] != rune('f') {
							goto l307
						}
						position++
						if buffer[position] != rune('f') {
							goto l307
						}
						position++
						goto l303
					l307:
						position, tokenIndex = position303, tokenIndex303
						if buffer[position] != rune('s') {
							goto l292
						}
						position++
						if buffer[position] != rune('p') {
							goto l292
						}
						position++
						if buffer[position] != rune('e') {
							goto l292
						}
						position++
						if buffer[position] != rune('c') {
							goto l292
						}
						position++
						if buffer[position] != rune('i') {
							goto l292
						}
						position++
						if buffer[position] != rune('e') {
							goto l292
						}
						position++
						if buffer[position] != rune('s') {
							goto l292
						}
						position++
					}
				l303:
					{
						position308, tokenIndex308 := position, tokenIndex
						{
							position310, tokenIndex310 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l309
							}
							position, tokenIndex = position310, tokenIndex310
						}
						goto l308
					l309:
						position, tokenIndex = position308, tokenIndex308
						if buffer[position] != rune('.') {
							goto l292
						}
						position++
					}
				l308:
				}
			l294:
				add(ruleApproximation, position293)
			}
			return true
		l292:
			position, tokenIndex = position292, tokenIndex292
			return false
		},
		/* 48 Authorship <- <((Authorship1 / Authorship2) &(SpaceCharEOI / ('\\' / '(' / ',' / ':')))> */
		func() bool {
			position311, tokenIndex311 := position, tokenIndex
			{
				position312 := position
				{
					position313, tokenIndex313 := position, tokenIndex
					if !_rules[ruleAuthorship1]() {
						goto l314
					}
					goto l313
				l314:
					position, tokenIndex = position313, tokenIndex313
					if !_rules[ruleAuthorship2]() {
						goto l311
					}
				}
			l313:
				{
					position315, tokenIndex315 := position, tokenIndex
					{
						position316, tokenIndex316 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l317
						}
						goto l316
					l317:
						position, tokenIndex = position316, tokenIndex316
						{
							position318, tokenIndex318 := position, tokenIndex
							if buffer[position] != rune('\\') {
								goto l319
							}
							position++
							goto l318
						l319:
							position, tokenIndex = position318, tokenIndex318
							if buffer[position] != rune('(') {
								goto l320
							}
							position++
							goto l318
						l320:
							position, tokenIndex = position318, tokenIndex318
							if buffer[position] != rune(',') {
								goto l321
							}
							position++
							goto l318
						l321:
							position, tokenIndex = position318, tokenIndex318
							if buffer[position] != rune(':') {
								goto l311
							}
							position++
						}
					l318:
					}
				l316:
					position, tokenIndex = position315, tokenIndex315
				}
				add(ruleAuthorship, position312)
			}
			return true
		l311:
			position, tokenIndex = position311, tokenIndex311
			return false
		},
		/* 49 Authorship1 <- <(Authorship2 _? AuthorsGroup)> */
		func() bool {
			position322, tokenIndex322 := position, tokenIndex
			{
				position323 := position
				if !_rules[ruleAuthorship2]() {
					goto l322
				}
				{
					position324, tokenIndex324 := position, tokenIndex
					if !_rules[rule_]() {
						goto l324
					}
					goto l325
				l324:
					position, tokenIndex = position324, tokenIndex324
				}
			l325:
				if !_rules[ruleAuthorsGroup]() {
					goto l322
				}
				add(ruleAuthorship1, position323)
			}
			return true
		l322:
			position, tokenIndex = position322, tokenIndex322
			return false
		},
		/* 50 Authorship2 <- <(AuthorsGroup / BasionymAuthorship / BasionymAuthorshipYearMisformed)> */
		func() bool {
			position326, tokenIndex326 := position, tokenIndex
			{
				position327 := position
				{
					position328, tokenIndex328 := position, tokenIndex
					if !_rules[ruleAuthorsGroup]() {
						goto l329
					}
					goto l328
				l329:
					position, tokenIndex = position328, tokenIndex328
					if !_rules[ruleBasionymAuthorship]() {
						goto l330
					}
					goto l328
				l330:
					position, tokenIndex = position328, tokenIndex328
					if !_rules[ruleBasionymAuthorshipYearMisformed]() {
						goto l326
					}
				}
			l328:
				add(ruleAuthorship2, position327)
			}
			return true
		l326:
			position, tokenIndex = position326, tokenIndex326
			return false
		},
		/* 51 BasionymAuthorshipYearMisformed <- <('(' _? AuthorsGroup _? ')' (_? ',')? _? Year)> */
		func() bool {
			position331, tokenIndex331 := position, tokenIndex
			{
				position332 := position
				if buffer[position] != rune('(') {
					goto l331
				}
				position++
				{
					position333, tokenIndex333 := position, tokenIndex
					if !_rules[rule_]() {
						goto l333
					}
					goto l334
				l333:
					position, tokenIndex = position333, tokenIndex333
				}
			l334:
				if !_rules[ruleAuthorsGroup]() {
					goto l331
				}
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
				if buffer[position] != rune(')') {
					goto l331
				}
				position++
				{
					position337, tokenIndex337 := position, tokenIndex
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
					if buffer[position] != rune(',') {
						goto l337
					}
					position++
					goto l338
				l337:
					position, tokenIndex = position337, tokenIndex337
				}
			l338:
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
				if !_rules[ruleYear]() {
					goto l331
				}
				add(ruleBasionymAuthorshipYearMisformed, position332)
			}
			return true
		l331:
			position, tokenIndex = position331, tokenIndex331
			return false
		},
		/* 52 BasionymAuthorship <- <(BasionymAuthorship1 / BasionymAuthorship2)> */
		func() bool {
			position343, tokenIndex343 := position, tokenIndex
			{
				position344 := position
				{
					position345, tokenIndex345 := position, tokenIndex
					if !_rules[ruleBasionymAuthorship1]() {
						goto l346
					}
					goto l345
				l346:
					position, tokenIndex = position345, tokenIndex345
					if !_rules[ruleBasionymAuthorship2]() {
						goto l343
					}
				}
			l345:
				add(ruleBasionymAuthorship, position344)
			}
			return true
		l343:
			position, tokenIndex = position343, tokenIndex343
			return false
		},
		/* 53 BasionymAuthorship1 <- <('(' _? AuthorsGroup _? ')')> */
		func() bool {
			position347, tokenIndex347 := position, tokenIndex
			{
				position348 := position
				if buffer[position] != rune('(') {
					goto l347
				}
				position++
				{
					position349, tokenIndex349 := position, tokenIndex
					if !_rules[rule_]() {
						goto l349
					}
					goto l350
				l349:
					position, tokenIndex = position349, tokenIndex349
				}
			l350:
				if !_rules[ruleAuthorsGroup]() {
					goto l347
				}
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
				if buffer[position] != rune(')') {
					goto l347
				}
				position++
				add(ruleBasionymAuthorship1, position348)
			}
			return true
		l347:
			position, tokenIndex = position347, tokenIndex347
			return false
		},
		/* 54 BasionymAuthorship2 <- <('(' _? '(' _? AuthorsGroup _? ')' _? ')')> */
		func() bool {
			position353, tokenIndex353 := position, tokenIndex
			{
				position354 := position
				if buffer[position] != rune('(') {
					goto l353
				}
				position++
				{
					position355, tokenIndex355 := position, tokenIndex
					if !_rules[rule_]() {
						goto l355
					}
					goto l356
				l355:
					position, tokenIndex = position355, tokenIndex355
				}
			l356:
				if buffer[position] != rune('(') {
					goto l353
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
				if !_rules[ruleAuthorsGroup]() {
					goto l353
				}
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
				if buffer[position] != rune(')') {
					goto l353
				}
				position++
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
					goto l353
				}
				position++
				add(ruleBasionymAuthorship2, position354)
			}
			return true
		l353:
			position, tokenIndex = position353, tokenIndex353
			return false
		},
		/* 55 AuthorsGroup <- <(AuthorsTeam (_? AuthorEmend? AuthorEx? AuthorsTeam)?)> */
		func() bool {
			position363, tokenIndex363 := position, tokenIndex
			{
				position364 := position
				if !_rules[ruleAuthorsTeam]() {
					goto l363
				}
				{
					position365, tokenIndex365 := position, tokenIndex
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
					{
						position369, tokenIndex369 := position, tokenIndex
						if !_rules[ruleAuthorEmend]() {
							goto l369
						}
						goto l370
					l369:
						position, tokenIndex = position369, tokenIndex369
					}
				l370:
					{
						position371, tokenIndex371 := position, tokenIndex
						if !_rules[ruleAuthorEx]() {
							goto l371
						}
						goto l372
					l371:
						position, tokenIndex = position371, tokenIndex371
					}
				l372:
					if !_rules[ruleAuthorsTeam]() {
						goto l365
					}
					goto l366
				l365:
					position, tokenIndex = position365, tokenIndex365
				}
			l366:
				add(ruleAuthorsGroup, position364)
			}
			return true
		l363:
			position, tokenIndex = position363, tokenIndex363
			return false
		},
		/* 56 AuthorsTeam <- <(Author (AuthorSep Author)* (','? _? Year)?)> */
		func() bool {
			position373, tokenIndex373 := position, tokenIndex
			{
				position374 := position
				if !_rules[ruleAuthor]() {
					goto l373
				}
			l375:
				{
					position376, tokenIndex376 := position, tokenIndex
					if !_rules[ruleAuthorSep]() {
						goto l376
					}
					if !_rules[ruleAuthor]() {
						goto l376
					}
					goto l375
				l376:
					position, tokenIndex = position376, tokenIndex376
				}
				{
					position377, tokenIndex377 := position, tokenIndex
					{
						position379, tokenIndex379 := position, tokenIndex
						if buffer[position] != rune(',') {
							goto l379
						}
						position++
						goto l380
					l379:
						position, tokenIndex = position379, tokenIndex379
					}
				l380:
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
					if !_rules[ruleYear]() {
						goto l377
					}
					goto l378
				l377:
					position, tokenIndex = position377, tokenIndex377
				}
			l378:
				add(ruleAuthorsTeam, position374)
			}
			return true
		l373:
			position, tokenIndex = position373, tokenIndex373
			return false
		},
		/* 57 AuthorSep <- <(AuthorSep1 / AuthorSep2)> */
		func() bool {
			position383, tokenIndex383 := position, tokenIndex
			{
				position384 := position
				{
					position385, tokenIndex385 := position, tokenIndex
					if !_rules[ruleAuthorSep1]() {
						goto l386
					}
					goto l385
				l386:
					position, tokenIndex = position385, tokenIndex385
					if !_rules[ruleAuthorSep2]() {
						goto l383
					}
				}
			l385:
				add(ruleAuthorSep, position384)
			}
			return true
		l383:
			position, tokenIndex = position383, tokenIndex383
			return false
		},
		/* 58 AuthorSep1 <- <(_? (',' _)? ('&' / ('e' 't') / ('a' 'p' 'u' 'd')) _?)> */
		func() bool {
			position387, tokenIndex387 := position, tokenIndex
			{
				position388 := position
				{
					position389, tokenIndex389 := position, tokenIndex
					if !_rules[rule_]() {
						goto l389
					}
					goto l390
				l389:
					position, tokenIndex = position389, tokenIndex389
				}
			l390:
				{
					position391, tokenIndex391 := position, tokenIndex
					if buffer[position] != rune(',') {
						goto l391
					}
					position++
					if !_rules[rule_]() {
						goto l391
					}
					goto l392
				l391:
					position, tokenIndex = position391, tokenIndex391
				}
			l392:
				{
					position393, tokenIndex393 := position, tokenIndex
					if buffer[position] != rune('&') {
						goto l394
					}
					position++
					goto l393
				l394:
					position, tokenIndex = position393, tokenIndex393
					if buffer[position] != rune('e') {
						goto l395
					}
					position++
					if buffer[position] != rune('t') {
						goto l395
					}
					position++
					goto l393
				l395:
					position, tokenIndex = position393, tokenIndex393
					if buffer[position] != rune('a') {
						goto l387
					}
					position++
					if buffer[position] != rune('p') {
						goto l387
					}
					position++
					if buffer[position] != rune('u') {
						goto l387
					}
					position++
					if buffer[position] != rune('d') {
						goto l387
					}
					position++
				}
			l393:
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
				add(ruleAuthorSep1, position388)
			}
			return true
		l387:
			position, tokenIndex = position387, tokenIndex387
			return false
		},
		/* 59 AuthorSep2 <- <(_? ',' _?)> */
		func() bool {
			position398, tokenIndex398 := position, tokenIndex
			{
				position399 := position
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
				if buffer[position] != rune(',') {
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
				add(ruleAuthorSep2, position399)
			}
			return true
		l398:
			position, tokenIndex = position398, tokenIndex398
			return false
		},
		/* 60 AuthorEx <- <((('e' 'x' '.'?) / ('i' 'n')) _)> */
		func() bool {
			position404, tokenIndex404 := position, tokenIndex
			{
				position405 := position
				{
					position406, tokenIndex406 := position, tokenIndex
					if buffer[position] != rune('e') {
						goto l407
					}
					position++
					if buffer[position] != rune('x') {
						goto l407
					}
					position++
					{
						position408, tokenIndex408 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l408
						}
						position++
						goto l409
					l408:
						position, tokenIndex = position408, tokenIndex408
					}
				l409:
					goto l406
				l407:
					position, tokenIndex = position406, tokenIndex406
					if buffer[position] != rune('i') {
						goto l404
					}
					position++
					if buffer[position] != rune('n') {
						goto l404
					}
					position++
				}
			l406:
				if !_rules[rule_]() {
					goto l404
				}
				add(ruleAuthorEx, position405)
			}
			return true
		l404:
			position, tokenIndex = position404, tokenIndex404
			return false
		},
		/* 61 AuthorEmend <- <('e' 'm' 'e' 'n' 'd' '.'? _)> */
		func() bool {
			position410, tokenIndex410 := position, tokenIndex
			{
				position411 := position
				if buffer[position] != rune('e') {
					goto l410
				}
				position++
				if buffer[position] != rune('m') {
					goto l410
				}
				position++
				if buffer[position] != rune('e') {
					goto l410
				}
				position++
				if buffer[position] != rune('n') {
					goto l410
				}
				position++
				if buffer[position] != rune('d') {
					goto l410
				}
				position++
				{
					position412, tokenIndex412 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l412
					}
					position++
					goto l413
				l412:
					position, tokenIndex = position412, tokenIndex412
				}
			l413:
				if !_rules[rule_]() {
					goto l410
				}
				add(ruleAuthorEmend, position411)
			}
			return true
		l410:
			position, tokenIndex = position410, tokenIndex410
			return false
		},
		/* 62 Author <- <(Author1 / Author2 / UnknownAuthor)> */
		func() bool {
			position414, tokenIndex414 := position, tokenIndex
			{
				position415 := position
				{
					position416, tokenIndex416 := position, tokenIndex
					if !_rules[ruleAuthor1]() {
						goto l417
					}
					goto l416
				l417:
					position, tokenIndex = position416, tokenIndex416
					if !_rules[ruleAuthor2]() {
						goto l418
					}
					goto l416
				l418:
					position, tokenIndex = position416, tokenIndex416
					if !_rules[ruleUnknownAuthor]() {
						goto l414
					}
				}
			l416:
				add(ruleAuthor, position415)
			}
			return true
		l414:
			position, tokenIndex = position414, tokenIndex414
			return false
		},
		/* 63 Author1 <- <(Author2 _ Filius)> */
		func() bool {
			position419, tokenIndex419 := position, tokenIndex
			{
				position420 := position
				if !_rules[ruleAuthor2]() {
					goto l419
				}
				if !_rules[rule_]() {
					goto l419
				}
				if !_rules[ruleFilius]() {
					goto l419
				}
				add(ruleAuthor1, position420)
			}
			return true
		l419:
			position, tokenIndex = position419, tokenIndex419
			return false
		},
		/* 64 Author2 <- <(AuthorWord (_? AuthorWord)*)> */
		func() bool {
			position421, tokenIndex421 := position, tokenIndex
			{
				position422 := position
				if !_rules[ruleAuthorWord]() {
					goto l421
				}
			l423:
				{
					position424, tokenIndex424 := position, tokenIndex
					{
						position425, tokenIndex425 := position, tokenIndex
						if !_rules[rule_]() {
							goto l425
						}
						goto l426
					l425:
						position, tokenIndex = position425, tokenIndex425
					}
				l426:
					if !_rules[ruleAuthorWord]() {
						goto l424
					}
					goto l423
				l424:
					position, tokenIndex = position424, tokenIndex424
				}
				add(ruleAuthor2, position422)
			}
			return true
		l421:
			position, tokenIndex = position421, tokenIndex421
			return false
		},
		/* 65 UnknownAuthor <- <('?' / ((('a' 'u' 'c' 't') / ('a' 'n' 'o' 'n')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position427, tokenIndex427 := position, tokenIndex
			{
				position428 := position
				{
					position429, tokenIndex429 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l430
					}
					position++
					goto l429
				l430:
					position, tokenIndex = position429, tokenIndex429
					{
						position431, tokenIndex431 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l432
						}
						position++
						if buffer[position] != rune('u') {
							goto l432
						}
						position++
						if buffer[position] != rune('c') {
							goto l432
						}
						position++
						if buffer[position] != rune('t') {
							goto l432
						}
						position++
						goto l431
					l432:
						position, tokenIndex = position431, tokenIndex431
						if buffer[position] != rune('a') {
							goto l427
						}
						position++
						if buffer[position] != rune('n') {
							goto l427
						}
						position++
						if buffer[position] != rune('o') {
							goto l427
						}
						position++
						if buffer[position] != rune('n') {
							goto l427
						}
						position++
					}
				l431:
					{
						position433, tokenIndex433 := position, tokenIndex
						{
							position435, tokenIndex435 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l434
							}
							position, tokenIndex = position435, tokenIndex435
						}
						goto l433
					l434:
						position, tokenIndex = position433, tokenIndex433
						if buffer[position] != rune('.') {
							goto l427
						}
						position++
					}
				l433:
				}
			l429:
				add(ruleUnknownAuthor, position428)
			}
			return true
		l427:
			position, tokenIndex = position427, tokenIndex427
			return false
		},
		/* 66 AuthorWord <- <(AuthorWord1 / AuthorWord2 / AuthorWord3 / AuthorPrefix)> */
		func() bool {
			position436, tokenIndex436 := position, tokenIndex
			{
				position437 := position
				{
					position438, tokenIndex438 := position, tokenIndex
					if !_rules[ruleAuthorWord1]() {
						goto l439
					}
					goto l438
				l439:
					position, tokenIndex = position438, tokenIndex438
					if !_rules[ruleAuthorWord2]() {
						goto l440
					}
					goto l438
				l440:
					position, tokenIndex = position438, tokenIndex438
					if !_rules[ruleAuthorWord3]() {
						goto l441
					}
					goto l438
				l441:
					position, tokenIndex = position438, tokenIndex438
					if !_rules[ruleAuthorPrefix]() {
						goto l436
					}
				}
			l438:
				add(ruleAuthorWord, position437)
			}
			return true
		l436:
			position, tokenIndex = position436, tokenIndex436
			return false
		},
		/* 67 AuthorWord1 <- <(('a' 'r' 'g' '.') / ('e' 't' ' ' 'a' 'l' '.' '{' '?' '}') / ((('e' 't') / '&') (' ' 'a' 'l') '.'?))> */
		func() bool {
			position442, tokenIndex442 := position, tokenIndex
			{
				position443 := position
				{
					position444, tokenIndex444 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l445
					}
					position++
					if buffer[position] != rune('r') {
						goto l445
					}
					position++
					if buffer[position] != rune('g') {
						goto l445
					}
					position++
					if buffer[position] != rune('.') {
						goto l445
					}
					position++
					goto l444
				l445:
					position, tokenIndex = position444, tokenIndex444
					if buffer[position] != rune('e') {
						goto l446
					}
					position++
					if buffer[position] != rune('t') {
						goto l446
					}
					position++
					if buffer[position] != rune(' ') {
						goto l446
					}
					position++
					if buffer[position] != rune('a') {
						goto l446
					}
					position++
					if buffer[position] != rune('l') {
						goto l446
					}
					position++
					if buffer[position] != rune('.') {
						goto l446
					}
					position++
					if buffer[position] != rune('{') {
						goto l446
					}
					position++
					if buffer[position] != rune('?') {
						goto l446
					}
					position++
					if buffer[position] != rune('}') {
						goto l446
					}
					position++
					goto l444
				l446:
					position, tokenIndex = position444, tokenIndex444
					{
						position447, tokenIndex447 := position, tokenIndex
						if buffer[position] != rune('e') {
							goto l448
						}
						position++
						if buffer[position] != rune('t') {
							goto l448
						}
						position++
						goto l447
					l448:
						position, tokenIndex = position447, tokenIndex447
						if buffer[position] != rune('&') {
							goto l442
						}
						position++
					}
				l447:
					if buffer[position] != rune(' ') {
						goto l442
					}
					position++
					if buffer[position] != rune('a') {
						goto l442
					}
					position++
					if buffer[position] != rune('l') {
						goto l442
					}
					position++
					{
						position449, tokenIndex449 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l449
						}
						position++
						goto l450
					l449:
						position, tokenIndex = position449, tokenIndex449
					}
				l450:
				}
			l444:
				add(ruleAuthorWord1, position443)
			}
			return true
		l442:
			position, tokenIndex = position442, tokenIndex442
			return false
		},
		/* 68 AuthorWord2 <- <(AuthorWord3 dash AuthorWordSoft)> */
		func() bool {
			position451, tokenIndex451 := position, tokenIndex
			{
				position452 := position
				if !_rules[ruleAuthorWord3]() {
					goto l451
				}
				if !_rules[ruledash]() {
					goto l451
				}
				if !_rules[ruleAuthorWordSoft]() {
					goto l451
				}
				add(ruleAuthorWord2, position452)
			}
			return true
		l451:
			position, tokenIndex = position451, tokenIndex451
			return false
		},
		/* 69 AuthorWord3 <- <(AuthorPrefixGlued? AuthorUpperChar (AuthorUpperChar / AuthorLowerChar)* '.'?)> */
		func() bool {
			position453, tokenIndex453 := position, tokenIndex
			{
				position454 := position
				{
					position455, tokenIndex455 := position, tokenIndex
					if !_rules[ruleAuthorPrefixGlued]() {
						goto l455
					}
					goto l456
				l455:
					position, tokenIndex = position455, tokenIndex455
				}
			l456:
				if !_rules[ruleAuthorUpperChar]() {
					goto l453
				}
			l457:
				{
					position458, tokenIndex458 := position, tokenIndex
					{
						position459, tokenIndex459 := position, tokenIndex
						if !_rules[ruleAuthorUpperChar]() {
							goto l460
						}
						goto l459
					l460:
						position, tokenIndex = position459, tokenIndex459
						if !_rules[ruleAuthorLowerChar]() {
							goto l458
						}
					}
				l459:
					goto l457
				l458:
					position, tokenIndex = position458, tokenIndex458
				}
				{
					position461, tokenIndex461 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l461
					}
					position++
					goto l462
				l461:
					position, tokenIndex = position461, tokenIndex461
				}
			l462:
				add(ruleAuthorWord3, position454)
			}
			return true
		l453:
			position, tokenIndex = position453, tokenIndex453
			return false
		},
		/* 70 AuthorWordSoft <- <(((AuthorUpperChar (AuthorUpperChar+ / AuthorLowerChar+)) / AuthorLowerChar+) '.'?)> */
		func() bool {
			position463, tokenIndex463 := position, tokenIndex
			{
				position464 := position
				{
					position465, tokenIndex465 := position, tokenIndex
					if !_rules[ruleAuthorUpperChar]() {
						goto l466
					}
					{
						position467, tokenIndex467 := position, tokenIndex
						if !_rules[ruleAuthorUpperChar]() {
							goto l468
						}
					l469:
						{
							position470, tokenIndex470 := position, tokenIndex
							if !_rules[ruleAuthorUpperChar]() {
								goto l470
							}
							goto l469
						l470:
							position, tokenIndex = position470, tokenIndex470
						}
						goto l467
					l468:
						position, tokenIndex = position467, tokenIndex467
						if !_rules[ruleAuthorLowerChar]() {
							goto l466
						}
					l471:
						{
							position472, tokenIndex472 := position, tokenIndex
							if !_rules[ruleAuthorLowerChar]() {
								goto l472
							}
							goto l471
						l472:
							position, tokenIndex = position472, tokenIndex472
						}
					}
				l467:
					goto l465
				l466:
					position, tokenIndex = position465, tokenIndex465
					if !_rules[ruleAuthorLowerChar]() {
						goto l463
					}
				l473:
					{
						position474, tokenIndex474 := position, tokenIndex
						if !_rules[ruleAuthorLowerChar]() {
							goto l474
						}
						goto l473
					l474:
						position, tokenIndex = position474, tokenIndex474
					}
				}
			l465:
				{
					position475, tokenIndex475 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l475
					}
					position++
					goto l476
				l475:
					position, tokenIndex = position475, tokenIndex475
				}
			l476:
				add(ruleAuthorWordSoft, position464)
			}
			return true
		l463:
			position, tokenIndex = position463, tokenIndex463
			return false
		},
		/* 71 Filius <- <(('f' '.') / ('f' 'i' 'l' '.') / ('f' 'i' 'l' 'i' 'u' 's'))> */
		func() bool {
			position477, tokenIndex477 := position, tokenIndex
			{
				position478 := position
				{
					position479, tokenIndex479 := position, tokenIndex
					if buffer[position] != rune('f') {
						goto l480
					}
					position++
					if buffer[position] != rune('.') {
						goto l480
					}
					position++
					goto l479
				l480:
					position, tokenIndex = position479, tokenIndex479
					if buffer[position] != rune('f') {
						goto l481
					}
					position++
					if buffer[position] != rune('i') {
						goto l481
					}
					position++
					if buffer[position] != rune('l') {
						goto l481
					}
					position++
					if buffer[position] != rune('.') {
						goto l481
					}
					position++
					goto l479
				l481:
					position, tokenIndex = position479, tokenIndex479
					if buffer[position] != rune('f') {
						goto l477
					}
					position++
					if buffer[position] != rune('i') {
						goto l477
					}
					position++
					if buffer[position] != rune('l') {
						goto l477
					}
					position++
					if buffer[position] != rune('i') {
						goto l477
					}
					position++
					if buffer[position] != rune('u') {
						goto l477
					}
					position++
					if buffer[position] != rune('s') {
						goto l477
					}
					position++
				}
			l479:
				add(ruleFilius, position478)
			}
			return true
		l477:
			position, tokenIndex = position477, tokenIndex477
			return false
		},
		/* 72 AuthorPrefixGlued <- <(('d' '\'') / ('O' '\''))> */
		func() bool {
			position482, tokenIndex482 := position, tokenIndex
			{
				position483 := position
				{
					position484, tokenIndex484 := position, tokenIndex
					if buffer[position] != rune('d') {
						goto l485
					}
					position++
					if buffer[position] != rune('\'') {
						goto l485
					}
					position++
					goto l484
				l485:
					position, tokenIndex = position484, tokenIndex484
					if buffer[position] != rune('O') {
						goto l482
					}
					position++
					if buffer[position] != rune('\'') {
						goto l482
					}
					position++
				}
			l484:
				add(ruleAuthorPrefixGlued, position483)
			}
			return true
		l482:
			position, tokenIndex = position482, tokenIndex482
			return false
		},
		/* 73 AuthorPrefix <- <(AuthorPrefix1 / AuthorPrefix2)> */
		func() bool {
			position486, tokenIndex486 := position, tokenIndex
			{
				position487 := position
				{
					position488, tokenIndex488 := position, tokenIndex
					if !_rules[ruleAuthorPrefix1]() {
						goto l489
					}
					goto l488
				l489:
					position, tokenIndex = position488, tokenIndex488
					if !_rules[ruleAuthorPrefix2]() {
						goto l486
					}
				}
			l488:
				add(ruleAuthorPrefix, position487)
			}
			return true
		l486:
			position, tokenIndex = position486, tokenIndex486
			return false
		},
		/* 74 AuthorPrefix2 <- <(('v' '.' (_? ('d' '.'))?) / ('\'' 't'))> */
		func() bool {
			position490, tokenIndex490 := position, tokenIndex
			{
				position491 := position
				{
					position492, tokenIndex492 := position, tokenIndex
					if buffer[position] != rune('v') {
						goto l493
					}
					position++
					if buffer[position] != rune('.') {
						goto l493
					}
					position++
					{
						position494, tokenIndex494 := position, tokenIndex
						{
							position496, tokenIndex496 := position, tokenIndex
							if !_rules[rule_]() {
								goto l496
							}
							goto l497
						l496:
							position, tokenIndex = position496, tokenIndex496
						}
					l497:
						if buffer[position] != rune('d') {
							goto l494
						}
						position++
						if buffer[position] != rune('.') {
							goto l494
						}
						position++
						goto l495
					l494:
						position, tokenIndex = position494, tokenIndex494
					}
				l495:
					goto l492
				l493:
					position, tokenIndex = position492, tokenIndex492
					if buffer[position] != rune('\'') {
						goto l490
					}
					position++
					if buffer[position] != rune('t') {
						goto l490
					}
					position++
				}
			l492:
				add(ruleAuthorPrefix2, position491)
			}
			return true
		l490:
			position, tokenIndex = position490, tokenIndex490
			return false
		},
		/* 75 AuthorPrefix1 <- <((('a' 'b') / ('a' 'f') / ('b' 'i' 's') / ('d' 'a') / ('d' 'e' 'r') / ('d' 'e' 's') / ('d' 'e' 'n') / ('d' 'e' 'l') / ('d' 'e' 'l' 'l' 'a') / ('d' 'e' 'l' 'a') / ('d' 'e') / ('d' 'i') / ('d' 'u') / ('e' 'l') / ('l' 'a') / ('l' 'e') / ('t' 'e' 'r') / ('v' 'a' 'n') / ('d' '\'') / ('i' 'n' '\'' 't') / ('z' 'u' 'r') / ('v' 'o' 'n' (_ ('d' '.'))?) / ('v' (_ 'd')?)) &_)> */
		func() bool {
			position498, tokenIndex498 := position, tokenIndex
			{
				position499 := position
				{
					position500, tokenIndex500 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l501
					}
					position++
					if buffer[position] != rune('b') {
						goto l501
					}
					position++
					goto l500
				l501:
					position, tokenIndex = position500, tokenIndex500
					if buffer[position] != rune('a') {
						goto l502
					}
					position++
					if buffer[position] != rune('f') {
						goto l502
					}
					position++
					goto l500
				l502:
					position, tokenIndex = position500, tokenIndex500
					if buffer[position] != rune('b') {
						goto l503
					}
					position++
					if buffer[position] != rune('i') {
						goto l503
					}
					position++
					if buffer[position] != rune('s') {
						goto l503
					}
					position++
					goto l500
				l503:
					position, tokenIndex = position500, tokenIndex500
					if buffer[position] != rune('d') {
						goto l504
					}
					position++
					if buffer[position] != rune('a') {
						goto l504
					}
					position++
					goto l500
				l504:
					position, tokenIndex = position500, tokenIndex500
					if buffer[position] != rune('d') {
						goto l505
					}
					position++
					if buffer[position] != rune('e') {
						goto l505
					}
					position++
					if buffer[position] != rune('r') {
						goto l505
					}
					position++
					goto l500
				l505:
					position, tokenIndex = position500, tokenIndex500
					if buffer[position] != rune('d') {
						goto l506
					}
					position++
					if buffer[position] != rune('e') {
						goto l506
					}
					position++
					if buffer[position] != rune('s') {
						goto l506
					}
					position++
					goto l500
				l506:
					position, tokenIndex = position500, tokenIndex500
					if buffer[position] != rune('d') {
						goto l507
					}
					position++
					if buffer[position] != rune('e') {
						goto l507
					}
					position++
					if buffer[position] != rune('n') {
						goto l507
					}
					position++
					goto l500
				l507:
					position, tokenIndex = position500, tokenIndex500
					if buffer[position] != rune('d') {
						goto l508
					}
					position++
					if buffer[position] != rune('e') {
						goto l508
					}
					position++
					if buffer[position] != rune('l') {
						goto l508
					}
					position++
					goto l500
				l508:
					position, tokenIndex = position500, tokenIndex500
					if buffer[position] != rune('d') {
						goto l509
					}
					position++
					if buffer[position] != rune('e') {
						goto l509
					}
					position++
					if buffer[position] != rune('l') {
						goto l509
					}
					position++
					if buffer[position] != rune('l') {
						goto l509
					}
					position++
					if buffer[position] != rune('a') {
						goto l509
					}
					position++
					goto l500
				l509:
					position, tokenIndex = position500, tokenIndex500
					if buffer[position] != rune('d') {
						goto l510
					}
					position++
					if buffer[position] != rune('e') {
						goto l510
					}
					position++
					if buffer[position] != rune('l') {
						goto l510
					}
					position++
					if buffer[position] != rune('a') {
						goto l510
					}
					position++
					goto l500
				l510:
					position, tokenIndex = position500, tokenIndex500
					if buffer[position] != rune('d') {
						goto l511
					}
					position++
					if buffer[position] != rune('e') {
						goto l511
					}
					position++
					goto l500
				l511:
					position, tokenIndex = position500, tokenIndex500
					if buffer[position] != rune('d') {
						goto l512
					}
					position++
					if buffer[position] != rune('i') {
						goto l512
					}
					position++
					goto l500
				l512:
					position, tokenIndex = position500, tokenIndex500
					if buffer[position] != rune('d') {
						goto l513
					}
					position++
					if buffer[position] != rune('u') {
						goto l513
					}
					position++
					goto l500
				l513:
					position, tokenIndex = position500, tokenIndex500
					if buffer[position] != rune('e') {
						goto l514
					}
					position++
					if buffer[position] != rune('l') {
						goto l514
					}
					position++
					goto l500
				l514:
					position, tokenIndex = position500, tokenIndex500
					if buffer[position] != rune('l') {
						goto l515
					}
					position++
					if buffer[position] != rune('a') {
						goto l515
					}
					position++
					goto l500
				l515:
					position, tokenIndex = position500, tokenIndex500
					if buffer[position] != rune('l') {
						goto l516
					}
					position++
					if buffer[position] != rune('e') {
						goto l516
					}
					position++
					goto l500
				l516:
					position, tokenIndex = position500, tokenIndex500
					if buffer[position] != rune('t') {
						goto l517
					}
					position++
					if buffer[position] != rune('e') {
						goto l517
					}
					position++
					if buffer[position] != rune('r') {
						goto l517
					}
					position++
					goto l500
				l517:
					position, tokenIndex = position500, tokenIndex500
					if buffer[position] != rune('v') {
						goto l518
					}
					position++
					if buffer[position] != rune('a') {
						goto l518
					}
					position++
					if buffer[position] != rune('n') {
						goto l518
					}
					position++
					goto l500
				l518:
					position, tokenIndex = position500, tokenIndex500
					if buffer[position] != rune('d') {
						goto l519
					}
					position++
					if buffer[position] != rune('\'') {
						goto l519
					}
					position++
					goto l500
				l519:
					position, tokenIndex = position500, tokenIndex500
					if buffer[position] != rune('i') {
						goto l520
					}
					position++
					if buffer[position] != rune('n') {
						goto l520
					}
					position++
					if buffer[position] != rune('\'') {
						goto l520
					}
					position++
					if buffer[position] != rune('t') {
						goto l520
					}
					position++
					goto l500
				l520:
					position, tokenIndex = position500, tokenIndex500
					if buffer[position] != rune('z') {
						goto l521
					}
					position++
					if buffer[position] != rune('u') {
						goto l521
					}
					position++
					if buffer[position] != rune('r') {
						goto l521
					}
					position++
					goto l500
				l521:
					position, tokenIndex = position500, tokenIndex500
					if buffer[position] != rune('v') {
						goto l522
					}
					position++
					if buffer[position] != rune('o') {
						goto l522
					}
					position++
					if buffer[position] != rune('n') {
						goto l522
					}
					position++
					{
						position523, tokenIndex523 := position, tokenIndex
						if !_rules[rule_]() {
							goto l523
						}
						if buffer[position] != rune('d') {
							goto l523
						}
						position++
						if buffer[position] != rune('.') {
							goto l523
						}
						position++
						goto l524
					l523:
						position, tokenIndex = position523, tokenIndex523
					}
				l524:
					goto l500
				l522:
					position, tokenIndex = position500, tokenIndex500
					if buffer[position] != rune('v') {
						goto l498
					}
					position++
					{
						position525, tokenIndex525 := position, tokenIndex
						if !_rules[rule_]() {
							goto l525
						}
						if buffer[position] != rune('d') {
							goto l525
						}
						position++
						goto l526
					l525:
						position, tokenIndex = position525, tokenIndex525
					}
				l526:
				}
			l500:
				{
					position527, tokenIndex527 := position, tokenIndex
					if !_rules[rule_]() {
						goto l498
					}
					position, tokenIndex = position527, tokenIndex527
				}
				add(ruleAuthorPrefix1, position499)
			}
			return true
		l498:
			position, tokenIndex = position498, tokenIndex498
			return false
		},
		/* 76 AuthorUpperChar <- <(hASCII / ('À' / 'Á' / 'Â' / 'Ã' / 'Ä' / 'Å' / 'Æ' / 'Ç' / 'È' / 'É' / 'Ê' / 'Ë' / 'Ì' / 'Í' / 'Î' / 'Ï' / 'Ð' / 'Ñ' / 'Ò' / 'Ó' / 'Ô' / 'Õ' / 'Ö' / 'Ø' / 'Ù' / 'Ú' / 'Û' / 'Ü' / 'Ý' / 'Ć' / 'Č' / 'Ď' / 'İ' / 'Ķ' / 'Ĺ' / 'ĺ' / 'Ľ' / 'ľ' / 'Ł' / 'ł' / 'Ņ' / 'Ō' / 'Ő' / 'Œ' / 'Ř' / 'Ś' / 'Ŝ' / 'Ş' / 'Š' / 'Ÿ' / 'Ź' / 'Ż' / 'Ž' / 'ƒ' / 'Ǿ' / 'Ș' / 'Ț'))> */
		func() bool {
			position528, tokenIndex528 := position, tokenIndex
			{
				position529 := position
				{
					position530, tokenIndex530 := position, tokenIndex
					if !_rules[rulehASCII]() {
						goto l531
					}
					goto l530
				l531:
					position, tokenIndex = position530, tokenIndex530
					{
						position532, tokenIndex532 := position, tokenIndex
						if buffer[position] != rune('À') {
							goto l533
						}
						position++
						goto l532
					l533:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Á') {
							goto l534
						}
						position++
						goto l532
					l534:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Â') {
							goto l535
						}
						position++
						goto l532
					l535:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ã') {
							goto l536
						}
						position++
						goto l532
					l536:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ä') {
							goto l537
						}
						position++
						goto l532
					l537:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Å') {
							goto l538
						}
						position++
						goto l532
					l538:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Æ') {
							goto l539
						}
						position++
						goto l532
					l539:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ç') {
							goto l540
						}
						position++
						goto l532
					l540:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('È') {
							goto l541
						}
						position++
						goto l532
					l541:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('É') {
							goto l542
						}
						position++
						goto l532
					l542:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ê') {
							goto l543
						}
						position++
						goto l532
					l543:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ë') {
							goto l544
						}
						position++
						goto l532
					l544:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ì') {
							goto l545
						}
						position++
						goto l532
					l545:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Í') {
							goto l546
						}
						position++
						goto l532
					l546:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Î') {
							goto l547
						}
						position++
						goto l532
					l547:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ï') {
							goto l548
						}
						position++
						goto l532
					l548:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ð') {
							goto l549
						}
						position++
						goto l532
					l549:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ñ') {
							goto l550
						}
						position++
						goto l532
					l550:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ò') {
							goto l551
						}
						position++
						goto l532
					l551:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ó') {
							goto l552
						}
						position++
						goto l532
					l552:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ô') {
							goto l553
						}
						position++
						goto l532
					l553:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Õ') {
							goto l554
						}
						position++
						goto l532
					l554:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ö') {
							goto l555
						}
						position++
						goto l532
					l555:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ø') {
							goto l556
						}
						position++
						goto l532
					l556:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ù') {
							goto l557
						}
						position++
						goto l532
					l557:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ú') {
							goto l558
						}
						position++
						goto l532
					l558:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Û') {
							goto l559
						}
						position++
						goto l532
					l559:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ü') {
							goto l560
						}
						position++
						goto l532
					l560:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ý') {
							goto l561
						}
						position++
						goto l532
					l561:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ć') {
							goto l562
						}
						position++
						goto l532
					l562:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Č') {
							goto l563
						}
						position++
						goto l532
					l563:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ď') {
							goto l564
						}
						position++
						goto l532
					l564:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('İ') {
							goto l565
						}
						position++
						goto l532
					l565:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ķ') {
							goto l566
						}
						position++
						goto l532
					l566:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ĺ') {
							goto l567
						}
						position++
						goto l532
					l567:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('ĺ') {
							goto l568
						}
						position++
						goto l532
					l568:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ľ') {
							goto l569
						}
						position++
						goto l532
					l569:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('ľ') {
							goto l570
						}
						position++
						goto l532
					l570:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ł') {
							goto l571
						}
						position++
						goto l532
					l571:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('ł') {
							goto l572
						}
						position++
						goto l532
					l572:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ņ') {
							goto l573
						}
						position++
						goto l532
					l573:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ō') {
							goto l574
						}
						position++
						goto l532
					l574:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ő') {
							goto l575
						}
						position++
						goto l532
					l575:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Œ') {
							goto l576
						}
						position++
						goto l532
					l576:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ř') {
							goto l577
						}
						position++
						goto l532
					l577:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ś') {
							goto l578
						}
						position++
						goto l532
					l578:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ŝ') {
							goto l579
						}
						position++
						goto l532
					l579:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ş') {
							goto l580
						}
						position++
						goto l532
					l580:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Š') {
							goto l581
						}
						position++
						goto l532
					l581:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ÿ') {
							goto l582
						}
						position++
						goto l532
					l582:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ź') {
							goto l583
						}
						position++
						goto l532
					l583:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ż') {
							goto l584
						}
						position++
						goto l532
					l584:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ž') {
							goto l585
						}
						position++
						goto l532
					l585:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('ƒ') {
							goto l586
						}
						position++
						goto l532
					l586:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ǿ') {
							goto l587
						}
						position++
						goto l532
					l587:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ș') {
							goto l588
						}
						position++
						goto l532
					l588:
						position, tokenIndex = position532, tokenIndex532
						if buffer[position] != rune('Ț') {
							goto l528
						}
						position++
					}
				l532:
				}
			l530:
				add(ruleAuthorUpperChar, position529)
			}
			return true
		l528:
			position, tokenIndex = position528, tokenIndex528
			return false
		},
		/* 77 AuthorLowerChar <- <(lASCII / ('à' / 'á' / 'â' / 'ã' / 'ä' / 'å' / 'æ' / 'ç' / 'è' / 'é' / 'ê' / 'ë' / 'ì' / 'í' / 'î' / 'ï' / 'ð' / 'ñ' / 'ò' / 'ó' / 'ó' / 'ô' / 'õ' / 'ö' / 'ø' / 'ù' / 'ú' / 'û' / 'ü' / 'ý' / 'ÿ' / 'ā' / 'ă' / 'ą' / 'ć' / 'ĉ' / 'č' / 'ď' / 'đ' / '\'' / 'ē' / 'ĕ' / 'ė' / 'ę' / 'ě' / 'ğ' / 'ī' / 'ĭ' / 'İ' / 'ı' / 'ĺ' / 'ľ' / 'ł' / 'ń' / 'ņ' / 'ň' / 'ŏ' / 'ő' / 'œ' / 'ŕ' / 'ř' / 'ś' / 'ş' / 'š' / 'ţ' / 'ť' / 'ũ' / 'ū' / 'ŭ' / 'ů' / 'ű' / 'ź' / 'ż' / 'ž' / 'ſ' / 'ǎ' / 'ǔ' / 'ǧ' / 'ș' / 'ț' / 'ȳ' / 'ß'))> */
		func() bool {
			position589, tokenIndex589 := position, tokenIndex
			{
				position590 := position
				{
					position591, tokenIndex591 := position, tokenIndex
					if !_rules[rulelASCII]() {
						goto l592
					}
					goto l591
				l592:
					position, tokenIndex = position591, tokenIndex591
					{
						position593, tokenIndex593 := position, tokenIndex
						if buffer[position] != rune('à') {
							goto l594
						}
						position++
						goto l593
					l594:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('á') {
							goto l595
						}
						position++
						goto l593
					l595:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('â') {
							goto l596
						}
						position++
						goto l593
					l596:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ã') {
							goto l597
						}
						position++
						goto l593
					l597:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ä') {
							goto l598
						}
						position++
						goto l593
					l598:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('å') {
							goto l599
						}
						position++
						goto l593
					l599:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('æ') {
							goto l600
						}
						position++
						goto l593
					l600:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ç') {
							goto l601
						}
						position++
						goto l593
					l601:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('è') {
							goto l602
						}
						position++
						goto l593
					l602:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('é') {
							goto l603
						}
						position++
						goto l593
					l603:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ê') {
							goto l604
						}
						position++
						goto l593
					l604:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ë') {
							goto l605
						}
						position++
						goto l593
					l605:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ì') {
							goto l606
						}
						position++
						goto l593
					l606:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('í') {
							goto l607
						}
						position++
						goto l593
					l607:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('î') {
							goto l608
						}
						position++
						goto l593
					l608:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ï') {
							goto l609
						}
						position++
						goto l593
					l609:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ð') {
							goto l610
						}
						position++
						goto l593
					l610:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ñ') {
							goto l611
						}
						position++
						goto l593
					l611:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ò') {
							goto l612
						}
						position++
						goto l593
					l612:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ó') {
							goto l613
						}
						position++
						goto l593
					l613:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ó') {
							goto l614
						}
						position++
						goto l593
					l614:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ô') {
							goto l615
						}
						position++
						goto l593
					l615:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('õ') {
							goto l616
						}
						position++
						goto l593
					l616:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ö') {
							goto l617
						}
						position++
						goto l593
					l617:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ø') {
							goto l618
						}
						position++
						goto l593
					l618:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ù') {
							goto l619
						}
						position++
						goto l593
					l619:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ú') {
							goto l620
						}
						position++
						goto l593
					l620:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('û') {
							goto l621
						}
						position++
						goto l593
					l621:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ü') {
							goto l622
						}
						position++
						goto l593
					l622:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ý') {
							goto l623
						}
						position++
						goto l593
					l623:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ÿ') {
							goto l624
						}
						position++
						goto l593
					l624:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ā') {
							goto l625
						}
						position++
						goto l593
					l625:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ă') {
							goto l626
						}
						position++
						goto l593
					l626:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ą') {
							goto l627
						}
						position++
						goto l593
					l627:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ć') {
							goto l628
						}
						position++
						goto l593
					l628:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ĉ') {
							goto l629
						}
						position++
						goto l593
					l629:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('č') {
							goto l630
						}
						position++
						goto l593
					l630:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ď') {
							goto l631
						}
						position++
						goto l593
					l631:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('đ') {
							goto l632
						}
						position++
						goto l593
					l632:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('\'') {
							goto l633
						}
						position++
						goto l593
					l633:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ē') {
							goto l634
						}
						position++
						goto l593
					l634:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ĕ') {
							goto l635
						}
						position++
						goto l593
					l635:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ė') {
							goto l636
						}
						position++
						goto l593
					l636:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ę') {
							goto l637
						}
						position++
						goto l593
					l637:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ě') {
							goto l638
						}
						position++
						goto l593
					l638:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ğ') {
							goto l639
						}
						position++
						goto l593
					l639:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ī') {
							goto l640
						}
						position++
						goto l593
					l640:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ĭ') {
							goto l641
						}
						position++
						goto l593
					l641:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('İ') {
							goto l642
						}
						position++
						goto l593
					l642:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ı') {
							goto l643
						}
						position++
						goto l593
					l643:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ĺ') {
							goto l644
						}
						position++
						goto l593
					l644:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ľ') {
							goto l645
						}
						position++
						goto l593
					l645:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ł') {
							goto l646
						}
						position++
						goto l593
					l646:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ń') {
							goto l647
						}
						position++
						goto l593
					l647:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ņ') {
							goto l648
						}
						position++
						goto l593
					l648:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ň') {
							goto l649
						}
						position++
						goto l593
					l649:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ŏ') {
							goto l650
						}
						position++
						goto l593
					l650:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ő') {
							goto l651
						}
						position++
						goto l593
					l651:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('œ') {
							goto l652
						}
						position++
						goto l593
					l652:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ŕ') {
							goto l653
						}
						position++
						goto l593
					l653:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ř') {
							goto l654
						}
						position++
						goto l593
					l654:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ś') {
							goto l655
						}
						position++
						goto l593
					l655:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ş') {
							goto l656
						}
						position++
						goto l593
					l656:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('š') {
							goto l657
						}
						position++
						goto l593
					l657:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ţ') {
							goto l658
						}
						position++
						goto l593
					l658:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ť') {
							goto l659
						}
						position++
						goto l593
					l659:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ũ') {
							goto l660
						}
						position++
						goto l593
					l660:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ū') {
							goto l661
						}
						position++
						goto l593
					l661:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ŭ') {
							goto l662
						}
						position++
						goto l593
					l662:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ů') {
							goto l663
						}
						position++
						goto l593
					l663:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ű') {
							goto l664
						}
						position++
						goto l593
					l664:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ź') {
							goto l665
						}
						position++
						goto l593
					l665:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ż') {
							goto l666
						}
						position++
						goto l593
					l666:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ž') {
							goto l667
						}
						position++
						goto l593
					l667:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ſ') {
							goto l668
						}
						position++
						goto l593
					l668:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ǎ') {
							goto l669
						}
						position++
						goto l593
					l669:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ǔ') {
							goto l670
						}
						position++
						goto l593
					l670:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ǧ') {
							goto l671
						}
						position++
						goto l593
					l671:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ș') {
							goto l672
						}
						position++
						goto l593
					l672:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ț') {
							goto l673
						}
						position++
						goto l593
					l673:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ȳ') {
							goto l674
						}
						position++
						goto l593
					l674:
						position, tokenIndex = position593, tokenIndex593
						if buffer[position] != rune('ß') {
							goto l589
						}
						position++
					}
				l593:
				}
			l591:
				add(ruleAuthorLowerChar, position590)
			}
			return true
		l589:
			position, tokenIndex = position589, tokenIndex589
			return false
		},
		/* 78 Year <- <(YearRange / YearApprox / YearWithParens / YearWithPage / YearWithDot / YearWithChar / YearNum)> */
		func() bool {
			position675, tokenIndex675 := position, tokenIndex
			{
				position676 := position
				{
					position677, tokenIndex677 := position, tokenIndex
					if !_rules[ruleYearRange]() {
						goto l678
					}
					goto l677
				l678:
					position, tokenIndex = position677, tokenIndex677
					if !_rules[ruleYearApprox]() {
						goto l679
					}
					goto l677
				l679:
					position, tokenIndex = position677, tokenIndex677
					if !_rules[ruleYearWithParens]() {
						goto l680
					}
					goto l677
				l680:
					position, tokenIndex = position677, tokenIndex677
					if !_rules[ruleYearWithPage]() {
						goto l681
					}
					goto l677
				l681:
					position, tokenIndex = position677, tokenIndex677
					if !_rules[ruleYearWithDot]() {
						goto l682
					}
					goto l677
				l682:
					position, tokenIndex = position677, tokenIndex677
					if !_rules[ruleYearWithChar]() {
						goto l683
					}
					goto l677
				l683:
					position, tokenIndex = position677, tokenIndex677
					if !_rules[ruleYearNum]() {
						goto l675
					}
				}
			l677:
				add(ruleYear, position676)
			}
			return true
		l675:
			position, tokenIndex = position675, tokenIndex675
			return false
		},
		/* 79 YearRange <- <(YearNum dash (nums+ ('a' / 'b' / 'c' / 'd' / 'e' / 'f' / 'g' / 'h' / 'i' / 'j' / 'k' / 'l' / 'm' / 'n' / 'o' / 'p' / 'q' / 'r' / 's' / 't' / 'u' / 'v' / 'w' / 'x' / 'y' / 'z' / '?')*))> */
		func() bool {
			position684, tokenIndex684 := position, tokenIndex
			{
				position685 := position
				if !_rules[ruleYearNum]() {
					goto l684
				}
				if !_rules[ruledash]() {
					goto l684
				}
				if !_rules[rulenums]() {
					goto l684
				}
			l686:
				{
					position687, tokenIndex687 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l687
					}
					goto l686
				l687:
					position, tokenIndex = position687, tokenIndex687
				}
			l688:
				{
					position689, tokenIndex689 := position, tokenIndex
					{
						position690, tokenIndex690 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l691
						}
						position++
						goto l690
					l691:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('b') {
							goto l692
						}
						position++
						goto l690
					l692:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('c') {
							goto l693
						}
						position++
						goto l690
					l693:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('d') {
							goto l694
						}
						position++
						goto l690
					l694:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('e') {
							goto l695
						}
						position++
						goto l690
					l695:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('f') {
							goto l696
						}
						position++
						goto l690
					l696:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('g') {
							goto l697
						}
						position++
						goto l690
					l697:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('h') {
							goto l698
						}
						position++
						goto l690
					l698:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('i') {
							goto l699
						}
						position++
						goto l690
					l699:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('j') {
							goto l700
						}
						position++
						goto l690
					l700:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('k') {
							goto l701
						}
						position++
						goto l690
					l701:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('l') {
							goto l702
						}
						position++
						goto l690
					l702:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('m') {
							goto l703
						}
						position++
						goto l690
					l703:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('n') {
							goto l704
						}
						position++
						goto l690
					l704:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('o') {
							goto l705
						}
						position++
						goto l690
					l705:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('p') {
							goto l706
						}
						position++
						goto l690
					l706:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('q') {
							goto l707
						}
						position++
						goto l690
					l707:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('r') {
							goto l708
						}
						position++
						goto l690
					l708:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('s') {
							goto l709
						}
						position++
						goto l690
					l709:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('t') {
							goto l710
						}
						position++
						goto l690
					l710:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('u') {
							goto l711
						}
						position++
						goto l690
					l711:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('v') {
							goto l712
						}
						position++
						goto l690
					l712:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('w') {
							goto l713
						}
						position++
						goto l690
					l713:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('x') {
							goto l714
						}
						position++
						goto l690
					l714:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('y') {
							goto l715
						}
						position++
						goto l690
					l715:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('z') {
							goto l716
						}
						position++
						goto l690
					l716:
						position, tokenIndex = position690, tokenIndex690
						if buffer[position] != rune('?') {
							goto l689
						}
						position++
					}
				l690:
					goto l688
				l689:
					position, tokenIndex = position689, tokenIndex689
				}
				add(ruleYearRange, position685)
			}
			return true
		l684:
			position, tokenIndex = position684, tokenIndex684
			return false
		},
		/* 80 YearWithDot <- <(YearNum '.')> */
		func() bool {
			position717, tokenIndex717 := position, tokenIndex
			{
				position718 := position
				if !_rules[ruleYearNum]() {
					goto l717
				}
				if buffer[position] != rune('.') {
					goto l717
				}
				position++
				add(ruleYearWithDot, position718)
			}
			return true
		l717:
			position, tokenIndex = position717, tokenIndex717
			return false
		},
		/* 81 YearApprox <- <('[' _? YearNum _? ']')> */
		func() bool {
			position719, tokenIndex719 := position, tokenIndex
			{
				position720 := position
				if buffer[position] != rune('[') {
					goto l719
				}
				position++
				{
					position721, tokenIndex721 := position, tokenIndex
					if !_rules[rule_]() {
						goto l721
					}
					goto l722
				l721:
					position, tokenIndex = position721, tokenIndex721
				}
			l722:
				if !_rules[ruleYearNum]() {
					goto l719
				}
				{
					position723, tokenIndex723 := position, tokenIndex
					if !_rules[rule_]() {
						goto l723
					}
					goto l724
				l723:
					position, tokenIndex = position723, tokenIndex723
				}
			l724:
				if buffer[position] != rune(']') {
					goto l719
				}
				position++
				add(ruleYearApprox, position720)
			}
			return true
		l719:
			position, tokenIndex = position719, tokenIndex719
			return false
		},
		/* 82 YearWithPage <- <((YearWithChar / YearNum) _? ':' _? nums+)> */
		func() bool {
			position725, tokenIndex725 := position, tokenIndex
			{
				position726 := position
				{
					position727, tokenIndex727 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l728
					}
					goto l727
				l728:
					position, tokenIndex = position727, tokenIndex727
					if !_rules[ruleYearNum]() {
						goto l725
					}
				}
			l727:
				{
					position729, tokenIndex729 := position, tokenIndex
					if !_rules[rule_]() {
						goto l729
					}
					goto l730
				l729:
					position, tokenIndex = position729, tokenIndex729
				}
			l730:
				if buffer[position] != rune(':') {
					goto l725
				}
				position++
				{
					position731, tokenIndex731 := position, tokenIndex
					if !_rules[rule_]() {
						goto l731
					}
					goto l732
				l731:
					position, tokenIndex = position731, tokenIndex731
				}
			l732:
				if !_rules[rulenums]() {
					goto l725
				}
			l733:
				{
					position734, tokenIndex734 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l734
					}
					goto l733
				l734:
					position, tokenIndex = position734, tokenIndex734
				}
				add(ruleYearWithPage, position726)
			}
			return true
		l725:
			position, tokenIndex = position725, tokenIndex725
			return false
		},
		/* 83 YearWithParens <- <('(' (YearWithChar / YearNum) ')')> */
		func() bool {
			position735, tokenIndex735 := position, tokenIndex
			{
				position736 := position
				if buffer[position] != rune('(') {
					goto l735
				}
				position++
				{
					position737, tokenIndex737 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l738
					}
					goto l737
				l738:
					position, tokenIndex = position737, tokenIndex737
					if !_rules[ruleYearNum]() {
						goto l735
					}
				}
			l737:
				if buffer[position] != rune(')') {
					goto l735
				}
				position++
				add(ruleYearWithParens, position736)
			}
			return true
		l735:
			position, tokenIndex = position735, tokenIndex735
			return false
		},
		/* 84 YearWithChar <- <(YearNum lASCII)> */
		func() bool {
			position739, tokenIndex739 := position, tokenIndex
			{
				position740 := position
				if !_rules[ruleYearNum]() {
					goto l739
				}
				if !_rules[rulelASCII]() {
					goto l739
				}
				add(ruleYearWithChar, position740)
			}
			return true
		l739:
			position, tokenIndex = position739, tokenIndex739
			return false
		},
		/* 85 YearNum <- <(('1' / '2') ('0' / '7' / '8' / '9') nums (nums / '?') '?'*)> */
		func() bool {
			position741, tokenIndex741 := position, tokenIndex
			{
				position742 := position
				{
					position743, tokenIndex743 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l744
					}
					position++
					goto l743
				l744:
					position, tokenIndex = position743, tokenIndex743
					if buffer[position] != rune('2') {
						goto l741
					}
					position++
				}
			l743:
				{
					position745, tokenIndex745 := position, tokenIndex
					if buffer[position] != rune('0') {
						goto l746
					}
					position++
					goto l745
				l746:
					position, tokenIndex = position745, tokenIndex745
					if buffer[position] != rune('7') {
						goto l747
					}
					position++
					goto l745
				l747:
					position, tokenIndex = position745, tokenIndex745
					if buffer[position] != rune('8') {
						goto l748
					}
					position++
					goto l745
				l748:
					position, tokenIndex = position745, tokenIndex745
					if buffer[position] != rune('9') {
						goto l741
					}
					position++
				}
			l745:
				if !_rules[rulenums]() {
					goto l741
				}
				{
					position749, tokenIndex749 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l750
					}
					goto l749
				l750:
					position, tokenIndex = position749, tokenIndex749
					if buffer[position] != rune('?') {
						goto l741
					}
					position++
				}
			l749:
			l751:
				{
					position752, tokenIndex752 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l752
					}
					position++
					goto l751
				l752:
					position, tokenIndex = position752, tokenIndex752
				}
				add(ruleYearNum, position742)
			}
			return true
		l741:
			position, tokenIndex = position741, tokenIndex741
			return false
		},
		/* 86 NameUpperChar <- <(UpperChar / UpperCharExtended)> */
		func() bool {
			position753, tokenIndex753 := position, tokenIndex
			{
				position754 := position
				{
					position755, tokenIndex755 := position, tokenIndex
					if !_rules[ruleUpperChar]() {
						goto l756
					}
					goto l755
				l756:
					position, tokenIndex = position755, tokenIndex755
					if !_rules[ruleUpperCharExtended]() {
						goto l753
					}
				}
			l755:
				add(ruleNameUpperChar, position754)
			}
			return true
		l753:
			position, tokenIndex = position753, tokenIndex753
			return false
		},
		/* 87 UpperCharExtended <- <('Æ' / 'Œ' / 'Ö')> */
		func() bool {
			position757, tokenIndex757 := position, tokenIndex
			{
				position758 := position
				{
					position759, tokenIndex759 := position, tokenIndex
					if buffer[position] != rune('Æ') {
						goto l760
					}
					position++
					goto l759
				l760:
					position, tokenIndex = position759, tokenIndex759
					if buffer[position] != rune('Œ') {
						goto l761
					}
					position++
					goto l759
				l761:
					position, tokenIndex = position759, tokenIndex759
					if buffer[position] != rune('Ö') {
						goto l757
					}
					position++
				}
			l759:
				add(ruleUpperCharExtended, position758)
			}
			return true
		l757:
			position, tokenIndex = position757, tokenIndex757
			return false
		},
		/* 88 UpperChar <- <hASCII> */
		func() bool {
			position762, tokenIndex762 := position, tokenIndex
			{
				position763 := position
				if !_rules[rulehASCII]() {
					goto l762
				}
				add(ruleUpperChar, position763)
			}
			return true
		l762:
			position, tokenIndex = position762, tokenIndex762
			return false
		},
		/* 89 NameLowerChar <- <(LowerChar / LowerCharExtended / MiscodedChar)> */
		func() bool {
			position764, tokenIndex764 := position, tokenIndex
			{
				position765 := position
				{
					position766, tokenIndex766 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l767
					}
					goto l766
				l767:
					position, tokenIndex = position766, tokenIndex766
					if !_rules[ruleLowerCharExtended]() {
						goto l768
					}
					goto l766
				l768:
					position, tokenIndex = position766, tokenIndex766
					if !_rules[ruleMiscodedChar]() {
						goto l764
					}
				}
			l766:
				add(ruleNameLowerChar, position765)
			}
			return true
		l764:
			position, tokenIndex = position764, tokenIndex764
			return false
		},
		/* 90 MiscodedChar <- <'�'> */
		func() bool {
			position769, tokenIndex769 := position, tokenIndex
			{
				position770 := position
				if buffer[position] != rune('�') {
					goto l769
				}
				position++
				add(ruleMiscodedChar, position770)
			}
			return true
		l769:
			position, tokenIndex = position769, tokenIndex769
			return false
		},
		/* 91 LowerCharExtended <- <('æ' / 'œ' / 'ſ' / 'à' / 'â' / 'å' / 'ã' / 'ä' / 'á' / 'ç' / 'č' / 'é' / 'è' / 'í' / 'ì' / 'ï' / 'ň' / 'ñ' / 'ñ' / 'ó' / 'ò' / 'ô' / 'ø' / 'õ' / 'ö' / 'ú' / 'ù' / 'ü' / 'ŕ' / 'ř' / 'ŗ' / 'š' / 'š' / 'ş' / 'ž')> */
		func() bool {
			position771, tokenIndex771 := position, tokenIndex
			{
				position772 := position
				{
					position773, tokenIndex773 := position, tokenIndex
					if buffer[position] != rune('æ') {
						goto l774
					}
					position++
					goto l773
				l774:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('œ') {
						goto l775
					}
					position++
					goto l773
				l775:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('ſ') {
						goto l776
					}
					position++
					goto l773
				l776:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('à') {
						goto l777
					}
					position++
					goto l773
				l777:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('â') {
						goto l778
					}
					position++
					goto l773
				l778:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('å') {
						goto l779
					}
					position++
					goto l773
				l779:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('ã') {
						goto l780
					}
					position++
					goto l773
				l780:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('ä') {
						goto l781
					}
					position++
					goto l773
				l781:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('á') {
						goto l782
					}
					position++
					goto l773
				l782:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('ç') {
						goto l783
					}
					position++
					goto l773
				l783:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('č') {
						goto l784
					}
					position++
					goto l773
				l784:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('é') {
						goto l785
					}
					position++
					goto l773
				l785:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('è') {
						goto l786
					}
					position++
					goto l773
				l786:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('í') {
						goto l787
					}
					position++
					goto l773
				l787:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('ì') {
						goto l788
					}
					position++
					goto l773
				l788:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('ï') {
						goto l789
					}
					position++
					goto l773
				l789:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('ň') {
						goto l790
					}
					position++
					goto l773
				l790:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('ñ') {
						goto l791
					}
					position++
					goto l773
				l791:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('ñ') {
						goto l792
					}
					position++
					goto l773
				l792:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('ó') {
						goto l793
					}
					position++
					goto l773
				l793:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('ò') {
						goto l794
					}
					position++
					goto l773
				l794:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('ô') {
						goto l795
					}
					position++
					goto l773
				l795:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('ø') {
						goto l796
					}
					position++
					goto l773
				l796:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('õ') {
						goto l797
					}
					position++
					goto l773
				l797:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('ö') {
						goto l798
					}
					position++
					goto l773
				l798:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('ú') {
						goto l799
					}
					position++
					goto l773
				l799:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('ù') {
						goto l800
					}
					position++
					goto l773
				l800:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('ü') {
						goto l801
					}
					position++
					goto l773
				l801:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('ŕ') {
						goto l802
					}
					position++
					goto l773
				l802:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('ř') {
						goto l803
					}
					position++
					goto l773
				l803:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('ŗ') {
						goto l804
					}
					position++
					goto l773
				l804:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('š') {
						goto l805
					}
					position++
					goto l773
				l805:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('š') {
						goto l806
					}
					position++
					goto l773
				l806:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('ş') {
						goto l807
					}
					position++
					goto l773
				l807:
					position, tokenIndex = position773, tokenIndex773
					if buffer[position] != rune('ž') {
						goto l771
					}
					position++
				}
			l773:
				add(ruleLowerCharExtended, position772)
			}
			return true
		l771:
			position, tokenIndex = position771, tokenIndex771
			return false
		},
		/* 92 LowerChar <- <([a-z] / 'ë')> */
		func() bool {
			position808, tokenIndex808 := position, tokenIndex
			{
				position809 := position
				{
					position810, tokenIndex810 := position, tokenIndex
					if c := buffer[position]; c < rune('a') || c > rune('z') {
						goto l811
					}
					position++
					goto l810
				l811:
					position, tokenIndex = position810, tokenIndex810
					if buffer[position] != rune('ë') {
						goto l808
					}
					position++
				}
			l810:
				add(ruleLowerChar, position809)
			}
			return true
		l808:
			position, tokenIndex = position808, tokenIndex808
			return false
		},
		/* 93 SpaceCharEOI <- <(_ / !.)> */
		func() bool {
			position812, tokenIndex812 := position, tokenIndex
			{
				position813 := position
				{
					position814, tokenIndex814 := position, tokenIndex
					if !_rules[rule_]() {
						goto l815
					}
					goto l814
				l815:
					position, tokenIndex = position814, tokenIndex814
					{
						position816, tokenIndex816 := position, tokenIndex
						if !matchDot() {
							goto l816
						}
						goto l812
					l816:
						position, tokenIndex = position816, tokenIndex816
					}
				}
			l814:
				add(ruleSpaceCharEOI, position813)
			}
			return true
		l812:
			position, tokenIndex = position812, tokenIndex812
			return false
		},
		/* 94 WordBorderChar <- <(_ / (';' / '.' / ',' / ';' / '(' / ')'))> */
		nil,
		/* 95 nums <- <[0-9]> */
		func() bool {
			position818, tokenIndex818 := position, tokenIndex
			{
				position819 := position
				if c := buffer[position]; c < rune('0') || c > rune('9') {
					goto l818
				}
				position++
				add(rulenums, position819)
			}
			return true
		l818:
			position, tokenIndex = position818, tokenIndex818
			return false
		},
		/* 96 lASCII <- <[a-z]> */
		func() bool {
			position820, tokenIndex820 := position, tokenIndex
			{
				position821 := position
				if c := buffer[position]; c < rune('a') || c > rune('z') {
					goto l820
				}
				position++
				add(rulelASCII, position821)
			}
			return true
		l820:
			position, tokenIndex = position820, tokenIndex820
			return false
		},
		/* 97 hASCII <- <[A-Z]> */
		func() bool {
			position822, tokenIndex822 := position, tokenIndex
			{
				position823 := position
				if c := buffer[position]; c < rune('A') || c > rune('Z') {
					goto l822
				}
				position++
				add(rulehASCII, position823)
			}
			return true
		l822:
			position, tokenIndex = position822, tokenIndex822
			return false
		},
		/* 98 apostr <- <'\''> */
		func() bool {
			position824, tokenIndex824 := position, tokenIndex
			{
				position825 := position
				if buffer[position] != rune('\'') {
					goto l824
				}
				position++
				add(ruleapostr, position825)
			}
			return true
		l824:
			position, tokenIndex = position824, tokenIndex824
			return false
		},
		/* 99 dash <- <'-'> */
		func() bool {
			position826, tokenIndex826 := position, tokenIndex
			{
				position827 := position
				if buffer[position] != rune('-') {
					goto l826
				}
				position++
				add(ruledash, position827)
			}
			return true
		l826:
			position, tokenIndex = position826, tokenIndex826
			return false
		},
		/* 100 _ <- <' '+> */
		func() bool {
			position828, tokenIndex828 := position, tokenIndex
			{
				position829 := position
				if buffer[position] != rune(' ') {
					goto l828
				}
				position++
			l830:
				{
					position831, tokenIndex831 := position, tokenIndex
					if buffer[position] != rune(' ') {
						goto l831
					}
					position++
					goto l830
				l831:
					position, tokenIndex = position831, tokenIndex831
				}
				add(rule_, position829)
			}
			return true
		l828:
			position, tokenIndex = position828, tokenIndex828
			return false
		},
	}
	p.rules = _rules
}
