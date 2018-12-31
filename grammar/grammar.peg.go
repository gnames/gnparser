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
		/* 25 SubGenusOrSuperspecies <- <('(' _? Word _? ')')> */
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
				if !_rules[ruleWord]() {
					goto l171
				}
				{
					position175, tokenIndex175 := position, tokenIndex
					if !_rules[rule_]() {
						goto l175
					}
					goto l176
				l175:
					position, tokenIndex = position175, tokenIndex175
				}
			l176:
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
			position177, tokenIndex177 := position, tokenIndex
			{
				position178 := position
				if buffer[position] != rune('(') {
					goto l177
				}
				position++
				{
					position179, tokenIndex179 := position, tokenIndex
					if !_rules[rule_]() {
						goto l179
					}
					goto l180
				l179:
					position, tokenIndex = position179, tokenIndex179
				}
			l180:
				if !_rules[ruleUninomialWord]() {
					goto l177
				}
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
				if buffer[position] != rune(')') {
					goto l177
				}
				position++
				add(ruleSubGenus, position178)
			}
			return true
		l177:
			position, tokenIndex = position177, tokenIndex177
			return false
		},
		/* 27 UninomialCombo <- <(UninomialCombo1 / UninomialCombo2)> */
		func() bool {
			position183, tokenIndex183 := position, tokenIndex
			{
				position184 := position
				{
					position185, tokenIndex185 := position, tokenIndex
					if !_rules[ruleUninomialCombo1]() {
						goto l186
					}
					goto l185
				l186:
					position, tokenIndex = position185, tokenIndex185
					if !_rules[ruleUninomialCombo2]() {
						goto l183
					}
				}
			l185:
				add(ruleUninomialCombo, position184)
			}
			return true
		l183:
			position, tokenIndex = position183, tokenIndex183
			return false
		},
		/* 28 UninomialCombo1 <- <(UninomialWord _? SubGenus _? Authorship .?)> */
		func() bool {
			position187, tokenIndex187 := position, tokenIndex
			{
				position188 := position
				if !_rules[ruleUninomialWord]() {
					goto l187
				}
				{
					position189, tokenIndex189 := position, tokenIndex
					if !_rules[rule_]() {
						goto l189
					}
					goto l190
				l189:
					position, tokenIndex = position189, tokenIndex189
				}
			l190:
				if !_rules[ruleSubGenus]() {
					goto l187
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
				if !_rules[ruleAuthorship]() {
					goto l187
				}
				{
					position193, tokenIndex193 := position, tokenIndex
					if !matchDot() {
						goto l193
					}
					goto l194
				l193:
					position, tokenIndex = position193, tokenIndex193
				}
			l194:
				add(ruleUninomialCombo1, position188)
			}
			return true
		l187:
			position, tokenIndex = position187, tokenIndex187
			return false
		},
		/* 29 UninomialCombo2 <- <(Uninomial _? RankUninomial _? Uninomial)> */
		func() bool {
			position195, tokenIndex195 := position, tokenIndex
			{
				position196 := position
				if !_rules[ruleUninomial]() {
					goto l195
				}
				{
					position197, tokenIndex197 := position, tokenIndex
					if !_rules[rule_]() {
						goto l197
					}
					goto l198
				l197:
					position, tokenIndex = position197, tokenIndex197
				}
			l198:
				if !_rules[ruleRankUninomial]() {
					goto l195
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
				if !_rules[ruleUninomial]() {
					goto l195
				}
				add(ruleUninomialCombo2, position196)
			}
			return true
		l195:
			position, tokenIndex = position195, tokenIndex195
			return false
		},
		/* 30 RankUninomial <- <((('s' 'e' 'c' 't') / ('s' 'u' 'b' 's' 'e' 'c' 't') / ('t' 'r' 'i' 'b') / ('s' 'u' 'b' 't' 'r' 'i' 'b') / ('s' 'u' 'b' 's' 'e' 'r') / ('s' 'e' 'r') / ('s' 'u' 'b' 'g' 'e' 'n') / ('f' 'a' 'm') / ('s' 'u' 'b' 'f' 'a' 'm') / ('s' 'u' 'p' 'e' 'r' 't' 'r' 'i' 'b')) '.'?)> */
		func() bool {
			position201, tokenIndex201 := position, tokenIndex
			{
				position202 := position
				{
					position203, tokenIndex203 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l204
					}
					position++
					if buffer[position] != rune('e') {
						goto l204
					}
					position++
					if buffer[position] != rune('c') {
						goto l204
					}
					position++
					if buffer[position] != rune('t') {
						goto l204
					}
					position++
					goto l203
				l204:
					position, tokenIndex = position203, tokenIndex203
					if buffer[position] != rune('s') {
						goto l205
					}
					position++
					if buffer[position] != rune('u') {
						goto l205
					}
					position++
					if buffer[position] != rune('b') {
						goto l205
					}
					position++
					if buffer[position] != rune('s') {
						goto l205
					}
					position++
					if buffer[position] != rune('e') {
						goto l205
					}
					position++
					if buffer[position] != rune('c') {
						goto l205
					}
					position++
					if buffer[position] != rune('t') {
						goto l205
					}
					position++
					goto l203
				l205:
					position, tokenIndex = position203, tokenIndex203
					if buffer[position] != rune('t') {
						goto l206
					}
					position++
					if buffer[position] != rune('r') {
						goto l206
					}
					position++
					if buffer[position] != rune('i') {
						goto l206
					}
					position++
					if buffer[position] != rune('b') {
						goto l206
					}
					position++
					goto l203
				l206:
					position, tokenIndex = position203, tokenIndex203
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
					if buffer[position] != rune('t') {
						goto l207
					}
					position++
					if buffer[position] != rune('r') {
						goto l207
					}
					position++
					if buffer[position] != rune('i') {
						goto l207
					}
					position++
					if buffer[position] != rune('b') {
						goto l207
					}
					position++
					goto l203
				l207:
					position, tokenIndex = position203, tokenIndex203
					if buffer[position] != rune('s') {
						goto l208
					}
					position++
					if buffer[position] != rune('u') {
						goto l208
					}
					position++
					if buffer[position] != rune('b') {
						goto l208
					}
					position++
					if buffer[position] != rune('s') {
						goto l208
					}
					position++
					if buffer[position] != rune('e') {
						goto l208
					}
					position++
					if buffer[position] != rune('r') {
						goto l208
					}
					position++
					goto l203
				l208:
					position, tokenIndex = position203, tokenIndex203
					if buffer[position] != rune('s') {
						goto l209
					}
					position++
					if buffer[position] != rune('e') {
						goto l209
					}
					position++
					if buffer[position] != rune('r') {
						goto l209
					}
					position++
					goto l203
				l209:
					position, tokenIndex = position203, tokenIndex203
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
					if buffer[position] != rune('g') {
						goto l210
					}
					position++
					if buffer[position] != rune('e') {
						goto l210
					}
					position++
					if buffer[position] != rune('n') {
						goto l210
					}
					position++
					goto l203
				l210:
					position, tokenIndex = position203, tokenIndex203
					if buffer[position] != rune('f') {
						goto l211
					}
					position++
					if buffer[position] != rune('a') {
						goto l211
					}
					position++
					if buffer[position] != rune('m') {
						goto l211
					}
					position++
					goto l203
				l211:
					position, tokenIndex = position203, tokenIndex203
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
					if buffer[position] != rune('f') {
						goto l212
					}
					position++
					if buffer[position] != rune('a') {
						goto l212
					}
					position++
					if buffer[position] != rune('m') {
						goto l212
					}
					position++
					goto l203
				l212:
					position, tokenIndex = position203, tokenIndex203
					if buffer[position] != rune('s') {
						goto l201
					}
					position++
					if buffer[position] != rune('u') {
						goto l201
					}
					position++
					if buffer[position] != rune('p') {
						goto l201
					}
					position++
					if buffer[position] != rune('e') {
						goto l201
					}
					position++
					if buffer[position] != rune('r') {
						goto l201
					}
					position++
					if buffer[position] != rune('t') {
						goto l201
					}
					position++
					if buffer[position] != rune('r') {
						goto l201
					}
					position++
					if buffer[position] != rune('i') {
						goto l201
					}
					position++
					if buffer[position] != rune('b') {
						goto l201
					}
					position++
				}
			l203:
				{
					position213, tokenIndex213 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l213
					}
					position++
					goto l214
				l213:
					position, tokenIndex = position213, tokenIndex213
				}
			l214:
				add(ruleRankUninomial, position202)
			}
			return true
		l201:
			position, tokenIndex = position201, tokenIndex201
			return false
		},
		/* 31 Uninomial <- <(UninomialWord (_ Authorship)?)> */
		func() bool {
			position215, tokenIndex215 := position, tokenIndex
			{
				position216 := position
				if !_rules[ruleUninomialWord]() {
					goto l215
				}
				{
					position217, tokenIndex217 := position, tokenIndex
					if !_rules[rule_]() {
						goto l217
					}
					if !_rules[ruleAuthorship]() {
						goto l217
					}
					goto l218
				l217:
					position, tokenIndex = position217, tokenIndex217
				}
			l218:
				add(ruleUninomial, position216)
			}
			return true
		l215:
			position, tokenIndex = position215, tokenIndex215
			return false
		},
		/* 32 UninomialWord <- <(CapWord / TwoLetterGenus)> */
		func() bool {
			position219, tokenIndex219 := position, tokenIndex
			{
				position220 := position
				{
					position221, tokenIndex221 := position, tokenIndex
					if !_rules[ruleCapWord]() {
						goto l222
					}
					goto l221
				l222:
					position, tokenIndex = position221, tokenIndex221
					if !_rules[ruleTwoLetterGenus]() {
						goto l219
					}
				}
			l221:
				add(ruleUninomialWord, position220)
			}
			return true
		l219:
			position, tokenIndex = position219, tokenIndex219
			return false
		},
		/* 33 AbbrGenus <- <(UpperChar LowerChar* '.')> */
		func() bool {
			position223, tokenIndex223 := position, tokenIndex
			{
				position224 := position
				if !_rules[ruleUpperChar]() {
					goto l223
				}
			l225:
				{
					position226, tokenIndex226 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l226
					}
					goto l225
				l226:
					position, tokenIndex = position226, tokenIndex226
				}
				if buffer[position] != rune('.') {
					goto l223
				}
				position++
				add(ruleAbbrGenus, position224)
			}
			return true
		l223:
			position, tokenIndex = position223, tokenIndex223
			return false
		},
		/* 34 CapWord <- <(CapWord2 / CapWord1)> */
		func() bool {
			position227, tokenIndex227 := position, tokenIndex
			{
				position228 := position
				{
					position229, tokenIndex229 := position, tokenIndex
					if !_rules[ruleCapWord2]() {
						goto l230
					}
					goto l229
				l230:
					position, tokenIndex = position229, tokenIndex229
					if !_rules[ruleCapWord1]() {
						goto l227
					}
				}
			l229:
				add(ruleCapWord, position228)
			}
			return true
		l227:
			position, tokenIndex = position227, tokenIndex227
			return false
		},
		/* 35 CapWord1 <- <(NameUpperChar NameLowerChar NameLowerChar+ '?'?)> */
		func() bool {
			position231, tokenIndex231 := position, tokenIndex
			{
				position232 := position
				if !_rules[ruleNameUpperChar]() {
					goto l231
				}
				if !_rules[ruleNameLowerChar]() {
					goto l231
				}
				if !_rules[ruleNameLowerChar]() {
					goto l231
				}
			l233:
				{
					position234, tokenIndex234 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l234
					}
					goto l233
				l234:
					position, tokenIndex = position234, tokenIndex234
				}
				{
					position235, tokenIndex235 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l235
					}
					position++
					goto l236
				l235:
					position, tokenIndex = position235, tokenIndex235
				}
			l236:
				add(ruleCapWord1, position232)
			}
			return true
		l231:
			position, tokenIndex = position231, tokenIndex231
			return false
		},
		/* 36 CapWord2 <- <(CapWord1 dash (CapWord1 / Word1))> */
		func() bool {
			position237, tokenIndex237 := position, tokenIndex
			{
				position238 := position
				if !_rules[ruleCapWord1]() {
					goto l237
				}
				if !_rules[ruledash]() {
					goto l237
				}
				{
					position239, tokenIndex239 := position, tokenIndex
					if !_rules[ruleCapWord1]() {
						goto l240
					}
					goto l239
				l240:
					position, tokenIndex = position239, tokenIndex239
					if !_rules[ruleWord1]() {
						goto l237
					}
				}
			l239:
				add(ruleCapWord2, position238)
			}
			return true
		l237:
			position, tokenIndex = position237, tokenIndex237
			return false
		},
		/* 37 TwoLetterGenus <- <(('C' 'a') / ('E' 'a') / ('G' 'e') / ('I' 'a') / ('I' 'o') / ('I' 'x') / ('L' 'o') / ('O' 'a') / ('R' 'a') / ('T' 'y') / ('U' 'a') / ('A' 'a') / ('J' 'a') / ('Z' 'u') / ('L' 'a') / ('Q' 'u') / ('A' 's') / ('B' 'a'))> */
		func() bool {
			position241, tokenIndex241 := position, tokenIndex
			{
				position242 := position
				{
					position243, tokenIndex243 := position, tokenIndex
					if buffer[position] != rune('C') {
						goto l244
					}
					position++
					if buffer[position] != rune('a') {
						goto l244
					}
					position++
					goto l243
				l244:
					position, tokenIndex = position243, tokenIndex243
					if buffer[position] != rune('E') {
						goto l245
					}
					position++
					if buffer[position] != rune('a') {
						goto l245
					}
					position++
					goto l243
				l245:
					position, tokenIndex = position243, tokenIndex243
					if buffer[position] != rune('G') {
						goto l246
					}
					position++
					if buffer[position] != rune('e') {
						goto l246
					}
					position++
					goto l243
				l246:
					position, tokenIndex = position243, tokenIndex243
					if buffer[position] != rune('I') {
						goto l247
					}
					position++
					if buffer[position] != rune('a') {
						goto l247
					}
					position++
					goto l243
				l247:
					position, tokenIndex = position243, tokenIndex243
					if buffer[position] != rune('I') {
						goto l248
					}
					position++
					if buffer[position] != rune('o') {
						goto l248
					}
					position++
					goto l243
				l248:
					position, tokenIndex = position243, tokenIndex243
					if buffer[position] != rune('I') {
						goto l249
					}
					position++
					if buffer[position] != rune('x') {
						goto l249
					}
					position++
					goto l243
				l249:
					position, tokenIndex = position243, tokenIndex243
					if buffer[position] != rune('L') {
						goto l250
					}
					position++
					if buffer[position] != rune('o') {
						goto l250
					}
					position++
					goto l243
				l250:
					position, tokenIndex = position243, tokenIndex243
					if buffer[position] != rune('O') {
						goto l251
					}
					position++
					if buffer[position] != rune('a') {
						goto l251
					}
					position++
					goto l243
				l251:
					position, tokenIndex = position243, tokenIndex243
					if buffer[position] != rune('R') {
						goto l252
					}
					position++
					if buffer[position] != rune('a') {
						goto l252
					}
					position++
					goto l243
				l252:
					position, tokenIndex = position243, tokenIndex243
					if buffer[position] != rune('T') {
						goto l253
					}
					position++
					if buffer[position] != rune('y') {
						goto l253
					}
					position++
					goto l243
				l253:
					position, tokenIndex = position243, tokenIndex243
					if buffer[position] != rune('U') {
						goto l254
					}
					position++
					if buffer[position] != rune('a') {
						goto l254
					}
					position++
					goto l243
				l254:
					position, tokenIndex = position243, tokenIndex243
					if buffer[position] != rune('A') {
						goto l255
					}
					position++
					if buffer[position] != rune('a') {
						goto l255
					}
					position++
					goto l243
				l255:
					position, tokenIndex = position243, tokenIndex243
					if buffer[position] != rune('J') {
						goto l256
					}
					position++
					if buffer[position] != rune('a') {
						goto l256
					}
					position++
					goto l243
				l256:
					position, tokenIndex = position243, tokenIndex243
					if buffer[position] != rune('Z') {
						goto l257
					}
					position++
					if buffer[position] != rune('u') {
						goto l257
					}
					position++
					goto l243
				l257:
					position, tokenIndex = position243, tokenIndex243
					if buffer[position] != rune('L') {
						goto l258
					}
					position++
					if buffer[position] != rune('a') {
						goto l258
					}
					position++
					goto l243
				l258:
					position, tokenIndex = position243, tokenIndex243
					if buffer[position] != rune('Q') {
						goto l259
					}
					position++
					if buffer[position] != rune('u') {
						goto l259
					}
					position++
					goto l243
				l259:
					position, tokenIndex = position243, tokenIndex243
					if buffer[position] != rune('A') {
						goto l260
					}
					position++
					if buffer[position] != rune('s') {
						goto l260
					}
					position++
					goto l243
				l260:
					position, tokenIndex = position243, tokenIndex243
					if buffer[position] != rune('B') {
						goto l241
					}
					position++
					if buffer[position] != rune('a') {
						goto l241
					}
					position++
				}
			l243:
				add(ruleTwoLetterGenus, position242)
			}
			return true
		l241:
			position, tokenIndex = position241, tokenIndex241
			return false
		},
		/* 38 Word <- <(!(AuthorPrefix / RankUninomial / Approximation / Word4) (Word3 / Word2StartDigit / Word2 / Word1) &(SpaceCharEOI / '('))> */
		func() bool {
			position261, tokenIndex261 := position, tokenIndex
			{
				position262 := position
				{
					position263, tokenIndex263 := position, tokenIndex
					{
						position264, tokenIndex264 := position, tokenIndex
						if !_rules[ruleAuthorPrefix]() {
							goto l265
						}
						goto l264
					l265:
						position, tokenIndex = position264, tokenIndex264
						if !_rules[ruleRankUninomial]() {
							goto l266
						}
						goto l264
					l266:
						position, tokenIndex = position264, tokenIndex264
						if !_rules[ruleApproximation]() {
							goto l267
						}
						goto l264
					l267:
						position, tokenIndex = position264, tokenIndex264
						if !_rules[ruleWord4]() {
							goto l263
						}
					}
				l264:
					goto l261
				l263:
					position, tokenIndex = position263, tokenIndex263
				}
				{
					position268, tokenIndex268 := position, tokenIndex
					if !_rules[ruleWord3]() {
						goto l269
					}
					goto l268
				l269:
					position, tokenIndex = position268, tokenIndex268
					if !_rules[ruleWord2StartDigit]() {
						goto l270
					}
					goto l268
				l270:
					position, tokenIndex = position268, tokenIndex268
					if !_rules[ruleWord2]() {
						goto l271
					}
					goto l268
				l271:
					position, tokenIndex = position268, tokenIndex268
					if !_rules[ruleWord1]() {
						goto l261
					}
				}
			l268:
				{
					position272, tokenIndex272 := position, tokenIndex
					{
						position273, tokenIndex273 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l274
						}
						goto l273
					l274:
						position, tokenIndex = position273, tokenIndex273
						if buffer[position] != rune('(') {
							goto l261
						}
						position++
					}
				l273:
					position, tokenIndex = position272, tokenIndex272
				}
				add(ruleWord, position262)
			}
			return true
		l261:
			position, tokenIndex = position261, tokenIndex261
			return false
		},
		/* 39 Word1 <- <((lASCII dash)? NameLowerChar NameLowerChar+)> */
		func() bool {
			position275, tokenIndex275 := position, tokenIndex
			{
				position276 := position
				{
					position277, tokenIndex277 := position, tokenIndex
					if !_rules[rulelASCII]() {
						goto l277
					}
					if !_rules[ruledash]() {
						goto l277
					}
					goto l278
				l277:
					position, tokenIndex = position277, tokenIndex277
				}
			l278:
				if !_rules[ruleNameLowerChar]() {
					goto l275
				}
				if !_rules[ruleNameLowerChar]() {
					goto l275
				}
			l279:
				{
					position280, tokenIndex280 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l280
					}
					goto l279
				l280:
					position, tokenIndex = position280, tokenIndex280
				}
				add(ruleWord1, position276)
			}
			return true
		l275:
			position, tokenIndex = position275, tokenIndex275
			return false
		},
		/* 40 Word2StartDigit <- <(('1' / '2' / '3' / '4' / '5' / '6' / '7' / '8' / '9') nums? ('.' / dash)? NameLowerChar NameLowerChar NameLowerChar NameLowerChar+)> */
		func() bool {
			position281, tokenIndex281 := position, tokenIndex
			{
				position282 := position
				{
					position283, tokenIndex283 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l284
					}
					position++
					goto l283
				l284:
					position, tokenIndex = position283, tokenIndex283
					if buffer[position] != rune('2') {
						goto l285
					}
					position++
					goto l283
				l285:
					position, tokenIndex = position283, tokenIndex283
					if buffer[position] != rune('3') {
						goto l286
					}
					position++
					goto l283
				l286:
					position, tokenIndex = position283, tokenIndex283
					if buffer[position] != rune('4') {
						goto l287
					}
					position++
					goto l283
				l287:
					position, tokenIndex = position283, tokenIndex283
					if buffer[position] != rune('5') {
						goto l288
					}
					position++
					goto l283
				l288:
					position, tokenIndex = position283, tokenIndex283
					if buffer[position] != rune('6') {
						goto l289
					}
					position++
					goto l283
				l289:
					position, tokenIndex = position283, tokenIndex283
					if buffer[position] != rune('7') {
						goto l290
					}
					position++
					goto l283
				l290:
					position, tokenIndex = position283, tokenIndex283
					if buffer[position] != rune('8') {
						goto l291
					}
					position++
					goto l283
				l291:
					position, tokenIndex = position283, tokenIndex283
					if buffer[position] != rune('9') {
						goto l281
					}
					position++
				}
			l283:
				{
					position292, tokenIndex292 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l292
					}
					goto l293
				l292:
					position, tokenIndex = position292, tokenIndex292
				}
			l293:
				{
					position294, tokenIndex294 := position, tokenIndex
					{
						position296, tokenIndex296 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l297
						}
						position++
						goto l296
					l297:
						position, tokenIndex = position296, tokenIndex296
						if !_rules[ruledash]() {
							goto l294
						}
					}
				l296:
					goto l295
				l294:
					position, tokenIndex = position294, tokenIndex294
				}
			l295:
				if !_rules[ruleNameLowerChar]() {
					goto l281
				}
				if !_rules[ruleNameLowerChar]() {
					goto l281
				}
				if !_rules[ruleNameLowerChar]() {
					goto l281
				}
				if !_rules[ruleNameLowerChar]() {
					goto l281
				}
			l298:
				{
					position299, tokenIndex299 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l299
					}
					goto l298
				l299:
					position, tokenIndex = position299, tokenIndex299
				}
				add(ruleWord2StartDigit, position282)
			}
			return true
		l281:
			position, tokenIndex = position281, tokenIndex281
			return false
		},
		/* 41 Word2 <- <(NameLowerChar+ dash? NameLowerChar+)> */
		func() bool {
			position300, tokenIndex300 := position, tokenIndex
			{
				position301 := position
				if !_rules[ruleNameLowerChar]() {
					goto l300
				}
			l302:
				{
					position303, tokenIndex303 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l303
					}
					goto l302
				l303:
					position, tokenIndex = position303, tokenIndex303
				}
				{
					position304, tokenIndex304 := position, tokenIndex
					if !_rules[ruledash]() {
						goto l304
					}
					goto l305
				l304:
					position, tokenIndex = position304, tokenIndex304
				}
			l305:
				if !_rules[ruleNameLowerChar]() {
					goto l300
				}
			l306:
				{
					position307, tokenIndex307 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l307
					}
					goto l306
				l307:
					position, tokenIndex = position307, tokenIndex307
				}
				add(ruleWord2, position301)
			}
			return true
		l300:
			position, tokenIndex = position300, tokenIndex300
			return false
		},
		/* 42 Word3 <- <(NameLowerChar NameLowerChar* apostr Word1)> */
		func() bool {
			position308, tokenIndex308 := position, tokenIndex
			{
				position309 := position
				if !_rules[ruleNameLowerChar]() {
					goto l308
				}
			l310:
				{
					position311, tokenIndex311 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l311
					}
					goto l310
				l311:
					position, tokenIndex = position311, tokenIndex311
				}
				if !_rules[ruleapostr]() {
					goto l308
				}
				if !_rules[ruleWord1]() {
					goto l308
				}
				add(ruleWord3, position309)
			}
			return true
		l308:
			position, tokenIndex = position308, tokenIndex308
			return false
		},
		/* 43 Word4 <- <(NameLowerChar+ '.' NameLowerChar)> */
		func() bool {
			position312, tokenIndex312 := position, tokenIndex
			{
				position313 := position
				if !_rules[ruleNameLowerChar]() {
					goto l312
				}
			l314:
				{
					position315, tokenIndex315 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l315
					}
					goto l314
				l315:
					position, tokenIndex = position315, tokenIndex315
				}
				if buffer[position] != rune('.') {
					goto l312
				}
				position++
				if !_rules[ruleNameLowerChar]() {
					goto l312
				}
				add(ruleWord4, position313)
			}
			return true
		l312:
			position, tokenIndex = position312, tokenIndex312
			return false
		},
		/* 44 HybridChar <- <'×'> */
		func() bool {
			position316, tokenIndex316 := position, tokenIndex
			{
				position317 := position
				if buffer[position] != rune('×') {
					goto l316
				}
				position++
				add(ruleHybridChar, position317)
			}
			return true
		l316:
			position, tokenIndex = position316, tokenIndex316
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
			position322, tokenIndex322 := position, tokenIndex
			{
				position323 := position
				{
					position324, tokenIndex324 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l325
					}
					position++
					if buffer[position] != rune('p') {
						goto l325
					}
					position++
					if buffer[position] != rune('.') {
						goto l325
					}
					position++
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
					if buffer[position] != rune('n') {
						goto l325
					}
					position++
					if buffer[position] != rune('r') {
						goto l325
					}
					position++
					if buffer[position] != rune('.') {
						goto l325
					}
					position++
					goto l324
				l325:
					position, tokenIndex = position324, tokenIndex324
					if buffer[position] != rune('s') {
						goto l328
					}
					position++
					if buffer[position] != rune('p') {
						goto l328
					}
					position++
					if buffer[position] != rune('.') {
						goto l328
					}
					position++
					{
						position329, tokenIndex329 := position, tokenIndex
						if !_rules[rule_]() {
							goto l329
						}
						goto l330
					l329:
						position, tokenIndex = position329, tokenIndex329
					}
				l330:
					if buffer[position] != rune('a') {
						goto l328
					}
					position++
					if buffer[position] != rune('f') {
						goto l328
					}
					position++
					if buffer[position] != rune('f') {
						goto l328
					}
					position++
					if buffer[position] != rune('.') {
						goto l328
					}
					position++
					goto l324
				l328:
					position, tokenIndex = position324, tokenIndex324
					if buffer[position] != rune('m') {
						goto l331
					}
					position++
					if buffer[position] != rune('o') {
						goto l331
					}
					position++
					if buffer[position] != rune('n') {
						goto l331
					}
					position++
					if buffer[position] != rune('s') {
						goto l331
					}
					position++
					if buffer[position] != rune('t') {
						goto l331
					}
					position++
					if buffer[position] != rune('.') {
						goto l331
					}
					position++
					goto l324
				l331:
					position, tokenIndex = position324, tokenIndex324
					if buffer[position] != rune('?') {
						goto l332
					}
					position++
					goto l324
				l332:
					position, tokenIndex = position324, tokenIndex324
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
						if buffer[position] != rune('p') {
							goto l334
						}
						position++
						goto l333
					l334:
						position, tokenIndex = position333, tokenIndex333
						if buffer[position] != rune('n') {
							goto l335
						}
						position++
						if buffer[position] != rune('r') {
							goto l335
						}
						position++
						goto l333
					l335:
						position, tokenIndex = position333, tokenIndex333
						if buffer[position] != rune('s') {
							goto l336
						}
						position++
						if buffer[position] != rune('p') {
							goto l336
						}
						position++
						goto l333
					l336:
						position, tokenIndex = position333, tokenIndex333
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
						goto l333
					l337:
						position, tokenIndex = position333, tokenIndex333
						if buffer[position] != rune('s') {
							goto l322
						}
						position++
						if buffer[position] != rune('p') {
							goto l322
						}
						position++
						if buffer[position] != rune('e') {
							goto l322
						}
						position++
						if buffer[position] != rune('c') {
							goto l322
						}
						position++
						if buffer[position] != rune('i') {
							goto l322
						}
						position++
						if buffer[position] != rune('e') {
							goto l322
						}
						position++
						if buffer[position] != rune('s') {
							goto l322
						}
						position++
					}
				l333:
					{
						position338, tokenIndex338 := position, tokenIndex
						{
							position340, tokenIndex340 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l339
							}
							position, tokenIndex = position340, tokenIndex340
						}
						goto l338
					l339:
						position, tokenIndex = position338, tokenIndex338
						if buffer[position] != rune('.') {
							goto l322
						}
						position++
					}
				l338:
				}
			l324:
				add(ruleApproximation, position323)
			}
			return true
		l322:
			position, tokenIndex = position322, tokenIndex322
			return false
		},
		/* 50 Authorship <- <((AuthorshipCombo / OriginalAuthorship) &(SpaceCharEOI / ('\\' / '(' / ',' / ':')))> */
		func() bool {
			position341, tokenIndex341 := position, tokenIndex
			{
				position342 := position
				{
					position343, tokenIndex343 := position, tokenIndex
					if !_rules[ruleAuthorshipCombo]() {
						goto l344
					}
					goto l343
				l344:
					position, tokenIndex = position343, tokenIndex343
					if !_rules[ruleOriginalAuthorship]() {
						goto l341
					}
				}
			l343:
				{
					position345, tokenIndex345 := position, tokenIndex
					{
						position346, tokenIndex346 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l347
						}
						goto l346
					l347:
						position, tokenIndex = position346, tokenIndex346
						{
							position348, tokenIndex348 := position, tokenIndex
							if buffer[position] != rune('\\') {
								goto l349
							}
							position++
							goto l348
						l349:
							position, tokenIndex = position348, tokenIndex348
							if buffer[position] != rune('(') {
								goto l350
							}
							position++
							goto l348
						l350:
							position, tokenIndex = position348, tokenIndex348
							if buffer[position] != rune(',') {
								goto l351
							}
							position++
							goto l348
						l351:
							position, tokenIndex = position348, tokenIndex348
							if buffer[position] != rune(':') {
								goto l341
							}
							position++
						}
					l348:
					}
				l346:
					position, tokenIndex = position345, tokenIndex345
				}
				add(ruleAuthorship, position342)
			}
			return true
		l341:
			position, tokenIndex = position341, tokenIndex341
			return false
		},
		/* 51 AuthorshipCombo <- <(OriginalAuthorship _? CombinationAuthorship)> */
		func() bool {
			position352, tokenIndex352 := position, tokenIndex
			{
				position353 := position
				if !_rules[ruleOriginalAuthorship]() {
					goto l352
				}
				{
					position354, tokenIndex354 := position, tokenIndex
					if !_rules[rule_]() {
						goto l354
					}
					goto l355
				l354:
					position, tokenIndex = position354, tokenIndex354
				}
			l355:
				if !_rules[ruleCombinationAuthorship]() {
					goto l352
				}
				add(ruleAuthorshipCombo, position353)
			}
			return true
		l352:
			position, tokenIndex = position352, tokenIndex352
			return false
		},
		/* 52 OriginalAuthorship <- <(AuthorsGroup / BasionymAuthorship / BasionymAuthorshipYearMisformed)> */
		func() bool {
			position356, tokenIndex356 := position, tokenIndex
			{
				position357 := position
				{
					position358, tokenIndex358 := position, tokenIndex
					if !_rules[ruleAuthorsGroup]() {
						goto l359
					}
					goto l358
				l359:
					position, tokenIndex = position358, tokenIndex358
					if !_rules[ruleBasionymAuthorship]() {
						goto l360
					}
					goto l358
				l360:
					position, tokenIndex = position358, tokenIndex358
					if !_rules[ruleBasionymAuthorshipYearMisformed]() {
						goto l356
					}
				}
			l358:
				add(ruleOriginalAuthorship, position357)
			}
			return true
		l356:
			position, tokenIndex = position356, tokenIndex356
			return false
		},
		/* 53 CombinationAuthorship <- <AuthorsGroup> */
		func() bool {
			position361, tokenIndex361 := position, tokenIndex
			{
				position362 := position
				if !_rules[ruleAuthorsGroup]() {
					goto l361
				}
				add(ruleCombinationAuthorship, position362)
			}
			return true
		l361:
			position, tokenIndex = position361, tokenIndex361
			return false
		},
		/* 54 BasionymAuthorshipYearMisformed <- <('(' _? AuthorsGroup _? ')' (_? ',')? _? Year)> */
		func() bool {
			position363, tokenIndex363 := position, tokenIndex
			{
				position364 := position
				if buffer[position] != rune('(') {
					goto l363
				}
				position++
				{
					position365, tokenIndex365 := position, tokenIndex
					if !_rules[rule_]() {
						goto l365
					}
					goto l366
				l365:
					position, tokenIndex = position365, tokenIndex365
				}
			l366:
				if !_rules[ruleAuthorsGroup]() {
					goto l363
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
				if buffer[position] != rune(')') {
					goto l363
				}
				position++
				{
					position369, tokenIndex369 := position, tokenIndex
					{
						position371, tokenIndex371 := position, tokenIndex
						if !_rules[rule_]() {
							goto l371
						}
						goto l372
					l371:
						position, tokenIndex = position371, tokenIndex371
					}
				l372:
					if buffer[position] != rune(',') {
						goto l369
					}
					position++
					goto l370
				l369:
					position, tokenIndex = position369, tokenIndex369
				}
			l370:
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
				if !_rules[ruleYear]() {
					goto l363
				}
				add(ruleBasionymAuthorshipYearMisformed, position364)
			}
			return true
		l363:
			position, tokenIndex = position363, tokenIndex363
			return false
		},
		/* 55 BasionymAuthorship <- <(BasionymAuthorship1 / BasionymAuthorship2)> */
		func() bool {
			position375, tokenIndex375 := position, tokenIndex
			{
				position376 := position
				{
					position377, tokenIndex377 := position, tokenIndex
					if !_rules[ruleBasionymAuthorship1]() {
						goto l378
					}
					goto l377
				l378:
					position, tokenIndex = position377, tokenIndex377
					if !_rules[ruleBasionymAuthorship2]() {
						goto l375
					}
				}
			l377:
				add(ruleBasionymAuthorship, position376)
			}
			return true
		l375:
			position, tokenIndex = position375, tokenIndex375
			return false
		},
		/* 56 BasionymAuthorship1 <- <('(' _? AuthorsGroup _? ')')> */
		func() bool {
			position379, tokenIndex379 := position, tokenIndex
			{
				position380 := position
				if buffer[position] != rune('(') {
					goto l379
				}
				position++
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
				if !_rules[ruleAuthorsGroup]() {
					goto l379
				}
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
				if buffer[position] != rune(')') {
					goto l379
				}
				position++
				add(ruleBasionymAuthorship1, position380)
			}
			return true
		l379:
			position, tokenIndex = position379, tokenIndex379
			return false
		},
		/* 57 BasionymAuthorship2 <- <('(' _? '(' _? AuthorsGroup _? ')' _? ')')> */
		func() bool {
			position385, tokenIndex385 := position, tokenIndex
			{
				position386 := position
				if buffer[position] != rune('(') {
					goto l385
				}
				position++
				{
					position387, tokenIndex387 := position, tokenIndex
					if !_rules[rule_]() {
						goto l387
					}
					goto l388
				l387:
					position, tokenIndex = position387, tokenIndex387
				}
			l388:
				if buffer[position] != rune('(') {
					goto l385
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
				if !_rules[ruleAuthorsGroup]() {
					goto l385
				}
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
				if buffer[position] != rune(')') {
					goto l385
				}
				position++
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
					goto l385
				}
				position++
				add(ruleBasionymAuthorship2, position386)
			}
			return true
		l385:
			position, tokenIndex = position385, tokenIndex385
			return false
		},
		/* 58 AuthorsGroup <- <(AuthorsTeam (_? AuthorEmend? AuthorEx? AuthorsTeam)?)> */
		func() bool {
			position395, tokenIndex395 := position, tokenIndex
			{
				position396 := position
				if !_rules[ruleAuthorsTeam]() {
					goto l395
				}
				{
					position397, tokenIndex397 := position, tokenIndex
					{
						position399, tokenIndex399 := position, tokenIndex
						if !_rules[rule_]() {
							goto l399
						}
						goto l400
					l399:
						position, tokenIndex = position399, tokenIndex399
					}
				l400:
					{
						position401, tokenIndex401 := position, tokenIndex
						if !_rules[ruleAuthorEmend]() {
							goto l401
						}
						goto l402
					l401:
						position, tokenIndex = position401, tokenIndex401
					}
				l402:
					{
						position403, tokenIndex403 := position, tokenIndex
						if !_rules[ruleAuthorEx]() {
							goto l403
						}
						goto l404
					l403:
						position, tokenIndex = position403, tokenIndex403
					}
				l404:
					if !_rules[ruleAuthorsTeam]() {
						goto l397
					}
					goto l398
				l397:
					position, tokenIndex = position397, tokenIndex397
				}
			l398:
				add(ruleAuthorsGroup, position396)
			}
			return true
		l395:
			position, tokenIndex = position395, tokenIndex395
			return false
		},
		/* 59 AuthorsTeam <- <(Author (AuthorSep Author)* (_? ','? _? Year)?)> */
		func() bool {
			position405, tokenIndex405 := position, tokenIndex
			{
				position406 := position
				if !_rules[ruleAuthor]() {
					goto l405
				}
			l407:
				{
					position408, tokenIndex408 := position, tokenIndex
					if !_rules[ruleAuthorSep]() {
						goto l408
					}
					if !_rules[ruleAuthor]() {
						goto l408
					}
					goto l407
				l408:
					position, tokenIndex = position408, tokenIndex408
				}
				{
					position409, tokenIndex409 := position, tokenIndex
					{
						position411, tokenIndex411 := position, tokenIndex
						if !_rules[rule_]() {
							goto l411
						}
						goto l412
					l411:
						position, tokenIndex = position411, tokenIndex411
					}
				l412:
					{
						position413, tokenIndex413 := position, tokenIndex
						if buffer[position] != rune(',') {
							goto l413
						}
						position++
						goto l414
					l413:
						position, tokenIndex = position413, tokenIndex413
					}
				l414:
					{
						position415, tokenIndex415 := position, tokenIndex
						if !_rules[rule_]() {
							goto l415
						}
						goto l416
					l415:
						position, tokenIndex = position415, tokenIndex415
					}
				l416:
					if !_rules[ruleYear]() {
						goto l409
					}
					goto l410
				l409:
					position, tokenIndex = position409, tokenIndex409
				}
			l410:
				add(ruleAuthorsTeam, position406)
			}
			return true
		l405:
			position, tokenIndex = position405, tokenIndex405
			return false
		},
		/* 60 AuthorSep <- <(AuthorSep1 / AuthorSep2)> */
		func() bool {
			position417, tokenIndex417 := position, tokenIndex
			{
				position418 := position
				{
					position419, tokenIndex419 := position, tokenIndex
					if !_rules[ruleAuthorSep1]() {
						goto l420
					}
					goto l419
				l420:
					position, tokenIndex = position419, tokenIndex419
					if !_rules[ruleAuthorSep2]() {
						goto l417
					}
				}
			l419:
				add(ruleAuthorSep, position418)
			}
			return true
		l417:
			position, tokenIndex = position417, tokenIndex417
			return false
		},
		/* 61 AuthorSep1 <- <(_? (',' _)? ('&' / ('e' 't') / ('a' 'n' 'd') / ('a' 'p' 'u' 'd')) _?)> */
		func() bool {
			position421, tokenIndex421 := position, tokenIndex
			{
				position422 := position
				{
					position423, tokenIndex423 := position, tokenIndex
					if !_rules[rule_]() {
						goto l423
					}
					goto l424
				l423:
					position, tokenIndex = position423, tokenIndex423
				}
			l424:
				{
					position425, tokenIndex425 := position, tokenIndex
					if buffer[position] != rune(',') {
						goto l425
					}
					position++
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
					if buffer[position] != rune('&') {
						goto l428
					}
					position++
					goto l427
				l428:
					position, tokenIndex = position427, tokenIndex427
					if buffer[position] != rune('e') {
						goto l429
					}
					position++
					if buffer[position] != rune('t') {
						goto l429
					}
					position++
					goto l427
				l429:
					position, tokenIndex = position427, tokenIndex427
					if buffer[position] != rune('a') {
						goto l430
					}
					position++
					if buffer[position] != rune('n') {
						goto l430
					}
					position++
					if buffer[position] != rune('d') {
						goto l430
					}
					position++
					goto l427
				l430:
					position, tokenIndex = position427, tokenIndex427
					if buffer[position] != rune('a') {
						goto l421
					}
					position++
					if buffer[position] != rune('p') {
						goto l421
					}
					position++
					if buffer[position] != rune('u') {
						goto l421
					}
					position++
					if buffer[position] != rune('d') {
						goto l421
					}
					position++
				}
			l427:
				{
					position431, tokenIndex431 := position, tokenIndex
					if !_rules[rule_]() {
						goto l431
					}
					goto l432
				l431:
					position, tokenIndex = position431, tokenIndex431
				}
			l432:
				add(ruleAuthorSep1, position422)
			}
			return true
		l421:
			position, tokenIndex = position421, tokenIndex421
			return false
		},
		/* 62 AuthorSep2 <- <(_? ',' _?)> */
		func() bool {
			position433, tokenIndex433 := position, tokenIndex
			{
				position434 := position
				{
					position435, tokenIndex435 := position, tokenIndex
					if !_rules[rule_]() {
						goto l435
					}
					goto l436
				l435:
					position, tokenIndex = position435, tokenIndex435
				}
			l436:
				if buffer[position] != rune(',') {
					goto l433
				}
				position++
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
				add(ruleAuthorSep2, position434)
			}
			return true
		l433:
			position, tokenIndex = position433, tokenIndex433
			return false
		},
		/* 63 AuthorEx <- <((('e' 'x' '.'?) / ('i' 'n')) _)> */
		func() bool {
			position439, tokenIndex439 := position, tokenIndex
			{
				position440 := position
				{
					position441, tokenIndex441 := position, tokenIndex
					if buffer[position] != rune('e') {
						goto l442
					}
					position++
					if buffer[position] != rune('x') {
						goto l442
					}
					position++
					{
						position443, tokenIndex443 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l443
						}
						position++
						goto l444
					l443:
						position, tokenIndex = position443, tokenIndex443
					}
				l444:
					goto l441
				l442:
					position, tokenIndex = position441, tokenIndex441
					if buffer[position] != rune('i') {
						goto l439
					}
					position++
					if buffer[position] != rune('n') {
						goto l439
					}
					position++
				}
			l441:
				if !_rules[rule_]() {
					goto l439
				}
				add(ruleAuthorEx, position440)
			}
			return true
		l439:
			position, tokenIndex = position439, tokenIndex439
			return false
		},
		/* 64 AuthorEmend <- <('e' 'm' 'e' 'n' 'd' '.'? _)> */
		func() bool {
			position445, tokenIndex445 := position, tokenIndex
			{
				position446 := position
				if buffer[position] != rune('e') {
					goto l445
				}
				position++
				if buffer[position] != rune('m') {
					goto l445
				}
				position++
				if buffer[position] != rune('e') {
					goto l445
				}
				position++
				if buffer[position] != rune('n') {
					goto l445
				}
				position++
				if buffer[position] != rune('d') {
					goto l445
				}
				position++
				{
					position447, tokenIndex447 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l447
					}
					position++
					goto l448
				l447:
					position, tokenIndex = position447, tokenIndex447
				}
			l448:
				if !_rules[rule_]() {
					goto l445
				}
				add(ruleAuthorEmend, position446)
			}
			return true
		l445:
			position, tokenIndex = position445, tokenIndex445
			return false
		},
		/* 65 Author <- <(Author1 / Author2 / UnknownAuthor)> */
		func() bool {
			position449, tokenIndex449 := position, tokenIndex
			{
				position450 := position
				{
					position451, tokenIndex451 := position, tokenIndex
					if !_rules[ruleAuthor1]() {
						goto l452
					}
					goto l451
				l452:
					position, tokenIndex = position451, tokenIndex451
					if !_rules[ruleAuthor2]() {
						goto l453
					}
					goto l451
				l453:
					position, tokenIndex = position451, tokenIndex451
					if !_rules[ruleUnknownAuthor]() {
						goto l449
					}
				}
			l451:
				add(ruleAuthor, position450)
			}
			return true
		l449:
			position, tokenIndex = position449, tokenIndex449
			return false
		},
		/* 66 Author1 <- <(Author2 _? Filius)> */
		func() bool {
			position454, tokenIndex454 := position, tokenIndex
			{
				position455 := position
				if !_rules[ruleAuthor2]() {
					goto l454
				}
				{
					position456, tokenIndex456 := position, tokenIndex
					if !_rules[rule_]() {
						goto l456
					}
					goto l457
				l456:
					position, tokenIndex = position456, tokenIndex456
				}
			l457:
				if !_rules[ruleFilius]() {
					goto l454
				}
				add(ruleAuthor1, position455)
			}
			return true
		l454:
			position, tokenIndex = position454, tokenIndex454
			return false
		},
		/* 67 Author2 <- <(AuthorWord (_? AuthorWord)*)> */
		func() bool {
			position458, tokenIndex458 := position, tokenIndex
			{
				position459 := position
				if !_rules[ruleAuthorWord]() {
					goto l458
				}
			l460:
				{
					position461, tokenIndex461 := position, tokenIndex
					{
						position462, tokenIndex462 := position, tokenIndex
						if !_rules[rule_]() {
							goto l462
						}
						goto l463
					l462:
						position, tokenIndex = position462, tokenIndex462
					}
				l463:
					if !_rules[ruleAuthorWord]() {
						goto l461
					}
					goto l460
				l461:
					position, tokenIndex = position461, tokenIndex461
				}
				add(ruleAuthor2, position459)
			}
			return true
		l458:
			position, tokenIndex = position458, tokenIndex458
			return false
		},
		/* 68 UnknownAuthor <- <('?' / ((('a' 'u' 'c' 't') / ('a' 'n' 'o' 'n')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position464, tokenIndex464 := position, tokenIndex
			{
				position465 := position
				{
					position466, tokenIndex466 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l467
					}
					position++
					goto l466
				l467:
					position, tokenIndex = position466, tokenIndex466
					{
						position468, tokenIndex468 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l469
						}
						position++
						if buffer[position] != rune('u') {
							goto l469
						}
						position++
						if buffer[position] != rune('c') {
							goto l469
						}
						position++
						if buffer[position] != rune('t') {
							goto l469
						}
						position++
						goto l468
					l469:
						position, tokenIndex = position468, tokenIndex468
						if buffer[position] != rune('a') {
							goto l464
						}
						position++
						if buffer[position] != rune('n') {
							goto l464
						}
						position++
						if buffer[position] != rune('o') {
							goto l464
						}
						position++
						if buffer[position] != rune('n') {
							goto l464
						}
						position++
					}
				l468:
					{
						position470, tokenIndex470 := position, tokenIndex
						{
							position472, tokenIndex472 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l471
							}
							position, tokenIndex = position472, tokenIndex472
						}
						goto l470
					l471:
						position, tokenIndex = position470, tokenIndex470
						if buffer[position] != rune('.') {
							goto l464
						}
						position++
					}
				l470:
				}
			l466:
				add(ruleUnknownAuthor, position465)
			}
			return true
		l464:
			position, tokenIndex = position464, tokenIndex464
			return false
		},
		/* 69 AuthorWord <- <(AuthorWord1 / AuthorWord2 / AuthorWord3 / AuthorPrefix)> */
		func() bool {
			position473, tokenIndex473 := position, tokenIndex
			{
				position474 := position
				{
					position475, tokenIndex475 := position, tokenIndex
					if !_rules[ruleAuthorWord1]() {
						goto l476
					}
					goto l475
				l476:
					position, tokenIndex = position475, tokenIndex475
					if !_rules[ruleAuthorWord2]() {
						goto l477
					}
					goto l475
				l477:
					position, tokenIndex = position475, tokenIndex475
					if !_rules[ruleAuthorWord3]() {
						goto l478
					}
					goto l475
				l478:
					position, tokenIndex = position475, tokenIndex475
					if !_rules[ruleAuthorPrefix]() {
						goto l473
					}
				}
			l475:
				add(ruleAuthorWord, position474)
			}
			return true
		l473:
			position, tokenIndex = position473, tokenIndex473
			return false
		},
		/* 70 AuthorWord1 <- <(('a' 'r' 'g' '.') / ('e' 't' ' ' 'a' 'l' '.' '{' '?' '}') / ((('e' 't') / '&') (' ' 'a' 'l') '.'?))> */
		func() bool {
			position479, tokenIndex479 := position, tokenIndex
			{
				position480 := position
				{
					position481, tokenIndex481 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l482
					}
					position++
					if buffer[position] != rune('r') {
						goto l482
					}
					position++
					if buffer[position] != rune('g') {
						goto l482
					}
					position++
					if buffer[position] != rune('.') {
						goto l482
					}
					position++
					goto l481
				l482:
					position, tokenIndex = position481, tokenIndex481
					if buffer[position] != rune('e') {
						goto l483
					}
					position++
					if buffer[position] != rune('t') {
						goto l483
					}
					position++
					if buffer[position] != rune(' ') {
						goto l483
					}
					position++
					if buffer[position] != rune('a') {
						goto l483
					}
					position++
					if buffer[position] != rune('l') {
						goto l483
					}
					position++
					if buffer[position] != rune('.') {
						goto l483
					}
					position++
					if buffer[position] != rune('{') {
						goto l483
					}
					position++
					if buffer[position] != rune('?') {
						goto l483
					}
					position++
					if buffer[position] != rune('}') {
						goto l483
					}
					position++
					goto l481
				l483:
					position, tokenIndex = position481, tokenIndex481
					{
						position484, tokenIndex484 := position, tokenIndex
						if buffer[position] != rune('e') {
							goto l485
						}
						position++
						if buffer[position] != rune('t') {
							goto l485
						}
						position++
						goto l484
					l485:
						position, tokenIndex = position484, tokenIndex484
						if buffer[position] != rune('&') {
							goto l479
						}
						position++
					}
				l484:
					if buffer[position] != rune(' ') {
						goto l479
					}
					position++
					if buffer[position] != rune('a') {
						goto l479
					}
					position++
					if buffer[position] != rune('l') {
						goto l479
					}
					position++
					{
						position486, tokenIndex486 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l486
						}
						position++
						goto l487
					l486:
						position, tokenIndex = position486, tokenIndex486
					}
				l487:
				}
			l481:
				add(ruleAuthorWord1, position480)
			}
			return true
		l479:
			position, tokenIndex = position479, tokenIndex479
			return false
		},
		/* 71 AuthorWord2 <- <(AuthorWord3 dash AuthorWordSoft)> */
		func() bool {
			position488, tokenIndex488 := position, tokenIndex
			{
				position489 := position
				if !_rules[ruleAuthorWord3]() {
					goto l488
				}
				if !_rules[ruledash]() {
					goto l488
				}
				if !_rules[ruleAuthorWordSoft]() {
					goto l488
				}
				add(ruleAuthorWord2, position489)
			}
			return true
		l488:
			position, tokenIndex = position488, tokenIndex488
			return false
		},
		/* 72 AuthorWord3 <- <(AuthorPrefixGlued? (AllCapsAuthorWord / CapAuthorWord) '.'?)> */
		func() bool {
			position490, tokenIndex490 := position, tokenIndex
			{
				position491 := position
				{
					position492, tokenIndex492 := position, tokenIndex
					if !_rules[ruleAuthorPrefixGlued]() {
						goto l492
					}
					goto l493
				l492:
					position, tokenIndex = position492, tokenIndex492
				}
			l493:
				{
					position494, tokenIndex494 := position, tokenIndex
					if !_rules[ruleAllCapsAuthorWord]() {
						goto l495
					}
					goto l494
				l495:
					position, tokenIndex = position494, tokenIndex494
					if !_rules[ruleCapAuthorWord]() {
						goto l490
					}
				}
			l494:
				{
					position496, tokenIndex496 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l496
					}
					position++
					goto l497
				l496:
					position, tokenIndex = position496, tokenIndex496
				}
			l497:
				add(ruleAuthorWord3, position491)
			}
			return true
		l490:
			position, tokenIndex = position490, tokenIndex490
			return false
		},
		/* 73 AuthorWordSoft <- <(((AuthorUpperChar (AuthorUpperChar+ / AuthorLowerChar+)) / AuthorLowerChar+) '.'?)> */
		func() bool {
			position498, tokenIndex498 := position, tokenIndex
			{
				position499 := position
				{
					position500, tokenIndex500 := position, tokenIndex
					if !_rules[ruleAuthorUpperChar]() {
						goto l501
					}
					{
						position502, tokenIndex502 := position, tokenIndex
						if !_rules[ruleAuthorUpperChar]() {
							goto l503
						}
					l504:
						{
							position505, tokenIndex505 := position, tokenIndex
							if !_rules[ruleAuthorUpperChar]() {
								goto l505
							}
							goto l504
						l505:
							position, tokenIndex = position505, tokenIndex505
						}
						goto l502
					l503:
						position, tokenIndex = position502, tokenIndex502
						if !_rules[ruleAuthorLowerChar]() {
							goto l501
						}
					l506:
						{
							position507, tokenIndex507 := position, tokenIndex
							if !_rules[ruleAuthorLowerChar]() {
								goto l507
							}
							goto l506
						l507:
							position, tokenIndex = position507, tokenIndex507
						}
					}
				l502:
					goto l500
				l501:
					position, tokenIndex = position500, tokenIndex500
					if !_rules[ruleAuthorLowerChar]() {
						goto l498
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
			l500:
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
				add(ruleAuthorWordSoft, position499)
			}
			return true
		l498:
			position, tokenIndex = position498, tokenIndex498
			return false
		},
		/* 74 CapAuthorWord <- <(AuthorUpperChar AuthorLowerChar*)> */
		func() bool {
			position512, tokenIndex512 := position, tokenIndex
			{
				position513 := position
				if !_rules[ruleAuthorUpperChar]() {
					goto l512
				}
			l514:
				{
					position515, tokenIndex515 := position, tokenIndex
					if !_rules[ruleAuthorLowerChar]() {
						goto l515
					}
					goto l514
				l515:
					position, tokenIndex = position515, tokenIndex515
				}
				add(ruleCapAuthorWord, position513)
			}
			return true
		l512:
			position, tokenIndex = position512, tokenIndex512
			return false
		},
		/* 75 AllCapsAuthorWord <- <(AuthorUpperChar AuthorUpperChar+)> */
		func() bool {
			position516, tokenIndex516 := position, tokenIndex
			{
				position517 := position
				if !_rules[ruleAuthorUpperChar]() {
					goto l516
				}
				if !_rules[ruleAuthorUpperChar]() {
					goto l516
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
				add(ruleAllCapsAuthorWord, position517)
			}
			return true
		l516:
			position, tokenIndex = position516, tokenIndex516
			return false
		},
		/* 76 Filius <- <(('f' '.') / ('f' 'i' 'l' '.') / ('f' 'i' 'l' 'i' 'u' 's'))> */
		func() bool {
			position520, tokenIndex520 := position, tokenIndex
			{
				position521 := position
				{
					position522, tokenIndex522 := position, tokenIndex
					if buffer[position] != rune('f') {
						goto l523
					}
					position++
					if buffer[position] != rune('.') {
						goto l523
					}
					position++
					goto l522
				l523:
					position, tokenIndex = position522, tokenIndex522
					if buffer[position] != rune('f') {
						goto l524
					}
					position++
					if buffer[position] != rune('i') {
						goto l524
					}
					position++
					if buffer[position] != rune('l') {
						goto l524
					}
					position++
					if buffer[position] != rune('.') {
						goto l524
					}
					position++
					goto l522
				l524:
					position, tokenIndex = position522, tokenIndex522
					if buffer[position] != rune('f') {
						goto l520
					}
					position++
					if buffer[position] != rune('i') {
						goto l520
					}
					position++
					if buffer[position] != rune('l') {
						goto l520
					}
					position++
					if buffer[position] != rune('i') {
						goto l520
					}
					position++
					if buffer[position] != rune('u') {
						goto l520
					}
					position++
					if buffer[position] != rune('s') {
						goto l520
					}
					position++
				}
			l522:
				add(ruleFilius, position521)
			}
			return true
		l520:
			position, tokenIndex = position520, tokenIndex520
			return false
		},
		/* 77 AuthorPrefixGlued <- <(('d' '\'') / ('O' '\''))> */
		func() bool {
			position525, tokenIndex525 := position, tokenIndex
			{
				position526 := position
				{
					position527, tokenIndex527 := position, tokenIndex
					if buffer[position] != rune('d') {
						goto l528
					}
					position++
					if buffer[position] != rune('\'') {
						goto l528
					}
					position++
					goto l527
				l528:
					position, tokenIndex = position527, tokenIndex527
					if buffer[position] != rune('O') {
						goto l525
					}
					position++
					if buffer[position] != rune('\'') {
						goto l525
					}
					position++
				}
			l527:
				add(ruleAuthorPrefixGlued, position526)
			}
			return true
		l525:
			position, tokenIndex = position525, tokenIndex525
			return false
		},
		/* 78 AuthorPrefix <- <(AuthorPrefix1 / AuthorPrefix2)> */
		func() bool {
			position529, tokenIndex529 := position, tokenIndex
			{
				position530 := position
				{
					position531, tokenIndex531 := position, tokenIndex
					if !_rules[ruleAuthorPrefix1]() {
						goto l532
					}
					goto l531
				l532:
					position, tokenIndex = position531, tokenIndex531
					if !_rules[ruleAuthorPrefix2]() {
						goto l529
					}
				}
			l531:
				add(ruleAuthorPrefix, position530)
			}
			return true
		l529:
			position, tokenIndex = position529, tokenIndex529
			return false
		},
		/* 79 AuthorPrefix2 <- <(('v' '.' (_? ('d' '.'))?) / ('\'' 't'))> */
		func() bool {
			position533, tokenIndex533 := position, tokenIndex
			{
				position534 := position
				{
					position535, tokenIndex535 := position, tokenIndex
					if buffer[position] != rune('v') {
						goto l536
					}
					position++
					if buffer[position] != rune('.') {
						goto l536
					}
					position++
					{
						position537, tokenIndex537 := position, tokenIndex
						{
							position539, tokenIndex539 := position, tokenIndex
							if !_rules[rule_]() {
								goto l539
							}
							goto l540
						l539:
							position, tokenIndex = position539, tokenIndex539
						}
					l540:
						if buffer[position] != rune('d') {
							goto l537
						}
						position++
						if buffer[position] != rune('.') {
							goto l537
						}
						position++
						goto l538
					l537:
						position, tokenIndex = position537, tokenIndex537
					}
				l538:
					goto l535
				l536:
					position, tokenIndex = position535, tokenIndex535
					if buffer[position] != rune('\'') {
						goto l533
					}
					position++
					if buffer[position] != rune('t') {
						goto l533
					}
					position++
				}
			l535:
				add(ruleAuthorPrefix2, position534)
			}
			return true
		l533:
			position, tokenIndex = position533, tokenIndex533
			return false
		},
		/* 80 AuthorPrefix1 <- <((('a' 'b') / ('a' 'f') / ('b' 'i' 's') / ('d' 'a') / ('d' 'e' 'r') / ('d' 'e' 's') / ('d' 'e' 'n') / ('d' 'e' 'l') / ('d' 'e' 'l' 'l' 'a') / ('d' 'e' 'l' 'a') / ('d' 'e') / ('d' 'i') / ('d' 'u') / ('e' 'l') / ('l' 'a') / ('l' 'e') / ('t' 'e' 'r') / ('v' 'a' 'n') / ('d' '\'') / ('i' 'n' '\'' 't') / ('z' 'u' 'r') / ('v' 'o' 'n' (_ (('d' '.') / ('d' 'e' 'm')))?) / ('v' (_ 'd')?)) &_)> */
		func() bool {
			position541, tokenIndex541 := position, tokenIndex
			{
				position542 := position
				{
					position543, tokenIndex543 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l544
					}
					position++
					if buffer[position] != rune('b') {
						goto l544
					}
					position++
					goto l543
				l544:
					position, tokenIndex = position543, tokenIndex543
					if buffer[position] != rune('a') {
						goto l545
					}
					position++
					if buffer[position] != rune('f') {
						goto l545
					}
					position++
					goto l543
				l545:
					position, tokenIndex = position543, tokenIndex543
					if buffer[position] != rune('b') {
						goto l546
					}
					position++
					if buffer[position] != rune('i') {
						goto l546
					}
					position++
					if buffer[position] != rune('s') {
						goto l546
					}
					position++
					goto l543
				l546:
					position, tokenIndex = position543, tokenIndex543
					if buffer[position] != rune('d') {
						goto l547
					}
					position++
					if buffer[position] != rune('a') {
						goto l547
					}
					position++
					goto l543
				l547:
					position, tokenIndex = position543, tokenIndex543
					if buffer[position] != rune('d') {
						goto l548
					}
					position++
					if buffer[position] != rune('e') {
						goto l548
					}
					position++
					if buffer[position] != rune('r') {
						goto l548
					}
					position++
					goto l543
				l548:
					position, tokenIndex = position543, tokenIndex543
					if buffer[position] != rune('d') {
						goto l549
					}
					position++
					if buffer[position] != rune('e') {
						goto l549
					}
					position++
					if buffer[position] != rune('s') {
						goto l549
					}
					position++
					goto l543
				l549:
					position, tokenIndex = position543, tokenIndex543
					if buffer[position] != rune('d') {
						goto l550
					}
					position++
					if buffer[position] != rune('e') {
						goto l550
					}
					position++
					if buffer[position] != rune('n') {
						goto l550
					}
					position++
					goto l543
				l550:
					position, tokenIndex = position543, tokenIndex543
					if buffer[position] != rune('d') {
						goto l551
					}
					position++
					if buffer[position] != rune('e') {
						goto l551
					}
					position++
					if buffer[position] != rune('l') {
						goto l551
					}
					position++
					goto l543
				l551:
					position, tokenIndex = position543, tokenIndex543
					if buffer[position] != rune('d') {
						goto l552
					}
					position++
					if buffer[position] != rune('e') {
						goto l552
					}
					position++
					if buffer[position] != rune('l') {
						goto l552
					}
					position++
					if buffer[position] != rune('l') {
						goto l552
					}
					position++
					if buffer[position] != rune('a') {
						goto l552
					}
					position++
					goto l543
				l552:
					position, tokenIndex = position543, tokenIndex543
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
					if buffer[position] != rune('a') {
						goto l553
					}
					position++
					goto l543
				l553:
					position, tokenIndex = position543, tokenIndex543
					if buffer[position] != rune('d') {
						goto l554
					}
					position++
					if buffer[position] != rune('e') {
						goto l554
					}
					position++
					goto l543
				l554:
					position, tokenIndex = position543, tokenIndex543
					if buffer[position] != rune('d') {
						goto l555
					}
					position++
					if buffer[position] != rune('i') {
						goto l555
					}
					position++
					goto l543
				l555:
					position, tokenIndex = position543, tokenIndex543
					if buffer[position] != rune('d') {
						goto l556
					}
					position++
					if buffer[position] != rune('u') {
						goto l556
					}
					position++
					goto l543
				l556:
					position, tokenIndex = position543, tokenIndex543
					if buffer[position] != rune('e') {
						goto l557
					}
					position++
					if buffer[position] != rune('l') {
						goto l557
					}
					position++
					goto l543
				l557:
					position, tokenIndex = position543, tokenIndex543
					if buffer[position] != rune('l') {
						goto l558
					}
					position++
					if buffer[position] != rune('a') {
						goto l558
					}
					position++
					goto l543
				l558:
					position, tokenIndex = position543, tokenIndex543
					if buffer[position] != rune('l') {
						goto l559
					}
					position++
					if buffer[position] != rune('e') {
						goto l559
					}
					position++
					goto l543
				l559:
					position, tokenIndex = position543, tokenIndex543
					if buffer[position] != rune('t') {
						goto l560
					}
					position++
					if buffer[position] != rune('e') {
						goto l560
					}
					position++
					if buffer[position] != rune('r') {
						goto l560
					}
					position++
					goto l543
				l560:
					position, tokenIndex = position543, tokenIndex543
					if buffer[position] != rune('v') {
						goto l561
					}
					position++
					if buffer[position] != rune('a') {
						goto l561
					}
					position++
					if buffer[position] != rune('n') {
						goto l561
					}
					position++
					goto l543
				l561:
					position, tokenIndex = position543, tokenIndex543
					if buffer[position] != rune('d') {
						goto l562
					}
					position++
					if buffer[position] != rune('\'') {
						goto l562
					}
					position++
					goto l543
				l562:
					position, tokenIndex = position543, tokenIndex543
					if buffer[position] != rune('i') {
						goto l563
					}
					position++
					if buffer[position] != rune('n') {
						goto l563
					}
					position++
					if buffer[position] != rune('\'') {
						goto l563
					}
					position++
					if buffer[position] != rune('t') {
						goto l563
					}
					position++
					goto l543
				l563:
					position, tokenIndex = position543, tokenIndex543
					if buffer[position] != rune('z') {
						goto l564
					}
					position++
					if buffer[position] != rune('u') {
						goto l564
					}
					position++
					if buffer[position] != rune('r') {
						goto l564
					}
					position++
					goto l543
				l564:
					position, tokenIndex = position543, tokenIndex543
					if buffer[position] != rune('v') {
						goto l565
					}
					position++
					if buffer[position] != rune('o') {
						goto l565
					}
					position++
					if buffer[position] != rune('n') {
						goto l565
					}
					position++
					{
						position566, tokenIndex566 := position, tokenIndex
						if !_rules[rule_]() {
							goto l566
						}
						{
							position568, tokenIndex568 := position, tokenIndex
							if buffer[position] != rune('d') {
								goto l569
							}
							position++
							if buffer[position] != rune('.') {
								goto l569
							}
							position++
							goto l568
						l569:
							position, tokenIndex = position568, tokenIndex568
							if buffer[position] != rune('d') {
								goto l566
							}
							position++
							if buffer[position] != rune('e') {
								goto l566
							}
							position++
							if buffer[position] != rune('m') {
								goto l566
							}
							position++
						}
					l568:
						goto l567
					l566:
						position, tokenIndex = position566, tokenIndex566
					}
				l567:
					goto l543
				l565:
					position, tokenIndex = position543, tokenIndex543
					if buffer[position] != rune('v') {
						goto l541
					}
					position++
					{
						position570, tokenIndex570 := position, tokenIndex
						if !_rules[rule_]() {
							goto l570
						}
						if buffer[position] != rune('d') {
							goto l570
						}
						position++
						goto l571
					l570:
						position, tokenIndex = position570, tokenIndex570
					}
				l571:
				}
			l543:
				{
					position572, tokenIndex572 := position, tokenIndex
					if !_rules[rule_]() {
						goto l541
					}
					position, tokenIndex = position572, tokenIndex572
				}
				add(ruleAuthorPrefix1, position542)
			}
			return true
		l541:
			position, tokenIndex = position541, tokenIndex541
			return false
		},
		/* 81 AuthorUpperChar <- <(hASCII / ('À' / 'Á' / 'Â' / 'Ã' / 'Ä' / 'Å' / 'Æ' / 'Ç' / 'È' / 'É' / 'Ê' / 'Ë' / 'Ì' / 'Í' / 'Î' / 'Ï' / 'Ð' / 'Ñ' / 'Ò' / 'Ó' / 'Ô' / 'Õ' / 'Ö' / 'Ø' / 'Ù' / 'Ú' / 'Û' / 'Ü' / 'Ý' / 'Ć' / 'Č' / 'Ď' / 'İ' / 'Ķ' / 'Ĺ' / 'ĺ' / 'Ľ' / 'ľ' / 'Ł' / 'ł' / 'Ņ' / 'Ō' / 'Ő' / 'Œ' / 'Ř' / 'Ś' / 'Ŝ' / 'Ş' / 'Š' / 'Ÿ' / 'Ź' / 'Ż' / 'Ž' / 'ƒ' / 'Ǿ' / 'Ș' / 'Ț'))> */
		func() bool {
			position573, tokenIndex573 := position, tokenIndex
			{
				position574 := position
				{
					position575, tokenIndex575 := position, tokenIndex
					if !_rules[rulehASCII]() {
						goto l576
					}
					goto l575
				l576:
					position, tokenIndex = position575, tokenIndex575
					{
						position577, tokenIndex577 := position, tokenIndex
						if buffer[position] != rune('À') {
							goto l578
						}
						position++
						goto l577
					l578:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Á') {
							goto l579
						}
						position++
						goto l577
					l579:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Â') {
							goto l580
						}
						position++
						goto l577
					l580:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ã') {
							goto l581
						}
						position++
						goto l577
					l581:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ä') {
							goto l582
						}
						position++
						goto l577
					l582:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Å') {
							goto l583
						}
						position++
						goto l577
					l583:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Æ') {
							goto l584
						}
						position++
						goto l577
					l584:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ç') {
							goto l585
						}
						position++
						goto l577
					l585:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('È') {
							goto l586
						}
						position++
						goto l577
					l586:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('É') {
							goto l587
						}
						position++
						goto l577
					l587:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ê') {
							goto l588
						}
						position++
						goto l577
					l588:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ë') {
							goto l589
						}
						position++
						goto l577
					l589:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ì') {
							goto l590
						}
						position++
						goto l577
					l590:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Í') {
							goto l591
						}
						position++
						goto l577
					l591:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Î') {
							goto l592
						}
						position++
						goto l577
					l592:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ï') {
							goto l593
						}
						position++
						goto l577
					l593:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ð') {
							goto l594
						}
						position++
						goto l577
					l594:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ñ') {
							goto l595
						}
						position++
						goto l577
					l595:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ò') {
							goto l596
						}
						position++
						goto l577
					l596:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ó') {
							goto l597
						}
						position++
						goto l577
					l597:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ô') {
							goto l598
						}
						position++
						goto l577
					l598:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Õ') {
							goto l599
						}
						position++
						goto l577
					l599:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ö') {
							goto l600
						}
						position++
						goto l577
					l600:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ø') {
							goto l601
						}
						position++
						goto l577
					l601:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ù') {
							goto l602
						}
						position++
						goto l577
					l602:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ú') {
							goto l603
						}
						position++
						goto l577
					l603:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Û') {
							goto l604
						}
						position++
						goto l577
					l604:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ü') {
							goto l605
						}
						position++
						goto l577
					l605:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ý') {
							goto l606
						}
						position++
						goto l577
					l606:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ć') {
							goto l607
						}
						position++
						goto l577
					l607:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Č') {
							goto l608
						}
						position++
						goto l577
					l608:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ď') {
							goto l609
						}
						position++
						goto l577
					l609:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('İ') {
							goto l610
						}
						position++
						goto l577
					l610:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ķ') {
							goto l611
						}
						position++
						goto l577
					l611:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ĺ') {
							goto l612
						}
						position++
						goto l577
					l612:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('ĺ') {
							goto l613
						}
						position++
						goto l577
					l613:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ľ') {
							goto l614
						}
						position++
						goto l577
					l614:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('ľ') {
							goto l615
						}
						position++
						goto l577
					l615:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ł') {
							goto l616
						}
						position++
						goto l577
					l616:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('ł') {
							goto l617
						}
						position++
						goto l577
					l617:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ņ') {
							goto l618
						}
						position++
						goto l577
					l618:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ō') {
							goto l619
						}
						position++
						goto l577
					l619:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ő') {
							goto l620
						}
						position++
						goto l577
					l620:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Œ') {
							goto l621
						}
						position++
						goto l577
					l621:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ř') {
							goto l622
						}
						position++
						goto l577
					l622:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ś') {
							goto l623
						}
						position++
						goto l577
					l623:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ŝ') {
							goto l624
						}
						position++
						goto l577
					l624:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ş') {
							goto l625
						}
						position++
						goto l577
					l625:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Š') {
							goto l626
						}
						position++
						goto l577
					l626:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ÿ') {
							goto l627
						}
						position++
						goto l577
					l627:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ź') {
							goto l628
						}
						position++
						goto l577
					l628:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ż') {
							goto l629
						}
						position++
						goto l577
					l629:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ž') {
							goto l630
						}
						position++
						goto l577
					l630:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('ƒ') {
							goto l631
						}
						position++
						goto l577
					l631:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ǿ') {
							goto l632
						}
						position++
						goto l577
					l632:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ș') {
							goto l633
						}
						position++
						goto l577
					l633:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('Ț') {
							goto l573
						}
						position++
					}
				l577:
				}
			l575:
				add(ruleAuthorUpperChar, position574)
			}
			return true
		l573:
			position, tokenIndex = position573, tokenIndex573
			return false
		},
		/* 82 AuthorLowerChar <- <(lASCII / ('à' / 'á' / 'â' / 'ã' / 'ä' / 'å' / 'æ' / 'ç' / 'è' / 'é' / 'ê' / 'ë' / 'ì' / 'í' / 'î' / 'ï' / 'ð' / 'ñ' / 'ò' / 'ó' / 'ó' / 'ô' / 'õ' / 'ö' / 'ø' / 'ù' / 'ú' / 'û' / 'ü' / 'ý' / 'ÿ' / 'ā' / 'ă' / 'ą' / 'ć' / 'ĉ' / 'č' / 'ď' / 'đ' / '\'' / 'ē' / 'ĕ' / 'ė' / 'ę' / 'ě' / 'ğ' / 'ī' / 'ĭ' / 'İ' / 'ı' / 'ĺ' / 'ľ' / 'ł' / 'ń' / 'ņ' / 'ň' / 'ŏ' / 'ő' / 'œ' / 'ŕ' / 'ř' / 'ś' / 'ş' / 'š' / 'ţ' / 'ť' / 'ũ' / 'ū' / 'ŭ' / 'ů' / 'ű' / 'ź' / 'ż' / 'ž' / 'ſ' / 'ǎ' / 'ǔ' / 'ǧ' / 'ș' / 'ț' / 'ȳ' / 'ß'))> */
		func() bool {
			position634, tokenIndex634 := position, tokenIndex
			{
				position635 := position
				{
					position636, tokenIndex636 := position, tokenIndex
					if !_rules[rulelASCII]() {
						goto l637
					}
					goto l636
				l637:
					position, tokenIndex = position636, tokenIndex636
					{
						position638, tokenIndex638 := position, tokenIndex
						if buffer[position] != rune('à') {
							goto l639
						}
						position++
						goto l638
					l639:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('á') {
							goto l640
						}
						position++
						goto l638
					l640:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('â') {
							goto l641
						}
						position++
						goto l638
					l641:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ã') {
							goto l642
						}
						position++
						goto l638
					l642:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ä') {
							goto l643
						}
						position++
						goto l638
					l643:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('å') {
							goto l644
						}
						position++
						goto l638
					l644:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('æ') {
							goto l645
						}
						position++
						goto l638
					l645:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ç') {
							goto l646
						}
						position++
						goto l638
					l646:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('è') {
							goto l647
						}
						position++
						goto l638
					l647:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('é') {
							goto l648
						}
						position++
						goto l638
					l648:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ê') {
							goto l649
						}
						position++
						goto l638
					l649:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ë') {
							goto l650
						}
						position++
						goto l638
					l650:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ì') {
							goto l651
						}
						position++
						goto l638
					l651:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('í') {
							goto l652
						}
						position++
						goto l638
					l652:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('î') {
							goto l653
						}
						position++
						goto l638
					l653:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ï') {
							goto l654
						}
						position++
						goto l638
					l654:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ð') {
							goto l655
						}
						position++
						goto l638
					l655:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ñ') {
							goto l656
						}
						position++
						goto l638
					l656:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ò') {
							goto l657
						}
						position++
						goto l638
					l657:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ó') {
							goto l658
						}
						position++
						goto l638
					l658:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ó') {
							goto l659
						}
						position++
						goto l638
					l659:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ô') {
							goto l660
						}
						position++
						goto l638
					l660:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('õ') {
							goto l661
						}
						position++
						goto l638
					l661:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ö') {
							goto l662
						}
						position++
						goto l638
					l662:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ø') {
							goto l663
						}
						position++
						goto l638
					l663:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ù') {
							goto l664
						}
						position++
						goto l638
					l664:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ú') {
							goto l665
						}
						position++
						goto l638
					l665:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('û') {
							goto l666
						}
						position++
						goto l638
					l666:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ü') {
							goto l667
						}
						position++
						goto l638
					l667:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ý') {
							goto l668
						}
						position++
						goto l638
					l668:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ÿ') {
							goto l669
						}
						position++
						goto l638
					l669:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ā') {
							goto l670
						}
						position++
						goto l638
					l670:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ă') {
							goto l671
						}
						position++
						goto l638
					l671:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ą') {
							goto l672
						}
						position++
						goto l638
					l672:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ć') {
							goto l673
						}
						position++
						goto l638
					l673:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ĉ') {
							goto l674
						}
						position++
						goto l638
					l674:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('č') {
							goto l675
						}
						position++
						goto l638
					l675:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ď') {
							goto l676
						}
						position++
						goto l638
					l676:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('đ') {
							goto l677
						}
						position++
						goto l638
					l677:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('\'') {
							goto l678
						}
						position++
						goto l638
					l678:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ē') {
							goto l679
						}
						position++
						goto l638
					l679:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ĕ') {
							goto l680
						}
						position++
						goto l638
					l680:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ė') {
							goto l681
						}
						position++
						goto l638
					l681:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ę') {
							goto l682
						}
						position++
						goto l638
					l682:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ě') {
							goto l683
						}
						position++
						goto l638
					l683:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ğ') {
							goto l684
						}
						position++
						goto l638
					l684:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ī') {
							goto l685
						}
						position++
						goto l638
					l685:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ĭ') {
							goto l686
						}
						position++
						goto l638
					l686:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('İ') {
							goto l687
						}
						position++
						goto l638
					l687:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ı') {
							goto l688
						}
						position++
						goto l638
					l688:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ĺ') {
							goto l689
						}
						position++
						goto l638
					l689:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ľ') {
							goto l690
						}
						position++
						goto l638
					l690:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ł') {
							goto l691
						}
						position++
						goto l638
					l691:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ń') {
							goto l692
						}
						position++
						goto l638
					l692:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ņ') {
							goto l693
						}
						position++
						goto l638
					l693:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ň') {
							goto l694
						}
						position++
						goto l638
					l694:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ŏ') {
							goto l695
						}
						position++
						goto l638
					l695:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ő') {
							goto l696
						}
						position++
						goto l638
					l696:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('œ') {
							goto l697
						}
						position++
						goto l638
					l697:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ŕ') {
							goto l698
						}
						position++
						goto l638
					l698:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ř') {
							goto l699
						}
						position++
						goto l638
					l699:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ś') {
							goto l700
						}
						position++
						goto l638
					l700:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ş') {
							goto l701
						}
						position++
						goto l638
					l701:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('š') {
							goto l702
						}
						position++
						goto l638
					l702:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ţ') {
							goto l703
						}
						position++
						goto l638
					l703:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ť') {
							goto l704
						}
						position++
						goto l638
					l704:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ũ') {
							goto l705
						}
						position++
						goto l638
					l705:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ū') {
							goto l706
						}
						position++
						goto l638
					l706:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ŭ') {
							goto l707
						}
						position++
						goto l638
					l707:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ů') {
							goto l708
						}
						position++
						goto l638
					l708:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ű') {
							goto l709
						}
						position++
						goto l638
					l709:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ź') {
							goto l710
						}
						position++
						goto l638
					l710:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ż') {
							goto l711
						}
						position++
						goto l638
					l711:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ž') {
							goto l712
						}
						position++
						goto l638
					l712:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ſ') {
							goto l713
						}
						position++
						goto l638
					l713:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ǎ') {
							goto l714
						}
						position++
						goto l638
					l714:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ǔ') {
							goto l715
						}
						position++
						goto l638
					l715:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ǧ') {
							goto l716
						}
						position++
						goto l638
					l716:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ș') {
							goto l717
						}
						position++
						goto l638
					l717:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ț') {
							goto l718
						}
						position++
						goto l638
					l718:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ȳ') {
							goto l719
						}
						position++
						goto l638
					l719:
						position, tokenIndex = position638, tokenIndex638
						if buffer[position] != rune('ß') {
							goto l634
						}
						position++
					}
				l638:
				}
			l636:
				add(ruleAuthorLowerChar, position635)
			}
			return true
		l634:
			position, tokenIndex = position634, tokenIndex634
			return false
		},
		/* 83 Year <- <(YearRange / YearApprox / YearWithParens / YearWithPage / YearWithDot / YearWithChar / YearNum)> */
		func() bool {
			position720, tokenIndex720 := position, tokenIndex
			{
				position721 := position
				{
					position722, tokenIndex722 := position, tokenIndex
					if !_rules[ruleYearRange]() {
						goto l723
					}
					goto l722
				l723:
					position, tokenIndex = position722, tokenIndex722
					if !_rules[ruleYearApprox]() {
						goto l724
					}
					goto l722
				l724:
					position, tokenIndex = position722, tokenIndex722
					if !_rules[ruleYearWithParens]() {
						goto l725
					}
					goto l722
				l725:
					position, tokenIndex = position722, tokenIndex722
					if !_rules[ruleYearWithPage]() {
						goto l726
					}
					goto l722
				l726:
					position, tokenIndex = position722, tokenIndex722
					if !_rules[ruleYearWithDot]() {
						goto l727
					}
					goto l722
				l727:
					position, tokenIndex = position722, tokenIndex722
					if !_rules[ruleYearWithChar]() {
						goto l728
					}
					goto l722
				l728:
					position, tokenIndex = position722, tokenIndex722
					if !_rules[ruleYearNum]() {
						goto l720
					}
				}
			l722:
				add(ruleYear, position721)
			}
			return true
		l720:
			position, tokenIndex = position720, tokenIndex720
			return false
		},
		/* 84 YearRange <- <(YearNum dash (nums+ ('a' / 'b' / 'c' / 'd' / 'e' / 'f' / 'g' / 'h' / 'i' / 'j' / 'k' / 'l' / 'm' / 'n' / 'o' / 'p' / 'q' / 'r' / 's' / 't' / 'u' / 'v' / 'w' / 'x' / 'y' / 'z' / '?')*))> */
		func() bool {
			position729, tokenIndex729 := position, tokenIndex
			{
				position730 := position
				if !_rules[ruleYearNum]() {
					goto l729
				}
				if !_rules[ruledash]() {
					goto l729
				}
				if !_rules[rulenums]() {
					goto l729
				}
			l731:
				{
					position732, tokenIndex732 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l732
					}
					goto l731
				l732:
					position, tokenIndex = position732, tokenIndex732
				}
			l733:
				{
					position734, tokenIndex734 := position, tokenIndex
					{
						position735, tokenIndex735 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l736
						}
						position++
						goto l735
					l736:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('b') {
							goto l737
						}
						position++
						goto l735
					l737:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('c') {
							goto l738
						}
						position++
						goto l735
					l738:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('d') {
							goto l739
						}
						position++
						goto l735
					l739:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('e') {
							goto l740
						}
						position++
						goto l735
					l740:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('f') {
							goto l741
						}
						position++
						goto l735
					l741:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('g') {
							goto l742
						}
						position++
						goto l735
					l742:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('h') {
							goto l743
						}
						position++
						goto l735
					l743:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('i') {
							goto l744
						}
						position++
						goto l735
					l744:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('j') {
							goto l745
						}
						position++
						goto l735
					l745:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('k') {
							goto l746
						}
						position++
						goto l735
					l746:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('l') {
							goto l747
						}
						position++
						goto l735
					l747:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('m') {
							goto l748
						}
						position++
						goto l735
					l748:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('n') {
							goto l749
						}
						position++
						goto l735
					l749:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('o') {
							goto l750
						}
						position++
						goto l735
					l750:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('p') {
							goto l751
						}
						position++
						goto l735
					l751:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('q') {
							goto l752
						}
						position++
						goto l735
					l752:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('r') {
							goto l753
						}
						position++
						goto l735
					l753:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('s') {
							goto l754
						}
						position++
						goto l735
					l754:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('t') {
							goto l755
						}
						position++
						goto l735
					l755:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('u') {
							goto l756
						}
						position++
						goto l735
					l756:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('v') {
							goto l757
						}
						position++
						goto l735
					l757:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('w') {
							goto l758
						}
						position++
						goto l735
					l758:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('x') {
							goto l759
						}
						position++
						goto l735
					l759:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('y') {
							goto l760
						}
						position++
						goto l735
					l760:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('z') {
							goto l761
						}
						position++
						goto l735
					l761:
						position, tokenIndex = position735, tokenIndex735
						if buffer[position] != rune('?') {
							goto l734
						}
						position++
					}
				l735:
					goto l733
				l734:
					position, tokenIndex = position734, tokenIndex734
				}
				add(ruleYearRange, position730)
			}
			return true
		l729:
			position, tokenIndex = position729, tokenIndex729
			return false
		},
		/* 85 YearWithDot <- <(YearNum '.')> */
		func() bool {
			position762, tokenIndex762 := position, tokenIndex
			{
				position763 := position
				if !_rules[ruleYearNum]() {
					goto l762
				}
				if buffer[position] != rune('.') {
					goto l762
				}
				position++
				add(ruleYearWithDot, position763)
			}
			return true
		l762:
			position, tokenIndex = position762, tokenIndex762
			return false
		},
		/* 86 YearApprox <- <('[' _? YearNum _? ']')> */
		func() bool {
			position764, tokenIndex764 := position, tokenIndex
			{
				position765 := position
				if buffer[position] != rune('[') {
					goto l764
				}
				position++
				{
					position766, tokenIndex766 := position, tokenIndex
					if !_rules[rule_]() {
						goto l766
					}
					goto l767
				l766:
					position, tokenIndex = position766, tokenIndex766
				}
			l767:
				if !_rules[ruleYearNum]() {
					goto l764
				}
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
				if buffer[position] != rune(']') {
					goto l764
				}
				position++
				add(ruleYearApprox, position765)
			}
			return true
		l764:
			position, tokenIndex = position764, tokenIndex764
			return false
		},
		/* 87 YearWithPage <- <((YearWithChar / YearNum) _? ':' _? nums+)> */
		func() bool {
			position770, tokenIndex770 := position, tokenIndex
			{
				position771 := position
				{
					position772, tokenIndex772 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l773
					}
					goto l772
				l773:
					position, tokenIndex = position772, tokenIndex772
					if !_rules[ruleYearNum]() {
						goto l770
					}
				}
			l772:
				{
					position774, tokenIndex774 := position, tokenIndex
					if !_rules[rule_]() {
						goto l774
					}
					goto l775
				l774:
					position, tokenIndex = position774, tokenIndex774
				}
			l775:
				if buffer[position] != rune(':') {
					goto l770
				}
				position++
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
				if !_rules[rulenums]() {
					goto l770
				}
			l778:
				{
					position779, tokenIndex779 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l779
					}
					goto l778
				l779:
					position, tokenIndex = position779, tokenIndex779
				}
				add(ruleYearWithPage, position771)
			}
			return true
		l770:
			position, tokenIndex = position770, tokenIndex770
			return false
		},
		/* 88 YearWithParens <- <('(' (YearWithChar / YearNum) ')')> */
		func() bool {
			position780, tokenIndex780 := position, tokenIndex
			{
				position781 := position
				if buffer[position] != rune('(') {
					goto l780
				}
				position++
				{
					position782, tokenIndex782 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l783
					}
					goto l782
				l783:
					position, tokenIndex = position782, tokenIndex782
					if !_rules[ruleYearNum]() {
						goto l780
					}
				}
			l782:
				if buffer[position] != rune(')') {
					goto l780
				}
				position++
				add(ruleYearWithParens, position781)
			}
			return true
		l780:
			position, tokenIndex = position780, tokenIndex780
			return false
		},
		/* 89 YearWithChar <- <(YearNum lASCII Action0)> */
		func() bool {
			position784, tokenIndex784 := position, tokenIndex
			{
				position785 := position
				if !_rules[ruleYearNum]() {
					goto l784
				}
				if !_rules[rulelASCII]() {
					goto l784
				}
				if !_rules[ruleAction0]() {
					goto l784
				}
				add(ruleYearWithChar, position785)
			}
			return true
		l784:
			position, tokenIndex = position784, tokenIndex784
			return false
		},
		/* 90 YearNum <- <(('1' / '2') ('0' / '7' / '8' / '9') nums (nums / '?') '?'*)> */
		func() bool {
			position786, tokenIndex786 := position, tokenIndex
			{
				position787 := position
				{
					position788, tokenIndex788 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l789
					}
					position++
					goto l788
				l789:
					position, tokenIndex = position788, tokenIndex788
					if buffer[position] != rune('2') {
						goto l786
					}
					position++
				}
			l788:
				{
					position790, tokenIndex790 := position, tokenIndex
					if buffer[position] != rune('0') {
						goto l791
					}
					position++
					goto l790
				l791:
					position, tokenIndex = position790, tokenIndex790
					if buffer[position] != rune('7') {
						goto l792
					}
					position++
					goto l790
				l792:
					position, tokenIndex = position790, tokenIndex790
					if buffer[position] != rune('8') {
						goto l793
					}
					position++
					goto l790
				l793:
					position, tokenIndex = position790, tokenIndex790
					if buffer[position] != rune('9') {
						goto l786
					}
					position++
				}
			l790:
				if !_rules[rulenums]() {
					goto l786
				}
				{
					position794, tokenIndex794 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l795
					}
					goto l794
				l795:
					position, tokenIndex = position794, tokenIndex794
					if buffer[position] != rune('?') {
						goto l786
					}
					position++
				}
			l794:
			l796:
				{
					position797, tokenIndex797 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l797
					}
					position++
					goto l796
				l797:
					position, tokenIndex = position797, tokenIndex797
				}
				add(ruleYearNum, position787)
			}
			return true
		l786:
			position, tokenIndex = position786, tokenIndex786
			return false
		},
		/* 91 NameUpperChar <- <(UpperChar / UpperCharExtended)> */
		func() bool {
			position798, tokenIndex798 := position, tokenIndex
			{
				position799 := position
				{
					position800, tokenIndex800 := position, tokenIndex
					if !_rules[ruleUpperChar]() {
						goto l801
					}
					goto l800
				l801:
					position, tokenIndex = position800, tokenIndex800
					if !_rules[ruleUpperCharExtended]() {
						goto l798
					}
				}
			l800:
				add(ruleNameUpperChar, position799)
			}
			return true
		l798:
			position, tokenIndex = position798, tokenIndex798
			return false
		},
		/* 92 UpperCharExtended <- <('Æ' / 'Œ' / 'Ö')> */
		func() bool {
			position802, tokenIndex802 := position, tokenIndex
			{
				position803 := position
				{
					position804, tokenIndex804 := position, tokenIndex
					if buffer[position] != rune('Æ') {
						goto l805
					}
					position++
					goto l804
				l805:
					position, tokenIndex = position804, tokenIndex804
					if buffer[position] != rune('Œ') {
						goto l806
					}
					position++
					goto l804
				l806:
					position, tokenIndex = position804, tokenIndex804
					if buffer[position] != rune('Ö') {
						goto l802
					}
					position++
				}
			l804:
				add(ruleUpperCharExtended, position803)
			}
			return true
		l802:
			position, tokenIndex = position802, tokenIndex802
			return false
		},
		/* 93 UpperChar <- <hASCII> */
		func() bool {
			position807, tokenIndex807 := position, tokenIndex
			{
				position808 := position
				if !_rules[rulehASCII]() {
					goto l807
				}
				add(ruleUpperChar, position808)
			}
			return true
		l807:
			position, tokenIndex = position807, tokenIndex807
			return false
		},
		/* 94 NameLowerChar <- <(LowerChar / LowerCharExtended / MiscodedChar)> */
		func() bool {
			position809, tokenIndex809 := position, tokenIndex
			{
				position810 := position
				{
					position811, tokenIndex811 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l812
					}
					goto l811
				l812:
					position, tokenIndex = position811, tokenIndex811
					if !_rules[ruleLowerCharExtended]() {
						goto l813
					}
					goto l811
				l813:
					position, tokenIndex = position811, tokenIndex811
					if !_rules[ruleMiscodedChar]() {
						goto l809
					}
				}
			l811:
				add(ruleNameLowerChar, position810)
			}
			return true
		l809:
			position, tokenIndex = position809, tokenIndex809
			return false
		},
		/* 95 MiscodedChar <- <'�'> */
		func() bool {
			position814, tokenIndex814 := position, tokenIndex
			{
				position815 := position
				if buffer[position] != rune('�') {
					goto l814
				}
				position++
				add(ruleMiscodedChar, position815)
			}
			return true
		l814:
			position, tokenIndex = position814, tokenIndex814
			return false
		},
		/* 96 LowerCharExtended <- <('æ' / 'œ' / 'ſ' / 'à' / 'â' / 'å' / 'ã' / 'ä' / 'á' / 'ç' / 'č' / 'é' / 'è' / 'ë' / 'í' / 'ì' / 'ï' / 'ň' / 'ñ' / 'ñ' / 'ó' / 'ò' / 'ô' / 'ø' / 'õ' / 'ö' / 'ú' / 'ù' / 'ü' / 'ŕ' / 'ř' / 'ŗ' / 'š' / 'š' / 'ş' / 'ž')> */
		func() bool {
			position816, tokenIndex816 := position, tokenIndex
			{
				position817 := position
				{
					position818, tokenIndex818 := position, tokenIndex
					if buffer[position] != rune('æ') {
						goto l819
					}
					position++
					goto l818
				l819:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('œ') {
						goto l820
					}
					position++
					goto l818
				l820:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('ſ') {
						goto l821
					}
					position++
					goto l818
				l821:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('à') {
						goto l822
					}
					position++
					goto l818
				l822:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('â') {
						goto l823
					}
					position++
					goto l818
				l823:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('å') {
						goto l824
					}
					position++
					goto l818
				l824:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('ã') {
						goto l825
					}
					position++
					goto l818
				l825:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('ä') {
						goto l826
					}
					position++
					goto l818
				l826:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('á') {
						goto l827
					}
					position++
					goto l818
				l827:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('ç') {
						goto l828
					}
					position++
					goto l818
				l828:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('č') {
						goto l829
					}
					position++
					goto l818
				l829:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('é') {
						goto l830
					}
					position++
					goto l818
				l830:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('è') {
						goto l831
					}
					position++
					goto l818
				l831:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('ë') {
						goto l832
					}
					position++
					goto l818
				l832:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('í') {
						goto l833
					}
					position++
					goto l818
				l833:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('ì') {
						goto l834
					}
					position++
					goto l818
				l834:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('ï') {
						goto l835
					}
					position++
					goto l818
				l835:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('ň') {
						goto l836
					}
					position++
					goto l818
				l836:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('ñ') {
						goto l837
					}
					position++
					goto l818
				l837:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('ñ') {
						goto l838
					}
					position++
					goto l818
				l838:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('ó') {
						goto l839
					}
					position++
					goto l818
				l839:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('ò') {
						goto l840
					}
					position++
					goto l818
				l840:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('ô') {
						goto l841
					}
					position++
					goto l818
				l841:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('ø') {
						goto l842
					}
					position++
					goto l818
				l842:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('õ') {
						goto l843
					}
					position++
					goto l818
				l843:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('ö') {
						goto l844
					}
					position++
					goto l818
				l844:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('ú') {
						goto l845
					}
					position++
					goto l818
				l845:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('ù') {
						goto l846
					}
					position++
					goto l818
				l846:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('ü') {
						goto l847
					}
					position++
					goto l818
				l847:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('ŕ') {
						goto l848
					}
					position++
					goto l818
				l848:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('ř') {
						goto l849
					}
					position++
					goto l818
				l849:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('ŗ') {
						goto l850
					}
					position++
					goto l818
				l850:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('š') {
						goto l851
					}
					position++
					goto l818
				l851:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('š') {
						goto l852
					}
					position++
					goto l818
				l852:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('ş') {
						goto l853
					}
					position++
					goto l818
				l853:
					position, tokenIndex = position818, tokenIndex818
					if buffer[position] != rune('ž') {
						goto l816
					}
					position++
				}
			l818:
				add(ruleLowerCharExtended, position817)
			}
			return true
		l816:
			position, tokenIndex = position816, tokenIndex816
			return false
		},
		/* 97 LowerChar <- <lASCII> */
		func() bool {
			position854, tokenIndex854 := position, tokenIndex
			{
				position855 := position
				if !_rules[rulelASCII]() {
					goto l854
				}
				add(ruleLowerChar, position855)
			}
			return true
		l854:
			position, tokenIndex = position854, tokenIndex854
			return false
		},
		/* 98 SpaceCharEOI <- <(_ / !.)> */
		func() bool {
			position856, tokenIndex856 := position, tokenIndex
			{
				position857 := position
				{
					position858, tokenIndex858 := position, tokenIndex
					if !_rules[rule_]() {
						goto l859
					}
					goto l858
				l859:
					position, tokenIndex = position858, tokenIndex858
					{
						position860, tokenIndex860 := position, tokenIndex
						if !matchDot() {
							goto l860
						}
						goto l856
					l860:
						position, tokenIndex = position860, tokenIndex860
					}
				}
			l858:
				add(ruleSpaceCharEOI, position857)
			}
			return true
		l856:
			position, tokenIndex = position856, tokenIndex856
			return false
		},
		/* 99 WordBorderChar <- <(_ / (';' / '.' / ',' / ';' / '(' / ')'))> */
		nil,
		/* 100 nums <- <[0-9]> */
		func() bool {
			position862, tokenIndex862 := position, tokenIndex
			{
				position863 := position
				if c := buffer[position]; c < rune('0') || c > rune('9') {
					goto l862
				}
				position++
				add(rulenums, position863)
			}
			return true
		l862:
			position, tokenIndex = position862, tokenIndex862
			return false
		},
		/* 101 lASCII <- <[a-z]> */
		func() bool {
			position864, tokenIndex864 := position, tokenIndex
			{
				position865 := position
				if c := buffer[position]; c < rune('a') || c > rune('z') {
					goto l864
				}
				position++
				add(rulelASCII, position865)
			}
			return true
		l864:
			position, tokenIndex = position864, tokenIndex864
			return false
		},
		/* 102 hASCII <- <[A-Z]> */
		func() bool {
			position866, tokenIndex866 := position, tokenIndex
			{
				position867 := position
				if c := buffer[position]; c < rune('A') || c > rune('Z') {
					goto l866
				}
				position++
				add(rulehASCII, position867)
			}
			return true
		l866:
			position, tokenIndex = position866, tokenIndex866
			return false
		},
		/* 103 apostr <- <'\''> */
		func() bool {
			position868, tokenIndex868 := position, tokenIndex
			{
				position869 := position
				if buffer[position] != rune('\'') {
					goto l868
				}
				position++
				add(ruleapostr, position869)
			}
			return true
		l868:
			position, tokenIndex = position868, tokenIndex868
			return false
		},
		/* 104 dash <- <'-'> */
		func() bool {
			position870, tokenIndex870 := position, tokenIndex
			{
				position871 := position
				if buffer[position] != rune('-') {
					goto l870
				}
				position++
				add(ruledash, position871)
			}
			return true
		l870:
			position, tokenIndex = position870, tokenIndex870
			return false
		},
		/* 105 _ <- <(MultipleSpace / SingleSpace)> */
		func() bool {
			position872, tokenIndex872 := position, tokenIndex
			{
				position873 := position
				{
					position874, tokenIndex874 := position, tokenIndex
					if !_rules[ruleMultipleSpace]() {
						goto l875
					}
					goto l874
				l875:
					position, tokenIndex = position874, tokenIndex874
					if !_rules[ruleSingleSpace]() {
						goto l872
					}
				}
			l874:
				add(rule_, position873)
			}
			return true
		l872:
			position, tokenIndex = position872, tokenIndex872
			return false
		},
		/* 106 MultipleSpace <- <(SingleSpace SingleSpace+)> */
		func() bool {
			position876, tokenIndex876 := position, tokenIndex
			{
				position877 := position
				if !_rules[ruleSingleSpace]() {
					goto l876
				}
				if !_rules[ruleSingleSpace]() {
					goto l876
				}
			l878:
				{
					position879, tokenIndex879 := position, tokenIndex
					if !_rules[ruleSingleSpace]() {
						goto l879
					}
					goto l878
				l879:
					position, tokenIndex = position879, tokenIndex879
				}
				add(ruleMultipleSpace, position877)
			}
			return true
		l876:
			position, tokenIndex = position876, tokenIndex876
			return false
		},
		/* 107 SingleSpace <- <' '> */
		func() bool {
			position880, tokenIndex880 := position, tokenIndex
			{
				position881 := position
				if buffer[position] != rune(' ') {
					goto l880
				}
				position++
				add(ruleSingleSpace, position881)
			}
			return true
		l880:
			position, tokenIndex = position880, tokenIndex880
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
