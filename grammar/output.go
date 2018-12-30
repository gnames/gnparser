package grammar

import (
	"fmt"
	"strings"

	"gitlab.com/gogna/gnparser/str"
)

type UninomialOutput struct {
	Uninomial *uniDetails `json:"uninomial"`
}

type uniDetails struct {
	Value      string            `json:"value"`
	Rank       string            `json:"rank,omitempty"`
	Parent     string            `json:"parent,omitempty"`
	Authorship *authorshipOutput `json:"authorship,omitempty"`
}

type SpeciesOutput struct {
	Genus        *genusOutput            `json:"genus"`
	SpecEpithet  *specEpithetOutput      `json:"specificEpithet"`
	SubGenus     *subGenusOutput         `json:"infragenericEpithet,omitempty"`
	InfraSpecies []*infraSpEpithetOutput `json:"infraspecificEpithets,omitempty"`
}

type genusOutput struct {
	Value string `json:"value"`
}

type subGenusOutput struct {
	Value string `json:"value"`
}
type specEpithetOutput struct {
	Value      string            `json:"value"`
	Authorship *authorshipOutput `json:"authorship,omitempty"`
}

type infraSpEpithetOutput struct {
	Value      string            `json:"value"`
	Rank       string            `json:"rank,omitempty"`
	Authorship *authorshipOutput `json:"authorship,omitempty"`
}

type authorshipOutput struct {
	Value       string           `json:"value"`
	Original    *authGroupOutput `json:"basionymAuthorship,omitempty"`
	Combination *authGroupOutput `json:"combinationAuthorship,omitempty"`
}

type authGroupOutput struct {
	Authors   []string         `json:"authors"`
	Years     []yearOutput     `json:"years,omitempty"`
	ExAuthors *exAuthorsOutput `json:"exAuthors,omitempty"`
}

type exAuthorsOutput struct {
	Authors []string     `json:"authors"`
	Years   []yearOutput `json:"years,omitempty"`
}

type yearOutput struct {
	Value       string `json:"value,omitempty"`
	Approximate bool   `json:"approximate,omitempty"`
}

type Canonical struct {
	Value       string
	ValueRanked string
}

func appendCanonical(c1 *Canonical, c2 *Canonical, sep string) *Canonical {
	return &Canonical{
		Value:       str.JoinStrings(c1.Value, c2.Value, sep),
		ValueRanked: str.JoinStrings(c1.ValueRanked, c2.ValueRanked, sep),
	}
}

func (sn *ScientificNameNode) Pos() []Pos {
	var pos []Pos
	for _, v := range sn.NamesGroup {
		pos = append(pos, v.pos()...)
	}
	return pos
}

func (sn *ScientificNameNode) Value() string {
	ng := sn.NamesGroup
	if len(ng) == 1 {
		return ng[0].value()
	}
	values := make([]string, len(ng))
	for i, v := range ng {
		values[i] = v.value()
	}
	return strings.Join(values, " × ")
}

func (sn *ScientificNameNode) Canonical() *Canonical {
	var cs *Canonical
	ng := sn.NamesGroup
	if len(ng) == 1 {
		return ng[0].canonical()
	}
	for _, v := range ng {
		cs = appendCanonical(cs, v.canonical(), " × ")
	}
	return cs
}

func (sn *ScientificNameNode) Details() []interface{} {
	res := make([]interface{}, len(sn.NamesGroup))
	for i, v := range sn.NamesGroup {
		res[i] = v.details()
	}
	return res
}

func (sn *ScientificNameNode) LastAuthorship() *authorshipOutput {
	if len(sn.NamesGroup) != 1 {
		var ao *authorshipOutput
		return ao
	}
	an := sn.NamesGroup[0].lastAuthorship()
	return an.details()
}

func (sp *speciesNode) pos() []Pos {
	pos := []Pos{sp.Genus.Pos}
	if sp.SubGenus != nil {
		pos = append(pos, sp.SubGenus.Pos)
	}
	pos = append(pos, sp.Species.Word.Pos)
	pos = append(pos, sp.Species.Authorship.pos()...)
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
	res = str.JoinStrings(res, sp.Species.Word.NormValue, " ")
	res = str.JoinStrings(res, sp.Species.Authorship.value(), " ")
	for _, v := range sp.InfraSpecies {
		res = str.JoinStrings(res, v.value(), " ")
	}
	return res
}

