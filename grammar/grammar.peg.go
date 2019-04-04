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
	rules  [120]func() bool
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
		/* 19 Rank <- <(RankForma / RankVar / RankSsp / RankOther / RankOtherUncommon / RankAgamo / RankNotho)> */
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
				add(ruleRank, position113)
			}
			return true
		l112:
			position, tokenIndex = position112, tokenIndex112
			return false
		},
		/* 20 RankNotho <- <((('n' 'o' 't' 'h' 'o' (('v' 'a' 'r') / ('f' 'o') / 'f' / ('s' 'u' 'b' 's' 'p') / ('s' 's' 'p') / ('s' 'p') / ('m' 'o' 'r' 't' 'h') / ('s' 'u' 'p' 's' 'p') / ('s' 'u'))) / ('n' 'v' 'a' 'r')) ('.' / &SpaceCharEOI))> */
		func() bool {
			position121, tokenIndex121 := position, tokenIndex
			{
				position122 := position
				{
					position123, tokenIndex123 := position, tokenIndex
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
					{
						position125, tokenIndex125 := position, tokenIndex
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
						goto l125
					l126:
						position, tokenIndex = position125, tokenIndex125
						if buffer[position] != rune('f') {
							goto l127
						}
						position++
						if buffer[position] != rune('o') {
							goto l127
						}
						position++
						goto l125
					l127:
						position, tokenIndex = position125, tokenIndex125
						if buffer[position] != rune('f') {
							goto l128
						}
						position++
						goto l125
					l128:
						position, tokenIndex = position125, tokenIndex125
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
						if buffer[position] != rune('s') {
							goto l129
						}
						position++
						if buffer[position] != rune('p') {
							goto l129
						}
						position++
						goto l125
					l129:
						position, tokenIndex = position125, tokenIndex125
						if buffer[position] != rune('s') {
							goto l130
						}
						position++
						if buffer[position] != rune('s') {
							goto l130
						}
						position++
						if buffer[position] != rune('p') {
							goto l130
						}
						position++
						goto l125
					l130:
						position, tokenIndex = position125, tokenIndex125
						if buffer[position] != rune('s') {
							goto l131
						}
						position++
						if buffer[position] != rune('p') {
							goto l131
						}
						position++
						goto l125
					l131:
						position, tokenIndex = position125, tokenIndex125
						if buffer[position] != rune('m') {
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
						if buffer[position] != rune('t') {
							goto l132
						}
						position++
						if buffer[position] != rune('h') {
							goto l132
						}
						position++
						goto l125
					l132:
						position, tokenIndex = position125, tokenIndex125
						if buffer[position] != rune('s') {
							goto l133
						}
						position++
						if buffer[position] != rune('u') {
							goto l133
						}
						position++
						if buffer[position] != rune('p') {
							goto l133
						}
						position++
						if buffer[position] != rune('s') {
							goto l133
						}
						position++
						if buffer[position] != rune('p') {
							goto l133
						}
						position++
						goto l125
					l133:
						position, tokenIndex = position125, tokenIndex125
						if buffer[position] != rune('s') {
							goto l124
						}
						position++
						if buffer[position] != rune('u') {
							goto l124
						}
						position++
					}
				l125:
					goto l123
				l124:
					position, tokenIndex = position123, tokenIndex123
					if buffer[position] != rune('n') {
						goto l121
					}
					position++
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
				}
			l123:
				{
					position134, tokenIndex134 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l135
					}
					position++
					goto l134
				l135:
					position, tokenIndex = position134, tokenIndex134
					{
						position136, tokenIndex136 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l121
						}
						position, tokenIndex = position136, tokenIndex136
					}
				}
			l134:
				add(ruleRankNotho, position122)
			}
			return true
		l121:
			position, tokenIndex = position121, tokenIndex121
			return false
		},
		/* 21 RankOtherUncommon <- <(('*' / ('n' 'a' 't' 'i' 'o') / ('n' 'a' 't' '.') / ('n' 'a' 't') / ('f' '.' 's' 'p') / 'α' / ('β' 'β') / 'β' / 'γ' / 'δ' / 'ε' / 'φ' / 'θ' / 'μ' / ('a' '.') / ('b' '.') / ('c' '.') / ('d' '.') / ('e' '.') / ('g' '.') / ('k' '.') / ('m' 'u' 't' '.')) &SpaceCharEOI)> */
		func() bool {
			position137, tokenIndex137 := position, tokenIndex
			{
				position138 := position
				{
					position139, tokenIndex139 := position, tokenIndex
					if buffer[position] != rune('*') {
						goto l140
					}
					position++
					goto l139
				l140:
					position, tokenIndex = position139, tokenIndex139
					if buffer[position] != rune('n') {
						goto l141
					}
					position++
					if buffer[position] != rune('a') {
						goto l141
					}
					position++
					if buffer[position] != rune('t') {
						goto l141
					}
					position++
					if buffer[position] != rune('i') {
						goto l141
					}
					position++
					if buffer[position] != rune('o') {
						goto l141
					}
					position++
					goto l139
				l141:
					position, tokenIndex = position139, tokenIndex139
					if buffer[position] != rune('n') {
						goto l142
					}
					position++
					if buffer[position] != rune('a') {
						goto l142
					}
					position++
					if buffer[position] != rune('t') {
						goto l142
					}
					position++
					if buffer[position] != rune('.') {
						goto l142
					}
					position++
					goto l139
				l142:
					position, tokenIndex = position139, tokenIndex139
					if buffer[position] != rune('n') {
						goto l143
					}
					position++
					if buffer[position] != rune('a') {
						goto l143
					}
					position++
					if buffer[position] != rune('t') {
						goto l143
					}
					position++
					goto l139
				l143:
					position, tokenIndex = position139, tokenIndex139
					if buffer[position] != rune('f') {
						goto l144
					}
					position++
					if buffer[position] != rune('.') {
						goto l144
					}
					position++
					if buffer[position] != rune('s') {
						goto l144
					}
					position++
					if buffer[position] != rune('p') {
						goto l144
					}
					position++
					goto l139
				l144:
					position, tokenIndex = position139, tokenIndex139
					if buffer[position] != rune('α') {
						goto l145
					}
					position++
					goto l139
				l145:
					position, tokenIndex = position139, tokenIndex139
					if buffer[position] != rune('β') {
						goto l146
					}
					position++
					if buffer[position] != rune('β') {
						goto l146
					}
					position++
					goto l139
				l146:
					position, tokenIndex = position139, tokenIndex139
					if buffer[position] != rune('β') {
						goto l147
					}
					position++
					goto l139
				l147:
					position, tokenIndex = position139, tokenIndex139
					if buffer[position] != rune('γ') {
						goto l148
					}
					position++
					goto l139
				l148:
					position, tokenIndex = position139, tokenIndex139
					if buffer[position] != rune('δ') {
						goto l149
					}
					position++
					goto l139
				l149:
					position, tokenIndex = position139, tokenIndex139
					if buffer[position] != rune('ε') {
						goto l150
					}
					position++
					goto l139
				l150:
					position, tokenIndex = position139, tokenIndex139
					if buffer[position] != rune('φ') {
						goto l151
					}
					position++
					goto l139
				l151:
					position, tokenIndex = position139, tokenIndex139
					if buffer[position] != rune('θ') {
						goto l152
					}
					position++
					goto l139
				l152:
					position, tokenIndex = position139, tokenIndex139
					if buffer[position] != rune('μ') {
						goto l153
					}
					position++
					goto l139
				l153:
					position, tokenIndex = position139, tokenIndex139
					if buffer[position] != rune('a') {
						goto l154
					}
					position++
					if buffer[position] != rune('.') {
						goto l154
					}
					position++
					goto l139
				l154:
					position, tokenIndex = position139, tokenIndex139
					if buffer[position] != rune('b') {
						goto l155
					}
					position++
					if buffer[position] != rune('.') {
						goto l155
					}
					position++
					goto l139
				l155:
					position, tokenIndex = position139, tokenIndex139
					if buffer[position] != rune('c') {
						goto l156
					}
					position++
					if buffer[position] != rune('.') {
						goto l156
					}
					position++
					goto l139
				l156:
					position, tokenIndex = position139, tokenIndex139
					if buffer[position] != rune('d') {
						goto l157
					}
					position++
					if buffer[position] != rune('.') {
						goto l157
					}
					position++
					goto l139
				l157:
					position, tokenIndex = position139, tokenIndex139
					if buffer[position] != rune('e') {
						goto l158
					}
					position++
					if buffer[position] != rune('.') {
						goto l158
					}
					position++
					goto l139
				l158:
					position, tokenIndex = position139, tokenIndex139
					if buffer[position] != rune('g') {
						goto l159
					}
					position++
					if buffer[position] != rune('.') {
						goto l159
					}
					position++
					goto l139
				l159:
					position, tokenIndex = position139, tokenIndex139
					if buffer[position] != rune('k') {
						goto l160
					}
					position++
					if buffer[position] != rune('.') {
						goto l160
					}
					position++
					goto l139
				l160:
					position, tokenIndex = position139, tokenIndex139
					if buffer[position] != rune('m') {
						goto l137
					}
					position++
					if buffer[position] != rune('u') {
						goto l137
					}
					position++
					if buffer[position] != rune('t') {
						goto l137
					}
					position++
					if buffer[position] != rune('.') {
						goto l137
					}
					position++
				}
			l139:
				{
					position161, tokenIndex161 := position, tokenIndex
					if !_rules[ruleSpaceCharEOI]() {
						goto l137
					}
					position, tokenIndex = position161, tokenIndex161
				}
				add(ruleRankOtherUncommon, position138)
			}
			return true
		l137:
			position, tokenIndex = position137, tokenIndex137
			return false
		},
		/* 22 RankOther <- <((('m' 'o' 'r' 'p' 'h') / ('c' 'o' 'n' 'v' 'a' 'r') / ('p' 's' 'e' 'u' 'd' 'o' 'v' 'a' 'r') / ('s' 'e' 'c' 't') / ('s' 'e' 'r') / ('s' 'u' 'b' 'v' 'a' 'r') / ('s' 'u' 'b' 'f') / ('r' 'a' 'c' 'e') / ('p' 'v') / ('p' 'a' 't' 'h' 'o' 'v' 'a' 'r') / ('a' 'b' '.' (_? ('n' '.'))?) / ('s' 't')) ('.' / &SpaceCharEOI))> */
		func() bool {
			position162, tokenIndex162 := position, tokenIndex
			{
				position163 := position
				{
					position164, tokenIndex164 := position, tokenIndex
					if buffer[position] != rune('m') {
						goto l165
					}
					position++
					if buffer[position] != rune('o') {
						goto l165
					}
					position++
					if buffer[position] != rune('r') {
						goto l165
					}
					position++
					if buffer[position] != rune('p') {
						goto l165
					}
					position++
					if buffer[position] != rune('h') {
						goto l165
					}
					position++
					goto l164
				l165:
					position, tokenIndex = position164, tokenIndex164
					if buffer[position] != rune('c') {
						goto l166
					}
					position++
					if buffer[position] != rune('o') {
						goto l166
					}
					position++
					if buffer[position] != rune('n') {
						goto l166
					}
					position++
					if buffer[position] != rune('v') {
						goto l166
					}
					position++
					if buffer[position] != rune('a') {
						goto l166
					}
					position++
					if buffer[position] != rune('r') {
						goto l166
					}
					position++
					goto l164
				l166:
					position, tokenIndex = position164, tokenIndex164
					if buffer[position] != rune('p') {
						goto l167
					}
					position++
					if buffer[position] != rune('s') {
						goto l167
					}
					position++
					if buffer[position] != rune('e') {
						goto l167
					}
					position++
					if buffer[position] != rune('u') {
						goto l167
					}
					position++
					if buffer[position] != rune('d') {
						goto l167
					}
					position++
					if buffer[position] != rune('o') {
						goto l167
					}
					position++
					if buffer[position] != rune('v') {
						goto l167
					}
					position++
					if buffer[position] != rune('a') {
						goto l167
					}
					position++
					if buffer[position] != rune('r') {
						goto l167
					}
					position++
					goto l164
				l167:
					position, tokenIndex = position164, tokenIndex164
					if buffer[position] != rune('s') {
						goto l168
					}
					position++
					if buffer[position] != rune('e') {
						goto l168
					}
					position++
					if buffer[position] != rune('c') {
						goto l168
					}
					position++
					if buffer[position] != rune('t') {
						goto l168
					}
					position++
					goto l164
				l168:
					position, tokenIndex = position164, tokenIndex164
					if buffer[position] != rune('s') {
						goto l169
					}
					position++
					if buffer[position] != rune('e') {
						goto l169
					}
					position++
					if buffer[position] != rune('r') {
						goto l169
					}
					position++
					goto l164
				l169:
					position, tokenIndex = position164, tokenIndex164
					if buffer[position] != rune('s') {
						goto l170
					}
					position++
					if buffer[position] != rune('u') {
						goto l170
					}
					position++
					if buffer[position] != rune('b') {
						goto l170
					}
					position++
					if buffer[position] != rune('v') {
						goto l170
					}
					position++
					if buffer[position] != rune('a') {
						goto l170
					}
					position++
					if buffer[position] != rune('r') {
						goto l170
					}
					position++
					goto l164
				l170:
					position, tokenIndex = position164, tokenIndex164
					if buffer[position] != rune('s') {
						goto l171
					}
					position++
					if buffer[position] != rune('u') {
						goto l171
					}
					position++
					if buffer[position] != rune('b') {
						goto l171
					}
					position++
					if buffer[position] != rune('f') {
						goto l171
					}
					position++
					goto l164
				l171:
					position, tokenIndex = position164, tokenIndex164
					if buffer[position] != rune('r') {
						goto l172
					}
					position++
					if buffer[position] != rune('a') {
						goto l172
					}
					position++
					if buffer[position] != rune('c') {
						goto l172
					}
					position++
					if buffer[position] != rune('e') {
						goto l172
					}
					position++
					goto l164
				l172:
					position, tokenIndex = position164, tokenIndex164
					if buffer[position] != rune('p') {
						goto l173
					}
					position++
					if buffer[position] != rune('v') {
						goto l173
					}
					position++
					goto l164
				l173:
					position, tokenIndex = position164, tokenIndex164
					if buffer[position] != rune('p') {
						goto l174
					}
					position++
					if buffer[position] != rune('a') {
						goto l174
					}
					position++
					if buffer[position] != rune('t') {
						goto l174
					}
					position++
					if buffer[position] != rune('h') {
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
					goto l164
				l174:
					position, tokenIndex = position164, tokenIndex164
					if buffer[position] != rune('a') {
						goto l175
					}
					position++
					if buffer[position] != rune('b') {
						goto l175
					}
					position++
					if buffer[position] != rune('.') {
						goto l175
					}
					position++
					{
						position176, tokenIndex176 := position, tokenIndex
						{
							position178, tokenIndex178 := position, tokenIndex
							if !_rules[rule_]() {
								goto l178
							}
							goto l179
						l178:
							position, tokenIndex = position178, tokenIndex178
						}
					l179:
						if buffer[position] != rune('n') {
							goto l176
						}
						position++
						if buffer[position] != rune('.') {
							goto l176
						}
						position++
						goto l177
					l176:
						position, tokenIndex = position176, tokenIndex176
					}
				l177:
					goto l164
				l175:
					position, tokenIndex = position164, tokenIndex164
					if buffer[position] != rune('s') {
						goto l162
					}
					position++
					if buffer[position] != rune('t') {
						goto l162
					}
					position++
				}
			l164:
				{
					position180, tokenIndex180 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l181
					}
					position++
					goto l180
				l181:
					position, tokenIndex = position180, tokenIndex180
					{
						position182, tokenIndex182 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l162
						}
						position, tokenIndex = position182, tokenIndex182
					}
				}
			l180:
				add(ruleRankOther, position163)
			}
			return true
		l162:
			position, tokenIndex = position162, tokenIndex162
			return false
		},
		/* 23 RankVar <- <((('v' 'a' 'r' 'i' 'e' 't' 'y') / ('[' 'v' 'a' 'r' '.' ']') / ('v' 'a' 'r')) ('.' / &SpaceCharEOI))> */
		func() bool {
			position183, tokenIndex183 := position, tokenIndex
			{
				position184 := position
				{
					position185, tokenIndex185 := position, tokenIndex
					if buffer[position] != rune('v') {
						goto l186
					}
					position++
					if buffer[position] != rune('a') {
						goto l186
					}
					position++
					if buffer[position] != rune('r') {
						goto l186
					}
					position++
					if buffer[position] != rune('i') {
						goto l186
					}
					position++
					if buffer[position] != rune('e') {
						goto l186
					}
					position++
					if buffer[position] != rune('t') {
						goto l186
					}
					position++
					if buffer[position] != rune('y') {
						goto l186
					}
					position++
					goto l185
				l186:
					position, tokenIndex = position185, tokenIndex185
					if buffer[position] != rune('[') {
						goto l187
					}
					position++
					if buffer[position] != rune('v') {
						goto l187
					}
					position++
					if buffer[position] != rune('a') {
						goto l187
					}
					position++
					if buffer[position] != rune('r') {
						goto l187
					}
					position++
					if buffer[position] != rune('.') {
						goto l187
					}
					position++
					if buffer[position] != rune(']') {
						goto l187
					}
					position++
					goto l185
				l187:
					position, tokenIndex = position185, tokenIndex185
					if buffer[position] != rune('v') {
						goto l183
					}
					position++
					if buffer[position] != rune('a') {
						goto l183
					}
					position++
					if buffer[position] != rune('r') {
						goto l183
					}
					position++
				}
			l185:
				{
					position188, tokenIndex188 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l189
					}
					position++
					goto l188
				l189:
					position, tokenIndex = position188, tokenIndex188
					{
						position190, tokenIndex190 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l183
						}
						position, tokenIndex = position190, tokenIndex190
					}
				}
			l188:
				add(ruleRankVar, position184)
			}
			return true
		l183:
			position, tokenIndex = position183, tokenIndex183
			return false
		},
		/* 24 RankForma <- <((('f' 'o' 'r' 'm' 'a') / ('f' 'm' 'a') / ('f' 'o' 'r' 'm') / ('f' 'o') / 'f') ('.' / &SpaceCharEOI))> */
		func() bool {
			position191, tokenIndex191 := position, tokenIndex
			{
				position192 := position
				{
					position193, tokenIndex193 := position, tokenIndex
					if buffer[position] != rune('f') {
						goto l194
					}
					position++
					if buffer[position] != rune('o') {
						goto l194
					}
					position++
					if buffer[position] != rune('r') {
						goto l194
					}
					position++
					if buffer[position] != rune('m') {
						goto l194
					}
					position++
					if buffer[position] != rune('a') {
						goto l194
					}
					position++
					goto l193
				l194:
					position, tokenIndex = position193, tokenIndex193
					if buffer[position] != rune('f') {
						goto l195
					}
					position++
					if buffer[position] != rune('m') {
						goto l195
					}
					position++
					if buffer[position] != rune('a') {
						goto l195
					}
					position++
					goto l193
				l195:
					position, tokenIndex = position193, tokenIndex193
					if buffer[position] != rune('f') {
						goto l196
					}
					position++
					if buffer[position] != rune('o') {
						goto l196
					}
					position++
					if buffer[position] != rune('r') {
						goto l196
					}
					position++
					if buffer[position] != rune('m') {
						goto l196
					}
					position++
					goto l193
				l196:
					position, tokenIndex = position193, tokenIndex193
					if buffer[position] != rune('f') {
						goto l197
					}
					position++
					if buffer[position] != rune('o') {
						goto l197
					}
					position++
					goto l193
				l197:
					position, tokenIndex = position193, tokenIndex193
					if buffer[position] != rune('f') {
						goto l191
					}
					position++
				}
			l193:
				{
					position198, tokenIndex198 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l199
					}
					position++
					goto l198
				l199:
					position, tokenIndex = position198, tokenIndex198
					{
						position200, tokenIndex200 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l191
						}
						position, tokenIndex = position200, tokenIndex200
					}
				}
			l198:
				add(ruleRankForma, position192)
			}
			return true
		l191:
			position, tokenIndex = position191, tokenIndex191
			return false
		},
		/* 25 RankSsp <- <((('s' 's' 'p') / ('s' 'u' 'b' 's' 'p')) ('.' / &SpaceCharEOI))> */
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
					if buffer[position] != rune('s') {
						goto l204
					}
					position++
					if buffer[position] != rune('p') {
						goto l204
					}
					position++
					goto l203
				l204:
					position, tokenIndex = position203, tokenIndex203
					if buffer[position] != rune('s') {
						goto l201
					}
					position++
					if buffer[position] != rune('u') {
						goto l201
					}
					position++
					if buffer[position] != rune('b') {
						goto l201
					}
					position++
					if buffer[position] != rune('s') {
						goto l201
					}
					position++
					if buffer[position] != rune('p') {
						goto l201
					}
					position++
				}
			l203:
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
							goto l201
						}
						position, tokenIndex = position207, tokenIndex207
					}
				}
			l205:
				add(ruleRankSsp, position202)
			}
			return true
		l201:
			position, tokenIndex = position201, tokenIndex201
			return false
		},
		/* 26 RankAgamo <- <((('a' 'g' 'a' 'm' 'o' 's' 'p') / ('a' 'g' 'a' 'm' 'o' 's' 's' 'p') / ('a' 'g' 'a' 'm' 'o' 'v' 'a' 'r')) ('.' / &SpaceCharEOI))> */
		func() bool {
			position208, tokenIndex208 := position, tokenIndex
			{
				position209 := position
				{
					position210, tokenIndex210 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l211
					}
					position++
					if buffer[position] != rune('g') {
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
					if buffer[position] != rune('o') {
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
					if buffer[position] != rune('a') {
						goto l212
					}
					position++
					if buffer[position] != rune('g') {
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
					if buffer[position] != rune('o') {
						goto l212
					}
					position++
					if buffer[position] != rune('s') {
						goto l212
					}
					position++
					if buffer[position] != rune('s') {
						goto l212
					}
					position++
					if buffer[position] != rune('p') {
						goto l212
					}
					position++
					goto l210
				l212:
					position, tokenIndex = position210, tokenIndex210
					if buffer[position] != rune('a') {
						goto l208
					}
					position++
					if buffer[position] != rune('g') {
						goto l208
					}
					position++
					if buffer[position] != rune('a') {
						goto l208
					}
					position++
					if buffer[position] != rune('m') {
						goto l208
					}
					position++
					if buffer[position] != rune('o') {
						goto l208
					}
					position++
					if buffer[position] != rune('v') {
						goto l208
					}
					position++
					if buffer[position] != rune('a') {
						goto l208
					}
					position++
					if buffer[position] != rune('r') {
						goto l208
					}
					position++
				}
			l210:
				{
					position213, tokenIndex213 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l214
					}
					position++
					goto l213
				l214:
					position, tokenIndex = position213, tokenIndex213
					{
						position215, tokenIndex215 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l208
						}
						position, tokenIndex = position215, tokenIndex215
					}
				}
			l213:
				add(ruleRankAgamo, position209)
			}
			return true
		l208:
			position, tokenIndex = position208, tokenIndex208
			return false
		},
		/* 27 SubGenusOrSuperspecies <- <('(' _? NameLowerChar+ _? ')')> */
		func() bool {
			position216, tokenIndex216 := position, tokenIndex
			{
				position217 := position
				if buffer[position] != rune('(') {
					goto l216
				}
				position++
				{
					position218, tokenIndex218 := position, tokenIndex
					if !_rules[rule_]() {
						goto l218
					}
					goto l219
				l218:
					position, tokenIndex = position218, tokenIndex218
				}
			l219:
				if !_rules[ruleNameLowerChar]() {
					goto l216
				}
			l220:
				{
					position221, tokenIndex221 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l221
					}
					goto l220
				l221:
					position, tokenIndex = position221, tokenIndex221
				}
				{
					position222, tokenIndex222 := position, tokenIndex
					if !_rules[rule_]() {
						goto l222
					}
					goto l223
				l222:
					position, tokenIndex = position222, tokenIndex222
				}
			l223:
				if buffer[position] != rune(')') {
					goto l216
				}
				position++
				add(ruleSubGenusOrSuperspecies, position217)
			}
			return true
		l216:
			position, tokenIndex = position216, tokenIndex216
			return false
		},
		/* 28 SubGenus <- <('(' _? UninomialWord _? ')')> */
		func() bool {
			position224, tokenIndex224 := position, tokenIndex
			{
				position225 := position
				if buffer[position] != rune('(') {
					goto l224
				}
				position++
				{
					position226, tokenIndex226 := position, tokenIndex
					if !_rules[rule_]() {
						goto l226
					}
					goto l227
				l226:
					position, tokenIndex = position226, tokenIndex226
				}
			l227:
				if !_rules[ruleUninomialWord]() {
					goto l224
				}
				{
					position228, tokenIndex228 := position, tokenIndex
					if !_rules[rule_]() {
						goto l228
					}
					goto l229
				l228:
					position, tokenIndex = position228, tokenIndex228
				}
			l229:
				if buffer[position] != rune(')') {
					goto l224
				}
				position++
				add(ruleSubGenus, position225)
			}
			return true
		l224:
			position, tokenIndex = position224, tokenIndex224
			return false
		},
		/* 29 UninomialCombo <- <(UninomialCombo1 / UninomialCombo2)> */
		func() bool {
			position230, tokenIndex230 := position, tokenIndex
			{
				position231 := position
				{
					position232, tokenIndex232 := position, tokenIndex
					if !_rules[ruleUninomialCombo1]() {
						goto l233
					}
					goto l232
				l233:
					position, tokenIndex = position232, tokenIndex232
					if !_rules[ruleUninomialCombo2]() {
						goto l230
					}
				}
			l232:
				add(ruleUninomialCombo, position231)
			}
			return true
		l230:
			position, tokenIndex = position230, tokenIndex230
			return false
		},
		/* 30 UninomialCombo1 <- <(UninomialWord _? SubGenus (_? Authorship)?)> */
		func() bool {
			position234, tokenIndex234 := position, tokenIndex
			{
				position235 := position
				if !_rules[ruleUninomialWord]() {
					goto l234
				}
				{
					position236, tokenIndex236 := position, tokenIndex
					if !_rules[rule_]() {
						goto l236
					}
					goto l237
				l236:
					position, tokenIndex = position236, tokenIndex236
				}
			l237:
				if !_rules[ruleSubGenus]() {
					goto l234
				}
				{
					position238, tokenIndex238 := position, tokenIndex
					{
						position240, tokenIndex240 := position, tokenIndex
						if !_rules[rule_]() {
							goto l240
						}
						goto l241
					l240:
						position, tokenIndex = position240, tokenIndex240
					}
				l241:
					if !_rules[ruleAuthorship]() {
						goto l238
					}
					goto l239
				l238:
					position, tokenIndex = position238, tokenIndex238
				}
			l239:
				add(ruleUninomialCombo1, position235)
			}
			return true
		l234:
			position, tokenIndex = position234, tokenIndex234
			return false
		},
		/* 31 UninomialCombo2 <- <(Uninomial _ RankUninomial _ Uninomial)> */
		func() bool {
			position242, tokenIndex242 := position, tokenIndex
			{
				position243 := position
				if !_rules[ruleUninomial]() {
					goto l242
				}
				if !_rules[rule_]() {
					goto l242
				}
				if !_rules[ruleRankUninomial]() {
					goto l242
				}
				if !_rules[rule_]() {
					goto l242
				}
				if !_rules[ruleUninomial]() {
					goto l242
				}
				add(ruleUninomialCombo2, position243)
			}
			return true
		l242:
			position, tokenIndex = position242, tokenIndex242
			return false
		},
		/* 32 RankUninomial <- <(RankUninomialPlain / RankUninomialNotho)> */
		func() bool {
			position244, tokenIndex244 := position, tokenIndex
			{
				position245 := position
				{
					position246, tokenIndex246 := position, tokenIndex
					if !_rules[ruleRankUninomialPlain]() {
						goto l247
					}
					goto l246
				l247:
					position, tokenIndex = position246, tokenIndex246
					if !_rules[ruleRankUninomialNotho]() {
						goto l244
					}
				}
			l246:
				add(ruleRankUninomial, position245)
			}
			return true
		l244:
			position, tokenIndex = position244, tokenIndex244
			return false
		},
		/* 33 RankUninomialPlain <- <((('s' 'e' 'c' 't') / ('s' 'u' 'b' 's' 'e' 'c' 't') / ('t' 'r' 'i' 'b') / ('s' 'u' 'b' 't' 'r' 'i' 'b') / ('s' 'u' 'b' 's' 'e' 'r') / ('s' 'e' 'r') / ('s' 'u' 'b' 'g' 'e' 'n') / ('s' 'u' 'b' 'g') / ('f' 'a' 'm') / ('s' 'u' 'b' 'f' 'a' 'm') / ('s' 'u' 'p' 'e' 'r' 't' 'r' 'i' 'b')) ('.' / &SpaceCharEOI))> */
		func() bool {
			position248, tokenIndex248 := position, tokenIndex
			{
				position249 := position
				{
					position250, tokenIndex250 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l251
					}
					position++
					if buffer[position] != rune('e') {
						goto l251
					}
					position++
					if buffer[position] != rune('c') {
						goto l251
					}
					position++
					if buffer[position] != rune('t') {
						goto l251
					}
					position++
					goto l250
				l251:
					position, tokenIndex = position250, tokenIndex250
					if buffer[position] != rune('s') {
						goto l252
					}
					position++
					if buffer[position] != rune('u') {
						goto l252
					}
					position++
					if buffer[position] != rune('b') {
						goto l252
					}
					position++
					if buffer[position] != rune('s') {
						goto l252
					}
					position++
					if buffer[position] != rune('e') {
						goto l252
					}
					position++
					if buffer[position] != rune('c') {
						goto l252
					}
					position++
					if buffer[position] != rune('t') {
						goto l252
					}
					position++
					goto l250
				l252:
					position, tokenIndex = position250, tokenIndex250
					if buffer[position] != rune('t') {
						goto l253
					}
					position++
					if buffer[position] != rune('r') {
						goto l253
					}
					position++
					if buffer[position] != rune('i') {
						goto l253
					}
					position++
					if buffer[position] != rune('b') {
						goto l253
					}
					position++
					goto l250
				l253:
					position, tokenIndex = position250, tokenIndex250
					if buffer[position] != rune('s') {
						goto l254
					}
					position++
					if buffer[position] != rune('u') {
						goto l254
					}
					position++
					if buffer[position] != rune('b') {
						goto l254
					}
					position++
					if buffer[position] != rune('t') {
						goto l254
					}
					position++
					if buffer[position] != rune('r') {
						goto l254
					}
					position++
					if buffer[position] != rune('i') {
						goto l254
					}
					position++
					if buffer[position] != rune('b') {
						goto l254
					}
					position++
					goto l250
				l254:
					position, tokenIndex = position250, tokenIndex250
					if buffer[position] != rune('s') {
						goto l255
					}
					position++
					if buffer[position] != rune('u') {
						goto l255
					}
					position++
					if buffer[position] != rune('b') {
						goto l255
					}
					position++
					if buffer[position] != rune('s') {
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
					goto l250
				l255:
					position, tokenIndex = position250, tokenIndex250
					if buffer[position] != rune('s') {
						goto l256
					}
					position++
					if buffer[position] != rune('e') {
						goto l256
					}
					position++
					if buffer[position] != rune('r') {
						goto l256
					}
					position++
					goto l250
				l256:
					position, tokenIndex = position250, tokenIndex250
					if buffer[position] != rune('s') {
						goto l257
					}
					position++
					if buffer[position] != rune('u') {
						goto l257
					}
					position++
					if buffer[position] != rune('b') {
						goto l257
					}
					position++
					if buffer[position] != rune('g') {
						goto l257
					}
					position++
					if buffer[position] != rune('e') {
						goto l257
					}
					position++
					if buffer[position] != rune('n') {
						goto l257
					}
					position++
					goto l250
				l257:
					position, tokenIndex = position250, tokenIndex250
					if buffer[position] != rune('s') {
						goto l258
					}
					position++
					if buffer[position] != rune('u') {
						goto l258
					}
					position++
					if buffer[position] != rune('b') {
						goto l258
					}
					position++
					if buffer[position] != rune('g') {
						goto l258
					}
					position++
					goto l250
				l258:
					position, tokenIndex = position250, tokenIndex250
					if buffer[position] != rune('f') {
						goto l259
					}
					position++
					if buffer[position] != rune('a') {
						goto l259
					}
					position++
					if buffer[position] != rune('m') {
						goto l259
					}
					position++
					goto l250
				l259:
					position, tokenIndex = position250, tokenIndex250
					if buffer[position] != rune('s') {
						goto l260
					}
					position++
					if buffer[position] != rune('u') {
						goto l260
					}
					position++
					if buffer[position] != rune('b') {
						goto l260
					}
					position++
					if buffer[position] != rune('f') {
						goto l260
					}
					position++
					if buffer[position] != rune('a') {
						goto l260
					}
					position++
					if buffer[position] != rune('m') {
						goto l260
					}
					position++
					goto l250
				l260:
					position, tokenIndex = position250, tokenIndex250
					if buffer[position] != rune('s') {
						goto l248
					}
					position++
					if buffer[position] != rune('u') {
						goto l248
					}
					position++
					if buffer[position] != rune('p') {
						goto l248
					}
					position++
					if buffer[position] != rune('e') {
						goto l248
					}
					position++
					if buffer[position] != rune('r') {
						goto l248
					}
					position++
					if buffer[position] != rune('t') {
						goto l248
					}
					position++
					if buffer[position] != rune('r') {
						goto l248
					}
					position++
					if buffer[position] != rune('i') {
						goto l248
					}
					position++
					if buffer[position] != rune('b') {
						goto l248
					}
					position++
				}
			l250:
				{
					position261, tokenIndex261 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l262
					}
					position++
					goto l261
				l262:
					position, tokenIndex = position261, tokenIndex261
					{
						position263, tokenIndex263 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l248
						}
						position, tokenIndex = position263, tokenIndex263
					}
				}
			l261:
				add(ruleRankUninomialPlain, position249)
			}
			return true
		l248:
			position, tokenIndex = position248, tokenIndex248
			return false
		},
		/* 34 RankUninomialNotho <- <('n' 'o' 't' 'h' 'o' _? (('s' 'e' 'c' 't') / ('g' 'e' 'n') / ('s' 'e' 'r') / ('s' 'u' 'b' 'g' 'e' 'e' 'n') / ('s' 'u' 'b' 'g' 'e' 'n') / ('s' 'u' 'b' 'g') / ('s' 'u' 'b' 's' 'e' 'c' 't') / ('s' 'u' 'b' 't' 'r' 'i' 'b')) ('.' / &SpaceCharEOI))> */
		func() bool {
			position264, tokenIndex264 := position, tokenIndex
			{
				position265 := position
				if buffer[position] != rune('n') {
					goto l264
				}
				position++
				if buffer[position] != rune('o') {
					goto l264
				}
				position++
				if buffer[position] != rune('t') {
					goto l264
				}
				position++
				if buffer[position] != rune('h') {
					goto l264
				}
				position++
				if buffer[position] != rune('o') {
					goto l264
				}
				position++
				{
					position266, tokenIndex266 := position, tokenIndex
					if !_rules[rule_]() {
						goto l266
					}
					goto l267
				l266:
					position, tokenIndex = position266, tokenIndex266
				}
			l267:
				{
					position268, tokenIndex268 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l269
					}
					position++
					if buffer[position] != rune('e') {
						goto l269
					}
					position++
					if buffer[position] != rune('c') {
						goto l269
					}
					position++
					if buffer[position] != rune('t') {
						goto l269
					}
					position++
					goto l268
				l269:
					position, tokenIndex = position268, tokenIndex268
					if buffer[position] != rune('g') {
						goto l270
					}
					position++
					if buffer[position] != rune('e') {
						goto l270
					}
					position++
					if buffer[position] != rune('n') {
						goto l270
					}
					position++
					goto l268
				l270:
					position, tokenIndex = position268, tokenIndex268
					if buffer[position] != rune('s') {
						goto l271
					}
					position++
					if buffer[position] != rune('e') {
						goto l271
					}
					position++
					if buffer[position] != rune('r') {
						goto l271
					}
					position++
					goto l268
				l271:
					position, tokenIndex = position268, tokenIndex268
					if buffer[position] != rune('s') {
						goto l272
					}
					position++
					if buffer[position] != rune('u') {
						goto l272
					}
					position++
					if buffer[position] != rune('b') {
						goto l272
					}
					position++
					if buffer[position] != rune('g') {
						goto l272
					}
					position++
					if buffer[position] != rune('e') {
						goto l272
					}
					position++
					if buffer[position] != rune('e') {
						goto l272
					}
					position++
					if buffer[position] != rune('n') {
						goto l272
					}
					position++
					goto l268
				l272:
					position, tokenIndex = position268, tokenIndex268
					if buffer[position] != rune('s') {
						goto l273
					}
					position++
					if buffer[position] != rune('u') {
						goto l273
					}
					position++
					if buffer[position] != rune('b') {
						goto l273
					}
					position++
					if buffer[position] != rune('g') {
						goto l273
					}
					position++
					if buffer[position] != rune('e') {
						goto l273
					}
					position++
					if buffer[position] != rune('n') {
						goto l273
					}
					position++
					goto l268
				l273:
					position, tokenIndex = position268, tokenIndex268
					if buffer[position] != rune('s') {
						goto l274
					}
					position++
					if buffer[position] != rune('u') {
						goto l274
					}
					position++
					if buffer[position] != rune('b') {
						goto l274
					}
					position++
					if buffer[position] != rune('g') {
						goto l274
					}
					position++
					goto l268
				l274:
					position, tokenIndex = position268, tokenIndex268
					if buffer[position] != rune('s') {
						goto l275
					}
					position++
					if buffer[position] != rune('u') {
						goto l275
					}
					position++
					if buffer[position] != rune('b') {
						goto l275
					}
					position++
					if buffer[position] != rune('s') {
						goto l275
					}
					position++
					if buffer[position] != rune('e') {
						goto l275
					}
					position++
					if buffer[position] != rune('c') {
						goto l275
					}
					position++
					if buffer[position] != rune('t') {
						goto l275
					}
					position++
					goto l268
				l275:
					position, tokenIndex = position268, tokenIndex268
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
					if buffer[position] != rune('t') {
						goto l264
					}
					position++
					if buffer[position] != rune('r') {
						goto l264
					}
					position++
					if buffer[position] != rune('i') {
						goto l264
					}
					position++
					if buffer[position] != rune('b') {
						goto l264
					}
					position++
				}
			l268:
				{
					position276, tokenIndex276 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l277
					}
					position++
					goto l276
				l277:
					position, tokenIndex = position276, tokenIndex276
					{
						position278, tokenIndex278 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l264
						}
						position, tokenIndex = position278, tokenIndex278
					}
				}
			l276:
				add(ruleRankUninomialNotho, position265)
			}
			return true
		l264:
			position, tokenIndex = position264, tokenIndex264
			return false
		},
		/* 35 Uninomial <- <(UninomialWord (_ Authorship)?)> */
		func() bool {
			position279, tokenIndex279 := position, tokenIndex
			{
				position280 := position
				if !_rules[ruleUninomialWord]() {
					goto l279
				}
				{
					position281, tokenIndex281 := position, tokenIndex
					if !_rules[rule_]() {
						goto l281
					}
					if !_rules[ruleAuthorship]() {
						goto l281
					}
					goto l282
				l281:
					position, tokenIndex = position281, tokenIndex281
				}
			l282:
				add(ruleUninomial, position280)
			}
			return true
		l279:
			position, tokenIndex = position279, tokenIndex279
			return false
		},
		/* 36 UninomialWord <- <(CapWord / TwoLetterGenus)> */
		func() bool {
			position283, tokenIndex283 := position, tokenIndex
			{
				position284 := position
				{
					position285, tokenIndex285 := position, tokenIndex
					if !_rules[ruleCapWord]() {
						goto l286
					}
					goto l285
				l286:
					position, tokenIndex = position285, tokenIndex285
					if !_rules[ruleTwoLetterGenus]() {
						goto l283
					}
				}
			l285:
				add(ruleUninomialWord, position284)
			}
			return true
		l283:
			position, tokenIndex = position283, tokenIndex283
			return false
		},
		/* 37 AbbrGenus <- <(UpperChar LowerChar? '.')> */
		func() bool {
			position287, tokenIndex287 := position, tokenIndex
			{
				position288 := position
				if !_rules[ruleUpperChar]() {
					goto l287
				}
				{
					position289, tokenIndex289 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l289
					}
					goto l290
				l289:
					position, tokenIndex = position289, tokenIndex289
				}
			l290:
				if buffer[position] != rune('.') {
					goto l287
				}
				position++
				add(ruleAbbrGenus, position288)
			}
			return true
		l287:
			position, tokenIndex = position287, tokenIndex287
			return false
		},
		/* 38 CapWord <- <(CapWordWithDash / CapWord1)> */
		func() bool {
			position291, tokenIndex291 := position, tokenIndex
			{
				position292 := position
				{
					position293, tokenIndex293 := position, tokenIndex
					if !_rules[ruleCapWordWithDash]() {
						goto l294
					}
					goto l293
				l294:
					position, tokenIndex = position293, tokenIndex293
					if !_rules[ruleCapWord1]() {
						goto l291
					}
				}
			l293:
				add(ruleCapWord, position292)
			}
			return true
		l291:
			position, tokenIndex = position291, tokenIndex291
			return false
		},
		/* 39 CapWord1 <- <(NameUpperChar NameLowerChar NameLowerChar+ '?'?)> */
		func() bool {
			position295, tokenIndex295 := position, tokenIndex
			{
				position296 := position
				if !_rules[ruleNameUpperChar]() {
					goto l295
				}
				if !_rules[ruleNameLowerChar]() {
					goto l295
				}
				if !_rules[ruleNameLowerChar]() {
					goto l295
				}
			l297:
				{
					position298, tokenIndex298 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l298
					}
					goto l297
				l298:
					position, tokenIndex = position298, tokenIndex298
				}
				{
					position299, tokenIndex299 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l299
					}
					position++
					goto l300
				l299:
					position, tokenIndex = position299, tokenIndex299
				}
			l300:
				add(ruleCapWord1, position296)
			}
			return true
		l295:
			position, tokenIndex = position295, tokenIndex295
			return false
		},
		/* 40 CapWordWithDash <- <(CapWord1 Dash (UpperAfterDash / LowerAfterDash))> */
		func() bool {
			position301, tokenIndex301 := position, tokenIndex
			{
				position302 := position
				if !_rules[ruleCapWord1]() {
					goto l301
				}
				if !_rules[ruleDash]() {
					goto l301
				}
				{
					position303, tokenIndex303 := position, tokenIndex
					if !_rules[ruleUpperAfterDash]() {
						goto l304
					}
					goto l303
				l304:
					position, tokenIndex = position303, tokenIndex303
					if !_rules[ruleLowerAfterDash]() {
						goto l301
					}
				}
			l303:
				add(ruleCapWordWithDash, position302)
			}
			return true
		l301:
			position, tokenIndex = position301, tokenIndex301
			return false
		},
		/* 41 UpperAfterDash <- <CapWord1> */
		func() bool {
			position305, tokenIndex305 := position, tokenIndex
			{
				position306 := position
				if !_rules[ruleCapWord1]() {
					goto l305
				}
				add(ruleUpperAfterDash, position306)
			}
			return true
		l305:
			position, tokenIndex = position305, tokenIndex305
			return false
		},
		/* 42 LowerAfterDash <- <Word1> */
		func() bool {
			position307, tokenIndex307 := position, tokenIndex
			{
				position308 := position
				if !_rules[ruleWord1]() {
					goto l307
				}
				add(ruleLowerAfterDash, position308)
			}
			return true
		l307:
			position, tokenIndex = position307, tokenIndex307
			return false
		},
		/* 43 TwoLetterGenus <- <(('C' 'a') / ('E' 'a') / ('G' 'e') / ('I' 'a') / ('I' 'o') / ('I' 'x') / ('L' 'o') / ('O' 'a') / ('R' 'a') / ('T' 'y') / ('U' 'a') / ('A' 'a') / ('J' 'a') / ('Z' 'u') / ('L' 'a') / ('Q' 'u') / ('A' 's') / ('B' 'a'))> */
		func() bool {
			position309, tokenIndex309 := position, tokenIndex
			{
				position310 := position
				{
					position311, tokenIndex311 := position, tokenIndex
					if buffer[position] != rune('C') {
						goto l312
					}
					position++
					if buffer[position] != rune('a') {
						goto l312
					}
					position++
					goto l311
				l312:
					position, tokenIndex = position311, tokenIndex311
					if buffer[position] != rune('E') {
						goto l313
					}
					position++
					if buffer[position] != rune('a') {
						goto l313
					}
					position++
					goto l311
				l313:
					position, tokenIndex = position311, tokenIndex311
					if buffer[position] != rune('G') {
						goto l314
					}
					position++
					if buffer[position] != rune('e') {
						goto l314
					}
					position++
					goto l311
				l314:
					position, tokenIndex = position311, tokenIndex311
					if buffer[position] != rune('I') {
						goto l315
					}
					position++
					if buffer[position] != rune('a') {
						goto l315
					}
					position++
					goto l311
				l315:
					position, tokenIndex = position311, tokenIndex311
					if buffer[position] != rune('I') {
						goto l316
					}
					position++
					if buffer[position] != rune('o') {
						goto l316
					}
					position++
					goto l311
				l316:
					position, tokenIndex = position311, tokenIndex311
					if buffer[position] != rune('I') {
						goto l317
					}
					position++
					if buffer[position] != rune('x') {
						goto l317
					}
					position++
					goto l311
				l317:
					position, tokenIndex = position311, tokenIndex311
					if buffer[position] != rune('L') {
						goto l318
					}
					position++
					if buffer[position] != rune('o') {
						goto l318
					}
					position++
					goto l311
				l318:
					position, tokenIndex = position311, tokenIndex311
					if buffer[position] != rune('O') {
						goto l319
					}
					position++
					if buffer[position] != rune('a') {
						goto l319
					}
					position++
					goto l311
				l319:
					position, tokenIndex = position311, tokenIndex311
					if buffer[position] != rune('R') {
						goto l320
					}
					position++
					if buffer[position] != rune('a') {
						goto l320
					}
					position++
					goto l311
				l320:
					position, tokenIndex = position311, tokenIndex311
					if buffer[position] != rune('T') {
						goto l321
					}
					position++
					if buffer[position] != rune('y') {
						goto l321
					}
					position++
					goto l311
				l321:
					position, tokenIndex = position311, tokenIndex311
					if buffer[position] != rune('U') {
						goto l322
					}
					position++
					if buffer[position] != rune('a') {
						goto l322
					}
					position++
					goto l311
				l322:
					position, tokenIndex = position311, tokenIndex311
					if buffer[position] != rune('A') {
						goto l323
					}
					position++
					if buffer[position] != rune('a') {
						goto l323
					}
					position++
					goto l311
				l323:
					position, tokenIndex = position311, tokenIndex311
					if buffer[position] != rune('J') {
						goto l324
					}
					position++
					if buffer[position] != rune('a') {
						goto l324
					}
					position++
					goto l311
				l324:
					position, tokenIndex = position311, tokenIndex311
					if buffer[position] != rune('Z') {
						goto l325
					}
					position++
					if buffer[position] != rune('u') {
						goto l325
					}
					position++
					goto l311
				l325:
					position, tokenIndex = position311, tokenIndex311
					if buffer[position] != rune('L') {
						goto l326
					}
					position++
					if buffer[position] != rune('a') {
						goto l326
					}
					position++
					goto l311
				l326:
					position, tokenIndex = position311, tokenIndex311
					if buffer[position] != rune('Q') {
						goto l327
					}
					position++
					if buffer[position] != rune('u') {
						goto l327
					}
					position++
					goto l311
				l327:
					position, tokenIndex = position311, tokenIndex311
					if buffer[position] != rune('A') {
						goto l328
					}
					position++
					if buffer[position] != rune('s') {
						goto l328
					}
					position++
					goto l311
				l328:
					position, tokenIndex = position311, tokenIndex311
					if buffer[position] != rune('B') {
						goto l309
					}
					position++
					if buffer[position] != rune('a') {
						goto l309
					}
					position++
				}
			l311:
				add(ruleTwoLetterGenus, position310)
			}
			return true
		l309:
			position, tokenIndex = position309, tokenIndex309
			return false
		},
		/* 44 Word <- <(!((AuthorPrefix / RankUninomial / Approximation / Word4) SpaceCharEOI) (WordApostr / WordStartsWithDigit / MultiDashedWord / Word2 / Word1) &(SpaceCharEOI / '('))> */
		func() bool {
			position329, tokenIndex329 := position, tokenIndex
			{
				position330 := position
				{
					position331, tokenIndex331 := position, tokenIndex
					{
						position332, tokenIndex332 := position, tokenIndex
						if !_rules[ruleAuthorPrefix]() {
							goto l333
						}
						goto l332
					l333:
						position, tokenIndex = position332, tokenIndex332
						if !_rules[ruleRankUninomial]() {
							goto l334
						}
						goto l332
					l334:
						position, tokenIndex = position332, tokenIndex332
						if !_rules[ruleApproximation]() {
							goto l335
						}
						goto l332
					l335:
						position, tokenIndex = position332, tokenIndex332
						if !_rules[ruleWord4]() {
							goto l331
						}
					}
				l332:
					if !_rules[ruleSpaceCharEOI]() {
						goto l331
					}
					goto l329
				l331:
					position, tokenIndex = position331, tokenIndex331
				}
				{
					position336, tokenIndex336 := position, tokenIndex
					if !_rules[ruleWordApostr]() {
						goto l337
					}
					goto l336
				l337:
					position, tokenIndex = position336, tokenIndex336
					if !_rules[ruleWordStartsWithDigit]() {
						goto l338
					}
					goto l336
				l338:
					position, tokenIndex = position336, tokenIndex336
					if !_rules[ruleMultiDashedWord]() {
						goto l339
					}
					goto l336
				l339:
					position, tokenIndex = position336, tokenIndex336
					if !_rules[ruleWord2]() {
						goto l340
					}
					goto l336
				l340:
					position, tokenIndex = position336, tokenIndex336
					if !_rules[ruleWord1]() {
						goto l329
					}
				}
			l336:
				{
					position341, tokenIndex341 := position, tokenIndex
					{
						position342, tokenIndex342 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l343
						}
						goto l342
					l343:
						position, tokenIndex = position342, tokenIndex342
						if buffer[position] != rune('(') {
							goto l329
						}
						position++
					}
				l342:
					position, tokenIndex = position341, tokenIndex341
				}
				add(ruleWord, position330)
			}
			return true
		l329:
			position, tokenIndex = position329, tokenIndex329
			return false
		},
		/* 45 Word1 <- <((LowerASCII Dash)? NameLowerChar NameLowerChar+)> */
		func() bool {
			position344, tokenIndex344 := position, tokenIndex
			{
				position345 := position
				{
					position346, tokenIndex346 := position, tokenIndex
					if !_rules[ruleLowerASCII]() {
						goto l346
					}
					if !_rules[ruleDash]() {
						goto l346
					}
					goto l347
				l346:
					position, tokenIndex = position346, tokenIndex346
				}
			l347:
				if !_rules[ruleNameLowerChar]() {
					goto l344
				}
				if !_rules[ruleNameLowerChar]() {
					goto l344
				}
			l348:
				{
					position349, tokenIndex349 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l349
					}
					goto l348
				l349:
					position, tokenIndex = position349, tokenIndex349
				}
				add(ruleWord1, position345)
			}
			return true
		l344:
			position, tokenIndex = position344, tokenIndex344
			return false
		},
		/* 46 WordStartsWithDigit <- <(('1' / '2' / '3' / '4' / '5' / '6' / '7' / '8' / '9') Nums? ('.' / Dash)? NameLowerChar NameLowerChar NameLowerChar NameLowerChar+)> */
		func() bool {
			position350, tokenIndex350 := position, tokenIndex
			{
				position351 := position
				{
					position352, tokenIndex352 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l353
					}
					position++
					goto l352
				l353:
					position, tokenIndex = position352, tokenIndex352
					if buffer[position] != rune('2') {
						goto l354
					}
					position++
					goto l352
				l354:
					position, tokenIndex = position352, tokenIndex352
					if buffer[position] != rune('3') {
						goto l355
					}
					position++
					goto l352
				l355:
					position, tokenIndex = position352, tokenIndex352
					if buffer[position] != rune('4') {
						goto l356
					}
					position++
					goto l352
				l356:
					position, tokenIndex = position352, tokenIndex352
					if buffer[position] != rune('5') {
						goto l357
					}
					position++
					goto l352
				l357:
					position, tokenIndex = position352, tokenIndex352
					if buffer[position] != rune('6') {
						goto l358
					}
					position++
					goto l352
				l358:
					position, tokenIndex = position352, tokenIndex352
					if buffer[position] != rune('7') {
						goto l359
					}
					position++
					goto l352
				l359:
					position, tokenIndex = position352, tokenIndex352
					if buffer[position] != rune('8') {
						goto l360
					}
					position++
					goto l352
				l360:
					position, tokenIndex = position352, tokenIndex352
					if buffer[position] != rune('9') {
						goto l350
					}
					position++
				}
			l352:
				{
					position361, tokenIndex361 := position, tokenIndex
					if !_rules[ruleNums]() {
						goto l361
					}
					goto l362
				l361:
					position, tokenIndex = position361, tokenIndex361
				}
			l362:
				{
					position363, tokenIndex363 := position, tokenIndex
					{
						position365, tokenIndex365 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l366
						}
						position++
						goto l365
					l366:
						position, tokenIndex = position365, tokenIndex365
						if !_rules[ruleDash]() {
							goto l363
						}
					}
				l365:
					goto l364
				l363:
					position, tokenIndex = position363, tokenIndex363
				}
			l364:
				if !_rules[ruleNameLowerChar]() {
					goto l350
				}
				if !_rules[ruleNameLowerChar]() {
					goto l350
				}
				if !_rules[ruleNameLowerChar]() {
					goto l350
				}
				if !_rules[ruleNameLowerChar]() {
					goto l350
				}
			l367:
				{
					position368, tokenIndex368 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l368
					}
					goto l367
				l368:
					position, tokenIndex = position368, tokenIndex368
				}
				add(ruleWordStartsWithDigit, position351)
			}
			return true
		l350:
			position, tokenIndex = position350, tokenIndex350
			return false
		},
		/* 47 Word2 <- <(NameLowerChar+ Dash? NameLowerChar+)> */
		func() bool {
			position369, tokenIndex369 := position, tokenIndex
			{
				position370 := position
				if !_rules[ruleNameLowerChar]() {
					goto l369
				}
			l371:
				{
					position372, tokenIndex372 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l372
					}
					goto l371
				l372:
					position, tokenIndex = position372, tokenIndex372
				}
				{
					position373, tokenIndex373 := position, tokenIndex
					if !_rules[ruleDash]() {
						goto l373
					}
					goto l374
				l373:
					position, tokenIndex = position373, tokenIndex373
				}
			l374:
				if !_rules[ruleNameLowerChar]() {
					goto l369
				}
			l375:
				{
					position376, tokenIndex376 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l376
					}
					goto l375
				l376:
					position, tokenIndex = position376, tokenIndex376
				}
				add(ruleWord2, position370)
			}
			return true
		l369:
			position, tokenIndex = position369, tokenIndex369
			return false
		},
		/* 48 WordApostr <- <(NameLowerChar NameLowerChar* Apostrophe Word1)> */
		func() bool {
			position377, tokenIndex377 := position, tokenIndex
			{
				position378 := position
				if !_rules[ruleNameLowerChar]() {
					goto l377
				}
			l379:
				{
					position380, tokenIndex380 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l380
					}
					goto l379
				l380:
					position, tokenIndex = position380, tokenIndex380
				}
				if !_rules[ruleApostrophe]() {
					goto l377
				}
				if !_rules[ruleWord1]() {
					goto l377
				}
				add(ruleWordApostr, position378)
			}
			return true
		l377:
			position, tokenIndex = position377, tokenIndex377
			return false
		},
		/* 49 Word4 <- <(NameLowerChar+ '.' NameLowerChar)> */
		func() bool {
			position381, tokenIndex381 := position, tokenIndex
			{
				position382 := position
				if !_rules[ruleNameLowerChar]() {
					goto l381
				}
			l383:
				{
					position384, tokenIndex384 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l384
					}
					goto l383
				l384:
					position, tokenIndex = position384, tokenIndex384
				}
				if buffer[position] != rune('.') {
					goto l381
				}
				position++
				if !_rules[ruleNameLowerChar]() {
					goto l381
				}
				add(ruleWord4, position382)
			}
			return true
		l381:
			position, tokenIndex = position381, tokenIndex381
			return false
		},
		/* 50 MultiDashedWord <- <(NameLowerChar+ Dash NameLowerChar+ Dash NameLowerChar+ (Dash NameLowerChar+)?)> */
		func() bool {
			position385, tokenIndex385 := position, tokenIndex
			{
				position386 := position
				if !_rules[ruleNameLowerChar]() {
					goto l385
				}
			l387:
				{
					position388, tokenIndex388 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l388
					}
					goto l387
				l388:
					position, tokenIndex = position388, tokenIndex388
				}
				if !_rules[ruleDash]() {
					goto l385
				}
				if !_rules[ruleNameLowerChar]() {
					goto l385
				}
			l389:
				{
					position390, tokenIndex390 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l390
					}
					goto l389
				l390:
					position, tokenIndex = position390, tokenIndex390
				}
				if !_rules[ruleDash]() {
					goto l385
				}
				if !_rules[ruleNameLowerChar]() {
					goto l385
				}
			l391:
				{
					position392, tokenIndex392 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l392
					}
					goto l391
				l392:
					position, tokenIndex = position392, tokenIndex392
				}
				{
					position393, tokenIndex393 := position, tokenIndex
					if !_rules[ruleDash]() {
						goto l393
					}
					if !_rules[ruleNameLowerChar]() {
						goto l393
					}
				l395:
					{
						position396, tokenIndex396 := position, tokenIndex
						if !_rules[ruleNameLowerChar]() {
							goto l396
						}
						goto l395
					l396:
						position, tokenIndex = position396, tokenIndex396
					}
					goto l394
				l393:
					position, tokenIndex = position393, tokenIndex393
				}
			l394:
				add(ruleMultiDashedWord, position386)
			}
			return true
		l385:
			position, tokenIndex = position385, tokenIndex385
			return false
		},
		/* 51 HybridChar <- <'×'> */
		func() bool {
			position397, tokenIndex397 := position, tokenIndex
			{
				position398 := position
				if buffer[position] != rune('×') {
					goto l397
				}
				position++
				add(ruleHybridChar, position398)
			}
			return true
		l397:
			position, tokenIndex = position397, tokenIndex397
			return false
		},
		/* 52 ApproxNameIgnored <- <.*> */
		func() bool {
			{
				position400 := position
			l401:
				{
					position402, tokenIndex402 := position, tokenIndex
					if !matchDot() {
						goto l402
					}
					goto l401
				l402:
					position, tokenIndex = position402, tokenIndex402
				}
				add(ruleApproxNameIgnored, position400)
			}
			return true
		},
		/* 53 Approximation <- <(('s' 'p' '.' _? ('n' 'r' '.')) / ('s' 'p' '.' _? ('a' 'f' 'f' '.')) / ('m' 'o' 'n' 's' 't' '.') / '?' / ((('s' 'p' 'p') / ('n' 'r') / ('s' 'p') / ('a' 'f' 'f') / ('s' 'p' 'e' 'c' 'i' 'e' 's')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position403, tokenIndex403 := position, tokenIndex
			{
				position404 := position
				{
					position405, tokenIndex405 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l406
					}
					position++
					if buffer[position] != rune('p') {
						goto l406
					}
					position++
					if buffer[position] != rune('.') {
						goto l406
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
					if buffer[position] != rune('n') {
						goto l406
					}
					position++
					if buffer[position] != rune('r') {
						goto l406
					}
					position++
					if buffer[position] != rune('.') {
						goto l406
					}
					position++
					goto l405
				l406:
					position, tokenIndex = position405, tokenIndex405
					if buffer[position] != rune('s') {
						goto l409
					}
					position++
					if buffer[position] != rune('p') {
						goto l409
					}
					position++
					if buffer[position] != rune('.') {
						goto l409
					}
					position++
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
					if buffer[position] != rune('a') {
						goto l409
					}
					position++
					if buffer[position] != rune('f') {
						goto l409
					}
					position++
					if buffer[position] != rune('f') {
						goto l409
					}
					position++
					if buffer[position] != rune('.') {
						goto l409
					}
					position++
					goto l405
				l409:
					position, tokenIndex = position405, tokenIndex405
					if buffer[position] != rune('m') {
						goto l412
					}
					position++
					if buffer[position] != rune('o') {
						goto l412
					}
					position++
					if buffer[position] != rune('n') {
						goto l412
					}
					position++
					if buffer[position] != rune('s') {
						goto l412
					}
					position++
					if buffer[position] != rune('t') {
						goto l412
					}
					position++
					if buffer[position] != rune('.') {
						goto l412
					}
					position++
					goto l405
				l412:
					position, tokenIndex = position405, tokenIndex405
					if buffer[position] != rune('?') {
						goto l413
					}
					position++
					goto l405
				l413:
					position, tokenIndex = position405, tokenIndex405
					{
						position414, tokenIndex414 := position, tokenIndex
						if buffer[position] != rune('s') {
							goto l415
						}
						position++
						if buffer[position] != rune('p') {
							goto l415
						}
						position++
						if buffer[position] != rune('p') {
							goto l415
						}
						position++
						goto l414
					l415:
						position, tokenIndex = position414, tokenIndex414
						if buffer[position] != rune('n') {
							goto l416
						}
						position++
						if buffer[position] != rune('r') {
							goto l416
						}
						position++
						goto l414
					l416:
						position, tokenIndex = position414, tokenIndex414
						if buffer[position] != rune('s') {
							goto l417
						}
						position++
						if buffer[position] != rune('p') {
							goto l417
						}
						position++
						goto l414
					l417:
						position, tokenIndex = position414, tokenIndex414
						if buffer[position] != rune('a') {
							goto l418
						}
						position++
						if buffer[position] != rune('f') {
							goto l418
						}
						position++
						if buffer[position] != rune('f') {
							goto l418
						}
						position++
						goto l414
					l418:
						position, tokenIndex = position414, tokenIndex414
						if buffer[position] != rune('s') {
							goto l403
						}
						position++
						if buffer[position] != rune('p') {
							goto l403
						}
						position++
						if buffer[position] != rune('e') {
							goto l403
						}
						position++
						if buffer[position] != rune('c') {
							goto l403
						}
						position++
						if buffer[position] != rune('i') {
							goto l403
						}
						position++
						if buffer[position] != rune('e') {
							goto l403
						}
						position++
						if buffer[position] != rune('s') {
							goto l403
						}
						position++
					}
				l414:
					{
						position419, tokenIndex419 := position, tokenIndex
						{
							position421, tokenIndex421 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l420
							}
							position, tokenIndex = position421, tokenIndex421
						}
						goto l419
					l420:
						position, tokenIndex = position419, tokenIndex419
						if buffer[position] != rune('.') {
							goto l403
						}
						position++
					}
				l419:
				}
			l405:
				add(ruleApproximation, position404)
			}
			return true
		l403:
			position, tokenIndex = position403, tokenIndex403
			return false
		},
		/* 54 Authorship <- <((AuthorshipCombo / OriginalAuthorship) &(SpaceCharEOI / ';' / ','))> */
		func() bool {
			position422, tokenIndex422 := position, tokenIndex
			{
				position423 := position
				{
					position424, tokenIndex424 := position, tokenIndex
					if !_rules[ruleAuthorshipCombo]() {
						goto l425
					}
					goto l424
				l425:
					position, tokenIndex = position424, tokenIndex424
					if !_rules[ruleOriginalAuthorship]() {
						goto l422
					}
				}
			l424:
				{
					position426, tokenIndex426 := position, tokenIndex
					{
						position427, tokenIndex427 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l428
						}
						goto l427
					l428:
						position, tokenIndex = position427, tokenIndex427
						if buffer[position] != rune(';') {
							goto l429
						}
						position++
						goto l427
					l429:
						position, tokenIndex = position427, tokenIndex427
						if buffer[position] != rune(',') {
							goto l422
						}
						position++
					}
				l427:
					position, tokenIndex = position426, tokenIndex426
				}
				add(ruleAuthorship, position423)
			}
			return true
		l422:
			position, tokenIndex = position422, tokenIndex422
			return false
		},
		/* 55 AuthorshipCombo <- <(OriginalAuthorshipComb (_? CombinationAuthorship)?)> */
		func() bool {
			position430, tokenIndex430 := position, tokenIndex
			{
				position431 := position
				if !_rules[ruleOriginalAuthorshipComb]() {
					goto l430
				}
				{
					position432, tokenIndex432 := position, tokenIndex
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
					if !_rules[ruleCombinationAuthorship]() {
						goto l432
					}
					goto l433
				l432:
					position, tokenIndex = position432, tokenIndex432
				}
			l433:
				add(ruleAuthorshipCombo, position431)
			}
			return true
		l430:
			position, tokenIndex = position430, tokenIndex430
			return false
		},
		/* 56 OriginalAuthorship <- <AuthorsGroup> */
		func() bool {
			position436, tokenIndex436 := position, tokenIndex
			{
				position437 := position
				if !_rules[ruleAuthorsGroup]() {
					goto l436
				}
				add(ruleOriginalAuthorship, position437)
			}
			return true
		l436:
			position, tokenIndex = position436, tokenIndex436
			return false
		},
		/* 57 OriginalAuthorshipComb <- <(BasionymAuthorshipYearMisformed / BasionymAuthorship / BasionymAuthorshipMissingParens)> */
		func() bool {
			position438, tokenIndex438 := position, tokenIndex
			{
				position439 := position
				{
					position440, tokenIndex440 := position, tokenIndex
					if !_rules[ruleBasionymAuthorshipYearMisformed]() {
						goto l441
					}
					goto l440
				l441:
					position, tokenIndex = position440, tokenIndex440
					if !_rules[ruleBasionymAuthorship]() {
						goto l442
					}
					goto l440
				l442:
					position, tokenIndex = position440, tokenIndex440
					if !_rules[ruleBasionymAuthorshipMissingParens]() {
						goto l438
					}
				}
			l440:
				add(ruleOriginalAuthorshipComb, position439)
			}
			return true
		l438:
			position, tokenIndex = position438, tokenIndex438
			return false
		},
		/* 58 CombinationAuthorship <- <AuthorsGroup> */
		func() bool {
			position443, tokenIndex443 := position, tokenIndex
			{
				position444 := position
				if !_rules[ruleAuthorsGroup]() {
					goto l443
				}
				add(ruleCombinationAuthorship, position444)
			}
			return true
		l443:
			position, tokenIndex = position443, tokenIndex443
			return false
		},
		/* 59 BasionymAuthorshipMissingParens <- <(MissingParensStart / MissingParensEnd)> */
		func() bool {
			position445, tokenIndex445 := position, tokenIndex
			{
				position446 := position
				{
					position447, tokenIndex447 := position, tokenIndex
					if !_rules[ruleMissingParensStart]() {
						goto l448
					}
					goto l447
				l448:
					position, tokenIndex = position447, tokenIndex447
					if !_rules[ruleMissingParensEnd]() {
						goto l445
					}
				}
			l447:
				add(ruleBasionymAuthorshipMissingParens, position446)
			}
			return true
		l445:
			position, tokenIndex = position445, tokenIndex445
			return false
		},
		/* 60 MissingParensStart <- <('(' _? AuthorsGroup)> */
		func() bool {
			position449, tokenIndex449 := position, tokenIndex
			{
				position450 := position
				if buffer[position] != rune('(') {
					goto l449
				}
				position++
				{
					position451, tokenIndex451 := position, tokenIndex
					if !_rules[rule_]() {
						goto l451
					}
					goto l452
				l451:
					position, tokenIndex = position451, tokenIndex451
				}
			l452:
				if !_rules[ruleAuthorsGroup]() {
					goto l449
				}
				add(ruleMissingParensStart, position450)
			}
			return true
		l449:
			position, tokenIndex = position449, tokenIndex449
			return false
		},
		/* 61 MissingParensEnd <- <(AuthorsGroup _? ')')> */
		func() bool {
			position453, tokenIndex453 := position, tokenIndex
			{
				position454 := position
				if !_rules[ruleAuthorsGroup]() {
					goto l453
				}
				{
					position455, tokenIndex455 := position, tokenIndex
					if !_rules[rule_]() {
						goto l455
					}
					goto l456
				l455:
					position, tokenIndex = position455, tokenIndex455
				}
			l456:
				if buffer[position] != rune(')') {
					goto l453
				}
				position++
				add(ruleMissingParensEnd, position454)
			}
			return true
		l453:
			position, tokenIndex = position453, tokenIndex453
			return false
		},
		/* 62 BasionymAuthorshipYearMisformed <- <('(' _? AuthorsGroup _? ')' (_? ',')? _? Year)> */
		func() bool {
			position457, tokenIndex457 := position, tokenIndex
			{
				position458 := position
				if buffer[position] != rune('(') {
					goto l457
				}
				position++
				{
					position459, tokenIndex459 := position, tokenIndex
					if !_rules[rule_]() {
						goto l459
					}
					goto l460
				l459:
					position, tokenIndex = position459, tokenIndex459
				}
			l460:
				if !_rules[ruleAuthorsGroup]() {
					goto l457
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
				if buffer[position] != rune(')') {
					goto l457
				}
				position++
				{
					position463, tokenIndex463 := position, tokenIndex
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
					if buffer[position] != rune(',') {
						goto l463
					}
					position++
					goto l464
				l463:
					position, tokenIndex = position463, tokenIndex463
				}
			l464:
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
				if !_rules[ruleYear]() {
					goto l457
				}
				add(ruleBasionymAuthorshipYearMisformed, position458)
			}
			return true
		l457:
			position, tokenIndex = position457, tokenIndex457
			return false
		},
		/* 63 BasionymAuthorship <- <(BasionymAuthorship1 / BasionymAuthorship2Parens)> */
		func() bool {
			position469, tokenIndex469 := position, tokenIndex
			{
				position470 := position
				{
					position471, tokenIndex471 := position, tokenIndex
					if !_rules[ruleBasionymAuthorship1]() {
						goto l472
					}
					goto l471
				l472:
					position, tokenIndex = position471, tokenIndex471
					if !_rules[ruleBasionymAuthorship2Parens]() {
						goto l469
					}
				}
			l471:
				add(ruleBasionymAuthorship, position470)
			}
			return true
		l469:
			position, tokenIndex = position469, tokenIndex469
			return false
		},
		/* 64 BasionymAuthorship1 <- <('(' _? AuthorsGroup _? ')')> */
		func() bool {
			position473, tokenIndex473 := position, tokenIndex
			{
				position474 := position
				if buffer[position] != rune('(') {
					goto l473
				}
				position++
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
				if !_rules[ruleAuthorsGroup]() {
					goto l473
				}
				{
					position477, tokenIndex477 := position, tokenIndex
					if !_rules[rule_]() {
						goto l477
					}
					goto l478
				l477:
					position, tokenIndex = position477, tokenIndex477
				}
			l478:
				if buffer[position] != rune(')') {
					goto l473
				}
				position++
				add(ruleBasionymAuthorship1, position474)
			}
			return true
		l473:
			position, tokenIndex = position473, tokenIndex473
			return false
		},
		/* 65 BasionymAuthorship2Parens <- <('(' _? '(' _? AuthorsGroup _? ')' _? ')')> */
		func() bool {
			position479, tokenIndex479 := position, tokenIndex
			{
				position480 := position
				if buffer[position] != rune('(') {
					goto l479
				}
				position++
				{
					position481, tokenIndex481 := position, tokenIndex
					if !_rules[rule_]() {
						goto l481
					}
					goto l482
				l481:
					position, tokenIndex = position481, tokenIndex481
				}
			l482:
				if buffer[position] != rune('(') {
					goto l479
				}
				position++
				{
					position483, tokenIndex483 := position, tokenIndex
					if !_rules[rule_]() {
						goto l483
					}
					goto l484
				l483:
					position, tokenIndex = position483, tokenIndex483
				}
			l484:
				if !_rules[ruleAuthorsGroup]() {
					goto l479
				}
				{
					position485, tokenIndex485 := position, tokenIndex
					if !_rules[rule_]() {
						goto l485
					}
					goto l486
				l485:
					position, tokenIndex = position485, tokenIndex485
				}
			l486:
				if buffer[position] != rune(')') {
					goto l479
				}
				position++
				{
					position487, tokenIndex487 := position, tokenIndex
					if !_rules[rule_]() {
						goto l487
					}
					goto l488
				l487:
					position, tokenIndex = position487, tokenIndex487
				}
			l488:
				if buffer[position] != rune(')') {
					goto l479
				}
				position++
				add(ruleBasionymAuthorship2Parens, position480)
			}
			return true
		l479:
			position, tokenIndex = position479, tokenIndex479
			return false
		},
		/* 66 AuthorsGroup <- <(AuthorsTeam (_ (AuthorEmend / AuthorEx) AuthorsTeam)?)> */
		func() bool {
			position489, tokenIndex489 := position, tokenIndex
			{
				position490 := position
				if !_rules[ruleAuthorsTeam]() {
					goto l489
				}
				{
					position491, tokenIndex491 := position, tokenIndex
					if !_rules[rule_]() {
						goto l491
					}
					{
						position493, tokenIndex493 := position, tokenIndex
						if !_rules[ruleAuthorEmend]() {
							goto l494
						}
						goto l493
					l494:
						position, tokenIndex = position493, tokenIndex493
						if !_rules[ruleAuthorEx]() {
							goto l491
						}
					}
				l493:
					if !_rules[ruleAuthorsTeam]() {
						goto l491
					}
					goto l492
				l491:
					position, tokenIndex = position491, tokenIndex491
				}
			l492:
				add(ruleAuthorsGroup, position490)
			}
			return true
		l489:
			position, tokenIndex = position489, tokenIndex489
			return false
		},
		/* 67 AuthorsTeam <- <(Author (AuthorSep Author)* (_? ','? _? Year)?)> */
		func() bool {
			position495, tokenIndex495 := position, tokenIndex
			{
				position496 := position
				if !_rules[ruleAuthor]() {
					goto l495
				}
			l497:
				{
					position498, tokenIndex498 := position, tokenIndex
					if !_rules[ruleAuthorSep]() {
						goto l498
					}
					if !_rules[ruleAuthor]() {
						goto l498
					}
					goto l497
				l498:
					position, tokenIndex = position498, tokenIndex498
				}
				{
					position499, tokenIndex499 := position, tokenIndex
					{
						position501, tokenIndex501 := position, tokenIndex
						if !_rules[rule_]() {
							goto l501
						}
						goto l502
					l501:
						position, tokenIndex = position501, tokenIndex501
					}
				l502:
					{
						position503, tokenIndex503 := position, tokenIndex
						if buffer[position] != rune(',') {
							goto l503
						}
						position++
						goto l504
					l503:
						position, tokenIndex = position503, tokenIndex503
					}
				l504:
					{
						position505, tokenIndex505 := position, tokenIndex
						if !_rules[rule_]() {
							goto l505
						}
						goto l506
					l505:
						position, tokenIndex = position505, tokenIndex505
					}
				l506:
					if !_rules[ruleYear]() {
						goto l499
					}
					goto l500
				l499:
					position, tokenIndex = position499, tokenIndex499
				}
			l500:
				add(ruleAuthorsTeam, position496)
			}
			return true
		l495:
			position, tokenIndex = position495, tokenIndex495
			return false
		},
		/* 68 AuthorSep <- <(AuthorSep1 / AuthorSep2)> */
		func() bool {
			position507, tokenIndex507 := position, tokenIndex
			{
				position508 := position
				{
					position509, tokenIndex509 := position, tokenIndex
					if !_rules[ruleAuthorSep1]() {
						goto l510
					}
					goto l509
				l510:
					position, tokenIndex = position509, tokenIndex509
					if !_rules[ruleAuthorSep2]() {
						goto l507
					}
				}
			l509:
				add(ruleAuthorSep, position508)
			}
			return true
		l507:
			position, tokenIndex = position507, tokenIndex507
			return false
		},
		/* 69 AuthorSep1 <- <(_? (',' _)? ('&' / ('e' 't') / ('a' 'n' 'd') / ('a' 'p' 'u' 'd')) _?)> */
		func() bool {
			position511, tokenIndex511 := position, tokenIndex
			{
				position512 := position
				{
					position513, tokenIndex513 := position, tokenIndex
					if !_rules[rule_]() {
						goto l513
					}
					goto l514
				l513:
					position, tokenIndex = position513, tokenIndex513
				}
			l514:
				{
					position515, tokenIndex515 := position, tokenIndex
					if buffer[position] != rune(',') {
						goto l515
					}
					position++
					if !_rules[rule_]() {
						goto l515
					}
					goto l516
				l515:
					position, tokenIndex = position515, tokenIndex515
				}
			l516:
				{
					position517, tokenIndex517 := position, tokenIndex
					if buffer[position] != rune('&') {
						goto l518
					}
					position++
					goto l517
				l518:
					position, tokenIndex = position517, tokenIndex517
					if buffer[position] != rune('e') {
						goto l519
					}
					position++
					if buffer[position] != rune('t') {
						goto l519
					}
					position++
					goto l517
				l519:
					position, tokenIndex = position517, tokenIndex517
					if buffer[position] != rune('a') {
						goto l520
					}
					position++
					if buffer[position] != rune('n') {
						goto l520
					}
					position++
					if buffer[position] != rune('d') {
						goto l520
					}
					position++
					goto l517
				l520:
					position, tokenIndex = position517, tokenIndex517
					if buffer[position] != rune('a') {
						goto l511
					}
					position++
					if buffer[position] != rune('p') {
						goto l511
					}
					position++
					if buffer[position] != rune('u') {
						goto l511
					}
					position++
					if buffer[position] != rune('d') {
						goto l511
					}
					position++
				}
			l517:
				{
					position521, tokenIndex521 := position, tokenIndex
					if !_rules[rule_]() {
						goto l521
					}
					goto l522
				l521:
					position, tokenIndex = position521, tokenIndex521
				}
			l522:
				add(ruleAuthorSep1, position512)
			}
			return true
		l511:
			position, tokenIndex = position511, tokenIndex511
			return false
		},
		/* 70 AuthorSep2 <- <(_? ',' _?)> */
		func() bool {
			position523, tokenIndex523 := position, tokenIndex
			{
				position524 := position
				{
					position525, tokenIndex525 := position, tokenIndex
					if !_rules[rule_]() {
						goto l525
					}
					goto l526
				l525:
					position, tokenIndex = position525, tokenIndex525
				}
			l526:
				if buffer[position] != rune(',') {
					goto l523
				}
				position++
				{
					position527, tokenIndex527 := position, tokenIndex
					if !_rules[rule_]() {
						goto l527
					}
					goto l528
				l527:
					position, tokenIndex = position527, tokenIndex527
				}
			l528:
				add(ruleAuthorSep2, position524)
			}
			return true
		l523:
			position, tokenIndex = position523, tokenIndex523
			return false
		},
		/* 71 AuthorEx <- <((('e' 'x' '.'?) / ('i' 'n')) _)> */
		func() bool {
			position529, tokenIndex529 := position, tokenIndex
			{
				position530 := position
				{
					position531, tokenIndex531 := position, tokenIndex
					if buffer[position] != rune('e') {
						goto l532
					}
					position++
					if buffer[position] != rune('x') {
						goto l532
					}
					position++
					{
						position533, tokenIndex533 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l533
						}
						position++
						goto l534
					l533:
						position, tokenIndex = position533, tokenIndex533
					}
				l534:
					goto l531
				l532:
					position, tokenIndex = position531, tokenIndex531
					if buffer[position] != rune('i') {
						goto l529
					}
					position++
					if buffer[position] != rune('n') {
						goto l529
					}
					position++
				}
			l531:
				if !_rules[rule_]() {
					goto l529
				}
				add(ruleAuthorEx, position530)
			}
			return true
		l529:
			position, tokenIndex = position529, tokenIndex529
			return false
		},
		/* 72 AuthorEmend <- <('e' 'm' 'e' 'n' 'd' '.'? _)> */
		func() bool {
			position535, tokenIndex535 := position, tokenIndex
			{
				position536 := position
				if buffer[position] != rune('e') {
					goto l535
				}
				position++
				if buffer[position] != rune('m') {
					goto l535
				}
				position++
				if buffer[position] != rune('e') {
					goto l535
				}
				position++
				if buffer[position] != rune('n') {
					goto l535
				}
				position++
				if buffer[position] != rune('d') {
					goto l535
				}
				position++
				{
					position537, tokenIndex537 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l537
					}
					position++
					goto l538
				l537:
					position, tokenIndex = position537, tokenIndex537
				}
			l538:
				if !_rules[rule_]() {
					goto l535
				}
				add(ruleAuthorEmend, position536)
			}
			return true
		l535:
			position, tokenIndex = position535, tokenIndex535
			return false
		},
		/* 73 Author <- <(Author1 / Author2 / UnknownAuthor)> */
		func() bool {
			position539, tokenIndex539 := position, tokenIndex
			{
				position540 := position
				{
					position541, tokenIndex541 := position, tokenIndex
					if !_rules[ruleAuthor1]() {
						goto l542
					}
					goto l541
				l542:
					position, tokenIndex = position541, tokenIndex541
					if !_rules[ruleAuthor2]() {
						goto l543
					}
					goto l541
				l543:
					position, tokenIndex = position541, tokenIndex541
					if !_rules[ruleUnknownAuthor]() {
						goto l539
					}
				}
			l541:
				add(ruleAuthor, position540)
			}
			return true
		l539:
			position, tokenIndex = position539, tokenIndex539
			return false
		},
		/* 74 Author1 <- <(Author2 _? Filius)> */
		func() bool {
			position544, tokenIndex544 := position, tokenIndex
			{
				position545 := position
				if !_rules[ruleAuthor2]() {
					goto l544
				}
				{
					position546, tokenIndex546 := position, tokenIndex
					if !_rules[rule_]() {
						goto l546
					}
					goto l547
				l546:
					position, tokenIndex = position546, tokenIndex546
				}
			l547:
				if !_rules[ruleFilius]() {
					goto l544
				}
				add(ruleAuthor1, position545)
			}
			return true
		l544:
			position, tokenIndex = position544, tokenIndex544
			return false
		},
		/* 75 Author2 <- <(AuthorWord (_? AuthorWord)*)> */
		func() bool {
			position548, tokenIndex548 := position, tokenIndex
			{
				position549 := position
				if !_rules[ruleAuthorWord]() {
					goto l548
				}
			l550:
				{
					position551, tokenIndex551 := position, tokenIndex
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
					if !_rules[ruleAuthorWord]() {
						goto l551
					}
					goto l550
				l551:
					position, tokenIndex = position551, tokenIndex551
				}
				add(ruleAuthor2, position549)
			}
			return true
		l548:
			position, tokenIndex = position548, tokenIndex548
			return false
		},
		/* 76 UnknownAuthor <- <('?' / ((('a' 'u' 'c' 't') / ('a' 'n' 'o' 'n')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position554, tokenIndex554 := position, tokenIndex
			{
				position555 := position
				{
					position556, tokenIndex556 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l557
					}
					position++
					goto l556
				l557:
					position, tokenIndex = position556, tokenIndex556
					{
						position558, tokenIndex558 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l559
						}
						position++
						if buffer[position] != rune('u') {
							goto l559
						}
						position++
						if buffer[position] != rune('c') {
							goto l559
						}
						position++
						if buffer[position] != rune('t') {
							goto l559
						}
						position++
						goto l558
					l559:
						position, tokenIndex = position558, tokenIndex558
						if buffer[position] != rune('a') {
							goto l554
						}
						position++
						if buffer[position] != rune('n') {
							goto l554
						}
						position++
						if buffer[position] != rune('o') {
							goto l554
						}
						position++
						if buffer[position] != rune('n') {
							goto l554
						}
						position++
					}
				l558:
					{
						position560, tokenIndex560 := position, tokenIndex
						{
							position562, tokenIndex562 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l561
							}
							position, tokenIndex = position562, tokenIndex562
						}
						goto l560
					l561:
						position, tokenIndex = position560, tokenIndex560
						if buffer[position] != rune('.') {
							goto l554
						}
						position++
					}
				l560:
				}
			l556:
				add(ruleUnknownAuthor, position555)
			}
			return true
		l554:
			position, tokenIndex = position554, tokenIndex554
			return false
		},
		/* 77 AuthorWord <- <(!(('b' / 'B') ('o' / 'O') ('l' / 'L') ('d' / 'D') ':') (AuthorEtAl / AuthorWord2 / AuthorWord3 / AuthorPrefix))> */
		func() bool {
			position563, tokenIndex563 := position, tokenIndex
			{
				position564 := position
				{
					position565, tokenIndex565 := position, tokenIndex
					{
						position566, tokenIndex566 := position, tokenIndex
						if buffer[position] != rune('b') {
							goto l567
						}
						position++
						goto l566
					l567:
						position, tokenIndex = position566, tokenIndex566
						if buffer[position] != rune('B') {
							goto l565
						}
						position++
					}
				l566:
					{
						position568, tokenIndex568 := position, tokenIndex
						if buffer[position] != rune('o') {
							goto l569
						}
						position++
						goto l568
					l569:
						position, tokenIndex = position568, tokenIndex568
						if buffer[position] != rune('O') {
							goto l565
						}
						position++
					}
				l568:
					{
						position570, tokenIndex570 := position, tokenIndex
						if buffer[position] != rune('l') {
							goto l571
						}
						position++
						goto l570
					l571:
						position, tokenIndex = position570, tokenIndex570
						if buffer[position] != rune('L') {
							goto l565
						}
						position++
					}
				l570:
					{
						position572, tokenIndex572 := position, tokenIndex
						if buffer[position] != rune('d') {
							goto l573
						}
						position++
						goto l572
					l573:
						position, tokenIndex = position572, tokenIndex572
						if buffer[position] != rune('D') {
							goto l565
						}
						position++
					}
				l572:
					if buffer[position] != rune(':') {
						goto l565
					}
					position++
					goto l563
				l565:
					position, tokenIndex = position565, tokenIndex565
				}
				{
					position574, tokenIndex574 := position, tokenIndex
					if !_rules[ruleAuthorEtAl]() {
						goto l575
					}
					goto l574
				l575:
					position, tokenIndex = position574, tokenIndex574
					if !_rules[ruleAuthorWord2]() {
						goto l576
					}
					goto l574
				l576:
					position, tokenIndex = position574, tokenIndex574
					if !_rules[ruleAuthorWord3]() {
						goto l577
					}
					goto l574
				l577:
					position, tokenIndex = position574, tokenIndex574
					if !_rules[ruleAuthorPrefix]() {
						goto l563
					}
				}
			l574:
				add(ruleAuthorWord, position564)
			}
			return true
		l563:
			position, tokenIndex = position563, tokenIndex563
			return false
		},
		/* 78 AuthorEtAl <- <(('a' 'r' 'g' '.') / ('e' 't' ' ' 'a' 'l' '.' '{' '?' '}') / ((('e' 't') / '&') (' ' 'a' 'l') '.'?))> */
		func() bool {
			position578, tokenIndex578 := position, tokenIndex
			{
				position579 := position
				{
					position580, tokenIndex580 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l581
					}
					position++
					if buffer[position] != rune('r') {
						goto l581
					}
					position++
					if buffer[position] != rune('g') {
						goto l581
					}
					position++
					if buffer[position] != rune('.') {
						goto l581
					}
					position++
					goto l580
				l581:
					position, tokenIndex = position580, tokenIndex580
					if buffer[position] != rune('e') {
						goto l582
					}
					position++
					if buffer[position] != rune('t') {
						goto l582
					}
					position++
					if buffer[position] != rune(' ') {
						goto l582
					}
					position++
					if buffer[position] != rune('a') {
						goto l582
					}
					position++
					if buffer[position] != rune('l') {
						goto l582
					}
					position++
					if buffer[position] != rune('.') {
						goto l582
					}
					position++
					if buffer[position] != rune('{') {
						goto l582
					}
					position++
					if buffer[position] != rune('?') {
						goto l582
					}
					position++
					if buffer[position] != rune('}') {
						goto l582
					}
					position++
					goto l580
				l582:
					position, tokenIndex = position580, tokenIndex580
					{
						position583, tokenIndex583 := position, tokenIndex
						if buffer[position] != rune('e') {
							goto l584
						}
						position++
						if buffer[position] != rune('t') {
							goto l584
						}
						position++
						goto l583
					l584:
						position, tokenIndex = position583, tokenIndex583
						if buffer[position] != rune('&') {
							goto l578
						}
						position++
					}
				l583:
					if buffer[position] != rune(' ') {
						goto l578
					}
					position++
					if buffer[position] != rune('a') {
						goto l578
					}
					position++
					if buffer[position] != rune('l') {
						goto l578
					}
					position++
					{
						position585, tokenIndex585 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l585
						}
						position++
						goto l586
					l585:
						position, tokenIndex = position585, tokenIndex585
					}
				l586:
				}
			l580:
				add(ruleAuthorEtAl, position579)
			}
			return true
		l578:
			position, tokenIndex = position578, tokenIndex578
			return false
		},
		/* 79 AuthorWord2 <- <(AuthorWord3 Dash AuthorWordSoft)> */
		func() bool {
			position587, tokenIndex587 := position, tokenIndex
			{
				position588 := position
				if !_rules[ruleAuthorWord3]() {
					goto l587
				}
				if !_rules[ruleDash]() {
					goto l587
				}
				if !_rules[ruleAuthorWordSoft]() {
					goto l587
				}
				add(ruleAuthorWord2, position588)
			}
			return true
		l587:
			position, tokenIndex = position587, tokenIndex587
			return false
		},
		/* 80 AuthorWord3 <- <(AuthorPrefixGlued? (AllCapsAuthorWord / CapAuthorWord) '.'?)> */
		func() bool {
			position589, tokenIndex589 := position, tokenIndex
			{
				position590 := position
				{
					position591, tokenIndex591 := position, tokenIndex
					if !_rules[ruleAuthorPrefixGlued]() {
						goto l591
					}
					goto l592
				l591:
					position, tokenIndex = position591, tokenIndex591
				}
			l592:
				{
					position593, tokenIndex593 := position, tokenIndex
					if !_rules[ruleAllCapsAuthorWord]() {
						goto l594
					}
					goto l593
				l594:
					position, tokenIndex = position593, tokenIndex593
					if !_rules[ruleCapAuthorWord]() {
						goto l589
					}
				}
			l593:
				{
					position595, tokenIndex595 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l595
					}
					position++
					goto l596
				l595:
					position, tokenIndex = position595, tokenIndex595
				}
			l596:
				add(ruleAuthorWord3, position590)
			}
			return true
		l589:
			position, tokenIndex = position589, tokenIndex589
			return false
		},
		/* 81 AuthorWordSoft <- <(((AuthorUpperChar (AuthorUpperChar+ / AuthorLowerChar+)) / AuthorLowerChar+) '.'?)> */
		func() bool {
			position597, tokenIndex597 := position, tokenIndex
			{
				position598 := position
				{
					position599, tokenIndex599 := position, tokenIndex
					if !_rules[ruleAuthorUpperChar]() {
						goto l600
					}
					{
						position601, tokenIndex601 := position, tokenIndex
						if !_rules[ruleAuthorUpperChar]() {
							goto l602
						}
					l603:
						{
							position604, tokenIndex604 := position, tokenIndex
							if !_rules[ruleAuthorUpperChar]() {
								goto l604
							}
							goto l603
						l604:
							position, tokenIndex = position604, tokenIndex604
						}
						goto l601
					l602:
						position, tokenIndex = position601, tokenIndex601
						if !_rules[ruleAuthorLowerChar]() {
							goto l600
						}
					l605:
						{
							position606, tokenIndex606 := position, tokenIndex
							if !_rules[ruleAuthorLowerChar]() {
								goto l606
							}
							goto l605
						l606:
							position, tokenIndex = position606, tokenIndex606
						}
					}
				l601:
					goto l599
				l600:
					position, tokenIndex = position599, tokenIndex599
					if !_rules[ruleAuthorLowerChar]() {
						goto l597
					}
				l607:
					{
						position608, tokenIndex608 := position, tokenIndex
						if !_rules[ruleAuthorLowerChar]() {
							goto l608
						}
						goto l607
					l608:
						position, tokenIndex = position608, tokenIndex608
					}
				}
			l599:
				{
					position609, tokenIndex609 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l609
					}
					position++
					goto l610
				l609:
					position, tokenIndex = position609, tokenIndex609
				}
			l610:
				add(ruleAuthorWordSoft, position598)
			}
			return true
		l597:
			position, tokenIndex = position597, tokenIndex597
			return false
		},
		/* 82 CapAuthorWord <- <(AuthorUpperChar AuthorLowerChar*)> */
		func() bool {
			position611, tokenIndex611 := position, tokenIndex
			{
				position612 := position
				if !_rules[ruleAuthorUpperChar]() {
					goto l611
				}
			l613:
				{
					position614, tokenIndex614 := position, tokenIndex
					if !_rules[ruleAuthorLowerChar]() {
						goto l614
					}
					goto l613
				l614:
					position, tokenIndex = position614, tokenIndex614
				}
				add(ruleCapAuthorWord, position612)
			}
			return true
		l611:
			position, tokenIndex = position611, tokenIndex611
			return false
		},
		/* 83 AllCapsAuthorWord <- <(AuthorUpperChar AuthorUpperChar+)> */
		func() bool {
			position615, tokenIndex615 := position, tokenIndex
			{
				position616 := position
				if !_rules[ruleAuthorUpperChar]() {
					goto l615
				}
				if !_rules[ruleAuthorUpperChar]() {
					goto l615
				}
			l617:
				{
					position618, tokenIndex618 := position, tokenIndex
					if !_rules[ruleAuthorUpperChar]() {
						goto l618
					}
					goto l617
				l618:
					position, tokenIndex = position618, tokenIndex618
				}
				add(ruleAllCapsAuthorWord, position616)
			}
			return true
		l615:
			position, tokenIndex = position615, tokenIndex615
			return false
		},
		/* 84 Filius <- <(('f' '.') / ('f' 'i' 'l' '.') / ('f' 'i' 'l' 'i' 'u' 's'))> */
		func() bool {
			position619, tokenIndex619 := position, tokenIndex
			{
				position620 := position
				{
					position621, tokenIndex621 := position, tokenIndex
					if buffer[position] != rune('f') {
						goto l622
					}
					position++
					if buffer[position] != rune('.') {
						goto l622
					}
					position++
					goto l621
				l622:
					position, tokenIndex = position621, tokenIndex621
					if buffer[position] != rune('f') {
						goto l623
					}
					position++
					if buffer[position] != rune('i') {
						goto l623
					}
					position++
					if buffer[position] != rune('l') {
						goto l623
					}
					position++
					if buffer[position] != rune('.') {
						goto l623
					}
					position++
					goto l621
				l623:
					position, tokenIndex = position621, tokenIndex621
					if buffer[position] != rune('f') {
						goto l619
					}
					position++
					if buffer[position] != rune('i') {
						goto l619
					}
					position++
					if buffer[position] != rune('l') {
						goto l619
					}
					position++
					if buffer[position] != rune('i') {
						goto l619
					}
					position++
					if buffer[position] != rune('u') {
						goto l619
					}
					position++
					if buffer[position] != rune('s') {
						goto l619
					}
					position++
				}
			l621:
				add(ruleFilius, position620)
			}
			return true
		l619:
			position, tokenIndex = position619, tokenIndex619
			return false
		},
		/* 85 AuthorPrefixGlued <- <(('d' / 'O' / 'L') Apostrophe)> */
		func() bool {
			position624, tokenIndex624 := position, tokenIndex
			{
				position625 := position
				{
					position626, tokenIndex626 := position, tokenIndex
					if buffer[position] != rune('d') {
						goto l627
					}
					position++
					goto l626
				l627:
					position, tokenIndex = position626, tokenIndex626
					if buffer[position] != rune('O') {
						goto l628
					}
					position++
					goto l626
				l628:
					position, tokenIndex = position626, tokenIndex626
					if buffer[position] != rune('L') {
						goto l624
					}
					position++
				}
			l626:
				if !_rules[ruleApostrophe]() {
					goto l624
				}
				add(ruleAuthorPrefixGlued, position625)
			}
			return true
		l624:
			position, tokenIndex = position624, tokenIndex624
			return false
		},
		/* 86 AuthorPrefix <- <(AuthorPrefix1 / AuthorPrefix2)> */
		func() bool {
			position629, tokenIndex629 := position, tokenIndex
			{
				position630 := position
				{
					position631, tokenIndex631 := position, tokenIndex
					if !_rules[ruleAuthorPrefix1]() {
						goto l632
					}
					goto l631
				l632:
					position, tokenIndex = position631, tokenIndex631
					if !_rules[ruleAuthorPrefix2]() {
						goto l629
					}
				}
			l631:
				add(ruleAuthorPrefix, position630)
			}
			return true
		l629:
			position, tokenIndex = position629, tokenIndex629
			return false
		},
		/* 87 AuthorPrefix2 <- <(('v' '.' (_? ('d' '.'))?) / (Apostrophe 't'))> */
		func() bool {
			position633, tokenIndex633 := position, tokenIndex
			{
				position634 := position
				{
					position635, tokenIndex635 := position, tokenIndex
					if buffer[position] != rune('v') {
						goto l636
					}
					position++
					if buffer[position] != rune('.') {
						goto l636
					}
					position++
					{
						position637, tokenIndex637 := position, tokenIndex
						{
							position639, tokenIndex639 := position, tokenIndex
							if !_rules[rule_]() {
								goto l639
							}
							goto l640
						l639:
							position, tokenIndex = position639, tokenIndex639
						}
					l640:
						if buffer[position] != rune('d') {
							goto l637
						}
						position++
						if buffer[position] != rune('.') {
							goto l637
						}
						position++
						goto l638
					l637:
						position, tokenIndex = position637, tokenIndex637
					}
				l638:
					goto l635
				l636:
					position, tokenIndex = position635, tokenIndex635
					if !_rules[ruleApostrophe]() {
						goto l633
					}
					if buffer[position] != rune('t') {
						goto l633
					}
					position++
				}
			l635:
				add(ruleAuthorPrefix2, position634)
			}
			return true
		l633:
			position, tokenIndex = position633, tokenIndex633
			return false
		},
		/* 88 AuthorPrefix1 <- <((('a' 'b') / ('a' 'f') / ('b' 'i' 's') / ('d' 'a') / ('d' 'e' 'r') / ('d' 'e' 's') / ('d' 'e' 'n') / ('d' 'e' 'l') / ('d' 'e' 'l' 'l' 'a') / ('d' 'e' 'l' 'a') / ('d' 'e') / ('d' 'i') / ('d' 'u') / ('e' 'l') / ('l' 'a') / ('l' 'e') / ('t' 'e' 'r') / ('v' 'a' 'n') / ('d' Apostrophe) / ('i' 'n' Apostrophe 't') / ('z' 'u' 'r') / ('v' 'o' 'n' (_ (('d' '.') / ('d' 'e' 'm')))?) / ('v' (_ 'd')?)) &_)> */
		func() bool {
			position641, tokenIndex641 := position, tokenIndex
			{
				position642 := position
				{
					position643, tokenIndex643 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l644
					}
					position++
					if buffer[position] != rune('b') {
						goto l644
					}
					position++
					goto l643
				l644:
					position, tokenIndex = position643, tokenIndex643
					if buffer[position] != rune('a') {
						goto l645
					}
					position++
					if buffer[position] != rune('f') {
						goto l645
					}
					position++
					goto l643
				l645:
					position, tokenIndex = position643, tokenIndex643
					if buffer[position] != rune('b') {
						goto l646
					}
					position++
					if buffer[position] != rune('i') {
						goto l646
					}
					position++
					if buffer[position] != rune('s') {
						goto l646
					}
					position++
					goto l643
				l646:
					position, tokenIndex = position643, tokenIndex643
					if buffer[position] != rune('d') {
						goto l647
					}
					position++
					if buffer[position] != rune('a') {
						goto l647
					}
					position++
					goto l643
				l647:
					position, tokenIndex = position643, tokenIndex643
					if buffer[position] != rune('d') {
						goto l648
					}
					position++
					if buffer[position] != rune('e') {
						goto l648
					}
					position++
					if buffer[position] != rune('r') {
						goto l648
					}
					position++
					goto l643
				l648:
					position, tokenIndex = position643, tokenIndex643
					if buffer[position] != rune('d') {
						goto l649
					}
					position++
					if buffer[position] != rune('e') {
						goto l649
					}
					position++
					if buffer[position] != rune('s') {
						goto l649
					}
					position++
					goto l643
				l649:
					position, tokenIndex = position643, tokenIndex643
					if buffer[position] != rune('d') {
						goto l650
					}
					position++
					if buffer[position] != rune('e') {
						goto l650
					}
					position++
					if buffer[position] != rune('n') {
						goto l650
					}
					position++
					goto l643
				l650:
					position, tokenIndex = position643, tokenIndex643
					if buffer[position] != rune('d') {
						goto l651
					}
					position++
					if buffer[position] != rune('e') {
						goto l651
					}
					position++
					if buffer[position] != rune('l') {
						goto l651
					}
					position++
					goto l643
				l651:
					position, tokenIndex = position643, tokenIndex643
					if buffer[position] != rune('d') {
						goto l652
					}
					position++
					if buffer[position] != rune('e') {
						goto l652
					}
					position++
					if buffer[position] != rune('l') {
						goto l652
					}
					position++
					if buffer[position] != rune('l') {
						goto l652
					}
					position++
					if buffer[position] != rune('a') {
						goto l652
					}
					position++
					goto l643
				l652:
					position, tokenIndex = position643, tokenIndex643
					if buffer[position] != rune('d') {
						goto l653
					}
					position++
					if buffer[position] != rune('e') {
						goto l653
					}
					position++
					if buffer[position] != rune('l') {
						goto l653
					}
					position++
					if buffer[position] != rune('a') {
						goto l653
					}
					position++
					goto l643
				l653:
					position, tokenIndex = position643, tokenIndex643
					if buffer[position] != rune('d') {
						goto l654
					}
					position++
					if buffer[position] != rune('e') {
						goto l654
					}
					position++
					goto l643
				l654:
					position, tokenIndex = position643, tokenIndex643
					if buffer[position] != rune('d') {
						goto l655
					}
					position++
					if buffer[position] != rune('i') {
						goto l655
					}
					position++
					goto l643
				l655:
					position, tokenIndex = position643, tokenIndex643
					if buffer[position] != rune('d') {
						goto l656
					}
					position++
					if buffer[position] != rune('u') {
						goto l656
					}
					position++
					goto l643
				l656:
					position, tokenIndex = position643, tokenIndex643
					if buffer[position] != rune('e') {
						goto l657
					}
					position++
					if buffer[position] != rune('l') {
						goto l657
					}
					position++
					goto l643
				l657:
					position, tokenIndex = position643, tokenIndex643
					if buffer[position] != rune('l') {
						goto l658
					}
					position++
					if buffer[position] != rune('a') {
						goto l658
					}
					position++
					goto l643
				l658:
					position, tokenIndex = position643, tokenIndex643
					if buffer[position] != rune('l') {
						goto l659
					}
					position++
					if buffer[position] != rune('e') {
						goto l659
					}
					position++
					goto l643
				l659:
					position, tokenIndex = position643, tokenIndex643
					if buffer[position] != rune('t') {
						goto l660
					}
					position++
					if buffer[position] != rune('e') {
						goto l660
					}
					position++
					if buffer[position] != rune('r') {
						goto l660
					}
					position++
					goto l643
				l660:
					position, tokenIndex = position643, tokenIndex643
					if buffer[position] != rune('v') {
						goto l661
					}
					position++
					if buffer[position] != rune('a') {
						goto l661
					}
					position++
					if buffer[position] != rune('n') {
						goto l661
					}
					position++
					goto l643
				l661:
					position, tokenIndex = position643, tokenIndex643
					if buffer[position] != rune('d') {
						goto l662
					}
					position++
					if !_rules[ruleApostrophe]() {
						goto l662
					}
					goto l643
				l662:
					position, tokenIndex = position643, tokenIndex643
					if buffer[position] != rune('i') {
						goto l663
					}
					position++
					if buffer[position] != rune('n') {
						goto l663
					}
					position++
					if !_rules[ruleApostrophe]() {
						goto l663
					}
					if buffer[position] != rune('t') {
						goto l663
					}
					position++
					goto l643
				l663:
					position, tokenIndex = position643, tokenIndex643
					if buffer[position] != rune('z') {
						goto l664
					}
					position++
					if buffer[position] != rune('u') {
						goto l664
					}
					position++
					if buffer[position] != rune('r') {
						goto l664
					}
					position++
					goto l643
				l664:
					position, tokenIndex = position643, tokenIndex643
					if buffer[position] != rune('v') {
						goto l665
					}
					position++
					if buffer[position] != rune('o') {
						goto l665
					}
					position++
					if buffer[position] != rune('n') {
						goto l665
					}
					position++
					{
						position666, tokenIndex666 := position, tokenIndex
						if !_rules[rule_]() {
							goto l666
						}
						{
							position668, tokenIndex668 := position, tokenIndex
							if buffer[position] != rune('d') {
								goto l669
							}
							position++
							if buffer[position] != rune('.') {
								goto l669
							}
							position++
							goto l668
						l669:
							position, tokenIndex = position668, tokenIndex668
							if buffer[position] != rune('d') {
								goto l666
							}
							position++
							if buffer[position] != rune('e') {
								goto l666
							}
							position++
							if buffer[position] != rune('m') {
								goto l666
							}
							position++
						}
					l668:
						goto l667
					l666:
						position, tokenIndex = position666, tokenIndex666
					}
				l667:
					goto l643
				l665:
					position, tokenIndex = position643, tokenIndex643
					if buffer[position] != rune('v') {
						goto l641
					}
					position++
					{
						position670, tokenIndex670 := position, tokenIndex
						if !_rules[rule_]() {
							goto l670
						}
						if buffer[position] != rune('d') {
							goto l670
						}
						position++
						goto l671
					l670:
						position, tokenIndex = position670, tokenIndex670
					}
				l671:
				}
			l643:
				{
					position672, tokenIndex672 := position, tokenIndex
					if !_rules[rule_]() {
						goto l641
					}
					position, tokenIndex = position672, tokenIndex672
				}
				add(ruleAuthorPrefix1, position642)
			}
			return true
		l641:
			position, tokenIndex = position641, tokenIndex641
			return false
		},
		/* 89 AuthorUpperChar <- <(UpperASCII / MiscodedChar / ('À' / 'Á' / 'Â' / 'Ã' / 'Ä' / 'Å' / 'Æ' / 'Ç' / 'È' / 'É' / 'Ê' / 'Ë' / 'Ì' / 'Í' / 'Î' / 'Ï' / 'Ð' / 'Ñ' / 'Ò' / 'Ó' / 'Ô' / 'Õ' / 'Ö' / 'Ø' / 'Ù' / 'Ú' / 'Û' / 'Ü' / 'Ý' / 'Ć' / 'Č' / 'Ď' / 'İ' / 'Ķ' / 'Ĺ' / 'ĺ' / 'Ľ' / 'ľ' / 'Ł' / 'ł' / 'Ņ' / 'Ō' / 'Ő' / 'Œ' / 'Ř' / 'Ś' / 'Ŝ' / 'Ş' / 'Š' / 'Ÿ' / 'Ź' / 'Ż' / 'Ž' / 'ƒ' / 'Ǿ' / 'Ș' / 'Ț'))> */
		func() bool {
			position673, tokenIndex673 := position, tokenIndex
			{
				position674 := position
				{
					position675, tokenIndex675 := position, tokenIndex
					if !_rules[ruleUpperASCII]() {
						goto l676
					}
					goto l675
				l676:
					position, tokenIndex = position675, tokenIndex675
					if !_rules[ruleMiscodedChar]() {
						goto l677
					}
					goto l675
				l677:
					position, tokenIndex = position675, tokenIndex675
					{
						position678, tokenIndex678 := position, tokenIndex
						if buffer[position] != rune('À') {
							goto l679
						}
						position++
						goto l678
					l679:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Á') {
							goto l680
						}
						position++
						goto l678
					l680:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Â') {
							goto l681
						}
						position++
						goto l678
					l681:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ã') {
							goto l682
						}
						position++
						goto l678
					l682:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ä') {
							goto l683
						}
						position++
						goto l678
					l683:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Å') {
							goto l684
						}
						position++
						goto l678
					l684:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Æ') {
							goto l685
						}
						position++
						goto l678
					l685:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ç') {
							goto l686
						}
						position++
						goto l678
					l686:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('È') {
							goto l687
						}
						position++
						goto l678
					l687:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('É') {
							goto l688
						}
						position++
						goto l678
					l688:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ê') {
							goto l689
						}
						position++
						goto l678
					l689:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ë') {
							goto l690
						}
						position++
						goto l678
					l690:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ì') {
							goto l691
						}
						position++
						goto l678
					l691:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Í') {
							goto l692
						}
						position++
						goto l678
					l692:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Î') {
							goto l693
						}
						position++
						goto l678
					l693:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ï') {
							goto l694
						}
						position++
						goto l678
					l694:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ð') {
							goto l695
						}
						position++
						goto l678
					l695:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ñ') {
							goto l696
						}
						position++
						goto l678
					l696:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ò') {
							goto l697
						}
						position++
						goto l678
					l697:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ó') {
							goto l698
						}
						position++
						goto l678
					l698:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ô') {
							goto l699
						}
						position++
						goto l678
					l699:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Õ') {
							goto l700
						}
						position++
						goto l678
					l700:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ö') {
							goto l701
						}
						position++
						goto l678
					l701:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ø') {
							goto l702
						}
						position++
						goto l678
					l702:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ù') {
							goto l703
						}
						position++
						goto l678
					l703:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ú') {
							goto l704
						}
						position++
						goto l678
					l704:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Û') {
							goto l705
						}
						position++
						goto l678
					l705:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ü') {
							goto l706
						}
						position++
						goto l678
					l706:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ý') {
							goto l707
						}
						position++
						goto l678
					l707:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ć') {
							goto l708
						}
						position++
						goto l678
					l708:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Č') {
							goto l709
						}
						position++
						goto l678
					l709:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ď') {
							goto l710
						}
						position++
						goto l678
					l710:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('İ') {
							goto l711
						}
						position++
						goto l678
					l711:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ķ') {
							goto l712
						}
						position++
						goto l678
					l712:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ĺ') {
							goto l713
						}
						position++
						goto l678
					l713:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('ĺ') {
							goto l714
						}
						position++
						goto l678
					l714:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ľ') {
							goto l715
						}
						position++
						goto l678
					l715:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('ľ') {
							goto l716
						}
						position++
						goto l678
					l716:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ł') {
							goto l717
						}
						position++
						goto l678
					l717:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('ł') {
							goto l718
						}
						position++
						goto l678
					l718:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ņ') {
							goto l719
						}
						position++
						goto l678
					l719:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ō') {
							goto l720
						}
						position++
						goto l678
					l720:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ő') {
							goto l721
						}
						position++
						goto l678
					l721:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Œ') {
							goto l722
						}
						position++
						goto l678
					l722:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ř') {
							goto l723
						}
						position++
						goto l678
					l723:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ś') {
							goto l724
						}
						position++
						goto l678
					l724:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ŝ') {
							goto l725
						}
						position++
						goto l678
					l725:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ş') {
							goto l726
						}
						position++
						goto l678
					l726:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Š') {
							goto l727
						}
						position++
						goto l678
					l727:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ÿ') {
							goto l728
						}
						position++
						goto l678
					l728:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ź') {
							goto l729
						}
						position++
						goto l678
					l729:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ż') {
							goto l730
						}
						position++
						goto l678
					l730:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ž') {
							goto l731
						}
						position++
						goto l678
					l731:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('ƒ') {
							goto l732
						}
						position++
						goto l678
					l732:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ǿ') {
							goto l733
						}
						position++
						goto l678
					l733:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ș') {
							goto l734
						}
						position++
						goto l678
					l734:
						position, tokenIndex = position678, tokenIndex678
						if buffer[position] != rune('Ț') {
							goto l673
						}
						position++
					}
				l678:
				}
			l675:
				add(ruleAuthorUpperChar, position674)
			}
			return true
		l673:
			position, tokenIndex = position673, tokenIndex673
			return false
		},
		/* 90 AuthorLowerChar <- <(LowerASCII / MiscodedChar / ('à' / 'á' / 'â' / 'ã' / 'ä' / 'å' / 'æ' / 'ç' / 'è' / 'é' / 'ê' / 'ë' / 'ì' / 'í' / 'î' / 'ï' / 'ð' / 'ñ' / 'ò' / 'ó' / 'ó' / 'ô' / 'õ' / 'ö' / 'ø' / 'ù' / 'ú' / 'û' / 'ü' / 'ý' / 'ÿ' / 'ā' / 'ă' / 'ą' / 'ć' / 'ĉ' / 'č' / 'ď' / 'đ' / '\'' / 'ē' / 'ĕ' / 'ė' / 'ę' / 'ě' / 'ğ' / 'ī' / 'ĭ' / 'İ' / 'ı' / 'ĺ' / 'ľ' / 'ł' / 'ń' / 'ņ' / 'ň' / 'ŏ' / 'ő' / 'œ' / 'ŕ' / 'ř' / 'ś' / 'ş' / 'š' / 'ţ' / 'ť' / 'ũ' / 'ū' / 'ŭ' / 'ů' / 'ű' / 'ź' / 'ż' / 'ž' / 'ſ' / 'ǎ' / 'ǔ' / 'ǧ' / 'ș' / 'ț' / 'ȳ' / 'ß'))> */
		func() bool {
			position735, tokenIndex735 := position, tokenIndex
			{
				position736 := position
				{
					position737, tokenIndex737 := position, tokenIndex
					if !_rules[ruleLowerASCII]() {
						goto l738
					}
					goto l737
				l738:
					position, tokenIndex = position737, tokenIndex737
					if !_rules[ruleMiscodedChar]() {
						goto l739
					}
					goto l737
				l739:
					position, tokenIndex = position737, tokenIndex737
					{
						position740, tokenIndex740 := position, tokenIndex
						if buffer[position] != rune('à') {
							goto l741
						}
						position++
						goto l740
					l741:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('á') {
							goto l742
						}
						position++
						goto l740
					l742:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('â') {
							goto l743
						}
						position++
						goto l740
					l743:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ã') {
							goto l744
						}
						position++
						goto l740
					l744:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ä') {
							goto l745
						}
						position++
						goto l740
					l745:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('å') {
							goto l746
						}
						position++
						goto l740
					l746:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('æ') {
							goto l747
						}
						position++
						goto l740
					l747:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ç') {
							goto l748
						}
						position++
						goto l740
					l748:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('è') {
							goto l749
						}
						position++
						goto l740
					l749:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('é') {
							goto l750
						}
						position++
						goto l740
					l750:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ê') {
							goto l751
						}
						position++
						goto l740
					l751:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ë') {
							goto l752
						}
						position++
						goto l740
					l752:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ì') {
							goto l753
						}
						position++
						goto l740
					l753:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('í') {
							goto l754
						}
						position++
						goto l740
					l754:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('î') {
							goto l755
						}
						position++
						goto l740
					l755:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ï') {
							goto l756
						}
						position++
						goto l740
					l756:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ð') {
							goto l757
						}
						position++
						goto l740
					l757:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ñ') {
							goto l758
						}
						position++
						goto l740
					l758:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ò') {
							goto l759
						}
						position++
						goto l740
					l759:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ó') {
							goto l760
						}
						position++
						goto l740
					l760:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ó') {
							goto l761
						}
						position++
						goto l740
					l761:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ô') {
							goto l762
						}
						position++
						goto l740
					l762:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('õ') {
							goto l763
						}
						position++
						goto l740
					l763:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ö') {
							goto l764
						}
						position++
						goto l740
					l764:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ø') {
							goto l765
						}
						position++
						goto l740
					l765:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ù') {
							goto l766
						}
						position++
						goto l740
					l766:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ú') {
							goto l767
						}
						position++
						goto l740
					l767:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('û') {
							goto l768
						}
						position++
						goto l740
					l768:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ü') {
							goto l769
						}
						position++
						goto l740
					l769:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ý') {
							goto l770
						}
						position++
						goto l740
					l770:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ÿ') {
							goto l771
						}
						position++
						goto l740
					l771:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ā') {
							goto l772
						}
						position++
						goto l740
					l772:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ă') {
							goto l773
						}
						position++
						goto l740
					l773:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ą') {
							goto l774
						}
						position++
						goto l740
					l774:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ć') {
							goto l775
						}
						position++
						goto l740
					l775:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ĉ') {
							goto l776
						}
						position++
						goto l740
					l776:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('č') {
							goto l777
						}
						position++
						goto l740
					l777:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ď') {
							goto l778
						}
						position++
						goto l740
					l778:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('đ') {
							goto l779
						}
						position++
						goto l740
					l779:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('\'') {
							goto l780
						}
						position++
						goto l740
					l780:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ē') {
							goto l781
						}
						position++
						goto l740
					l781:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ĕ') {
							goto l782
						}
						position++
						goto l740
					l782:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ė') {
							goto l783
						}
						position++
						goto l740
					l783:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ę') {
							goto l784
						}
						position++
						goto l740
					l784:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ě') {
							goto l785
						}
						position++
						goto l740
					l785:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ğ') {
							goto l786
						}
						position++
						goto l740
					l786:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ī') {
							goto l787
						}
						position++
						goto l740
					l787:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ĭ') {
							goto l788
						}
						position++
						goto l740
					l788:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('İ') {
							goto l789
						}
						position++
						goto l740
					l789:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ı') {
							goto l790
						}
						position++
						goto l740
					l790:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ĺ') {
							goto l791
						}
						position++
						goto l740
					l791:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ľ') {
							goto l792
						}
						position++
						goto l740
					l792:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ł') {
							goto l793
						}
						position++
						goto l740
					l793:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ń') {
							goto l794
						}
						position++
						goto l740
					l794:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ņ') {
							goto l795
						}
						position++
						goto l740
					l795:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ň') {
							goto l796
						}
						position++
						goto l740
					l796:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ŏ') {
							goto l797
						}
						position++
						goto l740
					l797:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ő') {
							goto l798
						}
						position++
						goto l740
					l798:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('œ') {
							goto l799
						}
						position++
						goto l740
					l799:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ŕ') {
							goto l800
						}
						position++
						goto l740
					l800:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ř') {
							goto l801
						}
						position++
						goto l740
					l801:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ś') {
							goto l802
						}
						position++
						goto l740
					l802:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ş') {
							goto l803
						}
						position++
						goto l740
					l803:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('š') {
							goto l804
						}
						position++
						goto l740
					l804:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ţ') {
							goto l805
						}
						position++
						goto l740
					l805:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ť') {
							goto l806
						}
						position++
						goto l740
					l806:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ũ') {
							goto l807
						}
						position++
						goto l740
					l807:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ū') {
							goto l808
						}
						position++
						goto l740
					l808:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ŭ') {
							goto l809
						}
						position++
						goto l740
					l809:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ů') {
							goto l810
						}
						position++
						goto l740
					l810:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ű') {
							goto l811
						}
						position++
						goto l740
					l811:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ź') {
							goto l812
						}
						position++
						goto l740
					l812:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ż') {
							goto l813
						}
						position++
						goto l740
					l813:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ž') {
							goto l814
						}
						position++
						goto l740
					l814:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ſ') {
							goto l815
						}
						position++
						goto l740
					l815:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ǎ') {
							goto l816
						}
						position++
						goto l740
					l816:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ǔ') {
							goto l817
						}
						position++
						goto l740
					l817:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ǧ') {
							goto l818
						}
						position++
						goto l740
					l818:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ș') {
							goto l819
						}
						position++
						goto l740
					l819:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ț') {
							goto l820
						}
						position++
						goto l740
					l820:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ȳ') {
							goto l821
						}
						position++
						goto l740
					l821:
						position, tokenIndex = position740, tokenIndex740
						if buffer[position] != rune('ß') {
							goto l735
						}
						position++
					}
				l740:
				}
			l737:
				add(ruleAuthorLowerChar, position736)
			}
			return true
		l735:
			position, tokenIndex = position735, tokenIndex735
			return false
		},
		/* 91 Year <- <(YearRange / YearApprox / YearWithParens / YearWithPage / YearWithDot / YearWithChar / YearNum)> */
		func() bool {
			position822, tokenIndex822 := position, tokenIndex
			{
				position823 := position
				{
					position824, tokenIndex824 := position, tokenIndex
					if !_rules[ruleYearRange]() {
						goto l825
					}
					goto l824
				l825:
					position, tokenIndex = position824, tokenIndex824
					if !_rules[ruleYearApprox]() {
						goto l826
					}
					goto l824
				l826:
					position, tokenIndex = position824, tokenIndex824
					if !_rules[ruleYearWithParens]() {
						goto l827
					}
					goto l824
				l827:
					position, tokenIndex = position824, tokenIndex824
					if !_rules[ruleYearWithPage]() {
						goto l828
					}
					goto l824
				l828:
					position, tokenIndex = position824, tokenIndex824
					if !_rules[ruleYearWithDot]() {
						goto l829
					}
					goto l824
				l829:
					position, tokenIndex = position824, tokenIndex824
					if !_rules[ruleYearWithChar]() {
						goto l830
					}
					goto l824
				l830:
					position, tokenIndex = position824, tokenIndex824
					if !_rules[ruleYearNum]() {
						goto l822
					}
				}
			l824:
				add(ruleYear, position823)
			}
			return true
		l822:
			position, tokenIndex = position822, tokenIndex822
			return false
		},
		/* 92 YearRange <- <(YearNum Dash (Nums+ ('a' / 'b' / 'c' / 'd' / 'e' / 'f' / 'g' / 'h' / 'i' / 'j' / 'k' / 'l' / 'm' / 'n' / 'o' / 'p' / 'q' / 'r' / 's' / 't' / 'u' / 'v' / 'w' / 'x' / 'y' / 'z' / '?')*))> */
		func() bool {
			position831, tokenIndex831 := position, tokenIndex
			{
				position832 := position
				if !_rules[ruleYearNum]() {
					goto l831
				}
				if !_rules[ruleDash]() {
					goto l831
				}
				if !_rules[ruleNums]() {
					goto l831
				}
			l833:
				{
					position834, tokenIndex834 := position, tokenIndex
					if !_rules[ruleNums]() {
						goto l834
					}
					goto l833
				l834:
					position, tokenIndex = position834, tokenIndex834
				}
			l835:
				{
					position836, tokenIndex836 := position, tokenIndex
					{
						position837, tokenIndex837 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l838
						}
						position++
						goto l837
					l838:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('b') {
							goto l839
						}
						position++
						goto l837
					l839:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('c') {
							goto l840
						}
						position++
						goto l837
					l840:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('d') {
							goto l841
						}
						position++
						goto l837
					l841:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('e') {
							goto l842
						}
						position++
						goto l837
					l842:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('f') {
							goto l843
						}
						position++
						goto l837
					l843:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('g') {
							goto l844
						}
						position++
						goto l837
					l844:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('h') {
							goto l845
						}
						position++
						goto l837
					l845:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('i') {
							goto l846
						}
						position++
						goto l837
					l846:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('j') {
							goto l847
						}
						position++
						goto l837
					l847:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('k') {
							goto l848
						}
						position++
						goto l837
					l848:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('l') {
							goto l849
						}
						position++
						goto l837
					l849:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('m') {
							goto l850
						}
						position++
						goto l837
					l850:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('n') {
							goto l851
						}
						position++
						goto l837
					l851:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('o') {
							goto l852
						}
						position++
						goto l837
					l852:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('p') {
							goto l853
						}
						position++
						goto l837
					l853:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('q') {
							goto l854
						}
						position++
						goto l837
					l854:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('r') {
							goto l855
						}
						position++
						goto l837
					l855:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('s') {
							goto l856
						}
						position++
						goto l837
					l856:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('t') {
							goto l857
						}
						position++
						goto l837
					l857:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('u') {
							goto l858
						}
						position++
						goto l837
					l858:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('v') {
							goto l859
						}
						position++
						goto l837
					l859:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('w') {
							goto l860
						}
						position++
						goto l837
					l860:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('x') {
							goto l861
						}
						position++
						goto l837
					l861:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('y') {
							goto l862
						}
						position++
						goto l837
					l862:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('z') {
							goto l863
						}
						position++
						goto l837
					l863:
						position, tokenIndex = position837, tokenIndex837
						if buffer[position] != rune('?') {
							goto l836
						}
						position++
					}
				l837:
					goto l835
				l836:
					position, tokenIndex = position836, tokenIndex836
				}
				add(ruleYearRange, position832)
			}
			return true
		l831:
			position, tokenIndex = position831, tokenIndex831
			return false
		},
		/* 93 YearWithDot <- <(YearNum '.')> */
		func() bool {
			position864, tokenIndex864 := position, tokenIndex
			{
				position865 := position
				if !_rules[ruleYearNum]() {
					goto l864
				}
				if buffer[position] != rune('.') {
					goto l864
				}
				position++
				add(ruleYearWithDot, position865)
			}
			return true
		l864:
			position, tokenIndex = position864, tokenIndex864
			return false
		},
		/* 94 YearApprox <- <('[' _? YearNum _? ']')> */
		func() bool {
			position866, tokenIndex866 := position, tokenIndex
			{
				position867 := position
				if buffer[position] != rune('[') {
					goto l866
				}
				position++
				{
					position868, tokenIndex868 := position, tokenIndex
					if !_rules[rule_]() {
						goto l868
					}
					goto l869
				l868:
					position, tokenIndex = position868, tokenIndex868
				}
			l869:
				if !_rules[ruleYearNum]() {
					goto l866
				}
				{
					position870, tokenIndex870 := position, tokenIndex
					if !_rules[rule_]() {
						goto l870
					}
					goto l871
				l870:
					position, tokenIndex = position870, tokenIndex870
				}
			l871:
				if buffer[position] != rune(']') {
					goto l866
				}
				position++
				add(ruleYearApprox, position867)
			}
			return true
		l866:
			position, tokenIndex = position866, tokenIndex866
			return false
		},
		/* 95 YearWithPage <- <((YearWithChar / YearNum) _? ':' _? Nums+)> */
		func() bool {
			position872, tokenIndex872 := position, tokenIndex
			{
				position873 := position
				{
					position874, tokenIndex874 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l875
					}
					goto l874
				l875:
					position, tokenIndex = position874, tokenIndex874
					if !_rules[ruleYearNum]() {
						goto l872
					}
				}
			l874:
				{
					position876, tokenIndex876 := position, tokenIndex
					if !_rules[rule_]() {
						goto l876
					}
					goto l877
				l876:
					position, tokenIndex = position876, tokenIndex876
				}
			l877:
				if buffer[position] != rune(':') {
					goto l872
				}
				position++
				{
					position878, tokenIndex878 := position, tokenIndex
					if !_rules[rule_]() {
						goto l878
					}
					goto l879
				l878:
					position, tokenIndex = position878, tokenIndex878
				}
			l879:
				if !_rules[ruleNums]() {
					goto l872
				}
			l880:
				{
					position881, tokenIndex881 := position, tokenIndex
					if !_rules[ruleNums]() {
						goto l881
					}
					goto l880
				l881:
					position, tokenIndex = position881, tokenIndex881
				}
				add(ruleYearWithPage, position873)
			}
			return true
		l872:
			position, tokenIndex = position872, tokenIndex872
			return false
		},
		/* 96 YearWithParens <- <('(' (YearWithChar / YearNum) ')')> */
		func() bool {
			position882, tokenIndex882 := position, tokenIndex
			{
				position883 := position
				if buffer[position] != rune('(') {
					goto l882
				}
				position++
				{
					position884, tokenIndex884 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l885
					}
					goto l884
				l885:
					position, tokenIndex = position884, tokenIndex884
					if !_rules[ruleYearNum]() {
						goto l882
					}
				}
			l884:
				if buffer[position] != rune(')') {
					goto l882
				}
				position++
				add(ruleYearWithParens, position883)
			}
			return true
		l882:
			position, tokenIndex = position882, tokenIndex882
			return false
		},
		/* 97 YearWithChar <- <(YearNum LowerASCII Action0)> */
		func() bool {
			position886, tokenIndex886 := position, tokenIndex
			{
				position887 := position
				if !_rules[ruleYearNum]() {
					goto l886
				}
				if !_rules[ruleLowerASCII]() {
					goto l886
				}
				if !_rules[ruleAction0]() {
					goto l886
				}
				add(ruleYearWithChar, position887)
			}
			return true
		l886:
			position, tokenIndex = position886, tokenIndex886
			return false
		},
		/* 98 YearNum <- <(('1' / '2') ('0' / '7' / '8' / '9') Nums (Nums / '?') '?'*)> */
		func() bool {
			position888, tokenIndex888 := position, tokenIndex
			{
				position889 := position
				{
					position890, tokenIndex890 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l891
					}
					position++
					goto l890
				l891:
					position, tokenIndex = position890, tokenIndex890
					if buffer[position] != rune('2') {
						goto l888
					}
					position++
				}
			l890:
				{
					position892, tokenIndex892 := position, tokenIndex
					if buffer[position] != rune('0') {
						goto l893
					}
					position++
					goto l892
				l893:
					position, tokenIndex = position892, tokenIndex892
					if buffer[position] != rune('7') {
						goto l894
					}
					position++
					goto l892
				l894:
					position, tokenIndex = position892, tokenIndex892
					if buffer[position] != rune('8') {
						goto l895
					}
					position++
					goto l892
				l895:
					position, tokenIndex = position892, tokenIndex892
					if buffer[position] != rune('9') {
						goto l888
					}
					position++
				}
			l892:
				if !_rules[ruleNums]() {
					goto l888
				}
				{
					position896, tokenIndex896 := position, tokenIndex
					if !_rules[ruleNums]() {
						goto l897
					}
					goto l896
				l897:
					position, tokenIndex = position896, tokenIndex896
					if buffer[position] != rune('?') {
						goto l888
					}
					position++
				}
			l896:
			l898:
				{
					position899, tokenIndex899 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l899
					}
					position++
					goto l898
				l899:
					position, tokenIndex = position899, tokenIndex899
				}
				add(ruleYearNum, position889)
			}
			return true
		l888:
			position, tokenIndex = position888, tokenIndex888
			return false
		},
		/* 99 NameUpperChar <- <(UpperChar / UpperCharExtended)> */
		func() bool {
			position900, tokenIndex900 := position, tokenIndex
			{
				position901 := position
				{
					position902, tokenIndex902 := position, tokenIndex
					if !_rules[ruleUpperChar]() {
						goto l903
					}
					goto l902
				l903:
					position, tokenIndex = position902, tokenIndex902
					if !_rules[ruleUpperCharExtended]() {
						goto l900
					}
				}
			l902:
				add(ruleNameUpperChar, position901)
			}
			return true
		l900:
			position, tokenIndex = position900, tokenIndex900
			return false
		},
		/* 100 UpperCharExtended <- <('Æ' / 'Œ' / 'Ö')> */
		func() bool {
			position904, tokenIndex904 := position, tokenIndex
			{
				position905 := position
				{
					position906, tokenIndex906 := position, tokenIndex
					if buffer[position] != rune('Æ') {
						goto l907
					}
					position++
					goto l906
				l907:
					position, tokenIndex = position906, tokenIndex906
					if buffer[position] != rune('Œ') {
						goto l908
					}
					position++
					goto l906
				l908:
					position, tokenIndex = position906, tokenIndex906
					if buffer[position] != rune('Ö') {
						goto l904
					}
					position++
				}
			l906:
				add(ruleUpperCharExtended, position905)
			}
			return true
		l904:
			position, tokenIndex = position904, tokenIndex904
			return false
		},
		/* 101 UpperChar <- <UpperASCII> */
		func() bool {
			position909, tokenIndex909 := position, tokenIndex
			{
				position910 := position
				if !_rules[ruleUpperASCII]() {
					goto l909
				}
				add(ruleUpperChar, position910)
			}
			return true
		l909:
			position, tokenIndex = position909, tokenIndex909
			return false
		},
		/* 102 NameLowerChar <- <(LowerChar / LowerCharExtended / MiscodedChar)> */
		func() bool {
			position911, tokenIndex911 := position, tokenIndex
			{
				position912 := position
				{
					position913, tokenIndex913 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l914
					}
					goto l913
				l914:
					position, tokenIndex = position913, tokenIndex913
					if !_rules[ruleLowerCharExtended]() {
						goto l915
					}
					goto l913
				l915:
					position, tokenIndex = position913, tokenIndex913
					if !_rules[ruleMiscodedChar]() {
						goto l911
					}
				}
			l913:
				add(ruleNameLowerChar, position912)
			}
			return true
		l911:
			position, tokenIndex = position911, tokenIndex911
			return false
		},
		/* 103 MiscodedChar <- <'�'> */
		func() bool {
			position916, tokenIndex916 := position, tokenIndex
			{
				position917 := position
				if buffer[position] != rune('�') {
					goto l916
				}
				position++
				add(ruleMiscodedChar, position917)
			}
			return true
		l916:
			position, tokenIndex = position916, tokenIndex916
			return false
		},
		/* 104 LowerCharExtended <- <('æ' / 'œ' / 'à' / 'â' / 'å' / 'ã' / 'ä' / 'á' / 'ç' / 'č' / 'é' / 'è' / 'ë' / 'í' / 'ì' / 'ï' / 'ň' / 'ñ' / 'ñ' / 'ó' / 'ò' / 'ô' / 'ø' / 'õ' / 'ö' / 'ú' / 'ù' / 'ü' / 'ŕ' / 'ř' / 'ŗ' / 'ſ' / 'š' / 'š' / 'ş' / 'ž')> */
		func() bool {
			position918, tokenIndex918 := position, tokenIndex
			{
				position919 := position
				{
					position920, tokenIndex920 := position, tokenIndex
					if buffer[position] != rune('æ') {
						goto l921
					}
					position++
					goto l920
				l921:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('œ') {
						goto l922
					}
					position++
					goto l920
				l922:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('à') {
						goto l923
					}
					position++
					goto l920
				l923:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('â') {
						goto l924
					}
					position++
					goto l920
				l924:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('å') {
						goto l925
					}
					position++
					goto l920
				l925:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('ã') {
						goto l926
					}
					position++
					goto l920
				l926:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('ä') {
						goto l927
					}
					position++
					goto l920
				l927:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('á') {
						goto l928
					}
					position++
					goto l920
				l928:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('ç') {
						goto l929
					}
					position++
					goto l920
				l929:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('č') {
						goto l930
					}
					position++
					goto l920
				l930:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('é') {
						goto l931
					}
					position++
					goto l920
				l931:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('è') {
						goto l932
					}
					position++
					goto l920
				l932:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('ë') {
						goto l933
					}
					position++
					goto l920
				l933:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('í') {
						goto l934
					}
					position++
					goto l920
				l934:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('ì') {
						goto l935
					}
					position++
					goto l920
				l935:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('ï') {
						goto l936
					}
					position++
					goto l920
				l936:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('ň') {
						goto l937
					}
					position++
					goto l920
				l937:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('ñ') {
						goto l938
					}
					position++
					goto l920
				l938:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('ñ') {
						goto l939
					}
					position++
					goto l920
				l939:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('ó') {
						goto l940
					}
					position++
					goto l920
				l940:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('ò') {
						goto l941
					}
					position++
					goto l920
				l941:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('ô') {
						goto l942
					}
					position++
					goto l920
				l942:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('ø') {
						goto l943
					}
					position++
					goto l920
				l943:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('õ') {
						goto l944
					}
					position++
					goto l920
				l944:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('ö') {
						goto l945
					}
					position++
					goto l920
				l945:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('ú') {
						goto l946
					}
					position++
					goto l920
				l946:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('ù') {
						goto l947
					}
					position++
					goto l920
				l947:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('ü') {
						goto l948
					}
					position++
					goto l920
				l948:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('ŕ') {
						goto l949
					}
					position++
					goto l920
				l949:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('ř') {
						goto l950
					}
					position++
					goto l920
				l950:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('ŗ') {
						goto l951
					}
					position++
					goto l920
				l951:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('ſ') {
						goto l952
					}
					position++
					goto l920
				l952:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('š') {
						goto l953
					}
					position++
					goto l920
				l953:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('š') {
						goto l954
					}
					position++
					goto l920
				l954:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('ş') {
						goto l955
					}
					position++
					goto l920
				l955:
					position, tokenIndex = position920, tokenIndex920
					if buffer[position] != rune('ž') {
						goto l918
					}
					position++
				}
			l920:
				add(ruleLowerCharExtended, position919)
			}
			return true
		l918:
			position, tokenIndex = position918, tokenIndex918
			return false
		},
		/* 105 LowerChar <- <LowerASCII> */
		func() bool {
			position956, tokenIndex956 := position, tokenIndex
			{
				position957 := position
				if !_rules[ruleLowerASCII]() {
					goto l956
				}
				add(ruleLowerChar, position957)
			}
			return true
		l956:
			position, tokenIndex = position956, tokenIndex956
			return false
		},
		/* 106 SpaceCharEOI <- <(_ / !.)> */
		func() bool {
			position958, tokenIndex958 := position, tokenIndex
			{
				position959 := position
				{
					position960, tokenIndex960 := position, tokenIndex
					if !_rules[rule_]() {
						goto l961
					}
					goto l960
				l961:
					position, tokenIndex = position960, tokenIndex960
					{
						position962, tokenIndex962 := position, tokenIndex
						if !matchDot() {
							goto l962
						}
						goto l958
					l962:
						position, tokenIndex = position962, tokenIndex962
					}
				}
			l960:
				add(ruleSpaceCharEOI, position959)
			}
			return true
		l958:
			position, tokenIndex = position958, tokenIndex958
			return false
		},
		/* 107 Nums <- <[0-9]> */
		func() bool {
			position963, tokenIndex963 := position, tokenIndex
			{
				position964 := position
				if c := buffer[position]; c < rune('0') || c > rune('9') {
					goto l963
				}
				position++
				add(ruleNums, position964)
			}
			return true
		l963:
			position, tokenIndex = position963, tokenIndex963
			return false
		},
		/* 108 LowerASCII <- <[a-z]> */
		func() bool {
			position965, tokenIndex965 := position, tokenIndex
			{
				position966 := position
				if c := buffer[position]; c < rune('a') || c > rune('z') {
					goto l965
				}
				position++
				add(ruleLowerASCII, position966)
			}
			return true
		l965:
			position, tokenIndex = position965, tokenIndex965
			return false
		},
		/* 109 UpperASCII <- <[A-Z]> */
		func() bool {
			position967, tokenIndex967 := position, tokenIndex
			{
				position968 := position
				if c := buffer[position]; c < rune('A') || c > rune('Z') {
					goto l967
				}
				position++
				add(ruleUpperASCII, position968)
			}
			return true
		l967:
			position, tokenIndex = position967, tokenIndex967
			return false
		},
		/* 110 Apostrophe <- <(ApostrOther / ApostrASCII)> */
		func() bool {
			position969, tokenIndex969 := position, tokenIndex
			{
				position970 := position
				{
					position971, tokenIndex971 := position, tokenIndex
					if !_rules[ruleApostrOther]() {
						goto l972
					}
					goto l971
				l972:
					position, tokenIndex = position971, tokenIndex971
					if !_rules[ruleApostrASCII]() {
						goto l969
					}
				}
			l971:
				add(ruleApostrophe, position970)
			}
			return true
		l969:
			position, tokenIndex = position969, tokenIndex969
			return false
		},
		/* 111 ApostrASCII <- <'\''> */
		func() bool {
			position973, tokenIndex973 := position, tokenIndex
			{
				position974 := position
				if buffer[position] != rune('\'') {
					goto l973
				}
				position++
				add(ruleApostrASCII, position974)
			}
			return true
		l973:
			position, tokenIndex = position973, tokenIndex973
			return false
		},
		/* 112 ApostrOther <- <('‘' / '’')> */
		func() bool {
			position975, tokenIndex975 := position, tokenIndex
			{
				position976 := position
				{
					position977, tokenIndex977 := position, tokenIndex
					if buffer[position] != rune('‘') {
						goto l978
					}
					position++
					goto l977
				l978:
					position, tokenIndex = position977, tokenIndex977
					if buffer[position] != rune('’') {
						goto l975
					}
					position++
				}
			l977:
				add(ruleApostrOther, position976)
			}
			return true
		l975:
			position, tokenIndex = position975, tokenIndex975
			return false
		},
		/* 113 Dash <- <'-'> */
		func() bool {
			position979, tokenIndex979 := position, tokenIndex
			{
				position980 := position
				if buffer[position] != rune('-') {
					goto l979
				}
				position++
				add(ruleDash, position980)
			}
			return true
		l979:
			position, tokenIndex = position979, tokenIndex979
			return false
		},
		/* 114 _ <- <(MultipleSpace / SingleSpace)> */
		func() bool {
			position981, tokenIndex981 := position, tokenIndex
			{
				position982 := position
				{
					position983, tokenIndex983 := position, tokenIndex
					if !_rules[ruleMultipleSpace]() {
						goto l984
					}
					goto l983
				l984:
					position, tokenIndex = position983, tokenIndex983
					if !_rules[ruleSingleSpace]() {
						goto l981
					}
				}
			l983:
				add(rule_, position982)
			}
			return true
		l981:
			position, tokenIndex = position981, tokenIndex981
			return false
		},
		/* 115 MultipleSpace <- <(SingleSpace SingleSpace+)> */
		func() bool {
			position985, tokenIndex985 := position, tokenIndex
			{
				position986 := position
				if !_rules[ruleSingleSpace]() {
					goto l985
				}
				if !_rules[ruleSingleSpace]() {
					goto l985
				}
			l987:
				{
					position988, tokenIndex988 := position, tokenIndex
					if !_rules[ruleSingleSpace]() {
						goto l988
					}
					goto l987
				l988:
					position, tokenIndex = position988, tokenIndex988
				}
				add(ruleMultipleSpace, position986)
			}
			return true
		l985:
			position, tokenIndex = position985, tokenIndex985
			return false
		},
		/* 116 SingleSpace <- <(' ' / OtherSpace)> */
		func() bool {
			position989, tokenIndex989 := position, tokenIndex
			{
				position990 := position
				{
					position991, tokenIndex991 := position, tokenIndex
					if buffer[position] != rune(' ') {
						goto l992
					}
					position++
					goto l991
				l992:
					position, tokenIndex = position991, tokenIndex991
					if !_rules[ruleOtherSpace]() {
						goto l989
					}
				}
			l991:
				add(ruleSingleSpace, position990)
			}
			return true
		l989:
			position, tokenIndex = position989, tokenIndex989
			return false
		},
		/* 117 OtherSpace <- <('\u3000' / '\u00a0' / '\t' / '\r' / '\n' / '\f' / '\v')> */
		func() bool {
			position993, tokenIndex993 := position, tokenIndex
			{
				position994 := position
				{
					position995, tokenIndex995 := position, tokenIndex
					if buffer[position] != rune('\u3000') {
						goto l996
					}
					position++
					goto l995
				l996:
					position, tokenIndex = position995, tokenIndex995
					if buffer[position] != rune('\u00a0') {
						goto l997
					}
					position++
					goto l995
				l997:
					position, tokenIndex = position995, tokenIndex995
					if buffer[position] != rune('\t') {
						goto l998
					}
					position++
					goto l995
				l998:
					position, tokenIndex = position995, tokenIndex995
					if buffer[position] != rune('\r') {
						goto l999
					}
					position++
					goto l995
				l999:
					position, tokenIndex = position995, tokenIndex995
					if buffer[position] != rune('\n') {
						goto l1000
					}
					position++
					goto l995
				l1000:
					position, tokenIndex = position995, tokenIndex995
					if buffer[position] != rune('\f') {
						goto l1001
					}
					position++
					goto l995
				l1001:
					position, tokenIndex = position995, tokenIndex995
					if buffer[position] != rune('\v') {
						goto l993
					}
					position++
				}
			l995:
				add(ruleOtherSpace, position994)
			}
			return true
		l993:
			position, tokenIndex = position993, tokenIndex993
			return false
		},
		/* 119 Action0 <- <{ p.AddWarn(YearCharWarn) }> */
		func() bool {
			{
				add(ruleAction0, position)
			}
			return true
		},
	}
	p.rules = _rules
}
