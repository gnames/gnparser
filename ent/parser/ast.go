package parser

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/gnames/gnparser/ent/internal/preprocess"

	"github.com/gnames/gnparser/ent/parsed"
	"github.com/gnames/gnparser/ent/str"
	"github.com/gnames/gnparser/io/dict"
	"github.com/gnames/gnuuid"
	"github.com/gnames/tribool"
)

type scientificNameNode struct {
	nameData
	verbatim         string
	verbatimID       string
	cardinality      int
	rank             string
	virus            bool
	daggerChar       bool
	hybrid           *parsed.Annotation
	surrogate        *parsed.Annotation
	cultivar         bool
	bacteria         *tribool.Tribool
	candidatus       bool
	tail             string
	parserVersion    string
	ambiguousEpithet string
	ambiguousModif   string
	warnings         map[parsed.Warning]struct{}
	withSpGroup      bool
}

func (p *Engine) newScientificNameNode() {
	n := p.root.up
	var name nameData
	var tail string

	for n != nil {
		switch n.pegRule {
		case ruleName:
			name = p.newName(n)
		case ruleTail:
			tail = p.tailValue(n)
		}
		n = n.next
	}
	if p.tail != "" && tail == "" {
		tail = p.tail
	}
	if p.cardinality == 2 {
		p.rank = "sp."
	}
	if p.cultivar {
		p.rank = ""
	}
	sn := scientificNameNode{
		nameData:    name,
		cardinality: p.cardinality,
		rank:        p.rank,
		hybrid:      p.hybrid,
		surrogate:   p.surrogate,
		bacteria:    p.bacteria,
		candidatus:  p.candidatus,
		cultivar:    p.cultivar,
		tail:        tail,
	}
	p.sn = &sn
}

func (p *Engine) newNotParsedScientificNameNode(pp *preprocess.Preprocessor) {
	sn := &scientificNameNode{virus: pp.Virus}
	p.sn = sn
}

func (sn *scientificNameNode) addVerbatim(s string) {
	sn.verbatim = s
	sn.verbatimID = gnuuid.New(s).String()
}

func (p *Engine) tailValue(n *node32) string {
	t := n.token32
	tail := string(p.buffer[t.begin:t.end])
	tail = strings.TrimRight(tail, " ")
	return tail
}

func (p *Engine) newName(n *node32) nameData {
	var name nameData
	var annot parsed.Annotation
	n = n.up
	switch n.pegRule {
	case ruleHybridFormula:
		annot = parsed.HybridFormulaAnnot
		p.hybrid = &annot
		name = p.newHybridFormulaNode(n)
	case ruleNamedGenusHybrid:
		annot = parsed.NamedHybridAnnot
		p.hybrid = &annot
		name = p.newNamedGenusHybridNode(n)
	case ruleNamedSpeciesHybrid:
		annot = parsed.NamedHybridAnnot
		p.hybrid = &annot
		name = p.newNamedSpeciesHybridNode(n)
	case ruleGraftChimeraFormula:
		if p.enableCultivars {
			annot = parsed.GraftChimeraFormulaAnnot
			p.hybrid = &annot
			name = p.newGraftChimeraFormulaNode(n)
		}
	case ruleNamedGenusGraftChimera:
		if p.enableCultivars {
			annot = parsed.NamedGraftChimeraAnnot
			p.hybrid = &annot
			name = p.newNamedGenusGraftChimeraNode(n)
		}
	case ruleCandidatusName:
		name = p.newCandidatusName(n)
	case ruleSingleName:
		name = p.newSingleName(n)
	}
	return name
}

type hybridFormulaNode struct {
	FirstSpecies   nameData
	HybridElements []*hybridElement
}

type hybridElement struct {
	HybridChar *parsed.Word
	Species    nameData
}

type graftChimeraFormulaNode struct {
	FirstSpecies         nameData
	GraftChimeraElements []*graftChimeraElement
}

type graftChimeraElement struct {
	GraftChimeraChar *parsed.Word
	Species          nameData
}

