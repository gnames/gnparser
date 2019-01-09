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
		/* 1 Tail <- <(_ .*)?> */
		func() bool {
			{
				position6 := position
				{
					position7, tokenIndex7 := position, tokenIndex
					if !_rules[rule_]() {
						goto l7
					}
				l9:
					{
						position10, tokenIndex10 := position, tokenIndex
						if !matchDot() {
							goto l10
						}
						goto l9
					l10:
						position, tokenIndex = position10, tokenIndex10
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
			position11, tokenIndex11 := position, tokenIndex
			{
				position12 := position
				{
					position13, tokenIndex13 := position, tokenIndex
					if !_rules[ruleNamedHybrid]() {
						goto l14
					}
					goto l13
				l14:
					position, tokenIndex = position13, tokenIndex13
					if !_rules[ruleHybridFormula]() {
						goto l15
					}
					goto l13
				l15:
					position, tokenIndex = position13, tokenIndex13
					if !_rules[ruleSingleName]() {
						goto l11
					}
				}
			l13:
				add(ruleName, position12)
			}
			return true
		l11:
			position, tokenIndex = position11, tokenIndex11
			return false
		},
		/* 3 HybridFormula <- <(SingleName (_ (HybridFormulaPart / HybridFormulaFull))+)> */
		func() bool {
			position16, tokenIndex16 := position, tokenIndex
			{
				position17 := position
				if !_rules[ruleSingleName]() {
					goto l16
				}
				if !_rules[rule_]() {
					goto l16
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
						goto l16
					}
				}
			l20:
			l18:
				{
					position19, tokenIndex19 := position, tokenIndex
					if !_rules[rule_]() {
						goto l19
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
							goto l19
						}
					}
				l22:
					goto l18
				l19:
					position, tokenIndex = position19, tokenIndex19
				}
				add(ruleHybridFormula, position17)
			}
			return true
		l16:
			position, tokenIndex = position16, tokenIndex16
			return false
		},
		/* 4 HybridFormulaFull <- <(HybridChar (_ SingleName)?)> */
		func() bool {
			position24, tokenIndex24 := position, tokenIndex
			{
				position25 := position
				if !_rules[ruleHybridChar]() {
					goto l24
				}
				{
					position26, tokenIndex26 := position, tokenIndex
					if !_rules[rule_]() {
						goto l26
					}
					if !_rules[ruleSingleName]() {
						goto l26
					}
					goto l27
				l26:
					position, tokenIndex = position26, tokenIndex26
				}
			l27:
				add(ruleHybridFormulaFull, position25)
			}
			return true
		l24:
			position, tokenIndex = position24, tokenIndex24
			return false
		},
		/* 5 HybridFormulaPart <- <(HybridChar _ SpeciesEpithet (_ InfraspGroup)?)> */
		func() bool {
			position28, tokenIndex28 := position, tokenIndex
			{
				position29 := position
				if !_rules[ruleHybridChar]() {
					goto l28
				}
				if !_rules[rule_]() {
					goto l28
				}
				if !_rules[ruleSpeciesEpithet]() {
					goto l28
				}
				{
					position30, tokenIndex30 := position, tokenIndex
					if !_rules[rule_]() {
						goto l30
					}
					if !_rules[ruleInfraspGroup]() {
						goto l30
					}
					goto l31
				l30:
					position, tokenIndex = position30, tokenIndex30
				}
			l31:
				add(ruleHybridFormulaPart, position29)
			}
			return true
		l28:
			position, tokenIndex = position28, tokenIndex28
			return false
		},
		/* 6 NamedHybrid <- <(NamedGenusHybrid / NamedSpeciesHybrid)> */
		func() bool {
			position32, tokenIndex32 := position, tokenIndex
			{
				position33 := position
				{
					position34, tokenIndex34 := position, tokenIndex
					if !_rules[ruleNamedGenusHybrid]() {
						goto l35
					}
					goto l34
				l35:
					position, tokenIndex = position34, tokenIndex34
					if !_rules[ruleNamedSpeciesHybrid]() {
						goto l32
					}
				}
			l34:
				add(ruleNamedHybrid, position33)
			}
			return true
		l32:
			position, tokenIndex = position32, tokenIndex32
			return false
		},
		/* 7 NamedSpeciesHybrid <- <(GenusWord _ HybridChar _? SpeciesEpithet)> */
		func() bool {
			position36, tokenIndex36 := position, tokenIndex
			{
				position37 := position
				if !_rules[ruleGenusWord]() {
					goto l36
				}
				if !_rules[rule_]() {
					goto l36
				}
				if !_rules[ruleHybridChar]() {
					goto l36
				}
				{
					position38, tokenIndex38 := position, tokenIndex
					if !_rules[rule_]() {
						goto l38
					}
					goto l39
				l38:
					position, tokenIndex = position38, tokenIndex38
				}
			l39:
				if !_rules[ruleSpeciesEpithet]() {
					goto l36
				}
				add(ruleNamedSpeciesHybrid, position37)
			}
			return true
		l36:
			position, tokenIndex = position36, tokenIndex36
			return false
		},
		/* 8 NamedGenusHybrid <- <(HybridChar _? SingleName)> */
		func() bool {
			position40, tokenIndex40 := position, tokenIndex
			{
				position41 := position
				if !_rules[ruleHybridChar]() {
					goto l40
				}
				{
					position42, tokenIndex42 := position, tokenIndex
					if !_rules[rule_]() {
						goto l42
					}
					goto l43
				l42:
					position, tokenIndex = position42, tokenIndex42
				}
			l43:
				if !_rules[ruleSingleName]() {
					goto l40
				}
				add(ruleNamedGenusHybrid, position41)
			}
			return true
		l40:
			position, tokenIndex = position40, tokenIndex40
			return false
		},
		/* 9 SingleName <- <(NameComp / NameApprox / NameSpecies / NameUninomial)> */
		func() bool {
			position44, tokenIndex44 := position, tokenIndex
			{
				position45 := position
				{
					position46, tokenIndex46 := position, tokenIndex
					if !_rules[ruleNameComp]() {
						goto l47
					}
					goto l46
				l47:
					position, tokenIndex = position46, tokenIndex46
					if !_rules[ruleNameApprox]() {
						goto l48
					}
					goto l46
				l48:
					position, tokenIndex = position46, tokenIndex46
					if !_rules[ruleNameSpecies]() {
						goto l49
					}
					goto l46
				l49:
					position, tokenIndex = position46, tokenIndex46
					if !_rules[ruleNameUninomial]() {
						goto l44
					}
				}
			l46:
				add(ruleSingleName, position45)
			}
			return true
		l44:
			position, tokenIndex = position44, tokenIndex44
			return false
		},
		/* 10 NameUninomial <- <(UninomialCombo / Uninomial)> */
		func() bool {
			position50, tokenIndex50 := position, tokenIndex
			{
				position51 := position
				{
					position52, tokenIndex52 := position, tokenIndex
					if !_rules[ruleUninomialCombo]() {
						goto l53
					}
					goto l52
				l53:
					position, tokenIndex = position52, tokenIndex52
					if !_rules[ruleUninomial]() {
						goto l50
					}
				}
			l52:
				add(ruleNameUninomial, position51)
			}
			return true
		l50:
			position, tokenIndex = position50, tokenIndex50
			return false
		},
		/* 11 NameApprox <- <(GenusWord (_ SpeciesEpithet)? _ Approximation ApproxNameIgnored)> */
		func() bool {
			position54, tokenIndex54 := position, tokenIndex
			{
				position55 := position
				if !_rules[ruleGenusWord]() {
					goto l54
				}
				{
					position56, tokenIndex56 := position, tokenIndex
					if !_rules[rule_]() {
						goto l56
					}
					if !_rules[ruleSpeciesEpithet]() {
						goto l56
					}
					goto l57
				l56:
					position, tokenIndex = position56, tokenIndex56
				}
			l57:
				if !_rules[rule_]() {
					goto l54
				}
				if !_rules[ruleApproximation]() {
					goto l54
				}
				if !_rules[ruleApproxNameIgnored]() {
					goto l54
				}
				add(ruleNameApprox, position55)
			}
			return true
		l54:
			position, tokenIndex = position54, tokenIndex54
			return false
		},
		/* 12 NameComp <- <(GenusWord _ Comparison (_ SpeciesEpithet)?)> */
		func() bool {
			position58, tokenIndex58 := position, tokenIndex
			{
				position59 := position
				if !_rules[ruleGenusWord]() {
					goto l58
				}
				if !_rules[rule_]() {
					goto l58
				}
				if !_rules[ruleComparison]() {
					goto l58
				}
				{
					position60, tokenIndex60 := position, tokenIndex
					if !_rules[rule_]() {
						goto l60
					}
					if !_rules[ruleSpeciesEpithet]() {
						goto l60
					}
					goto l61
				l60:
					position, tokenIndex = position60, tokenIndex60
				}
			l61:
				add(ruleNameComp, position59)
			}
			return true
		l58:
			position, tokenIndex = position58, tokenIndex58
			return false
		},
		/* 13 NameSpecies <- <(GenusWord (_? (SubGenus / SubGenusOrSuperspecies))? _ SpeciesEpithet (_ InfraspGroup)?)> */
		func() bool {
			position62, tokenIndex62 := position, tokenIndex
			{
				position63 := position
				if !_rules[ruleGenusWord]() {
					goto l62
				}
				{
					position64, tokenIndex64 := position, tokenIndex
					{
						position66, tokenIndex66 := position, tokenIndex
						if !_rules[rule_]() {
							goto l66
						}
						goto l67
					l66:
						position, tokenIndex = position66, tokenIndex66
					}
				l67:
					{
						position68, tokenIndex68 := position, tokenIndex
						if !_rules[ruleSubGenus]() {
							goto l69
						}
						goto l68
					l69:
						position, tokenIndex = position68, tokenIndex68
						if !_rules[ruleSubGenusOrSuperspecies]() {
							goto l64
						}
					}
				l68:
					goto l65
				l64:
					position, tokenIndex = position64, tokenIndex64
				}
			l65:
				if !_rules[rule_]() {
					goto l62
				}
				if !_rules[ruleSpeciesEpithet]() {
					goto l62
				}
				{
					position70, tokenIndex70 := position, tokenIndex
					if !_rules[rule_]() {
						goto l70
					}
					if !_rules[ruleInfraspGroup]() {
						goto l70
					}
					goto l71
				l70:
					position, tokenIndex = position70, tokenIndex70
				}
			l71:
				add(ruleNameSpecies, position63)
			}
			return true
		l62:
			position, tokenIndex = position62, tokenIndex62
			return false
		},
		/* 14 GenusWord <- <((AbbrGenus / UninomialWord) !(_ AuthorWord))> */
		func() bool {
			position72, tokenIndex72 := position, tokenIndex
			{
				position73 := position
				{
					position74, tokenIndex74 := position, tokenIndex
					if !_rules[ruleAbbrGenus]() {
						goto l75
					}
					goto l74
				l75:
					position, tokenIndex = position74, tokenIndex74
					if !_rules[ruleUninomialWord]() {
						goto l72
					}
				}
			l74:
				{
					position76, tokenIndex76 := position, tokenIndex
					if !_rules[rule_]() {
						goto l76
					}
					if !_rules[ruleAuthorWord]() {
						goto l76
					}
					goto l72
				l76:
					position, tokenIndex = position76, tokenIndex76
				}
				add(ruleGenusWord, position73)
			}
			return true
		l72:
			position, tokenIndex = position72, tokenIndex72
			return false
		},
		/* 15 InfraspGroup <- <(InfraspEpithet (_ InfraspEpithet)? (_ InfraspEpithet)?)> */
		func() bool {
			position77, tokenIndex77 := position, tokenIndex
			{
				position78 := position
				if !_rules[ruleInfraspEpithet]() {
					goto l77
				}
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
				add(ruleInfraspGroup, position78)
			}
			return true
		l77:
			position, tokenIndex = position77, tokenIndex77
			return false
		},
		/* 16 InfraspEpithet <- <((Rank _?)? !AuthorEx Word (_ Authorship)?)> */
		func() bool {
			position83, tokenIndex83 := position, tokenIndex
			{
				position84 := position
				{
					position85, tokenIndex85 := position, tokenIndex
					if !_rules[ruleRank]() {
						goto l85
					}
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
					goto l86
				l85:
					position, tokenIndex = position85, tokenIndex85
				}
			l86:
				{
					position89, tokenIndex89 := position, tokenIndex
					if !_rules[ruleAuthorEx]() {
						goto l89
					}
					goto l83
				l89:
					position, tokenIndex = position89, tokenIndex89
				}
				if !_rules[ruleWord]() {
					goto l83
				}
				{
					position90, tokenIndex90 := position, tokenIndex
					if !_rules[rule_]() {
						goto l90
					}
					if !_rules[ruleAuthorship]() {
						goto l90
					}
					goto l91
				l90:
					position, tokenIndex = position90, tokenIndex90
				}
			l91:
				add(ruleInfraspEpithet, position84)
			}
			return true
		l83:
			position, tokenIndex = position83, tokenIndex83
			return false
		},
		/* 17 SpeciesEpithet <- <(!AuthorEx Word (_? Authorship)? ','? &(SpaceCharEOI / '('))> */
		func() bool {
			position92, tokenIndex92 := position, tokenIndex
			{
				position93 := position
				{
					position94, tokenIndex94 := position, tokenIndex
					if !_rules[ruleAuthorEx]() {
						goto l94
					}
					goto l92
				l94:
					position, tokenIndex = position94, tokenIndex94
				}
				if !_rules[ruleWord]() {
					goto l92
				}
				{
					position95, tokenIndex95 := position, tokenIndex
					{
						position97, tokenIndex97 := position, tokenIndex
						if !_rules[rule_]() {
							goto l97
						}
						goto l98
					l97:
						position, tokenIndex = position97, tokenIndex97
					}
				l98:
					if !_rules[ruleAuthorship]() {
						goto l95
					}
					goto l96
				l95:
					position, tokenIndex = position95, tokenIndex95
				}
			l96:
				{
					position99, tokenIndex99 := position, tokenIndex
					if buffer[position] != rune(',') {
						goto l99
					}
					position++
					goto l100
				l99:
					position, tokenIndex = position99, tokenIndex99
				}
			l100:
				{
					position101, tokenIndex101 := position, tokenIndex
					{
						position102, tokenIndex102 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l103
						}
						goto l102
					l103:
						position, tokenIndex = position102, tokenIndex102
						if buffer[position] != rune('(') {
							goto l92
						}
						position++
					}
				l102:
					position, tokenIndex = position101, tokenIndex101
				}
				add(ruleSpeciesEpithet, position93)
			}
			return true
		l92:
			position, tokenIndex = position92, tokenIndex92
			return false
		},
		/* 18 Comparison <- <('c' 'f' '.'?)> */
		func() bool {
			position104, tokenIndex104 := position, tokenIndex
			{
				position105 := position
				if buffer[position] != rune('c') {
					goto l104
				}
				position++
				if buffer[position] != rune('f') {
					goto l104
				}
				position++
				{
					position106, tokenIndex106 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l106
					}
					position++
					goto l107
				l106:
					position, tokenIndex = position106, tokenIndex106
				}
			l107:
				add(ruleComparison, position105)
			}
			return true
		l104:
			position, tokenIndex = position104, tokenIndex104
			return false
		},
		/* 19 Rank <- <(RankForma / RankVar / RankSsp / RankOther / RankOtherUncommon)> */
		func() bool {
			position108, tokenIndex108 := position, tokenIndex
			{
				position109 := position
				{
					position110, tokenIndex110 := position, tokenIndex
					if !_rules[ruleRankForma]() {
						goto l111
					}
					goto l110
				l111:
					position, tokenIndex = position110, tokenIndex110
					if !_rules[ruleRankVar]() {
						goto l112
					}
					goto l110
				l112:
					position, tokenIndex = position110, tokenIndex110
					if !_rules[ruleRankSsp]() {
						goto l113
					}
					goto l110
				l113:
					position, tokenIndex = position110, tokenIndex110
					if !_rules[ruleRankOther]() {
						goto l114
					}
					goto l110
				l114:
					position, tokenIndex = position110, tokenIndex110
					if !_rules[ruleRankOtherUncommon]() {
						goto l108
					}
				}
			l110:
				add(ruleRank, position109)
			}
			return true
		l108:
			position, tokenIndex = position108, tokenIndex108
			return false
		},
		/* 20 RankOtherUncommon <- <(('*' / ('n' 'a' 't') / ('f' '.' 's' 'p') / ('m' 'u' 't' '.')) &SpaceCharEOI)> */
		func() bool {
			position115, tokenIndex115 := position, tokenIndex
			{
				position116 := position
				{
					position117, tokenIndex117 := position, tokenIndex
					if buffer[position] != rune('*') {
						goto l118
					}
					position++
					goto l117
				l118:
					position, tokenIndex = position117, tokenIndex117
					if buffer[position] != rune('n') {
						goto l119
					}
					position++
					if buffer[position] != rune('a') {
						goto l119
					}
					position++
					if buffer[position] != rune('t') {
						goto l119
					}
					position++
					goto l117
				l119:
					position, tokenIndex = position117, tokenIndex117
					if buffer[position] != rune('f') {
						goto l120
					}
					position++
					if buffer[position] != rune('.') {
						goto l120
					}
					position++
					if buffer[position] != rune('s') {
						goto l120
					}
					position++
					if buffer[position] != rune('p') {
						goto l120
					}
					position++
					goto l117
				l120:
					position, tokenIndex = position117, tokenIndex117
					if buffer[position] != rune('m') {
						goto l115
					}
					position++
					if buffer[position] != rune('u') {
						goto l115
					}
					position++
					if buffer[position] != rune('t') {
						goto l115
					}
					position++
					if buffer[position] != rune('.') {
						goto l115
					}
					position++
				}
			l117:
				{
					position121, tokenIndex121 := position, tokenIndex
					if !_rules[ruleSpaceCharEOI]() {
						goto l115
					}
					position, tokenIndex = position121, tokenIndex121
				}
				add(ruleRankOtherUncommon, position116)
			}
			return true
		l115:
			position, tokenIndex = position115, tokenIndex115
			return false
		},
		/* 21 RankOther <- <((('m' 'o' 'r' 'p' 'h' '.') / ('n' 'o' 't' 'h' 'o' 's' 'u' 'b' 's' 'p' '.') / ('c' 'o' 'n' 'v' 'a' 'r' '.') / ('p' 's' 'e' 'u' 'd' 'o' 'v' 'a' 'r' '.') / ('s' 'e' 'c' 't' '.') / ('s' 'e' 'r' '.') / ('s' 'u' 'b' 'v' 'a' 'r' '.') / ('s' 'u' 'b' 'f' '.') / ('r' 'a' 'c' 'e') / 'α' / ('β' 'β') / 'β' / 'γ' / 'δ' / 'ε' / 'φ' / 'θ' / 'μ' / ('a' '.') / ('b' '.') / ('c' '.') / ('d' '.') / ('e' '.') / ('g' '.') / ('k' '.') / ('p' 'v' '.') / ('p' 'a' 't' 'h' 'o' 'v' 'a' 'r' '.') / ('a' 'b' '.' (_? ('n' '.'))?) / ('s' 't' '.')) &SpaceCharEOI)> */
		func() bool {
			position122, tokenIndex122 := position, tokenIndex
			{
				position123 := position
				{
					position124, tokenIndex124 := position, tokenIndex
					if buffer[position] != rune('m') {
						goto l125
					}
					position++
					if buffer[position] != rune('o') {
						goto l125
					}
					position++
					if buffer[position] != rune('r') {
						goto l125
					}
					position++
					if buffer[position] != rune('p') {
						goto l125
					}
					position++
					if buffer[position] != rune('h') {
						goto l125
					}
					position++
					if buffer[position] != rune('.') {
						goto l125
					}
					position++
					goto l124
				l125:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('n') {
						goto l126
					}
					position++
					if buffer[position] != rune('o') {
						goto l126
					}
					position++
					if buffer[position] != rune('t') {
						goto l126
					}
					position++
					if buffer[position] != rune('h') {
						goto l126
					}
					position++
					if buffer[position] != rune('o') {
						goto l126
					}
					position++
					if buffer[position] != rune('s') {
						goto l126
					}
					position++
					if buffer[position] != rune('u') {
						goto l126
					}
					position++
					if buffer[position] != rune('b') {
						goto l126
					}
					position++
					if buffer[position] != rune('s') {
						goto l126
					}
					position++
					if buffer[position] != rune('p') {
						goto l126
					}
					position++
					if buffer[position] != rune('.') {
						goto l126
					}
					position++
					goto l124
				l126:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('c') {
						goto l127
					}
					position++
					if buffer[position] != rune('o') {
						goto l127
					}
					position++
					if buffer[position] != rune('n') {
						goto l127
					}
					position++
					if buffer[position] != rune('v') {
						goto l127
					}
					position++
					if buffer[position] != rune('a') {
						goto l127
					}
					position++
					if buffer[position] != rune('r') {
						goto l127
					}
					position++
					if buffer[position] != rune('.') {
						goto l127
					}
					position++
					goto l124
				l127:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('p') {
						goto l128
					}
					position++
					if buffer[position] != rune('s') {
						goto l128
					}
					position++
					if buffer[position] != rune('e') {
						goto l128
					}
					position++
					if buffer[position] != rune('u') {
						goto l128
					}
					position++
					if buffer[position] != rune('d') {
						goto l128
					}
					position++
					if buffer[position] != rune('o') {
						goto l128
					}
					position++
					if buffer[position] != rune('v') {
						goto l128
					}
					position++
					if buffer[position] != rune('a') {
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
					goto l124
				l128:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('s') {
						goto l129
					}
					position++
					if buffer[position] != rune('e') {
						goto l129
					}
					position++
					if buffer[position] != rune('c') {
						goto l129
					}
					position++
					if buffer[position] != rune('t') {
						goto l129
					}
					position++
					if buffer[position] != rune('.') {
						goto l129
					}
					position++
					goto l124
				l129:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('s') {
						goto l130
					}
					position++
					if buffer[position] != rune('e') {
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
					goto l124
				l130:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('s') {
						goto l131
					}
					position++
					if buffer[position] != rune('u') {
						goto l131
					}
					position++
					if buffer[position] != rune('b') {
						goto l131
					}
					position++
					if buffer[position] != rune('v') {
						goto l131
					}
					position++
					if buffer[position] != rune('a') {
						goto l131
					}
					position++
					if buffer[position] != rune('r') {
						goto l131
					}
					position++
					if buffer[position] != rune('.') {
						goto l131
					}
					position++
					goto l124
				l131:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('s') {
						goto l132
					}
					position++
					if buffer[position] != rune('u') {
						goto l132
					}
					position++
					if buffer[position] != rune('b') {
						goto l132
					}
					position++
					if buffer[position] != rune('f') {
						goto l132
					}
					position++
					if buffer[position] != rune('.') {
						goto l132
					}
					position++
					goto l124
				l132:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('r') {
						goto l133
					}
					position++
					if buffer[position] != rune('a') {
						goto l133
					}
					position++
					if buffer[position] != rune('c') {
						goto l133
					}
					position++
					if buffer[position] != rune('e') {
						goto l133
					}
					position++
					goto l124
				l133:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('α') {
						goto l134
					}
					position++
					goto l124
				l134:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('β') {
						goto l135
					}
					position++
					if buffer[position] != rune('β') {
						goto l135
					}
					position++
					goto l124
				l135:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('β') {
						goto l136
					}
					position++
					goto l124
				l136:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('γ') {
						goto l137
					}
					position++
					goto l124
				l137:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('δ') {
						goto l138
					}
					position++
					goto l124
				l138:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('ε') {
						goto l139
					}
					position++
					goto l124
				l139:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('φ') {
						goto l140
					}
					position++
					goto l124
				l140:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('θ') {
						goto l141
					}
					position++
					goto l124
				l141:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('μ') {
						goto l142
					}
					position++
					goto l124
				l142:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('a') {
						goto l143
					}
					position++
					if buffer[position] != rune('.') {
						goto l143
					}
					position++
					goto l124
				l143:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('b') {
						goto l144
					}
					position++
					if buffer[position] != rune('.') {
						goto l144
					}
					position++
					goto l124
				l144:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('c') {
						goto l145
					}
					position++
					if buffer[position] != rune('.') {
						goto l145
					}
					position++
					goto l124
				l145:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('d') {
						goto l146
					}
					position++
					if buffer[position] != rune('.') {
						goto l146
					}
					position++
					goto l124
				l146:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('e') {
						goto l147
					}
					position++
					if buffer[position] != rune('.') {
						goto l147
					}
					position++
					goto l124
				l147:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('g') {
						goto l148
					}
					position++
					if buffer[position] != rune('.') {
						goto l148
					}
					position++
					goto l124
				l148:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('k') {
						goto l149
					}
					position++
					if buffer[position] != rune('.') {
						goto l149
					}
					position++
					goto l124
				l149:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('p') {
						goto l150
					}
					position++
					if buffer[position] != rune('v') {
						goto l150
					}
					position++
					if buffer[position] != rune('.') {
						goto l150
					}
					position++
					goto l124
				l150:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('p') {
						goto l151
					}
					position++
					if buffer[position] != rune('a') {
						goto l151
					}
					position++
					if buffer[position] != rune('t') {
						goto l151
					}
					position++
					if buffer[position] != rune('h') {
						goto l151
					}
					position++
					if buffer[position] != rune('o') {
						goto l151
					}
					position++
					if buffer[position] != rune('v') {
						goto l151
					}
					position++
					if buffer[position] != rune('a') {
						goto l151
					}
					position++
					if buffer[position] != rune('r') {
						goto l151
					}
					position++
					if buffer[position] != rune('.') {
						goto l151
					}
					position++
					goto l124
				l151:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('a') {
						goto l152
					}
					position++
					if buffer[position] != rune('b') {
						goto l152
					}
					position++
					if buffer[position] != rune('.') {
						goto l152
					}
					position++
					{
						position153, tokenIndex153 := position, tokenIndex
						{
							position155, tokenIndex155 := position, tokenIndex
							if !_rules[rule_]() {
								goto l155
							}
							goto l156
						l155:
							position, tokenIndex = position155, tokenIndex155
						}
					l156:
						if buffer[position] != rune('n') {
							goto l153
						}
						position++
						if buffer[position] != rune('.') {
							goto l153
						}
						position++
						goto l154
					l153:
						position, tokenIndex = position153, tokenIndex153
					}
				l154:
					goto l124
				l152:
					position, tokenIndex = position124, tokenIndex124
					if buffer[position] != rune('s') {
						goto l122
					}
					position++
					if buffer[position] != rune('t') {
						goto l122
					}
					position++
					if buffer[position] != rune('.') {
						goto l122
					}
					position++
				}
			l124:
				{
					position157, tokenIndex157 := position, tokenIndex
					if !_rules[ruleSpaceCharEOI]() {
						goto l122
					}
					position, tokenIndex = position157, tokenIndex157
				}
				add(ruleRankOther, position123)
			}
			return true
		l122:
			position, tokenIndex = position122, tokenIndex122
			return false
		},
		/* 22 RankVar <- <(('v' 'a' 'r' 'i' 'e' 't' 'y') / ('[' 'v' 'a' 'r' '.' ']') / ('n' 'v' 'a' 'r' '.') / ('v' 'a' 'r' (&SpaceCharEOI / '.')))> */
		func() bool {
			position158, tokenIndex158 := position, tokenIndex
			{
				position159 := position
				{
					position160, tokenIndex160 := position, tokenIndex
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
					if buffer[position] != rune('i') {
						goto l161
					}
					position++
					if buffer[position] != rune('e') {
						goto l161
					}
					position++
					if buffer[position] != rune('t') {
						goto l161
					}
					position++
					if buffer[position] != rune('y') {
						goto l161
					}
					position++
					goto l160
				l161:
					position, tokenIndex = position160, tokenIndex160
					if buffer[position] != rune('[') {
						goto l162
					}
					position++
					if buffer[position] != rune('v') {
						goto l162
					}
					position++
					if buffer[position] != rune('a') {
						goto l162
					}
					position++
					if buffer[position] != rune('r') {
						goto l162
					}
					position++
					if buffer[position] != rune('.') {
						goto l162
					}
					position++
					if buffer[position] != rune(']') {
						goto l162
					}
					position++
					goto l160
				l162:
					position, tokenIndex = position160, tokenIndex160
					if buffer[position] != rune('n') {
						goto l163
					}
					position++
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
					if buffer[position] != rune('.') {
						goto l163
					}
					position++
					goto l160
				l163:
					position, tokenIndex = position160, tokenIndex160
					if buffer[position] != rune('v') {
						goto l158
					}
					position++
					if buffer[position] != rune('a') {
						goto l158
					}
					position++
					if buffer[position] != rune('r') {
						goto l158
					}
					position++
					{
						position164, tokenIndex164 := position, tokenIndex
						{
							position166, tokenIndex166 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l165
							}
							position, tokenIndex = position166, tokenIndex166
						}
						goto l164
					l165:
						position, tokenIndex = position164, tokenIndex164
						if buffer[position] != rune('.') {
							goto l158
						}
						position++
					}
				l164:
				}
			l160:
				add(ruleRankVar, position159)
			}
			return true
		l158:
			position, tokenIndex = position158, tokenIndex158
			return false
		},
		/* 23 RankForma <- <((('f' 'o' 'r' 'm' 'a') / ('f' 'm' 'a') / ('f' 'o' 'r' 'm') / ('f' 'o') / 'f') (&SpaceCharEOI / '.'))> */
		func() bool {
			position167, tokenIndex167 := position, tokenIndex
			{
				position168 := position
				{
					position169, tokenIndex169 := position, tokenIndex
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
					if buffer[position] != rune('a') {
						goto l170
					}
					position++
					goto l169
				l170:
					position, tokenIndex = position169, tokenIndex169
					if buffer[position] != rune('f') {
						goto l171
					}
					position++
					if buffer[position] != rune('m') {
						goto l171
					}
					position++
					if buffer[position] != rune('a') {
						goto l171
					}
					position++
					goto l169
				l171:
					position, tokenIndex = position169, tokenIndex169
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
					goto l169
				l172:
					position, tokenIndex = position169, tokenIndex169
					if buffer[position] != rune('f') {
						goto l173
					}
					position++
					if buffer[position] != rune('o') {
						goto l173
					}
					position++
					goto l169
				l173:
					position, tokenIndex = position169, tokenIndex169
					if buffer[position] != rune('f') {
						goto l167
					}
					position++
				}
			l169:
				{
					position174, tokenIndex174 := position, tokenIndex
					{
						position176, tokenIndex176 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l175
						}
						position, tokenIndex = position176, tokenIndex176
					}
					goto l174
				l175:
					position, tokenIndex = position174, tokenIndex174
					if buffer[position] != rune('.') {
						goto l167
					}
					position++
				}
			l174:
				add(ruleRankForma, position168)
			}
			return true
		l167:
			position, tokenIndex = position167, tokenIndex167
			return false
		},
		/* 24 RankSsp <- <((('s' 's' 'p') / ('s' 'u' 'b' 's' 'p')) (&SpaceCharEOI / '.'))> */
		func() bool {
			position177, tokenIndex177 := position, tokenIndex
			{
				position178 := position
				{
					position179, tokenIndex179 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l180
					}
					position++
					if buffer[position] != rune('s') {
						goto l180
					}
					position++
					if buffer[position] != rune('p') {
						goto l180
					}
					position++
					goto l179
				l180:
					position, tokenIndex = position179, tokenIndex179
					if buffer[position] != rune('s') {
						goto l177
					}
					position++
					if buffer[position] != rune('u') {
						goto l177
					}
					position++
					if buffer[position] != rune('b') {
						goto l177
					}
					position++
					if buffer[position] != rune('s') {
						goto l177
					}
					position++
					if buffer[position] != rune('p') {
						goto l177
					}
					position++
				}
			l179:
				{
					position181, tokenIndex181 := position, tokenIndex
					{
						position183, tokenIndex183 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l182
						}
						position, tokenIndex = position183, tokenIndex183
					}
					goto l181
				l182:
					position, tokenIndex = position181, tokenIndex181
					if buffer[position] != rune('.') {
						goto l177
					}
					position++
				}
			l181:
				add(ruleRankSsp, position178)
			}
			return true
		l177:
			position, tokenIndex = position177, tokenIndex177
			return false
		},
		/* 25 SubGenusOrSuperspecies <- <('(' _? NameLowerChar+ _? ')')> */
		func() bool {
			position184, tokenIndex184 := position, tokenIndex
			{
				position185 := position
				if buffer[position] != rune('(') {
					goto l184
				}
				position++
				{
					position186, tokenIndex186 := position, tokenIndex
					if !_rules[rule_]() {
						goto l186
					}
					goto l187
				l186:
					position, tokenIndex = position186, tokenIndex186
				}
			l187:
				if !_rules[ruleNameLowerChar]() {
					goto l184
				}
			l188:
				{
					position189, tokenIndex189 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l189
					}
					goto l188
				l189:
					position, tokenIndex = position189, tokenIndex189
				}
				{
					position190, tokenIndex190 := position, tokenIndex
					if !_rules[rule_]() {
						goto l190
					}
					goto l191
				l190:
					position, tokenIndex = position190, tokenIndex190
				}
			l191:
				if buffer[position] != rune(')') {
					goto l184
				}
				position++
				add(ruleSubGenusOrSuperspecies, position185)
			}
			return true
		l184:
			position, tokenIndex = position184, tokenIndex184
			return false
		},
		/* 26 SubGenus <- <('(' _? UninomialWord _? ')')> */
		func() bool {
			position192, tokenIndex192 := position, tokenIndex
			{
				position193 := position
				if buffer[position] != rune('(') {
					goto l192
				}
				position++
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
				if !_rules[ruleUninomialWord]() {
					goto l192
				}
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
				if buffer[position] != rune(')') {
					goto l192
				}
				position++
				add(ruleSubGenus, position193)
			}
			return true
		l192:
			position, tokenIndex = position192, tokenIndex192
			return false
		},
		/* 27 UninomialCombo <- <(UninomialCombo1 / UninomialCombo2)> */
		func() bool {
			position198, tokenIndex198 := position, tokenIndex
			{
				position199 := position
				{
					position200, tokenIndex200 := position, tokenIndex
					if !_rules[ruleUninomialCombo1]() {
						goto l201
					}
					goto l200
				l201:
					position, tokenIndex = position200, tokenIndex200
					if !_rules[ruleUninomialCombo2]() {
						goto l198
					}
				}
			l200:
				add(ruleUninomialCombo, position199)
			}
			return true
		l198:
			position, tokenIndex = position198, tokenIndex198
			return false
		},
		/* 28 UninomialCombo1 <- <(UninomialWord _? SubGenus _? Authorship .?)> */
		func() bool {
			position202, tokenIndex202 := position, tokenIndex
			{
				position203 := position
				if !_rules[ruleUninomialWord]() {
					goto l202
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
				if !_rules[ruleSubGenus]() {
					goto l202
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
				if !_rules[ruleAuthorship]() {
					goto l202
				}
				{
					position208, tokenIndex208 := position, tokenIndex
					if !matchDot() {
						goto l208
					}
					goto l209
				l208:
					position, tokenIndex = position208, tokenIndex208
				}
			l209:
				add(ruleUninomialCombo1, position203)
			}
			return true
		l202:
			position, tokenIndex = position202, tokenIndex202
			return false
		},
		/* 29 UninomialCombo2 <- <(Uninomial _? RankUninomial _? Uninomial)> */
		func() bool {
			position210, tokenIndex210 := position, tokenIndex
			{
				position211 := position
				if !_rules[ruleUninomial]() {
					goto l210
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
				if !_rules[ruleRankUninomial]() {
					goto l210
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
				if !_rules[ruleUninomial]() {
					goto l210
				}
				add(ruleUninomialCombo2, position211)
			}
			return true
		l210:
			position, tokenIndex = position210, tokenIndex210
			return false
		},
		/* 30 RankUninomial <- <((('s' 'e' 'c' 't') / ('s' 'u' 'b' 's' 'e' 'c' 't') / ('t' 'r' 'i' 'b') / ('s' 'u' 'b' 't' 'r' 'i' 'b') / ('s' 'u' 'b' 's' 'e' 'r') / ('s' 'e' 'r') / ('s' 'u' 'b' 'g' 'e' 'n') / ('f' 'a' 'm') / ('s' 'u' 'b' 'f' 'a' 'm') / ('s' 'u' 'p' 'e' 'r' 't' 'r' 'i' 'b')) '.'?)> */
		func() bool {
			position216, tokenIndex216 := position, tokenIndex
			{
				position217 := position
				{
					position218, tokenIndex218 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l219
					}
					position++
					if buffer[position] != rune('e') {
						goto l219
					}
					position++
					if buffer[position] != rune('c') {
						goto l219
					}
					position++
					if buffer[position] != rune('t') {
						goto l219
					}
					position++
					goto l218
				l219:
					position, tokenIndex = position218, tokenIndex218
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
					if buffer[position] != rune('s') {
						goto l220
					}
					position++
					if buffer[position] != rune('e') {
						goto l220
					}
					position++
					if buffer[position] != rune('c') {
						goto l220
					}
					position++
					if buffer[position] != rune('t') {
						goto l220
					}
					position++
					goto l218
				l220:
					position, tokenIndex = position218, tokenIndex218
					if buffer[position] != rune('t') {
						goto l221
					}
					position++
					if buffer[position] != rune('r') {
						goto l221
					}
					position++
					if buffer[position] != rune('i') {
						goto l221
					}
					position++
					if buffer[position] != rune('b') {
						goto l221
					}
					position++
					goto l218
				l221:
					position, tokenIndex = position218, tokenIndex218
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
					if buffer[position] != rune('t') {
						goto l222
					}
					position++
					if buffer[position] != rune('r') {
						goto l222
					}
					position++
					if buffer[position] != rune('i') {
						goto l222
					}
					position++
					if buffer[position] != rune('b') {
						goto l222
					}
					position++
					goto l218
				l222:
					position, tokenIndex = position218, tokenIndex218
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
					if buffer[position] != rune('s') {
						goto l223
					}
					position++
					if buffer[position] != rune('e') {
						goto l223
					}
					position++
					if buffer[position] != rune('r') {
						goto l223
					}
					position++
					goto l218
				l223:
					position, tokenIndex = position218, tokenIndex218
					if buffer[position] != rune('s') {
						goto l224
					}
					position++
					if buffer[position] != rune('e') {
						goto l224
					}
					position++
					if buffer[position] != rune('r') {
						goto l224
					}
					position++
					goto l218
				l224:
					position, tokenIndex = position218, tokenIndex218
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
					if buffer[position] != rune('g') {
						goto l225
					}
					position++
					if buffer[position] != rune('e') {
						goto l225
					}
					position++
					if buffer[position] != rune('n') {
						goto l225
					}
					position++
					goto l218
				l225:
					position, tokenIndex = position218, tokenIndex218
					if buffer[position] != rune('f') {
						goto l226
					}
					position++
					if buffer[position] != rune('a') {
						goto l226
					}
					position++
					if buffer[position] != rune('m') {
						goto l226
					}
					position++
					goto l218
				l226:
					position, tokenIndex = position218, tokenIndex218
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
					if buffer[position] != rune('f') {
						goto l227
					}
					position++
					if buffer[position] != rune('a') {
						goto l227
					}
					position++
					if buffer[position] != rune('m') {
						goto l227
					}
					position++
					goto l218
				l227:
					position, tokenIndex = position218, tokenIndex218
					if buffer[position] != rune('s') {
						goto l216
					}
					position++
					if buffer[position] != rune('u') {
						goto l216
					}
					position++
					if buffer[position] != rune('p') {
						goto l216
					}
					position++
					if buffer[position] != rune('e') {
						goto l216
					}
					position++
					if buffer[position] != rune('r') {
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
				}
			l218:
				{
					position228, tokenIndex228 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l228
					}
					position++
					goto l229
				l228:
					position, tokenIndex = position228, tokenIndex228
				}
			l229:
				add(ruleRankUninomial, position217)
			}
			return true
		l216:
			position, tokenIndex = position216, tokenIndex216
			return false
		},
		/* 31 Uninomial <- <(UninomialWord (_ Authorship)?)> */
		func() bool {
			position230, tokenIndex230 := position, tokenIndex
			{
				position231 := position
				if !_rules[ruleUninomialWord]() {
					goto l230
				}
				{
					position232, tokenIndex232 := position, tokenIndex
					if !_rules[rule_]() {
						goto l232
					}
					if !_rules[ruleAuthorship]() {
						goto l232
					}
					goto l233
				l232:
					position, tokenIndex = position232, tokenIndex232
				}
			l233:
				add(ruleUninomial, position231)
			}
			return true
		l230:
			position, tokenIndex = position230, tokenIndex230
			return false
		},
		/* 32 UninomialWord <- <(CapWord / TwoLetterGenus)> */
		func() bool {
			position234, tokenIndex234 := position, tokenIndex
			{
				position235 := position
				{
					position236, tokenIndex236 := position, tokenIndex
					if !_rules[ruleCapWord]() {
						goto l237
					}
					goto l236
				l237:
					position, tokenIndex = position236, tokenIndex236
					if !_rules[ruleTwoLetterGenus]() {
						goto l234
					}
				}
			l236:
				add(ruleUninomialWord, position235)
			}
			return true
		l234:
			position, tokenIndex = position234, tokenIndex234
			return false
		},
		/* 33 AbbrGenus <- <(UpperChar LowerChar* '.')> */
		func() bool {
			position238, tokenIndex238 := position, tokenIndex
			{
				position239 := position
				if !_rules[ruleUpperChar]() {
					goto l238
				}
			l240:
				{
					position241, tokenIndex241 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l241
					}
					goto l240
				l241:
					position, tokenIndex = position241, tokenIndex241
				}
				if buffer[position] != rune('.') {
					goto l238
				}
				position++
				add(ruleAbbrGenus, position239)
			}
			return true
		l238:
			position, tokenIndex = position238, tokenIndex238
			return false
		},
		/* 34 CapWord <- <(CapWord2 / CapWord1)> */
		func() bool {
			position242, tokenIndex242 := position, tokenIndex
			{
				position243 := position
				{
					position244, tokenIndex244 := position, tokenIndex
					if !_rules[ruleCapWord2]() {
						goto l245
					}
					goto l244
				l245:
					position, tokenIndex = position244, tokenIndex244
					if !_rules[ruleCapWord1]() {
						goto l242
					}
				}
			l244:
				add(ruleCapWord, position243)
			}
			return true
		l242:
			position, tokenIndex = position242, tokenIndex242
			return false
		},
		/* 35 CapWord1 <- <(NameUpperChar NameLowerChar NameLowerChar+ '?'?)> */
		func() bool {
			position246, tokenIndex246 := position, tokenIndex
			{
				position247 := position
				if !_rules[ruleNameUpperChar]() {
					goto l246
				}
				if !_rules[ruleNameLowerChar]() {
					goto l246
				}
				if !_rules[ruleNameLowerChar]() {
					goto l246
				}
			l248:
				{
					position249, tokenIndex249 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l249
					}
					goto l248
				l249:
					position, tokenIndex = position249, tokenIndex249
				}
				{
					position250, tokenIndex250 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l250
					}
					position++
					goto l251
				l250:
					position, tokenIndex = position250, tokenIndex250
				}
			l251:
				add(ruleCapWord1, position247)
			}
			return true
		l246:
			position, tokenIndex = position246, tokenIndex246
			return false
		},
		/* 36 CapWord2 <- <(CapWord1 dash (CapWord1 / Word1))> */
		func() bool {
			position252, tokenIndex252 := position, tokenIndex
			{
				position253 := position
				if !_rules[ruleCapWord1]() {
					goto l252
				}
				if !_rules[ruledash]() {
					goto l252
				}
				{
					position254, tokenIndex254 := position, tokenIndex
					if !_rules[ruleCapWord1]() {
						goto l255
					}
					goto l254
				l255:
					position, tokenIndex = position254, tokenIndex254
					if !_rules[ruleWord1]() {
						goto l252
					}
				}
			l254:
				add(ruleCapWord2, position253)
			}
			return true
		l252:
			position, tokenIndex = position252, tokenIndex252
			return false
		},
		/* 37 TwoLetterGenus <- <(('C' 'a') / ('E' 'a') / ('G' 'e') / ('I' 'a') / ('I' 'o') / ('I' 'x') / ('L' 'o') / ('O' 'a') / ('R' 'a') / ('T' 'y') / ('U' 'a') / ('A' 'a') / ('J' 'a') / ('Z' 'u') / ('L' 'a') / ('Q' 'u') / ('A' 's') / ('B' 'a'))> */
		func() bool {
			position256, tokenIndex256 := position, tokenIndex
			{
				position257 := position
				{
					position258, tokenIndex258 := position, tokenIndex
					if buffer[position] != rune('C') {
						goto l259
					}
					position++
					if buffer[position] != rune('a') {
						goto l259
					}
					position++
					goto l258
				l259:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('E') {
						goto l260
					}
					position++
					if buffer[position] != rune('a') {
						goto l260
					}
					position++
					goto l258
				l260:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('G') {
						goto l261
					}
					position++
					if buffer[position] != rune('e') {
						goto l261
					}
					position++
					goto l258
				l261:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('I') {
						goto l262
					}
					position++
					if buffer[position] != rune('a') {
						goto l262
					}
					position++
					goto l258
				l262:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('I') {
						goto l263
					}
					position++
					if buffer[position] != rune('o') {
						goto l263
					}
					position++
					goto l258
				l263:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('I') {
						goto l264
					}
					position++
					if buffer[position] != rune('x') {
						goto l264
					}
					position++
					goto l258
				l264:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('L') {
						goto l265
					}
					position++
					if buffer[position] != rune('o') {
						goto l265
					}
					position++
					goto l258
				l265:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('O') {
						goto l266
					}
					position++
					if buffer[position] != rune('a') {
						goto l266
					}
					position++
					goto l258
				l266:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('R') {
						goto l267
					}
					position++
					if buffer[position] != rune('a') {
						goto l267
					}
					position++
					goto l258
				l267:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('T') {
						goto l268
					}
					position++
					if buffer[position] != rune('y') {
						goto l268
					}
					position++
					goto l258
				l268:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('U') {
						goto l269
					}
					position++
					if buffer[position] != rune('a') {
						goto l269
					}
					position++
					goto l258
				l269:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('A') {
						goto l270
					}
					position++
					if buffer[position] != rune('a') {
						goto l270
					}
					position++
					goto l258
				l270:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('J') {
						goto l271
					}
					position++
					if buffer[position] != rune('a') {
						goto l271
					}
					position++
					goto l258
				l271:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('Z') {
						goto l272
					}
					position++
					if buffer[position] != rune('u') {
						goto l272
					}
					position++
					goto l258
				l272:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('L') {
						goto l273
					}
					position++
					if buffer[position] != rune('a') {
						goto l273
					}
					position++
					goto l258
				l273:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('Q') {
						goto l274
					}
					position++
					if buffer[position] != rune('u') {
						goto l274
					}
					position++
					goto l258
				l274:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('A') {
						goto l275
					}
					position++
					if buffer[position] != rune('s') {
						goto l275
					}
					position++
					goto l258
				l275:
					position, tokenIndex = position258, tokenIndex258
					if buffer[position] != rune('B') {
						goto l256
					}
					position++
					if buffer[position] != rune('a') {
						goto l256
					}
					position++
				}
			l258:
				add(ruleTwoLetterGenus, position257)
			}
			return true
		l256:
			position, tokenIndex = position256, tokenIndex256
			return false
		},
		/* 38 Word <- <(!(AuthorPrefix / RankUninomial / Approximation / Word4) (WordApostr / WordStartsWithDigit / Word2 / Word1) &(SpaceCharEOI / '('))> */
		func() bool {
			position276, tokenIndex276 := position, tokenIndex
			{
				position277 := position
				{
					position278, tokenIndex278 := position, tokenIndex
					{
						position279, tokenIndex279 := position, tokenIndex
						if !_rules[ruleAuthorPrefix]() {
							goto l280
						}
						goto l279
					l280:
						position, tokenIndex = position279, tokenIndex279
						if !_rules[ruleRankUninomial]() {
							goto l281
						}
						goto l279
					l281:
						position, tokenIndex = position279, tokenIndex279
						if !_rules[ruleApproximation]() {
							goto l282
						}
						goto l279
					l282:
						position, tokenIndex = position279, tokenIndex279
						if !_rules[ruleWord4]() {
							goto l278
						}
					}
				l279:
					goto l276
				l278:
					position, tokenIndex = position278, tokenIndex278
				}
				{
					position283, tokenIndex283 := position, tokenIndex
					if !_rules[ruleWordApostr]() {
						goto l284
					}
					goto l283
				l284:
					position, tokenIndex = position283, tokenIndex283
					if !_rules[ruleWordStartsWithDigit]() {
						goto l285
					}
					goto l283
				l285:
					position, tokenIndex = position283, tokenIndex283
					if !_rules[ruleWord2]() {
						goto l286
					}
					goto l283
				l286:
					position, tokenIndex = position283, tokenIndex283
					if !_rules[ruleWord1]() {
						goto l276
					}
				}
			l283:
				{
					position287, tokenIndex287 := position, tokenIndex
					{
						position288, tokenIndex288 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l289
						}
						goto l288
					l289:
						position, tokenIndex = position288, tokenIndex288
						if buffer[position] != rune('(') {
							goto l276
						}
						position++
					}
				l288:
					position, tokenIndex = position287, tokenIndex287
				}
				add(ruleWord, position277)
			}
			return true
		l276:
			position, tokenIndex = position276, tokenIndex276
			return false
		},
		/* 39 Word1 <- <((lASCII dash)? NameLowerChar NameLowerChar+)> */
		func() bool {
			position290, tokenIndex290 := position, tokenIndex
			{
				position291 := position
				{
					position292, tokenIndex292 := position, tokenIndex
					if !_rules[rulelASCII]() {
						goto l292
					}
					if !_rules[ruledash]() {
						goto l292
					}
					goto l293
				l292:
					position, tokenIndex = position292, tokenIndex292
				}
			l293:
				if !_rules[ruleNameLowerChar]() {
					goto l290
				}
				if !_rules[ruleNameLowerChar]() {
					goto l290
				}
			l294:
				{
					position295, tokenIndex295 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l295
					}
					goto l294
				l295:
					position, tokenIndex = position295, tokenIndex295
				}
				add(ruleWord1, position291)
			}
			return true
		l290:
			position, tokenIndex = position290, tokenIndex290
			return false
		},
		/* 40 WordStartsWithDigit <- <(('1' / '2' / '3' / '4' / '5' / '6' / '7' / '8' / '9') nums? ('.' / dash)? NameLowerChar NameLowerChar NameLowerChar NameLowerChar+)> */
		func() bool {
			position296, tokenIndex296 := position, tokenIndex
			{
				position297 := position
				{
					position298, tokenIndex298 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l299
					}
					position++
					goto l298
				l299:
					position, tokenIndex = position298, tokenIndex298
					if buffer[position] != rune('2') {
						goto l300
					}
					position++
					goto l298
				l300:
					position, tokenIndex = position298, tokenIndex298
					if buffer[position] != rune('3') {
						goto l301
					}
					position++
					goto l298
				l301:
					position, tokenIndex = position298, tokenIndex298
					if buffer[position] != rune('4') {
						goto l302
					}
					position++
					goto l298
				l302:
					position, tokenIndex = position298, tokenIndex298
					if buffer[position] != rune('5') {
						goto l303
					}
					position++
					goto l298
				l303:
					position, tokenIndex = position298, tokenIndex298
					if buffer[position] != rune('6') {
						goto l304
					}
					position++
					goto l298
				l304:
					position, tokenIndex = position298, tokenIndex298
					if buffer[position] != rune('7') {
						goto l305
					}
					position++
					goto l298
				l305:
					position, tokenIndex = position298, tokenIndex298
					if buffer[position] != rune('8') {
						goto l306
					}
					position++
					goto l298
				l306:
					position, tokenIndex = position298, tokenIndex298
					if buffer[position] != rune('9') {
						goto l296
					}
					position++
				}
			l298:
				{
					position307, tokenIndex307 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l307
					}
					goto l308
				l307:
					position, tokenIndex = position307, tokenIndex307
				}
			l308:
				{
					position309, tokenIndex309 := position, tokenIndex
					{
						position311, tokenIndex311 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l312
						}
						position++
						goto l311
					l312:
						position, tokenIndex = position311, tokenIndex311
						if !_rules[ruledash]() {
							goto l309
						}
					}
				l311:
					goto l310
				l309:
					position, tokenIndex = position309, tokenIndex309
				}
			l310:
				if !_rules[ruleNameLowerChar]() {
					goto l296
				}
				if !_rules[ruleNameLowerChar]() {
					goto l296
				}
				if !_rules[ruleNameLowerChar]() {
					goto l296
				}
				if !_rules[ruleNameLowerChar]() {
					goto l296
				}
			l313:
				{
					position314, tokenIndex314 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l314
					}
					goto l313
				l314:
					position, tokenIndex = position314, tokenIndex314
				}
				add(ruleWordStartsWithDigit, position297)
			}
			return true
		l296:
			position, tokenIndex = position296, tokenIndex296
			return false
		},
		/* 41 Word2 <- <(NameLowerChar+ dash? NameLowerChar+)> */
		func() bool {
			position315, tokenIndex315 := position, tokenIndex
			{
				position316 := position
				if !_rules[ruleNameLowerChar]() {
					goto l315
				}
			l317:
				{
					position318, tokenIndex318 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l318
					}
					goto l317
				l318:
					position, tokenIndex = position318, tokenIndex318
				}
				{
					position319, tokenIndex319 := position, tokenIndex
					if !_rules[ruledash]() {
						goto l319
					}
					goto l320
				l319:
					position, tokenIndex = position319, tokenIndex319
				}
			l320:
				if !_rules[ruleNameLowerChar]() {
					goto l315
				}
			l321:
				{
					position322, tokenIndex322 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l322
					}
					goto l321
				l322:
					position, tokenIndex = position322, tokenIndex322
				}
				add(ruleWord2, position316)
			}
			return true
		l315:
			position, tokenIndex = position315, tokenIndex315
			return false
		},
		/* 42 WordApostr <- <(NameLowerChar NameLowerChar* apostr Word1)> */
		func() bool {
			position323, tokenIndex323 := position, tokenIndex
			{
				position324 := position
				if !_rules[ruleNameLowerChar]() {
					goto l323
				}
			l325:
				{
					position326, tokenIndex326 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l326
					}
					goto l325
				l326:
					position, tokenIndex = position326, tokenIndex326
				}
				if !_rules[ruleapostr]() {
					goto l323
				}
				if !_rules[ruleWord1]() {
					goto l323
				}
				add(ruleWordApostr, position324)
			}
			return true
		l323:
			position, tokenIndex = position323, tokenIndex323
			return false
		},
		/* 43 Word4 <- <(NameLowerChar+ '.' NameLowerChar)> */
		func() bool {
			position327, tokenIndex327 := position, tokenIndex
			{
				position328 := position
				if !_rules[ruleNameLowerChar]() {
					goto l327
				}
			l329:
				{
					position330, tokenIndex330 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l330
					}
					goto l329
				l330:
					position, tokenIndex = position330, tokenIndex330
				}
				if buffer[position] != rune('.') {
					goto l327
				}
				position++
				if !_rules[ruleNameLowerChar]() {
					goto l327
				}
				add(ruleWord4, position328)
			}
			return true
		l327:
			position, tokenIndex = position327, tokenIndex327
			return false
		},
		/* 44 HybridChar <- <'×'> */
		func() bool {
			position331, tokenIndex331 := position, tokenIndex
			{
				position332 := position
				if buffer[position] != rune('×') {
					goto l331
				}
				position++
				add(ruleHybridChar, position332)
			}
			return true
		l331:
			position, tokenIndex = position331, tokenIndex331
			return false
		},
		/* 45 ApproxNameIgnored <- <.*> */
		func() bool {
			{
				position334 := position
			l335:
				{
					position336, tokenIndex336 := position, tokenIndex
					if !matchDot() {
						goto l336
					}
					goto l335
				l336:
					position, tokenIndex = position336, tokenIndex336
				}
				add(ruleApproxNameIgnored, position334)
			}
			return true
		},
		/* 46 Approximation <- <(('s' 'p' '.' _? ('n' 'r' '.')) / ('s' 'p' '.' _? ('a' 'f' 'f' '.')) / ('m' 'o' 'n' 's' 't' '.') / '?' / ((('s' 'p' 'p') / ('n' 'r') / ('s' 'p') / ('a' 'f' 'f') / ('s' 'p' 'e' 'c' 'i' 'e' 's')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position337, tokenIndex337 := position, tokenIndex
			{
				position338 := position
				{
					position339, tokenIndex339 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l340
					}
					position++
					if buffer[position] != rune('p') {
						goto l340
					}
					position++
					if buffer[position] != rune('.') {
						goto l340
					}
					position++
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
					if buffer[position] != rune('n') {
						goto l340
					}
					position++
					if buffer[position] != rune('r') {
						goto l340
					}
					position++
					if buffer[position] != rune('.') {
						goto l340
					}
					position++
					goto l339
				l340:
					position, tokenIndex = position339, tokenIndex339
					if buffer[position] != rune('s') {
						goto l343
					}
					position++
					if buffer[position] != rune('p') {
						goto l343
					}
					position++
					if buffer[position] != rune('.') {
						goto l343
					}
					position++
					{
						position344, tokenIndex344 := position, tokenIndex
						if !_rules[rule_]() {
							goto l344
						}
						goto l345
					l344:
						position, tokenIndex = position344, tokenIndex344
					}
				l345:
					if buffer[position] != rune('a') {
						goto l343
					}
					position++
					if buffer[position] != rune('f') {
						goto l343
					}
					position++
					if buffer[position] != rune('f') {
						goto l343
					}
					position++
					if buffer[position] != rune('.') {
						goto l343
					}
					position++
					goto l339
				l343:
					position, tokenIndex = position339, tokenIndex339
					if buffer[position] != rune('m') {
						goto l346
					}
					position++
					if buffer[position] != rune('o') {
						goto l346
					}
					position++
					if buffer[position] != rune('n') {
						goto l346
					}
					position++
					if buffer[position] != rune('s') {
						goto l346
					}
					position++
					if buffer[position] != rune('t') {
						goto l346
					}
					position++
					if buffer[position] != rune('.') {
						goto l346
					}
					position++
					goto l339
				l346:
					position, tokenIndex = position339, tokenIndex339
					if buffer[position] != rune('?') {
						goto l347
					}
					position++
					goto l339
				l347:
					position, tokenIndex = position339, tokenIndex339
					{
						position348, tokenIndex348 := position, tokenIndex
						if buffer[position] != rune('s') {
							goto l349
						}
						position++
						if buffer[position] != rune('p') {
							goto l349
						}
						position++
						if buffer[position] != rune('p') {
							goto l349
						}
						position++
						goto l348
					l349:
						position, tokenIndex = position348, tokenIndex348
						if buffer[position] != rune('n') {
							goto l350
						}
						position++
						if buffer[position] != rune('r') {
							goto l350
						}
						position++
						goto l348
					l350:
						position, tokenIndex = position348, tokenIndex348
						if buffer[position] != rune('s') {
							goto l351
						}
						position++
						if buffer[position] != rune('p') {
							goto l351
						}
						position++
						goto l348
					l351:
						position, tokenIndex = position348, tokenIndex348
						if buffer[position] != rune('a') {
							goto l352
						}
						position++
						if buffer[position] != rune('f') {
							goto l352
						}
						position++
						if buffer[position] != rune('f') {
							goto l352
						}
						position++
						goto l348
					l352:
						position, tokenIndex = position348, tokenIndex348
						if buffer[position] != rune('s') {
							goto l337
						}
						position++
						if buffer[position] != rune('p') {
							goto l337
						}
						position++
						if buffer[position] != rune('e') {
							goto l337
						}
						position++
						if buffer[position] != rune('c') {
							goto l337
						}
						position++
						if buffer[position] != rune('i') {
							goto l337
						}
						position++
						if buffer[position] != rune('e') {
							goto l337
						}
						position++
						if buffer[position] != rune('s') {
							goto l337
						}
						position++
					}
				l348:
					{
						position353, tokenIndex353 := position, tokenIndex
						{
							position355, tokenIndex355 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l354
							}
							position, tokenIndex = position355, tokenIndex355
						}
						goto l353
					l354:
						position, tokenIndex = position353, tokenIndex353
						if buffer[position] != rune('.') {
							goto l337
						}
						position++
					}
				l353:
				}
			l339:
				add(ruleApproximation, position338)
			}
			return true
		l337:
			position, tokenIndex = position337, tokenIndex337
			return false
		},
		/* 47 Authorship <- <((AuthorshipCombo / OriginalAuthorship) &(SpaceCharEOI / ','))> */
		func() bool {
			position356, tokenIndex356 := position, tokenIndex
			{
				position357 := position
				{
					position358, tokenIndex358 := position, tokenIndex
					if !_rules[ruleAuthorshipCombo]() {
						goto l359
					}
					goto l358
				l359:
					position, tokenIndex = position358, tokenIndex358
					if !_rules[ruleOriginalAuthorship]() {
						goto l356
					}
				}
			l358:
				{
					position360, tokenIndex360 := position, tokenIndex
					{
						position361, tokenIndex361 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l362
						}
						goto l361
					l362:
						position, tokenIndex = position361, tokenIndex361
						if buffer[position] != rune(',') {
							goto l356
						}
						position++
					}
				l361:
					position, tokenIndex = position360, tokenIndex360
				}
				add(ruleAuthorship, position357)
			}
			return true
		l356:
			position, tokenIndex = position356, tokenIndex356
			return false
		},
		/* 48 AuthorshipCombo <- <(OriginalAuthorship _? CombinationAuthorship)> */
		func() bool {
			position363, tokenIndex363 := position, tokenIndex
			{
				position364 := position
				if !_rules[ruleOriginalAuthorship]() {
					goto l363
				}
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
				if !_rules[ruleCombinationAuthorship]() {
					goto l363
				}
				add(ruleAuthorshipCombo, position364)
			}
			return true
		l363:
			position, tokenIndex = position363, tokenIndex363
			return false
		},
		/* 49 OriginalAuthorship <- <(BasionymAuthorshipYearMisformed / AuthorsGroup / BasionymAuthorship)> */
		func() bool {
			position367, tokenIndex367 := position, tokenIndex
			{
				position368 := position
				{
					position369, tokenIndex369 := position, tokenIndex
					if !_rules[ruleBasionymAuthorshipYearMisformed]() {
						goto l370
					}
					goto l369
				l370:
					position, tokenIndex = position369, tokenIndex369
					if !_rules[ruleAuthorsGroup]() {
						goto l371
					}
					goto l369
				l371:
					position, tokenIndex = position369, tokenIndex369
					if !_rules[ruleBasionymAuthorship]() {
						goto l367
					}
				}
			l369:
				add(ruleOriginalAuthorship, position368)
			}
			return true
		l367:
			position, tokenIndex = position367, tokenIndex367
			return false
		},
		/* 50 CombinationAuthorship <- <AuthorsGroup> */
		func() bool {
			position372, tokenIndex372 := position, tokenIndex
			{
				position373 := position
				if !_rules[ruleAuthorsGroup]() {
					goto l372
				}
				add(ruleCombinationAuthorship, position373)
			}
			return true
		l372:
			position, tokenIndex = position372, tokenIndex372
			return false
		},
		/* 51 BasionymAuthorshipYearMisformed <- <('(' _? AuthorsGroup _? ')' (_? ',')? _? Year)> */
		func() bool {
			position374, tokenIndex374 := position, tokenIndex
			{
				position375 := position
				if buffer[position] != rune('(') {
					goto l374
				}
				position++
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
				if !_rules[ruleAuthorsGroup]() {
					goto l374
				}
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
				if buffer[position] != rune(')') {
					goto l374
				}
				position++
				{
					position380, tokenIndex380 := position, tokenIndex
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
					if buffer[position] != rune(',') {
						goto l380
					}
					position++
					goto l381
				l380:
					position, tokenIndex = position380, tokenIndex380
				}
			l381:
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
				if !_rules[ruleYear]() {
					goto l374
				}
				add(ruleBasionymAuthorshipYearMisformed, position375)
			}
			return true
		l374:
			position, tokenIndex = position374, tokenIndex374
			return false
		},
		/* 52 BasionymAuthorship <- <(BasionymAuthorship1 / BasionymAuthorship2)> */
		func() bool {
			position386, tokenIndex386 := position, tokenIndex
			{
				position387 := position
				{
					position388, tokenIndex388 := position, tokenIndex
					if !_rules[ruleBasionymAuthorship1]() {
						goto l389
					}
					goto l388
				l389:
					position, tokenIndex = position388, tokenIndex388
					if !_rules[ruleBasionymAuthorship2]() {
						goto l386
					}
				}
			l388:
				add(ruleBasionymAuthorship, position387)
			}
			return true
		l386:
			position, tokenIndex = position386, tokenIndex386
			return false
		},
		/* 53 BasionymAuthorship1 <- <('(' _? AuthorsGroup _? ')')> */
		func() bool {
			position390, tokenIndex390 := position, tokenIndex
			{
				position391 := position
				if buffer[position] != rune('(') {
					goto l390
				}
				position++
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
				if !_rules[ruleAuthorsGroup]() {
					goto l390
				}
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
				if buffer[position] != rune(')') {
					goto l390
				}
				position++
				add(ruleBasionymAuthorship1, position391)
			}
			return true
		l390:
			position, tokenIndex = position390, tokenIndex390
			return false
		},
		/* 54 BasionymAuthorship2 <- <('(' _? '(' _? AuthorsGroup _? ')' _? ')')> */
		func() bool {
			position396, tokenIndex396 := position, tokenIndex
			{
				position397 := position
				if buffer[position] != rune('(') {
					goto l396
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
				if buffer[position] != rune('(') {
					goto l396
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
				if !_rules[ruleAuthorsGroup]() {
					goto l396
				}
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
					goto l396
				}
				position++
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
					goto l396
				}
				position++
				add(ruleBasionymAuthorship2, position397)
			}
			return true
		l396:
			position, tokenIndex = position396, tokenIndex396
			return false
		},
		/* 55 AuthorsGroup <- <(AuthorsTeam (_? AuthorEmend? AuthorEx? AuthorsTeam)?)> */
		func() bool {
			position406, tokenIndex406 := position, tokenIndex
			{
				position407 := position
				if !_rules[ruleAuthorsTeam]() {
					goto l406
				}
				{
					position408, tokenIndex408 := position, tokenIndex
					{
						position410, tokenIndex410 := position, tokenIndex
						if !_rules[rule_]() {
							goto l410
						}
						goto l411
					l410:
						position, tokenIndex = position410, tokenIndex410
					}
				l411:
					{
						position412, tokenIndex412 := position, tokenIndex
						if !_rules[ruleAuthorEmend]() {
							goto l412
						}
						goto l413
					l412:
						position, tokenIndex = position412, tokenIndex412
					}
				l413:
					{
						position414, tokenIndex414 := position, tokenIndex
						if !_rules[ruleAuthorEx]() {
							goto l414
						}
						goto l415
					l414:
						position, tokenIndex = position414, tokenIndex414
					}
				l415:
					if !_rules[ruleAuthorsTeam]() {
						goto l408
					}
					goto l409
				l408:
					position, tokenIndex = position408, tokenIndex408
				}
			l409:
				add(ruleAuthorsGroup, position407)
			}
			return true
		l406:
			position, tokenIndex = position406, tokenIndex406
			return false
		},
		/* 56 AuthorsTeam <- <(Author (AuthorSep Author)* (_? ','? _? Year)?)> */
		func() bool {
			position416, tokenIndex416 := position, tokenIndex
			{
				position417 := position
				if !_rules[ruleAuthor]() {
					goto l416
				}
			l418:
				{
					position419, tokenIndex419 := position, tokenIndex
					if !_rules[ruleAuthorSep]() {
						goto l419
					}
					if !_rules[ruleAuthor]() {
						goto l419
					}
					goto l418
				l419:
					position, tokenIndex = position419, tokenIndex419
				}
				{
					position420, tokenIndex420 := position, tokenIndex
					{
						position422, tokenIndex422 := position, tokenIndex
						if !_rules[rule_]() {
							goto l422
						}
						goto l423
					l422:
						position, tokenIndex = position422, tokenIndex422
					}
				l423:
					{
						position424, tokenIndex424 := position, tokenIndex
						if buffer[position] != rune(',') {
							goto l424
						}
						position++
						goto l425
					l424:
						position, tokenIndex = position424, tokenIndex424
					}
				l425:
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
					if !_rules[ruleYear]() {
						goto l420
					}
					goto l421
				l420:
					position, tokenIndex = position420, tokenIndex420
				}
			l421:
				add(ruleAuthorsTeam, position417)
			}
			return true
		l416:
			position, tokenIndex = position416, tokenIndex416
			return false
		},
		/* 57 AuthorSep <- <(AuthorSep1 / AuthorSep2)> */
		func() bool {
			position428, tokenIndex428 := position, tokenIndex
			{
				position429 := position
				{
					position430, tokenIndex430 := position, tokenIndex
					if !_rules[ruleAuthorSep1]() {
						goto l431
					}
					goto l430
				l431:
					position, tokenIndex = position430, tokenIndex430
					if !_rules[ruleAuthorSep2]() {
						goto l428
					}
				}
			l430:
				add(ruleAuthorSep, position429)
			}
			return true
		l428:
			position, tokenIndex = position428, tokenIndex428
			return false
		},
		/* 58 AuthorSep1 <- <(_? (',' _)? ('&' / ('e' 't') / ('a' 'n' 'd') / ('a' 'p' 'u' 'd')) _?)> */
		func() bool {
			position432, tokenIndex432 := position, tokenIndex
			{
				position433 := position
				{
					position434, tokenIndex434 := position, tokenIndex
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
					if buffer[position] != rune(',') {
						goto l436
					}
					position++
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
					if buffer[position] != rune('&') {
						goto l439
					}
					position++
					goto l438
				l439:
					position, tokenIndex = position438, tokenIndex438
					if buffer[position] != rune('e') {
						goto l440
					}
					position++
					if buffer[position] != rune('t') {
						goto l440
					}
					position++
					goto l438
				l440:
					position, tokenIndex = position438, tokenIndex438
					if buffer[position] != rune('a') {
						goto l441
					}
					position++
					if buffer[position] != rune('n') {
						goto l441
					}
					position++
					if buffer[position] != rune('d') {
						goto l441
					}
					position++
					goto l438
				l441:
					position, tokenIndex = position438, tokenIndex438
					if buffer[position] != rune('a') {
						goto l432
					}
					position++
					if buffer[position] != rune('p') {
						goto l432
					}
					position++
					if buffer[position] != rune('u') {
						goto l432
					}
					position++
					if buffer[position] != rune('d') {
						goto l432
					}
					position++
				}
			l438:
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
				add(ruleAuthorSep1, position433)
			}
			return true
		l432:
			position, tokenIndex = position432, tokenIndex432
			return false
		},
		/* 59 AuthorSep2 <- <(_? ',' _?)> */
		func() bool {
			position444, tokenIndex444 := position, tokenIndex
			{
				position445 := position
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
				if buffer[position] != rune(',') {
					goto l444
				}
				position++
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
				add(ruleAuthorSep2, position445)
			}
			return true
		l444:
			position, tokenIndex = position444, tokenIndex444
			return false
		},
		/* 60 AuthorEx <- <((('e' 'x' '.'?) / ('i' 'n')) _)> */
		func() bool {
			position450, tokenIndex450 := position, tokenIndex
			{
				position451 := position
				{
					position452, tokenIndex452 := position, tokenIndex
					if buffer[position] != rune('e') {
						goto l453
					}
					position++
					if buffer[position] != rune('x') {
						goto l453
					}
					position++
					{
						position454, tokenIndex454 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l454
						}
						position++
						goto l455
					l454:
						position, tokenIndex = position454, tokenIndex454
					}
				l455:
					goto l452
				l453:
					position, tokenIndex = position452, tokenIndex452
					if buffer[position] != rune('i') {
						goto l450
					}
					position++
					if buffer[position] != rune('n') {
						goto l450
					}
					position++
				}
			l452:
				if !_rules[rule_]() {
					goto l450
				}
				add(ruleAuthorEx, position451)
			}
			return true
		l450:
			position, tokenIndex = position450, tokenIndex450
			return false
		},
		/* 61 AuthorEmend <- <('e' 'm' 'e' 'n' 'd' '.'? _)> */
		func() bool {
			position456, tokenIndex456 := position, tokenIndex
			{
				position457 := position
				if buffer[position] != rune('e') {
					goto l456
				}
				position++
				if buffer[position] != rune('m') {
					goto l456
				}
				position++
				if buffer[position] != rune('e') {
					goto l456
				}
				position++
				if buffer[position] != rune('n') {
					goto l456
				}
				position++
				if buffer[position] != rune('d') {
					goto l456
				}
				position++
				{
					position458, tokenIndex458 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l458
					}
					position++
					goto l459
				l458:
					position, tokenIndex = position458, tokenIndex458
				}
			l459:
				if !_rules[rule_]() {
					goto l456
				}
				add(ruleAuthorEmend, position457)
			}
			return true
		l456:
			position, tokenIndex = position456, tokenIndex456
			return false
		},
		/* 62 Author <- <(Author1 / Author2 / UnknownAuthor)> */
		func() bool {
			position460, tokenIndex460 := position, tokenIndex
			{
				position461 := position
				{
					position462, tokenIndex462 := position, tokenIndex
					if !_rules[ruleAuthor1]() {
						goto l463
					}
					goto l462
				l463:
					position, tokenIndex = position462, tokenIndex462
					if !_rules[ruleAuthor2]() {
						goto l464
					}
					goto l462
				l464:
					position, tokenIndex = position462, tokenIndex462
					if !_rules[ruleUnknownAuthor]() {
						goto l460
					}
				}
			l462:
				add(ruleAuthor, position461)
			}
			return true
		l460:
			position, tokenIndex = position460, tokenIndex460
			return false
		},
		/* 63 Author1 <- <(Author2 _? Filius)> */
		func() bool {
			position465, tokenIndex465 := position, tokenIndex
			{
				position466 := position
				if !_rules[ruleAuthor2]() {
					goto l465
				}
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
				if !_rules[ruleFilius]() {
					goto l465
				}
				add(ruleAuthor1, position466)
			}
			return true
		l465:
			position, tokenIndex = position465, tokenIndex465
			return false
		},
		/* 64 Author2 <- <(AuthorWord (_? AuthorWord)*)> */
		func() bool {
			position469, tokenIndex469 := position, tokenIndex
			{
				position470 := position
				if !_rules[ruleAuthorWord]() {
					goto l469
				}
			l471:
				{
					position472, tokenIndex472 := position, tokenIndex
					{
						position473, tokenIndex473 := position, tokenIndex
						if !_rules[rule_]() {
							goto l473
						}
						goto l474
					l473:
						position, tokenIndex = position473, tokenIndex473
					}
				l474:
					if !_rules[ruleAuthorWord]() {
						goto l472
					}
					goto l471
				l472:
					position, tokenIndex = position472, tokenIndex472
				}
				add(ruleAuthor2, position470)
			}
			return true
		l469:
			position, tokenIndex = position469, tokenIndex469
			return false
		},
		/* 65 UnknownAuthor <- <('?' / ((('a' 'u' 'c' 't') / ('a' 'n' 'o' 'n')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position475, tokenIndex475 := position, tokenIndex
			{
				position476 := position
				{
					position477, tokenIndex477 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l478
					}
					position++
					goto l477
				l478:
					position, tokenIndex = position477, tokenIndex477
					{
						position479, tokenIndex479 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l480
						}
						position++
						if buffer[position] != rune('u') {
							goto l480
						}
						position++
						if buffer[position] != rune('c') {
							goto l480
						}
						position++
						if buffer[position] != rune('t') {
							goto l480
						}
						position++
						goto l479
					l480:
						position, tokenIndex = position479, tokenIndex479
						if buffer[position] != rune('a') {
							goto l475
						}
						position++
						if buffer[position] != rune('n') {
							goto l475
						}
						position++
						if buffer[position] != rune('o') {
							goto l475
						}
						position++
						if buffer[position] != rune('n') {
							goto l475
						}
						position++
					}
				l479:
					{
						position481, tokenIndex481 := position, tokenIndex
						{
							position483, tokenIndex483 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l482
							}
							position, tokenIndex = position483, tokenIndex483
						}
						goto l481
					l482:
						position, tokenIndex = position481, tokenIndex481
						if buffer[position] != rune('.') {
							goto l475
						}
						position++
					}
				l481:
				}
			l477:
				add(ruleUnknownAuthor, position476)
			}
			return true
		l475:
			position, tokenIndex = position475, tokenIndex475
			return false
		},
		/* 66 AuthorWord <- <(AuthorWord1 / AuthorWord2 / AuthorWord3 / AuthorPrefix)> */
		func() bool {
			position484, tokenIndex484 := position, tokenIndex
			{
				position485 := position
				{
					position486, tokenIndex486 := position, tokenIndex
					if !_rules[ruleAuthorWord1]() {
						goto l487
					}
					goto l486
				l487:
					position, tokenIndex = position486, tokenIndex486
					if !_rules[ruleAuthorWord2]() {
						goto l488
					}
					goto l486
				l488:
					position, tokenIndex = position486, tokenIndex486
					if !_rules[ruleAuthorWord3]() {
						goto l489
					}
					goto l486
				l489:
					position, tokenIndex = position486, tokenIndex486
					if !_rules[ruleAuthorPrefix]() {
						goto l484
					}
				}
			l486:
				add(ruleAuthorWord, position485)
			}
			return true
		l484:
			position, tokenIndex = position484, tokenIndex484
			return false
		},
		/* 67 AuthorWord1 <- <(('a' 'r' 'g' '.') / ('e' 't' ' ' 'a' 'l' '.' '{' '?' '}') / ((('e' 't') / '&') (' ' 'a' 'l') '.'?))> */
		func() bool {
			position490, tokenIndex490 := position, tokenIndex
			{
				position491 := position
				{
					position492, tokenIndex492 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l493
					}
					position++
					if buffer[position] != rune('r') {
						goto l493
					}
					position++
					if buffer[position] != rune('g') {
						goto l493
					}
					position++
					if buffer[position] != rune('.') {
						goto l493
					}
					position++
					goto l492
				l493:
					position, tokenIndex = position492, tokenIndex492
					if buffer[position] != rune('e') {
						goto l494
					}
					position++
					if buffer[position] != rune('t') {
						goto l494
					}
					position++
					if buffer[position] != rune(' ') {
						goto l494
					}
					position++
					if buffer[position] != rune('a') {
						goto l494
					}
					position++
					if buffer[position] != rune('l') {
						goto l494
					}
					position++
					if buffer[position] != rune('.') {
						goto l494
					}
					position++
					if buffer[position] != rune('{') {
						goto l494
					}
					position++
					if buffer[position] != rune('?') {
						goto l494
					}
					position++
					if buffer[position] != rune('}') {
						goto l494
					}
					position++
					goto l492
				l494:
					position, tokenIndex = position492, tokenIndex492
					{
						position495, tokenIndex495 := position, tokenIndex
						if buffer[position] != rune('e') {
							goto l496
						}
						position++
						if buffer[position] != rune('t') {
							goto l496
						}
						position++
						goto l495
					l496:
						position, tokenIndex = position495, tokenIndex495
						if buffer[position] != rune('&') {
							goto l490
						}
						position++
					}
				l495:
					if buffer[position] != rune(' ') {
						goto l490
					}
					position++
					if buffer[position] != rune('a') {
						goto l490
					}
					position++
					if buffer[position] != rune('l') {
						goto l490
					}
					position++
					{
						position497, tokenIndex497 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l497
						}
						position++
						goto l498
					l497:
						position, tokenIndex = position497, tokenIndex497
					}
				l498:
				}
			l492:
				add(ruleAuthorWord1, position491)
			}
			return true
		l490:
			position, tokenIndex = position490, tokenIndex490
			return false
		},
		/* 68 AuthorWord2 <- <(AuthorWord3 dash AuthorWordSoft)> */
		func() bool {
			position499, tokenIndex499 := position, tokenIndex
			{
				position500 := position
				if !_rules[ruleAuthorWord3]() {
					goto l499
				}
				if !_rules[ruledash]() {
					goto l499
				}
				if !_rules[ruleAuthorWordSoft]() {
					goto l499
				}
				add(ruleAuthorWord2, position500)
			}
			return true
		l499:
			position, tokenIndex = position499, tokenIndex499
			return false
		},
		/* 69 AuthorWord3 <- <(AuthorPrefixGlued? (AllCapsAuthorWord / CapAuthorWord) '.'?)> */
		func() bool {
			position501, tokenIndex501 := position, tokenIndex
			{
				position502 := position
				{
					position503, tokenIndex503 := position, tokenIndex
					if !_rules[ruleAuthorPrefixGlued]() {
						goto l503
					}
					goto l504
				l503:
					position, tokenIndex = position503, tokenIndex503
				}
			l504:
				{
					position505, tokenIndex505 := position, tokenIndex
					if !_rules[ruleAllCapsAuthorWord]() {
						goto l506
					}
					goto l505
				l506:
					position, tokenIndex = position505, tokenIndex505
					if !_rules[ruleCapAuthorWord]() {
						goto l501
					}
				}
			l505:
				{
					position507, tokenIndex507 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l507
					}
					position++
					goto l508
				l507:
					position, tokenIndex = position507, tokenIndex507
				}
			l508:
				add(ruleAuthorWord3, position502)
			}
			return true
		l501:
			position, tokenIndex = position501, tokenIndex501
			return false
		},
		/* 70 AuthorWordSoft <- <(((AuthorUpperChar (AuthorUpperChar+ / AuthorLowerChar+)) / AuthorLowerChar+) '.'?)> */
		func() bool {
			position509, tokenIndex509 := position, tokenIndex
			{
				position510 := position
				{
					position511, tokenIndex511 := position, tokenIndex
					if !_rules[ruleAuthorUpperChar]() {
						goto l512
					}
					{
						position513, tokenIndex513 := position, tokenIndex
						if !_rules[ruleAuthorUpperChar]() {
							goto l514
						}
					l515:
						{
							position516, tokenIndex516 := position, tokenIndex
							if !_rules[ruleAuthorUpperChar]() {
								goto l516
							}
							goto l515
						l516:
							position, tokenIndex = position516, tokenIndex516
						}
						goto l513
					l514:
						position, tokenIndex = position513, tokenIndex513
						if !_rules[ruleAuthorLowerChar]() {
							goto l512
						}
					l517:
						{
							position518, tokenIndex518 := position, tokenIndex
							if !_rules[ruleAuthorLowerChar]() {
								goto l518
							}
							goto l517
						l518:
							position, tokenIndex = position518, tokenIndex518
						}
					}
				l513:
					goto l511
				l512:
					position, tokenIndex = position511, tokenIndex511
					if !_rules[ruleAuthorLowerChar]() {
						goto l509
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
			l511:
				{
					position521, tokenIndex521 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l521
					}
					position++
					goto l522
				l521:
					position, tokenIndex = position521, tokenIndex521
				}
			l522:
				add(ruleAuthorWordSoft, position510)
			}
			return true
		l509:
			position, tokenIndex = position509, tokenIndex509
			return false
		},
		/* 71 CapAuthorWord <- <(AuthorUpperChar AuthorLowerChar*)> */
		func() bool {
			position523, tokenIndex523 := position, tokenIndex
			{
				position524 := position
				if !_rules[ruleAuthorUpperChar]() {
					goto l523
				}
			l525:
				{
					position526, tokenIndex526 := position, tokenIndex
					if !_rules[ruleAuthorLowerChar]() {
						goto l526
					}
					goto l525
				l526:
					position, tokenIndex = position526, tokenIndex526
				}
				add(ruleCapAuthorWord, position524)
			}
			return true
		l523:
			position, tokenIndex = position523, tokenIndex523
			return false
		},
		/* 72 AllCapsAuthorWord <- <(AuthorUpperChar AuthorUpperChar+)> */
		func() bool {
			position527, tokenIndex527 := position, tokenIndex
			{
				position528 := position
				if !_rules[ruleAuthorUpperChar]() {
					goto l527
				}
				if !_rules[ruleAuthorUpperChar]() {
					goto l527
				}
			l529:
				{
					position530, tokenIndex530 := position, tokenIndex
					if !_rules[ruleAuthorUpperChar]() {
						goto l530
					}
					goto l529
				l530:
					position, tokenIndex = position530, tokenIndex530
				}
				add(ruleAllCapsAuthorWord, position528)
			}
			return true
		l527:
			position, tokenIndex = position527, tokenIndex527
			return false
		},
		/* 73 Filius <- <(('f' '.') / ('f' 'i' 'l' '.') / ('f' 'i' 'l' 'i' 'u' 's'))> */
		func() bool {
			position531, tokenIndex531 := position, tokenIndex
			{
				position532 := position
				{
					position533, tokenIndex533 := position, tokenIndex
					if buffer[position] != rune('f') {
						goto l534
					}
					position++
					if buffer[position] != rune('.') {
						goto l534
					}
					position++
					goto l533
				l534:
					position, tokenIndex = position533, tokenIndex533
					if buffer[position] != rune('f') {
						goto l535
					}
					position++
					if buffer[position] != rune('i') {
						goto l535
					}
					position++
					if buffer[position] != rune('l') {
						goto l535
					}
					position++
					if buffer[position] != rune('.') {
						goto l535
					}
					position++
					goto l533
				l535:
					position, tokenIndex = position533, tokenIndex533
					if buffer[position] != rune('f') {
						goto l531
					}
					position++
					if buffer[position] != rune('i') {
						goto l531
					}
					position++
					if buffer[position] != rune('l') {
						goto l531
					}
					position++
					if buffer[position] != rune('i') {
						goto l531
					}
					position++
					if buffer[position] != rune('u') {
						goto l531
					}
					position++
					if buffer[position] != rune('s') {
						goto l531
					}
					position++
				}
			l533:
				add(ruleFilius, position532)
			}
			return true
		l531:
			position, tokenIndex = position531, tokenIndex531
			return false
		},
		/* 74 AuthorPrefixGlued <- <(('d' '\'') / ('O' '\''))> */
		func() bool {
			position536, tokenIndex536 := position, tokenIndex
			{
				position537 := position
				{
					position538, tokenIndex538 := position, tokenIndex
					if buffer[position] != rune('d') {
						goto l539
					}
					position++
					if buffer[position] != rune('\'') {
						goto l539
					}
					position++
					goto l538
				l539:
					position, tokenIndex = position538, tokenIndex538
					if buffer[position] != rune('O') {
						goto l536
					}
					position++
					if buffer[position] != rune('\'') {
						goto l536
					}
					position++
				}
			l538:
				add(ruleAuthorPrefixGlued, position537)
			}
			return true
		l536:
			position, tokenIndex = position536, tokenIndex536
			return false
		},
		/* 75 AuthorPrefix <- <(AuthorPrefix1 / AuthorPrefix2)> */
		func() bool {
			position540, tokenIndex540 := position, tokenIndex
			{
				position541 := position
				{
					position542, tokenIndex542 := position, tokenIndex
					if !_rules[ruleAuthorPrefix1]() {
						goto l543
					}
					goto l542
				l543:
					position, tokenIndex = position542, tokenIndex542
					if !_rules[ruleAuthorPrefix2]() {
						goto l540
					}
				}
			l542:
				add(ruleAuthorPrefix, position541)
			}
			return true
		l540:
			position, tokenIndex = position540, tokenIndex540
			return false
		},
		/* 76 AuthorPrefix2 <- <(('v' '.' (_? ('d' '.'))?) / ('\'' 't'))> */
		func() bool {
			position544, tokenIndex544 := position, tokenIndex
			{
				position545 := position
				{
					position546, tokenIndex546 := position, tokenIndex
					if buffer[position] != rune('v') {
						goto l547
					}
					position++
					if buffer[position] != rune('.') {
						goto l547
					}
					position++
					{
						position548, tokenIndex548 := position, tokenIndex
						{
							position550, tokenIndex550 := position, tokenIndex
							if !_rules[rule_]() {
								goto l550
							}
							goto l551
						l550:
							position, tokenIndex = position550, tokenIndex550
						}
					l551:
						if buffer[position] != rune('d') {
							goto l548
						}
						position++
						if buffer[position] != rune('.') {
							goto l548
						}
						position++
						goto l549
					l548:
						position, tokenIndex = position548, tokenIndex548
					}
				l549:
					goto l546
				l547:
					position, tokenIndex = position546, tokenIndex546
					if buffer[position] != rune('\'') {
						goto l544
					}
					position++
					if buffer[position] != rune('t') {
						goto l544
					}
					position++
				}
			l546:
				add(ruleAuthorPrefix2, position545)
			}
			return true
		l544:
			position, tokenIndex = position544, tokenIndex544
			return false
		},
		/* 77 AuthorPrefix1 <- <((('a' 'b') / ('a' 'f') / ('b' 'i' 's') / ('d' 'a') / ('d' 'e' 'r') / ('d' 'e' 's') / ('d' 'e' 'n') / ('d' 'e' 'l') / ('d' 'e' 'l' 'l' 'a') / ('d' 'e' 'l' 'a') / ('d' 'e') / ('d' 'i') / ('d' 'u') / ('e' 'l') / ('l' 'a') / ('l' 'e') / ('t' 'e' 'r') / ('v' 'a' 'n') / ('d' '\'') / ('i' 'n' '\'' 't') / ('z' 'u' 'r') / ('v' 'o' 'n' (_ (('d' '.') / ('d' 'e' 'm')))?) / ('v' (_ 'd')?)) &_)> */
		func() bool {
			position552, tokenIndex552 := position, tokenIndex
			{
				position553 := position
				{
					position554, tokenIndex554 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l555
					}
					position++
					if buffer[position] != rune('b') {
						goto l555
					}
					position++
					goto l554
				l555:
					position, tokenIndex = position554, tokenIndex554
					if buffer[position] != rune('a') {
						goto l556
					}
					position++
					if buffer[position] != rune('f') {
						goto l556
					}
					position++
					goto l554
				l556:
					position, tokenIndex = position554, tokenIndex554
					if buffer[position] != rune('b') {
						goto l557
					}
					position++
					if buffer[position] != rune('i') {
						goto l557
					}
					position++
					if buffer[position] != rune('s') {
						goto l557
					}
					position++
					goto l554
				l557:
					position, tokenIndex = position554, tokenIndex554
					if buffer[position] != rune('d') {
						goto l558
					}
					position++
					if buffer[position] != rune('a') {
						goto l558
					}
					position++
					goto l554
				l558:
					position, tokenIndex = position554, tokenIndex554
					if buffer[position] != rune('d') {
						goto l559
					}
					position++
					if buffer[position] != rune('e') {
						goto l559
					}
					position++
					if buffer[position] != rune('r') {
						goto l559
					}
					position++
					goto l554
				l559:
					position, tokenIndex = position554, tokenIndex554
					if buffer[position] != rune('d') {
						goto l560
					}
					position++
					if buffer[position] != rune('e') {
						goto l560
					}
					position++
					if buffer[position] != rune('s') {
						goto l560
					}
					position++
					goto l554
				l560:
					position, tokenIndex = position554, tokenIndex554
					if buffer[position] != rune('d') {
						goto l561
					}
					position++
					if buffer[position] != rune('e') {
						goto l561
					}
					position++
					if buffer[position] != rune('n') {
						goto l561
					}
					position++
					goto l554
				l561:
					position, tokenIndex = position554, tokenIndex554
					if buffer[position] != rune('d') {
						goto l562
					}
					position++
					if buffer[position] != rune('e') {
						goto l562
					}
					position++
					if buffer[position] != rune('l') {
						goto l562
					}
					position++
					goto l554
				l562:
					position, tokenIndex = position554, tokenIndex554
					if buffer[position] != rune('d') {
						goto l563
					}
					position++
					if buffer[position] != rune('e') {
						goto l563
					}
					position++
					if buffer[position] != rune('l') {
						goto l563
					}
					position++
					if buffer[position] != rune('l') {
						goto l563
					}
					position++
					if buffer[position] != rune('a') {
						goto l563
					}
					position++
					goto l554
				l563:
					position, tokenIndex = position554, tokenIndex554
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
					if buffer[position] != rune('a') {
						goto l564
					}
					position++
					goto l554
				l564:
					position, tokenIndex = position554, tokenIndex554
					if buffer[position] != rune('d') {
						goto l565
					}
					position++
					if buffer[position] != rune('e') {
						goto l565
					}
					position++
					goto l554
				l565:
					position, tokenIndex = position554, tokenIndex554
					if buffer[position] != rune('d') {
						goto l566
					}
					position++
					if buffer[position] != rune('i') {
						goto l566
					}
					position++
					goto l554
				l566:
					position, tokenIndex = position554, tokenIndex554
					if buffer[position] != rune('d') {
						goto l567
					}
					position++
					if buffer[position] != rune('u') {
						goto l567
					}
					position++
					goto l554
				l567:
					position, tokenIndex = position554, tokenIndex554
					if buffer[position] != rune('e') {
						goto l568
					}
					position++
					if buffer[position] != rune('l') {
						goto l568
					}
					position++
					goto l554
				l568:
					position, tokenIndex = position554, tokenIndex554
					if buffer[position] != rune('l') {
						goto l569
					}
					position++
					if buffer[position] != rune('a') {
						goto l569
					}
					position++
					goto l554
				l569:
					position, tokenIndex = position554, tokenIndex554
					if buffer[position] != rune('l') {
						goto l570
					}
					position++
					if buffer[position] != rune('e') {
						goto l570
					}
					position++
					goto l554
				l570:
					position, tokenIndex = position554, tokenIndex554
					if buffer[position] != rune('t') {
						goto l571
					}
					position++
					if buffer[position] != rune('e') {
						goto l571
					}
					position++
					if buffer[position] != rune('r') {
						goto l571
					}
					position++
					goto l554
				l571:
					position, tokenIndex = position554, tokenIndex554
					if buffer[position] != rune('v') {
						goto l572
					}
					position++
					if buffer[position] != rune('a') {
						goto l572
					}
					position++
					if buffer[position] != rune('n') {
						goto l572
					}
					position++
					goto l554
				l572:
					position, tokenIndex = position554, tokenIndex554
					if buffer[position] != rune('d') {
						goto l573
					}
					position++
					if buffer[position] != rune('\'') {
						goto l573
					}
					position++
					goto l554
				l573:
					position, tokenIndex = position554, tokenIndex554
					if buffer[position] != rune('i') {
						goto l574
					}
					position++
					if buffer[position] != rune('n') {
						goto l574
					}
					position++
					if buffer[position] != rune('\'') {
						goto l574
					}
					position++
					if buffer[position] != rune('t') {
						goto l574
					}
					position++
					goto l554
				l574:
					position, tokenIndex = position554, tokenIndex554
					if buffer[position] != rune('z') {
						goto l575
					}
					position++
					if buffer[position] != rune('u') {
						goto l575
					}
					position++
					if buffer[position] != rune('r') {
						goto l575
					}
					position++
					goto l554
				l575:
					position, tokenIndex = position554, tokenIndex554
					if buffer[position] != rune('v') {
						goto l576
					}
					position++
					if buffer[position] != rune('o') {
						goto l576
					}
					position++
					if buffer[position] != rune('n') {
						goto l576
					}
					position++
					{
						position577, tokenIndex577 := position, tokenIndex
						if !_rules[rule_]() {
							goto l577
						}
						{
							position579, tokenIndex579 := position, tokenIndex
							if buffer[position] != rune('d') {
								goto l580
							}
							position++
							if buffer[position] != rune('.') {
								goto l580
							}
							position++
							goto l579
						l580:
							position, tokenIndex = position579, tokenIndex579
							if buffer[position] != rune('d') {
								goto l577
							}
							position++
							if buffer[position] != rune('e') {
								goto l577
							}
							position++
							if buffer[position] != rune('m') {
								goto l577
							}
							position++
						}
					l579:
						goto l578
					l577:
						position, tokenIndex = position577, tokenIndex577
					}
				l578:
					goto l554
				l576:
					position, tokenIndex = position554, tokenIndex554
					if buffer[position] != rune('v') {
						goto l552
					}
					position++
					{
						position581, tokenIndex581 := position, tokenIndex
						if !_rules[rule_]() {
							goto l581
						}
						if buffer[position] != rune('d') {
							goto l581
						}
						position++
						goto l582
					l581:
						position, tokenIndex = position581, tokenIndex581
					}
				l582:
				}
			l554:
				{
					position583, tokenIndex583 := position, tokenIndex
					if !_rules[rule_]() {
						goto l552
					}
					position, tokenIndex = position583, tokenIndex583
				}
				add(ruleAuthorPrefix1, position553)
			}
			return true
		l552:
			position, tokenIndex = position552, tokenIndex552
			return false
		},
		/* 78 AuthorUpperChar <- <(hASCII / ('À' / 'Á' / 'Â' / 'Ã' / 'Ä' / 'Å' / 'Æ' / 'Ç' / 'È' / 'É' / 'Ê' / 'Ë' / 'Ì' / 'Í' / 'Î' / 'Ï' / 'Ð' / 'Ñ' / 'Ò' / 'Ó' / 'Ô' / 'Õ' / 'Ö' / 'Ø' / 'Ù' / 'Ú' / 'Û' / 'Ü' / 'Ý' / 'Ć' / 'Č' / 'Ď' / 'İ' / 'Ķ' / 'Ĺ' / 'ĺ' / 'Ľ' / 'ľ' / 'Ł' / 'ł' / 'Ņ' / 'Ō' / 'Ő' / 'Œ' / 'Ř' / 'Ś' / 'Ŝ' / 'Ş' / 'Š' / 'Ÿ' / 'Ź' / 'Ż' / 'Ž' / 'ƒ' / 'Ǿ' / 'Ș' / 'Ț'))> */
		func() bool {
			position584, tokenIndex584 := position, tokenIndex
			{
				position585 := position
				{
					position586, tokenIndex586 := position, tokenIndex
					if !_rules[rulehASCII]() {
						goto l587
					}
					goto l586
				l587:
					position, tokenIndex = position586, tokenIndex586
					{
						position588, tokenIndex588 := position, tokenIndex
						if buffer[position] != rune('À') {
							goto l589
						}
						position++
						goto l588
					l589:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Á') {
							goto l590
						}
						position++
						goto l588
					l590:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Â') {
							goto l591
						}
						position++
						goto l588
					l591:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ã') {
							goto l592
						}
						position++
						goto l588
					l592:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ä') {
							goto l593
						}
						position++
						goto l588
					l593:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Å') {
							goto l594
						}
						position++
						goto l588
					l594:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Æ') {
							goto l595
						}
						position++
						goto l588
					l595:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ç') {
							goto l596
						}
						position++
						goto l588
					l596:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('È') {
							goto l597
						}
						position++
						goto l588
					l597:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('É') {
							goto l598
						}
						position++
						goto l588
					l598:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ê') {
							goto l599
						}
						position++
						goto l588
					l599:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ë') {
							goto l600
						}
						position++
						goto l588
					l600:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ì') {
							goto l601
						}
						position++
						goto l588
					l601:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Í') {
							goto l602
						}
						position++
						goto l588
					l602:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Î') {
							goto l603
						}
						position++
						goto l588
					l603:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ï') {
							goto l604
						}
						position++
						goto l588
					l604:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ð') {
							goto l605
						}
						position++
						goto l588
					l605:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ñ') {
							goto l606
						}
						position++
						goto l588
					l606:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ò') {
							goto l607
						}
						position++
						goto l588
					l607:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ó') {
							goto l608
						}
						position++
						goto l588
					l608:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ô') {
							goto l609
						}
						position++
						goto l588
					l609:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Õ') {
							goto l610
						}
						position++
						goto l588
					l610:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ö') {
							goto l611
						}
						position++
						goto l588
					l611:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ø') {
							goto l612
						}
						position++
						goto l588
					l612:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ù') {
							goto l613
						}
						position++
						goto l588
					l613:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ú') {
							goto l614
						}
						position++
						goto l588
					l614:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Û') {
							goto l615
						}
						position++
						goto l588
					l615:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ü') {
							goto l616
						}
						position++
						goto l588
					l616:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ý') {
							goto l617
						}
						position++
						goto l588
					l617:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ć') {
							goto l618
						}
						position++
						goto l588
					l618:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Č') {
							goto l619
						}
						position++
						goto l588
					l619:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ď') {
							goto l620
						}
						position++
						goto l588
					l620:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('İ') {
							goto l621
						}
						position++
						goto l588
					l621:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ķ') {
							goto l622
						}
						position++
						goto l588
					l622:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ĺ') {
							goto l623
						}
						position++
						goto l588
					l623:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('ĺ') {
							goto l624
						}
						position++
						goto l588
					l624:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ľ') {
							goto l625
						}
						position++
						goto l588
					l625:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('ľ') {
							goto l626
						}
						position++
						goto l588
					l626:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ł') {
							goto l627
						}
						position++
						goto l588
					l627:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('ł') {
							goto l628
						}
						position++
						goto l588
					l628:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ņ') {
							goto l629
						}
						position++
						goto l588
					l629:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ō') {
							goto l630
						}
						position++
						goto l588
					l630:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ő') {
							goto l631
						}
						position++
						goto l588
					l631:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Œ') {
							goto l632
						}
						position++
						goto l588
					l632:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ř') {
							goto l633
						}
						position++
						goto l588
					l633:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ś') {
							goto l634
						}
						position++
						goto l588
					l634:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ŝ') {
							goto l635
						}
						position++
						goto l588
					l635:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ş') {
							goto l636
						}
						position++
						goto l588
					l636:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Š') {
							goto l637
						}
						position++
						goto l588
					l637:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ÿ') {
							goto l638
						}
						position++
						goto l588
					l638:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ź') {
							goto l639
						}
						position++
						goto l588
					l639:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ż') {
							goto l640
						}
						position++
						goto l588
					l640:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ž') {
							goto l641
						}
						position++
						goto l588
					l641:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('ƒ') {
							goto l642
						}
						position++
						goto l588
					l642:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ǿ') {
							goto l643
						}
						position++
						goto l588
					l643:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ș') {
							goto l644
						}
						position++
						goto l588
					l644:
						position, tokenIndex = position588, tokenIndex588
						if buffer[position] != rune('Ț') {
							goto l584
						}
						position++
					}
				l588:
				}
			l586:
				add(ruleAuthorUpperChar, position585)
			}
			return true
		l584:
			position, tokenIndex = position584, tokenIndex584
			return false
		},
		/* 79 AuthorLowerChar <- <(lASCII / ('à' / 'á' / 'â' / 'ã' / 'ä' / 'å' / 'æ' / 'ç' / 'è' / 'é' / 'ê' / 'ë' / 'ì' / 'í' / 'î' / 'ï' / 'ð' / 'ñ' / 'ò' / 'ó' / 'ó' / 'ô' / 'õ' / 'ö' / 'ø' / 'ù' / 'ú' / 'û' / 'ü' / 'ý' / 'ÿ' / 'ā' / 'ă' / 'ą' / 'ć' / 'ĉ' / 'č' / 'ď' / 'đ' / '\'' / 'ē' / 'ĕ' / 'ė' / 'ę' / 'ě' / 'ğ' / 'ī' / 'ĭ' / 'İ' / 'ı' / 'ĺ' / 'ľ' / 'ł' / 'ń' / 'ņ' / 'ň' / 'ŏ' / 'ő' / 'œ' / 'ŕ' / 'ř' / 'ś' / 'ş' / 'š' / 'ţ' / 'ť' / 'ũ' / 'ū' / 'ŭ' / 'ů' / 'ű' / 'ź' / 'ż' / 'ž' / 'ſ' / 'ǎ' / 'ǔ' / 'ǧ' / 'ș' / 'ț' / 'ȳ' / 'ß'))> */
		func() bool {
			position645, tokenIndex645 := position, tokenIndex
			{
				position646 := position
				{
					position647, tokenIndex647 := position, tokenIndex
					if !_rules[rulelASCII]() {
						goto l648
					}
					goto l647
				l648:
					position, tokenIndex = position647, tokenIndex647
					{
						position649, tokenIndex649 := position, tokenIndex
						if buffer[position] != rune('à') {
							goto l650
						}
						position++
						goto l649
					l650:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('á') {
							goto l651
						}
						position++
						goto l649
					l651:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('â') {
							goto l652
						}
						position++
						goto l649
					l652:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ã') {
							goto l653
						}
						position++
						goto l649
					l653:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ä') {
							goto l654
						}
						position++
						goto l649
					l654:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('å') {
							goto l655
						}
						position++
						goto l649
					l655:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('æ') {
							goto l656
						}
						position++
						goto l649
					l656:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ç') {
							goto l657
						}
						position++
						goto l649
					l657:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('è') {
							goto l658
						}
						position++
						goto l649
					l658:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('é') {
							goto l659
						}
						position++
						goto l649
					l659:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ê') {
							goto l660
						}
						position++
						goto l649
					l660:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ë') {
							goto l661
						}
						position++
						goto l649
					l661:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ì') {
							goto l662
						}
						position++
						goto l649
					l662:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('í') {
							goto l663
						}
						position++
						goto l649
					l663:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('î') {
							goto l664
						}
						position++
						goto l649
					l664:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ï') {
							goto l665
						}
						position++
						goto l649
					l665:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ð') {
							goto l666
						}
						position++
						goto l649
					l666:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ñ') {
							goto l667
						}
						position++
						goto l649
					l667:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ò') {
							goto l668
						}
						position++
						goto l649
					l668:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ó') {
							goto l669
						}
						position++
						goto l649
					l669:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ó') {
							goto l670
						}
						position++
						goto l649
					l670:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ô') {
							goto l671
						}
						position++
						goto l649
					l671:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('õ') {
							goto l672
						}
						position++
						goto l649
					l672:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ö') {
							goto l673
						}
						position++
						goto l649
					l673:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ø') {
							goto l674
						}
						position++
						goto l649
					l674:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ù') {
							goto l675
						}
						position++
						goto l649
					l675:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ú') {
							goto l676
						}
						position++
						goto l649
					l676:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('û') {
							goto l677
						}
						position++
						goto l649
					l677:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ü') {
							goto l678
						}
						position++
						goto l649
					l678:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ý') {
							goto l679
						}
						position++
						goto l649
					l679:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ÿ') {
							goto l680
						}
						position++
						goto l649
					l680:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ā') {
							goto l681
						}
						position++
						goto l649
					l681:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ă') {
							goto l682
						}
						position++
						goto l649
					l682:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ą') {
							goto l683
						}
						position++
						goto l649
					l683:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ć') {
							goto l684
						}
						position++
						goto l649
					l684:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ĉ') {
							goto l685
						}
						position++
						goto l649
					l685:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('č') {
							goto l686
						}
						position++
						goto l649
					l686:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ď') {
							goto l687
						}
						position++
						goto l649
					l687:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('đ') {
							goto l688
						}
						position++
						goto l649
					l688:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('\'') {
							goto l689
						}
						position++
						goto l649
					l689:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ē') {
							goto l690
						}
						position++
						goto l649
					l690:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ĕ') {
							goto l691
						}
						position++
						goto l649
					l691:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ė') {
							goto l692
						}
						position++
						goto l649
					l692:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ę') {
							goto l693
						}
						position++
						goto l649
					l693:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ě') {
							goto l694
						}
						position++
						goto l649
					l694:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ğ') {
							goto l695
						}
						position++
						goto l649
					l695:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ī') {
							goto l696
						}
						position++
						goto l649
					l696:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ĭ') {
							goto l697
						}
						position++
						goto l649
					l697:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('İ') {
							goto l698
						}
						position++
						goto l649
					l698:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ı') {
							goto l699
						}
						position++
						goto l649
					l699:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ĺ') {
							goto l700
						}
						position++
						goto l649
					l700:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ľ') {
							goto l701
						}
						position++
						goto l649
					l701:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ł') {
							goto l702
						}
						position++
						goto l649
					l702:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ń') {
							goto l703
						}
						position++
						goto l649
					l703:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ņ') {
							goto l704
						}
						position++
						goto l649
					l704:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ň') {
							goto l705
						}
						position++
						goto l649
					l705:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ŏ') {
							goto l706
						}
						position++
						goto l649
					l706:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ő') {
							goto l707
						}
						position++
						goto l649
					l707:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('œ') {
							goto l708
						}
						position++
						goto l649
					l708:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ŕ') {
							goto l709
						}
						position++
						goto l649
					l709:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ř') {
							goto l710
						}
						position++
						goto l649
					l710:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ś') {
							goto l711
						}
						position++
						goto l649
					l711:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ş') {
							goto l712
						}
						position++
						goto l649
					l712:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('š') {
							goto l713
						}
						position++
						goto l649
					l713:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ţ') {
							goto l714
						}
						position++
						goto l649
					l714:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ť') {
							goto l715
						}
						position++
						goto l649
					l715:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ũ') {
							goto l716
						}
						position++
						goto l649
					l716:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ū') {
							goto l717
						}
						position++
						goto l649
					l717:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ŭ') {
							goto l718
						}
						position++
						goto l649
					l718:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ů') {
							goto l719
						}
						position++
						goto l649
					l719:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ű') {
							goto l720
						}
						position++
						goto l649
					l720:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ź') {
							goto l721
						}
						position++
						goto l649
					l721:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ż') {
							goto l722
						}
						position++
						goto l649
					l722:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ž') {
							goto l723
						}
						position++
						goto l649
					l723:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ſ') {
							goto l724
						}
						position++
						goto l649
					l724:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ǎ') {
							goto l725
						}
						position++
						goto l649
					l725:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ǔ') {
							goto l726
						}
						position++
						goto l649
					l726:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ǧ') {
							goto l727
						}
						position++
						goto l649
					l727:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ș') {
							goto l728
						}
						position++
						goto l649
					l728:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ț') {
							goto l729
						}
						position++
						goto l649
					l729:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ȳ') {
							goto l730
						}
						position++
						goto l649
					l730:
						position, tokenIndex = position649, tokenIndex649
						if buffer[position] != rune('ß') {
							goto l645
						}
						position++
					}
				l649:
				}
			l647:
				add(ruleAuthorLowerChar, position646)
			}
			return true
		l645:
			position, tokenIndex = position645, tokenIndex645
			return false
		},
		/* 80 Year <- <(YearRange / YearApprox / YearWithParens / YearWithPage / YearWithDot / YearWithChar / YearNum)> */
		func() bool {
			position731, tokenIndex731 := position, tokenIndex
			{
				position732 := position
				{
					position733, tokenIndex733 := position, tokenIndex
					if !_rules[ruleYearRange]() {
						goto l734
					}
					goto l733
				l734:
					position, tokenIndex = position733, tokenIndex733
					if !_rules[ruleYearApprox]() {
						goto l735
					}
					goto l733
				l735:
					position, tokenIndex = position733, tokenIndex733
					if !_rules[ruleYearWithParens]() {
						goto l736
					}
					goto l733
				l736:
					position, tokenIndex = position733, tokenIndex733
					if !_rules[ruleYearWithPage]() {
						goto l737
					}
					goto l733
				l737:
					position, tokenIndex = position733, tokenIndex733
					if !_rules[ruleYearWithDot]() {
						goto l738
					}
					goto l733
				l738:
					position, tokenIndex = position733, tokenIndex733
					if !_rules[ruleYearWithChar]() {
						goto l739
					}
					goto l733
				l739:
					position, tokenIndex = position733, tokenIndex733
					if !_rules[ruleYearNum]() {
						goto l731
					}
				}
			l733:
				add(ruleYear, position732)
			}
			return true
		l731:
			position, tokenIndex = position731, tokenIndex731
			return false
		},
		/* 81 YearRange <- <(YearNum dash (nums+ ('a' / 'b' / 'c' / 'd' / 'e' / 'f' / 'g' / 'h' / 'i' / 'j' / 'k' / 'l' / 'm' / 'n' / 'o' / 'p' / 'q' / 'r' / 's' / 't' / 'u' / 'v' / 'w' / 'x' / 'y' / 'z' / '?')*))> */
		func() bool {
			position740, tokenIndex740 := position, tokenIndex
			{
				position741 := position
				if !_rules[ruleYearNum]() {
					goto l740
				}
				if !_rules[ruledash]() {
					goto l740
				}
				if !_rules[rulenums]() {
					goto l740
				}
			l742:
				{
					position743, tokenIndex743 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l743
					}
					goto l742
				l743:
					position, tokenIndex = position743, tokenIndex743
				}
			l744:
				{
					position745, tokenIndex745 := position, tokenIndex
					{
						position746, tokenIndex746 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l747
						}
						position++
						goto l746
					l747:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('b') {
							goto l748
						}
						position++
						goto l746
					l748:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('c') {
							goto l749
						}
						position++
						goto l746
					l749:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('d') {
							goto l750
						}
						position++
						goto l746
					l750:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('e') {
							goto l751
						}
						position++
						goto l746
					l751:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('f') {
							goto l752
						}
						position++
						goto l746
					l752:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('g') {
							goto l753
						}
						position++
						goto l746
					l753:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('h') {
							goto l754
						}
						position++
						goto l746
					l754:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('i') {
							goto l755
						}
						position++
						goto l746
					l755:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('j') {
							goto l756
						}
						position++
						goto l746
					l756:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('k') {
							goto l757
						}
						position++
						goto l746
					l757:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('l') {
							goto l758
						}
						position++
						goto l746
					l758:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('m') {
							goto l759
						}
						position++
						goto l746
					l759:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('n') {
							goto l760
						}
						position++
						goto l746
					l760:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('o') {
							goto l761
						}
						position++
						goto l746
					l761:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('p') {
							goto l762
						}
						position++
						goto l746
					l762:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('q') {
							goto l763
						}
						position++
						goto l746
					l763:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('r') {
							goto l764
						}
						position++
						goto l746
					l764:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('s') {
							goto l765
						}
						position++
						goto l746
					l765:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('t') {
							goto l766
						}
						position++
						goto l746
					l766:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('u') {
							goto l767
						}
						position++
						goto l746
					l767:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('v') {
							goto l768
						}
						position++
						goto l746
					l768:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('w') {
							goto l769
						}
						position++
						goto l746
					l769:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('x') {
							goto l770
						}
						position++
						goto l746
					l770:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('y') {
							goto l771
						}
						position++
						goto l746
					l771:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('z') {
							goto l772
						}
						position++
						goto l746
					l772:
						position, tokenIndex = position746, tokenIndex746
						if buffer[position] != rune('?') {
							goto l745
						}
						position++
					}
				l746:
					goto l744
				l745:
					position, tokenIndex = position745, tokenIndex745
				}
				add(ruleYearRange, position741)
			}
			return true
		l740:
			position, tokenIndex = position740, tokenIndex740
			return false
		},
		/* 82 YearWithDot <- <(YearNum '.')> */
		func() bool {
			position773, tokenIndex773 := position, tokenIndex
			{
				position774 := position
				if !_rules[ruleYearNum]() {
					goto l773
				}
				if buffer[position] != rune('.') {
					goto l773
				}
				position++
				add(ruleYearWithDot, position774)
			}
			return true
		l773:
			position, tokenIndex = position773, tokenIndex773
			return false
		},
		/* 83 YearApprox <- <('[' _? YearNum _? ']')> */
		func() bool {
			position775, tokenIndex775 := position, tokenIndex
			{
				position776 := position
				if buffer[position] != rune('[') {
					goto l775
				}
				position++
				{
					position777, tokenIndex777 := position, tokenIndex
					if !_rules[rule_]() {
						goto l777
					}
					goto l778
				l777:
					position, tokenIndex = position777, tokenIndex777
				}
			l778:
				if !_rules[ruleYearNum]() {
					goto l775
				}
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
				if buffer[position] != rune(']') {
					goto l775
				}
				position++
				add(ruleYearApprox, position776)
			}
			return true
		l775:
			position, tokenIndex = position775, tokenIndex775
			return false
		},
		/* 84 YearWithPage <- <((YearWithChar / YearNum) _? ':' _? nums+)> */
		func() bool {
			position781, tokenIndex781 := position, tokenIndex
			{
				position782 := position
				{
					position783, tokenIndex783 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l784
					}
					goto l783
				l784:
					position, tokenIndex = position783, tokenIndex783
					if !_rules[ruleYearNum]() {
						goto l781
					}
				}
			l783:
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
				if buffer[position] != rune(':') {
					goto l781
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
				if !_rules[rulenums]() {
					goto l781
				}
			l789:
				{
					position790, tokenIndex790 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l790
					}
					goto l789
				l790:
					position, tokenIndex = position790, tokenIndex790
				}
				add(ruleYearWithPage, position782)
			}
			return true
		l781:
			position, tokenIndex = position781, tokenIndex781
			return false
		},
		/* 85 YearWithParens <- <('(' (YearWithChar / YearNum) ')')> */
		func() bool {
			position791, tokenIndex791 := position, tokenIndex
			{
				position792 := position
				if buffer[position] != rune('(') {
					goto l791
				}
				position++
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
				if buffer[position] != rune(')') {
					goto l791
				}
				position++
				add(ruleYearWithParens, position792)
			}
			return true
		l791:
			position, tokenIndex = position791, tokenIndex791
			return false
		},
		/* 86 YearWithChar <- <(YearNum lASCII Action0)> */
		func() bool {
			position795, tokenIndex795 := position, tokenIndex
			{
				position796 := position
				if !_rules[ruleYearNum]() {
					goto l795
				}
				if !_rules[rulelASCII]() {
					goto l795
				}
				if !_rules[ruleAction0]() {
					goto l795
				}
				add(ruleYearWithChar, position796)
			}
			return true
		l795:
			position, tokenIndex = position795, tokenIndex795
			return false
		},
		/* 87 YearNum <- <(('1' / '2') ('0' / '7' / '8' / '9') nums (nums / '?') '?'*)> */
		func() bool {
			position797, tokenIndex797 := position, tokenIndex
			{
				position798 := position
				{
					position799, tokenIndex799 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l800
					}
					position++
					goto l799
				l800:
					position, tokenIndex = position799, tokenIndex799
					if buffer[position] != rune('2') {
						goto l797
					}
					position++
				}
			l799:
				{
					position801, tokenIndex801 := position, tokenIndex
					if buffer[position] != rune('0') {
						goto l802
					}
					position++
					goto l801
				l802:
					position, tokenIndex = position801, tokenIndex801
					if buffer[position] != rune('7') {
						goto l803
					}
					position++
					goto l801
				l803:
					position, tokenIndex = position801, tokenIndex801
					if buffer[position] != rune('8') {
						goto l804
					}
					position++
					goto l801
				l804:
					position, tokenIndex = position801, tokenIndex801
					if buffer[position] != rune('9') {
						goto l797
					}
					position++
				}
			l801:
				if !_rules[rulenums]() {
					goto l797
				}
				{
					position805, tokenIndex805 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l806
					}
					goto l805
				l806:
					position, tokenIndex = position805, tokenIndex805
					if buffer[position] != rune('?') {
						goto l797
					}
					position++
				}
			l805:
			l807:
				{
					position808, tokenIndex808 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l808
					}
					position++
					goto l807
				l808:
					position, tokenIndex = position808, tokenIndex808
				}
				add(ruleYearNum, position798)
			}
			return true
		l797:
			position, tokenIndex = position797, tokenIndex797
			return false
		},
		/* 88 NameUpperChar <- <(UpperChar / UpperCharExtended)> */
		func() bool {
			position809, tokenIndex809 := position, tokenIndex
			{
				position810 := position
				{
					position811, tokenIndex811 := position, tokenIndex
					if !_rules[ruleUpperChar]() {
						goto l812
					}
					goto l811
				l812:
					position, tokenIndex = position811, tokenIndex811
					if !_rules[ruleUpperCharExtended]() {
						goto l809
					}
				}
			l811:
				add(ruleNameUpperChar, position810)
			}
			return true
		l809:
			position, tokenIndex = position809, tokenIndex809
			return false
		},
		/* 89 UpperCharExtended <- <('Æ' / 'Œ' / 'Ö')> */
		func() bool {
			position813, tokenIndex813 := position, tokenIndex
			{
				position814 := position
				{
					position815, tokenIndex815 := position, tokenIndex
					if buffer[position] != rune('Æ') {
						goto l816
					}
					position++
					goto l815
				l816:
					position, tokenIndex = position815, tokenIndex815
					if buffer[position] != rune('Œ') {
						goto l817
					}
					position++
					goto l815
				l817:
					position, tokenIndex = position815, tokenIndex815
					if buffer[position] != rune('Ö') {
						goto l813
					}
					position++
				}
			l815:
				add(ruleUpperCharExtended, position814)
			}
			return true
		l813:
			position, tokenIndex = position813, tokenIndex813
			return false
		},
		/* 90 UpperChar <- <hASCII> */
		func() bool {
			position818, tokenIndex818 := position, tokenIndex
			{
				position819 := position
				if !_rules[rulehASCII]() {
					goto l818
				}
				add(ruleUpperChar, position819)
			}
			return true
		l818:
			position, tokenIndex = position818, tokenIndex818
			return false
		},
		/* 91 NameLowerChar <- <(LowerChar / LowerCharExtended / MiscodedChar)> */
		func() bool {
			position820, tokenIndex820 := position, tokenIndex
			{
				position821 := position
				{
					position822, tokenIndex822 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l823
					}
					goto l822
				l823:
					position, tokenIndex = position822, tokenIndex822
					if !_rules[ruleLowerCharExtended]() {
						goto l824
					}
					goto l822
				l824:
					position, tokenIndex = position822, tokenIndex822
					if !_rules[ruleMiscodedChar]() {
						goto l820
					}
				}
			l822:
				add(ruleNameLowerChar, position821)
			}
			return true
		l820:
			position, tokenIndex = position820, tokenIndex820
			return false
		},
		/* 92 MiscodedChar <- <'�'> */
		func() bool {
			position825, tokenIndex825 := position, tokenIndex
			{
				position826 := position
				if buffer[position] != rune('�') {
					goto l825
				}
				position++
				add(ruleMiscodedChar, position826)
			}
			return true
		l825:
			position, tokenIndex = position825, tokenIndex825
			return false
		},
		/* 93 LowerCharExtended <- <('æ' / 'œ' / 'à' / 'â' / 'å' / 'ã' / 'ä' / 'á' / 'ç' / 'č' / 'é' / 'è' / 'ë' / 'í' / 'ì' / 'ï' / 'ň' / 'ñ' / 'ñ' / 'ó' / 'ò' / 'ô' / 'ø' / 'õ' / 'ö' / 'ú' / 'ù' / 'ü' / 'ŕ' / 'ř' / 'ŗ' / 'ſ' / 'š' / 'š' / 'ş' / 'ž')> */
		func() bool {
			position827, tokenIndex827 := position, tokenIndex
			{
				position828 := position
				{
					position829, tokenIndex829 := position, tokenIndex
					if buffer[position] != rune('æ') {
						goto l830
					}
					position++
					goto l829
				l830:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('œ') {
						goto l831
					}
					position++
					goto l829
				l831:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('à') {
						goto l832
					}
					position++
					goto l829
				l832:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('â') {
						goto l833
					}
					position++
					goto l829
				l833:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('å') {
						goto l834
					}
					position++
					goto l829
				l834:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('ã') {
						goto l835
					}
					position++
					goto l829
				l835:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('ä') {
						goto l836
					}
					position++
					goto l829
				l836:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('á') {
						goto l837
					}
					position++
					goto l829
				l837:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('ç') {
						goto l838
					}
					position++
					goto l829
				l838:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('č') {
						goto l839
					}
					position++
					goto l829
				l839:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('é') {
						goto l840
					}
					position++
					goto l829
				l840:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('è') {
						goto l841
					}
					position++
					goto l829
				l841:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('ë') {
						goto l842
					}
					position++
					goto l829
				l842:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('í') {
						goto l843
					}
					position++
					goto l829
				l843:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('ì') {
						goto l844
					}
					position++
					goto l829
				l844:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('ï') {
						goto l845
					}
					position++
					goto l829
				l845:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('ň') {
						goto l846
					}
					position++
					goto l829
				l846:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('ñ') {
						goto l847
					}
					position++
					goto l829
				l847:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('ñ') {
						goto l848
					}
					position++
					goto l829
				l848:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('ó') {
						goto l849
					}
					position++
					goto l829
				l849:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('ò') {
						goto l850
					}
					position++
					goto l829
				l850:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('ô') {
						goto l851
					}
					position++
					goto l829
				l851:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('ø') {
						goto l852
					}
					position++
					goto l829
				l852:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('õ') {
						goto l853
					}
					position++
					goto l829
				l853:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('ö') {
						goto l854
					}
					position++
					goto l829
				l854:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('ú') {
						goto l855
					}
					position++
					goto l829
				l855:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('ù') {
						goto l856
					}
					position++
					goto l829
				l856:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('ü') {
						goto l857
					}
					position++
					goto l829
				l857:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('ŕ') {
						goto l858
					}
					position++
					goto l829
				l858:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('ř') {
						goto l859
					}
					position++
					goto l829
				l859:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('ŗ') {
						goto l860
					}
					position++
					goto l829
				l860:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('ſ') {
						goto l861
					}
					position++
					goto l829
				l861:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('š') {
						goto l862
					}
					position++
					goto l829
				l862:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('š') {
						goto l863
					}
					position++
					goto l829
				l863:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('ş') {
						goto l864
					}
					position++
					goto l829
				l864:
					position, tokenIndex = position829, tokenIndex829
					if buffer[position] != rune('ž') {
						goto l827
					}
					position++
				}
			l829:
				add(ruleLowerCharExtended, position828)
			}
			return true
		l827:
			position, tokenIndex = position827, tokenIndex827
			return false
		},
		/* 94 LowerChar <- <lASCII> */
		func() bool {
			position865, tokenIndex865 := position, tokenIndex
			{
				position866 := position
				if !_rules[rulelASCII]() {
					goto l865
				}
				add(ruleLowerChar, position866)
			}
			return true
		l865:
			position, tokenIndex = position865, tokenIndex865
			return false
		},
		/* 95 SpaceCharEOI <- <(_ / !.)> */
		func() bool {
			position867, tokenIndex867 := position, tokenIndex
			{
				position868 := position
				{
					position869, tokenIndex869 := position, tokenIndex
					if !_rules[rule_]() {
						goto l870
					}
					goto l869
				l870:
					position, tokenIndex = position869, tokenIndex869
					{
						position871, tokenIndex871 := position, tokenIndex
						if !matchDot() {
							goto l871
						}
						goto l867
					l871:
						position, tokenIndex = position871, tokenIndex871
					}
				}
			l869:
				add(ruleSpaceCharEOI, position868)
			}
			return true
		l867:
			position, tokenIndex = position867, tokenIndex867
			return false
		},
		/* 96 WordBorderChar <- <(_ / ';' / '.' / ',' / ':' / '(' / ')' / ']')> */
		nil,
		/* 97 nums <- <[0-9]> */
		func() bool {
			position873, tokenIndex873 := position, tokenIndex
			{
				position874 := position
				if c := buffer[position]; c < rune('0') || c > rune('9') {
					goto l873
				}
				position++
				add(rulenums, position874)
			}
			return true
		l873:
			position, tokenIndex = position873, tokenIndex873
			return false
		},
		/* 98 lASCII <- <[a-z]> */
		func() bool {
			position875, tokenIndex875 := position, tokenIndex
			{
				position876 := position
				if c := buffer[position]; c < rune('a') || c > rune('z') {
					goto l875
				}
				position++
				add(rulelASCII, position876)
			}
			return true
		l875:
			position, tokenIndex = position875, tokenIndex875
			return false
		},
		/* 99 hASCII <- <[A-Z]> */
		func() bool {
			position877, tokenIndex877 := position, tokenIndex
			{
				position878 := position
				if c := buffer[position]; c < rune('A') || c > rune('Z') {
					goto l877
				}
				position++
				add(rulehASCII, position878)
			}
			return true
		l877:
			position, tokenIndex = position877, tokenIndex877
			return false
		},
		/* 100 apostr <- <'\''> */
		func() bool {
			position879, tokenIndex879 := position, tokenIndex
			{
				position880 := position
				if buffer[position] != rune('\'') {
					goto l879
				}
				position++
				add(ruleapostr, position880)
			}
			return true
		l879:
			position, tokenIndex = position879, tokenIndex879
			return false
		},
		/* 101 dash <- <'-'> */
		func() bool {
			position881, tokenIndex881 := position, tokenIndex
			{
				position882 := position
				if buffer[position] != rune('-') {
					goto l881
				}
				position++
				add(ruledash, position882)
			}
			return true
		l881:
			position, tokenIndex = position881, tokenIndex881
			return false
		},
		/* 102 _ <- <(MultipleSpace / SingleSpace)> */
		func() bool {
			position883, tokenIndex883 := position, tokenIndex
			{
				position884 := position
				{
					position885, tokenIndex885 := position, tokenIndex
					if !_rules[ruleMultipleSpace]() {
						goto l886
					}
					goto l885
				l886:
					position, tokenIndex = position885, tokenIndex885
					if !_rules[ruleSingleSpace]() {
						goto l883
					}
				}
			l885:
				add(rule_, position884)
			}
			return true
		l883:
			position, tokenIndex = position883, tokenIndex883
			return false
		},
		/* 103 MultipleSpace <- <(SingleSpace SingleSpace+)> */
		func() bool {
			position887, tokenIndex887 := position, tokenIndex
			{
				position888 := position
				if !_rules[ruleSingleSpace]() {
					goto l887
				}
				if !_rules[ruleSingleSpace]() {
					goto l887
				}
			l889:
				{
					position890, tokenIndex890 := position, tokenIndex
					if !_rules[ruleSingleSpace]() {
						goto l890
					}
					goto l889
				l890:
					position, tokenIndex = position890, tokenIndex890
				}
				add(ruleMultipleSpace, position888)
			}
			return true
		l887:
			position, tokenIndex = position887, tokenIndex887
			return false
		},
		/* 104 SingleSpace <- <' '> */
		func() bool {
			position891, tokenIndex891 := position, tokenIndex
			{
				position892 := position
				if buffer[position] != rune(' ') {
					goto l891
				}
				position++
				add(ruleSingleSpace, position892)
			}
			return true
		l891:
			position, tokenIndex = position891, tokenIndex891
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
