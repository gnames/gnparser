package grammar

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/gnames/gnparser/preprocess"

	"github.com/gnames/gnparser/dict"
	"github.com/gnames/gnparser/str"
	"github.com/gnames/uuid5"
)

type ScientificNameNode struct {
	Name
	Verbatim      string
	VerbatimID    string
	Cardinality   int
	Hybrid        bool
	Virus         bool
	Bacteria      bool
	Surrogate     bool
	Tail          string
	ParserVersion string
	Warnings      []Warning
}

func (p *Engine) NewScientificNameNode() {
	n := p.root.up
	var name Name
	var tail string

	for n != nil {
		switch n.token32.pegRule {
		case ruleName:
			name = p.newName(n)
		case ruleTail:
			tail = p.tailValue(n)
		}
		n = n.next
	}
	warns := make([]Warning, len(p.Warnings))
	i := 0
	for k := range p.Warnings {
		warns[i] = k
		i++
	}
	if str.IsBoldSurrogate(tail) {
		p.Cardinality = 0
		p.Surrogate = true
	}
	if p.Tail != "" && tail == "" {
		tail = p.Tail
	}
	sn := ScientificNameNode{
		Name:        name,
		Cardinality: p.Cardinality,
		Hybrid:      p.Hybrid,
		Surrogate:   p.Surrogate,
		Bacteria:    p.Bacteria,
		Tail:        tail,
		Warnings:    warns,
	}
	p.SN = &sn
}

func (p *Engine) NewNotParsedScientificNameNode(pp *preprocess.Preprocessor) {
	sn := &ScientificNameNode{Virus: pp.Virus}
	p.SN = sn
}

func (sn *ScientificNameNode) AddVerbatim(s string) {
	sn.Verbatim = s
	sn.VerbatimID = uuid5.UUID5(s).String()
}

func (p *Engine) tailValue(n *node32) string {
	t := n.token32
	if t.begin == t.end {
		return ""
	}
	p.AddWarn(TailWarn)
	return string([]rune(p.Buffer)[t.begin:t.end])
}

func (p *Engine) newName(n *node32) Name {
	var name Name
	n = n.up
	switch n.token32.pegRule {
	case ruleHybridFormula:
		name = p.newHybridFormulaNode(n)
	case ruleNamedGenusHybrid:
		name = p.newNamedGenusHybridNode(n)
	case ruleNamedSpeciesHybrid:
		name = p.newNamedSpeciesHybridNode(n)
	case ruleSingleName:
		name = p.newSingleName(n)
	}
	return name
}

type hybridFormulaNode struct {
	FirstSpecies   Name
	HybridElements []*hybridElement
}

type hybridElement struct {
	HybridChar *wordNode
	Species    Name
}

func (p *Engine) newHybridFormulaNode(n *node32) *hybridFormulaNode {
	var hf *hybridFormulaNode
	p.AddWarn(HybridFormulaWarn)
	n = n.up
	firstName := p.newSingleName(n)
	n = n.next
	var hes []*hybridElement
	var he *hybridElement
	for n != nil {
		switch n.pegRule {
		case ruleHybridChar:
			he = &hybridElement{
				HybridChar: p.newWordNode(n, HybridCharType),
			}
		case ruleSingleName:
			he.Species = p.newSingleName(n)
			hes = append(hes, he)
		case ruleSpeciesEpithet:
			p.AddWarn(HybridFormulaIncompleteWarn)
			var g *wordNode
			switch node := firstName.(type) {
			case *speciesNode:
				g = node.Genus
			case *uninomialNode:
				g = node.Word
			case *comparisonNode:
				g = node.Genus
			}
			spe := p.newSpeciesEpithetNode(n)
			g = &wordNode{Value: g.Value, NormValue: g.NormValue}
			he.Species = &speciesNode{Genus: g, SpEpithet: spe}
			hes = append(hes, he)
		}
		n = n.next
	}
	if he.Species == nil {
		p.AddWarn(HybridFormulaProbIncompleteWarn)
		hes = append(hes, he)
	}
	hf = &hybridFormulaNode{
		FirstSpecies:   firstName,
		HybridElements: hes,
	}
	hf.normalizeAbbreviated()
	p.Cardinality = 0
	return hf
}

