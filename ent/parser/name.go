package parser

import (
	"fmt"

	"github.com/gnames/gnparser/ent/parsed"
	"github.com/gnames/gnparser/ent/stemmer"
	"github.com/gnames/gnparser/ent/str"
)

type canonical struct {
	Value       string
	ValueRanked string
}

func appendCanonical(c1 *canonical, c2 *canonical, sep string) *canonical {
	return &canonical{
		Value:       str.JoinStrings(c1.Value, c2.Value, sep),
		ValueRanked: str.JoinStrings(c1.ValueRanked, c2.ValueRanked, sep),
	}
}

// Words returns a slice of output.Word objects, where each element
// contains the value of the word, its semantic meaning and its
// position in the string.
func (sn *scientificNameNode) Words() []parsed.Word {
	return sn.words()
}

// Normalized returns a normalized version of a scientific name.
func (sn *scientificNameNode) Normalized() string {
	if sn.nameData == nil {
		return ""
	}
	return sn.value()
}

// Canonical returns canonical forms of scientific name. There are
// three forms: Stemmed, the most normalized, Simple, and Full (the least
// normalized).
func (sn *scientificNameNode) Canonical() *parsed.Canonical {
	var res *parsed.Canonical
	if sn.nameData == nil {
		return res
	}
	c := sn.canonical()
	return &parsed.Canonical{
		Stemmed: stemmer.StemCanonical(c.Value),
		Simple:  c.Value,
		Full:    c.ValueRanked,
	}
}

// Details returns additional details of about a scientific names.
// This function is called only if config.Config.WithDetails is true.
func (sn *scientificNameNode) Details() parsed.Details {
	if sn.nameData == nil {
		return nil
	}
	return sn.details()
}

// LastAuthorship returns the authorshop of the smallest element of a name.
// For example for a variation, it returns the authors of the variation, and
// ignores authors of genus, species etc.
func (sn *scientificNameNode) LastAuthorship(withDetails bool) *parsed.Authorship {
	var ao *parsed.Authorship
	if sn.nameData == nil {
		return ao
	}
	an := sn.lastAuthorship()
	if an == nil {
		return ao
	}
	res := an.details()
	if !withDetails {
		res.Original = nil
		res.Combination = nil
	}
	return res
}

func (nf *hybridFormulaNode) words() []parsed.Word {
	words := nf.FirstSpecies.words()
	for _, v := range nf.HybridElements {
		words = append(words, *v.HybridChar)
		if v.Species != nil {
			words = append(words, v.Species.words()...)
		}
	}
	return words
}

func (nf *hybridFormulaNode) value() string {
	val := nf.FirstSpecies.value()
	for _, v := range nf.HybridElements {
		val = str.JoinStrings(val, v.HybridChar.Normalized, " ")
		if v.Species != nil {
			val = str.JoinStrings(val, v.Species.value(), " ")
		}
	}
	return val
}

func (nf *hybridFormulaNode) canonical() *canonical {
	c := nf.FirstSpecies.canonical()
	for _, v := range nf.HybridElements {
		hc := &canonical{
			Value:       v.HybridChar.Normalized,
			ValueRanked: v.HybridChar.Normalized,
		}
		c = appendCanonical(c, hc, " ")
		if v.Species != nil {
			sc := v.Species.canonical()
			c = appendCanonical(c, sc, " ")
		}
	}
	return c
}

func (nf *hybridFormulaNode) lastAuthorship() *authorshipNode {
	var au *authorshipNode
	return au
}

func (nf *hybridFormulaNode) details() parsed.Details {
	dets := make([]parsed.Details, 0, len(nf.HybridElements)+1)
	dets = append(dets, nf.FirstSpecies.details())
	for _, v := range nf.HybridElements {
		if v.Species != nil {
			dets = append(dets, v.Species.details())
		}
	}
	return parsed.DetailsHybridFormula{HybridFormula: dets}
}

func (nf *graftChimeraFormulaNode) words() []parsed.Word {
	words := nf.FirstSpecies.words()
	for _, v := range nf.GraftChimeraElements {
		words = append(words, *v.GraftChimeraChar)
		if v.Species != nil {
			words = append(words, v.Species.words()...)
		}
	}
	return words
}

