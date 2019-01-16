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
	ruleRankAgamo
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
	rules  [111]func() bool
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
		/* 7 NamedSpeciesHybrid <- <(GenusWord (_ SubGenus)? (_ Comparison)? _ HybridChar _? SpeciesEpithet)> */
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
				add(ruleNamedSpeciesHybrid, position40)
			}
			return true
		l39:
			position, tokenIndex = position39, tokenIndex39
			return false
		},
		/* 8 NamedGenusHybrid <- <(HybridChar _? SingleName)> */
		func() bool {
			position47, tokenIndex47 := position, tokenIndex
			{
				position48 := position
				if !_rules[ruleHybridChar]() {
					goto l47
				}
				{
					position49, tokenIndex49 := position, tokenIndex
					if !_rules[rule_]() {
						goto l49
					}
					goto l50
				l49:
					position, tokenIndex = position49, tokenIndex49
				}
			l50:
				if !_rules[ruleSingleName]() {
					goto l47
				}
				add(ruleNamedGenusHybrid, position48)
			}
			return true
		l47:
			position, tokenIndex = position47, tokenIndex47
			return false
		},
		/* 9 SingleName <- <(NameComp / NameApprox / NameSpecies / NameUninomial)> */
		func() bool {
			position51, tokenIndex51 := position, tokenIndex
			{
				position52 := position
				{
					position53, tokenIndex53 := position, tokenIndex
					if !_rules[ruleNameComp]() {
						goto l54
					}
					goto l53
				l54:
					position, tokenIndex = position53, tokenIndex53
					if !_rules[ruleNameApprox]() {
						goto l55
					}
					goto l53
				l55:
					position, tokenIndex = position53, tokenIndex53
					if !_rules[ruleNameSpecies]() {
						goto l56
					}
					goto l53
				l56:
					position, tokenIndex = position53, tokenIndex53
					if !_rules[ruleNameUninomial]() {
						goto l51
					}
				}
			l53:
				add(ruleSingleName, position52)
			}
			return true
		l51:
			position, tokenIndex = position51, tokenIndex51
			return false
		},
		/* 10 NameUninomial <- <(UninomialCombo / Uninomial)> */
		func() bool {
			position57, tokenIndex57 := position, tokenIndex
			{
				position58 := position
				{
					position59, tokenIndex59 := position, tokenIndex
					if !_rules[ruleUninomialCombo]() {
						goto l60
					}
					goto l59
				l60:
					position, tokenIndex = position59, tokenIndex59
					if !_rules[ruleUninomial]() {
						goto l57
					}
				}
			l59:
				add(ruleNameUninomial, position58)
			}
			return true
		l57:
			position, tokenIndex = position57, tokenIndex57
			return false
		},
		/* 11 NameApprox <- <(GenusWord (_ SpeciesEpithet)? _ Approximation ApproxNameIgnored)> */
		func() bool {
			position61, tokenIndex61 := position, tokenIndex
			{
				position62 := position
				if !_rules[ruleGenusWord]() {
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
				if !_rules[rule_]() {
					goto l61
				}
				if !_rules[ruleApproximation]() {
					goto l61
				}
				if !_rules[ruleApproxNameIgnored]() {
					goto l61
				}
				add(ruleNameApprox, position62)
			}
			return true
		l61:
			position, tokenIndex = position61, tokenIndex61
			return false
		},
		/* 12 NameComp <- <(GenusWord _ Comparison (_ SpeciesEpithet)?)> */
		func() bool {
			position65, tokenIndex65 := position, tokenIndex
			{
				position66 := position
				if !_rules[ruleGenusWord]() {
					goto l65
				}
				if !_rules[rule_]() {
					goto l65
				}
				if !_rules[ruleComparison]() {
					goto l65
				}
				{
					position67, tokenIndex67 := position, tokenIndex
					if !_rules[rule_]() {
						goto l67
					}
					if !_rules[ruleSpeciesEpithet]() {
						goto l67
					}
					goto l68
				l67:
					position, tokenIndex = position67, tokenIndex67
				}
			l68:
				add(ruleNameComp, position66)
			}
			return true
		l65:
			position, tokenIndex = position65, tokenIndex65
			return false
		},
		/* 13 NameSpecies <- <(GenusWord (_? (SubGenus / SubGenusOrSuperspecies))? _ SpeciesEpithet (_ InfraspGroup)?)> */
		func() bool {
			position69, tokenIndex69 := position, tokenIndex
			{
				position70 := position
				if !_rules[ruleGenusWord]() {
					goto l69
				}
				{
					position71, tokenIndex71 := position, tokenIndex
					{
						position73, tokenIndex73 := position, tokenIndex
						if !_rules[rule_]() {
							goto l73
						}
						goto l74
					l73:
						position, tokenIndex = position73, tokenIndex73
					}
				l74:
					{
						position75, tokenIndex75 := position, tokenIndex
						if !_rules[ruleSubGenus]() {
							goto l76
						}
						goto l75
					l76:
						position, tokenIndex = position75, tokenIndex75
						if !_rules[ruleSubGenusOrSuperspecies]() {
							goto l71
						}
					}
				l75:
					goto l72
				l71:
					position, tokenIndex = position71, tokenIndex71
				}
			l72:
				if !_rules[rule_]() {
					goto l69
				}
				if !_rules[ruleSpeciesEpithet]() {
					goto l69
				}
				{
					position77, tokenIndex77 := position, tokenIndex
					if !_rules[rule_]() {
						goto l77
					}
					if !_rules[ruleInfraspGroup]() {
						goto l77
					}
					goto l78
				l77:
					position, tokenIndex = position77, tokenIndex77
				}
			l78:
				add(ruleNameSpecies, position70)
			}
			return true
		l69:
			position, tokenIndex = position69, tokenIndex69
			return false
		},
		/* 14 GenusWord <- <((AbbrGenus / UninomialWord) !(_ AuthorWord))> */
		func() bool {
			position79, tokenIndex79 := position, tokenIndex
			{
				position80 := position
				{
					position81, tokenIndex81 := position, tokenIndex
					if !_rules[ruleAbbrGenus]() {
						goto l82
					}
					goto l81
				l82:
					position, tokenIndex = position81, tokenIndex81
					if !_rules[ruleUninomialWord]() {
						goto l79
					}
				}
			l81:
				{
					position83, tokenIndex83 := position, tokenIndex
					if !_rules[rule_]() {
						goto l83
					}
					if !_rules[ruleAuthorWord]() {
						goto l83
					}
					goto l79
				l83:
					position, tokenIndex = position83, tokenIndex83
				}
				add(ruleGenusWord, position80)
			}
			return true
		l79:
			position, tokenIndex = position79, tokenIndex79
			return false
		},
		/* 15 InfraspGroup <- <(InfraspEpithet (_ InfraspEpithet)? (_ InfraspEpithet)?)> */
		func() bool {
			position84, tokenIndex84 := position, tokenIndex
			{
				position85 := position
				if !_rules[ruleInfraspEpithet]() {
					goto l84
				}
				{
					position86, tokenIndex86 := position, tokenIndex
					if !_rules[rule_]() {
						goto l86
					}
					if !_rules[ruleInfraspEpithet]() {
						goto l86
					}
					goto l87
				l86:
					position, tokenIndex = position86, tokenIndex86
				}
			l87:
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
				add(ruleInfraspGroup, position85)
			}
			return true
		l84:
			position, tokenIndex = position84, tokenIndex84
			return false
		},
		/* 16 InfraspEpithet <- <((Rank _?)? !AuthorEx Word (_ Authorship)?)> */
		func() bool {
			position90, tokenIndex90 := position, tokenIndex
			{
				position91 := position
				{
					position92, tokenIndex92 := position, tokenIndex
					if !_rules[ruleRank]() {
						goto l92
					}
					{
						position94, tokenIndex94 := position, tokenIndex
						if !_rules[rule_]() {
							goto l94
						}
						goto l95
					l94:
						position, tokenIndex = position94, tokenIndex94
					}
				l95:
					goto l93
				l92:
					position, tokenIndex = position92, tokenIndex92
				}
			l93:
				{
					position96, tokenIndex96 := position, tokenIndex
					if !_rules[ruleAuthorEx]() {
						goto l96
					}
					goto l90
				l96:
					position, tokenIndex = position96, tokenIndex96
				}
				if !_rules[ruleWord]() {
					goto l90
				}
				{
					position97, tokenIndex97 := position, tokenIndex
					if !_rules[rule_]() {
						goto l97
					}
					if !_rules[ruleAuthorship]() {
						goto l97
					}
					goto l98
				l97:
					position, tokenIndex = position97, tokenIndex97
				}
			l98:
				add(ruleInfraspEpithet, position91)
			}
			return true
		l90:
			position, tokenIndex = position90, tokenIndex90
			return false
		},
		/* 17 SpeciesEpithet <- <(!AuthorEx Word (_? Authorship)?)> */
		func() bool {
			position99, tokenIndex99 := position, tokenIndex
			{
				position100 := position
				{
					position101, tokenIndex101 := position, tokenIndex
					if !_rules[ruleAuthorEx]() {
						goto l101
					}
					goto l99
				l101:
					position, tokenIndex = position101, tokenIndex101
				}
				if !_rules[ruleWord]() {
					goto l99
				}
				{
					position102, tokenIndex102 := position, tokenIndex
					{
						position104, tokenIndex104 := position, tokenIndex
						if !_rules[rule_]() {
							goto l104
						}
						goto l105
					l104:
						position, tokenIndex = position104, tokenIndex104
					}
				l105:
					if !_rules[ruleAuthorship]() {
						goto l102
					}
					goto l103
				l102:
					position, tokenIndex = position102, tokenIndex102
				}
			l103:
				add(ruleSpeciesEpithet, position100)
			}
			return true
		l99:
			position, tokenIndex = position99, tokenIndex99
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
		/* 19 Rank <- <(RankForma / RankVar / RankSsp / RankOther / RankOtherUncommon / RankAgamo)> */
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
						goto l117
					}
					goto l112
				l117:
					position, tokenIndex = position112, tokenIndex112
					if !_rules[ruleRankAgamo]() {
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
			position118, tokenIndex118 := position, tokenIndex
			{
				position119 := position
				{
					position120, tokenIndex120 := position, tokenIndex
					if buffer[position] != rune('*') {
						goto l121
					}
					position++
					goto l120
				l121:
					position, tokenIndex = position120, tokenIndex120
					if buffer[position] != rune('n') {
						goto l122
					}
					position++
					if buffer[position] != rune('a') {
						goto l122
					}
					position++
					if buffer[position] != rune('t') {
						goto l122
					}
					position++
					goto l120
				l122:
					position, tokenIndex = position120, tokenIndex120
					if buffer[position] != rune('f') {
						goto l123
					}
					position++
					if buffer[position] != rune('.') {
						goto l123
					}
					position++
					if buffer[position] != rune('s') {
						goto l123
					}
					position++
					if buffer[position] != rune('p') {
						goto l123
					}
					position++
					goto l120
				l123:
					position, tokenIndex = position120, tokenIndex120
					if buffer[position] != rune('m') {
						goto l118
					}
					position++
					if buffer[position] != rune('u') {
						goto l118
					}
					position++
					if buffer[position] != rune('t') {
						goto l118
					}
					position++
					if buffer[position] != rune('.') {
						goto l118
					}
					position++
				}
			l120:
				{
					position124, tokenIndex124 := position, tokenIndex
					if !_rules[ruleSpaceCharEOI]() {
						goto l118
					}
					position, tokenIndex = position124, tokenIndex124
				}
				add(ruleRankOtherUncommon, position119)
			}
			return true
		l118:
			position, tokenIndex = position118, tokenIndex118
			return false
		},
		/* 21 RankOther <- <((('m' 'o' 'r' 'p' 'h' '.') / ('n' 'o' 't' 'h' 'o' 's' 'u' 'b' 's' 'p' '.') / ('c' 'o' 'n' 'v' 'a' 'r' '.') / ('p' 's' 'e' 'u' 'd' 'o' 'v' 'a' 'r' '.') / ('s' 'e' 'c' 't' '.') / ('s' 'e' 'r' '.') / ('s' 'u' 'b' 'v' 'a' 'r' '.') / ('s' 'u' 'b' 'f' '.') / ('r' 'a' 'c' 'e') / 'α' / ('β' 'β') / 'β' / 'γ' / 'δ' / 'ε' / 'φ' / 'θ' / 'μ' / ('a' '.') / ('b' '.') / ('c' '.') / ('d' '.') / ('e' '.') / ('g' '.') / ('k' '.') / ('p' 'v' '.') / ('p' 'a' 't' 'h' 'o' 'v' 'a' 'r' '.') / ('a' 'b' '.' (_? ('n' '.'))?) / ('s' 't' '.')) &SpaceCharEOI)> */
		func() bool {
			position125, tokenIndex125 := position, tokenIndex
			{
				position126 := position
				{
					position127, tokenIndex127 := position, tokenIndex
					if buffer[position] != rune('m') {
						goto l128
					}
					position++
					if buffer[position] != rune('o') {
						goto l128
					}
					position++
					if buffer[position] != rune('r') {
						goto l128
					}
					position++
					if buffer[position] != rune('p') {
						goto l128
					}
					position++
					if buffer[position] != rune('h') {
						goto l128
					}
					position++
					if buffer[position] != rune('.') {
						goto l128
					}
					position++
					goto l127
				l128:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('n') {
						goto l129
					}
					position++
					if buffer[position] != rune('o') {
						goto l129
					}
					position++
					if buffer[position] != rune('t') {
						goto l129
					}
					position++
					if buffer[position] != rune('h') {
						goto l129
					}
					position++
					if buffer[position] != rune('o') {
						goto l129
					}
					position++
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
					if buffer[position] != rune('.') {
						goto l129
					}
					position++
					goto l127
				l129:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('c') {
						goto l130
					}
					position++
					if buffer[position] != rune('o') {
						goto l130
					}
					position++
					if buffer[position] != rune('n') {
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
					goto l127
				l130:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('p') {
						goto l131
					}
					position++
					if buffer[position] != rune('s') {
						goto l131
					}
					position++
					if buffer[position] != rune('e') {
						goto l131
					}
					position++
					if buffer[position] != rune('u') {
						goto l131
					}
					position++
					if buffer[position] != rune('d') {
						goto l131
					}
					position++
					if buffer[position] != rune('o') {
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
					goto l127
				l131:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('s') {
						goto l132
					}
					position++
					if buffer[position] != rune('e') {
						goto l132
					}
					position++
					if buffer[position] != rune('c') {
						goto l132
					}
					position++
					if buffer[position] != rune('t') {
						goto l132
					}
					position++
					if buffer[position] != rune('.') {
						goto l132
					}
					position++
					goto l127
				l132:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('s') {
						goto l133
					}
					position++
					if buffer[position] != rune('e') {
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
					goto l127
				l133:
					position, tokenIndex = position127, tokenIndex127
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
					if buffer[position] != rune('v') {
						goto l134
					}
					position++
					if buffer[position] != rune('a') {
						goto l134
					}
					position++
					if buffer[position] != rune('r') {
						goto l134
					}
					position++
					if buffer[position] != rune('.') {
						goto l134
					}
					position++
					goto l127
				l134:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('s') {
						goto l135
					}
					position++
					if buffer[position] != rune('u') {
						goto l135
					}
					position++
					if buffer[position] != rune('b') {
						goto l135
					}
					position++
					if buffer[position] != rune('f') {
						goto l135
					}
					position++
					if buffer[position] != rune('.') {
						goto l135
					}
					position++
					goto l127
				l135:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('r') {
						goto l136
					}
					position++
					if buffer[position] != rune('a') {
						goto l136
					}
					position++
					if buffer[position] != rune('c') {
						goto l136
					}
					position++
					if buffer[position] != rune('e') {
						goto l136
					}
					position++
					goto l127
				l136:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('α') {
						goto l137
					}
					position++
					goto l127
				l137:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('β') {
						goto l138
					}
					position++
					if buffer[position] != rune('β') {
						goto l138
					}
					position++
					goto l127
				l138:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('β') {
						goto l139
					}
					position++
					goto l127
				l139:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('γ') {
						goto l140
					}
					position++
					goto l127
				l140:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('δ') {
						goto l141
					}
					position++
					goto l127
				l141:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('ε') {
						goto l142
					}
					position++
					goto l127
				l142:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('φ') {
						goto l143
					}
					position++
					goto l127
				l143:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('θ') {
						goto l144
					}
					position++
					goto l127
				l144:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('μ') {
						goto l145
					}
					position++
					goto l127
				l145:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('a') {
						goto l146
					}
					position++
					if buffer[position] != rune('.') {
						goto l146
					}
					position++
					goto l127
				l146:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('b') {
						goto l147
					}
					position++
					if buffer[position] != rune('.') {
						goto l147
					}
					position++
					goto l127
				l147:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('c') {
						goto l148
					}
					position++
					if buffer[position] != rune('.') {
						goto l148
					}
					position++
					goto l127
				l148:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('d') {
						goto l149
					}
					position++
					if buffer[position] != rune('.') {
						goto l149
					}
					position++
					goto l127
				l149:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('e') {
						goto l150
					}
					position++
					if buffer[position] != rune('.') {
						goto l150
					}
					position++
					goto l127
				l150:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('g') {
						goto l151
					}
					position++
					if buffer[position] != rune('.') {
						goto l151
					}
					position++
					goto l127
				l151:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('k') {
						goto l152
					}
					position++
					if buffer[position] != rune('.') {
						goto l152
					}
					position++
					goto l127
				l152:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('p') {
						goto l153
					}
					position++
					if buffer[position] != rune('v') {
						goto l153
					}
					position++
					if buffer[position] != rune('.') {
						goto l153
					}
					position++
					goto l127
				l153:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('p') {
						goto l154
					}
					position++
					if buffer[position] != rune('a') {
						goto l154
					}
					position++
					if buffer[position] != rune('t') {
						goto l154
					}
					position++
					if buffer[position] != rune('h') {
						goto l154
					}
					position++
					if buffer[position] != rune('o') {
						goto l154
					}
					position++
					if buffer[position] != rune('v') {
						goto l154
					}
					position++
					if buffer[position] != rune('a') {
						goto l154
					}
					position++
					if buffer[position] != rune('r') {
						goto l154
					}
					position++
					if buffer[position] != rune('.') {
						goto l154
					}
					position++
					goto l127
				l154:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('a') {
						goto l155
					}
					position++
					if buffer[position] != rune('b') {
						goto l155
					}
					position++
					if buffer[position] != rune('.') {
						goto l155
					}
					position++
					{
						position156, tokenIndex156 := position, tokenIndex
						{
							position158, tokenIndex158 := position, tokenIndex
							if !_rules[rule_]() {
								goto l158
							}
							goto l159
						l158:
							position, tokenIndex = position158, tokenIndex158
						}
					l159:
						if buffer[position] != rune('n') {
							goto l156
						}
						position++
						if buffer[position] != rune('.') {
							goto l156
						}
						position++
						goto l157
					l156:
						position, tokenIndex = position156, tokenIndex156
					}
				l157:
					goto l127
				l155:
					position, tokenIndex = position127, tokenIndex127
					if buffer[position] != rune('s') {
						goto l125
					}
					position++
					if buffer[position] != rune('t') {
						goto l125
					}
					position++
					if buffer[position] != rune('.') {
						goto l125
					}
					position++
				}
			l127:
				{
					position160, tokenIndex160 := position, tokenIndex
					if !_rules[ruleSpaceCharEOI]() {
						goto l125
					}
					position, tokenIndex = position160, tokenIndex160
				}
				add(ruleRankOther, position126)
			}
			return true
		l125:
			position, tokenIndex = position125, tokenIndex125
			return false
		},
		/* 22 RankVar <- <(('v' 'a' 'r' 'i' 'e' 't' 'y') / ('[' 'v' 'a' 'r' '.' ']') / ('n' 'v' 'a' 'r' '.') / ('v' 'a' 'r' (&SpaceCharEOI / '.')))> */
		func() bool {
			position161, tokenIndex161 := position, tokenIndex
			{
				position162 := position
				{
					position163, tokenIndex163 := position, tokenIndex
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
					if buffer[position] != rune('i') {
						goto l164
					}
					position++
					if buffer[position] != rune('e') {
						goto l164
					}
					position++
					if buffer[position] != rune('t') {
						goto l164
					}
					position++
					if buffer[position] != rune('y') {
						goto l164
					}
					position++
					goto l163
				l164:
					position, tokenIndex = position163, tokenIndex163
					if buffer[position] != rune('[') {
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
					if buffer[position] != rune(']') {
						goto l165
					}
					position++
					goto l163
				l165:
					position, tokenIndex = position163, tokenIndex163
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
					if buffer[position] != rune('.') {
						goto l166
					}
					position++
					goto l163
				l166:
					position, tokenIndex = position163, tokenIndex163
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
					{
						position167, tokenIndex167 := position, tokenIndex
						{
							position169, tokenIndex169 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l168
							}
							position, tokenIndex = position169, tokenIndex169
						}
						goto l167
					l168:
						position, tokenIndex = position167, tokenIndex167
						if buffer[position] != rune('.') {
							goto l161
						}
						position++
					}
				l167:
				}
			l163:
				add(ruleRankVar, position162)
			}
			return true
		l161:
			position, tokenIndex = position161, tokenIndex161
			return false
		},
		/* 23 RankForma <- <((('f' 'o' 'r' 'm' 'a') / ('f' 'm' 'a') / ('f' 'o' 'r' 'm') / ('f' 'o') / 'f') (&SpaceCharEOI / '.'))> */
		func() bool {
			position170, tokenIndex170 := position, tokenIndex
			{
				position171 := position
				{
					position172, tokenIndex172 := position, tokenIndex
					if buffer[position] != rune('f') {
						goto l173
					}
					position++
					if buffer[position] != rune('o') {
						goto l173
					}
					position++
					if buffer[position] != rune('r') {
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
					goto l172
				l173:
					position, tokenIndex = position172, tokenIndex172
					if buffer[position] != rune('f') {
						goto l174
					}
					position++
					if buffer[position] != rune('m') {
						goto l174
					}
					position++
					if buffer[position] != rune('a') {
						goto l174
					}
					position++
					goto l172
				l174:
					position, tokenIndex = position172, tokenIndex172
					if buffer[position] != rune('f') {
						goto l175
					}
					position++
					if buffer[position] != rune('o') {
						goto l175
					}
					position++
					if buffer[position] != rune('r') {
						goto l175
					}
					position++
					if buffer[position] != rune('m') {
						goto l175
					}
					position++
					goto l172
				l175:
					position, tokenIndex = position172, tokenIndex172
					if buffer[position] != rune('f') {
						goto l176
					}
					position++
					if buffer[position] != rune('o') {
						goto l176
					}
					position++
					goto l172
				l176:
					position, tokenIndex = position172, tokenIndex172
					if buffer[position] != rune('f') {
						goto l170
					}
					position++
				}
			l172:
				{
					position177, tokenIndex177 := position, tokenIndex
					{
						position179, tokenIndex179 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l178
						}
						position, tokenIndex = position179, tokenIndex179
					}
					goto l177
				l178:
					position, tokenIndex = position177, tokenIndex177
					if buffer[position] != rune('.') {
						goto l170
					}
					position++
				}
			l177:
				add(ruleRankForma, position171)
			}
			return true
		l170:
			position, tokenIndex = position170, tokenIndex170
			return false
		},
		/* 24 RankSsp <- <((('s' 's' 'p') / ('s' 'u' 'b' 's' 'p')) (&SpaceCharEOI / '.'))> */
		func() bool {
			position180, tokenIndex180 := position, tokenIndex
			{
				position181 := position
				{
					position182, tokenIndex182 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l183
					}
					position++
					if buffer[position] != rune('s') {
						goto l183
					}
					position++
					if buffer[position] != rune('p') {
						goto l183
					}
					position++
					goto l182
				l183:
					position, tokenIndex = position182, tokenIndex182
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
					if buffer[position] != rune('p') {
						goto l180
					}
					position++
				}
			l182:
				{
					position184, tokenIndex184 := position, tokenIndex
					{
						position186, tokenIndex186 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l185
						}
						position, tokenIndex = position186, tokenIndex186
					}
					goto l184
				l185:
					position, tokenIndex = position184, tokenIndex184
					if buffer[position] != rune('.') {
						goto l180
					}
					position++
				}
			l184:
				add(ruleRankSsp, position181)
			}
			return true
		l180:
			position, tokenIndex = position180, tokenIndex180
			return false
		},
		/* 25 RankAgamo <- <((('a' 'g' 'a' 'm' 'o' 's' 'p') / ('a' 'g' 'a' 'm' 'o' 's' 's' 'p') / ('a' 'g' 'a' 'm' 'o' 'v' 'a' 'r')) '.'?)> */
		func() bool {
			position187, tokenIndex187 := position, tokenIndex
			{
				position188 := position
				{
					position189, tokenIndex189 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l190
					}
					position++
					if buffer[position] != rune('g') {
						goto l190
					}
					position++
					if buffer[position] != rune('a') {
						goto l190
					}
					position++
					if buffer[position] != rune('m') {
						goto l190
					}
					position++
					if buffer[position] != rune('o') {
						goto l190
					}
					position++
					if buffer[position] != rune('s') {
						goto l190
					}
					position++
					if buffer[position] != rune('p') {
						goto l190
					}
					position++
					goto l189
				l190:
					position, tokenIndex = position189, tokenIndex189
					if buffer[position] != rune('a') {
						goto l191
					}
					position++
					if buffer[position] != rune('g') {
						goto l191
					}
					position++
					if buffer[position] != rune('a') {
						goto l191
					}
					position++
					if buffer[position] != rune('m') {
						goto l191
					}
					position++
					if buffer[position] != rune('o') {
						goto l191
					}
					position++
					if buffer[position] != rune('s') {
						goto l191
					}
					position++
					if buffer[position] != rune('s') {
						goto l191
					}
					position++
					if buffer[position] != rune('p') {
						goto l191
					}
					position++
					goto l189
				l191:
					position, tokenIndex = position189, tokenIndex189
					if buffer[position] != rune('a') {
						goto l187
					}
					position++
					if buffer[position] != rune('g') {
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
					if buffer[position] != rune('o') {
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
				}
			l189:
				{
					position192, tokenIndex192 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l192
					}
					position++
					goto l193
				l192:
					position, tokenIndex = position192, tokenIndex192
				}
			l193:
				add(ruleRankAgamo, position188)
			}
			return true
		l187:
			position, tokenIndex = position187, tokenIndex187
			return false
		},
		/* 26 SubGenusOrSuperspecies <- <('(' _? NameLowerChar+ _? ')')> */
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
				if !_rules[ruleNameLowerChar]() {
					goto l194
				}
			l198:
				{
					position199, tokenIndex199 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l199
					}
					goto l198
				l199:
					position, tokenIndex = position199, tokenIndex199
				}
				{
					position200, tokenIndex200 := position, tokenIndex
					if !_rules[rule_]() {
						goto l200
					}
					goto l201
				l200:
					position, tokenIndex = position200, tokenIndex200
				}
			l201:
				if buffer[position] != rune(')') {
					goto l194
				}
				position++
				add(ruleSubGenusOrSuperspecies, position195)
			}
			return true
		l194:
			position, tokenIndex = position194, tokenIndex194
			return false
		},
		/* 27 SubGenus <- <('(' _? UninomialWord _? ')')> */
		func() bool {
			position202, tokenIndex202 := position, tokenIndex
			{
				position203 := position
				if buffer[position] != rune('(') {
					goto l202
				}
				position++
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
				if !_rules[ruleUninomialWord]() {
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
				if buffer[position] != rune(')') {
					goto l202
				}
				position++
				add(ruleSubGenus, position203)
			}
			return true
		l202:
			position, tokenIndex = position202, tokenIndex202
			return false
		},
		/* 28 UninomialCombo <- <(UninomialCombo1 / UninomialCombo2)> */
		func() bool {
			position208, tokenIndex208 := position, tokenIndex
			{
				position209 := position
				{
					position210, tokenIndex210 := position, tokenIndex
					if !_rules[ruleUninomialCombo1]() {
						goto l211
					}
					goto l210
				l211:
					position, tokenIndex = position210, tokenIndex210
					if !_rules[ruleUninomialCombo2]() {
						goto l208
					}
				}
			l210:
				add(ruleUninomialCombo, position209)
			}
			return true
		l208:
			position, tokenIndex = position208, tokenIndex208
			return false
		},
		/* 29 UninomialCombo1 <- <(UninomialWord _? SubGenus (_? Authorship)?)> */
		func() bool {
			position212, tokenIndex212 := position, tokenIndex
			{
				position213 := position
				if !_rules[ruleUninomialWord]() {
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
				if !_rules[ruleSubGenus]() {
					goto l212
				}
				{
					position216, tokenIndex216 := position, tokenIndex
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
					if !_rules[ruleAuthorship]() {
						goto l216
					}
					goto l217
				l216:
					position, tokenIndex = position216, tokenIndex216
				}
			l217:
				add(ruleUninomialCombo1, position213)
			}
			return true
		l212:
			position, tokenIndex = position212, tokenIndex212
			return false
		},
		/* 30 UninomialCombo2 <- <(Uninomial _ RankUninomial _ Uninomial)> */
		func() bool {
			position220, tokenIndex220 := position, tokenIndex
			{
				position221 := position
				if !_rules[ruleUninomial]() {
					goto l220
				}
				if !_rules[rule_]() {
					goto l220
				}
				if !_rules[ruleRankUninomial]() {
					goto l220
				}
				if !_rules[rule_]() {
					goto l220
				}
				if !_rules[ruleUninomial]() {
					goto l220
				}
				add(ruleUninomialCombo2, position221)
			}
			return true
		l220:
			position, tokenIndex = position220, tokenIndex220
			return false
		},
		/* 31 RankUninomial <- <((('s' 'e' 'c' 't') / ('s' 'u' 'b' 's' 'e' 'c' 't') / ('t' 'r' 'i' 'b') / ('s' 'u' 'b' 't' 'r' 'i' 'b') / ('s' 'u' 'b' 's' 'e' 'r') / ('s' 'e' 'r' '.') / ('s' 'u' 'b' 'g' 'e' 'n') / ('f' 'a' 'm') / ('s' 'u' 'b' 'f' 'a' 'm') / ('s' 'u' 'p' 'e' 'r' 't' 'r' 'i' 'b')) '.'?)> */
		func() bool {
			position222, tokenIndex222 := position, tokenIndex
			{
				position223 := position
				{
					position224, tokenIndex224 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l225
					}
					position++
					if buffer[position] != rune('e') {
						goto l225
					}
					position++
					if buffer[position] != rune('c') {
						goto l225
					}
					position++
					if buffer[position] != rune('t') {
						goto l225
					}
					position++
					goto l224
				l225:
					position, tokenIndex = position224, tokenIndex224
					if buffer[position] != rune('s') {
						goto l226
					}
					position++
					if buffer[position] != rune('u') {
						goto l226
					}
					position++
					if buffer[position] != rune('b') {
						goto l226
					}
					position++
					if buffer[position] != rune('s') {
						goto l226
					}
					position++
					if buffer[position] != rune('e') {
						goto l226
					}
					position++
					if buffer[position] != rune('c') {
						goto l226
					}
					position++
					if buffer[position] != rune('t') {
						goto l226
					}
					position++
					goto l224
				l226:
					position, tokenIndex = position224, tokenIndex224
					if buffer[position] != rune('t') {
						goto l227
					}
					position++
					if buffer[position] != rune('r') {
						goto l227
					}
					position++
					if buffer[position] != rune('i') {
						goto l227
					}
					position++
					if buffer[position] != rune('b') {
						goto l227
					}
					position++
					goto l224
				l227:
					position, tokenIndex = position224, tokenIndex224
					if buffer[position] != rune('s') {
						goto l228
					}
					position++
					if buffer[position] != rune('u') {
						goto l228
					}
					position++
					if buffer[position] != rune('b') {
						goto l228
					}
					position++
					if buffer[position] != rune('t') {
						goto l228
					}
					position++
					if buffer[position] != rune('r') {
						goto l228
					}
					position++
					if buffer[position] != rune('i') {
						goto l228
					}
					position++
					if buffer[position] != rune('b') {
						goto l228
					}
					position++
					goto l224
				l228:
					position, tokenIndex = position224, tokenIndex224
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
					if buffer[position] != rune('s') {
						goto l229
					}
					position++
					if buffer[position] != rune('e') {
						goto l229
					}
					position++
					if buffer[position] != rune('r') {
						goto l229
					}
					position++
					goto l224
				l229:
					position, tokenIndex = position224, tokenIndex224
					if buffer[position] != rune('s') {
						goto l230
					}
					position++
					if buffer[position] != rune('e') {
						goto l230
					}
					position++
					if buffer[position] != rune('r') {
						goto l230
					}
					position++
					if buffer[position] != rune('.') {
						goto l230
					}
					position++
					goto l224
				l230:
					position, tokenIndex = position224, tokenIndex224
					if buffer[position] != rune('s') {
						goto l231
					}
					position++
					if buffer[position] != rune('u') {
						goto l231
					}
					position++
					if buffer[position] != rune('b') {
						goto l231
					}
					position++
					if buffer[position] != rune('g') {
						goto l231
					}
					position++
					if buffer[position] != rune('e') {
						goto l231
					}
					position++
					if buffer[position] != rune('n') {
						goto l231
					}
					position++
					goto l224
				l231:
					position, tokenIndex = position224, tokenIndex224
					if buffer[position] != rune('f') {
						goto l232
					}
					position++
					if buffer[position] != rune('a') {
						goto l232
					}
					position++
					if buffer[position] != rune('m') {
						goto l232
					}
					position++
					goto l224
				l232:
					position, tokenIndex = position224, tokenIndex224
					if buffer[position] != rune('s') {
						goto l233
					}
					position++
					if buffer[position] != rune('u') {
						goto l233
					}
					position++
					if buffer[position] != rune('b') {
						goto l233
					}
					position++
					if buffer[position] != rune('f') {
						goto l233
					}
					position++
					if buffer[position] != rune('a') {
						goto l233
					}
					position++
					if buffer[position] != rune('m') {
						goto l233
					}
					position++
					goto l224
				l233:
					position, tokenIndex = position224, tokenIndex224
					if buffer[position] != rune('s') {
						goto l222
					}
					position++
					if buffer[position] != rune('u') {
						goto l222
					}
					position++
					if buffer[position] != rune('p') {
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
				}
			l224:
				{
					position234, tokenIndex234 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l234
					}
					position++
					goto l235
				l234:
					position, tokenIndex = position234, tokenIndex234
				}
			l235:
				add(ruleRankUninomial, position223)
			}
			return true
		l222:
			position, tokenIndex = position222, tokenIndex222
			return false
		},
		/* 32 Uninomial <- <(UninomialWord (_ Authorship)?)> */
		func() bool {
			position236, tokenIndex236 := position, tokenIndex
			{
				position237 := position
				if !_rules[ruleUninomialWord]() {
					goto l236
				}
				{
					position238, tokenIndex238 := position, tokenIndex
					if !_rules[rule_]() {
						goto l238
					}
					if !_rules[ruleAuthorship]() {
						goto l238
					}
					goto l239
				l238:
					position, tokenIndex = position238, tokenIndex238
				}
			l239:
				add(ruleUninomial, position237)
			}
			return true
		l236:
			position, tokenIndex = position236, tokenIndex236
			return false
		},
		/* 33 UninomialWord <- <(CapWord / TwoLetterGenus)> */
		func() bool {
			position240, tokenIndex240 := position, tokenIndex
			{
				position241 := position
				{
					position242, tokenIndex242 := position, tokenIndex
					if !_rules[ruleCapWord]() {
						goto l243
					}
					goto l242
				l243:
					position, tokenIndex = position242, tokenIndex242
					if !_rules[ruleTwoLetterGenus]() {
						goto l240
					}
				}
			l242:
				add(ruleUninomialWord, position241)
			}
			return true
		l240:
			position, tokenIndex = position240, tokenIndex240
			return false
		},
		/* 34 AbbrGenus <- <(UpperChar LowerChar? '.')> */
		func() bool {
			position244, tokenIndex244 := position, tokenIndex
			{
				position245 := position
				if !_rules[ruleUpperChar]() {
					goto l244
				}
				{
					position246, tokenIndex246 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l246
					}
					goto l247
				l246:
					position, tokenIndex = position246, tokenIndex246
				}
			l247:
				if buffer[position] != rune('.') {
					goto l244
				}
				position++
				add(ruleAbbrGenus, position245)
			}
			return true
		l244:
			position, tokenIndex = position244, tokenIndex244
			return false
		},
		/* 35 CapWord <- <(CapWord2 / CapWord1)> */
		func() bool {
			position248, tokenIndex248 := position, tokenIndex
			{
				position249 := position
				{
					position250, tokenIndex250 := position, tokenIndex
					if !_rules[ruleCapWord2]() {
						goto l251
					}
					goto l250
				l251:
					position, tokenIndex = position250, tokenIndex250
					if !_rules[ruleCapWord1]() {
						goto l248
					}
				}
			l250:
				add(ruleCapWord, position249)
			}
			return true
		l248:
			position, tokenIndex = position248, tokenIndex248
			return false
		},
		/* 36 CapWord1 <- <(NameUpperChar NameLowerChar NameLowerChar+ '?'?)> */
		func() bool {
			position252, tokenIndex252 := position, tokenIndex
			{
				position253 := position
				if !_rules[ruleNameUpperChar]() {
					goto l252
				}
				if !_rules[ruleNameLowerChar]() {
					goto l252
				}
				if !_rules[ruleNameLowerChar]() {
					goto l252
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
				{
					position256, tokenIndex256 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l256
					}
					position++
					goto l257
				l256:
					position, tokenIndex = position256, tokenIndex256
				}
			l257:
				add(ruleCapWord1, position253)
			}
			return true
		l252:
			position, tokenIndex = position252, tokenIndex252
			return false
		},
		/* 37 CapWord2 <- <(CapWord1 Dash (CapWord1 / Word1))> */
		func() bool {
			position258, tokenIndex258 := position, tokenIndex
			{
				position259 := position
				if !_rules[ruleCapWord1]() {
					goto l258
				}
				if !_rules[ruleDash]() {
					goto l258
				}
				{
					position260, tokenIndex260 := position, tokenIndex
					if !_rules[ruleCapWord1]() {
						goto l261
					}
					goto l260
				l261:
					position, tokenIndex = position260, tokenIndex260
					if !_rules[ruleWord1]() {
						goto l258
					}
				}
			l260:
				add(ruleCapWord2, position259)
			}
			return true
		l258:
			position, tokenIndex = position258, tokenIndex258
			return false
		},
		/* 38 TwoLetterGenus <- <(('C' 'a') / ('E' 'a') / ('G' 'e') / ('I' 'a') / ('I' 'o') / ('I' 'x') / ('L' 'o') / ('O' 'a') / ('R' 'a') / ('T' 'y') / ('U' 'a') / ('A' 'a') / ('J' 'a') / ('Z' 'u') / ('L' 'a') / ('Q' 'u') / ('A' 's') / ('B' 'a'))> */
		func() bool {
			position262, tokenIndex262 := position, tokenIndex
			{
				position263 := position
				{
					position264, tokenIndex264 := position, tokenIndex
					if buffer[position] != rune('C') {
						goto l265
					}
					position++
					if buffer[position] != rune('a') {
						goto l265
					}
					position++
					goto l264
				l265:
					position, tokenIndex = position264, tokenIndex264
					if buffer[position] != rune('E') {
						goto l266
					}
					position++
					if buffer[position] != rune('a') {
						goto l266
					}
					position++
					goto l264
				l266:
					position, tokenIndex = position264, tokenIndex264
					if buffer[position] != rune('G') {
						goto l267
					}
					position++
					if buffer[position] != rune('e') {
						goto l267
					}
					position++
					goto l264
				l267:
					position, tokenIndex = position264, tokenIndex264
					if buffer[position] != rune('I') {
						goto l268
					}
					position++
					if buffer[position] != rune('a') {
						goto l268
					}
					position++
					goto l264
				l268:
					position, tokenIndex = position264, tokenIndex264
					if buffer[position] != rune('I') {
						goto l269
					}
					position++
					if buffer[position] != rune('o') {
						goto l269
					}
					position++
					goto l264
				l269:
					position, tokenIndex = position264, tokenIndex264
					if buffer[position] != rune('I') {
						goto l270
					}
					position++
					if buffer[position] != rune('x') {
						goto l270
					}
					position++
					goto l264
				l270:
					position, tokenIndex = position264, tokenIndex264
					if buffer[position] != rune('L') {
						goto l271
					}
					position++
					if buffer[position] != rune('o') {
						goto l271
					}
					position++
					goto l264
				l271:
					position, tokenIndex = position264, tokenIndex264
					if buffer[position] != rune('O') {
						goto l272
					}
					position++
					if buffer[position] != rune('a') {
						goto l272
					}
					position++
					goto l264
				l272:
					position, tokenIndex = position264, tokenIndex264
					if buffer[position] != rune('R') {
						goto l273
					}
					position++
					if buffer[position] != rune('a') {
						goto l273
					}
					position++
					goto l264
				l273:
					position, tokenIndex = position264, tokenIndex264
					if buffer[position] != rune('T') {
						goto l274
					}
					position++
					if buffer[position] != rune('y') {
						goto l274
					}
					position++
					goto l264
				l274:
					position, tokenIndex = position264, tokenIndex264
					if buffer[position] != rune('U') {
						goto l275
					}
					position++
					if buffer[position] != rune('a') {
						goto l275
					}
					position++
					goto l264
				l275:
					position, tokenIndex = position264, tokenIndex264
					if buffer[position] != rune('A') {
						goto l276
					}
					position++
					if buffer[position] != rune('a') {
						goto l276
					}
					position++
					goto l264
				l276:
					position, tokenIndex = position264, tokenIndex264
					if buffer[position] != rune('J') {
						goto l277
					}
					position++
					if buffer[position] != rune('a') {
						goto l277
					}
					position++
					goto l264
				l277:
					position, tokenIndex = position264, tokenIndex264
					if buffer[position] != rune('Z') {
						goto l278
					}
					position++
					if buffer[position] != rune('u') {
						goto l278
					}
					position++
					goto l264
				l278:
					position, tokenIndex = position264, tokenIndex264
					if buffer[position] != rune('L') {
						goto l279
					}
					position++
					if buffer[position] != rune('a') {
						goto l279
					}
					position++
					goto l264
				l279:
					position, tokenIndex = position264, tokenIndex264
					if buffer[position] != rune('Q') {
						goto l280
					}
					position++
					if buffer[position] != rune('u') {
						goto l280
					}
					position++
					goto l264
				l280:
					position, tokenIndex = position264, tokenIndex264
					if buffer[position] != rune('A') {
						goto l281
					}
					position++
					if buffer[position] != rune('s') {
						goto l281
					}
					position++
					goto l264
				l281:
					position, tokenIndex = position264, tokenIndex264
					if buffer[position] != rune('B') {
						goto l262
					}
					position++
					if buffer[position] != rune('a') {
						goto l262
					}
					position++
				}
			l264:
				add(ruleTwoLetterGenus, position263)
			}
			return true
		l262:
			position, tokenIndex = position262, tokenIndex262
			return false
		},
		/* 39 Word <- <(!((AuthorPrefix / RankUninomial / Approximation / Word4) SpaceCharEOI) (WordApostr / WordStartsWithDigit / Word2 / Word1) &(SpaceCharEOI / '('))> */
		func() bool {
			position282, tokenIndex282 := position, tokenIndex
			{
				position283 := position
				{
					position284, tokenIndex284 := position, tokenIndex
					{
						position285, tokenIndex285 := position, tokenIndex
						if !_rules[ruleAuthorPrefix]() {
							goto l286
						}
						goto l285
					l286:
						position, tokenIndex = position285, tokenIndex285
						if !_rules[ruleRankUninomial]() {
							goto l287
						}
						goto l285
					l287:
						position, tokenIndex = position285, tokenIndex285
						if !_rules[ruleApproximation]() {
							goto l288
						}
						goto l285
					l288:
						position, tokenIndex = position285, tokenIndex285
						if !_rules[ruleWord4]() {
							goto l284
						}
					}
				l285:
					if !_rules[ruleSpaceCharEOI]() {
						goto l284
					}
					goto l282
				l284:
					position, tokenIndex = position284, tokenIndex284
				}
				{
					position289, tokenIndex289 := position, tokenIndex
					if !_rules[ruleWordApostr]() {
						goto l290
					}
					goto l289
				l290:
					position, tokenIndex = position289, tokenIndex289
					if !_rules[ruleWordStartsWithDigit]() {
						goto l291
					}
					goto l289
				l291:
					position, tokenIndex = position289, tokenIndex289
					if !_rules[ruleWord2]() {
						goto l292
					}
					goto l289
				l292:
					position, tokenIndex = position289, tokenIndex289
					if !_rules[ruleWord1]() {
						goto l282
					}
				}
			l289:
				{
					position293, tokenIndex293 := position, tokenIndex
					{
						position294, tokenIndex294 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l295
						}
						goto l294
					l295:
						position, tokenIndex = position294, tokenIndex294
						if buffer[position] != rune('(') {
							goto l282
						}
						position++
					}
				l294:
					position, tokenIndex = position293, tokenIndex293
				}
				add(ruleWord, position283)
			}
			return true
		l282:
			position, tokenIndex = position282, tokenIndex282
			return false
		},
		/* 40 Word1 <- <((LowerASCII Dash)? NameLowerChar NameLowerChar+)> */
		func() bool {
			position296, tokenIndex296 := position, tokenIndex
			{
				position297 := position
				{
					position298, tokenIndex298 := position, tokenIndex
					if !_rules[ruleLowerASCII]() {
						goto l298
					}
					if !_rules[ruleDash]() {
						goto l298
					}
					goto l299
				l298:
					position, tokenIndex = position298, tokenIndex298
				}
			l299:
				if !_rules[ruleNameLowerChar]() {
					goto l296
				}
				if !_rules[ruleNameLowerChar]() {
					goto l296
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
				add(ruleWord1, position297)
			}
			return true
		l296:
			position, tokenIndex = position296, tokenIndex296
			return false
		},
		/* 41 WordStartsWithDigit <- <(('1' / '2' / '3' / '4' / '5' / '6' / '7' / '8' / '9') Nums? ('.' / Dash)? NameLowerChar NameLowerChar NameLowerChar NameLowerChar+)> */
		func() bool {
			position302, tokenIndex302 := position, tokenIndex
			{
				position303 := position
				{
					position304, tokenIndex304 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l305
					}
					position++
					goto l304
				l305:
					position, tokenIndex = position304, tokenIndex304
					if buffer[position] != rune('2') {
						goto l306
					}
					position++
					goto l304
				l306:
					position, tokenIndex = position304, tokenIndex304
					if buffer[position] != rune('3') {
						goto l307
					}
					position++
					goto l304
				l307:
					position, tokenIndex = position304, tokenIndex304
					if buffer[position] != rune('4') {
						goto l308
					}
					position++
					goto l304
				l308:
					position, tokenIndex = position304, tokenIndex304
					if buffer[position] != rune('5') {
						goto l309
					}
					position++
					goto l304
				l309:
					position, tokenIndex = position304, tokenIndex304
					if buffer[position] != rune('6') {
						goto l310
					}
					position++
					goto l304
				l310:
					position, tokenIndex = position304, tokenIndex304
					if buffer[position] != rune('7') {
						goto l311
					}
					position++
					goto l304
				l311:
					position, tokenIndex = position304, tokenIndex304
					if buffer[position] != rune('8') {
						goto l312
					}
					position++
					goto l304
				l312:
					position, tokenIndex = position304, tokenIndex304
					if buffer[position] != rune('9') {
						goto l302
					}
					position++
				}
			l304:
				{
					position313, tokenIndex313 := position, tokenIndex
					if !_rules[ruleNums]() {
						goto l313
					}
					goto l314
				l313:
					position, tokenIndex = position313, tokenIndex313
				}
			l314:
				{
					position315, tokenIndex315 := position, tokenIndex
					{
						position317, tokenIndex317 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l318
						}
						position++
						goto l317
					l318:
						position, tokenIndex = position317, tokenIndex317
						if !_rules[ruleDash]() {
							goto l315
						}
					}
				l317:
					goto l316
				l315:
					position, tokenIndex = position315, tokenIndex315
				}
			l316:
				if !_rules[ruleNameLowerChar]() {
					goto l302
				}
				if !_rules[ruleNameLowerChar]() {
					goto l302
				}
				if !_rules[ruleNameLowerChar]() {
					goto l302
				}
				if !_rules[ruleNameLowerChar]() {
					goto l302
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
				add(ruleWordStartsWithDigit, position303)
			}
			return true
		l302:
			position, tokenIndex = position302, tokenIndex302
			return false
		},
		/* 42 Word2 <- <(NameLowerChar+ Dash? NameLowerChar+)> */
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
				{
					position325, tokenIndex325 := position, tokenIndex
					if !_rules[ruleDash]() {
						goto l325
					}
					goto l326
				l325:
					position, tokenIndex = position325, tokenIndex325
				}
			l326:
				if !_rules[ruleNameLowerChar]() {
					goto l321
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
				add(ruleWord2, position322)
			}
			return true
		l321:
			position, tokenIndex = position321, tokenIndex321
			return false
		},
		/* 43 WordApostr <- <(NameLowerChar NameLowerChar* Apostrophe Word1)> */
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
				if !_rules[ruleApostrophe]() {
					goto l329
				}
				if !_rules[ruleWord1]() {
					goto l329
				}
				add(ruleWordApostr, position330)
			}
			return true
		l329:
			position, tokenIndex = position329, tokenIndex329
			return false
		},
		/* 44 Word4 <- <(NameLowerChar+ '.' NameLowerChar)> */
		func() bool {
			position333, tokenIndex333 := position, tokenIndex
			{
				position334 := position
				if !_rules[ruleNameLowerChar]() {
					goto l333
				}
			l335:
				{
					position336, tokenIndex336 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l336
					}
					goto l335
				l336:
					position, tokenIndex = position336, tokenIndex336
				}
				if buffer[position] != rune('.') {
					goto l333
				}
				position++
				if !_rules[ruleNameLowerChar]() {
					goto l333
				}
				add(ruleWord4, position334)
			}
			return true
		l333:
			position, tokenIndex = position333, tokenIndex333
			return false
		},
		/* 45 HybridChar <- <'×'> */
		func() bool {
			position337, tokenIndex337 := position, tokenIndex
			{
				position338 := position
				if buffer[position] != rune('×') {
					goto l337
				}
				position++
				add(ruleHybridChar, position338)
			}
			return true
		l337:
			position, tokenIndex = position337, tokenIndex337
			return false
		},
		/* 46 ApproxNameIgnored <- <.*> */
		func() bool {
			{
				position340 := position
			l341:
				{
					position342, tokenIndex342 := position, tokenIndex
					if !matchDot() {
						goto l342
					}
					goto l341
				l342:
					position, tokenIndex = position342, tokenIndex342
				}
				add(ruleApproxNameIgnored, position340)
			}
			return true
		},
		/* 47 Approximation <- <(('s' 'p' '.' _? ('n' 'r' '.')) / ('s' 'p' '.' _? ('a' 'f' 'f' '.')) / ('m' 'o' 'n' 's' 't' '.') / '?' / ((('s' 'p' 'p') / ('n' 'r') / ('s' 'p') / ('a' 'f' 'f') / ('s' 'p' 'e' 'c' 'i' 'e' 's')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position343, tokenIndex343 := position, tokenIndex
			{
				position344 := position
				{
					position345, tokenIndex345 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l346
					}
					position++
					if buffer[position] != rune('p') {
						goto l346
					}
					position++
					if buffer[position] != rune('.') {
						goto l346
					}
					position++
					{
						position347, tokenIndex347 := position, tokenIndex
						if !_rules[rule_]() {
							goto l347
						}
						goto l348
					l347:
						position, tokenIndex = position347, tokenIndex347
					}
				l348:
					if buffer[position] != rune('n') {
						goto l346
					}
					position++
					if buffer[position] != rune('r') {
						goto l346
					}
					position++
					if buffer[position] != rune('.') {
						goto l346
					}
					position++
					goto l345
				l346:
					position, tokenIndex = position345, tokenIndex345
					if buffer[position] != rune('s') {
						goto l349
					}
					position++
					if buffer[position] != rune('p') {
						goto l349
					}
					position++
					if buffer[position] != rune('.') {
						goto l349
					}
					position++
					{
						position350, tokenIndex350 := position, tokenIndex
						if !_rules[rule_]() {
							goto l350
						}
						goto l351
					l350:
						position, tokenIndex = position350, tokenIndex350
					}
				l351:
					if buffer[position] != rune('a') {
						goto l349
					}
					position++
					if buffer[position] != rune('f') {
						goto l349
					}
					position++
					if buffer[position] != rune('f') {
						goto l349
					}
					position++
					if buffer[position] != rune('.') {
						goto l349
					}
					position++
					goto l345
				l349:
					position, tokenIndex = position345, tokenIndex345
					if buffer[position] != rune('m') {
						goto l352
					}
					position++
					if buffer[position] != rune('o') {
						goto l352
					}
					position++
					if buffer[position] != rune('n') {
						goto l352
					}
					position++
					if buffer[position] != rune('s') {
						goto l352
					}
					position++
					if buffer[position] != rune('t') {
						goto l352
					}
					position++
					if buffer[position] != rune('.') {
						goto l352
					}
					position++
					goto l345
				l352:
					position, tokenIndex = position345, tokenIndex345
					if buffer[position] != rune('?') {
						goto l353
					}
					position++
					goto l345
				l353:
					position, tokenIndex = position345, tokenIndex345
					{
						position354, tokenIndex354 := position, tokenIndex
						if buffer[position] != rune('s') {
							goto l355
						}
						position++
						if buffer[position] != rune('p') {
							goto l355
						}
						position++
						if buffer[position] != rune('p') {
							goto l355
						}
						position++
						goto l354
					l355:
						position, tokenIndex = position354, tokenIndex354
						if buffer[position] != rune('n') {
							goto l356
						}
						position++
						if buffer[position] != rune('r') {
							goto l356
						}
						position++
						goto l354
					l356:
						position, tokenIndex = position354, tokenIndex354
						if buffer[position] != rune('s') {
							goto l357
						}
						position++
						if buffer[position] != rune('p') {
							goto l357
						}
						position++
						goto l354
					l357:
						position, tokenIndex = position354, tokenIndex354
						if buffer[position] != rune('a') {
							goto l358
						}
						position++
						if buffer[position] != rune('f') {
							goto l358
						}
						position++
						if buffer[position] != rune('f') {
							goto l358
						}
						position++
						goto l354
					l358:
						position, tokenIndex = position354, tokenIndex354
						if buffer[position] != rune('s') {
							goto l343
						}
						position++
						if buffer[position] != rune('p') {
							goto l343
						}
						position++
						if buffer[position] != rune('e') {
							goto l343
						}
						position++
						if buffer[position] != rune('c') {
							goto l343
						}
						position++
						if buffer[position] != rune('i') {
							goto l343
						}
						position++
						if buffer[position] != rune('e') {
							goto l343
						}
						position++
						if buffer[position] != rune('s') {
							goto l343
						}
						position++
					}
				l354:
					{
						position359, tokenIndex359 := position, tokenIndex
						{
							position361, tokenIndex361 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l360
							}
							position, tokenIndex = position361, tokenIndex361
						}
						goto l359
					l360:
						position, tokenIndex = position359, tokenIndex359
						if buffer[position] != rune('.') {
							goto l343
						}
						position++
					}
				l359:
				}
			l345:
				add(ruleApproximation, position344)
			}
			return true
		l343:
			position, tokenIndex = position343, tokenIndex343
			return false
		},
		/* 48 Authorship <- <((AuthorshipCombo / OriginalAuthorship) &(SpaceCharEOI / ';' / ','))> */
		func() bool {
			position362, tokenIndex362 := position, tokenIndex
			{
				position363 := position
				{
					position364, tokenIndex364 := position, tokenIndex
					if !_rules[ruleAuthorshipCombo]() {
						goto l365
					}
					goto l364
				l365:
					position, tokenIndex = position364, tokenIndex364
					if !_rules[ruleOriginalAuthorship]() {
						goto l362
					}
				}
			l364:
				{
					position366, tokenIndex366 := position, tokenIndex
					{
						position367, tokenIndex367 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l368
						}
						goto l367
					l368:
						position, tokenIndex = position367, tokenIndex367
						if buffer[position] != rune(';') {
							goto l369
						}
						position++
						goto l367
					l369:
						position, tokenIndex = position367, tokenIndex367
						if buffer[position] != rune(',') {
							goto l362
						}
						position++
					}
				l367:
					position, tokenIndex = position366, tokenIndex366
				}
				add(ruleAuthorship, position363)
			}
			return true
		l362:
			position, tokenIndex = position362, tokenIndex362
			return false
		},
		/* 49 AuthorshipCombo <- <(OriginalAuthorshipComb (_? CombinationAuthorship)?)> */
		func() bool {
			position370, tokenIndex370 := position, tokenIndex
			{
				position371 := position
				if !_rules[ruleOriginalAuthorshipComb]() {
					goto l370
				}
				{
					position372, tokenIndex372 := position, tokenIndex
					{
						position374, tokenIndex374 := position, tokenIndex
						if !_rules[rule_]() {
							goto l374
						}
						goto l375
					l374:
						position, tokenIndex = position374, tokenIndex374
					}
				l375:
					if !_rules[ruleCombinationAuthorship]() {
						goto l372
					}
					goto l373
				l372:
					position, tokenIndex = position372, tokenIndex372
				}
			l373:
				add(ruleAuthorshipCombo, position371)
			}
			return true
		l370:
			position, tokenIndex = position370, tokenIndex370
			return false
		},
		/* 50 OriginalAuthorship <- <AuthorsGroup> */
		func() bool {
			position376, tokenIndex376 := position, tokenIndex
			{
				position377 := position
				if !_rules[ruleAuthorsGroup]() {
					goto l376
				}
				add(ruleOriginalAuthorship, position377)
			}
			return true
		l376:
			position, tokenIndex = position376, tokenIndex376
			return false
		},
		/* 51 OriginalAuthorshipComb <- <(BasionymAuthorshipYearMisformed / BasionymAuthorship)> */
		func() bool {
			position378, tokenIndex378 := position, tokenIndex
			{
				position379 := position
				{
					position380, tokenIndex380 := position, tokenIndex
					if !_rules[ruleBasionymAuthorshipYearMisformed]() {
						goto l381
					}
					goto l380
				l381:
					position, tokenIndex = position380, tokenIndex380
					if !_rules[ruleBasionymAuthorship]() {
						goto l378
					}
				}
			l380:
				add(ruleOriginalAuthorshipComb, position379)
			}
			return true
		l378:
			position, tokenIndex = position378, tokenIndex378
			return false
		},
		/* 52 CombinationAuthorship <- <AuthorsGroup> */
		func() bool {
			position382, tokenIndex382 := position, tokenIndex
			{
				position383 := position
				if !_rules[ruleAuthorsGroup]() {
					goto l382
				}
				add(ruleCombinationAuthorship, position383)
			}
			return true
		l382:
			position, tokenIndex = position382, tokenIndex382
			return false
		},
		/* 53 BasionymAuthorshipYearMisformed <- <('(' _? AuthorsGroup _? ')' (_? ',')? _? Year)> */
		func() bool {
			position384, tokenIndex384 := position, tokenIndex
			{
				position385 := position
				if buffer[position] != rune('(') {
					goto l384
				}
				position++
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
				if !_rules[ruleAuthorsGroup]() {
					goto l384
				}
				{
					position388, tokenIndex388 := position, tokenIndex
					if !_rules[rule_]() {
						goto l388
					}
					goto l389
				l388:
					position, tokenIndex = position388, tokenIndex388
				}
			l389:
				if buffer[position] != rune(')') {
					goto l384
				}
				position++
				{
					position390, tokenIndex390 := position, tokenIndex
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
					if buffer[position] != rune(',') {
						goto l390
					}
					position++
					goto l391
				l390:
					position, tokenIndex = position390, tokenIndex390
				}
			l391:
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
				if !_rules[ruleYear]() {
					goto l384
				}
				add(ruleBasionymAuthorshipYearMisformed, position385)
			}
			return true
		l384:
			position, tokenIndex = position384, tokenIndex384
			return false
		},
		/* 54 BasionymAuthorship <- <(BasionymAuthorship1 / BasionymAuthorship2Parens)> */
		func() bool {
			position396, tokenIndex396 := position, tokenIndex
			{
				position397 := position
				{
					position398, tokenIndex398 := position, tokenIndex
					if !_rules[ruleBasionymAuthorship1]() {
						goto l399
					}
					goto l398
				l399:
					position, tokenIndex = position398, tokenIndex398
					if !_rules[ruleBasionymAuthorship2Parens]() {
						goto l396
					}
				}
			l398:
				add(ruleBasionymAuthorship, position397)
			}
			return true
		l396:
			position, tokenIndex = position396, tokenIndex396
			return false
		},
		/* 55 BasionymAuthorship1 <- <('(' _? AuthorsGroup _? ')')> */
		func() bool {
			position400, tokenIndex400 := position, tokenIndex
			{
				position401 := position
				if buffer[position] != rune('(') {
					goto l400
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
					goto l400
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
					goto l400
				}
				position++
				add(ruleBasionymAuthorship1, position401)
			}
			return true
		l400:
			position, tokenIndex = position400, tokenIndex400
			return false
		},
		/* 56 BasionymAuthorship2Parens <- <('(' _? '(' _? AuthorsGroup _? ')' _? ')')> */
		func() bool {
			position406, tokenIndex406 := position, tokenIndex
			{
				position407 := position
				if buffer[position] != rune('(') {
					goto l406
				}
				position++
				{
					position408, tokenIndex408 := position, tokenIndex
					if !_rules[rule_]() {
						goto l408
					}
					goto l409
				l408:
					position, tokenIndex = position408, tokenIndex408
				}
			l409:
				if buffer[position] != rune('(') {
					goto l406
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
				if !_rules[ruleAuthorsGroup]() {
					goto l406
				}
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
				if buffer[position] != rune(')') {
					goto l406
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
				if buffer[position] != rune(')') {
					goto l406
				}
				position++
				add(ruleBasionymAuthorship2Parens, position407)
			}
			return true
		l406:
			position, tokenIndex = position406, tokenIndex406
			return false
		},
		/* 57 AuthorsGroup <- <(AuthorsTeam (_ (AuthorEmend / AuthorEx) AuthorsTeam)?)> */
		func() bool {
			position416, tokenIndex416 := position, tokenIndex
			{
				position417 := position
				if !_rules[ruleAuthorsTeam]() {
					goto l416
				}
				{
					position418, tokenIndex418 := position, tokenIndex
					if !_rules[rule_]() {
						goto l418
					}
					{
						position420, tokenIndex420 := position, tokenIndex
						if !_rules[ruleAuthorEmend]() {
							goto l421
						}
						goto l420
					l421:
						position, tokenIndex = position420, tokenIndex420
						if !_rules[ruleAuthorEx]() {
							goto l418
						}
					}
				l420:
					if !_rules[ruleAuthorsTeam]() {
						goto l418
					}
					goto l419
				l418:
					position, tokenIndex = position418, tokenIndex418
				}
			l419:
				add(ruleAuthorsGroup, position417)
			}
			return true
		l416:
			position, tokenIndex = position416, tokenIndex416
			return false
		},
		/* 58 AuthorsTeam <- <(Author (AuthorSep Author)* (_? ','? _? Year)?)> */
		func() bool {
			position422, tokenIndex422 := position, tokenIndex
			{
				position423 := position
				if !_rules[ruleAuthor]() {
					goto l422
				}
			l424:
				{
					position425, tokenIndex425 := position, tokenIndex
					if !_rules[ruleAuthorSep]() {
						goto l425
					}
					if !_rules[ruleAuthor]() {
						goto l425
					}
					goto l424
				l425:
					position, tokenIndex = position425, tokenIndex425
				}
				{
					position426, tokenIndex426 := position, tokenIndex
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
					{
						position430, tokenIndex430 := position, tokenIndex
						if buffer[position] != rune(',') {
							goto l430
						}
						position++
						goto l431
					l430:
						position, tokenIndex = position430, tokenIndex430
					}
				l431:
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
					if !_rules[ruleYear]() {
						goto l426
					}
					goto l427
				l426:
					position, tokenIndex = position426, tokenIndex426
				}
			l427:
				add(ruleAuthorsTeam, position423)
			}
			return true
		l422:
			position, tokenIndex = position422, tokenIndex422
			return false
		},
		/* 59 AuthorSep <- <(AuthorSep1 / AuthorSep2)> */
		func() bool {
			position434, tokenIndex434 := position, tokenIndex
			{
				position435 := position
				{
					position436, tokenIndex436 := position, tokenIndex
					if !_rules[ruleAuthorSep1]() {
						goto l437
					}
					goto l436
				l437:
					position, tokenIndex = position436, tokenIndex436
					if !_rules[ruleAuthorSep2]() {
						goto l434
					}
				}
			l436:
				add(ruleAuthorSep, position435)
			}
			return true
		l434:
			position, tokenIndex = position434, tokenIndex434
			return false
		},
		/* 60 AuthorSep1 <- <(_? (',' _)? ('&' / ('e' 't') / ('a' 'n' 'd') / ('a' 'p' 'u' 'd')) _?)> */
		func() bool {
			position438, tokenIndex438 := position, tokenIndex
			{
				position439 := position
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
				{
					position442, tokenIndex442 := position, tokenIndex
					if buffer[position] != rune(',') {
						goto l442
					}
					position++
					if !_rules[rule_]() {
						goto l442
					}
					goto l443
				l442:
					position, tokenIndex = position442, tokenIndex442
				}
			l443:
				{
					position444, tokenIndex444 := position, tokenIndex
					if buffer[position] != rune('&') {
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
					goto l444
				l446:
					position, tokenIndex = position444, tokenIndex444
					if buffer[position] != rune('a') {
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
					goto l444
				l447:
					position, tokenIndex = position444, tokenIndex444
					if buffer[position] != rune('a') {
						goto l438
					}
					position++
					if buffer[position] != rune('p') {
						goto l438
					}
					position++
					if buffer[position] != rune('u') {
						goto l438
					}
					position++
					if buffer[position] != rune('d') {
						goto l438
					}
					position++
				}
			l444:
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
				add(ruleAuthorSep1, position439)
			}
			return true
		l438:
			position, tokenIndex = position438, tokenIndex438
			return false
		},
		/* 61 AuthorSep2 <- <(_? ',' _?)> */
		func() bool {
			position450, tokenIndex450 := position, tokenIndex
			{
				position451 := position
				{
					position452, tokenIndex452 := position, tokenIndex
					if !_rules[rule_]() {
						goto l452
					}
					goto l453
				l452:
					position, tokenIndex = position452, tokenIndex452
				}
			l453:
				if buffer[position] != rune(',') {
					goto l450
				}
				position++
				{
					position454, tokenIndex454 := position, tokenIndex
					if !_rules[rule_]() {
						goto l454
					}
					goto l455
				l454:
					position, tokenIndex = position454, tokenIndex454
				}
			l455:
				add(ruleAuthorSep2, position451)
			}
			return true
		l450:
			position, tokenIndex = position450, tokenIndex450
			return false
		},
		/* 62 AuthorEx <- <((('e' 'x' '.'?) / ('i' 'n')) _)> */
		func() bool {
			position456, tokenIndex456 := position, tokenIndex
			{
				position457 := position
				{
					position458, tokenIndex458 := position, tokenIndex
					if buffer[position] != rune('e') {
						goto l459
					}
					position++
					if buffer[position] != rune('x') {
						goto l459
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
					goto l458
				l459:
					position, tokenIndex = position458, tokenIndex458
					if buffer[position] != rune('i') {
						goto l456
					}
					position++
					if buffer[position] != rune('n') {
						goto l456
					}
					position++
				}
			l458:
				if !_rules[rule_]() {
					goto l456
				}
				add(ruleAuthorEx, position457)
			}
			return true
		l456:
			position, tokenIndex = position456, tokenIndex456
			return false
		},
		/* 63 AuthorEmend <- <('e' 'm' 'e' 'n' 'd' '.'? _)> */
		func() bool {
			position462, tokenIndex462 := position, tokenIndex
			{
				position463 := position
				if buffer[position] != rune('e') {
					goto l462
				}
				position++
				if buffer[position] != rune('m') {
					goto l462
				}
				position++
				if buffer[position] != rune('e') {
					goto l462
				}
				position++
				if buffer[position] != rune('n') {
					goto l462
				}
				position++
				if buffer[position] != rune('d') {
					goto l462
				}
				position++
				{
					position464, tokenIndex464 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l464
					}
					position++
					goto l465
				l464:
					position, tokenIndex = position464, tokenIndex464
				}
			l465:
				if !_rules[rule_]() {
					goto l462
				}
				add(ruleAuthorEmend, position463)
			}
			return true
		l462:
			position, tokenIndex = position462, tokenIndex462
			return false
		},
		/* 64 Author <- <(Author1 / Author2 / UnknownAuthor)> */
		func() bool {
			position466, tokenIndex466 := position, tokenIndex
			{
				position467 := position
				{
					position468, tokenIndex468 := position, tokenIndex
					if !_rules[ruleAuthor1]() {
						goto l469
					}
					goto l468
				l469:
					position, tokenIndex = position468, tokenIndex468
					if !_rules[ruleAuthor2]() {
						goto l470
					}
					goto l468
				l470:
					position, tokenIndex = position468, tokenIndex468
					if !_rules[ruleUnknownAuthor]() {
						goto l466
					}
				}
			l468:
				add(ruleAuthor, position467)
			}
			return true
		l466:
			position, tokenIndex = position466, tokenIndex466
			return false
		},
		/* 65 Author1 <- <(Author2 _? Filius)> */
		func() bool {
			position471, tokenIndex471 := position, tokenIndex
			{
				position472 := position
				if !_rules[ruleAuthor2]() {
					goto l471
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
				if !_rules[ruleFilius]() {
					goto l471
				}
				add(ruleAuthor1, position472)
			}
			return true
		l471:
			position, tokenIndex = position471, tokenIndex471
			return false
		},
		/* 66 Author2 <- <(AuthorWord (_? AuthorWord)*)> */
		func() bool {
			position475, tokenIndex475 := position, tokenIndex
			{
				position476 := position
				if !_rules[ruleAuthorWord]() {
					goto l475
				}
			l477:
				{
					position478, tokenIndex478 := position, tokenIndex
					{
						position479, tokenIndex479 := position, tokenIndex
						if !_rules[rule_]() {
							goto l479
						}
						goto l480
					l479:
						position, tokenIndex = position479, tokenIndex479
					}
				l480:
					if !_rules[ruleAuthorWord]() {
						goto l478
					}
					goto l477
				l478:
					position, tokenIndex = position478, tokenIndex478
				}
				add(ruleAuthor2, position476)
			}
			return true
		l475:
			position, tokenIndex = position475, tokenIndex475
			return false
		},
		/* 67 UnknownAuthor <- <('?' / ((('a' 'u' 'c' 't') / ('a' 'n' 'o' 'n')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position481, tokenIndex481 := position, tokenIndex
			{
				position482 := position
				{
					position483, tokenIndex483 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l484
					}
					position++
					goto l483
				l484:
					position, tokenIndex = position483, tokenIndex483
					{
						position485, tokenIndex485 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l486
						}
						position++
						if buffer[position] != rune('u') {
							goto l486
						}
						position++
						if buffer[position] != rune('c') {
							goto l486
						}
						position++
						if buffer[position] != rune('t') {
							goto l486
						}
						position++
						goto l485
					l486:
						position, tokenIndex = position485, tokenIndex485
						if buffer[position] != rune('a') {
							goto l481
						}
						position++
						if buffer[position] != rune('n') {
							goto l481
						}
						position++
						if buffer[position] != rune('o') {
							goto l481
						}
						position++
						if buffer[position] != rune('n') {
							goto l481
						}
						position++
					}
				l485:
					{
						position487, tokenIndex487 := position, tokenIndex
						{
							position489, tokenIndex489 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l488
							}
							position, tokenIndex = position489, tokenIndex489
						}
						goto l487
					l488:
						position, tokenIndex = position487, tokenIndex487
						if buffer[position] != rune('.') {
							goto l481
						}
						position++
					}
				l487:
				}
			l483:
				add(ruleUnknownAuthor, position482)
			}
			return true
		l481:
			position, tokenIndex = position481, tokenIndex481
			return false
		},
		/* 68 AuthorWord <- <(!(('b' / 'B') ('o' / 'O') ('l' / 'L') ('d' / 'D') ':') (AuthorEtAl / AuthorWord2 / AuthorWord3 / AuthorPrefix))> */
		func() bool {
			position490, tokenIndex490 := position, tokenIndex
			{
				position491 := position
				{
					position492, tokenIndex492 := position, tokenIndex
					{
						position493, tokenIndex493 := position, tokenIndex
						if buffer[position] != rune('b') {
							goto l494
						}
						position++
						goto l493
					l494:
						position, tokenIndex = position493, tokenIndex493
						if buffer[position] != rune('B') {
							goto l492
						}
						position++
					}
				l493:
					{
						position495, tokenIndex495 := position, tokenIndex
						if buffer[position] != rune('o') {
							goto l496
						}
						position++
						goto l495
					l496:
						position, tokenIndex = position495, tokenIndex495
						if buffer[position] != rune('O') {
							goto l492
						}
						position++
					}
				l495:
					{
						position497, tokenIndex497 := position, tokenIndex
						if buffer[position] != rune('l') {
							goto l498
						}
						position++
						goto l497
					l498:
						position, tokenIndex = position497, tokenIndex497
						if buffer[position] != rune('L') {
							goto l492
						}
						position++
					}
				l497:
					{
						position499, tokenIndex499 := position, tokenIndex
						if buffer[position] != rune('d') {
							goto l500
						}
						position++
						goto l499
					l500:
						position, tokenIndex = position499, tokenIndex499
						if buffer[position] != rune('D') {
							goto l492
						}
						position++
					}
				l499:
					if buffer[position] != rune(':') {
						goto l492
					}
					position++
					goto l490
				l492:
					position, tokenIndex = position492, tokenIndex492
				}
				{
					position501, tokenIndex501 := position, tokenIndex
					if !_rules[ruleAuthorEtAl]() {
						goto l502
					}
					goto l501
				l502:
					position, tokenIndex = position501, tokenIndex501
					if !_rules[ruleAuthorWord2]() {
						goto l503
					}
					goto l501
				l503:
					position, tokenIndex = position501, tokenIndex501
					if !_rules[ruleAuthorWord3]() {
						goto l504
					}
					goto l501
				l504:
					position, tokenIndex = position501, tokenIndex501
					if !_rules[ruleAuthorPrefix]() {
						goto l490
					}
				}
			l501:
				add(ruleAuthorWord, position491)
			}
			return true
		l490:
			position, tokenIndex = position490, tokenIndex490
			return false
		},
		/* 69 AuthorEtAl <- <(('a' 'r' 'g' '.') / ('e' 't' ' ' 'a' 'l' '.' '{' '?' '}') / ((('e' 't') / '&') (' ' 'a' 'l') '.'?))> */
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
					if buffer[position] != rune('r') {
						goto l508
					}
					position++
					if buffer[position] != rune('g') {
						goto l508
					}
					position++
					if buffer[position] != rune('.') {
						goto l508
					}
					position++
					goto l507
				l508:
					position, tokenIndex = position507, tokenIndex507
					if buffer[position] != rune('e') {
						goto l509
					}
					position++
					if buffer[position] != rune('t') {
						goto l509
					}
					position++
					if buffer[position] != rune(' ') {
						goto l509
					}
					position++
					if buffer[position] != rune('a') {
						goto l509
					}
					position++
					if buffer[position] != rune('l') {
						goto l509
					}
					position++
					if buffer[position] != rune('.') {
						goto l509
					}
					position++
					if buffer[position] != rune('{') {
						goto l509
					}
					position++
					if buffer[position] != rune('?') {
						goto l509
					}
					position++
					if buffer[position] != rune('}') {
						goto l509
					}
					position++
					goto l507
				l509:
					position, tokenIndex = position507, tokenIndex507
					{
						position510, tokenIndex510 := position, tokenIndex
						if buffer[position] != rune('e') {
							goto l511
						}
						position++
						if buffer[position] != rune('t') {
							goto l511
						}
						position++
						goto l510
					l511:
						position, tokenIndex = position510, tokenIndex510
						if buffer[position] != rune('&') {
							goto l505
						}
						position++
					}
				l510:
					if buffer[position] != rune(' ') {
						goto l505
					}
					position++
					if buffer[position] != rune('a') {
						goto l505
					}
					position++
					if buffer[position] != rune('l') {
						goto l505
					}
					position++
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
				}
			l507:
				add(ruleAuthorEtAl, position506)
			}
			return true
		l505:
			position, tokenIndex = position505, tokenIndex505
			return false
		},
		/* 70 AuthorWord2 <- <(AuthorWord3 Dash AuthorWordSoft)> */
		func() bool {
			position514, tokenIndex514 := position, tokenIndex
			{
				position515 := position
				if !_rules[ruleAuthorWord3]() {
					goto l514
				}
				if !_rules[ruleDash]() {
					goto l514
				}
				if !_rules[ruleAuthorWordSoft]() {
					goto l514
				}
				add(ruleAuthorWord2, position515)
			}
			return true
		l514:
			position, tokenIndex = position514, tokenIndex514
			return false
		},
		/* 71 AuthorWord3 <- <(AuthorPrefixGlued? (AllCapsAuthorWord / CapAuthorWord) '.'?)> */
		func() bool {
			position516, tokenIndex516 := position, tokenIndex
			{
				position517 := position
				{
					position518, tokenIndex518 := position, tokenIndex
					if !_rules[ruleAuthorPrefixGlued]() {
						goto l518
					}
					goto l519
				l518:
					position, tokenIndex = position518, tokenIndex518
				}
			l519:
				{
					position520, tokenIndex520 := position, tokenIndex
					if !_rules[ruleAllCapsAuthorWord]() {
						goto l521
					}
					goto l520
				l521:
					position, tokenIndex = position520, tokenIndex520
					if !_rules[ruleCapAuthorWord]() {
						goto l516
					}
				}
			l520:
				{
					position522, tokenIndex522 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l522
					}
					position++
					goto l523
				l522:
					position, tokenIndex = position522, tokenIndex522
				}
			l523:
				add(ruleAuthorWord3, position517)
			}
			return true
		l516:
			position, tokenIndex = position516, tokenIndex516
			return false
		},
		/* 72 AuthorWordSoft <- <(((AuthorUpperChar (AuthorUpperChar+ / AuthorLowerChar+)) / AuthorLowerChar+) '.'?)> */
		func() bool {
			position524, tokenIndex524 := position, tokenIndex
			{
				position525 := position
				{
					position526, tokenIndex526 := position, tokenIndex
					if !_rules[ruleAuthorUpperChar]() {
						goto l527
					}
					{
						position528, tokenIndex528 := position, tokenIndex
						if !_rules[ruleAuthorUpperChar]() {
							goto l529
						}
					l530:
						{
							position531, tokenIndex531 := position, tokenIndex
							if !_rules[ruleAuthorUpperChar]() {
								goto l531
							}
							goto l530
						l531:
							position, tokenIndex = position531, tokenIndex531
						}
						goto l528
					l529:
						position, tokenIndex = position528, tokenIndex528
						if !_rules[ruleAuthorLowerChar]() {
							goto l527
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
					}
				l528:
					goto l526
				l527:
					position, tokenIndex = position526, tokenIndex526
					if !_rules[ruleAuthorLowerChar]() {
						goto l524
					}
				l534:
					{
						position535, tokenIndex535 := position, tokenIndex
						if !_rules[ruleAuthorLowerChar]() {
							goto l535
						}
						goto l534
					l535:
						position, tokenIndex = position535, tokenIndex535
					}
				}
			l526:
				{
					position536, tokenIndex536 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l536
					}
					position++
					goto l537
				l536:
					position, tokenIndex = position536, tokenIndex536
				}
			l537:
				add(ruleAuthorWordSoft, position525)
			}
			return true
		l524:
			position, tokenIndex = position524, tokenIndex524
			return false
		},
		/* 73 CapAuthorWord <- <(AuthorUpperChar AuthorLowerChar*)> */
		func() bool {
			position538, tokenIndex538 := position, tokenIndex
			{
				position539 := position
				if !_rules[ruleAuthorUpperChar]() {
					goto l538
				}
			l540:
				{
					position541, tokenIndex541 := position, tokenIndex
					if !_rules[ruleAuthorLowerChar]() {
						goto l541
					}
					goto l540
				l541:
					position, tokenIndex = position541, tokenIndex541
				}
				add(ruleCapAuthorWord, position539)
			}
			return true
		l538:
			position, tokenIndex = position538, tokenIndex538
			return false
		},
		/* 74 AllCapsAuthorWord <- <(AuthorUpperChar AuthorUpperChar+)> */
		func() bool {
			position542, tokenIndex542 := position, tokenIndex
			{
				position543 := position
				if !_rules[ruleAuthorUpperChar]() {
					goto l542
				}
				if !_rules[ruleAuthorUpperChar]() {
					goto l542
				}
			l544:
				{
					position545, tokenIndex545 := position, tokenIndex
					if !_rules[ruleAuthorUpperChar]() {
						goto l545
					}
					goto l544
				l545:
					position, tokenIndex = position545, tokenIndex545
				}
				add(ruleAllCapsAuthorWord, position543)
			}
			return true
		l542:
			position, tokenIndex = position542, tokenIndex542
			return false
		},
		/* 75 Filius <- <(('f' '.') / ('f' 'i' 'l' '.') / ('f' 'i' 'l' 'i' 'u' 's'))> */
		func() bool {
			position546, tokenIndex546 := position, tokenIndex
			{
				position547 := position
				{
					position548, tokenIndex548 := position, tokenIndex
					if buffer[position] != rune('f') {
						goto l549
					}
					position++
					if buffer[position] != rune('.') {
						goto l549
					}
					position++
					goto l548
				l549:
					position, tokenIndex = position548, tokenIndex548
					if buffer[position] != rune('f') {
						goto l550
					}
					position++
					if buffer[position] != rune('i') {
						goto l550
					}
					position++
					if buffer[position] != rune('l') {
						goto l550
					}
					position++
					if buffer[position] != rune('.') {
						goto l550
					}
					position++
					goto l548
				l550:
					position, tokenIndex = position548, tokenIndex548
					if buffer[position] != rune('f') {
						goto l546
					}
					position++
					if buffer[position] != rune('i') {
						goto l546
					}
					position++
					if buffer[position] != rune('l') {
						goto l546
					}
					position++
					if buffer[position] != rune('i') {
						goto l546
					}
					position++
					if buffer[position] != rune('u') {
						goto l546
					}
					position++
					if buffer[position] != rune('s') {
						goto l546
					}
					position++
				}
			l548:
				add(ruleFilius, position547)
			}
			return true
		l546:
			position, tokenIndex = position546, tokenIndex546
			return false
		},
		/* 76 AuthorPrefixGlued <- <(('d' / 'O' / 'L') Apostrophe)> */
		func() bool {
			position551, tokenIndex551 := position, tokenIndex
			{
				position552 := position
				{
					position553, tokenIndex553 := position, tokenIndex
					if buffer[position] != rune('d') {
						goto l554
					}
					position++
					goto l553
				l554:
					position, tokenIndex = position553, tokenIndex553
					if buffer[position] != rune('O') {
						goto l555
					}
					position++
					goto l553
				l555:
					position, tokenIndex = position553, tokenIndex553
					if buffer[position] != rune('L') {
						goto l551
					}
					position++
				}
			l553:
				if !_rules[ruleApostrophe]() {
					goto l551
				}
				add(ruleAuthorPrefixGlued, position552)
			}
			return true
		l551:
			position, tokenIndex = position551, tokenIndex551
			return false
		},
		/* 77 AuthorPrefix <- <(AuthorPrefix1 / AuthorPrefix2)> */
		func() bool {
			position556, tokenIndex556 := position, tokenIndex
			{
				position557 := position
				{
					position558, tokenIndex558 := position, tokenIndex
					if !_rules[ruleAuthorPrefix1]() {
						goto l559
					}
					goto l558
				l559:
					position, tokenIndex = position558, tokenIndex558
					if !_rules[ruleAuthorPrefix2]() {
						goto l556
					}
				}
			l558:
				add(ruleAuthorPrefix, position557)
			}
			return true
		l556:
			position, tokenIndex = position556, tokenIndex556
			return false
		},
		/* 78 AuthorPrefix2 <- <(('v' '.' (_? ('d' '.'))?) / (Apostrophe 't'))> */
		func() bool {
			position560, tokenIndex560 := position, tokenIndex
			{
				position561 := position
				{
					position562, tokenIndex562 := position, tokenIndex
					if buffer[position] != rune('v') {
						goto l563
					}
					position++
					if buffer[position] != rune('.') {
						goto l563
					}
					position++
					{
						position564, tokenIndex564 := position, tokenIndex
						{
							position566, tokenIndex566 := position, tokenIndex
							if !_rules[rule_]() {
								goto l566
							}
							goto l567
						l566:
							position, tokenIndex = position566, tokenIndex566
						}
					l567:
						if buffer[position] != rune('d') {
							goto l564
						}
						position++
						if buffer[position] != rune('.') {
							goto l564
						}
						position++
						goto l565
					l564:
						position, tokenIndex = position564, tokenIndex564
					}
				l565:
					goto l562
				l563:
					position, tokenIndex = position562, tokenIndex562
					if !_rules[ruleApostrophe]() {
						goto l560
					}
					if buffer[position] != rune('t') {
						goto l560
					}
					position++
				}
			l562:
				add(ruleAuthorPrefix2, position561)
			}
			return true
		l560:
			position, tokenIndex = position560, tokenIndex560
			return false
		},
		/* 79 AuthorPrefix1 <- <((('a' 'b') / ('a' 'f') / ('b' 'i' 's') / ('d' 'a') / ('d' 'e' 'r') / ('d' 'e' 's') / ('d' 'e' 'n') / ('d' 'e' 'l') / ('d' 'e' 'l' 'l' 'a') / ('d' 'e' 'l' 'a') / ('d' 'e') / ('d' 'i') / ('d' 'u') / ('e' 'l') / ('l' 'a') / ('l' 'e') / ('t' 'e' 'r') / ('v' 'a' 'n') / ('d' Apostrophe) / ('i' 'n' Apostrophe 't') / ('z' 'u' 'r') / ('v' 'o' 'n' (_ (('d' '.') / ('d' 'e' 'm')))?) / ('v' (_ 'd')?)) &_)> */
		func() bool {
			position568, tokenIndex568 := position, tokenIndex
			{
				position569 := position
				{
					position570, tokenIndex570 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l571
					}
					position++
					if buffer[position] != rune('b') {
						goto l571
					}
					position++
					goto l570
				l571:
					position, tokenIndex = position570, tokenIndex570
					if buffer[position] != rune('a') {
						goto l572
					}
					position++
					if buffer[position] != rune('f') {
						goto l572
					}
					position++
					goto l570
				l572:
					position, tokenIndex = position570, tokenIndex570
					if buffer[position] != rune('b') {
						goto l573
					}
					position++
					if buffer[position] != rune('i') {
						goto l573
					}
					position++
					if buffer[position] != rune('s') {
						goto l573
					}
					position++
					goto l570
				l573:
					position, tokenIndex = position570, tokenIndex570
					if buffer[position] != rune('d') {
						goto l574
					}
					position++
					if buffer[position] != rune('a') {
						goto l574
					}
					position++
					goto l570
				l574:
					position, tokenIndex = position570, tokenIndex570
					if buffer[position] != rune('d') {
						goto l575
					}
					position++
					if buffer[position] != rune('e') {
						goto l575
					}
					position++
					if buffer[position] != rune('r') {
						goto l575
					}
					position++
					goto l570
				l575:
					position, tokenIndex = position570, tokenIndex570
					if buffer[position] != rune('d') {
						goto l576
					}
					position++
					if buffer[position] != rune('e') {
						goto l576
					}
					position++
					if buffer[position] != rune('s') {
						goto l576
					}
					position++
					goto l570
				l576:
					position, tokenIndex = position570, tokenIndex570
					if buffer[position] != rune('d') {
						goto l577
					}
					position++
					if buffer[position] != rune('e') {
						goto l577
					}
					position++
					if buffer[position] != rune('n') {
						goto l577
					}
					position++
					goto l570
				l577:
					position, tokenIndex = position570, tokenIndex570
					if buffer[position] != rune('d') {
						goto l578
					}
					position++
					if buffer[position] != rune('e') {
						goto l578
					}
					position++
					if buffer[position] != rune('l') {
						goto l578
					}
					position++
					goto l570
				l578:
					position, tokenIndex = position570, tokenIndex570
					if buffer[position] != rune('d') {
						goto l579
					}
					position++
					if buffer[position] != rune('e') {
						goto l579
					}
					position++
					if buffer[position] != rune('l') {
						goto l579
					}
					position++
					if buffer[position] != rune('l') {
						goto l579
					}
					position++
					if buffer[position] != rune('a') {
						goto l579
					}
					position++
					goto l570
				l579:
					position, tokenIndex = position570, tokenIndex570
					if buffer[position] != rune('d') {
						goto l580
					}
					position++
					if buffer[position] != rune('e') {
						goto l580
					}
					position++
					if buffer[position] != rune('l') {
						goto l580
					}
					position++
					if buffer[position] != rune('a') {
						goto l580
					}
					position++
					goto l570
				l580:
					position, tokenIndex = position570, tokenIndex570
					if buffer[position] != rune('d') {
						goto l581
					}
					position++
					if buffer[position] != rune('e') {
						goto l581
					}
					position++
					goto l570
				l581:
					position, tokenIndex = position570, tokenIndex570
					if buffer[position] != rune('d') {
						goto l582
					}
					position++
					if buffer[position] != rune('i') {
						goto l582
					}
					position++
					goto l570
				l582:
					position, tokenIndex = position570, tokenIndex570
					if buffer[position] != rune('d') {
						goto l583
					}
					position++
					if buffer[position] != rune('u') {
						goto l583
					}
					position++
					goto l570
				l583:
					position, tokenIndex = position570, tokenIndex570
					if buffer[position] != rune('e') {
						goto l584
					}
					position++
					if buffer[position] != rune('l') {
						goto l584
					}
					position++
					goto l570
				l584:
					position, tokenIndex = position570, tokenIndex570
					if buffer[position] != rune('l') {
						goto l585
					}
					position++
					if buffer[position] != rune('a') {
						goto l585
					}
					position++
					goto l570
				l585:
					position, tokenIndex = position570, tokenIndex570
					if buffer[position] != rune('l') {
						goto l586
					}
					position++
					if buffer[position] != rune('e') {
						goto l586
					}
					position++
					goto l570
				l586:
					position, tokenIndex = position570, tokenIndex570
					if buffer[position] != rune('t') {
						goto l587
					}
					position++
					if buffer[position] != rune('e') {
						goto l587
					}
					position++
					if buffer[position] != rune('r') {
						goto l587
					}
					position++
					goto l570
				l587:
					position, tokenIndex = position570, tokenIndex570
					if buffer[position] != rune('v') {
						goto l588
					}
					position++
					if buffer[position] != rune('a') {
						goto l588
					}
					position++
					if buffer[position] != rune('n') {
						goto l588
					}
					position++
					goto l570
				l588:
					position, tokenIndex = position570, tokenIndex570
					if buffer[position] != rune('d') {
						goto l589
					}
					position++
					if !_rules[ruleApostrophe]() {
						goto l589
					}
					goto l570
				l589:
					position, tokenIndex = position570, tokenIndex570
					if buffer[position] != rune('i') {
						goto l590
					}
					position++
					if buffer[position] != rune('n') {
						goto l590
					}
					position++
					if !_rules[ruleApostrophe]() {
						goto l590
					}
					if buffer[position] != rune('t') {
						goto l590
					}
					position++
					goto l570
				l590:
					position, tokenIndex = position570, tokenIndex570
					if buffer[position] != rune('z') {
						goto l591
					}
					position++
					if buffer[position] != rune('u') {
						goto l591
					}
					position++
					if buffer[position] != rune('r') {
						goto l591
					}
					position++
					goto l570
				l591:
					position, tokenIndex = position570, tokenIndex570
					if buffer[position] != rune('v') {
						goto l592
					}
					position++
					if buffer[position] != rune('o') {
						goto l592
					}
					position++
					if buffer[position] != rune('n') {
						goto l592
					}
					position++
					{
						position593, tokenIndex593 := position, tokenIndex
						if !_rules[rule_]() {
							goto l593
						}
						{
							position595, tokenIndex595 := position, tokenIndex
							if buffer[position] != rune('d') {
								goto l596
							}
							position++
							if buffer[position] != rune('.') {
								goto l596
							}
							position++
							goto l595
						l596:
							position, tokenIndex = position595, tokenIndex595
							if buffer[position] != rune('d') {
								goto l593
							}
							position++
							if buffer[position] != rune('e') {
								goto l593
							}
							position++
							if buffer[position] != rune('m') {
								goto l593
							}
							position++
						}
					l595:
						goto l594
					l593:
						position, tokenIndex = position593, tokenIndex593
					}
				l594:
					goto l570
				l592:
					position, tokenIndex = position570, tokenIndex570
					if buffer[position] != rune('v') {
						goto l568
					}
					position++
					{
						position597, tokenIndex597 := position, tokenIndex
						if !_rules[rule_]() {
							goto l597
						}
						if buffer[position] != rune('d') {
							goto l597
						}
						position++
						goto l598
					l597:
						position, tokenIndex = position597, tokenIndex597
					}
				l598:
				}
			l570:
				{
					position599, tokenIndex599 := position, tokenIndex
					if !_rules[rule_]() {
						goto l568
					}
					position, tokenIndex = position599, tokenIndex599
				}
				add(ruleAuthorPrefix1, position569)
			}
			return true
		l568:
			position, tokenIndex = position568, tokenIndex568
			return false
		},
		/* 80 AuthorUpperChar <- <(UpperASCII / MiscodedChar / ('À' / 'Á' / 'Â' / 'Ã' / 'Ä' / 'Å' / 'Æ' / 'Ç' / 'È' / 'É' / 'Ê' / 'Ë' / 'Ì' / 'Í' / 'Î' / 'Ï' / 'Ð' / 'Ñ' / 'Ò' / 'Ó' / 'Ô' / 'Õ' / 'Ö' / 'Ø' / 'Ù' / 'Ú' / 'Û' / 'Ü' / 'Ý' / 'Ć' / 'Č' / 'Ď' / 'İ' / 'Ķ' / 'Ĺ' / 'ĺ' / 'Ľ' / 'ľ' / 'Ł' / 'ł' / 'Ņ' / 'Ō' / 'Ő' / 'Œ' / 'Ř' / 'Ś' / 'Ŝ' / 'Ş' / 'Š' / 'Ÿ' / 'Ź' / 'Ż' / 'Ž' / 'ƒ' / 'Ǿ' / 'Ș' / 'Ț'))> */
		func() bool {
			position600, tokenIndex600 := position, tokenIndex
			{
				position601 := position
				{
					position602, tokenIndex602 := position, tokenIndex
					if !_rules[ruleUpperASCII]() {
						goto l603
					}
					goto l602
				l603:
					position, tokenIndex = position602, tokenIndex602
					if !_rules[ruleMiscodedChar]() {
						goto l604
					}
					goto l602
				l604:
					position, tokenIndex = position602, tokenIndex602
					{
						position605, tokenIndex605 := position, tokenIndex
						if buffer[position] != rune('À') {
							goto l606
						}
						position++
						goto l605
					l606:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Á') {
							goto l607
						}
						position++
						goto l605
					l607:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Â') {
							goto l608
						}
						position++
						goto l605
					l608:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ã') {
							goto l609
						}
						position++
						goto l605
					l609:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ä') {
							goto l610
						}
						position++
						goto l605
					l610:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Å') {
							goto l611
						}
						position++
						goto l605
					l611:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Æ') {
							goto l612
						}
						position++
						goto l605
					l612:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ç') {
							goto l613
						}
						position++
						goto l605
					l613:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('È') {
							goto l614
						}
						position++
						goto l605
					l614:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('É') {
							goto l615
						}
						position++
						goto l605
					l615:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ê') {
							goto l616
						}
						position++
						goto l605
					l616:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ë') {
							goto l617
						}
						position++
						goto l605
					l617:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ì') {
							goto l618
						}
						position++
						goto l605
					l618:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Í') {
							goto l619
						}
						position++
						goto l605
					l619:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Î') {
							goto l620
						}
						position++
						goto l605
					l620:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ï') {
							goto l621
						}
						position++
						goto l605
					l621:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ð') {
							goto l622
						}
						position++
						goto l605
					l622:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ñ') {
							goto l623
						}
						position++
						goto l605
					l623:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ò') {
							goto l624
						}
						position++
						goto l605
					l624:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ó') {
							goto l625
						}
						position++
						goto l605
					l625:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ô') {
							goto l626
						}
						position++
						goto l605
					l626:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Õ') {
							goto l627
						}
						position++
						goto l605
					l627:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ö') {
							goto l628
						}
						position++
						goto l605
					l628:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ø') {
							goto l629
						}
						position++
						goto l605
					l629:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ù') {
							goto l630
						}
						position++
						goto l605
					l630:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ú') {
							goto l631
						}
						position++
						goto l605
					l631:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Û') {
							goto l632
						}
						position++
						goto l605
					l632:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ü') {
							goto l633
						}
						position++
						goto l605
					l633:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ý') {
							goto l634
						}
						position++
						goto l605
					l634:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ć') {
							goto l635
						}
						position++
						goto l605
					l635:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Č') {
							goto l636
						}
						position++
						goto l605
					l636:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ď') {
							goto l637
						}
						position++
						goto l605
					l637:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('İ') {
							goto l638
						}
						position++
						goto l605
					l638:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ķ') {
							goto l639
						}
						position++
						goto l605
					l639:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ĺ') {
							goto l640
						}
						position++
						goto l605
					l640:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('ĺ') {
							goto l641
						}
						position++
						goto l605
					l641:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ľ') {
							goto l642
						}
						position++
						goto l605
					l642:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('ľ') {
							goto l643
						}
						position++
						goto l605
					l643:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ł') {
							goto l644
						}
						position++
						goto l605
					l644:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('ł') {
							goto l645
						}
						position++
						goto l605
					l645:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ņ') {
							goto l646
						}
						position++
						goto l605
					l646:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ō') {
							goto l647
						}
						position++
						goto l605
					l647:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ő') {
							goto l648
						}
						position++
						goto l605
					l648:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Œ') {
							goto l649
						}
						position++
						goto l605
					l649:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ř') {
							goto l650
						}
						position++
						goto l605
					l650:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ś') {
							goto l651
						}
						position++
						goto l605
					l651:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ŝ') {
							goto l652
						}
						position++
						goto l605
					l652:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ş') {
							goto l653
						}
						position++
						goto l605
					l653:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Š') {
							goto l654
						}
						position++
						goto l605
					l654:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ÿ') {
							goto l655
						}
						position++
						goto l605
					l655:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ź') {
							goto l656
						}
						position++
						goto l605
					l656:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ż') {
							goto l657
						}
						position++
						goto l605
					l657:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ž') {
							goto l658
						}
						position++
						goto l605
					l658:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('ƒ') {
							goto l659
						}
						position++
						goto l605
					l659:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ǿ') {
							goto l660
						}
						position++
						goto l605
					l660:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ș') {
							goto l661
						}
						position++
						goto l605
					l661:
						position, tokenIndex = position605, tokenIndex605
						if buffer[position] != rune('Ț') {
							goto l600
						}
						position++
					}
				l605:
				}
			l602:
				add(ruleAuthorUpperChar, position601)
			}
			return true
		l600:
			position, tokenIndex = position600, tokenIndex600
			return false
		},
		/* 81 AuthorLowerChar <- <(LowerASCII / MiscodedChar / ('à' / 'á' / 'â' / 'ã' / 'ä' / 'å' / 'æ' / 'ç' / 'è' / 'é' / 'ê' / 'ë' / 'ì' / 'í' / 'î' / 'ï' / 'ð' / 'ñ' / 'ò' / 'ó' / 'ó' / 'ô' / 'õ' / 'ö' / 'ø' / 'ù' / 'ú' / 'û' / 'ü' / 'ý' / 'ÿ' / 'ā' / 'ă' / 'ą' / 'ć' / 'ĉ' / 'č' / 'ď' / 'đ' / '\'' / 'ē' / 'ĕ' / 'ė' / 'ę' / 'ě' / 'ğ' / 'ī' / 'ĭ' / 'İ' / 'ı' / 'ĺ' / 'ľ' / 'ł' / 'ń' / 'ņ' / 'ň' / 'ŏ' / 'ő' / 'œ' / 'ŕ' / 'ř' / 'ś' / 'ş' / 'š' / 'ţ' / 'ť' / 'ũ' / 'ū' / 'ŭ' / 'ů' / 'ű' / 'ź' / 'ż' / 'ž' / 'ſ' / 'ǎ' / 'ǔ' / 'ǧ' / 'ș' / 'ț' / 'ȳ' / 'ß'))> */
		func() bool {
			position662, tokenIndex662 := position, tokenIndex
			{
				position663 := position
				{
					position664, tokenIndex664 := position, tokenIndex
					if !_rules[ruleLowerASCII]() {
						goto l665
					}
					goto l664
				l665:
					position, tokenIndex = position664, tokenIndex664
					if !_rules[ruleMiscodedChar]() {
						goto l666
					}
					goto l664
				l666:
					position, tokenIndex = position664, tokenIndex664
					{
						position667, tokenIndex667 := position, tokenIndex
						if buffer[position] != rune('à') {
							goto l668
						}
						position++
						goto l667
					l668:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('á') {
							goto l669
						}
						position++
						goto l667
					l669:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('â') {
							goto l670
						}
						position++
						goto l667
					l670:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ã') {
							goto l671
						}
						position++
						goto l667
					l671:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ä') {
							goto l672
						}
						position++
						goto l667
					l672:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('å') {
							goto l673
						}
						position++
						goto l667
					l673:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('æ') {
							goto l674
						}
						position++
						goto l667
					l674:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ç') {
							goto l675
						}
						position++
						goto l667
					l675:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('è') {
							goto l676
						}
						position++
						goto l667
					l676:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('é') {
							goto l677
						}
						position++
						goto l667
					l677:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ê') {
							goto l678
						}
						position++
						goto l667
					l678:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ë') {
							goto l679
						}
						position++
						goto l667
					l679:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ì') {
							goto l680
						}
						position++
						goto l667
					l680:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('í') {
							goto l681
						}
						position++
						goto l667
					l681:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('î') {
							goto l682
						}
						position++
						goto l667
					l682:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ï') {
							goto l683
						}
						position++
						goto l667
					l683:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ð') {
							goto l684
						}
						position++
						goto l667
					l684:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ñ') {
							goto l685
						}
						position++
						goto l667
					l685:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ò') {
							goto l686
						}
						position++
						goto l667
					l686:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ó') {
							goto l687
						}
						position++
						goto l667
					l687:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ó') {
							goto l688
						}
						position++
						goto l667
					l688:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ô') {
							goto l689
						}
						position++
						goto l667
					l689:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('õ') {
							goto l690
						}
						position++
						goto l667
					l690:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ö') {
							goto l691
						}
						position++
						goto l667
					l691:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ø') {
							goto l692
						}
						position++
						goto l667
					l692:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ù') {
							goto l693
						}
						position++
						goto l667
					l693:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ú') {
							goto l694
						}
						position++
						goto l667
					l694:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('û') {
							goto l695
						}
						position++
						goto l667
					l695:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ü') {
							goto l696
						}
						position++
						goto l667
					l696:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ý') {
							goto l697
						}
						position++
						goto l667
					l697:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ÿ') {
							goto l698
						}
						position++
						goto l667
					l698:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ā') {
							goto l699
						}
						position++
						goto l667
					l699:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ă') {
							goto l700
						}
						position++
						goto l667
					l700:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ą') {
							goto l701
						}
						position++
						goto l667
					l701:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ć') {
							goto l702
						}
						position++
						goto l667
					l702:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ĉ') {
							goto l703
						}
						position++
						goto l667
					l703:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('č') {
							goto l704
						}
						position++
						goto l667
					l704:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ď') {
							goto l705
						}
						position++
						goto l667
					l705:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('đ') {
							goto l706
						}
						position++
						goto l667
					l706:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('\'') {
							goto l707
						}
						position++
						goto l667
					l707:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ē') {
							goto l708
						}
						position++
						goto l667
					l708:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ĕ') {
							goto l709
						}
						position++
						goto l667
					l709:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ė') {
							goto l710
						}
						position++
						goto l667
					l710:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ę') {
							goto l711
						}
						position++
						goto l667
					l711:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ě') {
							goto l712
						}
						position++
						goto l667
					l712:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ğ') {
							goto l713
						}
						position++
						goto l667
					l713:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ī') {
							goto l714
						}
						position++
						goto l667
					l714:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ĭ') {
							goto l715
						}
						position++
						goto l667
					l715:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('İ') {
							goto l716
						}
						position++
						goto l667
					l716:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ı') {
							goto l717
						}
						position++
						goto l667
					l717:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ĺ') {
							goto l718
						}
						position++
						goto l667
					l718:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ľ') {
							goto l719
						}
						position++
						goto l667
					l719:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ł') {
							goto l720
						}
						position++
						goto l667
					l720:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ń') {
							goto l721
						}
						position++
						goto l667
					l721:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ņ') {
							goto l722
						}
						position++
						goto l667
					l722:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ň') {
							goto l723
						}
						position++
						goto l667
					l723:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ŏ') {
							goto l724
						}
						position++
						goto l667
					l724:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ő') {
							goto l725
						}
						position++
						goto l667
					l725:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('œ') {
							goto l726
						}
						position++
						goto l667
					l726:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ŕ') {
							goto l727
						}
						position++
						goto l667
					l727:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ř') {
							goto l728
						}
						position++
						goto l667
					l728:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ś') {
							goto l729
						}
						position++
						goto l667
					l729:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ş') {
							goto l730
						}
						position++
						goto l667
					l730:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('š') {
							goto l731
						}
						position++
						goto l667
					l731:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ţ') {
							goto l732
						}
						position++
						goto l667
					l732:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ť') {
							goto l733
						}
						position++
						goto l667
					l733:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ũ') {
							goto l734
						}
						position++
						goto l667
					l734:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ū') {
							goto l735
						}
						position++
						goto l667
					l735:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ŭ') {
							goto l736
						}
						position++
						goto l667
					l736:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ů') {
							goto l737
						}
						position++
						goto l667
					l737:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ű') {
							goto l738
						}
						position++
						goto l667
					l738:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ź') {
							goto l739
						}
						position++
						goto l667
					l739:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ż') {
							goto l740
						}
						position++
						goto l667
					l740:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ž') {
							goto l741
						}
						position++
						goto l667
					l741:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ſ') {
							goto l742
						}
						position++
						goto l667
					l742:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ǎ') {
							goto l743
						}
						position++
						goto l667
					l743:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ǔ') {
							goto l744
						}
						position++
						goto l667
					l744:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ǧ') {
							goto l745
						}
						position++
						goto l667
					l745:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ș') {
							goto l746
						}
						position++
						goto l667
					l746:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ț') {
							goto l747
						}
						position++
						goto l667
					l747:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ȳ') {
							goto l748
						}
						position++
						goto l667
					l748:
						position, tokenIndex = position667, tokenIndex667
						if buffer[position] != rune('ß') {
							goto l662
						}
						position++
					}
				l667:
				}
			l664:
				add(ruleAuthorLowerChar, position663)
			}
			return true
		l662:
			position, tokenIndex = position662, tokenIndex662
			return false
		},
		/* 82 Year <- <(YearRange / YearApprox / YearWithParens / YearWithPage / YearWithDot / YearWithChar / YearNum)> */
		func() bool {
			position749, tokenIndex749 := position, tokenIndex
			{
				position750 := position
				{
					position751, tokenIndex751 := position, tokenIndex
					if !_rules[ruleYearRange]() {
						goto l752
					}
					goto l751
				l752:
					position, tokenIndex = position751, tokenIndex751
					if !_rules[ruleYearApprox]() {
						goto l753
					}
					goto l751
				l753:
					position, tokenIndex = position751, tokenIndex751
					if !_rules[ruleYearWithParens]() {
						goto l754
					}
					goto l751
				l754:
					position, tokenIndex = position751, tokenIndex751
					if !_rules[ruleYearWithPage]() {
						goto l755
					}
					goto l751
				l755:
					position, tokenIndex = position751, tokenIndex751
					if !_rules[ruleYearWithDot]() {
						goto l756
					}
					goto l751
				l756:
					position, tokenIndex = position751, tokenIndex751
					if !_rules[ruleYearWithChar]() {
						goto l757
					}
					goto l751
				l757:
					position, tokenIndex = position751, tokenIndex751
					if !_rules[ruleYearNum]() {
						goto l749
					}
				}
			l751:
				add(ruleYear, position750)
			}
			return true
		l749:
			position, tokenIndex = position749, tokenIndex749
			return false
		},
		/* 83 YearRange <- <(YearNum Dash (Nums+ ('a' / 'b' / 'c' / 'd' / 'e' / 'f' / 'g' / 'h' / 'i' / 'j' / 'k' / 'l' / 'm' / 'n' / 'o' / 'p' / 'q' / 'r' / 's' / 't' / 'u' / 'v' / 'w' / 'x' / 'y' / 'z' / '?')*))> */
		func() bool {
			position758, tokenIndex758 := position, tokenIndex
			{
				position759 := position
				if !_rules[ruleYearNum]() {
					goto l758
				}
				if !_rules[ruleDash]() {
					goto l758
				}
				if !_rules[ruleNums]() {
					goto l758
				}
			l760:
				{
					position761, tokenIndex761 := position, tokenIndex
					if !_rules[ruleNums]() {
						goto l761
					}
					goto l760
				l761:
					position, tokenIndex = position761, tokenIndex761
				}
			l762:
				{
					position763, tokenIndex763 := position, tokenIndex
					{
						position764, tokenIndex764 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l765
						}
						position++
						goto l764
					l765:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('b') {
							goto l766
						}
						position++
						goto l764
					l766:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('c') {
							goto l767
						}
						position++
						goto l764
					l767:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('d') {
							goto l768
						}
						position++
						goto l764
					l768:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('e') {
							goto l769
						}
						position++
						goto l764
					l769:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('f') {
							goto l770
						}
						position++
						goto l764
					l770:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('g') {
							goto l771
						}
						position++
						goto l764
					l771:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('h') {
							goto l772
						}
						position++
						goto l764
					l772:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('i') {
							goto l773
						}
						position++
						goto l764
					l773:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('j') {
							goto l774
						}
						position++
						goto l764
					l774:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('k') {
							goto l775
						}
						position++
						goto l764
					l775:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('l') {
							goto l776
						}
						position++
						goto l764
					l776:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('m') {
							goto l777
						}
						position++
						goto l764
					l777:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('n') {
							goto l778
						}
						position++
						goto l764
					l778:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('o') {
							goto l779
						}
						position++
						goto l764
					l779:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('p') {
							goto l780
						}
						position++
						goto l764
					l780:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('q') {
							goto l781
						}
						position++
						goto l764
					l781:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('r') {
							goto l782
						}
						position++
						goto l764
					l782:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('s') {
							goto l783
						}
						position++
						goto l764
					l783:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('t') {
							goto l784
						}
						position++
						goto l764
					l784:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('u') {
							goto l785
						}
						position++
						goto l764
					l785:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('v') {
							goto l786
						}
						position++
						goto l764
					l786:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('w') {
							goto l787
						}
						position++
						goto l764
					l787:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('x') {
							goto l788
						}
						position++
						goto l764
					l788:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('y') {
							goto l789
						}
						position++
						goto l764
					l789:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('z') {
							goto l790
						}
						position++
						goto l764
					l790:
						position, tokenIndex = position764, tokenIndex764
						if buffer[position] != rune('?') {
							goto l763
						}
						position++
					}
				l764:
					goto l762
				l763:
					position, tokenIndex = position763, tokenIndex763
				}
				add(ruleYearRange, position759)
			}
			return true
		l758:
			position, tokenIndex = position758, tokenIndex758
			return false
		},
		/* 84 YearWithDot <- <(YearNum '.')> */
		func() bool {
			position791, tokenIndex791 := position, tokenIndex
			{
				position792 := position
				if !_rules[ruleYearNum]() {
					goto l791
				}
				if buffer[position] != rune('.') {
					goto l791
				}
				position++
				add(ruleYearWithDot, position792)
			}
			return true
		l791:
			position, tokenIndex = position791, tokenIndex791
			return false
		},
		/* 85 YearApprox <- <('[' _? YearNum _? ']')> */
		func() bool {
			position793, tokenIndex793 := position, tokenIndex
			{
				position794 := position
				if buffer[position] != rune('[') {
					goto l793
				}
				position++
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
				if !_rules[ruleYearNum]() {
					goto l793
				}
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
				if buffer[position] != rune(']') {
					goto l793
				}
				position++
				add(ruleYearApprox, position794)
			}
			return true
		l793:
			position, tokenIndex = position793, tokenIndex793
			return false
		},
		/* 86 YearWithPage <- <((YearWithChar / YearNum) _? ':' _? Nums+)> */
		func() bool {
			position799, tokenIndex799 := position, tokenIndex
			{
				position800 := position
				{
					position801, tokenIndex801 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l802
					}
					goto l801
				l802:
					position, tokenIndex = position801, tokenIndex801
					if !_rules[ruleYearNum]() {
						goto l799
					}
				}
			l801:
				{
					position803, tokenIndex803 := position, tokenIndex
					if !_rules[rule_]() {
						goto l803
					}
					goto l804
				l803:
					position, tokenIndex = position803, tokenIndex803
				}
			l804:
				if buffer[position] != rune(':') {
					goto l799
				}
				position++
				{
					position805, tokenIndex805 := position, tokenIndex
					if !_rules[rule_]() {
						goto l805
					}
					goto l806
				l805:
					position, tokenIndex = position805, tokenIndex805
				}
			l806:
				if !_rules[ruleNums]() {
					goto l799
				}
			l807:
				{
					position808, tokenIndex808 := position, tokenIndex
					if !_rules[ruleNums]() {
						goto l808
					}
					goto l807
				l808:
					position, tokenIndex = position808, tokenIndex808
				}
				add(ruleYearWithPage, position800)
			}
			return true
		l799:
			position, tokenIndex = position799, tokenIndex799
			return false
		},
		/* 87 YearWithParens <- <('(' (YearWithChar / YearNum) ')')> */
		func() bool {
			position809, tokenIndex809 := position, tokenIndex
			{
				position810 := position
				if buffer[position] != rune('(') {
					goto l809
				}
				position++
				{
					position811, tokenIndex811 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l812
					}
					goto l811
				l812:
					position, tokenIndex = position811, tokenIndex811
					if !_rules[ruleYearNum]() {
						goto l809
					}
				}
			l811:
				if buffer[position] != rune(')') {
					goto l809
				}
				position++
				add(ruleYearWithParens, position810)
			}
			return true
		l809:
			position, tokenIndex = position809, tokenIndex809
			return false
		},
		/* 88 YearWithChar <- <(YearNum LowerASCII Action0)> */
		func() bool {
			position813, tokenIndex813 := position, tokenIndex
			{
				position814 := position
				if !_rules[ruleYearNum]() {
					goto l813
				}
				if !_rules[ruleLowerASCII]() {
					goto l813
				}
				if !_rules[ruleAction0]() {
					goto l813
				}
				add(ruleYearWithChar, position814)
			}
			return true
		l813:
			position, tokenIndex = position813, tokenIndex813
			return false
		},
		/* 89 YearNum <- <(('1' / '2') ('0' / '7' / '8' / '9') Nums (Nums / '?') '?'*)> */
		func() bool {
			position815, tokenIndex815 := position, tokenIndex
			{
				position816 := position
				{
					position817, tokenIndex817 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l818
					}
					position++
					goto l817
				l818:
					position, tokenIndex = position817, tokenIndex817
					if buffer[position] != rune('2') {
						goto l815
					}
					position++
				}
			l817:
				{
					position819, tokenIndex819 := position, tokenIndex
					if buffer[position] != rune('0') {
						goto l820
					}
					position++
					goto l819
				l820:
					position, tokenIndex = position819, tokenIndex819
					if buffer[position] != rune('7') {
						goto l821
					}
					position++
					goto l819
				l821:
					position, tokenIndex = position819, tokenIndex819
					if buffer[position] != rune('8') {
						goto l822
					}
					position++
					goto l819
				l822:
					position, tokenIndex = position819, tokenIndex819
					if buffer[position] != rune('9') {
						goto l815
					}
					position++
				}
			l819:
				if !_rules[ruleNums]() {
					goto l815
				}
				{
					position823, tokenIndex823 := position, tokenIndex
					if !_rules[ruleNums]() {
						goto l824
					}
					goto l823
				l824:
					position, tokenIndex = position823, tokenIndex823
					if buffer[position] != rune('?') {
						goto l815
					}
					position++
				}
			l823:
			l825:
				{
					position826, tokenIndex826 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l826
					}
					position++
					goto l825
				l826:
					position, tokenIndex = position826, tokenIndex826
				}
				add(ruleYearNum, position816)
			}
			return true
		l815:
			position, tokenIndex = position815, tokenIndex815
			return false
		},
		/* 90 NameUpperChar <- <(UpperChar / UpperCharExtended)> */
		func() bool {
			position827, tokenIndex827 := position, tokenIndex
			{
				position828 := position
				{
					position829, tokenIndex829 := position, tokenIndex
					if !_rules[ruleUpperChar]() {
						goto l830
					}
					goto l829
				l830:
					position, tokenIndex = position829, tokenIndex829
					if !_rules[ruleUpperCharExtended]() {
						goto l827
					}
				}
			l829:
				add(ruleNameUpperChar, position828)
			}
			return true
		l827:
			position, tokenIndex = position827, tokenIndex827
			return false
		},
		/* 91 UpperCharExtended <- <('Æ' / 'Œ' / 'Ö')> */
		func() bool {
			position831, tokenIndex831 := position, tokenIndex
			{
				position832 := position
				{
					position833, tokenIndex833 := position, tokenIndex
					if buffer[position] != rune('Æ') {
						goto l834
					}
					position++
					goto l833
				l834:
					position, tokenIndex = position833, tokenIndex833
					if buffer[position] != rune('Œ') {
						goto l835
					}
					position++
					goto l833
				l835:
					position, tokenIndex = position833, tokenIndex833
					if buffer[position] != rune('Ö') {
						goto l831
					}
					position++
				}
			l833:
				add(ruleUpperCharExtended, position832)
			}
			return true
		l831:
			position, tokenIndex = position831, tokenIndex831
			return false
		},
		/* 92 UpperChar <- <UpperASCII> */
		func() bool {
			position836, tokenIndex836 := position, tokenIndex
			{
				position837 := position
				if !_rules[ruleUpperASCII]() {
					goto l836
				}
				add(ruleUpperChar, position837)
			}
			return true
		l836:
			position, tokenIndex = position836, tokenIndex836
			return false
		},
		/* 93 NameLowerChar <- <(LowerChar / LowerCharExtended / MiscodedChar)> */
		func() bool {
			position838, tokenIndex838 := position, tokenIndex
			{
				position839 := position
				{
					position840, tokenIndex840 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l841
					}
					goto l840
				l841:
					position, tokenIndex = position840, tokenIndex840
					if !_rules[ruleLowerCharExtended]() {
						goto l842
					}
					goto l840
				l842:
					position, tokenIndex = position840, tokenIndex840
					if !_rules[ruleMiscodedChar]() {
						goto l838
					}
				}
			l840:
				add(ruleNameLowerChar, position839)
			}
			return true
		l838:
			position, tokenIndex = position838, tokenIndex838
			return false
		},
		/* 94 MiscodedChar <- <'�'> */
		func() bool {
			position843, tokenIndex843 := position, tokenIndex
			{
				position844 := position
				if buffer[position] != rune('�') {
					goto l843
				}
				position++
				add(ruleMiscodedChar, position844)
			}
			return true
		l843:
			position, tokenIndex = position843, tokenIndex843
			return false
		},
		/* 95 LowerCharExtended <- <('æ' / 'œ' / 'à' / 'â' / 'å' / 'ã' / 'ä' / 'á' / 'ç' / 'č' / 'é' / 'è' / 'ë' / 'í' / 'ì' / 'ï' / 'ň' / 'ñ' / 'ñ' / 'ó' / 'ò' / 'ô' / 'ø' / 'õ' / 'ö' / 'ú' / 'ù' / 'ü' / 'ŕ' / 'ř' / 'ŗ' / 'ſ' / 'š' / 'š' / 'ş' / 'ž')> */
		func() bool {
			position845, tokenIndex845 := position, tokenIndex
			{
				position846 := position
				{
					position847, tokenIndex847 := position, tokenIndex
					if buffer[position] != rune('æ') {
						goto l848
					}
					position++
					goto l847
				l848:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('œ') {
						goto l849
					}
					position++
					goto l847
				l849:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('à') {
						goto l850
					}
					position++
					goto l847
				l850:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('â') {
						goto l851
					}
					position++
					goto l847
				l851:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('å') {
						goto l852
					}
					position++
					goto l847
				l852:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('ã') {
						goto l853
					}
					position++
					goto l847
				l853:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('ä') {
						goto l854
					}
					position++
					goto l847
				l854:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('á') {
						goto l855
					}
					position++
					goto l847
				l855:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('ç') {
						goto l856
					}
					position++
					goto l847
				l856:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('č') {
						goto l857
					}
					position++
					goto l847
				l857:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('é') {
						goto l858
					}
					position++
					goto l847
				l858:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('è') {
						goto l859
					}
					position++
					goto l847
				l859:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('ë') {
						goto l860
					}
					position++
					goto l847
				l860:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('í') {
						goto l861
					}
					position++
					goto l847
				l861:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('ì') {
						goto l862
					}
					position++
					goto l847
				l862:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('ï') {
						goto l863
					}
					position++
					goto l847
				l863:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('ň') {
						goto l864
					}
					position++
					goto l847
				l864:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('ñ') {
						goto l865
					}
					position++
					goto l847
				l865:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('ñ') {
						goto l866
					}
					position++
					goto l847
				l866:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('ó') {
						goto l867
					}
					position++
					goto l847
				l867:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('ò') {
						goto l868
					}
					position++
					goto l847
				l868:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('ô') {
						goto l869
					}
					position++
					goto l847
				l869:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('ø') {
						goto l870
					}
					position++
					goto l847
				l870:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('õ') {
						goto l871
					}
					position++
					goto l847
				l871:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('ö') {
						goto l872
					}
					position++
					goto l847
				l872:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('ú') {
						goto l873
					}
					position++
					goto l847
				l873:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('ù') {
						goto l874
					}
					position++
					goto l847
				l874:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('ü') {
						goto l875
					}
					position++
					goto l847
				l875:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('ŕ') {
						goto l876
					}
					position++
					goto l847
				l876:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('ř') {
						goto l877
					}
					position++
					goto l847
				l877:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('ŗ') {
						goto l878
					}
					position++
					goto l847
				l878:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('ſ') {
						goto l879
					}
					position++
					goto l847
				l879:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('š') {
						goto l880
					}
					position++
					goto l847
				l880:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('š') {
						goto l881
					}
					position++
					goto l847
				l881:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('ş') {
						goto l882
					}
					position++
					goto l847
				l882:
					position, tokenIndex = position847, tokenIndex847
					if buffer[position] != rune('ž') {
						goto l845
					}
					position++
				}
			l847:
				add(ruleLowerCharExtended, position846)
			}
			return true
		l845:
			position, tokenIndex = position845, tokenIndex845
			return false
		},
		/* 96 LowerChar <- <LowerASCII> */
		func() bool {
			position883, tokenIndex883 := position, tokenIndex
			{
				position884 := position
				if !_rules[ruleLowerASCII]() {
					goto l883
				}
				add(ruleLowerChar, position884)
			}
			return true
		l883:
			position, tokenIndex = position883, tokenIndex883
			return false
		},
		/* 97 SpaceCharEOI <- <(_ / !.)> */
		func() bool {
			position885, tokenIndex885 := position, tokenIndex
			{
				position886 := position
				{
					position887, tokenIndex887 := position, tokenIndex
					if !_rules[rule_]() {
						goto l888
					}
					goto l887
				l888:
					position, tokenIndex = position887, tokenIndex887
					{
						position889, tokenIndex889 := position, tokenIndex
						if !matchDot() {
							goto l889
						}
						goto l885
					l889:
						position, tokenIndex = position889, tokenIndex889
					}
				}
			l887:
				add(ruleSpaceCharEOI, position886)
			}
			return true
		l885:
			position, tokenIndex = position885, tokenIndex885
			return false
		},
		/* 98 Nums <- <[0-9]> */
		func() bool {
			position890, tokenIndex890 := position, tokenIndex
			{
				position891 := position
				if c := buffer[position]; c < rune('0') || c > rune('9') {
					goto l890
				}
				position++
				add(ruleNums, position891)
			}
			return true
		l890:
			position, tokenIndex = position890, tokenIndex890
			return false
		},
		/* 99 LowerASCII <- <[a-z]> */
		func() bool {
			position892, tokenIndex892 := position, tokenIndex
			{
				position893 := position
				if c := buffer[position]; c < rune('a') || c > rune('z') {
					goto l892
				}
				position++
				add(ruleLowerASCII, position893)
			}
			return true
		l892:
			position, tokenIndex = position892, tokenIndex892
			return false
		},
		/* 100 UpperASCII <- <[A-Z]> */
		func() bool {
			position894, tokenIndex894 := position, tokenIndex
			{
				position895 := position
				if c := buffer[position]; c < rune('A') || c > rune('Z') {
					goto l894
				}
				position++
				add(ruleUpperASCII, position895)
			}
			return true
		l894:
			position, tokenIndex = position894, tokenIndex894
			return false
		},
		/* 101 Apostrophe <- <(ApostrOther / ApostrASCII)> */
		func() bool {
			position896, tokenIndex896 := position, tokenIndex
			{
				position897 := position
				{
					position898, tokenIndex898 := position, tokenIndex
					if !_rules[ruleApostrOther]() {
						goto l899
					}
					goto l898
				l899:
					position, tokenIndex = position898, tokenIndex898
					if !_rules[ruleApostrASCII]() {
						goto l896
					}
				}
			l898:
				add(ruleApostrophe, position897)
			}
			return true
		l896:
			position, tokenIndex = position896, tokenIndex896
			return false
		},
		/* 102 ApostrASCII <- <'\''> */
		func() bool {
			position900, tokenIndex900 := position, tokenIndex
			{
				position901 := position
				if buffer[position] != rune('\'') {
					goto l900
				}
				position++
				add(ruleApostrASCII, position901)
			}
			return true
		l900:
			position, tokenIndex = position900, tokenIndex900
			return false
		},
		/* 103 ApostrOther <- <('‘' / '’')> */
		func() bool {
			position902, tokenIndex902 := position, tokenIndex
			{
				position903 := position
				{
					position904, tokenIndex904 := position, tokenIndex
					if buffer[position] != rune('‘') {
						goto l905
					}
					position++
					goto l904
				l905:
					position, tokenIndex = position904, tokenIndex904
					if buffer[position] != rune('’') {
						goto l902
					}
					position++
				}
			l904:
				add(ruleApostrOther, position903)
			}
			return true
		l902:
			position, tokenIndex = position902, tokenIndex902
			return false
		},
		/* 104 Dash <- <'-'> */
		func() bool {
			position906, tokenIndex906 := position, tokenIndex
			{
				position907 := position
				if buffer[position] != rune('-') {
					goto l906
				}
				position++
				add(ruleDash, position907)
			}
			return true
		l906:
			position, tokenIndex = position906, tokenIndex906
			return false
		},
		/* 105 _ <- <(MultipleSpace / SingleSpace)> */
		func() bool {
			position908, tokenIndex908 := position, tokenIndex
			{
				position909 := position
				{
					position910, tokenIndex910 := position, tokenIndex
					if !_rules[ruleMultipleSpace]() {
						goto l911
					}
					goto l910
				l911:
					position, tokenIndex = position910, tokenIndex910
					if !_rules[ruleSingleSpace]() {
						goto l908
					}
				}
			l910:
				add(rule_, position909)
			}
			return true
		l908:
			position, tokenIndex = position908, tokenIndex908
			return false
		},
		/* 106 MultipleSpace <- <(SingleSpace SingleSpace+)> */
		func() bool {
			position912, tokenIndex912 := position, tokenIndex
			{
				position913 := position
				if !_rules[ruleSingleSpace]() {
					goto l912
				}
				if !_rules[ruleSingleSpace]() {
					goto l912
				}
			l914:
				{
					position915, tokenIndex915 := position, tokenIndex
					if !_rules[ruleSingleSpace]() {
						goto l915
					}
					goto l914
				l915:
					position, tokenIndex = position915, tokenIndex915
				}
				add(ruleMultipleSpace, position913)
			}
			return true
		l912:
			position, tokenIndex = position912, tokenIndex912
			return false
		},
		/* 107 SingleSpace <- <(' ' / OtherSpace)> */
		func() bool {
			position916, tokenIndex916 := position, tokenIndex
			{
				position917 := position
				{
					position918, tokenIndex918 := position, tokenIndex
					if buffer[position] != rune(' ') {
						goto l919
					}
					position++
					goto l918
				l919:
					position, tokenIndex = position918, tokenIndex918
					if !_rules[ruleOtherSpace]() {
						goto l916
					}
				}
			l918:
				add(ruleSingleSpace, position917)
			}
			return true
		l916:
			position, tokenIndex = position916, tokenIndex916
			return false
		},
		/* 108 OtherSpace <- <('\u3000' / '\u00a0' / '\t' / '\r' / '\n' / '\f' / '\v')> */
		func() bool {
			position920, tokenIndex920 := position, tokenIndex
			{
				position921 := position
				{
					position922, tokenIndex922 := position, tokenIndex
					if buffer[position] != rune('\u3000') {
						goto l923
					}
					position++
					goto l922
				l923:
					position, tokenIndex = position922, tokenIndex922
					if buffer[position] != rune('\u00a0') {
						goto l924
					}
					position++
					goto l922
				l924:
					position, tokenIndex = position922, tokenIndex922
					if buffer[position] != rune('\t') {
						goto l925
					}
					position++
					goto l922
				l925:
					position, tokenIndex = position922, tokenIndex922
					if buffer[position] != rune('\r') {
						goto l926
					}
					position++
					goto l922
				l926:
					position, tokenIndex = position922, tokenIndex922
					if buffer[position] != rune('\n') {
						goto l927
					}
					position++
					goto l922
				l927:
					position, tokenIndex = position922, tokenIndex922
					if buffer[position] != rune('\f') {
						goto l928
					}
					position++
					goto l922
				l928:
					position, tokenIndex = position922, tokenIndex922
					if buffer[position] != rune('\v') {
						goto l920
					}
					position++
				}
			l922:
				add(ruleOtherSpace, position921)
			}
			return true
		l920:
			position, tokenIndex = position920, tokenIndex920
			return false
		},
		/* 110 Action0 <- <{ p.AddWarn(YearCharWarn) }> */
		func() bool {
			{
				add(ruleAction0, position)
			}
			return true
		},
	}
	p.rules = _rules
}
