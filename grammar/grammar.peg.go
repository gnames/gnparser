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
	ruleWordBorderChar
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
	"WordBorderChar",
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
	rules  [107]func() bool
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
		/* 1 Tail <- <((_ / ',') .*)?> */
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
						if buffer[position] != rune(',') {
							goto l7
						}
						position++
					}
				l9:
				l11:
					{
						position12, tokenIndex12 := position, tokenIndex
						if !matchDot() {
							goto l12
						}
						goto l11
					l12:
						position, tokenIndex = position12, tokenIndex12
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
			position13, tokenIndex13 := position, tokenIndex
			{
				position14 := position
				{
					position15, tokenIndex15 := position, tokenIndex
					if !_rules[ruleNamedHybrid]() {
						goto l16
					}
					goto l15
				l16:
					position, tokenIndex = position15, tokenIndex15
					if !_rules[ruleHybridFormula]() {
						goto l17
					}
					goto l15
				l17:
					position, tokenIndex = position15, tokenIndex15
					if !_rules[ruleSingleName]() {
						goto l13
					}
				}
			l15:
				add(ruleName, position14)
			}
			return true
		l13:
			position, tokenIndex = position13, tokenIndex13
			return false
		},
		/* 3 HybridFormula <- <(SingleName (_ (HybridFormulaPart / HybridFormulaFull))+)> */
		func() bool {
			position18, tokenIndex18 := position, tokenIndex
			{
				position19 := position
				if !_rules[ruleSingleName]() {
					goto l18
				}
				if !_rules[rule_]() {
					goto l18
				}
				{
					position22, tokenIndex22 := position, tokenIndex
					if !_rules[ruleHybridFormulaPart]() {
						goto l23
					}
					goto l22
				l23:
					position, tokenIndex = position22, tokenIndex22
					if !_rules[ruleHybridFormulaFull]() {
						goto l18
					}
				}
			l22:
			l20:
				{
					position21, tokenIndex21 := position, tokenIndex
					if !_rules[rule_]() {
						goto l21
					}
					{
						position24, tokenIndex24 := position, tokenIndex
						if !_rules[ruleHybridFormulaPart]() {
							goto l25
						}
						goto l24
					l25:
						position, tokenIndex = position24, tokenIndex24
						if !_rules[ruleHybridFormulaFull]() {
							goto l21
						}
					}
				l24:
					goto l20
				l21:
					position, tokenIndex = position21, tokenIndex21
				}
				add(ruleHybridFormula, position19)
			}
			return true
		l18:
			position, tokenIndex = position18, tokenIndex18
			return false
		},
		/* 4 HybridFormulaFull <- <(HybridChar (_ SingleName)?)> */
		func() bool {
			position26, tokenIndex26 := position, tokenIndex
			{
				position27 := position
				if !_rules[ruleHybridChar]() {
					goto l26
				}
				{
					position28, tokenIndex28 := position, tokenIndex
					if !_rules[rule_]() {
						goto l28
					}
					if !_rules[ruleSingleName]() {
						goto l28
					}
					goto l29
				l28:
					position, tokenIndex = position28, tokenIndex28
				}
			l29:
				add(ruleHybridFormulaFull, position27)
			}
			return true
		l26:
			position, tokenIndex = position26, tokenIndex26
			return false
		},
		/* 5 HybridFormulaPart <- <(HybridChar _ SpeciesEpithet (_ InfraspGroup)?)> */
		func() bool {
			position30, tokenIndex30 := position, tokenIndex
			{
				position31 := position
				if !_rules[ruleHybridChar]() {
					goto l30
				}
				if !_rules[rule_]() {
					goto l30
				}
				if !_rules[ruleSpeciesEpithet]() {
					goto l30
				}
				{
					position32, tokenIndex32 := position, tokenIndex
					if !_rules[rule_]() {
						goto l32
					}
					if !_rules[ruleInfraspGroup]() {
						goto l32
					}
					goto l33
				l32:
					position, tokenIndex = position32, tokenIndex32
				}
			l33:
				add(ruleHybridFormulaPart, position31)
			}
			return true
		l30:
			position, tokenIndex = position30, tokenIndex30
			return false
		},
		/* 6 NamedHybrid <- <(NamedGenusHybrid / NamedSpeciesHybrid)> */
		func() bool {
			position34, tokenIndex34 := position, tokenIndex
			{
				position35 := position
				{
					position36, tokenIndex36 := position, tokenIndex
					if !_rules[ruleNamedGenusHybrid]() {
						goto l37
					}
					goto l36
				l37:
					position, tokenIndex = position36, tokenIndex36
					if !_rules[ruleNamedSpeciesHybrid]() {
						goto l34
					}
				}
			l36:
				add(ruleNamedHybrid, position35)
			}
			return true
		l34:
			position, tokenIndex = position34, tokenIndex34
			return false
		},
		/* 7 NamedSpeciesHybrid <- <(GenusWord _ HybridChar _? SpeciesEpithet)> */
		func() bool {
			position38, tokenIndex38 := position, tokenIndex
			{
				position39 := position
				if !_rules[ruleGenusWord]() {
					goto l38
				}
				if !_rules[rule_]() {
					goto l38
				}
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
				if !_rules[ruleSpeciesEpithet]() {
					goto l38
				}
				add(ruleNamedSpeciesHybrid, position39)
			}
			return true
		l38:
			position, tokenIndex = position38, tokenIndex38
			return false
		},
		/* 8 NamedGenusHybrid <- <(HybridChar _? SingleName)> */
		func() bool {
			position42, tokenIndex42 := position, tokenIndex
			{
				position43 := position
				if !_rules[ruleHybridChar]() {
					goto l42
				}
				{
					position44, tokenIndex44 := position, tokenIndex
					if !_rules[rule_]() {
						goto l44
					}
					goto l45
				l44:
					position, tokenIndex = position44, tokenIndex44
				}
			l45:
				if !_rules[ruleSingleName]() {
					goto l42
				}
				add(ruleNamedGenusHybrid, position43)
			}
			return true
		l42:
			position, tokenIndex = position42, tokenIndex42
			return false
		},
		/* 9 SingleName <- <(NameComp / NameApprox / NameSpecies / NameUninomial)> */
		func() bool {
			position46, tokenIndex46 := position, tokenIndex
			{
				position47 := position
				{
					position48, tokenIndex48 := position, tokenIndex
					if !_rules[ruleNameComp]() {
						goto l49
					}
					goto l48
				l49:
					position, tokenIndex = position48, tokenIndex48
					if !_rules[ruleNameApprox]() {
						goto l50
					}
					goto l48
				l50:
					position, tokenIndex = position48, tokenIndex48
					if !_rules[ruleNameSpecies]() {
						goto l51
					}
					goto l48
				l51:
					position, tokenIndex = position48, tokenIndex48
					if !_rules[ruleNameUninomial]() {
						goto l46
					}
				}
			l48:
				add(ruleSingleName, position47)
			}
			return true
		l46:
			position, tokenIndex = position46, tokenIndex46
			return false
		},
		/* 10 NameUninomial <- <(UninomialCombo / Uninomial)> */
		func() bool {
			position52, tokenIndex52 := position, tokenIndex
			{
				position53 := position
				{
					position54, tokenIndex54 := position, tokenIndex
					if !_rules[ruleUninomialCombo]() {
						goto l55
					}
					goto l54
				l55:
					position, tokenIndex = position54, tokenIndex54
					if !_rules[ruleUninomial]() {
						goto l52
					}
				}
			l54:
				add(ruleNameUninomial, position53)
			}
			return true
		l52:
			position, tokenIndex = position52, tokenIndex52
			return false
		},
		/* 11 NameApprox <- <(GenusWord (_ SpeciesEpithet)? _ Approximation ApproxNameIgnored)> */
		func() bool {
			position56, tokenIndex56 := position, tokenIndex
			{
				position57 := position
				if !_rules[ruleGenusWord]() {
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
				if !_rules[rule_]() {
					goto l56
				}
				if !_rules[ruleApproximation]() {
					goto l56
				}
				if !_rules[ruleApproxNameIgnored]() {
					goto l56
				}
				add(ruleNameApprox, position57)
			}
			return true
		l56:
			position, tokenIndex = position56, tokenIndex56
			return false
		},
		/* 12 NameComp <- <(GenusWord _ Comparison (_ SpeciesEpithet)?)> */
		func() bool {
			position60, tokenIndex60 := position, tokenIndex
			{
				position61 := position
				if !_rules[ruleGenusWord]() {
					goto l60
				}
				if !_rules[rule_]() {
					goto l60
				}
				if !_rules[ruleComparison]() {
					goto l60
				}
				{
					position62, tokenIndex62 := position, tokenIndex
					if !_rules[rule_]() {
						goto l62
					}
					if !_rules[ruleSpeciesEpithet]() {
						goto l62
					}
					goto l63
				l62:
					position, tokenIndex = position62, tokenIndex62
				}
			l63:
				add(ruleNameComp, position61)
			}
			return true
		l60:
			position, tokenIndex = position60, tokenIndex60
			return false
		},
		/* 13 NameSpecies <- <(GenusWord (_? (SubGenus / SubGenusOrSuperspecies))? _ SpeciesEpithet (_ InfraspGroup)?)> */
		func() bool {
			position64, tokenIndex64 := position, tokenIndex
			{
				position65 := position
				if !_rules[ruleGenusWord]() {
					goto l64
				}
				{
					position66, tokenIndex66 := position, tokenIndex
					{
						position68, tokenIndex68 := position, tokenIndex
						if !_rules[rule_]() {
							goto l68
						}
						goto l69
					l68:
						position, tokenIndex = position68, tokenIndex68
					}
				l69:
					{
						position70, tokenIndex70 := position, tokenIndex
						if !_rules[ruleSubGenus]() {
							goto l71
						}
						goto l70
					l71:
						position, tokenIndex = position70, tokenIndex70
						if !_rules[ruleSubGenusOrSuperspecies]() {
							goto l66
						}
					}
				l70:
					goto l67
				l66:
					position, tokenIndex = position66, tokenIndex66
				}
			l67:
				if !_rules[rule_]() {
					goto l64
				}
				if !_rules[ruleSpeciesEpithet]() {
					goto l64
				}
				{
					position72, tokenIndex72 := position, tokenIndex
					if !_rules[rule_]() {
						goto l72
					}
					if !_rules[ruleInfraspGroup]() {
						goto l72
					}
					goto l73
				l72:
					position, tokenIndex = position72, tokenIndex72
				}
			l73:
				add(ruleNameSpecies, position65)
			}
			return true
		l64:
			position, tokenIndex = position64, tokenIndex64
			return false
		},
		/* 14 GenusWord <- <((AbbrGenus / UninomialWord) !(_ AuthorWord))> */
		func() bool {
			position74, tokenIndex74 := position, tokenIndex
			{
				position75 := position
				{
					position76, tokenIndex76 := position, tokenIndex
					if !_rules[ruleAbbrGenus]() {
						goto l77
					}
					goto l76
				l77:
					position, tokenIndex = position76, tokenIndex76
					if !_rules[ruleUninomialWord]() {
						goto l74
					}
				}
			l76:
				{
					position78, tokenIndex78 := position, tokenIndex
					if !_rules[rule_]() {
						goto l78
					}
					if !_rules[ruleAuthorWord]() {
						goto l78
					}
					goto l74
				l78:
					position, tokenIndex = position78, tokenIndex78
				}
				add(ruleGenusWord, position75)
			}
			return true
		l74:
			position, tokenIndex = position74, tokenIndex74
			return false
		},
		/* 15 InfraspGroup <- <(InfraspEpithet (_ InfraspEpithet)? (_ InfraspEpithet)?)> */
		func() bool {
			position79, tokenIndex79 := position, tokenIndex
			{
				position80 := position
				if !_rules[ruleInfraspEpithet]() {
					goto l79
				}
				{
					position81, tokenIndex81 := position, tokenIndex
					if !_rules[rule_]() {
						goto l81
					}
					if !_rules[ruleInfraspEpithet]() {
						goto l81
					}
					goto l82
				l81:
					position, tokenIndex = position81, tokenIndex81
				}
			l82:
				{
					position83, tokenIndex83 := position, tokenIndex
					if !_rules[rule_]() {
						goto l83
					}
					if !_rules[ruleInfraspEpithet]() {
						goto l83
					}
					goto l84
				l83:
					position, tokenIndex = position83, tokenIndex83
				}
			l84:
				add(ruleInfraspGroup, position80)
			}
			return true
		l79:
			position, tokenIndex = position79, tokenIndex79
			return false
		},
		/* 16 InfraspEpithet <- <((Rank _?)? !AuthorEx Word (_ Authorship)?)> */
		func() bool {
			position85, tokenIndex85 := position, tokenIndex
			{
				position86 := position
				{
					position87, tokenIndex87 := position, tokenIndex
					if !_rules[ruleRank]() {
						goto l87
					}
					{
						position89, tokenIndex89 := position, tokenIndex
						if !_rules[rule_]() {
							goto l89
						}
						goto l90
					l89:
						position, tokenIndex = position89, tokenIndex89
					}
				l90:
					goto l88
				l87:
					position, tokenIndex = position87, tokenIndex87
				}
			l88:
				{
					position91, tokenIndex91 := position, tokenIndex
					if !_rules[ruleAuthorEx]() {
						goto l91
					}
					goto l85
				l91:
					position, tokenIndex = position91, tokenIndex91
				}
				if !_rules[ruleWord]() {
					goto l85
				}
				{
					position92, tokenIndex92 := position, tokenIndex
					if !_rules[rule_]() {
						goto l92
					}
					if !_rules[ruleAuthorship]() {
						goto l92
					}
					goto l93
				l92:
					position, tokenIndex = position92, tokenIndex92
				}
			l93:
				add(ruleInfraspEpithet, position86)
			}
			return true
		l85:
			position, tokenIndex = position85, tokenIndex85
			return false
		},
		/* 17 SpeciesEpithet <- <(!AuthorEx Word (_? Authorship)? ','? &(SpaceCharEOI / '('))> */
		func() bool {
			position94, tokenIndex94 := position, tokenIndex
			{
				position95 := position
				{
					position96, tokenIndex96 := position, tokenIndex
					if !_rules[ruleAuthorEx]() {
						goto l96
					}
					goto l94
				l96:
					position, tokenIndex = position96, tokenIndex96
				}
				if !_rules[ruleWord]() {
					goto l94
				}
				{
					position97, tokenIndex97 := position, tokenIndex
					{
						position99, tokenIndex99 := position, tokenIndex
						if !_rules[rule_]() {
							goto l99
						}
						goto l100
					l99:
						position, tokenIndex = position99, tokenIndex99
					}
				l100:
					if !_rules[ruleAuthorship]() {
						goto l97
					}
					goto l98
				l97:
					position, tokenIndex = position97, tokenIndex97
				}
			l98:
				{
					position101, tokenIndex101 := position, tokenIndex
					if buffer[position] != rune(',') {
						goto l101
					}
					position++
					goto l102
				l101:
					position, tokenIndex = position101, tokenIndex101
				}
			l102:
				{
					position103, tokenIndex103 := position, tokenIndex
					{
						position104, tokenIndex104 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l105
						}
						goto l104
					l105:
						position, tokenIndex = position104, tokenIndex104
						if buffer[position] != rune('(') {
							goto l94
						}
						position++
					}
				l104:
					position, tokenIndex = position103, tokenIndex103
				}
				add(ruleSpeciesEpithet, position95)
			}
			return true
		l94:
			position, tokenIndex = position94, tokenIndex94
			return false
		},
		/* 18 Comparison <- <('c' 'f' '.'?)> */
		func() bool {
			position106, tokenIndex106 := position, tokenIndex
			{
				position107 := position
				if buffer[position] != rune('c') {
					goto l106
				}
				position++
				if buffer[position] != rune('f') {
					goto l106
				}
				position++
				{
					position108, tokenIndex108 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l108
					}
					position++
					goto l109
				l108:
					position, tokenIndex = position108, tokenIndex108
				}
			l109:
				add(ruleComparison, position107)
			}
			return true
		l106:
			position, tokenIndex = position106, tokenIndex106
			return false
		},
		/* 19 Rank <- <(RankForma / RankVar / RankSsp / RankOther / RankOtherUncommon)> */
		func() bool {
			position110, tokenIndex110 := position, tokenIndex
			{
				position111 := position
				{
					position112, tokenIndex112 := position, tokenIndex
					if !_rules[ruleRankForma]() {
						goto l113
					}
					goto l112
				l113:
					position, tokenIndex = position112, tokenIndex112
					if !_rules[ruleRankVar]() {
						goto l114
					}
					goto l112
				l114:
					position, tokenIndex = position112, tokenIndex112
					if !_rules[ruleRankSsp]() {
						goto l115
					}
					goto l112
				l115:
					position, tokenIndex = position112, tokenIndex112
					if !_rules[ruleRankOther]() {
						goto l116
					}
					goto l112
				l116:
					position, tokenIndex = position112, tokenIndex112
					if !_rules[ruleRankOtherUncommon]() {
						goto l110
					}
				}
			l112:
				add(ruleRank, position111)
			}
			return true
		l110:
			position, tokenIndex = position110, tokenIndex110
			return false
		},
		/* 20 RankOtherUncommon <- <(('*' / ('n' 'a' 't') / ('f' '.' 's' 'p') / ('m' 'u' 't' '.')) &SpaceCharEOI)> */
		func() bool {
			position117, tokenIndex117 := position, tokenIndex
			{
				position118 := position
				{
					position119, tokenIndex119 := position, tokenIndex
					if buffer[position] != rune('*') {
						goto l120
					}
					position++
					goto l119
				l120:
					position, tokenIndex = position119, tokenIndex119
					if buffer[position] != rune('n') {
						goto l121
					}
					position++
					if buffer[position] != rune('a') {
						goto l121
					}
					position++
					if buffer[position] != rune('t') {
						goto l121
					}
					position++
					goto l119
				l121:
					position, tokenIndex = position119, tokenIndex119
					if buffer[position] != rune('f') {
						goto l122
					}
					position++
					if buffer[position] != rune('.') {
						goto l122
					}
					position++
					if buffer[position] != rune('s') {
						goto l122
					}
					position++
					if buffer[position] != rune('p') {
						goto l122
					}
					position++
					goto l119
				l122:
					position, tokenIndex = position119, tokenIndex119
					if buffer[position] != rune('m') {
						goto l117
					}
					position++
					if buffer[position] != rune('u') {
						goto l117
					}
					position++
					if buffer[position] != rune('t') {
						goto l117
					}
					position++
					if buffer[position] != rune('.') {
						goto l117
					}
					position++
				}
			l119:
				{
					position123, tokenIndex123 := position, tokenIndex
					if !_rules[ruleSpaceCharEOI]() {
						goto l117
					}
					position, tokenIndex = position123, tokenIndex123
				}
				add(ruleRankOtherUncommon, position118)
			}
			return true
		l117:
			position, tokenIndex = position117, tokenIndex117
			return false
		},
		/* 21 RankOther <- <((('m' 'o' 'r' 'p' 'h' '.') / ('n' 'o' 't' 'h' 'o' 's' 'u' 'b' 's' 'p' '.') / ('c' 'o' 'n' 'v' 'a' 'r' '.') / ('p' 's' 'e' 'u' 'd' 'o' 'v' 'a' 'r' '.') / ('s' 'e' 'c' 't' '.') / ('s' 'e' 'r' '.') / ('s' 'u' 'b' 'v' 'a' 'r' '.') / ('s' 'u' 'b' 'f' '.') / ('r' 'a' 'c' 'e') / 'α' / ('β' 'β') / 'β' / 'γ' / 'δ' / 'ε' / 'φ' / 'θ' / 'μ' / ('a' '.') / ('b' '.') / ('c' '.') / ('d' '.') / ('e' '.') / ('g' '.') / ('k' '.') / ('p' 'v' '.') / ('p' 'a' 't' 'h' 'o' 'v' 'a' 'r' '.') / ('a' 'b' '.' (_? ('n' '.'))?) / ('s' 't' '.')) &SpaceCharEOI)> */
		func() bool {
			position124, tokenIndex124 := position, tokenIndex
			{
				position125 := position
				{
					position126, tokenIndex126 := position, tokenIndex
					if buffer[position] != rune('m') {
						goto l127
					}
					position++
					if buffer[position] != rune('o') {
						goto l127
					}
					position++
					if buffer[position] != rune('r') {
						goto l127
					}
					position++
					if buffer[position] != rune('p') {
						goto l127
					}
					position++
					if buffer[position] != rune('h') {
						goto l127
					}
					position++
					if buffer[position] != rune('.') {
						goto l127
					}
					position++
					goto l126
				l127:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('n') {
						goto l128
					}
					position++
					if buffer[position] != rune('o') {
						goto l128
					}
					position++
					if buffer[position] != rune('t') {
						goto l128
					}
					position++
					if buffer[position] != rune('h') {
						goto l128
					}
					position++
					if buffer[position] != rune('o') {
						goto l128
					}
					position++
					if buffer[position] != rune('s') {
						goto l128
					}
					position++
					if buffer[position] != rune('u') {
						goto l128
					}
					position++
					if buffer[position] != rune('b') {
						goto l128
					}
					position++
					if buffer[position] != rune('s') {
						goto l128
					}
					position++
					if buffer[position] != rune('p') {
						goto l128
					}
					position++
					if buffer[position] != rune('.') {
						goto l128
					}
					position++
					goto l126
				l128:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('c') {
						goto l129
					}
					position++
					if buffer[position] != rune('o') {
						goto l129
					}
					position++
					if buffer[position] != rune('n') {
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
					goto l126
				l129:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('p') {
						goto l130
					}
					position++
					if buffer[position] != rune('s') {
						goto l130
					}
					position++
					if buffer[position] != rune('e') {
						goto l130
					}
					position++
					if buffer[position] != rune('u') {
						goto l130
					}
					position++
					if buffer[position] != rune('d') {
						goto l130
					}
					position++
					if buffer[position] != rune('o') {
						goto l130
					}
					position++
					if buffer[position] != rune('v') {
						goto l130
					}
					position++
					if buffer[position] != rune('a') {
						goto l130
					}
					position++
					if buffer[position] != rune('r') {
						goto l130
					}
					position++
					if buffer[position] != rune('.') {
						goto l130
					}
					position++
					goto l126
				l130:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('s') {
						goto l131
					}
					position++
					if buffer[position] != rune('e') {
						goto l131
					}
					position++
					if buffer[position] != rune('c') {
						goto l131
					}
					position++
					if buffer[position] != rune('t') {
						goto l131
					}
					position++
					if buffer[position] != rune('.') {
						goto l131
					}
					position++
					goto l126
				l131:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('s') {
						goto l132
					}
					position++
					if buffer[position] != rune('e') {
						goto l132
					}
					position++
					if buffer[position] != rune('r') {
						goto l132
					}
					position++
					if buffer[position] != rune('.') {
						goto l132
					}
					position++
					goto l126
				l132:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('s') {
						goto l133
					}
					position++
					if buffer[position] != rune('u') {
						goto l133
					}
					position++
					if buffer[position] != rune('b') {
						goto l133
					}
					position++
					if buffer[position] != rune('v') {
						goto l133
					}
					position++
					if buffer[position] != rune('a') {
						goto l133
					}
					position++
					if buffer[position] != rune('r') {
						goto l133
					}
					position++
					if buffer[position] != rune('.') {
						goto l133
					}
					position++
					goto l126
				l133:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('s') {
						goto l134
					}
					position++
					if buffer[position] != rune('u') {
						goto l134
					}
					position++
					if buffer[position] != rune('b') {
						goto l134
					}
					position++
					if buffer[position] != rune('f') {
						goto l134
					}
					position++
					if buffer[position] != rune('.') {
						goto l134
					}
					position++
					goto l126
				l134:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('r') {
						goto l135
					}
					position++
					if buffer[position] != rune('a') {
						goto l135
					}
					position++
					if buffer[position] != rune('c') {
						goto l135
					}
					position++
					if buffer[position] != rune('e') {
						goto l135
					}
					position++
					goto l126
				l135:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('α') {
						goto l136
					}
					position++
					goto l126
				l136:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('β') {
						goto l137
					}
					position++
					if buffer[position] != rune('β') {
						goto l137
					}
					position++
					goto l126
				l137:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('β') {
						goto l138
					}
					position++
					goto l126
				l138:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('γ') {
						goto l139
					}
					position++
					goto l126
				l139:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('δ') {
						goto l140
					}
					position++
					goto l126
				l140:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('ε') {
						goto l141
					}
					position++
					goto l126
				l141:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('φ') {
						goto l142
					}
					position++
					goto l126
				l142:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('θ') {
						goto l143
					}
					position++
					goto l126
				l143:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('μ') {
						goto l144
					}
					position++
					goto l126
				l144:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('a') {
						goto l145
					}
					position++
					if buffer[position] != rune('.') {
						goto l145
					}
					position++
					goto l126
				l145:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('b') {
						goto l146
					}
					position++
					if buffer[position] != rune('.') {
						goto l146
					}
					position++
					goto l126
				l146:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('c') {
						goto l147
					}
					position++
					if buffer[position] != rune('.') {
						goto l147
					}
					position++
					goto l126
				l147:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('d') {
						goto l148
					}
					position++
					if buffer[position] != rune('.') {
						goto l148
					}
					position++
					goto l126
				l148:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('e') {
						goto l149
					}
					position++
					if buffer[position] != rune('.') {
						goto l149
					}
					position++
					goto l126
				l149:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('g') {
						goto l150
					}
					position++
					if buffer[position] != rune('.') {
						goto l150
					}
					position++
					goto l126
				l150:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('k') {
						goto l151
					}
					position++
					if buffer[position] != rune('.') {
						goto l151
					}
					position++
					goto l126
				l151:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('p') {
						goto l152
					}
					position++
					if buffer[position] != rune('v') {
						goto l152
					}
					position++
					if buffer[position] != rune('.') {
						goto l152
					}
					position++
					goto l126
				l152:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('p') {
						goto l153
					}
					position++
					if buffer[position] != rune('a') {
						goto l153
					}
					position++
					if buffer[position] != rune('t') {
						goto l153
					}
					position++
					if buffer[position] != rune('h') {
						goto l153
					}
					position++
					if buffer[position] != rune('o') {
						goto l153
					}
					position++
					if buffer[position] != rune('v') {
						goto l153
					}
					position++
					if buffer[position] != rune('a') {
						goto l153
					}
					position++
					if buffer[position] != rune('r') {
						goto l153
					}
					position++
					if buffer[position] != rune('.') {
						goto l153
					}
					position++
					goto l126
				l153:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('a') {
						goto l154
					}
					position++
					if buffer[position] != rune('b') {
						goto l154
					}
					position++
					if buffer[position] != rune('.') {
						goto l154
					}
					position++
					{
						position155, tokenIndex155 := position, tokenIndex
						{
							position157, tokenIndex157 := position, tokenIndex
							if !_rules[rule_]() {
								goto l157
							}
							goto l158
						l157:
							position, tokenIndex = position157, tokenIndex157
						}
					l158:
						if buffer[position] != rune('n') {
							goto l155
						}
						position++
						if buffer[position] != rune('.') {
							goto l155
						}
						position++
						goto l156
					l155:
						position, tokenIndex = position155, tokenIndex155
					}
				l156:
					goto l126
				l154:
					position, tokenIndex = position126, tokenIndex126
					if buffer[position] != rune('s') {
						goto l124
					}
					position++
					if buffer[position] != rune('t') {
						goto l124
					}
					position++
					if buffer[position] != rune('.') {
						goto l124
					}
					position++
				}
			l126:
				{
					position159, tokenIndex159 := position, tokenIndex
					if !_rules[ruleSpaceCharEOI]() {
						goto l124
					}
					position, tokenIndex = position159, tokenIndex159
				}
				add(ruleRankOther, position125)
			}
			return true
		l124:
			position, tokenIndex = position124, tokenIndex124
			return false
		},
		/* 22 RankVar <- <(('v' 'a' 'r' 'i' 'e' 't' 'y') / ('[' 'v' 'a' 'r' '.' ']') / ('n' 'v' 'a' 'r' '.') / ('v' 'a' 'r' (&SpaceCharEOI / '.')))> */
		func() bool {
			position160, tokenIndex160 := position, tokenIndex
			{
				position161 := position
				{
					position162, tokenIndex162 := position, tokenIndex
					if buffer[position] != rune('v') {
						goto l163
					}
					position++
					if buffer[position] != rune('a') {
						goto l163
					}
					position++
					if buffer[position] != rune('r') {
						goto l163
					}
					position++
					if buffer[position] != rune('i') {
						goto l163
					}
					position++
					if buffer[position] != rune('e') {
						goto l163
					}
					position++
					if buffer[position] != rune('t') {
						goto l163
					}
					position++
					if buffer[position] != rune('y') {
						goto l163
					}
					position++
					goto l162
				l163:
					position, tokenIndex = position162, tokenIndex162
					if buffer[position] != rune('[') {
						goto l164
					}
					position++
					if buffer[position] != rune('v') {
						goto l164
					}
					position++
					if buffer[position] != rune('a') {
						goto l164
					}
					position++
					if buffer[position] != rune('r') {
						goto l164
					}
					position++
					if buffer[position] != rune('.') {
						goto l164
					}
					position++
					if buffer[position] != rune(']') {
						goto l164
					}
					position++
					goto l162
				l164:
					position, tokenIndex = position162, tokenIndex162
					if buffer[position] != rune('n') {
						goto l165
					}
					position++
					if buffer[position] != rune('v') {
						goto l165
					}
					position++
					if buffer[position] != rune('a') {
						goto l165
					}
					position++
					if buffer[position] != rune('r') {
						goto l165
					}
					position++
					if buffer[position] != rune('.') {
						goto l165
					}
					position++
					goto l162
				l165:
					position, tokenIndex = position162, tokenIndex162
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
					{
						position166, tokenIndex166 := position, tokenIndex
						{
							position168, tokenIndex168 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l167
							}
							position, tokenIndex = position168, tokenIndex168
						}
						goto l166
					l167:
						position, tokenIndex = position166, tokenIndex166
						if buffer[position] != rune('.') {
							goto l160
						}
						position++
					}
				l166:
				}
			l162:
				add(ruleRankVar, position161)
			}
			return true
		l160:
			position, tokenIndex = position160, tokenIndex160
			return false
		},
		/* 23 RankForma <- <((('f' 'o' 'r' 'm' 'a') / ('f' 'm' 'a') / ('f' 'o' 'r' 'm') / ('f' 'o') / 'f') (&SpaceCharEOI / '.'))> */
		func() bool {
			position169, tokenIndex169 := position, tokenIndex
			{
				position170 := position
				{
					position171, tokenIndex171 := position, tokenIndex
					if buffer[position] != rune('f') {
						goto l172
					}
					position++
					if buffer[position] != rune('o') {
						goto l172
					}
					position++
					if buffer[position] != rune('r') {
						goto l172
					}
					position++
					if buffer[position] != rune('m') {
						goto l172
					}
					position++
					if buffer[position] != rune('a') {
						goto l172
					}
					position++
					goto l171
				l172:
					position, tokenIndex = position171, tokenIndex171
					if buffer[position] != rune('f') {
						goto l173
					}
					position++
					if buffer[position] != rune('m') {
						goto l173
					}
					position++
					if buffer[position] != rune('a') {
						goto l173
					}
					position++
					goto l171
				l173:
					position, tokenIndex = position171, tokenIndex171
					if buffer[position] != rune('f') {
						goto l174
					}
					position++
					if buffer[position] != rune('o') {
						goto l174
					}
					position++
					if buffer[position] != rune('r') {
						goto l174
					}
					position++
					if buffer[position] != rune('m') {
						goto l174
					}
					position++
					goto l171
				l174:
					position, tokenIndex = position171, tokenIndex171
					if buffer[position] != rune('f') {
						goto l175
					}
					position++
					if buffer[position] != rune('o') {
						goto l175
					}
					position++
					goto l171
				l175:
					position, tokenIndex = position171, tokenIndex171
					if buffer[position] != rune('f') {
						goto l169
					}
					position++
				}
			l171:
				{
					position176, tokenIndex176 := position, tokenIndex
					{
						position178, tokenIndex178 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l177
						}
						position, tokenIndex = position178, tokenIndex178
					}
					goto l176
				l177:
					position, tokenIndex = position176, tokenIndex176
					if buffer[position] != rune('.') {
						goto l169
					}
					position++
				}
			l176:
				add(ruleRankForma, position170)
			}
			return true
		l169:
			position, tokenIndex = position169, tokenIndex169
			return false
		},
		/* 24 RankSsp <- <((('s' 's' 'p') / ('s' 'u' 'b' 's' 'p')) (&SpaceCharEOI / '.'))> */
		func() bool {
			position179, tokenIndex179 := position, tokenIndex
			{
				position180 := position
				{
					position181, tokenIndex181 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l182
					}
					position++
					if buffer[position] != rune('s') {
						goto l182
					}
					position++
					if buffer[position] != rune('p') {
						goto l182
					}
					position++
					goto l181
				l182:
					position, tokenIndex = position181, tokenIndex181
					if buffer[position] != rune('s') {
						goto l179
					}
					position++
					if buffer[position] != rune('u') {
						goto l179
					}
					position++
					if buffer[position] != rune('b') {
						goto l179
					}
					position++
					if buffer[position] != rune('s') {
						goto l179
					}
					position++
					if buffer[position] != rune('p') {
						goto l179
					}
					position++
				}
			l181:
				{
					position183, tokenIndex183 := position, tokenIndex
					{
						position185, tokenIndex185 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l184
						}
						position, tokenIndex = position185, tokenIndex185
					}
					goto l183
				l184:
					position, tokenIndex = position183, tokenIndex183
					if buffer[position] != rune('.') {
						goto l179
					}
					position++
				}
			l183:
				add(ruleRankSsp, position180)
			}
			return true
		l179:
			position, tokenIndex = position179, tokenIndex179
			return false
		},
		/* 25 SubGenusOrSuperspecies <- <('(' _? NameLowerChar+ _? ')')> */
		func() bool {
			position186, tokenIndex186 := position, tokenIndex
			{
				position187 := position
				if buffer[position] != rune('(') {
					goto l186
				}
				position++
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
				if !_rules[ruleNameLowerChar]() {
					goto l186
				}
			l190:
				{
					position191, tokenIndex191 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l191
					}
					goto l190
				l191:
					position, tokenIndex = position191, tokenIndex191
				}
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
				if buffer[position] != rune(')') {
					goto l186
				}
				position++
				add(ruleSubGenusOrSuperspecies, position187)
			}
			return true
		l186:
			position, tokenIndex = position186, tokenIndex186
			return false
		},
		/* 26 SubGenus <- <('(' _? UninomialWord _? ')')> */
		func() bool {
			position194, tokenIndex194 := position, tokenIndex
			{
				position195 := position
				if buffer[position] != rune('(') {
					goto l194
				}
				position++
				{
					position196, tokenIndex196 := position, tokenIndex
					if !_rules[rule_]() {
						goto l196
					}
					goto l197
				l196:
					position, tokenIndex = position196, tokenIndex196
				}
			l197:
				if !_rules[ruleUninomialWord]() {
					goto l194
				}
				{
					position198, tokenIndex198 := position, tokenIndex
					if !_rules[rule_]() {
						goto l198
					}
					goto l199
				l198:
					position, tokenIndex = position198, tokenIndex198
				}
			l199:
				if buffer[position] != rune(')') {
					goto l194
				}
				position++
				add(ruleSubGenus, position195)
			}
			return true
		l194:
			position, tokenIndex = position194, tokenIndex194
			return false
		},
		/* 27 UninomialCombo <- <(UninomialCombo1 / UninomialCombo2)> */
		func() bool {
			position200, tokenIndex200 := position, tokenIndex
			{
				position201 := position
				{
					position202, tokenIndex202 := position, tokenIndex
					if !_rules[ruleUninomialCombo1]() {
						goto l203
					}
					goto l202
				l203:
					position, tokenIndex = position202, tokenIndex202
					if !_rules[ruleUninomialCombo2]() {
						goto l200
					}
				}
			l202:
				add(ruleUninomialCombo, position201)
			}
			return true
		l200:
			position, tokenIndex = position200, tokenIndex200
			return false
		},
		/* 28 UninomialCombo1 <- <(UninomialWord _? SubGenus _? Authorship .?)> */
		func() bool {
			position204, tokenIndex204 := position, tokenIndex
			{
				position205 := position
				if !_rules[ruleUninomialWord]() {
					goto l204
				}
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
				if !_rules[ruleSubGenus]() {
					goto l204
				}
				{
					position208, tokenIndex208 := position, tokenIndex
					if !_rules[rule_]() {
						goto l208
					}
					goto l209
				l208:
					position, tokenIndex = position208, tokenIndex208
				}
			l209:
				if !_rules[ruleAuthorship]() {
					goto l204
				}
				{
					position210, tokenIndex210 := position, tokenIndex
					if !matchDot() {
						goto l210
					}
					goto l211
				l210:
					position, tokenIndex = position210, tokenIndex210
				}
			l211:
				add(ruleUninomialCombo1, position205)
			}
			return true
		l204:
			position, tokenIndex = position204, tokenIndex204
			return false
		},
		/* 29 UninomialCombo2 <- <(Uninomial _? RankUninomial _? Uninomial)> */
		func() bool {
			position212, tokenIndex212 := position, tokenIndex
			{
				position213 := position
				if !_rules[ruleUninomial]() {
					goto l212
				}
				{
					position214, tokenIndex214 := position, tokenIndex
					if !_rules[rule_]() {
						goto l214
					}
					goto l215
				l214:
					position, tokenIndex = position214, tokenIndex214
				}
			l215:
				if !_rules[ruleRankUninomial]() {
					goto l212
				}
				{
					position216, tokenIndex216 := position, tokenIndex
					if !_rules[rule_]() {
						goto l216
					}
					goto l217
				l216:
					position, tokenIndex = position216, tokenIndex216
				}
			l217:
				if !_rules[ruleUninomial]() {
					goto l212
				}
				add(ruleUninomialCombo2, position213)
			}
			return true
		l212:
			position, tokenIndex = position212, tokenIndex212
			return false
		},
		/* 30 RankUninomial <- <((('s' 'e' 'c' 't') / ('s' 'u' 'b' 's' 'e' 'c' 't') / ('t' 'r' 'i' 'b') / ('s' 'u' 'b' 't' 'r' 'i' 'b') / ('s' 'u' 'b' 's' 'e' 'r') / ('s' 'e' 'r') / ('s' 'u' 'b' 'g' 'e' 'n') / ('f' 'a' 'm') / ('s' 'u' 'b' 'f' 'a' 'm') / ('s' 'u' 'p' 'e' 'r' 't' 'r' 'i' 'b')) '.'?)> */
		func() bool {
			position218, tokenIndex218 := position, tokenIndex
			{
				position219 := position
				{
					position220, tokenIndex220 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l221
					}
					position++
					if buffer[position] != rune('e') {
						goto l221
					}
					position++
					if buffer[position] != rune('c') {
						goto l221
					}
					position++
					if buffer[position] != rune('t') {
						goto l221
					}
					position++
					goto l220
				l221:
					position, tokenIndex = position220, tokenIndex220
					if buffer[position] != rune('s') {
						goto l222
					}
					position++
					if buffer[position] != rune('u') {
						goto l222
					}
					position++
					if buffer[position] != rune('b') {
						goto l222
					}
					position++
					if buffer[position] != rune('s') {
						goto l222
					}
					position++
					if buffer[position] != rune('e') {
						goto l222
					}
					position++
					if buffer[position] != rune('c') {
						goto l222
					}
					position++
					if buffer[position] != rune('t') {
						goto l222
					}
					position++
					goto l220
				l222:
					position, tokenIndex = position220, tokenIndex220
					if buffer[position] != rune('t') {
						goto l223
					}
					position++
					if buffer[position] != rune('r') {
						goto l223
					}
					position++
					if buffer[position] != rune('i') {
						goto l223
					}
					position++
					if buffer[position] != rune('b') {
						goto l223
					}
					position++
					goto l220
				l223:
					position, tokenIndex = position220, tokenIndex220
					if buffer[position] != rune('s') {
						goto l224
					}
					position++
					if buffer[position] != rune('u') {
						goto l224
					}
					position++
					if buffer[position] != rune('b') {
						goto l224
					}
					position++
					if buffer[position] != rune('t') {
						goto l224
					}
					position++
					if buffer[position] != rune('r') {
						goto l224
					}
					position++
					if buffer[position] != rune('i') {
						goto l224
					}
					position++
					if buffer[position] != rune('b') {
						goto l224
					}
					position++
					goto l220
				l224:
					position, tokenIndex = position220, tokenIndex220
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
					if buffer[position] != rune('s') {
						goto l225
					}
					position++
					if buffer[position] != rune('e') {
						goto l225
					}
					position++
					if buffer[position] != rune('r') {
						goto l225
					}
					position++
					goto l220
				l225:
					position, tokenIndex = position220, tokenIndex220
					if buffer[position] != rune('s') {
						goto l226
					}
					position++
					if buffer[position] != rune('e') {
						goto l226
					}
					position++
					if buffer[position] != rune('r') {
						goto l226
					}
					position++
					goto l220
				l226:
					position, tokenIndex = position220, tokenIndex220
					if buffer[position] != rune('s') {
						goto l227
					}
					position++
					if buffer[position] != rune('u') {
						goto l227
					}
					position++
					if buffer[position] != rune('b') {
						goto l227
					}
					position++
					if buffer[position] != rune('g') {
						goto l227
					}
					position++
					if buffer[position] != rune('e') {
						goto l227
					}
					position++
					if buffer[position] != rune('n') {
						goto l227
					}
					position++
					goto l220
				l227:
					position, tokenIndex = position220, tokenIndex220
					if buffer[position] != rune('f') {
						goto l228
					}
					position++
					if buffer[position] != rune('a') {
						goto l228
					}
					position++
					if buffer[position] != rune('m') {
						goto l228
					}
					position++
					goto l220
				l228:
					position, tokenIndex = position220, tokenIndex220
					if buffer[position] != rune('s') {
						goto l229
					}
					position++
					if buffer[position] != rune('u') {
						goto l229
					}
					position++
					if buffer[position] != rune('b') {
						goto l229
					}
					position++
					if buffer[position] != rune('f') {
						goto l229
					}
					position++
					if buffer[position] != rune('a') {
						goto l229
					}
					position++
					if buffer[position] != rune('m') {
						goto l229
					}
					position++
					goto l220
				l229:
					position, tokenIndex = position220, tokenIndex220
					if buffer[position] != rune('s') {
						goto l218
					}
					position++
					if buffer[position] != rune('u') {
						goto l218
					}
					position++
					if buffer[position] != rune('p') {
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
					if buffer[position] != rune('t') {
						goto l218
					}
					position++
					if buffer[position] != rune('r') {
						goto l218
					}
					position++
					if buffer[position] != rune('i') {
						goto l218
					}
					position++
					if buffer[position] != rune('b') {
						goto l218
					}
					position++
				}
			l220:
				{
					position230, tokenIndex230 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l230
					}
					position++
					goto l231
				l230:
					position, tokenIndex = position230, tokenIndex230
				}
			l231:
				add(ruleRankUninomial, position219)
			}
			return true
		l218:
			position, tokenIndex = position218, tokenIndex218
			return false
		},
		/* 31 Uninomial <- <(UninomialWord (_ Authorship)?)> */
		func() bool {
			position232, tokenIndex232 := position, tokenIndex
			{
				position233 := position
				if !_rules[ruleUninomialWord]() {
					goto l232
				}
				{
					position234, tokenIndex234 := position, tokenIndex
					if !_rules[rule_]() {
						goto l234
					}
					if !_rules[ruleAuthorship]() {
						goto l234
					}
					goto l235
				l234:
					position, tokenIndex = position234, tokenIndex234
				}
			l235:
				add(ruleUninomial, position233)
			}
			return true
		l232:
			position, tokenIndex = position232, tokenIndex232
			return false
		},
		/* 32 UninomialWord <- <(CapWord / TwoLetterGenus)> */
		func() bool {
			position236, tokenIndex236 := position, tokenIndex
			{
				position237 := position
				{
					position238, tokenIndex238 := position, tokenIndex
					if !_rules[ruleCapWord]() {
						goto l239
					}
					goto l238
				l239:
					position, tokenIndex = position238, tokenIndex238
					if !_rules[ruleTwoLetterGenus]() {
						goto l236
					}
				}
			l238:
				add(ruleUninomialWord, position237)
			}
			return true
		l236:
			position, tokenIndex = position236, tokenIndex236
			return false
		},
		/* 33 AbbrGenus <- <(UpperChar LowerChar* '.')> */
		func() bool {
			position240, tokenIndex240 := position, tokenIndex
			{
				position241 := position
				if !_rules[ruleUpperChar]() {
					goto l240
				}
			l242:
				{
					position243, tokenIndex243 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l243
					}
					goto l242
				l243:
					position, tokenIndex = position243, tokenIndex243
				}
				if buffer[position] != rune('.') {
					goto l240
				}
				position++
				add(ruleAbbrGenus, position241)
			}
			return true
		l240:
			position, tokenIndex = position240, tokenIndex240
			return false
		},
		/* 34 CapWord <- <(CapWord2 / CapWord1)> */
		func() bool {
			position244, tokenIndex244 := position, tokenIndex
			{
				position245 := position
				{
					position246, tokenIndex246 := position, tokenIndex
					if !_rules[ruleCapWord2]() {
						goto l247
					}
					goto l246
				l247:
					position, tokenIndex = position246, tokenIndex246
					if !_rules[ruleCapWord1]() {
						goto l244
					}
				}
			l246:
				add(ruleCapWord, position245)
			}
			return true
		l244:
			position, tokenIndex = position244, tokenIndex244
			return false
		},
		/* 35 CapWord1 <- <(NameUpperChar NameLowerChar NameLowerChar+ '?'?)> */
		func() bool {
			position248, tokenIndex248 := position, tokenIndex
			{
				position249 := position
				if !_rules[ruleNameUpperChar]() {
					goto l248
				}
				if !_rules[ruleNameLowerChar]() {
					goto l248
				}
				if !_rules[ruleNameLowerChar]() {
					goto l248
				}
			l250:
				{
					position251, tokenIndex251 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l251
					}
					goto l250
				l251:
					position, tokenIndex = position251, tokenIndex251
				}
				{
					position252, tokenIndex252 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l252
					}
					position++
					goto l253
				l252:
					position, tokenIndex = position252, tokenIndex252
				}
			l253:
				add(ruleCapWord1, position249)
			}
			return true
		l248:
			position, tokenIndex = position248, tokenIndex248
			return false
		},
		/* 36 CapWord2 <- <(CapWord1 dash (CapWord1 / Word1))> */
		func() bool {
			position254, tokenIndex254 := position, tokenIndex
			{
				position255 := position
				if !_rules[ruleCapWord1]() {
					goto l254
				}
				if !_rules[ruledash]() {
					goto l254
				}
				{
					position256, tokenIndex256 := position, tokenIndex
					if !_rules[ruleCapWord1]() {
						goto l257
					}
					goto l256
				l257:
					position, tokenIndex = position256, tokenIndex256
					if !_rules[ruleWord1]() {
						goto l254
					}
				}
			l256:
				add(ruleCapWord2, position255)
			}
			return true
		l254:
			position, tokenIndex = position254, tokenIndex254
			return false
		},
		/* 37 TwoLetterGenus <- <(('C' 'a') / ('E' 'a') / ('G' 'e') / ('I' 'a') / ('I' 'o') / ('I' 'x') / ('L' 'o') / ('O' 'a') / ('R' 'a') / ('T' 'y') / ('U' 'a') / ('A' 'a') / ('J' 'a') / ('Z' 'u') / ('L' 'a') / ('Q' 'u') / ('A' 's') / ('B' 'a'))> */
		func() bool {
			position258, tokenIndex258 := position, tokenIndex
			{
				position259 := position
				{
					position260, tokenIndex260 := position, tokenIndex
					if buffer[position] != rune('C') {
						goto l261
					}
					position++
					if buffer[position] != rune('a') {
						goto l261
					}
					position++
					goto l260
				l261:
					position, tokenIndex = position260, tokenIndex260
					if buffer[position] != rune('E') {
						goto l262
					}
					position++
					if buffer[position] != rune('a') {
						goto l262
					}
					position++
					goto l260
				l262:
					position, tokenIndex = position260, tokenIndex260
					if buffer[position] != rune('G') {
						goto l263
					}
					position++
					if buffer[position] != rune('e') {
						goto l263
					}
					position++
					goto l260
				l263:
					position, tokenIndex = position260, tokenIndex260
					if buffer[position] != rune('I') {
						goto l264
					}
					position++
					if buffer[position] != rune('a') {
						goto l264
					}
					position++
					goto l260
				l264:
					position, tokenIndex = position260, tokenIndex260
					if buffer[position] != rune('I') {
						goto l265
					}
					position++
					if buffer[position] != rune('o') {
						goto l265
					}
					position++
					goto l260
				l265:
					position, tokenIndex = position260, tokenIndex260
					if buffer[position] != rune('I') {
						goto l266
					}
					position++
					if buffer[position] != rune('x') {
						goto l266
					}
					position++
					goto l260
				l266:
					position, tokenIndex = position260, tokenIndex260
					if buffer[position] != rune('L') {
						goto l267
					}
					position++
					if buffer[position] != rune('o') {
						goto l267
					}
					position++
					goto l260
				l267:
					position, tokenIndex = position260, tokenIndex260
					if buffer[position] != rune('O') {
						goto l268
					}
					position++
					if buffer[position] != rune('a') {
						goto l268
					}
					position++
					goto l260
				l268:
					position, tokenIndex = position260, tokenIndex260
					if buffer[position] != rune('R') {
						goto l269
					}
					position++
					if buffer[position] != rune('a') {
						goto l269
					}
					position++
					goto l260
				l269:
					position, tokenIndex = position260, tokenIndex260
					if buffer[position] != rune('T') {
						goto l270
					}
					position++
					if buffer[position] != rune('y') {
						goto l270
					}
					position++
					goto l260
				l270:
					position, tokenIndex = position260, tokenIndex260
					if buffer[position] != rune('U') {
						goto l271
					}
					position++
					if buffer[position] != rune('a') {
						goto l271
					}
					position++
					goto l260
				l271:
					position, tokenIndex = position260, tokenIndex260
					if buffer[position] != rune('A') {
						goto l272
					}
					position++
					if buffer[position] != rune('a') {
						goto l272
					}
					position++
					goto l260
				l272:
					position, tokenIndex = position260, tokenIndex260
					if buffer[position] != rune('J') {
						goto l273
					}
					position++
					if buffer[position] != rune('a') {
						goto l273
					}
					position++
					goto l260
				l273:
					position, tokenIndex = position260, tokenIndex260
					if buffer[position] != rune('Z') {
						goto l274
					}
					position++
					if buffer[position] != rune('u') {
						goto l274
					}
					position++
					goto l260
				l274:
					position, tokenIndex = position260, tokenIndex260
					if buffer[position] != rune('L') {
						goto l275
					}
					position++
					if buffer[position] != rune('a') {
						goto l275
					}
					position++
					goto l260
				l275:
					position, tokenIndex = position260, tokenIndex260
					if buffer[position] != rune('Q') {
						goto l276
					}
					position++
					if buffer[position] != rune('u') {
						goto l276
					}
					position++
					goto l260
				l276:
					position, tokenIndex = position260, tokenIndex260
					if buffer[position] != rune('A') {
						goto l277
					}
					position++
					if buffer[position] != rune('s') {
						goto l277
					}
					position++
					goto l260
				l277:
					position, tokenIndex = position260, tokenIndex260
					if buffer[position] != rune('B') {
						goto l258
					}
					position++
					if buffer[position] != rune('a') {
						goto l258
					}
					position++
				}
			l260:
				add(ruleTwoLetterGenus, position259)
			}
			return true
		l258:
			position, tokenIndex = position258, tokenIndex258
			return false
		},
		/* 38 Word <- <(!(AuthorPrefix / RankUninomial / Approximation / Word4) (WordApostr / WordStartsWithDigit / Word2 / Word1) &(SpaceCharEOI / '('))> */
		func() bool {
			position278, tokenIndex278 := position, tokenIndex
			{
				position279 := position
				{
					position280, tokenIndex280 := position, tokenIndex
					{
						position281, tokenIndex281 := position, tokenIndex
						if !_rules[ruleAuthorPrefix]() {
							goto l282
						}
						goto l281
					l282:
						position, tokenIndex = position281, tokenIndex281
						if !_rules[ruleRankUninomial]() {
							goto l283
						}
						goto l281
					l283:
						position, tokenIndex = position281, tokenIndex281
						if !_rules[ruleApproximation]() {
							goto l284
						}
						goto l281
					l284:
						position, tokenIndex = position281, tokenIndex281
						if !_rules[ruleWord4]() {
							goto l280
						}
					}
				l281:
					goto l278
				l280:
					position, tokenIndex = position280, tokenIndex280
				}
				{
					position285, tokenIndex285 := position, tokenIndex
					if !_rules[ruleWordApostr]() {
						goto l286
					}
					goto l285
				l286:
					position, tokenIndex = position285, tokenIndex285
					if !_rules[ruleWordStartsWithDigit]() {
						goto l287
					}
					goto l285
				l287:
					position, tokenIndex = position285, tokenIndex285
					if !_rules[ruleWord2]() {
						goto l288
					}
					goto l285
				l288:
					position, tokenIndex = position285, tokenIndex285
					if !_rules[ruleWord1]() {
						goto l278
					}
				}
			l285:
				{
					position289, tokenIndex289 := position, tokenIndex
					{
						position290, tokenIndex290 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l291
						}
						goto l290
					l291:
						position, tokenIndex = position290, tokenIndex290
						if buffer[position] != rune('(') {
							goto l278
						}
						position++
					}
				l290:
					position, tokenIndex = position289, tokenIndex289
				}
				add(ruleWord, position279)
			}
			return true
		l278:
			position, tokenIndex = position278, tokenIndex278
			return false
		},
		/* 39 Word1 <- <((lASCII dash)? NameLowerChar NameLowerChar+)> */
		func() bool {
			position292, tokenIndex292 := position, tokenIndex
			{
				position293 := position
				{
					position294, tokenIndex294 := position, tokenIndex
					if !_rules[rulelASCII]() {
						goto l294
					}
					if !_rules[ruledash]() {
						goto l294
					}
					goto l295
				l294:
					position, tokenIndex = position294, tokenIndex294
				}
			l295:
				if !_rules[ruleNameLowerChar]() {
					goto l292
				}
				if !_rules[ruleNameLowerChar]() {
					goto l292
				}
			l296:
				{
					position297, tokenIndex297 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l297
					}
					goto l296
				l297:
					position, tokenIndex = position297, tokenIndex297
				}
				add(ruleWord1, position293)
			}
			return true
		l292:
			position, tokenIndex = position292, tokenIndex292
			return false
		},
		/* 40 WordStartsWithDigit <- <(('1' / '2' / '3' / '4' / '5' / '6' / '7' / '8' / '9') nums? ('.' / dash)? NameLowerChar NameLowerChar NameLowerChar NameLowerChar+)> */
		func() bool {
			position298, tokenIndex298 := position, tokenIndex
			{
				position299 := position
				{
					position300, tokenIndex300 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l301
					}
					position++
					goto l300
				l301:
					position, tokenIndex = position300, tokenIndex300
					if buffer[position] != rune('2') {
						goto l302
					}
					position++
					goto l300
				l302:
					position, tokenIndex = position300, tokenIndex300
					if buffer[position] != rune('3') {
						goto l303
					}
					position++
					goto l300
				l303:
					position, tokenIndex = position300, tokenIndex300
					if buffer[position] != rune('4') {
						goto l304
					}
					position++
					goto l300
				l304:
					position, tokenIndex = position300, tokenIndex300
					if buffer[position] != rune('5') {
						goto l305
					}
					position++
					goto l300
				l305:
					position, tokenIndex = position300, tokenIndex300
					if buffer[position] != rune('6') {
						goto l306
					}
					position++
					goto l300
				l306:
					position, tokenIndex = position300, tokenIndex300
					if buffer[position] != rune('7') {
						goto l307
					}
					position++
					goto l300
				l307:
					position, tokenIndex = position300, tokenIndex300
					if buffer[position] != rune('8') {
						goto l308
					}
					position++
					goto l300
				l308:
					position, tokenIndex = position300, tokenIndex300
					if buffer[position] != rune('9') {
						goto l298
					}
					position++
				}
			l300:
				{
					position309, tokenIndex309 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l309
					}
					goto l310
				l309:
					position, tokenIndex = position309, tokenIndex309
				}
			l310:
				{
					position311, tokenIndex311 := position, tokenIndex
					{
						position313, tokenIndex313 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l314
						}
						position++
						goto l313
					l314:
						position, tokenIndex = position313, tokenIndex313
						if !_rules[ruledash]() {
							goto l311
						}
					}
				l313:
					goto l312
				l311:
					position, tokenIndex = position311, tokenIndex311
				}
			l312:
				if !_rules[ruleNameLowerChar]() {
					goto l298
				}
				if !_rules[ruleNameLowerChar]() {
					goto l298
				}
				if !_rules[ruleNameLowerChar]() {
					goto l298
				}
				if !_rules[ruleNameLowerChar]() {
					goto l298
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
				add(ruleWordStartsWithDigit, position299)
			}
			return true
		l298:
			position, tokenIndex = position298, tokenIndex298
			return false
		},
		/* 41 Word2 <- <(NameLowerChar+ dash? NameLowerChar+)> */
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
				{
					position321, tokenIndex321 := position, tokenIndex
					if !_rules[ruledash]() {
						goto l321
					}
					goto l322
				l321:
					position, tokenIndex = position321, tokenIndex321
				}
			l322:
				if !_rules[ruleNameLowerChar]() {
					goto l317
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
				add(ruleWord2, position318)
			}
			return true
		l317:
			position, tokenIndex = position317, tokenIndex317
			return false
		},
		/* 42 WordApostr <- <(NameLowerChar NameLowerChar* apostr Word1)> */
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
				if !_rules[ruleapostr]() {
					goto l325
				}
				if !_rules[ruleWord1]() {
					goto l325
				}
				add(ruleWordApostr, position326)
			}
			return true
		l325:
			position, tokenIndex = position325, tokenIndex325
			return false
		},
		/* 43 Word4 <- <(NameLowerChar+ '.' NameLowerChar)> */
		func() bool {
			position329, tokenIndex329 := position, tokenIndex
			{
				position330 := position
				if !_rules[ruleNameLowerChar]() {
					goto l329
				}
			l331:
				{
					position332, tokenIndex332 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l332
					}
					goto l331
				l332:
					position, tokenIndex = position332, tokenIndex332
				}
				if buffer[position] != rune('.') {
					goto l329
				}
				position++
				if !_rules[ruleNameLowerChar]() {
					goto l329
				}
				add(ruleWord4, position330)
			}
			return true
		l329:
			position, tokenIndex = position329, tokenIndex329
			return false
		},
		/* 44 HybridChar <- <'×'> */
		func() bool {
			position333, tokenIndex333 := position, tokenIndex
			{
				position334 := position
				if buffer[position] != rune('×') {
					goto l333
				}
				position++
				add(ruleHybridChar, position334)
			}
			return true
		l333:
			position, tokenIndex = position333, tokenIndex333
			return false
		},
		/* 45 ApproxNameIgnored <- <.*> */
		func() bool {
			{
				position336 := position
			l337:
				{
					position338, tokenIndex338 := position, tokenIndex
					if !matchDot() {
						goto l338
					}
					goto l337
				l338:
					position, tokenIndex = position338, tokenIndex338
				}
				add(ruleApproxNameIgnored, position336)
			}
			return true
		},
		/* 46 Approximation <- <(('s' 'p' '.' _? ('n' 'r' '.')) / ('s' 'p' '.' _? ('a' 'f' 'f' '.')) / ('m' 'o' 'n' 's' 't' '.') / '?' / ((('s' 'p' 'p') / ('n' 'r') / ('s' 'p') / ('a' 'f' 'f') / ('s' 'p' 'e' 'c' 'i' 'e' 's')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position339, tokenIndex339 := position, tokenIndex
			{
				position340 := position
				{
					position341, tokenIndex341 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l342
					}
					position++
					if buffer[position] != rune('p') {
						goto l342
					}
					position++
					if buffer[position] != rune('.') {
						goto l342
					}
					position++
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
					if buffer[position] != rune('n') {
						goto l342
					}
					position++
					if buffer[position] != rune('r') {
						goto l342
					}
					position++
					if buffer[position] != rune('.') {
						goto l342
					}
					position++
					goto l341
				l342:
					position, tokenIndex = position341, tokenIndex341
					if buffer[position] != rune('s') {
						goto l345
					}
					position++
					if buffer[position] != rune('p') {
						goto l345
					}
					position++
					if buffer[position] != rune('.') {
						goto l345
					}
					position++
					{
						position346, tokenIndex346 := position, tokenIndex
						if !_rules[rule_]() {
							goto l346
						}
						goto l347
					l346:
						position, tokenIndex = position346, tokenIndex346
					}
				l347:
					if buffer[position] != rune('a') {
						goto l345
					}
					position++
					if buffer[position] != rune('f') {
						goto l345
					}
					position++
					if buffer[position] != rune('f') {
						goto l345
					}
					position++
					if buffer[position] != rune('.') {
						goto l345
					}
					position++
					goto l341
				l345:
					position, tokenIndex = position341, tokenIndex341
					if buffer[position] != rune('m') {
						goto l348
					}
					position++
					if buffer[position] != rune('o') {
						goto l348
					}
					position++
					if buffer[position] != rune('n') {
						goto l348
					}
					position++
					if buffer[position] != rune('s') {
						goto l348
					}
					position++
					if buffer[position] != rune('t') {
						goto l348
					}
					position++
					if buffer[position] != rune('.') {
						goto l348
					}
					position++
					goto l341
				l348:
					position, tokenIndex = position341, tokenIndex341
					if buffer[position] != rune('?') {
						goto l349
					}
					position++
					goto l341
				l349:
					position, tokenIndex = position341, tokenIndex341
					{
						position350, tokenIndex350 := position, tokenIndex
						if buffer[position] != rune('s') {
							goto l351
						}
						position++
						if buffer[position] != rune('p') {
							goto l351
						}
						position++
						if buffer[position] != rune('p') {
							goto l351
						}
						position++
						goto l350
					l351:
						position, tokenIndex = position350, tokenIndex350
						if buffer[position] != rune('n') {
							goto l352
						}
						position++
						if buffer[position] != rune('r') {
							goto l352
						}
						position++
						goto l350
					l352:
						position, tokenIndex = position350, tokenIndex350
						if buffer[position] != rune('s') {
							goto l353
						}
						position++
						if buffer[position] != rune('p') {
							goto l353
						}
						position++
						goto l350
					l353:
						position, tokenIndex = position350, tokenIndex350
						if buffer[position] != rune('a') {
							goto l354
						}
						position++
						if buffer[position] != rune('f') {
							goto l354
						}
						position++
						if buffer[position] != rune('f') {
							goto l354
						}
						position++
						goto l350
					l354:
						position, tokenIndex = position350, tokenIndex350
						if buffer[position] != rune('s') {
							goto l339
						}
						position++
						if buffer[position] != rune('p') {
							goto l339
						}
						position++
						if buffer[position] != rune('e') {
							goto l339
						}
						position++
						if buffer[position] != rune('c') {
							goto l339
						}
						position++
						if buffer[position] != rune('i') {
							goto l339
						}
						position++
						if buffer[position] != rune('e') {
							goto l339
						}
						position++
						if buffer[position] != rune('s') {
							goto l339
						}
						position++
					}
				l350:
					{
						position355, tokenIndex355 := position, tokenIndex
						{
							position357, tokenIndex357 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l356
							}
							position, tokenIndex = position357, tokenIndex357
						}
						goto l355
					l356:
						position, tokenIndex = position355, tokenIndex355
						if buffer[position] != rune('.') {
							goto l339
						}
						position++
					}
				l355:
				}
			l341:
				add(ruleApproximation, position340)
			}
			return true
		l339:
			position, tokenIndex = position339, tokenIndex339
			return false
		},
		/* 47 Authorship <- <((AuthorshipCombo / OriginalAuthorship) &(SpaceCharEOI / ','))> */
		func() bool {
			position358, tokenIndex358 := position, tokenIndex
			{
				position359 := position
				{
					position360, tokenIndex360 := position, tokenIndex
					if !_rules[ruleAuthorshipCombo]() {
						goto l361
					}
					goto l360
				l361:
					position, tokenIndex = position360, tokenIndex360
					if !_rules[ruleOriginalAuthorship]() {
						goto l358
					}
				}
			l360:
				{
					position362, tokenIndex362 := position, tokenIndex
					{
						position363, tokenIndex363 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l364
						}
						goto l363
					l364:
						position, tokenIndex = position363, tokenIndex363
						if buffer[position] != rune(',') {
							goto l358
						}
						position++
					}
				l363:
					position, tokenIndex = position362, tokenIndex362
				}
				add(ruleAuthorship, position359)
			}
			return true
		l358:
			position, tokenIndex = position358, tokenIndex358
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
		/* 66 AuthorWord <- <(!('B' 'o' 'l' 'd' ':') (AuthorWord1 / AuthorWord2 / AuthorWord3 / AuthorPrefix))> */
		func() bool {
			position486, tokenIndex486 := position, tokenIndex
			{
				position487 := position
				{
					position488, tokenIndex488 := position, tokenIndex
					if buffer[position] != rune('B') {
						goto l488
					}
					position++
					if buffer[position] != rune('o') {
						goto l488
					}
					position++
					if buffer[position] != rune('l') {
						goto l488
					}
					position++
					if buffer[position] != rune('d') {
						goto l488
					}
					position++
					if buffer[position] != rune(':') {
						goto l488
					}
					position++
					goto l486
				l488:
					position, tokenIndex = position488, tokenIndex488
				}
				{
					position489, tokenIndex489 := position, tokenIndex
					if !_rules[ruleAuthorWord1]() {
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
						goto l486
					}
				}
			l489:
				add(ruleAuthorWord, position487)
			}
			return true
		l486:
			position, tokenIndex = position486, tokenIndex486
			return false
		},
		/* 67 AuthorWord1 <- <(('a' 'r' 'g' '.') / ('e' 't' ' ' 'a' 'l' '.' '{' '?' '}') / ((('e' 't') / '&') (' ' 'a' 'l') '.'?))> */
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
				add(ruleAuthorWord1, position494)
			}
			return true
		l493:
			position, tokenIndex = position493, tokenIndex493
			return false
		},
		/* 68 AuthorWord2 <- <(AuthorWord3 dash AuthorWordSoft)> */
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
		/* 69 AuthorWord3 <- <(AuthorPrefixGlued? (AllCapsAuthorWord / CapAuthorWord) '.'?)> */
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
		/* 70 AuthorWordSoft <- <(((AuthorUpperChar (AuthorUpperChar+ / AuthorLowerChar+)) / AuthorLowerChar+) '.'?)> */
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
		/* 71 CapAuthorWord <- <(AuthorUpperChar AuthorLowerChar*)> */
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
		/* 72 AllCapsAuthorWord <- <(AuthorUpperChar AuthorUpperChar+)> */
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
		/* 73 Filius <- <(('f' '.') / ('f' 'i' 'l' '.') / ('f' 'i' 'l' 'i' 'u' 's'))> */
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
		/* 74 AuthorPrefixGlued <- <(('d' '\'') / ('O' '\''))> */
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
		/* 75 AuthorPrefix <- <(AuthorPrefix1 / AuthorPrefix2)> */
		func() bool {
			position543, tokenIndex543 := position, tokenIndex
			{
				position544 := position
				{
					position545, tokenIndex545 := position, tokenIndex
					if !_rules[ruleAuthorPrefix1]() {
						goto l546
					}
					goto l545
				l546:
					position, tokenIndex = position545, tokenIndex545
					if !_rules[ruleAuthorPrefix2]() {
						goto l543
					}
				}
			l545:
				add(ruleAuthorPrefix, position544)
			}
			return true
		l543:
			position, tokenIndex = position543, tokenIndex543
			return false
		},
		/* 76 AuthorPrefix2 <- <(('v' '.' (_? ('d' '.'))?) / ('\'' 't'))> */
		func() bool {
			position547, tokenIndex547 := position, tokenIndex
			{
				position548 := position
				{
					position549, tokenIndex549 := position, tokenIndex
					if buffer[position] != rune('v') {
						goto l550
					}
					position++
					if buffer[position] != rune('.') {
						goto l550
					}
					position++
					{
						position551, tokenIndex551 := position, tokenIndex
						{
							position553, tokenIndex553 := position, tokenIndex
							if !_rules[rule_]() {
								goto l553
							}
							goto l554
						l553:
							position, tokenIndex = position553, tokenIndex553
						}
					l554:
						if buffer[position] != rune('d') {
							goto l551
						}
						position++
						if buffer[position] != rune('.') {
							goto l551
						}
						position++
						goto l552
					l551:
						position, tokenIndex = position551, tokenIndex551
					}
				l552:
					goto l549
				l550:
					position, tokenIndex = position549, tokenIndex549
					if buffer[position] != rune('\'') {
						goto l547
					}
					position++
					if buffer[position] != rune('t') {
						goto l547
					}
					position++
				}
			l549:
				add(ruleAuthorPrefix2, position548)
			}
			return true
		l547:
			position, tokenIndex = position547, tokenIndex547
			return false
		},
		/* 77 AuthorPrefix1 <- <((('a' 'b') / ('a' 'f') / ('b' 'i' 's') / ('d' 'a') / ('d' 'e' 'r') / ('d' 'e' 's') / ('d' 'e' 'n') / ('d' 'e' 'l') / ('d' 'e' 'l' 'l' 'a') / ('d' 'e' 'l' 'a') / ('d' 'e') / ('d' 'i') / ('d' 'u') / ('e' 'l') / ('l' 'a') / ('l' 'e') / ('t' 'e' 'r') / ('v' 'a' 'n') / ('d' '\'') / ('i' 'n' '\'' 't') / ('z' 'u' 'r') / ('v' 'o' 'n' (_ (('d' '.') / ('d' 'e' 'm')))?) / ('v' (_ 'd')?)) &_)> */
		func() bool {
			position555, tokenIndex555 := position, tokenIndex
			{
				position556 := position
				{
					position557, tokenIndex557 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l558
					}
					position++
					if buffer[position] != rune('b') {
						goto l558
					}
					position++
					goto l557
				l558:
					position, tokenIndex = position557, tokenIndex557
					if buffer[position] != rune('a') {
						goto l559
					}
					position++
					if buffer[position] != rune('f') {
						goto l559
					}
					position++
					goto l557
				l559:
					position, tokenIndex = position557, tokenIndex557
					if buffer[position] != rune('b') {
						goto l560
					}
					position++
					if buffer[position] != rune('i') {
						goto l560
					}
					position++
					if buffer[position] != rune('s') {
						goto l560
					}
					position++
					goto l557
				l560:
					position, tokenIndex = position557, tokenIndex557
					if buffer[position] != rune('d') {
						goto l561
					}
					position++
					if buffer[position] != rune('a') {
						goto l561
					}
					position++
					goto l557
				l561:
					position, tokenIndex = position557, tokenIndex557
					if buffer[position] != rune('d') {
						goto l562
					}
					position++
					if buffer[position] != rune('e') {
						goto l562
					}
					position++
					if buffer[position] != rune('r') {
						goto l562
					}
					position++
					goto l557
				l562:
					position, tokenIndex = position557, tokenIndex557
					if buffer[position] != rune('d') {
						goto l563
					}
					position++
					if buffer[position] != rune('e') {
						goto l563
					}
					position++
					if buffer[position] != rune('s') {
						goto l563
					}
					position++
					goto l557
				l563:
					position, tokenIndex = position557, tokenIndex557
					if buffer[position] != rune('d') {
						goto l564
					}
					position++
					if buffer[position] != rune('e') {
						goto l564
					}
					position++
					if buffer[position] != rune('n') {
						goto l564
					}
					position++
					goto l557
				l564:
					position, tokenIndex = position557, tokenIndex557
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
					goto l557
				l565:
					position, tokenIndex = position557, tokenIndex557
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
					if buffer[position] != rune('l') {
						goto l566
					}
					position++
					if buffer[position] != rune('a') {
						goto l566
					}
					position++
					goto l557
				l566:
					position, tokenIndex = position557, tokenIndex557
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
					if buffer[position] != rune('a') {
						goto l567
					}
					position++
					goto l557
				l567:
					position, tokenIndex = position557, tokenIndex557
					if buffer[position] != rune('d') {
						goto l568
					}
					position++
					if buffer[position] != rune('e') {
						goto l568
					}
					position++
					goto l557
				l568:
					position, tokenIndex = position557, tokenIndex557
					if buffer[position] != rune('d') {
						goto l569
					}
					position++
					if buffer[position] != rune('i') {
						goto l569
					}
					position++
					goto l557
				l569:
					position, tokenIndex = position557, tokenIndex557
					if buffer[position] != rune('d') {
						goto l570
					}
					position++
					if buffer[position] != rune('u') {
						goto l570
					}
					position++
					goto l557
				l570:
					position, tokenIndex = position557, tokenIndex557
					if buffer[position] != rune('e') {
						goto l571
					}
					position++
					if buffer[position] != rune('l') {
						goto l571
					}
					position++
					goto l557
				l571:
					position, tokenIndex = position557, tokenIndex557
					if buffer[position] != rune('l') {
						goto l572
					}
					position++
					if buffer[position] != rune('a') {
						goto l572
					}
					position++
					goto l557
				l572:
					position, tokenIndex = position557, tokenIndex557
					if buffer[position] != rune('l') {
						goto l573
					}
					position++
					if buffer[position] != rune('e') {
						goto l573
					}
					position++
					goto l557
				l573:
					position, tokenIndex = position557, tokenIndex557
					if buffer[position] != rune('t') {
						goto l574
					}
					position++
					if buffer[position] != rune('e') {
						goto l574
					}
					position++
					if buffer[position] != rune('r') {
						goto l574
					}
					position++
					goto l557
				l574:
					position, tokenIndex = position557, tokenIndex557
					if buffer[position] != rune('v') {
						goto l575
					}
					position++
					if buffer[position] != rune('a') {
						goto l575
					}
					position++
					if buffer[position] != rune('n') {
						goto l575
					}
					position++
					goto l557
				l575:
					position, tokenIndex = position557, tokenIndex557
					if buffer[position] != rune('d') {
						goto l576
					}
					position++
					if buffer[position] != rune('\'') {
						goto l576
					}
					position++
					goto l557
				l576:
					position, tokenIndex = position557, tokenIndex557
					if buffer[position] != rune('i') {
						goto l577
					}
					position++
					if buffer[position] != rune('n') {
						goto l577
					}
					position++
					if buffer[position] != rune('\'') {
						goto l577
					}
					position++
					if buffer[position] != rune('t') {
						goto l577
					}
					position++
					goto l557
				l577:
					position, tokenIndex = position557, tokenIndex557
					if buffer[position] != rune('z') {
						goto l578
					}
					position++
					if buffer[position] != rune('u') {
						goto l578
					}
					position++
					if buffer[position] != rune('r') {
						goto l578
					}
					position++
					goto l557
				l578:
					position, tokenIndex = position557, tokenIndex557
					if buffer[position] != rune('v') {
						goto l579
					}
					position++
					if buffer[position] != rune('o') {
						goto l579
					}
					position++
					if buffer[position] != rune('n') {
						goto l579
					}
					position++
					{
						position580, tokenIndex580 := position, tokenIndex
						if !_rules[rule_]() {
							goto l580
						}
						{
							position582, tokenIndex582 := position, tokenIndex
							if buffer[position] != rune('d') {
								goto l583
							}
							position++
							if buffer[position] != rune('.') {
								goto l583
							}
							position++
							goto l582
						l583:
							position, tokenIndex = position582, tokenIndex582
							if buffer[position] != rune('d') {
								goto l580
							}
							position++
							if buffer[position] != rune('e') {
								goto l580
							}
							position++
							if buffer[position] != rune('m') {
								goto l580
							}
							position++
						}
					l582:
						goto l581
					l580:
						position, tokenIndex = position580, tokenIndex580
					}
				l581:
					goto l557
				l579:
					position, tokenIndex = position557, tokenIndex557
					if buffer[position] != rune('v') {
						goto l555
					}
					position++
					{
						position584, tokenIndex584 := position, tokenIndex
						if !_rules[rule_]() {
							goto l584
						}
						if buffer[position] != rune('d') {
							goto l584
						}
						position++
						goto l585
					l584:
						position, tokenIndex = position584, tokenIndex584
					}
				l585:
				}
			l557:
				{
					position586, tokenIndex586 := position, tokenIndex
					if !_rules[rule_]() {
						goto l555
					}
					position, tokenIndex = position586, tokenIndex586
				}
				add(ruleAuthorPrefix1, position556)
			}
			return true
		l555:
			position, tokenIndex = position555, tokenIndex555
			return false
		},
		/* 78 AuthorUpperChar <- <(hASCII / ('À' / 'Á' / 'Â' / 'Ã' / 'Ä' / 'Å' / 'Æ' / 'Ç' / 'È' / 'É' / 'Ê' / 'Ë' / 'Ì' / 'Í' / 'Î' / 'Ï' / 'Ð' / 'Ñ' / 'Ò' / 'Ó' / 'Ô' / 'Õ' / 'Ö' / 'Ø' / 'Ù' / 'Ú' / 'Û' / 'Ü' / 'Ý' / 'Ć' / 'Č' / 'Ď' / 'İ' / 'Ķ' / 'Ĺ' / 'ĺ' / 'Ľ' / 'ľ' / 'Ł' / 'ł' / 'Ņ' / 'Ō' / 'Ő' / 'Œ' / 'Ř' / 'Ś' / 'Ŝ' / 'Ş' / 'Š' / 'Ÿ' / 'Ź' / 'Ż' / 'Ž' / 'ƒ' / 'Ǿ' / 'Ș' / 'Ț'))> */
		func() bool {
			position587, tokenIndex587 := position, tokenIndex
			{
				position588 := position
				{
					position589, tokenIndex589 := position, tokenIndex
					if !_rules[rulehASCII]() {
						goto l590
					}
					goto l589
				l590:
					position, tokenIndex = position589, tokenIndex589
					{
						position591, tokenIndex591 := position, tokenIndex
						if buffer[position] != rune('À') {
							goto l592
						}
						position++
						goto l591
					l592:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Á') {
							goto l593
						}
						position++
						goto l591
					l593:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Â') {
							goto l594
						}
						position++
						goto l591
					l594:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ã') {
							goto l595
						}
						position++
						goto l591
					l595:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ä') {
							goto l596
						}
						position++
						goto l591
					l596:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Å') {
							goto l597
						}
						position++
						goto l591
					l597:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Æ') {
							goto l598
						}
						position++
						goto l591
					l598:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ç') {
							goto l599
						}
						position++
						goto l591
					l599:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('È') {
							goto l600
						}
						position++
						goto l591
					l600:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('É') {
							goto l601
						}
						position++
						goto l591
					l601:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ê') {
							goto l602
						}
						position++
						goto l591
					l602:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ë') {
							goto l603
						}
						position++
						goto l591
					l603:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ì') {
							goto l604
						}
						position++
						goto l591
					l604:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Í') {
							goto l605
						}
						position++
						goto l591
					l605:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Î') {
							goto l606
						}
						position++
						goto l591
					l606:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ï') {
							goto l607
						}
						position++
						goto l591
					l607:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ð') {
							goto l608
						}
						position++
						goto l591
					l608:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ñ') {
							goto l609
						}
						position++
						goto l591
					l609:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ò') {
							goto l610
						}
						position++
						goto l591
					l610:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ó') {
							goto l611
						}
						position++
						goto l591
					l611:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ô') {
							goto l612
						}
						position++
						goto l591
					l612:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Õ') {
							goto l613
						}
						position++
						goto l591
					l613:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ö') {
							goto l614
						}
						position++
						goto l591
					l614:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ø') {
							goto l615
						}
						position++
						goto l591
					l615:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ù') {
							goto l616
						}
						position++
						goto l591
					l616:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ú') {
							goto l617
						}
						position++
						goto l591
					l617:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Û') {
							goto l618
						}
						position++
						goto l591
					l618:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ü') {
							goto l619
						}
						position++
						goto l591
					l619:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ý') {
							goto l620
						}
						position++
						goto l591
					l620:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ć') {
							goto l621
						}
						position++
						goto l591
					l621:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Č') {
							goto l622
						}
						position++
						goto l591
					l622:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ď') {
							goto l623
						}
						position++
						goto l591
					l623:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('İ') {
							goto l624
						}
						position++
						goto l591
					l624:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ķ') {
							goto l625
						}
						position++
						goto l591
					l625:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ĺ') {
							goto l626
						}
						position++
						goto l591
					l626:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('ĺ') {
							goto l627
						}
						position++
						goto l591
					l627:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ľ') {
							goto l628
						}
						position++
						goto l591
					l628:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('ľ') {
							goto l629
						}
						position++
						goto l591
					l629:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ł') {
							goto l630
						}
						position++
						goto l591
					l630:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('ł') {
							goto l631
						}
						position++
						goto l591
					l631:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ņ') {
							goto l632
						}
						position++
						goto l591
					l632:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ō') {
							goto l633
						}
						position++
						goto l591
					l633:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ő') {
							goto l634
						}
						position++
						goto l591
					l634:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Œ') {
							goto l635
						}
						position++
						goto l591
					l635:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ř') {
							goto l636
						}
						position++
						goto l591
					l636:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ś') {
							goto l637
						}
						position++
						goto l591
					l637:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ŝ') {
							goto l638
						}
						position++
						goto l591
					l638:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ş') {
							goto l639
						}
						position++
						goto l591
					l639:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Š') {
							goto l640
						}
						position++
						goto l591
					l640:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ÿ') {
							goto l641
						}
						position++
						goto l591
					l641:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ź') {
							goto l642
						}
						position++
						goto l591
					l642:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ż') {
							goto l643
						}
						position++
						goto l591
					l643:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ž') {
							goto l644
						}
						position++
						goto l591
					l644:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('ƒ') {
							goto l645
						}
						position++
						goto l591
					l645:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ǿ') {
							goto l646
						}
						position++
						goto l591
					l646:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ș') {
							goto l647
						}
						position++
						goto l591
					l647:
						position, tokenIndex = position591, tokenIndex591
						if buffer[position] != rune('Ț') {
							goto l587
						}
						position++
					}
				l591:
				}
			l589:
				add(ruleAuthorUpperChar, position588)
			}
			return true
		l587:
			position, tokenIndex = position587, tokenIndex587
			return false
		},
		/* 79 AuthorLowerChar <- <(lASCII / ('à' / 'á' / 'â' / 'ã' / 'ä' / 'å' / 'æ' / 'ç' / 'è' / 'é' / 'ê' / 'ë' / 'ì' / 'í' / 'î' / 'ï' / 'ð' / 'ñ' / 'ò' / 'ó' / 'ó' / 'ô' / 'õ' / 'ö' / 'ø' / 'ù' / 'ú' / 'û' / 'ü' / 'ý' / 'ÿ' / 'ā' / 'ă' / 'ą' / 'ć' / 'ĉ' / 'č' / 'ď' / 'đ' / '\'' / 'ē' / 'ĕ' / 'ė' / 'ę' / 'ě' / 'ğ' / 'ī' / 'ĭ' / 'İ' / 'ı' / 'ĺ' / 'ľ' / 'ł' / 'ń' / 'ņ' / 'ň' / 'ŏ' / 'ő' / 'œ' / 'ŕ' / 'ř' / 'ś' / 'ş' / 'š' / 'ţ' / 'ť' / 'ũ' / 'ū' / 'ŭ' / 'ů' / 'ű' / 'ź' / 'ż' / 'ž' / 'ſ' / 'ǎ' / 'ǔ' / 'ǧ' / 'ș' / 'ț' / 'ȳ' / 'ß'))> */
		func() bool {
			position648, tokenIndex648 := position, tokenIndex
			{
				position649 := position
				{
					position650, tokenIndex650 := position, tokenIndex
					if !_rules[rulelASCII]() {
						goto l651
					}
					goto l650
				l651:
					position, tokenIndex = position650, tokenIndex650
					{
						position652, tokenIndex652 := position, tokenIndex
						if buffer[position] != rune('à') {
							goto l653
						}
						position++
						goto l652
					l653:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('á') {
							goto l654
						}
						position++
						goto l652
					l654:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('â') {
							goto l655
						}
						position++
						goto l652
					l655:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ã') {
							goto l656
						}
						position++
						goto l652
					l656:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ä') {
							goto l657
						}
						position++
						goto l652
					l657:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('å') {
							goto l658
						}
						position++
						goto l652
					l658:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('æ') {
							goto l659
						}
						position++
						goto l652
					l659:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ç') {
							goto l660
						}
						position++
						goto l652
					l660:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('è') {
							goto l661
						}
						position++
						goto l652
					l661:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('é') {
							goto l662
						}
						position++
						goto l652
					l662:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ê') {
							goto l663
						}
						position++
						goto l652
					l663:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ë') {
							goto l664
						}
						position++
						goto l652
					l664:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ì') {
							goto l665
						}
						position++
						goto l652
					l665:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('í') {
							goto l666
						}
						position++
						goto l652
					l666:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('î') {
							goto l667
						}
						position++
						goto l652
					l667:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ï') {
							goto l668
						}
						position++
						goto l652
					l668:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ð') {
							goto l669
						}
						position++
						goto l652
					l669:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ñ') {
							goto l670
						}
						position++
						goto l652
					l670:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ò') {
							goto l671
						}
						position++
						goto l652
					l671:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ó') {
							goto l672
						}
						position++
						goto l652
					l672:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ó') {
							goto l673
						}
						position++
						goto l652
					l673:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ô') {
							goto l674
						}
						position++
						goto l652
					l674:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('õ') {
							goto l675
						}
						position++
						goto l652
					l675:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ö') {
							goto l676
						}
						position++
						goto l652
					l676:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ø') {
							goto l677
						}
						position++
						goto l652
					l677:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ù') {
							goto l678
						}
						position++
						goto l652
					l678:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ú') {
							goto l679
						}
						position++
						goto l652
					l679:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('û') {
							goto l680
						}
						position++
						goto l652
					l680:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ü') {
							goto l681
						}
						position++
						goto l652
					l681:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ý') {
							goto l682
						}
						position++
						goto l652
					l682:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ÿ') {
							goto l683
						}
						position++
						goto l652
					l683:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ā') {
							goto l684
						}
						position++
						goto l652
					l684:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ă') {
							goto l685
						}
						position++
						goto l652
					l685:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ą') {
							goto l686
						}
						position++
						goto l652
					l686:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ć') {
							goto l687
						}
						position++
						goto l652
					l687:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ĉ') {
							goto l688
						}
						position++
						goto l652
					l688:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('č') {
							goto l689
						}
						position++
						goto l652
					l689:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ď') {
							goto l690
						}
						position++
						goto l652
					l690:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('đ') {
							goto l691
						}
						position++
						goto l652
					l691:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('\'') {
							goto l692
						}
						position++
						goto l652
					l692:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ē') {
							goto l693
						}
						position++
						goto l652
					l693:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ĕ') {
							goto l694
						}
						position++
						goto l652
					l694:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ė') {
							goto l695
						}
						position++
						goto l652
					l695:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ę') {
							goto l696
						}
						position++
						goto l652
					l696:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ě') {
							goto l697
						}
						position++
						goto l652
					l697:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ğ') {
							goto l698
						}
						position++
						goto l652
					l698:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ī') {
							goto l699
						}
						position++
						goto l652
					l699:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ĭ') {
							goto l700
						}
						position++
						goto l652
					l700:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('İ') {
							goto l701
						}
						position++
						goto l652
					l701:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ı') {
							goto l702
						}
						position++
						goto l652
					l702:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ĺ') {
							goto l703
						}
						position++
						goto l652
					l703:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ľ') {
							goto l704
						}
						position++
						goto l652
					l704:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ł') {
							goto l705
						}
						position++
						goto l652
					l705:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ń') {
							goto l706
						}
						position++
						goto l652
					l706:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ņ') {
							goto l707
						}
						position++
						goto l652
					l707:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ň') {
							goto l708
						}
						position++
						goto l652
					l708:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ŏ') {
							goto l709
						}
						position++
						goto l652
					l709:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ő') {
							goto l710
						}
						position++
						goto l652
					l710:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('œ') {
							goto l711
						}
						position++
						goto l652
					l711:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ŕ') {
							goto l712
						}
						position++
						goto l652
					l712:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ř') {
							goto l713
						}
						position++
						goto l652
					l713:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ś') {
							goto l714
						}
						position++
						goto l652
					l714:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ş') {
							goto l715
						}
						position++
						goto l652
					l715:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('š') {
							goto l716
						}
						position++
						goto l652
					l716:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ţ') {
							goto l717
						}
						position++
						goto l652
					l717:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ť') {
							goto l718
						}
						position++
						goto l652
					l718:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ũ') {
							goto l719
						}
						position++
						goto l652
					l719:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ū') {
							goto l720
						}
						position++
						goto l652
					l720:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ŭ') {
							goto l721
						}
						position++
						goto l652
					l721:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ů') {
							goto l722
						}
						position++
						goto l652
					l722:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ű') {
							goto l723
						}
						position++
						goto l652
					l723:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ź') {
							goto l724
						}
						position++
						goto l652
					l724:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ż') {
							goto l725
						}
						position++
						goto l652
					l725:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ž') {
							goto l726
						}
						position++
						goto l652
					l726:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ſ') {
							goto l727
						}
						position++
						goto l652
					l727:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ǎ') {
							goto l728
						}
						position++
						goto l652
					l728:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ǔ') {
							goto l729
						}
						position++
						goto l652
					l729:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ǧ') {
							goto l730
						}
						position++
						goto l652
					l730:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ș') {
							goto l731
						}
						position++
						goto l652
					l731:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ț') {
							goto l732
						}
						position++
						goto l652
					l732:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ȳ') {
							goto l733
						}
						position++
						goto l652
					l733:
						position, tokenIndex = position652, tokenIndex652
						if buffer[position] != rune('ß') {
							goto l648
						}
						position++
					}
				l652:
				}
			l650:
				add(ruleAuthorLowerChar, position649)
			}
			return true
		l648:
			position, tokenIndex = position648, tokenIndex648
			return false
		},
		/* 80 Year <- <(YearRange / YearApprox / YearWithParens / YearWithPage / YearWithDot / YearWithChar / YearNum)> */
		func() bool {
			position734, tokenIndex734 := position, tokenIndex
			{
				position735 := position
				{
					position736, tokenIndex736 := position, tokenIndex
					if !_rules[ruleYearRange]() {
						goto l737
					}
					goto l736
				l737:
					position, tokenIndex = position736, tokenIndex736
					if !_rules[ruleYearApprox]() {
						goto l738
					}
					goto l736
				l738:
					position, tokenIndex = position736, tokenIndex736
					if !_rules[ruleYearWithParens]() {
						goto l739
					}
					goto l736
				l739:
					position, tokenIndex = position736, tokenIndex736
					if !_rules[ruleYearWithPage]() {
						goto l740
					}
					goto l736
				l740:
					position, tokenIndex = position736, tokenIndex736
					if !_rules[ruleYearWithDot]() {
						goto l741
					}
					goto l736
				l741:
					position, tokenIndex = position736, tokenIndex736
					if !_rules[ruleYearWithChar]() {
						goto l742
					}
					goto l736
				l742:
					position, tokenIndex = position736, tokenIndex736
					if !_rules[ruleYearNum]() {
						goto l734
					}
				}
			l736:
				add(ruleYear, position735)
			}
			return true
		l734:
			position, tokenIndex = position734, tokenIndex734
			return false
		},
		/* 81 YearRange <- <(YearNum dash (nums+ ('a' / 'b' / 'c' / 'd' / 'e' / 'f' / 'g' / 'h' / 'i' / 'j' / 'k' / 'l' / 'm' / 'n' / 'o' / 'p' / 'q' / 'r' / 's' / 't' / 'u' / 'v' / 'w' / 'x' / 'y' / 'z' / '?')*))> */
		func() bool {
			position743, tokenIndex743 := position, tokenIndex
			{
				position744 := position
				if !_rules[ruleYearNum]() {
					goto l743
				}
				if !_rules[ruledash]() {
					goto l743
				}
				if !_rules[rulenums]() {
					goto l743
				}
			l745:
				{
					position746, tokenIndex746 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l746
					}
					goto l745
				l746:
					position, tokenIndex = position746, tokenIndex746
				}
			l747:
				{
					position748, tokenIndex748 := position, tokenIndex
					{
						position749, tokenIndex749 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l750
						}
						position++
						goto l749
					l750:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('b') {
							goto l751
						}
						position++
						goto l749
					l751:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('c') {
							goto l752
						}
						position++
						goto l749
					l752:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('d') {
							goto l753
						}
						position++
						goto l749
					l753:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('e') {
							goto l754
						}
						position++
						goto l749
					l754:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('f') {
							goto l755
						}
						position++
						goto l749
					l755:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('g') {
							goto l756
						}
						position++
						goto l749
					l756:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('h') {
							goto l757
						}
						position++
						goto l749
					l757:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('i') {
							goto l758
						}
						position++
						goto l749
					l758:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('j') {
							goto l759
						}
						position++
						goto l749
					l759:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('k') {
							goto l760
						}
						position++
						goto l749
					l760:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('l') {
							goto l761
						}
						position++
						goto l749
					l761:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('m') {
							goto l762
						}
						position++
						goto l749
					l762:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('n') {
							goto l763
						}
						position++
						goto l749
					l763:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('o') {
							goto l764
						}
						position++
						goto l749
					l764:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('p') {
							goto l765
						}
						position++
						goto l749
					l765:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('q') {
							goto l766
						}
						position++
						goto l749
					l766:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('r') {
							goto l767
						}
						position++
						goto l749
					l767:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('s') {
							goto l768
						}
						position++
						goto l749
					l768:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('t') {
							goto l769
						}
						position++
						goto l749
					l769:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('u') {
							goto l770
						}
						position++
						goto l749
					l770:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('v') {
							goto l771
						}
						position++
						goto l749
					l771:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('w') {
							goto l772
						}
						position++
						goto l749
					l772:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('x') {
							goto l773
						}
						position++
						goto l749
					l773:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('y') {
							goto l774
						}
						position++
						goto l749
					l774:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('z') {
							goto l775
						}
						position++
						goto l749
					l775:
						position, tokenIndex = position749, tokenIndex749
						if buffer[position] != rune('?') {
							goto l748
						}
						position++
					}
				l749:
					goto l747
				l748:
					position, tokenIndex = position748, tokenIndex748
				}
				add(ruleYearRange, position744)
			}
			return true
		l743:
			position, tokenIndex = position743, tokenIndex743
			return false
		},
		/* 82 YearWithDot <- <(YearNum '.')> */
		func() bool {
			position776, tokenIndex776 := position, tokenIndex
			{
				position777 := position
				if !_rules[ruleYearNum]() {
					goto l776
				}
				if buffer[position] != rune('.') {
					goto l776
				}
				position++
				add(ruleYearWithDot, position777)
			}
			return true
		l776:
			position, tokenIndex = position776, tokenIndex776
			return false
		},
		/* 83 YearApprox <- <('[' _? YearNum _? ']')> */
		func() bool {
			position778, tokenIndex778 := position, tokenIndex
			{
				position779 := position
				if buffer[position] != rune('[') {
					goto l778
				}
				position++
				{
					position780, tokenIndex780 := position, tokenIndex
					if !_rules[rule_]() {
						goto l780
					}
					goto l781
				l780:
					position, tokenIndex = position780, tokenIndex780
				}
			l781:
				if !_rules[ruleYearNum]() {
					goto l778
				}
				{
					position782, tokenIndex782 := position, tokenIndex
					if !_rules[rule_]() {
						goto l782
					}
					goto l783
				l782:
					position, tokenIndex = position782, tokenIndex782
				}
			l783:
				if buffer[position] != rune(']') {
					goto l778
				}
				position++
				add(ruleYearApprox, position779)
			}
			return true
		l778:
			position, tokenIndex = position778, tokenIndex778
			return false
		},
		/* 84 YearWithPage <- <((YearWithChar / YearNum) _? ':' _? nums+)> */
		func() bool {
			position784, tokenIndex784 := position, tokenIndex
			{
				position785 := position
				{
					position786, tokenIndex786 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l787
					}
					goto l786
				l787:
					position, tokenIndex = position786, tokenIndex786
					if !_rules[ruleYearNum]() {
						goto l784
					}
				}
			l786:
				{
					position788, tokenIndex788 := position, tokenIndex
					if !_rules[rule_]() {
						goto l788
					}
					goto l789
				l788:
					position, tokenIndex = position788, tokenIndex788
				}
			l789:
				if buffer[position] != rune(':') {
					goto l784
				}
				position++
				{
					position790, tokenIndex790 := position, tokenIndex
					if !_rules[rule_]() {
						goto l790
					}
					goto l791
				l790:
					position, tokenIndex = position790, tokenIndex790
				}
			l791:
				if !_rules[rulenums]() {
					goto l784
				}
			l792:
				{
					position793, tokenIndex793 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l793
					}
					goto l792
				l793:
					position, tokenIndex = position793, tokenIndex793
				}
				add(ruleYearWithPage, position785)
			}
			return true
		l784:
			position, tokenIndex = position784, tokenIndex784
			return false
		},
		/* 85 YearWithParens <- <('(' (YearWithChar / YearNum) ')')> */
		func() bool {
			position794, tokenIndex794 := position, tokenIndex
			{
				position795 := position
				if buffer[position] != rune('(') {
					goto l794
				}
				position++
				{
					position796, tokenIndex796 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l797
					}
					goto l796
				l797:
					position, tokenIndex = position796, tokenIndex796
					if !_rules[ruleYearNum]() {
						goto l794
					}
				}
			l796:
				if buffer[position] != rune(')') {
					goto l794
				}
				position++
				add(ruleYearWithParens, position795)
			}
			return true
		l794:
			position, tokenIndex = position794, tokenIndex794
			return false
		},
		/* 86 YearWithChar <- <(YearNum lASCII Action0)> */
		func() bool {
			position798, tokenIndex798 := position, tokenIndex
			{
				position799 := position
				if !_rules[ruleYearNum]() {
					goto l798
				}
				if !_rules[rulelASCII]() {
					goto l798
				}
				if !_rules[ruleAction0]() {
					goto l798
				}
				add(ruleYearWithChar, position799)
			}
			return true
		l798:
			position, tokenIndex = position798, tokenIndex798
			return false
		},
		/* 87 YearNum <- <(('1' / '2') ('0' / '7' / '8' / '9') nums (nums / '?') '?'*)> */
		func() bool {
			position800, tokenIndex800 := position, tokenIndex
			{
				position801 := position
				{
					position802, tokenIndex802 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l803
					}
					position++
					goto l802
				l803:
					position, tokenIndex = position802, tokenIndex802
					if buffer[position] != rune('2') {
						goto l800
					}
					position++
				}
			l802:
				{
					position804, tokenIndex804 := position, tokenIndex
					if buffer[position] != rune('0') {
						goto l805
					}
					position++
					goto l804
				l805:
					position, tokenIndex = position804, tokenIndex804
					if buffer[position] != rune('7') {
						goto l806
					}
					position++
					goto l804
				l806:
					position, tokenIndex = position804, tokenIndex804
					if buffer[position] != rune('8') {
						goto l807
					}
					position++
					goto l804
				l807:
					position, tokenIndex = position804, tokenIndex804
					if buffer[position] != rune('9') {
						goto l800
					}
					position++
				}
			l804:
				if !_rules[rulenums]() {
					goto l800
				}
				{
					position808, tokenIndex808 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l809
					}
					goto l808
				l809:
					position, tokenIndex = position808, tokenIndex808
					if buffer[position] != rune('?') {
						goto l800
					}
					position++
				}
			l808:
			l810:
				{
					position811, tokenIndex811 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l811
					}
					position++
					goto l810
				l811:
					position, tokenIndex = position811, tokenIndex811
				}
				add(ruleYearNum, position801)
			}
			return true
		l800:
			position, tokenIndex = position800, tokenIndex800
			return false
		},
		/* 88 NameUpperChar <- <(UpperChar / UpperCharExtended)> */
		func() bool {
			position812, tokenIndex812 := position, tokenIndex
			{
				position813 := position
				{
					position814, tokenIndex814 := position, tokenIndex
					if !_rules[ruleUpperChar]() {
						goto l815
					}
					goto l814
				l815:
					position, tokenIndex = position814, tokenIndex814
					if !_rules[ruleUpperCharExtended]() {
						goto l812
					}
				}
			l814:
				add(ruleNameUpperChar, position813)
			}
			return true
		l812:
			position, tokenIndex = position812, tokenIndex812
			return false
		},
		/* 89 UpperCharExtended <- <('Æ' / 'Œ' / 'Ö')> */
		func() bool {
			position816, tokenIndex816 := position, tokenIndex
			{
				position817 := position
				{
					position818, tokenIndex818 := position, tokenIndex
					if buffer[position] != rune('Æ') {
						goto l819
					}
					position++
					goto l818
				l819:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('Œ') {
						goto l820
					}
					position++
					goto l818
				l820:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('Ö') {
						goto l816
					}
					position++
				}
			l818:
				add(ruleUpperCharExtended, position817)
			}
			return true
		l816:
			position, tokenIndex = position816, tokenIndex816
			return false
		},
		/* 90 UpperChar <- <hASCII> */
		func() bool {
			position821, tokenIndex821 := position, tokenIndex
			{
				position822 := position
				if !_rules[rulehASCII]() {
					goto l821
				}
				add(ruleUpperChar, position822)
			}
			return true
		l821:
			position, tokenIndex = position821, tokenIndex821
			return false
		},
		/* 91 NameLowerChar <- <(LowerChar / LowerCharExtended / MiscodedChar)> */
		func() bool {
			position823, tokenIndex823 := position, tokenIndex
			{
				position824 := position
				{
					position825, tokenIndex825 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l826
					}
					goto l825
				l826:
					position, tokenIndex = position825, tokenIndex825
					if !_rules[ruleLowerCharExtended]() {
						goto l827
					}
					goto l825
				l827:
					position, tokenIndex = position825, tokenIndex825
					if !_rules[ruleMiscodedChar]() {
						goto l823
					}
				}
			l825:
				add(ruleNameLowerChar, position824)
			}
			return true
		l823:
			position, tokenIndex = position823, tokenIndex823
			return false
		},
		/* 92 MiscodedChar <- <'�'> */
		func() bool {
			position828, tokenIndex828 := position, tokenIndex
			{
				position829 := position
				if buffer[position] != rune('�') {
					goto l828
				}
				position++
				add(ruleMiscodedChar, position829)
			}
			return true
		l828:
			position, tokenIndex = position828, tokenIndex828
			return false
		},
		/* 93 LowerCharExtended <- <('æ' / 'œ' / 'à' / 'â' / 'å' / 'ã' / 'ä' / 'á' / 'ç' / 'č' / 'é' / 'è' / 'ë' / 'í' / 'ì' / 'ï' / 'ň' / 'ñ' / 'ñ' / 'ó' / 'ò' / 'ô' / 'ø' / 'õ' / 'ö' / 'ú' / 'ù' / 'ü' / 'ŕ' / 'ř' / 'ŗ' / 'ſ' / 'š' / 'š' / 'ş' / 'ž')> */
		func() bool {
			position830, tokenIndex830 := position, tokenIndex
			{
				position831 := position
				{
					position832, tokenIndex832 := position, tokenIndex
					if buffer[position] != rune('æ') {
						goto l833
					}
					position++
					goto l832
				l833:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('œ') {
						goto l834
					}
					position++
					goto l832
				l834:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('à') {
						goto l835
					}
					position++
					goto l832
				l835:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('â') {
						goto l836
					}
					position++
					goto l832
				l836:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('å') {
						goto l837
					}
					position++
					goto l832
				l837:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('ã') {
						goto l838
					}
					position++
					goto l832
				l838:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('ä') {
						goto l839
					}
					position++
					goto l832
				l839:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('á') {
						goto l840
					}
					position++
					goto l832
				l840:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('ç') {
						goto l841
					}
					position++
					goto l832
				l841:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('č') {
						goto l842
					}
					position++
					goto l832
				l842:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('é') {
						goto l843
					}
					position++
					goto l832
				l843:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('è') {
						goto l844
					}
					position++
					goto l832
				l844:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('ë') {
						goto l845
					}
					position++
					goto l832
				l845:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('í') {
						goto l846
					}
					position++
					goto l832
				l846:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('ì') {
						goto l847
					}
					position++
					goto l832
				l847:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('ï') {
						goto l848
					}
					position++
					goto l832
				l848:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('ň') {
						goto l849
					}
					position++
					goto l832
				l849:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('ñ') {
						goto l850
					}
					position++
					goto l832
				l850:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('ñ') {
						goto l851
					}
					position++
					goto l832
				l851:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('ó') {
						goto l852
					}
					position++
					goto l832
				l852:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('ò') {
						goto l853
					}
					position++
					goto l832
				l853:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('ô') {
						goto l854
					}
					position++
					goto l832
				l854:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('ø') {
						goto l855
					}
					position++
					goto l832
				l855:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('õ') {
						goto l856
					}
					position++
					goto l832
				l856:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('ö') {
						goto l857
					}
					position++
					goto l832
				l857:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('ú') {
						goto l858
					}
					position++
					goto l832
				l858:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('ù') {
						goto l859
					}
					position++
					goto l832
				l859:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('ü') {
						goto l860
					}
					position++
					goto l832
				l860:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('ŕ') {
						goto l861
					}
					position++
					goto l832
				l861:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('ř') {
						goto l862
					}
					position++
					goto l832
				l862:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('ŗ') {
						goto l863
					}
					position++
					goto l832
				l863:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('ſ') {
						goto l864
					}
					position++
					goto l832
				l864:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('š') {
						goto l865
					}
					position++
					goto l832
				l865:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('š') {
						goto l866
					}
					position++
					goto l832
				l866:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('ş') {
						goto l867
					}
					position++
					goto l832
				l867:
					position, tokenIndex = position832, tokenIndex832
					if buffer[position] != rune('ž') {
						goto l830
					}
					position++
				}
			l832:
				add(ruleLowerCharExtended, position831)
			}
			return true
		l830:
			position, tokenIndex = position830, tokenIndex830
			return false
		},
		/* 94 LowerChar <- <lASCII> */
		func() bool {
			position868, tokenIndex868 := position, tokenIndex
			{
				position869 := position
				if !_rules[rulelASCII]() {
					goto l868
				}
				add(ruleLowerChar, position869)
			}
			return true
		l868:
			position, tokenIndex = position868, tokenIndex868
			return false
		},
		/* 95 SpaceCharEOI <- <(_ / !.)> */
		func() bool {
			position870, tokenIndex870 := position, tokenIndex
			{
				position871 := position
				{
					position872, tokenIndex872 := position, tokenIndex
					if !_rules[rule_]() {
						goto l873
					}
					goto l872
				l873:
					position, tokenIndex = position872, tokenIndex872
					{
						position874, tokenIndex874 := position, tokenIndex
						if !matchDot() {
							goto l874
						}
						goto l870
					l874:
						position, tokenIndex = position874, tokenIndex874
					}
				}
			l872:
				add(ruleSpaceCharEOI, position871)
			}
			return true
		l870:
			position, tokenIndex = position870, tokenIndex870
			return false
		},
		/* 96 WordBorderChar <- <(_ / ';' / '.' / ',' / ':' / '(' / ')' / ']')> */
		nil,
		/* 97 nums <- <[0-9]> */
		func() bool {
			position876, tokenIndex876 := position, tokenIndex
			{
				position877 := position
				if c := buffer[position]; c < rune('0') || c > rune('9') {
					goto l876
				}
				position++
				add(rulenums, position877)
			}
			return true
		l876:
			position, tokenIndex = position876, tokenIndex876
			return false
		},
		/* 98 lASCII <- <[a-z]> */
		func() bool {
			position878, tokenIndex878 := position, tokenIndex
			{
				position879 := position
				if c := buffer[position]; c < rune('a') || c > rune('z') {
					goto l878
				}
				position++
				add(rulelASCII, position879)
			}
			return true
		l878:
			position, tokenIndex = position878, tokenIndex878
			return false
		},
		/* 99 hASCII <- <[A-Z]> */
		func() bool {
			position880, tokenIndex880 := position, tokenIndex
			{
				position881 := position
				if c := buffer[position]; c < rune('A') || c > rune('Z') {
					goto l880
				}
				position++
				add(rulehASCII, position881)
			}
			return true
		l880:
			position, tokenIndex = position880, tokenIndex880
			return false
		},
		/* 100 apostr <- <'\''> */
		func() bool {
			position882, tokenIndex882 := position, tokenIndex
			{
				position883 := position
				if buffer[position] != rune('\'') {
					goto l882
				}
				position++
				add(ruleapostr, position883)
			}
			return true
		l882:
			position, tokenIndex = position882, tokenIndex882
			return false
		},
		/* 101 dash <- <'-'> */
		func() bool {
			position884, tokenIndex884 := position, tokenIndex
			{
				position885 := position
				if buffer[position] != rune('-') {
					goto l884
				}
				position++
				add(ruledash, position885)
			}
			return true
		l884:
			position, tokenIndex = position884, tokenIndex884
			return false
		},
		/* 102 _ <- <(MultipleSpace / SingleSpace)> */
		func() bool {
			position886, tokenIndex886 := position, tokenIndex
			{
				position887 := position
				{
					position888, tokenIndex888 := position, tokenIndex
					if !_rules[ruleMultipleSpace]() {
						goto l889
					}
					goto l888
				l889:
					position, tokenIndex = position888, tokenIndex888
					if !_rules[ruleSingleSpace]() {
						goto l886
					}
				}
			l888:
				add(rule_, position887)
			}
			return true
		l886:
			position, tokenIndex = position886, tokenIndex886
			return false
		},
		/* 103 MultipleSpace <- <(SingleSpace SingleSpace+)> */
		func() bool {
			position890, tokenIndex890 := position, tokenIndex
			{
				position891 := position
				if !_rules[ruleSingleSpace]() {
					goto l890
				}
				if !_rules[ruleSingleSpace]() {
					goto l890
				}
			l892:
				{
					position893, tokenIndex893 := position, tokenIndex
					if !_rules[ruleSingleSpace]() {
						goto l893
					}
					goto l892
				l893:
					position, tokenIndex = position893, tokenIndex893
				}
				add(ruleMultipleSpace, position891)
			}
			return true
		l890:
			position, tokenIndex = position890, tokenIndex890
			return false
		},
		/* 104 SingleSpace <- <' '> */
		func() bool {
			position894, tokenIndex894 := position, tokenIndex
			{
				position895 := position
				if buffer[position] != rune(' ') {
					goto l894
				}
				position++
				add(ruleSingleSpace, position895)
			}
			return true
		l894:
			position, tokenIndex = position894, tokenIndex894
			return false
		},
		/* 106 Action0 <- <{ p.AddWarn(YearCharWarn) }> */
		func() bool {
			{
				add(ruleAction0, position)
			}
			return true
		},
	}
	p.rules = _rules
}
