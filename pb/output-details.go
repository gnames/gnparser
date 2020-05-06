package pb

import (
	"gitlab.com/gogna/gnparser/grammar"
	"gitlab.com/gogna/gnparser/output"
)

func details(po *Parsed, o *output.Output) {
	switch len(o.Details) {
	case 0:
		return
	case 1:
		simpleName(po, o)
	default:
		hybridName(po, o)
	}
}

func simpleName(po *Parsed, o *output.Output) {
	switch d := o.Details[0].(type) {
	case *grammar.UninomialOutput:
		res := uninomial(po, o, d)
		po.Details = &Parsed_Uninomial{res}
	case *grammar.SpeciesOutput:
		res := species(po, o, d)
		po.Details = &Parsed_Species{res}
	case *grammar.ComparisonOutput:
		res := comparison(po, o, d)
		po.Details = &Parsed_Comparison{res}
	case *grammar.ApproxOutput:
		res := approx(po, o, d)
		po.Details = &Parsed_Approximation{res}
	}
	if po.Hybrid {
		po.NameType = NameType_NAMED_HYBRID
	}
}

func hybridName(po *Parsed, o *output.Output) {
	hf := make([]*HybridFormula, len(o.Details))
	for i, v := range o.Details {
		switch d := v.(type) {
		case *grammar.UninomialOutput:
			res := uninomial(po, o, d)
			hf[i] = &HybridFormula{Element: &HybridFormula_Uninomial{res}}
		case *grammar.SpeciesOutput:
			res := species(po, o, d)
			hf[i] = &HybridFormula{Element: &HybridFormula_Species{res}}
		case *grammar.ComparisonOutput:
			res := comparison(po, o, d)
			hf[i] = &HybridFormula{Element: &HybridFormula_Comparison{res}}
		case *grammar.ApproxOutput:
			res := approx(po, o, d)
			hf[i] = &HybridFormula{Element: &HybridFormula_Approximation{res}}
		}
	}
	po.Authorship = nil
	po.NameType = NameType_HYBRID_FORMULA
	po.DetailsHybridFormula = hf
}

func uninomial(po *Parsed, o *output.Output,
	uo *grammar.UninomialOutput) *Uninomial {
	u := &Uninomial{
		Value:  uo.Uninomial.Value,
		Rank:   uo.Uninomial.Rank,
		Parent: uo.Uninomial.Parent,
	}

	if uo.Uninomial.Authorship != nil {
		au := authorship(uo.Uninomial.Authorship)
		u.Authorship = au
		po.Authorship = au
	}
	po.NameType = NameType_UNINOMIAL
	return u
}

func species(po *Parsed, o *output.Output,
	so *grammar.SpeciesOutput) *Species {
	var au *Authorship
	s := &Species{
		Genus:   so.Genus.Value,
		Species: so.SpecEpithet.Value,
	}
	if so.SubGenus != nil {
		s.SubGenus = so.SubGenus.Value
	}
	if so.SpecEpithet.Authorship != nil {
		au = authorship(so.SpecEpithet.Authorship)
		s.SpeciesAuthorship = au
	}

	if len(so.InfraSpecies) > 0 {
		au = nil
		inf := make([]*InfraSpecies, len(so.InfraSpecies))
		for i, v := range so.InfraSpecies {
			inf[i] = infraspecies(v)
		}
		if inf[len(inf)-1].Authorship != nil {
			au = inf[len(inf)-1].Authorship
		}
	} else {
		au = s.SpeciesAuthorship
	}
	po.Authorship = au
	po.NameType = NameType_SPECIES
	return s
}

func comparison(po *Parsed, o *output.Output,
	co *grammar.ComparisonOutput) *Comparison {
	c := &Comparison{
		Genus: co.Genus.Value,
	}

	if co.SpecEpithet != nil {
		c.Species = co.SpecEpithet.Value
		if co.SpecEpithet.Authorship != nil {
			c.SpeciesAuthorship = authorship(co.SpecEpithet.Authorship)
		}
	}

	if co.Comparison != "" {
		c.Comparison = co.Comparison
	}
	po.NameType = NameType_COMPARISON
	return c
}

func approx(po *Parsed, o *output.Output,
	ao *grammar.ApproxOutput) *Approximation {
	po.NameType = NameType_APPROX_SURROGATE
	a := &Approximation{Genus: ao.Genus.Value}
	if ao.SpecEpithet != nil {
		a.Species = ao.SpecEpithet.Value
		if ao.SpecEpithet.Authorship != nil {
			a.SpeciesAuthorship = authorship(ao.SpecEpithet.Authorship)
		}
	}
	if ao.Approx != "" {
		a.Approximation = ao.Approx
	}
	if ao.Ignored != nil {
		a.Ignored = ao.Ignored.Value
	}
	return a
}

func infraspecies(inf *grammar.InfraSpEpithetOutput) *InfraSpecies {
	res := &InfraSpecies{Value: inf.Value}
	if inf.Rank != "" {
		res.Rank = inf.Rank
	}
	if inf.Authorship != nil {
		res.Authorship = authorship(inf.Authorship)
	}
	return res
}

func authorship(a *grammar.AuthorshipOutput) *Authorship {
	var allAuth []string
	var authList []string
	var au *Authorship
	var orig, comb *AuthGroup

	if a == nil {
		return au
	}
	if a.Original != nil {
		orig, authList = authGroup(a.Original)
		allAuth = append(allAuth, authList...)
	}
	if a.Combination != nil {
		comb, authList = authGroup(a.Combination)
		allAuth = append(allAuth, authList...)
	}

	au = &Authorship{
		Value:       a.Value,
		AllAuthors:  allAuth,
		Original:    orig,
		Combination: comb,
	}
	return au
}

func authGroup(ago *grammar.AuthGroupOutput) (*AuthGroup, []string) {
	var exAu, emendAu *Authors
	var authList []string
	allAuth := ago.Authors

	if ago.ExAuthors != nil {
		exAu, authList = authors(ago.ExAuthors)
		allAuth = append(allAuth, authList...)
	}
	if ago.EmendAuthors != nil {
		emendAu, authList = authors(ago.EmendAuthors)
		allAuth = append(allAuth, authList...)
	}

	ag := &AuthGroup{
		Authors:      ago.Authors,
		ExAuthors:    exAu,
		EmendAuthors: emendAu,
	}
	if ago.Year != nil {
		ag.Year = ago.Year.Value
		ag.ApproximateYear = ago.Year.Approximate
	}
	return ag, allAuth
}

func authors(aso *grammar.AuthorsOutput) (*Authors, []string) {
	as := &Authors{Authors: aso.Authors}
	if aso.Year != nil {
		as.Year = aso.Year.Value
		as.ApproximateYear = aso.Year.Approximate
	}
	return as, aso.Authors
}