func (nf *graftChimeraFormulaNode) value() string {
	val := nf.FirstSpecies.value()
	for _, v := range nf.GraftChimeraElements {
		val = str.JoinStrings(val, v.GraftChimeraChar.Normalized, " ")
		if v.Species != nil {
			val = str.JoinStrings(val, v.Species.value(), " ")
		}
	}
	return val
}

func (nf *graftChimeraFormulaNode) canonical() *canonical {
	c := nf.FirstSpecies.canonical()
	for _, v := range nf.GraftChimeraElements {
		hc := &canonical{
			Value:       v.GraftChimeraChar.Normalized,
			ValueRanked: v.GraftChimeraChar.Normalized,
		}
		c = appendCanonical(c, hc, " ")
		if v.Species != nil {
			sc := v.Species.canonical()
			c = appendCanonical(c, sc, " ")
		}
	}
	return c
}

func (nf *graftChimeraFormulaNode) lastAuthorship() *authorshipNode {
	var au *authorshipNode
	return au
}

func (nf *graftChimeraFormulaNode) details() parsed.Details {
	dets := make([]parsed.Details, 0, len(nf.GraftChimeraElements)+1)
	dets = append(dets, nf.FirstSpecies.details())
	for _, v := range nf.GraftChimeraElements {
		if v.Species != nil {
			dets = append(dets, v.Species.details())
		}
	}
	return parsed.DetailsGraftChimeraFormula{GraftChimeraFormula: dets}
}

func (nh *namedGenusHybridNode) words() []parsed.Word {
	words := []parsed.Word{*nh.Hybrid}
	words = append(words, nh.nameData.words()...)
	return words
}

func (nh *namedGenusHybridNode) value() string {
	v := nh.nameData.value()
	v = "× " + v
	return v
}

func (nh *namedGenusHybridNode) canonical() *canonical {
	c := &canonical{
		Value:       "",
		ValueRanked: "×",
	}

	c1 := nh.nameData.canonical()
	c = appendCanonical(c, c1, " ")
	return c
}

func (nh *namedGenusHybridNode) details() parsed.Details {
	d := nh.nameData.details()
	return d
}

func (nh *namedGenusHybridNode) lastAuthorship() *authorshipNode {
	au := nh.nameData.lastAuthorship()
	return au
}

func (nh *namedSpeciesHybridNode) words() []parsed.Word {
	var wrd parsed.Word
	wrd = *nh.Genus
	words := []parsed.Word{wrd}
	if nh.Comparison != nil {
		wrd = *nh.Comparison
		words = append(words, wrd)
	}
	wrd = *nh.Hybrid
	words = append(words, wrd)
	words = append(words, nh.SpEpithet.words()...)

	for _, v := range nh.Infraspecies {
		words = append(words, v.words()...)
	}
	return words
}

func (nh *namedSpeciesHybridNode) value() string {
	res := nh.Genus.Normalized
	res = res + " × " + nh.SpEpithet.value()
	for _, v := range nh.Infraspecies {
		res = str.JoinStrings(res, v.value(), " ")
	}
	return res
}

func (nh *namedSpeciesHybridNode) canonical() *canonical {
	g := nh.Genus.Normalized
	c := &canonical{Value: g, ValueRanked: g}
	hCan := &canonical{Value: "", ValueRanked: "×"}
	c = appendCanonical(c, hCan, " ")
	cSp := nh.SpEpithet.canonical()
	c = appendCanonical(c, cSp, " ")

	for _, v := range nh.Infraspecies {
		c1 := v.canonical()
		c = appendCanonical(c, c1, " ")
	}
	return c
}

func (nh *namedSpeciesHybridNode) lastAuthorship() *authorshipNode {
	if len(nh.Infraspecies) == 0 {
		return nh.SpEpithet.Authorship
	}
	return nh.Infraspecies[len(nh.Infraspecies)-1].Authorship
}

func (nh *namedSpeciesHybridNode) details() parsed.Details {
	g := nh.Genus.Normalized
	so := parsed.Species{
		Genus:   g,
		Species: nh.SpEpithet.value(),
	}
	if nh.SpEpithet.Authorship != nil {
		so.Authorship = nh.SpEpithet.Authorship.details()
	}

	if len(nh.Infraspecies) == 0 {
		return parsed.DetailsSpecies{Species: so}
	}
	infs := make([]parsed.InfraspeciesElem, 0, len(nh.Infraspecies))
	for _, v := range nh.Infraspecies {
		if v == nil {
			continue
		}
		infs = append(infs, v.details())
	}
	iso := parsed.Infraspecies{
		Species:      so,
		Infraspecies: infs,
	}

	return parsed.DetailsInfraspecies{Infraspecies: iso}
}