func (p *Engine) newHybridFormulaNode(n *node32) *hybridFormulaNode {
	var hf *hybridFormulaNode
	p.addWarn(parsed.HybridFormulaWarn)
	n = n.up
	firstName := p.newSingleName(n)
	n = n.next
	var hes []*hybridElement
	var he *hybridElement
	for n != nil {
		switch n.pegRule {
		case ruleHybridChar:
			he = &hybridElement{
				HybridChar: p.newWordNode(n, parsed.HybridCharType),
			}
		case ruleSingleName:
			he.Species = p.newSingleName(n)
			hes = append(hes, he)
		case ruleSpeciesEpithet:
			p.addWarn(parsed.HybridFormulaIncompleteWarn)
			var g *parsed.Word
			switch node := firstName.(type) {
			case *speciesNode:
				g = node.Genus
			case *uninomialNode:
				g = node.Word
			case *comparisonNode:
				g = node.Genus
			}
			spe := p.newSpeciesEpithetNode(n)
			g = &parsed.Word{Verbatim: g.Verbatim, Normalized: g.Normalized}
			he.Species = &speciesNode{Genus: g, SpEpithet: spe}
			hes = append(hes, he)
		}
		n = n.next
	}
	if he.Species == nil {
		p.addWarn(parsed.HybridFormulaProbIncompleteWarn)
		hes = append(hes, he)
	}
	hf = &hybridFormulaNode{
		FirstSpecies:   firstName,
		HybridElements: hes,
	}
	hf.normalizeAbbreviated()
	p.cardinality = 0
	return hf
}

func (p *Engine) newGraftChimeraFormulaNode(n *node32) *graftChimeraFormulaNode {
	var gcf *graftChimeraFormulaNode
	p.addWarn(parsed.GraftChimeraFormulaWarn)
	n = n.up
	firstName := p.newSingleName(n)
	n = n.next
	var gces []*graftChimeraElement
	var gce *graftChimeraElement
	for n != nil {
		switch n.pegRule {
		case ruleGraftChimeraChar:
			gce = &graftChimeraElement{
				GraftChimeraChar: p.newWordNode(n, parsed.GraftChimeraCharType),
			}
		case ruleSingleName:
			gce.Species = p.newSingleName(n)
			gces = append(gces, gce)
		case ruleSpeciesEpithet:
			p.addWarn(parsed.GraftChimeraFormulaIncompleteWarn)
			var g *parsed.Word
			switch node := firstName.(type) {
			case *speciesNode:
				g = node.Genus
			case *uninomialNode:
				g = node.Word
			case *comparisonNode:
				g = node.Genus
			}
			spe := p.newSpeciesEpithetNode(n)
			g = &parsed.Word{Verbatim: g.Verbatim, Normalized: g.Normalized}
			gce.Species = &speciesNode{Genus: g, SpEpithet: spe}
			gces = append(gces, gce)
		}
		n = n.next
	}
	if gce.Species == nil {
		p.addWarn(parsed.GraftChimeraFormulaProbIncompleteWarn)
		gces = append(gces, gce)
	}
	gcf = &graftChimeraFormulaNode{
		FirstSpecies:         firstName,
		GraftChimeraElements: gces,
	}
	gcf.normalizeAbbreviated()
	p.cardinality = 0
	return gcf
}

func (hf *hybridFormulaNode) normalizeAbbreviated() {
	var fsv string
	if fsp, ok := hf.FirstSpecies.(*speciesNode); ok {
		fsv = fsp.Genus.Normalized
	} else {
		return
	}
	for _, v := range hf.HybridElements {
		if sp, ok := v.Species.(*speciesNode); ok {
			val := sp.Genus.Normalized
			if val[len(val)-1] == '.' && fsv[0:len(val)-1] == val[0:len(val)-1] {
				sp.Genus.Normalized = fsv
				v.Species = sp
			}
		} else {
			continue
		}
	}
}

func (gcf *graftChimeraFormulaNode) normalizeAbbreviated() {
	var fsv string
	if fsp, ok := gcf.FirstSpecies.(*speciesNode); ok {
		fsv = fsp.Genus.Normalized
	} else {
		return
	}
	for _, v := range gcf.GraftChimeraElements {
		if sp, ok := v.Species.(*speciesNode); ok {
			val := sp.Genus.Normalized
			if val[len(val)-1] == '.' && fsv[0:len(val)-1] == val[0:len(val)-1] {
				sp.Genus.Normalized = fsv
				v.Species = sp
			}
		} else {
			continue
		}
	}
}

type namedGenusHybridNode struct {
	Hybrid *parsed.Word
	nameData
}

func (p *Engine) newNamedGenusHybridNode(n *node32) *namedGenusHybridNode {
	var nhn *namedGenusHybridNode
	var name nameData
	n = n.up
	if n.pegRule != ruleHybridChar {
		return nhn
	}
	hybr := p.newWordNode(n, parsed.HybridCharType)
	n = n.next
	n = n.up
	p.addWarn(parsed.HybridNamedWarn)
	if n.begin == 1 {
		p.addWarn(parsed.HybridCharNoSpaceWarn)
	}
	switch n.pegRule {
	case ruleUninomial:
		name = p.newUninomialNode(n)
	case ruleUninomialCombo:
		p.addWarn(parsed.UninomialComboWarn)
		name = p.newUninomialComboNode(n)
	case ruleNameSpecies:
		name = p.newSpeciesNode(n)
	case ruleNameApprox:
		name = p.newApproxNode(n)
	}
	nhn = &namedGenusHybridNode{
		Hybrid:   hybr,
		nameData: name,
	}
	return nhn
}