func (sp *speciesNode) canonical() *Canonical {
	spPart := str.JoinStrings(sp.Genus.NormValue, sp.Species.Word.NormValue, " ")
	c := &Canonical{Value: spPart, ValueRanked: spPart}
	for _, v := range sp.InfraSpecies {
		c = appendCanonical(c, v.canonical(), " ")
	}
	return c
}

func (sp *speciesNode) lastAuthorship() *authorshipNode {
	if len(sp.InfraSpecies) == 0 {
		return sp.Species.Authorship
	}
	return sp.InfraSpecies[len(sp.InfraSpecies)-1].Authorship
}

func (sp *speciesNode) details() interface{} {
	se := specEpithetOutput{
		Value: sp.Species.Word.NormValue,
	}
	if sp.Species.Authorship != nil {
		se.Authorship = sp.Species.Authorship.details()
	}

	g := sp.Genus.NormValue
	so := SpeciesOutput{
		Genus:       &genusOutput{Value: g},
		SpecEpithet: &se,
	}

	if sp.SubGenus != nil {
		sg := sp.SubGenus.NormValue
		so.SubGenus = &subGenusOutput{Value: sg}
	}
	if len(sp.InfraSpecies) == 0 {
		return &so
	}
	infs := make([]*infraSpEpithetOutput, len(sp.InfraSpecies))
	for i, v := range sp.InfraSpecies {
		infs[i] = v.details()
	}
	so.InfraSpecies = infs

	return &so
}