func (nc *namedGenusGraftChimeraNode) words() []parsed.Word {
	words := []parsed.Word{*nc.GraftChimera}
	words = append(words, nc.nameData.words()...)
	return words
}

func (nc *namedGenusGraftChimeraNode) value() string {
	v := nc.nameData.value()
	v = "+ " + v
	return v
}

func (nc *namedGenusGraftChimeraNode) canonical() *canonical {
	c := &canonical{
		Value:       "",
		ValueRanked: "+",
	}

	c1 := nc.nameData.canonical()
	c = appendCanonical(c, c1, " ")
	return c
}

func (nc *namedGenusGraftChimeraNode) details() parsed.Details {
	d := nc.nameData.details()
	return d
}

func (nc *namedGenusGraftChimeraNode) lastAuthorship() *authorshipNode {
	au := nc.nameData.lastAuthorship()
	return au
}

func (cnd *candidatusNameNode) words() []parsed.Word {
	wrd := *cnd.Candidatus
	words := []parsed.Word{wrd}
	words = append(words, cnd.SingleName.words()...)
	return words
}

func (cnd *candidatusNameNode) value() string {
	val := cnd.Candidatus.Normalized
	val = str.JoinStrings(val, cnd.SingleName.value(), " ")
	return val
}

func (cnd *candidatusNameNode) canonical() *canonical {
	var c *canonical
	if cnd == nil {
		return c
	}
	can := cnd.SingleName.canonical()
	c = &canonical{
		Value:       can.Value,
		ValueRanked: "Candidatus " + can.ValueRanked,
	}
	return c
}

func (cnd *candidatusNameNode) lastAuthorship() *authorshipNode {
	return cnd.SingleName.lastAuthorship()
}

func (cnd *candidatusNameNode) details() parsed.Details {
	return cnd.SingleName.details()
}

func (apr *approxNode) words() []parsed.Word {
	var words []parsed.Word
	var wrd parsed.Word
	if apr == nil {
		return words
	}
	wrd = *apr.Genus
	words = append(words, wrd)
	if apr.SpEpithet != nil {
		words = append(words, apr.SpEpithet.words()...)
	}
	if apr.Approx != nil {
		wrd = *apr.Approx
		words = append(words, wrd)
	}
	return words
}

func (apr *approxNode) value() string {
	if apr == nil {
		return ""
	}
	val := apr.Genus.Normalized
	if apr.SpEpithet != nil {
		val = str.JoinStrings(val, apr.SpEpithet.value(), " ")
	}
	return val
}

func (apr *approxNode) canonical() *canonical {
	var c *canonical
	if apr == nil {
		return c
	}
	c = &canonical{
		Value:       apr.Genus.Normalized,
		ValueRanked: apr.Genus.Normalized,
	}
	if apr.SpEpithet != nil {
		spCan := apr.SpEpithet.canonical()
		c = appendCanonical(c, spCan, " ")
	}
	return c
}

func (apr *approxNode) lastAuthorship() *authorshipNode {
	var au *authorshipNode
	if apr == nil || apr.SpEpithet == nil {
		return au
	}
	return apr.SpEpithet.Authorship
}

func (apr *approxNode) details() parsed.Details {
	if apr == nil {
		return nil
	}
	ao := parsed.Approximation{
		Genus:        apr.Genus.Normalized,
		ApproxMarker: apr.Approx.Normalized,
		Ignored:      apr.Ignored,
	}
	if apr.SpEpithet == nil {
		return parsed.DetailsApproximation{Approximation: ao}
	}
	ao.Species = apr.SpEpithet.Word.Normalized

	if apr.SpEpithet.Authorship != nil {
		ao.SpeciesAuthorship = apr.SpEpithet.Authorship.details()
	}
	return parsed.DetailsApproximation{Approximation: ao}
}