func (hf *hybridFormulaNode) normalizeAbbreviated() {
	var fsv string
	if fsp, ok := hf.FirstSpecies.(*speciesNode); ok {
		fsv = fsp.Genus.NormValue
	} else {
		return
	}
	for _, v := range hf.HybridElements {
		if sp, ok := v.Species.(*speciesNode); ok {
			val := sp.Genus.NormValue
			if val[len(val)-1] == '.' && fsv[0:len(val)-1] == val[0:len(val)-1] {
				sp.Genus.NormValue = fsv
				v.Species = sp
			}
		} else {
			continue
		}
	}
}

type namedGenusHybridNode struct {
	Hybrid *wordNode
	Name
}

func (p *Engine) newNamedGenusHybridNode(n *node32) *namedGenusHybridNode {
	var nhn *namedGenusHybridNode
	var name Name
	n = n.up
	if n.token32.pegRule != ruleHybridChar {
		return nhn
	}
	hybr := p.newWordNode(n, HybridCharType)
	n = n.next
	n = n.up
	p.AddWarn(HybridNamedWarn)
	if n.token32.begin == 1 {
		p.AddWarn(HybridCharNoSpaceWarn)
	}
	switch n.token32.pegRule {
	case ruleUninomial:
		name = p.newUninomialNode(n)
	case ruleUninomialCombo:
		p.AddWarn(UninomialComboWarn)
		name = p.newUninomialComboNode(n)
	case ruleNameSpecies:
		name = p.newSpeciesNode(n)
	case ruleNameApprox:
		p.Surrogate = true
		p.AddWarn(NameApproxWarn)
		name = p.newApproxNode(n)
	}
	nhn = &namedGenusHybridNode{
		Hybrid: hybr,
		Name:   name,
	}
	return nhn
}

type namedSpeciesHybridNode struct {
	Genus        *wordNode
	Comparison   *wordNode
	Hybrid       *wordNode
	SpEpithet    *spEpithetNode
	InfraSpecies []*infraspEpithetNode
}

func (p *Engine) newNamedSpeciesHybridNode(n *node32) *namedSpeciesHybridNode {
	var nhl *namedSpeciesHybridNode
	n = n.up
	var gen, hybrid, cf *wordNode
	var sp *spEpithetNode
	var infs []*infraspEpithetNode
	for n != nil {
		switch n.pegRule {
		case ruleGenusWord:
			gen = p.newWordNode(n, GenusType)
		case ruleComparison:
			cf = p.newWordNode(n, ComparisonType)
			p.Surrogate = true
			p.AddWarn(NameComparisonWarn)
		case ruleHybridChar:
			hybrid = p.newWordNode(n, HybridCharType)
		case ruleSpeciesEpithet:
			sp = p.newSpeciesEpithetNode(n)
		case ruleInfraspGroup:
			infs = p.newInfraspeciesGroup(n)
		}
		n = n.next
	}

	p.AddWarn(HybridNamedWarn)
	if hybrid.Pos.End == sp.Word.Pos.Start {
		p.AddWarn(HybridCharNoSpaceWarn)
	}
	p.Cardinality = 2 + len(infs)
	nhl = &namedSpeciesHybridNode{
		Genus:        gen,
		Comparison:   cf,
		Hybrid:       hybrid,
		SpEpithet:    sp,
		InfraSpecies: infs,
	}
	return nhl
}

func (p *Engine) botanicalUninomial(n *node32) bool {
	n = n.up
	if n.token32.pegRule == ruleUninomial {
		return false
	}
	n = n.next
	n = n.up
	if n.token32.pegRule != ruleUninomialWord {
		return false
	}
	w := p.newWordNode(n, UnknownType)

	if _, ok := dict.Dict.AuthorICN[w.NormValue]; ok {
		return true
	}
	return false
}

