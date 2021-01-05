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

func (sn *scientificNameNode) Pos() []o.Position {
	return sn.nameData.pos()
}

func (sn *scientificNameNode) Normalized() string {
	if sn.nameData == nil {
		return ""
	}
	return sn.nameData.value()
}

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

func (sn *scientificNameNode) Details() o.Details {
	if sn.nameData == nil {
		return nil
	}
	return sn.nameData.details()
}

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

func (nf *hybridFormulaNode) pos() []o.Position {
	pos := nf.FirstSpecies.pos()
	for _, v := range nf.HybridElements {
		pos = append(pos, v.HybridChar.Pos)
		if v.Species != nil {
			pos = append(pos, v.Species.pos()...)
		}
	}
	return pos
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

func (nh *namedGenusHybridNode) pos() []o.Position {
	pos := []o.Position{nh.Hybrid.Pos}
	pos = append(pos, nh.nameData.pos()...)
	return pos
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

func (nh *namedSpeciesHybridNode) pos() []o.Position {
	pos := []o.Position{nh.Genus.Pos}
	if nh.Comparison != nil {
		pos = append(pos, nh.Comparison.Pos)
	}
	pos = append(pos, nh.Hybrid.Pos)
	pos = append(pos, nh.SpEpithet.pos()...)

	for _, v := range nh.InfraSpecies {
		pos = append(pos, v.pos()...)
	}
	return pos
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

func (apr *approxNode) pos() []o.Position {
	var pos []o.Position
	if apr == nil {
		return pos
	}
	pos = append(pos, apr.Genus.Pos)
	if apr.SpEpithet != nil {
		pos = append(pos, apr.SpEpithet.pos()...)
	}
	if apr.Approx != nil {
		pos = append(pos, apr.Approx.Pos)
	}
	return pos
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

func (comp *comparisonNode) pos() []o.Position {
	var pos []o.Position
	if comp == nil {
		return nil
	}
	pos = []o.Position{comp.Genus.Pos}
	pos = append(pos, comp.Comparison.Pos)
	if comp.SpEpithet != nil {
		pos = append(pos, comp.SpEpithet.pos()...)
	}
	return pos
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

func (sp *speciesNode) pos() []o.Position {
	var pos []o.Position
	if sp.Genus.Pos.End != 0 {
		pos = append(pos, sp.Genus.Pos)
	}
	if sp.SubGenus != nil {
		pos = append(pos, sp.SubGenus.Pos)
	}
	pos = append(pos, sp.SpEpithet.pos()...)
	for _, v := range sp.InfraSpecies {
		pos = append(pos, v.pos()...)
	}
	return pos
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

func (sep *spEpithetNode) pos() []o.Position {
	pos := []o.Position{sep.Word.Pos}
	pos = append(pos, sep.Authorship.pos()...)
	return pos
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

func (inf *infraspEpithetNode) pos() []o.Position {
	var pos []o.Position

	if inf.Rank != nil && inf.Rank.Word.Pos.Start != 0 {
		pos = append(pos, inf.Rank.Word.Pos)
	}
	pos = append(pos, inf.Word.Pos)
	if inf.Authorship != nil {
		pos = append(pos, inf.Authorship.pos()...)
	}
	return pos
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

func (u *uninomialNode) pos() []o.Position {
	pos := []o.Position{u.Word.Pos}
	pos = append(pos, u.Authorship.pos()...)
	return pos
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
	ud := o.Uninomial{Uninomial: u.Word.NormValue}
	if u.Authorship != nil {
		ud.Authorship = u.Authorship.details()
	}
	uo := o.DetailsUninomial{Uninomial: ud}
	return uo
}

func (u *uninomialComboNode) pos() []o.Position {
	pos := []o.Position{u.Uninomial1.Word.Pos}
	pos = append(pos, u.Uninomial1.Authorship.pos()...)
	if u.Rank.Word.Pos.Start != 0 {
		pos = append(pos, u.Rank.Word.Pos)
	}
	pos = append(pos, u.Uninomial2.Word.Pos)
	pos = append(pos, u.Uninomial2.Authorship.pos()...)
	return pos
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
		Uninomial: u.Uninomial2.Word.NormValue,
		Rank:      u.Rank.Word.NormValue,
		Parent:    u.Uninomial1.Word.NormValue,
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

func (a *authorshipNode) pos() []o.Position {
	if a == nil {
		var p []o.Position
		return p
	}
	p := a.OriginalAuthors.pos()
	return append(p, a.CombinationAuthors.pos()...)
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

func (ag *authorsGroupNode) pos() []o.Position {
	if ag == nil {
		var p []o.Position
		return p
	}
	p := ag.Team1.pos()
	return append(p, ag.Team2.pos()...)
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

func (aut *authorsTeamNode) pos() []o.Position {
	var res []o.Position
	if aut == nil {
		return res
	}
	for _, v := range aut.Authors {
		res = append(res, v.pos()...)
	}
	if aut.Year != nil {
		res = append(res, aut.Year.Word.Pos)
	}
	return res
}

func (aun *authorNode) pos() []o.Position {
	p := make([]o.Position, len(aun.Words))
	for i, v := range aun.Words {
		p[i] = v.Pos
	}
	return p
}