func (comp *comparisonNode) words() []parsed.Word {
	var words []parsed.Word
	var wrd parsed.Word
	if comp == nil {
		return nil
	}
	wrd = *comp.Genus
	words = []parsed.Word{wrd}
	if comp.Cardinality == 2 {
		words = append(words, *comp.Comparison)
	}
	if comp.SpEpithet != nil {
		words = append(words, comp.SpEpithet.words()...)
	}
	if comp.Cardinality == 3 {
		words = append(words, *comp.Comparison)
	}
	if comp.InfraSpEpithet != nil {
		words = append(words, comp.InfraSpEpithet.words()...)
	}
	return words
}

func (comp *comparisonNode) value() string {
	if comp == nil {
		return ""
	}
	val := comp.Genus.Normalized
	if comp.Cardinality == 2 {
		val = str.JoinStrings(val, comp.Comparison.Normalized, " ")
	}
	if comp.SpEpithet != nil {
		val = str.JoinStrings(val, comp.SpEpithet.value(), " ")
	}
	if comp.Cardinality == 3 {
		val = str.JoinStrings(val, comp.Comparison.Normalized, " ")
	}
	if comp.InfraSpEpithet != nil {
		val = str.JoinStrings(val, comp.InfraSpEpithet.value(), " ")
	}
	return val
}

func (comp *comparisonNode) canonical() *canonical {
	if comp == nil {
		return &canonical{}
	}
	gen := comp.Genus.Normalized
	c := &canonical{Value: gen, ValueRanked: gen}
	if comp.SpEpithet != nil {
		sCan := comp.SpEpithet.canonical()
		c = appendCanonical(c, sCan, " ")
	}
	if comp.InfraSpEpithet != nil {
		ispCan := comp.InfraSpEpithet.canonical()
		c = appendCanonical(c, ispCan, " ")

	}
	return c
}

func (comp *comparisonNode) lastAuthorship() *authorshipNode {
	var au *authorshipNode
	if comp.Cardinality == 2 && comp.SpEpithet != nil {
		return comp.SpEpithet.Authorship
	}
	if comp.Cardinality == 3 && comp.InfraSpEpithet != nil {
		return comp.InfraSpEpithet.Authorship
	}
	return au
}

func (comp *comparisonNode) details() parsed.Details {
	if comp == nil {
		return nil
	}
	co := parsed.Comparison{
		Genus:      comp.Genus.Normalized,
		CompMarker: comp.Comparison.Normalized,
	}
	if comp.SpEpithet != nil {
		co.Species = &parsed.Species{
			Genus:      comp.Genus.Normalized,
			Species:    comp.SpEpithet.Word.Normalized,
			Authorship: comp.SpEpithet.Authorship.details(),
		}
	}
	if comp.InfraSpEpithet != nil {
		co.InfraSpecies = &parsed.InfraspeciesElem{
			Value:      comp.InfraSpEpithet.Word.Normalized,
			Authorship: comp.InfraSpEpithet.Authorship.details(),
		}
		if comp.InfraSpEpithet.Rank != nil {
			co.InfraSpecies.Rank = comp.InfraSpEpithet.Rank.Word.Normalized
		}

	}

	return parsed.DetailsComparison{Comparison: co}
}

func (sp *speciesNode) words() []parsed.Word {
	var words []parsed.Word
	var wrd parsed.Word
	if sp.Genus.End != 0 {
		wrd = *sp.Genus
		words = append(words, wrd)
	}
	if sp.Subgenus != nil {
		wrd = *sp.Subgenus
		words = append(words, wrd)
	}
	words = append(words, sp.SpEpithet.words()...)
	for _, v := range sp.Infraspecies {
		words = append(words, v.words()...)
	}
	if sp.CultivarEpithet != nil {
		wrd = *sp.CultivarEpithet.Word
		words = append(words, wrd)
	}
	return words
}

func (sp *speciesNode) value() string {
	gen := sp.Genus.Normalized
	sgen := ""
	if sp.Subgenus != nil {
		sgen = "(" + sp.Subgenus.Normalized + ")"
	}
	res := str.JoinStrings(gen, sgen, " ")
	res = str.JoinStrings(res, sp.SpEpithet.value(), " ")
	for _, v := range sp.Infraspecies {
		res = str.JoinStrings(res, v.value(), " ")
	}
	if sp.CultivarEpithet != nil && sp.CultivarEpithet.enableCultivars {
		res = str.JoinStrings(res, sp.CultivarEpithet.Word.Normalized, " ")
	}
	return res
}