func (p *Engine) newBotanicalUninomialNode(n *node32) *uninomialNode {
	var at2 *authorsGroupNode
	n = n.up
	w := p.newWordNode(n, UninomialType)
	n = n.next // fake Subgenus
	au := p.newWordNode(n.up, AuthorWordType)
	an := &authorNode{Value: au.NormValue, Words: []*wordNode{au}}
	at := &authorsTeamNode{Authors: []*authorNode{an}}
	ag := &authorsGroupNode{Team1: at, Parens: true}
	n = n.next
	if n != nil {
		n = n.up // fake OriginalAuthorship
		switch n.token32.pegRule {
		case ruleOriginalAuthorship:
			at2 = p.newAuthorsGroupNode(n.up)
		default:
			p.Tail = p.tailValue(n)
		}
	}
	authorship := &authorshipNode{OriginalAuthors: ag, CombinationAuthors: at2}
	u := &uninomialNode{Word: w, Authorship: authorship}
	p.AddWarn(BotanyAuthorNotSubgenWarn)
	p.Cardinality = 1
	return u
}

func (p *Engine) newSingleName(n *node32) Name {
	var name Name
	n = n.up
	switch n.token32.pegRule {
	case ruleNameSpecies:
		name = p.newSpeciesNode(n)
	case ruleNameApprox:
		p.AddWarn(NameApproxWarn)
		p.Surrogate = true
		name = p.newApproxNode(n)
	case ruleNameComp:
		p.AddWarn(NameComparisonWarn)
		p.Surrogate = true
		name = p.newComparisonNode(n)
	case ruleUninomial:
		name = p.newUninomialNode(n)
	case ruleUninomialCombo:
		if p.botanicalUninomial(n) {
			return p.newBotanicalUninomialNode(n)
		}
		p.AddWarn(UninomialComboWarn)
		name = p.newUninomialComboNode(n)
	}
	return name
}

type approxNode struct {
	Genus     *wordNode
	SpEpithet *spEpithetNode
	Approx    *wordNode
	Ignored   string
}

func (p *Engine) newApproxNode(n *node32) *approxNode {
	var an *approxNode
	if n.token32.pegRule != ruleNameApprox {
		return an
	}
	var gen *wordNode
	var spEp *spEpithetNode
	var annot *wordNode
	var ign string
	n = n.up
	for n != nil {
		switch n.token32.pegRule {
		case ruleGenusWord:
			gen = p.newWordNode(n, GenusType)
		case ruleSpeciesEpithet:
			spEp = p.newSpeciesEpithetNode(n)
		case ruleApproximation:
			annot = p.newWordNode(n, ApproxType)
		case ruleApproxNameIgnored:
			ign = p.nodeValue(n)
		}
		n = n.next
	}
	an = &approxNode{
		Genus:     gen,
		SpEpithet: spEp,
		Approx:    annot,
		Ignored:   ign,
	}
	p.Cardinality = 0
	return an
}

type comparisonNode struct {
	Genus      *wordNode
	SpEpithet  *spEpithetNode
	Comparison *wordNode
}

func (p *Engine) newComparisonNode(n *node32) *comparisonNode {
	var cn *comparisonNode
	if n.pegRule != ruleNameComp {
		return cn
	}
	n = n.up
	var gen, comp *wordNode
	var spEp *spEpithetNode
	for n != nil {
		switch n.pegRule {
		case ruleGenusWord:
			gen = p.newWordNode(n, GenusType)
			p.Cardinality = 1
		case ruleComparison:
			comp = p.newWordNode(n, ComparisonType)
		case ruleSpeciesEpithet:
			spEp = p.newSpeciesEpithetNode(n)
			p.Cardinality = 2
		}
		n = n.next
	}
	cn = &comparisonNode{
		Genus:      gen,
		Comparison: comp,
		SpEpithet:  spEp,
	}
	return cn
}

type speciesNode struct {
	Genus        *wordNode
	SubGenus     *wordNode
	SpEpithet    *spEpithetNode
	InfraSpecies []*infraspEpithetNode
}