type namedSpeciesHybridNode struct {
	Genus        *parsed.Word
	Comparison   *parsed.Word
	Hybrid       *parsed.Word
	SpEpithet    *spEpithetNode
	Infraspecies []*infraspEpithetNode
}

func (p *Engine) newNamedSpeciesHybridNode(n *node32) *namedSpeciesHybridNode {
	var nhl *namedSpeciesHybridNode
	var annot parsed.Annotation
	n = n.up
	var gen, hybrid, cf *parsed.Word
	var sp *spEpithetNode
	var infs []*infraspEpithetNode
	for n != nil {
		switch n.pegRule {
		case ruleGenusWord:
			gen = p.newWordNode(n, parsed.GenusType)
		case ruleComparison:
			cf = p.newWordNode(n, parsed.ComparisonMarkerType)
			annot = parsed.ComparisonAnnot
			p.surrogate = &annot
			p.addWarn(parsed.NameComparisonWarn)
		case ruleHybridChar:
			hybrid = p.newWordNode(n, parsed.HybridCharType)
		case ruleSpeciesEpithet:
			sp = p.newSpeciesEpithetNode(n)
		case ruleInfraspGroup:
			infs = p.newInfraspeciesGroup(n)
		}
		n = n.next
	}

	p.addWarn(parsed.HybridNamedWarn)
	if hybrid.End == sp.Word.Start {
		p.addWarn(parsed.HybridCharNoSpaceWarn)
	}
	p.cardinality = 2 + len(infs)
	nhl = &namedSpeciesHybridNode{
		Genus:        gen,
		Comparison:   cf,
		Hybrid:       hybrid,
		SpEpithet:    sp,
		Infraspecies: infs,
	}
	return nhl
}

type namedGenusGraftChimeraNode struct {
	GraftChimera *parsed.Word
	nameData
}

func (p *Engine) newNamedGenusGraftChimeraNode(n *node32) *namedGenusGraftChimeraNode {
	var nhn *namedGenusGraftChimeraNode
	var name nameData
	n = n.up
	if n.pegRule != ruleGraftChimeraChar {
		return nhn
	}
	gc := p.newWordNode(n, parsed.GraftChimeraCharType)
	n = n.next
	n = n.up
	p.addWarn(parsed.GraftChimeraNamedWarn)
	if n.begin == 1 {
		p.addWarn(parsed.GraftChimeraCharNoSpaceWarn)
	}
	switch n.pegRule {
	case ruleUninomial:
		name = p.newUninomialNode(n)
	case ruleUninomialCombo:
		p.addWarn(parsed.UninomialComboWarn)
		name = p.newUninomialComboNode(n)
	case ruleNameSpecies:
		name = p.newSpeciesNode(n)
	case ruleNameApprox:
		name = p.newApproxNode(n)
	}
	nhn = &namedGenusGraftChimeraNode{
		GraftChimera: gc,
		nameData:     name,
	}
	return nhn
}

func (p *Engine) botanicalUninomial(n *node32) bool {
	n = n.up
	if n.pegRule == ruleUninomial {
		return false
	}
	n = n.next
	n = n.up
	if n.pegRule != ruleUninomialWord {
		return false
	}
	w := p.newWordNode(n, parsed.AuthorWordType)

	if _, ok := dict.Dict.AuthorICN[w.Normalized]; ok {
		return true
	}
	return false
}

func (p *Engine) newBotanicalUninomialNode(n *node32) *uninomialNode {
	var at2 *authorsGroupNode
	n = n.up
	w := p.newWordNode(n, parsed.UninomialType)
	n = n.next // fake Subgenus
	verbatim := p.nodeValue(n)
	au := p.newWordNode(n.up, parsed.AuthorWordType)
	an := &authorNode{Value: au.Normalized, Words: []*parsed.Word{au}}
	at := &authorsTeamNode{Authors: []*authorNode{an}}
	ag := &authorsGroupNode{Team1: at, Parens: true}
	n = n.next
	if n != nil {
		// cheating by adding space by hand here, so it is not real verbatim
		// as a result.
		verbatim += " " + p.nodeValue(n)
		n = n.up // fake OriginalAuthorship
		switch n.pegRule {
		case ruleOriginalAuthorship:
			at2 = p.newAuthorsGroupNode(n.up)
		default:
			p.tail = p.tailValue(n)
		}
	}
	authorship := &authorshipNode{
		Verbatim:           verbatim,
		OriginalAuthors:    ag,
		CombinationAuthors: at2,
	}
	u := &uninomialNode{Word: w, Authorship: authorship}
	p.addWarn(parsed.BotanyAuthorNotSubgenWarn)
	p.cardinality = 1
	return u
}