func (sp *speciesNode) canonical() *canonical {
	spPart := str.JoinStrings(
		sp.Genus.Normalized,
		sp.SpEpithet.Word.Normalized,
		" ",
	)
	c := &canonical{Value: spPart, ValueRanked: spPart}
	for _, v := range sp.Infraspecies {
		c1 := v.canonical()
		c = appendCanonical(c, c1, " ")
	}
	if sp.CultivarEpithet != nil && sp.CultivarEpithet.enableCultivars {
		c2 := &canonical{
			Value:       sp.CultivarEpithet.Word.Normalized,
			ValueRanked: sp.CultivarEpithet.Word.Normalized,
		}
		c = appendCanonical(c, c2, " ")
	}
	return c
}

func (sp *speciesNode) lastAuthorship() *authorshipNode {
	if len(sp.Infraspecies) == 0 {
		return sp.SpEpithet.Authorship
	}
	return sp.Infraspecies[len(sp.Infraspecies)-1].Authorship
}

func (sp *speciesNode) details() parsed.Details {
	so := parsed.Species{
		Genus:   sp.Genus.Normalized,
		Species: sp.SpEpithet.Word.Normalized,
	}
	if sp.CultivarEpithet != nil {
		so.Cultivar = sp.CultivarEpithet.Word.Normalized
	}
	if sp.SpEpithet.Authorship != nil {
		so.Authorship = sp.SpEpithet.Authorship.details()
	}

	if sp.Subgenus != nil {
		so.Subgenus = sp.Subgenus.Normalized
	}
	if len(sp.Infraspecies) == 0 {
		return parsed.DetailsSpecies{Species: so}
	}
	infs := make([]parsed.InfraspeciesElem, 0, len(sp.Infraspecies))
	for _, v := range sp.Infraspecies {
		if v == nil {
			continue
		}
		infs = append(infs, v.details())
	}
	sio := parsed.Infraspecies{
		Species:      so,
		Infraspecies: infs,
	}

	return parsed.DetailsInfraspecies{Infraspecies: sio}
}

func (sep *spEpithetNode) words() []parsed.Word {
	wrd := *sep.Word
	words := []parsed.Word{wrd}
	words = append(words, sep.Authorship.words()...)
	return words
}

func (sep *spEpithetNode) value() string {
	val := sep.Word.Normalized
	val = str.JoinStrings(val, sep.Authorship.value(), " ")
	return val
}

func (sep *spEpithetNode) canonical() *canonical {
	c := &canonical{
		Value:       sep.Word.Normalized,
		ValueRanked: sep.Word.Normalized,
	}
	return c
}

func (inf *infraspEpithetNode) words() []parsed.Word {
	var words []parsed.Word
	var wrd parsed.Word
	if inf.Rank != nil && inf.Rank.Word.Start != 0 {
		wrd = *inf.Rank.Word
		words = append(words, wrd)
	}
	wrd = *inf.Word
	words = append(words, wrd)
	if inf.Authorship != nil {
		words = append(words, inf.Authorship.words()...)
	}
	return words
}

func (inf *infraspEpithetNode) value() string {
	val := inf.Word.Normalized
	rank := ""
	if inf.Rank != nil {
		rank = inf.Rank.Word.Normalized
	}
	au := inf.Authorship.value()
	res := str.JoinStrings(rank, val, " ")
	res = str.JoinStrings(res, au, " ")
	return res
}

func (inf *infraspEpithetNode) canonical() *canonical {
	val := inf.Word.Normalized
	rank := ""
	if inf.Rank != nil {
		rank = inf.Rank.Word.Normalized
	}
	rankedVal := str.JoinStrings(rank, val, " ")
	c := canonical{
		Value:       val,
		ValueRanked: rankedVal,
	}
	return &c
}

func (inf *infraspEpithetNode) details() parsed.InfraspeciesElem {
	rank := ""
	if inf.Rank != nil && inf.Rank.Word != nil {
		rank = inf.Rank.Word.Normalized
	}
	res := parsed.InfraspeciesElem{
		Value:      inf.Word.Normalized,
		Rank:       rank,
		Authorship: inf.Authorship.details(),
	}
	return res
}

func (u *uninomialNode) words() []parsed.Word {
	wrd := *u.Word
	words := []parsed.Word{wrd}

	words = append(words, u.Authorship.words()...)

	if u.CultivarEpithet != nil {
		wrd = *u.CultivarEpithet.Word
		words = append(words, wrd)
	}

	return words
}

