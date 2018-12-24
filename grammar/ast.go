package grammar

import (
	"fmt"
	"io"
	"strconv"

	"github.com/gnames/uuid5"
	"gitlab.com/gogna/gnparser/str"
)

var empty = struct{}{}
var nodeRules = map[pegRule]struct{}{
	ruleSciName:                         empty,
	ruleTail:                            empty,
	ruleName:                            empty,
	ruleNameApprox:                      empty,
	ruleNameComp:                        empty,
	ruleNameSpecies:                     empty,
	ruleGenusWord:                       empty,
	ruleInfraspGroup:                    empty,
	ruleInfraspEpithet:                  empty,
	ruleSpeciesEpithet:                  empty,
	ruleComparison:                      empty,
	ruleRank:                            empty,
	ruleRankOtherUncommon:               empty,
	ruleRankVar:                         empty,
	ruleRankForma:                       empty,
	ruleRankSsp:                         empty,
	ruleSubGenusOrSuperspecies:          empty,
	ruleSubGenus:                        empty,
	ruleUninomialCombo:                  empty,
	ruleRankUninomial:                   empty,
	ruleUninomial:                       empty,
	ruleUninomialWord:                   empty,
	ruleAbbrGenus:                       empty,
	ruleWord:                            empty,
	ruleWord2StartDigit:                 empty,
	ruleHybridChar:                      empty,
	ruleApproxName:                      empty,
	ruleApproxNameIgnored:               empty,
	ruleApproximation:                   empty,
	ruleAuthorship:                      empty,
	ruleBasionymAuthorshipYearMisformed: empty,
	ruleBasionymAuthorship:              empty,
	ruleAuthorsGroup:                    empty,
	ruleAuthorsTeam:                     empty,
	ruleAuthorSep:                       empty,
	ruleAuthorEx:                        empty,
	ruleAuthorEmend:                     empty,
	ruleAuthor:                          empty,
	ruleUnknownAuthor:                   empty,
	ruleAuthorWord:                      empty,
	ruleFilius:                          empty,
	ruleAuthorPrefix:                    empty,
	ruleYear:                            empty,
	ruleYearRange:                       empty,
	ruleYearWithDot:                     empty,
	ruleYearApprox:                      empty,
	ruleYearWithPage:                    empty,
	ruleYearWithParens:                  empty,
	ruleYearWithChar:                    empty,
	ruleYearNum:                         empty,
	ruleUpperCharExtended:               empty,
	ruleMiscodedChar:                    empty,
	ruleLowerCharExtended:               empty,
}

type BaseEngine struct {
	SN       *ScientificNameNode
	root     *nodegn
	Warnings map[Warning]struct{}
}

func (p *Engine) addWarn(w Warning) {
	if p.Warnings == nil {
		p.Warnings = make(map[Warning]struct{})
	}
	if _, ok := p.Warnings[w]; !ok {
		p.Warnings[w] = struct{}{}
	}
}

type nodegn struct {
	token32  token32
	vals     map[string]string
	up, next *nodegn
}

func (node *nodegn) print(w io.Writer, buffer string) {
	var print func(node *nodegn, depth int)
	print = func(node *nodegn, depth int) {
		for node != nil {
			for c := 0; c < depth; c++ {
				fmt.Fprintf(w, "  ")
			}
			rule := rul3s[node.token32.pegRule]
			quote := strconv.Quote(
				string(([]rune(buffer)[node.token32.begin:node.token32.end])),
			)
			fmt.Fprintf(w, "%v %v\n", rule, quote)
			if node.up != nil {
				print(node.up, depth+1)
			}
			node = node.next
		}
	}
	print(node, 0)
}

func (p *Engine) ASTfactory() {
	type element struct {
		node *nodegn
		down *element
	}
	var node *nodegn
	var skip bool
	var stack *element
	for _, token := range p.Tokens() {
		if node, skip = p.newNode(token); skip {
			continue
		}
		for stack != nil && stackNodeIsWithin(stack.node, token) {
			stack.node.next = node.up
			node.up = stack.node
			stack = stack.down
		}
		stack = &element{node: node, down: stack}
	}
	if stack != nil {
		p.root = stack.node
	}
}

func (p *Engine) PrintAST(w io.Writer) {
	p.root.print(w, p.Buffer)
}

func stackNodeIsWithin(n *nodegn, t token32) bool {
	return n.token32.begin >= t.begin && n.token32.end <= t.end
}

func (p *Engine) newNode(t token32) (*nodegn, bool) {
	var node *nodegn
	if _, ok := nodeRules[t.pegRule]; ok {
		node := &nodegn{token32: t}
		return node, false
	}
	return node, true
}

type ScientificNameNode struct {
	Verbatim   string
	VerbatimID string
	NamesGroup []Name
	Tail       string
	Warnings   []Warning
}