func (p *Engine) newSingleName(n *node32) nameData {
	var name nameData
	var annot parsed.Annotation
	n = n.up
	switch n.pegRule {
	case ruleNameSpecies:
		name = p.newSpeciesNode(n)
	case ruleNameApprox:
		name = p.newApproxNode(n)
	case ruleNameComp:
		p.addWarn(parsed.NameComparisonWarn)
		annot = parsed.ComparisonAnnot
		p.surrogate = &annot
		name = p.newComparisonNode(n)
	case ruleUninomial:
		name = p.newUninomialNode(n)
	case ruleUninomialCombo:
		if p.botanicalUninomial(n) {
			return p.newBotanicalUninomialNode(n)
		}
		name = p.newUninomialComboNode(n)
	}
	return name
}

type candidatusNameNode struct {
	Candidatus *parsed.Word
	SingleName nameData
}

func (p *Engine) newCandidatusName(n *node32) nameData {
	bac := tribool.New(1)
	p.bacteria = &bac
	p.candidatus = true
	p.addWarn(parsed.CandidatusName)

	var cand *parsed.Word
	var singName nameData

	n = n.up
	for n != nil {
		switch n.pegRule {
		case ruleCandidatus:
			cand = p.newWordNode(n, parsed.CandidatusType)
		case ruleSingleName:
			singName = p.newSingleName(n)
		}
		n = n.next
	}

	candName := &candidatusNameNode{Candidatus: cand, SingleName: singName}
	return candName
}

type approxNode struct {
	Genus     *parsed.Word
	SpEpithet *spEpithetNode
	Approx    *parsed.Word
	Ignored   string
}

func (p *Engine) newApproxNode(n *node32) *approxNode {
	var an *approxNode
	annot := parsed.ApproximationAnnot
	p.surrogate = &annot
	p.addWarn(parsed.NameApproxWarn)
	if n.pegRule != ruleNameApprox {
		return an
	}
	var gen, appr *parsed.Word
	var spEp *spEpithetNode
	var ign string
	n = n.up
	for n != nil {
		switch n.pegRule {
		case ruleGenusWord:
			gen = p.newWordNode(n, parsed.GenusType)
		case ruleSpeciesEpithet:
			spEp = p.newSpeciesEpithetNode(n)
		case ruleApproximation:
			appr = p.newWordNode(n, parsed.ApproxMarkerType)
		case ruleApproxNameIgnored:
			ign = p.nodeValue(n)
		}
		n = n.next
	}
	an = &approxNode{
		Genus:     gen,
		SpEpithet: spEp,
		Approx:    appr,
		Ignored:   ign,
	}
	p.cardinality = 0
	return an
}

type comparisonNode struct {
	Genus          *parsed.Word
	SpEpithet      *spEpithetNode
	InfraSpEpithet *infraspEpithetNode
	Comparison     *parsed.Word
	Cardinality    int
}

func (p *Engine) newComparisonNode(n *node32) *comparisonNode {
	n = n.up
	switch n.pegRule {
	case ruleNameCompIsp:
		return p.newCompIspNode(n)
	case ruleNameCompSp:
		return p.newCompSpNode(n)
	default:
		var res *comparisonNode
		return res
	}
}

func (p *Engine) newCompIspNode(n *node32) *comparisonNode {
	var cn *comparisonNode
	n = n.up
	var gen, comp *parsed.Word
	var spEp *spEpithetNode
	var ispEp *infraspEpithetNode
	for n != nil {
		switch n.pegRule {
		case ruleGenusWord:
			gen = p.newWordNode(n, parsed.GenusType)
		case ruleComparison:
			comp = p.newWordNode(n, parsed.ComparisonMarkerType)
		case ruleSpeciesEpithet:
			spEp = p.newSpeciesEpithetNode(n)
			p.cardinality = 2
		case ruleInfraspEpithet:
			ispEp = p.newInfraspEpithetNode(n)
			p.cardinality = 3
		}
		n = n.next
	}
	cn = &comparisonNode{
		Genus:          gen,
		Comparison:     comp,
		SpEpithet:      spEp,
		InfraSpEpithet: ispEp,
		Cardinality:    3,
	}
	return cn

}