func (u *uninomialNode) value() string {
	res := str.JoinStrings(u.Word.Normalized, u.Authorship.value(), " ")
	if u.CultivarEpithet != nil && u.CultivarEpithet.enableCultivars {
		res = str.JoinStrings(res, u.CultivarEpithet.Word.Normalized, " ")
	}
	return res
}

func (u *uninomialNode) canonical() *canonical {
	c := &canonical{
		Value:       u.Word.Normalized,
		ValueRanked: u.Word.Normalized,
	}
	if u.CultivarEpithet != nil && u.CultivarEpithet.enableCultivars {
		c2 := &canonical{
			Value:       u.CultivarEpithet.Word.Normalized,
			ValueRanked: u.CultivarEpithet.Word.Normalized,
		}
		c = appendCanonical(c, c2, " ")
	}
	return c
}

func (u *uninomialNode) lastAuthorship() *authorshipNode {
	return u.Authorship
}

func (u *uninomialNode) details() parsed.Details {
	ud := parsed.Uninomial{Value: u.Word.Normalized}
	if u.Authorship != nil {
		ud.Authorship = u.Authorship.details()
	}
	if u.CultivarEpithet != nil {
		ud.Cultivar = u.CultivarEpithet.Word.Normalized
	}
	uo := parsed.DetailsUninomial{Uninomial: ud}
	return uo
}

func (u *uninomialComboNode) words() []parsed.Word {
	var wrd parsed.Word
	wrd = *u.Uninomial1.Word
	words := []parsed.Word{wrd}
	words = append(words, u.Uninomial1.Authorship.words()...)
	if u.Rank.Word.Start != 0 {
		wrd = *u.Rank.Word
		words = append(words, wrd)
	}
	wrd = *u.Uninomial2.Word
	words = append(words, wrd)
	words = append(words, u.Uninomial2.Authorship.words()...)
	return words
}

func (u *uninomialComboNode) value() string {
	vl := str.JoinStrings(
		u.Uninomial1.Word.Normalized,
		u.Rank.Word.Normalized,
		" ",
	)
	tail := str.JoinStrings(
		u.Uninomial2.Word.Normalized,
		u.Uninomial2.Authorship.value(),
		" ",
	)
	return str.JoinStrings(vl, tail, " ")
}

func (u *uninomialComboNode) canonical() *canonical {
	ranked := str.JoinStrings(
		u.Uninomial1.Word.Normalized,
		u.Rank.Word.Normalized,
		" ",
	)
	ranked = str.JoinStrings(ranked, u.Uninomial2.Word.Normalized, " ")

	c := canonical{
		Value:       u.Uninomial2.Word.Normalized,
		ValueRanked: ranked,
	}
	return &c
}

func (u *uninomialComboNode) lastAuthorship() *authorshipNode {
	return u.Uninomial2.Authorship
}

func (u *uninomialComboNode) details() parsed.Details {
	ud := parsed.Uninomial{
		Value:  u.Uninomial2.Word.Normalized,
		Rank:   u.Rank.Word.Normalized,
		Parent: u.Uninomial1.Word.Normalized,
	}
	if u.Uninomial2.Authorship != nil {
		ud.Authorship = u.Uninomial2.Authorship.details()
	}
	uo := parsed.DetailsUninomial{Uninomial: ud}
	return uo
}

func (au *authorshipNode) details() *parsed.Authorship {
	if au == nil {
		var ao *parsed.Authorship
		return ao
	}
	ao := parsed.Authorship{Verbatim: au.Verbatim, Normalized: au.value()}
	ao.Original = authGroupDetail(au.OriginalAuthors)

	if au.CombinationAuthors != nil {
		ao.Combination = authGroupDetail(au.CombinationAuthors)
	}
	yr := ""
	if ao.Original != nil && ao.Original.Year != nil {
		yr = ao.Original.Year.Value
		if ao.Original.Year.IsApproximate {
			yr = fmt.Sprintf("(%s)", yr)
		}
	}
	if ao.Original != nil && ao.Original.ExAuthors != nil &&
		ao.Original.ExAuthors.Year != nil && yr == "" {
		yr = ao.Original.ExAuthors.Year.Value
		if ao.Original.ExAuthors.Year.IsApproximate {
			yr = fmt.Sprintf("(%s)", yr)
		}
	}
	var aus []string
	if ao.Original != nil {
		aus = ao.Original.Authors
		if ao.Original.ExAuthors != nil {
			aus = append(aus, ao.Original.ExAuthors.Authors...)
		}
	}
	if ao.Combination != nil {
		aus = append(aus, ao.Combination.Authors...)
		if ao.Combination.ExAuthors != nil {
			aus = append(aus, ao.Combination.ExAuthors.Authors...)
		}
	}
	ao.Authors = str.Uniq(aus)
	ao.Year = yr
	return &ao
}