func (p *Engine) NewScientificNameNode() {
	n := p.root.up
	var nameNodes []*nodegn
	var tail string

	for n != nil {
		switch n.token32.pegRule {
		case ruleName:
			nameNodes = append(nameNodes, n)
		case ruleTail:
			tail = p.tailValue(n)
		}
		n = n.next
	}

	ng := make([]Name, len(nameNodes))
	for i, v := range nameNodes {
		ng[i] = p.newName(v)
	}

	warns := make([]Warning, len(p.Warnings))
	i := 0
	for k, _ := range p.Warnings {
		warns[i] = k
		i++
	}

	sn := ScientificNameNode{
		Verbatim:   p.Buffer,
		VerbatimID: uuid5.UUID5(p.Buffer).String(),
		NamesGroup: ng,
		Tail:       tail,
		Warnings:   warns,
	}
	p.SN = &sn
}

func (p *Engine) tailValue(n *nodegn) string {
	t := n.token32
	if t.begin == t.end {
		return ""
	}
	p.addWarn(TailWarn)
	return string([]rune(p.Buffer)[t.begin:t.end])
}

func (p *Engine) newName(n *nodegn) Name {
	var name Name
	node := n.up
	switch node.token32.pegRule {
	case ruleUninomial:
		return p.newUninomialNode(node)
		// case ruleNameSpecies:
		// 	return p.newSpeciesNode(node)
	}
	return name
}

type uninomialNode struct {
	Word       *wordNode
	Authorship *authorshipNode
}

func (p *Engine) newUninomialNode(n *nodegn) *uninomialNode {
	var au *authorshipNode
	wn := n.up
	w := p.newWordNode(wn, UninomialType)
	if an := wn.next; an != nil {
		au = p.newAuthorship(an)
	}
	un := uninomialNode{
		Word:       w,
		Authorship: au,
	}
	return &un
}

type authorshipNode struct {
	OriginalAuthors    *authorsGroupNode
	CombinationAuthors *authorsGroupNode
}

func (p *Engine) newAuthorship(an *nodegn) *authorshipNode {
	var oa, ca *authorsGroupNode
	oan := an.up
	oa = p.newAuthorsGroupNode(oan)
	can := oan.next
	if can != nil {
		ca = p.newAuthorsGroupNode(can)
	}
	a := authorshipNode{
		OriginalAuthors:    oa,
		CombinationAuthors: ca,
	}
	return &a
}

type authorsGroupNode struct {
	Team1     *authorsTeamNode
	Team2Type *wordNode
	Team2     *authorsTeamNode
	Parens    bool
}

func (p *Engine) newAuthorsGroupNode(agn *nodegn) *authorsGroupNode {
	var t1, t2 *authorsTeamNode
	var t2t *wordNode
	t1n := agn.up
	t1 = p.newAuthorTeam(t1n)
	// TODO the rest
	ag := authorsGroupNode{
		Team1:     t1,
		Team2Type: t2t,
		Team2:     t2,
	}
	return &ag
}

type authorsTeamNode struct {
	Authors []*authorNode
	Years   []*yearNode
}

func (p *Engine) newAuthorTeam(at *nodegn) *authorsTeamNode {
	var anodes []*nodegn
	var ynodes []*nodegn
	n := at.up
	for n != nil {
		switch n.token32.pegRule {
		case ruleAuthor:
			anodes = append(anodes, n)
		case ruleYear:
			ynodes = append(ynodes, n)
		}
		n = n.next
	}
	aus := make([]*authorNode, len(anodes))
	yrs := make([]*yearNode, len(ynodes))

	for i, v := range anodes {
		aus[i] = p.newAuthorNode(v)
	}
	for i, v := range ynodes {
		yrs[i] = p.newYearNode(v)
	}
	atn := authorsTeamNode{
		Authors: aus,
		Years:   yrs,
	}
	return &atn
}

type authorSepNode struct {
	Value string
}

type authorNode struct {
	Value string
	Words []*wordNode
}

func (p *Engine) newAuthorNode(an *nodegn) *authorNode {
	var ws []*wordNode
	val := ""
	n := an.up
	for n != nil {
		w := p.newWordNode(n, AuthorWordType)
		ws = append(ws, w)
		val = str.JoinStrings(val, w.Value, " ")
		n = n.next
	}
	au := authorNode{
		Value: val,
		Words: ws,
	}
	return &au
}

type yearNode struct {
	Word        *wordNode
	Approximate bool
}

func (p *Engine) newYearNode(ngn *nodegn) *yearNode {
	w := p.newWordNode(ngn, YearType)
	yr := yearNode{
		Word: w,
	}
	return &yr
}

type wordNode struct {
	Value     string
	NormValue string
	Pos       Pos
}

func (p *Engine) newWordNode(n *nodegn, wt WordType) *wordNode {
	t := n.token32
	val := string([]rune(p.Buffer)[t.begin:t.end])
	pos := Pos{Type: wt, Start: int(t.begin), End: int(t.end)}
	wrd := wordNode{Value: val, NormValue: val, Pos: pos}
	up := n.up
	for up != nil {
		switch up.token32.pegRule {
		case ruleUpperCharExtended, ruleLowerCharExtended:
			p.addWarn(BadCharsWarn)
			wrd.normalize()
		}
		up = up.next
	}
	return &wrd
}

func (w *wordNode) normalize() error {
	if w == nil {
		return nil
	}
	nv, err := str.ToASCII(w.Value)
	if err != nil {
		return err
	}
	w.NormValue = nv
	return nil
}

type Pos struct {
	Type  WordType
	Start int
	End   int
}