func (p *Engine) newCompSpNode(n *node32) *comparisonNode {
	var cn *comparisonNode
	n = n.up
	var gen, comp *parsed.Word
	var spEp *spEpithetNode
	for n != nil {
		switch n.pegRule {
		case ruleGenusWord:
			gen = p.newWordNode(n, parsed.GenusType)
			p.cardinality = 1
		case ruleComparison:
			comp = p.newWordNode(n, parsed.ComparisonMarkerType)
		case ruleSpeciesEpithet:
			spEp = p.newSpeciesEpithetNode(n)
			p.cardinality = 2
		}
		n = n.next
	}
	cn = &comparisonNode{
		Genus:       gen,
		Comparison:  comp,
		SpEpithet:   spEp,
		Cardinality: 2,
	}
	return cn
}

type speciesNode struct {
	Genus           *parsed.Word
	Subgenus        *parsed.Word
	SpEpithet       *spEpithetNode
	Infraspecies    []*infraspEpithetNode
	CultivarEpithet *cultivarEpithetNode
}

type cultivarEpithetNode struct {
	Word            *parsed.Word
	enableCultivars bool
}

func (p *Engine) newSpeciesNode(n *node32) *speciesNode {
	var sp *spEpithetNode
	var sg *parsed.Word
	var infs []*infraspEpithetNode
	var cultivar *cultivarEpithetNode
	n = n.up
	gen := p.newWordNode(n, parsed.GenusType)
	if n.up.pegRule == ruleAbbrGenus {
		p.addWarn(parsed.GenusAbbrWarn)
	}
	n = n.next
	for n != nil {
		switch n.pegRule {
		case ruleSubgenus:
			w := p.newWordNode(n.up, parsed.SubgenusType)
			if _, ok := dict.Dict.AuthorICN[w.Normalized]; ok {
				p.addWarn(parsed.BotanyAuthorNotSubgenWarn)
			} else {
				sg = w
			}
		case ruleSubgenusOrSuperspecies:
			p.addWarn(parsed.SuperspeciesWarn)
		case ruleSpeciesEpithet:
			sp = p.newSpeciesEpithetNode(n)
		case ruleInfraspGroup:
			infs = p.newInfraspeciesGroup(n)
		case ruleCultivar, ruleCultivarRecursive:
			cultivar = p.newCultivarEpithetNode(n, parsed.CultivarType)
		}
		n = n.next
	}
	p.cardinality = 2 + len(infs)
	if cultivar != nil && p.enableCultivars {
		p.cultivar = true
		p.cardinality += 1
	}
	sn := speciesNode{
		Genus:           gen,
		Subgenus:        sg,
		SpEpithet:       sp,
		Infraspecies:    infs,
		CultivarEpithet: cultivar,
	}
	if len(infs) > 0 && infs[0].Rank == nil && sp.Authorship != nil &&
		sp.Authorship.TerminalFilius {
		p.addWarn(parsed.AuthAmbiguousFiliusWarn)

	}
	return &sn
}

type spEpithetNode struct {
	Word       *parsed.Word
	Authorship *authorshipNode
}

func (p *Engine) newSpeciesEpithetNode(n *node32) *spEpithetNode {
	var au *authorshipNode
	n = n.up
	se := p.newWordNode(n, parsed.SpEpithetType)
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
	Word       *parsed.Word
	Rank       *rankNode
	Authorship *authorshipNode
}

func (p *Engine) newInfraspeciesGroup(n *node32) []*infraspEpithetNode {
	var infs []*infraspEpithetNode
	n = n.up
	if n == nil || n.pegRule != ruleInfraspEpithet {
		return infs
	}
	var currentInf *infraspEpithetNode
	for n != nil {
		inf := p.newInfraspEpithetNode(n)
		if len(infs) > 0 && inf.Rank == nil {
			infPrev := infs[len(infs)-1]
			if infPrev.Authorship != nil && infPrev.Authorship.TerminalFilius {
				p.addWarn(parsed.AuthAmbiguousFiliusWarn)
			}
		}
		infs = append(infs, inf)
		currentInf = inf
		n = n.next
	}
	if currentInf != nil && currentInf.Rank != nil {
		p.rank = currentInf.Rank.Word.Normalized
	}
	return infs
}