func (p *Engine) newSpeciesNode(n *node32) *speciesNode {
	var sp *spEpithetNode
	var sg *wordNode
	var infs []*infraspEpithetNode
	n = n.up
	gen := p.newWordNode(n, GenusType)
	if n.up.token32.pegRule == ruleAbbrGenus {
		p.AddWarn(GenusAbbrWarn)
	}
	n = n.next
	for n != nil {
		switch n.token32.pegRule {
		case ruleSubGenus:
			w := p.newWordNode(n.up, SubGenusType)
			if _, ok := dict.Dict.AuthorICN[w.NormValue]; ok {
				p.AddWarn(BotanyAuthorNotSubgenWarn)
			} else {
				sg = w
			}
		case ruleSubGenusOrSuperspecies:
			p.AddWarn(SuperSpeciesWarn)
		case ruleSpeciesEpithet:
			sp = p.newSpeciesEpithetNode(n)
		case ruleInfraspGroup:
			infs = p.newInfraspeciesGroup(n)
		}
		n = n.next
	}
	p.Cardinality = 2 + len(infs)
	sn := speciesNode{
		Genus:        gen,
		SubGenus:     sg,
		SpEpithet:    sp,
		InfraSpecies: infs,
	}
	if len(infs) > 0 && infs[0].Rank == nil && sp.Authorship != nil &&
		sp.Authorship.TerminalFilius {
		p.AddWarn(AuthAmbiguousFiliusWarn)
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
		if len(infs) > 0 && inf.Rank == nil {
			infPrev := infs[len(infs)-1]
			if infPrev.Authorship != nil && infPrev.Authorship.TerminalFilius {
				p.AddWarn(AuthAmbiguousFiliusWarn)
			}
		}
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
		w.NormValue = "f."
	case ruleRankVar:
		w.NormValue = "var."
	case ruleRankSsp:
		w.NormValue = "subsp."
	case ruleRankOtherUncommon:
		p.AddWarn(RankUncommonWarn)
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
	p.Cardinality = 1
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
		rw := &wordNode{Value: "subgen.", NormValue: "subgen.",
			Pos: Pos{Type: RankUniType}}
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
	p.Cardinality = 1
	return &ucn
}

type rankUninomialNode struct {
	Word *wordNode
}

func (p *Engine) newRankUninomialNode(n *node32) *rankUninomialNode {
	r := p.newWordNode(n, RankUniType)
	run := rankUninomialNode{Word: r}
	switch {
	case strings.HasPrefix(run.Word.Value, "subg"):
		run.Word.NormValue = "subgen."
	case strings.HasPrefix(run.Word.Value, "fam"):
		run.Word.NormValue = "fam."
	}
	return &run
}

type authorshipNode struct {
	OriginalAuthors    *authorsGroupNode
	CombinationAuthors *authorsGroupNode
	TerminalFilius     bool
}

func (p *Engine) newAuthorshipNode(n *node32) *authorshipNode {
	var a *authorshipNode
	if n == nil {
		return a
	}
	var oa, ca *authorsGroupNode
	var misplacedYear bool
	var fil bool
	n = n.up
	for n != nil {
		switch n.token32.pegRule {
		case ruleOriginalAuthorship:
			oa = p.newAuthorsGroupNode(n.up)
		case ruleOriginalAuthorshipComb:
			on := n.up
			if on.token32.pegRule == ruleBasionymAuthorshipYearMisformed {
				p.AddWarn(YearOrigMisplacedWarn)
				on = on.up
				misplacedYear = true
			} else {
				on = on.up
			}
			oa = p.newAuthorsGroupNode(on)
			oa.Parens = true
			if misplacedYear {
				yr := p.newYearNode(on.next)
				oa.Team1.Year = yr
			}
		case ruleCombinationAuthorship:
			ca = p.newAuthorsGroupNode(n.up)
		}
		n = n.next
	}
	fil = oa.TerminalFilius && !oa.Parens
	if ca != nil {
		fil = ca.TerminalFilius
	}

	a = &authorshipNode{
		OriginalAuthors:    oa,
		CombinationAuthors: ca,
		TerminalFilius:     fil,
	}
	return a
}

type authorsGroupNode struct {
	Team1          *authorsTeamNode
	Team2Type      *wordNode
	Team2          *authorsTeamNode
	Parens         bool
	TerminalFilius bool
}

func (p *Engine) newAuthorsGroupNode(n *node32) *authorsGroupNode {
	var t1, t2 *authorsTeamNode
	var t2t *wordNode
	n = n.up
	t1 = p.newAuthorTeam(n)
	fil := t1.TerminalFilius
	ag := authorsGroupNode{
		Team1:          t1,
		Team2Type:      t2t,
		Team2:          t2,
		TerminalFilius: fil,
	}
	n = n.next
	if n == nil {
		return &ag
	}
	switch n.token32.pegRule {
	case ruleAuthorEx:
		p.AddWarn(AuthExWarn)
		t2t = p.newWordNode(n, AuthorWordExType)
		ex := strings.TrimSpace(t2t.Value)
		if ex[len(ex)-1] == '.' {
			p.AddWarn(AuthExWithDotWarn)
		}
		t2t.NormValue = "ex"
	case ruleAuthorEmend:
		p.AddWarn(AuthEmendWarn)
		t2t = p.newWordNode(n, AuthorWordEmendType)
		emend := strings.TrimSpace(t2t.Value)
		if emend[len(emend)-1] != '.' {
			p.AddWarn(AuthEmendWithoutDotWarn)
		}
		t2t.NormValue = "emend."
	default:
		return &ag
	}
	n = n.next
	if n == nil || n.token32.pegRule != ruleAuthorsTeam {
		return &ag
	}
	t2 = p.newAuthorTeam(n)
	ag.Team2Type = t2t
	ag.Team2 = t2
	ag.TerminalFilius = ag.Team2.TerminalFilius
	return &ag
}

type authorsTeamNode struct {
	Authors        []*authorNode
	TerminalFilius bool
	Year           *yearNode
}

func (p *Engine) newAuthorTeam(n *node32) *authorsTeamNode {
	var anodes []*node32
	var seps []string
	var yr *yearNode
	n = n.up
	for n != nil {
		switch n.token32.pegRule {
		case ruleAuthor:
			anodes = append(anodes, n)
		case ruleAuthorSep:
			seps = append(seps, p.nodeValue(n))
		case ruleYear:
			yr = p.newYearNode(n)
		}
		n = n.next
	}
	aus := make([]*authorNode, len(anodes))
	for i, v := range anodes {
		aus[i] = p.newAuthorNode(v)
		if i < len(seps) {
			switch {
			case strings.Contains(seps[i], "apud"):
				seps[i] = " apud "
			case i < len(seps)-1:
				seps[i] = ", "
			case i == len(seps)-1:
				seps[i] = " & "
			}
			aus[i].Sep = seps[i]
		}
	}
	atn := authorsTeamNode{
		Authors:        aus,
		TerminalFilius: aus[len(aus)-1].Filius,
		Year:           yr,
	}
	return &atn
}

type authorNode struct {
	Value  string
	Sep    string
	Words  []*wordNode
	Filius bool
}

func (p *Engine) newAuthorNode(n *node32) *authorNode {
	var w *wordNode
	var fil bool
	var ws []*wordNode
	val := ""
	rawVal := ""
	n = n.up
	for n != nil {
		switch n.token32.pegRule {
		case ruleFilius:
			w = p.newWordNode(n, AuthorWordFiliusType)
			w.NormValue = "fil."
			fil = true
		case ruleUnknownAuthor:
			p.AddWarn(AuthUnknownWarn)
			w = p.authorWord(n)
			if w.Value == "?" {
				p.AddWarn(AuthQuestionWarn)
			}
			w.NormValue = "anon."
		default:
			w = p.authorWord(n)
		}
		ws = append(ws, w)
		val = str.JoinStrings(val, w.NormValue, " ")
		rawVal = str.JoinStrings(rawVal, w.Value, " ")
		n = n.next
	}
	if len(rawVal) < 2 {
		p.AddWarn(AuthShortWarn)
	}
	au := authorNode{
		Value:  val,
		Words:  ws,
		Filius: fil,
	}
	return &au
}

func (p *Engine) authorWord(n *node32) *wordNode {
	w := p.newWordNode(n, AuthorWordType)
	if n.up != nil && n.up.token32.pegRule == ruleAllCapsAuthorWord {
		count := 0
		for _, v := range w.Value {
			if unicode.IsUpper(v) {
				count++
			}
		}
		if count > 2 {
			w.NormValue = str.FixAllCaps(w.NormValue)
			p.AddWarn(AuthUpperCaseWarn)
		}
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
		case ruleYearWithPage:
			p.AddWarn(YearPageWarn)
		case ruleYearRange:
			p.AddWarn(YearRangeWarn)
			appr = true
		case ruleYearWithParens:
			p.AddWarn(YearParensWarn)
			appr = true
		case ruleYearApprox:
			p.AddWarn(YearSqBraketsWarn)
			appr = true
		case ruleYearWithChar:
			p.AddWarn(YearCharWarn)
			w = p.newWordNode(v, YearType)
			w.NormValue = w.Value[0 : len(w.Value)-1]
		case ruleYearNum:
			if w == nil {
				w = p.newWordNode(v, YearType)
			}
			if w.Value[len(w.Value)-1] == '?' {
				p.AddWarn(YearQuestionWarn)
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
	var canApostrophe bool
	for _, v := range children {
		switch v.token32.pegRule {
		case ruleAuthorEtAl:
			if strings.Contains(wrd.NormValue, "&") {
				wrd.NormValue = "et al."
			}
		case ruleUpperCharExtended, ruleLowerCharExtended:
			p.AddWarn(CharBadWarn)
			_ = wrd.normalize()
		case ruleWordApostr:
			p.AddWarn(CanonicalApostropheWarn)
			canApostrophe = true
			_ = wrd.normalize()
		case ruleWordStartsWithDigit:
			p.AddWarn(SpeciesNumericWarn)
			wrd.normalizeNums()
		case ruleApostrOther:
			p.AddWarn(ApostrOtherWarn)
			if !canApostrophe {
				nv, _ := str.ToASCII([]byte(wrd.Value), str.GlobalTransliterations)
				wrd.NormValue = string(nv)
			}
		}
	}
	if wt == GenusType || wt == UninomialType {
		if val[len(val)-1] == '?' {
			p.AddWarn(CapWordQuestionWarn)
			wrd.NormValue = wrd.NormValue[0 : len(wrd.NormValue)-1]
		}
		if _, ok := p.Warnings[GenusUpperCharAfterDash]; ok {
			runes := []rune(wrd.Value)
			nv := make([]rune, len(runes))
			var afterDash bool
			for i, v := range runes {
				switch {
				case v == '-':
					afterDash = true
				case afterDash:
					v = unicode.ToLower(v)
					afterDash = false
				}
				nv[i] = v
			}
			wrd.NormValue = string(nv)
		}
		p.IsBacteria(wrd.NormValue)
	}
	return &wrd
}

func (w *wordNode) normalize() error {
	if w == nil {
		return nil
	}
	nv, err := str.ToASCII([]byte(w.Value), str.Transliterations)
	if err != nil {
		return err
	}
	w.NormValue = string(nv)
	return nil
}

func (w *wordNode) normalizeNums() {
	match := numWord.FindAllStringSubmatch(w.Value, 1)
	if len(match) == 0 {
		return
	}
	num := match[0][1]
	wrd := match[0][2]
	w.NormValue = str.NumToStr(num) + wrd
}

type Pos struct {
	Type  WordType
	Start int
	End   int
}

var numWord = regexp.MustCompile(`^([0-9]+)[-\.]?(.+)$`)
