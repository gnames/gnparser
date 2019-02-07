package grammar

import (
	"fmt"

	"gitlab.com/gogna/gnparser/str"
)

type ApproxOutput struct {
	Genus       *genusOutput       `json:"genus"`
	SpecEpithet *specEpithetOutput `json:"specificEpithet,omitempty"`
	Approx      string             `json:"annotationIdentification"`
	Ignored     *ignoredOutput     `json:"ignored,omitempty"`
}

type ignoredOutput struct {
	Value string `json:"value"`
}

type ComparisonOutput struct {
	Genus       *genusOutput       `json:"genus"`
	SpecEpithet *specEpithetOutput `json:"specificEpithet"`
	Comparison  string             `json:"annotationIdentification"`
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

type UninomialOutput struct {
	Uninomial *uniDetails `json:"uninomial"`
}

type uniDetails struct {
	Value      string            `json:"value"`
	Rank       string            `json:"rank,omitempty"`
	Parent     string            `json:"parent,omitempty"`
	Authorship *authorshipOutput `json:"authorship,omitempty"`
}

type authorshipOutput struct {
	Value       string           `json:"value"`
	Original    *authGroupOutput `json:"basionymAuthorship,omitempty"`
	Combination *authGroupOutput `json:"combinationAuthorship,omitempty"`
}

type authGroupOutput struct {
	Authors      []string            `json:"authors"`
	Year         *yearOutput         `json:"year,omitempty"`
	ExAuthors    *exAuthorsOutput    `json:"exAuthors,omitempty"`
	EmendAuthors *emendAuthorsOutput `json:"emendAuthors,omitempty"`
}

type exAuthorsOutput struct {
	Authors []string    `json:"authors"`
	Year    *yearOutput `json:"year,omitempty"`
}

type emendAuthorsOutput struct {
	Authors []string    `json:"authors"`
	Year    *yearOutput `json:"year,omitempty"`
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
	return sn.Name.pos()
}

func (sn *ScientificNameNode) Value() string {
	if sn.Name == nil {
		return ""
	}
	return sn.Name.value()
}

func (sn *ScientificNameNode) Canonical() *Canonical {
	if sn.Name == nil {
		var c *Canonical
		return c
	}
	return sn.Name.canonical()
}

func (sn *ScientificNameNode) Details() []interface{} {
	if sn.Name == nil {
		return []interface{}{}
	}
	return sn.Name.details()
}

func (sn *ScientificNameNode) LastAuthorship() *authorshipOutput {
	var ao *authorshipOutput
	if sn.Name == nil {
		return ao
	}
	an := sn.Name.lastAuthorship()
	if an == nil {
		return ao
	}
	return an.details()
}

func (nf *hybridFormulaNode) pos() []Pos {
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

func (nf *hybridFormulaNode) canonical() *Canonical {
	c := nf.FirstSpecies.canonical()
	for _, v := range nf.HybridElements {
		hc := &Canonical{
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

func (nf *hybridFormulaNode) details() []interface{} {
	ds := nf.FirstSpecies.details()
	for _, v := range nf.HybridElements {
		if v.Species != nil {
			ds = append(ds, v.Species.details()[0])
		}
	}
	return ds
}

func (nh *namedGenusHybridNode) pos() []Pos {
	pos := []Pos{nh.Hybrid.Pos}
	pos = append(pos, nh.Name.pos()...)
	return pos
}

func (nh *namedGenusHybridNode) value() string {
	v := nh.Name.value()
	v = "× " + v
	return v
}

func (nh *namedGenusHybridNode) canonical() *Canonical {
	c := &Canonical{
		Value:       "",
		ValueRanked: "×",
	}

	c1 := nh.Name.canonical()
	c = appendCanonical(c, c1, " ")
	return c
}

func (nh *namedGenusHybridNode) details() []interface{} {
	d := nh.Name.details()
	return d
}

func (nh *namedGenusHybridNode) lastAuthorship() *authorshipNode {
	au := nh.Name.lastAuthorship()
	return au
}

func (nh *namedSpeciesHybridNode) pos() []Pos {
	pos := []Pos{nh.Genus.Pos}
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

func (nh *namedSpeciesHybridNode) canonical() *Canonical {
	g := nh.Genus.NormValue
	c := &Canonical{Value: g, ValueRanked: g}
	hCan := &Canonical{Value: "", ValueRanked: "×"}
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

func (nh *namedSpeciesHybridNode) details() []interface{} {
	g := &genusOutput{Value: nh.Genus.NormValue}
	sp := nh.SpEpithet.details()
	so := &SpeciesOutput{
		Genus:       g,
		SpecEpithet: sp,
	}
	if len(nh.InfraSpecies) == 0 {
		return []interface{}{so}
	}
	infs := make([]*infraSpEpithetOutput, len(nh.InfraSpecies))
	for i, v := range nh.InfraSpecies {
		infs[i] = v.details()
	}
	so.InfraSpecies = infs

	return []interface{}{so}
}

func (apr *approxNode) pos() []Pos {
	var pos []Pos
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

func (apr *approxNode) canonical() *Canonical {
	var c *Canonical
	if apr == nil {
		return c
	}
	c = &Canonical{Value: apr.Genus.NormValue, ValueRanked: apr.Genus.NormValue}
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

func (apr *approxNode) details() []interface{} {
	if apr == nil {
		return []interface{}{}
	}
	g := apr.Genus.NormValue
	ao := &ApproxOutput{
		Genus:   &genusOutput{Value: g},
		Approx:  apr.Approx.NormValue,
		Ignored: &ignoredOutput{Value: apr.Ignored},
	}
	if apr.SpEpithet == nil {
		return []interface{}{ao}
	}
	se := &specEpithetOutput{
		Value: apr.SpEpithet.Word.NormValue,
	}
	if apr.SpEpithet.Authorship != nil {
		se.Authorship = apr.SpEpithet.Authorship.details()
	}
	ao.SpecEpithet = se
	return []interface{}{ao}
}

func (comp *comparisonNode) pos() []Pos {
	var pos []Pos
	if comp == nil {
		return pos
	}
	pos = []Pos{comp.Genus.Pos}
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

func (comp *comparisonNode) canonical() *Canonical {
	if comp == nil {
		return &Canonical{}
	}
	gen := comp.Genus.NormValue
	c := &Canonical{Value: gen, ValueRanked: gen}
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

func (comp *comparisonNode) details() []interface{} {
	if comp == nil {
		return []interface{}{}
	}
	var se *specEpithetOutput
	if comp.SpEpithet != nil {
		se = comp.SpEpithet.details()
	}

	co := &ComparisonOutput{
		Genus:       &genusOutput{Value: comp.Genus.NormValue},
		Comparison:  comp.Comparison.NormValue,
		SpecEpithet: se,
	}
	return []interface{}{co}
}

func (sp *speciesNode) pos() []Pos {
	var pos []Pos
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

func (sp *speciesNode) canonical() *Canonical {
	spPart := str.JoinStrings(sp.Genus.NormValue, sp.SpEpithet.Word.NormValue, " ")
	c := &Canonical{Value: spPart, ValueRanked: spPart}
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

func (sp *speciesNode) details() []interface{} {
	se := specEpithetOutput{
		Value: sp.SpEpithet.Word.NormValue,
	}
	if sp.SpEpithet.Authorship != nil {
		se.Authorship = sp.SpEpithet.Authorship.details()
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
		return []interface{}{&so}
	}
	infs := make([]*infraSpEpithetOutput, len(sp.InfraSpecies))
	for i, v := range sp.InfraSpecies {
		infs[i] = v.details()
	}
	so.InfraSpecies = infs

	return []interface{}{&so}
}

func (sep *spEpithetNode) pos() []Pos {
	pos := []Pos{sep.Word.Pos}
	pos = append(pos, sep.Authorship.pos()...)
	return pos
}

func (sep *spEpithetNode) value() string {
	val := sep.Word.NormValue
	val = str.JoinStrings(val, sep.Authorship.value(), " ")
	return val
}

func (sep *spEpithetNode) canonical() *Canonical {
	c := &Canonical{Value: sep.Word.NormValue, ValueRanked: sep.Word.NormValue}
	return c
}

func (sep *spEpithetNode) details() *specEpithetOutput {
	val := sep.Word.NormValue
	au := sep.Authorship.details()
	seo := specEpithetOutput{
		Value:      val,
		Authorship: au,
	}
	return &seo
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

func (u *uninomialNode) details() []interface{} {
	ud := uniDetails{Value: u.Word.NormValue}
	if u.Authorship != nil {
		ud.Authorship = u.Authorship.details()
	}
	uo := UninomialOutput{Uninomial: &ud}
	return []interface{}{&uo}
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
	vl := str.JoinStrings(u.Uninomial1.Word.NormValue, u.Rank.Word.NormValue, " ")
	tail := str.JoinStrings(u.Uninomial2.Word.NormValue,
		u.Uninomial2.Authorship.value(), " ")
	return str.JoinStrings(vl, tail, " ")
}

func (u *uninomialComboNode) canonical() *Canonical {
	ranked := str.JoinStrings(u.Uninomial1.Word.NormValue, u.Rank.Word.NormValue, " ")
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

func (u *uninomialComboNode) details() []interface{} {
	ud := uniDetails{
		Value:  u.Uninomial2.Word.NormValue,
		Rank:   u.Rank.Word.NormValue,
		Parent: u.Uninomial1.Word.NormValue,
	}
	if u.Uninomial2.Authorship != nil {
		ud.Authorship = u.Uninomial2.Authorship.details()
	}
	uo := UninomialOutput{Uninomial: &ud}
	return []interface{}{&uo}
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
	aus, yr := ag.Team1.details()
	ago = authGroupOutput{
		Authors: aus,
		Year:    yr,
	}
	if ag.Team2 == nil {
		return &ago
	}
	aus, yr = ag.Team2.details()
	switch ag.Team2Type.Pos.Type {
	case AuthorWordExType:
		eao := exAuthorsOutput{
			Authors: aus,
			Year:    yr,
		}
		ago.ExAuthors = &eao
	case AuthorWordEmendType:
		eao := emendAuthorsOutput{
			Authors: aus,
			Year:    yr,
		}
		ago.EmendAuthors = &eao
	}
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
	v = fmt.Sprintf("%s %s %s", v, ag.Team2Type.NormValue, ag.Team2.value())
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

func (at *authorsTeamNode) details() ([]string, *yearOutput) {
	var yr *yearOutput
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
	yr = &yearOutput{
		Value:       at.Year.Word.NormValue,
		Approximate: at.Year.Approximate,
	}
	return aus, yr
}

func (aut *authorsTeamNode) pos() []Pos {
	var res []Pos
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

func (aun *authorNode) pos() []Pos {
	p := make([]Pos, len(aun.Words))
	for i, v := range aun.Words {
		p[i] = v.Pos
	}
	return p
}
