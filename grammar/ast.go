package grammar

import (
	"io"
	"strings"

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
	ruleOriginalAuthorship:              empty,
	ruleCombinationAuthorship:           empty,
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
	ruleAllCapsAuthorWord:               empty,
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
	root     *node32
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

func (p *Engine) OutputAST() {
	type element struct {
		node *node32
		down *element
	}
	var node *node32
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

func (p *Engine) PrintOutputSyntaxTree(w io.Writer) {
	p.root.print(w, true, p.Buffer)
}

func stackNodeIsWithin(n *node32, t token32) bool {
	return n.token32.begin >= t.begin && n.token32.end <= t.end
}

func (p *Engine) newNode(t token32) (*node32, bool) {
	var node *node32
	if _, ok := nodeRules[t.pegRule]; ok {
		node := &node32{token32: t}
		return node, false
	}
	switch t.pegRule {
	case ruleMultipleSpace:
		p.addWarn(SpaceMultipleWarn)
	}
	return node, true
}

func (p *Engine) nodeValue(n *node32) string {
	t := n.token32
	v := string([]rune(p.Buffer)[t.begin:t.end])
	return v
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
	var nameNodes []*node32
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
	var warnReset map[Warning]struct{}
	p.Warnings = warnReset

	sn := ScientificNameNode{
		Verbatim:   p.Buffer,
		VerbatimID: uuid5.UUID5(p.Buffer).String(),
		NamesGroup: ng,
		Tail:       tail,
		Warnings:   warns,
	}
	p.SN = &sn
}

func (p *Engine) tailValue(n *node32) string {
	t := n.token32
	if t.begin == t.end {
		return ""
	}
	p.addWarn(TailWarn)
	return string([]rune(p.Buffer)[t.begin:t.end])
}

func (p *Engine) newName(n *node32) Name {
	var name Name
	n = n.up
	switch n.token32.pegRule {
	case ruleUninomialCombo:
		p.addWarn(UninomialComboWarn)
		return p.newUninomialComboNode(n)
	case ruleUninomial:
		return p.newUninomialNode(n)
	case ruleNameSpecies:
		return p.newSpeciesNode(n)
	}
	return name
}

type speciesNode struct {
	Genus        *wordNode
	SubGenus     *wordNode
	Species      *spEpithetNode
	InfraSpecies []*infraspEpithetNode
}

func (p *Engine) newSpeciesNode(n *node32) *speciesNode {
	var sp *spEpithetNode
	var sg *wordNode
	var infs []*infraspEpithetNode
	n = n.up
	gen := p.newWordNode(n, GenusType)
	if n.up.token32.pegRule == ruleAbbrGenus {
		p.addWarn(GenusAbbrWarn)
	}
	n = n.next
	for n != nil {
		switch n.token32.pegRule {
		case ruleSubGenus:
			sg = p.newWordNode(n.up, SubGenusType)
		case ruleSpeciesEpithet:
			sp = p.newSpeciesEpithetNode(n)
		case ruleInfraspGroup:
			infs = p.newInfraspeciesGroup(n)
		}
		n = n.next
	}
	sn := speciesNode{
		Genus:        gen,
		SubGenus:     sg,
		Species:      sp,
		InfraSpecies: infs,
	}
	return &sn
}

type spEpithetNode struct {
	Word       *wordNode
	Authorship *authorshipNode
}

func (p *Engine) newSpeciesEpithetNode(n *node32) *spEpithetNode {
	var au *authorshipNode
	n = n.up
	se := p.newWordNode(n, SpEpithetType)
	n = n.next
	if n != nil {
		au = p.newAuthorshipNode(n)
	}
	sen := spEpithetNode{
		Word:       se,
		Authorship: au,
	}
	return &sen
}

type infraspEpithetNode struct {
	Word       *wordNode
	Rank       *rankNode
	Authorship *authorshipNode
}

func (p *Engine) newInfraspeciesGroup(n *node32) []*infraspEpithetNode {
	var infs []*infraspEpithetNode
	n = n.up
	if n == nil || n.token32.pegRule != ruleInfraspEpithet {
		return infs
	}
	for n != nil {
		inf := p.newInfraspEpithetNode(n)
		infs = append(infs, inf)
		n = n.next
	}
	return infs
}

func (p *Engine) newInfraspEpithetNode(n *node32) *infraspEpithetNode {
	var inf infraspEpithetNode
	var r *rankNode
	var w *wordNode
	var au *authorshipNode
	n = n.up
	if n == nil {
		return &inf
	}

	for n != nil {
		switch n.token32.pegRule {
		case ruleWord:
			w = p.newWordNode(n, InfraSpEpithetType)
		case ruleRank:
			r = p.newRankNode(n)
		case ruleAuthorship:
			au = p.newAuthorshipNode(n)
		}
		n = n.next
	}
	inf = infraspEpithetNode{
		Word:       w,
		Rank:       r,
		Authorship: au,
	}
	return &inf
}

type rankNode struct {
	Word *wordNode
}

func (p *Engine) newRankNode(n *node32) *rankNode {
	if n.up == nil {
		w := p.newWordNode(n, RankType)
		r := rankNode{Word: w}
		return &r
	}
	n = n.up
	w := p.newWordNode(n, RankType)
	switch n.token32.pegRule {
	case ruleRankForma:
		w.NormValue = "fm."
	case ruleRankVar:
		if w.Value[0] == 'n' {
			w.NormValue = "nvar."
		} else {
			w.NormValue = "var."
		}
	case ruleRankSsp:
		w.NormValue = "ssp."
	case ruleRankOtherUncommon:
		p.addWarn(RankUncommonWarn)
	}
	r := rankNode{Word: w}
	return &r
}

type uninomialNode struct {
	Word       *wordNode
	Authorship *authorshipNode
}

func (p *Engine) newUninomialNode(n *node32) *uninomialNode {
	var au *authorshipNode
	wn := n.up
	w := p.newWordNode(wn, UninomialType)
	if an := wn.next; an != nil {
		au = p.newAuthorshipNode(an)
	}
	un := uninomialNode{
		Word:       w,
		Authorship: au,
	}
	return &un
}

type uninomialComboNode struct {
	Uninomial1 *uninomialNode
	Uninomial2 *uninomialNode
	Rank       *rankUninomialNode
}

func (p *Engine) newUninomialComboNode(n *node32) *uninomialComboNode {
	var u1, u2 *uninomialNode
	var r *rankUninomialNode
	n = n.up
	switch n.token32.pegRule {
	case ruleUninomial:
		u1n := n
		u1 = p.newUninomialNode(u1n)
		rn := u1n.next
		r = p.newRankUninomialNode(rn)
		u2n := rn.next
		u2 = p.newUninomialNode(u2n)
	case ruleUninomialWord:
		uw := p.newWordNode(n, UninomialType)
		u1 = &uninomialNode{Word: uw}
		n = n.next
		u2w := p.newWordNode(n.up, UninomialType)
		n := n.next
		au2 := p.newAuthorshipNode(n)
		rw := &wordNode{Value: "subgen.", Pos: Pos{Type: RankUniType}}
		r = &rankUninomialNode{Word: rw}
		u2 = &uninomialNode{
			Word:       u2w,
			Authorship: au2,
		}
	}
	ucn := uninomialComboNode{
		Uninomial1: u1,
		Rank:       r,
		Uninomial2: u2,
	}
	return &ucn
}

type rankUninomialNode struct {
	Word *wordNode
}

func (p *Engine) newRankUninomialNode(n *node32) *rankUninomialNode {
	r := p.newWordNode(n, RankUniType)
	run := rankUninomialNode{Word: r}
	return &run
}

type authorshipNode struct {
	OriginalAuthors    *authorsGroupNode
	CombinationAuthors *authorsGroupNode
}

func (p *Engine) newAuthorshipNode(n *node32) *authorshipNode {
	var oa, ca *authorsGroupNode
	var parens bool
	n = n.up
	for n != nil {
		switch n.token32.pegRule {
		case ruleOriginalAuthorship:
			on := n.up
			if on.token32.pegRule == ruleBasionymAuthorshipYearMisformed {
				p.addWarn(AuthMisformedYearWarn)
				on = on.up
				parens = true
			} else if on.token32.pegRule == ruleBasionymAuthorship {
				on = on.up
				parens = true
			}
			oa = p.newAuthorsGroupNode(on)
			oa.Parens = parens
		case ruleCombinationAuthorship:
			ca = p.newAuthorsGroupNode(n.up)
		}
		n = n.next
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

func (p *Engine) newAuthorsGroupNode(n *node32) *authorsGroupNode {
	var t1, t2 *authorsTeamNode
	var t2t *wordNode
	n = n.up
	t1 = p.newAuthorTeam(n)
	ag := authorsGroupNode{
		Team1:     t1,
		Team2Type: t2t,
		Team2:     t2,
	}
	n = n.next
	if n == nil || n.token32.pegRule != ruleAuthorEx {
		return &ag
	}
	p.addWarn(AuthExWarn)
	t2t = p.newWordNode(n, AuthorWordType)
	t2t.NormValue = "ex"
	n = n.next
	if n == nil || n.token32.pegRule != ruleAuthorsTeam {
		return &ag
	}
	t2 = p.newAuthorTeam(n)
	ag.Team2Type = t2t
	ag.Team2 = t2
	return &ag
}

type authorsTeamNode struct {
	Authors []*authorNode
	Years   []*yearNode
}

func (p *Engine) newAuthorTeam(n *node32) *authorsTeamNode {
	var anodes []*node32
	var ynodes []*node32
	n = n.up
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

func (p *Engine) newAuthorNode(n *node32) *authorNode {
	var w *wordNode
	var ws []*wordNode
	val := ""
	n = n.up
	for n != nil {
		switch n.token32.pegRule {
		case ruleFilius:
			w = p.newWordNode(n, AuthorWordFiliusType)
			w.NormValue = "fil."
		default:
			w = p.authorWord(n)
		}
		ws = append(ws, w)
		val = str.JoinStrings(val, w.NormValue, " ")
		n = n.next
	}
	if len(val) < 2 {
		p.addWarn(AuthShortWarn)
	}
	au := authorNode{
		Value: val,
		Words: ws,
	}
	return &au
}

func (p *Engine) authorWord(n *node32) *wordNode {
	w := p.newWordNode(n, AuthorWordType)
	if n.up != nil && n.up.token32.pegRule == ruleAllCapsAuthorWord {
		nv := w.NormValue
		w.NormValue = string(nv[0]) + strings.ToLower(nv[1:len(nv)])
		p.addWarn(AuthUpperCaseWarn)
	}
	return w
}

type yearNode struct {
	Word        *wordNode
	Approximate bool
}

func (p *Engine) newYearNode(nd *node32) *yearNode {
	var w *wordNode
	appr := false
	nodes := nd.flatChildren()
	for _, v := range nodes {
		switch v.token32.pegRule {
		case ruleYearWithParens:
			p.addWarn(YearParensWarn)
			appr = true
		case ruleYearWithChar:
			p.addWarn(YearCharWarn)
			w = p.newWordNode(v, YearType)
			w.Value = w.Value[0 : len(w.Value)-1]
		case ruleYearNum:
			if w == nil {
				w = p.newWordNode(v, YearType)
			}
			if w.Value[len(w.Value)-1] == '?' {
				p.addWarn(YearQuestionWarn)
				appr = true
			}
		}
	}
	if w == nil {
		w = p.newWordNode(nd, YearType)
	}
	if appr {
		w.Pos.Type = YearApproximateType
	}
	yr := yearNode{
		Word:        w,
		Approximate: appr,
	}
	return &yr
}

func (n *node32) flatChildren() []*node32 {
	var ns []*node32
	if n.up == nil {
		return ns
	}
	n = n.up
	for n != nil {
		ns = append(ns, n)
		nn := n.next
		for nn != nil {
			ns = append(ns, nn)
			nn = nn.next
		}
		n = n.up
	}
	return ns
}

type wordNode struct {
	Value     string
	NormValue string
	Pos       Pos
}

func (p *Engine) newWordNode(n *node32, wt WordType) *wordNode {
	t := n.token32
	val := p.nodeValue(n)
	pos := Pos{Type: wt, Start: int(t.begin), End: int(t.end)}
	wrd := wordNode{Value: val, NormValue: val, Pos: pos}
	children := n.flatChildren()
	for _, v := range children {
		switch v.token32.pegRule {
		case ruleUpperCharExtended, ruleLowerCharExtended:
			p.addWarn(CharBadWarn)
			wrd.normalize()
		}
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
