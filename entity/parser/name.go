package parser

import (
	"fmt"

	o "github.com/gnames/gnparser/entity/output"
	"github.com/gnames/gnparser/entity/stemmer"
	"github.com/gnames/gnparser/entity/str"
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
func (sn *scientificNameNode) Words() []o.Word {
	return sn.nameData.words()
}

// Normalized returns a normalized version of a scientific name.
func (sn *scientificNameNode) Normalized() string {
	if sn.nameData == nil {
		return ""
	}
	return sn.nameData.value()
}

// Canonical returns canonical forms of scientific name. There are
// three forms: Stemmed, the most normalized, Simple, and Full (the least
// normalized).
func (sn *scientificNameNode) Canonical() *o.Canonical {
	var res *o.Canonical
	if sn.nameData == nil {
		return res
	}
	c := sn.nameData.canonical()
	return &o.Canonical{
		Stemmed: stemmer.StemCanonical(c.Value),
		Simple:  c.Value,
		Full:    c.ValueRanked,
	}
}

// Details returns additional details of about a scientific names.
// This function is called only if config.Config.WithDetails is true.
func (sn *scientificNameNode) Details() o.Details {
	if sn.nameData == nil {
		return nil
	}
	return sn.nameData.details()
}

// LastAuthorship returns the authorshop of the smallest element of a name.
// For example for a variation, it returns the authors of the variation, and
// ignores authors of genus, species etc.
func (sn *scientificNameNode) LastAuthorship(withDetails bool) *o.Authorship {
	var ao *o.Authorship
	if sn.nameData == nil {
		return ao
	}
	an := sn.nameData.lastAuthorship()
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

func (nf *hybridFormulaNode) words() []o.Word {
	words := nf.FirstSpecies.words()
	for _, v := range nf.HybridElements {
		words = append(words, v.HybridChar.Pos)
		if v.Species != nil {
			words = append(words, v.Species.words()...)
		}
	}
	return words
}

func (nf *hybridFormulaNode) value() string {
	val := nf.FirstSpecies.value()
	for _, v := range nf.HybridElements {
		val = str.JoinStrings(val, v.HybridChar.Value, " ")
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
			Value:       v.HybridChar.NormValue,
			ValueRanked: v.HybridChar.NormValue,
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

func (nf *hybridFormulaNode) details() o.Details {
	dets := make([]o.Details, 0, len(nf.HybridElements)+1)
	dets = append(dets, nf.FirstSpecies.details())
	for _, v := range nf.HybridElements {
		if v.Species != nil {
			dets = append(dets, v.Species.details())
		}
	}
	return o.DetailsHybridFormula{HybridFormula: dets}
}

func (nh *namedGenusHybridNode) words() []o.Word {
	words := []o.Word{nh.Hybrid.Pos}
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

func (nh *namedGenusHybridNode) details() o.Details {
	d := nh.nameData.details()
	return d
}

func (nh *namedGenusHybridNode) lastAuthorship() *authorshipNode {
	au := nh.nameData.lastAuthorship()
	return au
}

func (nh *namedSpeciesHybridNode) words() []o.Word {
	var wrd o.Word
	wrd = nh.Genus.Pos
	wrd.Verbatim = nh.Genus.Value
	wrd.Normalized = nh.Genus.NormValue
	words := []o.Word{wrd}
	if nh.Comparison != nil {
		wrd = nh.Comparison.Pos
		wrd.Verbatim = nh.Comparison.Value
		wrd.Normalized = nh.Comparison.NormValue
		words = append(words, wrd)
	}
	wrd = nh.Hybrid.Pos
	wrd.Verbatim = nh.Hybrid.Value
	wrd.Normalized = nh.Hybrid.NormValue
	words = append(words, wrd)
	words = append(words, nh.SpEpithet.words()...)

	for _, v := range nh.InfraSpecies {
		words = append(words, v.words()...)
	}
	return words
}

func (nh *namedSpeciesHybridNode) value() string {
	res := nh.Genus.NormValue
	res = res + " × " + nh.SpEpithet.value()
	for _, v := range nh.InfraSpecies {
		res = str.JoinStrings(res, v.value(), " ")
	}
	return res
}

func (nh *namedSpeciesHybridNode) canonical() *canonical {
	g := nh.Genus.NormValue
	c := &canonical{Value: g, ValueRanked: g}
	hCan := &canonical{Value: "", ValueRanked: "×"}
	c = appendCanonical(c, hCan, " ")
	cSp := nh.SpEpithet.canonical()
	c = appendCanonical(c, cSp, " ")

	for _, v := range nh.InfraSpecies {
		c1 := v.canonical()
		c = appendCanonical(c, c1, " ")
	}
	return c
}

func (nh *namedSpeciesHybridNode) lastAuthorship() *authorshipNode {
	if len(nh.InfraSpecies) == 0 {
		return nh.SpEpithet.Authorship
	}
	return nh.InfraSpecies[len(nh.InfraSpecies)-1].Authorship
}

func (nh *namedSpeciesHybridNode) details() o.Details {
	g := nh.Genus.NormValue
	so := o.Species{
		Genus:   g,
		Species: nh.SpEpithet.value(),
	}
	if nh.SpEpithet.Authorship != nil {
		so.Authorship = nh.SpEpithet.Authorship.details()
	}

	if len(nh.InfraSpecies) == 0 {
		return o.DetailsSpecies{Species: so}
	}
	infs := make([]o.InfraSpeciesElem, 0, len(nh.InfraSpecies))
	for _, v := range nh.InfraSpecies {
		if v == nil {
			continue
		}
		infs = append(infs, v.details())
	}
	iso := o.InfraSpecies{
		Species:      so,
		InfraSpecies: infs,
	}

	return o.DetailsInfraSpecies{InfraSpecies: iso}
}

func (apr *approxNode) words() []o.Word {
	var words []o.Word
	var wrd o.Word
	if apr == nil {
		return words
	}
	wrd = apr.Genus.Pos
	wrd.Verbatim = apr.Genus.Value
	wrd.Normalized = apr.Genus.NormValue
	words = append(words, wrd)
	if apr.SpEpithet != nil {
		words = append(words, apr.SpEpithet.words()...)
	}
	if apr.Approx != nil {
		wrd = apr.Approx.Pos
		wrd.Verbatim = apr.Approx.Value
		wrd.Normalized = apr.Approx.NormValue
		words = append(words, wrd)
	}
	return words
}

func (apr *approxNode) value() string {
	if apr == nil {
		return ""
	}
	val := apr.Genus.NormValue
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
	c = &canonical{Value: apr.Genus.NormValue, ValueRanked: apr.Genus.NormValue}
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

func (apr *approxNode) details() o.Details {
	if apr == nil {
		return nil
	}
	ao := o.Approximation{
		Genus:        apr.Genus.NormValue,
		ApproxMarker: apr.Approx.NormValue,
		Ignored:      apr.Ignored,
	}
	if apr.SpEpithet == nil {
		return o.DetailsApproximation{Approximation: ao}
	}
	ao.Species = apr.SpEpithet.Word.NormValue

	if apr.SpEpithet.Authorship != nil {
		ao.SpeciesAuthorship = apr.SpEpithet.Authorship.details()
	}
	return o.DetailsApproximation{Approximation: ao}
}

func (comp *comparisonNode) words() []o.Word {
	var words []o.Word
	var wrd o.Word
	if comp == nil {
		return nil
	}
	wrd = comp.Genus.Pos
	wrd.Verbatim = comp.Genus.Value
	wrd.Normalized = comp.Genus.NormValue
	words = []o.Word{wrd}
	wrd = comp.Comparison.Pos
	wrd.Verbatim = comp.Comparison.Value
	wrd.Normalized = comp.Comparison.NormValue
	words = append(words, wrd)
	if comp.SpEpithet != nil {
		words = append(words, comp.SpEpithet.words()...)
	}
	return words
}

func (comp *comparisonNode) value() string {
	if comp == nil {
		return ""
	}
	val := comp.Genus.NormValue
	val = str.JoinStrings(val, comp.Comparison.NormValue, " ")
	if comp.SpEpithet != nil {
		val = str.JoinStrings(val, comp.SpEpithet.value(), " ")
	}
	return val
}

func (comp *comparisonNode) canonical() *canonical {
	if comp == nil {
		return &canonical{}
	}
	gen := comp.Genus.NormValue
	c := &canonical{Value: gen, ValueRanked: gen}
	if comp.SpEpithet != nil {
		sCan := comp.SpEpithet.canonical()
		c = appendCanonical(c, sCan, " ")
	}
	return c
}

func (comp *comparisonNode) lastAuthorship() *authorshipNode {
	var au *authorshipNode
	if comp == nil || comp.SpEpithet == nil {
		return au
	}
	return comp.SpEpithet.Authorship
}

func (comp *comparisonNode) details() o.Details {
	if comp == nil {
		return nil
	}
	co := o.Comparison{
		Genus:      comp.Genus.NormValue,
		CompMarker: comp.Comparison.NormValue,
	}
	if comp.SpEpithet == nil {
		return o.DetailsComparison{Comparison: co}
	}

	co.Species = comp.SpEpithet.value()
	if comp.SpEpithet.Authorship != nil {
		co.SpeciesAuthorship = comp.SpEpithet.Authorship.details()
	}
	return o.DetailsComparison{Comparison: co}
}

func (sp *speciesNode) words() []o.Word {
	var words []o.Word
	var wrd o.Word
	if sp.Genus.Pos.End != 0 {
		wrd = sp.Genus.Pos
		wrd.Verbatim = sp.Genus.Value
		wrd.Normalized = sp.Genus.NormValue
		words = append(words, wrd)
	}
	if sp.SubGenus != nil {
		wrd = sp.SubGenus.Pos
		wrd.Verbatim = sp.SubGenus.Value
		wrd.Normalized = sp.SubGenus.NormValue
		words = append(words, wrd)
	}
	words = append(words, sp.SpEpithet.words()...)
	for _, v := range sp.InfraSpecies {
		words = append(words, v.words()...)
	}
	return words
}

func (sp *speciesNode) value() string {
	gen := sp.Genus.NormValue
	sgen := ""
	if sp.SubGenus != nil {
		sgen = "(" + sp.SubGenus.NormValue + ")"
	}
	res := str.JoinStrings(gen, sgen, " ")
	res = str.JoinStrings(res, sp.SpEpithet.value(), " ")
	for _, v := range sp.InfraSpecies {
		res = str.JoinStrings(res, v.value(), " ")
	}
	return res
}

func (sp *speciesNode) canonical() *canonical {
	spPart := str.JoinStrings(sp.Genus.NormValue, sp.SpEpithet.Word.NormValue, " ")
	c := &canonical{Value: spPart, ValueRanked: spPart}
	for _, v := range sp.InfraSpecies {
		c1 := v.canonical()
		c = appendCanonical(c, c1, " ")
	}
	return c
}

func (sp *speciesNode) lastAuthorship() *authorshipNode {
	if len(sp.InfraSpecies) == 0 {
		return sp.SpEpithet.Authorship
	}
	return sp.InfraSpecies[len(sp.InfraSpecies)-1].Authorship
}

func (sp *speciesNode) details() o.Details {
	so := o.Species{
		Genus:   sp.Genus.NormValue,
		Species: sp.SpEpithet.Word.NormValue,
	}
	if sp.SpEpithet.Authorship != nil {
		so.Authorship = sp.SpEpithet.Authorship.details()
	}

	if sp.SubGenus != nil {
		so.SubGenus = sp.SubGenus.NormValue
	}
	if len(sp.InfraSpecies) == 0 {
		return o.DetailsSpecies{Species: so}
	}
	infs := make([]o.InfraSpeciesElem, 0, len(sp.InfraSpecies))
	for _, v := range sp.InfraSpecies {
		if v == nil {
			continue
		}
		infs = append(infs, v.details())
	}
	sio := o.InfraSpecies{
		Species:      so,
		InfraSpecies: infs,
	}

	return o.DetailsInfraSpecies{InfraSpecies: sio}
}

func (sep *spEpithetNode) words() []o.Word {
	wrd := sep.Word.Pos
	wrd.Verbatim = sep.Word.Value
	wrd.Normalized = sep.Word.NormValue
	words := []o.Word{wrd}
	words = append(words, sep.Authorship.words()...)
	return words
}

func (sep *spEpithetNode) value() string {
	val := sep.Word.NormValue
	val = str.JoinStrings(val, sep.Authorship.value(), " ")
	return val
}

func (sep *spEpithetNode) canonical() *canonical {
	c := &canonical{Value: sep.Word.NormValue, ValueRanked: sep.Word.NormValue}
	return c
}

func (inf *infraspEpithetNode) words() []o.Word {
	var words []o.Word
	var wrd o.Word
	if inf.Rank != nil && inf.Rank.Word.Pos.Start != 0 {
		wrd = inf.Rank.Word.Pos
		wrd.Verbatim = inf.Rank.Word.Value
		wrd.Normalized = inf.Rank.Word.NormValue
		words = append(words, wrd)
	}
	wrd = inf.Word.Pos
	wrd.Verbatim = inf.Word.Value
	wrd.Normalized = inf.Word.NormValue
	words = append(words, wrd)
	if inf.Authorship != nil {
		words = append(words, inf.Authorship.words()...)
	}
	return words
}

func (inf *infraspEpithetNode) value() string {
	val := inf.Word.NormValue
	rank := ""
	if inf.Rank != nil {
		rank = inf.Rank.Word.NormValue
	}
	au := inf.Authorship.value()
	res := str.JoinStrings(rank, val, " ")
	res = str.JoinStrings(res, au, " ")
	return res
}

func (inf *infraspEpithetNode) canonical() *canonical {
	val := inf.Word.NormValue
	rank := ""
	if inf.Rank != nil {
		rank = inf.Rank.Word.NormValue
	}
	rankedVal := str.JoinStrings(rank, val, " ")
	c := canonical{
		Value:       val,
		ValueRanked: rankedVal,
	}
	return &c
}

func (inf *infraspEpithetNode) details() o.InfraSpeciesElem {
	rank := ""
	if inf.Rank != nil && inf.Rank.Word != nil {
		rank = inf.Rank.Word.NormValue
	}
	res := o.InfraSpeciesElem{
		Value:      inf.Word.NormValue,
		Rank:       rank,
		Authorship: inf.Authorship.details(),
	}
	return res
}

func (u *uninomialNode) words() []o.Word {
	wrd := u.Word.Pos
	wrd.Verbatim = u.Word.Value
	wrd.Normalized = u.Word.NormValue
	words := []o.Word{wrd}
	words = append(words, u.Authorship.words()...)
	return words
}

func (u *uninomialNode) value() string {
	return str.JoinStrings(u.Word.NormValue, u.Authorship.value(), " ")
}

func (u *uninomialNode) canonical() *canonical {
	c := canonical{Value: u.Word.NormValue, ValueRanked: u.Word.NormValue}
	return &c
}

func (u *uninomialNode) lastAuthorship() *authorshipNode {
	return u.Authorship
}

func (u *uninomialNode) details() o.Details {
	ud := o.Uninomial{Value: u.Word.NormValue}
	if u.Authorship != nil {
		ud.Authorship = u.Authorship.details()
	}
	uo := o.DetailsUninomial{Uninomial: ud}
	return uo
}

func (u *uninomialComboNode) words() []o.Word {
	var wrd o.Word
	wrd = u.Uninomial1.Word.Pos
	wrd.Verbatim = u.Uninomial1.Word.Value
	wrd.Normalized = u.Uninomial1.Word.NormValue
	words := []o.Word{wrd}
	words = append(words, u.Uninomial1.Authorship.words()...)
	if u.Rank.Word.Pos.Start != 0 {
		wrd = u.Rank.Word.Pos
		wrd.Verbatim = u.Rank.Word.Value
		wrd.Normalized = u.Rank.Word.NormValue
		words = append(words, wrd)
	}
	wrd = u.Uninomial2.Word.Pos
	wrd.Verbatim = u.Uninomial2.Word.Value
	wrd.Normalized = u.Uninomial2.Word.NormValue
	words = append(words, wrd)
	words = append(words, u.Uninomial2.Authorship.words()...)
	return words
}

func (u *uninomialComboNode) value() string {
	vl := str.JoinStrings(u.Uninomial1.Word.NormValue, u.Rank.Word.NormValue, " ")
	tail := str.JoinStrings(u.Uninomial2.Word.NormValue,
		u.Uninomial2.Authorship.value(), " ")
	return str.JoinStrings(vl, tail, " ")
}

func (u *uninomialComboNode) canonical() *canonical {
	ranked := str.JoinStrings(u.Uninomial1.Word.NormValue, u.Rank.Word.NormValue, " ")
	ranked = str.JoinStrings(ranked, u.Uninomial2.Word.NormValue, " ")

	c := canonical{
		Value:       u.Uninomial2.Word.NormValue,
		ValueRanked: ranked,
	}
	return &c
}

func (u *uninomialComboNode) lastAuthorship() *authorshipNode {
	return u.Uninomial2.Authorship
}

func (u *uninomialComboNode) details() o.Details {
	ud := o.Uninomial{
		Value:  u.Uninomial2.Word.NormValue,
		Rank:   u.Rank.Word.NormValue,
		Parent: u.Uninomial1.Word.NormValue,
	}
	if u.Uninomial2.Authorship != nil {
		ud.Authorship = u.Uninomial2.Authorship.details()
	}
	uo := o.DetailsUninomial{Uninomial: ud}
	return uo
}

func (au *authorshipNode) details() *o.Authorship {
	if au == nil {
		var ao *o.Authorship
		return ao
	}
	ao := o.Authorship{Verbatim: au.Verbatim, Normalized: au.value()}
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
	var aus []string
	if ao.Original != nil {
		aus = ao.Original.Authors
	}
	if ao.Combination != nil {
		aus = append(aus, ao.Combination.Authors...)
	}
	ao.Authors = aus
	ao.Year = yr
	return &ao
}

func authGroupDetail(ag *authorsGroupNode) *o.AuthGroup {
	var ago o.AuthGroup
	if ag == nil {
		return &ago
	}
	aus, yr := ag.Team1.details()
	ago = o.AuthGroup{
		Authors: aus,
		Year:    yr,
	}
	if ag.Team2 == nil {
		return &ago
	}
	aus, yr = ag.Team2.details()
	switch ag.Team2Type {
	case teamEx:
		eao := o.Authors{
			Authors: aus,
			Year:    yr,
		}
		ago.ExAuthors = &eao
	case teamEmend:
		eao := o.Authors{
			Authors: aus,
			Year:    yr,
		}
		ago.EmendAuthors = &eao
	}
	return &ago
}

func (a *authorshipNode) words() []o.Word {
	if a == nil {
		var p []o.Word
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
	v = fmt.Sprintf("%s %s %s", v, ag.Team2Word.NormValue, ag.Team2.value())
	return v
}

func (ag *authorsGroupNode) words() []o.Word {
	if ag == nil {
		var p []o.Word
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

	yr := aut.Year.Word.NormValue
	if aut.Year.Approximate {
		yr = fmt.Sprintf("(%s)", yr)
	}
	value = str.JoinStrings(value, yr, " ")
	return value
}

func (at *authorsTeamNode) details() ([]string, *o.Year) {
	var yr *o.Year
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
	yr = &o.Year{
		Value:         at.Year.Word.NormValue,
		IsApproximate: at.Year.Approximate,
	}
	return aus, yr
}

func (aut *authorsTeamNode) words() []o.Word {
	var res []o.Word
	if aut == nil {
		return res
	}
	for _, v := range aut.Authors {
		res = append(res, v.words()...)
	}
	if aut.Year != nil {
		wrd := aut.Year.Word.Pos
		wrd.Verbatim = aut.Year.Word.Value
		wrd.Normalized = aut.Year.Word.NormValue
		res = append(res, wrd)
	}
	return res
}

func (aun *authorNode) words() []o.Word {
	p := make([]o.Word, len(aun.Words))
	for i, v := range aun.Words {
		p[i] = v.Pos
		p[i].Verbatim = v.Value
		p[i].Normalized = v.NormValue
	}
	return p
}
