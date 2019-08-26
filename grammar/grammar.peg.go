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
	ruleRankNotho
	ruleRankOtherUncommon
	ruleRankOther
	ruleRankVar
	ruleRankForma
	ruleRankSsp
	ruleRankAgamo
	ruleSubGenusOrSuperspecies
	ruleSubGenus
	ruleUninomialCombo
	ruleUninomialCombo1
	ruleUninomialCombo2
	ruleRankUninomial
	ruleRankUninomialPlain
	ruleRankUninomialNotho
	ruleUninomial
	ruleUninomialWord
	ruleAbbrGenus
	ruleCapWord
	ruleCapWord1
	ruleCapWordWithDash
	ruleUpperAfterDash
	ruleLowerAfterDash
	ruleTwoLetterGenus
	ruleWord
	ruleWord1
	ruleWordStartsWithDigit
	ruleWord2
	ruleWordApostr
	ruleWord4
	ruleMultiDashedWord
	ruleHybridChar
	ruleApproxNameIgnored
	ruleApproximation
	ruleAuthorship
	ruleAuthorshipCombo
	ruleOriginalAuthorship
	ruleOriginalAuthorshipComb
	ruleCombinationAuthorship
	ruleBasionymAuthorshipMissingParens
	ruleMissingParensStart
	ruleMissingParensEnd
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
	ruleNums
	ruleLowerGreek
	ruleLowerASCII
	ruleUpperASCII
	ruleApostrophe
	ruleApostrASCII
	ruleApostrOther
	ruleDash
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
	"RankNotho",
	"RankOtherUncommon",
	"RankOther",
	"RankVar",
	"RankForma",
	"RankSsp",
	"RankAgamo",
	"SubGenusOrSuperspecies",
	"SubGenus",
	"UninomialCombo",
	"UninomialCombo1",
	"UninomialCombo2",
	"RankUninomial",
	"RankUninomialPlain",
	"RankUninomialNotho",
	"Uninomial",
	"UninomialWord",
	"AbbrGenus",
	"CapWord",
	"CapWord1",
	"CapWordWithDash",
	"UpperAfterDash",
	"LowerAfterDash",
	"TwoLetterGenus",
	"Word",
	"Word1",
	"WordStartsWithDigit",
	"Word2",
	"WordApostr",
	"Word4",
	"MultiDashedWord",
	"HybridChar",
	"ApproxNameIgnored",
	"Approximation",
	"Authorship",
	"AuthorshipCombo",
	"OriginalAuthorship",
	"OriginalAuthorshipComb",
	"CombinationAuthorship",
	"BasionymAuthorshipMissingParens",
	"MissingParensStart",
	"MissingParensEnd",
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
	"Nums",
	"LowerGreek",
	"LowerASCII",
	"UpperASCII",
	"Apostrophe",
	"ApostrASCII",
	"ApostrOther",
	"Dash",
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
	rules  [121]func() bool
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
		/* 7 NamedSpeciesHybrid <- <(GenusWord (_ SubGenus)? (_ Comparison)? _ HybridChar _? SpeciesEpithet (_ InfraspGroup)?)> */
		func() bool {
			position39, tokenIndex39 := position, tokenIndex
			{
				position40 := position
				if !_rules[ruleGenusWord]() {
					goto l39
				}
				{
					position41, tokenIndex41 := position, tokenIndex
					if !_rules[rule_]() {
						goto l41
					}
					if !_rules[ruleSubGenus]() {
						goto l41
					}
					goto l42
				l41:
					position, tokenIndex = position41, tokenIndex41
				}
			l42:
				{
					position43, tokenIndex43 := position, tokenIndex
					if !_rules[rule_]() {
						goto l43
					}
					if !_rules[ruleComparison]() {
						goto l43
					}
					goto l44
				l43:
					position, tokenIndex = position43, tokenIndex43
				}
			l44:
				if !_rules[rule_]() {
					goto l39
				}
				if !_rules[ruleHybridChar]() {
					goto l39
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
				if !_rules[ruleSpeciesEpithet]() {
					goto l39
				}
				{
					position47, tokenIndex47 := position, tokenIndex
					if !_rules[rule_]() {
						goto l47
					}
					if !_rules[ruleInfraspGroup]() {
						goto l47
					}
					goto l48
				l47:
					position, tokenIndex = position47, tokenIndex47
				}
			l48:
				add(ruleNamedSpeciesHybrid, position40)
			}
			return true
		l39:
			position, tokenIndex = position39, tokenIndex39
			return false
		},
		/* 8 NamedGenusHybrid <- <(HybridChar _? SingleName)> */
		func() bool {
			position49, tokenIndex49 := position, tokenIndex
			{
				position50 := position
				if !_rules[ruleHybridChar]() {
					goto l49
				}
				{
					position51, tokenIndex51 := position, tokenIndex
					if !_rules[rule_]() {
						goto l51
					}
					goto l52
				l51:
					position, tokenIndex = position51, tokenIndex51
				}
			l52:
				if !_rules[ruleSingleName]() {
					goto l49
				}
				add(ruleNamedGenusHybrid, position50)
			}
			return true
		l49:
			position, tokenIndex = position49, tokenIndex49
			return false
		},
		/* 9 SingleName <- <(NameComp / NameApprox / NameSpecies / NameUninomial)> */
		func() bool {
			position53, tokenIndex53 := position, tokenIndex
			{
				position54 := position
				{
					position55, tokenIndex55 := position, tokenIndex
					if !_rules[ruleNameComp]() {
						goto l56
					}
					goto l55
				l56:
					position, tokenIndex = position55, tokenIndex55
					if !_rules[ruleNameApprox]() {
						goto l57
					}
					goto l55
				l57:
					position, tokenIndex = position55, tokenIndex55
					if !_rules[ruleNameSpecies]() {
						goto l58
					}
					goto l55
				l58:
					position, tokenIndex = position55, tokenIndex55
					if !_rules[ruleNameUninomial]() {
						goto l53
					}
				}
			l55:
				add(ruleSingleName, position54)
			}
			return true
		l53:
			position, tokenIndex = position53, tokenIndex53
			return false
		},
		/* 10 NameUninomial <- <(UninomialCombo / Uninomial)> */
		func() bool {
			position59, tokenIndex59 := position, tokenIndex
			{
				position60 := position
				{
					position61, tokenIndex61 := position, tokenIndex
					if !_rules[ruleUninomialCombo]() {
						goto l62
					}
					goto l61
				l62:
					position, tokenIndex = position61, tokenIndex61
					if !_rules[ruleUninomial]() {
						goto l59
					}
				}
			l61:
				add(ruleNameUninomial, position60)
			}
			return true
		l59:
			position, tokenIndex = position59, tokenIndex59
			return false
		},
		/* 11 NameApprox <- <(GenusWord (_ SpeciesEpithet)? _ Approximation ApproxNameIgnored)> */
		func() bool {
			position63, tokenIndex63 := position, tokenIndex
			{
				position64 := position
				if !_rules[ruleGenusWord]() {
					goto l63
				}
				{
					position65, tokenIndex65 := position, tokenIndex
					if !_rules[rule_]() {
						goto l65
					}
					if !_rules[ruleSpeciesEpithet]() {
						goto l65
					}
					goto l66
				l65:
					position, tokenIndex = position65, tokenIndex65
				}
			l66:
				if !_rules[rule_]() {
					goto l63
				}
				if !_rules[ruleApproximation]() {
					goto l63
				}
				if !_rules[ruleApproxNameIgnored]() {
					goto l63
				}
				add(ruleNameApprox, position64)
			}
			return true
		l63:
			position, tokenIndex = position63, tokenIndex63
			return false
		},
		/* 12 NameComp <- <(GenusWord _ Comparison (_ SpeciesEpithet)?)> */
		func() bool {
			position67, tokenIndex67 := position, tokenIndex
			{
				position68 := position
				if !_rules[ruleGenusWord]() {
					goto l67
				}
				if !_rules[rule_]() {
					goto l67
				}
				if !_rules[ruleComparison]() {
					goto l67
				}
				{
					position69, tokenIndex69 := position, tokenIndex
					if !_rules[rule_]() {
						goto l69
					}
					if !_rules[ruleSpeciesEpithet]() {
						goto l69
					}
					goto l70
				l69:
					position, tokenIndex = position69, tokenIndex69
				}
			l70:
				add(ruleNameComp, position68)
			}
			return true
		l67:
			position, tokenIndex = position67, tokenIndex67
			return false
		},
		/* 13 NameSpecies <- <(GenusWord (_? (SubGenus / SubGenusOrSuperspecies))? _ SpeciesEpithet (_ InfraspGroup)?)> */
		func() bool {
			position71, tokenIndex71 := position, tokenIndex
			{
				position72 := position
				if !_rules[ruleGenusWord]() {
					goto l71
				}
				{
					position73, tokenIndex73 := position, tokenIndex
					{
						position75, tokenIndex75 := position, tokenIndex
						if !_rules[rule_]() {
							goto l75
						}
						goto l76
					l75:
						position, tokenIndex = position75, tokenIndex75
					}
				l76:
					{
						position77, tokenIndex77 := position, tokenIndex
						if !_rules[ruleSubGenus]() {
							goto l78
						}
						goto l77
					l78:
						position, tokenIndex = position77, tokenIndex77
						if !_rules[ruleSubGenusOrSuperspecies]() {
							goto l73
						}
					}
				l77:
					goto l74
				l73:
					position, tokenIndex = position73, tokenIndex73
				}
			l74:
				if !_rules[rule_]() {
					goto l71
				}
				if !_rules[ruleSpeciesEpithet]() {
					goto l71
				}
				{
					position79, tokenIndex79 := position, tokenIndex
					if !_rules[rule_]() {
						goto l79
					}
					if !_rules[ruleInfraspGroup]() {
						goto l79
					}
					goto l80
				l79:
					position, tokenIndex = position79, tokenIndex79
				}
			l80:
				add(ruleNameSpecies, position72)
			}
			return true
		l71:
			position, tokenIndex = position71, tokenIndex71
			return false
		},
		/* 14 GenusWord <- <((AbbrGenus / UninomialWord) !(_ AuthorWord))> */
		func() bool {
			position81, tokenIndex81 := position, tokenIndex
			{
				position82 := position
				{
					position83, tokenIndex83 := position, tokenIndex
					if !_rules[ruleAbbrGenus]() {
						goto l84
					}
					goto l83
				l84:
					position, tokenIndex = position83, tokenIndex83
					if !_rules[ruleUninomialWord]() {
						goto l81
					}
				}
			l83:
				{
					position85, tokenIndex85 := position, tokenIndex
					if !_rules[rule_]() {
						goto l85
					}
					if !_rules[ruleAuthorWord]() {
						goto l85
					}
					goto l81
				l85:
					position, tokenIndex = position85, tokenIndex85
				}
				add(ruleGenusWord, position82)
			}
			return true
		l81:
			position, tokenIndex = position81, tokenIndex81
			return false
		},
		/* 15 InfraspGroup <- <(InfraspEpithet (_ InfraspEpithet)? (_ InfraspEpithet)?)> */
		func() bool {
			position86, tokenIndex86 := position, tokenIndex
			{
				position87 := position
				if !_rules[ruleInfraspEpithet]() {
					goto l86
				}
				{
					position88, tokenIndex88 := position, tokenIndex
					if !_rules[rule_]() {
						goto l88
					}
					if !_rules[ruleInfraspEpithet]() {
						goto l88
					}
					goto l89
				l88:
					position, tokenIndex = position88, tokenIndex88
				}
			l89:
				{
					position90, tokenIndex90 := position, tokenIndex
					if !_rules[rule_]() {
						goto l90
					}
					if !_rules[ruleInfraspEpithet]() {
						goto l90
					}
					goto l91
				l90:
					position, tokenIndex = position90, tokenIndex90
				}
			l91:
				add(ruleInfraspGroup, position87)
			}
			return true
		l86:
			position, tokenIndex = position86, tokenIndex86
			return false
		},
		/* 16 InfraspEpithet <- <((Rank _?)? !AuthorEx Word (_ Authorship)?)> */
		func() bool {
			position92, tokenIndex92 := position, tokenIndex
			{
				position93 := position
				{
					position94, tokenIndex94 := position, tokenIndex
					if !_rules[ruleRank]() {
						goto l94
					}
					{
						position96, tokenIndex96 := position, tokenIndex
						if !_rules[rule_]() {
							goto l96
						}
						goto l97
					l96:
						position, tokenIndex = position96, tokenIndex96
					}
				l97:
					goto l95
				l94:
					position, tokenIndex = position94, tokenIndex94
				}
			l95:
				{
					position98, tokenIndex98 := position, tokenIndex
					if !_rules[ruleAuthorEx]() {
						goto l98
					}
					goto l92
				l98:
					position, tokenIndex = position98, tokenIndex98
				}
				if !_rules[ruleWord]() {
					goto l92
				}
				{
					position99, tokenIndex99 := position, tokenIndex
					if !_rules[rule_]() {
						goto l99
					}
					if !_rules[ruleAuthorship]() {
						goto l99
					}
					goto l100
				l99:
					position, tokenIndex = position99, tokenIndex99
				}
			l100:
				add(ruleInfraspEpithet, position93)
			}
			return true
		l92:
			position, tokenIndex = position92, tokenIndex92
			return false
		},
		/* 17 SpeciesEpithet <- <(!AuthorEx Word (_? Authorship)?)> */
		func() bool {
			position101, tokenIndex101 := position, tokenIndex
			{
				position102 := position
				{
					position103, tokenIndex103 := position, tokenIndex
					if !_rules[ruleAuthorEx]() {
						goto l103
					}
					goto l101
				l103:
					position, tokenIndex = position103, tokenIndex103
				}
				if !_rules[ruleWord]() {
					goto l101
				}
				{
					position104, tokenIndex104 := position, tokenIndex
					{
						position106, tokenIndex106 := position, tokenIndex
						if !_rules[rule_]() {
							goto l106
						}
						goto l107
					l106:
						position, tokenIndex = position106, tokenIndex106
					}
				l107:
					if !_rules[ruleAuthorship]() {
						goto l104
					}
					goto l105
				l104:
					position, tokenIndex = position104, tokenIndex104
				}
			l105:
				add(ruleSpeciesEpithet, position102)
			}
			return true
		l101:
			position, tokenIndex = position101, tokenIndex101
			return false
		},
		/* 18 Comparison <- <('c' 'f' '.'?)> */
		func() bool {
			position108, tokenIndex108 := position, tokenIndex
			{
				position109 := position
				if buffer[position] != rune('c') {
					goto l108
				}
				position++
				if buffer[position] != rune('f') {
					goto l108
				}
				position++
				{
					position110, tokenIndex110 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l110
					}
					position++
					goto l111
				l110:
					position, tokenIndex = position110, tokenIndex110
				}
			l111:
				add(ruleComparison, position109)
			}
			return true
		l108:
			position, tokenIndex = position108, tokenIndex108
			return false
		},
		/* 19 Rank <- <((RankForma / RankVar / RankSsp / RankOther / RankOtherUncommon / RankAgamo / RankNotho) (_? LowerGreek ('.' / &SpaceCharEOI))?)> */
		func() bool {
			position112, tokenIndex112 := position, tokenIndex
			{
				position113 := position
				{
					position114, tokenIndex114 := position, tokenIndex
					if !_rules[ruleRankForma]() {
						goto l115
					}
					goto l114
				l115:
					position, tokenIndex = position114, tokenIndex114
					if !_rules[ruleRankVar]() {
						goto l116
					}
					goto l114
				l116:
					position, tokenIndex = position114, tokenIndex114
					if !_rules[ruleRankSsp]() {
						goto l117
					}
					goto l114
				l117:
					position, tokenIndex = position114, tokenIndex114
					if !_rules[ruleRankOther]() {
						goto l118
					}
					goto l114
				l118:
					position, tokenIndex = position114, tokenIndex114
					if !_rules[ruleRankOtherUncommon]() {
						goto l119
					}
					goto l114
				l119:
					position, tokenIndex = position114, tokenIndex114
					if !_rules[ruleRankAgamo]() {
						goto l120
					}
					goto l114
				l120:
					position, tokenIndex = position114, tokenIndex114
					if !_rules[ruleRankNotho]() {
						goto l112
					}
				}
			l114:
				{
					position121, tokenIndex121 := position, tokenIndex
					{
						position123, tokenIndex123 := position, tokenIndex
						if !_rules[rule_]() {
							goto l123
						}
						goto l124
					l123:
						position, tokenIndex = position123, tokenIndex123
					}
				l124:
					if !_rules[ruleLowerGreek]() {
						goto l121
					}
					{
						position125, tokenIndex125 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l126
						}
						position++
						goto l125
					l126:
						position, tokenIndex = position125, tokenIndex125
						{
							position127, tokenIndex127 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l121
							}
							position, tokenIndex = position127, tokenIndex127
						}
					}
				l125:
					goto l122
				l121:
					position, tokenIndex = position121, tokenIndex121
				}
			l122:
				add(ruleRank, position113)
			}
			return true
		l112:
			position, tokenIndex = position112, tokenIndex112
			return false
		},
		/* 20 RankNotho <- <((('n' 'o' 't' 'h' 'o' (('v' 'a' 'r') / ('f' 'o') / 'f' / ('s' 'u' 'b' 's' 'p') / ('s' 's' 'p') / ('s' 'p') / ('m' 'o' 'r' 't' 'h') / ('s' 'u' 'p' 's' 'p') / ('s' 'u'))) / ('n' 'v' 'a' 'r')) ('.' / &SpaceCharEOI))> */
		func() bool {
			position128, tokenIndex128 := position, tokenIndex
			{
				position129 := position
				{
					position130, tokenIndex130 := position, tokenIndex
					if buffer[position] != rune('n') {
						goto l131
					}
					position++
					if buffer[position] != rune('o') {
						goto l131
					}
					position++
					if buffer[position] != rune('t') {
						goto l131
					}
					position++
					if buffer[position] != rune('h') {
						goto l131
					}
					position++
					if buffer[position] != rune('o') {
						goto l131
					}
					position++
					{
						position132, tokenIndex132 := position, tokenIndex
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
						goto l132
					l133:
						position, tokenIndex = position132, tokenIndex132
						if buffer[position] != rune('f') {
							goto l134
						}
						position++
						if buffer[position] != rune('o') {
							goto l134
						}
						position++
						goto l132
					l134:
						position, tokenIndex = position132, tokenIndex132
						if buffer[position] != rune('f') {
							goto l135
						}
						position++
						goto l132
					l135:
						position, tokenIndex = position132, tokenIndex132
						if buffer[position] != rune('s') {
							goto l136
						}
						position++
						if buffer[position] != rune('u') {
							goto l136
						}
						position++
						if buffer[position] != rune('b') {
							goto l136
						}
						position++
						if buffer[position] != rune('s') {
							goto l136
						}
						position++
						if buffer[position] != rune('p') {
							goto l136
						}
						position++
						goto l132
					l136:
						position, tokenIndex = position132, tokenIndex132
						if buffer[position] != rune('s') {
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
						goto l132
					l137:
						position, tokenIndex = position132, tokenIndex132
						if buffer[position] != rune('s') {
							goto l138
						}
						position++
						if buffer[position] != rune('p') {
							goto l138
						}
						position++
						goto l132
					l138:
						position, tokenIndex = position132, tokenIndex132
						if buffer[position] != rune('m') {
							goto l139
						}
						position++
						if buffer[position] != rune('o') {
							goto l139
						}
						position++
						if buffer[position] != rune('r') {
							goto l139
						}
						position++
						if buffer[position] != rune('t') {
							goto l139
						}
						position++
						if buffer[position] != rune('h') {
							goto l139
						}
						position++
						goto l132
					l139:
						position, tokenIndex = position132, tokenIndex132
						if buffer[position] != rune('s') {
							goto l140
						}
						position++
						if buffer[position] != rune('u') {
							goto l140
						}
						position++
						if buffer[position] != rune('p') {
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
						goto l132
					l140:
						position, tokenIndex = position132, tokenIndex132
						if buffer[position] != rune('s') {
							goto l131
						}
						position++
						if buffer[position] != rune('u') {
							goto l131
						}
						position++
					}
				l132:
					goto l130
				l131:
					position, tokenIndex = position130, tokenIndex130
					if buffer[position] != rune('n') {
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
				}
			l130:
				{
					position141, tokenIndex141 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l142
					}
					position++
					goto l141
				l142:
					position, tokenIndex = position141, tokenIndex141
					{
						position143, tokenIndex143 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l128
						}
						position, tokenIndex = position143, tokenIndex143
					}
				}
			l141:
				add(ruleRankNotho, position129)
			}
			return true
		l128:
			position, tokenIndex = position128, tokenIndex128
			return false
		},
		/* 21 RankOtherUncommon <- <(('*' / ('n' 'a' 't' 'i' 'o') / ('n' 'a' 't' '.') / ('n' 'a' 't') / ('f' '.' 's' 'p') / 'α' / ('β' 'β') / 'β' / 'γ' / 'δ' / 'ε' / 'φ' / 'θ' / 'μ' / ('a' '.') / ('b' '.') / ('c' '.') / ('d' '.') / ('e' '.') / ('g' '.') / ('k' '.') / ('m' 'u' 't' '.')) &SpaceCharEOI)> */
		func() bool {
			position144, tokenIndex144 := position, tokenIndex
			{
				position145 := position
				{
					position146, tokenIndex146 := position, tokenIndex
					if buffer[position] != rune('*') {
						goto l147
					}
					position++
					goto l146
				l147:
					position, tokenIndex = position146, tokenIndex146
					if buffer[position] != rune('n') {
						goto l148
					}
					position++
					if buffer[position] != rune('a') {
						goto l148
					}
					position++
					if buffer[position] != rune('t') {
						goto l148
					}
					position++
					if buffer[position] != rune('i') {
						goto l148
					}
					position++
					if buffer[position] != rune('o') {
						goto l148
					}
					position++
					goto l146
				l148:
					position, tokenIndex = position146, tokenIndex146
					if buffer[position] != rune('n') {
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
					if buffer[position] != rune('.') {
						goto l149
					}
					position++
					goto l146
				l149:
					position, tokenIndex = position146, tokenIndex146
					if buffer[position] != rune('n') {
						goto l150
					}
					position++
					if buffer[position] != rune('a') {
						goto l150
					}
					position++
					if buffer[position] != rune('t') {
						goto l150
					}
					position++
					goto l146
				l150:
					position, tokenIndex = position146, tokenIndex146
					if buffer[position] != rune('f') {
						goto l151
					}
					position++
					if buffer[position] != rune('.') {
						goto l151
					}
					position++
					if buffer[position] != rune('s') {
						goto l151
					}
					position++
					if buffer[position] != rune('p') {
						goto l151
					}
					position++
					goto l146
				l151:
					position, tokenIndex = position146, tokenIndex146
					if buffer[position] != rune('α') {
						goto l152
					}
					position++
					goto l146
				l152:
					position, tokenIndex = position146, tokenIndex146
					if buffer[position] != rune('β') {
						goto l153
					}
					position++
					if buffer[position] != rune('β') {
						goto l153
					}
					position++
					goto l146
				l153:
					position, tokenIndex = position146, tokenIndex146
					if buffer[position] != rune('β') {
						goto l154
					}
					position++
					goto l146
				l154:
					position, tokenIndex = position146, tokenIndex146
					if buffer[position] != rune('γ') {
						goto l155
					}
					position++
					goto l146
				l155:
					position, tokenIndex = position146, tokenIndex146
					if buffer[position] != rune('δ') {
						goto l156
					}
					position++
					goto l146
				l156:
					position, tokenIndex = position146, tokenIndex146
					if buffer[position] != rune('ε') {
						goto l157
					}
					position++
					goto l146
				l157:
					position, tokenIndex = position146, tokenIndex146
					if buffer[position] != rune('φ') {
						goto l158
					}
					position++
					goto l146
				l158:
					position, tokenIndex = position146, tokenIndex146
					if buffer[position] != rune('θ') {
						goto l159
					}
					position++
					goto l146
				l159:
					position, tokenIndex = position146, tokenIndex146
					if buffer[position] != rune('μ') {
						goto l160
					}
					position++
					goto l146
				l160:
					position, tokenIndex = position146, tokenIndex146
					if buffer[position] != rune('a') {
						goto l161
					}
					position++
					if buffer[position] != rune('.') {
						goto l161
					}
					position++
					goto l146
				l161:
					position, tokenIndex = position146, tokenIndex146
					if buffer[position] != rune('b') {
						goto l162
					}
					position++
					if buffer[position] != rune('.') {
						goto l162
					}
					position++
					goto l146
				l162:
					position, tokenIndex = position146, tokenIndex146
					if buffer[position] != rune('c') {
						goto l163
					}
					position++
					if buffer[position] != rune('.') {
						goto l163
					}
					position++
					goto l146
				l163:
					position, tokenIndex = position146, tokenIndex146
					if buffer[position] != rune('d') {
						goto l164
					}
					position++
					if buffer[position] != rune('.') {
						goto l164
					}
					position++
					goto l146
				l164:
					position, tokenIndex = position146, tokenIndex146
					if buffer[position] != rune('e') {
						goto l165
					}
					position++
					if buffer[position] != rune('.') {
						goto l165
					}
					position++
					goto l146
				l165:
					position, tokenIndex = position146, tokenIndex146
					if buffer[position] != rune('g') {
						goto l166
					}
					position++
					if buffer[position] != rune('.') {
						goto l166
					}
					position++
					goto l146
				l166:
					position, tokenIndex = position146, tokenIndex146
					if buffer[position] != rune('k') {
						goto l167
					}
					position++
					if buffer[position] != rune('.') {
						goto l167
					}
					position++
					goto l146
				l167:
					position, tokenIndex = position146, tokenIndex146
					if buffer[position] != rune('m') {
						goto l144
					}
					position++
					if buffer[position] != rune('u') {
						goto l144
					}
					position++
					if buffer[position] != rune('t') {
						goto l144
					}
					position++
					if buffer[position] != rune('.') {
						goto l144
					}
					position++
				}
			l146:
				{
					position168, tokenIndex168 := position, tokenIndex
					if !_rules[ruleSpaceCharEOI]() {
						goto l144
					}
					position, tokenIndex = position168, tokenIndex168
				}
				add(ruleRankOtherUncommon, position145)
			}
			return true
		l144:
			position, tokenIndex = position144, tokenIndex144
			return false
		},
		/* 22 RankOther <- <((('m' 'o' 'r' 'p' 'h') / ('c' 'o' 'n' 'v' 'a' 'r') / ('p' 's' 'e' 'u' 'd' 'o' 'v' 'a' 'r') / ('s' 'e' 'c' 't') / ('s' 'e' 'r') / ('s' 'u' 'b' 'v' 'a' 'r') / ('s' 'u' 'b' 'f') / ('r' 'a' 'c' 'e') / ('p' 'v') / ('p' 'a' 't' 'h' 'o' 'v' 'a' 'r') / ('a' 'b' '.' (_? ('n' '.'))?) / ('s' 't')) ('.' / &SpaceCharEOI))> */
		func() bool {
			position169, tokenIndex169 := position, tokenIndex
			{
				position170 := position
				{
					position171, tokenIndex171 := position, tokenIndex
					if buffer[position] != rune('m') {
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
					if buffer[position] != rune('p') {
						goto l172
					}
					position++
					if buffer[position] != rune('h') {
						goto l172
					}
					position++
					goto l171
				l172:
					position, tokenIndex = position171, tokenIndex171
					if buffer[position] != rune('c') {
						goto l173
					}
					position++
					if buffer[position] != rune('o') {
						goto l173
					}
					position++
					if buffer[position] != rune('n') {
						goto l173
					}
					position++
					if buffer[position] != rune('v') {
						goto l173
					}
					position++
					if buffer[position] != rune('a') {
						goto l173
					}
					position++
					if buffer[position] != rune('r') {
						goto l173
					}
					position++
					goto l171
				l173:
					position, tokenIndex = position171, tokenIndex171
					if buffer[position] != rune('p') {
						goto l174
					}
					position++
					if buffer[position] != rune('s') {
						goto l174
					}
					position++
					if buffer[position] != rune('e') {
						goto l174
					}
					position++
					if buffer[position] != rune('u') {
						goto l174
					}
					position++
					if buffer[position] != rune('d') {
						goto l174
					}
					position++
					if buffer[position] != rune('o') {
						goto l174
					}
					position++
					if buffer[position] != rune('v') {
						goto l174
					}
					position++
					if buffer[position] != rune('a') {
						goto l174
					}
					position++
					if buffer[position] != rune('r') {
						goto l174
					}
					position++
					goto l171
				l174:
					position, tokenIndex = position171, tokenIndex171
					if buffer[position] != rune('s') {
						goto l175
					}
					position++
					if buffer[position] != rune('e') {
						goto l175
					}
					position++
					if buffer[position] != rune('c') {
						goto l175
					}
					position++
					if buffer[position] != rune('t') {
						goto l175
					}
					position++
					goto l171
				l175:
					position, tokenIndex = position171, tokenIndex171
					if buffer[position] != rune('s') {
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
					goto l171
				l176:
					position, tokenIndex = position171, tokenIndex171
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
					if buffer[position] != rune('v') {
						goto l177
					}
					position++
					if buffer[position] != rune('a') {
						goto l177
					}
					position++
					if buffer[position] != rune('r') {
						goto l177
					}
					position++
					goto l171
				l177:
					position, tokenIndex = position171, tokenIndex171
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
					if buffer[position] != rune('f') {
						goto l178
					}
					position++
					goto l171
				l178:
					position, tokenIndex = position171, tokenIndex171
					if buffer[position] != rune('r') {
						goto l179
					}
					position++
					if buffer[position] != rune('a') {
						goto l179
					}
					position++
					if buffer[position] != rune('c') {
						goto l179
					}
					position++
					if buffer[position] != rune('e') {
						goto l179
					}
					position++
					goto l171
				l179:
					position, tokenIndex = position171, tokenIndex171
					if buffer[position] != rune('p') {
						goto l180
					}
					position++
					if buffer[position] != rune('v') {
						goto l180
					}
					position++
					goto l171
				l180:
					position, tokenIndex = position171, tokenIndex171
					if buffer[position] != rune('p') {
						goto l181
					}
					position++
					if buffer[position] != rune('a') {
						goto l181
					}
					position++
					if buffer[position] != rune('t') {
						goto l181
					}
					position++
					if buffer[position] != rune('h') {
						goto l181
					}
					position++
					if buffer[position] != rune('o') {
						goto l181
					}
					position++
					if buffer[position] != rune('v') {
						goto l181
					}
					position++
					if buffer[position] != rune('a') {
						goto l181
					}
					position++
					if buffer[position] != rune('r') {
						goto l181
					}
					position++
					goto l171
				l181:
					position, tokenIndex = position171, tokenIndex171
					if buffer[position] != rune('a') {
						goto l182
					}
					position++
					if buffer[position] != rune('b') {
						goto l182
					}
					position++
					if buffer[position] != rune('.') {
						goto l182
					}
					position++
					{
						position183, tokenIndex183 := position, tokenIndex
						{
							position185, tokenIndex185 := position, tokenIndex
							if !_rules[rule_]() {
								goto l185
							}
							goto l186
						l185:
							position, tokenIndex = position185, tokenIndex185
						}
					l186:
						if buffer[position] != rune('n') {
							goto l183
						}
						position++
						if buffer[position] != rune('.') {
							goto l183
						}
						position++
						goto l184
					l183:
						position, tokenIndex = position183, tokenIndex183
					}
				l184:
					goto l171
				l182:
					position, tokenIndex = position171, tokenIndex171
					if buffer[position] != rune('s') {
						goto l169
					}
					position++
					if buffer[position] != rune('t') {
						goto l169
					}
					position++
				}
			l171:
				{
					position187, tokenIndex187 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l188
					}
					position++
					goto l187
				l188:
					position, tokenIndex = position187, tokenIndex187
					{
						position189, tokenIndex189 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l169
						}
						position, tokenIndex = position189, tokenIndex189
					}
				}
			l187:
				add(ruleRankOther, position170)
			}
			return true
		l169:
			position, tokenIndex = position169, tokenIndex169
			return false
		},
		/* 23 RankVar <- <((('v' 'a' 'r' 'i' 'e' 't' 'y') / ('[' 'v' 'a' 'r' '.' ']') / ('v' 'a' 'r')) ('.' / &SpaceCharEOI))> */
		func() bool {
			position190, tokenIndex190 := position, tokenIndex
			{
				position191 := position
				{
					position192, tokenIndex192 := position, tokenIndex
					if buffer[position] != rune('v') {
						goto l193
					}
					position++
					if buffer[position] != rune('a') {
						goto l193
					}
					position++
					if buffer[position] != rune('r') {
						goto l193
					}
					position++
					if buffer[position] != rune('i') {
						goto l193
					}
					position++
					if buffer[position] != rune('e') {
						goto l193
					}
					position++
					if buffer[position] != rune('t') {
						goto l193
					}
					position++
					if buffer[position] != rune('y') {
						goto l193
					}
					position++
					goto l192
				l193:
					position, tokenIndex = position192, tokenIndex192
					if buffer[position] != rune('[') {
						goto l194
					}
					position++
					if buffer[position] != rune('v') {
						goto l194
					}
					position++
					if buffer[position] != rune('a') {
						goto l194
					}
					position++
					if buffer[position] != rune('r') {
						goto l194
					}
					position++
					if buffer[position] != rune('.') {
						goto l194
					}
					position++
					if buffer[position] != rune(']') {
						goto l194
					}
					position++
					goto l192
				l194:
					position, tokenIndex = position192, tokenIndex192
					if buffer[position] != rune('v') {
						goto l190
					}
					position++
					if buffer[position] != rune('a') {
						goto l190
					}
					position++
					if buffer[position] != rune('r') {
						goto l190
					}
					position++
				}
			l192:
				{
					position195, tokenIndex195 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l196
					}
					position++
					goto l195
				l196:
					position, tokenIndex = position195, tokenIndex195
					{
						position197, tokenIndex197 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l190
						}
						position, tokenIndex = position197, tokenIndex197
					}
				}
			l195:
				add(ruleRankVar, position191)
			}
			return true
		l190:
			position, tokenIndex = position190, tokenIndex190
			return false
		},
		/* 24 RankForma <- <((('f' 'o' 'r' 'm' 'a') / ('f' 'm' 'a') / ('f' 'o' 'r' 'm') / ('f' 'o') / 'f') ('.' / &SpaceCharEOI))> */
		func() bool {
			position198, tokenIndex198 := position, tokenIndex
			{
				position199 := position
				{
					position200, tokenIndex200 := position, tokenIndex
					if buffer[position] != rune('f') {
						goto l201
					}
					position++
					if buffer[position] != rune('o') {
						goto l201
					}
					position++
					if buffer[position] != rune('r') {
						goto l201
					}
					position++
					if buffer[position] != rune('m') {
						goto l201
					}
					position++
					if buffer[position] != rune('a') {
						goto l201
					}
					position++
					goto l200
				l201:
					position, tokenIndex = position200, tokenIndex200
					if buffer[position] != rune('f') {
						goto l202
					}
					position++
					if buffer[position] != rune('m') {
						goto l202
					}
					position++
					if buffer[position] != rune('a') {
						goto l202
					}
					position++
					goto l200
				l202:
					position, tokenIndex = position200, tokenIndex200
					if buffer[position] != rune('f') {
						goto l203
					}
					position++
					if buffer[position] != rune('o') {
						goto l203
					}
					position++
					if buffer[position] != rune('r') {
						goto l203
					}
					position++
					if buffer[position] != rune('m') {
						goto l203
					}
					position++
					goto l200
				l203:
					position, tokenIndex = position200, tokenIndex200
					if buffer[position] != rune('f') {
						goto l204
					}
					position++
					if buffer[position] != rune('o') {
						goto l204
					}
					position++
					goto l200
				l204:
					position, tokenIndex = position200, tokenIndex200
					if buffer[position] != rune('f') {
						goto l198
					}
					position++
				}
			l200:
				{
					position205, tokenIndex205 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l206
					}
					position++
					goto l205
				l206:
					position, tokenIndex = position205, tokenIndex205
					{
						position207, tokenIndex207 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l198
						}
						position, tokenIndex = position207, tokenIndex207
					}
				}
			l205:
				add(ruleRankForma, position199)
			}
			return true
		l198:
			position, tokenIndex = position198, tokenIndex198
			return false
		},
		/* 25 RankSsp <- <((('s' 's' 'p') / ('s' 'u' 'b' 's' 'p')) ('.' / &SpaceCharEOI))> */
		func() bool {
			position208, tokenIndex208 := position, tokenIndex
			{
				position209 := position
				{
					position210, tokenIndex210 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l211
					}
					position++
					if buffer[position] != rune('s') {
						goto l211
					}
					position++
					if buffer[position] != rune('p') {
						goto l211
					}
					position++
					goto l210
				l211:
					position, tokenIndex = position210, tokenIndex210
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
					if buffer[position] != rune('p') {
						goto l208
					}
					position++
				}
			l210:
				{
					position212, tokenIndex212 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l213
					}
					position++
					goto l212
				l213:
					position, tokenIndex = position212, tokenIndex212
					{
						position214, tokenIndex214 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l208
						}
						position, tokenIndex = position214, tokenIndex214
					}
				}
			l212:
				add(ruleRankSsp, position209)
			}
			return true
		l208:
			position, tokenIndex = position208, tokenIndex208
			return false
		},
		/* 26 RankAgamo <- <((('a' 'g' 'a' 'm' 'o' 's' 'p') / ('a' 'g' 'a' 'm' 'o' 's' 's' 'p') / ('a' 'g' 'a' 'm' 'o' 'v' 'a' 'r')) ('.' / &SpaceCharEOI))> */
		func() bool {
			position215, tokenIndex215 := position, tokenIndex
			{
				position216 := position
				{
					position217, tokenIndex217 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l218
					}
					position++
					if buffer[position] != rune('g') {
						goto l218
					}
					position++
					if buffer[position] != rune('a') {
						goto l218
					}
					position++
					if buffer[position] != rune('m') {
						goto l218
					}
					position++
					if buffer[position] != rune('o') {
						goto l218
					}
					position++
					if buffer[position] != rune('s') {
						goto l218
					}
					position++
					if buffer[position] != rune('p') {
						goto l218
					}
					position++
					goto l217
				l218:
					position, tokenIndex = position217, tokenIndex217
					if buffer[position] != rune('a') {
						goto l219
					}
					position++
					if buffer[position] != rune('g') {
						goto l219
					}
					position++
					if buffer[position] != rune('a') {
						goto l219
					}
					position++
					if buffer[position] != rune('m') {
						goto l219
					}
					position++
					if buffer[position] != rune('o') {
						goto l219
					}
					position++
					if buffer[position] != rune('s') {
						goto l219
					}
					position++
					if buffer[position] != rune('s') {
						goto l219
					}
					position++
					if buffer[position] != rune('p') {
						goto l219
					}
					position++
					goto l217
				l219:
					position, tokenIndex = position217, tokenIndex217
					if buffer[position] != rune('a') {
						goto l215
					}
					position++
					if buffer[position] != rune('g') {
						goto l215
					}
					position++
					if buffer[position] != rune('a') {
						goto l215
					}
					position++
					if buffer[position] != rune('m') {
						goto l215
					}
					position++
					if buffer[position] != rune('o') {
						goto l215
					}
					position++
					if buffer[position] != rune('v') {
						goto l215
					}
					position++
					if buffer[position] != rune('a') {
						goto l215
					}
					position++
					if buffer[position] != rune('r') {
						goto l215
					}
					position++
				}
			l217:
				{
					position220, tokenIndex220 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l221
					}
					position++
					goto l220
				l221:
					position, tokenIndex = position220, tokenIndex220
					{
						position222, tokenIndex222 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l215
						}
						position, tokenIndex = position222, tokenIndex222
					}
				}
			l220:
				add(ruleRankAgamo, position216)
			}
			return true
		l215:
			position, tokenIndex = position215, tokenIndex215
			return false
		},
		/* 27 SubGenusOrSuperspecies <- <('(' _? NameLowerChar+ _? ')')> */
		func() bool {
			position223, tokenIndex223 := position, tokenIndex
			{
				position224 := position
				if buffer[position] != rune('(') {
					goto l223
				}
				position++
				{
					position225, tokenIndex225 := position, tokenIndex
					if !_rules[rule_]() {
						goto l225
					}
					goto l226
				l225:
					position, tokenIndex = position225, tokenIndex225
				}
			l226:
				if !_rules[ruleNameLowerChar]() {
					goto l223
				}
			l227:
				{
					position228, tokenIndex228 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l228
					}
					goto l227
				l228:
					position, tokenIndex = position228, tokenIndex228
				}
				{
					position229, tokenIndex229 := position, tokenIndex
					if !_rules[rule_]() {
						goto l229
					}
					goto l230
				l229:
					position, tokenIndex = position229, tokenIndex229
				}
			l230:
				if buffer[position] != rune(')') {
					goto l223
				}
				position++
				add(ruleSubGenusOrSuperspecies, position224)
			}
			return true
		l223:
			position, tokenIndex = position223, tokenIndex223
			return false
		},
		/* 28 SubGenus <- <('(' _? UninomialWord _? ')')> */
		func() bool {
			position231, tokenIndex231 := position, tokenIndex
			{
				position232 := position
				if buffer[position] != rune('(') {
					goto l231
				}
				position++
				{
					position233, tokenIndex233 := position, tokenIndex
					if !_rules[rule_]() {
						goto l233
					}
					goto l234
				l233:
					position, tokenIndex = position233, tokenIndex233
				}
			l234:
				if !_rules[ruleUninomialWord]() {
					goto l231
				}
				{
					position235, tokenIndex235 := position, tokenIndex
					if !_rules[rule_]() {
						goto l235
					}
					goto l236
				l235:
					position, tokenIndex = position235, tokenIndex235
				}
			l236:
				if buffer[position] != rune(')') {
					goto l231
				}
				position++
				add(ruleSubGenus, position232)
			}
			return true
		l231:
			position, tokenIndex = position231, tokenIndex231
			return false
		},
		/* 29 UninomialCombo <- <(UninomialCombo1 / UninomialCombo2)> */
		func() bool {
			position237, tokenIndex237 := position, tokenIndex
			{
				position238 := position
				{
					position239, tokenIndex239 := position, tokenIndex
					if !_rules[ruleUninomialCombo1]() {
						goto l240
					}
					goto l239
				l240:
					position, tokenIndex = position239, tokenIndex239
					if !_rules[ruleUninomialCombo2]() {
						goto l237
					}
				}
			l239:
				add(ruleUninomialCombo, position238)
			}
			return true
		l237:
			position, tokenIndex = position237, tokenIndex237
			return false
		},
		/* 30 UninomialCombo1 <- <(UninomialWord _? SubGenus (_? Authorship)?)> */
		func() bool {
			position241, tokenIndex241 := position, tokenIndex
			{
				position242 := position
				if !_rules[ruleUninomialWord]() {
					goto l241
				}
				{
					position243, tokenIndex243 := position, tokenIndex
					if !_rules[rule_]() {
						goto l243
					}
					goto l244
				l243:
					position, tokenIndex = position243, tokenIndex243
				}
			l244:
				if !_rules[ruleSubGenus]() {
					goto l241
				}
				{
					position245, tokenIndex245 := position, tokenIndex
					{
						position247, tokenIndex247 := position, tokenIndex
						if !_rules[rule_]() {
							goto l247
						}
						goto l248
					l247:
						position, tokenIndex = position247, tokenIndex247
					}
				l248:
					if !_rules[ruleAuthorship]() {
						goto l245
					}
					goto l246
				l245:
					position, tokenIndex = position245, tokenIndex245
				}
			l246:
				add(ruleUninomialCombo1, position242)
			}
			return true
		l241:
			position, tokenIndex = position241, tokenIndex241
			return false
		},
		/* 31 UninomialCombo2 <- <(Uninomial _ RankUninomial _ Uninomial)> */
		func() bool {
			position249, tokenIndex249 := position, tokenIndex
			{
				position250 := position
				if !_rules[ruleUninomial]() {
					goto l249
				}
				if !_rules[rule_]() {
					goto l249
				}
				if !_rules[ruleRankUninomial]() {
					goto l249
				}
				if !_rules[rule_]() {
					goto l249
				}
				if !_rules[ruleUninomial]() {
					goto l249
				}
				add(ruleUninomialCombo2, position250)
			}
			return true
		l249:
			position, tokenIndex = position249, tokenIndex249
			return false
		},
		/* 32 RankUninomial <- <(RankUninomialPlain / RankUninomialNotho)> */
		func() bool {
			position251, tokenIndex251 := position, tokenIndex
			{
				position252 := position
				{
					position253, tokenIndex253 := position, tokenIndex
					if !_rules[ruleRankUninomialPlain]() {
						goto l254
					}
					goto l253
				l254:
					position, tokenIndex = position253, tokenIndex253
					if !_rules[ruleRankUninomialNotho]() {
						goto l251
					}
				}
			l253:
				add(ruleRankUninomial, position252)
			}
			return true
		l251:
			position, tokenIndex = position251, tokenIndex251
			return false
		},
		/* 33 RankUninomialPlain <- <((('s' 'e' 'c' 't') / ('s' 'u' 'b' 's' 'e' 'c' 't') / ('t' 'r' 'i' 'b') / ('s' 'u' 'b' 't' 'r' 'i' 'b') / ('s' 'u' 'b' 's' 'e' 'r') / ('s' 'e' 'r') / ('s' 'u' 'b' 'g' 'e' 'n') / ('s' 'u' 'b' 'g') / ('f' 'a' 'm') / ('s' 'u' 'b' 'f' 'a' 'm') / ('s' 'u' 'p' 'e' 'r' 't' 'r' 'i' 'b')) ('.' / &SpaceCharEOI))> */
		func() bool {
			position255, tokenIndex255 := position, tokenIndex
			{
				position256 := position
				{
					position257, tokenIndex257 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l258
					}
					position++
					if buffer[position] != rune('e') {
						goto l258
					}
					position++
					if buffer[position] != rune('c') {
						goto l258
					}
					position++
					if buffer[position] != rune('t') {
						goto l258
					}
					position++
					goto l257
				l258:
					position, tokenIndex = position257, tokenIndex257
					if buffer[position] != rune('s') {
						goto l259
					}
					position++
					if buffer[position] != rune('u') {
						goto l259
					}
					position++
					if buffer[position] != rune('b') {
						goto l259
					}
					position++
					if buffer[position] != rune('s') {
						goto l259
					}
					position++
					if buffer[position] != rune('e') {
						goto l259
					}
					position++
					if buffer[position] != rune('c') {
						goto l259
					}
					position++
					if buffer[position] != rune('t') {
						goto l259
					}
					position++
					goto l257
				l259:
					position, tokenIndex = position257, tokenIndex257
					if buffer[position] != rune('t') {
						goto l260
					}
					position++
					if buffer[position] != rune('r') {
						goto l260
					}
					position++
					if buffer[position] != rune('i') {
						goto l260
					}
					position++
					if buffer[position] != rune('b') {
						goto l260
					}
					position++
					goto l257
				l260:
					position, tokenIndex = position257, tokenIndex257
					if buffer[position] != rune('s') {
						goto l261
					}
					position++
					if buffer[position] != rune('u') {
						goto l261
					}
					position++
					if buffer[position] != rune('b') {
						goto l261
					}
					position++
					if buffer[position] != rune('t') {
						goto l261
					}
					position++
					if buffer[position] != rune('r') {
						goto l261
					}
					position++
					if buffer[position] != rune('i') {
						goto l261
					}
					position++
					if buffer[position] != rune('b') {
						goto l261
					}
					position++
					goto l257
				l261:
					position, tokenIndex = position257, tokenIndex257
					if buffer[position] != rune('s') {
						goto l262
					}
					position++
					if buffer[position] != rune('u') {
						goto l262
					}
					position++
					if buffer[position] != rune('b') {
						goto l262
					}
					position++
					if buffer[position] != rune('s') {
						goto l262
					}
					position++
					if buffer[position] != rune('e') {
						goto l262
					}
					position++
					if buffer[position] != rune('r') {
						goto l262
					}
					position++
					goto l257
				l262:
					position, tokenIndex = position257, tokenIndex257
					if buffer[position] != rune('s') {
						goto l263
					}
					position++
					if buffer[position] != rune('e') {
						goto l263
					}
					position++
					if buffer[position] != rune('r') {
						goto l263
					}
					position++
					goto l257
				l263:
					position, tokenIndex = position257, tokenIndex257
					if buffer[position] != rune('s') {
						goto l264
					}
					position++
					if buffer[position] != rune('u') {
						goto l264
					}
					position++
					if buffer[position] != rune('b') {
						goto l264
					}
					position++
					if buffer[position] != rune('g') {
						goto l264
					}
					position++
					if buffer[position] != rune('e') {
						goto l264
					}
					position++
					if buffer[position] != rune('n') {
						goto l264
					}
					position++
					goto l257
				l264:
					position, tokenIndex = position257, tokenIndex257
					if buffer[position] != rune('s') {
						goto l265
					}
					position++
					if buffer[position] != rune('u') {
						goto l265
					}
					position++
					if buffer[position] != rune('b') {
						goto l265
					}
					position++
					if buffer[position] != rune('g') {
						goto l265
					}
					position++
					goto l257
				l265:
					position, tokenIndex = position257, tokenIndex257
					if buffer[position] != rune('f') {
						goto l266
					}
					position++
					if buffer[position] != rune('a') {
						goto l266
					}
					position++
					if buffer[position] != rune('m') {
						goto l266
					}
					position++
					goto l257
				l266:
					position, tokenIndex = position257, tokenIndex257
					if buffer[position] != rune('s') {
						goto l267
					}
					position++
					if buffer[position] != rune('u') {
						goto l267
					}
					position++
					if buffer[position] != rune('b') {
						goto l267
					}
					position++
					if buffer[position] != rune('f') {
						goto l267
					}
					position++
					if buffer[position] != rune('a') {
						goto l267
					}
					position++
					if buffer[position] != rune('m') {
						goto l267
					}
					position++
					goto l257
				l267:
					position, tokenIndex = position257, tokenIndex257
					if buffer[position] != rune('s') {
						goto l255
					}
					position++
					if buffer[position] != rune('u') {
						goto l255
					}
					position++
					if buffer[position] != rune('p') {
						goto l255
					}
					position++
					if buffer[position] != rune('e') {
						goto l255
					}
					position++
					if buffer[position] != rune('r') {
						goto l255
					}
					position++
					if buffer[position] != rune('t') {
						goto l255
					}
					position++
					if buffer[position] != rune('r') {
						goto l255
					}
					position++
					if buffer[position] != rune('i') {
						goto l255
					}
					position++
					if buffer[position] != rune('b') {
						goto l255
					}
					position++
				}
			l257:
				{
					position268, tokenIndex268 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l269
					}
					position++
					goto l268
				l269:
					position, tokenIndex = position268, tokenIndex268
					{
						position270, tokenIndex270 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l255
						}
						position, tokenIndex = position270, tokenIndex270
					}
				}
			l268:
				add(ruleRankUninomialPlain, position256)
			}
			return true
		l255:
			position, tokenIndex = position255, tokenIndex255
			return false
		},
		/* 34 RankUninomialNotho <- <('n' 'o' 't' 'h' 'o' _? (('s' 'e' 'c' 't') / ('g' 'e' 'n') / ('s' 'e' 'r') / ('s' 'u' 'b' 'g' 'e' 'e' 'n') / ('s' 'u' 'b' 'g' 'e' 'n') / ('s' 'u' 'b' 'g') / ('s' 'u' 'b' 's' 'e' 'c' 't') / ('s' 'u' 'b' 't' 'r' 'i' 'b')) ('.' / &SpaceCharEOI))> */
		func() bool {
			position271, tokenIndex271 := position, tokenIndex
			{
				position272 := position
				if buffer[position] != rune('n') {
					goto l271
				}
				position++
				if buffer[position] != rune('o') {
					goto l271
				}
				position++
				if buffer[position] != rune('t') {
					goto l271
				}
				position++
				if buffer[position] != rune('h') {
					goto l271
				}
				position++
				if buffer[position] != rune('o') {
					goto l271
				}
				position++
				{
					position273, tokenIndex273 := position, tokenIndex
					if !_rules[rule_]() {
						goto l273
					}
					goto l274
				l273:
					position, tokenIndex = position273, tokenIndex273
				}
			l274:
				{
					position275, tokenIndex275 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l276
					}
					position++
					if buffer[position] != rune('e') {
						goto l276
					}
					position++
					if buffer[position] != rune('c') {
						goto l276
					}
					position++
					if buffer[position] != rune('t') {
						goto l276
					}
					position++
					goto l275
				l276:
					position, tokenIndex = position275, tokenIndex275
					if buffer[position] != rune('g') {
						goto l277
					}
					position++
					if buffer[position] != rune('e') {
						goto l277
					}
					position++
					if buffer[position] != rune('n') {
						goto l277
					}
					position++
					goto l275
				l277:
					position, tokenIndex = position275, tokenIndex275
					if buffer[position] != rune('s') {
						goto l278
					}
					position++
					if buffer[position] != rune('e') {
						goto l278
					}
					position++
					if buffer[position] != rune('r') {
						goto l278
					}
					position++
					goto l275
				l278:
					position, tokenIndex = position275, tokenIndex275
					if buffer[position] != rune('s') {
						goto l279
					}
					position++
					if buffer[position] != rune('u') {
						goto l279
					}
					position++
					if buffer[position] != rune('b') {
						goto l279
					}
					position++
					if buffer[position] != rune('g') {
						goto l279
					}
					position++
					if buffer[position] != rune('e') {
						goto l279
					}
					position++
					if buffer[position] != rune('e') {
						goto l279
					}
					position++
					if buffer[position] != rune('n') {
						goto l279
					}
					position++
					goto l275
				l279:
					position, tokenIndex = position275, tokenIndex275
					if buffer[position] != rune('s') {
						goto l280
					}
					position++
					if buffer[position] != rune('u') {
						goto l280
					}
					position++
					if buffer[position] != rune('b') {
						goto l280
					}
					position++
					if buffer[position] != rune('g') {
						goto l280
					}
					position++
					if buffer[position] != rune('e') {
						goto l280
					}
					position++
					if buffer[position] != rune('n') {
						goto l280
					}
					position++
					goto l275
				l280:
					position, tokenIndex = position275, tokenIndex275
					if buffer[position] != rune('s') {
						goto l281
					}
					position++
					if buffer[position] != rune('u') {
						goto l281
					}
					position++
					if buffer[position] != rune('b') {
						goto l281
					}
					position++
					if buffer[position] != rune('g') {
						goto l281
					}
					position++
					goto l275
				l281:
					position, tokenIndex = position275, tokenIndex275
					if buffer[position] != rune('s') {
						goto l282
					}
					position++
					if buffer[position] != rune('u') {
						goto l282
					}
					position++
					if buffer[position] != rune('b') {
						goto l282
					}
					position++
					if buffer[position] != rune('s') {
						goto l282
					}
					position++
					if buffer[position] != rune('e') {
						goto l282
					}
					position++
					if buffer[position] != rune('c') {
						goto l282
					}
					position++
					if buffer[position] != rune('t') {
						goto l282
					}
					position++
					goto l275
				l282:
					position, tokenIndex = position275, tokenIndex275
					if buffer[position] != rune('s') {
						goto l271
					}
					position++
					if buffer[position] != rune('u') {
						goto l271
					}
					position++
					if buffer[position] != rune('b') {
						goto l271
					}
					position++
					if buffer[position] != rune('t') {
						goto l271
					}
					position++
					if buffer[position] != rune('r') {
						goto l271
					}
					position++
					if buffer[position] != rune('i') {
						goto l271
					}
					position++
					if buffer[position] != rune('b') {
						goto l271
					}
					position++
				}
			l275:
				{
					position283, tokenIndex283 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l284
					}
					position++
					goto l283
				l284:
					position, tokenIndex = position283, tokenIndex283
					{
						position285, tokenIndex285 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l271
						}
						position, tokenIndex = position285, tokenIndex285
					}
				}
			l283:
				add(ruleRankUninomialNotho, position272)
			}
			return true
		l271:
			position, tokenIndex = position271, tokenIndex271
			return false
		},
		/* 35 Uninomial <- <(UninomialWord (_ Authorship)?)> */
		func() bool {
			position286, tokenIndex286 := position, tokenIndex
			{
				position287 := position
				if !_rules[ruleUninomialWord]() {
					goto l286
				}
				{
					position288, tokenIndex288 := position, tokenIndex
					if !_rules[rule_]() {
						goto l288
					}
					if !_rules[ruleAuthorship]() {
						goto l288
					}
					goto l289
				l288:
					position, tokenIndex = position288, tokenIndex288
				}
			l289:
				add(ruleUninomial, position287)
			}
			return true
		l286:
			position, tokenIndex = position286, tokenIndex286
			return false
		},
		/* 36 UninomialWord <- <(CapWord / TwoLetterGenus)> */
		func() bool {
			position290, tokenIndex290 := position, tokenIndex
			{
				position291 := position
				{
					position292, tokenIndex292 := position, tokenIndex
					if !_rules[ruleCapWord]() {
						goto l293
					}
					goto l292
				l293:
					position, tokenIndex = position292, tokenIndex292
					if !_rules[ruleTwoLetterGenus]() {
						goto l290
					}
				}
			l292:
				add(ruleUninomialWord, position291)
			}
			return true
		l290:
			position, tokenIndex = position290, tokenIndex290
			return false
		},
		/* 37 AbbrGenus <- <(UpperChar LowerChar? '.')> */
		func() bool {
			position294, tokenIndex294 := position, tokenIndex
			{
				position295 := position
				if !_rules[ruleUpperChar]() {
					goto l294
				}
				{
					position296, tokenIndex296 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l296
					}
					goto l297
				l296:
					position, tokenIndex = position296, tokenIndex296
				}
			l297:
				if buffer[position] != rune('.') {
					goto l294
				}
				position++
				add(ruleAbbrGenus, position295)
			}
			return true
		l294:
			position, tokenIndex = position294, tokenIndex294
			return false
		},
		/* 38 CapWord <- <(CapWordWithDash / CapWord1)> */
		func() bool {
			position298, tokenIndex298 := position, tokenIndex
			{
				position299 := position
				{
					position300, tokenIndex300 := position, tokenIndex
					if !_rules[ruleCapWordWithDash]() {
						goto l301
					}
					goto l300
				l301:
					position, tokenIndex = position300, tokenIndex300
					if !_rules[ruleCapWord1]() {
						goto l298
					}
				}
			l300:
				add(ruleCapWord, position299)
			}
			return true
		l298:
			position, tokenIndex = position298, tokenIndex298
			return false
		},
		/* 39 CapWord1 <- <(NameUpperChar NameLowerChar NameLowerChar+ '?'?)> */
		func() bool {
			position302, tokenIndex302 := position, tokenIndex
			{
				position303 := position
				if !_rules[ruleNameUpperChar]() {
					goto l302
				}
				if !_rules[ruleNameLowerChar]() {
					goto l302
				}
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
					if buffer[position] != rune('?') {
						goto l306
					}
					position++
					goto l307
				l306:
					position, tokenIndex = position306, tokenIndex306
				}
			l307:
				add(ruleCapWord1, position303)
			}
			return true
		l302:
			position, tokenIndex = position302, tokenIndex302
			return false
		},
		/* 40 CapWordWithDash <- <(CapWord1 Dash (UpperAfterDash / LowerAfterDash))> */
		func() bool {
			position308, tokenIndex308 := position, tokenIndex
			{
				position309 := position
				if !_rules[ruleCapWord1]() {
					goto l308
				}
				if !_rules[ruleDash]() {
					goto l308
				}
				{
					position310, tokenIndex310 := position, tokenIndex
					if !_rules[ruleUpperAfterDash]() {
						goto l311
					}
					goto l310
				l311:
					position, tokenIndex = position310, tokenIndex310
					if !_rules[ruleLowerAfterDash]() {
						goto l308
					}
				}
			l310:
				add(ruleCapWordWithDash, position309)
			}
			return true
		l308:
			position, tokenIndex = position308, tokenIndex308
			return false
		},
		/* 41 UpperAfterDash <- <CapWord1> */
		func() bool {
			position312, tokenIndex312 := position, tokenIndex
			{
				position313 := position
				if !_rules[ruleCapWord1]() {
					goto l312
				}
				add(ruleUpperAfterDash, position313)
			}
			return true
		l312:
			position, tokenIndex = position312, tokenIndex312
			return false
		},
		/* 42 LowerAfterDash <- <Word1> */
		func() bool {
			position314, tokenIndex314 := position, tokenIndex
			{
				position315 := position
				if !_rules[ruleWord1]() {
					goto l314
				}
				add(ruleLowerAfterDash, position315)
			}
			return true
		l314:
			position, tokenIndex = position314, tokenIndex314
			return false
		},
		/* 43 TwoLetterGenus <- <(('C' 'a') / ('E' 'a') / ('G' 'e') / ('I' 'a') / ('I' 'o') / ('I' 'x') / ('L' 'o') / ('O' 'a') / ('R' 'a') / ('T' 'y') / ('U' 'a') / ('A' 'a') / ('J' 'a') / ('Z' 'u') / ('L' 'a') / ('Q' 'u') / ('A' 's') / ('B' 'a'))> */
		func() bool {
			position316, tokenIndex316 := position, tokenIndex
			{
				position317 := position
				{
					position318, tokenIndex318 := position, tokenIndex
					if buffer[position] != rune('C') {
						goto l319
					}
					position++
					if buffer[position] != rune('a') {
						goto l319
					}
					position++
					goto l318
				l319:
					position, tokenIndex = position318, tokenIndex318
					if buffer[position] != rune('E') {
						goto l320
					}
					position++
					if buffer[position] != rune('a') {
						goto l320
					}
					position++
					goto l318
				l320:
					position, tokenIndex = position318, tokenIndex318
					if buffer[position] != rune('G') {
						goto l321
					}
					position++
					if buffer[position] != rune('e') {
						goto l321
					}
					position++
					goto l318
				l321:
					position, tokenIndex = position318, tokenIndex318
					if buffer[position] != rune('I') {
						goto l322
					}
					position++
					if buffer[position] != rune('a') {
						goto l322
					}
					position++
					goto l318
				l322:
					position, tokenIndex = position318, tokenIndex318
					if buffer[position] != rune('I') {
						goto l323
					}
					position++
					if buffer[position] != rune('o') {
						goto l323
					}
					position++
					goto l318
				l323:
					position, tokenIndex = position318, tokenIndex318
					if buffer[position] != rune('I') {
						goto l324
					}
					position++
					if buffer[position] != rune('x') {
						goto l324
					}
					position++
					goto l318
				l324:
					position, tokenIndex = position318, tokenIndex318
					if buffer[position] != rune('L') {
						goto l325
					}
					position++
					if buffer[position] != rune('o') {
						goto l325
					}
					position++
					goto l318
				l325:
					position, tokenIndex = position318, tokenIndex318
					if buffer[position] != rune('O') {
						goto l326
					}
					position++
					if buffer[position] != rune('a') {
						goto l326
					}
					position++
					goto l318
				l326:
					position, tokenIndex = position318, tokenIndex318
					if buffer[position] != rune('R') {
						goto l327
					}
					position++
					if buffer[position] != rune('a') {
						goto l327
					}
					position++
					goto l318
				l327:
					position, tokenIndex = position318, tokenIndex318
					if buffer[position] != rune('T') {
						goto l328
					}
					position++
					if buffer[position] != rune('y') {
						goto l328
					}
					position++
					goto l318
				l328:
					position, tokenIndex = position318, tokenIndex318
					if buffer[position] != rune('U') {
						goto l329
					}
					position++
					if buffer[position] != rune('a') {
						goto l329
					}
					position++
					goto l318
				l329:
					position, tokenIndex = position318, tokenIndex318
					if buffer[position] != rune('A') {
						goto l330
					}
					position++
					if buffer[position] != rune('a') {
						goto l330
					}
					position++
					goto l318
				l330:
					position, tokenIndex = position318, tokenIndex318
					if buffer[position] != rune('J') {
						goto l331
					}
					position++
					if buffer[position] != rune('a') {
						goto l331
					}
					position++
					goto l318
				l331:
					position, tokenIndex = position318, tokenIndex318
					if buffer[position] != rune('Z') {
						goto l332
					}
					position++
					if buffer[position] != rune('u') {
						goto l332
					}
					position++
					goto l318
				l332:
					position, tokenIndex = position318, tokenIndex318
					if buffer[position] != rune('L') {
						goto l333
					}
					position++
					if buffer[position] != rune('a') {
						goto l333
					}
					position++
					goto l318
				l333:
					position, tokenIndex = position318, tokenIndex318
					if buffer[position] != rune('Q') {
						goto l334
					}
					position++
					if buffer[position] != rune('u') {
						goto l334
					}
					position++
					goto l318
				l334:
					position, tokenIndex = position318, tokenIndex318
					if buffer[position] != rune('A') {
						goto l335
					}
					position++
					if buffer[position] != rune('s') {
						goto l335
					}
					position++
					goto l318
				l335:
					position, tokenIndex = position318, tokenIndex318
					if buffer[position] != rune('B') {
						goto l316
					}
					position++
					if buffer[position] != rune('a') {
						goto l316
					}
					position++
				}
			l318:
				add(ruleTwoLetterGenus, position317)
			}
			return true
		l316:
			position, tokenIndex = position316, tokenIndex316
			return false
		},
		/* 44 Word <- <(!((AuthorPrefix / RankUninomial / Approximation / Word4) SpaceCharEOI) (WordApostr / WordStartsWithDigit / MultiDashedWord / Word2 / Word1) &(SpaceCharEOI / '('))> */
		func() bool {
			position336, tokenIndex336 := position, tokenIndex
			{
				position337 := position
				{
					position338, tokenIndex338 := position, tokenIndex
					{
						position339, tokenIndex339 := position, tokenIndex
						if !_rules[ruleAuthorPrefix]() {
							goto l340
						}
						goto l339
					l340:
						position, tokenIndex = position339, tokenIndex339
						if !_rules[ruleRankUninomial]() {
							goto l341
						}
						goto l339
					l341:
						position, tokenIndex = position339, tokenIndex339
						if !_rules[ruleApproximation]() {
							goto l342
						}
						goto l339
					l342:
						position, tokenIndex = position339, tokenIndex339
						if !_rules[ruleWord4]() {
							goto l338
						}
					}
				l339:
					if !_rules[ruleSpaceCharEOI]() {
						goto l338
					}
					goto l336
				l338:
					position, tokenIndex = position338, tokenIndex338
				}
				{
					position343, tokenIndex343 := position, tokenIndex
					if !_rules[ruleWordApostr]() {
						goto l344
					}
					goto l343
				l344:
					position, tokenIndex = position343, tokenIndex343
					if !_rules[ruleWordStartsWithDigit]() {
						goto l345
					}
					goto l343
				l345:
					position, tokenIndex = position343, tokenIndex343
					if !_rules[ruleMultiDashedWord]() {
						goto l346
					}
					goto l343
				l346:
					position, tokenIndex = position343, tokenIndex343
					if !_rules[ruleWord2]() {
						goto l347
					}
					goto l343
				l347:
					position, tokenIndex = position343, tokenIndex343
					if !_rules[ruleWord1]() {
						goto l336
					}
				}
			l343:
				{
					position348, tokenIndex348 := position, tokenIndex
					{
						position349, tokenIndex349 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l350
						}
						goto l349
					l350:
						position, tokenIndex = position349, tokenIndex349
						if buffer[position] != rune('(') {
							goto l336
						}
						position++
					}
				l349:
					position, tokenIndex = position348, tokenIndex348
				}
				add(ruleWord, position337)
			}
			return true
		l336:
			position, tokenIndex = position336, tokenIndex336
			return false
		},
		/* 45 Word1 <- <((LowerASCII Dash)? NameLowerChar NameLowerChar+)> */
		func() bool {
			position351, tokenIndex351 := position, tokenIndex
			{
				position352 := position
				{
					position353, tokenIndex353 := position, tokenIndex
					if !_rules[ruleLowerASCII]() {
						goto l353
					}
					if !_rules[ruleDash]() {
						goto l353
					}
					goto l354
				l353:
					position, tokenIndex = position353, tokenIndex353
				}
			l354:
				if !_rules[ruleNameLowerChar]() {
					goto l351
				}
				if !_rules[ruleNameLowerChar]() {
					goto l351
				}
			l355:
				{
					position356, tokenIndex356 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l356
					}
					goto l355
				l356:
					position, tokenIndex = position356, tokenIndex356
				}
				add(ruleWord1, position352)
			}
			return true
		l351:
			position, tokenIndex = position351, tokenIndex351
			return false
		},
		/* 46 WordStartsWithDigit <- <(('1' / '2' / '3' / '4' / '5' / '6' / '7' / '8' / '9') Nums? ('.' / Dash)? NameLowerChar NameLowerChar NameLowerChar NameLowerChar+)> */
		func() bool {
			position357, tokenIndex357 := position, tokenIndex
			{
				position358 := position
				{
					position359, tokenIndex359 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l360
					}
					position++
					goto l359
				l360:
					position, tokenIndex = position359, tokenIndex359
					if buffer[position] != rune('2') {
						goto l361
					}
					position++
					goto l359
				l361:
					position, tokenIndex = position359, tokenIndex359
					if buffer[position] != rune('3') {
						goto l362
					}
					position++
					goto l359
				l362:
					position, tokenIndex = position359, tokenIndex359
					if buffer[position] != rune('4') {
						goto l363
					}
					position++
					goto l359
				l363:
					position, tokenIndex = position359, tokenIndex359
					if buffer[position] != rune('5') {
						goto l364
					}
					position++
					goto l359
				l364:
					position, tokenIndex = position359, tokenIndex359
					if buffer[position] != rune('6') {
						goto l365
					}
					position++
					goto l359
				l365:
					position, tokenIndex = position359, tokenIndex359
					if buffer[position] != rune('7') {
						goto l366
					}
					position++
					goto l359
				l366:
					position, tokenIndex = position359, tokenIndex359
					if buffer[position] != rune('8') {
						goto l367
					}
					position++
					goto l359
				l367:
					position, tokenIndex = position359, tokenIndex359
					if buffer[position] != rune('9') {
						goto l357
					}
					position++
				}
			l359:
				{
					position368, tokenIndex368 := position, tokenIndex
					if !_rules[ruleNums]() {
						goto l368
					}
					goto l369
				l368:
					position, tokenIndex = position368, tokenIndex368
				}
			l369:
				{
					position370, tokenIndex370 := position, tokenIndex
					{
						position372, tokenIndex372 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l373
						}
						position++
						goto l372
					l373:
						position, tokenIndex = position372, tokenIndex372
						if !_rules[ruleDash]() {
							goto l370
						}
					}
				l372:
					goto l371
				l370:
					position, tokenIndex = position370, tokenIndex370
				}
			l371:
				if !_rules[ruleNameLowerChar]() {
					goto l357
				}
				if !_rules[ruleNameLowerChar]() {
					goto l357
				}
				if !_rules[ruleNameLowerChar]() {
					goto l357
				}
				if !_rules[ruleNameLowerChar]() {
					goto l357
				}
			l374:
				{
					position375, tokenIndex375 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l375
					}
					goto l374
				l375:
					position, tokenIndex = position375, tokenIndex375
				}
				add(ruleWordStartsWithDigit, position358)
			}
			return true
		l357:
			position, tokenIndex = position357, tokenIndex357
			return false
		},
		/* 47 Word2 <- <(NameLowerChar+ Dash? NameLowerChar+)> */
		func() bool {
			position376, tokenIndex376 := position, tokenIndex
			{
				position377 := position
				if !_rules[ruleNameLowerChar]() {
					goto l376
				}
			l378:
				{
					position379, tokenIndex379 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l379
					}
					goto l378
				l379:
					position, tokenIndex = position379, tokenIndex379
				}
				{
					position380, tokenIndex380 := position, tokenIndex
					if !_rules[ruleDash]() {
						goto l380
					}
					goto l381
				l380:
					position, tokenIndex = position380, tokenIndex380
				}
			l381:
				if !_rules[ruleNameLowerChar]() {
					goto l376
				}
			l382:
				{
					position383, tokenIndex383 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l383
					}
					goto l382
				l383:
					position, tokenIndex = position383, tokenIndex383
				}
				add(ruleWord2, position377)
			}
			return true
		l376:
			position, tokenIndex = position376, tokenIndex376
			return false
		},
		/* 48 WordApostr <- <(NameLowerChar NameLowerChar* Apostrophe Word1)> */
		func() bool {
			position384, tokenIndex384 := position, tokenIndex
			{
				position385 := position
				if !_rules[ruleNameLowerChar]() {
					goto l384
				}
			l386:
				{
					position387, tokenIndex387 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l387
					}
					goto l386
				l387:
					position, tokenIndex = position387, tokenIndex387
				}
				if !_rules[ruleApostrophe]() {
					goto l384
				}
				if !_rules[ruleWord1]() {
					goto l384
				}
				add(ruleWordApostr, position385)
			}
			return true
		l384:
			position, tokenIndex = position384, tokenIndex384
			return false
		},
		/* 49 Word4 <- <(NameLowerChar+ '.' NameLowerChar)> */
		func() bool {
			position388, tokenIndex388 := position, tokenIndex
			{
				position389 := position
				if !_rules[ruleNameLowerChar]() {
					goto l388
				}
			l390:
				{
					position391, tokenIndex391 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l391
					}
					goto l390
				l391:
					position, tokenIndex = position391, tokenIndex391
				}
				if buffer[position] != rune('.') {
					goto l388
				}
				position++
				if !_rules[ruleNameLowerChar]() {
					goto l388
				}
				add(ruleWord4, position389)
			}
			return true
		l388:
			position, tokenIndex = position388, tokenIndex388
			return false
		},
		/* 50 MultiDashedWord <- <(NameLowerChar+ Dash NameLowerChar+ Dash NameLowerChar+ (Dash NameLowerChar+)?)> */
		func() bool {
			position392, tokenIndex392 := position, tokenIndex
			{
				position393 := position
				if !_rules[ruleNameLowerChar]() {
					goto l392
				}
			l394:
				{
					position395, tokenIndex395 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l395
					}
					goto l394
				l395:
					position, tokenIndex = position395, tokenIndex395
				}
				if !_rules[ruleDash]() {
					goto l392
				}
				if !_rules[ruleNameLowerChar]() {
					goto l392
				}
			l396:
				{
					position397, tokenIndex397 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l397
					}
					goto l396
				l397:
					position, tokenIndex = position397, tokenIndex397
				}
				if !_rules[ruleDash]() {
					goto l392
				}
				if !_rules[ruleNameLowerChar]() {
					goto l392
				}
			l398:
				{
					position399, tokenIndex399 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l399
					}
					goto l398
				l399:
					position, tokenIndex = position399, tokenIndex399
				}
				{
					position400, tokenIndex400 := position, tokenIndex
					if !_rules[ruleDash]() {
						goto l400
					}
					if !_rules[ruleNameLowerChar]() {
						goto l400
					}
				l402:
					{
						position403, tokenIndex403 := position, tokenIndex
						if !_rules[ruleNameLowerChar]() {
							goto l403
						}
						goto l402
					l403:
						position, tokenIndex = position403, tokenIndex403
					}
					goto l401
				l400:
					position, tokenIndex = position400, tokenIndex400
				}
			l401:
				add(ruleMultiDashedWord, position393)
			}
			return true
		l392:
			position, tokenIndex = position392, tokenIndex392
			return false
		},
		/* 51 HybridChar <- <'×'> */
		func() bool {
			position404, tokenIndex404 := position, tokenIndex
			{
				position405 := position
				if buffer[position] != rune('×') {
					goto l404
				}
				position++
				add(ruleHybridChar, position405)
			}
			return true
		l404:
			position, tokenIndex = position404, tokenIndex404
			return false
		},
		/* 52 ApproxNameIgnored <- <.*> */
		func() bool {
			{
				position407 := position
			l408:
				{
					position409, tokenIndex409 := position, tokenIndex
					if !matchDot() {
						goto l409
					}
					goto l408
				l409:
					position, tokenIndex = position409, tokenIndex409
				}
				add(ruleApproxNameIgnored, position407)
			}
			return true
		},
		/* 53 Approximation <- <(('s' 'p' '.' _? ('n' 'r' '.')) / ('s' 'p' '.' _? ('a' 'f' 'f' '.')) / ('m' 'o' 'n' 's' 't' '.') / '?' / ((('s' 'p' 'p') / ('n' 'r') / ('s' 'p') / ('a' 'f' 'f') / ('s' 'p' 'e' 'c' 'i' 'e' 's')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position410, tokenIndex410 := position, tokenIndex
			{
				position411 := position
				{
					position412, tokenIndex412 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l413
					}
					position++
					if buffer[position] != rune('p') {
						goto l413
					}
					position++
					if buffer[position] != rune('.') {
						goto l413
					}
					position++
					{
						position414, tokenIndex414 := position, tokenIndex
						if !_rules[rule_]() {
							goto l414
						}
						goto l415
					l414:
						position, tokenIndex = position414, tokenIndex414
					}
				l415:
					if buffer[position] != rune('n') {
						goto l413
					}
					position++
					if buffer[position] != rune('r') {
						goto l413
					}
					position++
					if buffer[position] != rune('.') {
						goto l413
					}
					position++
					goto l412
				l413:
					position, tokenIndex = position412, tokenIndex412
					if buffer[position] != rune('s') {
						goto l416
					}
					position++
					if buffer[position] != rune('p') {
						goto l416
					}
					position++
					if buffer[position] != rune('.') {
						goto l416
					}
					position++
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
					if buffer[position] != rune('a') {
						goto l416
					}
					position++
					if buffer[position] != rune('f') {
						goto l416
					}
					position++
					if buffer[position] != rune('f') {
						goto l416
					}
					position++
					if buffer[position] != rune('.') {
						goto l416
					}
					position++
					goto l412
				l416:
					position, tokenIndex = position412, tokenIndex412
					if buffer[position] != rune('m') {
						goto l419
					}
					position++
					if buffer[position] != rune('o') {
						goto l419
					}
					position++
					if buffer[position] != rune('n') {
						goto l419
					}
					position++
					if buffer[position] != rune('s') {
						goto l419
					}
					position++
					if buffer[position] != rune('t') {
						goto l419
					}
					position++
					if buffer[position] != rune('.') {
						goto l419
					}
					position++
					goto l412
				l419:
					position, tokenIndex = position412, tokenIndex412
					if buffer[position] != rune('?') {
						goto l420
					}
					position++
					goto l412
				l420:
					position, tokenIndex = position412, tokenIndex412
					{
						position421, tokenIndex421 := position, tokenIndex
						if buffer[position] != rune('s') {
							goto l422
						}
						position++
						if buffer[position] != rune('p') {
							goto l422
						}
						position++
						if buffer[position] != rune('p') {
							goto l422
						}
						position++
						goto l421
					l422:
						position, tokenIndex = position421, tokenIndex421
						if buffer[position] != rune('n') {
							goto l423
						}
						position++
						if buffer[position] != rune('r') {
							goto l423
						}
						position++
						goto l421
					l423:
						position, tokenIndex = position421, tokenIndex421
						if buffer[position] != rune('s') {
							goto l424
						}
						position++
						if buffer[position] != rune('p') {
							goto l424
						}
						position++
						goto l421
					l424:
						position, tokenIndex = position421, tokenIndex421
						if buffer[position] != rune('a') {
							goto l425
						}
						position++
						if buffer[position] != rune('f') {
							goto l425
						}
						position++
						if buffer[position] != rune('f') {
							goto l425
						}
						position++
						goto l421
					l425:
						position, tokenIndex = position421, tokenIndex421
						if buffer[position] != rune('s') {
							goto l410
						}
						position++
						if buffer[position] != rune('p') {
							goto l410
						}
						position++
						if buffer[position] != rune('e') {
							goto l410
						}
						position++
						if buffer[position] != rune('c') {
							goto l410
						}
						position++
						if buffer[position] != rune('i') {
							goto l410
						}
						position++
						if buffer[position] != rune('e') {
							goto l410
						}
						position++
						if buffer[position] != rune('s') {
							goto l410
						}
						position++
					}
				l421:
					{
						position426, tokenIndex426 := position, tokenIndex
						{
							position428, tokenIndex428 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l427
							}
							position, tokenIndex = position428, tokenIndex428
						}
						goto l426
					l427:
						position, tokenIndex = position426, tokenIndex426
						if buffer[position] != rune('.') {
							goto l410
						}
						position++
					}
				l426:
				}
			l412:
				add(ruleApproximation, position411)
			}
			return true
		l410:
			position, tokenIndex = position410, tokenIndex410
			return false
		},
		/* 54 Authorship <- <((AuthorshipCombo / OriginalAuthorship) &(SpaceCharEOI / ';' / ','))> */
		func() bool {
			position429, tokenIndex429 := position, tokenIndex
			{
				position430 := position
				{
					position431, tokenIndex431 := position, tokenIndex
					if !_rules[ruleAuthorshipCombo]() {
						goto l432
					}
					goto l431
				l432:
					position, tokenIndex = position431, tokenIndex431
					if !_rules[ruleOriginalAuthorship]() {
						goto l429
					}
				}
			l431:
				{
					position433, tokenIndex433 := position, tokenIndex
					{
						position434, tokenIndex434 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l435
						}
						goto l434
					l435:
						position, tokenIndex = position434, tokenIndex434
						if buffer[position] != rune(';') {
							goto l436
						}
						position++
						goto l434
					l436:
						position, tokenIndex = position434, tokenIndex434
						if buffer[position] != rune(',') {
							goto l429
						}
						position++
					}
				l434:
					position, tokenIndex = position433, tokenIndex433
				}
				add(ruleAuthorship, position430)
			}
			return true
		l429:
			position, tokenIndex = position429, tokenIndex429
			return false
		},
		/* 55 AuthorshipCombo <- <(OriginalAuthorshipComb (_? CombinationAuthorship)?)> */
		func() bool {
			position437, tokenIndex437 := position, tokenIndex
			{
				position438 := position
				if !_rules[ruleOriginalAuthorshipComb]() {
					goto l437
				}
				{
					position439, tokenIndex439 := position, tokenIndex
					{
						position441, tokenIndex441 := position, tokenIndex
						if !_rules[rule_]() {
							goto l441
						}
						goto l442
					l441:
						position, tokenIndex = position441, tokenIndex441
					}
				l442:
					if !_rules[ruleCombinationAuthorship]() {
						goto l439
					}
					goto l440
				l439:
					position, tokenIndex = position439, tokenIndex439
				}
			l440:
				add(ruleAuthorshipCombo, position438)
			}
			return true
		l437:
			position, tokenIndex = position437, tokenIndex437
			return false
		},
		/* 56 OriginalAuthorship <- <AuthorsGroup> */
		func() bool {
			position443, tokenIndex443 := position, tokenIndex
			{
				position444 := position
				if !_rules[ruleAuthorsGroup]() {
					goto l443
				}
				add(ruleOriginalAuthorship, position444)
			}
			return true
		l443:
			position, tokenIndex = position443, tokenIndex443
			return false
		},
		/* 57 OriginalAuthorshipComb <- <(BasionymAuthorshipYearMisformed / BasionymAuthorship / BasionymAuthorshipMissingParens)> */
		func() bool {
			position445, tokenIndex445 := position, tokenIndex
			{
				position446 := position
				{
					position447, tokenIndex447 := position, tokenIndex
					if !_rules[ruleBasionymAuthorshipYearMisformed]() {
						goto l448
					}
					goto l447
				l448:
					position, tokenIndex = position447, tokenIndex447
					if !_rules[ruleBasionymAuthorship]() {
						goto l449
					}
					goto l447
				l449:
					position, tokenIndex = position447, tokenIndex447
					if !_rules[ruleBasionymAuthorshipMissingParens]() {
						goto l445
					}
				}
			l447:
				add(ruleOriginalAuthorshipComb, position446)
			}
			return true
		l445:
			position, tokenIndex = position445, tokenIndex445
			return false
		},
		/* 58 CombinationAuthorship <- <AuthorsGroup> */
		func() bool {
			position450, tokenIndex450 := position, tokenIndex
			{
				position451 := position
				if !_rules[ruleAuthorsGroup]() {
					goto l450
				}
				add(ruleCombinationAuthorship, position451)
			}
			return true
		l450:
			position, tokenIndex = position450, tokenIndex450
			return false
		},
		/* 59 BasionymAuthorshipMissingParens <- <(MissingParensStart / MissingParensEnd)> */
		func() bool {
			position452, tokenIndex452 := position, tokenIndex
			{
				position453 := position
				{
					position454, tokenIndex454 := position, tokenIndex
					if !_rules[ruleMissingParensStart]() {
						goto l455
					}
					goto l454
				l455:
					position, tokenIndex = position454, tokenIndex454
					if !_rules[ruleMissingParensEnd]() {
						goto l452
					}
				}
			l454:
				add(ruleBasionymAuthorshipMissingParens, position453)
			}
			return true
		l452:
			position, tokenIndex = position452, tokenIndex452
			return false
		},
		/* 60 MissingParensStart <- <('(' _? AuthorsGroup)> */
		func() bool {
			position456, tokenIndex456 := position, tokenIndex
			{
				position457 := position
				if buffer[position] != rune('(') {
					goto l456
				}
				position++
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
				if !_rules[ruleAuthorsGroup]() {
					goto l456
				}
				add(ruleMissingParensStart, position457)
			}
			return true
		l456:
			position, tokenIndex = position456, tokenIndex456
			return false
		},
		/* 61 MissingParensEnd <- <(AuthorsGroup _? ')')> */
		func() bool {
			position460, tokenIndex460 := position, tokenIndex
			{
				position461 := position
				if !_rules[ruleAuthorsGroup]() {
					goto l460
				}
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
				if buffer[position] != rune(')') {
					goto l460
				}
				position++
				add(ruleMissingParensEnd, position461)
			}
			return true
		l460:
			position, tokenIndex = position460, tokenIndex460
			return false
		},
		/* 62 BasionymAuthorshipYearMisformed <- <('(' _? AuthorsGroup _? ')' (_? ',')? _? Year)> */
		func() bool {
			position464, tokenIndex464 := position, tokenIndex
			{
				position465 := position
				if buffer[position] != rune('(') {
					goto l464
				}
				position++
				{
					position466, tokenIndex466 := position, tokenIndex
					if !_rules[rule_]() {
						goto l466
					}
					goto l467
				l466:
					position, tokenIndex = position466, tokenIndex466
				}
			l467:
				if !_rules[ruleAuthorsGroup]() {
					goto l464
				}
				{
					position468, tokenIndex468 := position, tokenIndex
					if !_rules[rule_]() {
						goto l468
					}
					goto l469
				l468:
					position, tokenIndex = position468, tokenIndex468
				}
			l469:
				if buffer[position] != rune(')') {
					goto l464
				}
				position++
				{
					position470, tokenIndex470 := position, tokenIndex
					{
						position472, tokenIndex472 := position, tokenIndex
						if !_rules[rule_]() {
							goto l472
						}
						goto l473
					l472:
						position, tokenIndex = position472, tokenIndex472
					}
				l473:
					if buffer[position] != rune(',') {
						goto l470
					}
					position++
					goto l471
				l470:
					position, tokenIndex = position470, tokenIndex470
				}
			l471:
				{
					position474, tokenIndex474 := position, tokenIndex
					if !_rules[rule_]() {
						goto l474
					}
					goto l475
				l474:
					position, tokenIndex = position474, tokenIndex474
				}
			l475:
				if !_rules[ruleYear]() {
					goto l464
				}
				add(ruleBasionymAuthorshipYearMisformed, position465)
			}
			return true
		l464:
			position, tokenIndex = position464, tokenIndex464
			return false
		},
		/* 63 BasionymAuthorship <- <(BasionymAuthorship1 / BasionymAuthorship2Parens)> */
		func() bool {
			position476, tokenIndex476 := position, tokenIndex
			{
				position477 := position
				{
					position478, tokenIndex478 := position, tokenIndex
					if !_rules[ruleBasionymAuthorship1]() {
						goto l479
					}
					goto l478
				l479:
					position, tokenIndex = position478, tokenIndex478
					if !_rules[ruleBasionymAuthorship2Parens]() {
						goto l476
					}
				}
			l478:
				add(ruleBasionymAuthorship, position477)
			}
			return true
		l476:
			position, tokenIndex = position476, tokenIndex476
			return false
		},
		/* 64 BasionymAuthorship1 <- <('(' _? AuthorsGroup _? ')')> */
		func() bool {
			position480, tokenIndex480 := position, tokenIndex
			{
				position481 := position
				if buffer[position] != rune('(') {
					goto l480
				}
				position++
				{
					position482, tokenIndex482 := position, tokenIndex
					if !_rules[rule_]() {
						goto l482
					}
					goto l483
				l482:
					position, tokenIndex = position482, tokenIndex482
				}
			l483:
				if !_rules[ruleAuthorsGroup]() {
					goto l480
				}
				{
					position484, tokenIndex484 := position, tokenIndex
					if !_rules[rule_]() {
						goto l484
					}
					goto l485
				l484:
					position, tokenIndex = position484, tokenIndex484
				}
			l485:
				if buffer[position] != rune(')') {
					goto l480
				}
				position++
				add(ruleBasionymAuthorship1, position481)
			}
			return true
		l480:
			position, tokenIndex = position480, tokenIndex480
			return false
		},
		/* 65 BasionymAuthorship2Parens <- <('(' _? '(' _? AuthorsGroup _? ')' _? ')')> */
		func() bool {
			position486, tokenIndex486 := position, tokenIndex
			{
				position487 := position
				if buffer[position] != rune('(') {
					goto l486
				}
				position++
				{
					position488, tokenIndex488 := position, tokenIndex
					if !_rules[rule_]() {
						goto l488
					}
					goto l489
				l488:
					position, tokenIndex = position488, tokenIndex488
				}
			l489:
				if buffer[position] != rune('(') {
					goto l486
				}
				position++
				{
					position490, tokenIndex490 := position, tokenIndex
					if !_rules[rule_]() {
						goto l490
					}
					goto l491
				l490:
					position, tokenIndex = position490, tokenIndex490
				}
			l491:
				if !_rules[ruleAuthorsGroup]() {
					goto l486
				}
				{
					position492, tokenIndex492 := position, tokenIndex
					if !_rules[rule_]() {
						goto l492
					}
					goto l493
				l492:
					position, tokenIndex = position492, tokenIndex492
				}
			l493:
				if buffer[position] != rune(')') {
					goto l486
				}
				position++
				{
					position494, tokenIndex494 := position, tokenIndex
					if !_rules[rule_]() {
						goto l494
					}
					goto l495
				l494:
					position, tokenIndex = position494, tokenIndex494
				}
			l495:
				if buffer[position] != rune(')') {
					goto l486
				}
				position++
				add(ruleBasionymAuthorship2Parens, position487)
			}
			return true
		l486:
			position, tokenIndex = position486, tokenIndex486
			return false
		},
		/* 66 AuthorsGroup <- <(AuthorsTeam (_ (AuthorEmend / AuthorEx) AuthorsTeam)?)> */
		func() bool {
			position496, tokenIndex496 := position, tokenIndex
			{
				position497 := position
				if !_rules[ruleAuthorsTeam]() {
					goto l496
				}
				{
					position498, tokenIndex498 := position, tokenIndex
					if !_rules[rule_]() {
						goto l498
					}
					{
						position500, tokenIndex500 := position, tokenIndex
						if !_rules[ruleAuthorEmend]() {
							goto l501
						}
						goto l500
					l501:
						position, tokenIndex = position500, tokenIndex500
						if !_rules[ruleAuthorEx]() {
							goto l498
						}
					}
				l500:
					if !_rules[ruleAuthorsTeam]() {
						goto l498
					}
					goto l499
				l498:
					position, tokenIndex = position498, tokenIndex498
				}
			l499:
				add(ruleAuthorsGroup, position497)
			}
			return true
		l496:
			position, tokenIndex = position496, tokenIndex496
			return false
		},
		/* 67 AuthorsTeam <- <(Author (AuthorSep Author)* (_? ','? _? Year)?)> */
		func() bool {
			position502, tokenIndex502 := position, tokenIndex
			{
				position503 := position
				if !_rules[ruleAuthor]() {
					goto l502
				}
			l504:
				{
					position505, tokenIndex505 := position, tokenIndex
					if !_rules[ruleAuthorSep]() {
						goto l505
					}
					if !_rules[ruleAuthor]() {
						goto l505
					}
					goto l504
				l505:
					position, tokenIndex = position505, tokenIndex505
				}
				{
					position506, tokenIndex506 := position, tokenIndex
					{
						position508, tokenIndex508 := position, tokenIndex
						if !_rules[rule_]() {
							goto l508
						}
						goto l509
					l508:
						position, tokenIndex = position508, tokenIndex508
					}
				l509:
					{
						position510, tokenIndex510 := position, tokenIndex
						if buffer[position] != rune(',') {
							goto l510
						}
						position++
						goto l511
					l510:
						position, tokenIndex = position510, tokenIndex510
					}
				l511:
					{
						position512, tokenIndex512 := position, tokenIndex
						if !_rules[rule_]() {
							goto l512
						}
						goto l513
					l512:
						position, tokenIndex = position512, tokenIndex512
					}
				l513:
					if !_rules[ruleYear]() {
						goto l506
					}
					goto l507
				l506:
					position, tokenIndex = position506, tokenIndex506
				}
			l507:
				add(ruleAuthorsTeam, position503)
			}
			return true
		l502:
			position, tokenIndex = position502, tokenIndex502
			return false
		},
		/* 68 AuthorSep <- <(AuthorSep1 / AuthorSep2)> */
		func() bool {
			position514, tokenIndex514 := position, tokenIndex
			{
				position515 := position
				{
					position516, tokenIndex516 := position, tokenIndex
					if !_rules[ruleAuthorSep1]() {
						goto l517
					}
					goto l516
				l517:
					position, tokenIndex = position516, tokenIndex516
					if !_rules[ruleAuthorSep2]() {
						goto l514
					}
				}
			l516:
				add(ruleAuthorSep, position515)
			}
			return true
		l514:
			position, tokenIndex = position514, tokenIndex514
			return false
		},
		/* 69 AuthorSep1 <- <(_? (',' _)? ('&' / ('e' 't') / ('a' 'n' 'd') / ('a' 'p' 'u' 'd')) _?)> */
		func() bool {
			position518, tokenIndex518 := position, tokenIndex
			{
				position519 := position
				{
					position520, tokenIndex520 := position, tokenIndex
					if !_rules[rule_]() {
						goto l520
					}
					goto l521
				l520:
					position, tokenIndex = position520, tokenIndex520
				}
			l521:
				{
					position522, tokenIndex522 := position, tokenIndex
					if buffer[position] != rune(',') {
						goto l522
					}
					position++
					if !_rules[rule_]() {
						goto l522
					}
					goto l523
				l522:
					position, tokenIndex = position522, tokenIndex522
				}
			l523:
				{
					position524, tokenIndex524 := position, tokenIndex
					if buffer[position] != rune('&') {
						goto l525
					}
					position++
					goto l524
				l525:
					position, tokenIndex = position524, tokenIndex524
					if buffer[position] != rune('e') {
						goto l526
					}
					position++
					if buffer[position] != rune('t') {
						goto l526
					}
					position++
					goto l524
				l526:
					position, tokenIndex = position524, tokenIndex524
					if buffer[position] != rune('a') {
						goto l527
					}
					position++
					if buffer[position] != rune('n') {
						goto l527
					}
					position++
					if buffer[position] != rune('d') {
						goto l527
					}
					position++
					goto l524
				l527:
					position, tokenIndex = position524, tokenIndex524
					if buffer[position] != rune('a') {
						goto l518
					}
					position++
					if buffer[position] != rune('p') {
						goto l518
					}
					position++
					if buffer[position] != rune('u') {
						goto l518
					}
					position++
					if buffer[position] != rune('d') {
						goto l518
					}
					position++
				}
			l524:
				{
					position528, tokenIndex528 := position, tokenIndex
					if !_rules[rule_]() {
						goto l528
					}
					goto l529
				l528:
					position, tokenIndex = position528, tokenIndex528
				}
			l529:
				add(ruleAuthorSep1, position519)
			}
			return true
		l518:
			position, tokenIndex = position518, tokenIndex518
			return false
		},
		/* 70 AuthorSep2 <- <(_? ',' _?)> */
		func() bool {
			position530, tokenIndex530 := position, tokenIndex
			{
				position531 := position
				{
					position532, tokenIndex532 := position, tokenIndex
					if !_rules[rule_]() {
						goto l532
					}
					goto l533
				l532:
					position, tokenIndex = position532, tokenIndex532
				}
			l533:
				if buffer[position] != rune(',') {
					goto l530
				}
				position++
				{
					position534, tokenIndex534 := position, tokenIndex
					if !_rules[rule_]() {
						goto l534
					}
					goto l535
				l534:
					position, tokenIndex = position534, tokenIndex534
				}
			l535:
				add(ruleAuthorSep2, position531)
			}
			return true
		l530:
			position, tokenIndex = position530, tokenIndex530
			return false
		},
		/* 71 AuthorEx <- <((('e' 'x' '.'?) / ('i' 'n')) _)> */
		func() bool {
			position536, tokenIndex536 := position, tokenIndex
			{
				position537 := position
				{
					position538, tokenIndex538 := position, tokenIndex
					if buffer[position] != rune('e') {
						goto l539
					}
					position++
					if buffer[position] != rune('x') {
						goto l539
					}
					position++
					{
						position540, tokenIndex540 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l540
						}
						position++
						goto l541
					l540:
						position, tokenIndex = position540, tokenIndex540
					}
				l541:
					goto l538
				l539:
					position, tokenIndex = position538, tokenIndex538
					if buffer[position] != rune('i') {
						goto l536
					}
					position++
					if buffer[position] != rune('n') {
						goto l536
					}
					position++
				}
			l538:
				if !_rules[rule_]() {
					goto l536
				}
				add(ruleAuthorEx, position537)
			}
			return true
		l536:
			position, tokenIndex = position536, tokenIndex536
			return false
		},
		/* 72 AuthorEmend <- <('e' 'm' 'e' 'n' 'd' '.'? _)> */
		func() bool {
			position542, tokenIndex542 := position, tokenIndex
			{
				position543 := position
				if buffer[position] != rune('e') {
					goto l542
				}
				position++
				if buffer[position] != rune('m') {
					goto l542
				}
				position++
				if buffer[position] != rune('e') {
					goto l542
				}
				position++
				if buffer[position] != rune('n') {
					goto l542
				}
				position++
				if buffer[position] != rune('d') {
					goto l542
				}
				position++
				{
					position544, tokenIndex544 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l544
					}
					position++
					goto l545
				l544:
					position, tokenIndex = position544, tokenIndex544
				}
			l545:
				if !_rules[rule_]() {
					goto l542
				}
				add(ruleAuthorEmend, position543)
			}
			return true
		l542:
			position, tokenIndex = position542, tokenIndex542
			return false
		},
		/* 73 Author <- <(Author1 / Author2 / UnknownAuthor)> */
		func() bool {
			position546, tokenIndex546 := position, tokenIndex
			{
				position547 := position
				{
					position548, tokenIndex548 := position, tokenIndex
					if !_rules[ruleAuthor1]() {
						goto l549
					}
					goto l548
				l549:
					position, tokenIndex = position548, tokenIndex548
					if !_rules[ruleAuthor2]() {
						goto l550
					}
					goto l548
				l550:
					position, tokenIndex = position548, tokenIndex548
					if !_rules[ruleUnknownAuthor]() {
						goto l546
					}
				}
			l548:
				add(ruleAuthor, position547)
			}
			return true
		l546:
			position, tokenIndex = position546, tokenIndex546
			return false
		},
		/* 74 Author1 <- <(Author2 _? Filius)> */
		func() bool {
			position551, tokenIndex551 := position, tokenIndex
			{
				position552 := position
				if !_rules[ruleAuthor2]() {
					goto l551
				}
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
				if !_rules[ruleFilius]() {
					goto l551
				}
				add(ruleAuthor1, position552)
			}
			return true
		l551:
			position, tokenIndex = position551, tokenIndex551
			return false
		},
		/* 75 Author2 <- <(AuthorWord (_? AuthorWord)*)> */
		func() bool {
			position555, tokenIndex555 := position, tokenIndex
			{
				position556 := position
				if !_rules[ruleAuthorWord]() {
					goto l555
				}
			l557:
				{
					position558, tokenIndex558 := position, tokenIndex
					{
						position559, tokenIndex559 := position, tokenIndex
						if !_rules[rule_]() {
							goto l559
						}
						goto l560
					l559:
						position, tokenIndex = position559, tokenIndex559
					}
				l560:
					if !_rules[ruleAuthorWord]() {
						goto l558
					}
					goto l557
				l558:
					position, tokenIndex = position558, tokenIndex558
				}
				add(ruleAuthor2, position556)
			}
			return true
		l555:
			position, tokenIndex = position555, tokenIndex555
			return false
		},
		/* 76 UnknownAuthor <- <('?' / ((('a' 'u' 'c' 't') / ('a' 'n' 'o' 'n')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position561, tokenIndex561 := position, tokenIndex
			{
				position562 := position
				{
					position563, tokenIndex563 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l564
					}
					position++
					goto l563
				l564:
					position, tokenIndex = position563, tokenIndex563
					{
						position565, tokenIndex565 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l566
						}
						position++
						if buffer[position] != rune('u') {
							goto l566
						}
						position++
						if buffer[position] != rune('c') {
							goto l566
						}
						position++
						if buffer[position] != rune('t') {
							goto l566
						}
						position++
						goto l565
					l566:
						position, tokenIndex = position565, tokenIndex565
						if buffer[position] != rune('a') {
							goto l561
						}
						position++
						if buffer[position] != rune('n') {
							goto l561
						}
						position++
						if buffer[position] != rune('o') {
							goto l561
						}
						position++
						if buffer[position] != rune('n') {
							goto l561
						}
						position++
					}
				l565:
					{
						position567, tokenIndex567 := position, tokenIndex
						{
							position569, tokenIndex569 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l568
							}
							position, tokenIndex = position569, tokenIndex569
						}
						goto l567
					l568:
						position, tokenIndex = position567, tokenIndex567
						if buffer[position] != rune('.') {
							goto l561
						}
						position++
					}
				l567:
				}
			l563:
				add(ruleUnknownAuthor, position562)
			}
			return true
		l561:
			position, tokenIndex = position561, tokenIndex561
			return false
		},
		/* 77 AuthorWord <- <(!(('b' / 'B') ('o' / 'O') ('l' / 'L') ('d' / 'D') ':') (AuthorEtAl / AuthorWord2 / AuthorWord3 / AuthorPrefix))> */
		func() bool {
			position570, tokenIndex570 := position, tokenIndex
			{
				position571 := position
				{
					position572, tokenIndex572 := position, tokenIndex
					{
						position573, tokenIndex573 := position, tokenIndex
						if buffer[position] != rune('b') {
							goto l574
						}
						position++
						goto l573
					l574:
						position, tokenIndex = position573, tokenIndex573
						if buffer[position] != rune('B') {
							goto l572
						}
						position++
					}
				l573:
					{
						position575, tokenIndex575 := position, tokenIndex
						if buffer[position] != rune('o') {
							goto l576
						}
						position++
						goto l575
					l576:
						position, tokenIndex = position575, tokenIndex575
						if buffer[position] != rune('O') {
							goto l572
						}
						position++
					}
				l575:
					{
						position577, tokenIndex577 := position, tokenIndex
						if buffer[position] != rune('l') {
							goto l578
						}
						position++
						goto l577
					l578:
						position, tokenIndex = position577, tokenIndex577
						if buffer[position] != rune('L') {
							goto l572
						}
						position++
					}
				l577:
					{
						position579, tokenIndex579 := position, tokenIndex
						if buffer[position] != rune('d') {
							goto l580
						}
						position++
						goto l579
					l580:
						position, tokenIndex = position579, tokenIndex579
						if buffer[position] != rune('D') {
							goto l572
						}
						position++
					}
				l579:
					if buffer[position] != rune(':') {
						goto l572
					}
					position++
					goto l570
				l572:
					position, tokenIndex = position572, tokenIndex572
				}
				{
					position581, tokenIndex581 := position, tokenIndex
					if !_rules[ruleAuthorEtAl]() {
						goto l582
					}
					goto l581
				l582:
					position, tokenIndex = position581, tokenIndex581
					if !_rules[ruleAuthorWord2]() {
						goto l583
					}
					goto l581
				l583:
					position, tokenIndex = position581, tokenIndex581
					if !_rules[ruleAuthorWord3]() {
						goto l584
					}
					goto l581
				l584:
					position, tokenIndex = position581, tokenIndex581
					if !_rules[ruleAuthorPrefix]() {
						goto l570
					}
				}
			l581:
				add(ruleAuthorWord, position571)
			}
			return true
		l570:
			position, tokenIndex = position570, tokenIndex570
			return false
		},
		/* 78 AuthorEtAl <- <(('a' 'r' 'g' '.') / ('e' 't' ' ' 'a' 'l' '.' '{' '?' '}') / ((('e' 't') / '&') (' ' 'a' 'l') '.'?))> */
		func() bool {
			position585, tokenIndex585 := position, tokenIndex
			{
				position586 := position
				{
					position587, tokenIndex587 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l588
					}
					position++
					if buffer[position] != rune('r') {
						goto l588
					}
					position++
					if buffer[position] != rune('g') {
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
					if buffer[position] != rune('e') {
						goto l589
					}
					position++
					if buffer[position] != rune('t') {
						goto l589
					}
					position++
					if buffer[position] != rune(' ') {
						goto l589
					}
					position++
					if buffer[position] != rune('a') {
						goto l589
					}
					position++
					if buffer[position] != rune('l') {
						goto l589
					}
					position++
					if buffer[position] != rune('.') {
						goto l589
					}
					position++
					if buffer[position] != rune('{') {
						goto l589
					}
					position++
					if buffer[position] != rune('?') {
						goto l589
					}
					position++
					if buffer[position] != rune('}') {
						goto l589
					}
					position++
					goto l587
				l589:
					position, tokenIndex = position587, tokenIndex587
					{
						position590, tokenIndex590 := position, tokenIndex
						if buffer[position] != rune('e') {
							goto l591
						}
						position++
						if buffer[position] != rune('t') {
							goto l591
						}
						position++
						goto l590
					l591:
						position, tokenIndex = position590, tokenIndex590
						if buffer[position] != rune('&') {
							goto l585
						}
						position++
					}
				l590:
					if buffer[position] != rune(' ') {
						goto l585
					}
					position++
					if buffer[position] != rune('a') {
						goto l585
					}
					position++
					if buffer[position] != rune('l') {
						goto l585
					}
					position++
					{
						position592, tokenIndex592 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l592
						}
						position++
						goto l593
					l592:
						position, tokenIndex = position592, tokenIndex592
					}
				l593:
				}
			l587:
				add(ruleAuthorEtAl, position586)
			}
			return true
		l585:
			position, tokenIndex = position585, tokenIndex585
			return false
		},
		/* 79 AuthorWord2 <- <(AuthorWord3 Dash AuthorWordSoft)> */
		func() bool {
			position594, tokenIndex594 := position, tokenIndex
			{
				position595 := position
				if !_rules[ruleAuthorWord3]() {
					goto l594
				}
				if !_rules[ruleDash]() {
					goto l594
				}
				if !_rules[ruleAuthorWordSoft]() {
					goto l594
				}
				add(ruleAuthorWord2, position595)
			}
			return true
		l594:
			position, tokenIndex = position594, tokenIndex594
			return false
		},
		/* 80 AuthorWord3 <- <(AuthorPrefixGlued? (AllCapsAuthorWord / CapAuthorWord) '.'?)> */
		func() bool {
			position596, tokenIndex596 := position, tokenIndex
			{
				position597 := position
				{
					position598, tokenIndex598 := position, tokenIndex
					if !_rules[ruleAuthorPrefixGlued]() {
						goto l598
					}
					goto l599
				l598:
					position, tokenIndex = position598, tokenIndex598
				}
			l599:
				{
					position600, tokenIndex600 := position, tokenIndex
					if !_rules[ruleAllCapsAuthorWord]() {
						goto l601
					}
					goto l600
				l601:
					position, tokenIndex = position600, tokenIndex600
					if !_rules[ruleCapAuthorWord]() {
						goto l596
					}
				}
			l600:
				{
					position602, tokenIndex602 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l602
					}
					position++
					goto l603
				l602:
					position, tokenIndex = position602, tokenIndex602
				}
			l603:
				add(ruleAuthorWord3, position597)
			}
			return true
		l596:
			position, tokenIndex = position596, tokenIndex596
			return false
		},
		/* 81 AuthorWordSoft <- <(((AuthorUpperChar (AuthorUpperChar+ / AuthorLowerChar+)) / AuthorLowerChar+) '.'?)> */
		func() bool {
			position604, tokenIndex604 := position, tokenIndex
			{
				position605 := position
				{
					position606, tokenIndex606 := position, tokenIndex
					if !_rules[ruleAuthorUpperChar]() {
						goto l607
					}
					{
						position608, tokenIndex608 := position, tokenIndex
						if !_rules[ruleAuthorUpperChar]() {
							goto l609
						}
					l610:
						{
							position611, tokenIndex611 := position, tokenIndex
							if !_rules[ruleAuthorUpperChar]() {
								goto l611
							}
							goto l610
						l611:
							position, tokenIndex = position611, tokenIndex611
						}
						goto l608
					l609:
						position, tokenIndex = position608, tokenIndex608
						if !_rules[ruleAuthorLowerChar]() {
							goto l607
						}
					l612:
						{
							position613, tokenIndex613 := position, tokenIndex
							if !_rules[ruleAuthorLowerChar]() {
								goto l613
							}
							goto l612
						l613:
							position, tokenIndex = position613, tokenIndex613
						}
					}
				l608:
					goto l606
				l607:
					position, tokenIndex = position606, tokenIndex606
					if !_rules[ruleAuthorLowerChar]() {
						goto l604
					}
				l614:
					{
						position615, tokenIndex615 := position, tokenIndex
						if !_rules[ruleAuthorLowerChar]() {
							goto l615
						}
						goto l614
					l615:
						position, tokenIndex = position615, tokenIndex615
					}
				}
			l606:
				{
					position616, tokenIndex616 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l616
					}
					position++
					goto l617
				l616:
					position, tokenIndex = position616, tokenIndex616
				}
			l617:
				add(ruleAuthorWordSoft, position605)
			}
			return true
		l604:
			position, tokenIndex = position604, tokenIndex604
			return false
		},
		/* 82 CapAuthorWord <- <(AuthorUpperChar AuthorLowerChar*)> */
		func() bool {
			position618, tokenIndex618 := position, tokenIndex
			{
				position619 := position
				if !_rules[ruleAuthorUpperChar]() {
					goto l618
				}
			l620:
				{
					position621, tokenIndex621 := position, tokenIndex
					if !_rules[ruleAuthorLowerChar]() {
						goto l621
					}
					goto l620
				l621:
					position, tokenIndex = position621, tokenIndex621
				}
				add(ruleCapAuthorWord, position619)
			}
			return true
		l618:
			position, tokenIndex = position618, tokenIndex618
			return false
		},
		/* 83 AllCapsAuthorWord <- <(AuthorUpperChar AuthorUpperChar+)> */
		func() bool {
			position622, tokenIndex622 := position, tokenIndex
			{
				position623 := position
				if !_rules[ruleAuthorUpperChar]() {
					goto l622
				}
				if !_rules[ruleAuthorUpperChar]() {
					goto l622
				}
			l624:
				{
					position625, tokenIndex625 := position, tokenIndex
					if !_rules[ruleAuthorUpperChar]() {
						goto l625
					}
					goto l624
				l625:
					position, tokenIndex = position625, tokenIndex625
				}
				add(ruleAllCapsAuthorWord, position623)
			}
			return true
		l622:
			position, tokenIndex = position622, tokenIndex622
			return false
		},
		/* 84 Filius <- <(('f' '.') / ('f' 'i' 'l' '.') / ('f' 'i' 'l' 'i' 'u' 's'))> */
		func() bool {
			position626, tokenIndex626 := position, tokenIndex
			{
				position627 := position
				{
					position628, tokenIndex628 := position, tokenIndex
					if buffer[position] != rune('f') {
						goto l629
					}
					position++
					if buffer[position] != rune('.') {
						goto l629
					}
					position++
					goto l628
				l629:
					position, tokenIndex = position628, tokenIndex628
					if buffer[position] != rune('f') {
						goto l630
					}
					position++
					if buffer[position] != rune('i') {
						goto l630
					}
					position++
					if buffer[position] != rune('l') {
						goto l630
					}
					position++
					if buffer[position] != rune('.') {
						goto l630
					}
					position++
					goto l628
				l630:
					position, tokenIndex = position628, tokenIndex628
					if buffer[position] != rune('f') {
						goto l626
					}
					position++
					if buffer[position] != rune('i') {
						goto l626
					}
					position++
					if buffer[position] != rune('l') {
						goto l626
					}
					position++
					if buffer[position] != rune('i') {
						goto l626
					}
					position++
					if buffer[position] != rune('u') {
						goto l626
					}
					position++
					if buffer[position] != rune('s') {
						goto l626
					}
					position++
				}
			l628:
				add(ruleFilius, position627)
			}
			return true
		l626:
			position, tokenIndex = position626, tokenIndex626
			return false
		},
		/* 85 AuthorPrefixGlued <- <(('d' / 'O' / 'L') Apostrophe)> */
		func() bool {
			position631, tokenIndex631 := position, tokenIndex
			{
				position632 := position
				{
					position633, tokenIndex633 := position, tokenIndex
					if buffer[position] != rune('d') {
						goto l634
					}
					position++
					goto l633
				l634:
					position, tokenIndex = position633, tokenIndex633
					if buffer[position] != rune('O') {
						goto l635
					}
					position++
					goto l633
				l635:
					position, tokenIndex = position633, tokenIndex633
					if buffer[position] != rune('L') {
						goto l631
					}
					position++
				}
			l633:
				if !_rules[ruleApostrophe]() {
					goto l631
				}
				add(ruleAuthorPrefixGlued, position632)
			}
			return true
		l631:
			position, tokenIndex = position631, tokenIndex631
			return false
		},
		/* 86 AuthorPrefix <- <(AuthorPrefix1 / AuthorPrefix2)> */
		func() bool {
			position636, tokenIndex636 := position, tokenIndex
			{
				position637 := position
				{
					position638, tokenIndex638 := position, tokenIndex
					if !_rules[ruleAuthorPrefix1]() {
						goto l639
					}
					goto l638
				l639:
					position, tokenIndex = position638, tokenIndex638
					if !_rules[ruleAuthorPrefix2]() {
						goto l636
					}
				}
			l638:
				add(ruleAuthorPrefix, position637)
			}
			return true
		l636:
			position, tokenIndex = position636, tokenIndex636
			return false
		},
		/* 87 AuthorPrefix2 <- <(('v' '.' (_? ('d' '.'))?) / (Apostrophe 't'))> */
		func() bool {
			position640, tokenIndex640 := position, tokenIndex
			{
				position641 := position
				{
					position642, tokenIndex642 := position, tokenIndex
					if buffer[position] != rune('v') {
						goto l643
					}
					position++
					if buffer[position] != rune('.') {
						goto l643
					}
					position++
					{
						position644, tokenIndex644 := position, tokenIndex
						{
							position646, tokenIndex646 := position, tokenIndex
							if !_rules[rule_]() {
								goto l646
							}
							goto l647
						l646:
							position, tokenIndex = position646, tokenIndex646
						}
					l647:
						if buffer[position] != rune('d') {
							goto l644
						}
						position++
						if buffer[position] != rune('.') {
							goto l644
						}
						position++
						goto l645
					l644:
						position, tokenIndex = position644, tokenIndex644
					}
				l645:
					goto l642
				l643:
					position, tokenIndex = position642, tokenIndex642
					if !_rules[ruleApostrophe]() {
						goto l640
					}
					if buffer[position] != rune('t') {
						goto l640
					}
					position++
				}
			l642:
				add(ruleAuthorPrefix2, position641)
			}
			return true
		l640:
			position, tokenIndex = position640, tokenIndex640
			return false
		},
		/* 88 AuthorPrefix1 <- <((('a' 'b') / ('a' 'f') / ('b' 'i' 's') / ('d' 'a') / ('d' 'e' 'r') / ('d' 'e' 's') / ('d' 'e' 'n') / ('d' 'e' 'l') / ('d' 'e' 'l' 'l' 'a') / ('d' 'e' 'l' 'a') / ('d' 'e') / ('d' 'i') / ('d' 'u') / ('e' 'l') / ('l' 'a') / ('l' 'e') / ('t' 'e' 'r') / ('v' 'a' 'n') / ('d' Apostrophe) / ('i' 'n' Apostrophe 't') / ('z' 'u' 'r') / ('v' 'o' 'n' (_ (('d' '.') / ('d' 'e' 'm')))?) / ('v' (_ 'd')?)) &_)> */
		func() bool {
			position648, tokenIndex648 := position, tokenIndex
			{
				position649 := position
				{
					position650, tokenIndex650 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l651
					}
					position++
					if buffer[position] != rune('b') {
						goto l651
					}
					position++
					goto l650
				l651:
					position, tokenIndex = position650, tokenIndex650
					if buffer[position] != rune('a') {
						goto l652
					}
					position++
					if buffer[position] != rune('f') {
						goto l652
					}
					position++
					goto l650
				l652:
					position, tokenIndex = position650, tokenIndex650
					if buffer[position] != rune('b') {
						goto l653
					}
					position++
					if buffer[position] != rune('i') {
						goto l653
					}
					position++
					if buffer[position] != rune('s') {
						goto l653
					}
					position++
					goto l650
				l653:
					position, tokenIndex = position650, tokenIndex650
					if buffer[position] != rune('d') {
						goto l654
					}
					position++
					if buffer[position] != rune('a') {
						goto l654
					}
					position++
					goto l650
				l654:
					position, tokenIndex = position650, tokenIndex650
					if buffer[position] != rune('d') {
						goto l655
					}
					position++
					if buffer[position] != rune('e') {
						goto l655
					}
					position++
					if buffer[position] != rune('r') {
						goto l655
					}
					position++
					goto l650
				l655:
					position, tokenIndex = position650, tokenIndex650
					if buffer[position] != rune('d') {
						goto l656
					}
					position++
					if buffer[position] != rune('e') {
						goto l656
					}
					position++
					if buffer[position] != rune('s') {
						goto l656
					}
					position++
					goto l650
				l656:
					position, tokenIndex = position650, tokenIndex650
					if buffer[position] != rune('d') {
						goto l657
					}
					position++
					if buffer[position] != rune('e') {
						goto l657
					}
					position++
					if buffer[position] != rune('n') {
						goto l657
					}
					position++
					goto l650
				l657:
					position, tokenIndex = position650, tokenIndex650
					if buffer[position] != rune('d') {
						goto l658
					}
					position++
					if buffer[position] != rune('e') {
						goto l658
					}
					position++
					if buffer[position] != rune('l') {
						goto l658
					}
					position++
					goto l650
				l658:
					position, tokenIndex = position650, tokenIndex650
					if buffer[position] != rune('d') {
						goto l659
					}
					position++
					if buffer[position] != rune('e') {
						goto l659
					}
					position++
					if buffer[position] != rune('l') {
						goto l659
					}
					position++
					if buffer[position] != rune('l') {
						goto l659
					}
					position++
					if buffer[position] != rune('a') {
						goto l659
					}
					position++
					goto l650
				l659:
					position, tokenIndex = position650, tokenIndex650
					if buffer[position] != rune('d') {
						goto l660
					}
					position++
					if buffer[position] != rune('e') {
						goto l660
					}
					position++
					if buffer[position] != rune('l') {
						goto l660
					}
					position++
					if buffer[position] != rune('a') {
						goto l660
					}
					position++
					goto l650
				l660:
					position, tokenIndex = position650, tokenIndex650
					if buffer[position] != rune('d') {
						goto l661
					}
					position++
					if buffer[position] != rune('e') {
						goto l661
					}
					position++
					goto l650
				l661:
					position, tokenIndex = position650, tokenIndex650
					if buffer[position] != rune('d') {
						goto l662
					}
					position++
					if buffer[position] != rune('i') {
						goto l662
					}
					position++
					goto l650
				l662:
					position, tokenIndex = position650, tokenIndex650
					if buffer[position] != rune('d') {
						goto l663
					}
					position++
					if buffer[position] != rune('u') {
						goto l663
					}
					position++
					goto l650
				l663:
					position, tokenIndex = position650, tokenIndex650
					if buffer[position] != rune('e') {
						goto l664
					}
					position++
					if buffer[position] != rune('l') {
						goto l664
					}
					position++
					goto l650
				l664:
					position, tokenIndex = position650, tokenIndex650
					if buffer[position] != rune('l') {
						goto l665
					}
					position++
					if buffer[position] != rune('a') {
						goto l665
					}
					position++
					goto l650
				l665:
					position, tokenIndex = position650, tokenIndex650
					if buffer[position] != rune('l') {
						goto l666
					}
					position++
					if buffer[position] != rune('e') {
						goto l666
					}
					position++
					goto l650
				l666:
					position, tokenIndex = position650, tokenIndex650
					if buffer[position] != rune('t') {
						goto l667
					}
					position++
					if buffer[position] != rune('e') {
						goto l667
					}
					position++
					if buffer[position] != rune('r') {
						goto l667
					}
					position++
					goto l650
				l667:
					position, tokenIndex = position650, tokenIndex650
					if buffer[position] != rune('v') {
						goto l668
					}
					position++
					if buffer[position] != rune('a') {
						goto l668
					}
					position++
					if buffer[position] != rune('n') {
						goto l668
					}
					position++
					goto l650
				l668:
					position, tokenIndex = position650, tokenIndex650
					if buffer[position] != rune('d') {
						goto l669
					}
					position++
					if !_rules[ruleApostrophe]() {
						goto l669
					}
					goto l650
				l669:
					position, tokenIndex = position650, tokenIndex650
					if buffer[position] != rune('i') {
						goto l670
					}
					position++
					if buffer[position] != rune('n') {
						goto l670
					}
					position++
					if !_rules[ruleApostrophe]() {
						goto l670
					}
					if buffer[position] != rune('t') {
						goto l670
					}
					position++
					goto l650
				l670:
					position, tokenIndex = position650, tokenIndex650
					if buffer[position] != rune('z') {
						goto l671
					}
					position++
					if buffer[position] != rune('u') {
						goto l671
					}
					position++
					if buffer[position] != rune('r') {
						goto l671
					}
					position++
					goto l650
				l671:
					position, tokenIndex = position650, tokenIndex650
					if buffer[position] != rune('v') {
						goto l672
					}
					position++
					if buffer[position] != rune('o') {
						goto l672
					}
					position++
					if buffer[position] != rune('n') {
						goto l672
					}
					position++
					{
						position673, tokenIndex673 := position, tokenIndex
						if !_rules[rule_]() {
							goto l673
						}
						{
							position675, tokenIndex675 := position, tokenIndex
							if buffer[position] != rune('d') {
								goto l676
							}
							position++
							if buffer[position] != rune('.') {
								goto l676
							}
							position++
							goto l675
						l676:
							position, tokenIndex = position675, tokenIndex675
							if buffer[position] != rune('d') {
								goto l673
							}
							position++
							if buffer[position] != rune('e') {
								goto l673
							}
							position++
							if buffer[position] != rune('m') {
								goto l673
							}
							position++
						}
					l675:
						goto l674
					l673:
						position, tokenIndex = position673, tokenIndex673
					}
				l674:
					goto l650
				l672:
					position, tokenIndex = position650, tokenIndex650
					if buffer[position] != rune('v') {
						goto l648
					}
					position++
					{
						position677, tokenIndex677 := position, tokenIndex
						if !_rules[rule_]() {
							goto l677
						}
						if buffer[position] != rune('d') {
							goto l677
						}
						position++
						goto l678
					l677:
						position, tokenIndex = position677, tokenIndex677
					}
				l678:
				}
			l650:
				{
					position679, tokenIndex679 := position, tokenIndex
					if !_rules[rule_]() {
						goto l648
					}
					position, tokenIndex = position679, tokenIndex679
				}
				add(ruleAuthorPrefix1, position649)
			}
			return true
		l648:
			position, tokenIndex = position648, tokenIndex648
			return false
		},
		/* 89 AuthorUpperChar <- <(UpperASCII / MiscodedChar / ('À' / 'Á' / 'Â' / 'Ã' / 'Ä' / 'Å' / 'Æ' / 'Ç' / 'È' / 'É' / 'Ê' / 'Ë' / 'Ì' / 'Í' / 'Î' / 'Ï' / 'Ð' / 'Ñ' / 'Ò' / 'Ó' / 'Ô' / 'Õ' / 'Ö' / 'Ø' / 'Ù' / 'Ú' / 'Û' / 'Ü' / 'Ý' / 'Ć' / 'Č' / 'Ď' / 'İ' / 'Ķ' / 'Ĺ' / 'ĺ' / 'Ľ' / 'ľ' / 'Ł' / 'ł' / 'Ņ' / 'Ō' / 'Ő' / 'Œ' / 'Ř' / 'Ś' / 'Ŝ' / 'Ş' / 'Š' / 'Ÿ' / 'Ź' / 'Ż' / 'Ž' / 'ƒ' / 'Ǿ' / 'Ș' / 'Ț'))> */
		func() bool {
			position680, tokenIndex680 := position, tokenIndex
			{
				position681 := position
				{
					position682, tokenIndex682 := position, tokenIndex
					if !_rules[ruleUpperASCII]() {
						goto l683
					}
					goto l682
				l683:
					position, tokenIndex = position682, tokenIndex682
					if !_rules[ruleMiscodedChar]() {
						goto l684
					}
					goto l682
				l684:
					position, tokenIndex = position682, tokenIndex682
					{
						position685, tokenIndex685 := position, tokenIndex
						if buffer[position] != rune('À') {
							goto l686
						}
						position++
						goto l685
					l686:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Á') {
							goto l687
						}
						position++
						goto l685
					l687:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Â') {
							goto l688
						}
						position++
						goto l685
					l688:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ã') {
							goto l689
						}
						position++
						goto l685
					l689:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ä') {
							goto l690
						}
						position++
						goto l685
					l690:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Å') {
							goto l691
						}
						position++
						goto l685
					l691:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Æ') {
							goto l692
						}
						position++
						goto l685
					l692:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ç') {
							goto l693
						}
						position++
						goto l685
					l693:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('È') {
							goto l694
						}
						position++
						goto l685
					l694:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('É') {
							goto l695
						}
						position++
						goto l685
					l695:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ê') {
							goto l696
						}
						position++
						goto l685
					l696:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ë') {
							goto l697
						}
						position++
						goto l685
					l697:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ì') {
							goto l698
						}
						position++
						goto l685
					l698:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Í') {
							goto l699
						}
						position++
						goto l685
					l699:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Î') {
							goto l700
						}
						position++
						goto l685
					l700:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ï') {
							goto l701
						}
						position++
						goto l685
					l701:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ð') {
							goto l702
						}
						position++
						goto l685
					l702:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ñ') {
							goto l703
						}
						position++
						goto l685
					l703:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ò') {
							goto l704
						}
						position++
						goto l685
					l704:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ó') {
							goto l705
						}
						position++
						goto l685
					l705:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ô') {
							goto l706
						}
						position++
						goto l685
					l706:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Õ') {
							goto l707
						}
						position++
						goto l685
					l707:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ö') {
							goto l708
						}
						position++
						goto l685
					l708:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ø') {
							goto l709
						}
						position++
						goto l685
					l709:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ù') {
							goto l710
						}
						position++
						goto l685
					l710:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ú') {
							goto l711
						}
						position++
						goto l685
					l711:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Û') {
							goto l712
						}
						position++
						goto l685
					l712:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ü') {
							goto l713
						}
						position++
						goto l685
					l713:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ý') {
							goto l714
						}
						position++
						goto l685
					l714:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ć') {
							goto l715
						}
						position++
						goto l685
					l715:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Č') {
							goto l716
						}
						position++
						goto l685
					l716:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ď') {
							goto l717
						}
						position++
						goto l685
					l717:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('İ') {
							goto l718
						}
						position++
						goto l685
					l718:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ķ') {
							goto l719
						}
						position++
						goto l685
					l719:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ĺ') {
							goto l720
						}
						position++
						goto l685
					l720:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('ĺ') {
							goto l721
						}
						position++
						goto l685
					l721:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ľ') {
							goto l722
						}
						position++
						goto l685
					l722:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('ľ') {
							goto l723
						}
						position++
						goto l685
					l723:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ł') {
							goto l724
						}
						position++
						goto l685
					l724:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('ł') {
							goto l725
						}
						position++
						goto l685
					l725:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ņ') {
							goto l726
						}
						position++
						goto l685
					l726:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ō') {
							goto l727
						}
						position++
						goto l685
					l727:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ő') {
							goto l728
						}
						position++
						goto l685
					l728:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Œ') {
							goto l729
						}
						position++
						goto l685
					l729:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ř') {
							goto l730
						}
						position++
						goto l685
					l730:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ś') {
							goto l731
						}
						position++
						goto l685
					l731:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ŝ') {
							goto l732
						}
						position++
						goto l685
					l732:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ş') {
							goto l733
						}
						position++
						goto l685
					l733:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Š') {
							goto l734
						}
						position++
						goto l685
					l734:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ÿ') {
							goto l735
						}
						position++
						goto l685
					l735:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ź') {
							goto l736
						}
						position++
						goto l685
					l736:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ż') {
							goto l737
						}
						position++
						goto l685
					l737:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ž') {
							goto l738
						}
						position++
						goto l685
					l738:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('ƒ') {
							goto l739
						}
						position++
						goto l685
					l739:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ǿ') {
							goto l740
						}
						position++
						goto l685
					l740:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ș') {
							goto l741
						}
						position++
						goto l685
					l741:
						position, tokenIndex = position685, tokenIndex685
						if buffer[position] != rune('Ț') {
							goto l680
						}
						position++
					}
				l685:
				}
			l682:
				add(ruleAuthorUpperChar, position681)
			}
			return true
		l680:
			position, tokenIndex = position680, tokenIndex680
			return false
		},
		/* 90 AuthorLowerChar <- <(LowerASCII / MiscodedChar / ('à' / 'á' / 'â' / 'ã' / 'ä' / 'å' / 'æ' / 'ç' / 'è' / 'é' / 'ê' / 'ë' / 'ì' / 'í' / 'î' / 'ï' / 'ð' / 'ñ' / 'ò' / 'ó' / 'ó' / 'ô' / 'õ' / 'ö' / 'ø' / 'ù' / 'ú' / 'û' / 'ü' / 'ý' / 'ÿ' / 'ā' / 'ă' / 'ą' / 'ć' / 'ĉ' / 'č' / 'ď' / 'đ' / '\'' / 'ē' / 'ĕ' / 'ė' / 'ę' / 'ě' / 'ğ' / 'ī' / 'ĭ' / 'İ' / 'ı' / 'ĺ' / 'ľ' / 'ł' / 'ń' / 'ņ' / 'ň' / 'ŏ' / 'ő' / 'œ' / 'ŕ' / 'ř' / 'ś' / 'ş' / 'š' / 'ţ' / 'ť' / 'ũ' / 'ū' / 'ŭ' / 'ů' / 'ű' / 'ź' / 'ż' / 'ž' / 'ſ' / 'ǎ' / 'ǔ' / 'ǧ' / 'ș' / 'ț' / 'ȳ' / 'ß'))> */
		func() bool {
			position742, tokenIndex742 := position, tokenIndex
			{
				position743 := position
				{
					position744, tokenIndex744 := position, tokenIndex
					if !_rules[ruleLowerASCII]() {
						goto l745
					}
					goto l744
				l745:
					position, tokenIndex = position744, tokenIndex744
					if !_rules[ruleMiscodedChar]() {
						goto l746
					}
					goto l744
				l746:
					position, tokenIndex = position744, tokenIndex744
					{
						position747, tokenIndex747 := position, tokenIndex
						if buffer[position] != rune('à') {
							goto l748
						}
						position++
						goto l747
					l748:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('á') {
							goto l749
						}
						position++
						goto l747
					l749:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('â') {
							goto l750
						}
						position++
						goto l747
					l750:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ã') {
							goto l751
						}
						position++
						goto l747
					l751:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ä') {
							goto l752
						}
						position++
						goto l747
					l752:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('å') {
							goto l753
						}
						position++
						goto l747
					l753:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('æ') {
							goto l754
						}
						position++
						goto l747
					l754:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ç') {
							goto l755
						}
						position++
						goto l747
					l755:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('è') {
							goto l756
						}
						position++
						goto l747
					l756:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('é') {
							goto l757
						}
						position++
						goto l747
					l757:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ê') {
							goto l758
						}
						position++
						goto l747
					l758:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ë') {
							goto l759
						}
						position++
						goto l747
					l759:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ì') {
							goto l760
						}
						position++
						goto l747
					l760:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('í') {
							goto l761
						}
						position++
						goto l747
					l761:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('î') {
							goto l762
						}
						position++
						goto l747
					l762:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ï') {
							goto l763
						}
						position++
						goto l747
					l763:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ð') {
							goto l764
						}
						position++
						goto l747
					l764:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ñ') {
							goto l765
						}
						position++
						goto l747
					l765:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ò') {
							goto l766
						}
						position++
						goto l747
					l766:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ó') {
							goto l767
						}
						position++
						goto l747
					l767:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ó') {
							goto l768
						}
						position++
						goto l747
					l768:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ô') {
							goto l769
						}
						position++
						goto l747
					l769:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('õ') {
							goto l770
						}
						position++
						goto l747
					l770:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ö') {
							goto l771
						}
						position++
						goto l747
					l771:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ø') {
							goto l772
						}
						position++
						goto l747
					l772:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ù') {
							goto l773
						}
						position++
						goto l747
					l773:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ú') {
							goto l774
						}
						position++
						goto l747
					l774:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('û') {
							goto l775
						}
						position++
						goto l747
					l775:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ü') {
							goto l776
						}
						position++
						goto l747
					l776:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ý') {
							goto l777
						}
						position++
						goto l747
					l777:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ÿ') {
							goto l778
						}
						position++
						goto l747
					l778:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ā') {
							goto l779
						}
						position++
						goto l747
					l779:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ă') {
							goto l780
						}
						position++
						goto l747
					l780:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ą') {
							goto l781
						}
						position++
						goto l747
					l781:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ć') {
							goto l782
						}
						position++
						goto l747
					l782:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ĉ') {
							goto l783
						}
						position++
						goto l747
					l783:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('č') {
							goto l784
						}
						position++
						goto l747
					l784:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ď') {
							goto l785
						}
						position++
						goto l747
					l785:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('đ') {
							goto l786
						}
						position++
						goto l747
					l786:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('\'') {
							goto l787
						}
						position++
						goto l747
					l787:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ē') {
							goto l788
						}
						position++
						goto l747
					l788:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ĕ') {
							goto l789
						}
						position++
						goto l747
					l789:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ė') {
							goto l790
						}
						position++
						goto l747
					l790:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ę') {
							goto l791
						}
						position++
						goto l747
					l791:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ě') {
							goto l792
						}
						position++
						goto l747
					l792:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ğ') {
							goto l793
						}
						position++
						goto l747
					l793:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ī') {
							goto l794
						}
						position++
						goto l747
					l794:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ĭ') {
							goto l795
						}
						position++
						goto l747
					l795:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('İ') {
							goto l796
						}
						position++
						goto l747
					l796:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ı') {
							goto l797
						}
						position++
						goto l747
					l797:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ĺ') {
							goto l798
						}
						position++
						goto l747
					l798:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ľ') {
							goto l799
						}
						position++
						goto l747
					l799:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ł') {
							goto l800
						}
						position++
						goto l747
					l800:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ń') {
							goto l801
						}
						position++
						goto l747
					l801:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ņ') {
							goto l802
						}
						position++
						goto l747
					l802:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ň') {
							goto l803
						}
						position++
						goto l747
					l803:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ŏ') {
							goto l804
						}
						position++
						goto l747
					l804:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ő') {
							goto l805
						}
						position++
						goto l747
					l805:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('œ') {
							goto l806
						}
						position++
						goto l747
					l806:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ŕ') {
							goto l807
						}
						position++
						goto l747
					l807:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ř') {
							goto l808
						}
						position++
						goto l747
					l808:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ś') {
							goto l809
						}
						position++
						goto l747
					l809:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ş') {
							goto l810
						}
						position++
						goto l747
					l810:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('š') {
							goto l811
						}
						position++
						goto l747
					l811:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ţ') {
							goto l812
						}
						position++
						goto l747
					l812:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ť') {
							goto l813
						}
						position++
						goto l747
					l813:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ũ') {
							goto l814
						}
						position++
						goto l747
					l814:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ū') {
							goto l815
						}
						position++
						goto l747
					l815:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ŭ') {
							goto l816
						}
						position++
						goto l747
					l816:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ů') {
							goto l817
						}
						position++
						goto l747
					l817:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ű') {
							goto l818
						}
						position++
						goto l747
					l818:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ź') {
							goto l819
						}
						position++
						goto l747
					l819:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ż') {
							goto l820
						}
						position++
						goto l747
					l820:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ž') {
							goto l821
						}
						position++
						goto l747
					l821:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ſ') {
							goto l822
						}
						position++
						goto l747
					l822:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ǎ') {
							goto l823
						}
						position++
						goto l747
					l823:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ǔ') {
							goto l824
						}
						position++
						goto l747
					l824:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ǧ') {
							goto l825
						}
						position++
						goto l747
					l825:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ș') {
							goto l826
						}
						position++
						goto l747
					l826:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ț') {
							goto l827
						}
						position++
						goto l747
					l827:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ȳ') {
							goto l828
						}
						position++
						goto l747
					l828:
						position, tokenIndex = position747, tokenIndex747
						if buffer[position] != rune('ß') {
							goto l742
						}
						position++
					}
				l747:
				}
			l744:
				add(ruleAuthorLowerChar, position743)
			}
			return true
		l742:
			position, tokenIndex = position742, tokenIndex742
			return false
		},
		/* 91 Year <- <(YearRange / YearApprox / YearWithParens / YearWithPage / YearWithDot / YearWithChar / YearNum)> */
		func() bool {
			position829, tokenIndex829 := position, tokenIndex
			{
				position830 := position
				{
					position831, tokenIndex831 := position, tokenIndex
					if !_rules[ruleYearRange]() {
						goto l832
					}
					goto l831
				l832:
					position, tokenIndex = position831, tokenIndex831
					if !_rules[ruleYearApprox]() {
						goto l833
					}
					goto l831
				l833:
					position, tokenIndex = position831, tokenIndex831
					if !_rules[ruleYearWithParens]() {
						goto l834
					}
					goto l831
				l834:
					position, tokenIndex = position831, tokenIndex831
					if !_rules[ruleYearWithPage]() {
						goto l835
					}
					goto l831
				l835:
					position, tokenIndex = position831, tokenIndex831
					if !_rules[ruleYearWithDot]() {
						goto l836
					}
					goto l831
				l836:
					position, tokenIndex = position831, tokenIndex831
					if !_rules[ruleYearWithChar]() {
						goto l837
					}
					goto l831
				l837:
					position, tokenIndex = position831, tokenIndex831
					if !_rules[ruleYearNum]() {
						goto l829
					}
				}
			l831:
				add(ruleYear, position830)
			}
			return true
		l829:
			position, tokenIndex = position829, tokenIndex829
			return false
		},
		/* 92 YearRange <- <(YearNum Dash (Nums+ ('a' / 'b' / 'c' / 'd' / 'e' / 'f' / 'g' / 'h' / 'i' / 'j' / 'k' / 'l' / 'm' / 'n' / 'o' / 'p' / 'q' / 'r' / 's' / 't' / 'u' / 'v' / 'w' / 'x' / 'y' / 'z' / '?')*))> */
		func() bool {
			position838, tokenIndex838 := position, tokenIndex
			{
				position839 := position
				if !_rules[ruleYearNum]() {
					goto l838
				}
				if !_rules[ruleDash]() {
					goto l838
				}
				if !_rules[ruleNums]() {
					goto l838
				}
			l840:
				{
					position841, tokenIndex841 := position, tokenIndex
					if !_rules[ruleNums]() {
						goto l841
					}
					goto l840
				l841:
					position, tokenIndex = position841, tokenIndex841
				}
			l842:
				{
					position843, tokenIndex843 := position, tokenIndex
					{
						position844, tokenIndex844 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l845
						}
						position++
						goto l844
					l845:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('b') {
							goto l846
						}
						position++
						goto l844
					l846:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('c') {
							goto l847
						}
						position++
						goto l844
					l847:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('d') {
							goto l848
						}
						position++
						goto l844
					l848:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('e') {
							goto l849
						}
						position++
						goto l844
					l849:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('f') {
							goto l850
						}
						position++
						goto l844
					l850:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('g') {
							goto l851
						}
						position++
						goto l844
					l851:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('h') {
							goto l852
						}
						position++
						goto l844
					l852:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('i') {
							goto l853
						}
						position++
						goto l844
					l853:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('j') {
							goto l854
						}
						position++
						goto l844
					l854:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('k') {
							goto l855
						}
						position++
						goto l844
					l855:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('l') {
							goto l856
						}
						position++
						goto l844
					l856:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('m') {
							goto l857
						}
						position++
						goto l844
					l857:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('n') {
							goto l858
						}
						position++
						goto l844
					l858:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('o') {
							goto l859
						}
						position++
						goto l844
					l859:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('p') {
							goto l860
						}
						position++
						goto l844
					l860:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('q') {
							goto l861
						}
						position++
						goto l844
					l861:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('r') {
							goto l862
						}
						position++
						goto l844
					l862:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('s') {
							goto l863
						}
						position++
						goto l844
					l863:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('t') {
							goto l864
						}
						position++
						goto l844
					l864:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('u') {
							goto l865
						}
						position++
						goto l844
					l865:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('v') {
							goto l866
						}
						position++
						goto l844
					l866:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('w') {
							goto l867
						}
						position++
						goto l844
					l867:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('x') {
							goto l868
						}
						position++
						goto l844
					l868:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('y') {
							goto l869
						}
						position++
						goto l844
					l869:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('z') {
							goto l870
						}
						position++
						goto l844
					l870:
						position, tokenIndex = position844, tokenIndex844
						if buffer[position] != rune('?') {
							goto l843
						}
						position++
					}
				l844:
					goto l842
				l843:
					position, tokenIndex = position843, tokenIndex843
				}
				add(ruleYearRange, position839)
			}
			return true
		l838:
			position, tokenIndex = position838, tokenIndex838
			return false
		},
		/* 93 YearWithDot <- <(YearNum '.')> */
		func() bool {
			position871, tokenIndex871 := position, tokenIndex
			{
				position872 := position
				if !_rules[ruleYearNum]() {
					goto l871
				}
				if buffer[position] != rune('.') {
					goto l871
				}
				position++
				add(ruleYearWithDot, position872)
			}
			return true
		l871:
			position, tokenIndex = position871, tokenIndex871
			return false
		},
		/* 94 YearApprox <- <('[' _? YearNum _? ']')> */
		func() bool {
			position873, tokenIndex873 := position, tokenIndex
			{
				position874 := position
				if buffer[position] != rune('[') {
					goto l873
				}
				position++
				{
					position875, tokenIndex875 := position, tokenIndex
					if !_rules[rule_]() {
						goto l875
					}
					goto l876
				l875:
					position, tokenIndex = position875, tokenIndex875
				}
			l876:
				if !_rules[ruleYearNum]() {
					goto l873
				}
				{
					position877, tokenIndex877 := position, tokenIndex
					if !_rules[rule_]() {
						goto l877
					}
					goto l878
				l877:
					position, tokenIndex = position877, tokenIndex877
				}
			l878:
				if buffer[position] != rune(']') {
					goto l873
				}
				position++
				add(ruleYearApprox, position874)
			}
			return true
		l873:
			position, tokenIndex = position873, tokenIndex873
			return false
		},
		/* 95 YearWithPage <- <((YearWithChar / YearNum) _? ':' _? Nums+)> */
		func() bool {
			position879, tokenIndex879 := position, tokenIndex
			{
				position880 := position
				{
					position881, tokenIndex881 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l882
					}
					goto l881
				l882:
					position, tokenIndex = position881, tokenIndex881
					if !_rules[ruleYearNum]() {
						goto l879
					}
				}
			l881:
				{
					position883, tokenIndex883 := position, tokenIndex
					if !_rules[rule_]() {
						goto l883
					}
					goto l884
				l883:
					position, tokenIndex = position883, tokenIndex883
				}
			l884:
				if buffer[position] != rune(':') {
					goto l879
				}
				position++
				{
					position885, tokenIndex885 := position, tokenIndex
					if !_rules[rule_]() {
						goto l885
					}
					goto l886
				l885:
					position, tokenIndex = position885, tokenIndex885
				}
			l886:
				if !_rules[ruleNums]() {
					goto l879
				}
			l887:
				{
					position888, tokenIndex888 := position, tokenIndex
					if !_rules[ruleNums]() {
						goto l888
					}
					goto l887
				l888:
					position, tokenIndex = position888, tokenIndex888
				}
				add(ruleYearWithPage, position880)
			}
			return true
		l879:
			position, tokenIndex = position879, tokenIndex879
			return false
		},
		/* 96 YearWithParens <- <('(' (YearWithChar / YearNum) ')')> */
		func() bool {
			position889, tokenIndex889 := position, tokenIndex
			{
				position890 := position
				if buffer[position] != rune('(') {
					goto l889
				}
				position++
				{
					position891, tokenIndex891 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l892
					}
					goto l891
				l892:
					position, tokenIndex = position891, tokenIndex891
					if !_rules[ruleYearNum]() {
						goto l889
					}
				}
			l891:
				if buffer[position] != rune(')') {
					goto l889
				}
				position++
				add(ruleYearWithParens, position890)
			}
			return true
		l889:
			position, tokenIndex = position889, tokenIndex889
			return false
		},
		/* 97 YearWithChar <- <(YearNum LowerASCII Action0)> */
		func() bool {
			position893, tokenIndex893 := position, tokenIndex
			{
				position894 := position
				if !_rules[ruleYearNum]() {
					goto l893
				}
				if !_rules[ruleLowerASCII]() {
					goto l893
				}
				if !_rules[ruleAction0]() {
					goto l893
				}
				add(ruleYearWithChar, position894)
			}
			return true
		l893:
			position, tokenIndex = position893, tokenIndex893
			return false
		},
		/* 98 YearNum <- <(('1' / '2') ('0' / '7' / '8' / '9') Nums (Nums / '?') '?'*)> */
		func() bool {
			position895, tokenIndex895 := position, tokenIndex
			{
				position896 := position
				{
					position897, tokenIndex897 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l898
					}
					position++
					goto l897
				l898:
					position, tokenIndex = position897, tokenIndex897
					if buffer[position] != rune('2') {
						goto l895
					}
					position++
				}
			l897:
				{
					position899, tokenIndex899 := position, tokenIndex
					if buffer[position] != rune('0') {
						goto l900
					}
					position++
					goto l899
				l900:
					position, tokenIndex = position899, tokenIndex899
					if buffer[position] != rune('7') {
						goto l901
					}
					position++
					goto l899
				l901:
					position, tokenIndex = position899, tokenIndex899
					if buffer[position] != rune('8') {
						goto l902
					}
					position++
					goto l899
				l902:
					position, tokenIndex = position899, tokenIndex899
					if buffer[position] != rune('9') {
						goto l895
					}
					position++
				}
			l899:
				if !_rules[ruleNums]() {
					goto l895
				}
				{
					position903, tokenIndex903 := position, tokenIndex
					if !_rules[ruleNums]() {
						goto l904
					}
					goto l903
				l904:
					position, tokenIndex = position903, tokenIndex903
					if buffer[position] != rune('?') {
						goto l895
					}
					position++
				}
			l903:
			l905:
				{
					position906, tokenIndex906 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l906
					}
					position++
					goto l905
				l906:
					position, tokenIndex = position906, tokenIndex906
				}
				add(ruleYearNum, position896)
			}
			return true
		l895:
			position, tokenIndex = position895, tokenIndex895
			return false
		},
		/* 99 NameUpperChar <- <(UpperChar / UpperCharExtended)> */
		func() bool {
			position907, tokenIndex907 := position, tokenIndex
			{
				position908 := position
				{
					position909, tokenIndex909 := position, tokenIndex
					if !_rules[ruleUpperChar]() {
						goto l910
					}
					goto l909
				l910:
					position, tokenIndex = position909, tokenIndex909
					if !_rules[ruleUpperCharExtended]() {
						goto l907
					}
				}
			l909:
				add(ruleNameUpperChar, position908)
			}
			return true
		l907:
			position, tokenIndex = position907, tokenIndex907
			return false
		},
		/* 100 UpperCharExtended <- <('Æ' / 'Œ' / 'Ö')> */
		func() bool {
			position911, tokenIndex911 := position, tokenIndex
			{
				position912 := position
				{
					position913, tokenIndex913 := position, tokenIndex
					if buffer[position] != rune('Æ') {
						goto l914
					}
					position++
					goto l913
				l914:
					position, tokenIndex = position913, tokenIndex913
					if buffer[position] != rune('Œ') {
						goto l915
					}
					position++
					goto l913
				l915:
					position, tokenIndex = position913, tokenIndex913
					if buffer[position] != rune('Ö') {
						goto l911
					}
					position++
				}
			l913:
				add(ruleUpperCharExtended, position912)
			}
			return true
		l911:
			position, tokenIndex = position911, tokenIndex911
			return false
		},
		/* 101 UpperChar <- <UpperASCII> */
		func() bool {
			position916, tokenIndex916 := position, tokenIndex
			{
				position917 := position
				if !_rules[ruleUpperASCII]() {
					goto l916
				}
				add(ruleUpperChar, position917)
			}
			return true
		l916:
			position, tokenIndex = position916, tokenIndex916
			return false
		},
		/* 102 NameLowerChar <- <(LowerChar / LowerCharExtended / MiscodedChar)> */
		func() bool {
			position918, tokenIndex918 := position, tokenIndex
			{
				position919 := position
				{
					position920, tokenIndex920 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l921
					}
					goto l920
				l921:
					position, tokenIndex = position920, tokenIndex920
					if !_rules[ruleLowerCharExtended]() {
						goto l922
					}
					goto l920
				l922:
					position, tokenIndex = position920, tokenIndex920
					if !_rules[ruleMiscodedChar]() {
						goto l918
					}
				}
			l920:
				add(ruleNameLowerChar, position919)
			}
			return true
		l918:
			position, tokenIndex = position918, tokenIndex918
			return false
		},
		/* 103 MiscodedChar <- <'�'> */
		func() bool {
			position923, tokenIndex923 := position, tokenIndex
			{
				position924 := position
				if buffer[position] != rune('�') {
					goto l923
				}
				position++
				add(ruleMiscodedChar, position924)
			}
			return true
		l923:
			position, tokenIndex = position923, tokenIndex923
			return false
		},
		/* 104 LowerCharExtended <- <('æ' / 'œ' / 'à' / 'â' / 'å' / 'ã' / 'ä' / 'á' / 'ç' / 'č' / 'é' / 'è' / 'ë' / 'í' / 'ì' / 'ï' / 'ň' / 'ñ' / 'ñ' / 'ó' / 'ò' / 'ô' / 'ø' / 'õ' / 'ö' / 'ú' / 'ù' / 'ü' / 'ŕ' / 'ř' / 'ŗ' / 'ſ' / 'š' / 'š' / 'ş' / 'ž')> */
		func() bool {
			position925, tokenIndex925 := position, tokenIndex
			{
				position926 := position
				{
					position927, tokenIndex927 := position, tokenIndex
					if buffer[position] != rune('æ') {
						goto l928
					}
					position++
					goto l927
				l928:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('œ') {
						goto l929
					}
					position++
					goto l927
				l929:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('à') {
						goto l930
					}
					position++
					goto l927
				l930:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('â') {
						goto l931
					}
					position++
					goto l927
				l931:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('å') {
						goto l932
					}
					position++
					goto l927
				l932:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('ã') {
						goto l933
					}
					position++
					goto l927
				l933:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('ä') {
						goto l934
					}
					position++
					goto l927
				l934:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('á') {
						goto l935
					}
					position++
					goto l927
				l935:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('ç') {
						goto l936
					}
					position++
					goto l927
				l936:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('č') {
						goto l937
					}
					position++
					goto l927
				l937:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('é') {
						goto l938
					}
					position++
					goto l927
				l938:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('è') {
						goto l939
					}
					position++
					goto l927
				l939:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('ë') {
						goto l940
					}
					position++
					goto l927
				l940:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('í') {
						goto l941
					}
					position++
					goto l927
				l941:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('ì') {
						goto l942
					}
					position++
					goto l927
				l942:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('ï') {
						goto l943
					}
					position++
					goto l927
				l943:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('ň') {
						goto l944
					}
					position++
					goto l927
				l944:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('ñ') {
						goto l945
					}
					position++
					goto l927
				l945:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('ñ') {
						goto l946
					}
					position++
					goto l927
				l946:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('ó') {
						goto l947
					}
					position++
					goto l927
				l947:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('ò') {
						goto l948
					}
					position++
					goto l927
				l948:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('ô') {
						goto l949
					}
					position++
					goto l927
				l949:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('ø') {
						goto l950
					}
					position++
					goto l927
				l950:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('õ') {
						goto l951
					}
					position++
					goto l927
				l951:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('ö') {
						goto l952
					}
					position++
					goto l927
				l952:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('ú') {
						goto l953
					}
					position++
					goto l927
				l953:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('ù') {
						goto l954
					}
					position++
					goto l927
				l954:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('ü') {
						goto l955
					}
					position++
					goto l927
				l955:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('ŕ') {
						goto l956
					}
					position++
					goto l927
				l956:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('ř') {
						goto l957
					}
					position++
					goto l927
				l957:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('ŗ') {
						goto l958
					}
					position++
					goto l927
				l958:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('ſ') {
						goto l959
					}
					position++
					goto l927
				l959:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('š') {
						goto l960
					}
					position++
					goto l927
				l960:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('š') {
						goto l961
					}
					position++
					goto l927
				l961:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('ş') {
						goto l962
					}
					position++
					goto l927
				l962:
					position, tokenIndex = position927, tokenIndex927
					if buffer[position] != rune('ž') {
						goto l925
					}
					position++
				}
			l927:
				add(ruleLowerCharExtended, position926)
			}
			return true
		l925:
			position, tokenIndex = position925, tokenIndex925
			return false
		},
		/* 105 LowerChar <- <LowerASCII> */
		func() bool {
			position963, tokenIndex963 := position, tokenIndex
			{
				position964 := position
				if !_rules[ruleLowerASCII]() {
					goto l963
				}
				add(ruleLowerChar, position964)
			}
			return true
		l963:
			position, tokenIndex = position963, tokenIndex963
			return false
		},
		/* 106 SpaceCharEOI <- <(_ / !.)> */
		func() bool {
			position965, tokenIndex965 := position, tokenIndex
			{
				position966 := position
				{
					position967, tokenIndex967 := position, tokenIndex
					if !_rules[rule_]() {
						goto l968
					}
					goto l967
				l968:
					position, tokenIndex = position967, tokenIndex967
					{
						position969, tokenIndex969 := position, tokenIndex
						if !matchDot() {
							goto l969
						}
						goto l965
					l969:
						position, tokenIndex = position969, tokenIndex969
					}
				}
			l967:
				add(ruleSpaceCharEOI, position966)
			}
			return true
		l965:
			position, tokenIndex = position965, tokenIndex965
			return false
		},
		/* 107 Nums <- <[0-9]> */
		func() bool {
			position970, tokenIndex970 := position, tokenIndex
			{
				position971 := position
				if c := buffer[position]; c < rune('0') || c > rune('9') {
					goto l970
				}
				position++
				add(ruleNums, position971)
			}
			return true
		l970:
			position, tokenIndex = position970, tokenIndex970
			return false
		},
		/* 108 LowerGreek <- <[α-ω]> */
		func() bool {
			position972, tokenIndex972 := position, tokenIndex
			{
				position973 := position
				if c := buffer[position]; c < rune('α') || c > rune('ω') {
					goto l972
				}
				position++
				add(ruleLowerGreek, position973)
			}
			return true
		l972:
			position, tokenIndex = position972, tokenIndex972
			return false
		},
		/* 109 LowerASCII <- <[a-z]> */
		func() bool {
			position974, tokenIndex974 := position, tokenIndex
			{
				position975 := position
				if c := buffer[position]; c < rune('a') || c > rune('z') {
					goto l974
				}
				position++
				add(ruleLowerASCII, position975)
			}
			return true
		l974:
			position, tokenIndex = position974, tokenIndex974
			return false
		},
		/* 110 UpperASCII <- <[A-Z]> */
		func() bool {
			position976, tokenIndex976 := position, tokenIndex
			{
				position977 := position
				if c := buffer[position]; c < rune('A') || c > rune('Z') {
					goto l976
				}
				position++
				add(ruleUpperASCII, position977)
			}
			return true
		l976:
			position, tokenIndex = position976, tokenIndex976
			return false
		},
		/* 111 Apostrophe <- <(ApostrOther / ApostrASCII)> */
		func() bool {
			position978, tokenIndex978 := position, tokenIndex
			{
				position979 := position
				{
					position980, tokenIndex980 := position, tokenIndex
					if !_rules[ruleApostrOther]() {
						goto l981
					}
					goto l980
				l981:
					position, tokenIndex = position980, tokenIndex980
					if !_rules[ruleApostrASCII]() {
						goto l978
					}
				}
			l980:
				add(ruleApostrophe, position979)
			}
			return true
		l978:
			position, tokenIndex = position978, tokenIndex978
			return false
		},
		/* 112 ApostrASCII <- <'\''> */
		func() bool {
			position982, tokenIndex982 := position, tokenIndex
			{
				position983 := position
				if buffer[position] != rune('\'') {
					goto l982
				}
				position++
				add(ruleApostrASCII, position983)
			}
			return true
		l982:
			position, tokenIndex = position982, tokenIndex982
			return false
		},
		/* 113 ApostrOther <- <('‘' / '’')> */
		func() bool {
			position984, tokenIndex984 := position, tokenIndex
			{
				position985 := position
				{
					position986, tokenIndex986 := position, tokenIndex
					if buffer[position] != rune('‘') {
						goto l987
					}
					position++
					goto l986
				l987:
					position, tokenIndex = position986, tokenIndex986
					if buffer[position] != rune('’') {
						goto l984
					}
					position++
				}
			l986:
				add(ruleApostrOther, position985)
			}
			return true
		l984:
			position, tokenIndex = position984, tokenIndex984
			return false
		},
		/* 114 Dash <- <'-'> */
		func() bool {
			position988, tokenIndex988 := position, tokenIndex
			{
				position989 := position
				if buffer[position] != rune('-') {
					goto l988
				}
				position++
				add(ruleDash, position989)
			}
			return true
		l988:
			position, tokenIndex = position988, tokenIndex988
			return false
		},
		/* 115 _ <- <(MultipleSpace / SingleSpace)> */
		func() bool {
			position990, tokenIndex990 := position, tokenIndex
			{
				position991 := position
				{
					position992, tokenIndex992 := position, tokenIndex
					if !_rules[ruleMultipleSpace]() {
						goto l993
					}
					goto l992
				l993:
					position, tokenIndex = position992, tokenIndex992
					if !_rules[ruleSingleSpace]() {
						goto l990
					}
				}
			l992:
				add(rule_, position991)
			}
			return true
		l990:
			position, tokenIndex = position990, tokenIndex990
			return false
		},
		/* 116 MultipleSpace <- <(SingleSpace SingleSpace+)> */
		func() bool {
			position994, tokenIndex994 := position, tokenIndex
			{
				position995 := position
				if !_rules[ruleSingleSpace]() {
					goto l994
				}
				if !_rules[ruleSingleSpace]() {
					goto l994
				}
			l996:
				{
					position997, tokenIndex997 := position, tokenIndex
					if !_rules[ruleSingleSpace]() {
						goto l997
					}
					goto l996
				l997:
					position, tokenIndex = position997, tokenIndex997
				}
				add(ruleMultipleSpace, position995)
			}
			return true
		l994:
			position, tokenIndex = position994, tokenIndex994
			return false
		},
		/* 117 SingleSpace <- <(' ' / OtherSpace)> */
		func() bool {
			position998, tokenIndex998 := position, tokenIndex
			{
				position999 := position
				{
					position1000, tokenIndex1000 := position, tokenIndex
					if buffer[position] != rune(' ') {
						goto l1001
					}
					position++
					goto l1000
				l1001:
					position, tokenIndex = position1000, tokenIndex1000
					if !_rules[ruleOtherSpace]() {
						goto l998
					}
				}
			l1000:
				add(ruleSingleSpace, position999)
			}
			return true
		l998:
			position, tokenIndex = position998, tokenIndex998
			return false
		},
		/* 118 OtherSpace <- <('\u3000' / '\u00a0' / '\t' / '\r' / '\n' / '\f' / '\v')> */
		func() bool {
			position1002, tokenIndex1002 := position, tokenIndex
			{
				position1003 := position
				{
					position1004, tokenIndex1004 := position, tokenIndex
					if buffer[position] != rune('\u3000') {
						goto l1005
					}
					position++
					goto l1004
				l1005:
					position, tokenIndex = position1004, tokenIndex1004
					if buffer[position] != rune('\u00a0') {
						goto l1006
					}
					position++
					goto l1004
				l1006:
					position, tokenIndex = position1004, tokenIndex1004
					if buffer[position] != rune('\t') {
						goto l1007
					}
					position++
					goto l1004
				l1007:
					position, tokenIndex = position1004, tokenIndex1004
					if buffer[position] != rune('\r') {
						goto l1008
					}
					position++
					goto l1004
				l1008:
					position, tokenIndex = position1004, tokenIndex1004
					if buffer[position] != rune('\n') {
						goto l1009
					}
					position++
					goto l1004
				l1009:
					position, tokenIndex = position1004, tokenIndex1004
					if buffer[position] != rune('\f') {
						goto l1010
					}
					position++
					goto l1004
				l1010:
					position, tokenIndex = position1004, tokenIndex1004
					if buffer[position] != rune('\v') {
						goto l1002
					}
					position++
				}
			l1004:
				add(ruleOtherSpace, position1003)
			}
			return true
		l1002:
			position, tokenIndex = position1002, tokenIndex1002
			return false
		},
		/* 120 Action0 <- <{ p.AddWarn(YearCharWarn) }> */
		func() bool {
			{
				add(ruleAction0, position)
			}
			return true
		},
	}
	p.rules = _rules
}
