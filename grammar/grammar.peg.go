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
		/* 47 Authorship <- <((AuthorshipCombo / OriginalAuthorship) &(SpaceCharEOI / ';' / ','))> */
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
						if buffer[position] != rune(';') {
							goto l361
						}
						position++
						goto l359
					l361:
						position, tokenIndex = position359, tokenIndex359
						if buffer[position] != rune(',') {
							goto l354
						}
						position++
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
		/* 48 AuthorshipCombo <- <(OriginalAuthorshipComb (_? CombinationAuthorship)?)> */
		func() bool {
			position362, tokenIndex362 := position, tokenIndex
			{
				position363 := position
				if !_rules[ruleOriginalAuthorshipComb]() {
					goto l362
				}
				{
					position364, tokenIndex364 := position, tokenIndex
					{
						position366, tokenIndex366 := position, tokenIndex
						if !_rules[rule_]() {
							goto l366
						}
						goto l367
					l366:
						position, tokenIndex = position366, tokenIndex366
					}
				l367:
					if !_rules[ruleCombinationAuthorship]() {
						goto l364
					}
					goto l365
				l364:
					position, tokenIndex = position364, tokenIndex364
				}
			l365:
				add(ruleAuthorshipCombo, position363)
			}
			return true
		l362:
			position, tokenIndex = position362, tokenIndex362
			return false
		},
		/* 49 OriginalAuthorship <- <AuthorsGroup> */
		func() bool {
			position368, tokenIndex368 := position, tokenIndex
			{
				position369 := position
				if !_rules[ruleAuthorsGroup]() {
					goto l368
				}
				add(ruleOriginalAuthorship, position369)
			}
			return true
		l368:
			position, tokenIndex = position368, tokenIndex368
			return false
		},
		/* 50 OriginalAuthorshipComb <- <(BasionymAuthorshipYearMisformed / BasionymAuthorship)> */
		func() bool {
			position370, tokenIndex370 := position, tokenIndex
			{
				position371 := position
				{
					position372, tokenIndex372 := position, tokenIndex
					if !_rules[ruleBasionymAuthorshipYearMisformed]() {
						goto l373
					}
					goto l372
				l373:
					position, tokenIndex = position372, tokenIndex372
					if !_rules[ruleBasionymAuthorship]() {
						goto l370
					}
				}
			l372:
				add(ruleOriginalAuthorshipComb, position371)
			}
			return true
		l370:
			position, tokenIndex = position370, tokenIndex370
			return false
		},
		/* 51 CombinationAuthorship <- <AuthorsGroup> */
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
		/* 52 BasionymAuthorshipYearMisformed <- <('(' _? AuthorsGroup _? ')' (_? ',')? _? Year)> */
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
		/* 53 BasionymAuthorship <- <(BasionymAuthorship1 / BasionymAuthorship2Parens)> */
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
					if !_rules[ruleBasionymAuthorship2Parens]() {
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
		/* 54 BasionymAuthorship1 <- <('(' _? AuthorsGroup _? ')')> */
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
		/* 55 BasionymAuthorship2Parens <- <('(' _? '(' _? AuthorsGroup _? ')' _? ')')> */
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
				add(ruleBasionymAuthorship2Parens, position399)
			}
			return true
		l398:
			position, tokenIndex = position398, tokenIndex398
			return false
		},
		/* 56 AuthorsGroup <- <(AuthorsTeam (_ (AuthorEmend / AuthorEx) AuthorsTeam)?)> */
		func() bool {
			position408, tokenIndex408 := position, tokenIndex
			{
				position409 := position
				if !_rules[ruleAuthorsTeam]() {
					goto l408
				}
				{
					position410, tokenIndex410 := position, tokenIndex
					if !_rules[rule_]() {
						goto l410
					}
					{
						position412, tokenIndex412 := position, tokenIndex
						if !_rules[ruleAuthorEmend]() {
							goto l413
						}
						goto l412
					l413:
						position, tokenIndex = position412, tokenIndex412
						if !_rules[ruleAuthorEx]() {
							goto l410
						}
					}
				l412:
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
		/* 57 AuthorsTeam <- <(Author (AuthorSep Author)* (_? ','? _? Year)?)> */
		func() bool {
			position414, tokenIndex414 := position, tokenIndex
			{
				position415 := position
				if !_rules[ruleAuthor]() {
					goto l414
				}
			l416:
				{
					position417, tokenIndex417 := position, tokenIndex
					if !_rules[ruleAuthorSep]() {
						goto l417
					}
					if !_rules[ruleAuthor]() {
						goto l417
					}
					goto l416
				l417:
					position, tokenIndex = position417, tokenIndex417
				}
				{
					position418, tokenIndex418 := position, tokenIndex
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
					{
						position422, tokenIndex422 := position, tokenIndex
						if buffer[position] != rune(',') {
							goto l422
						}
						position++
						goto l423
					l422:
						position, tokenIndex = position422, tokenIndex422
					}
				l423:
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
					if !_rules[ruleYear]() {
						goto l418
					}
					goto l419
				l418:
					position, tokenIndex = position418, tokenIndex418
				}
			l419:
				add(ruleAuthorsTeam, position415)
			}
			return true
		l414:
			position, tokenIndex = position414, tokenIndex414
			return false
		},
		/* 58 AuthorSep <- <(AuthorSep1 / AuthorSep2)> */
		func() bool {
			position426, tokenIndex426 := position, tokenIndex
			{
				position427 := position
				{
					position428, tokenIndex428 := position, tokenIndex
					if !_rules[ruleAuthorSep1]() {
						goto l429
					}
					goto l428
				l429:
					position, tokenIndex = position428, tokenIndex428
					if !_rules[ruleAuthorSep2]() {
						goto l426
					}
				}
			l428:
				add(ruleAuthorSep, position427)
			}
			return true
		l426:
			position, tokenIndex = position426, tokenIndex426
			return false
		},
		/* 59 AuthorSep1 <- <(_? (',' _)? ('&' / ('e' 't') / ('a' 'n' 'd') / ('a' 'p' 'u' 'd')) _?)> */
		func() bool {
			position430, tokenIndex430 := position, tokenIndex
			{
				position431 := position
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
				{
					position434, tokenIndex434 := position, tokenIndex
					if buffer[position] != rune(',') {
						goto l434
					}
					position++
					if !_rules[rule_]() {
						goto l434
					}
					goto l435
				l434:
					position, tokenIndex = position434, tokenIndex434
				}
			l435:
				{
					position436, tokenIndex436 := position, tokenIndex
					if buffer[position] != rune('&') {
						goto l437
					}
					position++
					goto l436
				l437:
					position, tokenIndex = position436, tokenIndex436
					if buffer[position] != rune('e') {
						goto l438
					}
					position++
					if buffer[position] != rune('t') {
						goto l438
					}
					position++
					goto l436
				l438:
					position, tokenIndex = position436, tokenIndex436
					if buffer[position] != rune('a') {
						goto l439
					}
					position++
					if buffer[position] != rune('n') {
						goto l439
					}
					position++
					if buffer[position] != rune('d') {
						goto l439
					}
					position++
					goto l436
				l439:
					position, tokenIndex = position436, tokenIndex436
					if buffer[position] != rune('a') {
						goto l430
					}
					position++
					if buffer[position] != rune('p') {
						goto l430
					}
					position++
					if buffer[position] != rune('u') {
						goto l430
					}
					position++
					if buffer[position] != rune('d') {
						goto l430
					}
					position++
				}
			l436:
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
				add(ruleAuthorSep1, position431)
			}
			return true
		l430:
			position, tokenIndex = position430, tokenIndex430
			return false
		},
		/* 60 AuthorSep2 <- <(_? ',' _?)> */
		func() bool {
			position442, tokenIndex442 := position, tokenIndex
			{
				position443 := position
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
				if buffer[position] != rune(',') {
					goto l442
				}
				position++
				{
					position446, tokenIndex446 := position, tokenIndex
					if !_rules[rule_]() {
						goto l446
					}
					goto l447
				l446:
					position, tokenIndex = position446, tokenIndex446
				}
			l447:
				add(ruleAuthorSep2, position443)
			}
			return true
		l442:
			position, tokenIndex = position442, tokenIndex442
			return false
		},
		/* 61 AuthorEx <- <((('e' 'x' '.'?) / ('i' 'n')) _)> */
		func() bool {
			position448, tokenIndex448 := position, tokenIndex
			{
				position449 := position
				{
					position450, tokenIndex450 := position, tokenIndex
					if buffer[position] != rune('e') {
						goto l451
					}
					position++
					if buffer[position] != rune('x') {
						goto l451
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
					goto l450
				l451:
					position, tokenIndex = position450, tokenIndex450
					if buffer[position] != rune('i') {
						goto l448
					}
					position++
					if buffer[position] != rune('n') {
						goto l448
					}
					position++
				}
			l450:
				if !_rules[rule_]() {
					goto l448
				}
				add(ruleAuthorEx, position449)
			}
			return true
		l448:
			position, tokenIndex = position448, tokenIndex448
			return false
		},
		/* 62 AuthorEmend <- <('e' 'm' 'e' 'n' 'd' '.'? _)> */
		func() bool {
			position454, tokenIndex454 := position, tokenIndex
			{
				position455 := position
				if buffer[position] != rune('e') {
					goto l454
				}
				position++
				if buffer[position] != rune('m') {
					goto l454
				}
				position++
				if buffer[position] != rune('e') {
					goto l454
				}
				position++
				if buffer[position] != rune('n') {
					goto l454
				}
				position++
				if buffer[position] != rune('d') {
					goto l454
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
				if !_rules[rule_]() {
					goto l454
				}
				add(ruleAuthorEmend, position455)
			}
			return true
		l454:
			position, tokenIndex = position454, tokenIndex454
			return false
		},
		/* 63 Author <- <(Author1 / Author2 / UnknownAuthor)> */
		func() bool {
			position458, tokenIndex458 := position, tokenIndex
			{
				position459 := position
				{
					position460, tokenIndex460 := position, tokenIndex
					if !_rules[ruleAuthor1]() {
						goto l461
					}
					goto l460
				l461:
					position, tokenIndex = position460, tokenIndex460
					if !_rules[ruleAuthor2]() {
						goto l462
					}
					goto l460
				l462:
					position, tokenIndex = position460, tokenIndex460
					if !_rules[ruleUnknownAuthor]() {
						goto l458
					}
				}
			l460:
				add(ruleAuthor, position459)
			}
			return true
		l458:
			position, tokenIndex = position458, tokenIndex458
			return false
		},
		/* 64 Author1 <- <(Author2 _? Filius)> */
		func() bool {
			position463, tokenIndex463 := position, tokenIndex
			{
				position464 := position
				if !_rules[ruleAuthor2]() {
					goto l463
				}
				{
					position465, tokenIndex465 := position, tokenIndex
					if !_rules[rule_]() {
						goto l465
					}
					goto l466
				l465:
					position, tokenIndex = position465, tokenIndex465
				}
			l466:
				if !_rules[ruleFilius]() {
					goto l463
				}
				add(ruleAuthor1, position464)
			}
			return true
		l463:
			position, tokenIndex = position463, tokenIndex463
			return false
		},
		/* 65 Author2 <- <(AuthorWord (_? AuthorWord)*)> */
		func() bool {
			position467, tokenIndex467 := position, tokenIndex
			{
				position468 := position
				if !_rules[ruleAuthorWord]() {
					goto l467
				}
			l469:
				{
					position470, tokenIndex470 := position, tokenIndex
					{
						position471, tokenIndex471 := position, tokenIndex
						if !_rules[rule_]() {
							goto l471
						}
						goto l472
					l471:
						position, tokenIndex = position471, tokenIndex471
					}
				l472:
					if !_rules[ruleAuthorWord]() {
						goto l470
					}
					goto l469
				l470:
					position, tokenIndex = position470, tokenIndex470
				}
				add(ruleAuthor2, position468)
			}
			return true
		l467:
			position, tokenIndex = position467, tokenIndex467
			return false
		},
		/* 66 UnknownAuthor <- <('?' / ((('a' 'u' 'c' 't') / ('a' 'n' 'o' 'n')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position473, tokenIndex473 := position, tokenIndex
			{
				position474 := position
				{
					position475, tokenIndex475 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l476
					}
					position++
					goto l475
				l476:
					position, tokenIndex = position475, tokenIndex475
					{
						position477, tokenIndex477 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l478
						}
						position++
						if buffer[position] != rune('u') {
							goto l478
						}
						position++
						if buffer[position] != rune('c') {
							goto l478
						}
						position++
						if buffer[position] != rune('t') {
							goto l478
						}
						position++
						goto l477
					l478:
						position, tokenIndex = position477, tokenIndex477
						if buffer[position] != rune('a') {
							goto l473
						}
						position++
						if buffer[position] != rune('n') {
							goto l473
						}
						position++
						if buffer[position] != rune('o') {
							goto l473
						}
						position++
						if buffer[position] != rune('n') {
							goto l473
						}
						position++
					}
				l477:
					{
						position479, tokenIndex479 := position, tokenIndex
						{
							position481, tokenIndex481 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l480
							}
							position, tokenIndex = position481, tokenIndex481
						}
						goto l479
					l480:
						position, tokenIndex = position479, tokenIndex479
						if buffer[position] != rune('.') {
							goto l473
						}
						position++
					}
				l479:
				}
			l475:
				add(ruleUnknownAuthor, position474)
			}
			return true
		l473:
			position, tokenIndex = position473, tokenIndex473
			return false
		},
		/* 67 AuthorWord <- <(!(('b' / 'B') ('o' / 'O') ('l' / 'L') ('d' / 'D') ':') (AuthorWord1 / AuthorWord2 / AuthorWord3 / AuthorPrefix))> */
		func() bool {
			position482, tokenIndex482 := position, tokenIndex
			{
				position483 := position
				{
					position484, tokenIndex484 := position, tokenIndex
					{
						position485, tokenIndex485 := position, tokenIndex
						if buffer[position] != rune('b') {
							goto l486
						}
						position++
						goto l485
					l486:
						position, tokenIndex = position485, tokenIndex485
						if buffer[position] != rune('B') {
							goto l484
						}
						position++
					}
				l485:
					{
						position487, tokenIndex487 := position, tokenIndex
						if buffer[position] != rune('o') {
							goto l488
						}
						position++
						goto l487
					l488:
						position, tokenIndex = position487, tokenIndex487
						if buffer[position] != rune('O') {
							goto l484
						}
						position++
					}
				l487:
					{
						position489, tokenIndex489 := position, tokenIndex
						if buffer[position] != rune('l') {
							goto l490
						}
						position++
						goto l489
					l490:
						position, tokenIndex = position489, tokenIndex489
						if buffer[position] != rune('L') {
							goto l484
						}
						position++
					}
				l489:
					{
						position491, tokenIndex491 := position, tokenIndex
						if buffer[position] != rune('d') {
							goto l492
						}
						position++
						goto l491
					l492:
						position, tokenIndex = position491, tokenIndex491
						if buffer[position] != rune('D') {
							goto l484
						}
						position++
					}
				l491:
					if buffer[position] != rune(':') {
						goto l484
					}
					position++
					goto l482
				l484:
					position, tokenIndex = position484, tokenIndex484
				}
				{
					position493, tokenIndex493 := position, tokenIndex
					if !_rules[ruleAuthorWord1]() {
						goto l494
					}
					goto l493
				l494:
					position, tokenIndex = position493, tokenIndex493
					if !_rules[ruleAuthorWord2]() {
						goto l495
					}
					goto l493
				l495:
					position, tokenIndex = position493, tokenIndex493
					if !_rules[ruleAuthorWord3]() {
						goto l496
					}
					goto l493
				l496:
					position, tokenIndex = position493, tokenIndex493
					if !_rules[ruleAuthorPrefix]() {
						goto l482
					}
				}
			l493:
				add(ruleAuthorWord, position483)
			}
			return true
		l482:
			position, tokenIndex = position482, tokenIndex482
			return false
		},
		/* 68 AuthorWord1 <- <(('a' 'r' 'g' '.') / ('e' 't' ' ' 'a' 'l' '.' '{' '?' '}') / ((('e' 't') / '&') (' ' 'a' 'l') '.'?))> */
		func() bool {
			position497, tokenIndex497 := position, tokenIndex
			{
				position498 := position
				{
					position499, tokenIndex499 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l500
					}
					position++
					if buffer[position] != rune('r') {
						goto l500
					}
					position++
					if buffer[position] != rune('g') {
						goto l500
					}
					position++
					if buffer[position] != rune('.') {
						goto l500
					}
					position++
					goto l499
				l500:
					position, tokenIndex = position499, tokenIndex499
					if buffer[position] != rune('e') {
						goto l501
					}
					position++
					if buffer[position] != rune('t') {
						goto l501
					}
					position++
					if buffer[position] != rune(' ') {
						goto l501
					}
					position++
					if buffer[position] != rune('a') {
						goto l501
					}
					position++
					if buffer[position] != rune('l') {
						goto l501
					}
					position++
					if buffer[position] != rune('.') {
						goto l501
					}
					position++
					if buffer[position] != rune('{') {
						goto l501
					}
					position++
					if buffer[position] != rune('?') {
						goto l501
					}
					position++
					if buffer[position] != rune('}') {
						goto l501
					}
					position++
					goto l499
				l501:
					position, tokenIndex = position499, tokenIndex499
					{
						position502, tokenIndex502 := position, tokenIndex
						if buffer[position] != rune('e') {
							goto l503
						}
						position++
						if buffer[position] != rune('t') {
							goto l503
						}
						position++
						goto l502
					l503:
						position, tokenIndex = position502, tokenIndex502
						if buffer[position] != rune('&') {
							goto l497
						}
						position++
					}
				l502:
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
					{
						position504, tokenIndex504 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l504
						}
						position++
						goto l505
					l504:
						position, tokenIndex = position504, tokenIndex504
					}
				l505:
				}
			l499:
				add(ruleAuthorWord1, position498)
			}
			return true
		l497:
			position, tokenIndex = position497, tokenIndex497
			return false
		},
		/* 69 AuthorWord2 <- <(AuthorWord3 dash AuthorWordSoft)> */
		func() bool {
			position506, tokenIndex506 := position, tokenIndex
			{
				position507 := position
				if !_rules[ruleAuthorWord3]() {
					goto l506
				}
				if !_rules[ruledash]() {
					goto l506
				}
				if !_rules[ruleAuthorWordSoft]() {
					goto l506
				}
				add(ruleAuthorWord2, position507)
			}
			return true
		l506:
			position, tokenIndex = position506, tokenIndex506
			return false
		},
		/* 70 AuthorWord3 <- <(AuthorPrefixGlued? (AllCapsAuthorWord / CapAuthorWord) '.'?)> */
		func() bool {
			position508, tokenIndex508 := position, tokenIndex
			{
				position509 := position
				{
					position510, tokenIndex510 := position, tokenIndex
					if !_rules[ruleAuthorPrefixGlued]() {
						goto l510
					}
					goto l511
				l510:
					position, tokenIndex = position510, tokenIndex510
				}
			l511:
				{
					position512, tokenIndex512 := position, tokenIndex
					if !_rules[ruleAllCapsAuthorWord]() {
						goto l513
					}
					goto l512
				l513:
					position, tokenIndex = position512, tokenIndex512
					if !_rules[ruleCapAuthorWord]() {
						goto l508
					}
				}
			l512:
				{
					position514, tokenIndex514 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l514
					}
					position++
					goto l515
				l514:
					position, tokenIndex = position514, tokenIndex514
				}
			l515:
				add(ruleAuthorWord3, position509)
			}
			return true
		l508:
			position, tokenIndex = position508, tokenIndex508
			return false
		},
		/* 71 AuthorWordSoft <- <(((AuthorUpperChar (AuthorUpperChar+ / AuthorLowerChar+)) / AuthorLowerChar+) '.'?)> */
		func() bool {
			position516, tokenIndex516 := position, tokenIndex
			{
				position517 := position
				{
					position518, tokenIndex518 := position, tokenIndex
					if !_rules[ruleAuthorUpperChar]() {
						goto l519
					}
					{
						position520, tokenIndex520 := position, tokenIndex
						if !_rules[ruleAuthorUpperChar]() {
							goto l521
						}
					l522:
						{
							position523, tokenIndex523 := position, tokenIndex
							if !_rules[ruleAuthorUpperChar]() {
								goto l523
							}
							goto l522
						l523:
							position, tokenIndex = position523, tokenIndex523
						}
						goto l520
					l521:
						position, tokenIndex = position520, tokenIndex520
						if !_rules[ruleAuthorLowerChar]() {
							goto l519
						}
					l524:
						{
							position525, tokenIndex525 := position, tokenIndex
							if !_rules[ruleAuthorLowerChar]() {
								goto l525
							}
							goto l524
						l525:
							position, tokenIndex = position525, tokenIndex525
						}
					}
				l520:
					goto l518
				l519:
					position, tokenIndex = position518, tokenIndex518
					if !_rules[ruleAuthorLowerChar]() {
						goto l516
					}
				l526:
					{
						position527, tokenIndex527 := position, tokenIndex
						if !_rules[ruleAuthorLowerChar]() {
							goto l527
						}
						goto l526
					l527:
						position, tokenIndex = position527, tokenIndex527
					}
				}
			l518:
				{
					position528, tokenIndex528 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l528
					}
					position++
					goto l529
				l528:
					position, tokenIndex = position528, tokenIndex528
				}
			l529:
				add(ruleAuthorWordSoft, position517)
			}
			return true
		l516:
			position, tokenIndex = position516, tokenIndex516
			return false
		},
		/* 72 CapAuthorWord <- <(AuthorUpperChar AuthorLowerChar*)> */
		func() bool {
			position530, tokenIndex530 := position, tokenIndex
			{
				position531 := position
				if !_rules[ruleAuthorUpperChar]() {
					goto l530
				}
			l532:
				{
					position533, tokenIndex533 := position, tokenIndex
					if !_rules[ruleAuthorLowerChar]() {
						goto l533
					}
					goto l532
				l533:
					position, tokenIndex = position533, tokenIndex533
				}
				add(ruleCapAuthorWord, position531)
			}
			return true
		l530:
			position, tokenIndex = position530, tokenIndex530
			return false
		},
		/* 73 AllCapsAuthorWord <- <(AuthorUpperChar AuthorUpperChar+)> */
		func() bool {
			position534, tokenIndex534 := position, tokenIndex
			{
				position535 := position
				if !_rules[ruleAuthorUpperChar]() {
					goto l534
				}
				if !_rules[ruleAuthorUpperChar]() {
					goto l534
				}
			l536:
				{
					position537, tokenIndex537 := position, tokenIndex
					if !_rules[ruleAuthorUpperChar]() {
						goto l537
					}
					goto l536
				l537:
					position, tokenIndex = position537, tokenIndex537
				}
				add(ruleAllCapsAuthorWord, position535)
			}
			return true
		l534:
			position, tokenIndex = position534, tokenIndex534
			return false
		},
		/* 74 Filius <- <(('f' '.') / ('f' 'i' 'l' '.') / ('f' 'i' 'l' 'i' 'u' 's'))> */
		func() bool {
			position538, tokenIndex538 := position, tokenIndex
			{
				position539 := position
				{
					position540, tokenIndex540 := position, tokenIndex
					if buffer[position] != rune('f') {
						goto l541
					}
					position++
					if buffer[position] != rune('.') {
						goto l541
					}
					position++
					goto l540
				l541:
					position, tokenIndex = position540, tokenIndex540
					if buffer[position] != rune('f') {
						goto l542
					}
					position++
					if buffer[position] != rune('i') {
						goto l542
					}
					position++
					if buffer[position] != rune('l') {
						goto l542
					}
					position++
					if buffer[position] != rune('.') {
						goto l542
					}
					position++
					goto l540
				l542:
					position, tokenIndex = position540, tokenIndex540
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
					if buffer[position] != rune('i') {
						goto l538
					}
					position++
					if buffer[position] != rune('u') {
						goto l538
					}
					position++
					if buffer[position] != rune('s') {
						goto l538
					}
					position++
				}
			l540:
				add(ruleFilius, position539)
			}
			return true
		l538:
			position, tokenIndex = position538, tokenIndex538
			return false
		},
		/* 75 AuthorPrefixGlued <- <(('d' '\'') / ('O' '\'') / ('L' '\''))> */
		func() bool {
			position543, tokenIndex543 := position, tokenIndex
			{
				position544 := position
				{
					position545, tokenIndex545 := position, tokenIndex
					if buffer[position] != rune('d') {
						goto l546
					}
					position++
					if buffer[position] != rune('\'') {
						goto l546
					}
					position++
					goto l545
				l546:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('O') {
						goto l547
					}
					position++
					if buffer[position] != rune('\'') {
						goto l547
					}
					position++
					goto l545
				l547:
					position, tokenIndex = position545, tokenIndex545
					if buffer[position] != rune('L') {
						goto l543
					}
					position++
					if buffer[position] != rune('\'') {
						goto l543
					}
					position++
				}
			l545:
				add(ruleAuthorPrefixGlued, position544)
			}
			return true
		l543:
			position, tokenIndex = position543, tokenIndex543
			return false
		},
		/* 76 AuthorPrefix <- <(AuthorPrefix1 / AuthorPrefix2)> */
		func() bool {
			position548, tokenIndex548 := position, tokenIndex
			{
				position549 := position
				{
					position550, tokenIndex550 := position, tokenIndex
					if !_rules[ruleAuthorPrefix1]() {
						goto l551
					}
					goto l550
				l551:
					position, tokenIndex = position550, tokenIndex550
					if !_rules[ruleAuthorPrefix2]() {
						goto l548
					}
				}
			l550:
				add(ruleAuthorPrefix, position549)
			}
			return true
		l548:
			position, tokenIndex = position548, tokenIndex548
			return false
		},
		/* 77 AuthorPrefix2 <- <(('v' '.' (_? ('d' '.'))?) / ('\'' 't'))> */
		func() bool {
			position552, tokenIndex552 := position, tokenIndex
			{
				position553 := position
				{
					position554, tokenIndex554 := position, tokenIndex
					if buffer[position] != rune('v') {
						goto l555
					}
					position++
					if buffer[position] != rune('.') {
						goto l555
					}
					position++
					{
						position556, tokenIndex556 := position, tokenIndex
						{
							position558, tokenIndex558 := position, tokenIndex
							if !_rules[rule_]() {
								goto l558
							}
							goto l559
						l558:
							position, tokenIndex = position558, tokenIndex558
						}
					l559:
						if buffer[position] != rune('d') {
							goto l556
						}
						position++
						if buffer[position] != rune('.') {
							goto l556
						}
						position++
						goto l557
					l556:
						position, tokenIndex = position556, tokenIndex556
					}
				l557:
					goto l554
				l555:
					position, tokenIndex = position554, tokenIndex554
					if buffer[position] != rune('\'') {
						goto l552
					}
					position++
					if buffer[position] != rune('t') {
						goto l552
					}
					position++
				}
			l554:
				add(ruleAuthorPrefix2, position553)
			}
			return true
		l552:
			position, tokenIndex = position552, tokenIndex552
			return false
		},
		/* 78 AuthorPrefix1 <- <((('a' 'b') / ('a' 'f') / ('b' 'i' 's') / ('d' 'a') / ('d' 'e' 'r') / ('d' 'e' 's') / ('d' 'e' 'n') / ('d' 'e' 'l') / ('d' 'e' 'l' 'l' 'a') / ('d' 'e' 'l' 'a') / ('d' 'e') / ('d' 'i') / ('d' 'u') / ('e' 'l') / ('l' 'a') / ('l' 'e') / ('t' 'e' 'r') / ('v' 'a' 'n') / ('d' '\'') / ('i' 'n' '\'' 't') / ('z' 'u' 'r') / ('v' 'o' 'n' (_ (('d' '.') / ('d' 'e' 'm')))?) / ('v' (_ 'd')?)) &_)> */
		func() bool {
			position560, tokenIndex560 := position, tokenIndex
			{
				position561 := position
				{
					position562, tokenIndex562 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l563
					}
					position++
					if buffer[position] != rune('b') {
						goto l563
					}
					position++
					goto l562
				l563:
					position, tokenIndex = position562, tokenIndex562
					if buffer[position] != rune('a') {
						goto l564
					}
					position++
					if buffer[position] != rune('f') {
						goto l564
					}
					position++
					goto l562
				l564:
					position, tokenIndex = position562, tokenIndex562
					if buffer[position] != rune('b') {
						goto l565
					}
					position++
					if buffer[position] != rune('i') {
						goto l565
					}
					position++
					if buffer[position] != rune('s') {
						goto l565
					}
					position++
					goto l562
				l565:
					position, tokenIndex = position562, tokenIndex562
					if buffer[position] != rune('d') {
						goto l566
					}
					position++
					if buffer[position] != rune('a') {
						goto l566
					}
					position++
					goto l562
				l566:
					position, tokenIndex = position562, tokenIndex562
					if buffer[position] != rune('d') {
						goto l567
					}
					position++
					if buffer[position] != rune('e') {
						goto l567
					}
					position++
					if buffer[position] != rune('r') {
						goto l567
					}
					position++
					goto l562
				l567:
					position, tokenIndex = position562, tokenIndex562
					if buffer[position] != rune('d') {
						goto l568
					}
					position++
					if buffer[position] != rune('e') {
						goto l568
					}
					position++
					if buffer[position] != rune('s') {
						goto l568
					}
					position++
					goto l562
				l568:
					position, tokenIndex = position562, tokenIndex562
					if buffer[position] != rune('d') {
						goto l569
					}
					position++
					if buffer[position] != rune('e') {
						goto l569
					}
					position++
					if buffer[position] != rune('n') {
						goto l569
					}
					position++
					goto l562
				l569:
					position, tokenIndex = position562, tokenIndex562
					if buffer[position] != rune('d') {
						goto l570
					}
					position++
					if buffer[position] != rune('e') {
						goto l570
					}
					position++
					if buffer[position] != rune('l') {
						goto l570
					}
					position++
					goto l562
				l570:
					position, tokenIndex = position562, tokenIndex562
					if buffer[position] != rune('d') {
						goto l571
					}
					position++
					if buffer[position] != rune('e') {
						goto l571
					}
					position++
					if buffer[position] != rune('l') {
						goto l571
					}
					position++
					if buffer[position] != rune('l') {
						goto l571
					}
					position++
					if buffer[position] != rune('a') {
						goto l571
					}
					position++
					goto l562
				l571:
					position, tokenIndex = position562, tokenIndex562
					if buffer[position] != rune('d') {
						goto l572
					}
					position++
					if buffer[position] != rune('e') {
						goto l572
					}
					position++
					if buffer[position] != rune('l') {
						goto l572
					}
					position++
					if buffer[position] != rune('a') {
						goto l572
					}
					position++
					goto l562
				l572:
					position, tokenIndex = position562, tokenIndex562
					if buffer[position] != rune('d') {
						goto l573
					}
					position++
					if buffer[position] != rune('e') {
						goto l573
					}
					position++
					goto l562
				l573:
					position, tokenIndex = position562, tokenIndex562
					if buffer[position] != rune('d') {
						goto l574
					}
					position++
					if buffer[position] != rune('i') {
						goto l574
					}
					position++
					goto l562
				l574:
					position, tokenIndex = position562, tokenIndex562
					if buffer[position] != rune('d') {
						goto l575
					}
					position++
					if buffer[position] != rune('u') {
						goto l575
					}
					position++
					goto l562
				l575:
					position, tokenIndex = position562, tokenIndex562
					if buffer[position] != rune('e') {
						goto l576
					}
					position++
					if buffer[position] != rune('l') {
						goto l576
					}
					position++
					goto l562
				l576:
					position, tokenIndex = position562, tokenIndex562
					if buffer[position] != rune('l') {
						goto l577
					}
					position++
					if buffer[position] != rune('a') {
						goto l577
					}
					position++
					goto l562
				l577:
					position, tokenIndex = position562, tokenIndex562
					if buffer[position] != rune('l') {
						goto l578
					}
					position++
					if buffer[position] != rune('e') {
						goto l578
					}
					position++
					goto l562
				l578:
					position, tokenIndex = position562, tokenIndex562
					if buffer[position] != rune('t') {
						goto l579
					}
					position++
					if buffer[position] != rune('e') {
						goto l579
					}
					position++
					if buffer[position] != rune('r') {
						goto l579
					}
					position++
					goto l562
				l579:
					position, tokenIndex = position562, tokenIndex562
					if buffer[position] != rune('v') {
						goto l580
					}
					position++
					if buffer[position] != rune('a') {
						goto l580
					}
					position++
					if buffer[position] != rune('n') {
						goto l580
					}
					position++
					goto l562
				l580:
					position, tokenIndex = position562, tokenIndex562
					if buffer[position] != rune('d') {
						goto l581
					}
					position++
					if buffer[position] != rune('\'') {
						goto l581
					}
					position++
					goto l562
				l581:
					position, tokenIndex = position562, tokenIndex562
					if buffer[position] != rune('i') {
						goto l582
					}
					position++
					if buffer[position] != rune('n') {
						goto l582
					}
					position++
					if buffer[position] != rune('\'') {
						goto l582
					}
					position++
					if buffer[position] != rune('t') {
						goto l582
					}
					position++
					goto l562
				l582:
					position, tokenIndex = position562, tokenIndex562
					if buffer[position] != rune('z') {
						goto l583
					}
					position++
					if buffer[position] != rune('u') {
						goto l583
					}
					position++
					if buffer[position] != rune('r') {
						goto l583
					}
					position++
					goto l562
				l583:
					position, tokenIndex = position562, tokenIndex562
					if buffer[position] != rune('v') {
						goto l584
					}
					position++
					if buffer[position] != rune('o') {
						goto l584
					}
					position++
					if buffer[position] != rune('n') {
						goto l584
					}
					position++
					{
						position585, tokenIndex585 := position, tokenIndex
						if !_rules[rule_]() {
							goto l585
						}
						{
							position587, tokenIndex587 := position, tokenIndex
							if buffer[position] != rune('d') {
								goto l588
							}
							position++
							if buffer[position] != rune('.') {
								goto l588
							}
							position++
							goto l587
						l588:
							position, tokenIndex = position587, tokenIndex587
							if buffer[position] != rune('d') {
								goto l585
							}
							position++
							if buffer[position] != rune('e') {
								goto l585
							}
							position++
							if buffer[position] != rune('m') {
								goto l585
							}
							position++
						}
					l587:
						goto l586
					l585:
						position, tokenIndex = position585, tokenIndex585
					}
				l586:
					goto l562
				l584:
					position, tokenIndex = position562, tokenIndex562
					if buffer[position] != rune('v') {
						goto l560
					}
					position++
					{
						position589, tokenIndex589 := position, tokenIndex
						if !_rules[rule_]() {
							goto l589
						}
						if buffer[position] != rune('d') {
							goto l589
						}
						position++
						goto l590
					l589:
						position, tokenIndex = position589, tokenIndex589
					}
				l590:
				}
			l562:
				{
					position591, tokenIndex591 := position, tokenIndex
					if !_rules[rule_]() {
						goto l560
					}
					position, tokenIndex = position591, tokenIndex591
				}
				add(ruleAuthorPrefix1, position561)
			}
			return true
		l560:
			position, tokenIndex = position560, tokenIndex560
			return false
		},
		/* 79 AuthorUpperChar <- <(hASCII / MiscodedChar / ('À' / 'Á' / 'Â' / 'Ã' / 'Ä' / 'Å' / 'Æ' / 'Ç' / 'È' / 'É' / 'Ê' / 'Ë' / 'Ì' / 'Í' / 'Î' / 'Ï' / 'Ð' / 'Ñ' / 'Ò' / 'Ó' / 'Ô' / 'Õ' / 'Ö' / 'Ø' / 'Ù' / 'Ú' / 'Û' / 'Ü' / 'Ý' / 'Ć' / 'Č' / 'Ď' / 'İ' / 'Ķ' / 'Ĺ' / 'ĺ' / 'Ľ' / 'ľ' / 'Ł' / 'ł' / 'Ņ' / 'Ō' / 'Ő' / 'Œ' / 'Ř' / 'Ś' / 'Ŝ' / 'Ş' / 'Š' / 'Ÿ' / 'Ź' / 'Ż' / 'Ž' / 'ƒ' / 'Ǿ' / 'Ș' / 'Ț'))> */
		func() bool {
			position592, tokenIndex592 := position, tokenIndex
			{
				position593 := position
				{
					position594, tokenIndex594 := position, tokenIndex
					if !_rules[rulehASCII]() {
						goto l595
					}
					goto l594
				l595:
					position, tokenIndex = position594, tokenIndex594
					if !_rules[ruleMiscodedChar]() {
						goto l596
					}
					goto l594
				l596:
					position, tokenIndex = position594, tokenIndex594
					{
						position597, tokenIndex597 := position, tokenIndex
						if buffer[position] != rune('À') {
							goto l598
						}
						position++
						goto l597
					l598:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Á') {
							goto l599
						}
						position++
						goto l597
					l599:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Â') {
							goto l600
						}
						position++
						goto l597
					l600:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ã') {
							goto l601
						}
						position++
						goto l597
					l601:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ä') {
							goto l602
						}
						position++
						goto l597
					l602:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Å') {
							goto l603
						}
						position++
						goto l597
					l603:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Æ') {
							goto l604
						}
						position++
						goto l597
					l604:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ç') {
							goto l605
						}
						position++
						goto l597
					l605:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('È') {
							goto l606
						}
						position++
						goto l597
					l606:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('É') {
							goto l607
						}
						position++
						goto l597
					l607:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ê') {
							goto l608
						}
						position++
						goto l597
					l608:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ë') {
							goto l609
						}
						position++
						goto l597
					l609:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ì') {
							goto l610
						}
						position++
						goto l597
					l610:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Í') {
							goto l611
						}
						position++
						goto l597
					l611:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Î') {
							goto l612
						}
						position++
						goto l597
					l612:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ï') {
							goto l613
						}
						position++
						goto l597
					l613:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ð') {
							goto l614
						}
						position++
						goto l597
					l614:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ñ') {
							goto l615
						}
						position++
						goto l597
					l615:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ò') {
							goto l616
						}
						position++
						goto l597
					l616:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ó') {
							goto l617
						}
						position++
						goto l597
					l617:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ô') {
							goto l618
						}
						position++
						goto l597
					l618:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Õ') {
							goto l619
						}
						position++
						goto l597
					l619:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ö') {
							goto l620
						}
						position++
						goto l597
					l620:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ø') {
							goto l621
						}
						position++
						goto l597
					l621:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ù') {
							goto l622
						}
						position++
						goto l597
					l622:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ú') {
							goto l623
						}
						position++
						goto l597
					l623:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Û') {
							goto l624
						}
						position++
						goto l597
					l624:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ü') {
							goto l625
						}
						position++
						goto l597
					l625:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ý') {
							goto l626
						}
						position++
						goto l597
					l626:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ć') {
							goto l627
						}
						position++
						goto l597
					l627:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Č') {
							goto l628
						}
						position++
						goto l597
					l628:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ď') {
							goto l629
						}
						position++
						goto l597
					l629:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('İ') {
							goto l630
						}
						position++
						goto l597
					l630:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ķ') {
							goto l631
						}
						position++
						goto l597
					l631:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ĺ') {
							goto l632
						}
						position++
						goto l597
					l632:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('ĺ') {
							goto l633
						}
						position++
						goto l597
					l633:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ľ') {
							goto l634
						}
						position++
						goto l597
					l634:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('ľ') {
							goto l635
						}
						position++
						goto l597
					l635:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ł') {
							goto l636
						}
						position++
						goto l597
					l636:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('ł') {
							goto l637
						}
						position++
						goto l597
					l637:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ņ') {
							goto l638
						}
						position++
						goto l597
					l638:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ō') {
							goto l639
						}
						position++
						goto l597
					l639:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ő') {
							goto l640
						}
						position++
						goto l597
					l640:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Œ') {
							goto l641
						}
						position++
						goto l597
					l641:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ř') {
							goto l642
						}
						position++
						goto l597
					l642:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ś') {
							goto l643
						}
						position++
						goto l597
					l643:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ŝ') {
							goto l644
						}
						position++
						goto l597
					l644:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ş') {
							goto l645
						}
						position++
						goto l597
					l645:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Š') {
							goto l646
						}
						position++
						goto l597
					l646:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ÿ') {
							goto l647
						}
						position++
						goto l597
					l647:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ź') {
							goto l648
						}
						position++
						goto l597
					l648:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ż') {
							goto l649
						}
						position++
						goto l597
					l649:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ž') {
							goto l650
						}
						position++
						goto l597
					l650:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('ƒ') {
							goto l651
						}
						position++
						goto l597
					l651:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ǿ') {
							goto l652
						}
						position++
						goto l597
					l652:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ș') {
							goto l653
						}
						position++
						goto l597
					l653:
						position, tokenIndex = position597, tokenIndex597
						if buffer[position] != rune('Ț') {
							goto l592
						}
						position++
					}
				l597:
				}
			l594:
				add(ruleAuthorUpperChar, position593)
			}
			return true
		l592:
			position, tokenIndex = position592, tokenIndex592
			return false
		},
		/* 80 AuthorLowerChar <- <(lASCII / MiscodedChar / ('à' / 'á' / 'â' / 'ã' / 'ä' / 'å' / 'æ' / 'ç' / 'è' / 'é' / 'ê' / 'ë' / 'ì' / 'í' / 'î' / 'ï' / 'ð' / 'ñ' / 'ò' / 'ó' / 'ó' / 'ô' / 'õ' / 'ö' / 'ø' / 'ù' / 'ú' / 'û' / 'ü' / 'ý' / 'ÿ' / 'ā' / 'ă' / 'ą' / 'ć' / 'ĉ' / 'č' / 'ď' / 'đ' / '\'' / 'ē' / 'ĕ' / 'ė' / 'ę' / 'ě' / 'ğ' / 'ī' / 'ĭ' / 'İ' / 'ı' / 'ĺ' / 'ľ' / 'ł' / 'ń' / 'ņ' / 'ň' / 'ŏ' / 'ő' / 'œ' / 'ŕ' / 'ř' / 'ś' / 'ş' / 'š' / 'ţ' / 'ť' / 'ũ' / 'ū' / 'ŭ' / 'ů' / 'ű' / 'ź' / 'ż' / 'ž' / 'ſ' / 'ǎ' / 'ǔ' / 'ǧ' / 'ș' / 'ț' / 'ȳ' / 'ß'))> */
		func() bool {
			position654, tokenIndex654 := position, tokenIndex
			{
				position655 := position
				{
					position656, tokenIndex656 := position, tokenIndex
					if !_rules[rulelASCII]() {
						goto l657
					}
					goto l656
				l657:
					position, tokenIndex = position656, tokenIndex656
					if !_rules[ruleMiscodedChar]() {
						goto l658
					}
					goto l656
				l658:
					position, tokenIndex = position656, tokenIndex656
					{
						position659, tokenIndex659 := position, tokenIndex
						if buffer[position] != rune('à') {
							goto l660
						}
						position++
						goto l659
					l660:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('á') {
							goto l661
						}
						position++
						goto l659
					l661:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('â') {
							goto l662
						}
						position++
						goto l659
					l662:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ã') {
							goto l663
						}
						position++
						goto l659
					l663:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ä') {
							goto l664
						}
						position++
						goto l659
					l664:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('å') {
							goto l665
						}
						position++
						goto l659
					l665:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('æ') {
							goto l666
						}
						position++
						goto l659
					l666:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ç') {
							goto l667
						}
						position++
						goto l659
					l667:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('è') {
							goto l668
						}
						position++
						goto l659
					l668:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('é') {
							goto l669
						}
						position++
						goto l659
					l669:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ê') {
							goto l670
						}
						position++
						goto l659
					l670:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ë') {
							goto l671
						}
						position++
						goto l659
					l671:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ì') {
							goto l672
						}
						position++
						goto l659
					l672:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('í') {
							goto l673
						}
						position++
						goto l659
					l673:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('î') {
							goto l674
						}
						position++
						goto l659
					l674:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ï') {
							goto l675
						}
						position++
						goto l659
					l675:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ð') {
							goto l676
						}
						position++
						goto l659
					l676:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ñ') {
							goto l677
						}
						position++
						goto l659
					l677:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ò') {
							goto l678
						}
						position++
						goto l659
					l678:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ó') {
							goto l679
						}
						position++
						goto l659
					l679:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ó') {
							goto l680
						}
						position++
						goto l659
					l680:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ô') {
							goto l681
						}
						position++
						goto l659
					l681:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('õ') {
							goto l682
						}
						position++
						goto l659
					l682:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ö') {
							goto l683
						}
						position++
						goto l659
					l683:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ø') {
							goto l684
						}
						position++
						goto l659
					l684:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ù') {
							goto l685
						}
						position++
						goto l659
					l685:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ú') {
							goto l686
						}
						position++
						goto l659
					l686:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('û') {
							goto l687
						}
						position++
						goto l659
					l687:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ü') {
							goto l688
						}
						position++
						goto l659
					l688:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ý') {
							goto l689
						}
						position++
						goto l659
					l689:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ÿ') {
							goto l690
						}
						position++
						goto l659
					l690:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ā') {
							goto l691
						}
						position++
						goto l659
					l691:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ă') {
							goto l692
						}
						position++
						goto l659
					l692:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ą') {
							goto l693
						}
						position++
						goto l659
					l693:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ć') {
							goto l694
						}
						position++
						goto l659
					l694:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ĉ') {
							goto l695
						}
						position++
						goto l659
					l695:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('č') {
							goto l696
						}
						position++
						goto l659
					l696:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ď') {
							goto l697
						}
						position++
						goto l659
					l697:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('đ') {
							goto l698
						}
						position++
						goto l659
					l698:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('\'') {
							goto l699
						}
						position++
						goto l659
					l699:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ē') {
							goto l700
						}
						position++
						goto l659
					l700:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ĕ') {
							goto l701
						}
						position++
						goto l659
					l701:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ė') {
							goto l702
						}
						position++
						goto l659
					l702:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ę') {
							goto l703
						}
						position++
						goto l659
					l703:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ě') {
							goto l704
						}
						position++
						goto l659
					l704:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ğ') {
							goto l705
						}
						position++
						goto l659
					l705:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ī') {
							goto l706
						}
						position++
						goto l659
					l706:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ĭ') {
							goto l707
						}
						position++
						goto l659
					l707:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('İ') {
							goto l708
						}
						position++
						goto l659
					l708:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ı') {
							goto l709
						}
						position++
						goto l659
					l709:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ĺ') {
							goto l710
						}
						position++
						goto l659
					l710:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ľ') {
							goto l711
						}
						position++
						goto l659
					l711:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ł') {
							goto l712
						}
						position++
						goto l659
					l712:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ń') {
							goto l713
						}
						position++
						goto l659
					l713:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ņ') {
							goto l714
						}
						position++
						goto l659
					l714:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ň') {
							goto l715
						}
						position++
						goto l659
					l715:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ŏ') {
							goto l716
						}
						position++
						goto l659
					l716:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ő') {
							goto l717
						}
						position++
						goto l659
					l717:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('œ') {
							goto l718
						}
						position++
						goto l659
					l718:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ŕ') {
							goto l719
						}
						position++
						goto l659
					l719:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ř') {
							goto l720
						}
						position++
						goto l659
					l720:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ś') {
							goto l721
						}
						position++
						goto l659
					l721:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ş') {
							goto l722
						}
						position++
						goto l659
					l722:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('š') {
							goto l723
						}
						position++
						goto l659
					l723:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ţ') {
							goto l724
						}
						position++
						goto l659
					l724:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ť') {
							goto l725
						}
						position++
						goto l659
					l725:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ũ') {
							goto l726
						}
						position++
						goto l659
					l726:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ū') {
							goto l727
						}
						position++
						goto l659
					l727:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ŭ') {
							goto l728
						}
						position++
						goto l659
					l728:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ů') {
							goto l729
						}
						position++
						goto l659
					l729:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ű') {
							goto l730
						}
						position++
						goto l659
					l730:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ź') {
							goto l731
						}
						position++
						goto l659
					l731:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ż') {
							goto l732
						}
						position++
						goto l659
					l732:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ž') {
							goto l733
						}
						position++
						goto l659
					l733:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ſ') {
							goto l734
						}
						position++
						goto l659
					l734:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ǎ') {
							goto l735
						}
						position++
						goto l659
					l735:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ǔ') {
							goto l736
						}
						position++
						goto l659
					l736:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ǧ') {
							goto l737
						}
						position++
						goto l659
					l737:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ș') {
							goto l738
						}
						position++
						goto l659
					l738:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ț') {
							goto l739
						}
						position++
						goto l659
					l739:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ȳ') {
							goto l740
						}
						position++
						goto l659
					l740:
						position, tokenIndex = position659, tokenIndex659
						if buffer[position] != rune('ß') {
							goto l654
						}
						position++
					}
				l659:
				}
			l656:
				add(ruleAuthorLowerChar, position655)
			}
			return true
		l654:
			position, tokenIndex = position654, tokenIndex654
			return false
		},
		/* 81 Year <- <(YearRange / YearApprox / YearWithParens / YearWithPage / YearWithDot / YearWithChar / YearNum)> */
		func() bool {
			position741, tokenIndex741 := position, tokenIndex
			{
				position742 := position
				{
					position743, tokenIndex743 := position, tokenIndex
					if !_rules[ruleYearRange]() {
						goto l744
					}
					goto l743
				l744:
					position, tokenIndex = position743, tokenIndex743
					if !_rules[ruleYearApprox]() {
						goto l745
					}
					goto l743
				l745:
					position, tokenIndex = position743, tokenIndex743
					if !_rules[ruleYearWithParens]() {
						goto l746
					}
					goto l743
				l746:
					position, tokenIndex = position743, tokenIndex743
					if !_rules[ruleYearWithPage]() {
						goto l747
					}
					goto l743
				l747:
					position, tokenIndex = position743, tokenIndex743
					if !_rules[ruleYearWithDot]() {
						goto l748
					}
					goto l743
				l748:
					position, tokenIndex = position743, tokenIndex743
					if !_rules[ruleYearWithChar]() {
						goto l749
					}
					goto l743
				l749:
					position, tokenIndex = position743, tokenIndex743
					if !_rules[ruleYearNum]() {
						goto l741
					}
				}
			l743:
				add(ruleYear, position742)
			}
			return true
		l741:
			position, tokenIndex = position741, tokenIndex741
			return false
		},
		/* 82 YearRange <- <(YearNum dash (nums+ ('a' / 'b' / 'c' / 'd' / 'e' / 'f' / 'g' / 'h' / 'i' / 'j' / 'k' / 'l' / 'm' / 'n' / 'o' / 'p' / 'q' / 'r' / 's' / 't' / 'u' / 'v' / 'w' / 'x' / 'y' / 'z' / '?')*))> */
		func() bool {
			position750, tokenIndex750 := position, tokenIndex
			{
				position751 := position
				if !_rules[ruleYearNum]() {
					goto l750
				}
				if !_rules[ruledash]() {
					goto l750
				}
				if !_rules[rulenums]() {
					goto l750
				}
			l752:
				{
					position753, tokenIndex753 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l753
					}
					goto l752
				l753:
					position, tokenIndex = position753, tokenIndex753
				}
			l754:
				{
					position755, tokenIndex755 := position, tokenIndex
					{
						position756, tokenIndex756 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l757
						}
						position++
						goto l756
					l757:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('b') {
							goto l758
						}
						position++
						goto l756
					l758:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('c') {
							goto l759
						}
						position++
						goto l756
					l759:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('d') {
							goto l760
						}
						position++
						goto l756
					l760:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('e') {
							goto l761
						}
						position++
						goto l756
					l761:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('f') {
							goto l762
						}
						position++
						goto l756
					l762:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('g') {
							goto l763
						}
						position++
						goto l756
					l763:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('h') {
							goto l764
						}
						position++
						goto l756
					l764:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('i') {
							goto l765
						}
						position++
						goto l756
					l765:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('j') {
							goto l766
						}
						position++
						goto l756
					l766:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('k') {
							goto l767
						}
						position++
						goto l756
					l767:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('l') {
							goto l768
						}
						position++
						goto l756
					l768:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('m') {
							goto l769
						}
						position++
						goto l756
					l769:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('n') {
							goto l770
						}
						position++
						goto l756
					l770:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('o') {
							goto l771
						}
						position++
						goto l756
					l771:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('p') {
							goto l772
						}
						position++
						goto l756
					l772:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('q') {
							goto l773
						}
						position++
						goto l756
					l773:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('r') {
							goto l774
						}
						position++
						goto l756
					l774:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('s') {
							goto l775
						}
						position++
						goto l756
					l775:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('t') {
							goto l776
						}
						position++
						goto l756
					l776:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('u') {
							goto l777
						}
						position++
						goto l756
					l777:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('v') {
							goto l778
						}
						position++
						goto l756
					l778:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('w') {
							goto l779
						}
						position++
						goto l756
					l779:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('x') {
							goto l780
						}
						position++
						goto l756
					l780:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('y') {
							goto l781
						}
						position++
						goto l756
					l781:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('z') {
							goto l782
						}
						position++
						goto l756
					l782:
						position, tokenIndex = position756, tokenIndex756
						if buffer[position] != rune('?') {
							goto l755
						}
						position++
					}
				l756:
					goto l754
				l755:
					position, tokenIndex = position755, tokenIndex755
				}
				add(ruleYearRange, position751)
			}
			return true
		l750:
			position, tokenIndex = position750, tokenIndex750
			return false
		},
		/* 83 YearWithDot <- <(YearNum '.')> */
		func() bool {
			position783, tokenIndex783 := position, tokenIndex
			{
				position784 := position
				if !_rules[ruleYearNum]() {
					goto l783
				}
				if buffer[position] != rune('.') {
					goto l783
				}
				position++
				add(ruleYearWithDot, position784)
			}
			return true
		l783:
			position, tokenIndex = position783, tokenIndex783
			return false
		},
		/* 84 YearApprox <- <('[' _? YearNum _? ']')> */
		func() bool {
			position785, tokenIndex785 := position, tokenIndex
			{
				position786 := position
				if buffer[position] != rune('[') {
					goto l785
				}
				position++
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
				if !_rules[ruleYearNum]() {
					goto l785
				}
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
				if buffer[position] != rune(']') {
					goto l785
				}
				position++
				add(ruleYearApprox, position786)
			}
			return true
		l785:
			position, tokenIndex = position785, tokenIndex785
			return false
		},
		/* 85 YearWithPage <- <((YearWithChar / YearNum) _? ':' _? nums+)> */
		func() bool {
			position791, tokenIndex791 := position, tokenIndex
			{
				position792 := position
				{
					position793, tokenIndex793 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l794
					}
					goto l793
				l794:
					position, tokenIndex = position793, tokenIndex793
					if !_rules[ruleYearNum]() {
						goto l791
					}
				}
			l793:
				{
					position795, tokenIndex795 := position, tokenIndex
					if !_rules[rule_]() {
						goto l795
					}
					goto l796
				l795:
					position, tokenIndex = position795, tokenIndex795
				}
			l796:
				if buffer[position] != rune(':') {
					goto l791
				}
				position++
				{
					position797, tokenIndex797 := position, tokenIndex
					if !_rules[rule_]() {
						goto l797
					}
					goto l798
				l797:
					position, tokenIndex = position797, tokenIndex797
				}
			l798:
				if !_rules[rulenums]() {
					goto l791
				}
			l799:
				{
					position800, tokenIndex800 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l800
					}
					goto l799
				l800:
					position, tokenIndex = position800, tokenIndex800
				}
				add(ruleYearWithPage, position792)
			}
			return true
		l791:
			position, tokenIndex = position791, tokenIndex791
			return false
		},
		/* 86 YearWithParens <- <('(' (YearWithChar / YearNum) ')')> */
		func() bool {
			position801, tokenIndex801 := position, tokenIndex
			{
				position802 := position
				if buffer[position] != rune('(') {
					goto l801
				}
				position++
				{
					position803, tokenIndex803 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l804
					}
					goto l803
				l804:
					position, tokenIndex = position803, tokenIndex803
					if !_rules[ruleYearNum]() {
						goto l801
					}
				}
			l803:
				if buffer[position] != rune(')') {
					goto l801
				}
				position++
				add(ruleYearWithParens, position802)
			}
			return true
		l801:
			position, tokenIndex = position801, tokenIndex801
			return false
		},
		/* 87 YearWithChar <- <(YearNum lASCII Action0)> */
		func() bool {
			position805, tokenIndex805 := position, tokenIndex
			{
				position806 := position
				if !_rules[ruleYearNum]() {
					goto l805
				}
				if !_rules[rulelASCII]() {
					goto l805
				}
				if !_rules[ruleAction0]() {
					goto l805
				}
				add(ruleYearWithChar, position806)
			}
			return true
		l805:
			position, tokenIndex = position805, tokenIndex805
			return false
		},
		/* 88 YearNum <- <(('1' / '2') ('0' / '7' / '8' / '9') nums (nums / '?') '?'*)> */
		func() bool {
			position807, tokenIndex807 := position, tokenIndex
			{
				position808 := position
				{
					position809, tokenIndex809 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l810
					}
					position++
					goto l809
				l810:
					position, tokenIndex = position809, tokenIndex809
					if buffer[position] != rune('2') {
						goto l807
					}
					position++
				}
			l809:
				{
					position811, tokenIndex811 := position, tokenIndex
					if buffer[position] != rune('0') {
						goto l812
					}
					position++
					goto l811
				l812:
					position, tokenIndex = position811, tokenIndex811
					if buffer[position] != rune('7') {
						goto l813
					}
					position++
					goto l811
				l813:
					position, tokenIndex = position811, tokenIndex811
					if buffer[position] != rune('8') {
						goto l814
					}
					position++
					goto l811
				l814:
					position, tokenIndex = position811, tokenIndex811
					if buffer[position] != rune('9') {
						goto l807
					}
					position++
				}
			l811:
				if !_rules[rulenums]() {
					goto l807
				}
				{
					position815, tokenIndex815 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l816
					}
					goto l815
				l816:
					position, tokenIndex = position815, tokenIndex815
					if buffer[position] != rune('?') {
						goto l807
					}
					position++
				}
			l815:
			l817:
				{
					position818, tokenIndex818 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l818
					}
					position++
					goto l817
				l818:
					position, tokenIndex = position818, tokenIndex818
				}
				add(ruleYearNum, position808)
			}
			return true
		l807:
			position, tokenIndex = position807, tokenIndex807
			return false
		},
		/* 89 NameUpperChar <- <(UpperChar / UpperCharExtended)> */
		func() bool {
			position819, tokenIndex819 := position, tokenIndex
			{
				position820 := position
				{
					position821, tokenIndex821 := position, tokenIndex
					if !_rules[ruleUpperChar]() {
						goto l822
					}
					goto l821
				l822:
					position, tokenIndex = position821, tokenIndex821
					if !_rules[ruleUpperCharExtended]() {
						goto l819
					}
				}
			l821:
				add(ruleNameUpperChar, position820)
			}
			return true
		l819:
			position, tokenIndex = position819, tokenIndex819
			return false
		},
		/* 90 UpperCharExtended <- <('Æ' / 'Œ' / 'Ö')> */
		func() bool {
			position823, tokenIndex823 := position, tokenIndex
			{
				position824 := position
				{
					position825, tokenIndex825 := position, tokenIndex
					if buffer[position] != rune('Æ') {
						goto l826
					}
					position++
					goto l825
				l826:
					position, tokenIndex = position825, tokenIndex825
					if buffer[position] != rune('Œ') {
						goto l827
					}
					position++
					goto l825
				l827:
					position, tokenIndex = position825, tokenIndex825
					if buffer[position] != rune('Ö') {
						goto l823
					}
					position++
				}
			l825:
				add(ruleUpperCharExtended, position824)
			}
			return true
		l823:
			position, tokenIndex = position823, tokenIndex823
			return false
		},
		/* 91 UpperChar <- <hASCII> */
		func() bool {
			position828, tokenIndex828 := position, tokenIndex
			{
				position829 := position
				if !_rules[rulehASCII]() {
					goto l828
				}
				add(ruleUpperChar, position829)
			}
			return true
		l828:
			position, tokenIndex = position828, tokenIndex828
			return false
		},
		/* 92 NameLowerChar <- <(LowerChar / LowerCharExtended / MiscodedChar)> */
		func() bool {
			position830, tokenIndex830 := position, tokenIndex
			{
				position831 := position
				{
					position832, tokenIndex832 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l833
					}
					goto l832
				l833:
					position, tokenIndex = position832, tokenIndex832
					if !_rules[ruleLowerCharExtended]() {
						goto l834
					}
					goto l832
				l834:
					position, tokenIndex = position832, tokenIndex832
					if !_rules[ruleMiscodedChar]() {
						goto l830
					}
				}
			l832:
				add(ruleNameLowerChar, position831)
			}
			return true
		l830:
			position, tokenIndex = position830, tokenIndex830
			return false
		},
		/* 93 MiscodedChar <- <'�'> */
		func() bool {
			position835, tokenIndex835 := position, tokenIndex
			{
				position836 := position
				if buffer[position] != rune('�') {
					goto l835
				}
				position++
				add(ruleMiscodedChar, position836)
			}
			return true
		l835:
			position, tokenIndex = position835, tokenIndex835
			return false
		},
		/* 94 LowerCharExtended <- <('æ' / 'œ' / 'à' / 'â' / 'å' / 'ã' / 'ä' / 'á' / 'ç' / 'č' / 'é' / 'è' / 'ë' / 'í' / 'ì' / 'ï' / 'ň' / 'ñ' / 'ñ' / 'ó' / 'ò' / 'ô' / 'ø' / 'õ' / 'ö' / 'ú' / 'ù' / 'ü' / 'ŕ' / 'ř' / 'ŗ' / 'ſ' / 'š' / 'š' / 'ş' / 'ž')> */
		func() bool {
			position837, tokenIndex837 := position, tokenIndex
			{
				position838 := position
				{
					position839, tokenIndex839 := position, tokenIndex
					if buffer[position] != rune('æ') {
						goto l840
					}
					position++
					goto l839
				l840:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('œ') {
						goto l841
					}
					position++
					goto l839
				l841:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('à') {
						goto l842
					}
					position++
					goto l839
				l842:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('â') {
						goto l843
					}
					position++
					goto l839
				l843:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('å') {
						goto l844
					}
					position++
					goto l839
				l844:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('ã') {
						goto l845
					}
					position++
					goto l839
				l845:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('ä') {
						goto l846
					}
					position++
					goto l839
				l846:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('á') {
						goto l847
					}
					position++
					goto l839
				l847:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('ç') {
						goto l848
					}
					position++
					goto l839
				l848:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('č') {
						goto l849
					}
					position++
					goto l839
				l849:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('é') {
						goto l850
					}
					position++
					goto l839
				l850:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('è') {
						goto l851
					}
					position++
					goto l839
				l851:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('ë') {
						goto l852
					}
					position++
					goto l839
				l852:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('í') {
						goto l853
					}
					position++
					goto l839
				l853:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('ì') {
						goto l854
					}
					position++
					goto l839
				l854:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('ï') {
						goto l855
					}
					position++
					goto l839
				l855:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('ň') {
						goto l856
					}
					position++
					goto l839
				l856:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('ñ') {
						goto l857
					}
					position++
					goto l839
				l857:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('ñ') {
						goto l858
					}
					position++
					goto l839
				l858:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('ó') {
						goto l859
					}
					position++
					goto l839
				l859:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('ò') {
						goto l860
					}
					position++
					goto l839
				l860:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('ô') {
						goto l861
					}
					position++
					goto l839
				l861:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('ø') {
						goto l862
					}
					position++
					goto l839
				l862:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('õ') {
						goto l863
					}
					position++
					goto l839
				l863:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('ö') {
						goto l864
					}
					position++
					goto l839
				l864:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('ú') {
						goto l865
					}
					position++
					goto l839
				l865:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('ù') {
						goto l866
					}
					position++
					goto l839
				l866:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('ü') {
						goto l867
					}
					position++
					goto l839
				l867:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('ŕ') {
						goto l868
					}
					position++
					goto l839
				l868:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('ř') {
						goto l869
					}
					position++
					goto l839
				l869:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('ŗ') {
						goto l870
					}
					position++
					goto l839
				l870:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('ſ') {
						goto l871
					}
					position++
					goto l839
				l871:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('š') {
						goto l872
					}
					position++
					goto l839
				l872:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('š') {
						goto l873
					}
					position++
					goto l839
				l873:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('ş') {
						goto l874
					}
					position++
					goto l839
				l874:
					position, tokenIndex = position839, tokenIndex839
					if buffer[position] != rune('ž') {
						goto l837
					}
					position++
				}
			l839:
				add(ruleLowerCharExtended, position838)
			}
			return true
		l837:
			position, tokenIndex = position837, tokenIndex837
			return false
		},
		/* 95 LowerChar <- <lASCII> */
		func() bool {
			position875, tokenIndex875 := position, tokenIndex
			{
				position876 := position
				if !_rules[rulelASCII]() {
					goto l875
				}
				add(ruleLowerChar, position876)
			}
			return true
		l875:
			position, tokenIndex = position875, tokenIndex875
			return false
		},
		/* 96 SpaceCharEOI <- <(_ / !.)> */
		func() bool {
			position877, tokenIndex877 := position, tokenIndex
			{
				position878 := position
				{
					position879, tokenIndex879 := position, tokenIndex
					if !_rules[rule_]() {
						goto l880
					}
					goto l879
				l880:
					position, tokenIndex = position879, tokenIndex879
					{
						position881, tokenIndex881 := position, tokenIndex
						if !matchDot() {
							goto l881
						}
						goto l877
					l881:
						position, tokenIndex = position881, tokenIndex881
					}
				}
			l879:
				add(ruleSpaceCharEOI, position878)
			}
			return true
		l877:
			position, tokenIndex = position877, tokenIndex877
			return false
		},
		/* 97 nums <- <[0-9]> */
		func() bool {
			position882, tokenIndex882 := position, tokenIndex
			{
				position883 := position
				if c := buffer[position]; c < rune('0') || c > rune('9') {
					goto l882
				}
				position++
				add(rulenums, position883)
			}
			return true
		l882:
			position, tokenIndex = position882, tokenIndex882
			return false
		},
		/* 98 lASCII <- <[a-z]> */
		func() bool {
			position884, tokenIndex884 := position, tokenIndex
			{
				position885 := position
				if c := buffer[position]; c < rune('a') || c > rune('z') {
					goto l884
				}
				position++
				add(rulelASCII, position885)
			}
			return true
		l884:
			position, tokenIndex = position884, tokenIndex884
			return false
		},
		/* 99 hASCII <- <[A-Z]> */
		func() bool {
			position886, tokenIndex886 := position, tokenIndex
			{
				position887 := position
				if c := buffer[position]; c < rune('A') || c > rune('Z') {
					goto l886
				}
				position++
				add(rulehASCII, position887)
			}
			return true
		l886:
			position, tokenIndex = position886, tokenIndex886
			return false
		},
		/* 100 apostr <- <'\''> */
		func() bool {
			position888, tokenIndex888 := position, tokenIndex
			{
				position889 := position
				if buffer[position] != rune('\'') {
					goto l888
				}
				position++
				add(ruleapostr, position889)
			}
			return true
		l888:
			position, tokenIndex = position888, tokenIndex888
			return false
		},
		/* 101 dash <- <'-'> */
		func() bool {
			position890, tokenIndex890 := position, tokenIndex
			{
				position891 := position
				if buffer[position] != rune('-') {
					goto l890
				}
				position++
				add(ruledash, position891)
			}
			return true
		l890:
			position, tokenIndex = position890, tokenIndex890
			return false
		},
		/* 102 _ <- <(MultipleSpace / SingleSpace)> */
		func() bool {
			position892, tokenIndex892 := position, tokenIndex
			{
				position893 := position
				{
					position894, tokenIndex894 := position, tokenIndex
					if !_rules[ruleMultipleSpace]() {
						goto l895
					}
					goto l894
				l895:
					position, tokenIndex = position894, tokenIndex894
					if !_rules[ruleSingleSpace]() {
						goto l892
					}
				}
			l894:
				add(rule_, position893)
			}
			return true
		l892:
			position, tokenIndex = position892, tokenIndex892
			return false
		},
		/* 103 MultipleSpace <- <(SingleSpace SingleSpace+)> */
		func() bool {
			position896, tokenIndex896 := position, tokenIndex
			{
				position897 := position
				if !_rules[ruleSingleSpace]() {
					goto l896
				}
				if !_rules[ruleSingleSpace]() {
					goto l896
				}
			l898:
				{
					position899, tokenIndex899 := position, tokenIndex
					if !_rules[ruleSingleSpace]() {
						goto l899
					}
					goto l898
				l899:
					position, tokenIndex = position899, tokenIndex899
				}
				add(ruleMultipleSpace, position897)
			}
			return true
		l896:
			position, tokenIndex = position896, tokenIndex896
			return false
		},
		/* 104 SingleSpace <- <(' ' / OtherSpace)> */
		func() bool {
			position900, tokenIndex900 := position, tokenIndex
			{
				position901 := position
				{
					position902, tokenIndex902 := position, tokenIndex
					if buffer[position] != rune(' ') {
						goto l903
					}
					position++
					goto l902
				l903:
					position, tokenIndex = position902, tokenIndex902
					if !_rules[ruleOtherSpace]() {
						goto l900
					}
				}
			l902:
				add(ruleSingleSpace, position901)
			}
			return true
		l900:
			position, tokenIndex = position900, tokenIndex900
			return false
		},
		/* 105 OtherSpace <- <('\u3000' / '\u00a0' / '\t' / '\r' / '\n' / '\f' / '\v')> */
		func() bool {
			position904, tokenIndex904 := position, tokenIndex
			{
				position905 := position
				{
					position906, tokenIndex906 := position, tokenIndex
					if buffer[position] != rune('\u3000') {
						goto l907
					}
					position++
					goto l906
				l907:
					position, tokenIndex = position906, tokenIndex906
					if buffer[position] != rune('\u00a0') {
						goto l908
					}
					position++
					goto l906
				l908:
					position, tokenIndex = position906, tokenIndex906
					if buffer[position] != rune('\t') {
						goto l909
					}
					position++
					goto l906
				l909:
					position, tokenIndex = position906, tokenIndex906
					if buffer[position] != rune('\r') {
						goto l910
					}
					position++
					goto l906
				l910:
					position, tokenIndex = position906, tokenIndex906
					if buffer[position] != rune('\n') {
						goto l911
					}
					position++
					goto l906
				l911:
					position, tokenIndex = position906, tokenIndex906
					if buffer[position] != rune('\f') {
						goto l912
					}
					position++
					goto l906
				l912:
					position, tokenIndex = position906, tokenIndex906
					if buffer[position] != rune('\v') {
						goto l904
					}
					position++
				}
			l906:
				add(ruleOtherSpace, position905)
			}
			return true
		l904:
			position, tokenIndex = position904, tokenIndex904
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