func (p *Engine) newInfraspEpithetNode(n *node32) *infraspEpithetNode {
	var inf infraspEpithetNode
	var r *rankNode
	var w *parsed.Word
	var au *authorshipNode
	n = n.up
	if n == nil {
		return &inf
	}

	for n != nil {
		switch n.pegRule {
		case ruleWord:
			w = p.newWordNode(n, parsed.InfraspEpithetType)
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
	Word *parsed.Word
}

func (p *Engine) newRankNode(n *node32) *rankNode {
	if n.up == nil {
		w := p.newWordNode(n, parsed.RankType)
		r := rankNode{Word: w}
		return &r
	}
	n = n.up
	w := p.newWordNode(n, parsed.RankType)
	switch n.pegRule {
	case ruleRankForma:
		w.Normalized = "f."
	case ruleRankVar:
		w.Normalized = "var."
	case ruleRankSsp:
		w.Normalized = "subsp."
	case ruleRankOtherUncommon:
		p.addWarn(parsed.RankUncommonWarn)
	}
	r := rankNode{Word: w}
	return &r
}

type uninomialNode struct {
	Word            *parsed.Word
	CultivarEpithet *cultivarEpithetNode
	Authorship      *authorshipNode
}

func (p *Engine) newUninomialNode(n *node32) *uninomialNode {
	var au *authorshipNode
	var cultivar *cultivarEpithetNode
	wn := n.up
	w := p.newWordNode(wn, parsed.UninomialType)
	if an := wn.next; an != nil {
		au = p.newAuthorshipNode(an)
	}
	n = n.next
	for n != nil {
		switch n.pegRule {
		case ruleCultivar, ruleCultivarRecursive:
			cultivar = p.newCultivarEpithetNode(n, parsed.CultivarType)
		}
		n = n.next
	}
	un := uninomialNode{
		Word:            w,
		Authorship:      au,
		CultivarEpithet: cultivar,
	}
	p.cardinality = 1
	if cultivar != nil && p.enableCultivars {
		p.cultivar = true
		p.cardinality += 1
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
	switch n.pegRule {
	case ruleUninomial:
		u1n := n
		u1 = p.newUninomialNode(u1n)
		rn := u1n.next
		r = p.newRankUninomialNode(rn)
		u2n := rn.next
		u2 = p.newUninomialNode(u2n)
	case ruleRankUninomial:
		rn := n
		r = p.newRankUninomialNode(rn)
		u2n := rn.next
		u2 = p.newUninomialNode(u2n)
	case ruleUninomialWord:
		uw := p.newWordNode(n, parsed.UninomialType)
		u1 = &uninomialNode{Word: uw}
		n = n.next
		u2w := p.newWordNode(n.up, parsed.UninomialType)
		n = n.next
		au2 := p.newAuthorshipNode(n)
		rw := &parsed.Word{
			Verbatim:   "subgen.",
			Normalized: "subgen.",
			Type:       parsed.RankType,
		}
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
	if r != nil {
		p.rank = r.Word.Normalized
	}
	if u1 != nil {
		p.addWarn(parsed.UninomialComboWarn)
	} else {
		p.addWarn(parsed.UninomialWithRank)
	}
	p.cardinality = 1
	return &ucn
}

type rankUninomialNode struct {
	Word *parsed.Word
}

func (p *Engine) newRankUninomialNode(n *node32) *rankUninomialNode {
	r := p.newWordNode(n, parsed.RankType)
	run := rankUninomialNode{Word: r}
	switch {
	case strings.HasPrefix(run.Word.Verbatim, "subg"):
		run.Word.Normalized = "subgen."
	case strings.HasPrefix(run.Word.Verbatim, "fam"):
		run.Word.Normalized = "fam."
	case strings.HasPrefix(run.Word.Verbatim, "tr"):
		run.Word.Normalized = "trib."
	case strings.HasPrefix(run.Word.Verbatim, "subtr"):
		run.Word.Normalized = "subtrib."
	}
	return &run
}

type authorshipNode struct {
	Verbatim           string
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
	verbatim := p.buffer[n.begin:n.end]
	n = n.up
	for n != nil {
		switch n.pegRule {
		case ruleOriginalAuthorship:
			oa = p.newAuthorsGroupNode(n.up)
		case ruleOriginalAuthorshipComb:
			on := n.up
			if on.pegRule == ruleBasionymAuthorshipYearMisformed {
				on = on.up
				misplacedYear = true
			} else {
				on = on.up
			}
			oa = p.newAuthorsGroupNode(on)
			oa.Parens = true
			if misplacedYear {
				yr := p.newYearNode(on.next)
				if oa.Team1.Year == nil {
					p.addWarn(parsed.YearOrigMisplacedWarn)
					oa.Team1.Year = yr
				} else {
					p.addWarn(parsed.YearMisplacedWarn)
				}
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
		Verbatim:           string(verbatim),
		OriginalAuthors:    oa,
		CombinationAuthors: ca,
		TerminalFilius:     fil,
	}
	return a
}

type teamType int

const (
	teamDefault teamType = iota
	teamEx
	teamIn
	teamEmend
)

type authorsGroupNode struct {
	Team1          *authorsTeamNode
	Team2Type      teamType
	Team2Word      *parsed.Word
	Team2          *authorsTeamNode
	Parens         bool
	TerminalFilius bool
}

func (p *Engine) newAuthorsGroupNode(n *node32) *authorsGroupNode {
	var t1, t2 *authorsTeamNode
	var t2t teamType
	var t2wrd *parsed.Word
	n = n.up
	t1 = p.newAuthorTeam(n)
	fil := t1.TerminalFilius
	ag := authorsGroupNode{
		Team1:          t1,
		Team2Type:      t2t,
		Team2Word:      t2wrd,
		Team2:          t2,
		TerminalFilius: fil,
	}
	n = n.next
	if n == nil {
		return &ag
	}
	switch n.pegRule {
	case ruleAuthorEx:
		p.addWarn(parsed.AuthExWarn)
		t2t = teamEx
		t2wrd = p.newWordNode(n, parsed.AuthorWordType)
		ex := strings.TrimSpace(t2wrd.Verbatim)
		if ex[len(ex)-1] == '.' {
			p.addWarn(parsed.AuthExWithDotWarn)
		}
		t2wrd.Normalized = "ex"
	case ruleAuthorIn:
		p.addWarn(parsed.AuthInWarn)
		t2t = teamIn
		t2wrd = p.newWordNode(n, parsed.AuthorWordType)
		inWrd := strings.TrimSpace(t2wrd.Verbatim)
		if inWrd[len(inWrd)-1] == '.' {
			p.addWarn(parsed.AuthInWithDotWarn)
		}
		t2wrd.Normalized = "in"
	case ruleAuthorEmend:
		p.addWarn(parsed.AuthEmendWarn)
		t2t = teamEmend
		t2wrd = p.newWordNode(n, parsed.AuthorWordType)
		emend := strings.TrimSpace(t2wrd.Verbatim)
		if emend[len(emend)-1] != '.' {
			p.addWarn(parsed.AuthEmendWithoutDotWarn)
		}
		t2wrd.Normalized = "emend."
	default:
		return &ag
	}
	n = n.next
	if n == nil || n.pegRule != ruleAuthorsTeam {
		return &ag
	}
	t2 = p.newAuthorTeam(n)
	ag.Team2Type = t2t
	ag.Team2Word = t2wrd
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
		switch n.pegRule {
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
	Words  []*parsed.Word
	Filius bool
}

func (p *Engine) newAuthorNode(n *node32) *authorNode {
	var w *parsed.Word
	var fil bool
	var ws []*parsed.Word
	val := ""
	rawVal := ""
	n = n.up
	for n != nil {
		switch n.pegRule {
		case ruleFilius, ruleFiliusFNoSpace:
			w = p.newWordNode(n, parsed.AuthorWordFiliusType)
			w.Normalized = "fil."
			fil = true
		case ruleUnknownAuthor:
			p.addWarn(parsed.AuthUnknownWarn)
			w = p.authorWord(n)
			if w.Verbatim == "?" {
				p.addWarn(parsed.AuthQuestionWarn)
			}
			w.Normalized = "anon."
		case ruleAuthorEtAl:
			w = p.newWordNode(n, parsed.AuthorWordType)
			if strings.Contains(w.Normalized, "&") {
				w.Normalized = "et al."
			}
		default:
			w = p.authorWord(n)
		}
		ws = append(ws, w)
		val = str.JoinStrings(val, w.Normalized, " ")
		rawVal = str.JoinStrings(rawVal, w.Verbatim, " ")
		n = n.next
	}
	if len(rawVal) < 2 {
		p.addWarn(parsed.AuthShortWarn)
	}
	au := authorNode{
		Value:  val,
		Words:  ws,
		Filius: fil,
	}
	return &au
}

func (p *Engine) authorWord(n *node32) *parsed.Word {
	w := p.newWordNode(n, parsed.AuthorWordType)
	if n.up != nil && n.up.pegRule == ruleAllCapsAuthorWord {
		count := 0
		for _, v := range w.Verbatim {
			if unicode.IsUpper(v) {
				count++
			}
		}
		if count > 2 {
			w.Normalized = str.FixAllCaps(w.Normalized)
			p.addWarn(parsed.AuthUpperCaseWarn)
		}
	}
	return w
}

type yearNode struct {
	Word        *parsed.Word
	Approximate bool
}

func (p *Engine) newYearNode(nd *node32) *yearNode {
	var w *parsed.Word
	appr := false
	nodes := nd.flatChildren()
	for _, v := range nodes {
		switch v.pegRule {
		case ruleYearWithPage:
			p.addWarn(parsed.YearPageWarn)
		case ruleYearRange:
			p.addWarn(parsed.YearRangeWarn)
			appr = true
		case ruleYearWithParens:
			p.addWarn(parsed.YearParensWarn)
			appr = true
		case ruleYearApprox:
			p.addWarn(parsed.YearSqBracketsWarn)
			appr = true
		case ruleYearWithChar:
			p.addWarn(parsed.YearCharWarn)
			w = p.newWordNode(v, parsed.YearType)
			w.Normalized = w.Verbatim[0 : len(w.Verbatim)-1]
		case ruleYearNum:
			if w == nil {
				w = p.newWordNode(v, parsed.YearType)
			}
			if w.Verbatim[len(w.Verbatim)-1] == '?' {
				p.addWarn(parsed.YearQuestionWarn)
				appr = true
			}
		}
	}
	if w == nil {
		w = p.newWordNode(nd, parsed.YearType)
	}
	if appr {
		w.Type = parsed.YearApproximateType
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

func (p *Engine) newWordNode(n *node32, wt parsed.WordType) *parsed.Word {
	t := n.token32
	val := p.nodeValue(n)
	norm := val
	if n.pegRule == ruleComparison {
		norm = "cf."
	}
	wrd := parsed.Word{
		Verbatim:   val,
		Normalized: norm,
		Type:       wt,
		Start:      int(t.begin),
		End:        int(t.end),
	}
	children := n.flatChildren()
	var canonicalApostrophe bool
	for _, v := range children {
		switch v.pegRule {
		case ruleDotPrefix:
			p.addWarn(parsed.DotEpithetWarn)
			wrd.Normalized = str.Normalize(wrd.Verbatim)
		case ruleUpperCharExtended, ruleLowerCharExtended:
			if wt == parsed.AuthorWordType {
				// this can only happen if word looked like botanical uninomial with
				// parent, but happen to be an author.
			} else if p.preserveDiaereses {
				wrd.Normalized = str.NormalizePreservingDiaereses(wrd.Verbatim)
			} else {
				if wt != parsed.AuthorWordType {
					p.addWarn(parsed.CharBadWarn)
				}
				wrd.Normalized = str.Normalize(wrd.Verbatim)
			}
		case ruleWordApostr:
			p.addWarn(parsed.CanonicalApostropheWarn)
			canonicalApostrophe = true
			wrd.Normalized = str.Normalize(wrd.Verbatim)
		case ruleWordStartsWithDigit:
			p.addWarn(parsed.SpeciesNumericWarn)
			wrd.Normalized = normalizeNums(wrd.Verbatim)
		case ruleApostrOther:
			p.addWarn(parsed.ApostrOtherWarn)
			if !canonicalApostrophe {
				nv := str.ToASCII(wrd.Verbatim, str.GlobalTransliterations)
				wrd.Normalized = nv
			}
		case ruleDashOther:
			p.addWarn(parsed.DashOtherWarn)
			wrd.Normalized = normalizeDashes(wrd.Verbatim)
		}
	}

	if wt == parsed.HybridCharType {
		wrd.Normalized = "×"
	} else if wt == parsed.GraftChimeraCharType {
		wrd.Normalized = "+"
	} else if wt == parsed.GenusType || wt == parsed.UninomialType {
		if val[len(val)-1] == '?' {
			p.addWarn(parsed.CapWordQuestionWarn)
			wrd.Normalized = wrd.Normalized[0 : len(wrd.Normalized)-1]
		}
		if _, ok := p.warnings[parsed.GenusUpperCharAfterDash]; ok {
			runes := []rune(wrd.Verbatim)
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
			wrd.Normalized = string(nv)
		}
		p.isBacteria(wrd.Normalized)
	}
	return &wrd
}

func (p *Engine) newCultivarEpithetNode(n *node32, wt parsed.WordType) *cultivarEpithetNode {
	t := n.token32
	val := p.nodeValue(n)
	normval := "‘" + p.nodeValue(n) + "’"
	wrd := parsed.Word{
		Verbatim:   val,
		Normalized: normval,
		Type:       wt,
		Start:      int(t.begin),
		End:        int(t.end),
	}
	cv := cultivarEpithetNode{Word: &wrd, enableCultivars: p.enableCultivars}
	if !p.enableCultivars {
		p.addWarn(parsed.CultivarEpithetWarn)
	}
	return &cv
}

var numWord = regexp.MustCompile(`^([0-9]+)[-\.]?(.+)$`)

func normalizeDashes(s string) string {
	return strings.ReplaceAll(s, "‑", "-")
}

func normalizeNums(s string) string {
	res := s
	match := numWord.FindAllStringSubmatch(s, 1)
	if len(match) == 0 {
		return res
	}
	num := match[0][1]
	wrd := match[0][2]
	return str.NumToStr(num) + wrd
}