func (inf *infraspEpithetNode) pos() []Pos {
	var pos []Pos

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

func (inf *infraspEpithetNode) canonical() *Canonical {
	val := inf.Word.NormValue
	rank := ""
	if inf.Rank != nil {
		rank = inf.Rank.Word.NormValue
	}
	rankedVal := str.JoinStrings(rank, val, " ")
	c := Canonical{
		Value:       val,
		ValueRanked: rankedVal,
	}
	return &c
}

func (inf *infraspEpithetNode) details() *infraSpEpithetOutput {
	var info infraSpEpithetOutput
	if inf == nil {
		return &info
	}
	rank := ""
	if inf.Rank != nil && inf.Rank.Word != nil {
		rank = inf.Rank.Word.NormValue
	}
	info = infraSpEpithetOutput{
		Value:      inf.Word.NormValue,
		Rank:       rank,
		Authorship: inf.Authorship.details(),
	}
	return &info
}

func (u *uninomialNode) pos() []Pos {
	pos := []Pos{u.Word.Pos}
	pos = append(pos, u.Authorship.pos()...)
	return pos
}

func (u *uninomialNode) value() string {
	return str.JoinStrings(u.Word.NormValue, u.Authorship.value(), " ")
}

func (u *uninomialNode) canonical() *Canonical {
	c := Canonical{Value: u.Word.NormValue, ValueRanked: u.Word.NormValue}
	return &c
}

func (u *uninomialNode) lastAuthorship() *authorshipNode {
	return u.Authorship
}

func (u *uninomialNode) details() interface{} {
	ud := uniDetails{Value: u.Word.NormValue}
	if u.Authorship != nil {
		ud.Authorship = u.Authorship.details()
	}
	uo := UninomialOutput{Uninomial: &ud}
	return &uo
}

func (u *uninomialComboNode) pos() []Pos {
	pos := []Pos{u.Uninomial1.Word.Pos}
	pos = append(pos, u.Uninomial1.Authorship.pos()...)
	if u.Rank.Word.Pos.Start != 0 {
		pos = append(pos, u.Rank.Word.Pos)
	}
	pos = append(pos, u.Uninomial2.Word.Pos)
	pos = append(pos, u.Uninomial2.Authorship.pos()...)
	return pos
}

func (u *uninomialComboNode) value() string {
	vl := str.JoinStrings(u.Uninomial1.Word.NormValue, u.Rank.Word.Value, " ")
	tail := str.JoinStrings(u.Uninomial2.Word.NormValue,
		u.Uninomial2.Authorship.value(), " ")
	return str.JoinStrings(vl, tail, " ")
}

func (u *uninomialComboNode) canonical() *Canonical {
	ranked := str.JoinStrings(u.Uninomial1.Word.NormValue, u.Rank.Word.Value, " ")
	ranked = str.JoinStrings(ranked, u.Uninomial2.Word.NormValue, " ")

	c := Canonical{
		Value:       u.Uninomial2.Word.NormValue,
		ValueRanked: ranked,
	}
	return &c
}

func (u *uninomialComboNode) lastAuthorship() *authorshipNode {
	return u.Uninomial2.Authorship
}

func (u *uninomialComboNode) details() interface{} {
	ud := uniDetails{
		Value:  u.Uninomial2.Word.NormValue,
		Rank:   u.Rank.Word.Value,
		Parent: u.Uninomial1.Word.NormValue,
	}
	if u.Uninomial2.Authorship != nil {
		ud.Authorship = u.Uninomial2.Authorship.details()
	}
	uo := UninomialOutput{Uninomial: &ud}
	return &uo
}

func (au *authorshipNode) details() *authorshipOutput {
	if au == nil {
		var ao *authorshipOutput
		return ao
	}
	ao := authorshipOutput{Value: au.value()}
	ao.Original = authGroupDetail(au.OriginalAuthors)

	if au.CombinationAuthors != nil {
		ao.Combination = authGroupDetail(au.CombinationAuthors)
	}
	return &ao
}

func authGroupDetail(ag *authorsGroupNode) *authGroupOutput {
	var ago authGroupOutput
	if ag == nil {
		return &ago
	}
	aus, yrs := ag.Team1.details()
	ago = authGroupOutput{
		Authors: aus,
		Years:   yrs,
	}
	if ag.Team2 == nil {
		return &ago
	}

	ausEx, yrsEx := ag.Team2.details()
	eao := exAuthorsOutput{
		Authors: ausEx,
		Years:   yrsEx,
	}
	ago.ExAuthors = &eao
	return &ago
}

func (a *authorshipNode) pos() []Pos {
	if a == nil {
		var p []Pos
		return p
	}
	p := a.OriginalAuthors.pos()
	return append(p, a.CombinationAuthors.pos()...)
}

func (a *authorshipNode) value() string {
	if a == nil {
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
	v := ag.Team1.value()
	if ag.Team2 == nil {
		return v
	}
	v = fmt.Sprintf("%s ex %s", v, ag.Team2.value())
	return v
}

func (ag *authorsGroupNode) pos() []Pos {
	if ag == nil {
		var p []Pos
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
	for i, v := range aut.Authors {
		values[i] = v.Value
	}
	value := strings.Join(values[0:len(values)-1], ", ")
	value = str.JoinStrings(value, values[len(values)-1], " & ")
	if len(aut.Years) == 0 || aut.Years[0].Word.Value == "" {
		return value
	}

	years := make([]string, len(aut.Years))
	for i, v := range aut.Years {
		yr := v.Word.Value
		if v.Approximate {
			yr = fmt.Sprintf("(%s)", yr)
		}
		years[i] = yr
	}
	yrVal := strings.Join(years, ", ")
	value = str.JoinStrings(value, yrVal, " ")
	return value
}

func (at *authorsTeamNode) details() ([]string, []yearOutput) {
	var yrs []yearOutput
	var aus []string
	if at == nil {
		return aus, yrs
	}
	yrs = make([]yearOutput, len(at.Years))
	for i, v := range at.Years {
		yrs[i] = yearOutput{Value: v.Word.Value, Approximate: v.Approximate}
	}
	aus = make([]string, len(at.Authors))
	for i, v := range at.Authors {
		aus[i] = v.Value
	}
	return aus, yrs
}

func (aut *authorsTeamNode) pos() []Pos {
	var res []Pos
	if aut == nil {
		return res
	}
	for _, v := range aut.Authors {
		res = append(res, v.pos()...)
	}
	for _, v := range aut.Years {
		res = append(res, v.Word.Pos)
	}
	return res
}

func (aun *authorNode) pos() []Pos {
	p := make([]Pos, len(aun.Words))
	for i, v := range aun.Words {
		p[i] = v.Pos
	}
	return p
}
