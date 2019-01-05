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
	rules  [110]func() bool
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
		/* 9 SingleName <- <(NameSpecies / NameUninomial)> */
		func() bool {
			position42, tokenIndex42 := position, tokenIndex
			{
				position43 := position
				{
					position44, tokenIndex44 := position, tokenIndex
					if !_rules[ruleNameSpecies]() {
						goto l45
					}
					goto l44
				l45:
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
			position46, tokenIndex46 := position, tokenIndex
			{
				position47 := position
				{
					position48, tokenIndex48 := position, tokenIndex
					if !_rules[ruleUninomialCombo]() {
						goto l49
					}
					goto l48
				l49:
					position, tokenIndex = position48, tokenIndex48
					if !_rules[ruleUninomial]() {
						goto l46
					}
				}
			l48:
				add(ruleNameUninomial, position47)
			}
			return true
		l46:
			position, tokenIndex = position46, tokenIndex46
			return false
		},
		/* 11 NameApprox <- <(GenusWord _ Approximation (_ SpeciesEpithet)?)> */
		nil,
		/* 12 NameComp <- <(GenusWord _ Comparison (_ SpeciesEpithet)?)> */
		nil,
		/* 13 NameSpecies <- <(GenusWord (_? (SubGenus / SubGenusOrSuperspecies))? _ SpeciesEpithet (_ InfraspGroup)?)> */
		func() bool {
			position52, tokenIndex52 := position, tokenIndex
			{
				position53 := position
				if !_rules[ruleGenusWord]() {
					goto l52
				}
				{
					position54, tokenIndex54 := position, tokenIndex
					{
						position56, tokenIndex56 := position, tokenIndex
						if !_rules[rule_]() {
							goto l56
						}
						goto l57
					l56:
						position, tokenIndex = position56, tokenIndex56
					}
				l57:
					{
						position58, tokenIndex58 := position, tokenIndex
						if !_rules[ruleSubGenus]() {
							goto l59
						}
						goto l58
					l59:
						position, tokenIndex = position58, tokenIndex58
						if !_rules[ruleSubGenusOrSuperspecies]() {
							goto l54
						}
					}
				l58:
					goto l55
				l54:
					position, tokenIndex = position54, tokenIndex54
				}
			l55:
				if !_rules[rule_]() {
					goto l52
				}
				if !_rules[ruleSpeciesEpithet]() {
					goto l52
				}
				{
					position60, tokenIndex60 := position, tokenIndex
					if !_rules[rule_]() {
						goto l60
					}
					if !_rules[ruleInfraspGroup]() {
						goto l60
					}
					goto l61
				l60:
					position, tokenIndex = position60, tokenIndex60
				}
			l61:
				add(ruleNameSpecies, position53)
			}
			return true
		l52:
			position, tokenIndex = position52, tokenIndex52
			return false
		},
		/* 14 GenusWord <- <((AbbrGenus / UninomialWord) !(_ AuthorWord))> */
		func() bool {
			position62, tokenIndex62 := position, tokenIndex
			{
				position63 := position
				{
					position64, tokenIndex64 := position, tokenIndex
					if !_rules[ruleAbbrGenus]() {
						goto l65
					}
					goto l64
				l65:
					position, tokenIndex = position64, tokenIndex64
					if !_rules[ruleUninomialWord]() {
						goto l62
					}
				}
			l64:
				{
					position66, tokenIndex66 := position, tokenIndex
					if !_rules[rule_]() {
						goto l66
					}
					if !_rules[ruleAuthorWord]() {
						goto l66
					}
					goto l62
				l66:
					position, tokenIndex = position66, tokenIndex66
				}
				add(ruleGenusWord, position63)
			}
			return true
		l62:
			position, tokenIndex = position62, tokenIndex62
			return false
		},
		/* 15 InfraspGroup <- <(InfraspEpithet (_ InfraspEpithet)? (_ InfraspEpithet)?)> */
		func() bool {
			position67, tokenIndex67 := position, tokenIndex
			{
				position68 := position
				if !_rules[ruleInfraspEpithet]() {
					goto l67
				}
				{
					position69, tokenIndex69 := position, tokenIndex
					if !_rules[rule_]() {
						goto l69
					}
					if !_rules[ruleInfraspEpithet]() {
						goto l69
					}
					goto l70
				l69:
					position, tokenIndex = position69, tokenIndex69
				}
			l70:
				{
					position71, tokenIndex71 := position, tokenIndex
					if !_rules[rule_]() {
						goto l71
					}
					if !_rules[ruleInfraspEpithet]() {
						goto l71
					}
					goto l72
				l71:
					position, tokenIndex = position71, tokenIndex71
				}
			l72:
				add(ruleInfraspGroup, position68)
			}
			return true
		l67:
			position, tokenIndex = position67, tokenIndex67
			return false
		},
		/* 16 InfraspEpithet <- <((Rank _?)? !AuthorEx Word (_ Authorship)?)> */
		func() bool {
			position73, tokenIndex73 := position, tokenIndex
			{
				position74 := position
				{
					position75, tokenIndex75 := position, tokenIndex
					if !_rules[ruleRank]() {
						goto l75
					}
					{
						position77, tokenIndex77 := position, tokenIndex
						if !_rules[rule_]() {
							goto l77
						}
						goto l78
					l77:
						position, tokenIndex = position77, tokenIndex77
					}
				l78:
					goto l76
				l75:
					position, tokenIndex = position75, tokenIndex75
				}
			l76:
				{
					position79, tokenIndex79 := position, tokenIndex
					if !_rules[ruleAuthorEx]() {
						goto l79
					}
					goto l73
				l79:
					position, tokenIndex = position79, tokenIndex79
				}
				if !_rules[ruleWord]() {
					goto l73
				}
				{
					position80, tokenIndex80 := position, tokenIndex
					if !_rules[rule_]() {
						goto l80
					}
					if !_rules[ruleAuthorship]() {
						goto l80
					}
					goto l81
				l80:
					position, tokenIndex = position80, tokenIndex80
				}
			l81:
				add(ruleInfraspEpithet, position74)
			}
			return true
		l73:
			position, tokenIndex = position73, tokenIndex73
			return false
		},
		/* 17 SpeciesEpithet <- <(!AuthorEx Word (_? Authorship)? ','? &(SpaceCharEOI / '('))> */
		func() bool {
			position82, tokenIndex82 := position, tokenIndex
			{
				position83 := position
				{
					position84, tokenIndex84 := position, tokenIndex
					if !_rules[ruleAuthorEx]() {
						goto l84
					}
					goto l82
				l84:
					position, tokenIndex = position84, tokenIndex84
				}
				if !_rules[ruleWord]() {
					goto l82
				}
				{
					position85, tokenIndex85 := position, tokenIndex
					{
						position87, tokenIndex87 := position, tokenIndex
						if !_rules[rule_]() {
							goto l87
						}
						goto l88
					l87:
						position, tokenIndex = position87, tokenIndex87
					}
				l88:
					if !_rules[ruleAuthorship]() {
						goto l85
					}
					goto l86
				l85:
					position, tokenIndex = position85, tokenIndex85
				}
			l86:
				{
					position89, tokenIndex89 := position, tokenIndex
					if buffer[position] != rune(',') {
						goto l89
					}
					position++
					goto l90
				l89:
					position, tokenIndex = position89, tokenIndex89
				}
			l90:
				{
					position91, tokenIndex91 := position, tokenIndex
					{
						position92, tokenIndex92 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l93
						}
						goto l92
					l93:
						position, tokenIndex = position92, tokenIndex92
						if buffer[position] != rune('(') {
							goto l82
						}
						position++
					}
				l92:
					position, tokenIndex = position91, tokenIndex91
				}
				add(ruleSpeciesEpithet, position83)
			}
			return true
		l82:
			position, tokenIndex = position82, tokenIndex82
			return false
		},
		/* 18 Comparison <- <('c' 'f' '.'?)> */
		nil,
		/* 19 Rank <- <(RankForma / RankVar / RankSsp / RankOther / RankOtherUncommon)> */
		func() bool {
			position95, tokenIndex95 := position, tokenIndex
			{
				position96 := position
				{
					position97, tokenIndex97 := position, tokenIndex
					if !_rules[ruleRankForma]() {
						goto l98
					}
					goto l97
				l98:
					position, tokenIndex = position97, tokenIndex97
					if !_rules[ruleRankVar]() {
						goto l99
					}
					goto l97
				l99:
					position, tokenIndex = position97, tokenIndex97
					if !_rules[ruleRankSsp]() {
						goto l100
					}
					goto l97
				l100:
					position, tokenIndex = position97, tokenIndex97
					if !_rules[ruleRankOther]() {
						goto l101
					}
					goto l97
				l101:
					position, tokenIndex = position97, tokenIndex97
					if !_rules[ruleRankOtherUncommon]() {
						goto l95
					}
				}
			l97:
				add(ruleRank, position96)
			}
			return true
		l95:
			position, tokenIndex = position95, tokenIndex95
			return false
		},
		/* 20 RankOtherUncommon <- <(('*' / ('n' 'a' 't') / ('f' '.' 's' 'p') / ('m' 'u' 't' '.')) &SpaceCharEOI)> */
		func() bool {
			position102, tokenIndex102 := position, tokenIndex
			{
				position103 := position
				{
					position104, tokenIndex104 := position, tokenIndex
					if buffer[position] != rune('*') {
						goto l105
					}
					position++
					goto l104
				l105:
					position, tokenIndex = position104, tokenIndex104
					if buffer[position] != rune('n') {
						goto l106
					}
					position++
					if buffer[position] != rune('a') {
						goto l106
					}
					position++
					if buffer[position] != rune('t') {
						goto l106
					}
					position++
					goto l104
				l106:
					position, tokenIndex = position104, tokenIndex104
					if buffer[position] != rune('f') {
						goto l107
					}
					position++
					if buffer[position] != rune('.') {
						goto l107
					}
					position++
					if buffer[position] != rune('s') {
						goto l107
					}
					position++
					if buffer[position] != rune('p') {
						goto l107
					}
					position++
					goto l104
				l107:
					position, tokenIndex = position104, tokenIndex104
					if buffer[position] != rune('m') {
						goto l102
					}
					position++
					if buffer[position] != rune('u') {
						goto l102
					}
					position++
					if buffer[position] != rune('t') {
						goto l102
					}
					position++
					if buffer[position] != rune('.') {
						goto l102
					}
					position++
				}
			l104:
				{
					position108, tokenIndex108 := position, tokenIndex
					if !_rules[ruleSpaceCharEOI]() {
						goto l102
					}
					position, tokenIndex = position108, tokenIndex108
				}
				add(ruleRankOtherUncommon, position103)
			}
			return true
		l102:
			position, tokenIndex = position102, tokenIndex102
			return false
		},
		/* 21 RankOther <- <((('m' 'o' 'r' 'p' 'h' '.') / ('n' 'o' 't' 'h' 'o' 's' 'u' 'b' 's' 'p' '.') / ('c' 'o' 'n' 'v' 'a' 'r' '.') / ('p' 's' 'e' 'u' 'd' 'o' 'v' 'a' 'r' '.') / ('s' 'e' 'c' 't' '.') / ('s' 'e' 'r' '.') / ('s' 'u' 'b' 'v' 'a' 'r' '.') / ('s' 'u' 'b' 'f' '.') / ('r' 'a' 'c' 'e') / 'α' / ('β' 'β') / 'β' / 'γ' / 'δ' / 'ε' / 'φ' / 'θ' / 'μ' / ('a' '.') / ('b' '.') / ('c' '.') / ('d' '.') / ('e' '.') / ('g' '.') / ('k' '.') / ('p' 'v' '.') / ('p' 'a' 't' 'h' 'o' 'v' 'a' 'r' '.') / ('a' 'b' '.' (_? ('n' '.'))?) / ('s' 't' '.')) &SpaceCharEOI)> */
		func() bool {
			position109, tokenIndex109 := position, tokenIndex
			{
				position110 := position
				{
					position111, tokenIndex111 := position, tokenIndex
					if buffer[position] != rune('m') {
						goto l112
					}
					position++
					if buffer[position] != rune('o') {
						goto l112
					}
					position++
					if buffer[position] != rune('r') {
						goto l112
					}
					position++
					if buffer[position] != rune('p') {
						goto l112
					}
					position++
					if buffer[position] != rune('h') {
						goto l112
					}
					position++
					if buffer[position] != rune('.') {
						goto l112
					}
					position++
					goto l111
				l112:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('n') {
						goto l113
					}
					position++
					if buffer[position] != rune('o') {
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
					if buffer[position] != rune('s') {
						goto l113
					}
					position++
					if buffer[position] != rune('u') {
						goto l113
					}
					position++
					if buffer[position] != rune('b') {
						goto l113
					}
					position++
					if buffer[position] != rune('s') {
						goto l113
					}
					position++
					if buffer[position] != rune('p') {
						goto l113
					}
					position++
					if buffer[position] != rune('.') {
						goto l113
					}
					position++
					goto l111
				l113:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('c') {
						goto l114
					}
					position++
					if buffer[position] != rune('o') {
						goto l114
					}
					position++
					if buffer[position] != rune('n') {
						goto l114
					}
					position++
					if buffer[position] != rune('v') {
						goto l114
					}
					position++
					if buffer[position] != rune('a') {
						goto l114
					}
					position++
					if buffer[position] != rune('r') {
						goto l114
					}
					position++
					if buffer[position] != rune('.') {
						goto l114
					}
					position++
					goto l111
				l114:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('p') {
						goto l115
					}
					position++
					if buffer[position] != rune('s') {
						goto l115
					}
					position++
					if buffer[position] != rune('e') {
						goto l115
					}
					position++
					if buffer[position] != rune('u') {
						goto l115
					}
					position++
					if buffer[position] != rune('d') {
						goto l115
					}
					position++
					if buffer[position] != rune('o') {
						goto l115
					}
					position++
					if buffer[position] != rune('v') {
						goto l115
					}
					position++
					if buffer[position] != rune('a') {
						goto l115
					}
					position++
					if buffer[position] != rune('r') {
						goto l115
					}
					position++
					if buffer[position] != rune('.') {
						goto l115
					}
					position++
					goto l111
				l115:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('s') {
						goto l116
					}
					position++
					if buffer[position] != rune('e') {
						goto l116
					}
					position++
					if buffer[position] != rune('c') {
						goto l116
					}
					position++
					if buffer[position] != rune('t') {
						goto l116
					}
					position++
					if buffer[position] != rune('.') {
						goto l116
					}
					position++
					goto l111
				l116:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('s') {
						goto l117
					}
					position++
					if buffer[position] != rune('e') {
						goto l117
					}
					position++
					if buffer[position] != rune('r') {
						goto l117
					}
					position++
					if buffer[position] != rune('.') {
						goto l117
					}
					position++
					goto l111
				l117:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('s') {
						goto l118
					}
					position++
					if buffer[position] != rune('u') {
						goto l118
					}
					position++
					if buffer[position] != rune('b') {
						goto l118
					}
					position++
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
					if buffer[position] != rune('.') {
						goto l118
					}
					position++
					goto l111
				l118:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('s') {
						goto l119
					}
					position++
					if buffer[position] != rune('u') {
						goto l119
					}
					position++
					if buffer[position] != rune('b') {
						goto l119
					}
					position++
					if buffer[position] != rune('f') {
						goto l119
					}
					position++
					if buffer[position] != rune('.') {
						goto l119
					}
					position++
					goto l111
				l119:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('r') {
						goto l120
					}
					position++
					if buffer[position] != rune('a') {
						goto l120
					}
					position++
					if buffer[position] != rune('c') {
						goto l120
					}
					position++
					if buffer[position] != rune('e') {
						goto l120
					}
					position++
					goto l111
				l120:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('α') {
						goto l121
					}
					position++
					goto l111
				l121:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('β') {
						goto l122
					}
					position++
					if buffer[position] != rune('β') {
						goto l122
					}
					position++
					goto l111
				l122:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('β') {
						goto l123
					}
					position++
					goto l111
				l123:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('γ') {
						goto l124
					}
					position++
					goto l111
				l124:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('δ') {
						goto l125
					}
					position++
					goto l111
				l125:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('ε') {
						goto l126
					}
					position++
					goto l111
				l126:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('φ') {
						goto l127
					}
					position++
					goto l111
				l127:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('θ') {
						goto l128
					}
					position++
					goto l111
				l128:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('μ') {
						goto l129
					}
					position++
					goto l111
				l129:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('a') {
						goto l130
					}
					position++
					if buffer[position] != rune('.') {
						goto l130
					}
					position++
					goto l111
				l130:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('b') {
						goto l131
					}
					position++
					if buffer[position] != rune('.') {
						goto l131
					}
					position++
					goto l111
				l131:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('c') {
						goto l132
					}
					position++
					if buffer[position] != rune('.') {
						goto l132
					}
					position++
					goto l111
				l132:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('d') {
						goto l133
					}
					position++
					if buffer[position] != rune('.') {
						goto l133
					}
					position++
					goto l111
				l133:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('e') {
						goto l134
					}
					position++
					if buffer[position] != rune('.') {
						goto l134
					}
					position++
					goto l111
				l134:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('g') {
						goto l135
					}
					position++
					if buffer[position] != rune('.') {
						goto l135
					}
					position++
					goto l111
				l135:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('k') {
						goto l136
					}
					position++
					if buffer[position] != rune('.') {
						goto l136
					}
					position++
					goto l111
				l136:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('p') {
						goto l137
					}
					position++
					if buffer[position] != rune('v') {
						goto l137
					}
					position++
					if buffer[position] != rune('.') {
						goto l137
					}
					position++
					goto l111
				l137:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('p') {
						goto l138
					}
					position++
					if buffer[position] != rune('a') {
						goto l138
					}
					position++
					if buffer[position] != rune('t') {
						goto l138
					}
					position++
					if buffer[position] != rune('h') {
						goto l138
					}
					position++
					if buffer[position] != rune('o') {
						goto l138
					}
					position++
					if buffer[position] != rune('v') {
						goto l138
					}
					position++
					if buffer[position] != rune('a') {
						goto l138
					}
					position++
					if buffer[position] != rune('r') {
						goto l138
					}
					position++
					if buffer[position] != rune('.') {
						goto l138
					}
					position++
					goto l111
				l138:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('a') {
						goto l139
					}
					position++
					if buffer[position] != rune('b') {
						goto l139
					}
					position++
					if buffer[position] != rune('.') {
						goto l139
					}
					position++
					{
						position140, tokenIndex140 := position, tokenIndex
						{
							position142, tokenIndex142 := position, tokenIndex
							if !_rules[rule_]() {
								goto l142
							}
							goto l143
						l142:
							position, tokenIndex = position142, tokenIndex142
						}
					l143:
						if buffer[position] != rune('n') {
							goto l140
						}
						position++
						if buffer[position] != rune('.') {
							goto l140
						}
						position++
						goto l141
					l140:
						position, tokenIndex = position140, tokenIndex140
					}
				l141:
					goto l111
				l139:
					position, tokenIndex = position111, tokenIndex111
					if buffer[position] != rune('s') {
						goto l109
					}
					position++
					if buffer[position] != rune('t') {
						goto l109
					}
					position++
					if buffer[position] != rune('.') {
						goto l109
					}
					position++
				}
			l111:
				{
					position144, tokenIndex144 := position, tokenIndex
					if !_rules[ruleSpaceCharEOI]() {
						goto l109
					}
					position, tokenIndex = position144, tokenIndex144
				}
				add(ruleRankOther, position110)
			}
			return true
		l109:
			position, tokenIndex = position109, tokenIndex109
			return false
		},
		/* 22 RankVar <- <(('v' 'a' 'r' 'i' 'e' 't' 'y') / ('[' 'v' 'a' 'r' '.' ']') / ('n' 'v' 'a' 'r' '.') / ('v' 'a' 'r' (&SpaceCharEOI / '.')))> */
		func() bool {
			position145, tokenIndex145 := position, tokenIndex
			{
				position146 := position
				{
					position147, tokenIndex147 := position, tokenIndex
					if buffer[position] != rune('v') {
						goto l148
					}
					position++
					if buffer[position] != rune('a') {
						goto l148
					}
					position++
					if buffer[position] != rune('r') {
						goto l148
					}
					position++
					if buffer[position] != rune('i') {
						goto l148
					}
					position++
					if buffer[position] != rune('e') {
						goto l148
					}
					position++
					if buffer[position] != rune('t') {
						goto l148
					}
					position++
					if buffer[position] != rune('y') {
						goto l148
					}
					position++
					goto l147
				l148:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune('[') {
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
					if buffer[position] != rune(']') {
						goto l149
					}
					position++
					goto l147
				l149:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune('n') {
						goto l150
					}
					position++
					if buffer[position] != rune('v') {
						goto l150
					}
					position++
					if buffer[position] != rune('a') {
						goto l150
					}
					position++
					if buffer[position] != rune('r') {
						goto l150
					}
					position++
					if buffer[position] != rune('.') {
						goto l150
					}
					position++
					goto l147
				l150:
					position, tokenIndex = position147, tokenIndex147
					if buffer[position] != rune('v') {
						goto l145
					}
					position++
					if buffer[position] != rune('a') {
						goto l145
					}
					position++
					if buffer[position] != rune('r') {
						goto l145
					}
					position++
					{
						position151, tokenIndex151 := position, tokenIndex
						{
							position153, tokenIndex153 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l152
							}
							position, tokenIndex = position153, tokenIndex153
						}
						goto l151
					l152:
						position, tokenIndex = position151, tokenIndex151
						if buffer[position] != rune('.') {
							goto l145
						}
						position++
					}
				l151:
				}
			l147:
				add(ruleRankVar, position146)
			}
			return true
		l145:
			position, tokenIndex = position145, tokenIndex145
			return false
		},
		/* 23 RankForma <- <((('f' 'o' 'r' 'm' 'a') / ('f' 'm' 'a') / ('f' 'o' 'r' 'm') / ('f' 'o') / 'f') (&SpaceCharEOI / '.'))> */
		func() bool {
			position154, tokenIndex154 := position, tokenIndex
			{
				position155 := position
				{
					position156, tokenIndex156 := position, tokenIndex
					if buffer[position] != rune('f') {
						goto l157
					}
					position++
					if buffer[position] != rune('o') {
						goto l157
					}
					position++
					if buffer[position] != rune('r') {
						goto l157
					}
					position++
					if buffer[position] != rune('m') {
						goto l157
					}
					position++
					if buffer[position] != rune('a') {
						goto l157
					}
					position++
					goto l156
				l157:
					position, tokenIndex = position156, tokenIndex156
					if buffer[position] != rune('f') {
						goto l158
					}
					position++
					if buffer[position] != rune('m') {
						goto l158
					}
					position++
					if buffer[position] != rune('a') {
						goto l158
					}
					position++
					goto l156
				l158:
					position, tokenIndex = position156, tokenIndex156
					if buffer[position] != rune('f') {
						goto l159
					}
					position++
					if buffer[position] != rune('o') {
						goto l159
					}
					position++
					if buffer[position] != rune('r') {
						goto l159
					}
					position++
					if buffer[position] != rune('m') {
						goto l159
					}
					position++
					goto l156
				l159:
					position, tokenIndex = position156, tokenIndex156
					if buffer[position] != rune('f') {
						goto l160
					}
					position++
					if buffer[position] != rune('o') {
						goto l160
					}
					position++
					goto l156
				l160:
					position, tokenIndex = position156, tokenIndex156
					if buffer[position] != rune('f') {
						goto l154
					}
					position++
				}
			l156:
				{
					position161, tokenIndex161 := position, tokenIndex
					{
						position163, tokenIndex163 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l162
						}
						position, tokenIndex = position163, tokenIndex163
					}
					goto l161
				l162:
					position, tokenIndex = position161, tokenIndex161
					if buffer[position] != rune('.') {
						goto l154
					}
					position++
				}
			l161:
				add(ruleRankForma, position155)
			}
			return true
		l154:
			position, tokenIndex = position154, tokenIndex154
			return false
		},
		/* 24 RankSsp <- <((('s' 's' 'p') / ('s' 'u' 'b' 's' 'p')) (&SpaceCharEOI / '.'))> */
		func() bool {
			position164, tokenIndex164 := position, tokenIndex
			{
				position165 := position
				{
					position166, tokenIndex166 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l167
					}
					position++
					if buffer[position] != rune('s') {
						goto l167
					}
					position++
					if buffer[position] != rune('p') {
						goto l167
					}
					position++
					goto l166
				l167:
					position, tokenIndex = position166, tokenIndex166
					if buffer[position] != rune('s') {
						goto l164
					}
					position++
					if buffer[position] != rune('u') {
						goto l164
					}
					position++
					if buffer[position] != rune('b') {
						goto l164
					}
					position++
					if buffer[position] != rune('s') {
						goto l164
					}
					position++
					if buffer[position] != rune('p') {
						goto l164
					}
					position++
				}
			l166:
				{
					position168, tokenIndex168 := position, tokenIndex
					{
						position170, tokenIndex170 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l169
						}
						position, tokenIndex = position170, tokenIndex170
					}
					goto l168
				l169:
					position, tokenIndex = position168, tokenIndex168
					if buffer[position] != rune('.') {
						goto l164
					}
					position++
				}
			l168:
				add(ruleRankSsp, position165)
			}
			return true
		l164:
			position, tokenIndex = position164, tokenIndex164
			return false
		},
		/* 25 SubGenusOrSuperspecies <- <('(' _? NameLowerChar+ _? ')')> */
		func() bool {
			position171, tokenIndex171 := position, tokenIndex
			{
				position172 := position
				if buffer[position] != rune('(') {
					goto l171
				}
				position++
				{
					position173, tokenIndex173 := position, tokenIndex
					if !_rules[rule_]() {
						goto l173
					}
					goto l174
				l173:
					position, tokenIndex = position173, tokenIndex173
				}
			l174:
				if !_rules[ruleNameLowerChar]() {
					goto l171
				}
			l175:
				{
					position176, tokenIndex176 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l176
					}
					goto l175
				l176:
					position, tokenIndex = position176, tokenIndex176
				}
				{
					position177, tokenIndex177 := position, tokenIndex
					if !_rules[rule_]() {
						goto l177
					}
					goto l178
				l177:
					position, tokenIndex = position177, tokenIndex177
				}
			l178:
				if buffer[position] != rune(')') {
					goto l171
				}
				position++
				add(ruleSubGenusOrSuperspecies, position172)
			}
			return true
		l171:
			position, tokenIndex = position171, tokenIndex171
			return false
		},
		/* 26 SubGenus <- <('(' _? UninomialWord _? ')')> */
		func() bool {
			position179, tokenIndex179 := position, tokenIndex
			{
				position180 := position
				if buffer[position] != rune('(') {
					goto l179
				}
				position++
				{
					position181, tokenIndex181 := position, tokenIndex
					if !_rules[rule_]() {
						goto l181
					}
					goto l182
				l181:
					position, tokenIndex = position181, tokenIndex181
				}
			l182:
				if !_rules[ruleUninomialWord]() {
					goto l179
				}
				{
					position183, tokenIndex183 := position, tokenIndex
					if !_rules[rule_]() {
						goto l183
					}
					goto l184
				l183:
					position, tokenIndex = position183, tokenIndex183
				}
			l184:
				if buffer[position] != rune(')') {
					goto l179
				}
				position++
				add(ruleSubGenus, position180)
			}
			return true
		l179:
			position, tokenIndex = position179, tokenIndex179
			return false
		},
		/* 27 UninomialCombo <- <(UninomialCombo1 / UninomialCombo2)> */
		func() bool {
			position185, tokenIndex185 := position, tokenIndex
			{
				position186 := position
				{
					position187, tokenIndex187 := position, tokenIndex
					if !_rules[ruleUninomialCombo1]() {
						goto l188
					}
					goto l187
				l188:
					position, tokenIndex = position187, tokenIndex187
					if !_rules[ruleUninomialCombo2]() {
						goto l185
					}
				}
			l187:
				add(ruleUninomialCombo, position186)
			}
			return true
		l185:
			position, tokenIndex = position185, tokenIndex185
			return false
		},
		/* 28 UninomialCombo1 <- <(UninomialWord _? SubGenus _? Authorship .?)> */
		func() bool {
			position189, tokenIndex189 := position, tokenIndex
			{
				position190 := position
				if !_rules[ruleUninomialWord]() {
					goto l189
				}
				{
					position191, tokenIndex191 := position, tokenIndex
					if !_rules[rule_]() {
						goto l191
					}
					goto l192
				l191:
					position, tokenIndex = position191, tokenIndex191
				}
			l192:
				if !_rules[ruleSubGenus]() {
					goto l189
				}
				{
					position193, tokenIndex193 := position, tokenIndex
					if !_rules[rule_]() {
						goto l193
					}
					goto l194
				l193:
					position, tokenIndex = position193, tokenIndex193
				}
			l194:
				if !_rules[ruleAuthorship]() {
					goto l189
				}
				{
					position195, tokenIndex195 := position, tokenIndex
					if !matchDot() {
						goto l195
					}
					goto l196
				l195:
					position, tokenIndex = position195, tokenIndex195
				}
			l196:
				add(ruleUninomialCombo1, position190)
			}
			return true
		l189:
			position, tokenIndex = position189, tokenIndex189
			return false
		},
		/* 29 UninomialCombo2 <- <(Uninomial _? RankUninomial _? Uninomial)> */
		func() bool {
			position197, tokenIndex197 := position, tokenIndex
			{
				position198 := position
				if !_rules[ruleUninomial]() {
					goto l197
				}
				{
					position199, tokenIndex199 := position, tokenIndex
					if !_rules[rule_]() {
						goto l199
					}
					goto l200
				l199:
					position, tokenIndex = position199, tokenIndex199
				}
			l200:
				if !_rules[ruleRankUninomial]() {
					goto l197
				}
				{
					position201, tokenIndex201 := position, tokenIndex
					if !_rules[rule_]() {
						goto l201
					}
					goto l202
				l201:
					position, tokenIndex = position201, tokenIndex201
				}
			l202:
				if !_rules[ruleUninomial]() {
					goto l197
				}
				add(ruleUninomialCombo2, position198)
			}
			return true
		l197:
			position, tokenIndex = position197, tokenIndex197
			return false
		},
		/* 30 RankUninomial <- <((('s' 'e' 'c' 't') / ('s' 'u' 'b' 's' 'e' 'c' 't') / ('t' 'r' 'i' 'b') / ('s' 'u' 'b' 't' 'r' 'i' 'b') / ('s' 'u' 'b' 's' 'e' 'r') / ('s' 'e' 'r') / ('s' 'u' 'b' 'g' 'e' 'n') / ('f' 'a' 'm') / ('s' 'u' 'b' 'f' 'a' 'm') / ('s' 'u' 'p' 'e' 'r' 't' 'r' 'i' 'b')) '.'?)> */
		func() bool {
			position203, tokenIndex203 := position, tokenIndex
			{
				position204 := position
				{
					position205, tokenIndex205 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l206
					}
					position++
					if buffer[position] != rune('e') {
						goto l206
					}
					position++
					if buffer[position] != rune('c') {
						goto l206
					}
					position++
					if buffer[position] != rune('t') {
						goto l206
					}
					position++
					goto l205
				l206:
					position, tokenIndex = position205, tokenIndex205
					if buffer[position] != rune('s') {
						goto l207
					}
					position++
					if buffer[position] != rune('u') {
						goto l207
					}
					position++
					if buffer[position] != rune('b') {
						goto l207
					}
					position++
					if buffer[position] != rune('s') {
						goto l207
					}
					position++
					if buffer[position] != rune('e') {
						goto l207
					}
					position++
					if buffer[position] != rune('c') {
						goto l207
					}
					position++
					if buffer[position] != rune('t') {
						goto l207
					}
					position++
					goto l205
				l207:
					position, tokenIndex = position205, tokenIndex205
					if buffer[position] != rune('t') {
						goto l208
					}
					position++
					if buffer[position] != rune('r') {
						goto l208
					}
					position++
					if buffer[position] != rune('i') {
						goto l208
					}
					position++
					if buffer[position] != rune('b') {
						goto l208
					}
					position++
					goto l205
				l208:
					position, tokenIndex = position205, tokenIndex205
					if buffer[position] != rune('s') {
						goto l209
					}
					position++
					if buffer[position] != rune('u') {
						goto l209
					}
					position++
					if buffer[position] != rune('b') {
						goto l209
					}
					position++
					if buffer[position] != rune('t') {
						goto l209
					}
					position++
					if buffer[position] != rune('r') {
						goto l209
					}
					position++
					if buffer[position] != rune('i') {
						goto l209
					}
					position++
					if buffer[position] != rune('b') {
						goto l209
					}
					position++
					goto l205
				l209:
					position, tokenIndex = position205, tokenIndex205
					if buffer[position] != rune('s') {
						goto l210
					}
					position++
					if buffer[position] != rune('u') {
						goto l210
					}
					position++
					if buffer[position] != rune('b') {
						goto l210
					}
					position++
					if buffer[position] != rune('s') {
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
					goto l205
				l210:
					position, tokenIndex = position205, tokenIndex205
					if buffer[position] != rune('s') {
						goto l211
					}
					position++
					if buffer[position] != rune('e') {
						goto l211
					}
					position++
					if buffer[position] != rune('r') {
						goto l211
					}
					position++
					goto l205
				l211:
					position, tokenIndex = position205, tokenIndex205
					if buffer[position] != rune('s') {
						goto l212
					}
					position++
					if buffer[position] != rune('u') {
						goto l212
					}
					position++
					if buffer[position] != rune('b') {
						goto l212
					}
					position++
					if buffer[position] != rune('g') {
						goto l212
					}
					position++
					if buffer[position] != rune('e') {
						goto l212
					}
					position++
					if buffer[position] != rune('n') {
						goto l212
					}
					position++
					goto l205
				l212:
					position, tokenIndex = position205, tokenIndex205
					if buffer[position] != rune('f') {
						goto l213
					}
					position++
					if buffer[position] != rune('a') {
						goto l213
					}
					position++
					if buffer[position] != rune('m') {
						goto l213
					}
					position++
					goto l205
				l213:
					position, tokenIndex = position205, tokenIndex205
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
					if buffer[position] != rune('f') {
						goto l214
					}
					position++
					if buffer[position] != rune('a') {
						goto l214
					}
					position++
					if buffer[position] != rune('m') {
						goto l214
					}
					position++
					goto l205
				l214:
					position, tokenIndex = position205, tokenIndex205
					if buffer[position] != rune('s') {
						goto l203
					}
					position++
					if buffer[position] != rune('u') {
						goto l203
					}
					position++
					if buffer[position] != rune('p') {
						goto l203
					}
					position++
					if buffer[position] != rune('e') {
						goto l203
					}
					position++
					if buffer[position] != rune('r') {
						goto l203
					}
					position++
					if buffer[position] != rune('t') {
						goto l203
					}
					position++
					if buffer[position] != rune('r') {
						goto l203
					}
					position++
					if buffer[position] != rune('i') {
						goto l203
					}
					position++
					if buffer[position] != rune('b') {
						goto l203
					}
					position++
				}
			l205:
				{
					position215, tokenIndex215 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l215
					}
					position++
					goto l216
				l215:
					position, tokenIndex = position215, tokenIndex215
				}
			l216:
				add(ruleRankUninomial, position204)
			}
			return true
		l203:
			position, tokenIndex = position203, tokenIndex203
			return false
		},
		/* 31 Uninomial <- <(UninomialWord (_ Authorship)?)> */
		func() bool {
			position217, tokenIndex217 := position, tokenIndex
			{
				position218 := position
				if !_rules[ruleUninomialWord]() {
					goto l217
				}
				{
					position219, tokenIndex219 := position, tokenIndex
					if !_rules[rule_]() {
						goto l219
					}
					if !_rules[ruleAuthorship]() {
						goto l219
					}
					goto l220
				l219:
					position, tokenIndex = position219, tokenIndex219
				}
			l220:
				add(ruleUninomial, position218)
			}
			return true
		l217:
			position, tokenIndex = position217, tokenIndex217
			return false
		},
		/* 32 UninomialWord <- <(CapWord / TwoLetterGenus)> */
		func() bool {
			position221, tokenIndex221 := position, tokenIndex
			{
				position222 := position
				{
					position223, tokenIndex223 := position, tokenIndex
					if !_rules[ruleCapWord]() {
						goto l224
					}
					goto l223
				l224:
					position, tokenIndex = position223, tokenIndex223
					if !_rules[ruleTwoLetterGenus]() {
						goto l221
					}
				}
			l223:
				add(ruleUninomialWord, position222)
			}
			return true
		l221:
			position, tokenIndex = position221, tokenIndex221
			return false
		},
		/* 33 AbbrGenus <- <(UpperChar LowerChar* '.')> */
		func() bool {
			position225, tokenIndex225 := position, tokenIndex
			{
				position226 := position
				if !_rules[ruleUpperChar]() {
					goto l225
				}
			l227:
				{
					position228, tokenIndex228 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l228
					}
					goto l227
				l228:
					position, tokenIndex = position228, tokenIndex228
				}
				if buffer[position] != rune('.') {
					goto l225
				}
				position++
				add(ruleAbbrGenus, position226)
			}
			return true
		l225:
			position, tokenIndex = position225, tokenIndex225
			return false
		},
		/* 34 CapWord <- <(CapWord2 / CapWord1)> */
		func() bool {
			position229, tokenIndex229 := position, tokenIndex
			{
				position230 := position
				{
					position231, tokenIndex231 := position, tokenIndex
					if !_rules[ruleCapWord2]() {
						goto l232
					}
					goto l231
				l232:
					position, tokenIndex = position231, tokenIndex231
					if !_rules[ruleCapWord1]() {
						goto l229
					}
				}
			l231:
				add(ruleCapWord, position230)
			}
			return true
		l229:
			position, tokenIndex = position229, tokenIndex229
			return false
		},
		/* 35 CapWord1 <- <(NameUpperChar NameLowerChar NameLowerChar+ '?'?)> */
		func() bool {
			position233, tokenIndex233 := position, tokenIndex
			{
				position234 := position
				if !_rules[ruleNameUpperChar]() {
					goto l233
				}
				if !_rules[ruleNameLowerChar]() {
					goto l233
				}
				if !_rules[ruleNameLowerChar]() {
					goto l233
				}
			l235:
				{
					position236, tokenIndex236 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l236
					}
					goto l235
				l236:
					position, tokenIndex = position236, tokenIndex236
				}
				{
					position237, tokenIndex237 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l237
					}
					position++
					goto l238
				l237:
					position, tokenIndex = position237, tokenIndex237
				}
			l238:
				add(ruleCapWord1, position234)
			}
			return true
		l233:
			position, tokenIndex = position233, tokenIndex233
			return false
		},
		/* 36 CapWord2 <- <(CapWord1 dash (CapWord1 / Word1))> */
		func() bool {
			position239, tokenIndex239 := position, tokenIndex
			{
				position240 := position
				if !_rules[ruleCapWord1]() {
					goto l239
				}
				if !_rules[ruledash]() {
					goto l239
				}
				{
					position241, tokenIndex241 := position, tokenIndex
					if !_rules[ruleCapWord1]() {
						goto l242
					}
					goto l241
				l242:
					position, tokenIndex = position241, tokenIndex241
					if !_rules[ruleWord1]() {
						goto l239
					}
				}
			l241:
				add(ruleCapWord2, position240)
			}
			return true
		l239:
			position, tokenIndex = position239, tokenIndex239
			return false
		},
		/* 37 TwoLetterGenus <- <(('C' 'a') / ('E' 'a') / ('G' 'e') / ('I' 'a') / ('I' 'o') / ('I' 'x') / ('L' 'o') / ('O' 'a') / ('R' 'a') / ('T' 'y') / ('U' 'a') / ('A' 'a') / ('J' 'a') / ('Z' 'u') / ('L' 'a') / ('Q' 'u') / ('A' 's') / ('B' 'a'))> */
		func() bool {
			position243, tokenIndex243 := position, tokenIndex
			{
				position244 := position
				{
					position245, tokenIndex245 := position, tokenIndex
					if buffer[position] != rune('C') {
						goto l246
					}
					position++
					if buffer[position] != rune('a') {
						goto l246
					}
					position++
					goto l245
				l246:
					position, tokenIndex = position245, tokenIndex245
					if buffer[position] != rune('E') {
						goto l247
					}
					position++
					if buffer[position] != rune('a') {
						goto l247
					}
					position++
					goto l245
				l247:
					position, tokenIndex = position245, tokenIndex245
					if buffer[position] != rune('G') {
						goto l248
					}
					position++
					if buffer[position] != rune('e') {
						goto l248
					}
					position++
					goto l245
				l248:
					position, tokenIndex = position245, tokenIndex245
					if buffer[position] != rune('I') {
						goto l249
					}
					position++
					if buffer[position] != rune('a') {
						goto l249
					}
					position++
					goto l245
				l249:
					position, tokenIndex = position245, tokenIndex245
					if buffer[position] != rune('I') {
						goto l250
					}
					position++
					if buffer[position] != rune('o') {
						goto l250
					}
					position++
					goto l245
				l250:
					position, tokenIndex = position245, tokenIndex245
					if buffer[position] != rune('I') {
						goto l251
					}
					position++
					if buffer[position] != rune('x') {
						goto l251
					}
					position++
					goto l245
				l251:
					position, tokenIndex = position245, tokenIndex245
					if buffer[position] != rune('L') {
						goto l252
					}
					position++
					if buffer[position] != rune('o') {
						goto l252
					}
					position++
					goto l245
				l252:
					position, tokenIndex = position245, tokenIndex245
					if buffer[position] != rune('O') {
						goto l253
					}
					position++
					if buffer[position] != rune('a') {
						goto l253
					}
					position++
					goto l245
				l253:
					position, tokenIndex = position245, tokenIndex245
					if buffer[position] != rune('R') {
						goto l254
					}
					position++
					if buffer[position] != rune('a') {
						goto l254
					}
					position++
					goto l245
				l254:
					position, tokenIndex = position245, tokenIndex245
					if buffer[position] != rune('T') {
						goto l255
					}
					position++
					if buffer[position] != rune('y') {
						goto l255
					}
					position++
					goto l245
				l255:
					position, tokenIndex = position245, tokenIndex245
					if buffer[position] != rune('U') {
						goto l256
					}
					position++
					if buffer[position] != rune('a') {
						goto l256
					}
					position++
					goto l245
				l256:
					position, tokenIndex = position245, tokenIndex245
					if buffer[position] != rune('A') {
						goto l257
					}
					position++
					if buffer[position] != rune('a') {
						goto l257
					}
					position++
					goto l245
				l257:
					position, tokenIndex = position245, tokenIndex245
					if buffer[position] != rune('J') {
						goto l258
					}
					position++
					if buffer[position] != rune('a') {
						goto l258
					}
					position++
					goto l245
				l258:
					position, tokenIndex = position245, tokenIndex245
					if buffer[position] != rune('Z') {
						goto l259
					}
					position++
					if buffer[position] != rune('u') {
						goto l259
					}
					position++
					goto l245
				l259:
					position, tokenIndex = position245, tokenIndex245
					if buffer[position] != rune('L') {
						goto l260
					}
					position++
					if buffer[position] != rune('a') {
						goto l260
					}
					position++
					goto l245
				l260:
					position, tokenIndex = position245, tokenIndex245
					if buffer[position] != rune('Q') {
						goto l261
					}
					position++
					if buffer[position] != rune('u') {
						goto l261
					}
					position++
					goto l245
				l261:
					position, tokenIndex = position245, tokenIndex245
					if buffer[position] != rune('A') {
						goto l262
					}
					position++
					if buffer[position] != rune('s') {
						goto l262
					}
					position++
					goto l245
				l262:
					position, tokenIndex = position245, tokenIndex245
					if buffer[position] != rune('B') {
						goto l243
					}
					position++
					if buffer[position] != rune('a') {
						goto l243
					}
					position++
				}
			l245:
				add(ruleTwoLetterGenus, position244)
			}
			return true
		l243:
			position, tokenIndex = position243, tokenIndex243
			return false
		},
		/* 38 Word <- <(!(AuthorPrefix / RankUninomial / Approximation / Word4) (Word3 / Word2StartDigit / Word2 / Word1) &(SpaceCharEOI / '('))> */
		func() bool {
			position263, tokenIndex263 := position, tokenIndex
			{
				position264 := position
				{
					position265, tokenIndex265 := position, tokenIndex
					{
						position266, tokenIndex266 := position, tokenIndex
						if !_rules[ruleAuthorPrefix]() {
							goto l267
						}
						goto l266
					l267:
						position, tokenIndex = position266, tokenIndex266
						if !_rules[ruleRankUninomial]() {
							goto l268
						}
						goto l266
					l268:
						position, tokenIndex = position266, tokenIndex266
						if !_rules[ruleApproximation]() {
							goto l269
						}
						goto l266
					l269:
						position, tokenIndex = position266, tokenIndex266
						if !_rules[ruleWord4]() {
							goto l265
						}
					}
				l266:
					goto l263
				l265:
					position, tokenIndex = position265, tokenIndex265
				}
				{
					position270, tokenIndex270 := position, tokenIndex
					if !_rules[ruleWord3]() {
						goto l271
					}
					goto l270
				l271:
					position, tokenIndex = position270, tokenIndex270
					if !_rules[ruleWord2StartDigit]() {
						goto l272
					}
					goto l270
				l272:
					position, tokenIndex = position270, tokenIndex270
					if !_rules[ruleWord2]() {
						goto l273
					}
					goto l270
				l273:
					position, tokenIndex = position270, tokenIndex270
					if !_rules[ruleWord1]() {
						goto l263
					}
				}
			l270:
				{
					position274, tokenIndex274 := position, tokenIndex
					{
						position275, tokenIndex275 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l276
						}
						goto l275
					l276:
						position, tokenIndex = position275, tokenIndex275
						if buffer[position] != rune('(') {
							goto l263
						}
						position++
					}
				l275:
					position, tokenIndex = position274, tokenIndex274
				}
				add(ruleWord, position264)
			}
			return true
		l263:
			position, tokenIndex = position263, tokenIndex263
			return false
		},
		/* 39 Word1 <- <((lASCII dash)? NameLowerChar NameLowerChar+)> */
		func() bool {
			position277, tokenIndex277 := position, tokenIndex
			{
				position278 := position
				{
					position279, tokenIndex279 := position, tokenIndex
					if !_rules[rulelASCII]() {
						goto l279
					}
					if !_rules[ruledash]() {
						goto l279
					}
					goto l280
				l279:
					position, tokenIndex = position279, tokenIndex279
				}
			l280:
				if !_rules[ruleNameLowerChar]() {
					goto l277
				}
				if !_rules[ruleNameLowerChar]() {
					goto l277
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
				add(ruleWord1, position278)
			}
			return true
		l277:
			position, tokenIndex = position277, tokenIndex277
			return false
		},
		/* 40 Word2StartDigit <- <(('1' / '2' / '3' / '4' / '5' / '6' / '7' / '8' / '9') nums? ('.' / dash)? NameLowerChar NameLowerChar NameLowerChar NameLowerChar+)> */
		func() bool {
			position283, tokenIndex283 := position, tokenIndex
			{
				position284 := position
				{
					position285, tokenIndex285 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l286
					}
					position++
					goto l285
				l286:
					position, tokenIndex = position285, tokenIndex285
					if buffer[position] != rune('2') {
						goto l287
					}
					position++
					goto l285
				l287:
					position, tokenIndex = position285, tokenIndex285
					if buffer[position] != rune('3') {
						goto l288
					}
					position++
					goto l285
				l288:
					position, tokenIndex = position285, tokenIndex285
					if buffer[position] != rune('4') {
						goto l289
					}
					position++
					goto l285
				l289:
					position, tokenIndex = position285, tokenIndex285
					if buffer[position] != rune('5') {
						goto l290
					}
					position++
					goto l285
				l290:
					position, tokenIndex = position285, tokenIndex285
					if buffer[position] != rune('6') {
						goto l291
					}
					position++
					goto l285
				l291:
					position, tokenIndex = position285, tokenIndex285
					if buffer[position] != rune('7') {
						goto l292
					}
					position++
					goto l285
				l292:
					position, tokenIndex = position285, tokenIndex285
					if buffer[position] != rune('8') {
						goto l293
					}
					position++
					goto l285
				l293:
					position, tokenIndex = position285, tokenIndex285
					if buffer[position] != rune('9') {
						goto l283
					}
					position++
				}
			l285:
				{
					position294, tokenIndex294 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l294
					}
					goto l295
				l294:
					position, tokenIndex = position294, tokenIndex294
				}
			l295:
				{
					position296, tokenIndex296 := position, tokenIndex
					{
						position298, tokenIndex298 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l299
						}
						position++
						goto l298
					l299:
						position, tokenIndex = position298, tokenIndex298
						if !_rules[ruledash]() {
							goto l296
						}
					}
				l298:
					goto l297
				l296:
					position, tokenIndex = position296, tokenIndex296
				}
			l297:
				if !_rules[ruleNameLowerChar]() {
					goto l283
				}
				if !_rules[ruleNameLowerChar]() {
					goto l283
				}
				if !_rules[ruleNameLowerChar]() {
					goto l283
				}
				if !_rules[ruleNameLowerChar]() {
					goto l283
				}
			l300:
				{
					position301, tokenIndex301 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l301
					}
					goto l300
				l301:
					position, tokenIndex = position301, tokenIndex301
				}
				add(ruleWord2StartDigit, position284)
			}
			return true
		l283:
			position, tokenIndex = position283, tokenIndex283
			return false
		},
		/* 41 Word2 <- <(NameLowerChar+ dash? NameLowerChar+)> */
		func() bool {
			position302, tokenIndex302 := position, tokenIndex
			{
				position303 := position
				if !_rules[ruleNameLowerChar]() {
					goto l302
				}
			l304:
				{
					position305, tokenIndex305 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l305
					}
					goto l304
				l305:
					position, tokenIndex = position305, tokenIndex305
				}
				{
					position306, tokenIndex306 := position, tokenIndex
					if !_rules[ruledash]() {
						goto l306
					}
					goto l307
				l306:
					position, tokenIndex = position306, tokenIndex306
				}
			l307:
				if !_rules[ruleNameLowerChar]() {
					goto l302
				}
			l308:
				{
					position309, tokenIndex309 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l309
					}
					goto l308
				l309:
					position, tokenIndex = position309, tokenIndex309
				}
				add(ruleWord2, position303)
			}
			return true
		l302:
			position, tokenIndex = position302, tokenIndex302
			return false
		},
		/* 42 Word3 <- <(NameLowerChar NameLowerChar* apostr Word1)> */
		func() bool {
			position310, tokenIndex310 := position, tokenIndex
			{
				position311 := position
				if !_rules[ruleNameLowerChar]() {
					goto l310
				}
			l312:
				{
					position313, tokenIndex313 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l313
					}
					goto l312
				l313:
					position, tokenIndex = position313, tokenIndex313
				}
				if !_rules[ruleapostr]() {
					goto l310
				}
				if !_rules[ruleWord1]() {
					goto l310
				}
				add(ruleWord3, position311)
			}
			return true
		l310:
			position, tokenIndex = position310, tokenIndex310
			return false
		},
		/* 43 Word4 <- <(NameLowerChar+ '.' NameLowerChar)> */
		func() bool {
			position314, tokenIndex314 := position, tokenIndex
			{
				position315 := position
				if !_rules[ruleNameLowerChar]() {
					goto l314
				}
			l316:
				{
					position317, tokenIndex317 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l317
					}
					goto l316
				l317:
					position, tokenIndex = position317, tokenIndex317
				}
				if buffer[position] != rune('.') {
					goto l314
				}
				position++
				if !_rules[ruleNameLowerChar]() {
					goto l314
				}
				add(ruleWord4, position315)
			}
			return true
		l314:
			position, tokenIndex = position314, tokenIndex314
			return false
		},
		/* 44 HybridChar <- <'×'> */
		func() bool {
			position318, tokenIndex318 := position, tokenIndex
			{
				position319 := position
				if buffer[position] != rune('×') {
					goto l318
				}
				position++
				add(ruleHybridChar, position319)
			}
			return true
		l318:
			position, tokenIndex = position318, tokenIndex318
			return false
		},
		/* 45 ApproxName <- <(Uninomial _ (ApproxName1 / ApproxName2))> */
		nil,
		/* 46 ApproxName1 <- <(Approximation ApproxNameIgnored)> */
		nil,
		/* 47 ApproxName2 <- <(Word _ Approximation ApproxNameIgnored)> */
		nil,
		/* 48 ApproxNameIgnored <- <.*> */
		nil,
		/* 49 Approximation <- <(('s' 'p' '.' _? ('n' 'r' '.')) / ('s' 'p' '.' _? ('a' 'f' 'f' '.')) / ('m' 'o' 'n' 's' 't' '.') / '?' / ((('s' 'p' 'p') / ('n' 'r') / ('s' 'p') / ('a' 'f' 'f') / ('s' 'p' 'e' 'c' 'i' 'e' 's')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position324, tokenIndex324 := position, tokenIndex
			{
				position325 := position
				{
					position326, tokenIndex326 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l327
					}
					position++
					if buffer[position] != rune('p') {
						goto l327
					}
					position++
					if buffer[position] != rune('.') {
						goto l327
					}
					position++
					{
						position328, tokenIndex328 := position, tokenIndex
						if !_rules[rule_]() {
							goto l328
						}
						goto l329
					l328:
						position, tokenIndex = position328, tokenIndex328
					}
				l329:
					if buffer[position] != rune('n') {
						goto l327
					}
					position++
					if buffer[position] != rune('r') {
						goto l327
					}
					position++
					if buffer[position] != rune('.') {
						goto l327
					}
					position++
					goto l326
				l327:
					position, tokenIndex = position326, tokenIndex326
					if buffer[position] != rune('s') {
						goto l330
					}
					position++
					if buffer[position] != rune('p') {
						goto l330
					}
					position++
					if buffer[position] != rune('.') {
						goto l330
					}
					position++
					{
						position331, tokenIndex331 := position, tokenIndex
						if !_rules[rule_]() {
							goto l331
						}
						goto l332
					l331:
						position, tokenIndex = position331, tokenIndex331
					}
				l332:
					if buffer[position] != rune('a') {
						goto l330
					}
					position++
					if buffer[position] != rune('f') {
						goto l330
					}
					position++
					if buffer[position] != rune('f') {
						goto l330
					}
					position++
					if buffer[position] != rune('.') {
						goto l330
					}
					position++
					goto l326
				l330:
					position, tokenIndex = position326, tokenIndex326
					if buffer[position] != rune('m') {
						goto l333
					}
					position++
					if buffer[position] != rune('o') {
						goto l333
					}
					position++
					if buffer[position] != rune('n') {
						goto l333
					}
					position++
					if buffer[position] != rune('s') {
						goto l333
					}
					position++
					if buffer[position] != rune('t') {
						goto l333
					}
					position++
					if buffer[position] != rune('.') {
						goto l333
					}
					position++
					goto l326
				l333:
					position, tokenIndex = position326, tokenIndex326
					if buffer[position] != rune('?') {
						goto l334
					}
					position++
					goto l326
				l334:
					position, tokenIndex = position326, tokenIndex326
					{
						position335, tokenIndex335 := position, tokenIndex
						if buffer[position] != rune('s') {
							goto l336
						}
						position++
						if buffer[position] != rune('p') {
							goto l336
						}
						position++
						if buffer[position] != rune('p') {
							goto l336
						}
						position++
						goto l335
					l336:
						position, tokenIndex = position335, tokenIndex335
						if buffer[position] != rune('n') {
							goto l337
						}
						position++
						if buffer[position] != rune('r') {
							goto l337
						}
						position++
						goto l335
					l337:
						position, tokenIndex = position335, tokenIndex335
						if buffer[position] != rune('s') {
							goto l338
						}
						position++
						if buffer[position] != rune('p') {
							goto l338
						}
						position++
						goto l335
					l338:
						position, tokenIndex = position335, tokenIndex335
						if buffer[position] != rune('a') {
							goto l339
						}
						position++
						if buffer[position] != rune('f') {
							goto l339
						}
						position++
						if buffer[position] != rune('f') {
							goto l339
						}
						position++
						goto l335
					l339:
						position, tokenIndex = position335, tokenIndex335
						if buffer[position] != rune('s') {
							goto l324
						}
						position++
						if buffer[position] != rune('p') {
							goto l324
						}
						position++
						if buffer[position] != rune('e') {
							goto l324
						}
						position++
						if buffer[position] != rune('c') {
							goto l324
						}
						position++
						if buffer[position] != rune('i') {
							goto l324
						}
						position++
						if buffer[position] != rune('e') {
							goto l324
						}
						position++
						if buffer[position] != rune('s') {
							goto l324
						}
						position++
					}
				l335:
					{
						position340, tokenIndex340 := position, tokenIndex
						{
							position342, tokenIndex342 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l341
							}
							position, tokenIndex = position342, tokenIndex342
						}
						goto l340
					l341:
						position, tokenIndex = position340, tokenIndex340
						if buffer[position] != rune('.') {
							goto l324
						}
						position++
					}
				l340:
				}
			l326:
				add(ruleApproximation, position325)
			}
			return true
		l324:
			position, tokenIndex = position324, tokenIndex324
			return false
		},
		/* 50 Authorship <- <((AuthorshipCombo / OriginalAuthorship) &(SpaceCharEOI / ('\\' / '(' / ',' / ':')))> */
		func() bool {
			position343, tokenIndex343 := position, tokenIndex
			{
				position344 := position
				{
					position345, tokenIndex345 := position, tokenIndex
					if !_rules[ruleAuthorshipCombo]() {
						goto l346
					}
					goto l345
				l346:
					position, tokenIndex = position345, tokenIndex345
					if !_rules[ruleOriginalAuthorship]() {
						goto l343
					}
				}
			l345:
				{
					position347, tokenIndex347 := position, tokenIndex
					{
						position348, tokenIndex348 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l349
						}
						goto l348
					l349:
						position, tokenIndex = position348, tokenIndex348
						{
							position350, tokenIndex350 := position, tokenIndex
							if buffer[position] != rune('\\') {
								goto l351
							}
							position++
							goto l350
						l351:
							position, tokenIndex = position350, tokenIndex350
							if buffer[position] != rune('(') {
								goto l352
							}
							position++
							goto l350
						l352:
							position, tokenIndex = position350, tokenIndex350
							if buffer[position] != rune(',') {
								goto l353
							}
							position++
							goto l350
						l353:
							position, tokenIndex = position350, tokenIndex350
							if buffer[position] != rune(':') {
								goto l343
							}
							position++
						}
					l350:
					}
				l348:
					position, tokenIndex = position347, tokenIndex347
				}
				add(ruleAuthorship, position344)
			}
			return true
		l343:
			position, tokenIndex = position343, tokenIndex343
			return false
		},
		/* 51 AuthorshipCombo <- <(OriginalAuthorship _? CombinationAuthorship)> */
		func() bool {
			position354, tokenIndex354 := position, tokenIndex
			{
				position355 := position
				if !_rules[ruleOriginalAuthorship]() {
					goto l354
				}
				{
					position356, tokenIndex356 := position, tokenIndex
					if !_rules[rule_]() {
						goto l356
					}
					goto l357
				l356:
					position, tokenIndex = position356, tokenIndex356
				}
			l357:
				if !_rules[ruleCombinationAuthorship]() {
					goto l354
				}
				add(ruleAuthorshipCombo, position355)
			}
			return true
		l354:
			position, tokenIndex = position354, tokenIndex354
			return false
		},
		/* 52 OriginalAuthorship <- <(BasionymAuthorshipYearMisformed / AuthorsGroup / BasionymAuthorship)> */
		func() bool {
			position358, tokenIndex358 := position, tokenIndex
			{
				position359 := position
				{
					position360, tokenIndex360 := position, tokenIndex
					if !_rules[ruleBasionymAuthorshipYearMisformed]() {
						goto l361
					}
					goto l360
				l361:
					position, tokenIndex = position360, tokenIndex360
					if !_rules[ruleAuthorsGroup]() {
						goto l362
					}
					goto l360
				l362:
					position, tokenIndex = position360, tokenIndex360
					if !_rules[ruleBasionymAuthorship]() {
						goto l358
					}
				}
			l360:
				add(ruleOriginalAuthorship, position359)
			}
			return true
		l358:
			position, tokenIndex = position358, tokenIndex358
			return false
		},
		/* 53 CombinationAuthorship <- <AuthorsGroup> */
		func() bool {
			position363, tokenIndex363 := position, tokenIndex
			{
				position364 := position
				if !_rules[ruleAuthorsGroup]() {
					goto l363
				}
				add(ruleCombinationAuthorship, position364)
			}
			return true
		l363:
			position, tokenIndex = position363, tokenIndex363
			return false
		},
		/* 54 BasionymAuthorshipYearMisformed <- <('(' _? AuthorsGroup _? ')' (_? ',')? _? Year)> */
		func() bool {
			position365, tokenIndex365 := position, tokenIndex
			{
				position366 := position
				if buffer[position] != rune('(') {
					goto l365
				}
				position++
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
				if !_rules[ruleAuthorsGroup]() {
					goto l365
				}
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
				if buffer[position] != rune(')') {
					goto l365
				}
				position++
				{
					position371, tokenIndex371 := position, tokenIndex
					{
						position373, tokenIndex373 := position, tokenIndex
						if !_rules[rule_]() {
							goto l373
						}
						goto l374
					l373:
						position, tokenIndex = position373, tokenIndex373
					}
				l374:
					if buffer[position] != rune(',') {
						goto l371
					}
					position++
					goto l372
				l371:
					position, tokenIndex = position371, tokenIndex371
				}
			l372:
				{
					position375, tokenIndex375 := position, tokenIndex
					if !_rules[rule_]() {
						goto l375
					}
					goto l376
				l375:
					position, tokenIndex = position375, tokenIndex375
				}
			l376:
				if !_rules[ruleYear]() {
					goto l365
				}
				add(ruleBasionymAuthorshipYearMisformed, position366)
			}
			return true
		l365:
			position, tokenIndex = position365, tokenIndex365
			return false
		},
		/* 55 BasionymAuthorship <- <(BasionymAuthorship1 / BasionymAuthorship2)> */
		func() bool {
			position377, tokenIndex377 := position, tokenIndex
			{
				position378 := position
				{
					position379, tokenIndex379 := position, tokenIndex
					if !_rules[ruleBasionymAuthorship1]() {
						goto l380
					}
					goto l379
				l380:
					position, tokenIndex = position379, tokenIndex379
					if !_rules[ruleBasionymAuthorship2]() {
						goto l377
					}
				}
			l379:
				add(ruleBasionymAuthorship, position378)
			}
			return true
		l377:
			position, tokenIndex = position377, tokenIndex377
			return false
		},
		/* 56 BasionymAuthorship1 <- <('(' _? AuthorsGroup _? ')')> */
		func() bool {
			position381, tokenIndex381 := position, tokenIndex
			{
				position382 := position
				if buffer[position] != rune('(') {
					goto l381
				}
				position++
				{
					position383, tokenIndex383 := position, tokenIndex
					if !_rules[rule_]() {
						goto l383
					}
					goto l384
				l383:
					position, tokenIndex = position383, tokenIndex383
				}
			l384:
				if !_rules[ruleAuthorsGroup]() {
					goto l381
				}
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
				if buffer[position] != rune(')') {
					goto l381
				}
				position++
				add(ruleBasionymAuthorship1, position382)
			}
			return true
		l381:
			position, tokenIndex = position381, tokenIndex381
			return false
		},
		/* 57 BasionymAuthorship2 <- <('(' _? '(' _? AuthorsGroup _? ')' _? ')')> */
		func() bool {
			position387, tokenIndex387 := position, tokenIndex
			{
				position388 := position
				if buffer[position] != rune('(') {
					goto l387
				}
				position++
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
				if buffer[position] != rune('(') {
					goto l387
				}
				position++
				{
					position391, tokenIndex391 := position, tokenIndex
					if !_rules[rule_]() {
						goto l391
					}
					goto l392
				l391:
					position, tokenIndex = position391, tokenIndex391
				}
			l392:
				if !_rules[ruleAuthorsGroup]() {
					goto l387
				}
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
				if buffer[position] != rune(')') {
					goto l387
				}
				position++
				{
					position395, tokenIndex395 := position, tokenIndex
					if !_rules[rule_]() {
						goto l395
					}
					goto l396
				l395:
					position, tokenIndex = position395, tokenIndex395
				}
			l396:
				if buffer[position] != rune(')') {
					goto l387
				}
				position++
				add(ruleBasionymAuthorship2, position388)
			}
			return true
		l387:
			position, tokenIndex = position387, tokenIndex387
			return false
		},
		/* 58 AuthorsGroup <- <(AuthorsTeam (_? AuthorEmend? AuthorEx? AuthorsTeam)?)> */
		func() bool {
			position397, tokenIndex397 := position, tokenIndex
			{
				position398 := position
				if !_rules[ruleAuthorsTeam]() {
					goto l397
				}
				{
					position399, tokenIndex399 := position, tokenIndex
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
					{
						position403, tokenIndex403 := position, tokenIndex
						if !_rules[ruleAuthorEmend]() {
							goto l403
						}
						goto l404
					l403:
						position, tokenIndex = position403, tokenIndex403
					}
				l404:
					{
						position405, tokenIndex405 := position, tokenIndex
						if !_rules[ruleAuthorEx]() {
							goto l405
						}
						goto l406
					l405:
						position, tokenIndex = position405, tokenIndex405
					}
				l406:
					if !_rules[ruleAuthorsTeam]() {
						goto l399
					}
					goto l400
				l399:
					position, tokenIndex = position399, tokenIndex399
				}
			l400:
				add(ruleAuthorsGroup, position398)
			}
			return true
		l397:
			position, tokenIndex = position397, tokenIndex397
			return false
		},
		/* 59 AuthorsTeam <- <(Author (AuthorSep Author)* (_? ','? _? Year)?)> */
		func() bool {
			position407, tokenIndex407 := position, tokenIndex
			{
				position408 := position
				if !_rules[ruleAuthor]() {
					goto l407
				}
			l409:
				{
					position410, tokenIndex410 := position, tokenIndex
					if !_rules[ruleAuthorSep]() {
						goto l410
					}
					if !_rules[ruleAuthor]() {
						goto l410
					}
					goto l409
				l410:
					position, tokenIndex = position410, tokenIndex410
				}
				{
					position411, tokenIndex411 := position, tokenIndex
					{
						position413, tokenIndex413 := position, tokenIndex
						if !_rules[rule_]() {
							goto l413
						}
						goto l414
					l413:
						position, tokenIndex = position413, tokenIndex413
					}
				l414:
					{
						position415, tokenIndex415 := position, tokenIndex
						if buffer[position] != rune(',') {
							goto l415
						}
						position++
						goto l416
					l415:
						position, tokenIndex = position415, tokenIndex415
					}
				l416:
					{
						position417, tokenIndex417 := position, tokenIndex
						if !_rules[rule_]() {
							goto l417
						}
						goto l418
					l417:
						position, tokenIndex = position417, tokenIndex417
					}
				l418:
					if !_rules[ruleYear]() {
						goto l411
					}
					goto l412
				l411:
					position, tokenIndex = position411, tokenIndex411
				}
			l412:
				add(ruleAuthorsTeam, position408)
			}
			return true
		l407:
			position, tokenIndex = position407, tokenIndex407
			return false
		},
		/* 60 AuthorSep <- <(AuthorSep1 / AuthorSep2)> */
		func() bool {
			position419, tokenIndex419 := position, tokenIndex
			{
				position420 := position
				{
					position421, tokenIndex421 := position, tokenIndex
					if !_rules[ruleAuthorSep1]() {
						goto l422
					}
					goto l421
				l422:
					position, tokenIndex = position421, tokenIndex421
					if !_rules[ruleAuthorSep2]() {
						goto l419
					}
				}
			l421:
				add(ruleAuthorSep, position420)
			}
			return true
		l419:
			position, tokenIndex = position419, tokenIndex419
			return false
		},
		/* 61 AuthorSep1 <- <(_? (',' _)? ('&' / ('e' 't') / ('a' 'n' 'd') / ('a' 'p' 'u' 'd')) _?)> */
		func() bool {
			position423, tokenIndex423 := position, tokenIndex
			{
				position424 := position
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
				{
					position427, tokenIndex427 := position, tokenIndex
					if buffer[position] != rune(',') {
						goto l427
					}
					position++
					if !_rules[rule_]() {
						goto l427
					}
					goto l428
				l427:
					position, tokenIndex = position427, tokenIndex427
				}
			l428:
				{
					position429, tokenIndex429 := position, tokenIndex
					if buffer[position] != rune('&') {
						goto l430
					}
					position++
					goto l429
				l430:
					position, tokenIndex = position429, tokenIndex429
					if buffer[position] != rune('e') {
						goto l431
					}
					position++
					if buffer[position] != rune('t') {
						goto l431
					}
					position++
					goto l429
				l431:
					position, tokenIndex = position429, tokenIndex429
					if buffer[position] != rune('a') {
						goto l432
					}
					position++
					if buffer[position] != rune('n') {
						goto l432
					}
					position++
					if buffer[position] != rune('d') {
						goto l432
					}
					position++
					goto l429
				l432:
					position, tokenIndex = position429, tokenIndex429
					if buffer[position] != rune('a') {
						goto l423
					}
					position++
					if buffer[position] != rune('p') {
						goto l423
					}
					position++
					if buffer[position] != rune('u') {
						goto l423
					}
					position++
					if buffer[position] != rune('d') {
						goto l423
					}
					position++
				}
			l429:
				{
					position433, tokenIndex433 := position, tokenIndex
					if !_rules[rule_]() {
						goto l433
					}
					goto l434
				l433:
					position, tokenIndex = position433, tokenIndex433
				}
			l434:
				add(ruleAuthorSep1, position424)
			}
			return true
		l423:
			position, tokenIndex = position423, tokenIndex423
			return false
		},
		/* 62 AuthorSep2 <- <(_? ',' _?)> */
		func() bool {
			position435, tokenIndex435 := position, tokenIndex
			{
				position436 := position
				{
					position437, tokenIndex437 := position, tokenIndex
					if !_rules[rule_]() {
						goto l437
					}
					goto l438
				l437:
					position, tokenIndex = position437, tokenIndex437
				}
			l438:
				if buffer[position] != rune(',') {
					goto l435
				}
				position++
				{
					position439, tokenIndex439 := position, tokenIndex
					if !_rules[rule_]() {
						goto l439
					}
					goto l440
				l439:
					position, tokenIndex = position439, tokenIndex439
				}
			l440:
				add(ruleAuthorSep2, position436)
			}
			return true
		l435:
			position, tokenIndex = position435, tokenIndex435
			return false
		},
		/* 63 AuthorEx <- <((('e' 'x' '.'?) / ('i' 'n')) _)> */
		func() bool {
			position441, tokenIndex441 := position, tokenIndex
			{
				position442 := position
				{
					position443, tokenIndex443 := position, tokenIndex
					if buffer[position] != rune('e') {
						goto l444
					}
					position++
					if buffer[position] != rune('x') {
						goto l444
					}
					position++
					{
						position445, tokenIndex445 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l445
						}
						position++
						goto l446
					l445:
						position, tokenIndex = position445, tokenIndex445
					}
				l446:
					goto l443
				l444:
					position, tokenIndex = position443, tokenIndex443
					if buffer[position] != rune('i') {
						goto l441
					}
					position++
					if buffer[position] != rune('n') {
						goto l441
					}
					position++
				}
			l443:
				if !_rules[rule_]() {
					goto l441
				}
				add(ruleAuthorEx, position442)
			}
			return true
		l441:
			position, tokenIndex = position441, tokenIndex441
			return false
		},
		/* 64 AuthorEmend <- <('e' 'm' 'e' 'n' 'd' '.'? _)> */
		func() bool {
			position447, tokenIndex447 := position, tokenIndex
			{
				position448 := position
				if buffer[position] != rune('e') {
					goto l447
				}
				position++
				if buffer[position] != rune('m') {
					goto l447
				}
				position++
				if buffer[position] != rune('e') {
					goto l447
				}
				position++
				if buffer[position] != rune('n') {
					goto l447
				}
				position++
				if buffer[position] != rune('d') {
					goto l447
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
				if !_rules[rule_]() {
					goto l447
				}
				add(ruleAuthorEmend, position448)
			}
			return true
		l447:
			position, tokenIndex = position447, tokenIndex447
			return false
		},
		/* 65 Author <- <(Author1 / Author2 / UnknownAuthor)> */
		func() bool {
			position451, tokenIndex451 := position, tokenIndex
			{
				position452 := position
				{
					position453, tokenIndex453 := position, tokenIndex
					if !_rules[ruleAuthor1]() {
						goto l454
					}
					goto l453
				l454:
					position, tokenIndex = position453, tokenIndex453
					if !_rules[ruleAuthor2]() {
						goto l455
					}
					goto l453
				l455:
					position, tokenIndex = position453, tokenIndex453
					if !_rules[ruleUnknownAuthor]() {
						goto l451
					}
				}
			l453:
				add(ruleAuthor, position452)
			}
			return true
		l451:
			position, tokenIndex = position451, tokenIndex451
			return false
		},
		/* 66 Author1 <- <(Author2 _? Filius)> */
		func() bool {
			position456, tokenIndex456 := position, tokenIndex
			{
				position457 := position
				if !_rules[ruleAuthor2]() {
					goto l456
				}
				{
					position458, tokenIndex458 := position, tokenIndex
					if !_rules[rule_]() {
						goto l458
					}
					goto l459
				l458:
					position, tokenIndex = position458, tokenIndex458
				}
			l459:
				if !_rules[ruleFilius]() {
					goto l456
				}
				add(ruleAuthor1, position457)
			}
			return true
		l456:
			position, tokenIndex = position456, tokenIndex456
			return false
		},
		/* 67 Author2 <- <(AuthorWord (_? AuthorWord)*)> */
		func() bool {
			position460, tokenIndex460 := position, tokenIndex
			{
				position461 := position
				if !_rules[ruleAuthorWord]() {
					goto l460
				}
			l462:
				{
					position463, tokenIndex463 := position, tokenIndex
					{
						position464, tokenIndex464 := position, tokenIndex
						if !_rules[rule_]() {
							goto l464
						}
						goto l465
					l464:
						position, tokenIndex = position464, tokenIndex464
					}
				l465:
					if !_rules[ruleAuthorWord]() {
						goto l463
					}
					goto l462
				l463:
					position, tokenIndex = position463, tokenIndex463
				}
				add(ruleAuthor2, position461)
			}
			return true
		l460:
			position, tokenIndex = position460, tokenIndex460
			return false
		},
		/* 68 UnknownAuthor <- <('?' / ((('a' 'u' 'c' 't') / ('a' 'n' 'o' 'n')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position466, tokenIndex466 := position, tokenIndex
			{
				position467 := position
				{
					position468, tokenIndex468 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l469
					}
					position++
					goto l468
				l469:
					position, tokenIndex = position468, tokenIndex468
					{
						position470, tokenIndex470 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l471
						}
						position++
						if buffer[position] != rune('u') {
							goto l471
						}
						position++
						if buffer[position] != rune('c') {
							goto l471
						}
						position++
						if buffer[position] != rune('t') {
							goto l471
						}
						position++
						goto l470
					l471:
						position, tokenIndex = position470, tokenIndex470
						if buffer[position] != rune('a') {
							goto l466
						}
						position++
						if buffer[position] != rune('n') {
							goto l466
						}
						position++
						if buffer[position] != rune('o') {
							goto l466
						}
						position++
						if buffer[position] != rune('n') {
							goto l466
						}
						position++
					}
				l470:
					{
						position472, tokenIndex472 := position, tokenIndex
						{
							position474, tokenIndex474 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l473
							}
							position, tokenIndex = position474, tokenIndex474
						}
						goto l472
					l473:
						position, tokenIndex = position472, tokenIndex472
						if buffer[position] != rune('.') {
							goto l466
						}
						position++
					}
				l472:
				}
			l468:
				add(ruleUnknownAuthor, position467)
			}
			return true
		l466:
			position, tokenIndex = position466, tokenIndex466
			return false
		},
		/* 69 AuthorWord <- <(AuthorWord1 / AuthorWord2 / AuthorWord3 / AuthorPrefix)> */
		func() bool {
			position475, tokenIndex475 := position, tokenIndex
			{
				position476 := position
				{
					position477, tokenIndex477 := position, tokenIndex
					if !_rules[ruleAuthorWord1]() {
						goto l478
					}
					goto l477
				l478:
					position, tokenIndex = position477, tokenIndex477
					if !_rules[ruleAuthorWord2]() {
						goto l479
					}
					goto l477
				l479:
					position, tokenIndex = position477, tokenIndex477
					if !_rules[ruleAuthorWord3]() {
						goto l480
					}
					goto l477
				l480:
					position, tokenIndex = position477, tokenIndex477
					if !_rules[ruleAuthorPrefix]() {
						goto l475
					}
				}
			l477:
				add(ruleAuthorWord, position476)
			}
			return true
		l475:
			position, tokenIndex = position475, tokenIndex475
			return false
		},
		/* 70 AuthorWord1 <- <(('a' 'r' 'g' '.') / ('e' 't' ' ' 'a' 'l' '.' '{' '?' '}') / ((('e' 't') / '&') (' ' 'a' 'l') '.'?))> */
		func() bool {
			position481, tokenIndex481 := position, tokenIndex
			{
				position482 := position
				{
					position483, tokenIndex483 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l484
					}
					position++
					if buffer[position] != rune('r') {
						goto l484
					}
					position++
					if buffer[position] != rune('g') {
						goto l484
					}
					position++
					if buffer[position] != rune('.') {
						goto l484
					}
					position++
					goto l483
				l484:
					position, tokenIndex = position483, tokenIndex483
					if buffer[position] != rune('e') {
						goto l485
					}
					position++
					if buffer[position] != rune('t') {
						goto l485
					}
					position++
					if buffer[position] != rune(' ') {
						goto l485
					}
					position++
					if buffer[position] != rune('a') {
						goto l485
					}
					position++
					if buffer[position] != rune('l') {
						goto l485
					}
					position++
					if buffer[position] != rune('.') {
						goto l485
					}
					position++
					if buffer[position] != rune('{') {
						goto l485
					}
					position++
					if buffer[position] != rune('?') {
						goto l485
					}
					position++
					if buffer[position] != rune('}') {
						goto l485
					}
					position++
					goto l483
				l485:
					position, tokenIndex = position483, tokenIndex483
					{
						position486, tokenIndex486 := position, tokenIndex
						if buffer[position] != rune('e') {
							goto l487
						}
						position++
						if buffer[position] != rune('t') {
							goto l487
						}
						position++
						goto l486
					l487:
						position, tokenIndex = position486, tokenIndex486
						if buffer[position] != rune('&') {
							goto l481
						}
						position++
					}
				l486:
					if buffer[position] != rune(' ') {
						goto l481
					}
					position++
					if buffer[position] != rune('a') {
						goto l481
					}
					position++
					if buffer[position] != rune('l') {
						goto l481
					}
					position++
					{
						position488, tokenIndex488 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l488
						}
						position++
						goto l489
					l488:
						position, tokenIndex = position488, tokenIndex488
					}
				l489:
				}
			l483:
				add(ruleAuthorWord1, position482)
			}
			return true
		l481:
			position, tokenIndex = position481, tokenIndex481
			return false
		},
		/* 71 AuthorWord2 <- <(AuthorWord3 dash AuthorWordSoft)> */
		func() bool {
			position490, tokenIndex490 := position, tokenIndex
			{
				position491 := position
				if !_rules[ruleAuthorWord3]() {
					goto l490
				}
				if !_rules[ruledash]() {
					goto l490
				}
				if !_rules[ruleAuthorWordSoft]() {
					goto l490
				}
				add(ruleAuthorWord2, position491)
			}
			return true
		l490:
			position, tokenIndex = position490, tokenIndex490
			return false
		},
		/* 72 AuthorWord3 <- <(AuthorPrefixGlued? (AllCapsAuthorWord / CapAuthorWord) '.'?)> */
		func() bool {
			position492, tokenIndex492 := position, tokenIndex
			{
				position493 := position
				{
					position494, tokenIndex494 := position, tokenIndex
					if !_rules[ruleAuthorPrefixGlued]() {
						goto l494
					}
					goto l495
				l494:
					position, tokenIndex = position494, tokenIndex494
				}
			l495:
				{
					position496, tokenIndex496 := position, tokenIndex
					if !_rules[ruleAllCapsAuthorWord]() {
						goto l497
					}
					goto l496
				l497:
					position, tokenIndex = position496, tokenIndex496
					if !_rules[ruleCapAuthorWord]() {
						goto l492
					}
				}
			l496:
				{
					position498, tokenIndex498 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l498
					}
					position++
					goto l499
				l498:
					position, tokenIndex = position498, tokenIndex498
				}
			l499:
				add(ruleAuthorWord3, position493)
			}
			return true
		l492:
			position, tokenIndex = position492, tokenIndex492
			return false
		},
		/* 73 AuthorWordSoft <- <(((AuthorUpperChar (AuthorUpperChar+ / AuthorLowerChar+)) / AuthorLowerChar+) '.'?)> */
		func() bool {
			position500, tokenIndex500 := position, tokenIndex
			{
				position501 := position
				{
					position502, tokenIndex502 := position, tokenIndex
					if !_rules[ruleAuthorUpperChar]() {
						goto l503
					}
					{
						position504, tokenIndex504 := position, tokenIndex
						if !_rules[ruleAuthorUpperChar]() {
							goto l505
						}
					l506:
						{
							position507, tokenIndex507 := position, tokenIndex
							if !_rules[ruleAuthorUpperChar]() {
								goto l507
							}
							goto l506
						l507:
							position, tokenIndex = position507, tokenIndex507
						}
						goto l504
					l505:
						position, tokenIndex = position504, tokenIndex504
						if !_rules[ruleAuthorLowerChar]() {
							goto l503
						}
					l508:
						{
							position509, tokenIndex509 := position, tokenIndex
							if !_rules[ruleAuthorLowerChar]() {
								goto l509
							}
							goto l508
						l509:
							position, tokenIndex = position509, tokenIndex509
						}
					}
				l504:
					goto l502
				l503:
					position, tokenIndex = position502, tokenIndex502
					if !_rules[ruleAuthorLowerChar]() {
						goto l500
					}
				l510:
					{
						position511, tokenIndex511 := position, tokenIndex
						if !_rules[ruleAuthorLowerChar]() {
							goto l511
						}
						goto l510
					l511:
						position, tokenIndex = position511, tokenIndex511
					}
				}
			l502:
				{
					position512, tokenIndex512 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l512
					}
					position++
					goto l513
				l512:
					position, tokenIndex = position512, tokenIndex512
				}
			l513:
				add(ruleAuthorWordSoft, position501)
			}
			return true
		l500:
			position, tokenIndex = position500, tokenIndex500
			return false
		},
		/* 74 CapAuthorWord <- <(AuthorUpperChar AuthorLowerChar*)> */
		func() bool {
			position514, tokenIndex514 := position, tokenIndex
			{
				position515 := position
				if !_rules[ruleAuthorUpperChar]() {
					goto l514
				}
			l516:
				{
					position517, tokenIndex517 := position, tokenIndex
					if !_rules[ruleAuthorLowerChar]() {
						goto l517
					}
					goto l516
				l517:
					position, tokenIndex = position517, tokenIndex517
				}
				add(ruleCapAuthorWord, position515)
			}
			return true
		l514:
			position, tokenIndex = position514, tokenIndex514
			return false
		},
		/* 75 AllCapsAuthorWord <- <(AuthorUpperChar AuthorUpperChar+)> */
		func() bool {
			position518, tokenIndex518 := position, tokenIndex
			{
				position519 := position
				if !_rules[ruleAuthorUpperChar]() {
					goto l518
				}
				if !_rules[ruleAuthorUpperChar]() {
					goto l518
				}
			l520:
				{
					position521, tokenIndex521 := position, tokenIndex
					if !_rules[ruleAuthorUpperChar]() {
						goto l521
					}
					goto l520
				l521:
					position, tokenIndex = position521, tokenIndex521
				}
				add(ruleAllCapsAuthorWord, position519)
			}
			return true
		l518:
			position, tokenIndex = position518, tokenIndex518
			return false
		},
		/* 76 Filius <- <(('f' '.') / ('f' 'i' 'l' '.') / ('f' 'i' 'l' 'i' 'u' 's'))> */
		func() bool {
			position522, tokenIndex522 := position, tokenIndex
			{
				position523 := position
				{
					position524, tokenIndex524 := position, tokenIndex
					if buffer[position] != rune('f') {
						goto l525
					}
					position++
					if buffer[position] != rune('.') {
						goto l525
					}
					position++
					goto l524
				l525:
					position, tokenIndex = position524, tokenIndex524
					if buffer[position] != rune('f') {
						goto l526
					}
					position++
					if buffer[position] != rune('i') {
						goto l526
					}
					position++
					if buffer[position] != rune('l') {
						goto l526
					}
					position++
					if buffer[position] != rune('.') {
						goto l526
					}
					position++
					goto l524
				l526:
					position, tokenIndex = position524, tokenIndex524
					if buffer[position] != rune('f') {
						goto l522
					}
					position++
					if buffer[position] != rune('i') {
						goto l522
					}
					position++
					if buffer[position] != rune('l') {
						goto l522
					}
					position++
					if buffer[position] != rune('i') {
						goto l522
					}
					position++
					if buffer[position] != rune('u') {
						goto l522
					}
					position++
					if buffer[position] != rune('s') {
						goto l522
					}
					position++
				}
			l524:
				add(ruleFilius, position523)
			}
			return true
		l522:
			position, tokenIndex = position522, tokenIndex522
			return false
		},
		/* 77 AuthorPrefixGlued <- <(('d' '\'') / ('O' '\''))> */
		func() bool {
			position527, tokenIndex527 := position, tokenIndex
			{
				position528 := position
				{
					position529, tokenIndex529 := position, tokenIndex
					if buffer[position] != rune('d') {
						goto l530
					}
					position++
					if buffer[position] != rune('\'') {
						goto l530
					}
					position++
					goto l529
				l530:
					position, tokenIndex = position529, tokenIndex529
					if buffer[position] != rune('O') {
						goto l527
					}
					position++
					if buffer[position] != rune('\'') {
						goto l527
					}
					position++
				}
			l529:
				add(ruleAuthorPrefixGlued, position528)
			}
			return true
		l527:
			position, tokenIndex = position527, tokenIndex527
			return false
		},
		/* 78 AuthorPrefix <- <(AuthorPrefix1 / AuthorPrefix2)> */
		func() bool {
			position531, tokenIndex531 := position, tokenIndex
			{
				position532 := position
				{
					position533, tokenIndex533 := position, tokenIndex
					if !_rules[ruleAuthorPrefix1]() {
						goto l534
					}
					goto l533
				l534:
					position, tokenIndex = position533, tokenIndex533
					if !_rules[ruleAuthorPrefix2]() {
						goto l531
					}
				}
			l533:
				add(ruleAuthorPrefix, position532)
			}
			return true
		l531:
			position, tokenIndex = position531, tokenIndex531
			return false
		},
		/* 79 AuthorPrefix2 <- <(('v' '.' (_? ('d' '.'))?) / ('\'' 't'))> */
		func() bool {
			position535, tokenIndex535 := position, tokenIndex
			{
				position536 := position
				{
					position537, tokenIndex537 := position, tokenIndex
					if buffer[position] != rune('v') {
						goto l538
					}
					position++
					if buffer[position] != rune('.') {
						goto l538
					}
					position++
					{
						position539, tokenIndex539 := position, tokenIndex
						{
							position541, tokenIndex541 := position, tokenIndex
							if !_rules[rule_]() {
								goto l541
							}
							goto l542
						l541:
							position, tokenIndex = position541, tokenIndex541
						}
					l542:
						if buffer[position] != rune('d') {
							goto l539
						}
						position++
						if buffer[position] != rune('.') {
							goto l539
						}
						position++
						goto l540
					l539:
						position, tokenIndex = position539, tokenIndex539
					}
				l540:
					goto l537
				l538:
					position, tokenIndex = position537, tokenIndex537
					if buffer[position] != rune('\'') {
						goto l535
					}
					position++
					if buffer[position] != rune('t') {
						goto l535
					}
					position++
				}
			l537:
				add(ruleAuthorPrefix2, position536)
			}
			return true
		l535:
			position, tokenIndex = position535, tokenIndex535
			return false
		},
		/* 80 AuthorPrefix1 <- <((('a' 'b') / ('a' 'f') / ('b' 'i' 's') / ('d' 'a') / ('d' 'e' 'r') / ('d' 'e' 's') / ('d' 'e' 'n') / ('d' 'e' 'l') / ('d' 'e' 'l' 'l' 'a') / ('d' 'e' 'l' 'a') / ('d' 'e') / ('d' 'i') / ('d' 'u') / ('e' 'l') / ('l' 'a') / ('l' 'e') / ('t' 'e' 'r') / ('v' 'a' 'n') / ('d' '\'') / ('i' 'n' '\'' 't') / ('z' 'u' 'r') / ('v' 'o' 'n' (_ (('d' '.') / ('d' 'e' 'm')))?) / ('v' (_ 'd')?)) &_)> */
		func() bool {
			position543, tokenIndex543 := position, tokenIndex
			{
				position544 := position
				{
					position545, tokenIndex545 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l546
					}
					position++
					if buffer[position] != rune('b') {
						goto l546
					}
					position++
					goto l545
				l546:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('a') {
						goto l547
					}
					position++
					if buffer[position] != rune('f') {
						goto l547
					}
					position++
					goto l545
				l547:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('b') {
						goto l548
					}
					position++
					if buffer[position] != rune('i') {
						goto l548
					}
					position++
					if buffer[position] != rune('s') {
						goto l548
					}
					position++
					goto l545
				l548:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('d') {
						goto l549
					}
					position++
					if buffer[position] != rune('a') {
						goto l549
					}
					position++
					goto l545
				l549:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('d') {
						goto l550
					}
					position++
					if buffer[position] != rune('e') {
						goto l550
					}
					position++
					if buffer[position] != rune('r') {
						goto l550
					}
					position++
					goto l545
				l550:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('d') {
						goto l551
					}
					position++
					if buffer[position] != rune('e') {
						goto l551
					}
					position++
					if buffer[position] != rune('s') {
						goto l551
					}
					position++
					goto l545
				l551:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('d') {
						goto l552
					}
					position++
					if buffer[position] != rune('e') {
						goto l552
					}
					position++
					if buffer[position] != rune('n') {
						goto l552
					}
					position++
					goto l545
				l552:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('d') {
						goto l553
					}
					position++
					if buffer[position] != rune('e') {
						goto l553
					}
					position++
					if buffer[position] != rune('l') {
						goto l553
					}
					position++
					goto l545
				l553:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('d') {
						goto l554
					}
					position++
					if buffer[position] != rune('e') {
						goto l554
					}
					position++
					if buffer[position] != rune('l') {
						goto l554
					}
					position++
					if buffer[position] != rune('l') {
						goto l554
					}
					position++
					if buffer[position] != rune('a') {
						goto l554
					}
					position++
					goto l545
				l554:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('d') {
						goto l555
					}
					position++
					if buffer[position] != rune('e') {
						goto l555
					}
					position++
					if buffer[position] != rune('l') {
						goto l555
					}
					position++
					if buffer[position] != rune('a') {
						goto l555
					}
					position++
					goto l545
				l555:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('d') {
						goto l556
					}
					position++
					if buffer[position] != rune('e') {
						goto l556
					}
					position++
					goto l545
				l556:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('d') {
						goto l557
					}
					position++
					if buffer[position] != rune('i') {
						goto l557
					}
					position++
					goto l545
				l557:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('d') {
						goto l558
					}
					position++
					if buffer[position] != rune('u') {
						goto l558
					}
					position++
					goto l545
				l558:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('e') {
						goto l559
					}
					position++
					if buffer[position] != rune('l') {
						goto l559
					}
					position++
					goto l545
				l559:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('l') {
						goto l560
					}
					position++
					if buffer[position] != rune('a') {
						goto l560
					}
					position++
					goto l545
				l560:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('l') {
						goto l561
					}
					position++
					if buffer[position] != rune('e') {
						goto l561
					}
					position++
					goto l545
				l561:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('t') {
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
					goto l545
				l562:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('v') {
						goto l563
					}
					position++
					if buffer[position] != rune('a') {
						goto l563
					}
					position++
					if buffer[position] != rune('n') {
						goto l563
					}
					position++
					goto l545
				l563:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('d') {
						goto l564
					}
					position++
					if buffer[position] != rune('\'') {
						goto l564
					}
					position++
					goto l545
				l564:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('i') {
						goto l565
					}
					position++
					if buffer[position] != rune('n') {
						goto l565
					}
					position++
					if buffer[position] != rune('\'') {
						goto l565
					}
					position++
					if buffer[position] != rune('t') {
						goto l565
					}
					position++
					goto l545
				l565:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('z') {
						goto l566
					}
					position++
					if buffer[position] != rune('u') {
						goto l566
					}
					position++
					if buffer[position] != rune('r') {
						goto l566
					}
					position++
					goto l545
				l566:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('v') {
						goto l567
					}
					position++
					if buffer[position] != rune('o') {
						goto l567
					}
					position++
					if buffer[position] != rune('n') {
						goto l567
					}
					position++
					{
						position568, tokenIndex568 := position, tokenIndex
						if !_rules[rule_]() {
							goto l568
						}
						{
							position570, tokenIndex570 := position, tokenIndex
							if buffer[position] != rune('d') {
								goto l571
							}
							position++
							if buffer[position] != rune('.') {
								goto l571
							}
							position++
							goto l570
						l571:
							position, tokenIndex = position570, tokenIndex570
							if buffer[position] != rune('d') {
								goto l568
							}
							position++
							if buffer[position] != rune('e') {
								goto l568
							}
							position++
							if buffer[position] != rune('m') {
								goto l568
							}
							position++
						}
					l570:
						goto l569
					l568:
						position, tokenIndex = position568, tokenIndex568
					}
				l569:
					goto l545
				l567:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('v') {
						goto l543
					}
					position++
					{
						position572, tokenIndex572 := position, tokenIndex
						if !_rules[rule_]() {
							goto l572
						}
						if buffer[position] != rune('d') {
							goto l572
						}
						position++
						goto l573
					l572:
						position, tokenIndex = position572, tokenIndex572
					}
				l573:
				}
			l545:
				{
					position574, tokenIndex574 := position, tokenIndex
					if !_rules[rule_]() {
						goto l543
					}
					position, tokenIndex = position574, tokenIndex574
				}
				add(ruleAuthorPrefix1, position544)
			}
			return true
		l543:
			position, tokenIndex = position543, tokenIndex543
			return false
		},
		/* 81 AuthorUpperChar <- <(hASCII / ('À' / 'Á' / 'Â' / 'Ã' / 'Ä' / 'Å' / 'Æ' / 'Ç' / 'È' / 'É' / 'Ê' / 'Ë' / 'Ì' / 'Í' / 'Î' / 'Ï' / 'Ð' / 'Ñ' / 'Ò' / 'Ó' / 'Ô' / 'Õ' / 'Ö' / 'Ø' / 'Ù' / 'Ú' / 'Û' / 'Ü' / 'Ý' / 'Ć' / 'Č' / 'Ď' / 'İ' / 'Ķ' / 'Ĺ' / 'ĺ' / 'Ľ' / 'ľ' / 'Ł' / 'ł' / 'Ņ' / 'Ō' / 'Ő' / 'Œ' / 'Ř' / 'Ś' / 'Ŝ' / 'Ş' / 'Š' / 'Ÿ' / 'Ź' / 'Ż' / 'Ž' / 'ƒ' / 'Ǿ' / 'Ș' / 'Ț'))> */
		func() bool {
			position575, tokenIndex575 := position, tokenIndex
			{
				position576 := position
				{
					position577, tokenIndex577 := position, tokenIndex
					if !_rules[rulehASCII]() {
						goto l578
					}
					goto l577
				l578:
					position, tokenIndex = position577, tokenIndex577
					{
						position579, tokenIndex579 := position, tokenIndex
						if buffer[position] != rune('À') {
							goto l580
						}
						position++
						goto l579
					l580:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Á') {
							goto l581
						}
						position++
						goto l579
					l581:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Â') {
							goto l582
						}
						position++
						goto l579
					l582:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ã') {
							goto l583
						}
						position++
						goto l579
					l583:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ä') {
							goto l584
						}
						position++
						goto l579
					l584:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Å') {
							goto l585
						}
						position++
						goto l579
					l585:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Æ') {
							goto l586
						}
						position++
						goto l579
					l586:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ç') {
							goto l587
						}
						position++
						goto l579
					l587:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('È') {
							goto l588
						}
						position++
						goto l579
					l588:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('É') {
							goto l589
						}
						position++
						goto l579
					l589:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ê') {
							goto l590
						}
						position++
						goto l579
					l590:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ë') {
							goto l591
						}
						position++
						goto l579
					l591:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ì') {
							goto l592
						}
						position++
						goto l579
					l592:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Í') {
							goto l593
						}
						position++
						goto l579
					l593:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Î') {
							goto l594
						}
						position++
						goto l579
					l594:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ï') {
							goto l595
						}
						position++
						goto l579
					l595:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ð') {
							goto l596
						}
						position++
						goto l579
					l596:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ñ') {
							goto l597
						}
						position++
						goto l579
					l597:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ò') {
							goto l598
						}
						position++
						goto l579
					l598:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ó') {
							goto l599
						}
						position++
						goto l579
					l599:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ô') {
							goto l600
						}
						position++
						goto l579
					l600:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Õ') {
							goto l601
						}
						position++
						goto l579
					l601:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ö') {
							goto l602
						}
						position++
						goto l579
					l602:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ø') {
							goto l603
						}
						position++
						goto l579
					l603:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ù') {
							goto l604
						}
						position++
						goto l579
					l604:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ú') {
							goto l605
						}
						position++
						goto l579
					l605:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Û') {
							goto l606
						}
						position++
						goto l579
					l606:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ü') {
							goto l607
						}
						position++
						goto l579
					l607:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ý') {
							goto l608
						}
						position++
						goto l579
					l608:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ć') {
							goto l609
						}
						position++
						goto l579
					l609:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Č') {
							goto l610
						}
						position++
						goto l579
					l610:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ď') {
							goto l611
						}
						position++
						goto l579
					l611:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('İ') {
							goto l612
						}
						position++
						goto l579
					l612:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ķ') {
							goto l613
						}
						position++
						goto l579
					l613:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ĺ') {
							goto l614
						}
						position++
						goto l579
					l614:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('ĺ') {
							goto l615
						}
						position++
						goto l579
					l615:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ľ') {
							goto l616
						}
						position++
						goto l579
					l616:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('ľ') {
							goto l617
						}
						position++
						goto l579
					l617:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ł') {
							goto l618
						}
						position++
						goto l579
					l618:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('ł') {
							goto l619
						}
						position++
						goto l579
					l619:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ņ') {
							goto l620
						}
						position++
						goto l579
					l620:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ō') {
							goto l621
						}
						position++
						goto l579
					l621:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ő') {
							goto l622
						}
						position++
						goto l579
					l622:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Œ') {
							goto l623
						}
						position++
						goto l579
					l623:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ř') {
							goto l624
						}
						position++
						goto l579
					l624:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ś') {
							goto l625
						}
						position++
						goto l579
					l625:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ŝ') {
							goto l626
						}
						position++
						goto l579
					l626:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ş') {
							goto l627
						}
						position++
						goto l579
					l627:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Š') {
							goto l628
						}
						position++
						goto l579
					l628:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ÿ') {
							goto l629
						}
						position++
						goto l579
					l629:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ź') {
							goto l630
						}
						position++
						goto l579
					l630:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ż') {
							goto l631
						}
						position++
						goto l579
					l631:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ž') {
							goto l632
						}
						position++
						goto l579
					l632:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('ƒ') {
							goto l633
						}
						position++
						goto l579
					l633:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ǿ') {
							goto l634
						}
						position++
						goto l579
					l634:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ș') {
							goto l635
						}
						position++
						goto l579
					l635:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('Ț') {
							goto l575
						}
						position++
					}
				l579:
				}
			l577:
				add(ruleAuthorUpperChar, position576)
			}
			return true
		l575:
			position, tokenIndex = position575, tokenIndex575
			return false
		},
		/* 82 AuthorLowerChar <- <(lASCII / ('à' / 'á' / 'â' / 'ã' / 'ä' / 'å' / 'æ' / 'ç' / 'è' / 'é' / 'ê' / 'ë' / 'ì' / 'í' / 'î' / 'ï' / 'ð' / 'ñ' / 'ò' / 'ó' / 'ó' / 'ô' / 'õ' / 'ö' / 'ø' / 'ù' / 'ú' / 'û' / 'ü' / 'ý' / 'ÿ' / 'ā' / 'ă' / 'ą' / 'ć' / 'ĉ' / 'č' / 'ď' / 'đ' / '\'' / 'ē' / 'ĕ' / 'ė' / 'ę' / 'ě' / 'ğ' / 'ī' / 'ĭ' / 'İ' / 'ı' / 'ĺ' / 'ľ' / 'ł' / 'ń' / 'ņ' / 'ň' / 'ŏ' / 'ő' / 'œ' / 'ŕ' / 'ř' / 'ś' / 'ş' / 'š' / 'ţ' / 'ť' / 'ũ' / 'ū' / 'ŭ' / 'ů' / 'ű' / 'ź' / 'ż' / 'ž' / 'ſ' / 'ǎ' / 'ǔ' / 'ǧ' / 'ș' / 'ț' / 'ȳ' / 'ß'))> */
		func() bool {
			position636, tokenIndex636 := position, tokenIndex
			{
				position637 := position
				{
					position638, tokenIndex638 := position, tokenIndex
					if !_rules[rulelASCII]() {
						goto l639
					}
					goto l638
				l639:
					position, tokenIndex = position638, tokenIndex638
					{
						position640, tokenIndex640 := position, tokenIndex
						if buffer[position] != rune('à') {
							goto l641
						}
						position++
						goto l640
					l641:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('á') {
							goto l642
						}
						position++
						goto l640
					l642:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('â') {
							goto l643
						}
						position++
						goto l640
					l643:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ã') {
							goto l644
						}
						position++
						goto l640
					l644:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ä') {
							goto l645
						}
						position++
						goto l640
					l645:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('å') {
							goto l646
						}
						position++
						goto l640
					l646:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('æ') {
							goto l647
						}
						position++
						goto l640
					l647:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ç') {
							goto l648
						}
						position++
						goto l640
					l648:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('è') {
							goto l649
						}
						position++
						goto l640
					l649:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('é') {
							goto l650
						}
						position++
						goto l640
					l650:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ê') {
							goto l651
						}
						position++
						goto l640
					l651:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ë') {
							goto l652
						}
						position++
						goto l640
					l652:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ì') {
							goto l653
						}
						position++
						goto l640
					l653:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('í') {
							goto l654
						}
						position++
						goto l640
					l654:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('î') {
							goto l655
						}
						position++
						goto l640
					l655:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ï') {
							goto l656
						}
						position++
						goto l640
					l656:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ð') {
							goto l657
						}
						position++
						goto l640
					l657:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ñ') {
							goto l658
						}
						position++
						goto l640
					l658:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ò') {
							goto l659
						}
						position++
						goto l640
					l659:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ó') {
							goto l660
						}
						position++
						goto l640
					l660:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ó') {
							goto l661
						}
						position++
						goto l640
					l661:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ô') {
							goto l662
						}
						position++
						goto l640
					l662:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('õ') {
							goto l663
						}
						position++
						goto l640
					l663:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ö') {
							goto l664
						}
						position++
						goto l640
					l664:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ø') {
							goto l665
						}
						position++
						goto l640
					l665:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ù') {
							goto l666
						}
						position++
						goto l640
					l666:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ú') {
							goto l667
						}
						position++
						goto l640
					l667:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('û') {
							goto l668
						}
						position++
						goto l640
					l668:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ü') {
							goto l669
						}
						position++
						goto l640
					l669:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ý') {
							goto l670
						}
						position++
						goto l640
					l670:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ÿ') {
							goto l671
						}
						position++
						goto l640
					l671:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ā') {
							goto l672
						}
						position++
						goto l640
					l672:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ă') {
							goto l673
						}
						position++
						goto l640
					l673:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ą') {
							goto l674
						}
						position++
						goto l640
					l674:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ć') {
							goto l675
						}
						position++
						goto l640
					l675:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ĉ') {
							goto l676
						}
						position++
						goto l640
					l676:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('č') {
							goto l677
						}
						position++
						goto l640
					l677:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ď') {
							goto l678
						}
						position++
						goto l640
					l678:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('đ') {
							goto l679
						}
						position++
						goto l640
					l679:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('\'') {
							goto l680
						}
						position++
						goto l640
					l680:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ē') {
							goto l681
						}
						position++
						goto l640
					l681:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ĕ') {
							goto l682
						}
						position++
						goto l640
					l682:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ė') {
							goto l683
						}
						position++
						goto l640
					l683:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ę') {
							goto l684
						}
						position++
						goto l640
					l684:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ě') {
							goto l685
						}
						position++
						goto l640
					l685:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ğ') {
							goto l686
						}
						position++
						goto l640
					l686:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ī') {
							goto l687
						}
						position++
						goto l640
					l687:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ĭ') {
							goto l688
						}
						position++
						goto l640
					l688:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('İ') {
							goto l689
						}
						position++
						goto l640
					l689:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ı') {
							goto l690
						}
						position++
						goto l640
					l690:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ĺ') {
							goto l691
						}
						position++
						goto l640
					l691:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ľ') {
							goto l692
						}
						position++
						goto l640
					l692:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ł') {
							goto l693
						}
						position++
						goto l640
					l693:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ń') {
							goto l694
						}
						position++
						goto l640
					l694:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ņ') {
							goto l695
						}
						position++
						goto l640
					l695:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ň') {
							goto l696
						}
						position++
						goto l640
					l696:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ŏ') {
							goto l697
						}
						position++
						goto l640
					l697:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ő') {
							goto l698
						}
						position++
						goto l640
					l698:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('œ') {
							goto l699
						}
						position++
						goto l640
					l699:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ŕ') {
							goto l700
						}
						position++
						goto l640
					l700:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ř') {
							goto l701
						}
						position++
						goto l640
					l701:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ś') {
							goto l702
						}
						position++
						goto l640
					l702:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ş') {
							goto l703
						}
						position++
						goto l640
					l703:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('š') {
							goto l704
						}
						position++
						goto l640
					l704:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ţ') {
							goto l705
						}
						position++
						goto l640
					l705:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ť') {
							goto l706
						}
						position++
						goto l640
					l706:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ũ') {
							goto l707
						}
						position++
						goto l640
					l707:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ū') {
							goto l708
						}
						position++
						goto l640
					l708:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ŭ') {
							goto l709
						}
						position++
						goto l640
					l709:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ů') {
							goto l710
						}
						position++
						goto l640
					l710:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ű') {
							goto l711
						}
						position++
						goto l640
					l711:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ź') {
							goto l712
						}
						position++
						goto l640
					l712:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ż') {
							goto l713
						}
						position++
						goto l640
					l713:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ž') {
							goto l714
						}
						position++
						goto l640
					l714:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ſ') {
							goto l715
						}
						position++
						goto l640
					l715:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ǎ') {
							goto l716
						}
						position++
						goto l640
					l716:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ǔ') {
							goto l717
						}
						position++
						goto l640
					l717:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ǧ') {
							goto l718
						}
						position++
						goto l640
					l718:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ș') {
							goto l719
						}
						position++
						goto l640
					l719:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ț') {
							goto l720
						}
						position++
						goto l640
					l720:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ȳ') {
							goto l721
						}
						position++
						goto l640
					l721:
						position, tokenIndex = position640, tokenIndex640
						if buffer[position] != rune('ß') {
							goto l636
						}
						position++
					}
				l640:
				}
			l638:
				add(ruleAuthorLowerChar, position637)
			}
			return true
		l636:
			position, tokenIndex = position636, tokenIndex636
			return false
		},
		/* 83 Year <- <(YearRange / YearApprox / YearWithParens / YearWithPage / YearWithDot / YearWithChar / YearNum)> */
		func() bool {
			position722, tokenIndex722 := position, tokenIndex
			{
				position723 := position
				{
					position724, tokenIndex724 := position, tokenIndex
					if !_rules[ruleYearRange]() {
						goto l725
					}
					goto l724
				l725:
					position, tokenIndex = position724, tokenIndex724
					if !_rules[ruleYearApprox]() {
						goto l726
					}
					goto l724
				l726:
					position, tokenIndex = position724, tokenIndex724
					if !_rules[ruleYearWithParens]() {
						goto l727
					}
					goto l724
				l727:
					position, tokenIndex = position724, tokenIndex724
					if !_rules[ruleYearWithPage]() {
						goto l728
					}
					goto l724
				l728:
					position, tokenIndex = position724, tokenIndex724
					if !_rules[ruleYearWithDot]() {
						goto l729
					}
					goto l724
				l729:
					position, tokenIndex = position724, tokenIndex724
					if !_rules[ruleYearWithChar]() {
						goto l730
					}
					goto l724
				l730:
					position, tokenIndex = position724, tokenIndex724
					if !_rules[ruleYearNum]() {
						goto l722
					}
				}
			l724:
				add(ruleYear, position723)
			}
			return true
		l722:
			position, tokenIndex = position722, tokenIndex722
			return false
		},
		/* 84 YearRange <- <(YearNum dash (nums+ ('a' / 'b' / 'c' / 'd' / 'e' / 'f' / 'g' / 'h' / 'i' / 'j' / 'k' / 'l' / 'm' / 'n' / 'o' / 'p' / 'q' / 'r' / 's' / 't' / 'u' / 'v' / 'w' / 'x' / 'y' / 'z' / '?')*))> */
		func() bool {
			position731, tokenIndex731 := position, tokenIndex
			{
				position732 := position
				if !_rules[ruleYearNum]() {
					goto l731
				}
				if !_rules[ruledash]() {
					goto l731
				}
				if !_rules[rulenums]() {
					goto l731
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
			l735:
				{
					position736, tokenIndex736 := position, tokenIndex
					{
						position737, tokenIndex737 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l738
						}
						position++
						goto l737
					l738:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('b') {
							goto l739
						}
						position++
						goto l737
					l739:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('c') {
							goto l740
						}
						position++
						goto l737
					l740:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('d') {
							goto l741
						}
						position++
						goto l737
					l741:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('e') {
							goto l742
						}
						position++
						goto l737
					l742:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('f') {
							goto l743
						}
						position++
						goto l737
					l743:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('g') {
							goto l744
						}
						position++
						goto l737
					l744:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('h') {
							goto l745
						}
						position++
						goto l737
					l745:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('i') {
							goto l746
						}
						position++
						goto l737
					l746:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('j') {
							goto l747
						}
						position++
						goto l737
					l747:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('k') {
							goto l748
						}
						position++
						goto l737
					l748:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('l') {
							goto l749
						}
						position++
						goto l737
					l749:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('m') {
							goto l750
						}
						position++
						goto l737
					l750:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('n') {
							goto l751
						}
						position++
						goto l737
					l751:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('o') {
							goto l752
						}
						position++
						goto l737
					l752:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('p') {
							goto l753
						}
						position++
						goto l737
					l753:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('q') {
							goto l754
						}
						position++
						goto l737
					l754:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('r') {
							goto l755
						}
						position++
						goto l737
					l755:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('s') {
							goto l756
						}
						position++
						goto l737
					l756:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('t') {
							goto l757
						}
						position++
						goto l737
					l757:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('u') {
							goto l758
						}
						position++
						goto l737
					l758:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('v') {
							goto l759
						}
						position++
						goto l737
					l759:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('w') {
							goto l760
						}
						position++
						goto l737
					l760:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('x') {
							goto l761
						}
						position++
						goto l737
					l761:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('y') {
							goto l762
						}
						position++
						goto l737
					l762:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('z') {
							goto l763
						}
						position++
						goto l737
					l763:
						position, tokenIndex = position737, tokenIndex737
						if buffer[position] != rune('?') {
							goto l736
						}
						position++
					}
				l737:
					goto l735
				l736:
					position, tokenIndex = position736, tokenIndex736
				}
				add(ruleYearRange, position732)
			}
			return true
		l731:
			position, tokenIndex = position731, tokenIndex731
			return false
		},
		/* 85 YearWithDot <- <(YearNum '.')> */
		func() bool {
			position764, tokenIndex764 := position, tokenIndex
			{
				position765 := position
				if !_rules[ruleYearNum]() {
					goto l764
				}
				if buffer[position] != rune('.') {
					goto l764
				}
				position++
				add(ruleYearWithDot, position765)
			}
			return true
		l764:
			position, tokenIndex = position764, tokenIndex764
			return false
		},
		/* 86 YearApprox <- <('[' _? YearNum _? ']')> */
		func() bool {
			position766, tokenIndex766 := position, tokenIndex
			{
				position767 := position
				if buffer[position] != rune('[') {
					goto l766
				}
				position++
				{
					position768, tokenIndex768 := position, tokenIndex
					if !_rules[rule_]() {
						goto l768
					}
					goto l769
				l768:
					position, tokenIndex = position768, tokenIndex768
				}
			l769:
				if !_rules[ruleYearNum]() {
					goto l766
				}
				{
					position770, tokenIndex770 := position, tokenIndex
					if !_rules[rule_]() {
						goto l770
					}
					goto l771
				l770:
					position, tokenIndex = position770, tokenIndex770
				}
			l771:
				if buffer[position] != rune(']') {
					goto l766
				}
				position++
				add(ruleYearApprox, position767)
			}
			return true
		l766:
			position, tokenIndex = position766, tokenIndex766
			return false
		},
		/* 87 YearWithPage <- <((YearWithChar / YearNum) _? ':' _? nums+)> */
		func() bool {
			position772, tokenIndex772 := position, tokenIndex
			{
				position773 := position
				{
					position774, tokenIndex774 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l775
					}
					goto l774
				l775:
					position, tokenIndex = position774, tokenIndex774
					if !_rules[ruleYearNum]() {
						goto l772
					}
				}
			l774:
				{
					position776, tokenIndex776 := position, tokenIndex
					if !_rules[rule_]() {
						goto l776
					}
					goto l777
				l776:
					position, tokenIndex = position776, tokenIndex776
				}
			l777:
				if buffer[position] != rune(':') {
					goto l772
				}
				position++
				{
					position778, tokenIndex778 := position, tokenIndex
					if !_rules[rule_]() {
						goto l778
					}
					goto l779
				l778:
					position, tokenIndex = position778, tokenIndex778
				}
			l779:
				if !_rules[rulenums]() {
					goto l772
				}
			l780:
				{
					position781, tokenIndex781 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l781
					}
					goto l780
				l781:
					position, tokenIndex = position781, tokenIndex781
				}
				add(ruleYearWithPage, position773)
			}
			return true
		l772:
			position, tokenIndex = position772, tokenIndex772
			return false
		},
		/* 88 YearWithParens <- <('(' (YearWithChar / YearNum) ')')> */
		func() bool {
			position782, tokenIndex782 := position, tokenIndex
			{
				position783 := position
				if buffer[position] != rune('(') {
					goto l782
				}
				position++
				{
					position784, tokenIndex784 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l785
					}
					goto l784
				l785:
					position, tokenIndex = position784, tokenIndex784
					if !_rules[ruleYearNum]() {
						goto l782
					}
				}
			l784:
				if buffer[position] != rune(')') {
					goto l782
				}
				position++
				add(ruleYearWithParens, position783)
			}
			return true
		l782:
			position, tokenIndex = position782, tokenIndex782
			return false
		},
		/* 89 YearWithChar <- <(YearNum lASCII Action0)> */
		func() bool {
			position786, tokenIndex786 := position, tokenIndex
			{
				position787 := position
				if !_rules[ruleYearNum]() {
					goto l786
				}
				if !_rules[rulelASCII]() {
					goto l786
				}
				if !_rules[ruleAction0]() {
					goto l786
				}
				add(ruleYearWithChar, position787)
			}
			return true
		l786:
			position, tokenIndex = position786, tokenIndex786
			return false
		},
		/* 90 YearNum <- <(('1' / '2') ('0' / '7' / '8' / '9') nums (nums / '?') '?'*)> */
		func() bool {
			position788, tokenIndex788 := position, tokenIndex
			{
				position789 := position
				{
					position790, tokenIndex790 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l791
					}
					position++
					goto l790
				l791:
					position, tokenIndex = position790, tokenIndex790
					if buffer[position] != rune('2') {
						goto l788
					}
					position++
				}
			l790:
				{
					position792, tokenIndex792 := position, tokenIndex
					if buffer[position] != rune('0') {
						goto l793
					}
					position++
					goto l792
				l793:
					position, tokenIndex = position792, tokenIndex792
					if buffer[position] != rune('7') {
						goto l794
					}
					position++
					goto l792
				l794:
					position, tokenIndex = position792, tokenIndex792
					if buffer[position] != rune('8') {
						goto l795
					}
					position++
					goto l792
				l795:
					position, tokenIndex = position792, tokenIndex792
					if buffer[position] != rune('9') {
						goto l788
					}
					position++
				}
			l792:
				if !_rules[rulenums]() {
					goto l788
				}
				{
					position796, tokenIndex796 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l797
					}
					goto l796
				l797:
					position, tokenIndex = position796, tokenIndex796
					if buffer[position] != rune('?') {
						goto l788
					}
					position++
				}
			l796:
			l798:
				{
					position799, tokenIndex799 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l799
					}
					position++
					goto l798
				l799:
					position, tokenIndex = position799, tokenIndex799
				}
				add(ruleYearNum, position789)
			}
			return true
		l788:
			position, tokenIndex = position788, tokenIndex788
			return false
		},
		/* 91 NameUpperChar <- <(UpperChar / UpperCharExtended)> */
		func() bool {
			position800, tokenIndex800 := position, tokenIndex
			{
				position801 := position
				{
					position802, tokenIndex802 := position, tokenIndex
					if !_rules[ruleUpperChar]() {
						goto l803
					}
					goto l802
				l803:
					position, tokenIndex = position802, tokenIndex802
					if !_rules[ruleUpperCharExtended]() {
						goto l800
					}
				}
			l802:
				add(ruleNameUpperChar, position801)
			}
			return true
		l800:
			position, tokenIndex = position800, tokenIndex800
			return false
		},
		/* 92 UpperCharExtended <- <('Æ' / 'Œ' / 'Ö')> */
		func() bool {
			position804, tokenIndex804 := position, tokenIndex
			{
				position805 := position
				{
					position806, tokenIndex806 := position, tokenIndex
					if buffer[position] != rune('Æ') {
						goto l807
					}
					position++
					goto l806
				l807:
					position, tokenIndex = position806, tokenIndex806
					if buffer[position] != rune('Œ') {
						goto l808
					}
					position++
					goto l806
				l808:
					position, tokenIndex = position806, tokenIndex806
					if buffer[position] != rune('Ö') {
						goto l804
					}
					position++
				}
			l806:
				add(ruleUpperCharExtended, position805)
			}
			return true
		l804:
			position, tokenIndex = position804, tokenIndex804
			return false
		},
		/* 93 UpperChar <- <hASCII> */
		func() bool {
			position809, tokenIndex809 := position, tokenIndex
			{
				position810 := position
				if !_rules[rulehASCII]() {
					goto l809
				}
				add(ruleUpperChar, position810)
			}
			return true
		l809:
			position, tokenIndex = position809, tokenIndex809
			return false
		},
		/* 94 NameLowerChar <- <(LowerChar / LowerCharExtended / MiscodedChar)> */
		func() bool {
			position811, tokenIndex811 := position, tokenIndex
			{
				position812 := position
				{
					position813, tokenIndex813 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l814
					}
					goto l813
				l814:
					position, tokenIndex = position813, tokenIndex813
					if !_rules[ruleLowerCharExtended]() {
						goto l815
					}
					goto l813
				l815:
					position, tokenIndex = position813, tokenIndex813
					if !_rules[ruleMiscodedChar]() {
						goto l811
					}
				}
			l813:
				add(ruleNameLowerChar, position812)
			}
			return true
		l811:
			position, tokenIndex = position811, tokenIndex811
			return false
		},
		/* 95 MiscodedChar <- <'�'> */
		func() bool {
			position816, tokenIndex816 := position, tokenIndex
			{
				position817 := position
				if buffer[position] != rune('�') {
					goto l816
				}
				position++
				add(ruleMiscodedChar, position817)
			}
			return true
		l816:
			position, tokenIndex = position816, tokenIndex816
			return false
		},
		/* 96 LowerCharExtended <- <('æ' / 'œ' / 'ſ' / 'à' / 'â' / 'å' / 'ã' / 'ä' / 'á' / 'ç' / 'č' / 'é' / 'è' / 'ë' / 'í' / 'ì' / 'ï' / 'ň' / 'ñ' / 'ñ' / 'ó' / 'ò' / 'ô' / 'ø' / 'õ' / 'ö' / 'ú' / 'ù' / 'ü' / 'ŕ' / 'ř' / 'ŗ' / 'š' / 'š' / 'ş' / 'ž')> */
		func() bool {
			position818, tokenIndex818 := position, tokenIndex
			{
				position819 := position
				{
					position820, tokenIndex820 := position, tokenIndex
					if buffer[position] != rune('æ') {
						goto l821
					}
					position++
					goto l820
				l821:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('œ') {
						goto l822
					}
					position++
					goto l820
				l822:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('ſ') {
						goto l823
					}
					position++
					goto l820
				l823:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('à') {
						goto l824
					}
					position++
					goto l820
				l824:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('â') {
						goto l825
					}
					position++
					goto l820
				l825:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('å') {
						goto l826
					}
					position++
					goto l820
				l826:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('ã') {
						goto l827
					}
					position++
					goto l820
				l827:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('ä') {
						goto l828
					}
					position++
					goto l820
				l828:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('á') {
						goto l829
					}
					position++
					goto l820
				l829:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('ç') {
						goto l830
					}
					position++
					goto l820
				l830:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('č') {
						goto l831
					}
					position++
					goto l820
				l831:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('é') {
						goto l832
					}
					position++
					goto l820
				l832:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('è') {
						goto l833
					}
					position++
					goto l820
				l833:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('ë') {
						goto l834
					}
					position++
					goto l820
				l834:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('í') {
						goto l835
					}
					position++
					goto l820
				l835:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('ì') {
						goto l836
					}
					position++
					goto l820
				l836:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('ï') {
						goto l837
					}
					position++
					goto l820
				l837:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('ň') {
						goto l838
					}
					position++
					goto l820
				l838:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('ñ') {
						goto l839
					}
					position++
					goto l820
				l839:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('ñ') {
						goto l840
					}
					position++
					goto l820
				l840:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('ó') {
						goto l841
					}
					position++
					goto l820
				l841:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('ò') {
						goto l842
					}
					position++
					goto l820
				l842:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('ô') {
						goto l843
					}
					position++
					goto l820
				l843:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('ø') {
						goto l844
					}
					position++
					goto l820
				l844:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('õ') {
						goto l845
					}
					position++
					goto l820
				l845:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('ö') {
						goto l846
					}
					position++
					goto l820
				l846:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('ú') {
						goto l847
					}
					position++
					goto l820
				l847:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('ù') {
						goto l848
					}
					position++
					goto l820
				l848:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('ü') {
						goto l849
					}
					position++
					goto l820
				l849:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('ŕ') {
						goto l850
					}
					position++
					goto l820
				l850:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('ř') {
						goto l851
					}
					position++
					goto l820
				l851:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('ŗ') {
						goto l852
					}
					position++
					goto l820
				l852:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('š') {
						goto l853
					}
					position++
					goto l820
				l853:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('š') {
						goto l854
					}
					position++
					goto l820
				l854:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('ş') {
						goto l855
					}
					position++
					goto l820
				l855:
					position, tokenIndex = position820, tokenIndex820
					if buffer[position] != rune('ž') {
						goto l818
					}
					position++
				}
			l820:
				add(ruleLowerCharExtended, position819)
			}
			return true
		l818:
			position, tokenIndex = position818, tokenIndex818
			return false
		},
		/* 97 LowerChar <- <lASCII> */
		func() bool {
			position856, tokenIndex856 := position, tokenIndex
			{
				position857 := position
				if !_rules[rulelASCII]() {
					goto l856
				}
				add(ruleLowerChar, position857)
			}
			return true
		l856:
			position, tokenIndex = position856, tokenIndex856
			return false
		},
		/* 98 SpaceCharEOI <- <(_ / !.)> */
		func() bool {
			position858, tokenIndex858 := position, tokenIndex
			{
				position859 := position
				{
					position860, tokenIndex860 := position, tokenIndex
					if !_rules[rule_]() {
						goto l861
					}
					goto l860
				l861:
					position, tokenIndex = position860, tokenIndex860
					{
						position862, tokenIndex862 := position, tokenIndex
						if !matchDot() {
							goto l862
						}
						goto l858
					l862:
						position, tokenIndex = position862, tokenIndex862
					}
				}
			l860:
				add(ruleSpaceCharEOI, position859)
			}
			return true
		l858:
			position, tokenIndex = position858, tokenIndex858
			return false
		},
		/* 99 WordBorderChar <- <(_ / (';' / '.' / ',' / ';' / '(' / ')'))> */
		nil,
		/* 100 nums <- <[0-9]> */
		func() bool {
			position864, tokenIndex864 := position, tokenIndex
			{
				position865 := position
				if c := buffer[position]; c < rune('0') || c > rune('9') {
					goto l864
				}
				position++
				add(rulenums, position865)
			}
			return true
		l864:
			position, tokenIndex = position864, tokenIndex864
			return false
		},
		/* 101 lASCII <- <[a-z]> */
		func() bool {
			position866, tokenIndex866 := position, tokenIndex
			{
				position867 := position
				if c := buffer[position]; c < rune('a') || c > rune('z') {
					goto l866
				}
				position++
				add(rulelASCII, position867)
			}
			return true
		l866:
			position, tokenIndex = position866, tokenIndex866
			return false
		},
		/* 102 hASCII <- <[A-Z]> */
		func() bool {
			position868, tokenIndex868 := position, tokenIndex
			{
				position869 := position
				if c := buffer[position]; c < rune('A') || c > rune('Z') {
					goto l868
				}
				position++
				add(rulehASCII, position869)
			}
			return true
		l868:
			position, tokenIndex = position868, tokenIndex868
			return false
		},
		/* 103 apostr <- <'\''> */
		func() bool {
			position870, tokenIndex870 := position, tokenIndex
			{
				position871 := position
				if buffer[position] != rune('\'') {
					goto l870
				}
				position++
				add(ruleapostr, position871)
			}
			return true
		l870:
			position, tokenIndex = position870, tokenIndex870
			return false
		},
		/* 104 dash <- <'-'> */
		func() bool {
			position872, tokenIndex872 := position, tokenIndex
			{
				position873 := position
				if buffer[position] != rune('-') {
					goto l872
				}
				position++
				add(ruledash, position873)
			}
			return true
		l872:
			position, tokenIndex = position872, tokenIndex872
			return false
		},
		/* 105 _ <- <(MultipleSpace / SingleSpace)> */
		func() bool {
			position874, tokenIndex874 := position, tokenIndex
			{
				position875 := position
				{
					position876, tokenIndex876 := position, tokenIndex
					if !_rules[ruleMultipleSpace]() {
						goto l877
					}
					goto l876
				l877:
					position, tokenIndex = position876, tokenIndex876
					if !_rules[ruleSingleSpace]() {
						goto l874
					}
				}
			l876:
				add(rule_, position875)
			}
			return true
		l874:
			position, tokenIndex = position874, tokenIndex874
			return false
		},
		/* 106 MultipleSpace <- <(SingleSpace SingleSpace+)> */
		func() bool {
			position878, tokenIndex878 := position, tokenIndex
			{
				position879 := position
				if !_rules[ruleSingleSpace]() {
					goto l878
				}
				if !_rules[ruleSingleSpace]() {
					goto l878
				}
			l880:
				{
					position881, tokenIndex881 := position, tokenIndex
					if !_rules[ruleSingleSpace]() {
						goto l881
					}
					goto l880
				l881:
					position, tokenIndex = position881, tokenIndex881
				}
				add(ruleMultipleSpace, position879)
			}
			return true
		l878:
			position, tokenIndex = position878, tokenIndex878
			return false
		},
		/* 107 SingleSpace <- <' '> */
		func() bool {
			position882, tokenIndex882 := position, tokenIndex
			{
				position883 := position
				if buffer[position] != rune(' ') {
					goto l882
				}
				position++
				add(ruleSingleSpace, position883)
			}
			return true
		l882:
			position, tokenIndex = position882, tokenIndex882
			return false
		},
		/* 109 Action0 <- <{ p.addWarn(YearCharWarn) }> */
		func() bool {
			{
				add(ruleAction0, position)
			}
			return true
		},
	}
	p.rules = _rules
}
