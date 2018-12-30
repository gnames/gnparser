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
	ruleSciName1
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
	"SciName1",
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
		/* 0 SciName <- <(_? SciName1 Tail !.)> */
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
		/* 2 SciName1 <- <(HybridFormula / NamedHybrid / Name)> */
		func() bool {
			position9, tokenIndex9 := position, tokenIndex
			{
				position10 := position
				{
					position11, tokenIndex11 := position, tokenIndex
					if !_rules[ruleHybridFormula]() {
						goto l12
					}
					goto l11
				l12:
					position, tokenIndex = position11, tokenIndex11
					if !_rules[ruleNamedHybrid]() {
						goto l13
					}
					goto l11
				l13:
					position, tokenIndex = position11, tokenIndex11
					if !_rules[ruleName]() {
						goto l9
					}
				}
			l11:
				add(ruleSciName1, position10)
			}
			return true
		l9:
			position, tokenIndex = position9, tokenIndex9
			return false
		},
		/* 3 HybridFormula <- <(Name (_ (HybridFormula1 / HybridFormula2))+)> */
		func() bool {
			position14, tokenIndex14 := position, tokenIndex
			{
				position15 := position
				if !_rules[ruleName]() {
					goto l14
				}
				if !_rules[rule_]() {
					goto l14
				}
				{
					position18, tokenIndex18 := position, tokenIndex
					if !_rules[ruleHybridFormula1]() {
						goto l19
					}
					goto l18
				l19:
					position, tokenIndex = position18, tokenIndex18
					if !_rules[ruleHybridFormula2]() {
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
						if !_rules[ruleHybridFormula1]() {
							goto l21
						}
						goto l20
					l21:
						position, tokenIndex = position20, tokenIndex20
						if !_rules[ruleHybridFormula2]() {
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
		/* 4 HybridFormula1 <- <(HybridChar _? SpeciesEpithet (_ InfraspGroup)?)> */
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
					goto l25
				l24:
					position, tokenIndex = position24, tokenIndex24
				}
			l25:
				if !_rules[ruleSpeciesEpithet]() {
					goto l22
				}
				{
					position26, tokenIndex26 := position, tokenIndex
					if !_rules[rule_]() {
						goto l26
					}
					if !_rules[ruleInfraspGroup]() {
						goto l26
					}
					goto l27
				l26:
					position, tokenIndex = position26, tokenIndex26
				}
			l27:
				add(ruleHybridFormula1, position23)
			}
			return true
		l22:
			position, tokenIndex = position22, tokenIndex22
			return false
		},
		/* 5 HybridFormula2 <- <(HybridChar (_ Name)?)> */
		func() bool {
			position28, tokenIndex28 := position, tokenIndex
			{
				position29 := position
				if !_rules[ruleHybridChar]() {
					goto l28
				}
				{
					position30, tokenIndex30 := position, tokenIndex
					if !_rules[rule_]() {
						goto l30
					}
					if !_rules[ruleName]() {
						goto l30
					}
					goto l31
				l30:
					position, tokenIndex = position30, tokenIndex30
				}
			l31:
				add(ruleHybridFormula2, position29)
			}
			return true
		l28:
			position, tokenIndex = position28, tokenIndex28
			return false
		},
		/* 6 NamedHybrid <- <(HybridChar _? Name)> */
		func() bool {
			position32, tokenIndex32 := position, tokenIndex
			{
				position33 := position
				if !_rules[ruleHybridChar]() {
					goto l32
				}
				{
					position34, tokenIndex34 := position, tokenIndex
					if !_rules[rule_]() {
						goto l34
					}
					goto l35
				l34:
					position, tokenIndex = position34, tokenIndex34
				}
			l35:
				if !_rules[ruleName]() {
					goto l32
				}
				add(ruleNamedHybrid, position33)
			}
			return true
		l32:
			position, tokenIndex = position32, tokenIndex32
			return false
		},
		/* 7 Name <- <(NameSpecies / NameUninomial)> */
		func() bool {
			position36, tokenIndex36 := position, tokenIndex
			{
				position37 := position
				{
					position38, tokenIndex38 := position, tokenIndex
					if !_rules[ruleNameSpecies]() {
						goto l39
					}
					goto l38
				l39:
					position, tokenIndex = position38, tokenIndex38
					if !_rules[ruleNameUninomial]() {
						goto l36
					}
				}
			l38:
				add(ruleName, position37)
			}
			return true
		l36:
			position, tokenIndex = position36, tokenIndex36
			return false
		},
		/* 8 NameUninomial <- <(UninomialCombo / Uninomial)> */
		func() bool {
			position40, tokenIndex40 := position, tokenIndex
			{
				position41 := position
				{
					position42, tokenIndex42 := position, tokenIndex
					if !_rules[ruleUninomialCombo]() {
						goto l43
					}
					goto l42
				l43:
					position, tokenIndex = position42, tokenIndex42
					if !_rules[ruleUninomial]() {
						goto l40
					}
				}
			l42:
				add(ruleNameUninomial, position41)
			}
			return true
		l40:
			position, tokenIndex = position40, tokenIndex40
			return false
		},
		/* 9 NameApprox <- <(GenusWord _ Approximation (_ SpeciesEpithet)?)> */
		nil,
		/* 10 NameComp <- <(GenusWord _ Comparison (_ SpeciesEpithet)?)> */
		nil,
		/* 11 NameSpecies <- <(GenusWord (_? (SubGenus / SubGenusOrSuperspecies))? _ SpeciesEpithet (_ InfraspGroup)?)> */
		func() bool {
			position46, tokenIndex46 := position, tokenIndex
			{
				position47 := position
				if !_rules[ruleGenusWord]() {
					goto l46
				}
				{
					position48, tokenIndex48 := position, tokenIndex
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
					{
						position52, tokenIndex52 := position, tokenIndex
						if !_rules[ruleSubGenus]() {
							goto l53
						}
						goto l52
					l53:
						position, tokenIndex = position52, tokenIndex52
						if !_rules[ruleSubGenusOrSuperspecies]() {
							goto l48
						}
					}
				l52:
					goto l49
				l48:
					position, tokenIndex = position48, tokenIndex48
				}
			l49:
				if !_rules[rule_]() {
					goto l46
				}
				if !_rules[ruleSpeciesEpithet]() {
					goto l46
				}
				{
					position54, tokenIndex54 := position, tokenIndex
					if !_rules[rule_]() {
						goto l54
					}
					if !_rules[ruleInfraspGroup]() {
						goto l54
					}
					goto l55
				l54:
					position, tokenIndex = position54, tokenIndex54
				}
			l55:
				add(ruleNameSpecies, position47)
			}
			return true
		l46:
			position, tokenIndex = position46, tokenIndex46
			return false
		},
		/* 12 GenusWord <- <((AbbrGenus / UninomialWord) !(_ AuthorWord))> */
		func() bool {
			position56, tokenIndex56 := position, tokenIndex
			{
				position57 := position
				{
					position58, tokenIndex58 := position, tokenIndex
					if !_rules[ruleAbbrGenus]() {
						goto l59
					}
					goto l58
				l59:
					position, tokenIndex = position58, tokenIndex58
					if !_rules[ruleUninomialWord]() {
						goto l56
					}
				}
			l58:
				{
					position60, tokenIndex60 := position, tokenIndex
					if !_rules[rule_]() {
						goto l60
					}
					if !_rules[ruleAuthorWord]() {
						goto l60
					}
					goto l56
				l60:
					position, tokenIndex = position60, tokenIndex60
				}
				add(ruleGenusWord, position57)
			}
			return true
		l56:
			position, tokenIndex = position56, tokenIndex56
			return false
		},
		/* 13 InfraspGroup <- <(InfraspEpithet (_ InfraspEpithet)? (_ InfraspEpithet)?)> */
		func() bool {
			position61, tokenIndex61 := position, tokenIndex
			{
				position62 := position
				if !_rules[ruleInfraspEpithet]() {
					goto l61
				}
				{
					position63, tokenIndex63 := position, tokenIndex
					if !_rules[rule_]() {
						goto l63
					}
					if !_rules[ruleInfraspEpithet]() {
						goto l63
					}
					goto l64
				l63:
					position, tokenIndex = position63, tokenIndex63
				}
			l64:
				{
					position65, tokenIndex65 := position, tokenIndex
					if !_rules[rule_]() {
						goto l65
					}
					if !_rules[ruleInfraspEpithet]() {
						goto l65
					}
					goto l66
				l65:
					position, tokenIndex = position65, tokenIndex65
				}
			l66:
				add(ruleInfraspGroup, position62)
			}
			return true
		l61:
			position, tokenIndex = position61, tokenIndex61
			return false
		},
		/* 14 InfraspEpithet <- <((Rank _?)? !AuthorEx Word (_ Authorship)?)> */
		func() bool {
			position67, tokenIndex67 := position, tokenIndex
			{
				position68 := position
				{
					position69, tokenIndex69 := position, tokenIndex
					if !_rules[ruleRank]() {
						goto l69
					}
					{
						position71, tokenIndex71 := position, tokenIndex
						if !_rules[rule_]() {
							goto l71
						}
						goto l72
					l71:
						position, tokenIndex = position71, tokenIndex71
					}
				l72:
					goto l70
				l69:
					position, tokenIndex = position69, tokenIndex69
				}
			l70:
				{
					position73, tokenIndex73 := position, tokenIndex
					if !_rules[ruleAuthorEx]() {
						goto l73
					}
					goto l67
				l73:
					position, tokenIndex = position73, tokenIndex73
				}
				if !_rules[ruleWord]() {
					goto l67
				}
				{
					position74, tokenIndex74 := position, tokenIndex
					if !_rules[rule_]() {
						goto l74
					}
					if !_rules[ruleAuthorship]() {
						goto l74
					}
					goto l75
				l74:
					position, tokenIndex = position74, tokenIndex74
				}
			l75:
				add(ruleInfraspEpithet, position68)
			}
			return true
		l67:
			position, tokenIndex = position67, tokenIndex67
			return false
		},
		/* 15 SpeciesEpithet <- <(!AuthorEx Word (_? Authorship)? ','? &(SpaceCharEOI / '('))> */
		func() bool {
			position76, tokenIndex76 := position, tokenIndex
			{
				position77 := position
				{
					position78, tokenIndex78 := position, tokenIndex
					if !_rules[ruleAuthorEx]() {
						goto l78
					}
					goto l76
				l78:
					position, tokenIndex = position78, tokenIndex78
				}
				if !_rules[ruleWord]() {
					goto l76
				}
				{
					position79, tokenIndex79 := position, tokenIndex
					{
						position81, tokenIndex81 := position, tokenIndex
						if !_rules[rule_]() {
							goto l81
						}
						goto l82
					l81:
						position, tokenIndex = position81, tokenIndex81
					}
				l82:
					if !_rules[ruleAuthorship]() {
						goto l79
					}
					goto l80
				l79:
					position, tokenIndex = position79, tokenIndex79
				}
			l80:
				{
					position83, tokenIndex83 := position, tokenIndex
					if buffer[position] != rune(',') {
						goto l83
					}
					position++
					goto l84
				l83:
					position, tokenIndex = position83, tokenIndex83
				}
			l84:
				{
					position85, tokenIndex85 := position, tokenIndex
					{
						position86, tokenIndex86 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l87
						}
						goto l86
					l87:
						position, tokenIndex = position86, tokenIndex86
						if buffer[position] != rune('(') {
							goto l76
						}
						position++
					}
				l86:
					position, tokenIndex = position85, tokenIndex85
				}
				add(ruleSpeciesEpithet, position77)
			}
			return true
		l76:
			position, tokenIndex = position76, tokenIndex76
			return false
		},
		/* 16 Comparison <- <('c' 'f' '.'?)> */
		nil,
		/* 17 Rank <- <(RankForma / RankVar / RankSsp / RankOther / RankOtherUncommon)> */
		func() bool {
			position89, tokenIndex89 := position, tokenIndex
			{
				position90 := position
				{
					position91, tokenIndex91 := position, tokenIndex
					if !_rules[ruleRankForma]() {
						goto l92
					}
					goto l91
				l92:
					position, tokenIndex = position91, tokenIndex91
					if !_rules[ruleRankVar]() {
						goto l93
					}
					goto l91
				l93:
					position, tokenIndex = position91, tokenIndex91
					if !_rules[ruleRankSsp]() {
						goto l94
					}
					goto l91
				l94:
					position, tokenIndex = position91, tokenIndex91
					if !_rules[ruleRankOther]() {
						goto l95
					}
					goto l91
				l95:
					position, tokenIndex = position91, tokenIndex91
					if !_rules[ruleRankOtherUncommon]() {
						goto l89
					}
				}
			l91:
				add(ruleRank, position90)
			}
			return true
		l89:
			position, tokenIndex = position89, tokenIndex89
			return false
		},
		/* 18 RankOtherUncommon <- <(('*' / ('n' 'a' 't') / ('f' '.' 's' 'p') / ('m' 'u' 't' '.')) &SpaceCharEOI)> */
		func() bool {
			position96, tokenIndex96 := position, tokenIndex
			{
				position97 := position
				{
					position98, tokenIndex98 := position, tokenIndex
					if buffer[position] != rune('*') {
						goto l99
					}
					position++
					goto l98
				l99:
					position, tokenIndex = position98, tokenIndex98
					if buffer[position] != rune('n') {
						goto l100
					}
					position++
					if buffer[position] != rune('a') {
						goto l100
					}
					position++
					if buffer[position] != rune('t') {
						goto l100
					}
					position++
					goto l98
				l100:
					position, tokenIndex = position98, tokenIndex98
					if buffer[position] != rune('f') {
						goto l101
					}
					position++
					if buffer[position] != rune('.') {
						goto l101
					}
					position++
					if buffer[position] != rune('s') {
						goto l101
					}
					position++
					if buffer[position] != rune('p') {
						goto l101
					}
					position++
					goto l98
				l101:
					position, tokenIndex = position98, tokenIndex98
					if buffer[position] != rune('m') {
						goto l96
					}
					position++
					if buffer[position] != rune('u') {
						goto l96
					}
					position++
					if buffer[position] != rune('t') {
						goto l96
					}
					position++
					if buffer[position] != rune('.') {
						goto l96
					}
					position++
				}
			l98:
				{
					position102, tokenIndex102 := position, tokenIndex
					if !_rules[ruleSpaceCharEOI]() {
						goto l96
					}
					position, tokenIndex = position102, tokenIndex102
				}
				add(ruleRankOtherUncommon, position97)
			}
			return true
		l96:
			position, tokenIndex = position96, tokenIndex96
			return false
		},
		/* 19 RankOther <- <((('m' 'o' 'r' 'p' 'h' '.') / ('n' 'o' 't' 'h' 'o' 's' 'u' 'b' 's' 'p' '.') / ('c' 'o' 'n' 'v' 'a' 'r' '.') / ('p' 's' 'e' 'u' 'd' 'o' 'v' 'a' 'r' '.') / ('s' 'e' 'c' 't' '.') / ('s' 'e' 'r' '.') / ('s' 'u' 'b' 'v' 'a' 'r' '.') / ('s' 'u' 'b' 'f' '.') / ('r' 'a' 'c' 'e') / 'α' / ('β' 'β') / 'β' / 'γ' / 'δ' / 'ε' / 'φ' / 'θ' / 'μ' / ('a' '.') / ('b' '.') / ('c' '.') / ('d' '.') / ('e' '.') / ('g' '.') / ('k' '.') / ('p' 'v' '.') / ('p' 'a' 't' 'h' 'o' 'v' 'a' 'r' '.') / ('a' 'b' '.' (_? ('n' '.'))?) / ('s' 't' '.')) &SpaceCharEOI)> */
		func() bool {
			position103, tokenIndex103 := position, tokenIndex
			{
				position104 := position
				{
					position105, tokenIndex105 := position, tokenIndex
					if buffer[position] != rune('m') {
						goto l106
					}
					position++
					if buffer[position] != rune('o') {
						goto l106
					}
					position++
					if buffer[position] != rune('r') {
						goto l106
					}
					position++
					if buffer[position] != rune('p') {
						goto l106
					}
					position++
					if buffer[position] != rune('h') {
						goto l106
					}
					position++
					if buffer[position] != rune('.') {
						goto l106
					}
					position++
					goto l105
				l106:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('n') {
						goto l107
					}
					position++
					if buffer[position] != rune('o') {
						goto l107
					}
					position++
					if buffer[position] != rune('t') {
						goto l107
					}
					position++
					if buffer[position] != rune('h') {
						goto l107
					}
					position++
					if buffer[position] != rune('o') {
						goto l107
					}
					position++
					if buffer[position] != rune('s') {
						goto l107
					}
					position++
					if buffer[position] != rune('u') {
						goto l107
					}
					position++
					if buffer[position] != rune('b') {
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
					if buffer[position] != rune('.') {
						goto l107
					}
					position++
					goto l105
				l107:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('c') {
						goto l108
					}
					position++
					if buffer[position] != rune('o') {
						goto l108
					}
					position++
					if buffer[position] != rune('n') {
						goto l108
					}
					position++
					if buffer[position] != rune('v') {
						goto l108
					}
					position++
					if buffer[position] != rune('a') {
						goto l108
					}
					position++
					if buffer[position] != rune('r') {
						goto l108
					}
					position++
					if buffer[position] != rune('.') {
						goto l108
					}
					position++
					goto l105
				l108:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('p') {
						goto l109
					}
					position++
					if buffer[position] != rune('s') {
						goto l109
					}
					position++
					if buffer[position] != rune('e') {
						goto l109
					}
					position++
					if buffer[position] != rune('u') {
						goto l109
					}
					position++
					if buffer[position] != rune('d') {
						goto l109
					}
					position++
					if buffer[position] != rune('o') {
						goto l109
					}
					position++
					if buffer[position] != rune('v') {
						goto l109
					}
					position++
					if buffer[position] != rune('a') {
						goto l109
					}
					position++
					if buffer[position] != rune('r') {
						goto l109
					}
					position++
					if buffer[position] != rune('.') {
						goto l109
					}
					position++
					goto l105
				l109:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('s') {
						goto l110
					}
					position++
					if buffer[position] != rune('e') {
						goto l110
					}
					position++
					if buffer[position] != rune('c') {
						goto l110
					}
					position++
					if buffer[position] != rune('t') {
						goto l110
					}
					position++
					if buffer[position] != rune('.') {
						goto l110
					}
					position++
					goto l105
				l110:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('s') {
						goto l111
					}
					position++
					if buffer[position] != rune('e') {
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
					goto l105
				l111:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('s') {
						goto l112
					}
					position++
					if buffer[position] != rune('u') {
						goto l112
					}
					position++
					if buffer[position] != rune('b') {
						goto l112
					}
					position++
					if buffer[position] != rune('v') {
						goto l112
					}
					position++
					if buffer[position] != rune('a') {
						goto l112
					}
					position++
					if buffer[position] != rune('r') {
						goto l112
					}
					position++
					if buffer[position] != rune('.') {
						goto l112
					}
					position++
					goto l105
				l112:
					position, tokenIndex = position105, tokenIndex105
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
					if buffer[position] != rune('f') {
						goto l113
					}
					position++
					if buffer[position] != rune('.') {
						goto l113
					}
					position++
					goto l105
				l113:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('r') {
						goto l114
					}
					position++
					if buffer[position] != rune('a') {
						goto l114
					}
					position++
					if buffer[position] != rune('c') {
						goto l114
					}
					position++
					if buffer[position] != rune('e') {
						goto l114
					}
					position++
					goto l105
				l114:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('α') {
						goto l115
					}
					position++
					goto l105
				l115:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('β') {
						goto l116
					}
					position++
					if buffer[position] != rune('β') {
						goto l116
					}
					position++
					goto l105
				l116:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('β') {
						goto l117
					}
					position++
					goto l105
				l117:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('γ') {
						goto l118
					}
					position++
					goto l105
				l118:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('δ') {
						goto l119
					}
					position++
					goto l105
				l119:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('ε') {
						goto l120
					}
					position++
					goto l105
				l120:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('φ') {
						goto l121
					}
					position++
					goto l105
				l121:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('θ') {
						goto l122
					}
					position++
					goto l105
				l122:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('μ') {
						goto l123
					}
					position++
					goto l105
				l123:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('a') {
						goto l124
					}
					position++
					if buffer[position] != rune('.') {
						goto l124
					}
					position++
					goto l105
				l124:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('b') {
						goto l125
					}
					position++
					if buffer[position] != rune('.') {
						goto l125
					}
					position++
					goto l105
				l125:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('c') {
						goto l126
					}
					position++
					if buffer[position] != rune('.') {
						goto l126
					}
					position++
					goto l105
				l126:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('d') {
						goto l127
					}
					position++
					if buffer[position] != rune('.') {
						goto l127
					}
					position++
					goto l105
				l127:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('e') {
						goto l128
					}
					position++
					if buffer[position] != rune('.') {
						goto l128
					}
					position++
					goto l105
				l128:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('g') {
						goto l129
					}
					position++
					if buffer[position] != rune('.') {
						goto l129
					}
					position++
					goto l105
				l129:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('k') {
						goto l130
					}
					position++
					if buffer[position] != rune('.') {
						goto l130
					}
					position++
					goto l105
				l130:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('p') {
						goto l131
					}
					position++
					if buffer[position] != rune('v') {
						goto l131
					}
					position++
					if buffer[position] != rune('.') {
						goto l131
					}
					position++
					goto l105
				l131:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('p') {
						goto l132
					}
					position++
					if buffer[position] != rune('a') {
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
					if buffer[position] != rune('o') {
						goto l132
					}
					position++
					if buffer[position] != rune('v') {
						goto l132
					}
					position++
					if buffer[position] != rune('a') {
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
					goto l105
				l132:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('a') {
						goto l133
					}
					position++
					if buffer[position] != rune('b') {
						goto l133
					}
					position++
					if buffer[position] != rune('.') {
						goto l133
					}
					position++
					{
						position134, tokenIndex134 := position, tokenIndex
						{
							position136, tokenIndex136 := position, tokenIndex
							if !_rules[rule_]() {
								goto l136
							}
							goto l137
						l136:
							position, tokenIndex = position136, tokenIndex136
						}
					l137:
						if buffer[position] != rune('n') {
							goto l134
						}
						position++
						if buffer[position] != rune('.') {
							goto l134
						}
						position++
						goto l135
					l134:
						position, tokenIndex = position134, tokenIndex134
					}
				l135:
					goto l105
				l133:
					position, tokenIndex = position105, tokenIndex105
					if buffer[position] != rune('s') {
						goto l103
					}
					position++
					if buffer[position] != rune('t') {
						goto l103
					}
					position++
					if buffer[position] != rune('.') {
						goto l103
					}
					position++
				}
			l105:
				{
					position138, tokenIndex138 := position, tokenIndex
					if !_rules[ruleSpaceCharEOI]() {
						goto l103
					}
					position, tokenIndex = position138, tokenIndex138
				}
				add(ruleRankOther, position104)
			}
			return true
		l103:
			position, tokenIndex = position103, tokenIndex103
			return false
		},
		/* 20 RankVar <- <(('v' 'a' 'r' 'i' 'e' 't' 'y') / ('[' 'v' 'a' 'r' '.' ']') / ('n' 'v' 'a' 'r' '.') / ('v' 'a' 'r' (&SpaceCharEOI / '.')))> */
		func() bool {
			position139, tokenIndex139 := position, tokenIndex
			{
				position140 := position
				{
					position141, tokenIndex141 := position, tokenIndex
					if buffer[position] != rune('v') {
						goto l142
					}
					position++
					if buffer[position] != rune('a') {
						goto l142
					}
					position++
					if buffer[position] != rune('r') {
						goto l142
					}
					position++
					if buffer[position] != rune('i') {
						goto l142
					}
					position++
					if buffer[position] != rune('e') {
						goto l142
					}
					position++
					if buffer[position] != rune('t') {
						goto l142
					}
					position++
					if buffer[position] != rune('y') {
						goto l142
					}
					position++
					goto l141
				l142:
					position, tokenIndex = position141, tokenIndex141
					if buffer[position] != rune('[') {
						goto l143
					}
					position++
					if buffer[position] != rune('v') {
						goto l143
					}
					position++
					if buffer[position] != rune('a') {
						goto l143
					}
					position++
					if buffer[position] != rune('r') {
						goto l143
					}
					position++
					if buffer[position] != rune('.') {
						goto l143
					}
					position++
					if buffer[position] != rune(']') {
						goto l143
					}
					position++
					goto l141
				l143:
					position, tokenIndex = position141, tokenIndex141
					if buffer[position] != rune('n') {
						goto l144
					}
					position++
					if buffer[position] != rune('v') {
						goto l144
					}
					position++
					if buffer[position] != rune('a') {
						goto l144
					}
					position++
					if buffer[position] != rune('r') {
						goto l144
					}
					position++
					if buffer[position] != rune('.') {
						goto l144
					}
					position++
					goto l141
				l144:
					position, tokenIndex = position141, tokenIndex141
					if buffer[position] != rune('v') {
						goto l139
					}
					position++
					if buffer[position] != rune('a') {
						goto l139
					}
					position++
					if buffer[position] != rune('r') {
						goto l139
					}
					position++
					{
						position145, tokenIndex145 := position, tokenIndex
						{
							position147, tokenIndex147 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l146
							}
							position, tokenIndex = position147, tokenIndex147
						}
						goto l145
					l146:
						position, tokenIndex = position145, tokenIndex145
						if buffer[position] != rune('.') {
							goto l139
						}
						position++
					}
				l145:
				}
			l141:
				add(ruleRankVar, position140)
			}
			return true
		l139:
			position, tokenIndex = position139, tokenIndex139
			return false
		},
		/* 21 RankForma <- <((('f' 'o' 'r' 'm' 'a') / ('f' 'm' 'a') / ('f' 'o' 'r' 'm') / ('f' 'o') / 'f') (&SpaceCharEOI / '.'))> */
		func() bool {
			position148, tokenIndex148 := position, tokenIndex
			{
				position149 := position
				{
					position150, tokenIndex150 := position, tokenIndex
					if buffer[position] != rune('f') {
						goto l151
					}
					position++
					if buffer[position] != rune('o') {
						goto l151
					}
					position++
					if buffer[position] != rune('r') {
						goto l151
					}
					position++
					if buffer[position] != rune('m') {
						goto l151
					}
					position++
					if buffer[position] != rune('a') {
						goto l151
					}
					position++
					goto l150
				l151:
					position, tokenIndex = position150, tokenIndex150
					if buffer[position] != rune('f') {
						goto l152
					}
					position++
					if buffer[position] != rune('m') {
						goto l152
					}
					position++
					if buffer[position] != rune('a') {
						goto l152
					}
					position++
					goto l150
				l152:
					position, tokenIndex = position150, tokenIndex150
					if buffer[position] != rune('f') {
						goto l153
					}
					position++
					if buffer[position] != rune('o') {
						goto l153
					}
					position++
					if buffer[position] != rune('r') {
						goto l153
					}
					position++
					if buffer[position] != rune('m') {
						goto l153
					}
					position++
					goto l150
				l153:
					position, tokenIndex = position150, tokenIndex150
					if buffer[position] != rune('f') {
						goto l154
					}
					position++
					if buffer[position] != rune('o') {
						goto l154
					}
					position++
					goto l150
				l154:
					position, tokenIndex = position150, tokenIndex150
					if buffer[position] != rune('f') {
						goto l148
					}
					position++
				}
			l150:
				{
					position155, tokenIndex155 := position, tokenIndex
					{
						position157, tokenIndex157 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l156
						}
						position, tokenIndex = position157, tokenIndex157
					}
					goto l155
				l156:
					position, tokenIndex = position155, tokenIndex155
					if buffer[position] != rune('.') {
						goto l148
					}
					position++
				}
			l155:
				add(ruleRankForma, position149)
			}
			return true
		l148:
			position, tokenIndex = position148, tokenIndex148
			return false
		},
		/* 22 RankSsp <- <((('s' 's' 'p') / ('s' 'u' 'b' 's' 'p')) (&SpaceCharEOI / '.'))> */
		func() bool {
			position158, tokenIndex158 := position, tokenIndex
			{
				position159 := position
				{
					position160, tokenIndex160 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l161
					}
					position++
					if buffer[position] != rune('s') {
						goto l161
					}
					position++
					if buffer[position] != rune('p') {
						goto l161
					}
					position++
					goto l160
				l161:
					position, tokenIndex = position160, tokenIndex160
					if buffer[position] != rune('s') {
						goto l158
					}
					position++
					if buffer[position] != rune('u') {
						goto l158
					}
					position++
					if buffer[position] != rune('b') {
						goto l158
					}
					position++
					if buffer[position] != rune('s') {
						goto l158
					}
					position++
					if buffer[position] != rune('p') {
						goto l158
					}
					position++
				}
			l160:
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
						goto l158
					}
					position++
				}
			l162:
				add(ruleRankSsp, position159)
			}
			return true
		l158:
			position, tokenIndex = position158, tokenIndex158
			return false
		},
		/* 23 SubGenusOrSuperspecies <- <('(' _? Word _? ')')> */
		func() bool {
			position165, tokenIndex165 := position, tokenIndex
			{
				position166 := position
				if buffer[position] != rune('(') {
					goto l165
				}
				position++
				{
					position167, tokenIndex167 := position, tokenIndex
					if !_rules[rule_]() {
						goto l167
					}
					goto l168
				l167:
					position, tokenIndex = position167, tokenIndex167
				}
			l168:
				if !_rules[ruleWord]() {
					goto l165
				}
				{
					position169, tokenIndex169 := position, tokenIndex
					if !_rules[rule_]() {
						goto l169
					}
					goto l170
				l169:
					position, tokenIndex = position169, tokenIndex169
				}
			l170:
				if buffer[position] != rune(')') {
					goto l165
				}
				position++
				add(ruleSubGenusOrSuperspecies, position166)
			}
			return true
		l165:
			position, tokenIndex = position165, tokenIndex165
			return false
		},
		/* 24 SubGenus <- <('(' _? UninomialWord _? ')')> */
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
				if !_rules[ruleUninomialWord]() {
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
				add(ruleSubGenus, position172)
			}
			return true
		l171:
			position, tokenIndex = position171, tokenIndex171
			return false
		},
		/* 25 UninomialCombo <- <(UninomialCombo1 / UninomialCombo2)> */
		func() bool {
			position177, tokenIndex177 := position, tokenIndex
			{
				position178 := position
				{
					position179, tokenIndex179 := position, tokenIndex
					if !_rules[ruleUninomialCombo1]() {
						goto l180
					}
					goto l179
				l180:
					position, tokenIndex = position179, tokenIndex179
					if !_rules[ruleUninomialCombo2]() {
						goto l177
					}
				}
			l179:
				add(ruleUninomialCombo, position178)
			}
			return true
		l177:
			position, tokenIndex = position177, tokenIndex177
			return false
		},
		/* 26 UninomialCombo1 <- <(UninomialWord _? SubGenus _? Authorship .?)> */
		func() bool {
			position181, tokenIndex181 := position, tokenIndex
			{
				position182 := position
				if !_rules[ruleUninomialWord]() {
					goto l181
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
				if !_rules[ruleSubGenus]() {
					goto l181
				}
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
				if !_rules[ruleAuthorship]() {
					goto l181
				}
				{
					position187, tokenIndex187 := position, tokenIndex
					if !matchDot() {
						goto l187
					}
					goto l188
				l187:
					position, tokenIndex = position187, tokenIndex187
				}
			l188:
				add(ruleUninomialCombo1, position182)
			}
			return true
		l181:
			position, tokenIndex = position181, tokenIndex181
			return false
		},
		/* 27 UninomialCombo2 <- <(Uninomial _? RankUninomial _? Uninomial)> */
		func() bool {
			position189, tokenIndex189 := position, tokenIndex
			{
				position190 := position
				if !_rules[ruleUninomial]() {
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
				if !_rules[ruleRankUninomial]() {
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
				if !_rules[ruleUninomial]() {
					goto l189
				}
				add(ruleUninomialCombo2, position190)
			}
			return true
		l189:
			position, tokenIndex = position189, tokenIndex189
			return false
		},
		/* 28 RankUninomial <- <((('s' 'e' 'c' 't') / ('s' 'u' 'b' 's' 'e' 'c' 't') / ('t' 'r' 'i' 'b') / ('s' 'u' 'b' 't' 'r' 'i' 'b') / ('s' 'u' 'b' 's' 'e' 'r') / ('s' 'e' 'r') / ('s' 'u' 'b' 'g' 'e' 'n') / ('f' 'a' 'm') / ('s' 'u' 'b' 'f' 'a' 'm') / ('s' 'u' 'p' 'e' 'r' 't' 'r' 'i' 'b')) '.'?)> */
		func() bool {
			position195, tokenIndex195 := position, tokenIndex
			{
				position196 := position
				{
					position197, tokenIndex197 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l198
					}
					position++
					if buffer[position] != rune('e') {
						goto l198
					}
					position++
					if buffer[position] != rune('c') {
						goto l198
					}
					position++
					if buffer[position] != rune('t') {
						goto l198
					}
					position++
					goto l197
				l198:
					position, tokenIndex = position197, tokenIndex197
					if buffer[position] != rune('s') {
						goto l199
					}
					position++
					if buffer[position] != rune('u') {
						goto l199
					}
					position++
					if buffer[position] != rune('b') {
						goto l199
					}
					position++
					if buffer[position] != rune('s') {
						goto l199
					}
					position++
					if buffer[position] != rune('e') {
						goto l199
					}
					position++
					if buffer[position] != rune('c') {
						goto l199
					}
					position++
					if buffer[position] != rune('t') {
						goto l199
					}
					position++
					goto l197
				l199:
					position, tokenIndex = position197, tokenIndex197
					if buffer[position] != rune('t') {
						goto l200
					}
					position++
					if buffer[position] != rune('r') {
						goto l200
					}
					position++
					if buffer[position] != rune('i') {
						goto l200
					}
					position++
					if buffer[position] != rune('b') {
						goto l200
					}
					position++
					goto l197
				l200:
					position, tokenIndex = position197, tokenIndex197
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
					goto l197
				l201:
					position, tokenIndex = position197, tokenIndex197
					if buffer[position] != rune('s') {
						goto l202
					}
					position++
					if buffer[position] != rune('u') {
						goto l202
					}
					position++
					if buffer[position] != rune('b') {
						goto l202
					}
					position++
					if buffer[position] != rune('s') {
						goto l202
					}
					position++
					if buffer[position] != rune('e') {
						goto l202
					}
					position++
					if buffer[position] != rune('r') {
						goto l202
					}
					position++
					goto l197
				l202:
					position, tokenIndex = position197, tokenIndex197
					if buffer[position] != rune('s') {
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
					goto l197
				l203:
					position, tokenIndex = position197, tokenIndex197
					if buffer[position] != rune('s') {
						goto l204
					}
					position++
					if buffer[position] != rune('u') {
						goto l204
					}
					position++
					if buffer[position] != rune('b') {
						goto l204
					}
					position++
					if buffer[position] != rune('g') {
						goto l204
					}
					position++
					if buffer[position] != rune('e') {
						goto l204
					}
					position++
					if buffer[position] != rune('n') {
						goto l204
					}
					position++
					goto l197
				l204:
					position, tokenIndex = position197, tokenIndex197
					if buffer[position] != rune('f') {
						goto l205
					}
					position++
					if buffer[position] != rune('a') {
						goto l205
					}
					position++
					if buffer[position] != rune('m') {
						goto l205
					}
					position++
					goto l197
				l205:
					position, tokenIndex = position197, tokenIndex197
					if buffer[position] != rune('s') {
						goto l206
					}
					position++
					if buffer[position] != rune('u') {
						goto l206
					}
					position++
					if buffer[position] != rune('b') {
						goto l206
					}
					position++
					if buffer[position] != rune('f') {
						goto l206
					}
					position++
					if buffer[position] != rune('a') {
						goto l206
					}
					position++
					if buffer[position] != rune('m') {
						goto l206
					}
					position++
					goto l197
				l206:
					position, tokenIndex = position197, tokenIndex197
					if buffer[position] != rune('s') {
						goto l195
					}
					position++
					if buffer[position] != rune('u') {
						goto l195
					}
					position++
					if buffer[position] != rune('p') {
						goto l195
					}
					position++
					if buffer[position] != rune('e') {
						goto l195
					}
					position++
					if buffer[position] != rune('r') {
						goto l195
					}
					position++
					if buffer[position] != rune('t') {
						goto l195
					}
					position++
					if buffer[position] != rune('r') {
						goto l195
					}
					position++
					if buffer[position] != rune('i') {
						goto l195
					}
					position++
					if buffer[position] != rune('b') {
						goto l195
					}
					position++
				}
			l197:
				{
					position207, tokenIndex207 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l207
					}
					position++
					goto l208
				l207:
					position, tokenIndex = position207, tokenIndex207
				}
			l208:
				add(ruleRankUninomial, position196)
			}
			return true
		l195:
			position, tokenIndex = position195, tokenIndex195
			return false
		},
		/* 29 Uninomial <- <(UninomialWord (_ Authorship)?)> */
		func() bool {
			position209, tokenIndex209 := position, tokenIndex
			{
				position210 := position
				if !_rules[ruleUninomialWord]() {
					goto l209
				}
				{
					position211, tokenIndex211 := position, tokenIndex
					if !_rules[rule_]() {
						goto l211
					}
					if !_rules[ruleAuthorship]() {
						goto l211
					}
					goto l212
				l211:
					position, tokenIndex = position211, tokenIndex211
				}
			l212:
				add(ruleUninomial, position210)
			}
			return true
		l209:
			position, tokenIndex = position209, tokenIndex209
			return false
		},
		/* 30 UninomialWord <- <(CapWord / TwoLetterGenus)> */
		func() bool {
			position213, tokenIndex213 := position, tokenIndex
			{
				position214 := position
				{
					position215, tokenIndex215 := position, tokenIndex
					if !_rules[ruleCapWord]() {
						goto l216
					}
					goto l215
				l216:
					position, tokenIndex = position215, tokenIndex215
					if !_rules[ruleTwoLetterGenus]() {
						goto l213
					}
				}
			l215:
				add(ruleUninomialWord, position214)
			}
			return true
		l213:
			position, tokenIndex = position213, tokenIndex213
			return false
		},
		/* 31 AbbrGenus <- <(UpperChar LowerChar* '.')> */
		func() bool {
			position217, tokenIndex217 := position, tokenIndex
			{
				position218 := position
				if !_rules[ruleUpperChar]() {
					goto l217
				}
			l219:
				{
					position220, tokenIndex220 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l220
					}
					goto l219
				l220:
					position, tokenIndex = position220, tokenIndex220
				}
				if buffer[position] != rune('.') {
					goto l217
				}
				position++
				add(ruleAbbrGenus, position218)
			}
			return true
		l217:
			position, tokenIndex = position217, tokenIndex217
			return false
		},
		/* 32 CapWord <- <(CapWord2 / CapWord1)> */
		func() bool {
			position221, tokenIndex221 := position, tokenIndex
			{
				position222 := position
				{
					position223, tokenIndex223 := position, tokenIndex
					if !_rules[ruleCapWord2]() {
						goto l224
					}
					goto l223
				l224:
					position, tokenIndex = position223, tokenIndex223
					if !_rules[ruleCapWord1]() {
						goto l221
					}
				}
			l223:
				add(ruleCapWord, position222)
			}
			return true
		l221:
			position, tokenIndex = position221, tokenIndex221
			return false
		},
		/* 33 CapWord1 <- <(NameUpperChar NameLowerChar NameLowerChar+ '?'?)> */
		func() bool {
			position225, tokenIndex225 := position, tokenIndex
			{
				position226 := position
				if !_rules[ruleNameUpperChar]() {
					goto l225
				}
				if !_rules[ruleNameLowerChar]() {
					goto l225
				}
				if !_rules[ruleNameLowerChar]() {
					goto l225
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
					if buffer[position] != rune('?') {
						goto l229
					}
					position++
					goto l230
				l229:
					position, tokenIndex = position229, tokenIndex229
				}
			l230:
				add(ruleCapWord1, position226)
			}
			return true
		l225:
			position, tokenIndex = position225, tokenIndex225
			return false
		},
		/* 34 CapWord2 <- <(CapWord1 dash (CapWord1 / Word1))> */
		func() bool {
			position231, tokenIndex231 := position, tokenIndex
			{
				position232 := position
				if !_rules[ruleCapWord1]() {
					goto l231
				}
				if !_rules[ruledash]() {
					goto l231
				}
				{
					position233, tokenIndex233 := position, tokenIndex
					if !_rules[ruleCapWord1]() {
						goto l234
					}
					goto l233
				l234:
					position, tokenIndex = position233, tokenIndex233
					if !_rules[ruleWord1]() {
						goto l231
					}
				}
			l233:
				add(ruleCapWord2, position232)
			}
			return true
		l231:
			position, tokenIndex = position231, tokenIndex231
			return false
		},
		/* 35 TwoLetterGenus <- <(('C' 'a') / ('E' 'a') / ('G' 'e') / ('I' 'a') / ('I' 'o') / ('I' 'x') / ('L' 'o') / ('O' 'a') / ('R' 'a') / ('T' 'y') / ('U' 'a') / ('A' 'a') / ('J' 'a') / ('Z' 'u') / ('L' 'a') / ('Q' 'u') / ('A' 's') / ('B' 'a'))> */
		func() bool {
			position235, tokenIndex235 := position, tokenIndex
			{
				position236 := position
				{
					position237, tokenIndex237 := position, tokenIndex
					if buffer[position] != rune('C') {
						goto l238
					}
					position++
					if buffer[position] != rune('a') {
						goto l238
					}
					position++
					goto l237
				l238:
					position, tokenIndex = position237, tokenIndex237
					if buffer[position] != rune('E') {
						goto l239
					}
					position++
					if buffer[position] != rune('a') {
						goto l239
					}
					position++
					goto l237
				l239:
					position, tokenIndex = position237, tokenIndex237
					if buffer[position] != rune('G') {
						goto l240
					}
					position++
					if buffer[position] != rune('e') {
						goto l240
					}
					position++
					goto l237
				l240:
					position, tokenIndex = position237, tokenIndex237
					if buffer[position] != rune('I') {
						goto l241
					}
					position++
					if buffer[position] != rune('a') {
						goto l241
					}
					position++
					goto l237
				l241:
					position, tokenIndex = position237, tokenIndex237
					if buffer[position] != rune('I') {
						goto l242
					}
					position++
					if buffer[position] != rune('o') {
						goto l242
					}
					position++
					goto l237
				l242:
					position, tokenIndex = position237, tokenIndex237
					if buffer[position] != rune('I') {
						goto l243
					}
					position++
					if buffer[position] != rune('x') {
						goto l243
					}
					position++
					goto l237
				l243:
					position, tokenIndex = position237, tokenIndex237
					if buffer[position] != rune('L') {
						goto l244
					}
					position++
					if buffer[position] != rune('o') {
						goto l244
					}
					position++
					goto l237
				l244:
					position, tokenIndex = position237, tokenIndex237
					if buffer[position] != rune('O') {
						goto l245
					}
					position++
					if buffer[position] != rune('a') {
						goto l245
					}
					position++
					goto l237
				l245:
					position, tokenIndex = position237, tokenIndex237
					if buffer[position] != rune('R') {
						goto l246
					}
					position++
					if buffer[position] != rune('a') {
						goto l246
					}
					position++
					goto l237
				l246:
					position, tokenIndex = position237, tokenIndex237
					if buffer[position] != rune('T') {
						goto l247
					}
					position++
					if buffer[position] != rune('y') {
						goto l247
					}
					position++
					goto l237
				l247:
					position, tokenIndex = position237, tokenIndex237
					if buffer[position] != rune('U') {
						goto l248
					}
					position++
					if buffer[position] != rune('a') {
						goto l248
					}
					position++
					goto l237
				l248:
					position, tokenIndex = position237, tokenIndex237
					if buffer[position] != rune('A') {
						goto l249
					}
					position++
					if buffer[position] != rune('a') {
						goto l249
					}
					position++
					goto l237
				l249:
					position, tokenIndex = position237, tokenIndex237
					if buffer[position] != rune('J') {
						goto l250
					}
					position++
					if buffer[position] != rune('a') {
						goto l250
					}
					position++
					goto l237
				l250:
					position, tokenIndex = position237, tokenIndex237
					if buffer[position] != rune('Z') {
						goto l251
					}
					position++
					if buffer[position] != rune('u') {
						goto l251
					}
					position++
					goto l237
				l251:
					position, tokenIndex = position237, tokenIndex237
					if buffer[position] != rune('L') {
						goto l252
					}
					position++
					if buffer[position] != rune('a') {
						goto l252
					}
					position++
					goto l237
				l252:
					position, tokenIndex = position237, tokenIndex237
					if buffer[position] != rune('Q') {
						goto l253
					}
					position++
					if buffer[position] != rune('u') {
						goto l253
					}
					position++
					goto l237
				l253:
					position, tokenIndex = position237, tokenIndex237
					if buffer[position] != rune('A') {
						goto l254
					}
					position++
					if buffer[position] != rune('s') {
						goto l254
					}
					position++
					goto l237
				l254:
					position, tokenIndex = position237, tokenIndex237
					if buffer[position] != rune('B') {
						goto l235
					}
					position++
					if buffer[position] != rune('a') {
						goto l235
					}
					position++
				}
			l237:
				add(ruleTwoLetterGenus, position236)
			}
			return true
		l235:
			position, tokenIndex = position235, tokenIndex235
			return false
		},
		/* 36 Word <- <(!(AuthorPrefix / RankUninomial / Approximation / Word4) (Word3 / Word2StartDigit / Word2 / Word1) &(SpaceCharEOI / '('))> */
		func() bool {
			position255, tokenIndex255 := position, tokenIndex
			{
				position256 := position
				{
					position257, tokenIndex257 := position, tokenIndex
					{
						position258, tokenIndex258 := position, tokenIndex
						if !_rules[ruleAuthorPrefix]() {
							goto l259
						}
						goto l258
					l259:
						position, tokenIndex = position258, tokenIndex258
						if !_rules[ruleRankUninomial]() {
							goto l260
						}
						goto l258
					l260:
						position, tokenIndex = position258, tokenIndex258
						if !_rules[ruleApproximation]() {
							goto l261
						}
						goto l258
					l261:
						position, tokenIndex = position258, tokenIndex258
						if !_rules[ruleWord4]() {
							goto l257
						}
					}
				l258:
					goto l255
				l257:
					position, tokenIndex = position257, tokenIndex257
				}
				{
					position262, tokenIndex262 := position, tokenIndex
					if !_rules[ruleWord3]() {
						goto l263
					}
					goto l262
				l263:
					position, tokenIndex = position262, tokenIndex262
					if !_rules[ruleWord2StartDigit]() {
						goto l264
					}
					goto l262
				l264:
					position, tokenIndex = position262, tokenIndex262
					if !_rules[ruleWord2]() {
						goto l265
					}
					goto l262
				l265:
					position, tokenIndex = position262, tokenIndex262
					if !_rules[ruleWord1]() {
						goto l255
					}
				}
			l262:
				{
					position266, tokenIndex266 := position, tokenIndex
					{
						position267, tokenIndex267 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l268
						}
						goto l267
					l268:
						position, tokenIndex = position267, tokenIndex267
						if buffer[position] != rune('(') {
							goto l255
						}
						position++
					}
				l267:
					position, tokenIndex = position266, tokenIndex266
				}
				add(ruleWord, position256)
			}
			return true
		l255:
			position, tokenIndex = position255, tokenIndex255
			return false
		},
		/* 37 Word1 <- <((lASCII dash)? NameLowerChar NameLowerChar+)> */
		func() bool {
			position269, tokenIndex269 := position, tokenIndex
			{
				position270 := position
				{
					position271, tokenIndex271 := position, tokenIndex
					if !_rules[rulelASCII]() {
						goto l271
					}
					if !_rules[ruledash]() {
						goto l271
					}
					goto l272
				l271:
					position, tokenIndex = position271, tokenIndex271
				}
			l272:
				if !_rules[ruleNameLowerChar]() {
					goto l269
				}
				if !_rules[ruleNameLowerChar]() {
					goto l269
				}
			l273:
				{
					position274, tokenIndex274 := position, tokenIndex
					if !_rules[ruleNameLowerChar]() {
						goto l274
					}
					goto l273
				l274:
					position, tokenIndex = position274, tokenIndex274
				}
				add(ruleWord1, position270)
			}
			return true
		l269:
			position, tokenIndex = position269, tokenIndex269
			return false
		},
		/* 38 Word2StartDigit <- <(('1' / '2' / '3' / '4' / '5' / '6' / '7' / '8' / '9') nums? ('.' / dash)? NameLowerChar NameLowerChar NameLowerChar NameLowerChar+)> */
		func() bool {
			position275, tokenIndex275 := position, tokenIndex
			{
				position276 := position
				{
					position277, tokenIndex277 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l278
					}
					position++
					goto l277
				l278:
					position, tokenIndex = position277, tokenIndex277
					if buffer[position] != rune('2') {
						goto l279
					}
					position++
					goto l277
				l279:
					position, tokenIndex = position277, tokenIndex277
					if buffer[position] != rune('3') {
						goto l280
					}
					position++
					goto l277
				l280:
					position, tokenIndex = position277, tokenIndex277
					if buffer[position] != rune('4') {
						goto l281
					}
					position++
					goto l277
				l281:
					position, tokenIndex = position277, tokenIndex277
					if buffer[position] != rune('5') {
						goto l282
					}
					position++
					goto l277
				l282:
					position, tokenIndex = position277, tokenIndex277
					if buffer[position] != rune('6') {
						goto l283
					}
					position++
					goto l277
				l283:
					position, tokenIndex = position277, tokenIndex277
					if buffer[position] != rune('7') {
						goto l284
					}
					position++
					goto l277
				l284:
					position, tokenIndex = position277, tokenIndex277
					if buffer[position] != rune('8') {
						goto l285
					}
					position++
					goto l277
				l285:
					position, tokenIndex = position277, tokenIndex277
					if buffer[position] != rune('9') {
						goto l275
					}
					position++
				}
			l277:
				{
					position286, tokenIndex286 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l286
					}
					goto l287
				l286:
					position, tokenIndex = position286, tokenIndex286
				}
			l287:
				{
					position288, tokenIndex288 := position, tokenIndex
					{
						position290, tokenIndex290 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l291
						}
						position++
						goto l290
					l291:
						position, tokenIndex = position290, tokenIndex290
						if !_rules[ruledash]() {
							goto l288
						}
					}
				l290:
					goto l289
				l288:
					position, tokenIndex = position288, tokenIndex288
				}
			l289:
				if !_rules[ruleNameLowerChar]() {
					goto l275
				}
				if !_rules[ruleNameLowerChar]() {
					goto l275
				}
				if !_rules[ruleNameLowerChar]() {
					goto l275
				}
				if !_rules[ruleNameLowerChar]() {
					goto l275
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
				add(ruleWord2StartDigit, position276)
			}
			return true
		l275:
			position, tokenIndex = position275, tokenIndex275
			return false
		},
		/* 39 Word2 <- <(NameLowerChar+ dash? NameLowerChar+)> */
		func() bool {
			position294, tokenIndex294 := position, tokenIndex
			{
				position295 := position
				if !_rules[ruleNameLowerChar]() {
					goto l294
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
				{
					position298, tokenIndex298 := position, tokenIndex
					if !_rules[ruledash]() {
						goto l298
					}
					goto l299
				l298:
					position, tokenIndex = position298, tokenIndex298
				}
			l299:
				if !_rules[ruleNameLowerChar]() {
					goto l294
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
				add(ruleWord2, position295)
			}
			return true
		l294:
			position, tokenIndex = position294, tokenIndex294
			return false
		},
		/* 40 Word3 <- <(NameLowerChar NameLowerChar* apostr Word1)> */
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
				if !_rules[ruleapostr]() {
					goto l302
				}
				if !_rules[ruleWord1]() {
					goto l302
				}
				add(ruleWord3, position303)
			}
			return true
		l302:
			position, tokenIndex = position302, tokenIndex302
			return false
		},
		/* 41 Word4 <- <(NameLowerChar+ '.' NameLowerChar)> */
		func() bool {
			position306, tokenIndex306 := position, tokenIndex
			{
				position307 := position
				if !_rules[ruleNameLowerChar]() {
					goto l306
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
				if buffer[position] != rune('.') {
					goto l306
				}
				position++
				if !_rules[ruleNameLowerChar]() {
					goto l306
				}
				add(ruleWord4, position307)
			}
			return true
		l306:
			position, tokenIndex = position306, tokenIndex306
			return false
		},
		/* 42 HybridChar <- <'×'> */
		func() bool {
			position310, tokenIndex310 := position, tokenIndex
			{
				position311 := position
				if buffer[position] != rune('×') {
					goto l310
				}
				position++
				add(ruleHybridChar, position311)
			}
			return true
		l310:
			position, tokenIndex = position310, tokenIndex310
			return false
		},
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
			position316, tokenIndex316 := position, tokenIndex
			{
				position317 := position
				{
					position318, tokenIndex318 := position, tokenIndex
					if buffer[position] != rune('s') {
						goto l319
					}
					position++
					if buffer[position] != rune('p') {
						goto l319
					}
					position++
					if buffer[position] != rune('.') {
						goto l319
					}
					position++
					{
						position320, tokenIndex320 := position, tokenIndex
						if !_rules[rule_]() {
							goto l320
						}
						goto l321
					l320:
						position, tokenIndex = position320, tokenIndex320
					}
				l321:
					if buffer[position] != rune('n') {
						goto l319
					}
					position++
					if buffer[position] != rune('r') {
						goto l319
					}
					position++
					if buffer[position] != rune('.') {
						goto l319
					}
					position++
					goto l318
				l319:
					position, tokenIndex = position318, tokenIndex318
					if buffer[position] != rune('s') {
						goto l322
					}
					position++
					if buffer[position] != rune('p') {
						goto l322
					}
					position++
					if buffer[position] != rune('.') {
						goto l322
					}
					position++
					{
						position323, tokenIndex323 := position, tokenIndex
						if !_rules[rule_]() {
							goto l323
						}
						goto l324
					l323:
						position, tokenIndex = position323, tokenIndex323
					}
				l324:
					if buffer[position] != rune('a') {
						goto l322
					}
					position++
					if buffer[position] != rune('f') {
						goto l322
					}
					position++
					if buffer[position] != rune('f') {
						goto l322
					}
					position++
					if buffer[position] != rune('.') {
						goto l322
					}
					position++
					goto l318
				l322:
					position, tokenIndex = position318, tokenIndex318
					if buffer[position] != rune('m') {
						goto l325
					}
					position++
					if buffer[position] != rune('o') {
						goto l325
					}
					position++
					if buffer[position] != rune('n') {
						goto l325
					}
					position++
					if buffer[position] != rune('s') {
						goto l325
					}
					position++
					if buffer[position] != rune('t') {
						goto l325
					}
					position++
					if buffer[position] != rune('.') {
						goto l325
					}
					position++
					goto l318
				l325:
					position, tokenIndex = position318, tokenIndex318
					if buffer[position] != rune('?') {
						goto l326
					}
					position++
					goto l318
				l326:
					position, tokenIndex = position318, tokenIndex318
					{
						position327, tokenIndex327 := position, tokenIndex
						if buffer[position] != rune('s') {
							goto l328
						}
						position++
						if buffer[position] != rune('p') {
							goto l328
						}
						position++
						if buffer[position] != rune('p') {
							goto l328
						}
						position++
						goto l327
					l328:
						position, tokenIndex = position327, tokenIndex327
						if buffer[position] != rune('n') {
							goto l329
						}
						position++
						if buffer[position] != rune('r') {
							goto l329
						}
						position++
						goto l327
					l329:
						position, tokenIndex = position327, tokenIndex327
						if buffer[position] != rune('s') {
							goto l330
						}
						position++
						if buffer[position] != rune('p') {
							goto l330
						}
						position++
						goto l327
					l330:
						position, tokenIndex = position327, tokenIndex327
						if buffer[position] != rune('a') {
							goto l331
						}
						position++
						if buffer[position] != rune('f') {
							goto l331
						}
						position++
						if buffer[position] != rune('f') {
							goto l331
						}
						position++
						goto l327
					l331:
						position, tokenIndex = position327, tokenIndex327
						if buffer[position] != rune('s') {
							goto l316
						}
						position++
						if buffer[position] != rune('p') {
							goto l316
						}
						position++
						if buffer[position] != rune('e') {
							goto l316
						}
						position++
						if buffer[position] != rune('c') {
							goto l316
						}
						position++
						if buffer[position] != rune('i') {
							goto l316
						}
						position++
						if buffer[position] != rune('e') {
							goto l316
						}
						position++
						if buffer[position] != rune('s') {
							goto l316
						}
						position++
					}
				l327:
					{
						position332, tokenIndex332 := position, tokenIndex
						{
							position334, tokenIndex334 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l333
							}
							position, tokenIndex = position334, tokenIndex334
						}
						goto l332
					l333:
						position, tokenIndex = position332, tokenIndex332
						if buffer[position] != rune('.') {
							goto l316
						}
						position++
					}
				l332:
				}
			l318:
				add(ruleApproximation, position317)
			}
			return true
		l316:
			position, tokenIndex = position316, tokenIndex316
			return false
		},
		/* 48 Authorship <- <((AuthorshipCombo / OriginalAuthorship) &(SpaceCharEOI / ('\\' / '(' / ',' / ':')))> */
		func() bool {
			position335, tokenIndex335 := position, tokenIndex
			{
				position336 := position
				{
					position337, tokenIndex337 := position, tokenIndex
					if !_rules[ruleAuthorshipCombo]() {
						goto l338
					}
					goto l337
				l338:
					position, tokenIndex = position337, tokenIndex337
					if !_rules[ruleOriginalAuthorship]() {
						goto l335
					}
				}
			l337:
				{
					position339, tokenIndex339 := position, tokenIndex
					{
						position340, tokenIndex340 := position, tokenIndex
						if !_rules[ruleSpaceCharEOI]() {
							goto l341
						}
						goto l340
					l341:
						position, tokenIndex = position340, tokenIndex340
						{
							position342, tokenIndex342 := position, tokenIndex
							if buffer[position] != rune('\\') {
								goto l343
							}
							position++
							goto l342
						l343:
							position, tokenIndex = position342, tokenIndex342
							if buffer[position] != rune('(') {
								goto l344
							}
							position++
							goto l342
						l344:
							position, tokenIndex = position342, tokenIndex342
							if buffer[position] != rune(',') {
								goto l345
							}
							position++
							goto l342
						l345:
							position, tokenIndex = position342, tokenIndex342
							if buffer[position] != rune(':') {
								goto l335
							}
							position++
						}
					l342:
					}
				l340:
					position, tokenIndex = position339, tokenIndex339
				}
				add(ruleAuthorship, position336)
			}
			return true
		l335:
			position, tokenIndex = position335, tokenIndex335
			return false
		},
		/* 49 AuthorshipCombo <- <(OriginalAuthorship _? CombinationAuthorship)> */
		func() bool {
			position346, tokenIndex346 := position, tokenIndex
			{
				position347 := position
				if !_rules[ruleOriginalAuthorship]() {
					goto l346
				}
				{
					position348, tokenIndex348 := position, tokenIndex
					if !_rules[rule_]() {
						goto l348
					}
					goto l349
				l348:
					position, tokenIndex = position348, tokenIndex348
				}
			l349:
				if !_rules[ruleCombinationAuthorship]() {
					goto l346
				}
				add(ruleAuthorshipCombo, position347)
			}
			return true
		l346:
			position, tokenIndex = position346, tokenIndex346
			return false
		},
		/* 50 OriginalAuthorship <- <(AuthorsGroup / BasionymAuthorship / BasionymAuthorshipYearMisformed)> */
		func() bool {
			position350, tokenIndex350 := position, tokenIndex
			{
				position351 := position
				{
					position352, tokenIndex352 := position, tokenIndex
					if !_rules[ruleAuthorsGroup]() {
						goto l353
					}
					goto l352
				l353:
					position, tokenIndex = position352, tokenIndex352
					if !_rules[ruleBasionymAuthorship]() {
						goto l354
					}
					goto l352
				l354:
					position, tokenIndex = position352, tokenIndex352
					if !_rules[ruleBasionymAuthorshipYearMisformed]() {
						goto l350
					}
				}
			l352:
				add(ruleOriginalAuthorship, position351)
			}
			return true
		l350:
			position, tokenIndex = position350, tokenIndex350
			return false
		},
		/* 51 CombinationAuthorship <- <AuthorsGroup> */
		func() bool {
			position355, tokenIndex355 := position, tokenIndex
			{
				position356 := position
				if !_rules[ruleAuthorsGroup]() {
					goto l355
				}
				add(ruleCombinationAuthorship, position356)
			}
			return true
		l355:
			position, tokenIndex = position355, tokenIndex355
			return false
		},
		/* 52 BasionymAuthorshipYearMisformed <- <('(' _? AuthorsGroup _? ')' (_? ',')? _? Year)> */
		func() bool {
			position357, tokenIndex357 := position, tokenIndex
			{
				position358 := position
				if buffer[position] != rune('(') {
					goto l357
				}
				position++
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
				if !_rules[ruleAuthorsGroup]() {
					goto l357
				}
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
					goto l357
				}
				position++
				{
					position363, tokenIndex363 := position, tokenIndex
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
					if buffer[position] != rune(',') {
						goto l363
					}
					position++
					goto l364
				l363:
					position, tokenIndex = position363, tokenIndex363
				}
			l364:
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
				if !_rules[ruleYear]() {
					goto l357
				}
				add(ruleBasionymAuthorshipYearMisformed, position358)
			}
			return true
		l357:
			position, tokenIndex = position357, tokenIndex357
			return false
		},
		/* 53 BasionymAuthorship <- <(BasionymAuthorship1 / BasionymAuthorship2)> */
		func() bool {
			position369, tokenIndex369 := position, tokenIndex
			{
				position370 := position
				{
					position371, tokenIndex371 := position, tokenIndex
					if !_rules[ruleBasionymAuthorship1]() {
						goto l372
					}
					goto l371
				l372:
					position, tokenIndex = position371, tokenIndex371
					if !_rules[ruleBasionymAuthorship2]() {
						goto l369
					}
				}
			l371:
				add(ruleBasionymAuthorship, position370)
			}
			return true
		l369:
			position, tokenIndex = position369, tokenIndex369
			return false
		},
		/* 54 BasionymAuthorship1 <- <('(' _? AuthorsGroup _? ')')> */
		func() bool {
			position373, tokenIndex373 := position, tokenIndex
			{
				position374 := position
				if buffer[position] != rune('(') {
					goto l373
				}
				position++
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
				if !_rules[ruleAuthorsGroup]() {
					goto l373
				}
				{
					position377, tokenIndex377 := position, tokenIndex
					if !_rules[rule_]() {
						goto l377
					}
					goto l378
				l377:
					position, tokenIndex = position377, tokenIndex377
				}
			l378:
				if buffer[position] != rune(')') {
					goto l373
				}
				position++
				add(ruleBasionymAuthorship1, position374)
			}
			return true
		l373:
			position, tokenIndex = position373, tokenIndex373
			return false
		},
		/* 55 BasionymAuthorship2 <- <('(' _? '(' _? AuthorsGroup _? ')' _? ')')> */
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
				if buffer[position] != rune('(') {
					goto l379
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
					goto l379
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
					goto l379
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
				if buffer[position] != rune(')') {
					goto l379
				}
				position++
				add(ruleBasionymAuthorship2, position380)
			}
			return true
		l379:
			position, tokenIndex = position379, tokenIndex379
			return false
		},
		/* 56 AuthorsGroup <- <(AuthorsTeam (_? AuthorEmend? AuthorEx? AuthorsTeam)?)> */
		func() bool {
			position389, tokenIndex389 := position, tokenIndex
			{
				position390 := position
				if !_rules[ruleAuthorsTeam]() {
					goto l389
				}
				{
					position391, tokenIndex391 := position, tokenIndex
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
					{
						position395, tokenIndex395 := position, tokenIndex
						if !_rules[ruleAuthorEmend]() {
							goto l395
						}
						goto l396
					l395:
						position, tokenIndex = position395, tokenIndex395
					}
				l396:
					{
						position397, tokenIndex397 := position, tokenIndex
						if !_rules[ruleAuthorEx]() {
							goto l397
						}
						goto l398
					l397:
						position, tokenIndex = position397, tokenIndex397
					}
				l398:
					if !_rules[ruleAuthorsTeam]() {
						goto l391
					}
					goto l392
				l391:
					position, tokenIndex = position391, tokenIndex391
				}
			l392:
				add(ruleAuthorsGroup, position390)
			}
			return true
		l389:
			position, tokenIndex = position389, tokenIndex389
			return false
		},
		/* 57 AuthorsTeam <- <(Author (AuthorSep Author)* (_? ','? _? Year)?)> */
		func() bool {
			position399, tokenIndex399 := position, tokenIndex
			{
				position400 := position
				if !_rules[ruleAuthor]() {
					goto l399
				}
			l401:
				{
					position402, tokenIndex402 := position, tokenIndex
					if !_rules[ruleAuthorSep]() {
						goto l402
					}
					if !_rules[ruleAuthor]() {
						goto l402
					}
					goto l401
				l402:
					position, tokenIndex = position402, tokenIndex402
				}
				{
					position403, tokenIndex403 := position, tokenIndex
					{
						position405, tokenIndex405 := position, tokenIndex
						if !_rules[rule_]() {
							goto l405
						}
						goto l406
					l405:
						position, tokenIndex = position405, tokenIndex405
					}
				l406:
					{
						position407, tokenIndex407 := position, tokenIndex
						if buffer[position] != rune(',') {
							goto l407
						}
						position++
						goto l408
					l407:
						position, tokenIndex = position407, tokenIndex407
					}
				l408:
					{
						position409, tokenIndex409 := position, tokenIndex
						if !_rules[rule_]() {
							goto l409
						}
						goto l410
					l409:
						position, tokenIndex = position409, tokenIndex409
					}
				l410:
					if !_rules[ruleYear]() {
						goto l403
					}
					goto l404
				l403:
					position, tokenIndex = position403, tokenIndex403
				}
			l404:
				add(ruleAuthorsTeam, position400)
			}
			return true
		l399:
			position, tokenIndex = position399, tokenIndex399
			return false
		},
		/* 58 AuthorSep <- <(AuthorSep1 / AuthorSep2)> */
		func() bool {
			position411, tokenIndex411 := position, tokenIndex
			{
				position412 := position
				{
					position413, tokenIndex413 := position, tokenIndex
					if !_rules[ruleAuthorSep1]() {
						goto l414
					}
					goto l413
				l414:
					position, tokenIndex = position413, tokenIndex413
					if !_rules[ruleAuthorSep2]() {
						goto l411
					}
				}
			l413:
				add(ruleAuthorSep, position412)
			}
			return true
		l411:
			position, tokenIndex = position411, tokenIndex411
			return false
		},
		/* 59 AuthorSep1 <- <(_? (',' _)? ('&' / ('e' 't') / ('a' 'n' 'd') / ('a' 'p' 'u' 'd')) _?)> */
		func() bool {
			position415, tokenIndex415 := position, tokenIndex
			{
				position416 := position
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
				{
					position419, tokenIndex419 := position, tokenIndex
					if buffer[position] != rune(',') {
						goto l419
					}
					position++
					if !_rules[rule_]() {
						goto l419
					}
					goto l420
				l419:
					position, tokenIndex = position419, tokenIndex419
				}
			l420:
				{
					position421, tokenIndex421 := position, tokenIndex
					if buffer[position] != rune('&') {
						goto l422
					}
					position++
					goto l421
				l422:
					position, tokenIndex = position421, tokenIndex421
					if buffer[position] != rune('e') {
						goto l423
					}
					position++
					if buffer[position] != rune('t') {
						goto l423
					}
					position++
					goto l421
				l423:
					position, tokenIndex = position421, tokenIndex421
					if buffer[position] != rune('a') {
						goto l424
					}
					position++
					if buffer[position] != rune('n') {
						goto l424
					}
					position++
					if buffer[position] != rune('d') {
						goto l424
					}
					position++
					goto l421
				l424:
					position, tokenIndex = position421, tokenIndex421
					if buffer[position] != rune('a') {
						goto l415
					}
					position++
					if buffer[position] != rune('p') {
						goto l415
					}
					position++
					if buffer[position] != rune('u') {
						goto l415
					}
					position++
					if buffer[position] != rune('d') {
						goto l415
					}
					position++
				}
			l421:
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
				add(ruleAuthorSep1, position416)
			}
			return true
		l415:
			position, tokenIndex = position415, tokenIndex415
			return false
		},
		/* 60 AuthorSep2 <- <(_? ',' _?)> */
		func() bool {
			position427, tokenIndex427 := position, tokenIndex
			{
				position428 := position
				{
					position429, tokenIndex429 := position, tokenIndex
					if !_rules[rule_]() {
						goto l429
					}
					goto l430
				l429:
					position, tokenIndex = position429, tokenIndex429
				}
			l430:
				if buffer[position] != rune(',') {
					goto l427
				}
				position++
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
				add(ruleAuthorSep2, position428)
			}
			return true
		l427:
			position, tokenIndex = position427, tokenIndex427
			return false
		},
		/* 61 AuthorEx <- <((('e' 'x' '.'?) / ('i' 'n')) _)> */
		func() bool {
			position433, tokenIndex433 := position, tokenIndex
			{
				position434 := position
				{
					position435, tokenIndex435 := position, tokenIndex
					if buffer[position] != rune('e') {
						goto l436
					}
					position++
					if buffer[position] != rune('x') {
						goto l436
					}
					position++
					{
						position437, tokenIndex437 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l437
						}
						position++
						goto l438
					l437:
						position, tokenIndex = position437, tokenIndex437
					}
				l438:
					goto l435
				l436:
					position, tokenIndex = position435, tokenIndex435
					if buffer[position] != rune('i') {
						goto l433
					}
					position++
					if buffer[position] != rune('n') {
						goto l433
					}
					position++
				}
			l435:
				if !_rules[rule_]() {
					goto l433
				}
				add(ruleAuthorEx, position434)
			}
			return true
		l433:
			position, tokenIndex = position433, tokenIndex433
			return false
		},
		/* 62 AuthorEmend <- <('e' 'm' 'e' 'n' 'd' '.'? _)> */
		func() bool {
			position439, tokenIndex439 := position, tokenIndex
			{
				position440 := position
				if buffer[position] != rune('e') {
					goto l439
				}
				position++
				if buffer[position] != rune('m') {
					goto l439
				}
				position++
				if buffer[position] != rune('e') {
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
				{
					position441, tokenIndex441 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l441
					}
					position++
					goto l442
				l441:
					position, tokenIndex = position441, tokenIndex441
				}
			l442:
				if !_rules[rule_]() {
					goto l439
				}
				add(ruleAuthorEmend, position440)
			}
			return true
		l439:
			position, tokenIndex = position439, tokenIndex439
			return false
		},
		/* 63 Author <- <(Author1 / Author2 / UnknownAuthor)> */
		func() bool {
			position443, tokenIndex443 := position, tokenIndex
			{
				position444 := position
				{
					position445, tokenIndex445 := position, tokenIndex
					if !_rules[ruleAuthor1]() {
						goto l446
					}
					goto l445
				l446:
					position, tokenIndex = position445, tokenIndex445
					if !_rules[ruleAuthor2]() {
						goto l447
					}
					goto l445
				l447:
					position, tokenIndex = position445, tokenIndex445
					if !_rules[ruleUnknownAuthor]() {
						goto l443
					}
				}
			l445:
				add(ruleAuthor, position444)
			}
			return true
		l443:
			position, tokenIndex = position443, tokenIndex443
			return false
		},
		/* 64 Author1 <- <(Author2 _? Filius)> */
		func() bool {
			position448, tokenIndex448 := position, tokenIndex
			{
				position449 := position
				if !_rules[ruleAuthor2]() {
					goto l448
				}
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
				if !_rules[ruleFilius]() {
					goto l448
				}
				add(ruleAuthor1, position449)
			}
			return true
		l448:
			position, tokenIndex = position448, tokenIndex448
			return false
		},
		/* 65 Author2 <- <(AuthorWord (_? AuthorWord)*)> */
		func() bool {
			position452, tokenIndex452 := position, tokenIndex
			{
				position453 := position
				if !_rules[ruleAuthorWord]() {
					goto l452
				}
			l454:
				{
					position455, tokenIndex455 := position, tokenIndex
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
					if !_rules[ruleAuthorWord]() {
						goto l455
					}
					goto l454
				l455:
					position, tokenIndex = position455, tokenIndex455
				}
				add(ruleAuthor2, position453)
			}
			return true
		l452:
			position, tokenIndex = position452, tokenIndex452
			return false
		},
		/* 66 UnknownAuthor <- <('?' / ((('a' 'u' 'c' 't') / ('a' 'n' 'o' 'n')) (&SpaceCharEOI / '.')))> */
		func() bool {
			position458, tokenIndex458 := position, tokenIndex
			{
				position459 := position
				{
					position460, tokenIndex460 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l461
					}
					position++
					goto l460
				l461:
					position, tokenIndex = position460, tokenIndex460
					{
						position462, tokenIndex462 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l463
						}
						position++
						if buffer[position] != rune('u') {
							goto l463
						}
						position++
						if buffer[position] != rune('c') {
							goto l463
						}
						position++
						if buffer[position] != rune('t') {
							goto l463
						}
						position++
						goto l462
					l463:
						position, tokenIndex = position462, tokenIndex462
						if buffer[position] != rune('a') {
							goto l458
						}
						position++
						if buffer[position] != rune('n') {
							goto l458
						}
						position++
						if buffer[position] != rune('o') {
							goto l458
						}
						position++
						if buffer[position] != rune('n') {
							goto l458
						}
						position++
					}
				l462:
					{
						position464, tokenIndex464 := position, tokenIndex
						{
							position466, tokenIndex466 := position, tokenIndex
							if !_rules[ruleSpaceCharEOI]() {
								goto l465
							}
							position, tokenIndex = position466, tokenIndex466
						}
						goto l464
					l465:
						position, tokenIndex = position464, tokenIndex464
						if buffer[position] != rune('.') {
							goto l458
						}
						position++
					}
				l464:
				}
			l460:
				add(ruleUnknownAuthor, position459)
			}
			return true
		l458:
			position, tokenIndex = position458, tokenIndex458
			return false
		},
		/* 67 AuthorWord <- <(AuthorWord1 / AuthorWord2 / AuthorWord3 / AuthorPrefix)> */
		func() bool {
			position467, tokenIndex467 := position, tokenIndex
			{
				position468 := position
				{
					position469, tokenIndex469 := position, tokenIndex
					if !_rules[ruleAuthorWord1]() {
						goto l470
					}
					goto l469
				l470:
					position, tokenIndex = position469, tokenIndex469
					if !_rules[ruleAuthorWord2]() {
						goto l471
					}
					goto l469
				l471:
					position, tokenIndex = position469, tokenIndex469
					if !_rules[ruleAuthorWord3]() {
						goto l472
					}
					goto l469
				l472:
					position, tokenIndex = position469, tokenIndex469
					if !_rules[ruleAuthorPrefix]() {
						goto l467
					}
				}
			l469:
				add(ruleAuthorWord, position468)
			}
			return true
		l467:
			position, tokenIndex = position467, tokenIndex467
			return false
		},
		/* 68 AuthorWord1 <- <(('a' 'r' 'g' '.') / ('e' 't' ' ' 'a' 'l' '.' '{' '?' '}') / ((('e' 't') / '&') (' ' 'a' 'l') '.'?))> */
		func() bool {
			position473, tokenIndex473 := position, tokenIndex
			{
				position474 := position
				{
					position475, tokenIndex475 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l476
					}
					position++
					if buffer[position] != rune('r') {
						goto l476
					}
					position++
					if buffer[position] != rune('g') {
						goto l476
					}
					position++
					if buffer[position] != rune('.') {
						goto l476
					}
					position++
					goto l475
				l476:
					position, tokenIndex = position475, tokenIndex475
					if buffer[position] != rune('e') {
						goto l477
					}
					position++
					if buffer[position] != rune('t') {
						goto l477
					}
					position++
					if buffer[position] != rune(' ') {
						goto l477
					}
					position++
					if buffer[position] != rune('a') {
						goto l477
					}
					position++
					if buffer[position] != rune('l') {
						goto l477
					}
					position++
					if buffer[position] != rune('.') {
						goto l477
					}
					position++
					if buffer[position] != rune('{') {
						goto l477
					}
					position++
					if buffer[position] != rune('?') {
						goto l477
					}
					position++
					if buffer[position] != rune('}') {
						goto l477
					}
					position++
					goto l475
				l477:
					position, tokenIndex = position475, tokenIndex475
					{
						position478, tokenIndex478 := position, tokenIndex
						if buffer[position] != rune('e') {
							goto l479
						}
						position++
						if buffer[position] != rune('t') {
							goto l479
						}
						position++
						goto l478
					l479:
						position, tokenIndex = position478, tokenIndex478
						if buffer[position] != rune('&') {
							goto l473
						}
						position++
					}
				l478:
					if buffer[position] != rune(' ') {
						goto l473
					}
					position++
					if buffer[position] != rune('a') {
						goto l473
					}
					position++
					if buffer[position] != rune('l') {
						goto l473
					}
					position++
					{
						position480, tokenIndex480 := position, tokenIndex
						if buffer[position] != rune('.') {
							goto l480
						}
						position++
						goto l481
					l480:
						position, tokenIndex = position480, tokenIndex480
					}
				l481:
				}
			l475:
				add(ruleAuthorWord1, position474)
			}
			return true
		l473:
			position, tokenIndex = position473, tokenIndex473
			return false
		},
		/* 69 AuthorWord2 <- <(AuthorWord3 dash AuthorWordSoft)> */
		func() bool {
			position482, tokenIndex482 := position, tokenIndex
			{
				position483 := position
				if !_rules[ruleAuthorWord3]() {
					goto l482
				}
				if !_rules[ruledash]() {
					goto l482
				}
				if !_rules[ruleAuthorWordSoft]() {
					goto l482
				}
				add(ruleAuthorWord2, position483)
			}
			return true
		l482:
			position, tokenIndex = position482, tokenIndex482
			return false
		},
		/* 70 AuthorWord3 <- <(AuthorPrefixGlued? (AllCapsAuthorWord / CapAuthorWord) '.'?)> */
		func() bool {
			position484, tokenIndex484 := position, tokenIndex
			{
				position485 := position
				{
					position486, tokenIndex486 := position, tokenIndex
					if !_rules[ruleAuthorPrefixGlued]() {
						goto l486
					}
					goto l487
				l486:
					position, tokenIndex = position486, tokenIndex486
				}
			l487:
				{
					position488, tokenIndex488 := position, tokenIndex
					if !_rules[ruleAllCapsAuthorWord]() {
						goto l489
					}
					goto l488
				l489:
					position, tokenIndex = position488, tokenIndex488
					if !_rules[ruleCapAuthorWord]() {
						goto l484
					}
				}
			l488:
				{
					position490, tokenIndex490 := position, tokenIndex
					if buffer[position] != rune('.') {
						goto l490
					}
					position++
					goto l491
				l490:
					position, tokenIndex = position490, tokenIndex490
				}
			l491:
				add(ruleAuthorWord3, position485)
			}
			return true
		l484:
			position, tokenIndex = position484, tokenIndex484
			return false
		},
		/* 71 AuthorWordSoft <- <(((AuthorUpperChar (AuthorUpperChar+ / AuthorLowerChar+)) / AuthorLowerChar+) '.'?)> */
		func() bool {
			position492, tokenIndex492 := position, tokenIndex
			{
				position493 := position
				{
					position494, tokenIndex494 := position, tokenIndex
					if !_rules[ruleAuthorUpperChar]() {
						goto l495
					}
					{
						position496, tokenIndex496 := position, tokenIndex
						if !_rules[ruleAuthorUpperChar]() {
							goto l497
						}
					l498:
						{
							position499, tokenIndex499 := position, tokenIndex
							if !_rules[ruleAuthorUpperChar]() {
								goto l499
							}
							goto l498
						l499:
							position, tokenIndex = position499, tokenIndex499
						}
						goto l496
					l497:
						position, tokenIndex = position496, tokenIndex496
						if !_rules[ruleAuthorLowerChar]() {
							goto l495
						}
					l500:
						{
							position501, tokenIndex501 := position, tokenIndex
							if !_rules[ruleAuthorLowerChar]() {
								goto l501
							}
							goto l500
						l501:
							position, tokenIndex = position501, tokenIndex501
						}
					}
				l496:
					goto l494
				l495:
					position, tokenIndex = position494, tokenIndex494
					if !_rules[ruleAuthorLowerChar]() {
						goto l492
					}
				l502:
					{
						position503, tokenIndex503 := position, tokenIndex
						if !_rules[ruleAuthorLowerChar]() {
							goto l503
						}
						goto l502
					l503:
						position, tokenIndex = position503, tokenIndex503
					}
				}
			l494:
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
				add(ruleAuthorWordSoft, position493)
			}
			return true
		l492:
			position, tokenIndex = position492, tokenIndex492
			return false
		},
		/* 72 CapAuthorWord <- <(AuthorUpperChar AuthorLowerChar*)> */
		func() bool {
			position506, tokenIndex506 := position, tokenIndex
			{
				position507 := position
				if !_rules[ruleAuthorUpperChar]() {
					goto l506
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
				add(ruleCapAuthorWord, position507)
			}
			return true
		l506:
			position, tokenIndex = position506, tokenIndex506
			return false
		},
		/* 73 AllCapsAuthorWord <- <(AuthorUpperChar AuthorUpperChar+)> */
		func() bool {
			position510, tokenIndex510 := position, tokenIndex
			{
				position511 := position
				if !_rules[ruleAuthorUpperChar]() {
					goto l510
				}
				if !_rules[ruleAuthorUpperChar]() {
					goto l510
				}
			l512:
				{
					position513, tokenIndex513 := position, tokenIndex
					if !_rules[ruleAuthorUpperChar]() {
						goto l513
					}
					goto l512
				l513:
					position, tokenIndex = position513, tokenIndex513
				}
				add(ruleAllCapsAuthorWord, position511)
			}
			return true
		l510:
			position, tokenIndex = position510, tokenIndex510
			return false
		},
		/* 74 Filius <- <(('f' '.') / ('f' 'i' 'l' '.') / ('f' 'i' 'l' 'i' 'u' 's'))> */
		func() bool {
			position514, tokenIndex514 := position, tokenIndex
			{
				position515 := position
				{
					position516, tokenIndex516 := position, tokenIndex
					if buffer[position] != rune('f') {
						goto l517
					}
					position++
					if buffer[position] != rune('.') {
						goto l517
					}
					position++
					goto l516
				l517:
					position, tokenIndex = position516, tokenIndex516
					if buffer[position] != rune('f') {
						goto l518
					}
					position++
					if buffer[position] != rune('i') {
						goto l518
					}
					position++
					if buffer[position] != rune('l') {
						goto l518
					}
					position++
					if buffer[position] != rune('.') {
						goto l518
					}
					position++
					goto l516
				l518:
					position, tokenIndex = position516, tokenIndex516
					if buffer[position] != rune('f') {
						goto l514
					}
					position++
					if buffer[position] != rune('i') {
						goto l514
					}
					position++
					if buffer[position] != rune('l') {
						goto l514
					}
					position++
					if buffer[position] != rune('i') {
						goto l514
					}
					position++
					if buffer[position] != rune('u') {
						goto l514
					}
					position++
					if buffer[position] != rune('s') {
						goto l514
					}
					position++
				}
			l516:
				add(ruleFilius, position515)
			}
			return true
		l514:
			position, tokenIndex = position514, tokenIndex514
			return false
		},
		/* 75 AuthorPrefixGlued <- <(('d' '\'') / ('O' '\''))> */
		func() bool {
			position519, tokenIndex519 := position, tokenIndex
			{
				position520 := position
				{
					position521, tokenIndex521 := position, tokenIndex
					if buffer[position] != rune('d') {
						goto l522
					}
					position++
					if buffer[position] != rune('\'') {
						goto l522
					}
					position++
					goto l521
				l522:
					position, tokenIndex = position521, tokenIndex521
					if buffer[position] != rune('O') {
						goto l519
					}
					position++
					if buffer[position] != rune('\'') {
						goto l519
					}
					position++
				}
			l521:
				add(ruleAuthorPrefixGlued, position520)
			}
			return true
		l519:
			position, tokenIndex = position519, tokenIndex519
			return false
		},
		/* 76 AuthorPrefix <- <(AuthorPrefix1 / AuthorPrefix2)> */
		func() bool {
			position523, tokenIndex523 := position, tokenIndex
			{
				position524 := position
				{
					position525, tokenIndex525 := position, tokenIndex
					if !_rules[ruleAuthorPrefix1]() {
						goto l526
					}
					goto l525
				l526:
					position, tokenIndex = position525, tokenIndex525
					if !_rules[ruleAuthorPrefix2]() {
						goto l523
					}
				}
			l525:
				add(ruleAuthorPrefix, position524)
			}
			return true
		l523:
			position, tokenIndex = position523, tokenIndex523
			return false
		},
		/* 77 AuthorPrefix2 <- <(('v' '.' (_? ('d' '.'))?) / ('\'' 't'))> */
		func() bool {
			position527, tokenIndex527 := position, tokenIndex
			{
				position528 := position
				{
					position529, tokenIndex529 := position, tokenIndex
					if buffer[position] != rune('v') {
						goto l530
					}
					position++
					if buffer[position] != rune('.') {
						goto l530
					}
					position++
					{
						position531, tokenIndex531 := position, tokenIndex
						{
							position533, tokenIndex533 := position, tokenIndex
							if !_rules[rule_]() {
								goto l533
							}
							goto l534
						l533:
							position, tokenIndex = position533, tokenIndex533
						}
					l534:
						if buffer[position] != rune('d') {
							goto l531
						}
						position++
						if buffer[position] != rune('.') {
							goto l531
						}
						position++
						goto l532
					l531:
						position, tokenIndex = position531, tokenIndex531
					}
				l532:
					goto l529
				l530:
					position, tokenIndex = position529, tokenIndex529
					if buffer[position] != rune('\'') {
						goto l527
					}
					position++
					if buffer[position] != rune('t') {
						goto l527
					}
					position++
				}
			l529:
				add(ruleAuthorPrefix2, position528)
			}
			return true
		l527:
			position, tokenIndex = position527, tokenIndex527
			return false
		},
		/* 78 AuthorPrefix1 <- <((('a' 'b') / ('a' 'f') / ('b' 'i' 's') / ('d' 'a') / ('d' 'e' 'r') / ('d' 'e' 's') / ('d' 'e' 'n') / ('d' 'e' 'l') / ('d' 'e' 'l' 'l' 'a') / ('d' 'e' 'l' 'a') / ('d' 'e') / ('d' 'i') / ('d' 'u') / ('e' 'l') / ('l' 'a') / ('l' 'e') / ('t' 'e' 'r') / ('v' 'a' 'n') / ('d' '\'') / ('i' 'n' '\'' 't') / ('z' 'u' 'r') / ('v' 'o' 'n' (_ (('d' '.') / ('d' 'e' 'm')))?) / ('v' (_ 'd')?)) &_)> */
		func() bool {
			position535, tokenIndex535 := position, tokenIndex
			{
				position536 := position
				{
					position537, tokenIndex537 := position, tokenIndex
					if buffer[position] != rune('a') {
						goto l538
					}
					position++
					if buffer[position] != rune('b') {
						goto l538
					}
					position++
					goto l537
				l538:
					position, tokenIndex = position537, tokenIndex537
					if buffer[position] != rune('a') {
						goto l539
					}
					position++
					if buffer[position] != rune('f') {
						goto l539
					}
					position++
					goto l537
				l539:
					position, tokenIndex = position537, tokenIndex537
					if buffer[position] != rune('b') {
						goto l540
					}
					position++
					if buffer[position] != rune('i') {
						goto l540
					}
					position++
					if buffer[position] != rune('s') {
						goto l540
					}
					position++
					goto l537
				l540:
					position, tokenIndex = position537, tokenIndex537
					if buffer[position] != rune('d') {
						goto l541
					}
					position++
					if buffer[position] != rune('a') {
						goto l541
					}
					position++
					goto l537
				l541:
					position, tokenIndex = position537, tokenIndex537
					if buffer[position] != rune('d') {
						goto l542
					}
					position++
					if buffer[position] != rune('e') {
						goto l542
					}
					position++
					if buffer[position] != rune('r') {
						goto l542
					}
					position++
					goto l537
				l542:
					position, tokenIndex = position537, tokenIndex537
					if buffer[position] != rune('d') {
						goto l543
					}
					position++
					if buffer[position] != rune('e') {
						goto l543
					}
					position++
					if buffer[position] != rune('s') {
						goto l543
					}
					position++
					goto l537
				l543:
					position, tokenIndex = position537, tokenIndex537
					if buffer[position] != rune('d') {
						goto l544
					}
					position++
					if buffer[position] != rune('e') {
						goto l544
					}
					position++
					if buffer[position] != rune('n') {
						goto l544
					}
					position++
					goto l537
				l544:
					position, tokenIndex = position537, tokenIndex537
					if buffer[position] != rune('d') {
						goto l545
					}
					position++
					if buffer[position] != rune('e') {
						goto l545
					}
					position++
					if buffer[position] != rune('l') {
						goto l545
					}
					position++
					goto l537
				l545:
					position, tokenIndex = position537, tokenIndex537
					if buffer[position] != rune('d') {
						goto l546
					}
					position++
					if buffer[position] != rune('e') {
						goto l546
					}
					position++
					if buffer[position] != rune('l') {
						goto l546
					}
					position++
					if buffer[position] != rune('l') {
						goto l546
					}
					position++
					if buffer[position] != rune('a') {
						goto l546
					}
					position++
					goto l537
				l546:
					position, tokenIndex = position537, tokenIndex537
					if buffer[position] != rune('d') {
						goto l547
					}
					position++
					if buffer[position] != rune('e') {
						goto l547
					}
					position++
					if buffer[position] != rune('l') {
						goto l547
					}
					position++
					if buffer[position] != rune('a') {
						goto l547
					}
					position++
					goto l537
				l547:
					position, tokenIndex = position537, tokenIndex537
					if buffer[position] != rune('d') {
						goto l548
					}
					position++
					if buffer[position] != rune('e') {
						goto l548
					}
					position++
					goto l537
				l548:
					position, tokenIndex = position537, tokenIndex537
					if buffer[position] != rune('d') {
						goto l549
					}
					position++
					if buffer[position] != rune('i') {
						goto l549
					}
					position++
					goto l537
				l549:
					position, tokenIndex = position537, tokenIndex537
					if buffer[position] != rune('d') {
						goto l550
					}
					position++
					if buffer[position] != rune('u') {
						goto l550
					}
					position++
					goto l537
				l550:
					position, tokenIndex = position537, tokenIndex537
					if buffer[position] != rune('e') {
						goto l551
					}
					position++
					if buffer[position] != rune('l') {
						goto l551
					}
					position++
					goto l537
				l551:
					position, tokenIndex = position537, tokenIndex537
					if buffer[position] != rune('l') {
						goto l552
					}
					position++
					if buffer[position] != rune('a') {
						goto l552
					}
					position++
					goto l537
				l552:
					position, tokenIndex = position537, tokenIndex537
					if buffer[position] != rune('l') {
						goto l553
					}
					position++
					if buffer[position] != rune('e') {
						goto l553
					}
					position++
					goto l537
				l553:
					position, tokenIndex = position537, tokenIndex537
					if buffer[position] != rune('t') {
						goto l554
					}
					position++
					if buffer[position] != rune('e') {
						goto l554
					}
					position++
					if buffer[position] != rune('r') {
						goto l554
					}
					position++
					goto l537
				l554:
					position, tokenIndex = position537, tokenIndex537
					if buffer[position] != rune('v') {
						goto l555
					}
					position++
					if buffer[position] != rune('a') {
						goto l555
					}
					position++
					if buffer[position] != rune('n') {
						goto l555
					}
					position++
					goto l537
				l555:
					position, tokenIndex = position537, tokenIndex537
					if buffer[position] != rune('d') {
						goto l556
					}
					position++
					if buffer[position] != rune('\'') {
						goto l556
					}
					position++
					goto l537
				l556:
					position, tokenIndex = position537, tokenIndex537
					if buffer[position] != rune('i') {
						goto l557
					}
					position++
					if buffer[position] != rune('n') {
						goto l557
					}
					position++
					if buffer[position] != rune('\'') {
						goto l557
					}
					position++
					if buffer[position] != rune('t') {
						goto l557
					}
					position++
					goto l537
				l557:
					position, tokenIndex = position537, tokenIndex537
					if buffer[position] != rune('z') {
						goto l558
					}
					position++
					if buffer[position] != rune('u') {
						goto l558
					}
					position++
					if buffer[position] != rune('r') {
						goto l558
					}
					position++
					goto l537
				l558:
					position, tokenIndex = position537, tokenIndex537
					if buffer[position] != rune('v') {
						goto l559
					}
					position++
					if buffer[position] != rune('o') {
						goto l559
					}
					position++
					if buffer[position] != rune('n') {
						goto l559
					}
					position++
					{
						position560, tokenIndex560 := position, tokenIndex
						if !_rules[rule_]() {
							goto l560
						}
						{
							position562, tokenIndex562 := position, tokenIndex
							if buffer[position] != rune('d') {
								goto l563
							}
							position++
							if buffer[position] != rune('.') {
								goto l563
							}
							position++
							goto l562
						l563:
							position, tokenIndex = position562, tokenIndex562
							if buffer[position] != rune('d') {
								goto l560
							}
							position++
							if buffer[position] != rune('e') {
								goto l560
							}
							position++
							if buffer[position] != rune('m') {
								goto l560
							}
							position++
						}
					l562:
						goto l561
					l560:
						position, tokenIndex = position560, tokenIndex560
					}
				l561:
					goto l537
				l559:
					position, tokenIndex = position537, tokenIndex537
					if buffer[position] != rune('v') {
						goto l535
					}
					position++
					{
						position564, tokenIndex564 := position, tokenIndex
						if !_rules[rule_]() {
							goto l564
						}
						if buffer[position] != rune('d') {
							goto l564
						}
						position++
						goto l565
					l564:
						position, tokenIndex = position564, tokenIndex564
					}
				l565:
				}
			l537:
				{
					position566, tokenIndex566 := position, tokenIndex
					if !_rules[rule_]() {
						goto l535
					}
					position, tokenIndex = position566, tokenIndex566
				}
				add(ruleAuthorPrefix1, position536)
			}
			return true
		l535:
			position, tokenIndex = position535, tokenIndex535
			return false
		},
		/* 79 AuthorUpperChar <- <(hASCII / ('À' / 'Á' / 'Â' / 'Ã' / 'Ä' / 'Å' / 'Æ' / 'Ç' / 'È' / 'É' / 'Ê' / 'Ë' / 'Ì' / 'Í' / 'Î' / 'Ï' / 'Ð' / 'Ñ' / 'Ò' / 'Ó' / 'Ô' / 'Õ' / 'Ö' / 'Ø' / 'Ù' / 'Ú' / 'Û' / 'Ü' / 'Ý' / 'Ć' / 'Č' / 'Ď' / 'İ' / 'Ķ' / 'Ĺ' / 'ĺ' / 'Ľ' / 'ľ' / 'Ł' / 'ł' / 'Ņ' / 'Ō' / 'Ő' / 'Œ' / 'Ř' / 'Ś' / 'Ŝ' / 'Ş' / 'Š' / 'Ÿ' / 'Ź' / 'Ż' / 'Ž' / 'ƒ' / 'Ǿ' / 'Ș' / 'Ț'))> */
		func() bool {
			position567, tokenIndex567 := position, tokenIndex
			{
				position568 := position
				{
					position569, tokenIndex569 := position, tokenIndex
					if !_rules[rulehASCII]() {
						goto l570
					}
					goto l569
				l570:
					position, tokenIndex = position569, tokenIndex569
					{
						position571, tokenIndex571 := position, tokenIndex
						if buffer[position] != rune('À') {
							goto l572
						}
						position++
						goto l571
					l572:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Á') {
							goto l573
						}
						position++
						goto l571
					l573:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Â') {
							goto l574
						}
						position++
						goto l571
					l574:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ã') {
							goto l575
						}
						position++
						goto l571
					l575:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ä') {
							goto l576
						}
						position++
						goto l571
					l576:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Å') {
							goto l577
						}
						position++
						goto l571
					l577:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Æ') {
							goto l578
						}
						position++
						goto l571
					l578:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ç') {
							goto l579
						}
						position++
						goto l571
					l579:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('È') {
							goto l580
						}
						position++
						goto l571
					l580:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('É') {
							goto l581
						}
						position++
						goto l571
					l581:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ê') {
							goto l582
						}
						position++
						goto l571
					l582:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ë') {
							goto l583
						}
						position++
						goto l571
					l583:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ì') {
							goto l584
						}
						position++
						goto l571
					l584:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Í') {
							goto l585
						}
						position++
						goto l571
					l585:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Î') {
							goto l586
						}
						position++
						goto l571
					l586:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ï') {
							goto l587
						}
						position++
						goto l571
					l587:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ð') {
							goto l588
						}
						position++
						goto l571
					l588:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ñ') {
							goto l589
						}
						position++
						goto l571
					l589:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ò') {
							goto l590
						}
						position++
						goto l571
					l590:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ó') {
							goto l591
						}
						position++
						goto l571
					l591:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ô') {
							goto l592
						}
						position++
						goto l571
					l592:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Õ') {
							goto l593
						}
						position++
						goto l571
					l593:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ö') {
							goto l594
						}
						position++
						goto l571
					l594:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ø') {
							goto l595
						}
						position++
						goto l571
					l595:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ù') {
							goto l596
						}
						position++
						goto l571
					l596:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ú') {
							goto l597
						}
						position++
						goto l571
					l597:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Û') {
							goto l598
						}
						position++
						goto l571
					l598:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ü') {
							goto l599
						}
						position++
						goto l571
					l599:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ý') {
							goto l600
						}
						position++
						goto l571
					l600:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ć') {
							goto l601
						}
						position++
						goto l571
					l601:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Č') {
							goto l602
						}
						position++
						goto l571
					l602:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ď') {
							goto l603
						}
						position++
						goto l571
					l603:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('İ') {
							goto l604
						}
						position++
						goto l571
					l604:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ķ') {
							goto l605
						}
						position++
						goto l571
					l605:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ĺ') {
							goto l606
						}
						position++
						goto l571
					l606:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('ĺ') {
							goto l607
						}
						position++
						goto l571
					l607:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ľ') {
							goto l608
						}
						position++
						goto l571
					l608:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('ľ') {
							goto l609
						}
						position++
						goto l571
					l609:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ł') {
							goto l610
						}
						position++
						goto l571
					l610:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('ł') {
							goto l611
						}
						position++
						goto l571
					l611:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ņ') {
							goto l612
						}
						position++
						goto l571
					l612:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ō') {
							goto l613
						}
						position++
						goto l571
					l613:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ő') {
							goto l614
						}
						position++
						goto l571
					l614:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Œ') {
							goto l615
						}
						position++
						goto l571
					l615:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ř') {
							goto l616
						}
						position++
						goto l571
					l616:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ś') {
							goto l617
						}
						position++
						goto l571
					l617:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ŝ') {
							goto l618
						}
						position++
						goto l571
					l618:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ş') {
							goto l619
						}
						position++
						goto l571
					l619:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Š') {
							goto l620
						}
						position++
						goto l571
					l620:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ÿ') {
							goto l621
						}
						position++
						goto l571
					l621:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ź') {
							goto l622
						}
						position++
						goto l571
					l622:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ż') {
							goto l623
						}
						position++
						goto l571
					l623:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ž') {
							goto l624
						}
						position++
						goto l571
					l624:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('ƒ') {
							goto l625
						}
						position++
						goto l571
					l625:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ǿ') {
							goto l626
						}
						position++
						goto l571
					l626:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ș') {
							goto l627
						}
						position++
						goto l571
					l627:
						position, tokenIndex = position571, tokenIndex571
						if buffer[position] != rune('Ț') {
							goto l567
						}
						position++
					}
				l571:
				}
			l569:
				add(ruleAuthorUpperChar, position568)
			}
			return true
		l567:
			position, tokenIndex = position567, tokenIndex567
			return false
		},
		/* 80 AuthorLowerChar <- <(lASCII / ('à' / 'á' / 'â' / 'ã' / 'ä' / 'å' / 'æ' / 'ç' / 'è' / 'é' / 'ê' / 'ë' / 'ì' / 'í' / 'î' / 'ï' / 'ð' / 'ñ' / 'ò' / 'ó' / 'ó' / 'ô' / 'õ' / 'ö' / 'ø' / 'ù' / 'ú' / 'û' / 'ü' / 'ý' / 'ÿ' / 'ā' / 'ă' / 'ą' / 'ć' / 'ĉ' / 'č' / 'ď' / 'đ' / '\'' / 'ē' / 'ĕ' / 'ė' / 'ę' / 'ě' / 'ğ' / 'ī' / 'ĭ' / 'İ' / 'ı' / 'ĺ' / 'ľ' / 'ł' / 'ń' / 'ņ' / 'ň' / 'ŏ' / 'ő' / 'œ' / 'ŕ' / 'ř' / 'ś' / 'ş' / 'š' / 'ţ' / 'ť' / 'ũ' / 'ū' / 'ŭ' / 'ů' / 'ű' / 'ź' / 'ż' / 'ž' / 'ſ' / 'ǎ' / 'ǔ' / 'ǧ' / 'ș' / 'ț' / 'ȳ' / 'ß'))> */
		func() bool {
			position628, tokenIndex628 := position, tokenIndex
			{
				position629 := position
				{
					position630, tokenIndex630 := position, tokenIndex
					if !_rules[rulelASCII]() {
						goto l631
					}
					goto l630
				l631:
					position, tokenIndex = position630, tokenIndex630
					{
						position632, tokenIndex632 := position, tokenIndex
						if buffer[position] != rune('à') {
							goto l633
						}
						position++
						goto l632
					l633:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('á') {
							goto l634
						}
						position++
						goto l632
					l634:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('â') {
							goto l635
						}
						position++
						goto l632
					l635:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ã') {
							goto l636
						}
						position++
						goto l632
					l636:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ä') {
							goto l637
						}
						position++
						goto l632
					l637:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('å') {
							goto l638
						}
						position++
						goto l632
					l638:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('æ') {
							goto l639
						}
						position++
						goto l632
					l639:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ç') {
							goto l640
						}
						position++
						goto l632
					l640:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('è') {
							goto l641
						}
						position++
						goto l632
					l641:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('é') {
							goto l642
						}
						position++
						goto l632
					l642:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ê') {
							goto l643
						}
						position++
						goto l632
					l643:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ë') {
							goto l644
						}
						position++
						goto l632
					l644:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ì') {
							goto l645
						}
						position++
						goto l632
					l645:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('í') {
							goto l646
						}
						position++
						goto l632
					l646:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('î') {
							goto l647
						}
						position++
						goto l632
					l647:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ï') {
							goto l648
						}
						position++
						goto l632
					l648:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ð') {
							goto l649
						}
						position++
						goto l632
					l649:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ñ') {
							goto l650
						}
						position++
						goto l632
					l650:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ò') {
							goto l651
						}
						position++
						goto l632
					l651:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ó') {
							goto l652
						}
						position++
						goto l632
					l652:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ó') {
							goto l653
						}
						position++
						goto l632
					l653:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ô') {
							goto l654
						}
						position++
						goto l632
					l654:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('õ') {
							goto l655
						}
						position++
						goto l632
					l655:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ö') {
							goto l656
						}
						position++
						goto l632
					l656:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ø') {
							goto l657
						}
						position++
						goto l632
					l657:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ù') {
							goto l658
						}
						position++
						goto l632
					l658:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ú') {
							goto l659
						}
						position++
						goto l632
					l659:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('û') {
							goto l660
						}
						position++
						goto l632
					l660:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ü') {
							goto l661
						}
						position++
						goto l632
					l661:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ý') {
							goto l662
						}
						position++
						goto l632
					l662:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ÿ') {
							goto l663
						}
						position++
						goto l632
					l663:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ā') {
							goto l664
						}
						position++
						goto l632
					l664:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ă') {
							goto l665
						}
						position++
						goto l632
					l665:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ą') {
							goto l666
						}
						position++
						goto l632
					l666:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ć') {
							goto l667
						}
						position++
						goto l632
					l667:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ĉ') {
							goto l668
						}
						position++
						goto l632
					l668:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('č') {
							goto l669
						}
						position++
						goto l632
					l669:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ď') {
							goto l670
						}
						position++
						goto l632
					l670:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('đ') {
							goto l671
						}
						position++
						goto l632
					l671:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('\'') {
							goto l672
						}
						position++
						goto l632
					l672:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ē') {
							goto l673
						}
						position++
						goto l632
					l673:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ĕ') {
							goto l674
						}
						position++
						goto l632
					l674:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ė') {
							goto l675
						}
						position++
						goto l632
					l675:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ę') {
							goto l676
						}
						position++
						goto l632
					l676:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ě') {
							goto l677
						}
						position++
						goto l632
					l677:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ğ') {
							goto l678
						}
						position++
						goto l632
					l678:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ī') {
							goto l679
						}
						position++
						goto l632
					l679:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ĭ') {
							goto l680
						}
						position++
						goto l632
					l680:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('İ') {
							goto l681
						}
						position++
						goto l632
					l681:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ı') {
							goto l682
						}
						position++
						goto l632
					l682:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ĺ') {
							goto l683
						}
						position++
						goto l632
					l683:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ľ') {
							goto l684
						}
						position++
						goto l632
					l684:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ł') {
							goto l685
						}
						position++
						goto l632
					l685:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ń') {
							goto l686
						}
						position++
						goto l632
					l686:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ņ') {
							goto l687
						}
						position++
						goto l632
					l687:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ň') {
							goto l688
						}
						position++
						goto l632
					l688:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ŏ') {
							goto l689
						}
						position++
						goto l632
					l689:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ő') {
							goto l690
						}
						position++
						goto l632
					l690:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('œ') {
							goto l691
						}
						position++
						goto l632
					l691:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ŕ') {
							goto l692
						}
						position++
						goto l632
					l692:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ř') {
							goto l693
						}
						position++
						goto l632
					l693:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ś') {
							goto l694
						}
						position++
						goto l632
					l694:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ş') {
							goto l695
						}
						position++
						goto l632
					l695:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('š') {
							goto l696
						}
						position++
						goto l632
					l696:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ţ') {
							goto l697
						}
						position++
						goto l632
					l697:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ť') {
							goto l698
						}
						position++
						goto l632
					l698:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ũ') {
							goto l699
						}
						position++
						goto l632
					l699:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ū') {
							goto l700
						}
						position++
						goto l632
					l700:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ŭ') {
							goto l701
						}
						position++
						goto l632
					l701:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ů') {
							goto l702
						}
						position++
						goto l632
					l702:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ű') {
							goto l703
						}
						position++
						goto l632
					l703:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ź') {
							goto l704
						}
						position++
						goto l632
					l704:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ż') {
							goto l705
						}
						position++
						goto l632
					l705:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ž') {
							goto l706
						}
						position++
						goto l632
					l706:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ſ') {
							goto l707
						}
						position++
						goto l632
					l707:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ǎ') {
							goto l708
						}
						position++
						goto l632
					l708:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ǔ') {
							goto l709
						}
						position++
						goto l632
					l709:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ǧ') {
							goto l710
						}
						position++
						goto l632
					l710:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ș') {
							goto l711
						}
						position++
						goto l632
					l711:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ț') {
							goto l712
						}
						position++
						goto l632
					l712:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ȳ') {
							goto l713
						}
						position++
						goto l632
					l713:
						position, tokenIndex = position632, tokenIndex632
						if buffer[position] != rune('ß') {
							goto l628
						}
						position++
					}
				l632:
				}
			l630:
				add(ruleAuthorLowerChar, position629)
			}
			return true
		l628:
			position, tokenIndex = position628, tokenIndex628
			return false
		},
		/* 81 Year <- <(YearRange / YearApprox / YearWithParens / YearWithPage / YearWithDot / YearWithChar / YearNum)> */
		func() bool {
			position714, tokenIndex714 := position, tokenIndex
			{
				position715 := position
				{
					position716, tokenIndex716 := position, tokenIndex
					if !_rules[ruleYearRange]() {
						goto l717
					}
					goto l716
				l717:
					position, tokenIndex = position716, tokenIndex716
					if !_rules[ruleYearApprox]() {
						goto l718
					}
					goto l716
				l718:
					position, tokenIndex = position716, tokenIndex716
					if !_rules[ruleYearWithParens]() {
						goto l719
					}
					goto l716
				l719:
					position, tokenIndex = position716, tokenIndex716
					if !_rules[ruleYearWithPage]() {
						goto l720
					}
					goto l716
				l720:
					position, tokenIndex = position716, tokenIndex716
					if !_rules[ruleYearWithDot]() {
						goto l721
					}
					goto l716
				l721:
					position, tokenIndex = position716, tokenIndex716
					if !_rules[ruleYearWithChar]() {
						goto l722
					}
					goto l716
				l722:
					position, tokenIndex = position716, tokenIndex716
					if !_rules[ruleYearNum]() {
						goto l714
					}
				}
			l716:
				add(ruleYear, position715)
			}
			return true
		l714:
			position, tokenIndex = position714, tokenIndex714
			return false
		},
		/* 82 YearRange <- <(YearNum dash (nums+ ('a' / 'b' / 'c' / 'd' / 'e' / 'f' / 'g' / 'h' / 'i' / 'j' / 'k' / 'l' / 'm' / 'n' / 'o' / 'p' / 'q' / 'r' / 's' / 't' / 'u' / 'v' / 'w' / 'x' / 'y' / 'z' / '?')*))> */
		func() bool {
			position723, tokenIndex723 := position, tokenIndex
			{
				position724 := position
				if !_rules[ruleYearNum]() {
					goto l723
				}
				if !_rules[ruledash]() {
					goto l723
				}
				if !_rules[rulenums]() {
					goto l723
				}
			l725:
				{
					position726, tokenIndex726 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l726
					}
					goto l725
				l726:
					position, tokenIndex = position726, tokenIndex726
				}
			l727:
				{
					position728, tokenIndex728 := position, tokenIndex
					{
						position729, tokenIndex729 := position, tokenIndex
						if buffer[position] != rune('a') {
							goto l730
						}
						position++
						goto l729
					l730:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('b') {
							goto l731
						}
						position++
						goto l729
					l731:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('c') {
							goto l732
						}
						position++
						goto l729
					l732:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('d') {
							goto l733
						}
						position++
						goto l729
					l733:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('e') {
							goto l734
						}
						position++
						goto l729
					l734:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('f') {
							goto l735
						}
						position++
						goto l729
					l735:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('g') {
							goto l736
						}
						position++
						goto l729
					l736:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('h') {
							goto l737
						}
						position++
						goto l729
					l737:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('i') {
							goto l738
						}
						position++
						goto l729
					l738:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('j') {
							goto l739
						}
						position++
						goto l729
					l739:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('k') {
							goto l740
						}
						position++
						goto l729
					l740:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('l') {
							goto l741
						}
						position++
						goto l729
					l741:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('m') {
							goto l742
						}
						position++
						goto l729
					l742:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('n') {
							goto l743
						}
						position++
						goto l729
					l743:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('o') {
							goto l744
						}
						position++
						goto l729
					l744:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('p') {
							goto l745
						}
						position++
						goto l729
					l745:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('q') {
							goto l746
						}
						position++
						goto l729
					l746:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('r') {
							goto l747
						}
						position++
						goto l729
					l747:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('s') {
							goto l748
						}
						position++
						goto l729
					l748:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('t') {
							goto l749
						}
						position++
						goto l729
					l749:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('u') {
							goto l750
						}
						position++
						goto l729
					l750:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('v') {
							goto l751
						}
						position++
						goto l729
					l751:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('w') {
							goto l752
						}
						position++
						goto l729
					l752:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('x') {
							goto l753
						}
						position++
						goto l729
					l753:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('y') {
							goto l754
						}
						position++
						goto l729
					l754:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('z') {
							goto l755
						}
						position++
						goto l729
					l755:
						position, tokenIndex = position729, tokenIndex729
						if buffer[position] != rune('?') {
							goto l728
						}
						position++
					}
				l729:
					goto l727
				l728:
					position, tokenIndex = position728, tokenIndex728
				}
				add(ruleYearRange, position724)
			}
			return true
		l723:
			position, tokenIndex = position723, tokenIndex723
			return false
		},
		/* 83 YearWithDot <- <(YearNum '.')> */
		func() bool {
			position756, tokenIndex756 := position, tokenIndex
			{
				position757 := position
				if !_rules[ruleYearNum]() {
					goto l756
				}
				if buffer[position] != rune('.') {
					goto l756
				}
				position++
				add(ruleYearWithDot, position757)
			}
			return true
		l756:
			position, tokenIndex = position756, tokenIndex756
			return false
		},
		/* 84 YearApprox <- <('[' _? YearNum _? ']')> */
		func() bool {
			position758, tokenIndex758 := position, tokenIndex
			{
				position759 := position
				if buffer[position] != rune('[') {
					goto l758
				}
				position++
				{
					position760, tokenIndex760 := position, tokenIndex
					if !_rules[rule_]() {
						goto l760
					}
					goto l761
				l760:
					position, tokenIndex = position760, tokenIndex760
				}
			l761:
				if !_rules[ruleYearNum]() {
					goto l758
				}
				{
					position762, tokenIndex762 := position, tokenIndex
					if !_rules[rule_]() {
						goto l762
					}
					goto l763
				l762:
					position, tokenIndex = position762, tokenIndex762
				}
			l763:
				if buffer[position] != rune(']') {
					goto l758
				}
				position++
				add(ruleYearApprox, position759)
			}
			return true
		l758:
			position, tokenIndex = position758, tokenIndex758
			return false
		},
		/* 85 YearWithPage <- <((YearWithChar / YearNum) _? ':' _? nums+)> */
		func() bool {
			position764, tokenIndex764 := position, tokenIndex
			{
				position765 := position
				{
					position766, tokenIndex766 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l767
					}
					goto l766
				l767:
					position, tokenIndex = position766, tokenIndex766
					if !_rules[ruleYearNum]() {
						goto l764
					}
				}
			l766:
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
				if buffer[position] != rune(':') {
					goto l764
				}
				position++
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
				if !_rules[rulenums]() {
					goto l764
				}
			l772:
				{
					position773, tokenIndex773 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l773
					}
					goto l772
				l773:
					position, tokenIndex = position773, tokenIndex773
				}
				add(ruleYearWithPage, position765)
			}
			return true
		l764:
			position, tokenIndex = position764, tokenIndex764
			return false
		},
		/* 86 YearWithParens <- <('(' (YearWithChar / YearNum) ')')> */
		func() bool {
			position774, tokenIndex774 := position, tokenIndex
			{
				position775 := position
				if buffer[position] != rune('(') {
					goto l774
				}
				position++
				{
					position776, tokenIndex776 := position, tokenIndex
					if !_rules[ruleYearWithChar]() {
						goto l777
					}
					goto l776
				l777:
					position, tokenIndex = position776, tokenIndex776
					if !_rules[ruleYearNum]() {
						goto l774
					}
				}
			l776:
				if buffer[position] != rune(')') {
					goto l774
				}
				position++
				add(ruleYearWithParens, position775)
			}
			return true
		l774:
			position, tokenIndex = position774, tokenIndex774
			return false
		},
		/* 87 YearWithChar <- <(YearNum lASCII Action0)> */
		func() bool {
			position778, tokenIndex778 := position, tokenIndex
			{
				position779 := position
				if !_rules[ruleYearNum]() {
					goto l778
				}
				if !_rules[rulelASCII]() {
					goto l778
				}
				if !_rules[ruleAction0]() {
					goto l778
				}
				add(ruleYearWithChar, position779)
			}
			return true
		l778:
			position, tokenIndex = position778, tokenIndex778
			return false
		},
		/* 88 YearNum <- <(('1' / '2') ('0' / '7' / '8' / '9') nums (nums / '?') '?'*)> */
		func() bool {
			position780, tokenIndex780 := position, tokenIndex
			{
				position781 := position
				{
					position782, tokenIndex782 := position, tokenIndex
					if buffer[position] != rune('1') {
						goto l783
					}
					position++
					goto l782
				l783:
					position, tokenIndex = position782, tokenIndex782
					if buffer[position] != rune('2') {
						goto l780
					}
					position++
				}
			l782:
				{
					position784, tokenIndex784 := position, tokenIndex
					if buffer[position] != rune('0') {
						goto l785
					}
					position++
					goto l784
				l785:
					position, tokenIndex = position784, tokenIndex784
					if buffer[position] != rune('7') {
						goto l786
					}
					position++
					goto l784
				l786:
					position, tokenIndex = position784, tokenIndex784
					if buffer[position] != rune('8') {
						goto l787
					}
					position++
					goto l784
				l787:
					position, tokenIndex = position784, tokenIndex784
					if buffer[position] != rune('9') {
						goto l780
					}
					position++
				}
			l784:
				if !_rules[rulenums]() {
					goto l780
				}
				{
					position788, tokenIndex788 := position, tokenIndex
					if !_rules[rulenums]() {
						goto l789
					}
					goto l788
				l789:
					position, tokenIndex = position788, tokenIndex788
					if buffer[position] != rune('?') {
						goto l780
					}
					position++
				}
			l788:
			l790:
				{
					position791, tokenIndex791 := position, tokenIndex
					if buffer[position] != rune('?') {
						goto l791
					}
					position++
					goto l790
				l791:
					position, tokenIndex = position791, tokenIndex791
				}
				add(ruleYearNum, position781)
			}
			return true
		l780:
			position, tokenIndex = position780, tokenIndex780
			return false
		},
		/* 89 NameUpperChar <- <(UpperChar / UpperCharExtended)> */
		func() bool {
			position792, tokenIndex792 := position, tokenIndex
			{
				position793 := position
				{
					position794, tokenIndex794 := position, tokenIndex
					if !_rules[ruleUpperChar]() {
						goto l795
					}
					goto l794
				l795:
					position, tokenIndex = position794, tokenIndex794
					if !_rules[ruleUpperCharExtended]() {
						goto l792
					}
				}
			l794:
				add(ruleNameUpperChar, position793)
			}
			return true
		l792:
			position, tokenIndex = position792, tokenIndex792
			return false
		},
		/* 90 UpperCharExtended <- <('Æ' / 'Œ' / 'Ö')> */
		func() bool {
			position796, tokenIndex796 := position, tokenIndex
			{
				position797 := position
				{
					position798, tokenIndex798 := position, tokenIndex
					if buffer[position] != rune('Æ') {
						goto l799
					}
					position++
					goto l798
				l799:
					position, tokenIndex = position798, tokenIndex798
					if buffer[position] != rune('Œ') {
						goto l800
					}
					position++
					goto l798
				l800:
					position, tokenIndex = position798, tokenIndex798
					if buffer[position] != rune('Ö') {
						goto l796
					}
					position++
				}
			l798:
				add(ruleUpperCharExtended, position797)
			}
			return true
		l796:
			position, tokenIndex = position796, tokenIndex796
			return false
		},
		/* 91 UpperChar <- <hASCII> */
		func() bool {
			position801, tokenIndex801 := position, tokenIndex
			{
				position802 := position
				if !_rules[rulehASCII]() {
					goto l801
				}
				add(ruleUpperChar, position802)
			}
			return true
		l801:
			position, tokenIndex = position801, tokenIndex801
			return false
		},
		/* 92 NameLowerChar <- <(LowerChar / LowerCharExtended / MiscodedChar)> */
		func() bool {
			position803, tokenIndex803 := position, tokenIndex
			{
				position804 := position
				{
					position805, tokenIndex805 := position, tokenIndex
					if !_rules[ruleLowerChar]() {
						goto l806
					}
					goto l805
				l806:
					position, tokenIndex = position805, tokenIndex805
					if !_rules[ruleLowerCharExtended]() {
						goto l807
					}
					goto l805
				l807:
					position, tokenIndex = position805, tokenIndex805
					if !_rules[ruleMiscodedChar]() {
						goto l803
					}
				}
			l805:
				add(ruleNameLowerChar, position804)
			}
			return true
		l803:
			position, tokenIndex = position803, tokenIndex803
			return false
		},
		/* 93 MiscodedChar <- <'�'> */
		func() bool {
			position808, tokenIndex808 := position, tokenIndex
			{
				position809 := position
				if buffer[position] != rune('�') {
					goto l808
				}
				position++
				add(ruleMiscodedChar, position809)
			}
			return true
		l808:
			position, tokenIndex = position808, tokenIndex808
			return false
		},
		/* 94 LowerCharExtended <- <('æ' / 'œ' / 'ſ' / 'à' / 'â' / 'å' / 'ã' / 'ä' / 'á' / 'ç' / 'č' / 'é' / 'è' / 'ë' / 'í' / 'ì' / 'ï' / 'ň' / 'ñ' / 'ñ' / 'ó' / 'ò' / 'ô' / 'ø' / 'õ' / 'ö' / 'ú' / 'ù' / 'ü' / 'ŕ' / 'ř' / 'ŗ' / 'š' / 'š' / 'ş' / 'ž')> */
		func() bool {
			position810, tokenIndex810 := position, tokenIndex
			{
				position811 := position
				{
					position812, tokenIndex812 := position, tokenIndex
					if buffer[position] != rune('æ') {
						goto l813
					}
					position++
					goto l812
				l813:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('œ') {
						goto l814
					}
					position++
					goto l812
				l814:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('ſ') {
						goto l815
					}
					position++
					goto l812
				l815:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('à') {
						goto l816
					}
					position++
					goto l812
				l816:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('â') {
						goto l817
					}
					position++
					goto l812
				l817:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('å') {
						goto l818
					}
					position++
					goto l812
				l818:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('ã') {
						goto l819
					}
					position++
					goto l812
				l819:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('ä') {
						goto l820
					}
					position++
					goto l812
				l820:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('á') {
						goto l821
					}
					position++
					goto l812
				l821:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('ç') {
						goto l822
					}
					position++
					goto l812
				l822:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('č') {
						goto l823
					}
					position++
					goto l812
				l823:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('é') {
						goto l824
					}
					position++
					goto l812
				l824:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('è') {
						goto l825
					}
					position++
					goto l812
				l825:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('ë') {
						goto l826
					}
					position++
					goto l812
				l826:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('í') {
						goto l827
					}
					position++
					goto l812
				l827:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('ì') {
						goto l828
					}
					position++
					goto l812
				l828:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('ï') {
						goto l829
					}
					position++
					goto l812
				l829:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('ň') {
						goto l830
					}
					position++
					goto l812
				l830:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('ñ') {
						goto l831
					}
					position++
					goto l812
				l831:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('ñ') {
						goto l832
					}
					position++
					goto l812
				l832:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('ó') {
						goto l833
					}
					position++
					goto l812
				l833:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('ò') {
						goto l834
					}
					position++
					goto l812
				l834:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('ô') {
						goto l835
					}
					position++
					goto l812
				l835:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('ø') {
						goto l836
					}
					position++
					goto l812
				l836:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('õ') {
						goto l837
					}
					position++
					goto l812
				l837:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('ö') {
						goto l838
					}
					position++
					goto l812
				l838:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('ú') {
						goto l839
					}
					position++
					goto l812
				l839:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('ù') {
						goto l840
					}
					position++
					goto l812
				l840:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('ü') {
						goto l841
					}
					position++
					goto l812
				l841:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('ŕ') {
						goto l842
					}
					position++
					goto l812
				l842:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('ř') {
						goto l843
					}
					position++
					goto l812
				l843:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('ŗ') {
						goto l844
					}
					position++
					goto l812
				l844:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('š') {
						goto l845
					}
					position++
					goto l812
				l845:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('š') {
						goto l846
					}
					position++
					goto l812
				l846:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('ş') {
						goto l847
					}
					position++
					goto l812
				l847:
					position, tokenIndex = position812, tokenIndex812
					if buffer[position] != rune('ž') {
						goto l810
					}
					position++
				}
			l812:
				add(ruleLowerCharExtended, position811)
			}
			return true
		l810:
			position, tokenIndex = position810, tokenIndex810
			return false
		},
		/* 95 LowerChar <- <lASCII> */
		func() bool {
			position848, tokenIndex848 := position, tokenIndex
			{
				position849 := position
				if !_rules[rulelASCII]() {
					goto l848
				}
				add(ruleLowerChar, position849)
			}
			return true
		l848:
			position, tokenIndex = position848, tokenIndex848
			return false
		},
		/* 96 SpaceCharEOI <- <(_ / !.)> */
		func() bool {
			position850, tokenIndex850 := position, tokenIndex
			{
				position851 := position
				{
					position852, tokenIndex852 := position, tokenIndex
					if !_rules[rule_]() {
						goto l853
					}
					goto l852
				l853:
					position, tokenIndex = position852, tokenIndex852
					{
						position854, tokenIndex854 := position, tokenIndex
						if !matchDot() {
							goto l854
						}
						goto l850
					l854:
						position, tokenIndex = position854, tokenIndex854
					}
				}
			l852:
				add(ruleSpaceCharEOI, position851)
			}
			return true
		l850:
			position, tokenIndex = position850, tokenIndex850
			return false
		},
		/* 97 WordBorderChar <- <(_ / (';' / '.' / ',' / ';' / '(' / ')'))> */
		nil,
		/* 98 nums <- <[0-9]> */
		func() bool {
			position856, tokenIndex856 := position, tokenIndex
			{
				position857 := position
				if c := buffer[position]; c < rune('0') || c > rune('9') {
					goto l856
				}
				position++
				add(rulenums, position857)
			}
			return true
		l856:
			position, tokenIndex = position856, tokenIndex856
			return false
		},
		/* 99 lASCII <- <[a-z]> */
		func() bool {
			position858, tokenIndex858 := position, tokenIndex
			{
				position859 := position
				if c := buffer[position]; c < rune('a') || c > rune('z') {
					goto l858
				}
				position++
				add(rulelASCII, position859)
			}
			return true
		l858:
			position, tokenIndex = position858, tokenIndex858
			return false
		},
		/* 100 hASCII <- <[A-Z]> */
		func() bool {
			position860, tokenIndex860 := position, tokenIndex
			{
				position861 := position
				if c := buffer[position]; c < rune('A') || c > rune('Z') {
					goto l860
				}
				position++
				add(rulehASCII, position861)
			}
			return true
		l860:
			position, tokenIndex = position860, tokenIndex860
			return false
		},
		/* 101 apostr <- <'\''> */
		func() bool {
			position862, tokenIndex862 := position, tokenIndex
			{
				position863 := position
				if buffer[position] != rune('\'') {
					goto l862
				}
				position++
				add(ruleapostr, position863)
			}
			return true
		l862:
			position, tokenIndex = position862, tokenIndex862
			return false
		},
		/* 102 dash <- <'-'> */
		func() bool {
			position864, tokenIndex864 := position, tokenIndex
			{
				position865 := position
				if buffer[position] != rune('-') {
					goto l864
				}
				position++
				add(ruledash, position865)
			}
			return true
		l864:
			position, tokenIndex = position864, tokenIndex864
			return false
		},
		/* 103 _ <- <(MultipleSpace / SingleSpace)> */
		func() bool {
			position866, tokenIndex866 := position, tokenIndex
			{
				position867 := position
				{
					position868, tokenIndex868 := position, tokenIndex
					if !_rules[ruleMultipleSpace]() {
						goto l869
					}
					goto l868
				l869:
					position, tokenIndex = position868, tokenIndex868
					if !_rules[ruleSingleSpace]() {
						goto l866
					}
				}
			l868:
				add(rule_, position867)
			}
			return true
		l866:
			position, tokenIndex = position866, tokenIndex866
			return false
		},
		/* 104 MultipleSpace <- <(SingleSpace SingleSpace+)> */
		func() bool {
			position870, tokenIndex870 := position, tokenIndex
			{
				position871 := position
				if !_rules[ruleSingleSpace]() {
					goto l870
				}
				if !_rules[ruleSingleSpace]() {
					goto l870
				}
			l872:
				{
					position873, tokenIndex873 := position, tokenIndex
					if !_rules[ruleSingleSpace]() {
						goto l873
					}
					goto l872
				l873:
					position, tokenIndex = position873, tokenIndex873
				}
				add(ruleMultipleSpace, position871)
			}
			return true
		l870:
			position, tokenIndex = position870, tokenIndex870
			return false
		},
		/* 105 SingleSpace <- <' '> */
		func() bool {
			position874, tokenIndex874 := position, tokenIndex
			{
				position875 := position
				if buffer[position] != rune(' ') {
					goto l874
				}
				position++
				add(ruleSingleSpace, position875)
			}
			return true
		l874:
			position, tokenIndex = position874, tokenIndex874
			return false
		},
		/* 107 Action0 <- <{ p.addWarn(YearCharWarn) }> */
		func() bool {
			{
				add(ruleAction0, position)
			}
			return true
		},
	}
	p.rules = _rules
}
