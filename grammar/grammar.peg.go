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
	rules  [118]func() bool
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
		/* 29 UninomialCombo <- <(Uninomial _ RankUninomial _ Uninomial)> */
		func() bool {
			position230, tokenIndex230 := position, tokenIndex
			{
				position231 := position
				if !_rules[ruleUninomial]() {
					goto l230
				}
				if !_rules[rule_]() {
					goto l230
				}
				if !_rules[ruleRankUninomial]() {
					goto l230
				}
				if !_rules[rule_]() {
					goto l230
				}
				if !_rules[ruleUninomial]() {
					goto l230
				}
				add(ruleUninomialCombo, position231)
			}
			return true
		l230:
			position, tokenIndex = position230, tokenIndex230
			return false
		},
		/* 30 RankUninomial <- <(RankUninomialPlain / RankUninomialNotho)> */
		func() bool {
			position232, tokenIndex232 := position, tokenIndex
			{
				position233 := position
				{
					position234, tokenIndex234 := position, tokenIndex
					if !_rules[ruleRankUninomialPlain]() {
						goto l235
					}
					goto l234
				l235:
					position, tokenIndex = position234, tokenIndex234
					if !_rules[ruleRankUninomialNotho]() {
						goto l232
					}
				}
			l234:
				add(ruleRankUninomial, position233)
			}
			return true
		l232:
			position, tokenIndex = position232, tokenIndex232
			return false
		},
		/* 31 RankUninomialPlain <- <((('s' 'e' 'c' 't') / ('s' 'u' 'b' 's' 'e' 'c' 't') / ('t' 'r' 'i' 'b') / ('s' 'u' 'b' 't' 'r' 'i' 'b') / ('s' 'u' 'b' 's' 'e' 'r') / ('s' 'e' 'r') / ('s' 'u' 'b' 'g' 'e' 'n') / ('s' 'u' 'b' 'g') / ('f' 'a' 'm') / ('s' 'u' 'b' 'f' 'a' 'm') / ('s' 'u' 'p' 'e' 'r' 't' 'r' 'i' 'b')) ('.' / &SpaceCharEOI))> */
		func() bool {
			position236, tokenIndex236 := position, tokenIndex
			{
				position237 := position
				{
					position238, tokenIndex238 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l239
					}
					position++
					if buffer[position] != rune('e') {
						goto l239
					}
					position++
					if buffer[position] != rune('c') {
						goto l239
					}
					position++
					if buffer[position] != rune('t') {
						goto l239
					}
					position++
					goto l238
				l239:
					position, tokenIndex = position238, tokenIndex238
					if buffer[position] != rune('s') {
						goto l240
					}
					position++
					if buffer[position] != rune('u') {
						goto l240
					}
					position++
					if buffer[position] != rune('b') {
						goto l240
					}
					position++
					if buffer[position] != rune('s') {
						goto l240
					}
					position++
					if buffer[position] != rune('e') {
						goto l240
					}
					position++
					if buffer[position] != rune('c') {
						goto l240
					}
					position++
					if buffer[position] != rune('t') {
						goto l240
					}
					position++
					goto l238
				l240:
					position, tokenIndex = position238, tokenIndex238
					if buffer[position] != rune('t') {
						goto l241
					}
					position++
					if buffer[position] != rune('r') {
						goto l241
					}
					position++
					if buffer[position] != rune('i') {
						goto l241
					}
					position++
					if buffer[position] != rune('b') {
						goto l241
					}
					position++
					goto l238
				l241:
					position, tokenIndex = position238, tokenIndex238
					if buffer[position] != rune('s') {
						goto l242
					}
					position++
					if buffer[position] != rune('u') {
						goto l242
					}
					position++
					if buffer[position] != rune('b') {
						goto l242
					}
					position++
					if buffer[position] != rune('t') {
						goto l242
					}
					position++
					if buffer[position] != rune('r') {
						goto l242
					}
					position++
					if buffer[position] != rune('i') {
						goto l242
					}
					position++
					if buffer[position] != rune('b') {
						goto l242
					}
					position++
					goto l238
				l242:
					position, tokenIndex = position238, tokenIndex238
					if buffer[position] != rune('s') {
						goto l243
					}
					position++
					if buffer[position] != rune('u') {
						goto l243
					}
					position++
					if buffer[position] != rune('b') {
						goto l243
					}
					position++
					if buffer[position] != rune('s') {
						goto l243
					}
					position++
					if buffer[position] != rune('e') {
						goto l243
					}
					position++
					if buffer[position] != rune('r') {
						goto l243
					}
					position++
					goto l238
				l243:
					position, tokenIndex = position238, tokenIndex238
					if buffer[position] != rune('s') {
						goto l244
					}
					position++
					if buffer[position] != rune('e') {
						goto l244
					}
					position++
					if buffer[position] != rune('r') {
						goto l244
					}
					position++
					goto l238
				l244:
					position, tokenIndex = position238, tokenIndex238
					if buffer[position] != rune('s') {
						goto l245
					}
					position++
					if buffer[position] != rune('u') {
						goto l245
					}
					position++
					if buffer[position] != rune('b') {
						goto l245
					}
					position++
					if buffer[position] != rune('g') {
						goto l245
					}
					position++
					if buffer[position] != rune('e') {
						goto l245
					}
					position++
					if buffer[position] != rune('n') {
						goto l245
					}
					position++
					goto l238
				l245:
					position, tokenIndex = position238, tokenIndex238
					if buffer[position] != rune('s') {
						goto l246
					}
					position++
					if buffer[position] != rune('u') {
						goto l246
					}
					position++
					if buffer[position] != rune('b') {
						goto l246
					}
					position++
					if buffer[position] != rune('g') {
						goto l246
					}
					position++
					goto l238
				l246:
					position, tokenIndex = position238, tokenIndex238
					if buffer[position] != rune('f') {
						goto l247
					}
					position++
					if buffer[position] != rune('a') {
						goto l247
					}
					position++
					if buffer[position] != rune('m') {
						goto l247
					}
					position++
					goto l238
				l247:
					position, tokenIndex = position238, tokenIndex238
					if buffer[position] != rune('s') {
						goto l248
					}
					position++
					if buffer[position] != rune('u') {
						goto l248
					}
					position++
					if buffer[position] != rune('b') {
						goto l248
					}
					position++
					if buffer[position] != rune('f') {
						goto l248
					}
					position++
					if buffer[position] != rune('a') {
						goto l248
					}
					position++
					if buffer[position] != rune('m') {
						goto l248
					}
					position++
					goto l238
				l248:
					position, tokenIndex = position238, tokenIndex238
					if buffer[position] != rune('s') {
						goto l236
					}
					position++
					if buffer[position] != rune('u') {
						goto l236
					}
					position++
					if buffer[position] != rune('p') {
						goto l236
					}
					position++
					if buffer[position] != rune('e') {
						goto l236
					}
					position++
					if buffer[position] != rune('r') {
						goto l236
					}
					position++
					if buffer[position] != rune('t') {
						goto l236
					}
					position++
					if buffer[position] != rune('r') {
						goto l236
					}
					position++
					if buffer[position] != rune('i') {
						goto l236
					}
					position++
					if buffer[position] != rune('b') {
						goto l236
					}
					position++
				}
			l238:
				{
					position249, tokenIndex249 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l250
					}
					position++
					goto l249
				l250:
					position, tokenIndex = position249, tokenIndex249
					{
						position251, tokenIndex251 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l236
						}
						position, tokenIndex = position251, tokenIndex251
					}
				}
			l249:
				add(ruleRankUninomialPlain, position237)
			}
			return true
		l236:
			position, tokenIndex = position236, tokenIndex236
			return false
		},
		/* 32 RankUninomialNotho <- <('n' 'o' 't' 'h' 'o' _? (('s' 'e' 'c' 't') / ('g' 'e' 'n') / ('s' 'e' 'r') / ('s' 'u' 'b' 'g' 'e' 'e' 'n') / ('s' 'u' 'b' 'g' 'e' 'n') / ('s' 'u' 'b' 'g') / ('s' 'u' 'b' 's' 'e' 'c' 't') / ('s' 'u' 'b' 't' 'r' 'i' 'b')) ('.' / &SpaceCharEOI))> */
		func() bool {
			position252, tokenIndex252 := position, tokenIndex
			{
				position253 := position
				if buffer[position] != rune('n') {
					goto l252
				}
				position++
				if buffer[position] != rune('o') {
					goto l252
				}
				position++
				if buffer[position] != rune('t') {
					goto l252
				}
				position++
				if buffer[position] != rune('h') {
					goto l252
				}
				position++
				if buffer[position] != rune('o') {
					goto l252
				}
				position++
				{
					position254, tokenIndex254 := position, tokenIndex
					if !_rules[rule_]() {
						goto l254
					}
					goto l255
				l254:
					position, tokenIndex = position254, tokenIndex254
				}
			l255:
				{
					position256, tokenIndex256 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l257
					}
					position++
					if buffer[position] != rune('e') {
						goto l257
					}
					position++
					if buffer[position] != rune('c') {
						goto l257
					}
					position++
					if buffer[position] != rune('t') {
						goto l257
					}
					position++
					goto l256
				l257:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('g') {
						goto l258
					}
					position++
					if buffer[position] != rune('e') {
						goto l258
					}
					position++
					if buffer[position] != rune('n') {
						goto l258
					}
					position++
					goto l256
				l258:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('s') {
						goto l259
					}
					position++
					if buffer[position] != rune('e') {
						goto l259
					}
					position++
					if buffer[position] != rune('r') {
						goto l259
					}
					position++
					goto l256
				l259:
					position, tokenIndex = position256, tokenIndex256
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
					if buffer[position] != rune('g') {
						goto l260
					}
					position++
					if buffer[position] != rune('e') {
						goto l260
					}
					position++
					if buffer[position] != rune('e') {
						goto l260
					}
					position++
					if buffer[position] != rune('n') {
						goto l260
					}
					position++
					goto l256
				l260:
					position, tokenIndex = position256, tokenIndex256
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
					if buffer[position] != rune('g') {
						goto l261
					}
					position++
					if buffer[position] != rune('e') {
						goto l261
					}
					position++
					if buffer[position] != rune('n') {
						goto l261
					}
					position++
					goto l256
				l261:
					position, tokenIndex = position256, tokenIndex256
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
					if buffer[position] != rune('g') {
						goto l262
					}
					position++
					goto l256
				l262:
					position, tokenIndex = position256, tokenIndex256
					if buffer[position] != rune('s') {
						goto l263
					}
					position++
					if buffer[position] != rune('u') {
						goto l263
					}
					position++
					if buffer[position] != rune('b') {
						goto l263
					}
					position++
					if buffer[position] != rune('s') {
						goto l263
					}
					position++
					if buffer[position] != rune('e') {
						goto l263
					}
					position++
					if buffer[position] != rune('c') {
						goto l263
					}
					position++
					if buffer[position] != rune('t') {
						goto l263
					}
					position++
					goto l256
				l263:
					position, tokenIndex = position256, tokenIndex256
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
					if buffer[position] != rune('t') {
						goto l252
					}
					position++
					if buffer[position] != rune('r') {
						goto l252
					}
					position++
					if buffer[position] != rune('i') {
						goto l252
					}
					position++
					if buffer[position] != rune('b') {
						goto l252
					}
					position++
				}
			l256:
				{
					position264, tokenIndex264 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l265
					}
					position++
					goto l264
				l265:
					position, tokenIndex = position264, tokenIndex264
					{
						position266, tokenIndex266 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l252
						}
						position, tokenIndex = position266, tokenIndex266
					}
				}
			l264:
				add(ruleRankUninomialNotho, position253)
			}
			return true
		l252:
			position, tokenIndex = position252, tokenIndex252
			return false
		},
		/* 33 Uninomial <- <(UninomialWord (_ Authorship)?)> */
		func() bool {
			position267, tokenIndex267 := position, tokenIndex
			{
				position268 := position
				if !_rules[ruleUninomialWord]() {
					goto l267
				}
				{
					position269, tokenIndex269 := position, tokenIndex
					if !_rules[rule_]() {
						goto l269
					}
					if !_rules[ruleAuthorship]() {
						goto l269
					}
					goto l270
				l269:
					position, tokenIndex = position269, tokenIndex269
				}
			l270:
				add(ruleUninomial, position268)
			}
			return true
		l267:
			position, tokenIndex = position267, tokenIndex267
			return false
		},
		/* 34 UninomialWord <- <(CapWord / TwoLetterGenus)> */
		func() bool {
			position271, tokenIndex271 := position, tokenIndex
			{
				position272 := position
				{
					position273, tokenIndex273 := position, tokenIndex
					if !_rules[ruleCapWord]() {
						goto l274
					}
					goto l273
				l274:
					position, tokenIndex = position273, tokenIndex273
					if !_rules[ruleTwoLetterGenus]() {
						goto l271
					}
				}
			l273:
				add(ruleUninomialWord, position272)
			}
			return true
		l271:
			position, tokenIndex = position271, tokenIndex271
			return false
		},
		/* 35 AbbrGenus <- <(UpperChar LowerChar? '.')> */
		func() bool {
			position275, tokenIndex275 := position, tokenIndex
			{
				position276 := position
				if !_rules[ruleUpperChar]() {
					goto l275
				}
				{
					position277, tokenIndex277 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l277
					}
					goto l278
				l277:
					position, tokenIndex = position277, tokenIndex277
				}
			l278:
				if buffer[position] != rune('.') {
					goto l275
				}
				position++
				add(ruleAbbrGenus, position276)
			}
			return true
		l275:
			position, tokenIndex = position275, tokenIndex275
			return false
		},
		/* 36 CapWord <- <(CapWordWithDash / CapWord1)> */
		func() bool {
			position279, tokenIndex279 := position, tokenIndex
			{
				position280 := position
				{
					position281, tokenIndex281 := position, tokenIndex
					if !_rules[ruleCapWordWithDash]() {
						goto l282
					}
					goto l281
				l282:
					position, tokenIndex = position281, tokenIndex281
					if !_rules[ruleCapWord1]() {
						goto l279
					}
				}
			l281:
				add(ruleCapWord, position280)
			}
			return true
		l279:
			position, tokenIndex = position279, tokenIndex279
			return false
		},
		/* 37 CapWord1 <- <(NameUpperChar NameLowerChar NameLowerChar+ '?'?)> */
		func() bool {
			position283, tokenIndex283 := position, tokenIndex
			{
				position284 := position
				if !_rules[ruleNameUpperChar]() {
					goto l283
				}
				if !_rules[ruleNameLowerChar]() {
					goto l283
				}
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
				{
					position287, tokenIndex287 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l287
					}
					position++
					goto l288
				l287:
					position, tokenIndex = position287, tokenIndex287
				}
			l288:
				add(ruleCapWord1, position284)
			}
			return true
		l283:
			position, tokenIndex = position283, tokenIndex283
			return false
		},
		/* 38 CapWordWithDash <- <(CapWord1 Dash (UpperAfterDash / LowerAfterDash))> */
		func() bool {
			position289, tokenIndex289 := position, tokenIndex
			{
				position290 := position
				if !_rules[ruleCapWord1]() {
					goto l289
				}
				if !_rules[ruleDash]() {
					goto l289
				}
				{
					position291, tokenIndex291 := position, tokenIndex
					if !_rules[ruleUpperAfterDash]() {
						goto l292
					}
					goto l291
				l292:
					position, tokenIndex = position291, tokenIndex291
					if !_rules[ruleLowerAfterDash]() {
						goto l289
					}
				}
			l291:
				add(ruleCapWordWithDash, position290)
			}
			return true
		l289:
			position, tokenIndex = position289, tokenIndex289
			return false
		},
		/* 39 UpperAfterDash <- <CapWord1> */
		func() bool {
			position293, tokenIndex293 := position, tokenIndex
			{
				position294 := position
				if !_rules[ruleCapWord1]() {
					goto l293
				}
				add(ruleUpperAfterDash, position294)
			}
			return true
		l293:
			position, tokenIndex = position293, tokenIndex293
			return false
		},
		/* 40 LowerAfterDash <- <Word1> */
		func() bool {
			position295, tokenIndex295 := position, tokenIndex
			{
				position296 := position
				if !_rules[ruleWord1]() {
					goto l295
				}
				add(ruleLowerAfterDash, position296)
			}
			return true
		l295:
			position, tokenIndex = position295, tokenIndex295
			return false
		},
		/* 41 TwoLetterGenus <- <(('C' 'a') / ('E' 'a') / ('G' 'e') / ('I' 'a') / ('I' 'o') / ('I' 'x') / ('L' 'o') / ('O' 'a') / ('R' 'a') / ('T' 'y') / ('U' 'a') / ('A' 'a') / ('J' 'a') / ('Z' 'u') / ('L' 'a') / ('Q' 'u') / ('A' 's') / ('B' 'a'))> */
		func() bool {
			position297, tokenIndex297 := position, tokenIndex
			{
				position298 := position
				{
					position299, tokenIndex299 := position, tokenIndex
					if buffer[position] != rune('C') {
						goto l300
					}
					position++
					if buffer[position] != rune('a') {
						goto l300
					}
					position++
					goto l299
				l300:
					position, tokenIndex = position299, tokenIndex299
					if buffer[position] != rune('E') {
						goto l301
					}
					position++
					if buffer[position] != rune('a') {
						goto l301
					}
					position++
					goto l299
				l301:
					position, tokenIndex = position299, tokenIndex299
					if buffer[position] != rune('G') {
						goto l302
					}
					position++
					if buffer[position] != rune('e') {
						goto l302
					}
					position++
					goto l299
				l302:
					position, tokenIndex = position299, tokenIndex299
					if buffer[position] != rune('I') {
						goto l303
					}
					position++
					if buffer[position] != rune('a') {
						goto l303
					}
					position++
					goto l299
				l303:
					position, tokenIndex = position299, tokenIndex299
					if buffer[position] != rune('I') {
						goto l304
					}
					position++
					if buffer[position] != rune('o') {
						goto l304
					}
					position++
					goto l299
				l304:
					position, tokenIndex = position299, tokenIndex299
					if buffer[position] != rune('I') {
						goto l305
					}
					position++
					if buffer[position] != rune('x') {
						goto l305
					}
					position++
					goto l299
				l305:
					position, tokenIndex = position299, tokenIndex299
					if buffer[position] != rune('L') {
						goto l306
					}
					position++
					if buffer[position] != rune('o') {
						goto l306
					}
					position++
					goto l299
				l306:
					position, tokenIndex = position299, tokenIndex299
					if buffer[position] != rune('O') {
						goto l307
					}
					position++
					if buffer[position] != rune('a') {
						goto l307
					}
					position++
					goto l299
				l307:
					position, tokenIndex = position299, tokenIndex299
					if buffer[position] != rune('R') {
						goto l308
					}
					position++
					if buffer[position] != rune('a') {
						goto l308
					}
					position++
					goto l299
				l308:
					position, tokenIndex = position299, tokenIndex299
					if buffer[position] != rune('T') {
						goto l309
					}
					position++
					if buffer[position] != rune('y') {
						goto l309
					}
					position++
					goto l299
				l309:
					position, tokenIndex = position299, tokenIndex299
					if buffer[position] != rune('U') {
						goto l310
					}
					position++
					if buffer[position] != rune('a') {
						goto l310
					}
					position++
					goto l299
				l310:
					position, tokenIndex = position299, tokenIndex299
					if buffer[position] != rune('A') {
						goto l311
					}
					position++
					if buffer[position] != rune('a') {
						goto l311
					}
					position++
					goto l299
				l311:
					position, tokenIndex = position299, tokenIndex299
					if buffer[position] != rune('J') {
						goto l312
					}
					position++
					if buffer[position] != rune('a') {
						goto l312
					}
					position++
					goto l299
				l312:
					position, tokenIndex = position299, tokenIndex299
					if buffer[position] != rune('Z') {
						goto l313
					}
					position++
					if buffer[position] != rune('u') {
						goto l313
					}
					position++
					goto l299
				l313:
					position, tokenIndex = position299, tokenIndex299
					if buffer[position] != rune('L') {
						goto l314
					}
					position++
					if buffer[position] != rune('a') {
						goto l314
					}
					position++
					goto l299
				l314:
					position, tokenIndex = position299, tokenIndex299
					if buffer[position] != rune('Q') {
						goto l315
					}
					position++
					if buffer[position] != rune('u') {
						goto l315
					}
					position++
					goto l299
				l315:
					position, tokenIndex = position299, tokenIndex299
					if buffer[position] != rune('A') {
						goto l316
					}
					position++
					if buffer[position] != rune('s') {
						goto l316
					}
					position++
					goto l299
				l316:
					position, tokenIndex = position299, tokenIndex299
					if buffer[position] != rune('B') {
						goto l297
					}
					position++
					if buffer[position] != rune('a') {
						goto l297
					}
					position++
				}
			l299:
				add(ruleTwoLetterGenus, position298)
			}
			return true
		l297:
			position, tokenIndex = position297, tokenIndex297
			return false
		},
		/* 42 Word <- <(!((AuthorPrefix / RankUninomial / Approximation / Word4) SpaceCharEOI) (WordApostr / WordStartsWithDigit / MultiDashedWord / Word2 / Word1) &(SpaceCharEOI / '('))> */
		func() bool {
			position317, tokenIndex317 := position, tokenIndex
			{
				position318 := position
				{
					position319, tokenIndex319 := position, tokenIndex
					{
						position320, tokenIndex320 := position, tokenIndex
						if !_rules[ruleAuthorPrefix]() {
							goto l321
						}
						goto l320
					l321:
						position, tokenIndex = position320, tokenIndex320
						if !_rules[ruleRankUninomial]() {
							goto l322
						}
						goto l320
					l322:
						position, tokenIndex = position320, tokenIndex320
						if !_rules[ruleApproximation]() {
							goto l323
						}
						goto l320
					l323:
						position, tokenIndex = position320, tokenIndex320
						if !_rules[ruleWord4]() {
							goto l319
						}
					}
				l320:
					if !_rules[ruleSpaceCharEOI]() {
						goto l319
					}
					goto l317
				l319:
					position, tokenIndex = position319, tokenIndex319
				}
				{
					position324, tokenIndex324 := position, tokenIndex
					if !_rules[ruleWordApostr]() {
						goto l325
					}
					goto l324
				l325:
					position, tokenIndex = position324, tokenIndex324
					if !_rules[ruleWordStartsWithDigit]() {
						goto l326
					}
					goto l324
				l326:
					position, tokenIndex = position324, tokenIndex324
					if !_rules[ruleMultiDashedWord]() {
						goto l327
					}
					goto l324
				l327:
					position, tokenIndex = position324, tokenIndex324
					if !_rules[ruleWord2]() {
						goto l328
					}
					goto l324
				l328:
					position, tokenIndex = position324, tokenIndex324
					if !_rules[ruleWord1]() {
						goto l317
					}
				}
			l324:
				{
					position329, tokenIndex329 := position, tokenIndex
					{
						position330, tokenIndex330 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l331
						}
						goto l330
					l331:
						position, tokenIndex = position330, tokenIndex330
						if buffer[position] != rune('(') {
							goto l317
						}
						position++
					}
				l330:
					position, tokenIndex = position329, tokenIndex329
				}
				add(ruleWord, position318)
			}
			return true
		l317:
			position, tokenIndex = position317, tokenIndex317
			return false
		},
		/* 43 Word1 <- <((LowerASCII Dash)? NameLowerChar NameLowerChar+)> */
		func() bool {
			position332, tokenIndex332 := position, tokenIndex
			{
				position333 := position
				{
					position334, tokenIndex334 := position, tokenIndex
					if !_rules[ruleLowerASCII]() {
						goto l334
					}
					if !_rules[ruleDash]() {
						goto l334
					}
					goto l335
				l334:
					position, tokenIndex = position334, tokenIndex334
				}
			l335:
				if !_rules[ruleNameLowerChar]() {
					goto l332
				}
				if !_rules[ruleNameLowerChar]() {
					goto l332
				}
			l336:
				{
					position337, tokenIndex337 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l337
					}
					goto l336
				l337:
					position, tokenIndex = position337, tokenIndex337
				}
				add(ruleWord1, position333)
			}
			return true
		l332:
			position, tokenIndex = position332, tokenIndex332
			return false
		},
		/* 44 WordStartsWithDigit <- <(('1' / '2' / '3' / '4' / '5' / '6' / '7' / '8' / '9') Nums? ('.' / Dash)? NameLowerChar NameLowerChar NameLowerChar NameLowerChar+)> */
		func() bool {
			position338, tokenIndex338 := position, tokenIndex
			{
				position339 := position
				{
					position340, tokenIndex340 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l341
					}
					position++
					goto l340
				l341:
					position, tokenIndex = position340, tokenIndex340
					if buffer[position] != rune('2') {
						goto l342
					}
					position++
					goto l340
				l342:
					position, tokenIndex = position340, tokenIndex340
					if buffer[position] != rune('3') {
						goto l343
					}
					position++
					goto l340
				l343:
					position, tokenIndex = position340, tokenIndex340
					if buffer[position] != rune('4') {
						goto l344
					}
					position++
					goto l340
				l344:
					position, tokenIndex = position340, tokenIndex340
					if buffer[position] != rune('5') {
						goto l345
					}
					position++
					goto l340
				l345:
					position, tokenIndex = position340, tokenIndex340
					if buffer[position] != rune('6') {
						goto l346
					}
					position++
					goto l340
				l346:
					position, tokenIndex = position340, tokenIndex340
					if buffer[position] != rune('7') {
						goto l347
					}
					position++
					goto l340
				l347:
					position, tokenIndex = position340, tokenIndex340
					if buffer[position] != rune('8') {
						goto l348
					}
					position++
					goto l340
				l348:
					position, tokenIndex = position340, tokenIndex340
					if buffer[position] != rune('9') {
						goto l338
					}
					position++
				}
			l340:
				{
					position349, tokenIndex349 := position, tokenIndex
					if !_rules[ruleNums]() {
						goto l349
					}
					goto l350
				l349:
					position, tokenIndex = position349, tokenIndex349
				}
			l350:
				{
					position351, tokenIndex351 := position, tokenIndex
					{
						position353, tokenIndex353 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l354
						}
						position++
						goto l353
					l354:
						position, tokenIndex = position353, tokenIndex353
						if !_rules[ruleDash]() {
							goto l351
						}
					}
				l353:
					goto l352
				l351:
					position, tokenIndex = position351, tokenIndex351
				}
			l352:
				if !_rules[ruleNameLowerChar]() {
					goto l338
				}
				if !_rules[ruleNameLowerChar]() {
					goto l338
				}
				if !_rules[ruleNameLowerChar]() {
					goto l338
				}
				if !_rules[ruleNameLowerChar]() {
					goto l338
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
				add(ruleWordStartsWithDigit, position339)
			}
			return true
		l338:
			position, tokenIndex = position338, tokenIndex338
			return false
		},
		/* 45 Word2 <- <(NameLowerChar+ Dash? NameLowerChar+)> */
		func() bool {
			position357, tokenIndex357 := position, tokenIndex
			{
				position358 := position
				if !_rules[ruleNameLowerChar]() {
					goto l357
				}
			l359:
				{
					position360, tokenIndex360 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l360
					}
					goto l359
				l360:
					position, tokenIndex = position360, tokenIndex360
				}
				{
					position361, tokenIndex361 := position, tokenIndex
					if !_rules[ruleDash]() {
						goto l361
					}
					goto l362
				l361:
					position, tokenIndex = position361, tokenIndex361
				}
			l362:
				if !_rules[ruleNameLowerChar]() {
					goto l357
				}
			l363:
				{
					position364, tokenIndex364 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l364
					}
					goto l363
				l364:
					position, tokenIndex = position364, tokenIndex364
				}
				add(ruleWord2, position358)
			}
			return true
		l357:
			position, tokenIndex = position357, tokenIndex357
			return false
		},
		/* 46 WordApostr <- <(NameLowerChar NameLowerChar* Apostrophe Word1)> */
		func() bool {
			position365, tokenIndex365 := position, tokenIndex
			{
				position366 := position
				if !_rules[ruleNameLowerChar]() {
					goto l365
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
				if !_rules[ruleApostrophe]() {
					goto l365
				}
				if !_rules[ruleWord1]() {
					goto l365
				}
				add(ruleWordApostr, position366)
			}
			return true
		l365:
			position, tokenIndex = position365, tokenIndex365
			return false
		},
		/* 47 Word4 <- <(NameLowerChar+ '.' NameLowerChar)> */
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
				if buffer[position] != rune('.') {
					goto l369
				}
				position++
				if !_rules[ruleNameLowerChar]() {
					goto l369
				}
				add(ruleWord4, position370)
			}
			return true
		l369:
			position, tokenIndex = position369, tokenIndex369
			return false
		},
		/* 48 MultiDashedWord <- <(NameLowerChar+ Dash NameLowerChar+ Dash NameLowerChar+ (Dash NameLowerChar+)?)> */
		func() bool {
			position373, tokenIndex373 := position, tokenIndex
			{
				position374 := position
				if !_rules[ruleNameLowerChar]() {
					goto l373
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
				if !_rules[ruleDash]() {
					goto l373
				}
				if !_rules[ruleNameLowerChar]() {
					goto l373
				}
			l377:
				{
					position378, tokenIndex378 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l378
					}
					goto l377
				l378:
					position, tokenIndex = position378, tokenIndex378
				}
				if !_rules[ruleDash]() {
					goto l373
				}
				if !_rules[ruleNameLowerChar]() {
					goto l373
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
				{
					position381, tokenIndex381 := position, tokenIndex
					if !_rules[ruleDash]() {
						goto l381
					}
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
					goto l382
				l381:
					position, tokenIndex = position381, tokenIndex381
				}
			l382:
				add(ruleMultiDashedWord, position374)
			}
			return true
		l373:
			position, tokenIndex = position373, tokenIndex373
			return false
		},
		/* 49 HybridChar <- <'×'> */
		func() bool {
			position385, tokenIndex385 := position, tokenIndex
			{
				position386 := position
				if buffer[position] != rune('×') {
					goto l385
				}
				position++
				add(ruleHybridChar, position386)
			}
			return true
		l385:
			position, tokenIndex = position385, tokenIndex385
			return false
		},
		/* 50 ApproxNameIgnored <- <.*> */
		func() bool {
			{
				position388 := position
			l389:
				{
					position390, tokenIndex390 := position, tokenIndex
					if !matchDot() {
						goto l390
					}
					goto l389
				l390:
					position, tokenIndex = position390, tokenIndex390
				}
				add(ruleApproxNameIgnored, position388)
			}
			return true
		},
		/* 51 Approximation <- <(('s' 'p' '.' _? ('n' 'r' '.')) / ('s' 'p' '.' _? ('a' 'f' 'f' '.')) / ('m' 'o' 'n' 's' 't' '.') / '?' / ((('s' 'p' 'p') / ('n' 'r') / ('s' 'p') / ('a' 'f' 'f') / ('s' 'p' 'e' 'c' 'i' 'e' 's')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position391, tokenIndex391 := position, tokenIndex
			{
				position392 := position
				{
					position393, tokenIndex393 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l394
					}
					position++
					if buffer[position] != rune('p') {
						goto l394
					}
					position++
					if buffer[position] != rune('.') {
						goto l394
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
					if buffer[position] != rune('n') {
						goto l394
					}
					position++
					if buffer[position] != rune('r') {
						goto l394
					}
					position++
					if buffer[position] != rune('.') {
						goto l394
					}
					position++
					goto l393
				l394:
					position, tokenIndex = position393, tokenIndex393
					if buffer[position] != rune('s') {
						goto l397
					}
					position++
					if buffer[position] != rune('p') {
						goto l397
					}
					position++
					if buffer[position] != rune('.') {
						goto l397
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
					if buffer[position] != rune('a') {
						goto l397
					}
					position++
					if buffer[position] != rune('f') {
						goto l397
					}
					position++
					if buffer[position] != rune('f') {
						goto l397
					}
					position++
					if buffer[position] != rune('.') {
						goto l397
					}
					position++
					goto l393
				l397:
					position, tokenIndex = position393, tokenIndex393
					if buffer[position] != rune('m') {
						goto l400
					}
					position++
					if buffer[position] != rune('o') {
						goto l400
					}
					position++
					if buffer[position] != rune('n') {
						goto l400
					}
					position++
					if buffer[position] != rune('s') {
						goto l400
					}
					position++
					if buffer[position] != rune('t') {
						goto l400
					}
					position++
					if buffer[position] != rune('.') {
						goto l400
					}
					position++
					goto l393
				l400:
					position, tokenIndex = position393, tokenIndex393
					if buffer[position] != rune('?') {
						goto l401
					}
					position++
					goto l393
				l401:
					position, tokenIndex = position393, tokenIndex393
					{
						position402, tokenIndex402 := position, tokenIndex
						if buffer[position] != rune('s') {
							goto l403
						}
						position++
						if buffer[position] != rune('p') {
							goto l403
						}
						position++
						if buffer[position] != rune('p') {
							goto l403
						}
						position++
						goto l402
					l403:
						position, tokenIndex = position402, tokenIndex402
						if buffer[position] != rune('n') {
							goto l404
						}
						position++
						if buffer[position] != rune('r') {
							goto l404
						}
						position++
						goto l402
					l404:
						position, tokenIndex = position402, tokenIndex402
						if buffer[position] != rune('s') {
							goto l405
						}
						position++
						if buffer[position] != rune('p') {
							goto l405
						}
						position++
						goto l402
					l405:
						position, tokenIndex = position402, tokenIndex402
						if buffer[position] != rune('a') {
							goto l406
						}
						position++
						if buffer[position] != rune('f') {
							goto l406
						}
						position++
						if buffer[position] != rune('f') {
							goto l406
						}
						position++
						goto l402
					l406:
						position, tokenIndex = position402, tokenIndex402
						if buffer[position] != rune('s') {
							goto l391
						}
						position++
						if buffer[position] != rune('p') {
							goto l391
						}
						position++
						if buffer[position] != rune('e') {
							goto l391
						}
						position++
						if buffer[position] != rune('c') {
							goto l391
						}
						position++
						if buffer[position] != rune('i') {
							goto l391
						}
						position++
						if buffer[position] != rune('e') {
							goto l391
						}
						position++
						if buffer[position] != rune('s') {
							goto l391
						}
						position++
					}
				l402:
					{
						position407, tokenIndex407 := position, tokenIndex
						{
							position409, tokenIndex409 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l408
							}
							position, tokenIndex = position409, tokenIndex409
						}
						goto l407
					l408:
						position, tokenIndex = position407, tokenIndex407
						if buffer[position] != rune('.') {
							goto l391
						}
						position++
					}
				l407:
				}
			l393:
				add(ruleApproximation, position392)
			}
			return true
		l391:
			position, tokenIndex = position391, tokenIndex391
			return false
		},
		/* 52 Authorship <- <((AuthorshipCombo / OriginalAuthorship) &(SpaceCharEOI / ';' / ','))> */
		func() bool {
			position410, tokenIndex410 := position, tokenIndex
			{
				position411 := position
				{
					position412, tokenIndex412 := position, tokenIndex
					if !_rules[ruleAuthorshipCombo]() {
						goto l413
					}
					goto l412
				l413:
					position, tokenIndex = position412, tokenIndex412
					if !_rules[ruleOriginalAuthorship]() {
						goto l410
					}
				}
			l412:
				{
					position414, tokenIndex414 := position, tokenIndex
					{
						position415, tokenIndex415 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l416
						}
						goto l415
					l416:
						position, tokenIndex = position415, tokenIndex415
						if buffer[position] != rune(';') {
							goto l417
						}
						position++
						goto l415
					l417:
						position, tokenIndex = position415, tokenIndex415
						if buffer[position] != rune(',') {
							goto l410
						}
						position++
					}
				l415:
					position, tokenIndex = position414, tokenIndex414
				}
				add(ruleAuthorship, position411)
			}
			return true
		l410:
			position, tokenIndex = position410, tokenIndex410
			return false
		},
		/* 53 AuthorshipCombo <- <(OriginalAuthorshipComb (_? CombinationAuthorship)?)> */
		func() bool {
			position418, tokenIndex418 := position, tokenIndex
			{
				position419 := position
				if !_rules[ruleOriginalAuthorshipComb]() {
					goto l418
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
					if !_rules[ruleCombinationAuthorship]() {
						goto l420
					}
					goto l421
				l420:
					position, tokenIndex = position420, tokenIndex420
				}
			l421:
				add(ruleAuthorshipCombo, position419)
			}
			return true
		l418:
			position, tokenIndex = position418, tokenIndex418
			return false
		},
		/* 54 OriginalAuthorship <- <AuthorsGroup> */
		func() bool {
			position424, tokenIndex424 := position, tokenIndex
			{
				position425 := position
				if !_rules[ruleAuthorsGroup]() {
					goto l424
				}
				add(ruleOriginalAuthorship, position425)
			}
			return true
		l424:
			position, tokenIndex = position424, tokenIndex424
			return false
		},
		/* 55 OriginalAuthorshipComb <- <(BasionymAuthorshipYearMisformed / BasionymAuthorship / BasionymAuthorshipMissingParens)> */
		func() bool {
			position426, tokenIndex426 := position, tokenIndex
			{
				position427 := position
				{
					position428, tokenIndex428 := position, tokenIndex
					if !_rules[ruleBasionymAuthorshipYearMisformed]() {
						goto l429
					}
					goto l428
				l429:
					position, tokenIndex = position428, tokenIndex428
					if !_rules[ruleBasionymAuthorship]() {
						goto l430
					}
					goto l428
				l430:
					position, tokenIndex = position428, tokenIndex428
					if !_rules[ruleBasionymAuthorshipMissingParens]() {
						goto l426
					}
				}
			l428:
				add(ruleOriginalAuthorshipComb, position427)
			}
			return true
		l426:
			position, tokenIndex = position426, tokenIndex426
			return false
		},
		/* 56 CombinationAuthorship <- <AuthorsGroup> */
		func() bool {
			position431, tokenIndex431 := position, tokenIndex
			{
				position432 := position
				if !_rules[ruleAuthorsGroup]() {
					goto l431
				}
				add(ruleCombinationAuthorship, position432)
			}
			return true
		l431:
			position, tokenIndex = position431, tokenIndex431
			return false
		},
		/* 57 BasionymAuthorshipMissingParens <- <(MissingParensStart / MissingParensEnd)> */
		func() bool {
			position433, tokenIndex433 := position, tokenIndex
			{
				position434 := position
				{
					position435, tokenIndex435 := position, tokenIndex
					if !_rules[ruleMissingParensStart]() {
						goto l436
					}
					goto l435
				l436:
					position, tokenIndex = position435, tokenIndex435
					if !_rules[ruleMissingParensEnd]() {
						goto l433
					}
				}
			l435:
				add(ruleBasionymAuthorshipMissingParens, position434)
			}
			return true
		l433:
			position, tokenIndex = position433, tokenIndex433
			return false
		},
		/* 58 MissingParensStart <- <('(' _? AuthorsGroup)> */
		func() bool {
			position437, tokenIndex437 := position, tokenIndex
			{
				position438 := position
				if buffer[position] != rune('(') {
					goto l437
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
				if !_rules[ruleAuthorsGroup]() {
					goto l437
				}
				add(ruleMissingParensStart, position438)
			}
			return true
		l437:
			position, tokenIndex = position437, tokenIndex437
			return false
		},
		/* 59 MissingParensEnd <- <(AuthorsGroup _? ')')> */
		func() bool {
			position441, tokenIndex441 := position, tokenIndex
			{
				position442 := position
				if !_rules[ruleAuthorsGroup]() {
					goto l441
				}
				{
					position443, tokenIndex443 := position, tokenIndex
					if !_rules[rule_]() {
						goto l443
					}
					goto l444
				l443:
					position, tokenIndex = position443, tokenIndex443
				}
			l444:
				if buffer[position] != rune(')') {
					goto l441
				}
				position++
				add(ruleMissingParensEnd, position442)
			}
			return true
		l441:
			position, tokenIndex = position441, tokenIndex441
			return false
		},
		/* 60 BasionymAuthorshipYearMisformed <- <('(' _? AuthorsGroup _? ')' (_? ',')? _? Year)> */
		func() bool {
			position445, tokenIndex445 := position, tokenIndex
			{
				position446 := position
				if buffer[position] != rune('(') {
					goto l445
				}
				position++
				{
					position447, tokenIndex447 := position, tokenIndex
					if !_rules[rule_]() {
						goto l447
					}
					goto l448
				l447:
					position, tokenIndex = position447, tokenIndex447
				}
			l448:
				if !_rules[ruleAuthorsGroup]() {
					goto l445
				}
				{
					position449, tokenIndex449 := position, tokenIndex
					if !_rules[rule_]() {
						goto l449
					}
					goto l450
				l449:
					position, tokenIndex = position449, tokenIndex449
				}
			l450:
				if buffer[position] != rune(')') {
					goto l445
				}
				position++
				{
					position451, tokenIndex451 := position, tokenIndex
					{
						position453, tokenIndex453 := position, tokenIndex
						if !_rules[rule_]() {
							goto l453
						}
						goto l454
					l453:
						position, tokenIndex = position453, tokenIndex453
					}
				l454:
					if buffer[position] != rune(',') {
						goto l451
					}
					position++
					goto l452
				l451:
					position, tokenIndex = position451, tokenIndex451
				}
			l452:
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
				if !_rules[ruleYear]() {
					goto l445
				}
				add(ruleBasionymAuthorshipYearMisformed, position446)
			}
			return true
		l445:
			position, tokenIndex = position445, tokenIndex445
			return false
		},
		/* 61 BasionymAuthorship <- <(BasionymAuthorship1 / BasionymAuthorship2Parens)> */
		func() bool {
			position457, tokenIndex457 := position, tokenIndex
			{
				position458 := position
				{
					position459, tokenIndex459 := position, tokenIndex
					if !_rules[ruleBasionymAuthorship1]() {
						goto l460
					}
					goto l459
				l460:
					position, tokenIndex = position459, tokenIndex459
					if !_rules[ruleBasionymAuthorship2Parens]() {
						goto l457
					}
				}
			l459:
				add(ruleBasionymAuthorship, position458)
			}
			return true
		l457:
			position, tokenIndex = position457, tokenIndex457
			return false
		},
		/* 62 BasionymAuthorship1 <- <('(' _? AuthorsGroup _? ')')> */
		func() bool {
			position461, tokenIndex461 := position, tokenIndex
			{
				position462 := position
				if buffer[position] != rune('(') {
					goto l461
				}
				position++
				{
					position463, tokenIndex463 := position, tokenIndex
					if !_rules[rule_]() {
						goto l463
					}
					goto l464
				l463:
					position, tokenIndex = position463, tokenIndex463
				}
			l464:
				if !_rules[ruleAuthorsGroup]() {
					goto l461
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
				if buffer[position] != rune(')') {
					goto l461
				}
				position++
				add(ruleBasionymAuthorship1, position462)
			}
			return true
		l461:
			position, tokenIndex = position461, tokenIndex461
			return false
		},
		/* 63 BasionymAuthorship2Parens <- <('(' _? '(' _? AuthorsGroup _? ')' _? ')')> */
		func() bool {
			position467, tokenIndex467 := position, tokenIndex
			{
				position468 := position
				if buffer[position] != rune('(') {
					goto l467
				}
				position++
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
				if buffer[position] != rune('(') {
					goto l467
				}
				position++
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
				if !_rules[ruleAuthorsGroup]() {
					goto l467
				}
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
				if buffer[position] != rune(')') {
					goto l467
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
				if buffer[position] != rune(')') {
					goto l467
				}
				position++
				add(ruleBasionymAuthorship2Parens, position468)
			}
			return true
		l467:
			position, tokenIndex = position467, tokenIndex467
			return false
		},
		/* 64 AuthorsGroup <- <(AuthorsTeam (_ (AuthorEmend / AuthorEx) AuthorsTeam)?)> */
		func() bool {
			position477, tokenIndex477 := position, tokenIndex
			{
				position478 := position
				if !_rules[ruleAuthorsTeam]() {
					goto l477
				}
				{
					position479, tokenIndex479 := position, tokenIndex
					if !_rules[rule_]() {
						goto l479
					}
					{
						position481, tokenIndex481 := position, tokenIndex
						if !_rules[ruleAuthorEmend]() {
							goto l482
						}
						goto l481
					l482:
						position, tokenIndex = position481, tokenIndex481
						if !_rules[ruleAuthorEx]() {
							goto l479
						}
					}
				l481:
					if !_rules[ruleAuthorsTeam]() {
						goto l479
					}
					goto l480
				l479:
					position, tokenIndex = position479, tokenIndex479
				}
			l480:
				add(ruleAuthorsGroup, position478)
			}
			return true
		l477:
			position, tokenIndex = position477, tokenIndex477
			return false
		},
		/* 65 AuthorsTeam <- <(Author (AuthorSep Author)* (_? ','? _? Year)?)> */
		func() bool {
			position483, tokenIndex483 := position, tokenIndex
			{
				position484 := position
				if !_rules[ruleAuthor]() {
					goto l483
				}
			l485:
				{
					position486, tokenIndex486 := position, tokenIndex
					if !_rules[ruleAuthorSep]() {
						goto l486
					}
					if !_rules[ruleAuthor]() {
						goto l486
					}
					goto l485
				l486:
					position, tokenIndex = position486, tokenIndex486
				}
				{
					position487, tokenIndex487 := position, tokenIndex
					{
						position489, tokenIndex489 := position, tokenIndex
						if !_rules[rule_]() {
							goto l489
						}
						goto l490
					l489:
						position, tokenIndex = position489, tokenIndex489
					}
				l490:
					{
						position491, tokenIndex491 := position, tokenIndex
						if buffer[position] != rune(',') {
							goto l491
						}
						position++
						goto l492
					l491:
						position, tokenIndex = position491, tokenIndex491
					}
				l492:
					{
						position493, tokenIndex493 := position, tokenIndex
						if !_rules[rule_]() {
							goto l493
						}
						goto l494
					l493:
						position, tokenIndex = position493, tokenIndex493
					}
				l494:
					if !_rules[ruleYear]() {
						goto l487
					}
					goto l488
				l487:
					position, tokenIndex = position487, tokenIndex487
				}
			l488:
				add(ruleAuthorsTeam, position484)
			}
			return true
		l483:
			position, tokenIndex = position483, tokenIndex483
			return false
		},
		/* 66 AuthorSep <- <(AuthorSep1 / AuthorSep2)> */
		func() bool {
			position495, tokenIndex495 := position, tokenIndex
			{
				position496 := position
				{
					position497, tokenIndex497 := position, tokenIndex
					if !_rules[ruleAuthorSep1]() {
						goto l498
					}
					goto l497
				l498:
					position, tokenIndex = position497, tokenIndex497
					if !_rules[ruleAuthorSep2]() {
						goto l495
					}
				}
			l497:
				add(ruleAuthorSep, position496)
			}
			return true
		l495:
			position, tokenIndex = position495, tokenIndex495
			return false
		},
		/* 67 AuthorSep1 <- <(_? (',' _)? ('&' / ('e' 't') / ('a' 'n' 'd') / ('a' 'p' 'u' 'd')) _?)> */
		func() bool {
			position499, tokenIndex499 := position, tokenIndex
			{
				position500 := position
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
					if !_rules[rule_]() {
						goto l503
					}
					goto l504
				l503:
					position, tokenIndex = position503, tokenIndex503
				}
			l504:
				{
					position505, tokenIndex505 := position, tokenIndex
					if buffer[position] != rune('&') {
						goto l506
					}
					position++
					goto l505
				l506:
					position, tokenIndex = position505, tokenIndex505
					if buffer[position] != rune('e') {
						goto l507
					}
					position++
					if buffer[position] != rune('t') {
						goto l507
					}
					position++
					goto l505
				l507:
					position, tokenIndex = position505, tokenIndex505
					if buffer[position] != rune('a') {
						goto l508
					}
					position++
					if buffer[position] != rune('n') {
						goto l508
					}
					position++
					if buffer[position] != rune('d') {
						goto l508
					}
					position++
					goto l505
				l508:
					position, tokenIndex = position505, tokenIndex505
					if buffer[position] != rune('a') {
						goto l499
					}
					position++
					if buffer[position] != rune('p') {
						goto l499
					}
					position++
					if buffer[position] != rune('u') {
						goto l499
					}
					position++
					if buffer[position] != rune('d') {
						goto l499
					}
					position++
				}
			l505:
				{
					position509, tokenIndex509 := position, tokenIndex
					if !_rules[rule_]() {
						goto l509
					}
					goto l510
				l509:
					position, tokenIndex = position509, tokenIndex509
				}
			l510:
				add(ruleAuthorSep1, position500)
			}
			return true
		l499:
			position, tokenIndex = position499, tokenIndex499
			return false
		},
		/* 68 AuthorSep2 <- <(_? ',' _?)> */
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
				if buffer[position] != rune(',') {
					goto l511
				}
				position++
				{
					position515, tokenIndex515 := position, tokenIndex
					if !_rules[rule_]() {
						goto l515
					}
					goto l516
				l515:
					position, tokenIndex = position515, tokenIndex515
				}
			l516:
				add(ruleAuthorSep2, position512)
			}
			return true
		l511:
			position, tokenIndex = position511, tokenIndex511
			return false
		},
		/* 69 AuthorEx <- <((('e' 'x' '.'?) / ('i' 'n')) _)> */
		func() bool {
			position517, tokenIndex517 := position, tokenIndex
			{
				position518 := position
				{
					position519, tokenIndex519 := position, tokenIndex
					if buffer[position] != rune('e') {
						goto l520
					}
					position++
					if buffer[position] != rune('x') {
						goto l520
					}
					position++
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
					goto l519
				l520:
					position, tokenIndex = position519, tokenIndex519
					if buffer[position] != rune('i') {
						goto l517
					}
					position++
					if buffer[position] != rune('n') {
						goto l517
					}
					position++
				}
			l519:
				if !_rules[rule_]() {
					goto l517
				}
				add(ruleAuthorEx, position518)
			}
			return true
		l517:
			position, tokenIndex = position517, tokenIndex517
			return false
		},
		/* 70 AuthorEmend <- <('e' 'm' 'e' 'n' 'd' '.'? _)> */
		func() bool {
			position523, tokenIndex523 := position, tokenIndex
			{
				position524 := position
				if buffer[position] != rune('e') {
					goto l523
				}
				position++
				if buffer[position] != rune('m') {
					goto l523
				}
				position++
				if buffer[position] != rune('e') {
					goto l523
				}
				position++
				if buffer[position] != rune('n') {
					goto l523
				}
				position++
				if buffer[position] != rune('d') {
					goto l523
				}
				position++
				{
					position525, tokenIndex525 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l525
					}
					position++
					goto l526
				l525:
					position, tokenIndex = position525, tokenIndex525
				}
			l526:
				if !_rules[rule_]() {
					goto l523
				}
				add(ruleAuthorEmend, position524)
			}
			return true
		l523:
			position, tokenIndex = position523, tokenIndex523
			return false
		},
		/* 71 Author <- <(Author1 / Author2 / UnknownAuthor)> */
		func() bool {
			position527, tokenIndex527 := position, tokenIndex
			{
				position528 := position
				{
					position529, tokenIndex529 := position, tokenIndex
					if !_rules[ruleAuthor1]() {
						goto l530
					}
					goto l529
				l530:
					position, tokenIndex = position529, tokenIndex529
					if !_rules[ruleAuthor2]() {
						goto l531
					}
					goto l529
				l531:
					position, tokenIndex = position529, tokenIndex529
					if !_rules[ruleUnknownAuthor]() {
						goto l527
					}
				}
			l529:
				add(ruleAuthor, position528)
			}
			return true
		l527:
			position, tokenIndex = position527, tokenIndex527
			return false
		},
		/* 72 Author1 <- <(Author2 _? Filius)> */
		func() bool {
			position532, tokenIndex532 := position, tokenIndex
			{
				position533 := position
				if !_rules[ruleAuthor2]() {
					goto l532
				}
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
				if !_rules[ruleFilius]() {
					goto l532
				}
				add(ruleAuthor1, position533)
			}
			return true
		l532:
			position, tokenIndex = position532, tokenIndex532
			return false
		},
		/* 73 Author2 <- <(AuthorWord (_? AuthorWord)*)> */
		func() bool {
			position536, tokenIndex536 := position, tokenIndex
			{
				position537 := position
				if !_rules[ruleAuthorWord]() {
					goto l536
				}
			l538:
				{
					position539, tokenIndex539 := position, tokenIndex
					{
						position540, tokenIndex540 := position, tokenIndex
						if !_rules[rule_]() {
							goto l540
						}
						goto l541
					l540:
						position, tokenIndex = position540, tokenIndex540
					}
				l541:
					if !_rules[ruleAuthorWord]() {
						goto l539
					}
					goto l538
				l539:
					position, tokenIndex = position539, tokenIndex539
				}
				add(ruleAuthor2, position537)
			}
			return true
		l536:
			position, tokenIndex = position536, tokenIndex536
			return false
		},
		/* 74 UnknownAuthor <- <('?' / ((('a' 'u' 'c' 't') / ('a' 'n' 'o' 'n')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position542, tokenIndex542 := position, tokenIndex
			{
				position543 := position
				{
					position544, tokenIndex544 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l545
					}
					position++
					goto l544
				l545:
					position, tokenIndex = position544, tokenIndex544
					{
						position546, tokenIndex546 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l547
						}
						position++
						if buffer[position] != rune('u') {
							goto l547
						}
						position++
						if buffer[position] != rune('c') {
							goto l547
						}
						position++
						if buffer[position] != rune('t') {
							goto l547
						}
						position++
						goto l546
					l547:
						position, tokenIndex = position546, tokenIndex546
						if buffer[position] != rune('a') {
							goto l542
						}
						position++
						if buffer[position] != rune('n') {
							goto l542
						}
						position++
						if buffer[position] != rune('o') {
							goto l542
						}
						position++
						if buffer[position] != rune('n') {
							goto l542
						}
						position++
					}
				l546:
					{
						position548, tokenIndex548 := position, tokenIndex
						{
							position550, tokenIndex550 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l549
							}
							position, tokenIndex = position550, tokenIndex550
						}
						goto l548
					l549:
						position, tokenIndex = position548, tokenIndex548
						if buffer[position] != rune('.') {
							goto l542
						}
						position++
					}
				l548:
				}
			l544:
				add(ruleUnknownAuthor, position543)
			}
			return true
		l542:
			position, tokenIndex = position542, tokenIndex542
			return false
		},
		/* 75 AuthorWord <- <(!(('b' / 'B') ('o' / 'O') ('l' / 'L') ('d' / 'D') ':') (AuthorEtAl / AuthorWord2 / AuthorWord3 / AuthorPrefix))> */
		func() bool {
			position551, tokenIndex551 := position, tokenIndex
			{
				position552 := position
				{
					position553, tokenIndex553 := position, tokenIndex
					{
						position554, tokenIndex554 := position, tokenIndex
						if buffer[position] != rune('b') {
							goto l555
						}
						position++
						goto l554
					l555:
						position, tokenIndex = position554, tokenIndex554
						if buffer[position] != rune('B') {
							goto l553
						}
						position++
					}
				l554:
					{
						position556, tokenIndex556 := position, tokenIndex
						if buffer[position] != rune('o') {
							goto l557
						}
						position++
						goto l556
					l557:
						position, tokenIndex = position556, tokenIndex556
						if buffer[position] != rune('O') {
							goto l553
						}
						position++
					}
				l556:
					{
						position558, tokenIndex558 := position, tokenIndex
						if buffer[position] != rune('l') {
							goto l559
						}
						position++
						goto l558
					l559:
						position, tokenIndex = position558, tokenIndex558
						if buffer[position] != rune('L') {
							goto l553
						}
						position++
					}
				l558:
					{
						position560, tokenIndex560 := position, tokenIndex
						if buffer[position] != rune('d') {
							goto l561
						}
						position++
						goto l560
					l561:
						position, tokenIndex = position560, tokenIndex560
						if buffer[position] != rune('D') {
							goto l553
						}
						position++
					}
				l560:
					if buffer[position] != rune(':') {
						goto l553
					}
					position++
					goto l551
				l553:
					position, tokenIndex = position553, tokenIndex553
				}
				{
					position562, tokenIndex562 := position, tokenIndex
					if !_rules[ruleAuthorEtAl]() {
						goto l563
					}
					goto l562
				l563:
					position, tokenIndex = position562, tokenIndex562
					if !_rules[ruleAuthorWord2]() {
						goto l564
					}
					goto l562
				l564:
					position, tokenIndex = position562, tokenIndex562
					if !_rules[ruleAuthorWord3]() {
						goto l565
					}
					goto l562
				l565:
					position, tokenIndex = position562, tokenIndex562
					if !_rules[ruleAuthorPrefix]() {
						goto l551
					}
				}
			l562:
				add(ruleAuthorWord, position552)
			}
			return true
		l551:
			position, tokenIndex = position551, tokenIndex551
			return false
		},
		/* 76 AuthorEtAl <- <(('a' 'r' 'g' '.') / ('e' 't' ' ' 'a' 'l' '.' '{' '?' '}') / ((('e' 't') / '&') (' ' 'a' 'l') '.'?))> */
		func() bool {
			position566, tokenIndex566 := position, tokenIndex
			{
				position567 := position
				{
					position568, tokenIndex568 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l569
					}
					position++
					if buffer[position] != rune('r') {
						goto l569
					}
					position++
					if buffer[position] != rune('g') {
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
					if buffer[position] != rune('e') {
						goto l570
					}
					position++
					if buffer[position] != rune('t') {
						goto l570
					}
					position++
					if buffer[position] != rune(' ') {
						goto l570
					}
					position++
					if buffer[position] != rune('a') {
						goto l570
					}
					position++
					if buffer[position] != rune('l') {
						goto l570
					}
					position++
					if buffer[position] != rune('.') {
						goto l570
					}
					position++
					if buffer[position] != rune('{') {
						goto l570
					}
					position++
					if buffer[position] != rune('?') {
						goto l570
					}
					position++
					if buffer[position] != rune('}') {
						goto l570
					}
					position++
					goto l568
				l570:
					position, tokenIndex = position568, tokenIndex568
					{
						position571, tokenIndex571 := position, tokenIndex
						if buffer[position] != rune('e') {
							goto l572
						}
						position++
						if buffer[position] != rune('t') {
							goto l572
						}
						position++
						goto l571
					l572:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('&') {
							goto l566
						}
						position++
					}
				l571:
					if buffer[position] != rune(' ') {
						goto l566
					}
					position++
					if buffer[position] != rune('a') {
						goto l566
					}
					position++
					if buffer[position] != rune('l') {
						goto l566
					}
					position++
					{
						position573, tokenIndex573 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l573
						}
						position++
						goto l574
					l573:
						position, tokenIndex = position573, tokenIndex573
					}
				l574:
				}
			l568:
				add(ruleAuthorEtAl, position567)
			}
			return true
		l566:
			position, tokenIndex = position566, tokenIndex566
			return false
		},
		/* 77 AuthorWord2 <- <(AuthorWord3 Dash AuthorWordSoft)> */
		func() bool {
			position575, tokenIndex575 := position, tokenIndex
			{
				position576 := position
				if !_rules[ruleAuthorWord3]() {
					goto l575
				}
				if !_rules[ruleDash]() {
					goto l575
				}
				if !_rules[ruleAuthorWordSoft]() {
					goto l575
				}
				add(ruleAuthorWord2, position576)
			}
			return true
		l575:
			position, tokenIndex = position575, tokenIndex575
			return false
		},
		/* 78 AuthorWord3 <- <(AuthorPrefixGlued? (AllCapsAuthorWord / CapAuthorWord) '.'?)> */
		func() bool {
			position577, tokenIndex577 := position, tokenIndex
			{
				position578 := position
				{
					position579, tokenIndex579 := position, tokenIndex
					if !_rules[ruleAuthorPrefixGlued]() {
						goto l579
					}
					goto l580
				l579:
					position, tokenIndex = position579, tokenIndex579
				}
			l580:
				{
					position581, tokenIndex581 := position, tokenIndex
					if !_rules[ruleAllCapsAuthorWord]() {
						goto l582
					}
					goto l581
				l582:
					position, tokenIndex = position581, tokenIndex581
					if !_rules[ruleCapAuthorWord]() {
						goto l577
					}
				}
			l581:
				{
					position583, tokenIndex583 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l583
					}
					position++
					goto l584
				l583:
					position, tokenIndex = position583, tokenIndex583
				}
			l584:
				add(ruleAuthorWord3, position578)
			}
			return true
		l577:
			position, tokenIndex = position577, tokenIndex577
			return false
		},
		/* 79 AuthorWordSoft <- <(((AuthorUpperChar (AuthorUpperChar+ / AuthorLowerChar+)) / AuthorLowerChar+) '.'?)> */
		func() bool {
			position585, tokenIndex585 := position, tokenIndex
			{
				position586 := position
				{
					position587, tokenIndex587 := position, tokenIndex
					if !_rules[ruleAuthorUpperChar]() {
						goto l588
					}
					{
						position589, tokenIndex589 := position, tokenIndex
						if !_rules[ruleAuthorUpperChar]() {
							goto l590
						}
					l591:
						{
							position592, tokenIndex592 := position, tokenIndex
							if !_rules[ruleAuthorUpperChar]() {
								goto l592
							}
							goto l591
						l592:
							position, tokenIndex = position592, tokenIndex592
						}
						goto l589
					l590:
						position, tokenIndex = position589, tokenIndex589
						if !_rules[ruleAuthorLowerChar]() {
							goto l588
						}
					l593:
						{
							position594, tokenIndex594 := position, tokenIndex
							if !_rules[ruleAuthorLowerChar]() {
								goto l594
							}
							goto l593
						l594:
							position, tokenIndex = position594, tokenIndex594
						}
					}
				l589:
					goto l587
				l588:
					position, tokenIndex = position587, tokenIndex587
					if !_rules[ruleAuthorLowerChar]() {
						goto l585
					}
				l595:
					{
						position596, tokenIndex596 := position, tokenIndex
						if !_rules[ruleAuthorLowerChar]() {
							goto l596
						}
						goto l595
					l596:
						position, tokenIndex = position596, tokenIndex596
					}
				}
			l587:
				{
					position597, tokenIndex597 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l597
					}
					position++
					goto l598
				l597:
					position, tokenIndex = position597, tokenIndex597
				}
			l598:
				add(ruleAuthorWordSoft, position586)
			}
			return true
		l585:
			position, tokenIndex = position585, tokenIndex585
			return false
		},
		/* 80 CapAuthorWord <- <(AuthorUpperChar AuthorLowerChar*)> */
		func() bool {
			position599, tokenIndex599 := position, tokenIndex
			{
				position600 := position
				if !_rules[ruleAuthorUpperChar]() {
					goto l599
				}
			l601:
				{
					position602, tokenIndex602 := position, tokenIndex
					if !_rules[ruleAuthorLowerChar]() {
						goto l602
					}
					goto l601
				l602:
					position, tokenIndex = position602, tokenIndex602
				}
				add(ruleCapAuthorWord, position600)
			}
			return true
		l599:
			position, tokenIndex = position599, tokenIndex599
			return false
		},
		/* 81 AllCapsAuthorWord <- <(AuthorUpperChar AuthorUpperChar+)> */
		func() bool {
			position603, tokenIndex603 := position, tokenIndex
			{
				position604 := position
				if !_rules[ruleAuthorUpperChar]() {
					goto l603
				}
				if !_rules[ruleAuthorUpperChar]() {
					goto l603
				}
			l605:
				{
					position606, tokenIndex606 := position, tokenIndex
					if !_rules[ruleAuthorUpperChar]() {
						goto l606
					}
					goto l605
				l606:
					position, tokenIndex = position606, tokenIndex606
				}
				add(ruleAllCapsAuthorWord, position604)
			}
			return true
		l603:
			position, tokenIndex = position603, tokenIndex603
			return false
		},
		/* 82 Filius <- <(('f' '.') / ('f' 'i' 'l' '.') / ('f' 'i' 'l' 'i' 'u' 's'))> */
		func() bool {
			position607, tokenIndex607 := position, tokenIndex
			{
				position608 := position
				{
					position609, tokenIndex609 := position, tokenIndex
					if buffer[position] != rune('f') {
						goto l610
					}
					position++
					if buffer[position] != rune('.') {
						goto l610
					}
					position++
					goto l609
				l610:
					position, tokenIndex = position609, tokenIndex609
					if buffer[position] != rune('f') {
						goto l611
					}
					position++
					if buffer[position] != rune('i') {
						goto l611
					}
					position++
					if buffer[position] != rune('l') {
						goto l611
					}
					position++
					if buffer[position] != rune('.') {
						goto l611
					}
					position++
					goto l609
				l611:
					position, tokenIndex = position609, tokenIndex609
					if buffer[position] != rune('f') {
						goto l607
					}
					position++
					if buffer[position] != rune('i') {
						goto l607
					}
					position++
					if buffer[position] != rune('l') {
						goto l607
					}
					position++
					if buffer[position] != rune('i') {
						goto l607
					}
					position++
					if buffer[position] != rune('u') {
						goto l607
					}
					position++
					if buffer[position] != rune('s') {
						goto l607
					}
					position++
				}
			l609:
				add(ruleFilius, position608)
			}
			return true
		l607:
			position, tokenIndex = position607, tokenIndex607
			return false
		},
		/* 83 AuthorPrefixGlued <- <(('d' / 'O' / 'L') Apostrophe)> */
		func() bool {
			position612, tokenIndex612 := position, tokenIndex
			{
				position613 := position
				{
					position614, tokenIndex614 := position, tokenIndex
					if buffer[position] != rune('d') {
						goto l615
					}
					position++
					goto l614
				l615:
					position, tokenIndex = position614, tokenIndex614
					if buffer[position] != rune('O') {
						goto l616
					}
					position++
					goto l614
				l616:
					position, tokenIndex = position614, tokenIndex614
					if buffer[position] != rune('L') {
						goto l612
					}
					position++
				}
			l614:
				if !_rules[ruleApostrophe]() {
					goto l612
				}
				add(ruleAuthorPrefixGlued, position613)
			}
			return true
		l612:
			position, tokenIndex = position612, tokenIndex612
			return false
		},
		/* 84 AuthorPrefix <- <(AuthorPrefix1 / AuthorPrefix2)> */
		func() bool {
			position617, tokenIndex617 := position, tokenIndex
			{
				position618 := position
				{
					position619, tokenIndex619 := position, tokenIndex
					if !_rules[ruleAuthorPrefix1]() {
						goto l620
					}
					goto l619
				l620:
					position, tokenIndex = position619, tokenIndex619
					if !_rules[ruleAuthorPrefix2]() {
						goto l617
					}
				}
			l619:
				add(ruleAuthorPrefix, position618)
			}
			return true
		l617:
			position, tokenIndex = position617, tokenIndex617
			return false
		},
		/* 85 AuthorPrefix2 <- <(('v' '.' (_? ('d' '.'))?) / (Apostrophe 't'))> */
		func() bool {
			position621, tokenIndex621 := position, tokenIndex
			{
				position622 := position
				{
					position623, tokenIndex623 := position, tokenIndex
					if buffer[position] != rune('v') {
						goto l624
					}
					position++
					if buffer[position] != rune('.') {
						goto l624
					}
					position++
					{
						position625, tokenIndex625 := position, tokenIndex
						{
							position627, tokenIndex627 := position, tokenIndex
							if !_rules[rule_]() {
								goto l627
							}
							goto l628
						l627:
							position, tokenIndex = position627, tokenIndex627
						}
					l628:
						if buffer[position] != rune('d') {
							goto l625
						}
						position++
						if buffer[position] != rune('.') {
							goto l625
						}
						position++
						goto l626
					l625:
						position, tokenIndex = position625, tokenIndex625
					}
				l626:
					goto l623
				l624:
					position, tokenIndex = position623, tokenIndex623
					if !_rules[ruleApostrophe]() {
						goto l621
					}
					if buffer[position] != rune('t') {
						goto l621
					}
					position++
				}
			l623:
				add(ruleAuthorPrefix2, position622)
			}
			return true
		l621:
			position, tokenIndex = position621, tokenIndex621
			return false
		},
		/* 86 AuthorPrefix1 <- <((('a' 'b') / ('a' 'f') / ('b' 'i' 's') / ('d' 'a') / ('d' 'e' 'r') / ('d' 'e' 's') / ('d' 'e' 'n') / ('d' 'e' 'l') / ('d' 'e' 'l' 'l' 'a') / ('d' 'e' 'l' 'a') / ('d' 'e') / ('d' 'i') / ('d' 'u') / ('e' 'l') / ('l' 'a') / ('l' 'e') / ('t' 'e' 'r') / ('v' 'a' 'n') / ('d' Apostrophe) / ('i' 'n' Apostrophe 't') / ('z' 'u' 'r') / ('v' 'o' 'n' (_ (('d' '.') / ('d' 'e' 'm')))?) / ('v' (_ 'd')?)) &_)> */
		func() bool {
			position629, tokenIndex629 := position, tokenIndex
			{
				position630 := position
				{
					position631, tokenIndex631 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l632
					}
					position++
					if buffer[position] != rune('b') {
						goto l632
					}
					position++
					goto l631
				l632:
					position, tokenIndex = position631, tokenIndex631
					if buffer[position] != rune('a') {
						goto l633
					}
					position++
					if buffer[position] != rune('f') {
						goto l633
					}
					position++
					goto l631
				l633:
					position, tokenIndex = position631, tokenIndex631
					if buffer[position] != rune('b') {
						goto l634
					}
					position++
					if buffer[position] != rune('i') {
						goto l634
					}
					position++
					if buffer[position] != rune('s') {
						goto l634
					}
					position++
					goto l631
				l634:
					position, tokenIndex = position631, tokenIndex631
					if buffer[position] != rune('d') {
						goto l635
					}
					position++
					if buffer[position] != rune('a') {
						goto l635
					}
					position++
					goto l631
				l635:
					position, tokenIndex = position631, tokenIndex631
					if buffer[position] != rune('d') {
						goto l636
					}
					position++
					if buffer[position] != rune('e') {
						goto l636
					}
					position++
					if buffer[position] != rune('r') {
						goto l636
					}
					position++
					goto l631
				l636:
					position, tokenIndex = position631, tokenIndex631
					if buffer[position] != rune('d') {
						goto l637
					}
					position++
					if buffer[position] != rune('e') {
						goto l637
					}
					position++
					if buffer[position] != rune('s') {
						goto l637
					}
					position++
					goto l631
				l637:
					position, tokenIndex = position631, tokenIndex631
					if buffer[position] != rune('d') {
						goto l638
					}
					position++
					if buffer[position] != rune('e') {
						goto l638
					}
					position++
					if buffer[position] != rune('n') {
						goto l638
					}
					position++
					goto l631
				l638:
					position, tokenIndex = position631, tokenIndex631
					if buffer[position] != rune('d') {
						goto l639
					}
					position++
					if buffer[position] != rune('e') {
						goto l639
					}
					position++
					if buffer[position] != rune('l') {
						goto l639
					}
					position++
					goto l631
				l639:
					position, tokenIndex = position631, tokenIndex631
					if buffer[position] != rune('d') {
						goto l640
					}
					position++
					if buffer[position] != rune('e') {
						goto l640
					}
					position++
					if buffer[position] != rune('l') {
						goto l640
					}
					position++
					if buffer[position] != rune('l') {
						goto l640
					}
					position++
					if buffer[position] != rune('a') {
						goto l640
					}
					position++
					goto l631
				l640:
					position, tokenIndex = position631, tokenIndex631
					if buffer[position] != rune('d') {
						goto l641
					}
					position++
					if buffer[position] != rune('e') {
						goto l641
					}
					position++
					if buffer[position] != rune('l') {
						goto l641
					}
					position++
					if buffer[position] != rune('a') {
						goto l641
					}
					position++
					goto l631
				l641:
					position, tokenIndex = position631, tokenIndex631
					if buffer[position] != rune('d') {
						goto l642
					}
					position++
					if buffer[position] != rune('e') {
						goto l642
					}
					position++
					goto l631
				l642:
					position, tokenIndex = position631, tokenIndex631
					if buffer[position] != rune('d') {
						goto l643
					}
					position++
					if buffer[position] != rune('i') {
						goto l643
					}
					position++
					goto l631
				l643:
					position, tokenIndex = position631, tokenIndex631
					if buffer[position] != rune('d') {
						goto l644
					}
					position++
					if buffer[position] != rune('u') {
						goto l644
					}
					position++
					goto l631
				l644:
					position, tokenIndex = position631, tokenIndex631
					if buffer[position] != rune('e') {
						goto l645
					}
					position++
					if buffer[position] != rune('l') {
						goto l645
					}
					position++
					goto l631
				l645:
					position, tokenIndex = position631, tokenIndex631
					if buffer[position] != rune('l') {
						goto l646
					}
					position++
					if buffer[position] != rune('a') {
						goto l646
					}
					position++
					goto l631
				l646:
					position, tokenIndex = position631, tokenIndex631
					if buffer[position] != rune('l') {
						goto l647
					}
					position++
					if buffer[position] != rune('e') {
						goto l647
					}
					position++
					goto l631
				l647:
					position, tokenIndex = position631, tokenIndex631
					if buffer[position] != rune('t') {
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
					goto l631
				l648:
					position, tokenIndex = position631, tokenIndex631
					if buffer[position] != rune('v') {
						goto l649
					}
					position++
					if buffer[position] != rune('a') {
						goto l649
					}
					position++
					if buffer[position] != rune('n') {
						goto l649
					}
					position++
					goto l631
				l649:
					position, tokenIndex = position631, tokenIndex631
					if buffer[position] != rune('d') {
						goto l650
					}
					position++
					if !_rules[ruleApostrophe]() {
						goto l650
					}
					goto l631
				l650:
					position, tokenIndex = position631, tokenIndex631
					if buffer[position] != rune('i') {
						goto l651
					}
					position++
					if buffer[position] != rune('n') {
						goto l651
					}
					position++
					if !_rules[ruleApostrophe]() {
						goto l651
					}
					if buffer[position] != rune('t') {
						goto l651
					}
					position++
					goto l631
				l651:
					position, tokenIndex = position631, tokenIndex631
					if buffer[position] != rune('z') {
						goto l652
					}
					position++
					if buffer[position] != rune('u') {
						goto l652
					}
					position++
					if buffer[position] != rune('r') {
						goto l652
					}
					position++
					goto l631
				l652:
					position, tokenIndex = position631, tokenIndex631
					if buffer[position] != rune('v') {
						goto l653
					}
					position++
					if buffer[position] != rune('o') {
						goto l653
					}
					position++
					if buffer[position] != rune('n') {
						goto l653
					}
					position++
					{
						position654, tokenIndex654 := position, tokenIndex
						if !_rules[rule_]() {
							goto l654
						}
						{
							position656, tokenIndex656 := position, tokenIndex
							if buffer[position] != rune('d') {
								goto l657
							}
							position++
							if buffer[position] != rune('.') {
								goto l657
							}
							position++
							goto l656
						l657:
							position, tokenIndex = position656, tokenIndex656
							if buffer[position] != rune('d') {
								goto l654
							}
							position++
							if buffer[position] != rune('e') {
								goto l654
							}
							position++
							if buffer[position] != rune('m') {
								goto l654
							}
							position++
						}
					l656:
						goto l655
					l654:
						position, tokenIndex = position654, tokenIndex654
					}
				l655:
					goto l631
				l653:
					position, tokenIndex = position631, tokenIndex631
					if buffer[position] != rune('v') {
						goto l629
					}
					position++
					{
						position658, tokenIndex658 := position, tokenIndex
						if !_rules[rule_]() {
							goto l658
						}
						if buffer[position] != rune('d') {
							goto l658
						}
						position++
						goto l659
					l658:
						position, tokenIndex = position658, tokenIndex658
					}
				l659:
				}
			l631:
				{
					position660, tokenIndex660 := position, tokenIndex
					if !_rules[rule_]() {
						goto l629
					}
					position, tokenIndex = position660, tokenIndex660
				}
				add(ruleAuthorPrefix1, position630)
			}
			return true
		l629:
			position, tokenIndex = position629, tokenIndex629
			return false
		},
		/* 87 AuthorUpperChar <- <(UpperASCII / MiscodedChar / ('À' / 'Á' / 'Â' / 'Ã' / 'Ä' / 'Å' / 'Æ' / 'Ç' / 'È' / 'É' / 'Ê' / 'Ë' / 'Ì' / 'Í' / 'Î' / 'Ï' / 'Ð' / 'Ñ' / 'Ò' / 'Ó' / 'Ô' / 'Õ' / 'Ö' / 'Ø' / 'Ù' / 'Ú' / 'Û' / 'Ü' / 'Ý' / 'Ć' / 'Č' / 'Ď' / 'İ' / 'Ķ' / 'Ĺ' / 'ĺ' / 'Ľ' / 'ľ' / 'Ł' / 'ł' / 'Ņ' / 'Ō' / 'Ő' / 'Œ' / 'Ř' / 'Ś' / 'Ŝ' / 'Ş' / 'Š' / 'Ÿ' / 'Ź' / 'Ż' / 'Ž' / 'ƒ' / 'Ǿ' / 'Ș' / 'Ț'))> */
		func() bool {
			position661, tokenIndex661 := position, tokenIndex
			{
				position662 := position
				{
					position663, tokenIndex663 := position, tokenIndex
					if !_rules[ruleUpperASCII]() {
						goto l664
					}
					goto l663
				l664:
					position, tokenIndex = position663, tokenIndex663
					if !_rules[ruleMiscodedChar]() {
						goto l665
					}
					goto l663
				l665:
					position, tokenIndex = position663, tokenIndex663
					{
						position666, tokenIndex666 := position, tokenIndex
						if buffer[position] != rune('À') {
							goto l667
						}
						position++
						goto l666
					l667:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Á') {
							goto l668
						}
						position++
						goto l666
					l668:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Â') {
							goto l669
						}
						position++
						goto l666
					l669:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ã') {
							goto l670
						}
						position++
						goto l666
					l670:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ä') {
							goto l671
						}
						position++
						goto l666
					l671:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Å') {
							goto l672
						}
						position++
						goto l666
					l672:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Æ') {
							goto l673
						}
						position++
						goto l666
					l673:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ç') {
							goto l674
						}
						position++
						goto l666
					l674:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('È') {
							goto l675
						}
						position++
						goto l666
					l675:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('É') {
							goto l676
						}
						position++
						goto l666
					l676:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ê') {
							goto l677
						}
						position++
						goto l666
					l677:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ë') {
							goto l678
						}
						position++
						goto l666
					l678:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ì') {
							goto l679
						}
						position++
						goto l666
					l679:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Í') {
							goto l680
						}
						position++
						goto l666
					l680:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Î') {
							goto l681
						}
						position++
						goto l666
					l681:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ï') {
							goto l682
						}
						position++
						goto l666
					l682:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ð') {
							goto l683
						}
						position++
						goto l666
					l683:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ñ') {
							goto l684
						}
						position++
						goto l666
					l684:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ò') {
							goto l685
						}
						position++
						goto l666
					l685:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ó') {
							goto l686
						}
						position++
						goto l666
					l686:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ô') {
							goto l687
						}
						position++
						goto l666
					l687:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Õ') {
							goto l688
						}
						position++
						goto l666
					l688:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ö') {
							goto l689
						}
						position++
						goto l666
					l689:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ø') {
							goto l690
						}
						position++
						goto l666
					l690:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ù') {
							goto l691
						}
						position++
						goto l666
					l691:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ú') {
							goto l692
						}
						position++
						goto l666
					l692:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Û') {
							goto l693
						}
						position++
						goto l666
					l693:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ü') {
							goto l694
						}
						position++
						goto l666
					l694:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ý') {
							goto l695
						}
						position++
						goto l666
					l695:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ć') {
							goto l696
						}
						position++
						goto l666
					l696:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Č') {
							goto l697
						}
						position++
						goto l666
					l697:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ď') {
							goto l698
						}
						position++
						goto l666
					l698:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('İ') {
							goto l699
						}
						position++
						goto l666
					l699:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ķ') {
							goto l700
						}
						position++
						goto l666
					l700:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ĺ') {
							goto l701
						}
						position++
						goto l666
					l701:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('ĺ') {
							goto l702
						}
						position++
						goto l666
					l702:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ľ') {
							goto l703
						}
						position++
						goto l666
					l703:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('ľ') {
							goto l704
						}
						position++
						goto l666
					l704:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ł') {
							goto l705
						}
						position++
						goto l666
					l705:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('ł') {
							goto l706
						}
						position++
						goto l666
					l706:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ņ') {
							goto l707
						}
						position++
						goto l666
					l707:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ō') {
							goto l708
						}
						position++
						goto l666
					l708:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ő') {
							goto l709
						}
						position++
						goto l666
					l709:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Œ') {
							goto l710
						}
						position++
						goto l666
					l710:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ř') {
							goto l711
						}
						position++
						goto l666
					l711:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ś') {
							goto l712
						}
						position++
						goto l666
					l712:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ŝ') {
							goto l713
						}
						position++
						goto l666
					l713:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ş') {
							goto l714
						}
						position++
						goto l666
					l714:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Š') {
							goto l715
						}
						position++
						goto l666
					l715:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ÿ') {
							goto l716
						}
						position++
						goto l666
					l716:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ź') {
							goto l717
						}
						position++
						goto l666
					l717:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ż') {
							goto l718
						}
						position++
						goto l666
					l718:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ž') {
							goto l719
						}
						position++
						goto l666
					l719:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('ƒ') {
							goto l720
						}
						position++
						goto l666
					l720:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ǿ') {
							goto l721
						}
						position++
						goto l666
					l721:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ș') {
							goto l722
						}
						position++
						goto l666
					l722:
						position, tokenIndex = position666, tokenIndex666
						if buffer[position] != rune('Ț') {
							goto l661
						}
						position++
					}
				l666:
				}
			l663:
				add(ruleAuthorUpperChar, position662)
			}
			return true
		l661:
			position, tokenIndex = position661, tokenIndex661
			return false
		},
		/* 88 AuthorLowerChar <- <(LowerASCII / MiscodedChar / ('à' / 'á' / 'â' / 'ã' / 'ä' / 'å' / 'æ' / 'ç' / 'è' / 'é' / 'ê' / 'ë' / 'ì' / 'í' / 'î' / 'ï' / 'ð' / 'ñ' / 'ò' / 'ó' / 'ó' / 'ô' / 'õ' / 'ö' / 'ø' / 'ù' / 'ú' / 'û' / 'ü' / 'ý' / 'ÿ' / 'ā' / 'ă' / 'ą' / 'ć' / 'ĉ' / 'č' / 'ď' / 'đ' / '\'' / 'ē' / 'ĕ' / 'ė' / 'ę' / 'ě' / 'ğ' / 'ī' / 'ĭ' / 'İ' / 'ı' / 'ĺ' / 'ľ' / 'ł' / 'ń' / 'ņ' / 'ň' / 'ŏ' / 'ő' / 'œ' / 'ŕ' / 'ř' / 'ś' / 'ş' / 'š' / 'ţ' / 'ť' / 'ũ' / 'ū' / 'ŭ' / 'ů' / 'ű' / 'ź' / 'ż' / 'ž' / 'ſ' / 'ǎ' / 'ǔ' / 'ǧ' / 'ș' / 'ț' / 'ȳ' / 'ß'))> */
		func() bool {
			position723, tokenIndex723 := position, tokenIndex
			{
				position724 := position
				{
					position725, tokenIndex725 := position, tokenIndex
					if !_rules[ruleLowerASCII]() {
						goto l726
					}
					goto l725
				l726:
					position, tokenIndex = position725, tokenIndex725
					if !_rules[ruleMiscodedChar]() {
						goto l727
					}
					goto l725
				l727:
					position, tokenIndex = position725, tokenIndex725
					{
						position728, tokenIndex728 := position, tokenIndex
						if buffer[position] != rune('à') {
							goto l729
						}
						position++
						goto l728
					l729:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('á') {
							goto l730
						}
						position++
						goto l728
					l730:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('â') {
							goto l731
						}
						position++
						goto l728
					l731:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ã') {
							goto l732
						}
						position++
						goto l728
					l732:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ä') {
							goto l733
						}
						position++
						goto l728
					l733:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('å') {
							goto l734
						}
						position++
						goto l728
					l734:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('æ') {
							goto l735
						}
						position++
						goto l728
					l735:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ç') {
							goto l736
						}
						position++
						goto l728
					l736:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('è') {
							goto l737
						}
						position++
						goto l728
					l737:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('é') {
							goto l738
						}
						position++
						goto l728
					l738:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ê') {
							goto l739
						}
						position++
						goto l728
					l739:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ë') {
							goto l740
						}
						position++
						goto l728
					l740:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ì') {
							goto l741
						}
						position++
						goto l728
					l741:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('í') {
							goto l742
						}
						position++
						goto l728
					l742:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('î') {
							goto l743
						}
						position++
						goto l728
					l743:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ï') {
							goto l744
						}
						position++
						goto l728
					l744:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ð') {
							goto l745
						}
						position++
						goto l728
					l745:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ñ') {
							goto l746
						}
						position++
						goto l728
					l746:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ò') {
							goto l747
						}
						position++
						goto l728
					l747:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ó') {
							goto l748
						}
						position++
						goto l728
					l748:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ó') {
							goto l749
						}
						position++
						goto l728
					l749:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ô') {
							goto l750
						}
						position++
						goto l728
					l750:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('õ') {
							goto l751
						}
						position++
						goto l728
					l751:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ö') {
							goto l752
						}
						position++
						goto l728
					l752:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ø') {
							goto l753
						}
						position++
						goto l728
					l753:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ù') {
							goto l754
						}
						position++
						goto l728
					l754:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ú') {
							goto l755
						}
						position++
						goto l728
					l755:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('û') {
							goto l756
						}
						position++
						goto l728
					l756:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ü') {
							goto l757
						}
						position++
						goto l728
					l757:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ý') {
							goto l758
						}
						position++
						goto l728
					l758:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ÿ') {
							goto l759
						}
						position++
						goto l728
					l759:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ā') {
							goto l760
						}
						position++
						goto l728
					l760:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ă') {
							goto l761
						}
						position++
						goto l728
					l761:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ą') {
							goto l762
						}
						position++
						goto l728
					l762:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ć') {
							goto l763
						}
						position++
						goto l728
					l763:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ĉ') {
							goto l764
						}
						position++
						goto l728
					l764:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('č') {
							goto l765
						}
						position++
						goto l728
					l765:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ď') {
							goto l766
						}
						position++
						goto l728
					l766:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('đ') {
							goto l767
						}
						position++
						goto l728
					l767:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('\'') {
							goto l768
						}
						position++
						goto l728
					l768:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ē') {
							goto l769
						}
						position++
						goto l728
					l769:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ĕ') {
							goto l770
						}
						position++
						goto l728
					l770:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ė') {
							goto l771
						}
						position++
						goto l728
					l771:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ę') {
							goto l772
						}
						position++
						goto l728
					l772:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ě') {
							goto l773
						}
						position++
						goto l728
					l773:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ğ') {
							goto l774
						}
						position++
						goto l728
					l774:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ī') {
							goto l775
						}
						position++
						goto l728
					l775:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ĭ') {
							goto l776
						}
						position++
						goto l728
					l776:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('İ') {
							goto l777
						}
						position++
						goto l728
					l777:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ı') {
							goto l778
						}
						position++
						goto l728
					l778:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ĺ') {
							goto l779
						}
						position++
						goto l728
					l779:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ľ') {
							goto l780
						}
						position++
						goto l728
					l780:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ł') {
							goto l781
						}
						position++
						goto l728
					l781:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ń') {
							goto l782
						}
						position++
						goto l728
					l782:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ņ') {
							goto l783
						}
						position++
						goto l728
					l783:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ň') {
							goto l784
						}
						position++
						goto l728
					l784:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ŏ') {
							goto l785
						}
						position++
						goto l728
					l785:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ő') {
							goto l786
						}
						position++
						goto l728
					l786:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('œ') {
							goto l787
						}
						position++
						goto l728
					l787:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ŕ') {
							goto l788
						}
						position++
						goto l728
					l788:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ř') {
							goto l789
						}
						position++
						goto l728
					l789:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ś') {
							goto l790
						}
						position++
						goto l728
					l790:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ş') {
							goto l791
						}
						position++
						goto l728
					l791:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('š') {
							goto l792
						}
						position++
						goto l728
					l792:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ţ') {
							goto l793
						}
						position++
						goto l728
					l793:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ť') {
							goto l794
						}
						position++
						goto l728
					l794:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ũ') {
							goto l795
						}
						position++
						goto l728
					l795:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ū') {
							goto l796
						}
						position++
						goto l728
					l796:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ŭ') {
							goto l797
						}
						position++
						goto l728
					l797:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ů') {
							goto l798
						}
						position++
						goto l728
					l798:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ű') {
							goto l799
						}
						position++
						goto l728
					l799:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ź') {
							goto l800
						}
						position++
						goto l728
					l800:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ż') {
							goto l801
						}
						position++
						goto l728
					l801:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ž') {
							goto l802
						}
						position++
						goto l728
					l802:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ſ') {
							goto l803
						}
						position++
						goto l728
					l803:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ǎ') {
							goto l804
						}
						position++
						goto l728
					l804:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ǔ') {
							goto l805
						}
						position++
						goto l728
					l805:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ǧ') {
							goto l806
						}
						position++
						goto l728
					l806:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ș') {
							goto l807
						}
						position++
						goto l728
					l807:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ț') {
							goto l808
						}
						position++
						goto l728
					l808:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ȳ') {
							goto l809
						}
						position++
						goto l728
					l809:
						position, tokenIndex = position728, tokenIndex728
						if buffer[position] != rune('ß') {
							goto l723
						}
						position++
					}
				l728:
				}
			l725:
				add(ruleAuthorLowerChar, position724)
			}
			return true
		l723:
			position, tokenIndex = position723, tokenIndex723
			return false
		},
		/* 89 Year <- <(YearRange / YearApprox / YearWithParens / YearWithPage / YearWithDot / YearWithChar / YearNum)> */
		func() bool {
			position810, tokenIndex810 := position, tokenIndex
			{
				position811 := position
				{
					position812, tokenIndex812 := position, tokenIndex
					if !_rules[ruleYearRange]() {
						goto l813
					}
					goto l812
				l813:
					position, tokenIndex = position812, tokenIndex812
					if !_rules[ruleYearApprox]() {
						goto l814
					}
					goto l812
				l814:
					position, tokenIndex = position812, tokenIndex812
					if !_rules[ruleYearWithParens]() {
						goto l815
					}
					goto l812
				l815:
					position, tokenIndex = position812, tokenIndex812
					if !_rules[ruleYearWithPage]() {
						goto l816
					}
					goto l812
				l816:
					position, tokenIndex = position812, tokenIndex812
					if !_rules[ruleYearWithDot]() {
						goto l817
					}
					goto l812
				l817:
					position, tokenIndex = position812, tokenIndex812
					if !_rules[ruleYearWithChar]() {
						goto l818
					}
					goto l812
				l818:
					position, tokenIndex = position812, tokenIndex812
					if !_rules[ruleYearNum]() {
						goto l810
					}
				}
			l812:
				add(ruleYear, position811)
			}
			return true
		l810:
			position, tokenIndex = position810, tokenIndex810
			return false
		},
		/* 90 YearRange <- <(YearNum Dash (Nums+ ('a' / 'b' / 'c' / 'd' / 'e' / 'f' / 'g' / 'h' / 'i' / 'j' / 'k' / 'l' / 'm' / 'n' / 'o' / 'p' / 'q' / 'r' / 's' / 't' / 'u' / 'v' / 'w' / 'x' / 'y' / 'z' / '?')*))> */
		func() bool {
			position819, tokenIndex819 := position, tokenIndex
			{
				position820 := position
				if !_rules[ruleYearNum]() {
					goto l819
				}
				if !_rules[ruleDash]() {
					goto l819
				}
				if !_rules[ruleNums]() {
					goto l819
				}
			l821:
				{
					position822, tokenIndex822 := position, tokenIndex
					if !_rules[ruleNums]() {
						goto l822
					}
					goto l821
				l822:
					position, tokenIndex = position822, tokenIndex822
				}
			l823:
				{
					position824, tokenIndex824 := position, tokenIndex
					{
						position825, tokenIndex825 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l826
						}
						position++
						goto l825
					l826:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('b') {
							goto l827
						}
						position++
						goto l825
					l827:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('c') {
							goto l828
						}
						position++
						goto l825
					l828:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('d') {
							goto l829
						}
						position++
						goto l825
					l829:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('e') {
							goto l830
						}
						position++
						goto l825
					l830:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('f') {
							goto l831
						}
						position++
						goto l825
					l831:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('g') {
							goto l832
						}
						position++
						goto l825
					l832:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('h') {
							goto l833
						}
						position++
						goto l825
					l833:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('i') {
							goto l834
						}
						position++
						goto l825
					l834:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('j') {
							goto l835
						}
						position++
						goto l825
					l835:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('k') {
							goto l836
						}
						position++
						goto l825
					l836:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('l') {
							goto l837
						}
						position++
						goto l825
					l837:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('m') {
							goto l838
						}
						position++
						goto l825
					l838:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('n') {
							goto l839
						}
						position++
						goto l825
					l839:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('o') {
							goto l840
						}
						position++
						goto l825
					l840:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('p') {
							goto l841
						}
						position++
						goto l825
					l841:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('q') {
							goto l842
						}
						position++
						goto l825
					l842:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('r') {
							goto l843
						}
						position++
						goto l825
					l843:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('s') {
							goto l844
						}
						position++
						goto l825
					l844:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('t') {
							goto l845
						}
						position++
						goto l825
					l845:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('u') {
							goto l846
						}
						position++
						goto l825
					l846:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('v') {
							goto l847
						}
						position++
						goto l825
					l847:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('w') {
							goto l848
						}
						position++
						goto l825
					l848:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('x') {
							goto l849
						}
						position++
						goto l825
					l849:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('y') {
							goto l850
						}
						position++
						goto l825
					l850:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('z') {
							goto l851
						}
						position++
						goto l825
					l851:
						position, tokenIndex = position825, tokenIndex825
						if buffer[position] != rune('?') {
							goto l824
						}
						position++
					}
				l825:
					goto l823
				l824:
					position, tokenIndex = position824, tokenIndex824
				}
				add(ruleYearRange, position820)
			}
			return true
		l819:
			position, tokenIndex = position819, tokenIndex819
			return false
		},
		/* 91 YearWithDot <- <(YearNum '.')> */
		func() bool {
			position852, tokenIndex852 := position, tokenIndex
			{
				position853 := position
				if !_rules[ruleYearNum]() {
					goto l852
				}
				if buffer[position] != rune('.') {
					goto l852
				}
				position++
				add(ruleYearWithDot, position853)
			}
			return true
		l852:
			position, tokenIndex = position852, tokenIndex852
			return false
		},
		/* 92 YearApprox <- <('[' _? YearNum _? ']')> */
		func() bool {
			position854, tokenIndex854 := position, tokenIndex
			{
				position855 := position
				if buffer[position] != rune('[') {
					goto l854
				}
				position++
				{
					position856, tokenIndex856 := position, tokenIndex
					if !_rules[rule_]() {
						goto l856
					}
					goto l857
				l856:
					position, tokenIndex = position856, tokenIndex856
				}
			l857:
				if !_rules[ruleYearNum]() {
					goto l854
				}
				{
					position858, tokenIndex858 := position, tokenIndex
					if !_rules[rule_]() {
						goto l858
					}
					goto l859
				l858:
					position, tokenIndex = position858, tokenIndex858
				}
			l859:
				if buffer[position] != rune(']') {
					goto l854
				}
				position++
				add(ruleYearApprox, position855)
			}
			return true
		l854:
			position, tokenIndex = position854, tokenIndex854
			return false
		},
		/* 93 YearWithPage <- <((YearWithChar / YearNum) _? ':' _? Nums+)> */
		func() bool {
			position860, tokenIndex860 := position, tokenIndex
			{
				position861 := position
				{
					position862, tokenIndex862 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l863
					}
					goto l862
				l863:
					position, tokenIndex = position862, tokenIndex862
					if !_rules[ruleYearNum]() {
						goto l860
					}
				}
			l862:
				{
					position864, tokenIndex864 := position, tokenIndex
					if !_rules[rule_]() {
						goto l864
					}
					goto l865
				l864:
					position, tokenIndex = position864, tokenIndex864
				}
			l865:
				if buffer[position] != rune(':') {
					goto l860
				}
				position++
				{
					position866, tokenIndex866 := position, tokenIndex
					if !_rules[rule_]() {
						goto l866
					}
					goto l867
				l866:
					position, tokenIndex = position866, tokenIndex866
				}
			l867:
				if !_rules[ruleNums]() {
					goto l860
				}
			l868:
				{
					position869, tokenIndex869 := position, tokenIndex
					if !_rules[ruleNums]() {
						goto l869
					}
					goto l868
				l869:
					position, tokenIndex = position869, tokenIndex869
				}
				add(ruleYearWithPage, position861)
			}
			return true
		l860:
			position, tokenIndex = position860, tokenIndex860
			return false
		},
		/* 94 YearWithParens <- <('(' (YearWithChar / YearNum) ')')> */
		func() bool {
			position870, tokenIndex870 := position, tokenIndex
			{
				position871 := position
				if buffer[position] != rune('(') {
					goto l870
				}
				position++
				{
					position872, tokenIndex872 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l873
					}
					goto l872
				l873:
					position, tokenIndex = position872, tokenIndex872
					if !_rules[ruleYearNum]() {
						goto l870
					}
				}
			l872:
				if buffer[position] != rune(')') {
					goto l870
				}
				position++
				add(ruleYearWithParens, position871)
			}
			return true
		l870:
			position, tokenIndex = position870, tokenIndex870
			return false
		},
		/* 95 YearWithChar <- <(YearNum LowerASCII Action0)> */
		func() bool {
			position874, tokenIndex874 := position, tokenIndex
			{
				position875 := position
				if !_rules[ruleYearNum]() {
					goto l874
				}
				if !_rules[ruleLowerASCII]() {
					goto l874
				}
				if !_rules[ruleAction0]() {
					goto l874
				}
				add(ruleYearWithChar, position875)
			}
			return true
		l874:
			position, tokenIndex = position874, tokenIndex874
			return false
		},
		/* 96 YearNum <- <(('1' / '2') ('0' / '7' / '8' / '9') Nums (Nums / '?') '?'*)> */
		func() bool {
			position876, tokenIndex876 := position, tokenIndex
			{
				position877 := position
				{
					position878, tokenIndex878 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l879
					}
					position++
					goto l878
				l879:
					position, tokenIndex = position878, tokenIndex878
					if buffer[position] != rune('2') {
						goto l876
					}
					position++
				}
			l878:
				{
					position880, tokenIndex880 := position, tokenIndex
					if buffer[position] != rune('0') {
						goto l881
					}
					position++
					goto l880
				l881:
					position, tokenIndex = position880, tokenIndex880
					if buffer[position] != rune('7') {
						goto l882
					}
					position++
					goto l880
				l882:
					position, tokenIndex = position880, tokenIndex880
					if buffer[position] != rune('8') {
						goto l883
					}
					position++
					goto l880
				l883:
					position, tokenIndex = position880, tokenIndex880
					if buffer[position] != rune('9') {
						goto l876
					}
					position++
				}
			l880:
				if !_rules[ruleNums]() {
					goto l876
				}
				{
					position884, tokenIndex884 := position, tokenIndex
					if !_rules[ruleNums]() {
						goto l885
					}
					goto l884
				l885:
					position, tokenIndex = position884, tokenIndex884
					if buffer[position] != rune('?') {
						goto l876
					}
					position++
				}
			l884:
			l886:
				{
					position887, tokenIndex887 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l887
					}
					position++
					goto l886
				l887:
					position, tokenIndex = position887, tokenIndex887
				}
				add(ruleYearNum, position877)
			}
			return true
		l876:
			position, tokenIndex = position876, tokenIndex876
			return false
		},
		/* 97 NameUpperChar <- <(UpperChar / UpperCharExtended)> */
		func() bool {
			position888, tokenIndex888 := position, tokenIndex
			{
				position889 := position
				{
					position890, tokenIndex890 := position, tokenIndex
					if !_rules[ruleUpperChar]() {
						goto l891
					}
					goto l890
				l891:
					position, tokenIndex = position890, tokenIndex890
					if !_rules[ruleUpperCharExtended]() {
						goto l888
					}
				}
			l890:
				add(ruleNameUpperChar, position889)
			}
			return true
		l888:
			position, tokenIndex = position888, tokenIndex888
			return false
		},
		/* 98 UpperCharExtended <- <('Æ' / 'Œ' / 'Ö')> */
		func() bool {
			position892, tokenIndex892 := position, tokenIndex
			{
				position893 := position
				{
					position894, tokenIndex894 := position, tokenIndex
					if buffer[position] != rune('Æ') {
						goto l895
					}
					position++
					goto l894
				l895:
					position, tokenIndex = position894, tokenIndex894
					if buffer[position] != rune('Œ') {
						goto l896
					}
					position++
					goto l894
				l896:
					position, tokenIndex = position894, tokenIndex894
					if buffer[position] != rune('Ö') {
						goto l892
					}
					position++
				}
			l894:
				add(ruleUpperCharExtended, position893)
			}
			return true
		l892:
			position, tokenIndex = position892, tokenIndex892
			return false
		},
		/* 99 UpperChar <- <UpperASCII> */
		func() bool {
			position897, tokenIndex897 := position, tokenIndex
			{
				position898 := position
				if !_rules[ruleUpperASCII]() {
					goto l897
				}
				add(ruleUpperChar, position898)
			}
			return true
		l897:
			position, tokenIndex = position897, tokenIndex897
			return false
		},
		/* 100 NameLowerChar <- <(LowerChar / LowerCharExtended / MiscodedChar)> */
		func() bool {
			position899, tokenIndex899 := position, tokenIndex
			{
				position900 := position
				{
					position901, tokenIndex901 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l902
					}
					goto l901
				l902:
					position, tokenIndex = position901, tokenIndex901
					if !_rules[ruleLowerCharExtended]() {
						goto l903
					}
					goto l901
				l903:
					position, tokenIndex = position901, tokenIndex901
					if !_rules[ruleMiscodedChar]() {
						goto l899
					}
				}
			l901:
				add(ruleNameLowerChar, position900)
			}
			return true
		l899:
			position, tokenIndex = position899, tokenIndex899
			return false
		},
		/* 101 MiscodedChar <- <'�'> */
		func() bool {
			position904, tokenIndex904 := position, tokenIndex
			{
				position905 := position
				if buffer[position] != rune('�') {
					goto l904
				}
				position++
				add(ruleMiscodedChar, position905)
			}
			return true
		l904:
			position, tokenIndex = position904, tokenIndex904
			return false
		},
		/* 102 LowerCharExtended <- <('æ' / 'œ' / 'à' / 'â' / 'å' / 'ã' / 'ä' / 'á' / 'ç' / 'č' / 'é' / 'è' / 'ë' / 'í' / 'ì' / 'ï' / 'ň' / 'ñ' / 'ñ' / 'ó' / 'ò' / 'ô' / 'ø' / 'õ' / 'ö' / 'ú' / 'ù' / 'ü' / 'ŕ' / 'ř' / 'ŗ' / 'ſ' / 'š' / 'š' / 'ş' / 'ž')> */
		func() bool {
			position906, tokenIndex906 := position, tokenIndex
			{
				position907 := position
				{
					position908, tokenIndex908 := position, tokenIndex
					if buffer[position] != rune('æ') {
						goto l909
					}
					position++
					goto l908
				l909:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('œ') {
						goto l910
					}
					position++
					goto l908
				l910:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('à') {
						goto l911
					}
					position++
					goto l908
				l911:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('â') {
						goto l912
					}
					position++
					goto l908
				l912:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('å') {
						goto l913
					}
					position++
					goto l908
				l913:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('ã') {
						goto l914
					}
					position++
					goto l908
				l914:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('ä') {
						goto l915
					}
					position++
					goto l908
				l915:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('á') {
						goto l916
					}
					position++
					goto l908
				l916:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('ç') {
						goto l917
					}
					position++
					goto l908
				l917:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('č') {
						goto l918
					}
					position++
					goto l908
				l918:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('é') {
						goto l919
					}
					position++
					goto l908
				l919:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('è') {
						goto l920
					}
					position++
					goto l908
				l920:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('ë') {
						goto l921
					}
					position++
					goto l908
				l921:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('í') {
						goto l922
					}
					position++
					goto l908
				l922:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('ì') {
						goto l923
					}
					position++
					goto l908
				l923:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('ï') {
						goto l924
					}
					position++
					goto l908
				l924:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('ň') {
						goto l925
					}
					position++
					goto l908
				l925:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('ñ') {
						goto l926
					}
					position++
					goto l908
				l926:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('ñ') {
						goto l927
					}
					position++
					goto l908
				l927:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('ó') {
						goto l928
					}
					position++
					goto l908
				l928:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('ò') {
						goto l929
					}
					position++
					goto l908
				l929:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('ô') {
						goto l930
					}
					position++
					goto l908
				l930:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('ø') {
						goto l931
					}
					position++
					goto l908
				l931:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('õ') {
						goto l932
					}
					position++
					goto l908
				l932:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('ö') {
						goto l933
					}
					position++
					goto l908
				l933:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('ú') {
						goto l934
					}
					position++
					goto l908
				l934:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('ù') {
						goto l935
					}
					position++
					goto l908
				l935:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('ü') {
						goto l936
					}
					position++
					goto l908
				l936:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('ŕ') {
						goto l937
					}
					position++
					goto l908
				l937:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('ř') {
						goto l938
					}
					position++
					goto l908
				l938:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('ŗ') {
						goto l939
					}
					position++
					goto l908
				l939:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('ſ') {
						goto l940
					}
					position++
					goto l908
				l940:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('š') {
						goto l941
					}
					position++
					goto l908
				l941:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('š') {
						goto l942
					}
					position++
					goto l908
				l942:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('ş') {
						goto l943
					}
					position++
					goto l908
				l943:
					position, tokenIndex = position908, tokenIndex908
					if buffer[position] != rune('ž') {
						goto l906
					}
					position++
				}
			l908:
				add(ruleLowerCharExtended, position907)
			}
			return true
		l906:
			position, tokenIndex = position906, tokenIndex906
			return false
		},
		/* 103 LowerChar <- <LowerASCII> */
		func() bool {
			position944, tokenIndex944 := position, tokenIndex
			{
				position945 := position
				if !_rules[ruleLowerASCII]() {
					goto l944
				}
				add(ruleLowerChar, position945)
			}
			return true
		l944:
			position, tokenIndex = position944, tokenIndex944
			return false
		},
		/* 104 SpaceCharEOI <- <(_ / !.)> */
		func() bool {
			position946, tokenIndex946 := position, tokenIndex
			{
				position947 := position
				{
					position948, tokenIndex948 := position, tokenIndex
					if !_rules[rule_]() {
						goto l949
					}
					goto l948
				l949:
					position, tokenIndex = position948, tokenIndex948
					{
						position950, tokenIndex950 := position, tokenIndex
						if !matchDot() {
							goto l950
						}
						goto l946
					l950:
						position, tokenIndex = position950, tokenIndex950
					}
				}
			l948:
				add(ruleSpaceCharEOI, position947)
			}
			return true
		l946:
			position, tokenIndex = position946, tokenIndex946
			return false
		},
		/* 105 Nums <- <[0-9]> */
		func() bool {
			position951, tokenIndex951 := position, tokenIndex
			{
				position952 := position
				if c := buffer[position]; c < rune('0') || c > rune('9') {
					goto l951
				}
				position++
				add(ruleNums, position952)
			}
			return true
		l951:
			position, tokenIndex = position951, tokenIndex951
			return false
		},
		/* 106 LowerASCII <- <[a-z]> */
		func() bool {
			position953, tokenIndex953 := position, tokenIndex
			{
				position954 := position
				if c := buffer[position]; c < rune('a') || c > rune('z') {
					goto l953
				}
				position++
				add(ruleLowerASCII, position954)
			}
			return true
		l953:
			position, tokenIndex = position953, tokenIndex953
			return false
		},
		/* 107 UpperASCII <- <[A-Z]> */
		func() bool {
			position955, tokenIndex955 := position, tokenIndex
			{
				position956 := position
				if c := buffer[position]; c < rune('A') || c > rune('Z') {
					goto l955
				}
				position++
				add(ruleUpperASCII, position956)
			}
			return true
		l955:
			position, tokenIndex = position955, tokenIndex955
			return false
		},
		/* 108 Apostrophe <- <(ApostrOther / ApostrASCII)> */
		func() bool {
			position957, tokenIndex957 := position, tokenIndex
			{
				position958 := position
				{
					position959, tokenIndex959 := position, tokenIndex
					if !_rules[ruleApostrOther]() {
						goto l960
					}
					goto l959
				l960:
					position, tokenIndex = position959, tokenIndex959
					if !_rules[ruleApostrASCII]() {
						goto l957
					}
				}
			l959:
				add(ruleApostrophe, position958)
			}
			return true
		l957:
			position, tokenIndex = position957, tokenIndex957
			return false
		},
		/* 109 ApostrASCII <- <'\''> */
		func() bool {
			position961, tokenIndex961 := position, tokenIndex
			{
				position962 := position
				if buffer[position] != rune('\'') {
					goto l961
				}
				position++
				add(ruleApostrASCII, position962)
			}
			return true
		l961:
			position, tokenIndex = position961, tokenIndex961
			return false
		},
		/* 110 ApostrOther <- <('‘' / '’')> */
		func() bool {
			position963, tokenIndex963 := position, tokenIndex
			{
				position964 := position
				{
					position965, tokenIndex965 := position, tokenIndex
					if buffer[position] != rune('‘') {
						goto l966
					}
					position++
					goto l965
				l966:
					position, tokenIndex = position965, tokenIndex965
					if buffer[position] != rune('’') {
						goto l963
					}
					position++
				}
			l965:
				add(ruleApostrOther, position964)
			}
			return true
		l963:
			position, tokenIndex = position963, tokenIndex963
			return false
		},
		/* 111 Dash <- <'-'> */
		func() bool {
			position967, tokenIndex967 := position, tokenIndex
			{
				position968 := position
				if buffer[position] != rune('-') {
					goto l967
				}
				position++
				add(ruleDash, position968)
			}
			return true
		l967:
			position, tokenIndex = position967, tokenIndex967
			return false
		},
		/* 112 _ <- <(MultipleSpace / SingleSpace)> */
		func() bool {
			position969, tokenIndex969 := position, tokenIndex
			{
				position970 := position
				{
					position971, tokenIndex971 := position, tokenIndex
					if !_rules[ruleMultipleSpace]() {
						goto l972
					}
					goto l971
				l972:
					position, tokenIndex = position971, tokenIndex971
					if !_rules[ruleSingleSpace]() {
						goto l969
					}
				}
			l971:
				add(rule_, position970)
			}
			return true
		l969:
			position, tokenIndex = position969, tokenIndex969
			return false
		},
		/* 113 MultipleSpace <- <(SingleSpace SingleSpace+)> */
		func() bool {
			position973, tokenIndex973 := position, tokenIndex
			{
				position974 := position
				if !_rules[ruleSingleSpace]() {
					goto l973
				}
				if !_rules[ruleSingleSpace]() {
					goto l973
				}
			l975:
				{
					position976, tokenIndex976 := position, tokenIndex
					if !_rules[ruleSingleSpace]() {
						goto l976
					}
					goto l975
				l976:
					position, tokenIndex = position976, tokenIndex976
				}
				add(ruleMultipleSpace, position974)
			}
			return true
		l973:
			position, tokenIndex = position973, tokenIndex973
			return false
		},
		/* 114 SingleSpace <- <(' ' / OtherSpace)> */
		func() bool {
			position977, tokenIndex977 := position, tokenIndex
			{
				position978 := position
				{
					position979, tokenIndex979 := position, tokenIndex
					if buffer[position] != rune(' ') {
						goto l980
					}
					position++
					goto l979
				l980:
					position, tokenIndex = position979, tokenIndex979
					if !_rules[ruleOtherSpace]() {
						goto l977
					}
				}
			l979:
				add(ruleSingleSpace, position978)
			}
			return true
		l977:
			position, tokenIndex = position977, tokenIndex977
			return false
		},
		/* 115 OtherSpace <- <('\u3000' / '\u00a0' / '\t' / '\r' / '\n' / '\f' / '\v')> */
		func() bool {
			position981, tokenIndex981 := position, tokenIndex
			{
				position982 := position
				{
					position983, tokenIndex983 := position, tokenIndex
					if buffer[position] != rune('\u3000') {
						goto l984
					}
					position++
					goto l983
				l984:
					position, tokenIndex = position983, tokenIndex983
					if buffer[position] != rune('\u00a0') {
						goto l985
					}
					position++
					goto l983
				l985:
					position, tokenIndex = position983, tokenIndex983
					if buffer[position] != rune('\t') {
						goto l986
					}
					position++
					goto l983
				l986:
					position, tokenIndex = position983, tokenIndex983
					if buffer[position] != rune('\r') {
						goto l987
					}
					position++
					goto l983
				l987:
					position, tokenIndex = position983, tokenIndex983
					if buffer[position] != rune('\n') {
						goto l988
					}
					position++
					goto l983
				l988:
					position, tokenIndex = position983, tokenIndex983
					if buffer[position] != rune('\f') {
						goto l989
					}
					position++
					goto l983
				l989:
					position, tokenIndex = position983, tokenIndex983
					if buffer[position] != rune('\v') {
						goto l981
					}
					position++
				}
			l983:
				add(ruleOtherSpace, position982)
			}
			return true
		l981:
			position, tokenIndex = position981, tokenIndex981
			return false
		},
		/* 117 Action0 <- <{ p.AddWarn(YearCharWarn) }> */
		func() bool {
			{
				add(ruleAction0, position)
			}
			return true
		},
	}
	p.rules = _rules
}