func authGroupDetail(ag *authorsGroupNode) *parsed.AuthGroup {
	var ago parsed.AuthGroup
	if ag == nil {
		return &ago
	}
	aus, yr := ag.Team1.details()
	ago = parsed.AuthGroup{
		Authors: aus,
		Year:    yr,
	}
	if ag.Team2 == nil {
		return &ago
	}
	aus, yr = ag.Team2.details()
	switch ag.Team2Type {
	case teamEx:
		eao := parsed.Authors{
			Authors: aus,
			Year:    yr,
		}
		ago.ExAuthors = &eao
	case teamEmend:
		eao := parsed.Authors{
			Authors: aus,
			Year:    yr,
		}
		ago.EmendAuthors = &eao
	}
	return &ago
}

func (a *authorshipNode) words() []parsed.Word {
	if a == nil {
		var p []parsed.Word
		return p
	}
	p := a.OriginalAuthors.words()
	return append(p, a.CombinationAuthors.words()...)
}

func (a *authorshipNode) value() string {
	if a == nil || a.OriginalAuthors == nil {
		return ""
	}

	v := a.OriginalAuthors.value()
	if a.OriginalAuthors.Parens {
		v = fmt.Sprintf("(%s)", v)
	}
	if a.CombinationAuthors == nil {
		return v
	}
	cav := a.CombinationAuthors.value()
	v = v + " " + cav
	return v
}

func (ag *authorsGroupNode) value() string {
	if ag == nil || ag.Team1 == nil {
		return ""
	}
	v := ag.Team1.value()
	if ag.Team2 == nil {
		return v
	}
	v = fmt.Sprintf("%s %s %s", v, ag.Team2Word.Normalized, ag.Team2.value())
	return v
}

func (ag *authorsGroupNode) words() []parsed.Word {
	if ag == nil {
		var p []parsed.Word
		return p
	}
	p := ag.Team1.words()
	return append(p, ag.Team2.words()...)
}

func (aut *authorsTeamNode) value() string {
	if aut == nil {
		return ""
	}
	values := make([]string, len(aut.Authors))
	if len(values) == 0 {
		return ""
	}
	value := aut.Authors[0].Value
	sep := aut.Authors[0].Sep
	for _, v := range aut.Authors[1:] {
		value = str.JoinStrings(value, v.Value, sep)
		sep = v.Sep
	}
	if aut.Year == nil {
		return value
	}

	yr := aut.Year.Word.Normalized
	if aut.Year.Approximate {
		yr = fmt.Sprintf("(%s)", yr)
	}
	value = str.JoinStrings(value, yr, " ")
	return value
}

func (at *authorsTeamNode) details() ([]string, *parsed.Year) {
	var yr *parsed.Year
	var aus []string
	if at == nil {
		return aus, yr
	}
	aus = make([]string, len(at.Authors))
	for i, v := range at.Authors {
		aus[i] = v.Value
	}
	if at.Year == nil {
		return aus, yr
	}
	yr = &parsed.Year{
		Value:         at.Year.Word.Normalized,
		IsApproximate: at.Year.Approximate,
	}
	return aus, yr
}

func (aut *authorsTeamNode) words() []parsed.Word {
	var res []parsed.Word
	if aut == nil {
		return res
	}
	for _, v := range aut.Authors {
		res = append(res, v.words()...)
	}
	if aut.Year != nil {
		wrd := *aut.Year.Word
		res = append(res, wrd)
	}
	return res
}

func (aun *authorNode) words() []parsed.Word {
	p := make([]parsed.Word, len(aun.Words))
	for i := range aun.Words {
		p[i] = *aun.Words[i]
	}
	return p
}
