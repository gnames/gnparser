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
	Genus       *genusOutput       `json:"genus"`
	SpecEpithet *specEpithetOutput `json:"specificEpithet"`
}

type genusOutput struct {
	Value string `json:"value"`
}

type specEpithetOutput struct {
	Value      string            `json:"value"`
	Authorship *authorshipOutput `json:"authorship"`
}

type authorshipOutput struct {
	Value    string          `json:"value"`
	Original *originalOutput `json:"basionymAuthorship,omitempty"`
}

type originalOutput struct {
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

func appendCanonical(c1 Canonical, c2 Canonical, sep string) Canonical {
	return Canonical{
		Value:       str.JoinStrings(c1.Value, c2.Value, sep),
		ValueRanked: str.JoinStrings(c1.ValueRanked, c2.ValueRanked, sep),
	}
}

func nilCanonical() Canonical {
	var c Canonical
	return c
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

func (sn *ScientificNameNode) Canonical() Canonical {
	ng := sn.NamesGroup
	if len(ng) == 1 {
		return ng[0].canonical()
	}
	var cs Canonical
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
	if len(sn.NamesGroup) > 1 {
		var ao *authorshipOutput
		return ao
	}
	an := sn.NamesGroup[0].lastAuthorship()
	return an.details()
}

// func (sp *speciesNode) pos() []Pos {
// 	pos := []Pos{sp.Genus.Pos}
// 	pos = append(pos, sp.Species.Word.Pos)
// 	pos = append(pos, sp.Species.Authorship.pos()...)
// 	return pos
// }

// func (sp *speciesNode) value() string {
// 	res := str.JoinStrings(sp.Genus.NormValue, sp.Species.Word.NormValue, " ")
// 	res = str.JoinStrings(res, sp.Species.Authorship.value(), " ")
// 	return res
// }

// func (sp *speciesNode) canonical() Canonical {
// 	spPart := str.JoinStrings(sp.Genus.NormValue, sp.Species.Word.NormValue, " ")
// 	return Canonical{Value: spPart, ValueRanked: spPart}
// }

// func (sp *speciesNode) lastAuthorship() *authorshipNode {
// 	return sp.Species.Authorship
// }

// func (sp *speciesNode) details() interface{} {
// 	se := specEpithetOutput{
// 		Value: sp.Species.Word.Value,
// 	}
// 	if sp.Species.Authorship != nil {
// 		se.Authorship = sp.Species.Authorship.details()
// 	}

// 	g := sp.Genus.Value
// 	so := &SpeciesOutput{
// 		Genus:       &genusOutput{Value: g},
// 		SpecEpithet: &se,
// 	}
// 	return so
// }

func (u *uninomialNode) pos() []Pos {
	pos := []Pos{u.Word.Pos}
	pos = append(pos, u.Authorship.pos()...)
	return pos
}

func (u *uninomialNode) value() string {
	return str.JoinStrings(u.Word.NormValue, u.Authorship.value(), " ")
}

func (u *uninomialNode) canonical() Canonical {
	return Canonical{Value: u.Word.NormValue, ValueRanked: u.Word.NormValue}
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

// func (u *uninomialComboNode) pos() []Pos {
// 	pos := []Pos{u.Uninomial1.Word.Pos}
// 	if u.Rank.Word.Pos.Start != 0 {
// 		pos = append(pos, u.Rank.Word.Pos)
// 	}
// 	pos = append(pos, u.Uninomial2.Word.Pos)
// 	pos = append(pos, u.Uninomial2.Authorship.pos()...)
// 	return pos
// }

// func (u *uninomialComboNode) value() string {
// 	vl := str.JoinStrings(u.Uninomial1.Word.NormValue, u.Rank.Word.Value, " ")
// 	tail := str.JoinStrings(u.Uninomial2.Word.NormValue,
// 		u.Uninomial2.Authorship.value(), " ")
// 	return str.JoinStrings(vl, tail, " ")
// }

// func (u *uninomialComboNode) canonical() Canonical {
// 	ranked := str.JoinStrings(u.Uninomial1.Word.NormValue, u.Rank.Word.Value, " ")
// 	ranked = str.JoinStrings(ranked, u.Uninomial2.Word.NormValue, " ")

// 	return Canonical{
// 		Value:       u.Uninomial2.Word.NormValue,
// 		ValueRanked: ranked,
// 	}
// }

// func (u *uninomialComboNode) lastAuthorship() *authorshipNode {
// 	return u.Uninomial2.Authorship
// }

// func (u *uninomialComboNode) details() interface{} {
// 	ud := uniDetails{
// 		Value:  u.Uninomial2.Word.NormValue,
// 		Rank:   u.Rank.Word.Value,
// 		Parent: u.Uninomial1.Word.NormValue,
// 	}
// 	if u.Uninomial2.Authorship != nil {
// 		ud.Authorship = u.Uninomial2.Authorship.details()
// 	}
// 	uo := UninomialOutput{Uninomial: &ud}
// 	return &uo
// }

func (au *authorshipNode) details() *authorshipOutput {
	if au == nil {
		var ao *authorshipOutput
		return ao
	}
	ao := authorshipOutput{Value: au.value()}
	auYrs := au.OriginalAuthors.Team1.Years
	yrs := make([]yearOutput, len(auYrs))
	for i, v := range auYrs {
		yrs[i] = yearOutput{Value: v.Word.Value, Approximate: v.Approximate}
	}
	auAuths := au.OriginalAuthors.Team1.Authors
	aus := make([]string, len(auAuths))
	for i, v := range auAuths {
		aus[i] = v.Value
	}
	ao.Original = &originalOutput{
		Authors: aus,
		Years:   yrs,
	}
	// TODO more stuff
	return &ao
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
	// TODO more stuff
	return v
}

func (ag *authorsGroupNode) value() string {
	v := ag.Team1.value()
	if ag.Team2 == nil {
		return v
	}
	// TODO more stuff
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
